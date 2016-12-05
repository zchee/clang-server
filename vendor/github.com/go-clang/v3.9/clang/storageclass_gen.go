package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Represents the storage classes as declared in the source. CX_SC_Invalid was added for the case that the passed cursor in not a declaration.
type StorageClass uint32

const (
	SC_Invalid              StorageClass = C.CX_SC_Invalid
	SC_None                              = C.CX_SC_None
	SC_Extern                            = C.CX_SC_Extern
	SC_Static                            = C.CX_SC_Static
	SC_PrivateExtern                     = C.CX_SC_PrivateExtern
	SC_OpenCLWorkGroupLocal              = C.CX_SC_OpenCLWorkGroupLocal
	SC_Auto                              = C.CX_SC_Auto
	SC_Register                          = C.CX_SC_Register
)

func (sc StorageClass) Spelling() string {
	switch sc {
	case SC_Invalid:
		return "SC=Invalid"
	case SC_None:
		return "SC=None"
	case SC_Extern:
		return "SC=Extern"
	case SC_Static:
		return "SC=Static"
	case SC_PrivateExtern:
		return "SC=PrivateExtern"
	case SC_OpenCLWorkGroupLocal:
		return "SC=OpenCLWorkGroupLocal"
	case SC_Auto:
		return "SC=Auto"
	case SC_Register:
		return "SC=Register"
	}

	return fmt.Sprintf("StorageClass unkown %d", int(sc))
}

func (sc StorageClass) String() string {
	return sc.Spelling()
}
