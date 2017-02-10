// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package symbol

import (
	"log"
	"path/filepath"
	"time"

	"github.com/go-clang/v3.9/clang"
	flatbuffers "github.com/google/flatbuffers/go"
)

// DB represents a symbol database.
type DB struct {
	filename        string
	translationUnit []byte
	builder         *flatbuffers.Builder

	*SymbolDatabase
}

// NewDB returnn the new SymbolDB.
func NewDB(filename string) *DB {
	db := &SymbolDatabase{
		Symbols: make(map[ID]*Symbol),
	}
	return &DB{
		filename:       filename,
		builder:        flatbuffers.NewBuilder(0),
		SymbolDatabase: db,
	}
}

func (db *DB) MarshalSymbols() []byte {
	var symbols []flatbuffers.UOffsetT

	for id, sym := range db.Symbols {
		log.Printf("id: %T => %+v\n", id, id)
		log.Printf("sym: %T => %+v\n", sym, sym)
		callers := createCallers(db.builder, sym.Callers)
		decls := createDecls(db.builder, sym.Decls)
		def := createLocation(db.builder, sym.Def)
		symbols = append(symbols, createSymbol(db.builder, id, def, decls, callers))
	}

	empty := []flatbuffers.UOffsetT{}
	symbolDB := createSymbolDatabase(db.builder, symbols, empty, empty, time.Now())

	return createFile(db.builder, db.filename, db.translationUnit, symbolDB)
}

// AddSymbol adds the symbol data into SymbolDB.
func (db *DB) addSymbol(sym, def *Location) {
	id := ToID(sym.USR)

	syms, ok := db.Symbols[id]
	if !ok {
		syms = new(Symbol)
	}
	syms.Decls = append(syms.Decls, sym)

	if def != nil {
		syms.Def = def
	}

	db.Symbols[id] = syms
}

// AddDecl add decl data into SymbolDB.
func (db *DB) AddDecl(sym *Location) {
	db.addSymbol(sym, nil)
}

// AddDefinition add definition data into SymbolDB.
func (db *DB) AddDefinition(sym, def *Location) {
	db.addSymbol(sym, def)
}

// notExistHeaderName return the not exist header name magic words.
func notExistHeaderName(headPath string) string {
	// adding magic to filename to not confuse it with real files
	return "IDoNotReallyExist-" + filepath.Base(headPath)
}

// AddHeader add header data into SymbolDB.
func (db *DB) AddHeader(includePath string, headerFile clang.File) {
	hdr := new(Header)
	if headerFile.Name() == "" {
		hdr.File = ToFileID(notExistHeaderName(filepath.Clean(headerFile.Name())))
		hdr.Mtime = time.Time{}
	} else {
		hdr.File = ToFileID(filepath.Clean(headerFile.Name()))
		hdr.Mtime = headerFile.Time()
	}
	db.Headers = append(db.Headers, hdr)
}

// AddCaller add caller data into SymbolDB.
func (db *DB) AddCaller(sym, def *Location, funcCall bool) {
	id := ToID(sym.USR)

	syms, ok := db.Symbols[id]
	if !ok {
		syms = new(Symbol)
	}

	syms.Callers = append(syms.Callers, &Caller{
		Location: sym,
		FuncCall: funcCall,
	})

	db.Symbols[id] = syms
}
