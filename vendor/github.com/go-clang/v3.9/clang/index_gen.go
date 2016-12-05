package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

// An "index" that consists of a set of translation units that would typically be linked together into an executable or library.
type Index struct {
	c C.CXIndex
}

/*
	Provides a shared context for creating translation units.

	It provides two options:

	- excludeDeclarationsFromPCH: When non-zero, allows enumeration of "local"
	declarations (when loading any new translation units). A "local" declaration
	is one that belongs in the translation unit itself and not in a precompiled
	header that was used by the translation unit. If zero, all declarations
	will be enumerated.

	Here is an example:

	\code
	// excludeDeclsFromPCH = 1, displayDiagnostics=1
	Idx = clang_createIndex(1, 1);

	// IndexTest.pch was produced with the following command:
	// "clang -x c IndexTest.h -emit-ast -o IndexTest.pch"
	TU = clang_createTranslationUnit(Idx, "IndexTest.pch");

	// This will load all the symbols from 'IndexTest.pch'
	clang_visitChildren(clang_getTranslationUnitCursor(TU),
	TranslationUnitVisitor, 0);
	clang_disposeTranslationUnit(TU);

	// This will load all the symbols from 'IndexTest.c', excluding symbols
	// from 'IndexTest.pch'.
	char *args[] = { "-Xclang", "-include-pch=IndexTest.pch" };
	TU = clang_createTranslationUnitFromSourceFile(Idx, "IndexTest.c", 2, args,
	0, 0);
	clang_visitChildren(clang_getTranslationUnitCursor(TU),
	TranslationUnitVisitor, 0);
	clang_disposeTranslationUnit(TU);
	\endcode

	This process of creating the 'pch', loading it separately, and using it (via
	-include-pch) allows 'excludeDeclsFromPCH' to remove redundant callbacks
	(which gives the indexer the same performance benefit as the compiler).
*/
func NewIndex(excludeDeclarationsFromPCH int32, displayDiagnostics int32) Index {
	return Index{C.clang_createIndex(C.int(excludeDeclarationsFromPCH), C.int(displayDiagnostics))}
}

/*
	Destroy the given index.

	The index must not be destroyed until all of the translation units created
	within that index have been destroyed.
*/
func (i Index) Dispose() {
	C.clang_disposeIndex(i.c)
}

/*
	Sets general options associated with a CXIndex.

	For example:
	\code
	CXIndex idx = ...;
	clang_CXIndex_setGlobalOptions(idx,
	clang_CXIndex_getGlobalOptions(idx) |
	CXGlobalOpt_ThreadBackgroundPriorityForIndexing);
	\endcode

	Parameter options A bitmask of options, a bitwise OR of CXGlobalOpt_XXX flags.
*/
func (i Index) SetGlobalOptions(options uint32) {
	C.clang_CXIndex_setGlobalOptions(i.c, C.uint(options))
}

/*
	Gets the general options associated with a CXIndex.

	Returns A bitmask of options, a bitwise OR of CXGlobalOpt_XXX flags that
	are associated with the given CXIndex object.
*/
func (i Index) GlobalOptions() uint32 {
	return uint32(C.clang_CXIndex_getGlobalOptions(i.c))
}

