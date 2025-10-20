// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioGraphicsDevice] class.
var (
	VZVirtioGraphicsDeviceClass     _VZVirtioGraphicsDeviceClass
	VZVirtioGraphicsDeviceClassOnce sync.Once
)

func getVZVirtioGraphicsDeviceClass() _VZVirtioGraphicsDeviceClass {
	VZVirtioGraphicsDeviceClassOnce.Do(func() {
		VZVirtioGraphicsDeviceClass = _VZVirtioGraphicsDeviceClass{objc.GetClass("VZVirtioGraphicsDevice")}
	})
	return VZVirtioGraphicsDeviceClass
}

type _VZVirtioGraphicsDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioGraphicsDevice] class.
type IVZVirtioGraphicsDevice interface {
	IVZGraphicsDevice
}

// A Virtio graphics device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDevice
type VZVirtioGraphicsDevice struct {
	VZGraphicsDevice
}

// VZVirtioGraphicsDeviceFrom constructs a [VZVirtioGraphicsDevice] from an unsafe.Pointer.
//
// A Virtio graphics device.
func VZVirtioGraphicsDeviceFrom(ptr unsafe.Pointer) VZVirtioGraphicsDevice {
	return VZVirtioGraphicsDevice{
		VZGraphicsDevice: VZGraphicsDeviceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioGraphicsDeviceClass) Alloc() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioGraphicsDeviceClass) New() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioGraphicsDevice) Init() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioGraphicsDevice) Autorelease() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsDevice creates a new VZVirtioGraphicsDevice instance.
func NewVZVirtioGraphicsDevice() VZVirtioGraphicsDevice {
	return getVZVirtioGraphicsDeviceClass().New()
}




