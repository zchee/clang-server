package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

// Identifies an array of ranges.
type SourceRangeList struct {
	c C.CXSourceRangeList
}

// The number of ranges in the ranges array.
func (srl SourceRangeList) Count() uint32 {
	return uint32(srl.c.count)
}

// An array of CXSourceRanges.
func (srl SourceRangeList) Ranges() []SourceRange {
	var s []SourceRange
	gos_s := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	gos_s.Cap = int(srl.c.count)
	gos_s.Len = int(srl.c.count)
	gos_s.Data = uintptr(unsafe.Pointer(srl.c.ranges))

	return s
}
