package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type RefQualifierKind uint32

const (
	// No ref-qualifier was provided.
	RefQualifier_None RefQualifierKind = C.CXRefQualifier_None
	// An lvalue ref-qualifier was provided (&).
	RefQualifier_LValue = C.CXRefQualifier_LValue
	// An rvalue ref-qualifier was provided (&&).
	RefQualifier_RValue = C.CXRefQualifier_RValue
)

func (rqk RefQualifierKind) Spelling() string {
	switch rqk {
	case RefQualifier_None:
		return "RefQualifier=None"
	case RefQualifier_LValue:
		return "RefQualifier=LValue"
	case RefQualifier_RValue:
		return "RefQualifier=RValue"
	}

	return fmt.Sprintf("RefQualifierKind unkown %d", int(rqk))
}

func (rqk RefQualifierKind) String() string {
	return rqk.Spelling()
}
