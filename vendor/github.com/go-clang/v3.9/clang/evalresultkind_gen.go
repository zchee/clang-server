package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type EvalResultKind uint32

const (
	Eval_Int            EvalResultKind = C.CXEval_Int
	Eval_Float                         = C.CXEval_Float
	Eval_ObjCStrLiteral                = C.CXEval_ObjCStrLiteral
	Eval_StrLiteral                    = C.CXEval_StrLiteral
	Eval_CFStr                         = C.CXEval_CFStr
	Eval_Other                         = C.CXEval_Other
	Eval_UnExposed                     = C.CXEval_UnExposed
)

func (erk EvalResultKind) Spelling() string {
	switch erk {
	case Eval_Int:
		return "Eval=Int"
	case Eval_Float:
		return "Eval=Float"
	case Eval_ObjCStrLiteral:
		return "Eval=ObjCStrLiteral"
	case Eval_StrLiteral:
		return "Eval=StrLiteral"
	case Eval_CFStr:
		return "Eval=CFStr"
	case Eval_Other:
		return "Eval=Other"
	case Eval_UnExposed:
		return "Eval=UnExposed"
	}

	return fmt.Sprintf("EvalResultKind unkown %d", int(erk))
}

func (erk EvalResultKind) String() string {
	return erk.Spelling()
}
