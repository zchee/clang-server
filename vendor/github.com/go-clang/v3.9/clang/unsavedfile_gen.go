package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

/*
	Provides the contents of a file that has not yet been saved to disk.

	Each CXUnsavedFile instance provides the name of a file on the
	system along with the current contents of that file that have not
	yet been saved to disk.
*/
type UnsavedFile struct {
	c C.struct_CXUnsavedFile
}

/*
	The file whose contents have not yet been saved.

	This file must already exist in the file system.
*/
func (uf UnsavedFile) Filename() string {
	return C.GoString(uf.c.Filename)
}

// A buffer containing the unsaved contents of this file.
func (uf UnsavedFile) Contents() string {
	return C.GoString(uf.c.Contents)
}

// The length of the unsaved contents of this buffer.
func (uf UnsavedFile) Length() uint64 {
	return uint64(uf.c.Length)
}
