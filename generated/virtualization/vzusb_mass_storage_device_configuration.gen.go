// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZUSBMassStorageDeviceConfiguration] class.
var (
	VZUSBMassStorageDeviceConfigurationClass     _VZUSBMassStorageDeviceConfigurationClass
	VZUSBMassStorageDeviceConfigurationClassOnce sync.Once
)

func getVZUSBMassStorageDeviceConfigurationClass() _VZUSBMassStorageDeviceConfigurationClass {
	VZUSBMassStorageDeviceConfigurationClassOnce.Do(func() {
		VZUSBMassStorageDeviceConfigurationClass = _VZUSBMassStorageDeviceConfigurationClass{objc.GetClass("VZUSBMassStorageDeviceConfiguration")}
	})
	return VZUSBMassStorageDeviceConfigurationClass
}

type _VZUSBMassStorageDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZUSBMassStorageDeviceConfiguration] class.
type IVZUSBMassStorageDeviceConfiguration interface {
	IVZStorageDeviceConfiguration
}

// The configuration object that represents a USB Mass storage device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDeviceConfiguration
type VZUSBMassStorageDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZUSBMassStorageDeviceConfigurationFrom constructs a [VZUSBMassStorageDeviceConfiguration] from an unsafe.Pointer.
//
// The configuration object that represents a USB Mass storage device.
func VZUSBMassStorageDeviceConfigurationFrom(ptr unsafe.Pointer) VZUSBMassStorageDeviceConfiguration {
	return VZUSBMassStorageDeviceConfiguration{
		VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZUSBMassStorageDeviceConfigurationClass) Alloc() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZUSBMassStorageDeviceConfigurationClass) New() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZUSBMassStorageDeviceConfiguration) Init() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZUSBMassStorageDeviceConfiguration) Autorelease() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBMassStorageDeviceConfiguration creates a new VZUSBMassStorageDeviceConfiguration instance.
func NewVZUSBMassStorageDeviceConfiguration() VZUSBMassStorageDeviceConfiguration {
	return getVZUSBMassStorageDeviceConfigurationClass().New()
}




