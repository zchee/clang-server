// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package symbol

import (
	"time"

	"github.com/go-clang/v3.9/clang"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/zchee/clang-server/internal/hashutil"
	"github.com/zchee/clang-server/symbol/internal/symbol"
)

// ----------------------------------------------------------------------------

// File represents a C/C++ file.
//
// table File {
//   FileName: string (required, key);
//   TranslationUnit: string;
//   SymbolDatabase: SymbolDatabase;
// }
type File struct {
	Name            string
	TranslationUnit []byte
	SymbolDatabase  *SymbolDatabase
	file            *symbol.File
}

func GetRootAsFile(buf []byte, offset flatbuffers.UOffsetT) *File {
	return &File{file: symbol.GetRootAsFile(buf, offset)}
}

func (f *File) Init(buf []byte) {
	f.file.Init(buf, flatbuffers.GetUOffsetT(buf))
}

func (f *File) Table() flatbuffers.Table {
	return f.file.Table()
}

func (f *File) GetFileName() []byte {
	return f.file.FileName()
}

func (f *File) GetSymbolDatabase() *SymbolDatabase {
	obj := new(symbol.SymbolDatabase)
	f.file.SymbolDatabase(obj)
	return &SymbolDatabase{symbolDB: obj}
}

func (f *File) GetTranslationUnit() []byte {
	return f.file.TranslationUnit()
}

// serialize

func CreateFile(builder *flatbuffers.Builder, filename string, translationUnit []byte, symbolDB flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	f := builder.CreateString(filename)
	tu := builder.CreateByteString(translationUnit)

	symbol.FileStart(builder)

	symbol.FileAddSymbolDatabase(builder, symbolDB)
	symbol.FileAddTranslationUnit(builder, tu)
	symbol.FileAddFileName(builder, f)

	return symbol.FileEnd(builder)
}

// deserialize

// ----------------------------------------------------------------------------

// SymbolDatabase database of C/C++ symbols.
//
// table SymbolDatabase {
//   Symbols: [Symbol] (id: 0);
//   Headers: [Header] (id: 1);
//   Includes: [string] (id: 2);
//   LastModified: long (id : 3); // time.Time.Unix(): int64
// }
type SymbolDatabase struct {
	Symbols      map[string]*Symbol
	Headers      []*Header
	Includes     []string
	LastModified time.Time
	symbolDB     *symbol.SymbolDatabase
}

func GetRootAsSymbolDatabase(buf []byte, offset flatbuffers.UOffsetT) *SymbolDatabase {
	return &SymbolDatabase{symbolDB: symbol.GetRootAsSymbolDatabase(buf, offset)}
}

func (sd *SymbolDatabase) Init(buf []byte) {
	sd.symbolDB.Init(buf, flatbuffers.GetUOffsetT(buf))
}

func (sd *SymbolDatabase) Table() flatbuffers.Table {
	return sd.symbolDB.Table()
}

func (sd *SymbolDatabase) Symbol(i int) *Symbol {
	obj := new(symbol.Symbol)
	if sd.symbolDB.SymbolsLength() <= i || !sd.symbolDB.Symbols(obj, i) {
		return nil
	}
	return &Symbol{sym: obj}
}

func (sd *SymbolDatabase) Header(i int) *Header {
	obj := new(symbol.Header)
	if sd.symbolDB.HeadersLength() <= i || !sd.symbolDB.Headers(obj, i) {
		return nil
	}
	return &Header{header: obj}
}

func (sd *SymbolDatabase) HeadersLength() int {
	return sd.symbolDB.HeadersLength()
}

func (sd *SymbolDatabase) Include(i int) []byte {
	if sd.symbolDB.IncludesLength() <= i {
		return nil
	}
	return sd.symbolDB.Includes(i)
}

func (sd *SymbolDatabase) IncludesLength() int {
	return sd.symbolDB.IncludesLength()
}

func (sd *SymbolDatabase) ModifyTime() int64 {
	return sd.symbolDB.Mtime()
}

// serialize

func CreateSymbolDatabase(builder *flatbuffers.Builder, symbols, headers, includes flatbuffers.UOffsetT, mtime time.Time) flatbuffers.UOffsetT {
	symbol.SymbolDatabaseStart(builder)

	symbol.SymbolDatabaseAddMtime(builder, mtime.Unix())
	symbol.SymbolDatabaseAddIncludes(builder, includes) // includes: []string
	symbol.SymbolDatabaseAddHeaders(builder, headers)   // header: []symbol.Header
	symbol.SymbolDatabaseAddSymbols(builder, symbols)   // symbols: []symbol.Symbol

	return symbol.SymbolDatabaseEnd(builder)
}

func CreateIncludes(builder *flatbuffers.Builder, includes []string) flatbuffers.UOffsetT {
	includeVector := []flatbuffers.UOffsetT{}
	for _, include := range includes {
		filename := builder.CreateString(include)
		includeVector = append(includeVector, filename)
	}

	n := len(includes)
	symbol.SymbolDatabaseStartIncludesVector(builder, n)
	for i := n - 1; i > -1; i-- {
		builder.PrependUOffsetT(includeVector[i])
	}

	return builder.EndVector(n)
}

