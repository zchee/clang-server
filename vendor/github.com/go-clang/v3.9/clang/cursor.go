package clang

// #include "go-clang.h"
import "C"
import (
	"sync"
	"unsafe"
)

// Determine the availability of the entity that this cursor refers to on any platforms for which availability information is known. \param cursor The cursor to query. \param always_deprecated If non-NULL, will be set to indicate whether the entity is deprecated on all platforms. \param deprecated_message If non-NULL, will be set to the message text provided along with the unconditional deprecation of this entity. The client is responsible for deallocating this string. \param always_unavailable If non-NULL, will be set to indicate whether the entity is unavailable on all platforms. \param unavailable_message If non-NULL, will be set to the message text provided along with the unconditional unavailability of this entity. The client is responsible for deallocating this string. \param availability If non-NULL, an array of CXPlatformAvailability instances that will be populated with platform availability information, up to either the number of platforms for which availability information is available (as returned by this function) or \c availability_size, whichever is smaller. \param availability_size The number of elements available in the \c availability array. \returns The number of platforms (N) for which availability information is available (which is unrelated to \c availability_size). Note that the client is responsible for calling \c clang_disposeCXPlatformAvailability to free each of the platform-availability structures returned. There are \c min(N, availability_size) such structures.
func (c Cursor) PlatformAvailability(availabilitySize int) (always_deprecated bool, deprecated_msg string, always_unavailable bool, unavailable_msg string, availability []PlatformAvailability) {
	var c_always_deprecated C.int
	var c_deprecated_message cxstring
	defer c_deprecated_message.Dispose()
	var c_always_unavailable C.int
	var c_unavailable_message cxstring
	defer c_unavailable_message.Dispose()
	var cp_availability = make([]C.CXPlatformAvailability, availabilitySize)

	nn := int(C.clang_getCursorPlatformAvailability(c.c, &c_always_deprecated, &c_deprecated_message.c, &c_always_unavailable, &c_unavailable_message.c, &cp_availability[0], C.int(len(cp_availability))))

	if nn > availabilitySize {
		nn = availabilitySize
	}

	availability = make([]PlatformAvailability, nn)
	for i := 0; i < nn; i++ {
		availability[i] = PlatformAvailability{&cp_availability[i]}
	}

	return c_always_deprecated != 0, c_deprecated_message.String(), c_always_unavailable != 0, c_unavailable_message.String(), availability
}

// CursorVisitor does the following.
/**
 * \brief Visitor invoked for each cursor found by a traversal.
 *
 * This visitor function will be invoked for each cursor found by
 * clang_visitCursorChildren(). Its first argument is the cursor being
 * visited, its second argument is the parent visitor for that cursor,
 * and its third argument is the client data provided to
 * clang_visitCursorChildren().
 *
 * The visitor should return one of the \c CXChildVisitResult values
 * to direct clang_visitCursorChildren().
 */
type CursorVisitor func(cursor, parent Cursor) (status ChildVisitResult)

type funcRegistry struct {
	sync.RWMutex

	index int
	funcs map[int]*CursorVisitor
}

func (fm *funcRegistry) register(f *CursorVisitor) int {
	fm.Lock()
	defer fm.Unlock()

	fm.index++
	for fm.funcs[fm.index] != nil {
		fm.index++
	}

	fm.funcs[fm.index] = f

	return fm.index
}

func (fm *funcRegistry) lookup(index int) *CursorVisitor {
	fm.RLock()
	defer fm.RUnlock()

	return fm.funcs[index]
}

func (fm *funcRegistry) unregister(index int) {
	fm.Lock()

	delete(fm.funcs, index)

	fm.Unlock()
}

var visitors = &funcRegistry{
	funcs: map[int]*CursorVisitor{},
}

// GoClangCursorVisitor calls the cursor visitor
//export GoClangCursorVisitor
func GoClangCursorVisitor(cursor C.CXCursor, parent C.CXCursor, cfct unsafe.Pointer) (status ChildVisitResult) {
	i := *(*C.int)(cfct)
	f := visitors.lookup(int(i))

	return (*f)(Cursor{cursor}, Cursor{parent})
}

// Visit does the following.
/**
 * \brief Visit the children of a particular cursor.
 *
 * This function visits all the direct children of the given cursor,
 * invoking the given \p visitor function with the cursors of each
 * visited child. The traversal may be recursive, if the visitor returns
 * \c CXChildVisit_Recurse. The traversal may also be ended prematurely, if
 * the visitor returns \c CXChildVisit_Break.
 *
 * \param parent the cursor whose child may be visited. All kinds of
 * cursors can be visited, including invalid cursors (which, by
 * definition, have no children).
 *
 * \param visitor the visitor function that will be invoked for each
 * child of \p parent.
 *
 * \param client_data pointer data supplied by the client, which will
 * be passed to the visitor each time it is invoked.
 *
 * \returns a non-zero value if the traversal was terminated
 * prematurely by the visitor returning \c CXChildVisit_Break.
 */
func (c Cursor) Visit(visitor CursorVisitor) bool {
	i := visitors.register(&visitor)
	defer visitors.unregister(i)

	// We need a pointer to the index because clang_visitChildren data parameter is a void pointer.
	ci := C.int(i)

	o := C.go_clang_visit_children(c.c, unsafe.Pointer(&ci))

	return o == C.uint(0)
}
