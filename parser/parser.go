// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/go-clang/v3.9/clang"
	"github.com/pkg/errors"
	"github.com/zchee/clang-server/compilationdatabase"
	"github.com/zchee/clang-server/indexdb"
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
const defaultClangOption uint32 = 0x445 // Use all flags for now

func init() {
	// for debug
	log.SetFlags(log.Lshortfile)
}

// Parser represents a C/C++ AST parser.
type Parser struct {
	root        string
	clangOption uint32

	idx clang.Index
	cd  *compilationdatabase.CompilationDatabase
	db  *indexdb.IndexDB

	mu sync.Mutex

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

	return p
}

// CreateBulitinHeaders creates(dumps) a clang builtin header to cache directory.
func CreateBulitinHeaders() error {
	builtinHdrDir := filepath.Join(pathutil.CacheDir(), "clang", "include")
	if pathutil.IsNotExist(builtinHdrDir) {
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

// Walk walk project directories.
func (p *Parser) Walk() {
	ccs := p.cd.CompileCommands()
	if len(ccs) == 0 {
		log.Fatal("not walk")
	}

	flags := []string{}
	compilerConfig := p.cd.CompilerConfig
	flags = compilerConfig.SystemIncludeDir
	flags = append(flags, compilerConfig.SystemFrameworkDir...)
	builtinHdrDir := filepath.Join(pathutil.CacheDir(), "clang", "include")
	flags = append(flags, "-I"+builtinHdrDir,
		"-Wno-nullability-completeness", // TODO(zchee): stdlib.h,stdio.h: pointer is missing a nullability type specifier (_Nonnull, _Nullable, or _Null_unspecified)
		"-Wno-expansion-to-defined")     // TODO(zchee): macro expansion producing 'defined' has undefined behavior

	ch := make(chan struct{}, runtime.NumCPU()*3)

	var wg sync.WaitGroup
	for i := 0; i < len(ccs); i++ {
		wg.Add(1)
		i := i
		flags := flags
		go func(i int, flags []string) {
			defer wg.Done()
			flags = append(flags, ccs[i].Arguments...)
			flags = append(flags, "-std=c11")
			ch <- struct{}{}
			if err := p.ParseFile(ccs[i].File, flags); err != nil {
				log.Fatal(err)
			}
			<-ch
		}(i, flags)
	}
	wg.Wait()
}

// Parse parses the projects C/C++ files.
func (p *Parser) Parse() error {
	if err := CreateBulitinHeaders(); err != nil {
		return err
	}
	p.idx.SetGlobalOptions(clang.GlobalOpt_ThreadBackgroundPriorityForAll)
	p.Walk()
	log.Printf("done")

	return nil
}

// ParseFile parses the C/C++ file.
func (p *Parser) ParseFile(filename string, flags []string) error {
	var tu clang.TranslationUnit

	if p.db.Has(filename) {
		// if p.has(filename) {
		// log.Info("has")
		// tmpFile, err := ioutil.TempFile(os.TempDir(), filepath.Base(filename))
		// if err != nil {
		// 	return err
		// }
		// defer os.Remove(tmpFile.Name())
		//
		// buf, err := p.db.Get(filename)
		// if err != nil {
		// 	return err
		// }
		// file := symbol.GetRootAsFile(buf, 0)
		// tmpFile.Write(file.GetTranslationUnit())
		//
		// if cErr := p.idx.TranslationUnit2(tmpFile.Name(), &tu); clang.ErrorCode(cErr) != clang.Error_Success {
		// 	return errors.New(clang.ErrorCode(cErr).Spelling())
		// }
		return nil
	}

	if cErr := p.idx.ParseTranslationUnit2(filename, flags, nil, p.clangOption, &tu); clang.ErrorCode(cErr) != clang.Error_Success {
		return errors.New(clang.ErrorCode(cErr).Spelling())
	}
	defer tu.Dispose()

	// p.printDiagnostics(tu.Diagnostics())
	symDB := NewSymbolDB(filename)
	visitNode := func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		if cursor.IsNull() {
			log.Printf("cursor: <none>")
			return clang.ChildVisit_Continue
		}

		cursorLoc := symbol.FromCursor(&cursor)
		if cursorLoc.File == "" || cursorLoc.File == "." {
			// TODO(zchee): Ignore system header(?)
			return clang.ChildVisit_Continue
		}

		kind := cursor.Kind()
		switch kind {
		case clang.Cursor_FunctionDecl, clang.Cursor_StructDecl, clang.Cursor_FieldDecl, clang.Cursor_TypedefDecl, clang.Cursor_EnumDecl, clang.Cursor_EnumConstantDecl:
			defCursor := cursor.Definition()
			switch {
			case defCursor.IsNull():
				symbolDB.AddDecl(cursorLoc)
			default:
				defLoc := symbol.FromCursor(&defCursor)
				symbolDB.AddDefinition(cursorLoc, defLoc)
			}
		case clang.Cursor_MacroDefinition:
			symbolDB.AddDefinition(cursorLoc, cursorLoc)
		case clang.Cursor_VarDecl:
			symbolDB.AddDecl(cursorLoc)
		case clang.Cursor_ParmDecl:
			if cursor.Spelling() != "" {
				symbolDB.AddDecl(cursorLoc)
			}
		case clang.Cursor_CallExpr:
			refCursor := cursor.Referenced()
			refLoc := symbol.FromCursor(&refCursor)
			symbolDB.AddCaller(cursorLoc, refLoc, true)
		case clang.Cursor_DeclRefExpr, clang.Cursor_TypeRef, clang.Cursor_MemberRefExpr, clang.Cursor_MacroExpansion:
			refCursor := cursor.Referenced()
			refLoc := symbol.FromCursor(&refCursor)
			symbolDB.AddCaller(cursorLoc, refLoc, false)
		case clang.Cursor_InclusionDirective:
			incFile := cursor.IncludedFile()
			symbolDB.AddHeader(cursor.Spelling(), incFile)
		default:
			if p.debugUncatched {
				p.uncachedKind[kind]++
			}
		}

		return clang.ChildVisit_Recurse
	}

	tu.TranslationUnitCursor().Visit(visitNode)

	log.Printf("done: filename: %+v\n", filename)

	return p.db.Put(filename, []byte{})
}

// serializeTranslationUnit selialize the TranslationUnit to Clang serialized representation.
// TODO(zchee): Avoid ioutil.TempFile, get directly if possible.
func (p *Parser) serializeTranslationUnit(filename string, tu clang.TranslationUnit) []byte {
	tmpFile, err := ioutil.TempFile(os.TempDir(), filepath.Base(filename))
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if cErr := tu.SaveTranslationUnit(tmpFile.Name(), tu.DefaultSaveOptions()); clang.SaveError(cErr) != clang.SaveError_None {
		log.Fatal(clang.SaveError(cErr))
	}

	buf, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	return buf
}

// printDiagnostics prints a diagnostics information.
func (p *Parser) printDiagnostics(diags []clang.Diagnostic) {
	for _, d := range diags {
		file, line, col, offset := d.Location().FileLocation()
		fmt.Println("Location:", file.Name(), line, col, offset)
		fmt.Println("PROBLEM:", d.Spelling())
	}
}

// ClangVersion return the current clang version.
func ClangVersion() string {
	return clang.GetClangVersion()
}
