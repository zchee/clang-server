package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Evaluation result of a cursor
type EvalResult struct {
	c C.CXEvalResult
}

// Returns the kind of the evaluated result.
func (er EvalResult) Kind() EvalResultKind {
	return EvalResultKind(C.clang_EvalResult_getKind(er.c))
}

// Returns the evaluation result as integer if the kind is Int.
func (er EvalResult) AsInt() int32 {
	return int32(C.clang_EvalResult_getAsInt(er.c))
}

// Returns the evaluation result as double if the kind is double.
func (er EvalResult) AsDouble() float64 {
	return float64(C.clang_EvalResult_getAsDouble(er.c))
}

// Returns the evaluation result as a constant string if the kind is other than Int or float. User must not free this pointer, instead call clang_EvalResult_dispose on the CXEvalResult returned by clang_Cursor_Evaluate.
func (er EvalResult) AsStr() string {
	return C.GoString(C.clang_EvalResult_getAsStr(er.c))
}

// Disposes the created Eval memory.
func (er EvalResult) Dispose() {
	C.clang_EvalResult_dispose(er.c)
}
