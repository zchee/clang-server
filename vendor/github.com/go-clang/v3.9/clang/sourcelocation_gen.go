package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

/*
	Identifies a specific source location within a translation
	unit.

	Use clang_getExpansionLocation() or clang_getSpellingLocation()
	to map a source location to a particular file, line, and column.
*/
type SourceLocation struct {
	c C.CXSourceLocation
}

// Retrieve a NULL (invalid) source location.
func NewNullLocation() SourceLocation {
	return SourceLocation{C.clang_getNullLocation()}
}

/*
	Determine whether two source locations, which must refer into
	the same translation unit, refer to exactly the same point in the source
	code.

	Returns non-zero if the source locations refer to the same location, zero
	if they refer to different locations.
*/
func (sl SourceLocation) Equal(sl2 SourceLocation) bool {
	o := C.clang_equalLocations(sl.c, sl2.c)

	return o != C.uint(0)
}

// Returns non-zero if the given source location is in a system header.
func (sl SourceLocation) IsInSystemHeader() bool {
	o := C.clang_Location_isInSystemHeader(sl.c)

	return o != C.int(0)
}

// Returns non-zero if the given source location is in the main file of the corresponding translation unit.
func (sl SourceLocation) IsFromMainFile() bool {
	o := C.clang_Location_isFromMainFile(sl.c)

	return o != C.int(0)
}

// Retrieve a source range given the beginning and ending source locations.
func (sl SourceLocation) Range(end SourceLocation) SourceRange {
	return SourceRange{C.clang_getRange(sl.c, end.c)}
}

/*
	Retrieve the file, line, column, and offset represented by
	the given source location.

	If the location refers into a macro expansion, retrieves the
	location of the macro expansion.

	Parameter location the location within a source file that will be decomposed
	into its parts.

	Parameter file [out] if non-NULL, will be set to the file to which the given
	source location points.

	Parameter line [out] if non-NULL, will be set to the line to which the given
	source location points.

	Parameter column [out] if non-NULL, will be set to the column to which the given
	source location points.

	Parameter offset [out] if non-NULL, will be set to the offset into the
	buffer to which the given source location points.
*/
func (sl SourceLocation) ExpansionLocation() (File, uint32, uint32, uint32) {
	var file File
	var line C.uint
	var column C.uint
	var offset C.uint

	C.clang_getExpansionLocation(sl.c, &file.c, &line, &column, &offset)

	return file, uint32(line), uint32(column), uint32(offset)
}

/*
	Retrieve the file, line, column, and offset represented by
	the given source location, as specified in a # line directive.

	Example: given the following source code in a file somefile.c

	\code
	#123 "dummy.c" 1

	static int func(void)
	{
	return 0;
	}
	\endcode

	the location information returned by this function would be

	File: dummy.c Line: 124 Column: 12

	whereas clang_getExpansionLocation would have returned

	File: somefile.c Line: 3 Column: 12

	Parameter location the location within a source file that will be decomposed
	into its parts.

	Parameter filename [out] if non-NULL, will be set to the filename of the
	source location. Note that filenames returned will be for "virtual" files,
	which don't necessarily exist on the machine running clang - e.g. when
	parsing preprocessed output obtained from a different environment. If
	a non-NULL value is passed in, remember to dispose of the returned value
	using clang_disposeString() once you've finished with it. For an invalid
	source location, an empty string is returned.

	Parameter line [out] if non-NULL, will be set to the line number of the
	source location. For an invalid source location, zero is returned.

	Parameter column [out] if non-NULL, will be set to the column number of the
	source location. For an invalid source location, zero is returned.
*/
func (sl SourceLocation) PresumedLocation() (string, uint32, uint32) {
	var filename cxstring
	defer filename.Dispose()
	var line C.uint
	var column C.uint

	C.clang_getPresumedLocation(sl.c, &filename.c, &line, &column)

	return filename.String(), uint32(line), uint32(column)
}

/*
	Legacy API to retrieve the file, line, column, and offset represented
	by the given source location.

	This interface has been replaced by the newer interface
	#clang_getExpansionLocation(). See that interface's documentation for
	details.
*/
func (sl SourceLocation) InstantiationLocation() (File, uint32, uint32, uint32) {
	var file File
	var line C.uint
	var column C.uint
	var offset C.uint

	C.clang_getInstantiationLocation(sl.c, &file.c, &line, &column, &offset)

	return file, uint32(line), uint32(column), uint32(offset)
}

/*
	Retrieve the file, line, column, and offset represented by
	the given source location.

	If the location refers into a macro instantiation, return where the
	location was originally spelled in the source file.

	Parameter location the location within a source file that will be decomposed
	into its parts.

	Parameter file [out] if non-NULL, will be set to the file to which the given
	source location points.

	Parameter line [out] if non-NULL, will be set to the line to which the given
	source location points.

	Parameter column [out] if non-NULL, will be set to the column to which the given
	source location points.

	Parameter offset [out] if non-NULL, will be set to the offset into the
	buffer to which the given source location points.
*/
func (sl SourceLocation) SpellingLocation() (File, uint32, uint32, uint32) {
	var file File
	var line C.uint
	var column C.uint
	var offset C.uint

	C.clang_getSpellingLocation(sl.c, &file.c, &line, &column, &offset)

	return file, uint32(line), uint32(column), uint32(offset)
}

/*
	Retrieve the file, line, column, and offset represented by
	the given source location.

	If the location refers into a macro expansion, return where the macro was
	expanded or where the macro argument was written, if the location points at
	a macro argument.

	Parameter location the location within a source file that will be decomposed
	into its parts.

	Parameter file [out] if non-NULL, will be set to the file to which the given
	source location points.

	Parameter line [out] if non-NULL, will be set to the line to which the given
	source location points.

	Parameter column [out] if non-NULL, will be set to the column to which the given
	source location points.

	Parameter offset [out] if non-NULL, will be set to the offset into the
	buffer to which the given source location points.
*/
func (sl SourceLocation) FileLocation() (File, uint32, uint32, uint32) {
	var file File
	var line C.uint
	var column C.uint
	var offset C.uint

	C.clang_getFileLocation(sl.c, &file.c, &line, &column, &offset)

	return file, uint32(line), uint32(column), uint32(offset)
}
