package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "unsafe"

// A remapping of original source files and their translated files.
type Remapping struct {
	c C.CXRemapping
}

/*
	Retrieve a remapping.

	Parameter path the path that contains metadata about remappings.

	Returns the requested remapping. This remapping must be freed
	via a call to clang_remap_dispose(). Can return NULL if an error occurred.
*/
func NewRemappings(path string) Remapping {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	return Remapping{C.clang_getRemappings(c_path)}
}

/*
	Retrieve a remapping.

	Parameter filePaths pointer to an array of file paths containing remapping info.

	Parameter numFiles number of file paths.

	Returns the requested remapping. This remapping must be freed
	via a call to clang_remap_dispose(). Can return NULL if an error occurred.
*/
func NewRemappingsFromFileList(filePaths []string) Remapping {
	ca_filePaths := make([]*C.char, len(filePaths))
	var cp_filePaths **C.char
	if len(filePaths) > 0 {
		cp_filePaths = &ca_filePaths[0]
	}
	for i := range filePaths {
		ci_str := C.CString(filePaths[i])
		defer C.free(unsafe.Pointer(ci_str))
		ca_filePaths[i] = ci_str
	}

	return Remapping{C.clang_getRemappingsFromFileList(cp_filePaths, C.uint(len(filePaths)))}
}

// Determine the number of remappings.
func (r Remapping) NumFiles() uint32 {
	return uint32(C.clang_remap_getNumFiles(r.c))
}

/*
	Get the original and the associated filename from the remapping.

	Parameter original If non-NULL, will be set to the original filename.

	Parameter transformed If non-NULL, will be set to the filename that the original
	is associated with.
*/
func (r Remapping) Filenames(index uint32) (string, string) {
	var original cxstring
	defer original.Dispose()
	var transformed cxstring
	defer transformed.Dispose()

	C.clang_remap_getFilenames(r.c, C.uint(index), &original.c, &transformed.c)

	return original.String(), transformed.String()
}

// Dispose the remapping.
func (r Remapping) Dispose() {
	C.clang_remap_dispose(r.c)
}
