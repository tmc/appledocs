// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioSocketDeviceConfiguration */


/* debug [class_header]: Header for VZVirtioSocketDeviceConfiguration */
// The class instance for the [VZVirtioSocketDeviceConfiguration] class.
var (
	VZVirtioSocketDeviceConfigurationClass     _VZVirtioSocketDeviceConfigurationClass
	VZVirtioSocketDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioSocketDeviceConfigurationClass() _VZVirtioSocketDeviceConfigurationClass {
	VZVirtioSocketDeviceConfigurationClassOnce.Do(func() {
		VZVirtioSocketDeviceConfigurationClass = _VZVirtioSocketDeviceConfigurationClass{objc.GetClass("VZVirtioSocketDeviceConfiguration")}
	})
	return VZVirtioSocketDeviceConfigurationClass
}

type _VZVirtioSocketDeviceConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioSocketDeviceConfiguration */
// An interface definition for the [VZVirtioSocketDeviceConfiguration] class.
type IVZVirtioSocketDeviceConfiguration interface {
	IVZSocketDeviceConfiguration
	
/* debug [class_interface_properties]: Properties for VZVirtioSocketDeviceConfiguration */
	// properties:
	SocketDevices() IVZSocketDeviceConfiguration
	SetSocketDevices(value IVZSocketDeviceConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioSocketDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioSocketDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSocketDeviceConfigurationClass) Alloc() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioSocketDeviceConfigurationClass) New() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioSocketDeviceConfiguration) Init() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioSocketDeviceConfiguration) Autorelease() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketDeviceConfiguration creates a new VZVirtioSocketDeviceConfiguration instance.
func NewVZVirtioSocketDeviceConfiguration() VZVirtioSocketDeviceConfiguration {
	return getVZVirtioSocketDeviceConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioSocketDeviceConfiguration */
// A configuration object that requests the creation of a socket device to communicate with the guest system.
//
// Use a object to implement port-based communication between the guest operating system and the host computer. When you add this object to the property of your , the virtual machine provides a corresponding object to use to configure the ports. Add only one to your virtual machine’s configuration.


// A configuration object that requests the creation of a socket device to communicate with the guest system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDeviceConfiguration
type VZVirtioSocketDeviceConfiguration struct {
	VZSocketDeviceConfiguration
}

// VZVirtioSocketDeviceConfigurationFrom constructs a [VZVirtioSocketDeviceConfiguration] from an unsafe.Pointer.
//
// A configuration object that requests the creation of a socket device to communicate with the guest system.
func VZVirtioSocketDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioSocketDeviceConfiguration {
	return VZVirtioSocketDeviceConfiguration{
		VZSocketDeviceConfiguration: VZSocketDeviceConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioSocketDeviceConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioSocketDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioSocketDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioSocketDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioSocketDeviceConfiguration */

// The socket device that you use to implement port-based communication with the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/socketdevices
func (v_ VZVirtioSocketDeviceConfiguration) SocketDevices() IVZSocketDeviceConfiguration {
	rv := objc.Send[VZSocketDeviceConfiguration](v_.ID, objc.Sel("socketDevices"))
	return rv
}/* debug [instance_properties/getter]: socketDevices */


// The socket device that you use to implement port-based communication with the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/socketdevices
func (v_ VZVirtioSocketDeviceConfiguration) SetSocketDevices(value IVZSocketDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSocketDevices:"), value)
}/* debug [instance_properties/setter]: socketDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioSocketDeviceConfiguration */


