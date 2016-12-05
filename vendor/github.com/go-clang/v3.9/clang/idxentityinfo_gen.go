package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

type IdxEntityInfo struct {
	c *C.CXIdxEntityInfo
}

// For retrieving a custom CXIdxClientEntity attached to an entity.
func (iei *IdxEntityInfo) ClientEntity() IdxClientEntity {
	return IdxClientEntity{C.clang_index_getClientEntity(iei.c)}
}

// For setting a custom CXIdxClientEntity attached to an entity.
func (iei *IdxEntityInfo) SetClientEntity(ice IdxClientEntity) {
	C.clang_index_setClientEntity(iei.c, ice.c)
}

func (iei IdxEntityInfo) Kind() IdxEntityKind {
	return IdxEntityKind(iei.c.kind)
}

func (iei IdxEntityInfo) TemplateKind() IdxEntityCXXTemplateKind {
	return IdxEntityCXXTemplateKind(iei.c.templateKind)
}

func (iei IdxEntityInfo) Lang() IdxEntityLanguage {
	return IdxEntityLanguage(iei.c.lang)
}

func (iei IdxEntityInfo) Name() string {
	return C.GoString(iei.c.name)
}

func (iei IdxEntityInfo) USR() string {
	return C.GoString(iei.c.USR)
}

func (iei IdxEntityInfo) Cursor() Cursor {
	return Cursor{iei.c.cursor}
}

func (iei IdxEntityInfo) Attributes() []*IdxAttrInfo {
	var s []*IdxAttrInfo
	gos_s := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	gos_s.Cap = int(iei.c.numAttributes)
	gos_s.Len = int(iei.c.numAttributes)
	gos_s.Data = uintptr(unsafe.Pointer(iei.c.attributes))

	return s
}

func (iei IdxEntityInfo) NumAttributes() uint32 {
	return uint32(iei.c.numAttributes)
}
