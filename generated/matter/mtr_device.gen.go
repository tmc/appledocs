// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDevice] class.
var (
	MTRDeviceClass     _MTRDeviceClass
	MTRDeviceClassOnce sync.Once
)

func getMTRDeviceClass() _MTRDeviceClass {
	MTRDeviceClassOnce.Do(func() {
		MTRDeviceClass = _MTRDeviceClass{objc.GetClass("MTRDevice")}
	})
	return MTRDeviceClass
}

type _MTRDeviceClass struct {
	class objc.Class
}

// An interface definition for the [MTRDevice] class.
type IMTRDevice interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice
type MTRDevice struct {
	objectivec.Object
}

// MTRDeviceFrom constructs a [MTRDevice] from an unsafe.Pointer.
func MTRDeviceFrom(ptr unsafe.Pointer) MTRDevice {
	return MTRDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceClass) Alloc() MTRDevice {
	rv := objc.Send[MTRDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceClass) New() MTRDevice {
	rv := objc.Send[MTRDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDevice) Init() MTRDevice {
	rv := objc.Send[MTRDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDevice) Autorelease() MTRDevice {
	rv := objc.Send[MTRDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDevice creates a new MTRDevice instance.
func NewMTRDevice() MTRDevice {
	return getMTRDeviceClass().New()
}




