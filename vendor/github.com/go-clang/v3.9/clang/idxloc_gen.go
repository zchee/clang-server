package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Source location passed to index callbacks.
type IdxLoc struct {
	c C.CXIdxLoc
}

/*
	Retrieve the CXIdxFile, file, line, column, and offset represented by
	the given CXIdxLoc.

	If the location refers into a macro expansion, retrieves the
	location of the macro expansion and if it refers into a macro argument
	retrieves the location of the argument.
*/
func (il IdxLoc) FileLocation() (IdxClientFile, File, uint32, uint32, uint32) {
	var indexFile IdxClientFile
	var file File
	var line C.uint
	var column C.uint
	var offset C.uint

	C.clang_indexLoc_getFileLocation(il.c, &indexFile.c, &file.c, &line, &column, &offset)

	return indexFile, file, uint32(line), uint32(column), uint32(offset)
}

// Retrieve the CXSourceLocation represented by the given CXIdxLoc.
func (il IdxLoc) SourceLocation() SourceLocation {
	return SourceLocation{C.clang_indexLoc_getCXSourceLocation(il.c)}
}
