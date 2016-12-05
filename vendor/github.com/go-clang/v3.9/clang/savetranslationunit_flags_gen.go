package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

/*
	Flags that control how translation units are saved.

	The enumerators in this enumeration type are meant to be bitwise
	ORed together to specify which options should be used when
	saving the translation unit.
*/
type SaveTranslationUnit_Flags uint32

const (
	// Used to indicate that no special saving options are needed.
	SaveTranslationUnit_None SaveTranslationUnit_Flags = C.CXSaveTranslationUnit_None
)

func (stuf SaveTranslationUnit_Flags) Spelling() string {
	switch stuf {
	case SaveTranslationUnit_None:
		return "SaveTranslationUnit=None"
	}

	return fmt.Sprintf("SaveTranslationUnit_Flags unkown %d", int(stuf))
}

func (stuf SaveTranslationUnit_Flags) String() string {
	return stuf.Spelling()
}
