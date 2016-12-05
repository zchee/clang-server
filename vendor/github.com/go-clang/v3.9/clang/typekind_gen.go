package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Describes the kind of type
type TypeKind uint32

const (
	// Represents an invalid type (e.g., where no type is available).
	Type_Invalid TypeKind = C.CXType_Invalid
	// A type whose specific kind is not exposed via this interface.
	Type_Unexposed = C.CXType_Unexposed
	// A type whose specific kind is not exposed via this interface.
	Type_Void = C.CXType_Void
	// A type whose specific kind is not exposed via this interface.
	Type_Bool = C.CXType_Bool
	// A type whose specific kind is not exposed via this interface.
	Type_Char_U = C.CXType_Char_U
	// A type whose specific kind is not exposed via this interface.
	Type_UChar = C.CXType_UChar
	// A type whose specific kind is not exposed via this interface.
	Type_Char16 = C.CXType_Char16
	// A type whose specific kind is not exposed via this interface.
	Type_Char32 = C.CXType_Char32
	// A type whose specific kind is not exposed via this interface.
	Type_UShort = C.CXType_UShort
	// A type whose specific kind is not exposed via this interface.
	Type_UInt = C.CXType_UInt
	// A type whose specific kind is not exposed via this interface.
	Type_ULong = C.CXType_ULong
	// A type whose specific kind is not exposed via this interface.
	Type_ULongLong = C.CXType_ULongLong
	// A type whose specific kind is not exposed via this interface.
	Type_UInt128 = C.CXType_UInt128
	// A type whose specific kind is not exposed via this interface.
	Type_Char_S = C.CXType_Char_S
	// A type whose specific kind is not exposed via this interface.
	Type_SChar = C.CXType_SChar
	// A type whose specific kind is not exposed via this interface.
	Type_WChar = C.CXType_WChar
	// A type whose specific kind is not exposed via this interface.
	Type_Short = C.CXType_Short
	// A type whose specific kind is not exposed via this interface.
	Type_Int = C.CXType_Int
	// A type whose specific kind is not exposed via this interface.
	Type_Long = C.CXType_Long
	// A type whose specific kind is not exposed via this interface.
	Type_LongLong = C.CXType_LongLong
	// A type whose specific kind is not exposed via this interface.
	Type_Int128 = C.CXType_Int128
	// A type whose specific kind is not exposed via this interface.
	Type_Float = C.CXType_Float
	// A type whose specific kind is not exposed via this interface.
	Type_Double = C.CXType_Double
	// A type whose specific kind is not exposed via this interface.
	Type_LongDouble = C.CXType_LongDouble
	// A type whose specific kind is not exposed via this interface.
	Type_NullPtr = C.CXType_NullPtr
	// A type whose specific kind is not exposed via this interface.
	Type_Overload = C.CXType_Overload
	// A type whose specific kind is not exposed via this interface.
	Type_Dependent = C.CXType_Dependent
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCId = C.CXType_ObjCId
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCClass = C.CXType_ObjCClass
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCSel = C.CXType_ObjCSel
	// A type whose specific kind is not exposed via this interface.
	Type_Float128 = C.CXType_Float128
	// A type whose specific kind is not exposed via this interface.
	Type_FirstBuiltin = C.CXType_FirstBuiltin
	// A type whose specific kind is not exposed via this interface.
	Type_LastBuiltin = C.CXType_LastBuiltin
	// A type whose specific kind is not exposed via this interface.
	Type_Complex = C.CXType_Complex
	// A type whose specific kind is not exposed via this interface.
	Type_Pointer = C.CXType_Pointer
	// A type whose specific kind is not exposed via this interface.
	Type_BlockPointer = C.CXType_BlockPointer
	// A type whose specific kind is not exposed via this interface.
	Type_LValueReference = C.CXType_LValueReference
	// A type whose specific kind is not exposed via this interface.
	Type_RValueReference = C.CXType_RValueReference
	// A type whose specific kind is not exposed via this interface.
	Type_Record = C.CXType_Record
	// A type whose specific kind is not exposed via this interface.
	Type_Enum = C.CXType_Enum
	// A type whose specific kind is not exposed via this interface.
	Type_Typedef = C.CXType_Typedef
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCInterface = C.CXType_ObjCInterface
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCObjectPointer = C.CXType_ObjCObjectPointer
	// A type whose specific kind is not exposed via this interface.
	Type_FunctionNoProto = C.CXType_FunctionNoProto
	// A type whose specific kind is not exposed via this interface.
	Type_FunctionProto = C.CXType_FunctionProto
	// A type whose specific kind is not exposed via this interface.
	Type_ConstantArray = C.CXType_ConstantArray
	// A type whose specific kind is not exposed via this interface.
	Type_Vector = C.CXType_Vector
	// A type whose specific kind is not exposed via this interface.
	Type_IncompleteArray = C.CXType_IncompleteArray
	// A type whose specific kind is not exposed via this interface.
	Type_VariableArray = C.CXType_VariableArray
	// A type whose specific kind is not exposed via this interface.
	Type_DependentSizedArray = C.CXType_DependentSizedArray
	// A type whose specific kind is not exposed via this interface.
	Type_MemberPointer = C.CXType_MemberPointer
	// A type whose specific kind is not exposed via this interface.
	Type_Auto = C.CXType_Auto
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_Elaborated = C.CXType_Elaborated
)

// Retrieve the spelling of a given CXTypeKind.
func (tk TypeKind) Spelling() string {
	o := cxstring{C.clang_getTypeKindSpelling(C.enum_CXTypeKind(tk))}
	defer o.Dispose()

	return o.String()
}

func (tk TypeKind) String() string {
	return tk.Spelling()
}
