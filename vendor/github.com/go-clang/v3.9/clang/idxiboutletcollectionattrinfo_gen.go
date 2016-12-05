package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type IdxIBOutletCollectionAttrInfo struct {
	c C.CXIdxIBOutletCollectionAttrInfo
}

func (iibocai IdxIBOutletCollectionAttrInfo) AttrInfo() *IdxAttrInfo {
	o := iibocai.c.attrInfo

	var gop_o *IdxAttrInfo
	if o != nil {
		gop_o = &IdxAttrInfo{o}
	}

	return gop_o
}

func (iibocai IdxIBOutletCollectionAttrInfo) ObjcClass() *IdxEntityInfo {
	o := iibocai.c.objcClass

	var gop_o *IdxEntityInfo
	if o != nil {
		gop_o = &IdxEntityInfo{o}
	}

	return gop_o
}

func (iibocai IdxIBOutletCollectionAttrInfo) ClassCursor() Cursor {
	return Cursor{iibocai.c.classCursor}
}

func (iibocai IdxIBOutletCollectionAttrInfo) ClassLoc() IdxLoc {
	return IdxLoc{iibocai.c.classLoc}
}
