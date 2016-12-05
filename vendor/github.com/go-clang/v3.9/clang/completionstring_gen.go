package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

/*
	A semantic string that describes a code-completion result.

	A semantic string that describes the formatting of a code-completion
	result as a single "template" of text that should be inserted into the
	source buffer when a particular code-completion result is selected.
	Each semantic string is made up of some number of "chunks", each of which
	contains some text along with a description of what that text means, e.g.,
	the name of the entity being referenced, whether the text chunk is part of
	the template, or whether it is a "placeholder" that the user should replace
	with actual code,of a specific kind. See CXCompletionChunkKind for a
	description of the different kinds of chunks.
*/
type CompletionString struct {
	c C.CXCompletionString
}

/*
	Determine the kind of a particular chunk within a completion string.

	Parameter completion_string the completion string to query.

	Parameter chunk_number the 0-based index of the chunk in the completion string.

	Returns the kind of the chunk at the index chunk_number.
*/
func (cs CompletionString) ChunkKind(chunkNumber uint32) CompletionChunkKind {
	return CompletionChunkKind(C.clang_getCompletionChunkKind(cs.c, C.uint(chunkNumber)))
}

/*
	Retrieve the text associated with a particular chunk within a
	completion string.

	Parameter completion_string the completion string to query.

	Parameter chunk_number the 0-based index of the chunk in the completion string.

	Returns the text associated with the chunk at index chunk_number.
*/
func (cs CompletionString) ChunkText(chunkNumber uint32) string {
	o := cxstring{C.clang_getCompletionChunkText(cs.c, C.uint(chunkNumber))}
	defer o.Dispose()

	return o.String()
}

/*
	Retrieve the completion string associated with a particular chunk
	within a completion string.

	Parameter completion_string the completion string to query.

	Parameter chunk_number the 0-based index of the chunk in the completion string.

	Returns the completion string associated with the chunk at index
	chunk_number.
*/
func (cs CompletionString) ChunkCompletionString(chunkNumber uint32) CompletionString {
	return CompletionString{C.clang_getCompletionChunkCompletionString(cs.c, C.uint(chunkNumber))}
}

// Retrieve the number of chunks in the given code-completion string.
func (cs CompletionString) NumChunks() uint32 {
	return uint32(C.clang_getNumCompletionChunks(cs.c))
}

/*
	Determine the priority of this code completion.

	The priority of a code completion indicates how likely it is that this
	particular completion is the completion that the user will select. The
	priority is selected by various internal heuristics.

	Parameter completion_string The completion string to query.

	Returns The priority of this completion string. Smaller values indicate
	higher-priority (more likely) completions.
*/
func (cs CompletionString) Priority() uint32 {
	return uint32(C.clang_getCompletionPriority(cs.c))
}

/*
	Determine the availability of the entity that this code-completion
	string refers to.

	Parameter completion_string The completion string to query.

	Returns The availability of the completion string.
*/
func (cs CompletionString) Availability() AvailabilityKind {
	return AvailabilityKind(C.clang_getCompletionAvailability(cs.c))
}

/*
	Retrieve the number of annotations associated with the given
	completion string.

	Parameter completion_string the completion string to query.

	Returns the number of annotations associated with the given completion
	string.
*/
func (cs CompletionString) NumAnnotations() uint32 {
	return uint32(C.clang_getCompletionNumAnnotations(cs.c))
}

/*
	Retrieve the annotation associated with the given completion string.

	Parameter completion_string the completion string to query.

	Parameter annotation_number the 0-based index of the annotation of the
	completion string.

	Returns annotation string associated with the completion at index
	annotation_number, or a NULL string if that annotation is not available.
*/
func (cs CompletionString) Annotation(annotationNumber uint32) string {
	o := cxstring{C.clang_getCompletionAnnotation(cs.c, C.uint(annotationNumber))}
	defer o.Dispose()

	return o.String()
}

/*
	Retrieve the parent context of the given completion string.

	The parent context of a completion string is the semantic parent of
	the declaration (if any) that the code completion represents. For example,
	a code completion for an Objective-C method would have the method's class
	or protocol as its context.

	Parameter completion_string The code completion string whose parent is
	being queried.

	Parameter kind DEPRECATED: always set to CXCursor_NotImplemented if non-NULL.

	Returns The name of the completion parent, e.g., "NSObject" if
	the completion string represents a method in the NSObject class.
*/
func (cs CompletionString) Parent(kind *CursorKind) string {
	var cp_kind C.enum_CXCursorKind
	if kind != nil {
		cp_kind = C.enum_CXCursorKind(*kind)
	}

	o := cxstring{C.clang_getCompletionParent(cs.c, &cp_kind)}
	defer o.Dispose()

	return o.String()
}

// Retrieve the brief documentation comment attached to the declaration that corresponds to the given completion string.
func (cs CompletionString) BriefComment() string {
	o := cxstring{C.clang_getCompletionBriefComment(cs.c)}
	defer o.Dispose()

	return o.String()
}
