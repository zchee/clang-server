package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

// A single translation unit, which resides in an index.
type TranslationUnit struct {
	c C.CXTranslationUnit
}

// Determine whether the given header is guarded against multiple inclusions, either with the conventional \#ifndef/\#define/\#endif macro guards or with \#pragma once.
func (tu TranslationUnit) IsFileMultipleIncludeGuarded(file File) bool {
	o := C.clang_isFileMultipleIncludeGuarded(tu.c, file.c)

	return o != C.uint(0)
}

/*
	Retrieve a file handle within the given translation unit.

	Parameter tu the translation unit

	Parameter file_name the name of the file.

	Returns the file handle for the named file in the translation unit \p tu,
	or a NULL file handle if the file was not a part of this translation unit.
*/
func (tu TranslationUnit) File(fileName string) File {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	return File{C.clang_getFile(tu.c, c_fileName)}
}

// Retrieves the source location associated with a given file/line/column in a particular translation unit.
func (tu TranslationUnit) Location(file File, line uint32, column uint32) SourceLocation {
	return SourceLocation{C.clang_getLocation(tu.c, file.c, C.uint(line), C.uint(column))}
}

// Retrieves the source location associated with a given character offset in a particular translation unit.
func (tu TranslationUnit) LocationForOffset(file File, offset uint32) SourceLocation {
	return SourceLocation{C.clang_getLocationForOffset(tu.c, file.c, C.uint(offset))}
}

/*
	Retrieve all ranges that were skipped by the preprocessor.

	The preprocessor will skip lines when they are surrounded by an
	if/ifdef/ifndef directive whose condition does not evaluate to true.
*/
func (tu TranslationUnit) SkippedRanges(file File) *SourceRangeList {
	o := C.clang_getSkippedRanges(tu.c, file.c)

	var gop_o *SourceRangeList
	if o != nil {
		gop_o = &SourceRangeList{*o}
	}

	return gop_o
}

// Determine the number of diagnostics produced for the given translation unit.
func (tu TranslationUnit) NumDiagnostics() uint32 {
	return uint32(C.clang_getNumDiagnostics(tu.c))
}

/*
	Retrieve a diagnostic associated with the given translation unit.

	Parameter Unit the translation unit to query.
	Parameter Index the zero-based diagnostic number to retrieve.

	Returns the requested diagnostic. This diagnostic must be freed
	via a call to clang_disposeDiagnostic().
*/
func (tu TranslationUnit) Diagnostic(index uint32) Diagnostic {
	return Diagnostic{C.clang_getDiagnostic(tu.c, C.uint(index))}
}

/*
	Retrieve the complete set of diagnostics associated with a
	translation unit.

	Parameter Unit the translation unit to query.
*/
func (tu TranslationUnit) DiagnosticSetFromTU() DiagnosticSet {
	return DiagnosticSet{C.clang_getDiagnosticSetFromTU(tu.c)}
}

// Get the original translation unit source file name.
func (tu TranslationUnit) Spelling() string {
	o := cxstring{C.clang_getTranslationUnitSpelling(tu.c)}
	defer o.Dispose()

	return o.String()
}

/*
	Returns the set of flags that is suitable for saving a translation
	unit.

	The set of flags returned provide options for
	clang_saveTranslationUnit() by default. The returned flag
	set contains an unspecified set of options that save translation units with
	the most commonly-requested data.
*/
func (tu TranslationUnit) DefaultSaveOptions() uint32 {
	return uint32(C.clang_defaultSaveOptions(tu.c))
}

