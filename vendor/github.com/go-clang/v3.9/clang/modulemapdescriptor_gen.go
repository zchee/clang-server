package clang

// #include "./clang-c/BuildSystem.h"
// #include "go-clang.h"
import "C"
import "unsafe"

// Object encapsulating information about a module.map file.
type ModuleMapDescriptor struct {
	c C.CXModuleMapDescriptor
}

/*
	Create a CXModuleMapDescriptor object.
	Must be disposed with clang_ModuleMapDescriptor_dispose().

	Parameter options is reserved, always pass 0.
*/
func NewModuleMapDescriptor(options uint32) ModuleMapDescriptor {
	return ModuleMapDescriptor{C.clang_ModuleMapDescriptor_create(C.uint(options))}
}

// Sets the framework module name that the module.map describes. Returns 0 for success, non-zero to indicate an error.
func (mmd ModuleMapDescriptor) SetFrameworkModuleName(name string) ErrorCode {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return ErrorCode(C.clang_ModuleMapDescriptor_setFrameworkModuleName(mmd.c, c_name))
}

// Sets the umbrealla header name that the module.map describes. Returns 0 for success, non-zero to indicate an error.
func (mmd ModuleMapDescriptor) SetUmbrellaHeader(name string) ErrorCode {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return ErrorCode(C.clang_ModuleMapDescriptor_setUmbrellaHeader(mmd.c, c_name))
}

/*
	Write out the CXModuleMapDescriptor object to a char buffer.

	Parameter options is reserved, always pass 0.
	Parameter out_buffer_ptr pointer to receive the buffer pointer, which should be
	disposed using clang_free().
	Parameter out_buffer_size pointer to receive the buffer size.
	Returns 0 for success, non-zero to indicate an error.
*/
func (mmd ModuleMapDescriptor) WriteToBuffer(options uint32) (string, uint32, ErrorCode) {
	var outBufferPtr *C.char
	defer C.free(unsafe.Pointer(outBufferPtr))
	var outBufferSize C.uint

	o := ErrorCode(C.clang_ModuleMapDescriptor_writeToBuffer(mmd.c, C.uint(options), &outBufferPtr, &outBufferSize))

	return C.GoString(outBufferPtr), uint32(outBufferSize), o
}

// Dispose a CXModuleMapDescriptor object.
func (mmd ModuleMapDescriptor) Dispose() {
	C.clang_ModuleMapDescriptor_dispose(mmd.c)
}
