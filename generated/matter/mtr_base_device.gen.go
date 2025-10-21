// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBaseDevice] class.
var (
	MTRBaseDeviceClass     _MTRBaseDeviceClass
	MTRBaseDeviceClassOnce sync.Once
)

func getMTRBaseDeviceClass() _MTRBaseDeviceClass {
	MTRBaseDeviceClassOnce.Do(func() {
		MTRBaseDeviceClass = _MTRBaseDeviceClass{objc.GetClass("MTRBaseDevice")}
	})
	return MTRBaseDeviceClass
}

type _MTRBaseDeviceClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseDevice] class.
type IMTRBaseDevice interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice
type MTRBaseDevice struct {
	objectivec.Object
}

// MTRBaseDeviceFrom constructs a [MTRBaseDevice] from an unsafe.Pointer.
func MTRBaseDeviceFrom(ptr unsafe.Pointer) MTRBaseDevice {
	return MTRBaseDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseDeviceClass) Alloc() MTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseDeviceClass) New() MTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseDevice) Init() MTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseDevice) Autorelease() MTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseDevice creates a new MTRBaseDevice instance.
func NewMTRBaseDevice() MTRBaseDevice {
	return getMTRBaseDeviceClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasedevice/sessiontransporttype
func (m_ MTRBaseDevice) SessionTransportType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sessionTransportType"))
	return rv
}


// SetSessionTransportType sets the value of the sessionTransportType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasedevice/sessiontransporttype
func (m_ MTRBaseDevice) SetSessionTransportType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionTransportType:"), value)
}



