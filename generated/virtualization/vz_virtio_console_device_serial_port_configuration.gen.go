// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioConsoleDeviceSerialPortConfiguration */


/* debug [class_header]: Header for VZVirtioConsoleDeviceSerialPortConfiguration */
// The class instance for the [VZVirtioConsoleDeviceSerialPortConfiguration] class.
var (
	VZVirtioConsoleDeviceSerialPortConfigurationClass     _VZVirtioConsoleDeviceSerialPortConfigurationClass
	VZVirtioConsoleDeviceSerialPortConfigurationClassOnce sync.Once
)

func getVZVirtioConsoleDeviceSerialPortConfigurationClass() _VZVirtioConsoleDeviceSerialPortConfigurationClass {
	VZVirtioConsoleDeviceSerialPortConfigurationClassOnce.Do(func() {
		VZVirtioConsoleDeviceSerialPortConfigurationClass = _VZVirtioConsoleDeviceSerialPortConfigurationClass{objc.GetClass("VZVirtioConsoleDeviceSerialPortConfiguration")}
	})
	return VZVirtioConsoleDeviceSerialPortConfigurationClass
}

type _VZVirtioConsoleDeviceSerialPortConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioConsoleDeviceSerialPortConfiguration */
// An interface definition for the [VZVirtioConsoleDeviceSerialPortConfiguration] class.
type IVZVirtioConsoleDeviceSerialPortConfiguration interface {
	IVZSerialPortConfiguration
	
/* debug [class_interface_properties]: Properties for VZVirtioConsoleDeviceSerialPortConfiguration */
	// properties:
	Attachment() IVZSerialPortAttachment
	SetAttachment(value IVZSerialPortAttachment)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioConsoleDeviceSerialPortConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioConsoleDeviceSerialPortConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsoleDeviceSerialPortConfigurationClass) Alloc() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioConsoleDeviceSerialPortConfigurationClass) New() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsoleDeviceSerialPortConfiguration) Init() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsoleDeviceSerialPortConfiguration) Autorelease() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsoleDeviceSerialPortConfiguration creates a new VZVirtioConsoleDeviceSerialPortConfiguration instance.
func NewVZVirtioConsoleDeviceSerialPortConfiguration() VZVirtioConsoleDeviceSerialPortConfiguration {
	return getVZVirtioConsoleDeviceSerialPortConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioConsoleDeviceSerialPortConfiguration */
// A configuration object that requests the creation of a console device to communicate with the guest system.
//
// A object enables serial communication between the guest operating system and host computer through the Virtio interface. After you create this configuration object, configure its inherited property with an object that defines the type of serial communication you want to enable. Use a object to enable two-way communication between the guest and host, and use a object to enable one-way communication from the guest to the file you designate.


// A configuration object that requests the creation of a console device to communicate with the guest system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceSerialPortConfiguration
type VZVirtioConsoleDeviceSerialPortConfiguration struct {
	VZSerialPortConfiguration
}

// VZVirtioConsoleDeviceSerialPortConfigurationFrom constructs a [VZVirtioConsoleDeviceSerialPortConfiguration] from an unsafe.Pointer.
//
// A configuration object that requests the creation of a console device to communicate with the guest system.
func VZVirtioConsoleDeviceSerialPortConfigurationFrom(ptr unsafe.Pointer) VZVirtioConsoleDeviceSerialPortConfiguration {
	return VZVirtioConsoleDeviceSerialPortConfiguration{
		VZSerialPortConfiguration: VZSerialPortConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioConsoleDeviceSerialPortConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioConsoleDeviceSerialPortConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioConsoleDeviceSerialPortConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioConsoleDeviceSerialPortConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioConsoleDeviceSerialPortConfiguration */

// The object that defines how the configuration of the virtual machine’s serial port interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzserialportconfiguration/attachment
func (v_ VZVirtioConsoleDeviceSerialPortConfiguration) Attachment() IVZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}/* debug [instance_properties/getter]: attachment */


// The object that defines how the configuration of the virtual machine’s serial port interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzserialportconfiguration/attachment
func (v_ VZVirtioConsoleDeviceSerialPortConfiguration) SetAttachment(value IVZSerialPortAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}/* debug [instance_properties/setter]: attachment */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioConsoleDeviceSerialPortConfiguration */


