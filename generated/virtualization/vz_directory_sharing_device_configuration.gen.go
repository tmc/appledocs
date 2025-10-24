// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZDirectorySharingDeviceConfiguration */


/* debug [class_header]: Header for VZDirectorySharingDeviceConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZDirectorySharingDeviceConfiguration */
// An interface definition for the [VZDirectorySharingDeviceConfiguration] class.
type IVZDirectorySharingDeviceConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZDirectorySharingDeviceConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZDirectorySharingDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZDirectorySharingDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZDirectorySharingDeviceConfigurationClass) Alloc() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZDirectorySharingDeviceConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZDirectorySharingDeviceConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZDirectorySharingDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZDirectorySharingDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZDirectorySharingDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZDirectorySharingDeviceConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZDirectorySharingDeviceConfiguration */



