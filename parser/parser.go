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
	"sync"
	"time"

	"github.com/go-clang/v3.9/clang"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/pkg/errors"
	"github.com/zchee/clang-server/compilationdatabase"
	"github.com/zchee/clang-server/indexdb"
	"github.com/zchee/clang-server/internal/pathutil"
	"github.com/zchee/clang-server/parser/clang/builtinheader"
	"github.com/zchee/clang-server/symbol"
)

// clangOption global clang options.
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
const clangOption uint32 = 0x445 // use all flags

// Parser represents a C/C++ AST parser.
type Parser struct {
	root string

	idx clang.Index
	cd  *compilationdatabase.CompilationDatabase
	db  *indexdb.IndexDB

	mu sync.Mutex

	// for debug
	debugUncatched bool
	uncachedKind   map[clang.CursorKind]int
}

// Config represents a parser config.
type Config struct {
	JSONName    string
	PathRange   []string
	ClangOption uint32

	Debug bool
}

func init() {
	log.SetFlags(log.Lshortfile)
}

// CreateBulitinHeaders creates(dumps) a clang builtin header to cache directory.
func CreateBulitinHeaders() error {
	files, err := builtinheader.AssetDir("clang/include")
	if err != nil {
		return errors.WithStack(err)
	}
	cacheDir := pathutil.CacheDir()
	includeDir := filepath.Join(cacheDir, "clang", "include")
	if pathutil.IsNotExist(includeDir) {
		if err := os.MkdirAll(includeDir, 0700); err != nil {
			return err
		}
	}

	for _, f := range files {
		data, err := builtinheader.Asset(filepath.Join("clang/include", f))
		if err != nil {
			continue
		}
		if err := ioutil.WriteFile(filepath.Join(includeDir, f), data, 0600); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// NewParser return the new Parser.
func NewParser(path string, config *Config) *Parser {
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

	p := &Parser{
		root: root,
		idx:  clang.NewIndex(0, 1), // disable excludeDeclarationsFromPCH, enable displayDiagnostics
		cd:   cd,
		db:   db,
	}

	if config.Debug {
		p.debugUncatched = true
		p.uncachedKind = make(map[clang.CursorKind]int)
	}

	return p
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
	includeDir := filepath.Join(pathutil.CacheDir(), "clang", "include")
	flags = append(flags, "-I"+includeDir,
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

// ParseFile parses the C/C++ file.
func (p *Parser) ParseFile(filename string, flags []string) error {
	var tu clang.TranslationUnit
	var tubuf chan []byte

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
	} else {
		if cErr := p.idx.ParseTranslationUnit2(filename, flags, nil, clangOption, &tu); clang.ErrorCode(cErr) != clang.Error_Success {
			return errors.New(clang.ErrorCode(cErr).Spelling())
		}
		go func() {
			tubuf <- p.SerializeTranslationUnit(filename, tu)
		}()
	}
	defer tu.Dispose()

	// p.PrintDiagnostics(tu.Diagnostics())

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
				symDB.AddDecl(cursorLoc)
			default:
				defLoc := symbol.FromCursor(&defCursor)
				symDB.AddDefinition(cursorLoc, defLoc)
			}
		case clang.Cursor_MacroDefinition:
			symDB.AddDefinition(cursorLoc, cursorLoc)
		case clang.Cursor_VarDecl:
			symDB.AddDecl(cursorLoc)
		case clang.Cursor_ParmDecl:
			if cursor.Spelling() != "" {
				symDB.AddDecl(cursorLoc)
			}
		case clang.Cursor_CallExpr:
			refCursor := cursor.Referenced()
			refLoc := symbol.FromCursor(&refCursor)
			symDB.AddCaller(cursorLoc, refLoc, true)
		case clang.Cursor_DeclRefExpr, clang.Cursor_TypeRef, clang.Cursor_MemberRefExpr, clang.Cursor_MacroExpansion:
			refCursor := cursor.Referenced()
			refLoc := symbol.FromCursor(&refCursor)
			symDB.AddCaller(cursorLoc, refLoc, false)
		case clang.Cursor_InclusionDirective:
			incFile := cursor.IncludedFile()
			symDB.AddHeader(cursor.Spelling(), incFile)
		default:
			if p.debugUncatched {
				p.uncachedKind[kind]++
			}
		}

		return clang.ChildVisit_Recurse
	}

	tu.TranslationUnitCursor().Visit(visitNode)

	log.Printf("done: filename: %+v\n", filename)

	return p.db.Put(filename, <-tubuf)
}

