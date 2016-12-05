package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Represents the C++ access control level to a base class for a cursor with kind CX_CXXBaseSpecifier.
type AccessSpecifier uint32

const (
	AccessSpecifier_Invalid   AccessSpecifier = C.CX_CXXInvalidAccessSpecifier
	AccessSpecifier_Public                    = C.CX_CXXPublic
	AccessSpecifier_Protected                 = C.CX_CXXProtected
	AccessSpecifier_Private                   = C.CX_CXXPrivate
)

func (as AccessSpecifier) Spelling() string {
	switch as {
	case AccessSpecifier_Invalid:
		return "AccessSpecifier=Invalid"
	case AccessSpecifier_Public:
		return "AccessSpecifier=Public"
	case AccessSpecifier_Protected:
		return "AccessSpecifier=Protected"
	case AccessSpecifier_Private:
		return "AccessSpecifier=Private"
	}

	return fmt.Sprintf("AccessSpecifier unkown %d", int(as))
}

func (as AccessSpecifier) String() string {
	return as.Spelling()
}
