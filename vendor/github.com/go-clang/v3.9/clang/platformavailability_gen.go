package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Describes the availability of a given entity on a particular platform, e.g., a particular class might only be available on Mac OS 10.7 or newer.
type PlatformAvailability struct {
	c *C.CXPlatformAvailability
}

// Free the memory associated with a CXPlatformAvailability structure.
func (pa PlatformAvailability) Dispose() {
	C.clang_disposeCXPlatformAvailability(pa.c)
}

/*
	A string that describes the platform for which this structure
	provides availability information.

	Possible values are "ios" or "macos".
*/
func (pa PlatformAvailability) Platform() string {
	o := cxstring{pa.c.Platform}
	defer o.Dispose()

	return o.String()
}

// The version number in which this entity was introduced.
func (pa PlatformAvailability) Introduced() Version {
	return Version{pa.c.Introduced}
}

// The version number in which this entity was deprecated (but is still available).
func (pa PlatformAvailability) Deprecated() Version {
	return Version{pa.c.Deprecated}
}

// The version number in which this entity was obsoleted, and therefore is no longer available.
func (pa PlatformAvailability) Obsoleted() Version {
	return Version{pa.c.Obsoleted}
}

// Whether the entity is unconditionally unavailable on this platform.
func (pa PlatformAvailability) Unavailable() int32 {
	return int32(pa.c.Unavailable)
}

// An optional message to provide to a user of this API, e.g., to suggest replacement APIs.
func (pa PlatformAvailability) Message() string {
	o := cxstring{pa.c.Message}
	defer o.Dispose()

	return o.String()
}
