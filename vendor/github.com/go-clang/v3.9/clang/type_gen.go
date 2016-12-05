package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "unsafe"

// The type of an element in the abstract syntax tree.
type Type struct {
	c C.CXType
}

/*
	Pretty-print the underlying type using the rules of the
	language of the translation unit from which it came.

	If the type is invalid, an empty string is returned.
*/
func (t Type) Spelling() string {
	o := cxstring{C.clang_getTypeSpelling(t.c)}
	defer o.Dispose()

	return o.String()
}

/*
	Determine whether two CXTypes represent the same type.

	Returns non-zero if the CXTypes represent the same type and
	zero otherwise.
*/
func (t Type) Equal(t2 Type) bool {
	o := C.clang_equalTypes(t.c, t2.c)

	return o != C.uint(0)
}

/*
	Return the canonical type for a CXType.

	Clang's type system explicitly models typedefs and all the ways
	a specific type can be represented. The canonical type is the underlying
	type with all the "sugar" removed. For example, if 'T' is a typedef
	for 'int', the canonical type for 'T' would be 'int'.
*/
func (t Type) CanonicalType() Type {
	return Type{C.clang_getCanonicalType(t.c)}
}

// Determine whether a CXType has the "const" qualifier set, without looking through typedefs that may have added "const" at a different level.
func (t Type) IsConstQualifiedType() bool {
	o := C.clang_isConstQualifiedType(t.c)

	return o != C.uint(0)
}

// Determine whether a CXType has the "volatile" qualifier set, without looking through typedefs that may have added "volatile" at a different level.
func (t Type) IsVolatileQualifiedType() bool {
	o := C.clang_isVolatileQualifiedType(t.c)

	return o != C.uint(0)
}

// Determine whether a CXType has the "restrict" qualifier set, without looking through typedefs that may have added "restrict" at a different level.
func (t Type) IsRestrictQualifiedType() bool {
	o := C.clang_isRestrictQualifiedType(t.c)

	return o != C.uint(0)
}

// For pointer types, returns the type of the pointee.
func (t Type) PointeeType() Type {
	return Type{C.clang_getPointeeType(t.c)}
}

// Return the cursor for the declaration of the given type.
func (t Type) Declaration() Cursor {
	return Cursor{C.clang_getTypeDeclaration(t.c)}
}

// Returns the Objective-C type encoding for the specified CXType.
func (t Type) Encoding() string {
	o := cxstring{C.clang_Type_getObjCEncoding(t.c)}
	defer o.Dispose()

	return o.String()
}

/*
	Retrieve the calling convention associated with a function type.

	If a non-function type is passed in, CXCallingConv_Invalid is returned.
*/
func (t Type) FunctionTypeCallingConv() CallingConv {
	return CallingConv(C.clang_getFunctionTypeCallingConv(t.c))
}

/*
	Retrieve the return type associated with a function type.

	If a non-function type is passed in, an invalid type is returned.
*/
func (t Type) ResultType() Type {
	return Type{C.clang_getResultType(t.c)}
}

/*
	Retrieve the number of non-variadic parameters associated with a
	function type.

	If a non-function type is passed in, -1 is returned.
*/
func (t Type) NumArgTypes() int32 {
	return int32(C.clang_getNumArgTypes(t.c))
}

/*
	Retrieve the type of a parameter of a function type.

	If a non-function type is passed in or the function does not have enough
	parameters, an invalid type is returned.
*/
func (t Type) ArgType(i uint32) Type {
	return Type{C.clang_getArgType(t.c, C.uint(i))}
}

// Return 1 if the CXType is a variadic function type, and 0 otherwise.
func (t Type) IsFunctionTypeVariadic() bool {
	o := C.clang_isFunctionTypeVariadic(t.c)

	return o != C.uint(0)
}

// Return 1 if the CXType is a POD (plain old data) type, and 0 otherwise.
func (t Type) IsPODType() bool {
	o := C.clang_isPODType(t.c)

	return o != C.uint(0)
}

/*
	Return the element type of an array, complex, or vector type.

	If a type is passed in that is not an array, complex, or vector type,
	an invalid type is returned.
*/
func (t Type) ElementType() Type {
	return Type{C.clang_getElementType(t.c)}
}

/*
	Return the number of elements of an array or vector type.

	If a type is passed in that is not an array or vector type,
	-1 is returned.
*/
func (t Type) NumElements() int64 {
	return int64(C.clang_getNumElements(t.c))
}

/*
	Return the element type of an array type.

	If a non-array type is passed in, an invalid type is returned.
*/
func (t Type) ArrayElementType() Type {
	return Type{C.clang_getArrayElementType(t.c)}
}

/*
	Return the array size of a constant array.

	If a non-array type is passed in, -1 is returned.
*/
func (t Type) ArraySize() int64 {
	return int64(C.clang_getArraySize(t.c))
}

/*
	Retrieve the type named by the qualified-id.

	If a non-elaborated type is passed in, an invalid type is returned.
*/
func (t Type) NamedType() Type {
	return Type{C.clang_Type_getNamedType(t.c)}
}

/*
	Return the alignment of a type in bytes as per C++[expr.alignof]
	standard.

	If the type declaration is invalid, CXTypeLayoutError_Invalid is returned.
	If the type declaration is an incomplete type, CXTypeLayoutError_Incomplete
	is returned.
	If the type declaration is a dependent type, CXTypeLayoutError_Dependent is
	returned.
	If the type declaration is not a constant size type,
	CXTypeLayoutError_NotConstantSize is returned.
*/
func (t Type) AlignOf() int64 {
	return int64(C.clang_Type_getAlignOf(t.c))
}

/*
	Return the class type of an member pointer type.

	If a non-member-pointer type is passed in, an invalid type is returned.
*/
func (t Type) ClassType() Type {
	return Type{C.clang_Type_getClassType(t.c)}
}

/*
	Return the size of a type in bytes as per C++[expr.sizeof] standard.

	If the type declaration is invalid, CXTypeLayoutError_Invalid is returned.
	If the type declaration is an incomplete type, CXTypeLayoutError_Incomplete
	is returned.
	If the type declaration is a dependent type, CXTypeLayoutError_Dependent is
	returned.
*/
func (t Type) SizeOf() int64 {
	return int64(C.clang_Type_getSizeOf(t.c))
}

/*
	Return the offset of a field named S in a record of type T in bits
	as it would be returned by __offsetof__ as per C++11[18.2p4]

	If the cursor is not a record field declaration, CXTypeLayoutError_Invalid
	is returned.
	If the field's type declaration is an incomplete type,
	CXTypeLayoutError_Incomplete is returned.
	If the field's type declaration is a dependent type,
	CXTypeLayoutError_Dependent is returned.
	If the field's name S is not found,
	CXTypeLayoutError_InvalidFieldName is returned.
*/
func (t Type) OffsetOf(s string) int64 {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))

	return int64(C.clang_Type_getOffsetOf(t.c, c_s))
}

/*
	Returns the number of template arguments for given class template
	specialization, or -1 if type T is not a class template specialization.

	Variadic argument packs count as only one argument, and can not be inspected
	further.
*/
func (t Type) NumTemplateArguments() int32 {
	return int32(C.clang_Type_getNumTemplateArguments(t.c))
}

/*
	Returns the type template argument of a template class specialization
	at given index.

	This function only returns template type arguments and does not handle
	template template arguments or variadic packs.
*/
func (t Type) TemplateArgumentAsType(i uint32) Type {
	return Type{C.clang_Type_getTemplateArgumentAsType(t.c, C.uint(i))}
}

/*
	Retrieve the ref-qualifier kind of a function or method.

	The ref-qualifier is returned for C++ functions or methods. For other types
	or non-C++ declarations, CXRefQualifier_None is returned.
*/
func (t Type) RefQualifier() RefQualifierKind {
	return RefQualifierKind(C.clang_Type_getCXXRefQualifier(t.c))
}

func (t Type) Kind() TypeKind {
	return TypeKind(t.c.kind)
}
