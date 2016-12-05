package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type IdxBaseClassInfo struct {
	c C.CXIdxBaseClassInfo
}

func (ibci IdxBaseClassInfo) Base() *IdxEntityInfo {
	o := ibci.c.base

	var gop_o *IdxEntityInfo
	if o != nil {
		gop_o = &IdxEntityInfo{o}
	}

	return gop_o
}

func (ibci IdxBaseClassInfo) Cursor() Cursor {
	return Cursor{ibci.c.cursor}
}

func (ibci IdxBaseClassInfo) Loc() IdxLoc {
	return IdxLoc{ibci.c.loc}
}
