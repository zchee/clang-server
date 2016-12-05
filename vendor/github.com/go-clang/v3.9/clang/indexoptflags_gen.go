package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type IndexOptFlags uint32

const (
	// Used to indicate that no special indexing options are needed.
	IndexOpt_None IndexOptFlags = C.CXIndexOpt_None
	// Used to indicate that IndexerCallbacks#indexEntityReference should be invoked for only one reference of an entity per source file that does not also include a declaration/definition of the entity.
	IndexOpt_SuppressRedundantRefs = C.CXIndexOpt_SuppressRedundantRefs
	// Function-local symbols should be indexed. If this is not set function-local symbols will be ignored.
	IndexOpt_IndexFunctionLocalSymbols = C.CXIndexOpt_IndexFunctionLocalSymbols
	// Implicit function/class template instantiations should be indexed. If this is not set, implicit instantiations will be ignored.
	IndexOpt_IndexImplicitTemplateInstantiations = C.CXIndexOpt_IndexImplicitTemplateInstantiations
	// Suppress all compiler warnings when parsing for indexing.
	IndexOpt_SuppressWarnings = C.CXIndexOpt_SuppressWarnings
	// Skip a function/method body that was already parsed during an indexing session associated with a CXIndexAction object. Bodies in system headers are always skipped.
	IndexOpt_SkipParsedBodiesInSession = C.CXIndexOpt_SkipParsedBodiesInSession
)

func (iof IndexOptFlags) Spelling() string {
	switch iof {
	case IndexOpt_None:
		return "IndexOpt=None"
	case IndexOpt_SuppressRedundantRefs:
		return "IndexOpt=SuppressRedundantRefs"
	case IndexOpt_IndexFunctionLocalSymbols:
		return "IndexOpt=IndexFunctionLocalSymbols"
	case IndexOpt_IndexImplicitTemplateInstantiations:
		return "IndexOpt=IndexImplicitTemplateInstantiations"
	case IndexOpt_SuppressWarnings:
		return "IndexOpt=SuppressWarnings"
	case IndexOpt_SkipParsedBodiesInSession:
		return "IndexOpt=SkipParsedBodiesInSession"
	}

	return fmt.Sprintf("IndexOptFlags unkown %d", int(iof))
}

func (iof IndexOptFlags) String() string {
	return iof.Spelling()
}
