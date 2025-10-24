// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZKeyboardConfiguration */


/* debug [class_header]: Header for VZKeyboardConfiguration */
// The class instance for the [VZKeyboardConfiguration] class.
var (
	VZKeyboardConfigurationClass     _VZKeyboardConfigurationClass
	VZKeyboardConfigurationClassOnce sync.Once
)

func getVZKeyboardConfigurationClass() _VZKeyboardConfigurationClass {
	VZKeyboardConfigurationClassOnce.Do(func() {
		VZKeyboardConfigurationClass = _VZKeyboardConfigurationClass{objc.GetClass("VZKeyboardConfiguration")}
	})
	return VZKeyboardConfigurationClass
}

type _VZKeyboardConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZKeyboardConfiguration */
// An interface definition for the [VZKeyboardConfiguration] class.
type IVZKeyboardConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZKeyboardConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZKeyboardConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZKeyboardConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZKeyboardConfigurationClass) Alloc() VZKeyboardConfiguration {
	rv := objc.Send[VZKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZKeyboardConfigurationClass) New() VZKeyboardConfiguration {
	rv := objc.Send[VZKeyboardConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZKeyboardConfiguration) Init() VZKeyboardConfiguration {
	rv := objc.Send[VZKeyboardConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZKeyboardConfiguration) Autorelease() VZKeyboardConfiguration {
	rv := objc.Send[VZKeyboardConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZKeyboardConfiguration creates a new VZKeyboardConfiguration instance.
func NewVZKeyboardConfiguration() VZKeyboardConfiguration {
	return getVZKeyboardConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZKeyboardConfiguration */
// The base class for a configuring a keyboard.
//
// defines the abstract interface that defines a virtual keyboard that you connect to a guest operating system. Don’t instantiate directly, use one of its subclasses such as instead.


// The base class for a configuring a keyboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration
type VZKeyboardConfiguration struct {
	objectivec.Object
}

// VZKeyboardConfigurationFrom constructs a [VZKeyboardConfiguration] from an unsafe.Pointer.
//
// The base class for a configuring a keyboard.
func VZKeyboardConfigurationFrom(ptr unsafe.Pointer) VZKeyboardConfiguration {
	return VZKeyboardConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZKeyboardConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZKeyboardConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZKeyboardConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZKeyboardConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZKeyboardConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZKeyboardConfiguration */



