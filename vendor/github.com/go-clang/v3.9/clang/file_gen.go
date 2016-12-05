package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "time"

// A particular source file that is part of a translation unit.
type File struct {
	c C.CXFile
}

// Retrieve the complete file and path name of the given file.
func (f File) Name() string {
	o := cxstring{C.clang_getFileName(f.c)}
	defer o.Dispose()

	return o.String()
}

// Retrieve the last modification time of the given file.
func (f File) Time() time.Time {
	return time.Unix(int64(C.clang_getFileTime(f.c)), 0)
}

/*
	Retrieve the unique ID for the given file.

	Parameter file the file to get the ID for.
	Parameter outID stores the returned CXFileUniqueID.
	Returns If there was a failure getting the unique ID, returns non-zero,
	otherwise returns 0.
*/
func (f File) UniqueID() (FileUniqueID, int32) {
	var outID FileUniqueID

	o := int32(C.clang_getFileUniqueID(f.c, &outID.c))

	return outID, o
}

// Returns non-zero if the file1 and file2 point to the same file, or they are both NULL.
func (f File) IsEqual(file2 File) bool {
	o := C.clang_File_isEqual(f.c, file2.c)

	return o != C.int(0)
}
