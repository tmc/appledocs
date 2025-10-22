// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZGraphicsDevice] class.
var (
	VZGraphicsDeviceClass     _VZGraphicsDeviceClass
	VZGraphicsDeviceClassOnce sync.Once
)

func getVZGraphicsDeviceClass() _VZGraphicsDeviceClass {
	VZGraphicsDeviceClassOnce.Do(func() {
		VZGraphicsDeviceClass = _VZGraphicsDeviceClass{objc.GetClass("VZGraphicsDevice")}
	})
	return VZGraphicsDeviceClass
}

type _VZGraphicsDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZGraphicsDevice] class.
type IVZGraphicsDevice interface {
	objectivec.IObject
	Displays() []VZGraphicsDisplay
	GraphicsDevices() VZGraphicsDevice
	SetGraphicsDevices(value IVZGraphicsDevice)
}

// A class that represents a graphics device in a VM.
//
// You don’t instantiate a directly. Graphics devices are first configured on the through a subclass of . When the framework creates a from the configuration, the graphics devices are available through the property. The real type of corresponds to the type used by the configuration. For example, a leads to a device of type and a leads to a device of type .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice
type VZGraphicsDevice struct {
	objectivec.Object
}

// VZGraphicsDeviceFrom constructs a [VZGraphicsDevice] from an unsafe.Pointer.
//
// A class that represents a graphics device in a VM.
func VZGraphicsDeviceFrom(ptr unsafe.Pointer) VZGraphicsDevice {
	return VZGraphicsDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZGraphicsDeviceClass) Alloc() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZGraphicsDeviceClass) New() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZGraphicsDevice) Init() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZGraphicsDevice) Autorelease() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDevice creates a new VZGraphicsDevice instance.
func NewVZGraphicsDevice() VZGraphicsDevice {
	return getVZGraphicsDeviceClass().New()
}


// The list of graphics displays configured for this graphics device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice/displays
func (v_ VZGraphicsDevice) Displays() []VZGraphicsDisplay {
	rv := objc.Send[[]VZGraphicsDisplay](v_.ID, objc.Sel("displays"))
	return rv
}

// The list of configured graphics devices on the virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/graphicsdevices
func (v_ VZGraphicsDevice) GraphicsDevices() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](v_.ID, objc.Sel("graphicsDevices"))
	return rv
}


// SetGraphicsDevices sets the value of the graphicsDevices property.
// The list of configured graphics devices on the virtual machine.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/graphicsdevices
func (v_ VZGraphicsDevice) SetGraphicsDevices(value IVZGraphicsDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setGraphicsDevices:"), value)
}



