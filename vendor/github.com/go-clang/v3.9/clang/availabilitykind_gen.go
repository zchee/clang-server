package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Describes the availability of a particular entity, which indicates whether the use of this entity will result in a warning or error due to it being deprecated or unavailable.
type AvailabilityKind uint32

const (
	// The entity is available.
	Availability_Available AvailabilityKind = C.CXAvailability_Available
	// The entity is available, but has been deprecated (and its use is not recommended).
	Availability_Deprecated = C.CXAvailability_Deprecated
	// The entity is not available; any use of it will be an error.
	Availability_NotAvailable = C.CXAvailability_NotAvailable
	// The entity is available, but not accessible; any use of it will be an error.
	Availability_NotAccessible = C.CXAvailability_NotAccessible
)

func (ak AvailabilityKind) Spelling() string {
	switch ak {
	case Availability_Available:
		return "Availability=Available"
	case Availability_Deprecated:
		return "Availability=Deprecated"
	case Availability_NotAvailable:
		return "Availability=NotAvailable"
	case Availability_NotAccessible:
		return "Availability=NotAccessible"
	}

	return fmt.Sprintf("AvailabilityKind unkown %d", int(ak))
}

func (ak AvailabilityKind) String() string {
	return ak.Spelling()
}
