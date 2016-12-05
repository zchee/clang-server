package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Extra C++ template information for an entity. This can apply to: CXIdxEntity_Function CXIdxEntity_CXXClass CXIdxEntity_CXXStaticMethod CXIdxEntity_CXXInstanceMethod CXIdxEntity_CXXConstructor CXIdxEntity_CXXConversionFunction CXIdxEntity_CXXTypeAlias
type IdxEntityCXXTemplateKind uint32

const (
	IdxEntity_NonTemplate                   IdxEntityCXXTemplateKind = C.CXIdxEntity_NonTemplate
	IdxEntity_Template                                               = C.CXIdxEntity_Template
	IdxEntity_TemplatePartialSpecialization                          = C.CXIdxEntity_TemplatePartialSpecialization
	IdxEntity_TemplateSpecialization                                 = C.CXIdxEntity_TemplateSpecialization
)

func (iecxxtk IdxEntityCXXTemplateKind) Spelling() string {
	switch iecxxtk {
	case IdxEntity_NonTemplate:
		return "IdxEntity=NonTemplate"
	case IdxEntity_Template:
		return "IdxEntity=Template"
	case IdxEntity_TemplatePartialSpecialization:
		return "IdxEntity=TemplatePartialSpecialization"
	case IdxEntity_TemplateSpecialization:
		return "IdxEntity=TemplateSpecialization"
	}

	return fmt.Sprintf("IdxEntityCXXTemplateKind unkown %d", int(iecxxtk))
}

func (iecxxtk IdxEntityCXXTemplateKind) String() string {
	return iecxxtk.Spelling()
}
