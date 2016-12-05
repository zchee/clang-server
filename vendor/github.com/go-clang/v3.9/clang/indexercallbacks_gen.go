package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// A group of callbacks used by #clang_indexSourceFile and #clang_indexTranslationUnit.
type IndexerCallbacks struct {
	c C.IndexerCallbacks
}
