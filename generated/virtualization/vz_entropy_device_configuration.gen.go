// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZEntropyDeviceConfiguration */


/* debug [class_header]: Header for VZEntropyDeviceConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZEntropyDeviceConfiguration */
// An interface definition for the [VZEntropyDeviceConfiguration] class.
type IVZEntropyDeviceConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZEntropyDeviceConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZEntropyDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZEntropyDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZEntropyDeviceConfigurationClass) Alloc() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZEntropyDeviceConfiguration */
// The common configuration traits for entropy devices.
//
// Don’t create a VZEntropyDeviceConfiguration object directly. Instead, instantiate a subclass such as to configure a source of entropy for your virtual machine.


// The common configuration traits for entropy devices.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZEntropyDeviceConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZEntropyDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZEntropyDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZEntropyDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZEntropyDeviceConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZEntropyDeviceConfiguration */



