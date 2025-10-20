// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZGraphicsDeviceConfiguration] class.
var (
	VZGraphicsDeviceConfigurationClass     _VZGraphicsDeviceConfigurationClass
	VZGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZGraphicsDeviceConfigurationClass() _VZGraphicsDeviceConfigurationClass {
	VZGraphicsDeviceConfigurationClassOnce.Do(func() {
		VZGraphicsDeviceConfigurationClass = _VZGraphicsDeviceConfigurationClass{objc.GetClass("VZGraphicsDeviceConfiguration")}
	})
	return VZGraphicsDeviceConfigurationClass
}

type _VZGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZGraphicsDeviceConfiguration] class.
type IVZGraphicsDeviceConfiguration interface {
	objectivec.IObject
}

// The base class for a graphics device configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration
type VZGraphicsDeviceConfiguration struct {
	objectivec.Object
}

// VZGraphicsDeviceConfigurationFrom constructs a [VZGraphicsDeviceConfiguration] from an unsafe.Pointer.
//
// The base class for a graphics device configuration.
func VZGraphicsDeviceConfigurationFrom(ptr unsafe.Pointer) VZGraphicsDeviceConfiguration {
	return VZGraphicsDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZGraphicsDeviceConfigurationClass) Alloc() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZGraphicsDeviceConfigurationClass) New() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZGraphicsDeviceConfiguration) Init() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZGraphicsDeviceConfiguration) Autorelease() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDeviceConfiguration creates a new VZGraphicsDeviceConfiguration instance.
func NewVZGraphicsDeviceConfiguration() VZGraphicsDeviceConfiguration {
	return getVZGraphicsDeviceConfigurationClass().New()
}




