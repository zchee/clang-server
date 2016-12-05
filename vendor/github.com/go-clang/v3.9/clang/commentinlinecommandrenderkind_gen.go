package clang

// #include "./clang-c/Documentation.h"
// #include "go-clang.h"
import "C"
import "fmt"

// The most appropriate rendering mode for an inline command, chosen on command semantics in Doxygen.
type CommentInlineCommandRenderKind uint32

const (
	// Command argument should be rendered in a normal font.
	CommentInlineCommandRenderKind_Normal CommentInlineCommandRenderKind = C.CXCommentInlineCommandRenderKind_Normal
	// Command argument should be rendered in a bold font.
	CommentInlineCommandRenderKind_Bold = C.CXCommentInlineCommandRenderKind_Bold
	// Command argument should be rendered in a monospaced font.
	CommentInlineCommandRenderKind_Monospaced = C.CXCommentInlineCommandRenderKind_Monospaced
	// Command argument should be rendered emphasized (typically italic font).
	CommentInlineCommandRenderKind_Emphasized = C.CXCommentInlineCommandRenderKind_Emphasized
)

func (cicrk CommentInlineCommandRenderKind) Spelling() string {
	switch cicrk {
	case CommentInlineCommandRenderKind_Normal:
		return "CommentInlineCommandRenderKind=Normal"
	case CommentInlineCommandRenderKind_Bold:
		return "CommentInlineCommandRenderKind=Bold"
	case CommentInlineCommandRenderKind_Monospaced:
		return "CommentInlineCommandRenderKind=Monospaced"
	case CommentInlineCommandRenderKind_Emphasized:
		return "CommentInlineCommandRenderKind=Emphasized"
	}

	return fmt.Sprintf("CommentInlineCommandRenderKind unkown %d", int(cicrk))
}

func (cicrk CommentInlineCommandRenderKind) String() string {
	return cicrk.Spelling()
}
