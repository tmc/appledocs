// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] class.
var (
	VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass     _VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass
	VZVirtioTraditionalMemoryBalloonDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioTraditionalMemoryBalloonDeviceConfigurationClass() _VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass {
	VZVirtioTraditionalMemoryBalloonDeviceConfigurationClassOnce.Do(func() {
		VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass = _VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass{objc.GetClass("VZVirtioTraditionalMemoryBalloonDeviceConfiguration")}
	})
	return VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass
}

type _VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] class.
type IVZVirtioTraditionalMemoryBalloonDeviceConfiguration interface {
	IVZMemoryBalloonDeviceConfiguration
}

// A configuration object that provides a way to reclaim memory from the guest system.
//
// Create a object when you want the ability to reclaim memory from the guest operating system. After creating this object, add it to the property of your object. In response, the virtual machine provides a object, which you use to initiate memory-related requests with the guest system. Access that object from the property of .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDeviceConfiguration
type VZVirtioTraditionalMemoryBalloonDeviceConfiguration struct {
	VZMemoryBalloonDeviceConfiguration
}

// VZVirtioTraditionalMemoryBalloonDeviceConfigurationFrom constructs a [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] from an unsafe.Pointer.
//
// A configuration object that provides a way to reclaim memory from the guest system.
func VZVirtioTraditionalMemoryBalloonDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	return VZVirtioTraditionalMemoryBalloonDeviceConfiguration{
		VZMemoryBalloonDeviceConfiguration: VZMemoryBalloonDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass) Alloc() VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass) New() VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioTraditionalMemoryBalloonDeviceConfiguration) Init() VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioTraditionalMemoryBalloonDeviceConfiguration) Autorelease() VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioTraditionalMemoryBalloonDeviceConfiguration creates a new VZVirtioTraditionalMemoryBalloonDeviceConfiguration instance.
func NewVZVirtioTraditionalMemoryBalloonDeviceConfiguration() VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	return getVZVirtioTraditionalMemoryBalloonDeviceConfigurationClass().New()
}




