// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package symbol

import (
	"path/filepath"
	"reflect"
	"time"

	"github.com/go-clang/v3.9/clang"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/zchee/clang-server/symbol/internal/symbol"
)

// ----------------------------------------------------------------------------

// File represents a C/C++ source file.
//
//  table File {
//    Name: string;
//    TranslationUnit: string;
//    Symbols: [Info];
//    Headers: [Header];
//    Includes: [string];
//  }
type File struct {
	name            string
	flags           []string
	translationUnit []byte
	locations       map[Location]ID
	symbols         map[ID]*Info
	headers         []*Header

	builder *flatbuffers.Builder

	file *symbol.File
}

// NewFile return the new File.
func NewFile(name string, flags []string) *File {
	return &File{
		name:      name,
		flags:     flags,
		locations: make(map[Location]ID),
		symbols:   make(map[ID]*Info),
		builder:   flatbuffers.NewBuilder(0),
	}
}

// GetRootAsFile gets the root of flatbuffers binary.
func GetRootAsFile(buf []byte, offset flatbuffers.UOffsetT) *File {
	return &File{
		file: symbol.GetRootAsFile(buf, offset),
	}
}

// Name return the filename.
func (f *File) Name() string {
	if f.name != "" {
		return f.name
	}
	return string(f.file.Name())
}

// TranslationUnit return the libclang translation unit data.
func (f *File) TranslationUnit() []byte {
	if len(f.translationUnit) > 0 {
		return f.translationUnit
	}
	return f.file.TranslationUnit()
}

// Symbols return the C/C++ files symbols.
func (f *File) Symbols() []*Info {
	if len(f.symbols) > 0 {
		symbols := make([]*Info, len(f.symbols))
		for _, v := range f.symbols {
			symbols = append(symbols, v)
		}
		return symbols
	}

	n := f.file.SymbolsLength()
	symbols := make([]*Info, n)

	for i := 0; i < n; i++ {
		obj := new(symbol.Info)
		if f.file.Symbols(obj, i) {
			symbols[i] = &Info{info: obj}
		}
	}

	return symbols
}

// Header return the C/C++ files included header files.
func (f *File) Headers() []*Header {
	if len(f.headers) > 0 {
		return f.headers
	}

	n := f.file.HeadersLength()
	headers := make([]*Header, n)

	for i := 0; i < n; i++ {
		obj := new(symbol.Header)
		if f.file.Headers(obj, i) {
			headers[i] = &Header{header: obj}
		}
	}

	return headers
}

// AddTranslationUnit add TranslationUnit data to File.
func (f *File) AddTranslationUnit(buf []byte) {
	f.translationUnit = buf
}

// AddSymbol adds the symbol data into File.
func (f *File) addSymbol(loc, def Location) {
	id := ToID(loc.usr)

	sym, ok := f.symbols[id]
	if !ok {
		sym = &Info{id: id}
	}
	sym.decls = append(sym.decls, loc)

	if def.isExist() {
		sym.def = def
	}

	f.locations[loc] = id
	f.symbols[id] = sym
}

// AddDecl add decl data into File.
func (f *File) AddDecl(loc Location) {
	f.addSymbol(loc, Location{})
}

// AddDefinition add definition data into File.
func (f *File) AddDefinition(loc, def Location) {
	f.addSymbol(loc, def)
}

// notExistHeaderName return the not exist header name magic words.
func notExistHeaderName(headPath string) string {
	// adding magic to filename to not confuse it with real files
	return "IDoNotReallyExist-" + filepath.Base(headPath)
}

// AddHeader add header data into File.
func (f *File) AddHeader(includePath string, headerFile clang.File) {
	hdr := new(Header)
	if headerFile.Name() == "" {
		hdr.fileid = ToFileID(notExistHeaderName(filepath.Clean(headerFile.Name())))
		hdr.mtime = time.Now()
	} else {
		hdr.fileid = ToFileID(filepath.Clean(headerFile.Name()))
		hdr.mtime = headerFile.Time()
	}

	f.headers = append(f.headers, hdr)
}

// AddCaller add caller data into File.
func (f *File) AddCaller(sym, def Location, funcCall bool) {
	id := ToID(sym.usr)

	syms, ok := f.symbols[id]
	if !ok {
		syms = &Info{id: id}
	}

	syms.callers = append(syms.callers, &Caller{
		location: sym,
		funcCall: funcCall,
	})

	f.symbols[id] = syms
}

// Unmarshal parses the flatbuffers representation in f.
func (f *File) Unmarshal() {
	f.name = string(f.file.Name())
	f.translationUnit = f.file.TranslationUnit()
	f.symbols = make(map[ID]*Info)
	for _, s := range f.Symbols() {
		f.symbols[s.ID()] = &Info{
			id:      s.ID(),
			decls:   s.Decls(),
			def:     s.Def(),
			callers: s.Callers(),
			info:    s.info,
		}
	}
	headers := f.Headers()
	f.headers = make([]*Header, len(headers))
	for _, hdr := range headers {
		f.headers = append(f.headers, hdr)
	}
}

