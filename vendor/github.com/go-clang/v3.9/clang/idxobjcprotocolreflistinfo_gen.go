package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

type IdxObjCProtocolRefListInfo struct {
	c C.CXIdxObjCProtocolRefListInfo
}

func (iocprli IdxObjCProtocolRefListInfo) Protocols() []*IdxObjCProtocolRefInfo {
	var s []*IdxObjCProtocolRefInfo
	gos_s := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	gos_s.Cap = int(iocprli.c.numProtocols)
	gos_s.Len = int(iocprli.c.numProtocols)
	gos_s.Data = uintptr(unsafe.Pointer(iocprli.c.protocols))

	return s
}

func (iocprli IdxObjCProtocolRefListInfo) NumProtocols() uint32 {
	return uint32(iocprli.c.numProtocols)
}
