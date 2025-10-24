// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZGraphicsDisplayConfiguration] class.
var (
	VZGraphicsDisplayConfigurationClass     _VZGraphicsDisplayConfigurationClass
	VZGraphicsDisplayConfigurationClassOnce sync.Once
)

func getVZGraphicsDisplayConfigurationClass() _VZGraphicsDisplayConfigurationClass {
	VZGraphicsDisplayConfigurationClassOnce.Do(func() {
		VZGraphicsDisplayConfigurationClass = _VZGraphicsDisplayConfigurationClass{objc.GetClass("VZGraphicsDisplayConfiguration")}
	})
	return VZGraphicsDisplayConfigurationClass
}

type _VZGraphicsDisplayConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZGraphicsDisplayConfiguration] class.
type IVZGraphicsDisplayConfiguration interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The base class for a graphics display configuration.
//
// Don’t instantiate directly. Use one of its subclasses instead.


// The base class for a graphics display configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayConfiguration
type VZGraphicsDisplayConfiguration struct {
	objectivec.Object
}

// VZGraphicsDisplayConfigurationFrom constructs a [VZGraphicsDisplayConfiguration] from an unsafe.Pointer.
//
// The base class for a graphics display configuration.
func VZGraphicsDisplayConfigurationFrom(ptr unsafe.Pointer) VZGraphicsDisplayConfiguration {
	return VZGraphicsDisplayConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZGraphicsDisplayConfigurationClass) Alloc() VZGraphicsDisplayConfiguration {
	rv := objc.Send[VZGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZGraphicsDisplayConfigurationClass) New() VZGraphicsDisplayConfiguration {
	rv := objc.Send[VZGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZGraphicsDisplayConfiguration) Init() VZGraphicsDisplayConfiguration {
	rv := objc.Send[VZGraphicsDisplayConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZGraphicsDisplayConfiguration) Autorelease() VZGraphicsDisplayConfiguration {
	rv := objc.Send[VZGraphicsDisplayConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDisplayConfiguration creates a new VZGraphicsDisplayConfiguration instance.
func NewVZGraphicsDisplayConfiguration() VZGraphicsDisplayConfiguration {
	return getVZGraphicsDisplayConfigurationClass().New()
}




