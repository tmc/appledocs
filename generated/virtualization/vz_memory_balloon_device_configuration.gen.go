// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZMemoryBalloonDeviceConfiguration] class.
var (
	VZMemoryBalloonDeviceConfigurationClass     _VZMemoryBalloonDeviceConfigurationClass
	VZMemoryBalloonDeviceConfigurationClassOnce sync.Once
)

func getVZMemoryBalloonDeviceConfigurationClass() _VZMemoryBalloonDeviceConfigurationClass {
	VZMemoryBalloonDeviceConfigurationClassOnce.Do(func() {
		VZMemoryBalloonDeviceConfigurationClass = _VZMemoryBalloonDeviceConfigurationClass{objc.GetClass("VZMemoryBalloonDeviceConfiguration")}
	})
	return VZMemoryBalloonDeviceConfigurationClass
}

type _VZMemoryBalloonDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZMemoryBalloonDeviceConfiguration] class.
type IVZMemoryBalloonDeviceConfiguration interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The common configuration traits for memory balloon devices.
//
// Don’t instantiate this abstract class directly. Instead, instantiate one of its subclasses such as .


// The common configuration traits for memory balloon devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration
type VZMemoryBalloonDeviceConfiguration struct {
	objectivec.Object
}

// VZMemoryBalloonDeviceConfigurationFrom constructs a [VZMemoryBalloonDeviceConfiguration] from an unsafe.Pointer.
//
// The common configuration traits for memory balloon devices.
func VZMemoryBalloonDeviceConfigurationFrom(ptr unsafe.Pointer) VZMemoryBalloonDeviceConfiguration {
	return VZMemoryBalloonDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMemoryBalloonDeviceConfigurationClass) Alloc() VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMemoryBalloonDeviceConfigurationClass) New() VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMemoryBalloonDeviceConfiguration) Init() VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMemoryBalloonDeviceConfiguration) Autorelease() VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMemoryBalloonDeviceConfiguration creates a new VZMemoryBalloonDeviceConfiguration instance.
func NewVZMemoryBalloonDeviceConfiguration() VZMemoryBalloonDeviceConfiguration {
	return getVZMemoryBalloonDeviceConfigurationClass().New()
}




