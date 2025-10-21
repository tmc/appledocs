// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZSocketDeviceConfiguration] class.
var (
	VZSocketDeviceConfigurationClass     _VZSocketDeviceConfigurationClass
	VZSocketDeviceConfigurationClassOnce sync.Once
)

func getVZSocketDeviceConfigurationClass() _VZSocketDeviceConfigurationClass {
	VZSocketDeviceConfigurationClassOnce.Do(func() {
		VZSocketDeviceConfigurationClass = _VZSocketDeviceConfigurationClass{objc.GetClass("VZSocketDeviceConfiguration")}
	})
	return VZSocketDeviceConfigurationClass
}

type _VZSocketDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZSocketDeviceConfiguration] class.
type IVZSocketDeviceConfiguration interface {
	objectivec.IObject
}

// The common configuration traits for socket device requests.
//
// Don’t create a object directly. Instead, create a object and add it to your virtual machine’s configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSocketDeviceConfiguration
type VZSocketDeviceConfiguration struct {
	objectivec.Object
}

// VZSocketDeviceConfigurationFrom constructs a [VZSocketDeviceConfiguration] from an unsafe.Pointer.
//
// The common configuration traits for socket device requests.
func VZSocketDeviceConfigurationFrom(ptr unsafe.Pointer) VZSocketDeviceConfiguration {
	return VZSocketDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZSocketDeviceConfigurationClass) Alloc() VZSocketDeviceConfiguration {
	rv := objc.Send[VZSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZSocketDeviceConfigurationClass) New() VZSocketDeviceConfiguration {
	rv := objc.Send[VZSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZSocketDeviceConfiguration) Init() VZSocketDeviceConfiguration {
	rv := objc.Send[VZSocketDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZSocketDeviceConfiguration) Autorelease() VZSocketDeviceConfiguration {
	rv := objc.Send[VZSocketDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSocketDeviceConfiguration creates a new VZSocketDeviceConfiguration instance.
func NewVZSocketDeviceConfiguration() VZSocketDeviceConfiguration {
	return getVZSocketDeviceConfigurationClass().New()
}




