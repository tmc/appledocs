// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZMacGraphicsDeviceConfiguration] class.
var (
	VZMacGraphicsDeviceConfigurationClass     _VZMacGraphicsDeviceConfigurationClass
	VZMacGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZMacGraphicsDeviceConfigurationClass() _VZMacGraphicsDeviceConfigurationClass {
	VZMacGraphicsDeviceConfigurationClassOnce.Do(func() {
		VZMacGraphicsDeviceConfigurationClass = _VZMacGraphicsDeviceConfigurationClass{objc.GetClass("VZMacGraphicsDeviceConfiguration")}
	})
	return VZMacGraphicsDeviceConfigurationClass
}

type _VZMacGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZMacGraphicsDeviceConfiguration] class.
type IVZMacGraphicsDeviceConfiguration interface {
	objectivec.IObject
}

// Configuration for a display attached to a Mac graphics device.
//
// Use this device to attach a display that’s shown in a .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration
type VZMacGraphicsDeviceConfiguration struct {
	objectivec.Object
}

// VZMacGraphicsDeviceConfigurationFrom constructs a [VZMacGraphicsDeviceConfiguration] from an unsafe.Pointer.
//
// Configuration for a display attached to a Mac graphics device.
func VZMacGraphicsDeviceConfigurationFrom(ptr unsafe.Pointer) VZMacGraphicsDeviceConfiguration {
	return VZMacGraphicsDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMacGraphicsDeviceConfigurationClass) Alloc() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMacGraphicsDeviceConfigurationClass) New() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacGraphicsDeviceConfiguration) Init() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacGraphicsDeviceConfiguration) Autorelease() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDeviceConfiguration creates a new VZMacGraphicsDeviceConfiguration instance.
func NewVZMacGraphicsDeviceConfiguration() VZMacGraphicsDeviceConfiguration {
	return getVZMacGraphicsDeviceConfigurationClass().New()
}



// The displays associated with this graphics device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/displays
func (v_ VZMacGraphicsDeviceConfiguration) Displays() []VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[[]VZMacGraphicsDisplayConfiguration](v_.ID, objc.Sel("displays"))
	return rv
}


// SetDisplays sets the value of the displays property.
// The displays associated with this graphics device.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/displays
func (v_ VZMacGraphicsDeviceConfiguration) SetDisplays(value []VZMacGraphicsDisplayConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDisplays:"), value)
}

