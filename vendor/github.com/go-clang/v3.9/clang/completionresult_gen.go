package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

// A single result of code completion.
type CompletionResult struct {
	c C.CXCompletionResult
}

/*
	Sort the code-completion results in case-insensitive alphabetical
	order.

	Parameter Results The set of results to sort.
	Parameter NumResults The number of results in \p Results.
*/
func SortCodeCompletionResults(results []CompletionResult) {
	gos_results := (*reflect.SliceHeader)(unsafe.Pointer(&results))
	cp_results := (*C.CXCompletionResult)(unsafe.Pointer(gos_results.Data))

	C.clang_sortCodeCompletionResults(cp_results, C.uint(len(results)))
}

/*
	The kind of entity that this completion refers to.

	The cursor kind will be a macro, keyword, or a declaration (one of the
	*Decl cursor kinds), describing the entity that the completion is
	referring to.

	\todo In the future, we would like to provide a full cursor, to allow
	the client to extract additional information from declaration.
*/
func (cr CompletionResult) CursorKind() CursorKind {
	return CursorKind(cr.c.CursorKind)
}

// The code-completion string that describes how to insert this code-completion result into the editing buffer.
func (cr CompletionResult) CompletionString() CompletionString {
	return CompletionString{cr.c.CompletionString}
}
