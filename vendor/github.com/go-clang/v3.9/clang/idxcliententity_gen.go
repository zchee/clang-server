package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// The client's data object that is associated with a semantic entity.
type IdxClientEntity struct {
	c C.CXIdxClientEntity
}
