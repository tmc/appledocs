// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZGraphicsDisplayConfiguration */


/* debug [class_header]: Header for VZGraphicsDisplayConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZGraphicsDisplayConfiguration */
// An interface definition for the [VZGraphicsDisplayConfiguration] class.
type IVZGraphicsDisplayConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZGraphicsDisplayConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZGraphicsDisplayConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZGraphicsDisplayConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZGraphicsDisplayConfigurationClass) Alloc() VZGraphicsDisplayConfiguration {
	rv := objc.Send[VZGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZGraphicsDisplayConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZGraphicsDisplayConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZGraphicsDisplayConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZGraphicsDisplayConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZGraphicsDisplayConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZGraphicsDisplayConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZGraphicsDisplayConfiguration */



