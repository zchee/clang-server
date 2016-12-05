package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Describes the severity of a particular diagnostic.
type DiagnosticSeverity uint32

const (
	// A diagnostic that has been suppressed, e.g., by a command-line option.
	Diagnostic_Ignored DiagnosticSeverity = C.CXDiagnostic_Ignored
	// This diagnostic is a note that should be attached to the previous (non-note) diagnostic.
	Diagnostic_Note = C.CXDiagnostic_Note
	// This diagnostic indicates suspicious code that may not be wrong.
	Diagnostic_Warning = C.CXDiagnostic_Warning
	// This diagnostic indicates that the code is ill-formed.
	Diagnostic_Error = C.CXDiagnostic_Error
	// This diagnostic indicates that the code is ill-formed such that future parser recovery is unlikely to produce useful results.
	Diagnostic_Fatal = C.CXDiagnostic_Fatal
)

func (ds DiagnosticSeverity) Spelling() string {
	switch ds {
	case Diagnostic_Ignored:
		return "Diagnostic=Ignored"
	case Diagnostic_Note:
		return "Diagnostic=Note"
	case Diagnostic_Warning:
		return "Diagnostic=Warning"
	case Diagnostic_Error:
		return "Diagnostic=Error"
	case Diagnostic_Fatal:
		return "Diagnostic=Fatal"
	}

	return fmt.Sprintf("DiagnosticSeverity unkown %d", int(ds))
}

func (ds DiagnosticSeverity) String() string {
	return ds.Spelling()
}
