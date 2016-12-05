package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type VisitorResult uint32

const (
	Visit_Break    VisitorResult = C.CXVisit_Break
	Visit_Continue               = C.CXVisit_Continue
)

func (vr VisitorResult) Spelling() string {
	switch vr {
	case Visit_Break:
		return "Visit=Break"
	case Visit_Continue:
		return "Visit=Continue"
	}

	return fmt.Sprintf("VisitorResult unkown %d", int(vr))
}

func (vr VisitorResult) String() string {
	return vr.Spelling()
}