// SerializeTranslationUnit selialize the TranslationUnit to Clang serialized representation.
// TODO(zchee): Avoid ioutil.TempFile, get directly if possible.
func (p *Parser) SerializeTranslationUnit(filename string, tu clang.TranslationUnit) []byte {
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

// PrintDiagnostics prints a diagnostics information.
func (p *Parser) PrintDiagnostics(diags []clang.Diagnostic) {
	for _, d := range diags {
		file, line, col, offset := d.Location().FileLocation()
		fmt.Println("Location:", file.Name(), line, col, offset)
		fmt.Println("PROBLEM:", d.Spelling())
	}
}

// SymbolDB represents a symbol database.
type SymbolDB struct {
	name    string
	builder *flatbuffers.Builder
	db      *symbol.SymbolDatabase
}

// NewSymbolDB returnn the new SymbolDB.
func NewSymbolDB(filename string) *SymbolDB {
	db := &symbol.SymbolDatabase{
		Symbols: make(map[string]*symbol.Symbol),
	}
	return &SymbolDB{
		name:    filename,
		builder: flatbuffers.NewBuilder(0),
		db:      db,
	}
}

// addSymbol adds the symbol data into SymbolDB.
func (s *SymbolDB) addSymbol(sym, def *symbol.Location) {
	syms, ok := s.db.Symbols[s.name]
	if !ok {
		syms = new(symbol.Symbol)
	}
	syms.Decls = append(syms.Decls, sym)

	if def != nil {
		syms.Def = def
	}

	s.db.Symbols[s.name] = syms
}

// AddDecl add decl data into SymbolDB.
func (s *SymbolDB) AddDecl(sym *symbol.Location) {
	s.addSymbol(sym, nil)
}

// AddDefinition add definition data into SymbolDB.
func (s *SymbolDB) AddDefinition(sym, def *symbol.Location) {
	s.addSymbol(sym, def)
}

// AddCaller add caller data into SymbolDB.
func (s *SymbolDB) AddCaller(sym, def *symbol.Location, funcCall bool) {
	syms, ok := s.db.Symbols[s.name]
	if !ok {
		syms = new(symbol.Symbol)
	}

	syms.Callers = append(syms.Callers, &symbol.Caller{
		Location: sym,
		FuncCall: funcCall,
	})

	s.db.Symbols[s.name] = syms
}

// notExisHeaderName return the not exist header name magic words.
func notExistHeaderName(headPath string) string {
	// adding magic to filename to not confuse it with real files
	return "IDoNotReallyExist-" + filepath.Base(headPath)
}

// AddHeader add header data into SymbolDB.
func (s *SymbolDB) AddHeader(includePath string, headerFile clang.File) {
	hdr := new(symbol.Header)
	if headerFile.Name() == "" {
		hdr.File = symbol.ToFileID(notExistHeaderName(filepath.Clean(headerFile.Name())))
		hdr.Mtime = time.Time{}
	} else {
		hdr.File = symbol.ToFileID(filepath.Clean(headerFile.Name()))
		hdr.Mtime = headerFile.Time()
	}
	s.db.Headers = append(s.db.Headers, hdr)
}

// ClangVersion return the current clang version.
func ClangVersion() string {
	return clang.GetClangVersion()
}
