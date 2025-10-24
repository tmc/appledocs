// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZMacGraphicsDevice] class.
var (
	VZMacGraphicsDeviceClass     _VZMacGraphicsDeviceClass
	VZMacGraphicsDeviceClassOnce sync.Once
)

func getVZMacGraphicsDeviceClass() _VZMacGraphicsDeviceClass {
	VZMacGraphicsDeviceClassOnce.Do(func() {
		VZMacGraphicsDeviceClass = _VZMacGraphicsDeviceClass{objc.GetClass("VZMacGraphicsDevice")}
	})
	return VZMacGraphicsDeviceClass
}

type _VZMacGraphicsDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZMacGraphicsDevice] class.
type IVZMacGraphicsDevice interface {
	IVZGraphicsDevice
	// properties:
	// methods:
}

// An object that represents a Mac graphics device.
//
// You don’t instantiate a   directly. Graphics devices are first configured on the through a subclass of  .  When the framework creates a VZVirtualMachine from the configuration, the graphics devices are available through the property.


// An object that represents a Mac graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDevice
type VZMacGraphicsDevice struct {
	VZGraphicsDevice
}

// VZMacGraphicsDeviceFrom constructs a [VZMacGraphicsDevice] from an unsafe.Pointer.
//
// An object that represents a Mac graphics device.
func VZMacGraphicsDeviceFrom(ptr unsafe.Pointer) VZMacGraphicsDevice {
	return VZMacGraphicsDevice{
		VZGraphicsDevice: VZGraphicsDeviceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMacGraphicsDeviceClass) Alloc() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMacGraphicsDeviceClass) New() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacGraphicsDevice) Init() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacGraphicsDevice) Autorelease() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDevice creates a new VZMacGraphicsDevice instance.
func NewVZMacGraphicsDevice() VZMacGraphicsDevice {
	return getVZMacGraphicsDeviceClass().New()
}




