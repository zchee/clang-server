package clang

// #include "./clang-c/Documentation.h"
// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

/*
	A cursor representing some element in the abstract syntax tree for
	a translation unit.

	The cursor abstraction unifies the different kinds of entities in a
	program--declaration, statements, expressions, references to declarations,
	etc.--under a single "cursor" abstraction with a common set of operations.
	Common operation for a cursor include: getting the physical location in
	a source file where the cursor points, getting the name associated with a
	cursor, and retrieving cursors for any child nodes of a particular cursor.

	Cursors can be produced in two specific ways.
	clang_getTranslationUnitCursor() produces a cursor for a translation unit,
	from which one can use clang_visitChildren() to explore the rest of the
	translation unit. clang_getCursor() maps from a physical source location
	to the entity that resides at that location, allowing one to map from the
	source code into the AST.
*/
type Cursor struct {
	c C.CXCursor
}

// Given a cursor that represents a documentable entity (e.g., declaration), return the associated parsed comment as a CXComment_FullComment AST node.
func (c Cursor) ParsedComment() Comment {
	return Comment{C.clang_Cursor_getParsedComment(c.c)}
}

// Retrieve the NULL cursor, which represents no entity.
func NewNullCursor() Cursor {
	return Cursor{C.clang_getNullCursor()}
}

// Determine whether two cursors are equivalent.
func (c Cursor) Equal(c2 Cursor) bool {
	o := C.clang_equalCursors(c.c, c2.c)

	return o != C.uint(0)
}

// Returns non-zero if \p cursor is null.
func (c Cursor) IsNull() bool {
	o := C.clang_Cursor_isNull(c.c)

	return o != C.int(0)
}

// Compute a hash value for the given cursor.
func (c Cursor) HashCursor() uint32 {
	return uint32(C.clang_hashCursor(c.c))
}

// Retrieve the kind of the given cursor.
func (c Cursor) Kind() CursorKind {
	return CursorKind(C.clang_getCursorKind(c.c))
}

// Determine whether the given cursor has any attributes.
func (c Cursor) HasAttrs() bool {
	o := C.clang_Cursor_hasAttrs(c.c)

	return o != C.uint(0)
}

// Determine the linkage of the entity referred to by a given cursor.
func (c Cursor) Linkage() LinkageKind {
	return LinkageKind(C.clang_getCursorLinkage(c.c))
}

/*
	Describe the visibility of the entity referred to by a cursor.

	This returns the default visibility if not explicitly specified by
	a visibility attribute. The default visibility may be changed by
	commandline arguments.

	Parameter cursor The cursor to query.

	Returns The visibility of the cursor.
*/
func (c Cursor) Visibility() VisibilityKind {
	return VisibilityKind(C.clang_getCursorVisibility(c.c))
}

/*
	Determine the availability of the entity that this cursor refers to,
	taking the current target platform into account.

	Parameter cursor The cursor to query.

	Returns The availability of the cursor.
*/
func (c Cursor) Availability() AvailabilityKind {
	return AvailabilityKind(C.clang_getCursorAvailability(c.c))
}

// Determine the "language" of the entity referred to by a given cursor.
func (c Cursor) Language() LanguageKind {
	return LanguageKind(C.clang_getCursorLanguage(c.c))
}

// Returns the translation unit that a cursor originated from.
func (c Cursor) TranslationUnit() TranslationUnit {
	return TranslationUnit{C.clang_Cursor_getTranslationUnit(c.c)}
}

/*
	Determine the semantic parent of the given cursor.

	The semantic parent of a cursor is the cursor that semantically contains
	the given \p cursor. For many declarations, the lexical and semantic parents
	are equivalent (the lexical parent is returned by
	clang_getCursorLexicalParent()). They diverge when declarations or
	definitions are provided out-of-line. For example:

	\code
	class C {
	void f();
	};

	void C::f() { }
	\endcode

	In the out-of-line definition of C::f, the semantic parent is
	the class C, of which this function is a member. The lexical parent is
	the place where the declaration actually occurs in the source code; in this
	case, the definition occurs in the translation unit. In general, the
	lexical parent for a given entity can change without affecting the semantics
	of the program, and the lexical parent of different declarations of the
	same entity may be different. Changing the semantic parent of a declaration,
	on the other hand, can have a major impact on semantics, and redeclarations
	of a particular entity should all have the same semantic context.

	In the example above, both declarations of C::f have C as their
	semantic context, while the lexical context of the first C::f is C
	and the lexical context of the second C::f is the translation unit.

	For global declarations, the semantic parent is the translation unit.
*/
func (c Cursor) SemanticParent() Cursor {
	return Cursor{C.clang_getCursorSemanticParent(c.c)}
}

/*
	Determine the lexical parent of the given cursor.

	The lexical parent of a cursor is the cursor in which the given \p cursor
	was actually written. For many declarations, the lexical and semantic parents
	are equivalent (the semantic parent is returned by
	clang_getCursorSemanticParent()). They diverge when declarations or
	definitions are provided out-of-line. For example:

	\code
	class C {
	void f();
	};

	void C::f() { }
	\endcode

	In the out-of-line definition of C::f, the semantic parent is
	the class C, of which this function is a member. The lexical parent is
	the place where the declaration actually occurs in the source code; in this
	case, the definition occurs in the translation unit. In general, the
	lexical parent for a given entity can change without affecting the semantics
	of the program, and the lexical parent of different declarations of the
	same entity may be different. Changing the semantic parent of a declaration,
	on the other hand, can have a major impact on semantics, and redeclarations
	of a particular entity should all have the same semantic context.

	In the example above, both declarations of C::f have C as their
	semantic context, while the lexical context of the first C::f is C
	and the lexical context of the second C::f is the translation unit.

	For declarations written in the global scope, the lexical parent is
	the translation unit.
*/
func (c Cursor) LexicalParent() Cursor {
	return Cursor{C.clang_getCursorLexicalParent(c.c)}
}

/*
	Determine the set of methods that are overridden by the given
	method.

	In both Objective-C and C++, a method (aka virtual member function,
	in C++) can override a virtual method in a base class. For
	Objective-C, a method is said to override any method in the class's
	base class, its protocols, or its categories' protocols, that has the same
	selector and is of the same kind (class or instance).
	If no such method exists, the search continues to the class's superclass,
	its protocols, and its categories, and so on. A method from an Objective-C
	implementation is considered to override the same methods as its
	corresponding method in the interface.

	For C++, a virtual member function overrides any virtual member
	function with the same signature that occurs in its base
	classes. With multiple inheritance, a virtual member function can
	override several virtual member functions coming from different
	base classes.

	In all cases, this function determines the immediate overridden
	method, rather than all of the overridden methods. For example, if
	a method is originally declared in a class A, then overridden in B
	(which in inherits from A) and also in C (which inherited from B),
	then the only overridden method returned from this function when
	invoked on C's method will be B's method. The client may then
	invoke this function again, given the previously-found overridden
	methods, to map out the complete method-override set.

	Parameter cursor A cursor representing an Objective-C or C++
	method. This routine will compute the set of methods that this
	method overrides.

	Parameter overridden A pointer whose pointee will be replaced with a
	pointer to an array of cursors, representing the set of overridden
	methods. If there are no overridden methods, the pointee will be
	set to NULL. The pointee must be freed via a call to
	clang_disposeOverriddenCursors().

	Parameter num_overridden A pointer to the number of overridden
	functions, will be set to the number of overridden functions in the
	array pointed to by \p overridden.
*/
func (c Cursor) OverriddenCursors() []Cursor {
	var cp_overridden *C.CXCursor
	var overridden []Cursor
	var numOverridden C.uint

	C.clang_getOverriddenCursors(c.c, &cp_overridden, &numOverridden)

	gos_overridden := (*reflect.SliceHeader)(unsafe.Pointer(&overridden))
	gos_overridden.Cap = int(numOverridden)
	gos_overridden.Len = int(numOverridden)
	gos_overridden.Data = uintptr(unsafe.Pointer(cp_overridden))

	return overridden
}

// Free the set of overridden cursors returned by \c clang_getOverriddenCursors().
func Dispose(overridden []Cursor) {
	gos_overridden := (*reflect.SliceHeader)(unsafe.Pointer(&overridden))
	cp_overridden := (*C.CXCursor)(unsafe.Pointer(gos_overridden.Data))

	C.clang_disposeOverriddenCursors(cp_overridden)
}

// Retrieve the file that is included by the given inclusion directive cursor.
func (c Cursor) IncludedFile() File {
	return File{C.clang_getIncludedFile(c.c)}
}

/*
	Retrieve the physical location of the source constructor referenced
	by the given cursor.

	The location of a declaration is typically the location of the name of that
	declaration, where the name of that declaration would occur if it is
	unnamed, or some keyword that introduces that particular declaration.
	The location of a reference is where that reference occurs within the
	source code.
*/
func (c Cursor) Location() SourceLocation {
	return SourceLocation{C.clang_getCursorLocation(c.c)}
}

/*
	Retrieve the physical extent of the source construct referenced by
	the given cursor.

	The extent of a cursor starts with the file/line/column pointing at the
	first character within the source construct that the cursor refers to and
	ends with the last character within that source construct. For a
	declaration, the extent covers the declaration itself. For a reference,
	the extent covers the location of the reference (e.g., where the referenced
	entity was actually used).
*/
func (c Cursor) Extent() SourceRange {
	return SourceRange{C.clang_getCursorExtent(c.c)}
}

// Retrieve the type of a CXCursor (if any).
func (c Cursor) Type() Type {
	return Type{C.clang_getCursorType(c.c)}
}

/*
	Retrieve the underlying type of a typedef declaration.

	If the cursor does not reference a typedef declaration, an invalid type is
	returned.
*/
func (c Cursor) TypedefDeclUnderlyingType() Type {
	return Type{C.clang_getTypedefDeclUnderlyingType(c.c)}
}

/*
	Retrieve the integer type of an enum declaration.

	If the cursor does not reference an enum declaration, an invalid type is
	returned.
*/
func (c Cursor) EnumDeclIntegerType() Type {
	return Type{C.clang_getEnumDeclIntegerType(c.c)}
}

/*
	Retrieve the integer value of an enum constant declaration as a signed
	long long.

	If the cursor does not reference an enum constant declaration, LLONG_MIN is returned.
	Since this is also potentially a valid constant value, the kind of the cursor
	must be verified before calling this function.
*/
func (c Cursor) EnumConstantDeclValue() int64 {
	return int64(C.clang_getEnumConstantDeclValue(c.c))
}

/*
	Retrieve the integer value of an enum constant declaration as an unsigned
	long long.

	If the cursor does not reference an enum constant declaration, ULLONG_MAX is returned.
	Since this is also potentially a valid constant value, the kind of the cursor
	must be verified before calling this function.
*/
func (c Cursor) EnumConstantDeclUnsignedValue() uint64 {
	return uint64(C.clang_getEnumConstantDeclUnsignedValue(c.c))
}

/*
	Retrieve the bit width of a bit field declaration as an integer.

	If a cursor that is not a bit field declaration is passed in, -1 is returned.
*/
func (c Cursor) FieldDeclBitWidth() int32 {
	return int32(C.clang_getFieldDeclBitWidth(c.c))
}

/*
	Retrieve the number of non-variadic arguments associated with a given
	cursor.

	The number of arguments can be determined for calls as well as for
	declarations of functions or methods. For other cursors -1 is returned.
*/
func (c Cursor) NumArguments() int32 {
	return int32(C.clang_Cursor_getNumArguments(c.c))
}

/*
	Retrieve the argument cursor of a function or method.

	The argument cursor can be determined for calls as well as for declarations
	of functions or methods. For other cursors and for invalid indices, an
	invalid cursor is returned.
*/
func (c Cursor) Argument(i uint32) Cursor {
	return Cursor{C.clang_Cursor_getArgument(c.c, C.uint(i))}
}

/*
	Returns the number of template args of a function decl representing a
	template specialization.

	If the argument cursor cannot be converted into a template function
	declaration, -1 is returned.

	For example, for the following declaration and specialization:
	template <typename T, int kInt, bool kBool>
	void foo() { ... }

	template <>
	void foo<float, -7, true>();

	The value 3 would be returned from this call.
*/
func (c Cursor) NumTemplateArguments() int32 {
	return int32(C.clang_Cursor_getNumTemplateArguments(c.c))
}

/*
	Retrieve the kind of the I'th template argument of the CXCursor C.

	If the argument CXCursor does not represent a FunctionDecl, an invalid
	template argument kind is returned.

	For example, for the following declaration and specialization:
	template <typename T, int kInt, bool kBool>
	void foo() { ... }

	template <>
	void foo<float, -7, true>();

	For I = 0, 1, and 2, Type, Integral, and Integral will be returned,
	respectively.
*/
func (c Cursor) TemplateArgumentKind(i uint32) TemplateArgumentKind {
	return TemplateArgumentKind(C.clang_Cursor_getTemplateArgumentKind(c.c, C.uint(i)))
}

/*
	Retrieve a CXType representing the type of a TemplateArgument of a
	function decl representing a template specialization.

	If the argument CXCursor does not represent a FunctionDecl whose I'th
	template argument has a kind of CXTemplateArgKind_Integral, an invalid type
	is returned.

	For example, for the following declaration and specialization:
	template <typename T, int kInt, bool kBool>
	void foo() { ... }

	template <>
	void foo<float, -7, true>();

	If called with I = 0, "float", will be returned.
	Invalid types will be returned for I == 1 or 2.
*/
func (c Cursor) TemplateArgumentType(i uint32) Type {
	return Type{C.clang_Cursor_getTemplateArgumentType(c.c, C.uint(i))}
}

/*
	Retrieve the value of an Integral TemplateArgument (of a function
	decl representing a template specialization) as a signed long long.

	It is undefined to call this function on a CXCursor that does not represent a
	FunctionDecl or whose I'th template argument is not an integral value.

	For example, for the following declaration and specialization:
	template <typename T, int kInt, bool kBool>
	void foo() { ... }

	template <>
	void foo<float, -7, true>();

	If called with I = 1 or 2, -7 or true will be returned, respectively.
	For I == 0, this function's behavior is undefined.
*/
func (c Cursor) TemplateArgumentValue(i uint32) int64 {
	return int64(C.clang_Cursor_getTemplateArgumentValue(c.c, C.uint(i)))
}

/*
	Retrieve the value of an Integral TemplateArgument (of a function
	decl representing a template specialization) as an unsigned long long.

	It is undefined to call this function on a CXCursor that does not represent a
	FunctionDecl or whose I'th template argument is not an integral value.

	For example, for the following declaration and specialization:
	template <typename T, int kInt, bool kBool>
	void foo() { ... }

	template <>
	void foo<float, 2147483649, true>();

	If called with I = 1 or 2, 2147483649 or true will be returned, respectively.
	For I == 0, this function's behavior is undefined.
*/
func (c Cursor) TemplateArgumentUnsignedValue(i uint32) uint64 {
	return uint64(C.clang_Cursor_getTemplateArgumentUnsignedValue(c.c, C.uint(i)))
}

// Determine whether a CXCursor that is a macro, is function like.
func (c Cursor) IsMacroFunctionLike() bool {
	o := C.clang_Cursor_isMacroFunctionLike(c.c)

	return o != C.uint(0)
}

// Determine whether a CXCursor that is a macro, is a builtin one.
func (c Cursor) IsMacroBuiltin() bool {
	o := C.clang_Cursor_isMacroBuiltin(c.c)

	return o != C.uint(0)
}

// Determine whether a CXCursor that is a function declaration, is an inline declaration.
func (c Cursor) IsFunctionInlined() bool {
	o := C.clang_Cursor_isFunctionInlined(c.c)

	return o != C.uint(0)
}

// Returns the Objective-C type encoding for the specified declaration.
func (c Cursor) DeclObjCTypeEncoding() string {
	o := cxstring{C.clang_getDeclObjCTypeEncoding(c.c)}
	defer o.Dispose()

	return o.String()
}

/*
	Retrieve the return type associated with a given cursor.

	This only returns a valid type if the cursor refers to a function or method.
*/
func (c Cursor) ResultType() Type {
	return Type{C.clang_getCursorResultType(c.c)}
}

/*
	Return the offset of the field represented by the Cursor.

	If the cursor is not a field declaration, -1 is returned.
	If the cursor semantic parent is not a record field declaration,
	CXTypeLayoutError_Invalid is returned.
	If the field's type declaration is an incomplete type,
	CXTypeLayoutError_Incomplete is returned.
	If the field's type declaration is a dependent type,
	CXTypeLayoutError_Dependent is returned.
	If the field's name S is not found,
	CXTypeLayoutError_InvalidFieldName is returned.
*/
func (c Cursor) OffsetOfField() int64 {
	return int64(C.clang_Cursor_getOffsetOfField(c.c))
}

// Determine whether the given cursor represents an anonymous record declaration.
func (c Cursor) IsAnonymous() bool {
	o := C.clang_Cursor_isAnonymous(c.c)

	return o != C.uint(0)
}

// Returns non-zero if the cursor specifies a Record member that is a bitfield.
func (c Cursor) IsBitField() bool {
	o := C.clang_Cursor_isBitField(c.c)

	return o != C.uint(0)
}

// Returns 1 if the base class specified by the cursor with kind CX_CXXBaseSpecifier is virtual.
func (c Cursor) IsVirtualBase() bool {
	o := C.clang_isVirtualBase(c.c)

	return o != C.uint(0)
}

/*
	Returns the access control level for the referenced object.

	If the cursor refers to a C++ declaration, its access control level within its
	parent scope is returned. Otherwise, if the cursor refers to a base specifier or
	access specifier, the specifier itself is returned.
*/
func (c Cursor) AccessSpecifier() AccessSpecifier {
	return AccessSpecifier(C.clang_getCXXAccessSpecifier(c.c))
}

/*
	Returns the storage class for a function or variable declaration.

	If the passed in Cursor is not a function or variable declaration,
	CX_SC_Invalid is returned else the storage class.
*/
func (c Cursor) StorageClass() StorageClass {
	return StorageClass(C.clang_Cursor_getStorageClass(c.c))
}

/*
	Determine the number of overloaded declarations referenced by a
	CXCursor_OverloadedDeclRef cursor.

	Parameter cursor The cursor whose overloaded declarations are being queried.

	Returns The number of overloaded declarations referenced by cursor. If it
	is not a CXCursor_OverloadedDeclRef cursor, returns 0.
*/
func (c Cursor) NumOverloadedDecls() uint32 {
	return uint32(C.clang_getNumOverloadedDecls(c.c))
}

/*
	Retrieve a cursor for one of the overloaded declarations referenced
	by a CXCursor_OverloadedDeclRef cursor.

	Parameter cursor The cursor whose overloaded declarations are being queried.

	Parameter index The zero-based index into the set of overloaded declarations in
	the cursor.

	Returns A cursor representing the declaration referenced by the given
	cursor at the specified index. If the cursor does not have an
	associated set of overloaded declarations, or if the index is out of bounds,
	returns clang_getNullCursor();
*/
func (c Cursor) OverloadedDecl(index uint32) Cursor {
	return Cursor{C.clang_getOverloadedDecl(c.c, C.uint(index))}
}

/*
	For cursors representing an iboutletcollection attribute,
	this function returns the collection element type.
*/
func (c Cursor) IBOutletCollectionType() Type {
	return Type{C.clang_getIBOutletCollectionType(c.c)}
}

/*
	Retrieve a Unified Symbol Resolution (USR) for the entity referenced
	by the given cursor.

	A Unified Symbol Resolution (USR) is a string that identifies a particular
	entity (function, class, variable, etc.) within a program. USRs can be
	compared across translation units to determine, e.g., when references in
	one translation refer to an entity defined in another translation unit.
*/
func (c Cursor) USR() string {
	o := cxstring{C.clang_getCursorUSR(c.c)}
	defer o.Dispose()

	return o.String()
}

// Retrieve a name for the entity referenced by this cursor.
func (c Cursor) Spelling() string {
	o := cxstring{C.clang_getCursorSpelling(c.c)}
	defer o.Dispose()

	return o.String()
}

/*
	Retrieve a range for a piece that forms the cursors spelling name.
	Most of the times there is only one range for the complete spelling but for
	Objective-C methods and Objective-C message expressions, there are multiple
	pieces for each selector identifier.

	Parameter pieceIndex the index of the spelling name piece. If this is greater
	than the actual number of pieces, it will return a NULL (invalid) range.

	Parameter options Reserved.
*/
func (c Cursor) SpellingNameRange(pieceIndex uint32, options uint32) SourceRange {
	return SourceRange{C.clang_Cursor_getSpellingNameRange(c.c, C.uint(pieceIndex), C.uint(options))}
}

/*
	Retrieve the display name for the entity referenced by this cursor.

	The display name contains extra information that helps identify the cursor,
	such as the parameters of a function or template or the arguments of a
	class template specialization.
*/
func (c Cursor) DisplayName() string {
	o := cxstring{C.clang_getCursorDisplayName(c.c)}
	defer o.Dispose()

	return o.String()
}

/*
	For a cursor that is a reference, retrieve a cursor representing the
	entity that it references.

	Reference cursors refer to other entities in the AST. For example, an
	Objective-C superclass reference cursor refers to an Objective-C class.
	This function produces the cursor for the Objective-C class from the
	cursor for the superclass reference. If the input cursor is a declaration or
	definition, it returns that declaration or definition unchanged.
	Otherwise, returns the NULL cursor.
*/
func (c Cursor) Referenced() Cursor {
	return Cursor{C.clang_getCursorReferenced(c.c)}
}

/*
	For a cursor that is either a reference to or a declaration
	of some entity, retrieve a cursor that describes the definition of
	that entity.

	Some entities can be declared multiple times within a translation
	unit, but only one of those declarations can also be a
	definition. For example, given:

	\code
	int f(int, int);
	int g(int x, int y) { return f(x, y); }
	int f(int a, int b) { return a + b; }
	int f(int, int);
	\endcode

	there are three declarations of the function "f", but only the
	second one is a definition. The clang_getCursorDefinition()
	function will take any cursor pointing to a declaration of "f"
	(the first or fourth lines of the example) or a cursor referenced
	that uses "f" (the call to "f' inside "g") and will return a
	declaration cursor pointing to the definition (the second "f"
	declaration).

	If given a cursor for which there is no corresponding definition,
	e.g., because there is no definition of that entity within this
	translation unit, returns a NULL cursor.
*/
func (c Cursor) Definition() Cursor {
	return Cursor{C.clang_getCursorDefinition(c.c)}
}

// Determine whether the declaration pointed to by this cursor is also a definition of that entity.
func (c Cursor) IsCursorDefinition() bool {
	o := C.clang_isCursorDefinition(c.c)

	return o != C.uint(0)
}

/*
	Retrieve the canonical cursor corresponding to the given cursor.

	In the C family of languages, many kinds of entities can be declared several
	times within a single translation unit. For example, a structure type can
	be forward-declared (possibly multiple times) and later defined:

	\code
	struct X;
	struct X;
	struct X {
	int member;
	};
	\endcode

	The declarations and the definition of X are represented by three
	different cursors, all of which are declarations of the same underlying
	entity. One of these cursor is considered the "canonical" cursor, which
	is effectively the representative for the underlying entity. One can
	determine if two cursors are declarations of the same underlying entity by
	comparing their canonical cursors.

	Returns The canonical cursor for the entity referred to by the given cursor.
*/
func (c Cursor) CanonicalCursor() Cursor {
	return Cursor{C.clang_getCanonicalCursor(c.c)}
}

/*
	If the cursor points to a selector identifier in an Objective-C
	method or message expression, this returns the selector index.

	After getting a cursor with #clang_getCursor, this can be called to
	determine if the location points to a selector identifier.

	Returns The selector index if the cursor is an Objective-C method or message
	expression and the cursor is pointing to a selector identifier, or -1
	otherwise.
*/
func (c Cursor) SelectorIndex() int32 {
	return int32(C.clang_Cursor_getObjCSelectorIndex(c.c))
}

/*
	Given a cursor pointing to a C++ method call or an Objective-C
	message, returns non-zero if the method/message is "dynamic", meaning:

	For a C++ method: the call is virtual.
	For an Objective-C message: the receiver is an object instance, not 'super'
	or a specific class.

	If the method/message is "static" or the cursor does not point to a
	method/message, it will return zero.
*/
func (c Cursor) IsDynamicCall() bool {
	o := C.clang_Cursor_isDynamicCall(c.c)

	return o != C.int(0)
}

// Given a cursor pointing to an Objective-C message, returns the CXType of the receiver.
func (c Cursor) ReceiverType() Type {
	return Type{C.clang_Cursor_getReceiverType(c.c)}
}

/*
	Given a cursor that represents a property declaration, return the
	associated property attributes. The bits are formed from
	CXObjCPropertyAttrKind.

	Parameter reserved Reserved for future use, pass 0.
*/
func (c Cursor) PropertyAttributes(reserved uint32) uint32 {
	return uint32(C.clang_Cursor_getObjCPropertyAttributes(c.c, C.uint(reserved)))
}

// Given a cursor that represents an Objective-C method or parameter declaration, return the associated Objective-C qualifiers for the return type or the parameter respectively. The bits are formed from CXObjCDeclQualifierKind.
func (c Cursor) DeclQualifiers() uint32 {
	return uint32(C.clang_Cursor_getObjCDeclQualifiers(c.c))
}

// Given a cursor that represents an Objective-C method or property declaration, return non-zero if the declaration was affected by "@optional". Returns zero if the cursor is not such a declaration or it is "@required".
func (c Cursor) IsObjCOptional() bool {
	o := C.clang_Cursor_isObjCOptional(c.c)

	return o != C.uint(0)
}

// Returns non-zero if the given cursor is a variadic function or method.
func (c Cursor) IsVariadic() bool {
	o := C.clang_Cursor_isVariadic(c.c)

	return o != C.uint(0)
}

// Given a cursor that represents a declaration, return the associated comment's source range. The range may include multiple consecutive comments with whitespace in between.
func (c Cursor) CommentRange() SourceRange {
	return SourceRange{C.clang_Cursor_getCommentRange(c.c)}
}

// Given a cursor that represents a declaration, return the associated comment text, including comment markers.
func (c Cursor) RawCommentText() string {
	o := cxstring{C.clang_Cursor_getRawCommentText(c.c)}
	defer o.Dispose()

	return o.String()
}

// Given a cursor that represents a documentable entity (e.g., declaration), return the associated \paragraph; otherwise return the first paragraph.
func (c Cursor) BriefCommentText() string {
	o := cxstring{C.clang_Cursor_getBriefCommentText(c.c)}
	defer o.Dispose()

	return o.String()
}

// Retrieve the CXString representing the mangled name of the cursor.
func (c Cursor) Mangling() string {
	o := cxstring{C.clang_Cursor_getMangling(c.c)}
	defer o.Dispose()

	return o.String()
}

// Retrieve the CXStrings representing the mangled symbols of the C++ constructor or destructor at the cursor.
func (c Cursor) Manglings() *StringSet {
	o := C.clang_Cursor_getCXXManglings(c.c)

	var gop_o *StringSet
	if o != nil {
		gop_o = &StringSet{*o}
	}

	return gop_o
}

// Given a CXCursor_ModuleImportDecl cursor, return the associated module.
func (c Cursor) Module() Module {
	return Module{C.clang_Cursor_getModule(c.c)}
}

// Determine if a C++ constructor is a converting constructor.
func (c Cursor) CXXConstructor_IsConvertingConstructor() bool {
	o := C.clang_CXXConstructor_isConvertingConstructor(c.c)

	return o != C.uint(0)
}

// Determine if a C++ constructor is a copy constructor.
func (c Cursor) CXXConstructor_IsCopyConstructor() bool {
	o := C.clang_CXXConstructor_isCopyConstructor(c.c)

	return o != C.uint(0)
}

// Determine if a C++ constructor is the default constructor.
func (c Cursor) CXXConstructor_IsDefaultConstructor() bool {
	o := C.clang_CXXConstructor_isDefaultConstructor(c.c)

	return o != C.uint(0)
}

// Determine if a C++ constructor is a move constructor.
func (c Cursor) CXXConstructor_IsMoveConstructor() bool {
	o := C.clang_CXXConstructor_isMoveConstructor(c.c)

	return o != C.uint(0)
}

// Determine if a C++ field is declared 'mutable'.
func (c Cursor) CXXField_IsMutable() bool {
	o := C.clang_CXXField_isMutable(c.c)

	return o != C.uint(0)
}

// Determine if a C++ method is declared '= default'.
func (c Cursor) CXXMethod_IsDefaulted() bool {
	o := C.clang_CXXMethod_isDefaulted(c.c)

	return o != C.uint(0)
}

// Determine if a C++ member function or member function template is pure virtual.
func (c Cursor) CXXMethod_IsPureVirtual() bool {
	o := C.clang_CXXMethod_isPureVirtual(c.c)

	return o != C.uint(0)
}

// Determine if a C++ member function or member function template is declared 'static'.
func (c Cursor) CXXMethod_IsStatic() bool {
	o := C.clang_CXXMethod_isStatic(c.c)

	return o != C.uint(0)
}

// Determine if a C++ member function or member function template is explicitly declared 'virtual' or if it overrides a virtual method from one of the base classes.
func (c Cursor) CXXMethod_IsVirtual() bool {
	o := C.clang_CXXMethod_isVirtual(c.c)

	return o != C.uint(0)
}

// Determine if a C++ member function or member function template is declared 'const'.
func (c Cursor) CXXMethod_IsConst() bool {
	o := C.clang_CXXMethod_isConst(c.c)

	return o != C.uint(0)
}

/*
	Given a cursor that represents a template, determine
	the cursor kind of the specializations would be generated by instantiating
	the template.

	This routine can be used to determine what flavor of function template,
	class template, or class template partial specialization is stored in the
	cursor. For example, it can describe whether a class template cursor is
	declared with "struct", "class" or "union".

	Parameter C The cursor to query. This cursor should represent a template
	declaration.

	Returns The cursor kind of the specializations that would be generated
	by instantiating the template \p C. If \p C is not a template, returns
	CXCursor_NoDeclFound.
*/
func (c Cursor) TemplateCursorKind() CursorKind {
	return CursorKind(C.clang_getTemplateCursorKind(c.c))
}

/*
	Given a cursor that may represent a specialization or instantiation
	of a template, retrieve the cursor that represents the template that it
	specializes or from which it was instantiated.

	This routine determines the template involved both for explicit
	specializations of templates and for implicit instantiations of the template,
	both of which are referred to as "specializations". For a class template
	specialization (e.g., std::vector<bool>), this routine will return
	either the primary template (std::vector) or, if the specialization was
	instantiated from a class template partial specialization, the class template
	partial specialization. For a class template partial specialization and a
	function template specialization (including instantiations), this
	this routine will return the specialized template.

	For members of a class template (e.g., member functions, member classes, or
	static data members), returns the specialized or instantiated member.
	Although not strictly "templates" in the C++ language, members of class
	templates have the same notions of specializations and instantiations that
	templates do, so this routine treats them similarly.

	Parameter C A cursor that may be a specialization of a template or a member
	of a template.

	Returns If the given cursor is a specialization or instantiation of a
	template or a member thereof, the template or member that it specializes or
	from which it was instantiated. Otherwise, returns a NULL cursor.
*/
func (c Cursor) SpecializedCursorTemplate() Cursor {
	return Cursor{C.clang_getSpecializedCursorTemplate(c.c)}
}

/*
	Given a cursor that references something else, return the source range
	covering that reference.

	Parameter C A cursor pointing to a member reference, a declaration reference, or
	an operator call.
	Parameter NameFlags A bitset with three independent flags:
	CXNameRange_WantQualifier, CXNameRange_WantTemplateArgs, and
	CXNameRange_WantSinglePiece.
	Parameter PieceIndex For contiguous names or when passing the flag
	CXNameRange_WantSinglePiece, only one piece with index 0 is
	available. When the CXNameRange_WantSinglePiece flag is not passed for a
	non-contiguous names, this index can be used to retrieve the individual
	pieces of the name. See also CXNameRange_WantSinglePiece.

	Returns The piece of the name pointed to by the given cursor. If there is no
	name, or if the PieceIndex is out-of-range, a null-cursor will be returned.
*/
func (c Cursor) ReferenceNameRange(nameFlags uint32, pieceIndex uint32) SourceRange {
	return SourceRange{C.clang_getCursorReferenceNameRange(c.c, C.uint(nameFlags), C.uint(pieceIndex))}
}

func (c Cursor) DefinitionSpellingAndExtent() (string, string, uint32, uint32, uint32, uint32) {
	var startBuf *C.char
	defer C.free(unsafe.Pointer(startBuf))
	var endBuf *C.char
	defer C.free(unsafe.Pointer(endBuf))
	var startLine C.uint
	var startColumn C.uint
	var endLine C.uint
	var endColumn C.uint

	C.clang_getDefinitionSpellingAndExtent(c.c, &startBuf, &endBuf, &startLine, &startColumn, &endLine, &endColumn)

	return C.GoString(startBuf), C.GoString(endBuf), uint32(startLine), uint32(startColumn), uint32(endLine), uint32(endColumn)
}

/*
	Retrieve a completion string for an arbitrary declaration or macro
	definition cursor.

	Parameter cursor The cursor to query.

	Returns A non-context-sensitive completion string for declaration and macro
	definition cursors, or NULL for other kinds of cursors.
*/
func (c Cursor) CompletionString() CompletionString {
	return CompletionString{C.clang_getCursorCompletionString(c.c)}
}

// If cursor is a statement declaration tries to evaluate the statement and if its variable, tries to evaluate its initializer, into its corresponding type.
func (c Cursor) Evaluate() EvalResult {
	return EvalResult{C.clang_Cursor_Evaluate(c.c)}
}

/*
	Find references of a declaration in a specific file.

	Parameter cursor pointing to a declaration or a reference of one.

	Parameter file to search for references.

	Parameter visitor callback that will receive pairs of CXCursor/CXSourceRange for
	each reference found.
	The CXSourceRange will point inside the file; if the reference is inside
	a macro (and not a macro argument) the CXSourceRange will be invalid.

	Returns one of the CXResult enumerators.
*/
func (c Cursor) FindReferencesInFile(file File, visitor CursorAndRangeVisitor) Result {
	return Result(C.clang_findReferencesInFile(c.c, file.c, visitor.c))
}

func (c Cursor) Xdata() int32 {
	return int32(c.c.xdata)
}
