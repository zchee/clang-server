package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Describes the calling convention of a function type
type CallingConv uint32

const (
	CallingConv_Default       CallingConv = C.CXCallingConv_Default
	CallingConv_C                         = C.CXCallingConv_C
	CallingConv_X86StdCall                = C.CXCallingConv_X86StdCall
	CallingConv_X86FastCall               = C.CXCallingConv_X86FastCall
	CallingConv_X86ThisCall               = C.CXCallingConv_X86ThisCall
	CallingConv_X86Pascal                 = C.CXCallingConv_X86Pascal
	CallingConv_AAPCS                     = C.CXCallingConv_AAPCS
	CallingConv_AAPCS_VFP                 = C.CXCallingConv_AAPCS_VFP
	CallingConv_IntelOclBicc              = C.CXCallingConv_IntelOclBicc
	CallingConv_X86_64Win64               = C.CXCallingConv_X86_64Win64
	CallingConv_X86_64SysV                = C.CXCallingConv_X86_64SysV
	CallingConv_X86VectorCall             = C.CXCallingConv_X86VectorCall
	CallingConv_Swift                     = C.CXCallingConv_Swift
	CallingConv_PreserveMost              = C.CXCallingConv_PreserveMost
	CallingConv_PreserveAll               = C.CXCallingConv_PreserveAll
	CallingConv_Invalid                   = C.CXCallingConv_Invalid
	CallingConv_Unexposed                 = C.CXCallingConv_Unexposed
)

func (cc CallingConv) Spelling() string {
	switch cc {
	case CallingConv_Default:
		return "CallingConv=Default"
	case CallingConv_C:
		return "CallingConv=C"
	case CallingConv_X86StdCall:
		return "CallingConv=X86StdCall"
	case CallingConv_X86FastCall:
		return "CallingConv=X86FastCall"
	case CallingConv_X86ThisCall:
		return "CallingConv=X86ThisCall"
	case CallingConv_X86Pascal:
		return "CallingConv=X86Pascal"
	case CallingConv_AAPCS:
		return "CallingConv=AAPCS"
	case CallingConv_AAPCS_VFP:
		return "CallingConv=AAPCS_VFP"
	case CallingConv_IntelOclBicc:
		return "CallingConv=IntelOclBicc"
	case CallingConv_X86_64Win64:
		return "CallingConv=X86_64Win64"
	case CallingConv_X86_64SysV:
		return "CallingConv=X86_64SysV"
	case CallingConv_X86VectorCall:
		return "CallingConv=X86VectorCall"
	case CallingConv_Swift:
		return "CallingConv=Swift"
	case CallingConv_PreserveMost:
		return "CallingConv=PreserveMost"
	case CallingConv_PreserveAll:
		return "CallingConv=PreserveAll"
	case CallingConv_Invalid:
		return "CallingConv=Invalid"
	case CallingConv_Unexposed:
		return "CallingConv=Unexposed"
	}

	return fmt.Sprintf("CallingConv unkown %d", int(cc))
}

func (cc CallingConv) String() string {
	return cc.Spelling()
}
