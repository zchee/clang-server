package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

/*
	Flags that control the creation of translation units.

	The enumerators in this enumeration type are meant to be bitwise
	ORed together to specify which options should be used when
	constructing the translation unit.
*/
type TranslationUnit_Flags uint32

const (
	// Used to indicate that no special translation-unit options are needed.
	TranslationUnit_None TranslationUnit_Flags = C.CXTranslationUnit_None
	/*
		Used to indicate that the parser should construct a "detailed"
		preprocessing record, including all macro definitions and instantiations.

		Constructing a detailed preprocessing record requires more memory
		and time to parse, since the information contained in the record
		is usually not retained. However, it can be useful for
		applications that require more detailed information about the
		behavior of the preprocessor.
	*/
	TranslationUnit_DetailedPreprocessingRecord = C.CXTranslationUnit_DetailedPreprocessingRecord
	/*
		Used to indicate that the translation unit is incomplete.

		When a translation unit is considered "incomplete", semantic
		analysis that is typically performed at the end of the
		translation unit will be suppressed. For example, this suppresses
		the completion of tentative declarations in C and of
		instantiation of implicitly-instantiation function templates in
		C++. This option is typically used when parsing a header with the
		intent of producing a precompiled header.
	*/
	TranslationUnit_Incomplete = C.CXTranslationUnit_Incomplete
	/*
		Used to indicate that the translation unit should be built with an
		implicit precompiled header for the preamble.

		An implicit precompiled header is used as an optimization when a
		particular translation unit is likely to be reparsed many times
		when the sources aren't changing that often. In this case, an
		implicit precompiled header will be built containing all of the
		initial includes at the top of the main file (what we refer to as
		the "preamble" of the file). In subsequent parses, if the
		preamble or the files in it have not changed, \c
		clang_reparseTranslationUnit() will re-use the implicit
		precompiled header to improve parsing performance.
	*/
	TranslationUnit_PrecompiledPreamble = C.CXTranslationUnit_PrecompiledPreamble
	/*
		Used to indicate that the translation unit should cache some
		code-completion results with each reparse of the source file.

		Caching of code-completion results is a performance optimization that
		introduces some overhead to reparsing but improves the performance of
		code-completion operations.
	*/
	TranslationUnit_CacheCompletionResults = C.CXTranslationUnit_CacheCompletionResults
	/*
		Used to indicate that the translation unit will be serialized with
		clang_saveTranslationUnit.

		This option is typically used when parsing a header with the intent of
		producing a precompiled header.
	*/
	TranslationUnit_ForSerialization = C.CXTranslationUnit_ForSerialization
	/*
		DEPRECATED: Enabled chained precompiled preambles in C++.

		Note: this is a *temporary* option that is available only while
		we are testing C++ precompiled preamble support. It is deprecated.
	*/
	TranslationUnit_CXXChainedPCH = C.CXTranslationUnit_CXXChainedPCH
	/*
		Used to indicate that function/method bodies should be skipped while
		parsing.

		This option can be used to search for declarations/definitions while
		ignoring the usages.
	*/
	TranslationUnit_SkipFunctionBodies = C.CXTranslationUnit_SkipFunctionBodies
	// Used to indicate that brief documentation comments should be included into the set of code completions returned from this translation unit.
	TranslationUnit_IncludeBriefCommentsInCodeCompletion = C.CXTranslationUnit_IncludeBriefCommentsInCodeCompletion
	// Used to indicate that the precompiled preamble should be created on the first parse. Otherwise it will be created on the first reparse. This trades runtime on the first parse (serializing the preamble takes time) for reduced runtime on the second parse (can now reuse the preamble).
	TranslationUnit_CreatePreambleOnFirstParse = C.CXTranslationUnit_CreatePreambleOnFirstParse
	/*
		Do not stop processing when fatal errors are encountered.

		When fatal errors are encountered while parsing a translation unit,
		semantic analysis is typically stopped early when compiling code. A common
		source for fatal errors are unresolvable include files. For the
		purposes of an IDE, this is undesirable behavior and as much information
		as possible should be reported. Use this flag to enable this behavior.
	*/
	TranslationUnit_KeepGoing = C.CXTranslationUnit_KeepGoing
)

func (tuf TranslationUnit_Flags) Spelling() string {
	switch tuf {
	case TranslationUnit_None:
		return "TranslationUnit=None"
	case TranslationUnit_DetailedPreprocessingRecord:
		return "TranslationUnit=DetailedPreprocessingRecord"
	case TranslationUnit_Incomplete:
		return "TranslationUnit=Incomplete"
	case TranslationUnit_PrecompiledPreamble:
		return "TranslationUnit=PrecompiledPreamble"
	case TranslationUnit_CacheCompletionResults:
		return "TranslationUnit=CacheCompletionResults"
	case TranslationUnit_ForSerialization:
		return "TranslationUnit=ForSerialization"
	case TranslationUnit_CXXChainedPCH:
		return "TranslationUnit=CXXChainedPCH"
	case TranslationUnit_SkipFunctionBodies:
		return "TranslationUnit=SkipFunctionBodies"
	case TranslationUnit_IncludeBriefCommentsInCodeCompletion:
		return "TranslationUnit=IncludeBriefCommentsInCodeCompletion"
	case TranslationUnit_CreatePreambleOnFirstParse:
		return "TranslationUnit=CreatePreambleOnFirstParse"
	case TranslationUnit_KeepGoing:
		return "TranslationUnit=KeepGoing"
	}

	return fmt.Sprintf("TranslationUnit_Flags unkown %d", int(tuf))
}

func (tuf TranslationUnit_Flags) String() string {
	return tuf.Spelling()
}
