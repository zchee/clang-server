package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type IdxAttrKind uint32

const (
	IdxAttr_Unexposed          IdxAttrKind = C.CXIdxAttr_Unexposed
	IdxAttr_IBAction                       = C.CXIdxAttr_IBAction
	IdxAttr_IBOutlet                       = C.CXIdxAttr_IBOutlet
	IdxAttr_IBOutletCollection             = C.CXIdxAttr_IBOutletCollection
)

func (iak IdxAttrKind) Spelling() string {
	switch iak {
	case IdxAttr_Unexposed:
		return "IdxAttr=Unexposed"
	case IdxAttr_IBAction:
		return "IdxAttr=IBAction"
	case IdxAttr_IBOutlet:
		return "IdxAttr=IBOutlet"
	case IdxAttr_IBOutletCollection:
		return "IdxAttr=IBOutletCollection"
	}

	return fmt.Sprintf("IdxAttrKind unkown %d", int(iak))
}

func (iak IdxAttrKind) String() string {
	return iak.Spelling()
}
