package clang

// #include "./clang-c/Documentation.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Describes parameter passing direction for \Parameter or \\arg command.
type CommentParamPassDirection uint32

const (
	// The parameter is an input parameter.
	CommentParamPassDirection_In CommentParamPassDirection = C.CXCommentParamPassDirection_In
	// The parameter is an output parameter.
	CommentParamPassDirection_Out = C.CXCommentParamPassDirection_Out
	// The parameter is an input and output parameter.
	CommentParamPassDirection_InOut = C.CXCommentParamPassDirection_InOut
)

func (cppd CommentParamPassDirection) Spelling() string {
	switch cppd {
	case CommentParamPassDirection_In:
		return "CommentParamPassDirection=In"
	case CommentParamPassDirection_Out:
		return "CommentParamPassDirection=Out"
	case CommentParamPassDirection_InOut:
		return "CommentParamPassDirection=InOut"
	}

	return fmt.Sprintf("CommentParamPassDirection unkown %d", int(cppd))
}

func (cppd CommentParamPassDirection) String() string {
	return cppd.Spelling()
}
