package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "unsafe"

// A group of CXDiagnostics.
type DiagnosticSet struct {
	c C.CXDiagnosticSet
}

// Determine the number of diagnostics in a CXDiagnosticSet.
func (ds DiagnosticSet) NumDiagnosticsInSet() uint32 {
	return uint32(C.clang_getNumDiagnosticsInSet(ds.c))
}

/*
	Retrieve a diagnostic associated with the given CXDiagnosticSet.

	Parameter Diags the CXDiagnosticSet to query.
	Parameter Index the zero-based diagnostic number to retrieve.

	Returns the requested diagnostic. This diagnostic must be freed
	via a call to clang_disposeDiagnostic().
*/
func (ds DiagnosticSet) DiagnosticInSet(index uint32) Diagnostic {
	return Diagnostic{C.clang_getDiagnosticInSet(ds.c, C.uint(index))}
}

/*
	Deserialize a set of diagnostics from a Clang diagnostics bitcode
	file.

	Parameter file The name of the file to deserialize.
	Parameter error A pointer to a enum value recording if there was a problem
	deserializing the diagnostics.
	Parameter errorString A pointer to a CXString for recording the error string
	if the file was not successfully loaded.

	Returns A loaded CXDiagnosticSet if successful, and NULL otherwise. These
	diagnostics should be released using clang_disposeDiagnosticSet().
*/
func LoadDiagnostics(file string) (LoadDiag_Error, string, DiagnosticSet) {
	var error C.enum_CXLoadDiag_Error
	var errorString cxstring
	defer errorString.Dispose()

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	o := DiagnosticSet{C.clang_loadDiagnostics(c_file, &error, &errorString.c)}

	return LoadDiag_Error(error), errorString.String(), o
}

// Release a CXDiagnosticSet and all of its contained diagnostics.
func (ds DiagnosticSet) Dispose() {
	C.clang_disposeDiagnosticSet(ds.c)
}
