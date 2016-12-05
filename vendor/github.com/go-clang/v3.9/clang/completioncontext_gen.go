package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

/*
	Bits that represent the context under which completion is occurring.

	The enumerators in this enumeration may be bitwise-OR'd together if multiple
	contexts are occurring simultaneously.
*/
type CompletionContext uint32

const (
	// The context for completions is unexposed, as only Clang results should be included. (This is equivalent to having no context bits set.)
	CompletionContext_Unexposed CompletionContext = C.CXCompletionContext_Unexposed
	// Completions for any possible type should be included in the results.
	CompletionContext_AnyType = C.CXCompletionContext_AnyType
	// Completions for any possible value (variables, function calls, etc.) should be included in the results.
	CompletionContext_AnyValue = C.CXCompletionContext_AnyValue
	// Completions for values that resolve to an Objective-C object should be included in the results.
	CompletionContext_ObjCObjectValue = C.CXCompletionContext_ObjCObjectValue
	// Completions for values that resolve to an Objective-C selector should be included in the results.
	CompletionContext_ObjCSelectorValue = C.CXCompletionContext_ObjCSelectorValue
	// Completions for values that resolve to a C++ class type should be included in the results.
	CompletionContext_CXXClassTypeValue = C.CXCompletionContext_CXXClassTypeValue
	// Completions for fields of the member being accessed using the dot operator should be included in the results.
	CompletionContext_DotMemberAccess = C.CXCompletionContext_DotMemberAccess
	// Completions for fields of the member being accessed using the arrow operator should be included in the results.
	CompletionContext_ArrowMemberAccess = C.CXCompletionContext_ArrowMemberAccess
	// Completions for properties of the Objective-C object being accessed using the dot operator should be included in the results.
	CompletionContext_ObjCPropertyAccess = C.CXCompletionContext_ObjCPropertyAccess
	// Completions for enum tags should be included in the results.
	CompletionContext_EnumTag = C.CXCompletionContext_EnumTag
	// Completions for union tags should be included in the results.
	CompletionContext_UnionTag = C.CXCompletionContext_UnionTag
	// Completions for struct tags should be included in the results.
	CompletionContext_StructTag = C.CXCompletionContext_StructTag
	// Completions for C++ class names should be included in the results.
	CompletionContext_ClassTag = C.CXCompletionContext_ClassTag
	// Completions for C++ namespaces and namespace aliases should be included in the results.
	CompletionContext_Namespace = C.CXCompletionContext_Namespace
	// Completions for C++ nested name specifiers should be included in the results.
	CompletionContext_NestedNameSpecifier = C.CXCompletionContext_NestedNameSpecifier
	// Completions for Objective-C interfaces (classes) should be included in the results.
	CompletionContext_ObjCInterface = C.CXCompletionContext_ObjCInterface
	// Completions for Objective-C protocols should be included in the results.
	CompletionContext_ObjCProtocol = C.CXCompletionContext_ObjCProtocol
	// Completions for Objective-C categories should be included in the results.
	CompletionContext_ObjCCategory = C.CXCompletionContext_ObjCCategory
	// Completions for Objective-C instance messages should be included in the results.
	CompletionContext_ObjCInstanceMessage = C.CXCompletionContext_ObjCInstanceMessage
	// Completions for Objective-C class messages should be included in the results.
	CompletionContext_ObjCClassMessage = C.CXCompletionContext_ObjCClassMessage
	// Completions for Objective-C selector names should be included in the results.
	CompletionContext_ObjCSelectorName = C.CXCompletionContext_ObjCSelectorName
	// Completions for preprocessor macro names should be included in the results.
	CompletionContext_MacroName = C.CXCompletionContext_MacroName
	// Natural language completions should be included in the results.
	CompletionContext_NaturalLanguage = C.CXCompletionContext_NaturalLanguage
	// The current context is unknown, so set all contexts.
	CompletionContext_Unknown = C.CXCompletionContext_Unknown
)

func (cc CompletionContext) Spelling() string {
	switch cc {
	case CompletionContext_Unexposed:
		return "CompletionContext=Unexposed"
	case CompletionContext_AnyType:
		return "CompletionContext=AnyType"
	case CompletionContext_AnyValue:
		return "CompletionContext=AnyValue"
	case CompletionContext_ObjCObjectValue:
		return "CompletionContext=ObjCObjectValue"
	case CompletionContext_ObjCSelectorValue:
		return "CompletionContext=ObjCSelectorValue"
	case CompletionContext_CXXClassTypeValue:
		return "CompletionContext=CXXClassTypeValue"
	case CompletionContext_DotMemberAccess:
		return "CompletionContext=DotMemberAccess"
	case CompletionContext_ArrowMemberAccess:
		return "CompletionContext=ArrowMemberAccess"
	case CompletionContext_ObjCPropertyAccess:
		return "CompletionContext=ObjCPropertyAccess"
	case CompletionContext_EnumTag:
		return "CompletionContext=EnumTag"
	case CompletionContext_UnionTag:
		return "CompletionContext=UnionTag"
	case CompletionContext_StructTag:
		return "CompletionContext=StructTag"
	case CompletionContext_ClassTag:
		return "CompletionContext=ClassTag"
	case CompletionContext_Namespace:
		return "CompletionContext=Namespace"
	case CompletionContext_NestedNameSpecifier:
		return "CompletionContext=NestedNameSpecifier"
	case CompletionContext_ObjCInterface:
		return "CompletionContext=ObjCInterface"
	case CompletionContext_ObjCProtocol:
		return "CompletionContext=ObjCProtocol"
	case CompletionContext_ObjCCategory:
		return "CompletionContext=ObjCCategory"
	case CompletionContext_ObjCInstanceMessage:
		return "CompletionContext=ObjCInstanceMessage"
	case CompletionContext_ObjCClassMessage:
		return "CompletionContext=ObjCClassMessage"
	case CompletionContext_ObjCSelectorName:
		return "CompletionContext=ObjCSelectorName"
	case CompletionContext_MacroName:
		return "CompletionContext=MacroName"
	case CompletionContext_NaturalLanguage:
		return "CompletionContext=NaturalLanguage"
	case CompletionContext_Unknown:
		return "CompletionContext=Unknown"
	}

	return fmt.Sprintf("CompletionContext unkown %d", int(cc))
}

func (cc CompletionContext) String() string {
	return cc.Spelling()
}
