// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZConsoleDeviceConfiguration] class.
var (
	VZConsoleDeviceConfigurationClass     _VZConsoleDeviceConfigurationClass
	VZConsoleDeviceConfigurationClassOnce sync.Once
)

func getVZConsoleDeviceConfigurationClass() _VZConsoleDeviceConfigurationClass {
	VZConsoleDeviceConfigurationClassOnce.Do(func() {
		VZConsoleDeviceConfigurationClass = _VZConsoleDeviceConfigurationClass{objc.GetClass("VZConsoleDeviceConfiguration")}
	})
	return VZConsoleDeviceConfigurationClass
}

type _VZConsoleDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZConsoleDeviceConfiguration] class.
type IVZConsoleDeviceConfiguration interface {
	objectivec.IObject
}

// The base class for a console device configuration.
//
// Don’t instantiate VZConsoleDeviceConfiguration directly, instead use one of its subclasses like instead.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration
type VZConsoleDeviceConfiguration struct {
	objectivec.Object
}

// VZConsoleDeviceConfigurationFrom constructs a [VZConsoleDeviceConfiguration] from an unsafe.Pointer.
//
// The base class for a console device configuration.
func VZConsoleDeviceConfigurationFrom(ptr unsafe.Pointer) VZConsoleDeviceConfiguration {
	return VZConsoleDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZConsoleDeviceConfigurationClass) Alloc() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZConsoleDeviceConfigurationClass) New() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZConsoleDeviceConfiguration) Init() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZConsoleDeviceConfiguration) Autorelease() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsoleDeviceConfiguration creates a new VZConsoleDeviceConfiguration instance.
func NewVZConsoleDeviceConfiguration() VZConsoleDeviceConfiguration {
	return getVZConsoleDeviceConfigurationClass().New()
}




