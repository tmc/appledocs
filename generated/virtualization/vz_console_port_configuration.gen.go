// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZConsolePortConfiguration */


/* debug [class_header]: Header for VZConsolePortConfiguration */
// The class instance for the [VZConsolePortConfiguration] class.
var (
	VZConsolePortConfigurationClass     _VZConsolePortConfigurationClass
	VZConsolePortConfigurationClassOnce sync.Once
)

func getVZConsolePortConfigurationClass() _VZConsolePortConfigurationClass {
	VZConsolePortConfigurationClassOnce.Do(func() {
		VZConsolePortConfigurationClass = _VZConsolePortConfigurationClass{objc.GetClass("VZConsolePortConfiguration")}
	})
	return VZConsolePortConfigurationClass
}

type _VZConsolePortConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZConsolePortConfiguration */
// An interface definition for the [VZConsolePortConfiguration] class.
type IVZConsolePortConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZConsolePortConfiguration */
	// properties:
	Attachment() IVZSerialPortAttachment
	SetAttachment(value IVZSerialPortAttachment)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZConsolePortConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZConsolePortConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZConsolePortConfigurationClass) Alloc() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZConsolePortConfigurationClass) New() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZConsolePortConfiguration) Init() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZConsolePortConfiguration) Autorelease() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsolePortConfiguration creates a new VZConsolePortConfiguration instance.
func NewVZConsolePortConfiguration() VZConsolePortConfiguration {
	return getVZConsolePortConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZConsolePortConfiguration */
// The base class for a console port configuration.
//
// Don’t instantiate directly, instead use one of its subclasses like .


// The base class for a console port configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration
type VZConsolePortConfiguration struct {
	objectivec.Object
}

// VZConsolePortConfigurationFrom constructs a [VZConsolePortConfiguration] from an unsafe.Pointer.
//
// The base class for a console port configuration.
func VZConsolePortConfigurationFrom(ptr unsafe.Pointer) VZConsolePortConfiguration {
	return VZConsolePortConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZConsolePortConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZConsolePortConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZConsolePortConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZConsolePortConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZConsolePortConfiguration */

// The serial port attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration/attachment
func (v_ VZConsolePortConfiguration) Attachment() IVZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}/* debug [instance_properties/getter]: attachment */


// The serial port attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration/attachment
func (v_ VZConsolePortConfiguration) SetAttachment(value IVZSerialPortAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}/* debug [instance_properties/setter]: attachment */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZConsolePortConfiguration */



