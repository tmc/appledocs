// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZSerialPortConfiguration */


/* debug [class_header]: Header for VZSerialPortConfiguration */
// The class instance for the [VZSerialPortConfiguration] class.
var (
	VZSerialPortConfigurationClass     _VZSerialPortConfigurationClass
	VZSerialPortConfigurationClassOnce sync.Once
)

func getVZSerialPortConfigurationClass() _VZSerialPortConfigurationClass {
	VZSerialPortConfigurationClassOnce.Do(func() {
		VZSerialPortConfigurationClass = _VZSerialPortConfigurationClass{objc.GetClass("VZSerialPortConfiguration")}
	})
	return VZSerialPortConfigurationClass
}

type _VZSerialPortConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZSerialPortConfiguration */
// An interface definition for the [VZSerialPortConfiguration] class.
type IVZSerialPortConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZSerialPortConfiguration */
	// properties:
	Attachment() IVZSerialPortAttachment
	SetAttachment(value IVZSerialPortAttachment)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZSerialPortConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZSerialPortConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZSerialPortConfigurationClass) Alloc() VZSerialPortConfiguration {
	rv := objc.Send[VZSerialPortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZSerialPortConfigurationClass) New() VZSerialPortConfiguration {
	rv := objc.Send[VZSerialPortConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZSerialPortConfiguration) Init() VZSerialPortConfiguration {
	rv := objc.Send[VZSerialPortConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZSerialPortConfiguration) Autorelease() VZSerialPortConfiguration {
	rv := objc.Send[VZSerialPortConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSerialPortConfiguration creates a new VZSerialPortConfiguration instance.
func NewVZSerialPortConfiguration() VZSerialPortConfiguration {
	return getVZSerialPortConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZSerialPortConfiguration */
// The common configuration traits for serial port requests.
//
// Don’t create a object directly. Instead, instantiate a concrete instance of one of its subclasses, such as . Use the property of this class to configure the medium through which serial communication happens.


// The common configuration traits for serial port requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration
type VZSerialPortConfiguration struct {
	objectivec.Object
}

// VZSerialPortConfigurationFrom constructs a [VZSerialPortConfiguration] from an unsafe.Pointer.
//
// The common configuration traits for serial port requests.
func VZSerialPortConfigurationFrom(ptr unsafe.Pointer) VZSerialPortConfiguration {
	return VZSerialPortConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZSerialPortConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZSerialPortConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZSerialPortConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZSerialPortConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZSerialPortConfiguration */

// The object that defines how the configuration of the virtual machine’s serial port interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/attachment
func (v_ VZSerialPortConfiguration) Attachment() IVZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}/* debug [instance_properties/getter]: attachment */


// The object that defines how the configuration of the virtual machine’s serial port interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/attachment
func (v_ VZSerialPortConfiguration) SetAttachment(value IVZSerialPortAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}/* debug [instance_properties/setter]: attachment */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZSerialPortConfiguration */



