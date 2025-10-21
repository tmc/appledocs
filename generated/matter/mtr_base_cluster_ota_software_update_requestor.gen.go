// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterOtaSoftwareUpdateRequestor] class.
var (
	MTRBaseClusterOtaSoftwareUpdateRequestorClass     _MTRBaseClusterOtaSoftwareUpdateRequestorClass
	MTRBaseClusterOtaSoftwareUpdateRequestorClassOnce sync.Once
)

func getMTRBaseClusterOtaSoftwareUpdateRequestorClass() _MTRBaseClusterOtaSoftwareUpdateRequestorClass {
	MTRBaseClusterOtaSoftwareUpdateRequestorClassOnce.Do(func() {
		MTRBaseClusterOtaSoftwareUpdateRequestorClass = _MTRBaseClusterOtaSoftwareUpdateRequestorClass{objc.GetClass("MTRBaseClusterOtaSoftwareUpdateRequestor")}
	})
	return MTRBaseClusterOtaSoftwareUpdateRequestorClass
}

type _MTRBaseClusterOtaSoftwareUpdateRequestorClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterOtaSoftwareUpdateRequestor] class.
type IMTRBaseClusterOtaSoftwareUpdateRequestor interface {
	IMTRBaseClusterOTASoftwareUpdateRequestor
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOtaSoftwareUpdateRequestor-35vsy
type MTRBaseClusterOtaSoftwareUpdateRequestor struct {
	MTRBaseClusterOTASoftwareUpdateRequestor
}

// MTRBaseClusterOtaSoftwareUpdateRequestorFrom constructs a [MTRBaseClusterOtaSoftwareUpdateRequestor] from an unsafe.Pointer.
func MTRBaseClusterOtaSoftwareUpdateRequestorFrom(ptr unsafe.Pointer) MTRBaseClusterOtaSoftwareUpdateRequestor {
	return MTRBaseClusterOtaSoftwareUpdateRequestor{
		MTRBaseClusterOTASoftwareUpdateRequestor: MTRBaseClusterOTASoftwareUpdateRequestorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOtaSoftwareUpdateRequestorClass) Alloc() MTRBaseClusterOtaSoftwareUpdateRequestor {
	rv := objc.Send[MTRBaseClusterOtaSoftwareUpdateRequestor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterOtaSoftwareUpdateRequestorClass) New() MTRBaseClusterOtaSoftwareUpdateRequestor {
	rv := objc.Send[MTRBaseClusterOtaSoftwareUpdateRequestor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOtaSoftwareUpdateRequestor) Init() MTRBaseClusterOtaSoftwareUpdateRequestor {
	rv := objc.Send[MTRBaseClusterOtaSoftwareUpdateRequestor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOtaSoftwareUpdateRequestor) Autorelease() MTRBaseClusterOtaSoftwareUpdateRequestor {
	rv := objc.Send[MTRBaseClusterOtaSoftwareUpdateRequestor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOtaSoftwareUpdateRequestor creates a new MTRBaseClusterOtaSoftwareUpdateRequestor instance.
func NewMTRBaseClusterOtaSoftwareUpdateRequestor() MTRBaseClusterOtaSoftwareUpdateRequestor {
	return getMTRBaseClusterOtaSoftwareUpdateRequestorClass().New()
}




