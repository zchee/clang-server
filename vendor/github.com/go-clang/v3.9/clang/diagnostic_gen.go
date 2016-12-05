package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// A single diagnostic, containing the diagnostic's severity, location, text, source ranges, and fix-it hints.
type Diagnostic struct {
	c C.CXDiagnostic
}

/*
	Retrieve the child diagnostics of a CXDiagnostic.

	This CXDiagnosticSet does not need to be released by
	clang_disposeDiagnosticSet.
*/
func (d Diagnostic) ChildDiagnostics() DiagnosticSet {
	return DiagnosticSet{C.clang_getChildDiagnostics(d.c)}
}

// Destroy a diagnostic.
func (d Diagnostic) Dispose() {
	C.clang_disposeDiagnostic(d.c)
}

/*
	Format the given diagnostic in a manner that is suitable for display.

	This routine will format the given diagnostic to a string, rendering
	the diagnostic according to the various options given. The
	clang_defaultDiagnosticDisplayOptions() function returns the set of
	options that most closely mimics the behavior of the clang compiler.

	Parameter Diagnostic The diagnostic to print.

	Parameter Options A set of options that control the diagnostic display,
	created by combining CXDiagnosticDisplayOptions values.

	Returns A new string containing for formatted diagnostic.
*/
func (d Diagnostic) FormatDiagnostic(options uint32) string {
	o := cxstring{C.clang_formatDiagnostic(d.c, C.uint(options))}
	defer o.Dispose()

	return o.String()
}

// Determine the severity of the given diagnostic.
func (d Diagnostic) Severity() DiagnosticSeverity {
	return DiagnosticSeverity(C.clang_getDiagnosticSeverity(d.c))
}

/*
	Retrieve the source location of the given diagnostic.

	This location is where Clang would print the caret ('^') when
	displaying the diagnostic on the command line.
*/
func (d Diagnostic) Location() SourceLocation {
	return SourceLocation{C.clang_getDiagnosticLocation(d.c)}
}

// Retrieve the text of the given diagnostic.
func (d Diagnostic) Spelling() string {
	o := cxstring{C.clang_getDiagnosticSpelling(d.c)}
	defer o.Dispose()

	return o.String()
}

/*
	Retrieve the name of the command-line option that enabled this
	diagnostic.

	Parameter Diag The diagnostic to be queried.

	Parameter Disable If non-NULL, will be set to the option that disables this
	diagnostic (if any).

	Returns A string that contains the command-line option used to enable this
	warning, such as "-Wconversion" or "-pedantic".
*/
func (d Diagnostic) Option() (string, string) {
	var disable cxstring
	defer disable.Dispose()

	o := cxstring{C.clang_getDiagnosticOption(d.c, &disable.c)}
	defer o.Dispose()

	return disable.String(), o.String()
}

/*
	Retrieve the category number for this diagnostic.

	Diagnostics can be categorized into groups along with other, related
	diagnostics (e.g., diagnostics under the same warning flag). This routine
	retrieves the category number for the given diagnostic.

	Returns The number of the category that contains this diagnostic, or zero
	if this diagnostic is uncategorized.
*/
func (d Diagnostic) Category() uint32 {
	return uint32(C.clang_getDiagnosticCategory(d.c))
}

/*
	Retrieve the diagnostic category text for a given diagnostic.

	Returns The text of the given diagnostic category.
*/
func (d Diagnostic) CategoryText() string {
	o := cxstring{C.clang_getDiagnosticCategoryText(d.c)}
	defer o.Dispose()

	return o.String()
}

// Determine the number of source ranges associated with the given diagnostic.
func (d Diagnostic) NumRanges() uint32 {
	return uint32(C.clang_getDiagnosticNumRanges(d.c))
}

/*
	Retrieve a source range associated with the diagnostic.

	A diagnostic's source ranges highlight important elements in the source
	code. On the command line, Clang displays source ranges by
	underlining them with '~' characters.

	Parameter Diagnostic the diagnostic whose range is being extracted.

	Parameter Range the zero-based index specifying which range to

	Returns the requested source range.
*/
func (d Diagnostic) Range(r uint32) SourceRange {
	return SourceRange{C.clang_getDiagnosticRange(d.c, C.uint(r))}
}

// Determine the number of fix-it hints associated with the given diagnostic.
func (d Diagnostic) NumFixIts() uint32 {
	return uint32(C.clang_getDiagnosticNumFixIts(d.c))
}

/*
	Retrieve the replacement information for a given fix-it.

	Fix-its are described in terms of a source range whose contents
	should be replaced by a string. This approach generalizes over
	three kinds of operations: removal of source code (the range covers
	the code to be removed and the replacement string is empty),
	replacement of source code (the range covers the code to be
	replaced and the replacement string provides the new code), and
	insertion (both the start and end of the range point at the
	insertion location, and the replacement string provides the text to
	insert).

	Parameter Diagnostic The diagnostic whose fix-its are being queried.

	Parameter FixIt The zero-based index of the fix-it.

	Parameter ReplacementRange The source range whose contents will be
	replaced with the returned replacement string. Note that source
	ranges are half-open ranges [a, b), so the source code should be
	replaced from a and up to (but not including) b.

	Returns A string containing text that should be replace the source
	code indicated by the ReplacementRange.
*/
func (d Diagnostic) FixIt(fixIt uint32) (SourceRange, string) {
	var replacementRange SourceRange

	o := cxstring{C.clang_getDiagnosticFixIt(d.c, C.uint(fixIt), &replacementRange.c)}
	defer o.Dispose()

	return replacementRange, o.String()
}