// Serialize serializes the File.
func (f *File) Serialize() *flatbuffers.Builder {
	if f.builder == nil {
		f.builder = flatbuffers.NewBuilder(0)
	}

	fname := f.builder.CreateString(f.Name())
	tu := f.builder.CreateByteString(f.TranslationUnit())

	flagNum := len(f.flags)
	flagOffsets := make([]flatbuffers.UOffsetT, 0, flagNum)
	for _, flag := range f.flags {
		flagOffsets = append(flagOffsets, f.builder.CreateString(flag))
	}
	symbol.FileStartFlagsVector(f.builder, flagNum)
	for i := flagNum - 1; i >= 0; i-- {
		f.builder.PrependUOffsetT(flagOffsets[i])
	}
	flagVecOffset := f.builder.EndVector(flagNum)

	symbols := f.symbols
	symbolNum := len(symbols)
	symbolOffsets := make([]flatbuffers.UOffsetT, 0, symbolNum)
	for _, info := range symbols {
		symbolOffsets = append(symbolOffsets, info.serialize(f.builder))
	}
	symbol.FileStartSymbolsVector(f.builder, symbolNum)
	for i := symbolNum - 1; i >= 0; i-- {
		f.builder.PrependUOffsetT(symbolOffsets[i])
	}
	symbolVecOffset := f.builder.EndVector(symbolNum)

	hdrs := f.headers
	hdrNum := len(hdrs)
	hdrOffsets := make([]flatbuffers.UOffsetT, 0, hdrNum)
	for _, hdr := range hdrs {
		hdrOffsets = append(hdrOffsets, hdr.serialize(f.builder))
	}
	symbol.FileStartHeadersVector(f.builder, hdrNum)
	for i := hdrNum - 1; i >= 0; i-- {
		f.builder.PrependUOffsetT(hdrOffsets[i])
	}
	headerVecOffset := f.builder.EndVector(hdrNum)

	symbol.FileStart(f.builder)
	symbol.FileAddName(f.builder, fname)
	symbol.FileAddFlags(f.builder, flagVecOffset)
	symbol.FileAddTranslationUnit(f.builder, tu)
	symbol.FileAddSymbols(f.builder, symbolVecOffset)
	symbol.FileAddHeaders(f.builder, headerVecOffset)

	f.builder.Finish(symbol.FileEnd(f.builder))

	return f.builder
}

// ----------------------------------------------------------------------------

// Info represents a location of C/C++ cursor symbol information.
//
//  table Info {
//    ID: string;
//    Decls: [Location];
//    Def: Location;
//  }
type Info struct {
	id      ID
	decls   []Location
	def     Location
	callers []*Caller

	info *symbol.Info
}

// serialize serializes the Info.
func (info *Info) serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	id := builder.CreateString(info.id.String())

	declsNum := len(info.decls)
	var declVecOffset flatbuffers.UOffsetT
	if declsNum > 0 {
		declsOffsets := make([]flatbuffers.UOffsetT, 0, declsNum)
		for _, decl := range info.decls {
			declsOffsets = append(declsOffsets, decl.serialize(builder))
		}
		symbol.InfoStartDeclsVector(builder, declsNum)
		for i := declsNum - 1; i >= 0; i-- {
			builder.PrependUOffsetT(declsOffsets[i])
		}
		declVecOffset = builder.EndVector(declsNum)
	}

	defOffset := info.def.serialize(builder)

	callersNum := len(info.callers)
	var callerVecOffset flatbuffers.UOffsetT
	if callersNum > 0 {
		callersOffsets := make([]flatbuffers.UOffsetT, 0, callersNum)
		for _, caller := range info.callers {
			callersOffsets = append(callersOffsets, caller.serialize(builder))
		}
		symbol.InfoStartCallersVector(builder, callersNum)
		for i := callersNum - 1; i >= 0; i-- {
			builder.PrependUOffsetT(callersOffsets[i])
		}
		callerVecOffset = builder.EndVector(callersNum)
	}

	symbol.InfoStart(builder)
	symbol.InfoAddID(builder, id)
	symbol.InfoAddDecls(builder, declVecOffset)
	symbol.InfoAddDef(builder, defOffset)
	symbol.InfoAddCallers(builder, callerVecOffset)

	return symbol.InfoEnd(builder)
}

// ID return the symbol ID which hashed blake2b.
func (info *Info) ID() ID {
	return ToID(string(info.info.ID()))
}

// Decls return the symbol declarations information.
func (info *Info) Decls() []Location {
	n := info.info.DeclsLength()
	decls := make([]Location, n)

	for i := 0; i < n; i++ {
		obj := new(symbol.Location)
		if info.info.Decls(obj, i) {
			decls[i] = Location{location: obj}
		}
	}

	return decls
}

