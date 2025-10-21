// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterOtaSoftwareUpdateRequestor] class.
var (
	MTRClusterOtaSoftwareUpdateRequestorClass     _MTRClusterOtaSoftwareUpdateRequestorClass
	MTRClusterOtaSoftwareUpdateRequestorClassOnce sync.Once
)

func getMTRClusterOtaSoftwareUpdateRequestorClass() _MTRClusterOtaSoftwareUpdateRequestorClass {
	MTRClusterOtaSoftwareUpdateRequestorClassOnce.Do(func() {
		MTRClusterOtaSoftwareUpdateRequestorClass = _MTRClusterOtaSoftwareUpdateRequestorClass{objc.GetClass("MTRClusterOtaSoftwareUpdateRequestor")}
	})
	return MTRClusterOtaSoftwareUpdateRequestorClass
}

type _MTRClusterOtaSoftwareUpdateRequestorClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterOtaSoftwareUpdateRequestor] class.
type IMTRClusterOtaSoftwareUpdateRequestor interface {
	IMTRClusterOTASoftwareUpdateRequestor
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOtaSoftwareUpdateRequestor-22dfp
type MTRClusterOtaSoftwareUpdateRequestor struct {
	MTRClusterOTASoftwareUpdateRequestor
}

// MTRClusterOtaSoftwareUpdateRequestorFrom constructs a [MTRClusterOtaSoftwareUpdateRequestor] from an unsafe.Pointer.
func MTRClusterOtaSoftwareUpdateRequestorFrom(ptr unsafe.Pointer) MTRClusterOtaSoftwareUpdateRequestor {
	return MTRClusterOtaSoftwareUpdateRequestor{
		MTRClusterOTASoftwareUpdateRequestor: MTRClusterOTASoftwareUpdateRequestorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOtaSoftwareUpdateRequestorClass) Alloc() MTRClusterOtaSoftwareUpdateRequestor {
	rv := objc.Send[MTRClusterOtaSoftwareUpdateRequestor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterOtaSoftwareUpdateRequestorClass) New() MTRClusterOtaSoftwareUpdateRequestor {
	rv := objc.Send[MTRClusterOtaSoftwareUpdateRequestor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOtaSoftwareUpdateRequestor) Init() MTRClusterOtaSoftwareUpdateRequestor {
	rv := objc.Send[MTRClusterOtaSoftwareUpdateRequestor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOtaSoftwareUpdateRequestor) Autorelease() MTRClusterOtaSoftwareUpdateRequestor {
	rv := objc.Send[MTRClusterOtaSoftwareUpdateRequestor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOtaSoftwareUpdateRequestor creates a new MTRClusterOtaSoftwareUpdateRequestor instance.
func NewMTRClusterOtaSoftwareUpdateRequestor() MTRClusterOtaSoftwareUpdateRequestor {
	return getMTRClusterOtaSoftwareUpdateRequestorClass().New()
}




