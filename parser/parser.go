// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"encoding/binary"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-clang/v3.9/clang"
	"github.com/pkg/errors"
	"github.com/pkgutil/osutil"
	"github.com/pkgutil/stringsutil"
	"github.com/zchee/clang-server/compilationdatabase"
	"github.com/zchee/clang-server/indexdb"
	"github.com/zchee/clang-server/internal/log"
	"github.com/zchee/clang-server/internal/pathutil"
	"github.com/zchee/clang-server/parser/builtinheader"
	"github.com/zchee/clang-server/symbol"
)

// defaultClangOption defalut global clang options.
// clang.TranslationUnit_DetailedPreprocessingRecord = 0x01
// clang.TranslationUnit_Incomplete = 0x02
// clang.TranslationUnit_PrecompiledPreamble = 0x04
// clang.TranslationUnit_CacheCompletionResults = 0x08
// clang.TranslationUnit_ForSerialization = 0x10
// clang.TranslationUnit_CXXChainedPCH = 0x20
// clang.TranslationUnit_SkipFunctionBodies = 0x40
// clang.TranslationUnit_IncludeBriefCommentsInCodeCompletion = 0x80
// clang.TranslationUnit_CreatePreambleOnFirstParse = 0x100
// clang.TranslationUnit_KeepGoing = 0x200
// const defaultClangOption uint32 = 0x445 // Use all flags for now
var defaultClangOption = clang.DefaultEditingTranslationUnitOptions() | uint32(clang.TranslationUnit_KeepGoing)

// Parser represents a C/C++ AST parser.
type Parser struct {
	root        string
	clangOption uint32

	idx clang.Index
	cd  *compilationdatabase.CompilationDatabase
	db  *indexdb.IndexDB

	dispatcher *dispatcher

	debugUncatched bool                     // for debug
	uncachedKind   map[clang.CursorKind]int // for debug
}

// Config represents a parser config.
type Config struct {
	JSONName    string
	PathRange   []string
	ClangOption uint32

	Debug bool
}

// NewParser return the new Parser.
func NewParser(path string, config Config) *Parser {
	root, err := pathutil.FindProjectRoot(path)
	if err != nil {
		log.Fatal(err)
	}

	cd := compilationdatabase.NewCompilationDatabase(root)
	if err := cd.Parse(config.JSONName, config.PathRange); err != nil {
		log.Fatal(err)
	}

	db, err := indexdb.NewIndexDB(root)
	if err != nil {
		log.Fatal(err)
	}

	clangOption := config.ClangOption
	if clangOption == 0 {
		clangOption = defaultClangOption
	}

	p := &Parser{
		root:        root,
		clangOption: clangOption,
		idx:         clang.NewIndex(0, 1), // disable excludeDeclarationsFromPCH, enable displayDiagnostics
		cd:          cd,
		db:          db,
	}

	if config.Debug {
		p.debugUncatched = true
		p.uncachedKind = make(map[clang.CursorKind]int)
	}

	if err := CreateBulitinHeaders(); err != nil {
		log.Fatal(err)
	}

	return p
}

