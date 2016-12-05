package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type IdxObjCCategoryDeclInfo struct {
	c C.CXIdxObjCCategoryDeclInfo
}

func (ioccdi IdxObjCCategoryDeclInfo) ContainerInfo() *IdxObjCContainerDeclInfo {
	o := ioccdi.c.containerInfo

	var gop_o *IdxObjCContainerDeclInfo
	if o != nil {
		gop_o = &IdxObjCContainerDeclInfo{*o}
	}

	return gop_o
}

func (ioccdi IdxObjCCategoryDeclInfo) ObjcClass() *IdxEntityInfo {
	o := ioccdi.c.objcClass

	var gop_o *IdxEntityInfo
	if o != nil {
		gop_o = &IdxEntityInfo{o}
	}

	return gop_o
}

func (ioccdi IdxObjCCategoryDeclInfo) ClassCursor() Cursor {
	return Cursor{ioccdi.c.classCursor}
}

func (ioccdi IdxObjCCategoryDeclInfo) ClassLoc() IdxLoc {
	return IdxLoc{ioccdi.c.classLoc}
}

func (ioccdi IdxObjCCategoryDeclInfo) Protocols() *IdxObjCProtocolRefListInfo {
	o := ioccdi.c.protocols

	var gop_o *IdxObjCProtocolRefListInfo
	if o != nil {
		gop_o = &IdxObjCProtocolRefListInfo{*o}
	}

	return gop_o
}
