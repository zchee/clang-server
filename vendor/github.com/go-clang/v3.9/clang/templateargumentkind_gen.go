package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

/*
	Describes the kind of a template argument.

	See the definition of llvm::clang::TemplateArgument::ArgKind for full
	element descriptions.
*/
type TemplateArgumentKind uint32

const (
	TemplateArgumentKind_Null              TemplateArgumentKind = C.CXTemplateArgumentKind_Null
	TemplateArgumentKind_Type                                   = C.CXTemplateArgumentKind_Type
	TemplateArgumentKind_Declaration                            = C.CXTemplateArgumentKind_Declaration
	TemplateArgumentKind_NullPtr                                = C.CXTemplateArgumentKind_NullPtr
	TemplateArgumentKind_Integral                               = C.CXTemplateArgumentKind_Integral
	TemplateArgumentKind_Template                               = C.CXTemplateArgumentKind_Template
	TemplateArgumentKind_TemplateExpansion                      = C.CXTemplateArgumentKind_TemplateExpansion
	TemplateArgumentKind_Expression                             = C.CXTemplateArgumentKind_Expression
	TemplateArgumentKind_Pack                                   = C.CXTemplateArgumentKind_Pack
	TemplateArgumentKind_Invalid                                = C.CXTemplateArgumentKind_Invalid
)

func (tak TemplateArgumentKind) Spelling() string {
	switch tak {
	case TemplateArgumentKind_Null:
		return "TemplateArgumentKind=Null"
	case TemplateArgumentKind_Type:
		return "TemplateArgumentKind=Type"
	case TemplateArgumentKind_Declaration:
		return "TemplateArgumentKind=Declaration"
	case TemplateArgumentKind_NullPtr:
		return "TemplateArgumentKind=NullPtr"
	case TemplateArgumentKind_Integral:
		return "TemplateArgumentKind=Integral"
	case TemplateArgumentKind_Template:
		return "TemplateArgumentKind=Template"
	case TemplateArgumentKind_TemplateExpansion:
		return "TemplateArgumentKind=TemplateExpansion"
	case TemplateArgumentKind_Expression:
		return "TemplateArgumentKind=Expression"
	case TemplateArgumentKind_Pack:
		return "TemplateArgumentKind=Pack"
	case TemplateArgumentKind_Invalid:
		return "TemplateArgumentKind=Invalid"
	}

	return fmt.Sprintf("TemplateArgumentKind unkown %d", int(tak))
}

func (tak TemplateArgumentKind) String() string {
	return tak.Spelling()
}
