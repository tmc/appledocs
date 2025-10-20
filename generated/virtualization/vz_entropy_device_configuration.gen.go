// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZEntropyDeviceConfiguration] class.
var (
	VZEntropyDeviceConfigurationClass     _VZEntropyDeviceConfigurationClass
	VZEntropyDeviceConfigurationClassOnce sync.Once
)

func getVZEntropyDeviceConfigurationClass() _VZEntropyDeviceConfigurationClass {
	VZEntropyDeviceConfigurationClassOnce.Do(func() {
		VZEntropyDeviceConfigurationClass = _VZEntropyDeviceConfigurationClass{objc.GetClass("VZEntropyDeviceConfiguration")}
	})
	return VZEntropyDeviceConfigurationClass
}

type _VZEntropyDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZEntropyDeviceConfiguration] class.
type IVZEntropyDeviceConfiguration interface {
	objectivec.IObject
}

// The common configuration traits for entropy devices.
//
// Don’t create a VZEntropyDeviceConfiguration object directly. Instead, instantiate a subclass such as to configure a source of entropy for your virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration
type VZEntropyDeviceConfiguration struct {
	objectivec.Object
}

// VZEntropyDeviceConfigurationFrom constructs a [VZEntropyDeviceConfiguration] from an unsafe.Pointer.
//
// The common configuration traits for entropy devices.
func VZEntropyDeviceConfigurationFrom(ptr unsafe.Pointer) VZEntropyDeviceConfiguration {
	return VZEntropyDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZEntropyDeviceConfigurationClass) Alloc() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZEntropyDeviceConfigurationClass) New() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZEntropyDeviceConfiguration) Init() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZEntropyDeviceConfiguration) Autorelease() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEntropyDeviceConfiguration creates a new VZEntropyDeviceConfiguration instance.
func NewVZEntropyDeviceConfiguration() VZEntropyDeviceConfiguration {
	return getVZEntropyDeviceConfigurationClass().New()
}




