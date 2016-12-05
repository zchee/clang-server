package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

/*
	Describes how the traversal of the children of a particular
	cursor should proceed after visiting a particular child cursor.

	A value of this enumeration type should be returned by each
	CXCursorVisitor to indicate how clang_visitChildren() proceed.
*/
type ChildVisitResult uint32

const (
	// Terminates the cursor traversal.
	ChildVisit_Break ChildVisitResult = C.CXChildVisit_Break
	// Continues the cursor traversal with the next sibling of the cursor just visited, without visiting its children.
	ChildVisit_Continue = C.CXChildVisit_Continue
	// Recursively traverse the children of this cursor, using the same visitor and client data.
	ChildVisit_Recurse = C.CXChildVisit_Recurse
)

func (cvr ChildVisitResult) Spelling() string {
	switch cvr {
	case ChildVisit_Break:
		return "ChildVisit=Break"
	case ChildVisit_Continue:
		return "ChildVisit=Continue"
	case ChildVisit_Recurse:
		return "ChildVisit=Recurse"
	}

	return fmt.Sprintf("ChildVisitResult unkown %d", int(cvr))
}

func (cvr ChildVisitResult) String() string {
	return cvr.Spelling()
}
