// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioConsoleDeviceConfiguration */


/* debug [class_header]: Header for VZVirtioConsoleDeviceConfiguration */
// The class instance for the [VZVirtioConsoleDeviceConfiguration] class.
var (
	VZVirtioConsoleDeviceConfigurationClass     _VZVirtioConsoleDeviceConfigurationClass
	VZVirtioConsoleDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioConsoleDeviceConfigurationClass() _VZVirtioConsoleDeviceConfigurationClass {
	VZVirtioConsoleDeviceConfigurationClassOnce.Do(func() {
		VZVirtioConsoleDeviceConfigurationClass = _VZVirtioConsoleDeviceConfigurationClass{objc.GetClass("VZVirtioConsoleDeviceConfiguration")}
	})
	return VZVirtioConsoleDeviceConfigurationClass
}

type _VZVirtioConsoleDeviceConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioConsoleDeviceConfiguration */
// An interface definition for the [VZVirtioConsoleDeviceConfiguration] class.
type IVZVirtioConsoleDeviceConfiguration interface {
	IVZConsoleDeviceConfiguration
	
/* debug [class_interface_properties]: Properties for VZVirtioConsoleDeviceConfiguration */
	// properties:
	Ports() IVZVirtioConsolePortConfigurationArray
	ConsoleDevices() IVZConsoleDeviceConfiguration
	SetConsoleDevices(value IVZConsoleDeviceConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioConsoleDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioConsoleDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsoleDeviceConfigurationClass) Alloc() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioConsoleDeviceConfigurationClass) New() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsoleDeviceConfiguration) Init() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsoleDeviceConfiguration) Autorelease() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsoleDeviceConfiguration creates a new VZVirtioConsoleDeviceConfiguration instance.
func NewVZVirtioConsoleDeviceConfiguration() VZVirtioConsoleDeviceConfiguration {
	return getVZVirtioConsoleDeviceConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioConsoleDeviceConfiguration */
// A console device that enables communication between the host and the guest using console ports through a Virtio interface.
//
// A object enables serial communication between the guest-operating system and host computer through the Virtio interface. The device sets up one or more ports through on the Virtio console device.


// A console device that enables communication between the host and the guest using console ports through a Virtio interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceConfiguration
type VZVirtioConsoleDeviceConfiguration struct {
	VZConsoleDeviceConfiguration
}

// VZVirtioConsoleDeviceConfigurationFrom constructs a [VZVirtioConsoleDeviceConfiguration] from an unsafe.Pointer.
//
// A console device that enables communication between the host and the guest using console ports through a Virtio interface.
func VZVirtioConsoleDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioConsoleDeviceConfiguration {
	return VZVirtioConsoleDeviceConfiguration{
		VZConsoleDeviceConfiguration: VZConsoleDeviceConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioConsoleDeviceConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioConsoleDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioConsoleDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioConsoleDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioConsoleDeviceConfiguration */

// The list of Virtio port configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceConfiguration/ports
func (v_ VZVirtioConsoleDeviceConfiguration) Ports() IVZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](v_.ID, objc.Sel("ports"))
	return rv
}/* debug [instance_properties/getter]: ports */


// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/consoledevices
func (v_ VZVirtioConsoleDeviceConfiguration) ConsoleDevices() IVZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](v_.ID, objc.Sel("consoleDevices"))
	return rv
}/* debug [instance_properties/getter]: consoleDevices */


// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/consoledevices
func (v_ VZVirtioConsoleDeviceConfiguration) SetConsoleDevices(value IVZConsoleDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setConsoleDevices:"), value)
}/* debug [instance_properties/setter]: consoleDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioConsoleDeviceConfiguration */


