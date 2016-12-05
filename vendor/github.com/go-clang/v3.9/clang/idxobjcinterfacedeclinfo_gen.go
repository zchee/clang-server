package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type IdxObjCInterfaceDeclInfo struct {
	c C.CXIdxObjCInterfaceDeclInfo
}

func (iocidi IdxObjCInterfaceDeclInfo) ContainerInfo() *IdxObjCContainerDeclInfo {
	o := iocidi.c.containerInfo

	var gop_o *IdxObjCContainerDeclInfo
	if o != nil {
		gop_o = &IdxObjCContainerDeclInfo{*o}
	}

	return gop_o
}

func (iocidi IdxObjCInterfaceDeclInfo) SuperInfo() *IdxBaseClassInfo {
	o := iocidi.c.superInfo

	var gop_o *IdxBaseClassInfo
	if o != nil {
		gop_o = &IdxBaseClassInfo{*o}
	}

	return gop_o
}

func (iocidi IdxObjCInterfaceDeclInfo) Protocols() *IdxObjCProtocolRefListInfo {
	o := iocidi.c.protocols

	var gop_o *IdxObjCProtocolRefListInfo
	if o != nil {
		gop_o = &IdxObjCProtocolRefListInfo{*o}
	}

	return gop_o
}
