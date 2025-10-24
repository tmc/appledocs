// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZSocketDeviceConfiguration */


/* debug [class_header]: Header for VZSocketDeviceConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZSocketDeviceConfiguration */
// An interface definition for the [VZSocketDeviceConfiguration] class.
type IVZSocketDeviceConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZSocketDeviceConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZSocketDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZSocketDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZSocketDeviceConfigurationClass) Alloc() VZSocketDeviceConfiguration {
	rv := objc.Send[VZSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZSocketDeviceConfiguration */
// The common configuration traits for socket device requests.
//
// Don’t create a object directly. Instead, create a object and add it to your virtual machine’s configuration.


// The common configuration traits for socket device requests.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZSocketDeviceConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZSocketDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZSocketDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZSocketDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZSocketDeviceConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZSocketDeviceConfiguration */



