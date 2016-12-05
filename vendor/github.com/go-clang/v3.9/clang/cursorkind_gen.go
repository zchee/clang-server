package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Describes the kind of entity that a cursor refers to.
type CursorKind uint32

const (
	/*
		A declaration whose specific kind is not exposed via this
		interface.

		Unexposed declarations have the same operations as any other kind
		of declaration; one can extract their location information,
		spelling, find their definitions, etc. However, the specific kind
		of the declaration is not reported.
	*/
	Cursor_UnexposedDecl CursorKind = C.CXCursor_UnexposedDecl
	// A C or C++ struct.
	Cursor_StructDecl = C.CXCursor_StructDecl
	// A C or C++ union.
	Cursor_UnionDecl = C.CXCursor_UnionDecl
	// A C++ class.
	Cursor_ClassDecl = C.CXCursor_ClassDecl
	// An enumeration.
	Cursor_EnumDecl = C.CXCursor_EnumDecl
	// A field (in C) or non-static data member (in C++) in a struct, union, or C++ class.
	Cursor_FieldDecl = C.CXCursor_FieldDecl
	// An enumerator constant.
	Cursor_EnumConstantDecl = C.CXCursor_EnumConstantDecl
	// A function.
	Cursor_FunctionDecl = C.CXCursor_FunctionDecl
	// A variable.
	Cursor_VarDecl = C.CXCursor_VarDecl
	// A function or method parameter.
	Cursor_ParmDecl = C.CXCursor_ParmDecl
	// An Objective-C \@interface.
	Cursor_ObjCInterfaceDecl = C.CXCursor_ObjCInterfaceDecl
	// An Objective-C \@interface for a category.
	Cursor_ObjCCategoryDecl = C.CXCursor_ObjCCategoryDecl
	// An Objective-C \@protocol declaration.
	Cursor_ObjCProtocolDecl = C.CXCursor_ObjCProtocolDecl
	// An Objective-C \@property declaration.
	Cursor_ObjCPropertyDecl = C.CXCursor_ObjCPropertyDecl
	// An Objective-C instance variable.
	Cursor_ObjCIvarDecl = C.CXCursor_ObjCIvarDecl
	// An Objective-C instance method.
	Cursor_ObjCInstanceMethodDecl = C.CXCursor_ObjCInstanceMethodDecl
	// An Objective-C class method.
	Cursor_ObjCClassMethodDecl = C.CXCursor_ObjCClassMethodDecl
	// An Objective-C \@implementation.
	Cursor_ObjCImplementationDecl = C.CXCursor_ObjCImplementationDecl
	// An Objective-C \@implementation for a category.
	Cursor_ObjCCategoryImplDecl = C.CXCursor_ObjCCategoryImplDecl
	// A typedef.
	Cursor_TypedefDecl = C.CXCursor_TypedefDecl
	// A C++ class method.
	Cursor_CXXMethod = C.CXCursor_CXXMethod
	// A C++ namespace.
	Cursor_Namespace = C.CXCursor_Namespace
	// A linkage specification, e.g. 'extern "C"'.
	Cursor_LinkageSpec = C.CXCursor_LinkageSpec
	// A C++ constructor.
	Cursor_Constructor = C.CXCursor_Constructor
	// A C++ destructor.
	Cursor_Destructor = C.CXCursor_Destructor
	// A C++ conversion function.
	Cursor_ConversionFunction = C.CXCursor_ConversionFunction
	// A C++ template type parameter.
	Cursor_TemplateTypeParameter = C.CXCursor_TemplateTypeParameter
	// A C++ non-type template parameter.
	Cursor_NonTypeTemplateParameter = C.CXCursor_NonTypeTemplateParameter
	// A C++ template template parameter.
	Cursor_TemplateTemplateParameter = C.CXCursor_TemplateTemplateParameter
	// A C++ function template.
	Cursor_FunctionTemplate = C.CXCursor_FunctionTemplate
	// A C++ class template.
	Cursor_ClassTemplate = C.CXCursor_ClassTemplate
	// A C++ class template partial specialization.
	Cursor_ClassTemplatePartialSpecialization = C.CXCursor_ClassTemplatePartialSpecialization
	// A C++ namespace alias declaration.
	Cursor_NamespaceAlias = C.CXCursor_NamespaceAlias
	// A C++ using directive.
	Cursor_UsingDirective = C.CXCursor_UsingDirective
	// A C++ using declaration.
	Cursor_UsingDeclaration = C.CXCursor_UsingDeclaration
	// A C++ alias declaration
	Cursor_TypeAliasDecl = C.CXCursor_TypeAliasDecl
	// An Objective-C \@synthesize definition.
	Cursor_ObjCSynthesizeDecl = C.CXCursor_ObjCSynthesizeDecl
	// An Objective-C \@dynamic definition.
	Cursor_ObjCDynamicDecl = C.CXCursor_ObjCDynamicDecl
	// An access specifier.
	Cursor_CXXAccessSpecifier = C.CXCursor_CXXAccessSpecifier
	// An access specifier.
	Cursor_FirstDecl = C.CXCursor_FirstDecl
	// An access specifier.
	Cursor_LastDecl = C.CXCursor_LastDecl
	// An access specifier.
	Cursor_FirstRef = C.CXCursor_FirstRef
	// An access specifier.
	Cursor_ObjCSuperClassRef = C.CXCursor_ObjCSuperClassRef
	// An access specifier.
	Cursor_ObjCProtocolRef = C.CXCursor_ObjCProtocolRef
	// An access specifier.
	Cursor_ObjCClassRef = C.CXCursor_ObjCClassRef
	/*
		A reference to a type declaration.

		A type reference occurs anywhere where a type is named but not
		declared. For example, given:

		\code
		typedef unsigned size_type;
		size_type size;
		\endcode

		The typedef is a declaration of size_type (CXCursor_TypedefDecl),
		while the type of the variable "size" is referenced. The cursor
		referenced by the type of size is the typedef for size_type.
	*/
	Cursor_TypeRef = C.CXCursor_TypeRef
	/*
		A reference to a type declaration.

		A type reference occurs anywhere where a type is named but not
		declared. For example, given:

		\code
		typedef unsigned size_type;
		size_type size;
		\endcode

		The typedef is a declaration of size_type (CXCursor_TypedefDecl),
		while the type of the variable "size" is referenced. The cursor
		referenced by the type of size is the typedef for size_type.
	*/
	Cursor_CXXBaseSpecifier = C.CXCursor_CXXBaseSpecifier
	// A reference to a class template, function template, template template parameter, or class template partial specialization.
	Cursor_TemplateRef = C.CXCursor_TemplateRef
	// A reference to a namespace or namespace alias.
	Cursor_NamespaceRef = C.CXCursor_NamespaceRef
	// A reference to a member of a struct, union, or class that occurs in some non-expression context, e.g., a designated initializer.
	Cursor_MemberRef = C.CXCursor_MemberRef
	/*
		A reference to a labeled statement.

		This cursor kind is used to describe the jump to "start_over" in the
		goto statement in the following example:

		\code
		start_over:
		++counter;

		goto start_over;
		\endcode

		A label reference cursor refers to a label statement.
	*/
	Cursor_LabelRef = C.CXCursor_LabelRef
	/*
		A reference to a set of overloaded functions or function templates
		that has not yet been resolved to a specific function or function template.

		An overloaded declaration reference cursor occurs in C++ templates where
		a dependent name refers to a function. For example:

		\code
		template<typename T> void swap(T&, T&);

		struct X { ... };
		void swap(X&, X&);

		template<typename T>
		void reverse(T* first, T* last) {
		while (first < last - 1) {
		swap(*first, *--last);
		++first;
		}
		}

		struct Y { };
		void swap(Y&, Y&);
		\endcode

		Here, the identifier "swap" is associated with an overloaded declaration
		reference. In the template definition, "swap" refers to either of the two
		"swap" functions declared above, so both results will be available. At
		instantiation time, "swap" may also refer to other functions found via
		argument-dependent lookup (e.g., the "swap" function at the end of the
		example).

		The functions clang_getNumOverloadedDecls() and
		clang_getOverloadedDecl() can be used to retrieve the definitions
		referenced by this cursor.
	*/
	Cursor_OverloadedDeclRef = C.CXCursor_OverloadedDeclRef
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_VariableRef = C.CXCursor_VariableRef
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_LastRef = C.CXCursor_LastRef
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_FirstInvalid = C.CXCursor_FirstInvalid
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_InvalidFile = C.CXCursor_InvalidFile
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_NoDeclFound = C.CXCursor_NoDeclFound
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_NotImplemented = C.CXCursor_NotImplemented
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_InvalidCode = C.CXCursor_InvalidCode
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_LastInvalid = C.CXCursor_LastInvalid
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_FirstExpr = C.CXCursor_FirstExpr
	/*
		An expression whose specific kind is not exposed via this
		interface.

		Unexposed expressions have the same operations as any other kind
		of expression; one can extract their location information,
		spelling, children, etc. However, the specific kind of the
		expression is not reported.
	*/
	Cursor_UnexposedExpr = C.CXCursor_UnexposedExpr
	// An expression that refers to some value declaration, such as a function, variable, or enumerator.
	Cursor_DeclRefExpr = C.CXCursor_DeclRefExpr
	// An expression that refers to a member of a struct, union, class, Objective-C class, etc.
	Cursor_MemberRefExpr = C.CXCursor_MemberRefExpr
	// An expression that calls a function.
	Cursor_CallExpr = C.CXCursor_CallExpr
	// An expression that sends a message to an Objective-C  object or class.
	Cursor_ObjCMessageExpr = C.CXCursor_ObjCMessageExpr
	// An expression that represents a block literal.
	Cursor_BlockExpr = C.CXCursor_BlockExpr
	// An integer literal.
	Cursor_IntegerLiteral = C.CXCursor_IntegerLiteral
	// A floating point number literal.
	Cursor_FloatingLiteral = C.CXCursor_FloatingLiteral
	// An imaginary number literal.
	Cursor_ImaginaryLiteral = C.CXCursor_ImaginaryLiteral
	// A string literal.
	Cursor_StringLiteral = C.CXCursor_StringLiteral
	// A character literal.
	Cursor_CharacterLiteral = C.CXCursor_CharacterLiteral
	/*
		A parenthesized expression, e.g. "(1)".

		This AST node is only formed if full location information is requested.
	*/
	Cursor_ParenExpr = C.CXCursor_ParenExpr
	// This represents the unary-expression's (except sizeof and alignof).
	Cursor_UnaryOperator = C.CXCursor_UnaryOperator
	// [C99 6.5.2.1] Array Subscripting.
	Cursor_ArraySubscriptExpr = C.CXCursor_ArraySubscriptExpr
	// A builtin binary operation expression such as "x + y" or "x <= y".
	Cursor_BinaryOperator = C.CXCursor_BinaryOperator
	// Compound assignment such as "+=".
	Cursor_CompoundAssignOperator = C.CXCursor_CompoundAssignOperator
	// The ?: ternary operator.
	Cursor_ConditionalOperator = C.CXCursor_ConditionalOperator
	/*
		An explicit cast in C (C99 6.5.4) or a C-style cast in C++
		(C++ [expr.cast]), which uses the syntax (Type)expr.

		For example: (int)f.
	*/
	Cursor_CStyleCastExpr = C.CXCursor_CStyleCastExpr
	// [C99 6.5.2.5]
	Cursor_CompoundLiteralExpr = C.CXCursor_CompoundLiteralExpr
	// Describes an C or C++ initializer list.
	Cursor_InitListExpr = C.CXCursor_InitListExpr
	// The GNU address of label extension, representing &&label.
	Cursor_AddrLabelExpr = C.CXCursor_AddrLabelExpr
	// This is the GNU Statement Expression extension: ({int X=4; X;})
	Cursor_StmtExpr = C.CXCursor_StmtExpr
	// Represents a C11 generic selection.
	Cursor_GenericSelectionExpr = C.CXCursor_GenericSelectionExpr
	/*
		Implements the GNU __null extension, which is a name for a null
		pointer constant that has integral type (e.g., int or long) and is the same
		size and alignment as a pointer.

		The __null extension is typically only used by system headers, which define
		NULL as __null in C++ rather than using 0 (which is an integer that may not
		match the size of a pointer).
	*/
	Cursor_GNUNullExpr = C.CXCursor_GNUNullExpr
	// C++'s static_cast<> expression.
	Cursor_CXXStaticCastExpr = C.CXCursor_CXXStaticCastExpr
	// C++'s dynamic_cast<> expression.
	Cursor_CXXDynamicCastExpr = C.CXCursor_CXXDynamicCastExpr
	// C++'s reinterpret_cast<> expression.
	Cursor_CXXReinterpretCastExpr = C.CXCursor_CXXReinterpretCastExpr
	// C++'s const_cast<> expression.
	Cursor_CXXConstCastExpr = C.CXCursor_CXXConstCastExpr
	/*
		Represents an explicit C++ type conversion that uses "functional"
		notion (C++ [expr.type.conv]).

		Example:
		\code
		x = int(0.5);
		\endcode
	*/
	Cursor_CXXFunctionalCastExpr = C.CXCursor_CXXFunctionalCastExpr
	// A C++ typeid expression (C++ [expr.typeid]).
	Cursor_CXXTypeidExpr = C.CXCursor_CXXTypeidExpr
	// [C++ 2.13.5] C++ Boolean Literal.
	Cursor_CXXBoolLiteralExpr = C.CXCursor_CXXBoolLiteralExpr
	// [C++0x 2.14.7] C++ Pointer Literal.
	Cursor_CXXNullPtrLiteralExpr = C.CXCursor_CXXNullPtrLiteralExpr
	// Represents the "this" expression in C++
	Cursor_CXXThisExpr = C.CXCursor_CXXThisExpr
	/*
		[C++ 15] C++ Throw Expression.

		This handles 'throw' and 'throw' assignment-expression. When
		assignment-expression isn't present, Op will be null.
	*/
	Cursor_CXXThrowExpr = C.CXCursor_CXXThrowExpr
	// A new expression for memory allocation and constructor calls, e.g: "new CXXNewExpr(foo)".
	Cursor_CXXNewExpr = C.CXCursor_CXXNewExpr
	// A delete expression for memory deallocation and destructor calls, e.g. "delete[] pArray".
	Cursor_CXXDeleteExpr = C.CXCursor_CXXDeleteExpr
	// A unary expression. (noexcept, sizeof, or other traits)
	Cursor_UnaryExpr = C.CXCursor_UnaryExpr
	// An Objective-C string literal i.e. @"foo".
	Cursor_ObjCStringLiteral = C.CXCursor_ObjCStringLiteral
	// An Objective-C \@encode expression.
	Cursor_ObjCEncodeExpr = C.CXCursor_ObjCEncodeExpr
	// An Objective-C \@selector expression.
	Cursor_ObjCSelectorExpr = C.CXCursor_ObjCSelectorExpr
	// An Objective-C \@protocol expression.
	Cursor_ObjCProtocolExpr = C.CXCursor_ObjCProtocolExpr
	/*
		An Objective-C "bridged" cast expression, which casts between
		Objective-C pointers and C pointers, transferring ownership in the process.

		\code
		NSString *str = (__bridge_transfer NSString *)CFCreateString();
		\endcode
	*/
	Cursor_ObjCBridgedCastExpr = C.CXCursor_ObjCBridgedCastExpr
	/*
		Represents a C++0x pack expansion that produces a sequence of
		expressions.

		A pack expansion expression contains a pattern (which itself is an
		expression) followed by an ellipsis. For example:

		\code
		template<typename F, typename ...Types>
		void forward(F f, Types &&...args) {
		f(static_cast<Types&&>(args)...);
		}
		\endcode
	*/
	Cursor_PackExpansionExpr = C.CXCursor_PackExpansionExpr
	/*
		Represents an expression that computes the length of a parameter
		pack.

		\code
		template<typename ...Types>
		struct count {
		static const unsigned value = sizeof...(Types);
		};
		\endcode
	*/
	Cursor_SizeOfPackExpr = C.CXCursor_SizeOfPackExpr
	Cursor_LambdaExpr     = C.CXCursor_LambdaExpr
	// Objective-c Boolean Literal.
	Cursor_ObjCBoolLiteralExpr = C.CXCursor_ObjCBoolLiteralExpr
	// Represents the "self" expression in an Objective-C method.
	Cursor_ObjCSelfExpr = C.CXCursor_ObjCSelfExpr
	// OpenMP 4.0 [2.4, Array Section].
	Cursor_OMPArraySectionExpr = C.CXCursor_OMPArraySectionExpr
	// Represents an @available(...) check.
	Cursor_ObjCAvailabilityCheckExpr = C.CXCursor_ObjCAvailabilityCheckExpr
	// Represents an @available(...) check.
	Cursor_LastExpr = C.CXCursor_LastExpr
	// Represents an @available(...) check.
	Cursor_FirstStmt = C.CXCursor_FirstStmt
	/*
		A statement whose specific kind is not exposed via this
		interface.

		Unexposed statements have the same operations as any other kind of
		statement; one can extract their location information, spelling,
		children, etc. However, the specific kind of the statement is not
		reported.
	*/
	Cursor_UnexposedStmt = C.CXCursor_UnexposedStmt
	/*
		A labelled statement in a function.

		This cursor kind is used to describe the "start_over:" label statement in
		the following example:

		\code
		start_over:
		++counter;
		\endcode
	*/
	Cursor_LabelStmt = C.CXCursor_LabelStmt
	/*
		A group of statements like { stmt stmt }.

		This cursor kind is used to describe compound statements, e.g. function
		bodies.
	*/
	Cursor_CompoundStmt = C.CXCursor_CompoundStmt
	// A case statement.
	Cursor_CaseStmt = C.CXCursor_CaseStmt
	// A default statement.
	Cursor_DefaultStmt = C.CXCursor_DefaultStmt
	// An if statement
	Cursor_IfStmt = C.CXCursor_IfStmt
	// A switch statement.
	Cursor_SwitchStmt = C.CXCursor_SwitchStmt
	// A while statement.
	Cursor_WhileStmt = C.CXCursor_WhileStmt
	// A do statement.
	Cursor_DoStmt = C.CXCursor_DoStmt
	// A for statement.
	Cursor_ForStmt = C.CXCursor_ForStmt
	// A goto statement.
	Cursor_GotoStmt = C.CXCursor_GotoStmt
	// An indirect goto statement.
	Cursor_IndirectGotoStmt = C.CXCursor_IndirectGotoStmt
	// A continue statement.
	Cursor_ContinueStmt = C.CXCursor_ContinueStmt
	// A break statement.
	Cursor_BreakStmt = C.CXCursor_BreakStmt
	// A return statement.
	Cursor_ReturnStmt = C.CXCursor_ReturnStmt
	// A GCC inline assembly statement extension.
	Cursor_GCCAsmStmt = C.CXCursor_GCCAsmStmt
	// A GCC inline assembly statement extension.
	Cursor_AsmStmt = C.CXCursor_AsmStmt
	// Objective-C's overall \@try-\@catch-\@finally statement.
	Cursor_ObjCAtTryStmt = C.CXCursor_ObjCAtTryStmt
	// Objective-C's \@catch statement.
	Cursor_ObjCAtCatchStmt = C.CXCursor_ObjCAtCatchStmt
	// Objective-C's \@finally statement.
	Cursor_ObjCAtFinallyStmt = C.CXCursor_ObjCAtFinallyStmt
	// Objective-C's \@throw statement.
	Cursor_ObjCAtThrowStmt = C.CXCursor_ObjCAtThrowStmt
	// Objective-C's \@synchronized statement.
	Cursor_ObjCAtSynchronizedStmt = C.CXCursor_ObjCAtSynchronizedStmt
	// Objective-C's autorelease pool statement.
	Cursor_ObjCAutoreleasePoolStmt = C.CXCursor_ObjCAutoreleasePoolStmt
	// Objective-C's collection statement.
	Cursor_ObjCForCollectionStmt = C.CXCursor_ObjCForCollectionStmt
	// C++'s catch statement.
	Cursor_CXXCatchStmt = C.CXCursor_CXXCatchStmt
	// C++'s try statement.
	Cursor_CXXTryStmt = C.CXCursor_CXXTryStmt
	// C++'s for (* : *) statement.
	Cursor_CXXForRangeStmt = C.CXCursor_CXXForRangeStmt
	// Windows Structured Exception Handling's try statement.
	Cursor_SEHTryStmt = C.CXCursor_SEHTryStmt
	// Windows Structured Exception Handling's except statement.
	Cursor_SEHExceptStmt = C.CXCursor_SEHExceptStmt
	// Windows Structured Exception Handling's finally statement.
	Cursor_SEHFinallyStmt = C.CXCursor_SEHFinallyStmt
	// A MS inline assembly statement extension.
	Cursor_MSAsmStmt = C.CXCursor_MSAsmStmt
	/*
		The null statement ";": C99 6.8.3p3.

		This cursor kind is used to describe the null statement.
	*/
	Cursor_NullStmt = C.CXCursor_NullStmt
	// Adaptor class for mixing declarations with statements and expressions.
	Cursor_DeclStmt = C.CXCursor_DeclStmt
	// OpenMP parallel directive.
	Cursor_OMPParallelDirective = C.CXCursor_OMPParallelDirective
	// OpenMP SIMD directive.
	Cursor_OMPSimdDirective = C.CXCursor_OMPSimdDirective
	// OpenMP for directive.
	Cursor_OMPForDirective = C.CXCursor_OMPForDirective
	// OpenMP sections directive.
	Cursor_OMPSectionsDirective = C.CXCursor_OMPSectionsDirective
	// OpenMP section directive.
	Cursor_OMPSectionDirective = C.CXCursor_OMPSectionDirective
	// OpenMP single directive.
	Cursor_OMPSingleDirective = C.CXCursor_OMPSingleDirective
	// OpenMP parallel for directive.
	Cursor_OMPParallelForDirective = C.CXCursor_OMPParallelForDirective
	// OpenMP parallel sections directive.
	Cursor_OMPParallelSectionsDirective = C.CXCursor_OMPParallelSectionsDirective
	// OpenMP task directive.
	Cursor_OMPTaskDirective = C.CXCursor_OMPTaskDirective
	// OpenMP master directive.
	Cursor_OMPMasterDirective = C.CXCursor_OMPMasterDirective
	// OpenMP critical directive.
	Cursor_OMPCriticalDirective = C.CXCursor_OMPCriticalDirective
	// OpenMP taskyield directive.
	Cursor_OMPTaskyieldDirective = C.CXCursor_OMPTaskyieldDirective
	// OpenMP barrier directive.
	Cursor_OMPBarrierDirective = C.CXCursor_OMPBarrierDirective
	// OpenMP taskwait directive.
	Cursor_OMPTaskwaitDirective = C.CXCursor_OMPTaskwaitDirective
	// OpenMP flush directive.
	Cursor_OMPFlushDirective = C.CXCursor_OMPFlushDirective
	// Windows Structured Exception Handling's leave statement.
	Cursor_SEHLeaveStmt = C.CXCursor_SEHLeaveStmt
	// OpenMP ordered directive.
	Cursor_OMPOrderedDirective = C.CXCursor_OMPOrderedDirective
	// OpenMP atomic directive.
	Cursor_OMPAtomicDirective = C.CXCursor_OMPAtomicDirective
	// OpenMP for SIMD directive.
	Cursor_OMPForSimdDirective = C.CXCursor_OMPForSimdDirective
	// OpenMP parallel for SIMD directive.
	Cursor_OMPParallelForSimdDirective = C.CXCursor_OMPParallelForSimdDirective
	// OpenMP target directive.
	Cursor_OMPTargetDirective = C.CXCursor_OMPTargetDirective
	// OpenMP teams directive.
	Cursor_OMPTeamsDirective = C.CXCursor_OMPTeamsDirective
	// OpenMP taskgroup directive.
	Cursor_OMPTaskgroupDirective = C.CXCursor_OMPTaskgroupDirective
	// OpenMP cancellation point directive.
	Cursor_OMPCancellationPointDirective = C.CXCursor_OMPCancellationPointDirective
	// OpenMP cancel directive.
	Cursor_OMPCancelDirective = C.CXCursor_OMPCancelDirective
	// OpenMP target data directive.
	Cursor_OMPTargetDataDirective = C.CXCursor_OMPTargetDataDirective
	// OpenMP taskloop directive.
	Cursor_OMPTaskLoopDirective = C.CXCursor_OMPTaskLoopDirective
	// OpenMP taskloop simd directive.
	Cursor_OMPTaskLoopSimdDirective = C.CXCursor_OMPTaskLoopSimdDirective
	// OpenMP distribute directive.
	Cursor_OMPDistributeDirective = C.CXCursor_OMPDistributeDirective
	// OpenMP target enter data directive.
	Cursor_OMPTargetEnterDataDirective = C.CXCursor_OMPTargetEnterDataDirective
	// OpenMP target exit data directive.
	Cursor_OMPTargetExitDataDirective = C.CXCursor_OMPTargetExitDataDirective
	// OpenMP target parallel directive.
	Cursor_OMPTargetParallelDirective = C.CXCursor_OMPTargetParallelDirective
	// OpenMP target parallel for directive.
	Cursor_OMPTargetParallelForDirective = C.CXCursor_OMPTargetParallelForDirective
	// OpenMP target update directive.
	Cursor_OMPTargetUpdateDirective = C.CXCursor_OMPTargetUpdateDirective
	// OpenMP distribute parallel for directive.
	Cursor_OMPDistributeParallelForDirective = C.CXCursor_OMPDistributeParallelForDirective
	// OpenMP distribute parallel for simd directive.
	Cursor_OMPDistributeParallelForSimdDirective = C.CXCursor_OMPDistributeParallelForSimdDirective
	// OpenMP distribute simd directive.
	Cursor_OMPDistributeSimdDirective = C.CXCursor_OMPDistributeSimdDirective
	// OpenMP target parallel for simd directive.
	Cursor_OMPTargetParallelForSimdDirective = C.CXCursor_OMPTargetParallelForSimdDirective
	// OpenMP target parallel for simd directive.
	Cursor_LastStmt = C.CXCursor_LastStmt
	/*
		Cursor that represents the translation unit itself.

		The translation unit cursor exists primarily to act as the root
		cursor for traversing the contents of a translation unit.
	*/
	Cursor_TranslationUnit = C.CXCursor_TranslationUnit
	/*
		Cursor that represents the translation unit itself.

		The translation unit cursor exists primarily to act as the root
		cursor for traversing the contents of a translation unit.
	*/
	Cursor_FirstAttr = C.CXCursor_FirstAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_UnexposedAttr = C.CXCursor_UnexposedAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_IBActionAttr = C.CXCursor_IBActionAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_IBOutletAttr = C.CXCursor_IBOutletAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_IBOutletCollectionAttr = C.CXCursor_IBOutletCollectionAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CXXFinalAttr = C.CXCursor_CXXFinalAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CXXOverrideAttr = C.CXCursor_CXXOverrideAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_AnnotateAttr = C.CXCursor_AnnotateAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_AsmLabelAttr = C.CXCursor_AsmLabelAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_PackedAttr = C.CXCursor_PackedAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_PureAttr = C.CXCursor_PureAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_ConstAttr = C.CXCursor_ConstAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_NoDuplicateAttr = C.CXCursor_NoDuplicateAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CUDAConstantAttr = C.CXCursor_CUDAConstantAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CUDADeviceAttr = C.CXCursor_CUDADeviceAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CUDAGlobalAttr = C.CXCursor_CUDAGlobalAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CUDAHostAttr = C.CXCursor_CUDAHostAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CUDASharedAttr = C.CXCursor_CUDASharedAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_VisibilityAttr = C.CXCursor_VisibilityAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_DLLExport = C.CXCursor_DLLExport
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_DLLImport = C.CXCursor_DLLImport
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_LastAttr = C.CXCursor_LastAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_PreprocessingDirective = C.CXCursor_PreprocessingDirective
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_MacroDefinition = C.CXCursor_MacroDefinition
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_MacroExpansion = C.CXCursor_MacroExpansion
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_MacroInstantiation = C.CXCursor_MacroInstantiation
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_InclusionDirective = C.CXCursor_InclusionDirective
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_FirstPreprocessing = C.CXCursor_FirstPreprocessing
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_LastPreprocessing = C.CXCursor_LastPreprocessing
	// A module import declaration.
	Cursor_ModuleImportDecl = C.CXCursor_ModuleImportDecl
	// A module import declaration.
	Cursor_TypeAliasTemplateDecl = C.CXCursor_TypeAliasTemplateDecl
	// A static_assert or _Static_assert node
	Cursor_StaticAssert = C.CXCursor_StaticAssert
	// A static_assert or _Static_assert node
	Cursor_FirstExtraDecl = C.CXCursor_FirstExtraDecl
	// A static_assert or _Static_assert node
	Cursor_LastExtraDecl = C.CXCursor_LastExtraDecl
	// A code completion overload candidate.
	Cursor_OverloadCandidate = C.CXCursor_OverloadCandidate
)

// Determine whether the given cursor kind represents a declaration.
func (ck CursorKind) IsDeclaration() bool {
	o := C.clang_isDeclaration(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

/*
	Determine whether the given cursor kind represents a simple
	reference.

	Note that other kinds of cursors (such as expressions) can also refer to
	other cursors. Use clang_getCursorReferenced() to determine whether a
	particular cursor refers to another entity.
*/
func (ck CursorKind) IsReference() bool {
	o := C.clang_isReference(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// Determine whether the given cursor kind represents an expression.
func (ck CursorKind) IsExpression() bool {
	o := C.clang_isExpression(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// Determine whether the given cursor kind represents a statement.
func (ck CursorKind) IsStatement() bool {
	o := C.clang_isStatement(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// Determine whether the given cursor kind represents an attribute.
func (ck CursorKind) IsAttribute() bool {
	o := C.clang_isAttribute(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// Determine whether the given cursor kind represents an invalid cursor.
func (ck CursorKind) IsInvalid() bool {
	o := C.clang_isInvalid(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// Determine whether the given cursor kind represents a translation unit.
func (ck CursorKind) IsTranslationUnit() bool {
	o := C.clang_isTranslationUnit(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// * Determine whether the given cursor represents a preprocessing element, such as a preprocessor directive or macro instantiation.
func (ck CursorKind) IsPreprocessing() bool {
	o := C.clang_isPreprocessing(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// * Determine whether the given cursor represents a currently unexposed piece of the AST (e.g., CXCursor_UnexposedStmt).
func (ck CursorKind) IsUnexposed() bool {
	o := C.clang_isUnexposed(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

func (ck CursorKind) Spelling() string {
	o := cxstring{C.clang_getCursorKindSpelling(C.enum_CXCursorKind(ck))}
	defer o.Dispose()

	return o.String()
}

func (ck CursorKind) String() string {
	return ck.Spelling()
}
