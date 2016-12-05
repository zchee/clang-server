package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Data for ppIncludedFile callback.
type IdxIncludedFileInfo struct {
	c C.CXIdxIncludedFileInfo
}

// Location of '#' in the \#include/\#import directive.
func (iifi IdxIncludedFileInfo) HashLoc() IdxLoc {
	return IdxLoc{iifi.c.hashLoc}
}

// Filename as written in the \#include/\#import directive.
func (iifi IdxIncludedFileInfo) Filename() string {
	return C.GoString(iifi.c.filename)
}

// The actual file that the \#include/\#import directive resolved to.
func (iifi IdxIncludedFileInfo) File() File {
	return File{iifi.c.file}
}

func (iifi IdxIncludedFileInfo) IsImport() bool {
	o := iifi.c.isImport

	return o != C.int(0)
}

func (iifi IdxIncludedFileInfo) IsAngled() bool {
	o := iifi.c.isAngled

	return o != C.int(0)
}

// Non-zero if the directive was automatically turned into a module import.
func (iifi IdxIncludedFileInfo) IsModuleImport() bool {
	o := iifi.c.isModuleImport

	return o != C.int(0)
}
