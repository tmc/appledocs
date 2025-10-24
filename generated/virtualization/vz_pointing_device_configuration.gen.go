// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZPointingDeviceConfiguration] class.
var (
	VZPointingDeviceConfigurationClass     _VZPointingDeviceConfigurationClass
	VZPointingDeviceConfigurationClassOnce sync.Once
)

func getVZPointingDeviceConfigurationClass() _VZPointingDeviceConfigurationClass {
	VZPointingDeviceConfigurationClassOnce.Do(func() {
		VZPointingDeviceConfigurationClass = _VZPointingDeviceConfigurationClass{objc.GetClass("VZPointingDeviceConfiguration")}
	})
	return VZPointingDeviceConfigurationClass
}

type _VZPointingDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZPointingDeviceConfiguration] class.
type IVZPointingDeviceConfiguration interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The base class for a pointing device configuration.
//
// Don’t instantiate a directly, use one of its subclasses like instead.


// The base class for a pointing device configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration
type VZPointingDeviceConfiguration struct {
	objectivec.Object
}

// VZPointingDeviceConfigurationFrom constructs a [VZPointingDeviceConfiguration] from an unsafe.Pointer.
//
// The base class for a pointing device configuration.
func VZPointingDeviceConfigurationFrom(ptr unsafe.Pointer) VZPointingDeviceConfiguration {
	return VZPointingDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZPointingDeviceConfigurationClass) Alloc() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZPointingDeviceConfigurationClass) New() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZPointingDeviceConfiguration) Init() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZPointingDeviceConfiguration) Autorelease() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPointingDeviceConfiguration creates a new VZPointingDeviceConfiguration instance.
func NewVZPointingDeviceConfiguration() VZPointingDeviceConfiguration {
	return getVZPointingDeviceConfigurationClass().New()
}