/*
	Return the CXTranslationUnit for a given source file and the provided
	command line arguments one would pass to the compiler.

	Note: The 'source_filename' argument is optional. If the caller provides a
	NULL pointer, the name of the source file is expected to reside in the
	specified command line arguments.

	Note: When encountered in 'clang_command_line_args', the following options
	are ignored:

	'-c'
	'-emit-ast'
	'-fsyntax-only'
	'-o \<output file>' (both '-o' and '\<output file>' are ignored)

	Parameter CIdx The index object with which the translation unit will be
	associated.

	Parameter source_filename The name of the source file to load, or NULL if the
	source file is included in \p clang_command_line_args.

	Parameter num_clang_command_line_args The number of command-line arguments in
	\p clang_command_line_args.

	Parameter clang_command_line_args The command-line arguments that would be
	passed to the clang executable if it were being invoked out-of-process.
	These command-line options will be parsed and will affect how the translation
	unit is parsed. Note that the following options are ignored: '-c',
	'-emit-ast', '-fsyntax-only' (which is the default), and '-o \<output file>'.

	Parameter num_unsaved_files the number of unsaved file entries in \p
	unsaved_files.

	Parameter unsaved_files the files that have not yet been saved to disk
	but may be required for code completion, including the contents of
	those files. The contents and name of these files (as specified by
	CXUnsavedFile) are copied when necessary, so the client only needs to
	guarantee their validity until the call to this function returns.
*/
func (i Index) TranslationUnitFromSourceFile(sourceFilename string, clangCommandLineArgs []string, unsavedFiles []UnsavedFile) TranslationUnit {
	ca_clangCommandLineArgs := make([]*C.char, len(clangCommandLineArgs))
	var cp_clangCommandLineArgs **C.char
	if len(clangCommandLineArgs) > 0 {
		cp_clangCommandLineArgs = &ca_clangCommandLineArgs[0]
	}
	for i := range clangCommandLineArgs {
		ci_str := C.CString(clangCommandLineArgs[i])
		defer C.free(unsafe.Pointer(ci_str))
		ca_clangCommandLineArgs[i] = ci_str
	}
	gos_unsavedFiles := (*reflect.SliceHeader)(unsafe.Pointer(&unsavedFiles))
	cp_unsavedFiles := (*C.struct_CXUnsavedFile)(unsafe.Pointer(gos_unsavedFiles.Data))

	c_sourceFilename := C.CString(sourceFilename)
	defer C.free(unsafe.Pointer(c_sourceFilename))

	return TranslationUnit{C.clang_createTranslationUnitFromSourceFile(i.c, c_sourceFilename, C.int(len(clangCommandLineArgs)), cp_clangCommandLineArgs, C.uint(len(unsavedFiles)), cp_unsavedFiles)}
}

// Same as clang_createTranslationUnit2, but returns the CXTranslationUnit instead of an error code. In case of an error this routine returns a NULL CXTranslationUnit, without further detailed error codes.
func (i Index) TranslationUnit(astFilename string) TranslationUnit {
	c_astFilename := C.CString(astFilename)
	defer C.free(unsafe.Pointer(c_astFilename))

	return TranslationUnit{C.clang_createTranslationUnit(i.c, c_astFilename)}
}

/*
	Create a translation unit from an AST file (-emit-ast).

	\param[out] out_TU A non-NULL pointer to store the created
	CXTranslationUnit.

	Returns Zero on success, otherwise returns an error code.
*/
func (i Index) TranslationUnit2(astFilename string, outTU *TranslationUnit) ErrorCode {
	c_astFilename := C.CString(astFilename)
	defer C.free(unsafe.Pointer(c_astFilename))

	return ErrorCode(C.clang_createTranslationUnit2(i.c, c_astFilename, &outTU.c))
}

// Same as clang_parseTranslationUnit2, but returns the CXTranslationUnit instead of an error code. In case of an error this routine returns a NULL CXTranslationUnit, without further detailed error codes.
func (i Index) ParseTranslationUnit(sourceFilename string, commandLineArgs []string, unsavedFiles []UnsavedFile, options uint32) TranslationUnit {
	ca_commandLineArgs := make([]*C.char, len(commandLineArgs))
	var cp_commandLineArgs **C.char
	if len(commandLineArgs) > 0 {
		cp_commandLineArgs = &ca_commandLineArgs[0]
	}
	for i := range commandLineArgs {
		ci_str := C.CString(commandLineArgs[i])
		defer C.free(unsafe.Pointer(ci_str))
		ca_commandLineArgs[i] = ci_str
	}
	gos_unsavedFiles := (*reflect.SliceHeader)(unsafe.Pointer(&unsavedFiles))
	cp_unsavedFiles := (*C.struct_CXUnsavedFile)(unsafe.Pointer(gos_unsavedFiles.Data))

	c_sourceFilename := C.CString(sourceFilename)
	defer C.free(unsafe.Pointer(c_sourceFilename))

	return TranslationUnit{C.clang_parseTranslationUnit(i.c, c_sourceFilename, cp_commandLineArgs, C.int(len(commandLineArgs)), cp_unsavedFiles, C.uint(len(unsavedFiles)), C.uint(options))}
}

