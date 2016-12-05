package clang

// #include "./clang-c/CXString.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

type StringSet struct {
	c C.CXStringSet
}

func (ss StringSet) Strings() []cxstring {
	var s []cxstring
	gos_s := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	gos_s.Cap = int(ss.c.Count)
	gos_s.Len = int(ss.c.Count)
	gos_s.Data = uintptr(unsafe.Pointer(ss.c.Strings))

	return s
}

func (ss StringSet) Count() uint32 {
	return uint32(ss.c.Count)
}
