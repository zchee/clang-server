package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Describes a kind of token.
type TokenKind uint32

const (
	// A token that contains some kind of punctuation.
	Token_Punctuation TokenKind = C.CXToken_Punctuation
	// A language keyword.
	Token_Keyword = C.CXToken_Keyword
	// An identifier (that is not a keyword).
	Token_Identifier = C.CXToken_Identifier
	// A numeric, string, or character literal.
	Token_Literal = C.CXToken_Literal
	// A comment.
	Token_Comment = C.CXToken_Comment
)

func (tk TokenKind) Spelling() string {
	switch tk {
	case Token_Punctuation:
		return "Token=Punctuation"
	case Token_Keyword:
		return "Token=Keyword"
	case Token_Identifier:
		return "Token=Identifier"
	case Token_Literal:
		return "Token=Literal"
	case Token_Comment:
		return "Token=Comment"
	}

	return fmt.Sprintf("TokenKind unkown %d", int(tk))
}

func (tk TokenKind) String() string {
	return tk.Spelling()
}
