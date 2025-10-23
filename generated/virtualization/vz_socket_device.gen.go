// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZSocketDevice] class.
var (
	VZSocketDeviceClass     _VZSocketDeviceClass
	VZSocketDeviceClassOnce sync.Once
)

func getVZSocketDeviceClass() _VZSocketDeviceClass {
	VZSocketDeviceClassOnce.Do(func() {
		VZSocketDeviceClass = _VZSocketDeviceClass{objc.GetClass("VZSocketDevice")}
	})
	return VZSocketDeviceClass
}

type _VZSocketDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZSocketDevice] class.
type IVZSocketDevice interface {
	objectivec.IObject
}

// A parent class referenced by other Virtualization classes.


// A parent class referenced by other Virtualization classes. [Full Topic]
type VZSocketDevice struct {
	objectivec.Object
}

// VZSocketDeviceFrom constructs a [VZSocketDevice] from an unsafe.Pointer.
//
// A parent class referenced by other Virtualization classes.
func VZSocketDeviceFrom(ptr unsafe.Pointer) VZSocketDevice {
	return VZSocketDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZSocketDeviceClass) Alloc() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZSocketDeviceClass) New() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZSocketDevice) Init() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZSocketDevice) Autorelease() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSocketDevice creates a new VZSocketDevice instance.
func NewVZSocketDevice() VZSocketDevice {
	return getVZSocketDeviceClass().New()
}




