// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioEntropyDeviceConfiguration] class.
var (
	VZVirtioEntropyDeviceConfigurationClass     _VZVirtioEntropyDeviceConfigurationClass
	VZVirtioEntropyDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioEntropyDeviceConfigurationClass() _VZVirtioEntropyDeviceConfigurationClass {
	VZVirtioEntropyDeviceConfigurationClassOnce.Do(func() {
		VZVirtioEntropyDeviceConfigurationClass = _VZVirtioEntropyDeviceConfigurationClass{objc.GetClass("VZVirtioEntropyDeviceConfiguration")}
	})
	return VZVirtioEntropyDeviceConfigurationClass
}

type _VZVirtioEntropyDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioEntropyDeviceConfiguration] class.
type IVZVirtioEntropyDeviceConfiguration interface {
	IVZEntropyDeviceConfiguration
}

// A source of entropy for the guest’s random number generator.
//
// Use a object to expose a source of entropy for the guest operating system’s random-number generator. When you create this object and add it to your virtual machine’s configuration, the virtual machine configures a Virtio-compliant entropy device. The guest operating system uses this device as a seed to generate random numbers. Create a object and add it to the property of your virtual machine’s configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioEntropyDeviceConfiguration
type VZVirtioEntropyDeviceConfiguration struct {
	VZEntropyDeviceConfiguration
}

// VZVirtioEntropyDeviceConfigurationFrom constructs a [VZVirtioEntropyDeviceConfiguration] from an unsafe.Pointer.
//
// A source of entropy for the guest’s random number generator.
func VZVirtioEntropyDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioEntropyDeviceConfiguration {
	return VZVirtioEntropyDeviceConfiguration{
		VZEntropyDeviceConfiguration: VZEntropyDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioEntropyDeviceConfigurationClass) Alloc() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioEntropyDeviceConfigurationClass) New() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioEntropyDeviceConfiguration) Init() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioEntropyDeviceConfiguration) Autorelease() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioEntropyDeviceConfiguration creates a new VZVirtioEntropyDeviceConfiguration instance.
func NewVZVirtioEntropyDeviceConfiguration() VZVirtioEntropyDeviceConfiguration {
	return getVZVirtioEntropyDeviceConfigurationClass().New()
}



// The array of randomization devices that you expose to the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/entropydevices
func (v_ VZVirtioEntropyDeviceConfiguration) EntropyDevices() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("entropyDevices"))
	return rv
}


// SetEntropyDevices sets the value of the entropyDevices property.
// The array of randomization devices that you expose to the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/entropydevices
func (v_ VZVirtioEntropyDeviceConfiguration) SetEntropyDevices(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setEntropyDevices:"), value)
}


