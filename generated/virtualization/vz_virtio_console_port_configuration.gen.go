// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioConsolePortConfiguration */


/* debug [class_header]: Header for VZVirtioConsolePortConfiguration */
// The class instance for the [VZVirtioConsolePortConfiguration] class.
var (
	VZVirtioConsolePortConfigurationClass     _VZVirtioConsolePortConfigurationClass
	VZVirtioConsolePortConfigurationClassOnce sync.Once
)

func getVZVirtioConsolePortConfigurationClass() _VZVirtioConsolePortConfigurationClass {
	VZVirtioConsolePortConfigurationClassOnce.Do(func() {
		VZVirtioConsolePortConfigurationClass = _VZVirtioConsolePortConfigurationClass{objc.GetClass("VZVirtioConsolePortConfiguration")}
	})
	return VZVirtioConsolePortConfigurationClass
}

type _VZVirtioConsolePortConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioConsolePortConfiguration */
// An interface definition for the [VZVirtioConsolePortConfiguration] class.
type IVZVirtioConsolePortConfiguration interface {
	IVZConsolePortConfiguration
	
/* debug [class_interface_properties]: Properties for VZVirtioConsolePortConfiguration */
	// properties:
	IsConsole() bool
	SetIsConsole(value bool)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	ConsoleDevices() IVZConsoleDeviceConfiguration
	SetConsoleDevices(value IVZConsoleDeviceConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioConsolePortConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioConsolePortConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsolePortConfigurationClass) Alloc() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioConsolePortConfigurationClass) New() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsolePortConfiguration) Init() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsolePortConfiguration) Autorelease() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePortConfiguration creates a new VZVirtioConsolePortConfiguration instance.
func NewVZVirtioConsolePortConfiguration() VZVirtioConsolePortConfiguration {
	return getVZVirtioConsolePortConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioConsolePortConfiguration */
// A class that represents the configuration options you can set on a Virtio console port.
//
// A console port is a two-way communication channel between a host and a VM console port. A Virtio device can have one or more attached console devices. Optionally, you can set a name for a console port and also configure a console port that the guest can use as the system console.


// A class that represents the configuration options you can set on a Virtio console port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration
type VZVirtioConsolePortConfiguration struct {
	VZConsolePortConfiguration
}

// VZVirtioConsolePortConfigurationFrom constructs a [VZVirtioConsolePortConfiguration] from an unsafe.Pointer.
//
// A class that represents the configuration options you can set on a Virtio console port.
func VZVirtioConsolePortConfigurationFrom(ptr unsafe.Pointer) VZVirtioConsolePortConfiguration {
	return VZVirtioConsolePortConfiguration{
		VZConsolePortConfiguration: VZConsolePortConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioConsolePortConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioConsolePortConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioConsolePortConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioConsolePortConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioConsolePortConfiguration */

// A Boolean value that indicates whether this port is a console.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration/isConsole
func (v_ VZVirtioConsolePortConfiguration) IsConsole() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isConsole"))
	return rv
}/* debug [instance_properties/getter]: isConsole */


// A Boolean value that indicates whether this port is a console.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration/isConsole
func (v_ VZVirtioConsolePortConfiguration) SetIsConsole(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsConsole:"), value)
}/* debug [instance_properties/setter]: isConsole */


// The name of the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration/name
func (v_ VZVirtioConsolePortConfiguration) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration/name
func (v_ VZVirtioConsolePortConfiguration) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/consoledevices
func (v_ VZVirtioConsolePortConfiguration) ConsoleDevices() IVZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](v_.ID, objc.Sel("consoleDevices"))
	return rv
}/* debug [instance_properties/getter]: consoleDevices */


// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/consoledevices
func (v_ VZVirtioConsolePortConfiguration) SetConsoleDevices(value IVZConsoleDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setConsoleDevices:"), value)
}/* debug [instance_properties/setter]: consoleDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioConsolePortConfiguration */


