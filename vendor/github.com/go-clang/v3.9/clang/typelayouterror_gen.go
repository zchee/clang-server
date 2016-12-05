package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

/*
	List the possible error codes for clang_Type_getSizeOf,
	clang_Type_getAlignOf, clang_Type_getOffsetOf and
	clang_Cursor_getOffsetOf.

	A value of this enumeration type can be returned if the target type is not
	a valid argument to sizeof, alignof or offsetof.
*/
type TypeLayoutError int32

const (
	// Type is of kind CXType_Invalid.
	TypeLayoutError_Invalid TypeLayoutError = C.CXTypeLayoutError_Invalid
	// The type is an incomplete Type.
	TypeLayoutError_Incomplete = C.CXTypeLayoutError_Incomplete
	// The type is a dependent Type.
	TypeLayoutError_Dependent = C.CXTypeLayoutError_Dependent
	// The type is not a constant size type.
	TypeLayoutError_NotConstantSize = C.CXTypeLayoutError_NotConstantSize
	// The Field name is not valid for this record.
	TypeLayoutError_InvalidFieldName = C.CXTypeLayoutError_InvalidFieldName
)

func (tle TypeLayoutError) Spelling() string {
	switch tle {
	case TypeLayoutError_Invalid:
		return "TypeLayoutError=Invalid"
	case TypeLayoutError_Incomplete:
		return "TypeLayoutError=Incomplete"
	case TypeLayoutError_Dependent:
		return "TypeLayoutError=Dependent"
	case TypeLayoutError_NotConstantSize:
		return "TypeLayoutError=NotConstantSize"
	case TypeLayoutError_InvalidFieldName:
		return "TypeLayoutError=InvalidFieldName"
	}

	return fmt.Sprintf("TypeLayoutError unkown %d", int(tle))
}

func (tle TypeLayoutError) String() string {
	return tle.Spelling()
}

func (tle TypeLayoutError) Error() string {
	return tle.Spelling()
}
