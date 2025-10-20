// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZUSBKeyboardConfiguration] class.
var (
	VZUSBKeyboardConfigurationClass     _VZUSBKeyboardConfigurationClass
	VZUSBKeyboardConfigurationClassOnce sync.Once
)

func getVZUSBKeyboardConfigurationClass() _VZUSBKeyboardConfigurationClass {
	VZUSBKeyboardConfigurationClassOnce.Do(func() {
		VZUSBKeyboardConfigurationClass = _VZUSBKeyboardConfigurationClass{objc.GetClass("VZUSBKeyboardConfiguration")}
	})
	return VZUSBKeyboardConfigurationClass
}

type _VZUSBKeyboardConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZUSBKeyboardConfiguration] class.
type IVZUSBKeyboardConfiguration interface {
	IVZKeyboardConfiguration
}

// A device that defines the configuration for a USB keyboard.
//
// A can use this device to send key events to the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBKeyboardConfiguration
type VZUSBKeyboardConfiguration struct {
	VZKeyboardConfiguration
}

// VZUSBKeyboardConfigurationFrom constructs a [VZUSBKeyboardConfiguration] from an unsafe.Pointer.
//
// A device that defines the configuration for a USB keyboard.
func VZUSBKeyboardConfigurationFrom(ptr unsafe.Pointer) VZUSBKeyboardConfiguration {
	return VZUSBKeyboardConfiguration{
		VZKeyboardConfiguration: VZKeyboardConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZUSBKeyboardConfigurationClass) Alloc() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZUSBKeyboardConfigurationClass) New() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZUSBKeyboardConfiguration) Init() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZUSBKeyboardConfiguration) Autorelease() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBKeyboardConfiguration creates a new VZUSBKeyboardConfiguration instance.
func NewVZUSBKeyboardConfiguration() VZUSBKeyboardConfiguration {
	return getVZUSBKeyboardConfigurationClass().New()
}




