// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZUSBScreenCoordinatePointingDeviceConfiguration] class.
var (
	VZUSBScreenCoordinatePointingDeviceConfigurationClass     _VZUSBScreenCoordinatePointingDeviceConfigurationClass
	VZUSBScreenCoordinatePointingDeviceConfigurationClassOnce sync.Once
)

func getVZUSBScreenCoordinatePointingDeviceConfigurationClass() _VZUSBScreenCoordinatePointingDeviceConfigurationClass {
	VZUSBScreenCoordinatePointingDeviceConfigurationClassOnce.Do(func() {
		VZUSBScreenCoordinatePointingDeviceConfigurationClass = _VZUSBScreenCoordinatePointingDeviceConfigurationClass{objc.GetClass("VZUSBScreenCoordinatePointingDeviceConfiguration")}
	})
	return VZUSBScreenCoordinatePointingDeviceConfigurationClass
}

type _VZUSBScreenCoordinatePointingDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZUSBScreenCoordinatePointingDeviceConfiguration] class.
type IVZUSBScreenCoordinatePointingDeviceConfiguration interface {
	IVZPointingDeviceConfiguration
	// properties:
	// methods:
}

// An object that defines the configuration for a USB pointing device that reports absolute coordinates.
//
// A can use this device to send pointer events to the VM.


// An object that defines the configuration for a USB pointing device that reports absolute coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBScreenCoordinatePointingDeviceConfiguration
type VZUSBScreenCoordinatePointingDeviceConfiguration struct {
	VZPointingDeviceConfiguration
}

// VZUSBScreenCoordinatePointingDeviceConfigurationFrom constructs a [VZUSBScreenCoordinatePointingDeviceConfiguration] from an unsafe.Pointer.
//
// An object that defines the configuration for a USB pointing device that reports absolute coordinates.
func VZUSBScreenCoordinatePointingDeviceConfigurationFrom(ptr unsafe.Pointer) VZUSBScreenCoordinatePointingDeviceConfiguration {
	return VZUSBScreenCoordinatePointingDeviceConfiguration{
		VZPointingDeviceConfiguration: VZPointingDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZUSBScreenCoordinatePointingDeviceConfigurationClass) Alloc() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZUSBScreenCoordinatePointingDeviceConfigurationClass) New() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZUSBScreenCoordinatePointingDeviceConfiguration) Init() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZUSBScreenCoordinatePointingDeviceConfiguration) Autorelease() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBScreenCoordinatePointingDeviceConfiguration creates a new VZUSBScreenCoordinatePointingDeviceConfiguration instance.
func NewVZUSBScreenCoordinatePointingDeviceConfiguration() VZUSBScreenCoordinatePointingDeviceConfiguration {
	return getVZUSBScreenCoordinatePointingDeviceConfigurationClass().New()
}




