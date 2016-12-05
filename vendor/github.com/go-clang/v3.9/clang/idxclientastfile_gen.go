package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// The client's data object that is associated with an AST file (PCH or module).
type IdxClientASTFile struct {
	c C.CXIdxClientASTFile
}
