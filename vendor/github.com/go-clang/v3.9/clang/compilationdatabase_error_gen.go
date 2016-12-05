package clang

// #include "./clang-c/CXCompilationDatabase.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Error codes for Compilation Database
type CompilationDatabase_Error int32

const (
	CompilationDatabase_NoError            CompilationDatabase_Error = C.CXCompilationDatabase_NoError
	CompilationDatabase_CanNotLoadDatabase                           = C.CXCompilationDatabase_CanNotLoadDatabase
)

func (cde CompilationDatabase_Error) Spelling() string {
	switch cde {
	case CompilationDatabase_NoError:
		return "CompilationDatabase=NoError"
	case CompilationDatabase_CanNotLoadDatabase:
		return "CompilationDatabase=CanNotLoadDatabase"
	}

	return fmt.Sprintf("CompilationDatabase_Error unkown %d", int(cde))
}

func (cde CompilationDatabase_Error) String() string {
	return cde.Spelling()
}

func (cde CompilationDatabase_Error) Error() string {
	return cde.Spelling()
}
