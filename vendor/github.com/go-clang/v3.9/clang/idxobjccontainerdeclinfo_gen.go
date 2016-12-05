package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type IdxObjCContainerDeclInfo struct {
	c C.CXIdxObjCContainerDeclInfo
}

func (ioccdi IdxObjCContainerDeclInfo) DeclInfo() *IdxDeclInfo {
	o := ioccdi.c.declInfo

	var gop_o *IdxDeclInfo
	if o != nil {
		gop_o = &IdxDeclInfo{o}
	}

	return gop_o
}

func (ioccdi IdxObjCContainerDeclInfo) Kind() IdxObjCContainerKind {
	return IdxObjCContainerKind(ioccdi.c.kind)
}
