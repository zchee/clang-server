package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type IdxDeclInfoFlags uint32

const (
	IdxDeclFlag_Skipped IdxDeclInfoFlags = C.CXIdxDeclFlag_Skipped
)

func (idif IdxDeclInfoFlags) Spelling() string {
	switch idif {
	case IdxDeclFlag_Skipped:
		return "IdxDeclFlag=Skipped"
	}

	return fmt.Sprintf("IdxDeclInfoFlags unkown %d", int(idif))
}

func (idif IdxDeclInfoFlags) String() string {
	return idif.Spelling()
}