// Def return the symbol definition information.
func (info *Info) Def() Location {
	obj := new(symbol.Location)
	info.info.Def(obj)

	return Location{location: obj}
}

// Callers return the symbol callers information.
func (info *Info) Callers() []*Caller {
	n := info.info.CallersLength()
	callers := make([]*Caller, n)

	for i := 0; i < n; i++ {
		obj := new(symbol.Caller)
		if info.info.Callers(obj, i) {
			callers[i] = &Caller{caller: obj}
		}
	}

	return callers
}

// ----------------------------------------------------------------------------

// Header represents a location of include header file.
//
//  table Header {
//    FileID: string (id: 0, required, key); // -> []byte
//    Mtime: long (id: 1); // time.Time.Unix(): int64
//  }
type Header struct {
	fileid FileID
	mtime  time.Time

	header *symbol.Header
}

// FileID return the header FileID.
func (h *Header) FileID() FileID {
	return ToFileID(string(h.header.FileID()))
}

// Mtime return the header modified time.
func (h *Header) Mtime() int64 {
	return h.header.Mtime()
}

// serialize serializes the h data to flatbuffers.UOffsetT.
func (h *Header) serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	fid := builder.CreateString(h.fileid.String())

	symbol.HeaderStart(builder)

	symbol.HeaderAddFileID(builder, fid)
	symbol.HeaderAddMtime(builder, h.mtime.Unix())

	return symbol.HeaderEnd(builder)
}

// ----------------------------------------------------------------------------

// Caller represents a location of caller function.
//
//  table Caller {
//    Location: Location (required);
//    FuncCall: bool = false; // -> byte
//  }
type Caller struct {
	location Location
	funcCall bool

	caller *symbol.Caller
}

// Location return the location of caller function.
func (c *Caller) Location() Location {
	obj := new(symbol.Location)
	c.caller.Location(obj)

	return Location{location: obj}
}

// FuncCall reports whether caller is function call.
func (c *Caller) FuncCall() bool {
	return c.caller.FuncCall() != 0
}

// serialize serializes the c data to flatbuffers.UOffsetT.
func (c *Caller) serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	locOffset := c.location.serialize(builder)

	symbol.CallerStart(builder)

	symbol.CallerAddLocation(builder, locOffset)
	funcCall := byte(0)
	if c.funcCall {
		funcCall = byte(1)
	}
	symbol.CallerAddFuncCall(builder, funcCall)

	return symbol.CallerEnd(builder)
}

// ----------------------------------------------------------------------------

// Location location of symbol.
// TODO(zchee): method receiver is pointer for location?
//
//  table Location {
//    FileName: string;
//    Line: uint;
//    Col: uint = 0;
//    Offset: uint;
//    USR: string;
//  }
type Location struct {
	fileName string
	line     uint32
	col      uint32
	offset   uint32
	usr      string

	location *symbol.Location
}

// FileName return the filename of location.
func (l *Location) FileName() string {
	if l.location == nil {
		return l.fileName
	}
	return string(l.location.FileName())
}

// Line return the line number of symbol location.
func (l *Location) Line() uint32 {
	return l.location.Line()
}

// Col return the column number of symbol location.
func (l *Location) Col() uint32 {
	return l.location.Col()
}

// Offset return the byte offset of symbol location.
func (l *Location) Offset() uint32 {
	return l.location.Offset()
}

// USR return the Unified Symbol Resolution of symbol.
func (l *Location) USR() string {
	if l.location == nil {
		return l.usr
	}
	return string(l.location.USR())
}

// serialize serializes the l data to flatbuffers.UOffsetT.
func (l *Location) serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	fname := builder.CreateString(l.fileName)
	usr := builder.CreateString(l.usr)

	symbol.LocationStart(builder)

	symbol.LocationAddFileName(builder, fname)
	symbol.LocationAddLine(builder, l.line)
	symbol.LocationAddCol(builder, l.col)
	symbol.LocationAddOffset(builder, l.offset)
	symbol.LocationAddUSR(builder, usr)

	return symbol.LocationEnd(builder)
}

// TODO(zchee): avoid reflection
func (l *Location) isExist() bool {
	return !reflect.DeepEqual(l, Location{})
}

// CreateLocation creates location data using flatbuffers binary.
func CreateLocation(filename string, line, col uint32) *flatbuffers.Builder {
	l := Location{
		fileName: filename,
		line:     line,
		col:      col,
	}
	builder := flatbuffers.NewBuilder(0)

	fname := builder.CreateString(l.fileName)

	symbol.LocationStart(builder)
	symbol.LocationAddFileName(builder, fname)
	symbol.LocationAddLine(builder, l.line)
	symbol.LocationAddCol(builder, l.col)
	builder.Finish(symbol.LocationEnd(builder))

	return builder
}
