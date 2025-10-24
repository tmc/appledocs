// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZPlatformConfiguration */


/* debug [class_header]: Header for VZPlatformConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZPlatformConfiguration */
// An interface definition for the [VZPlatformConfiguration] class.
type IVZPlatformConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZPlatformConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZPlatformConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZPlatformConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZPlatformConfigurationClass) Alloc() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZPlatformConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZPlatformConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZPlatformConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZPlatformConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZPlatformConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZPlatformConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZPlatformConfiguration */



