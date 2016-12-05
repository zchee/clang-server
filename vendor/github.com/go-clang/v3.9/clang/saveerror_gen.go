package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Describes the kind of error that occurred (if any) in a call to clang_saveTranslationUnit().
type SaveError int32

const (
	// Indicates that no error occurred while saving a translation unit.
	SaveError_None SaveError = C.CXSaveError_None
	/*
		Indicates that an unknown error occurred while attempting to save
		the file.

		This error typically indicates that file I/O failed when attempting to
		write the file.
	*/
	SaveError_Unknown = C.CXSaveError_Unknown
	/*
		Indicates that errors during translation prevented this attempt
		to save the translation unit.

		Errors that prevent the translation unit from being saved can be
		extracted using clang_getNumDiagnostics() and clang_getDiagnostic().
	*/
	SaveError_TranslationErrors = C.CXSaveError_TranslationErrors
	// Indicates that the translation unit to be saved was somehow invalid (e.g., NULL).
	SaveError_InvalidTU = C.CXSaveError_InvalidTU
)

func (se SaveError) Spelling() string {
	switch se {
	case SaveError_None:
		return "SaveError=None"
	case SaveError_Unknown:
		return "SaveError=Unknown"
	case SaveError_TranslationErrors:
		return "SaveError=TranslationErrors"
	case SaveError_InvalidTU:
		return "SaveError=InvalidTU"
	}

	return fmt.Sprintf("SaveError unkown %d", int(se))
}

func (se SaveError) String() string {
	return se.Spelling()
}

func (se SaveError) Error() string {
	return se.Spelling()
}
