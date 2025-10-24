// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVirtioSocketDevice */


/* debug [class_header]: Header for VZVirtioSocketDevice */
// The class instance for the [VZVirtioSocketDevice] class.
var (
	VZVirtioSocketDeviceClass     _VZVirtioSocketDeviceClass
	VZVirtioSocketDeviceClassOnce sync.Once
)

func getVZVirtioSocketDeviceClass() _VZVirtioSocketDeviceClass {
	VZVirtioSocketDeviceClassOnce.Do(func() {
		VZVirtioSocketDeviceClass = _VZVirtioSocketDeviceClass{objc.GetClass("VZVirtioSocketDevice")}
	})
	return VZVirtioSocketDeviceClass
}

type _VZVirtioSocketDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioSocketDevice */
// An interface definition for the [VZVirtioSocketDevice] class.
type IVZVirtioSocketDevice interface {
	IVZSocketDevice
	
/* debug [class_interface_properties]: Properties for VZVirtioSocketDevice */
	// properties:
	SocketDevices() IVZSocketDevice
	SetSocketDevices(value IVZSocketDevice)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioSocketDevice */
	// methods:
	ConnectToPortCompletionHandler(port uint32 /* not a class type */, completionHandler unsafe.Pointer)
	RemoveSocketListenerForPort(port uint32 /* not a class type */)
	SetSocketListenerForPort(listener IVZVirtioSocketListener, port uint32 /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioSocketDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSocketDeviceClass) Alloc() VZVirtioSocketDevice {
	rv := objc.Send[VZVirtioSocketDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioSocketDeviceClass) New() VZVirtioSocketDevice {
	rv := objc.Send[VZVirtioSocketDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioSocketDevice) Init() VZVirtioSocketDevice {
	rv := objc.Send[VZVirtioSocketDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioSocketDevice) Autorelease() VZVirtioSocketDevice {
	rv := objc.Send[VZVirtioSocketDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketDevice creates a new VZVirtioSocketDevice instance.
func NewVZVirtioSocketDevice() VZVirtioSocketDevice {
	return getVZVirtioSocketDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioSocketDevice */
// A device that manages port-based connections between the guest system and the host computer.
//
// Use a object to configure services and other communication end points in your virtual machine. Host computers make services available using ports, which identify the type of service and the protocol to use when transmitting data. Use this object to specify the ports available to your guest operating system, and to register handlers to manage the communication on those ports. Don’t create a object directly. Instead, when you request a socket device in your configuration, the virtual machine creates it and stores it in the property. For each port you want to make available in your virtual machine, call the method and provide an object to manage the port connections.


// A device that manages port-based connections between the guest system and the host computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDevice
type VZVirtioSocketDevice struct {
	VZSocketDevice
}

// VZVirtioSocketDeviceFrom constructs a [VZVirtioSocketDevice] from an unsafe.Pointer.
//
// A device that manages port-based connections between the guest system and the host computer.
func VZVirtioSocketDeviceFrom(ptr unsafe.Pointer) VZVirtioSocketDevice {
	return VZVirtioSocketDevice{
		VZSocketDevice: VZSocketDeviceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioSocketDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioSocketDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioSocketDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioSocketDevice */

// Initiates a connection to the specified port of the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDevice/connect(toPort:)
func (v_ VZVirtioSocketDevice) ConnectToPortCompletionHandler(port uint32 /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("connectToPort:completionHandler:"), port, completionHandler)
}/* debug [instance_methods/method]: ConnectToPortCompletionHandler */


// Removes the listener object from the specfied port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDevice/removeSocketListener(forPort:)
func (v_ VZVirtioSocketDevice) RemoveSocketListenerForPort(port uint32 /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeSocketListenerForPort:"), port)
}/* debug [instance_methods/method]: RemoveSocketListenerForPort */


// Configures an object to monitor the specified port for new connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDevice/setSocketListener(_:forPort:)
func (v_ VZVirtioSocketDevice) SetSocketListenerForPort(listener IVZVirtioSocketListener, port uint32 /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSocketListener:forPort:"), listener, port)
}/* debug [instance_methods/method]: SetSocketListenerForPort */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioSocketDevice */

// The array of socket devices that the VM configures for use ports in the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/socketdevices
func (v_ VZVirtioSocketDevice) SocketDevices() IVZSocketDevice {
	rv := objc.Send[VZSocketDevice](v_.ID, objc.Sel("socketDevices"))
	return rv
}/* debug [instance_properties/getter]: socketDevices */


// The array of socket devices that the VM configures for use ports in the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/socketdevices
func (v_ VZVirtioSocketDevice) SetSocketDevices(value IVZSocketDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSocketDevices:"), value)
}/* debug [instance_properties/setter]: socketDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioSocketDevice */



