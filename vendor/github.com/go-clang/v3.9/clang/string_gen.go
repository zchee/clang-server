package clang

// #include "./clang-c/CXString.h"
// #include "go-clang.h"
import "C"

/*
	A character string.

	The CXString type is used to return strings from the interface when
	the ownership of that string might differ from one call to the next.
	Use clang_getCString() to retrieve the string data and, once finished
	with the string data, call clang_disposeString() to free the string.
*/
type String struct {
	c C.CXString
}

func (s String) Private_flags() uint32 {
	return uint32(s.c.private_flags)
}
