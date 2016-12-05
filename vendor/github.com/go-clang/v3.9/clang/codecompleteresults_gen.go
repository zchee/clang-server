package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

/*
	Contains the results of code-completion.

	This data structure contains the results of code completion, as
	produced by clang_codeCompleteAt(). Its contents must be freed by
	clang_disposeCodeCompleteResults.
*/
type CodeCompleteResults struct {
	c *C.CXCodeCompleteResults
}

// Free the given set of code-completion results.
func (ccr *CodeCompleteResults) Dispose() {
	C.clang_disposeCodeCompleteResults(ccr.c)
}

// Determine the number of diagnostics produced prior to the location where code completion was performed.
func (ccr *CodeCompleteResults) NumDiagnostics() uint32 {
	return uint32(C.clang_codeCompleteGetNumDiagnostics(ccr.c))
}

/*
	Retrieve a diagnostic associated with the given code completion.

	Parameter Results the code completion results to query.
	Parameter Index the zero-based diagnostic number to retrieve.

	Returns the requested diagnostic. This diagnostic must be freed
	via a call to clang_disposeDiagnostic().
*/
func (ccr *CodeCompleteResults) Diagnostic(index uint32) Diagnostic {
	return Diagnostic{C.clang_codeCompleteGetDiagnostic(ccr.c, C.uint(index))}
}

/*
	Determines what completions are appropriate for the context
	the given code completion.

	Parameter Results the code completion results to query

	Returns the kinds of completions that are appropriate for use
	along with the given code completion results.
*/
func (ccr *CodeCompleteResults) Contexts() uint64 {
	return uint64(C.clang_codeCompleteGetContexts(ccr.c))
}

/*
	Returns the cursor kind for the container for the current code
	completion context. The container is only guaranteed to be set for
	contexts where a container exists (i.e. member accesses or Objective-C
	message sends); if there is not a container, this function will return
	CXCursor_InvalidCode.

	Parameter Results the code completion results to query

	Parameter IsIncomplete on return, this value will be false if Clang has complete
	information about the container. If Clang does not have complete
	information, this value will be true.

	Returns the container kind, or CXCursor_InvalidCode if there is not a
	container
*/
func (ccr *CodeCompleteResults) ContainerKind() (uint32, CursorKind) {
	var isIncomplete C.uint

	o := CursorKind(C.clang_codeCompleteGetContainerKind(ccr.c, &isIncomplete))

	return uint32(isIncomplete), o
}

/*
	Returns the USR for the container for the current code completion
	context. If there is not a container for the current context, this
	function will return the empty string.

	Parameter Results the code completion results to query

	Returns the USR for the container
*/
func (ccr *CodeCompleteResults) ContainerUSR() string {
	o := cxstring{C.clang_codeCompleteGetContainerUSR(ccr.c)}
	defer o.Dispose()

	return o.String()
}

/*
	Returns the currently-entered selector for an Objective-C message
	send, formatted like "initWithFoo:bar:". Only guaranteed to return a
	non-empty string for CXCompletionContext_ObjCInstanceMessage and
	CXCompletionContext_ObjCClassMessage.

	Parameter Results the code completion results to query

	Returns the selector (or partial selector) that has been entered thus far
	for an Objective-C message send.
*/
func (ccr *CodeCompleteResults) Selector() string {
	o := cxstring{C.clang_codeCompleteGetObjCSelector(ccr.c)}
	defer o.Dispose()

	return o.String()
}

// The code-completion results.
func (ccr CodeCompleteResults) Results() []CompletionResult {
	var s []CompletionResult
	gos_s := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	gos_s.Cap = int(ccr.c.NumResults)
	gos_s.Len = int(ccr.c.NumResults)
	gos_s.Data = uintptr(unsafe.Pointer(ccr.c.Results))

	return s
}

// The number of code-completion results stored in the Results array.
func (ccr CodeCompleteResults) NumResults() uint32 {
	return uint32(ccr.c.NumResults)
}
