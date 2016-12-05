package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type IdxObjCPropertyDeclInfo struct {
	c C.CXIdxObjCPropertyDeclInfo
}

func (iocpdi IdxObjCPropertyDeclInfo) DeclInfo() *IdxDeclInfo {
	o := iocpdi.c.declInfo

	var gop_o *IdxDeclInfo
	if o != nil {
		gop_o = &IdxDeclInfo{o}
	}

	return gop_o
}

func (iocpdi IdxObjCPropertyDeclInfo) Getter() *IdxEntityInfo {
	o := iocpdi.c.getter

	var gop_o *IdxEntityInfo
	if o != nil {
		gop_o = &IdxEntityInfo{o}
	}

	return gop_o
}

func (iocpdi IdxObjCPropertyDeclInfo) Setter() *IdxEntityInfo {
	o := iocpdi.c.setter

	var gop_o *IdxEntityInfo
	if o != nil {
		gop_o = &IdxEntityInfo{o}
	}

	return gop_o
}
