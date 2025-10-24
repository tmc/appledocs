// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZNVMExpressControllerDeviceConfiguration] class.
var (
	VZNVMExpressControllerDeviceConfigurationClass     _VZNVMExpressControllerDeviceConfigurationClass
	VZNVMExpressControllerDeviceConfigurationClassOnce sync.Once
)

func getVZNVMExpressControllerDeviceConfigurationClass() _VZNVMExpressControllerDeviceConfigurationClass {
	VZNVMExpressControllerDeviceConfigurationClassOnce.Do(func() {
		VZNVMExpressControllerDeviceConfigurationClass = _VZNVMExpressControllerDeviceConfigurationClass{objc.GetClass("VZNVMExpressControllerDeviceConfiguration")}
	})
	return VZNVMExpressControllerDeviceConfigurationClass
}

type _VZNVMExpressControllerDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZNVMExpressControllerDeviceConfiguration] class.
type IVZNVMExpressControllerDeviceConfiguration interface {
	IVZStorageDeviceConfiguration
	// properties:
	// methods:
}

// The configuration object that represents an NVM Express Controller storage device.
//
// This device configuration creates a storage device that conforms to the . The device configuration is valid only if used with .


// The configuration object that represents an NVM Express Controller storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNVMExpressControllerDeviceConfiguration
type VZNVMExpressControllerDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZNVMExpressControllerDeviceConfigurationFrom constructs a [VZNVMExpressControllerDeviceConfiguration] from an unsafe.Pointer.
//
// The configuration object that represents an NVM Express Controller storage device.
func VZNVMExpressControllerDeviceConfigurationFrom(ptr unsafe.Pointer) VZNVMExpressControllerDeviceConfiguration {
	return VZNVMExpressControllerDeviceConfiguration{
		VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZNVMExpressControllerDeviceConfigurationClass) Alloc() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZNVMExpressControllerDeviceConfigurationClass) New() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZNVMExpressControllerDeviceConfiguration) Init() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZNVMExpressControllerDeviceConfiguration) Autorelease() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNVMExpressControllerDeviceConfiguration creates a new VZNVMExpressControllerDeviceConfiguration instance.
func NewVZNVMExpressControllerDeviceConfiguration() VZNVMExpressControllerDeviceConfiguration {
	return getVZNVMExpressControllerDeviceConfigurationClass().New()
}



// Creates a new NVM Express controller configuration with the storage device attachment you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNVMExpressControllerDeviceConfiguration/init(attachment:)
func NewVZNVMExpressControllerDeviceConfigurationWithAttachment(attachment IVZStorageDeviceAttachment) VZNVMExpressControllerDeviceConfiguration {
	instance := getVZNVMExpressControllerDeviceConfigurationClass().Alloc()
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](instance.ID, objc.Sel("initWithAttachment:"), attachment)
	rv.Autorelease()
	return rv
}



