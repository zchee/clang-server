package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

/*
	Flags that can be passed to clang_codeCompleteAt() to
	modify its behavior.

	The enumerators in this enumeration can be bitwise-OR'd together to
	provide multiple options to clang_codeCompleteAt().
*/
type CodeComplete_Flags uint32

const (
	// Whether to include macros within the set of code completions returned.
	CodeComplete_IncludeMacros CodeComplete_Flags = C.CXCodeComplete_IncludeMacros
	// Whether to include code patterns for language constructs within the set of code completions, e.g., for loops.
	CodeComplete_IncludeCodePatterns = C.CXCodeComplete_IncludeCodePatterns
	// Whether to include brief documentation within the set of code completions returned.
	CodeComplete_IncludeBriefComments = C.CXCodeComplete_IncludeBriefComments
)

func (ccf CodeComplete_Flags) Spelling() string {
	switch ccf {
	case CodeComplete_IncludeMacros:
		return "CodeComplete=IncludeMacros"
	case CodeComplete_IncludeCodePatterns:
		return "CodeComplete=IncludeCodePatterns"
	case CodeComplete_IncludeBriefComments:
		return "CodeComplete=IncludeBriefComments"
	}

	return fmt.Sprintf("CodeComplete_Flags unkown %d", int(ccf))
}

func (ccf CodeComplete_Flags) String() string {
	return ccf.Spelling()
}