// deserialize

// ----------------------------------------------------------------------------

// Symbol represents a location of C/C++ cursor symbol.
//
// table Symbol {
//   ID: string (id: 0, required, key);
//   Definition: Location (id: 1);
//   Decls: [Location] (id: 2);
//   Callers: [Caller] (id: 3);
// }
type Symbol struct {
	ID      ID
	Def     *Location
	Decls   []*Location
	Callers []*Caller
	sym     *symbol.Symbol
}

func GetRootAsSymbol(buf []byte, offset flatbuffers.UOffsetT) *Symbol {
	return &Symbol{sym: symbol.GetRootAsSymbol(buf, offset)}
}

func (s *Symbol) Init(buf []byte) {
	s.sym.Init(buf, flatbuffers.GetUOffsetT(buf))
}

func (s *Symbol) Table() flatbuffers.Table {
	return s.sym.Table()
}

func (s *Symbol) SymbolID() []byte {
	return s.sym.ID()
}

func (s *Symbol) Definition() *Location {
	loc := new(symbol.Location)
	loc = s.sym.Definition(loc)

	return &Location{location: loc}
}

func (s *Symbol) Decl(obj *symbol.Location, i int) *Location {
	if s.sym.DeclsLength() <= i || !s.sym.Decls(obj, i) {
		return nil
	}
	return &Location{location: obj}
}

func (s *Symbol) Caller(obj *symbol.Caller, i int) *Caller {
	if s.sym.CallersLength() <= i || !s.sym.Callers(obj, i) {
		return nil
	}
	return &Caller{caller: obj}
}

// serialize

func CreateSymbol(builder *flatbuffers.Builder, id string, def flatbuffers.UOffsetT, decls, callers []flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	fbsID := builder.CreateString(id)

	n := len(decls)
	symbol.SymbolStartDeclsVector(builder, n)
	for i := n - 1; i > -1; i-- {
		builder.PrependUOffsetT(decls[i])
	}
	endDecls := builder.EndVector(n)

	n = len(callers)
	symbol.SymbolStartDeclsVector(builder, n)
	for i := n - 1; i > -1; i-- {
		builder.PrependUOffsetT(callers[i])
	}
	endCallers := builder.EndVector(n)

	symbol.SymbolStart(builder)

	symbol.SymbolAddCallers(builder, endCallers) // endCallers: symbol.Caller
	symbol.SymbolAddDecls(builder, endDecls)     // endDecls: []symbol.Location
	symbol.SymbolAddDefinition(builder, def)     // def: symbol.Location
	symbol.SymbolAddID(builder, fbsID)

	return symbol.SymbolEnd(builder)
}

func CreateDecls(builder *flatbuffers.Builder, src []*Location) []flatbuffers.UOffsetT {
	var decls []flatbuffers.UOffsetT

	for i := 0; i < len(src); i++ {
		file := builder.CreateByteString(src[i].FileID.Bytes())

		symbol.LocationStart(builder)
		symbol.LocationAddFile(builder, file)
		symbol.LocationAddLine(builder, src[i].Line)
		symbol.LocationAddCol(builder, src[i].Col)
		symbol.LocationAddOffset(builder, src[i].Offset)

		location := symbol.LocationEnd(builder)
		decls = append(decls, location)
	}

	return decls
}

// deserialize

func (s *Symbol) GetDecls() map[FileID]*Location {
	n := s.sym.DeclsLength()
	decls := make(map[FileID]*Location)
	obj := new(symbol.Location)

	for i := 0; i < n; i++ {
		decl := s.Decl(obj, i)
		fid := ToFileID(string(decl.location.File()))
		decls[fid] = decl
	}

	return decls
}

func (s *Symbol) GetCallers() []*Caller {
	n := s.sym.CallersLength()
	callers := make([]*Caller, n)

	obj := new(symbol.Caller)
	for i := 0; i < n; i++ {
		callers[i] = s.Caller(obj, i)
	}

	return callers
}

// Marshal serializes symbols.
// WIP
func (s *Symbol) marshal() ([]byte, error) {
	b := flatbuffers.NewBuilder(0)
	n := s.sym.DeclsLength()

	off := make([]flatbuffers.UOffsetT, n)
	var v symbol.Location
	for i := 0; i < n; i++ {
		f := b.CreateByteVector(v.File())
		symbol.LocationStart(b)
		symbol.LocationAddFile(b, f)
		symbol.LocationAddLine(b, v.Line())
		symbol.LocationAddCol(b, v.Col())
		symbol.LocationAddOffset(b, v.Offset())
		off[i] = symbol.LocationEnd(b)
	}

	symbol.SymbolStartDeclsVector(b, n)
	for i := n - 1; i >= 0; i-- {
		b.PrependUOffsetT(off[i])
	}
	declVecOffset := b.EndVector(n)

	symbol.SymbolStart(b)
	symbol.SymbolAddDecls(b, declVecOffset)
	b.Finish(symbol.SymbolEnd(b))
	return b.FinishedBytes(), nil
}

