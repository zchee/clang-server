package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Describe the "language" of the entity referred to by a cursor.
type LanguageKind uint32

const (
	Language_Invalid   LanguageKind = C.CXLanguage_Invalid
	Language_C                      = C.CXLanguage_C
	Language_ObjC                   = C.CXLanguage_ObjC
	Language_CPlusPlus              = C.CXLanguage_CPlusPlus
)

func (lk LanguageKind) Spelling() string {
	switch lk {
	case Language_Invalid:
		return "Language=Invalid"
	case Language_C:
		return "Language=C"
	case Language_ObjC:
		return "Language=ObjC"
	case Language_CPlusPlus:
		return "Language=CPlusPlus"
	}

	return fmt.Sprintf("LanguageKind unkown %d", int(lk))
}

func (lk LanguageKind) String() string {
	return lk.Spelling()
}
