package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type TUResourceUsageEntry struct {
	c C.CXTUResourceUsageEntry
}

func (turue TUResourceUsageEntry) Kind() TUResourceUsageKind {
	return TUResourceUsageKind(turue.c.kind)
}

func (turue TUResourceUsageEntry) Amount() uint64 {
	return uint64(turue.c.amount)
}