/*
	Parse the given source file and the translation unit corresponding
	to that file.

	This routine is the main entry point for the Clang C API, providing the
	ability to parse a source file into a translation unit that can then be
	queried by other functions in the API. This routine accepts a set of
	command-line arguments so that the compilation can be configured in the same
	way that the compiler is configured on the command line.

	Parameter CIdx The index object with which the translation unit will be
	associated.

	Parameter source_filename The name of the source file to load, or NULL if the
	source file is included in command_line_args.

	Parameter command_line_args The command-line arguments that would be
	passed to the clang executable if it were being invoked out-of-process.
	These command-line options will be parsed and will affect how the translation
	unit is parsed. Note that the following options are ignored: '-c',
	'-emit-ast', '-fsyntax-only' (which is the default), and '-o \<output file>'.

	Parameter num_command_line_args The number of command-line arguments in
	command_line_args.

	Parameter unsaved_files the files that have not yet been saved to disk
	but may be required for parsing, including the contents of
	those files. The contents and name of these files (as specified by
	CXUnsavedFile) are copied when necessary, so the client only needs to
	guarantee their validity until the call to this function returns.

	Parameter num_unsaved_files the number of unsaved file entries in \p
	unsaved_files.

	Parameter options A bitmask of options that affects how the translation unit
	is managed but not its compilation. This should be a bitwise OR of the
	CXTranslationUnit_XXX flags.

	\param[out] out_TU A non-NULL pointer to store the created
	CXTranslationUnit, describing the parsed code and containing any
	diagnostics produced by the compiler.

	Returns Zero on success, otherwise returns an error code.
*/
func (i Index) ParseTranslationUnit2(sourceFilename string, commandLineArgs []string, unsavedFiles []UnsavedFile, options uint32, outTU *TranslationUnit) ErrorCode {
	ca_commandLineArgs := make([]*C.char, len(commandLineArgs))
	var cp_commandLineArgs **C.char
	if len(commandLineArgs) > 0 {
		cp_commandLineArgs = &ca_commandLineArgs[0]
	}
	for i := range commandLineArgs {
		ci_str := C.CString(commandLineArgs[i])
		defer C.free(unsafe.Pointer(ci_str))
		ca_commandLineArgs[i] = ci_str
	}
	gos_unsavedFiles := (*reflect.SliceHeader)(unsafe.Pointer(&unsavedFiles))
	cp_unsavedFiles := (*C.struct_CXUnsavedFile)(unsafe.Pointer(gos_unsavedFiles.Data))

	c_sourceFilename := C.CString(sourceFilename)
	defer C.free(unsafe.Pointer(c_sourceFilename))

	return ErrorCode(C.clang_parseTranslationUnit2(i.c, c_sourceFilename, cp_commandLineArgs, C.int(len(commandLineArgs)), cp_unsavedFiles, C.uint(len(unsavedFiles)), C.uint(options), &outTU.c))
}

// Same as clang_parseTranslationUnit2 but requires a full command line for command_line_args including argv[0]. This is useful if the standard library paths are relative to the binary.
func (i Index) ParseTranslationUnit2FullArgv(sourceFilename string, commandLineArgs []string, unsavedFiles []UnsavedFile, options uint32, outTU *TranslationUnit) ErrorCode {
	ca_commandLineArgs := make([]*C.char, len(commandLineArgs))
	var cp_commandLineArgs **C.char
	if len(commandLineArgs) > 0 {
		cp_commandLineArgs = &ca_commandLineArgs[0]
	}
	for i := range commandLineArgs {
		ci_str := C.CString(commandLineArgs[i])
		defer C.free(unsafe.Pointer(ci_str))
		ca_commandLineArgs[i] = ci_str
	}
	gos_unsavedFiles := (*reflect.SliceHeader)(unsafe.Pointer(&unsavedFiles))
	cp_unsavedFiles := (*C.struct_CXUnsavedFile)(unsafe.Pointer(gos_unsavedFiles.Data))

	c_sourceFilename := C.CString(sourceFilename)
	defer C.free(unsafe.Pointer(c_sourceFilename))

	return ErrorCode(C.clang_parseTranslationUnit2FullArgv(i.c, c_sourceFilename, cp_commandLineArgs, C.int(len(commandLineArgs)), cp_unsavedFiles, C.uint(len(unsavedFiles)), C.uint(options), &outTU.c))
}

/*
	An indexing action/session, to be applied to one or multiple
	translation units.

	Parameter CIdx The index object with which the index action will be associated.
*/
func (i Index) Action_create() IndexAction {
	return IndexAction{C.clang_IndexAction_create(i.c)}
}
