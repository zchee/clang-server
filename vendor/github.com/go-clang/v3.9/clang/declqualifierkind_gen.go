package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// 'Qualifiers' written next to the return and parameter types in Objective-C method declarations.
type DeclQualifierKind uint32

const (
	DeclQualifier_None   DeclQualifierKind = C.CXObjCDeclQualifier_None
	DeclQualifier_In                       = C.CXObjCDeclQualifier_In
	DeclQualifier_Inout                    = C.CXObjCDeclQualifier_Inout
	DeclQualifier_Out                      = C.CXObjCDeclQualifier_Out
	DeclQualifier_Bycopy                   = C.CXObjCDeclQualifier_Bycopy
	DeclQualifier_Byref                    = C.CXObjCDeclQualifier_Byref
	DeclQualifier_Oneway                   = C.CXObjCDeclQualifier_Oneway
)

func (dqk DeclQualifierKind) Spelling() string {
	switch dqk {
	case DeclQualifier_None:
		return "DeclQualifier=None"
	case DeclQualifier_In:
		return "DeclQualifier=In"
	case DeclQualifier_Inout:
		return "DeclQualifier=Inout"
	case DeclQualifier_Out:
		return "DeclQualifier=Out"
	case DeclQualifier_Bycopy:
		return "DeclQualifier=Bycopy"
	case DeclQualifier_Byref:
		return "DeclQualifier=Byref"
	case DeclQualifier_Oneway:
		return "DeclQualifier=Oneway"
	}

	return fmt.Sprintf("DeclQualifierKind unkown %d", int(dqk))
}

func (dqk DeclQualifierKind) String() string {
	return dqk.Spelling()
}
