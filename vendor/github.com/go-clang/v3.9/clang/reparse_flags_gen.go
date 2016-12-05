package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

/*
	Flags that control the reparsing of translation units.

	The enumerators in this enumeration type are meant to be bitwise
	ORed together to specify which options should be used when
	reparsing the translation unit.
*/
type Reparse_Flags uint32

const (
	// Used to indicate that no special reparsing options are needed.
	Reparse_None Reparse_Flags = C.CXReparse_None
)

func (rf Reparse_Flags) Spelling() string {
	switch rf {
	case Reparse_None:
		return "Reparse=None"
	}

	return fmt.Sprintf("Reparse_Flags unkown %d", int(rf))
}

func (rf Reparse_Flags) String() string {
	return rf.Spelling()
}
