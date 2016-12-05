package clang

// #include "./clang-c/CXCompilationDatabase.h"
// #include "go-clang.h"
import "C"

/*
	Contains the results of a search in the compilation database

	When searching for the compile command for a file, the compilation db can
	return several commands, as the file may have been compiled with
	different options in different places of the project. This choice of compile
	commands is wrapped in this opaque data structure. It must be freed by
	clang_CompileCommands_dispose.
*/
type CompileCommands struct {
	c C.CXCompileCommands
}

// Free the given CompileCommands
func (cc CompileCommands) Dispose() {
	C.clang_CompileCommands_dispose(cc.c)
}

// Get the number of CompileCommand we have for a file
func (cc CompileCommands) Size() uint32 {
	return uint32(C.clang_CompileCommands_getSize(cc.c))
}

/*
	Get the I'th CompileCommand for a file

	Note : 0 <= i < clang_CompileCommands_getSize(CXCompileCommands)
*/
func (cc CompileCommands) Command(i uint32) CompileCommand {
	return CompileCommand{C.clang_CompileCommands_getCommand(cc.c, C.uint(i))}
}
