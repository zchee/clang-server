package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type VisibilityKind uint32

const (
	// This value indicates that no visibility information is available for a provided CXCursor.
	Visibility_Invalid VisibilityKind = C.CXVisibility_Invalid
	// Symbol not seen by the linker.
	Visibility_Hidden = C.CXVisibility_Hidden
	// Symbol seen by the linker but resolves to a symbol inside this object.
	Visibility_Protected = C.CXVisibility_Protected
	// Symbol seen by the linker and acts like a normal symbol.
	Visibility_Default = C.CXVisibility_Default
)

func (vk VisibilityKind) Spelling() string {
	switch vk {
	case Visibility_Invalid:
		return "Visibility=Invalid"
	case Visibility_Hidden:
		return "Visibility=Hidden"
	case Visibility_Protected:
		return "Visibility=Protected"
	case Visibility_Default:
		return "Visibility=Default"
	}

	return fmt.Sprintf("VisibilityKind unkown %d", int(vk))
}

func (vk VisibilityKind) String() string {
	return vk.Spelling()
}
