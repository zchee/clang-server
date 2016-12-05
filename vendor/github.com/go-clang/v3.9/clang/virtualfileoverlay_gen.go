package clang

// #include "./clang-c/BuildSystem.h"
// #include "go-clang.h"
import "C"
import "unsafe"

// Object encapsulating information about overlaying virtual file/directories over the real file system.
type VirtualFileOverlay struct {
	c C.CXVirtualFileOverlay
}

/*
	Create a CXVirtualFileOverlay object.
	Must be disposed with clang_VirtualFileOverlay_dispose().

	Parameter options is reserved, always pass 0.
*/
func NewVirtualFileOverlay(options uint32) VirtualFileOverlay {
	return VirtualFileOverlay{C.clang_VirtualFileOverlay_create(C.uint(options))}
}

// Map an absolute virtual file path to an absolute real one. The virtual path must be canonicalized (not contain "."/".."). Returns 0 for success, non-zero to indicate an error.
func (vfo VirtualFileOverlay) AddFileMapping(virtualPath string, realPath string) ErrorCode {
	c_virtualPath := C.CString(virtualPath)
	defer C.free(unsafe.Pointer(c_virtualPath))
	c_realPath := C.CString(realPath)
	defer C.free(unsafe.Pointer(c_realPath))

	return ErrorCode(C.clang_VirtualFileOverlay_addFileMapping(vfo.c, c_virtualPath, c_realPath))
}

// Set the case sensitivity for the CXVirtualFileOverlay object. The CXVirtualFileOverlay object is case-sensitive by default, this option can be used to override the default. Returns 0 for success, non-zero to indicate an error.
func (vfo VirtualFileOverlay) SetCaseSensitivity(caseSensitive int32) ErrorCode {
	return ErrorCode(C.clang_VirtualFileOverlay_setCaseSensitivity(vfo.c, C.int(caseSensitive)))
}

/*
	Write out the CXVirtualFileOverlay object to a char buffer.

	Parameter options is reserved, always pass 0.
	Parameter out_buffer_ptr pointer to receive the buffer pointer, which should be
	disposed using clang_free().
	Parameter out_buffer_size pointer to receive the buffer size.
	Returns 0 for success, non-zero to indicate an error.
*/
func (vfo VirtualFileOverlay) WriteToBuffer(options uint32) (string, uint32, ErrorCode) {
	var outBufferPtr *C.char
	defer C.free(unsafe.Pointer(outBufferPtr))
	var outBufferSize C.uint

	o := ErrorCode(C.clang_VirtualFileOverlay_writeToBuffer(vfo.c, C.uint(options), &outBufferPtr, &outBufferSize))

	return C.GoString(outBufferPtr), uint32(outBufferSize), o
}

// Dispose a CXVirtualFileOverlay object.
func (vfo VirtualFileOverlay) Dispose() {
	C.clang_VirtualFileOverlay_dispose(vfo.c)
}
