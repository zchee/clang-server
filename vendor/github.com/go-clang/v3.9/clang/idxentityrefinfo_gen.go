package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Data for IndexerCallbacks#indexEntityReference.
type IdxEntityRefInfo struct {
	c C.CXIdxEntityRefInfo
}

func (ieri IdxEntityRefInfo) Kind() IdxEntityRefKind {
	return IdxEntityRefKind(ieri.c.kind)
}

// Reference cursor.
func (ieri IdxEntityRefInfo) Cursor() Cursor {
	return Cursor{ieri.c.cursor}
}

func (ieri IdxEntityRefInfo) Loc() IdxLoc {
	return IdxLoc{ieri.c.loc}
}

// The entity that gets referenced.
func (ieri IdxEntityRefInfo) ReferencedEntity() *IdxEntityInfo {
	o := ieri.c.referencedEntity

	var gop_o *IdxEntityInfo
	if o != nil {
		gop_o = &IdxEntityInfo{o}
	}

	return gop_o
}

/*
	Immediate "parent" of the reference. For example:

	\code
	Foo *var;
	\endcode

	The parent of reference of type 'Foo' is the variable 'var'.
	For references inside statement bodies of functions/methods,
	the parentEntity will be the function/method.
*/
func (ieri IdxEntityRefInfo) ParentEntity() *IdxEntityInfo {
	o := ieri.c.parentEntity

	var gop_o *IdxEntityInfo
	if o != nil {
		gop_o = &IdxEntityInfo{o}
	}

	return gop_o
}

// Lexical container context of the reference.
func (ieri IdxEntityRefInfo) Container() *IdxContainerInfo {
	o := ieri.c.container

	var gop_o *IdxContainerInfo
	if o != nil {
		gop_o = &IdxContainerInfo{o}
	}

	return gop_o
}