/*
	Saves a translation unit into a serialized representation of
	that translation unit on disk.

	Any translation unit that was parsed without error can be saved
	into a file. The translation unit can then be deserialized into a
	new CXTranslationUnit with clang_createTranslationUnit() or,
	if it is an incomplete translation unit that corresponds to a
	header, used as a precompiled header when parsing other translation
	units.

	Parameter TU The translation unit to save.

	Parameter FileName The file to which the translation unit will be saved.

	Parameter options A bitmask of options that affects how the translation unit
	is saved. This should be a bitwise OR of the
	CXSaveTranslationUnit_XXX flags.

	Returns A value that will match one of the enumerators of the CXSaveError
	enumeration. Zero (CXSaveError_None) indicates that the translation unit was
	saved successfully, while a non-zero value indicates that a problem occurred.
*/
func (tu TranslationUnit) SaveTranslationUnit(fileName string, options uint32) int32 {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	return int32(C.clang_saveTranslationUnit(tu.c, c_fileName, C.uint(options)))
}

// Destroy the specified CXTranslationUnit object.
func (tu TranslationUnit) Dispose() {
	C.clang_disposeTranslationUnit(tu.c)
}

/*
	Returns the set of flags that is suitable for reparsing a translation
	unit.

	The set of flags returned provide options for
	clang_reparseTranslationUnit() by default. The returned flag
	set contains an unspecified set of optimizations geared toward common uses
	of reparsing. The set of optimizations enabled may change from one version
	to the next.
*/
func (tu TranslationUnit) DefaultReparseOptions() uint32 {
	return uint32(C.clang_defaultReparseOptions(tu.c))
}

/*
	Reparse the source files that produced this translation unit.

	This routine can be used to re-parse the source files that originally
	created the given translation unit, for example because those source files
	have changed (either on disk or as passed via \p unsaved_files). The
	source code will be reparsed with the same command-line options as it
	was originally parsed.

	Reparsing a translation unit invalidates all cursors and source locations
	that refer into that translation unit. This makes reparsing a translation
	unit semantically equivalent to destroying the translation unit and then
	creating a new translation unit with the same command-line arguments.
	However, it may be more efficient to reparse a translation
	unit using this routine.

	Parameter TU The translation unit whose contents will be re-parsed. The
	translation unit must originally have been built with
	clang_createTranslationUnitFromSourceFile().

	Parameter num_unsaved_files The number of unsaved file entries in \p
	unsaved_files.

	Parameter unsaved_files The files that have not yet been saved to disk
	but may be required for parsing, including the contents of
	those files. The contents and name of these files (as specified by
	CXUnsavedFile) are copied when necessary, so the client only needs to
	guarantee their validity until the call to this function returns.

	Parameter options A bitset of options composed of the flags in CXReparse_Flags.
	The function clang_defaultReparseOptions() produces a default set of
	options recommended for most uses, based on the translation unit.

	Returns 0 if the sources could be reparsed. A non-zero error code will be
	returned if reparsing was impossible, such that the translation unit is
	invalid. In such cases, the only valid call for TU is
	clang_disposeTranslationUnit(TU). The error codes returned by this
	routine are described by the CXErrorCode enum.
*/
func (tu TranslationUnit) ReparseTranslationUnit(unsavedFiles []UnsavedFile, options uint32) int32 {
	gos_unsavedFiles := (*reflect.SliceHeader)(unsafe.Pointer(&unsavedFiles))
	cp_unsavedFiles := (*C.struct_CXUnsavedFile)(unsafe.Pointer(gos_unsavedFiles.Data))

	return int32(C.clang_reparseTranslationUnit(tu.c, C.uint(len(unsavedFiles)), cp_unsavedFiles, C.uint(options)))
}

// Return the memory usage of a translation unit. This object should be released with clang_disposeCXTUResourceUsage().
func (tu TranslationUnit) TUResourceUsage() TUResourceUsage {
	return TUResourceUsage{C.clang_getCXTUResourceUsage(tu.c)}
}

/*
	Retrieve the cursor that represents the given translation unit.

	The translation unit cursor can be used to start traversing the
	various declarations within the given translation unit.
*/
func (tu TranslationUnit) TranslationUnitCursor() Cursor {
	return Cursor{C.clang_getTranslationUnitCursor(tu.c)}
}

/*
	Map a source location to the cursor that describes the entity at that
	location in the source code.

	clang_getCursor() maps an arbitrary source location within a translation
	unit down to the most specific cursor that describes the entity at that
	location. For example, given an expression x + y, invoking
	clang_getCursor() with a source location pointing to "x" will return the
	cursor for "x"; similarly for "y". If the cursor points anywhere between
	"x" or "y" (e.g., on the + or the whitespace around it), clang_getCursor()
	will return a cursor referring to the "+" expression.

	Returns a cursor representing the entity at the given source location, or
	a NULL cursor if no such entity can be found.
*/
func (tu TranslationUnit) Cursor(sl SourceLocation) Cursor {
	return Cursor{C.clang_getCursor(tu.c, sl.c)}
}

// Given a CXFile header file, return the module that contains it, if one exists.
func (tu TranslationUnit) ModuleForFile(f File) Module {
	return Module{C.clang_getModuleForFile(tu.c, f.c)}
}

/*
	Parameter Module a module object.

	Returns the number of top level headers associated with this module.
*/
func (tu TranslationUnit) Module_getNumTopLevelHeaders(module Module) uint32 {
	return uint32(C.clang_Module_getNumTopLevelHeaders(tu.c, module.c))
}

/*
	Parameter Module a module object.

	Parameter Index top level header index (zero-based).

	Returns the specified top level header associated with the module.
*/
func (tu TranslationUnit) Module_getTopLevelHeader(module Module, index uint32) File {
	return File{C.clang_Module_getTopLevelHeader(tu.c, module.c, C.uint(index))}
}

/*
	Determine the spelling of the given token.

	The spelling of a token is the textual representation of that token, e.g.,
	the text of an identifier or keyword.
*/
func (tu TranslationUnit) TokenSpelling(t Token) string {
	o := cxstring{C.clang_getTokenSpelling(tu.c, t.c)}
	defer o.Dispose()

	return o.String()
}

// Retrieve the source location of the given token.
func (tu TranslationUnit) TokenLocation(t Token) SourceLocation {
	return SourceLocation{C.clang_getTokenLocation(tu.c, t.c)}
}

// Retrieve a source range that covers the given token.
func (tu TranslationUnit) TokenExtent(t Token) SourceRange {
	return SourceRange{C.clang_getTokenExtent(tu.c, t.c)}
}

/*
	Tokenize the source code described by the given range into raw
	lexical tokens.

	Parameter TU the translation unit whose text is being tokenized.

	Parameter Range the source range in which text should be tokenized. All of the
	tokens produced by tokenization will fall within this source range,

	Parameter Tokens this pointer will be set to point to the array of tokens
	that occur within the given source range. The returned pointer must be
	freed with clang_disposeTokens() before the translation unit is destroyed.

	Parameter NumTokens will be set to the number of tokens in the *Tokens
	array.
*/
func (tu TranslationUnit) Tokenize(r SourceRange) []Token {
	var cp_tokens *C.CXToken
	var tokens []Token
	var numTokens C.uint

	C.clang_tokenize(tu.c, r.c, &cp_tokens, &numTokens)

	gos_tokens := (*reflect.SliceHeader)(unsafe.Pointer(&tokens))
	gos_tokens.Cap = int(numTokens)
	gos_tokens.Len = int(numTokens)
	gos_tokens.Data = uintptr(unsafe.Pointer(cp_tokens))

	return tokens
}

// Free the given set of tokens.
func (tu TranslationUnit) DisposeTokens(tokens []Token) {
	gos_tokens := (*reflect.SliceHeader)(unsafe.Pointer(&tokens))
	cp_tokens := (*C.CXToken)(unsafe.Pointer(gos_tokens.Data))

	C.clang_disposeTokens(tu.c, cp_tokens, C.uint(len(tokens)))
}

