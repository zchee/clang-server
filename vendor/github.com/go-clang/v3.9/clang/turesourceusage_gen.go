package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

// The memory usage of a CXTranslationUnit, broken into categories.
type TUResourceUsage struct {
	c C.CXTUResourceUsage
}

func (turu TUResourceUsage) Dispose() {
	C.clang_disposeCXTUResourceUsage(turu.c)
}

func (turu TUResourceUsage) NumEntries() uint32 {
	return uint32(turu.c.numEntries)
}

func (turu TUResourceUsage) Entries() []TUResourceUsageEntry {
	var s []TUResourceUsageEntry
	gos_s := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	gos_s.Cap = int(turu.c.numEntries)
	gos_s.Len = int(turu.c.numEntries)
	gos_s.Data = uintptr(unsafe.Pointer(turu.c.entries))

	return s
}