// ----------------------------------------------------------------------------

// Header represents a location of include header file.
//
// table Header {
//   FileID: string (id: 0, required, key);
//   Mtime: long (id: 1);
// }
type Header struct {
	File   FileID
	Mtime  time.Time
	header *symbol.Header
}

func GetRootAsHeader(buf []byte, offset flatbuffers.UOffsetT) *Header {
	return &Header{header: symbol.GetRootAsHeader(buf, offset)}
}

func (h *Header) Init(buf []byte) {
	h.header.Init(buf, flatbuffers.GetUOffsetT(buf))
}

func (h *Header) Table() flatbuffers.Table {
	return h.header.Table()
}

func (h *Header) FileID() []byte {
	return h.header.FileID()
}

func (h *Header) ModifyTime() int64 {
	return h.header.Mtime()
}

// serialize

func CreateHeader(builder *flatbuffers.Builder, fileID []byte, mtime time.Time) flatbuffers.UOffsetT {
	id := builder.CreateByteString(fileID)

	symbol.HeaderStart(builder)
	symbol.HeaderAddMtime(builder, mtime.Unix())
	symbol.HeaderAddFileID(builder, id)

	return symbol.HeaderEnd(builder)
}

// ----------------------------------------------------------------------------

// Caller represents a location of caller function.
//
// table Caller {
//   Location: Location (required);
//   FuncCall: bool = false;
// }
type Caller struct {
	Location *Location
	FuncCall bool
	caller   *symbol.Caller
}

func GetRootAsCaller(buf []byte, offset flatbuffers.UOffsetT) *Caller {
	return &Caller{caller: symbol.GetRootAsCaller(buf, offset)}
}

func (c *Caller) Init(buf []byte) {
	c.caller.Init(buf, flatbuffers.GetUOffsetT(buf))
}

func (c *Caller) Table() flatbuffers.Table {
	return c.caller.Table()
}

func (c *Caller) IsFuncCall() bool {
	return c.caller.FuncCall() != 0
}

// serialize

func CreateCallers(builder *flatbuffers.Builder, src []Caller) []flatbuffers.UOffsetT {
	var callers []flatbuffers.UOffsetT

	for i := 0; i < len(src); i++ {
		location := CreateLocation(builder, src[i].Location)
		var fc byte
		if src[i].FuncCall {
			fc = byte(1)
		}

		symbol.CallerStart(builder)
		symbol.CallerAddFuncCall(builder, fc)
		symbol.CallerAddLocation(builder, location)

		caller := symbol.CallerEnd(builder)
		callers = append(callers, caller)
	}

	return callers
}

// deserialize

func (c *Caller) GetLocation() (string, uint32, uint32) {
	obj := new(symbol.Location)
	loc := &Location{location: c.caller.Location(obj)}
	return loc.GetLocation()
}

// ----------------------------------------------------------------------------

// Location location of symbol.
//
// table Location {
//   File: string (required);
//   Line: uint;    // clang.SourceLocation.Line: uint32
//   Col: uint = 0; // clang.SourceLocation.Col: uint32
//   Offset: uint;  // clang.SourceLocation.Offset: uint32
// }
type Location struct {
	File     string
	FileID   FileID
	Line     uint32
	Col      uint32
	Offset   uint32
	USR      string
	location *symbol.Location
}

func GetRootAsLocation(buf []byte, offset flatbuffers.UOffsetT) *Location {
	return &Location{location: symbol.GetRootAsLocation(buf, offset)}
}

func (l *Location) Init(buf []byte) {
	l.location.Init(buf, flatbuffers.GetUOffsetT(buf))
}

func (l *Location) Table() flatbuffers.Table {
	return l.location.Table()
}

// serialize

func CreateLocation(builder *flatbuffers.Builder, src *Location) flatbuffers.UOffsetT {
	file := builder.CreateString(src.FileID.String())

	symbol.LocationStart(builder)

	symbol.LocationAddOffset(builder, src.Offset)
	symbol.LocationAddCol(builder, src.Col)
	symbol.LocationAddLine(builder, src.Line)
	symbol.LocationAddFile(builder, file)

	return symbol.LocationEnd(builder)
}

// deserialize

func (l *Location) GetLocation() (string, uint32, uint32) {
	return hashutil.UnsafeString(l.location.File()), l.location.Line(), l.location.Col()
}

// ----------------------------------------------------------------------------
// WIP

func (sd *SymbolDatabase) decode() ([]*symbol.Symbol, int64) {
	n := sd.symbolDB.SymbolsLength()
	syms := make([]*symbol.Symbol, n)

	sym := new(symbol.Symbol)
	for i := 0; i < n; i++ {
		if ok := sd.symbolDB.Symbols(sym, i); ok {
			syms[i] = sym
		}
	}

	return syms, sd.ModifyTime()
}