// CreateBulitinHeaders creates(dumps) a clang builtin header to cache directory.
func CreateBulitinHeaders() error {
	builtinHdrDir := filepath.Join(pathutil.CacheDir(), "clang", "include")
	if !osutil.IsExist(builtinHdrDir) {
		if err := os.MkdirAll(builtinHdrDir, 0700); err != nil {
			return errors.WithStack(err)
		}
	}

	for _, fname := range builtinheader.AssetNames() {
		data, err := builtinheader.AssetInfo(fname)
		if err != nil {
			return errors.WithStack(err)
		}

		if strings.Contains(data.Name(), string(filepath.Separator)) {
			dir, _ := filepath.Split(data.Name())
			if err := os.MkdirAll(filepath.Join(builtinHdrDir, dir), 0700); err != nil {
				return errors.WithStack(err)
			}
		}

		buf, err := builtinheader.Asset(data.Name())
		if err != nil {
			return errors.WithStack(err)
		}

		if err := ioutil.WriteFile(filepath.Join(builtinHdrDir, data.Name()), buf, 0600); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// Parse parses the project directories.
func (p *Parser) Parse() {
	defer func() {
		p.db.Close()
		symbol.Serve()
	}()

	ccs := p.cd.CompileCommands()
	if len(ccs) == 0 {
		log.Fatal("not walk")
	}

	compilerConfig := p.cd.CompilerConfig
	flags := append(compilerConfig.SystemCIncludeDir, compilerConfig.SystemFrameworkDir...)

	// TODO(zchee): needs include stdint.h?
	if i := stringsutil.IndexContainsSlice(ccs[0].Arguments, "-std="); i > 0 {
		std := ccs[0].Arguments[i][5:]
		switch {
		case strings.HasPrefix(std, "c"), strings.HasPrefix(std, "gnu"):
			if std[len(std)-2] == '8' || std[len(std)-2] == '9' {
				flags = append(flags, "-include", "/usr/include/stdint.h")
			}
		}
	} else {
		flags = append(flags, "-include", "/usr/include/stdint.h")
	}
	if !(filepath.Ext(ccs[0].File) == ".c") {
		flags = append(flags, compilerConfig.SystemCXXIncludeDir...)
	}

	builtinHdrDir := filepath.Join(pathutil.CacheDir(), "clang", "include")
	flags = append(flags, "-I"+builtinHdrDir)

	p.dispatcher = newDispatcher(p.ParseFile)
	p.dispatcher.Start()
	for i := 0; i < len(ccs); i++ {
		args := ccs[i].Arguments
		args = append(flags, args...)
		p.dispatcher.Add(parseArg{ccs[i].File, args})
	}
	p.dispatcher.Wait()
}

type parseArg struct {
	filename string
	flag     []string
}

// ParseFile parses the C/C++ file.
func (p *Parser) ParseFile(arg parseArg) error {
	var tu clang.TranslationUnit

	if p.db.Has(arg.filename) {
		buf, err := p.db.Get(arg.filename)
		if err != nil {
			return err
		}

		tu, err = deserializeTranslationUnit(p.idx, buf)
		if err != nil {
			return err
		}
		defer tu.Dispose()

		log.Debugf("tu.Spelling(): %T => %+v\n", tu.Spelling(), tu.Spelling()) // for debug

		return nil
	}

	if cErr := p.idx.ParseTranslationUnit2(arg.filename, arg.flag, nil, p.clangOption, &tu); clang.ErrorCode(cErr) != clang.Error_Success {
		return errors.New(clang.ErrorCode(cErr).Spelling())
	}
	defer tu.Dispose()

	tuch := make(chan []byte)
	go func() {
		tuch <- serializeTranslationUnit(arg.filename, tu)
	}()

	// printDiagnostics(tu.Diagnostics())

	rootCursor := tu.TranslationUnitCursor()
	file := symbol.NewFile(arg.filename)
	visitNode := func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		if cursor.IsNull() {
			log.Debug("cursor: <none>")
			return clang.ChildVisit_Continue
		}

		cursorLoc := symbol.FromCursor(cursor)
		if cursorLoc.FileName() == "" || cursorLoc.FileName() == "." {
			// TODO(zchee): Ignore system header(?)
			return clang.ChildVisit_Continue
		}

		kind := cursor.Kind()
		switch kind {
		case clang.Cursor_FunctionDecl, clang.Cursor_StructDecl, clang.Cursor_FieldDecl, clang.Cursor_TypedefDecl, clang.Cursor_EnumDecl, clang.Cursor_EnumConstantDecl:
			defCursor := cursor.Definition()
			if defCursor.IsNull() {
				file.AddDecl(cursorLoc)
			} else {
				defLoc := symbol.FromCursor(defCursor)
				file.AddDefinition(cursorLoc, defLoc)
			}
		case clang.Cursor_MacroDefinition:
			file.AddDefinition(cursorLoc, cursorLoc)
		case clang.Cursor_VarDecl:
			file.AddDecl(cursorLoc)
		case clang.Cursor_ParmDecl:
			if cursor.Spelling() != "" {
				file.AddDecl(cursorLoc)
			}
		case clang.Cursor_CallExpr:
			refCursor := cursor.Referenced()
			refLoc := symbol.FromCursor(refCursor)
			file.AddCaller(cursorLoc, refLoc, true)
		case clang.Cursor_DeclRefExpr, clang.Cursor_TypeRef, clang.Cursor_MemberRefExpr, clang.Cursor_MacroExpansion:
			refCursor := cursor.Referenced()
			refLoc := symbol.FromCursor(refCursor)
			file.AddCaller(cursorLoc, refLoc, false)
		case clang.Cursor_InclusionDirective:
			incFile := cursor.IncludedFile()
			file.AddHeader(cursor.Spelling(), incFile)
		default:
			if p.debugUncatched {
				p.uncachedKind[kind]++
			}
		}

		return clang.ChildVisit_Recurse
	}

	rootCursor.Visit(visitNode)
	file.AddTranslationUnit(<-tuch)
	buf := file.Serialize()

	out := symbol.GetRootAsFile(buf.FinishedBytes(), 0)
	printFile(out) // for debug

	log.Debugf("Goroutine:%d", runtime.NumGoroutine())
	log.Debugf("\n================== DONE: filename: %+v ==================\n\n\n", arg.filename)

	return p.db.Put(arg.filename, buf.FinishedBytes())
}

// serializeTranslationUnit selialize the TranslationUnit to Clang serialized representation.
// TODO(zchee): Avoid ioutil.TempFile if possible.
func serializeTranslationUnit(filename string, tu clang.TranslationUnit) []byte {
	tmpFile, err := ioutil.TempFile(os.TempDir(), filepath.Base(filename))
	if err != nil {
		log.Fatal(err)
	}

	saveOptions := uint32(clang.TranslationUnit_KeepGoing)
	if cErr := tu.SaveTranslationUnit(tmpFile.Name(), saveOptions); clang.SaveError(cErr) != clang.SaveError_None {
		log.Fatal(clang.SaveError(cErr))
	}

	buf, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	os.Remove(tmpFile.Name())

	return buf
}

func deserializeTranslationUnit(idx clang.Index, buf []byte) (clang.TranslationUnit, error) {
	var tu clang.TranslationUnit

	tmpfile, err := ioutil.TempFile(os.TempDir(), "clang-server")
	if err != nil {
		return tu, err
	}
	binary.Write(tmpfile, binary.LittleEndian, buf)

	if err := idx.TranslationUnit2(tmpfile.Name(), &tu); clang.ErrorCode(err) != clang.Error_Success {
		return tu, errors.New(err.Spelling())
	}
	// finished create a translation unit from an AST file, remove tmpfile
	os.Remove(tmpfile.Name())

	return tu, nil
}

// ClangVersion return the current clang version.
func ClangVersion() string {
	return clang.GetClangVersion()
}
