// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class IOUSBHostDevice */


/* debug [class_header]: Header for IOUSBHostDevice */
// The class instance for the [USBHostDevice] class.
var (
	USBHostDeviceClass     _USBHostDeviceClass
	USBHostDeviceClassOnce sync.Once
)

func getUSBHostDeviceClass() _USBHostDeviceClass {
	USBHostDeviceClassOnce.Do(func() {
		USBHostDeviceClass = _USBHostDeviceClass{objc.GetClass("IOUSBHostDevice")}
	})
	return USBHostDeviceClass
}

type _USBHostDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for USBHostDevice */
// An interface definition for the [USBHostDevice] class.
type IUSBHostDevice interface {
	IUSBHostObject
	
/* debug [class_interface_properties]: Properties for USBHostDevice */
	// properties:
	ConfigurationDescriptor() USBConfigurationDescriptor /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for USBHostDevice */
	// methods:
	ConfigureWithValueError(value uint, error_ unsafe.Pointer) bool
	ConfigureWithValueMatchInterfacesError(value uint, matchInterfaces bool, error_ unsafe.Pointer) bool
	ResetWithError(error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for USBHostDevice */
// Alloc allocates a new instance without initialization.
func (uc _USBHostDeviceClass) Alloc() USBHostDevice {
	rv := objc.Send[USBHostDevice](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _USBHostDeviceClass) New() USBHostDevice {
	rv := objc.Send[USBHostDevice](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostDevice) Init() USBHostDevice {
	rv := objc.Send[USBHostDevice](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostDevice) Autorelease() USBHostDevice {
	rv := objc.Send[USBHostDevice](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostDevice creates a new USBHostDevice instance.
func NewUSBHostDevice() USBHostDevice {
	return getUSBHostDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for USBHostDevice */
// The class that claims and configures devices, retrieves descriptors, and sends device requests.
//
// This class enables management of the device state, including sending control requests to the default endpoint 0, configuring the device, and resetting the device. The interest handler also allows monitoring of the device state. The client creates the class and initializes it with .


// The class that claims and configures devices, retrieves descriptors, and sends device requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice
type USBHostDevice struct {
	USBHostObject
}

// USBHostDeviceFrom constructs a [USBHostDevice] from an unsafe.Pointer.
//
// The class that claims and configures devices, retrieves descriptors, and sends device requests.
func USBHostDeviceFrom(ptr unsafe.Pointer) USBHostDevice {
	return USBHostDevice{
		USBHostObject: USBHostObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for USBHostDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for USBHostDevice */

// Creates a matching dictionary to find a USB device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice/createMatchingDictionaryWithVendorID:productID:bcdDevice:deviceClass:deviceSubclass:deviceProtocol:speed:productIDArray:
func (uc _USBHostDeviceClass) CreateMatchingDictionaryWithVendorIDProductIDBcdDeviceDeviceClassDeviceSubclassDeviceProtocolSpeedProductIDArray(vendorID objc.IObject /* cross-framework: NSNumber */, productID objc.IObject /* cross-framework: NSNumber */, bcdDevice objc.IObject /* cross-framework: NSNumber */, deviceClass objc.IObject /* cross-framework: NSNumber */, deviceSubclass objc.IObject /* cross-framework: NSNumber */, deviceProtocol objc.IObject /* cross-framework: NSNumber */, speed objc.IObject /* cross-framework: NSNumber */, productIDArray objc.IObject /* cross-framework: NSArray */) MutableDictionaryRef /* not a class type */ {
	rv := objc.Send[MutableDictionaryRef](objc.ID(uc.class), objc.Sel("createMatchingDictionaryWithVendorID:productID:bcdDevice:deviceClass:deviceSubclass:deviceProtocol:speed:productIDArray:"), vendorID, productID, bcdDevice, deviceClass, deviceSubclass, deviceProtocol, speed, productIDArray)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateMatchingDictionaryWithVendorIDProductIDBcdDeviceDeviceClassDeviceSubclassDeviceProtocolSpeedProductIDArray) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for USBHostDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for USBHostDevice */

// Selects a new configuration for the device and registers the interfaces for matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice/configureWithValue:error:
func (u_ USBHostDevice) ConfigureWithValueError(value uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("configureWithValue:error:"), value, error_)
	return rv
}/* debug [instance_methods/method]: ConfigureWithValueError */


// Selects a new configuration for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice/configureWithValue:matchInterfaces:error:
func (u_ USBHostDevice) ConfigureWithValueMatchInterfacesError(value uint, matchInterfaces bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("configureWithValue:matchInterfaces:error:"), value, matchInterfaces, error_)
	return rv
}/* debug [instance_methods/method]: ConfigureWithValueMatchInterfacesError */


// Terminates the device and attempts to re-enumerate it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice/reset()
func (u_ USBHostDevice) ResetWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("resetWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: ResetWithError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for USBHostDevice */

// The currently selected configuration descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice/configurationDescriptor
func (u_ USBHostDevice) ConfigurationDescriptor() USBConfigurationDescriptor /* not a class type */ {
	rv := objc.Send[USBConfigurationDescriptor](u_.ID, objc.Sel("configurationDescriptor"))
	return rv
}/* debug [instance_properties/getter]: configurationDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOUSBHostDevice */



