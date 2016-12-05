package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type IdxEntityKind uint32

const (
	IdxEntity_Unexposed             IdxEntityKind = C.CXIdxEntity_Unexposed
	IdxEntity_Typedef                             = C.CXIdxEntity_Typedef
	IdxEntity_Function                            = C.CXIdxEntity_Function
	IdxEntity_Variable                            = C.CXIdxEntity_Variable
	IdxEntity_Field                               = C.CXIdxEntity_Field
	IdxEntity_EnumConstant                        = C.CXIdxEntity_EnumConstant
	IdxEntity_ObjCClass                           = C.CXIdxEntity_ObjCClass
	IdxEntity_ObjCProtocol                        = C.CXIdxEntity_ObjCProtocol
	IdxEntity_ObjCCategory                        = C.CXIdxEntity_ObjCCategory
	IdxEntity_ObjCInstanceMethod                  = C.CXIdxEntity_ObjCInstanceMethod
	IdxEntity_ObjCClassMethod                     = C.CXIdxEntity_ObjCClassMethod
	IdxEntity_ObjCProperty                        = C.CXIdxEntity_ObjCProperty
	IdxEntity_ObjCIvar                            = C.CXIdxEntity_ObjCIvar
	IdxEntity_Enum                                = C.CXIdxEntity_Enum
	IdxEntity_Struct                              = C.CXIdxEntity_Struct
	IdxEntity_Union                               = C.CXIdxEntity_Union
	IdxEntity_CXXClass                            = C.CXIdxEntity_CXXClass
	IdxEntity_CXXNamespace                        = C.CXIdxEntity_CXXNamespace
	IdxEntity_CXXNamespaceAlias                   = C.CXIdxEntity_CXXNamespaceAlias
	IdxEntity_CXXStaticVariable                   = C.CXIdxEntity_CXXStaticVariable
	IdxEntity_CXXStaticMethod                     = C.CXIdxEntity_CXXStaticMethod
	IdxEntity_CXXInstanceMethod                   = C.CXIdxEntity_CXXInstanceMethod
	IdxEntity_CXXConstructor                      = C.CXIdxEntity_CXXConstructor
	IdxEntity_CXXDestructor                       = C.CXIdxEntity_CXXDestructor
	IdxEntity_CXXConversionFunction               = C.CXIdxEntity_CXXConversionFunction
	IdxEntity_CXXTypeAlias                        = C.CXIdxEntity_CXXTypeAlias
	IdxEntity_CXXInterface                        = C.CXIdxEntity_CXXInterface
)

func (iek IdxEntityKind) IsEntityObjCContainerKind() bool {
	o := C.clang_index_isEntityObjCContainerKind(C.CXIdxEntityKind(iek))

	return o != C.int(0)
}

func (iek IdxEntityKind) Spelling() string {
	switch iek {
	case IdxEntity_Unexposed:
		return "IdxEntity=Unexposed"
	case IdxEntity_Typedef:
		return "IdxEntity=Typedef"
	case IdxEntity_Function:
		return "IdxEntity=Function"
	case IdxEntity_Variable:
		return "IdxEntity=Variable"
	case IdxEntity_Field:
		return "IdxEntity=Field"
	case IdxEntity_EnumConstant:
		return "IdxEntity=EnumConstant"
	case IdxEntity_ObjCClass:
		return "IdxEntity=ObjCClass"
	case IdxEntity_ObjCProtocol:
		return "IdxEntity=ObjCProtocol"
	case IdxEntity_ObjCCategory:
		return "IdxEntity=ObjCCategory"
	case IdxEntity_ObjCInstanceMethod:
		return "IdxEntity=ObjCInstanceMethod"
	case IdxEntity_ObjCClassMethod:
		return "IdxEntity=ObjCClassMethod"
	case IdxEntity_ObjCProperty:
		return "IdxEntity=ObjCProperty"
	case IdxEntity_ObjCIvar:
		return "IdxEntity=ObjCIvar"
	case IdxEntity_Enum:
		return "IdxEntity=Enum"
	case IdxEntity_Struct:
		return "IdxEntity=Struct"
	case IdxEntity_Union:
		return "IdxEntity=Union"
	case IdxEntity_CXXClass:
		return "IdxEntity=CXXClass"
	case IdxEntity_CXXNamespace:
		return "IdxEntity=CXXNamespace"
	case IdxEntity_CXXNamespaceAlias:
		return "IdxEntity=CXXNamespaceAlias"
	case IdxEntity_CXXStaticVariable:
		return "IdxEntity=CXXStaticVariable"
	case IdxEntity_CXXStaticMethod:
		return "IdxEntity=CXXStaticMethod"
	case IdxEntity_CXXInstanceMethod:
		return "IdxEntity=CXXInstanceMethod"
	case IdxEntity_CXXConstructor:
		return "IdxEntity=CXXConstructor"
	case IdxEntity_CXXDestructor:
		return "IdxEntity=CXXDestructor"
	case IdxEntity_CXXConversionFunction:
		return "IdxEntity=CXXConversionFunction"
	case IdxEntity_CXXTypeAlias:
		return "IdxEntity=CXXTypeAlias"
	case IdxEntity_CXXInterface:
		return "IdxEntity=CXXInterface"
	}

	return fmt.Sprintf("IdxEntityKind unkown %d", int(iek))
}

func (iek IdxEntityKind) String() string {
	return iek.Spelling()
}
