package clang

// #include "go-clang.h"
import "C"

// Annotate the given set of tokens by providing cursors for each token that can be mapped to a specific entity within the abstract syntax tree. This token-annotation routine is equivalent to invoking clang_getCursor() for the source locations of each of the tokens. The cursors provided are filtered, so that only those cursors that have a direct correspondence to the token are accepted. For example, given a function call \c f(x), clang_getCursor() would provide the following cursors: * when the cursor is over the 'f', a DeclRefExpr cursor referring to 'f'. * when the cursor is over the '(' or the ')', a CallExpr referring to 'f'. * when the cursor is over the 'x', a DeclRefExpr cursor referring to 'x'. Only the first and last of these cursors will occur within the annotate, since the tokens "f" and "x' directly refer to a function and a variable, respectively, but the parentheses are just a small part of the full syntax of the function call expression, which is not provided as an annotation. \param TU the translation unit that owns the given tokens. \param Tokens the set of tokens to annotate. \param NumTokens the number of tokens in \p Tokens. \param Cursors an array of \p NumTokens cursors, whose contents will be replaced with the cursors corresponding to each token.
func (tu TranslationUnit) AnnotateTokens(Tokens []Token) []Cursor {
	ca_Tokens := make([]C.CXToken, len(Tokens))
	var cp_Tokens *C.CXToken
	if len(Tokens) > 0 {
		cp_Tokens = &ca_Tokens[0]
	}
	for i := range Tokens {
		ca_Tokens[i] = Tokens[i].c
	}
	ca_Cursors := make([]C.CXCursor, len(Tokens))
	var cp_Cursors *C.CXCursor
	if len(Tokens) > 0 {
		cp_Cursors = &ca_Cursors[0]
	}

	C.clang_annotateTokens(tu.c, cp_Tokens, C.uint(len(Tokens)), cp_Cursors)

	cursors := make([]Cursor, len(Tokens))
	for i := range ca_Cursors {
		cursors[i] = Cursor{ca_Cursors[i]}
	}

	return cursors
}

func (tu TranslationUnit) IsValid() bool {
	return tu.c != nil
}

func (tu TranslationUnit) Diagnostics() []Diagnostic { // TODO this can be generated https://github.com/go-clang/gen/issues/47
	s := make([]Diagnostic, tu.NumDiagnostics())

	for i := range s {
		s[i] = tu.Diagnostic(uint32(i))
	}

	return s
}
