package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type IdxObjCProtocolRefInfo struct {
	c C.CXIdxObjCProtocolRefInfo
}

func (iocpri IdxObjCProtocolRefInfo) Protocol() *IdxEntityInfo {
	o := iocpri.c.protocol

	var gop_o *IdxEntityInfo
	if o != nil {
		gop_o = &IdxEntityInfo{o}
	}

	return gop_o
}

func (iocpri IdxObjCProtocolRefInfo) Cursor() Cursor {
	return Cursor{iocpri.c.cursor}
}

func (iocpri IdxObjCProtocolRefInfo) Loc() IdxLoc {
	return IdxLoc{iocpri.c.loc}
}
