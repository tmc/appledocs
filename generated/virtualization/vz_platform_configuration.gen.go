// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZPlatformConfiguration] class.
var (
	VZPlatformConfigurationClass     _VZPlatformConfigurationClass
	VZPlatformConfigurationClassOnce sync.Once
)

func getVZPlatformConfigurationClass() _VZPlatformConfigurationClass {
	VZPlatformConfigurationClassOnce.Do(func() {
		VZPlatformConfigurationClass = _VZPlatformConfigurationClass{objc.GetClass("VZPlatformConfiguration")}
	})
	return VZPlatformConfigurationClass
}

type _VZPlatformConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZPlatformConfiguration] class.
type IVZPlatformConfiguration interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The base class for a platform configuration.
//
// Don’t instantiate directly , use one of its subclasses, such as or instead.


// The base class for a platform configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration
type VZPlatformConfiguration struct {
	objectivec.Object
}

// VZPlatformConfigurationFrom constructs a [VZPlatformConfiguration] from an unsafe.Pointer.
//
// The base class for a platform configuration.
func VZPlatformConfigurationFrom(ptr unsafe.Pointer) VZPlatformConfiguration {
	return VZPlatformConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZPlatformConfigurationClass) Alloc() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZPlatformConfigurationClass) New() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZPlatformConfiguration) Init() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZPlatformConfiguration) Autorelease() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPlatformConfiguration creates a new VZPlatformConfiguration instance.
func NewVZPlatformConfiguration() VZPlatformConfiguration {
	return getVZPlatformConfigurationClass().New()
}




