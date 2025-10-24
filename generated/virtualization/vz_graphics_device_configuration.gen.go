// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZGraphicsDeviceConfiguration */

/* debug [class_header]: Header for VZGraphicsDeviceConfiguration */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZGraphicsDeviceConfiguration */
// An interface definition for the [VZGraphicsDeviceConfiguration] class.
type IVZGraphicsDeviceConfiguration interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZGraphicsDeviceConfiguration */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZGraphicsDeviceConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZGraphicsDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZGraphicsDeviceConfigurationClass) Alloc() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZGraphicsDeviceConfiguration */
// The base class for a graphics device configuration.

// The base class for a graphics device configuration.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZGraphicsDeviceConfiguration */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZGraphicsDeviceConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZGraphicsDeviceConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZGraphicsDeviceConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZGraphicsDeviceConfiguration */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZGraphicsDeviceConfiguration */
