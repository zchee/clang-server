package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type IdxObjCContainerKind uint32

const (
	IdxObjCContainer_ForwardRef     IdxObjCContainerKind = C.CXIdxObjCContainer_ForwardRef
	IdxObjCContainer_Interface                           = C.CXIdxObjCContainer_Interface
	IdxObjCContainer_Implementation                      = C.CXIdxObjCContainer_Implementation
)

func (iocck IdxObjCContainerKind) Spelling() string {
	switch iocck {
	case IdxObjCContainer_ForwardRef:
		return "IdxObjCContainer=ForwardRef"
	case IdxObjCContainer_Interface:
		return "IdxObjCContainer=Interface"
	case IdxObjCContainer_Implementation:
		return "IdxObjCContainer=Implementation"
	}

	return fmt.Sprintf("IdxObjCContainerKind unkown %d", int(iocck))
}

func (iocck IdxObjCContainerKind) String() string {
	return iocck.Spelling()
}