/*
	Perform code completion at a given location in a translation unit.

	This function performs code completion at a particular file, line, and
	column within source code, providing results that suggest potential
	code snippets based on the context of the completion. The basic model
	for code completion is that Clang will parse a complete source file,
	performing syntax checking up to the location where code-completion has
	been requested. At that point, a special code-completion token is passed
	to the parser, which recognizes this token and determines, based on the
	current location in the C/Objective-C/C++ grammar and the state of
	semantic analysis, what completions to provide. These completions are
	returned via a new CXCodeCompleteResults structure.

	Code completion itself is meant to be triggered by the client when the
	user types punctuation characters or whitespace, at which point the
	code-completion location will coincide with the cursor. For example, if p
	is a pointer, code-completion might be triggered after the "-" and then
	after the ">" in p->. When the code-completion location is afer the ">",
	the completion results will provide, e.g., the members of the struct that
	"p" points to. The client is responsible for placing the cursor at the
	beginning of the token currently being typed, then filtering the results
	based on the contents of the token. For example, when code-completing for
	the expression p->get, the client should provide the location just after
	the ">" (e.g., pointing at the "g") to this code-completion hook. Then, the
	client can filter the results based on the current token text ("get"), only
	showing those results that start with "get". The intent of this interface
	is to separate the relatively high-latency acquisition of code-completion
	results from the filtering of results on a per-character basis, which must
	have a lower latency.

	Parameter TU The translation unit in which code-completion should
	occur. The source files for this translation unit need not be
	completely up-to-date (and the contents of those source files may
	be overridden via \p unsaved_files). Cursors referring into the
	translation unit may be invalidated by this invocation.

	Parameter complete_filename The name of the source file where code
	completion should be performed. This filename may be any file
	included in the translation unit.

	Parameter complete_line The line at which code-completion should occur.

	Parameter complete_column The column at which code-completion should occur.
	Note that the column should point just after the syntactic construct that
	initiated code completion, and not in the middle of a lexical token.

	Parameter unsaved_files the Files that have not yet been saved to disk
	but may be required for parsing or code completion, including the
	contents of those files. The contents and name of these files (as
	specified by CXUnsavedFile) are copied when necessary, so the
	client only needs to guarantee their validity until the call to
	this function returns.

	Parameter num_unsaved_files The number of unsaved file entries in \p
	unsaved_files.

	Parameter options Extra options that control the behavior of code
	completion, expressed as a bitwise OR of the enumerators of the
	CXCodeComplete_Flags enumeration. The
	clang_defaultCodeCompleteOptions() function returns a default set
	of code-completion options.

	Returns If successful, a new CXCodeCompleteResults structure
	containing code-completion results, which should eventually be
	freed with clang_disposeCodeCompleteResults(). If code
	completion fails, returns NULL.
*/
func (tu TranslationUnit) CodeCompleteAt(completeFilename string, completeLine uint32, completeColumn uint32, unsavedFiles []UnsavedFile, options uint32) *CodeCompleteResults {
	gos_unsavedFiles := (*reflect.SliceHeader)(unsafe.Pointer(&unsavedFiles))
	cp_unsavedFiles := (*C.struct_CXUnsavedFile)(unsafe.Pointer(gos_unsavedFiles.Data))

	c_completeFilename := C.CString(completeFilename)
	defer C.free(unsafe.Pointer(c_completeFilename))

	o := C.clang_codeCompleteAt(tu.c, c_completeFilename, C.uint(completeLine), C.uint(completeColumn), cp_unsavedFiles, C.uint(len(unsavedFiles)), C.uint(options))

	var gop_o *CodeCompleteResults
	if o != nil {
		gop_o = &CodeCompleteResults{o}
	}

	return gop_o
}

/*
	Find #import/#include directives in a specific file.

	Parameter TU translation unit containing the file to query.

	Parameter file to search for #import/#include directives.

	Parameter visitor callback that will receive pairs of CXCursor/CXSourceRange for
	each directive found.

	Returns one of the CXResult enumerators.
*/
func (tu TranslationUnit) FindIncludesInFile(file File, visitor CursorAndRangeVisitor) Result {
	return Result(C.clang_findIncludesInFile(tu.c, file.c, visitor.c))
}
