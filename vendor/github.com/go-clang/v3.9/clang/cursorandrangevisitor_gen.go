package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type CursorAndRangeVisitor struct {
	c C.CXCursorAndRangeVisitor
}
