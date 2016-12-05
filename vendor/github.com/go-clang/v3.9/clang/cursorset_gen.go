package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// A fast container representing a set of CXCursors.
type CursorSet struct {
	c C.CXCursorSet
}

// Creates an empty CXCursorSet.
func NewCursorSet() CursorSet {
	return CursorSet{C.clang_createCXCursorSet()}
}

// Disposes a CXCursorSet and releases its associated memory.
func (cs CursorSet) Dispose() {
	C.clang_disposeCXCursorSet(cs.c)
}

/*
	Queries a CXCursorSet to see if it contains a specific CXCursor.

	Returns non-zero if the set contains the specified cursor.
*/
func (cs CursorSet) Contains(cursor Cursor) uint32 {
	return uint32(C.clang_CXCursorSet_contains(cs.c, cursor.c))
}

/*
	Inserts a CXCursor into a CXCursorSet.

	Returns zero if the CXCursor was already in the set, and non-zero otherwise.
*/
func (cs CursorSet) Insert(cursor Cursor) uint32 {
	return uint32(C.clang_CXCursorSet_insert(cs.c, cursor.c))
}
