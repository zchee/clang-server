package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

type IdxDeclInfo struct {
	c *C.CXIdxDeclInfo
}

func (idi *IdxDeclInfo) ContainerDeclInfo() *IdxObjCContainerDeclInfo {
	o := C.clang_index_getObjCContainerDeclInfo(idi.c)

	var gop_o *IdxObjCContainerDeclInfo
	if o != nil {
		gop_o = &IdxObjCContainerDeclInfo{*o}
	}

	return gop_o
}

func (idi *IdxDeclInfo) InterfaceDeclInfo() *IdxObjCInterfaceDeclInfo {
	o := C.clang_index_getObjCInterfaceDeclInfo(idi.c)

	var gop_o *IdxObjCInterfaceDeclInfo
	if o != nil {
		gop_o = &IdxObjCInterfaceDeclInfo{*o}
	}

	return gop_o
}

func (idi *IdxDeclInfo) CategoryDeclInfo() *IdxObjCCategoryDeclInfo {
	o := C.clang_index_getObjCCategoryDeclInfo(idi.c)

	var gop_o *IdxObjCCategoryDeclInfo
	if o != nil {
		gop_o = &IdxObjCCategoryDeclInfo{*o}
	}

	return gop_o
}

func (idi *IdxDeclInfo) ProtocolRefListInfo() *IdxObjCProtocolRefListInfo {
	o := C.clang_index_getObjCProtocolRefListInfo(idi.c)

	var gop_o *IdxObjCProtocolRefListInfo
	if o != nil {
		gop_o = &IdxObjCProtocolRefListInfo{*o}
	}

	return gop_o
}

func (idi *IdxDeclInfo) PropertyDeclInfo() *IdxObjCPropertyDeclInfo {
	o := C.clang_index_getObjCPropertyDeclInfo(idi.c)

	var gop_o *IdxObjCPropertyDeclInfo
	if o != nil {
		gop_o = &IdxObjCPropertyDeclInfo{*o}
	}

	return gop_o
}

func (idi *IdxDeclInfo) ClassDeclInfo() *IdxCXXClassDeclInfo {
	o := C.clang_index_getCXXClassDeclInfo(idi.c)

	var gop_o *IdxCXXClassDeclInfo
	if o != nil {
		gop_o = &IdxCXXClassDeclInfo{*o}
	}

	return gop_o
}

func (idi IdxDeclInfo) EntityInfo() *IdxEntityInfo {
	o := idi.c.entityInfo

	var gop_o *IdxEntityInfo
	if o != nil {
		gop_o = &IdxEntityInfo{o}
	}

	return gop_o
}

func (idi IdxDeclInfo) Cursor() Cursor {
	return Cursor{idi.c.cursor}
}

func (idi IdxDeclInfo) Loc() IdxLoc {
	return IdxLoc{idi.c.loc}
}

func (idi IdxDeclInfo) SemanticContainer() *IdxContainerInfo {
	o := idi.c.semanticContainer

	var gop_o *IdxContainerInfo
	if o != nil {
		gop_o = &IdxContainerInfo{o}
	}

	return gop_o
}

// Generally same as #semanticContainer but can be different in cases like out-of-line C++ member functions.
func (idi IdxDeclInfo) LexicalContainer() *IdxContainerInfo {
	o := idi.c.lexicalContainer

	var gop_o *IdxContainerInfo
	if o != nil {
		gop_o = &IdxContainerInfo{o}
	}

	return gop_o
}

func (idi IdxDeclInfo) IsRedeclaration() bool {
	o := idi.c.isRedeclaration

	return o != C.int(0)
}

func (idi IdxDeclInfo) IsDefinition() bool {
	o := idi.c.isDefinition

	return o != C.int(0)
}

func (idi IdxDeclInfo) IsContainer() bool {
	o := idi.c.isContainer

	return o != C.int(0)
}

func (idi IdxDeclInfo) DeclAsContainer() *IdxContainerInfo {
	o := idi.c.declAsContainer

	var gop_o *IdxContainerInfo
	if o != nil {
		gop_o = &IdxContainerInfo{o}
	}

	return gop_o
}

// Whether the declaration exists in code or was created implicitly by the compiler, e.g. implicit Objective-C methods for properties.
func (idi IdxDeclInfo) IsImplicit() bool {
	o := idi.c.isImplicit

	return o != C.int(0)
}

func (idi IdxDeclInfo) Attributes() []*IdxAttrInfo {
	var s []*IdxAttrInfo
	gos_s := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	gos_s.Cap = int(idi.c.numAttributes)
	gos_s.Len = int(idi.c.numAttributes)
	gos_s.Data = uintptr(unsafe.Pointer(idi.c.attributes))

	return s
}

func (idi IdxDeclInfo) NumAttributes() uint32 {
	return uint32(idi.c.numAttributes)
}

func (idi IdxDeclInfo) Flags() uint32 {
	return uint32(idi.c.flags)
}
