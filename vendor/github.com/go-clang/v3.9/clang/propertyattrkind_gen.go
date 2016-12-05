package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Property attributes for a CXCursor_ObjCPropertyDecl.
type PropertyAttrKind uint32

const (
	PropertyAttr_noattr            PropertyAttrKind = C.CXObjCPropertyAttr_noattr
	PropertyAttr_readonly                           = C.CXObjCPropertyAttr_readonly
	PropertyAttr_getter                             = C.CXObjCPropertyAttr_getter
	PropertyAttr_assign                             = C.CXObjCPropertyAttr_assign
	PropertyAttr_readwrite                          = C.CXObjCPropertyAttr_readwrite
	PropertyAttr_retain                             = C.CXObjCPropertyAttr_retain
	PropertyAttr_copy                               = C.CXObjCPropertyAttr_copy
	PropertyAttr_nonatomic                          = C.CXObjCPropertyAttr_nonatomic
	PropertyAttr_setter                             = C.CXObjCPropertyAttr_setter
	PropertyAttr_atomic                             = C.CXObjCPropertyAttr_atomic
	PropertyAttr_weak                               = C.CXObjCPropertyAttr_weak
	PropertyAttr_strong                             = C.CXObjCPropertyAttr_strong
	PropertyAttr_unsafe_unretained                  = C.CXObjCPropertyAttr_unsafe_unretained
	PropertyAttr_class                              = C.CXObjCPropertyAttr_class
)

func (pak PropertyAttrKind) Spelling() string {
	switch pak {
	case PropertyAttr_noattr:
		return "PropertyAttr=noattr"
	case PropertyAttr_readonly:
		return "PropertyAttr=readonly"
	case PropertyAttr_getter:
		return "PropertyAttr=getter"
	case PropertyAttr_assign:
		return "PropertyAttr=assign"
	case PropertyAttr_readwrite:
		return "PropertyAttr=readwrite"
	case PropertyAttr_retain:
		return "PropertyAttr=retain"
	case PropertyAttr_copy:
		return "PropertyAttr=copy"
	case PropertyAttr_nonatomic:
		return "PropertyAttr=nonatomic"
	case PropertyAttr_setter:
		return "PropertyAttr=setter"
	case PropertyAttr_atomic:
		return "PropertyAttr=atomic"
	case PropertyAttr_weak:
		return "PropertyAttr=weak"
	case PropertyAttr_strong:
		return "PropertyAttr=strong"
	case PropertyAttr_unsafe_unretained:
		return "PropertyAttr=unsafe_unretained"
	case PropertyAttr_class:
		return "PropertyAttr=class"
	}

	return fmt.Sprintf("PropertyAttrKind unkown %d", int(pak))
}

func (pak PropertyAttrKind) String() string {
	return pak.Spelling()
}
