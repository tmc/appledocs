// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZMacKeyboardConfiguration] class.
var (
	VZMacKeyboardConfigurationClass     _VZMacKeyboardConfigurationClass
	VZMacKeyboardConfigurationClassOnce sync.Once
)

func getVZMacKeyboardConfigurationClass() _VZMacKeyboardConfigurationClass {
	VZMacKeyboardConfigurationClassOnce.Do(func() {
		VZMacKeyboardConfigurationClass = _VZMacKeyboardConfigurationClass{objc.GetClass("VZMacKeyboardConfiguration")}
	})
	return VZMacKeyboardConfigurationClass
}

type _VZMacKeyboardConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZMacKeyboardConfiguration] class.
type IVZMacKeyboardConfiguration interface {
	IVZKeyboardConfiguration
}

// A device that defines the configuration for a Mac keyboard.
//
// Use this configuration to attach a Mac keyboard configuration to a VM. A can use this device to send key events to the VM, including the Mac-specific key events, such as the Globe key.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacKeyboardConfiguration
type VZMacKeyboardConfiguration struct {
	VZKeyboardConfiguration
}

// VZMacKeyboardConfigurationFrom constructs a [VZMacKeyboardConfiguration] from an unsafe.Pointer.
//
// A device that defines the configuration for a Mac keyboard.
func VZMacKeyboardConfigurationFrom(ptr unsafe.Pointer) VZMacKeyboardConfiguration {
	return VZMacKeyboardConfiguration{
		VZKeyboardConfiguration: VZKeyboardConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMacKeyboardConfigurationClass) Alloc() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMacKeyboardConfigurationClass) New() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacKeyboardConfiguration) Init() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacKeyboardConfiguration) Autorelease() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacKeyboardConfiguration creates a new VZMacKeyboardConfiguration instance.
func NewVZMacKeyboardConfiguration() VZMacKeyboardConfiguration {
	return getVZMacKeyboardConfigurationClass().New()
}




