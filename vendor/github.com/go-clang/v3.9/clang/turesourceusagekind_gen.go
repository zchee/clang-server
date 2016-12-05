package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Categorizes how memory is being used by a translation unit.
type TUResourceUsageKind uint32

const (
	TUResourceUsage_AST                                TUResourceUsageKind = C.CXTUResourceUsage_AST
	TUResourceUsage_Identifiers                                            = C.CXTUResourceUsage_Identifiers
	TUResourceUsage_Selectors                                              = C.CXTUResourceUsage_Selectors
	TUResourceUsage_GlobalCompletionResults                                = C.CXTUResourceUsage_GlobalCompletionResults
	TUResourceUsage_SourceManagerContentCache                              = C.CXTUResourceUsage_SourceManagerContentCache
	TUResourceUsage_AST_SideTables                                         = C.CXTUResourceUsage_AST_SideTables
	TUResourceUsage_SourceManager_Membuffer_Malloc                         = C.CXTUResourceUsage_SourceManager_Membuffer_Malloc
	TUResourceUsage_SourceManager_Membuffer_MMap                           = C.CXTUResourceUsage_SourceManager_Membuffer_MMap
	TUResourceUsage_ExternalASTSource_Membuffer_Malloc                     = C.CXTUResourceUsage_ExternalASTSource_Membuffer_Malloc
	TUResourceUsage_ExternalASTSource_Membuffer_MMap                       = C.CXTUResourceUsage_ExternalASTSource_Membuffer_MMap
	TUResourceUsage_Preprocessor                                           = C.CXTUResourceUsage_Preprocessor
	TUResourceUsage_PreprocessingRecord                                    = C.CXTUResourceUsage_PreprocessingRecord
	TUResourceUsage_SourceManager_DataStructures                           = C.CXTUResourceUsage_SourceManager_DataStructures
	TUResourceUsage_Preprocessor_HeaderSearch                              = C.CXTUResourceUsage_Preprocessor_HeaderSearch
	TUResourceUsage_MEMORY_IN_BYTES_BEGIN                                  = C.CXTUResourceUsage_MEMORY_IN_BYTES_BEGIN
	TUResourceUsage_MEMORY_IN_BYTES_END                                    = C.CXTUResourceUsage_MEMORY_IN_BYTES_END
	TUResourceUsage_First                                                  = C.CXTUResourceUsage_First
	TUResourceUsage_Last                                                   = C.CXTUResourceUsage_Last
)

// Returns the human-readable null-terminated C string that represents the name of the memory category. This string should never be freed.
func (turuk TUResourceUsageKind) Name() string {
	return C.GoString(C.clang_getTUResourceUsageName(C.enum_CXTUResourceUsageKind(turuk)))
}

func (turuk TUResourceUsageKind) Spelling() string {
	switch turuk {
	case TUResourceUsage_AST:
		return "TUResourceUsage=AST, MEMORY_IN_BYTES_BEGIN, First"
	case TUResourceUsage_Identifiers:
		return "TUResourceUsage=Identifiers"
	case TUResourceUsage_Selectors:
		return "TUResourceUsage=Selectors"
	case TUResourceUsage_GlobalCompletionResults:
		return "TUResourceUsage=GlobalCompletionResults"
	case TUResourceUsage_SourceManagerContentCache:
		return "TUResourceUsage=SourceManagerContentCache"
	case TUResourceUsage_AST_SideTables:
		return "TUResourceUsage=AST_SideTables"
	case TUResourceUsage_SourceManager_Membuffer_Malloc:
		return "TUResourceUsage=SourceManager_Membuffer_Malloc"
	case TUResourceUsage_SourceManager_Membuffer_MMap:
		return "TUResourceUsage=SourceManager_Membuffer_MMap"
	case TUResourceUsage_ExternalASTSource_Membuffer_Malloc:
		return "TUResourceUsage=ExternalASTSource_Membuffer_Malloc"
	case TUResourceUsage_ExternalASTSource_Membuffer_MMap:
		return "TUResourceUsage=ExternalASTSource_Membuffer_MMap"
	case TUResourceUsage_Preprocessor:
		return "TUResourceUsage=Preprocessor"
	case TUResourceUsage_PreprocessingRecord:
		return "TUResourceUsage=PreprocessingRecord"
	case TUResourceUsage_SourceManager_DataStructures:
		return "TUResourceUsage=SourceManager_DataStructures"
	case TUResourceUsage_Preprocessor_HeaderSearch:
		return "TUResourceUsage=Preprocessor_HeaderSearch, MEMORY_IN_BYTES_END, Last"
	}

	return fmt.Sprintf("TUResourceUsageKind unkown %d", int(turuk))
}

func (turuk TUResourceUsageKind) String() string {
	return turuk.Spelling()
}
