package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Describes a single preprocessing token.
type Token struct {
	c C.CXToken
}

// Determine the kind of the given token.
func (t Token) Kind() TokenKind {
	return TokenKind(C.clang_getTokenKind(t.c))
}
