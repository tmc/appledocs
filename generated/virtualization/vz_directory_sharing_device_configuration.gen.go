// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZDirectorySharingDeviceConfiguration] class.
var (
	VZDirectorySharingDeviceConfigurationClass     _VZDirectorySharingDeviceConfigurationClass
	VZDirectorySharingDeviceConfigurationClassOnce sync.Once
)

func getVZDirectorySharingDeviceConfigurationClass() _VZDirectorySharingDeviceConfigurationClass {
	VZDirectorySharingDeviceConfigurationClassOnce.Do(func() {
		VZDirectorySharingDeviceConfigurationClass = _VZDirectorySharingDeviceConfigurationClass{objc.GetClass("VZDirectorySharingDeviceConfiguration")}
	})
	return VZDirectorySharingDeviceConfigurationClass
}

type _VZDirectorySharingDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZDirectorySharingDeviceConfiguration] class.
type IVZDirectorySharingDeviceConfiguration interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The base class for a directory sharing device configuration.
//
// Don’t instantiate directly. Instead use one of its subclasses, like .


// The base class for a directory sharing device configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration
type VZDirectorySharingDeviceConfiguration struct {
	objectivec.Object
}

// VZDirectorySharingDeviceConfigurationFrom constructs a [VZDirectorySharingDeviceConfiguration] from an unsafe.Pointer.
//
// The base class for a directory sharing device configuration.
func VZDirectorySharingDeviceConfigurationFrom(ptr unsafe.Pointer) VZDirectorySharingDeviceConfiguration {
	return VZDirectorySharingDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZDirectorySharingDeviceConfigurationClass) Alloc() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZDirectorySharingDeviceConfigurationClass) New() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZDirectorySharingDeviceConfiguration) Init() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZDirectorySharingDeviceConfiguration) Autorelease() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectorySharingDeviceConfiguration creates a new VZDirectorySharingDeviceConfiguration instance.
func NewVZDirectorySharingDeviceConfiguration() VZDirectorySharingDeviceConfiguration {
	return getVZDirectorySharingDeviceConfigurationClass().New()
}




