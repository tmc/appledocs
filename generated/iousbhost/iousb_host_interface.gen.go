// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class IOUSBHostInterface */


/* debug [class_header]: Header for IOUSBHostInterface */
// The class instance for the [USBHostInterface] class.
var (
	USBHostInterfaceClass     _USBHostInterfaceClass
	USBHostInterfaceClassOnce sync.Once
)

func getUSBHostInterfaceClass() _USBHostInterfaceClass {
	USBHostInterfaceClassOnce.Do(func() {
		USBHostInterfaceClass = _USBHostInterfaceClass{objc.GetClass("IOUSBHostInterface")}
	})
	return USBHostInterfaceClass
}

type _USBHostInterfaceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for USBHostInterface */
// An interface definition for the [USBHostInterface] class.
type IUSBHostInterface interface {
	IUSBHostObject
	
/* debug [class_interface_properties]: Properties for USBHostInterface */
	// properties:
	ConfigurationDescriptor() USBConfigurationDescriptor /* not a class type */
	IdleTimeout() float64
	InterfaceDescriptor() USBInterfaceDescriptor /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for USBHostInterface */
	// methods:
	CopyPipeWithAddressError(address uint, error_ unsafe.Pointer) IUSBHostPipe
	SelectAlternateSettingError(alternateSetting uint, error_ unsafe.Pointer) bool
	SetIdleTimeoutError(idleTimeout float64, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for USBHostInterface */
// Alloc allocates a new instance without initialization.
func (uc _USBHostInterfaceClass) Alloc() USBHostInterface {
	rv := objc.Send[USBHostInterface](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _USBHostInterfaceClass) New() USBHostInterface {
	rv := objc.Send[USBHostInterface](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostInterface) Init() USBHostInterface {
	rv := objc.Send[USBHostInterface](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostInterface) Autorelease() USBHostInterface {
	rv := objc.Send[USBHostInterface](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostInterface creates a new USBHostInterface instance.
func NewUSBHostInterface() USBHostInterface {
	return getUSBHostInterfaceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for USBHostInterface */
// The class for accessing USB-related services.
//
// Use this class to create pipes, retrieve descriptors, send device requests, and enable power savings. Create an instance of the class with .


// The class for accessing USB-related services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface
type USBHostInterface struct {
	USBHostObject
}

// USBHostInterfaceFrom constructs a [USBHostInterface] from an unsafe.Pointer.
//
// The class for accessing USB-related services.
func USBHostInterfaceFrom(ptr unsafe.Pointer) USBHostInterface {
	return USBHostInterface{
		USBHostObject: USBHostObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for USBHostInterface */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/initWithIOService:options:queue:error:interestHandler:
func NewUSBHostInterfaceWithIOServiceOptionsQueueErrorInterestHandler(ioService unsafe.Pointer, options USBHostObjectInitOptions, queue unsafe.Pointer, error_ unsafe.Pointer, interestHandler USBHostInterestHandler /* not a class type */) USBHostInterface {
	instance := getUSBHostInterfaceClass().Alloc()
	rv := objc.Send[USBHostInterface](instance.ID, objc.Sel("initWithIOService:options:queue:error:interestHandler:"), ioService, options, queue, error_, interestHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUSBHostInterfaceWithIOServiceOptionsQueueErrorInterestHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for USBHostInterface */

// Creates a matching dictionary to find a USB interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/createMatchingDictionaryWithVendorID:productID:bcdDevice:interfaceNumber:configurationValue:interfaceClass:interfaceSubclass:interfaceProtocol:speed:productIDArray:
func (uc _USBHostInterfaceClass) CreateMatchingDictionaryWithVendorIDProductIDBcdDeviceInterfaceNumberConfigurationValueInterfaceClassInterfaceSubclassInterfaceProtocolSpeedProductIDArray(vendorID objc.IObject /* cross-framework: NSNumber */, productID objc.IObject /* cross-framework: NSNumber */, bcdDevice objc.IObject /* cross-framework: NSNumber */, interfaceNumber objc.IObject /* cross-framework: NSNumber */, configurationValue objc.IObject /* cross-framework: NSNumber */, interfaceClass objc.IObject /* cross-framework: NSNumber */, interfaceSubclass objc.IObject /* cross-framework: NSNumber */, interfaceProtocol objc.IObject /* cross-framework: NSNumber */, speed objc.IObject /* cross-framework: NSNumber */, productIDArray objc.IObject /* cross-framework: NSArray */) MutableDictionaryRef /* not a class type */ {
	rv := objc.Send[MutableDictionaryRef](objc.ID(uc.class), objc.Sel("createMatchingDictionaryWithVendorID:productID:bcdDevice:interfaceNumber:configurationValue:interfaceClass:interfaceSubclass:interfaceProtocol:speed:productIDArray:"), vendorID, productID, bcdDevice, interfaceNumber, configurationValue, interfaceClass, interfaceSubclass, interfaceProtocol, speed, productIDArray)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateMatchingDictionaryWithVendorIDProductIDBcdDeviceInterfaceNumberConfigurationValueInterfaceClassInterfaceSubclassInterfaceProtocolSpeedProductIDArray) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for USBHostInterface */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for USBHostInterface */

// Copies a pipe for a specific endpoint address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/copyPipe(withAddress:)
func (u_ USBHostInterface) CopyPipeWithAddressError(address uint, error_ unsafe.Pointer) IUSBHostPipe {
	rv := objc.Send[USBHostPipe](u_.ID, objc.Sel("copyPipeWithAddress:error:"), address, error_)
	return rv
}/* debug [instance_methods/method]: CopyPipeWithAddressError */


// Selects an alternative setting for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/selectAlternateSetting(_:)
func (u_ USBHostInterface) SelectAlternateSettingError(alternateSetting uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("selectAlternateSetting:error:"), alternateSetting, error_)
	return rv
}/* debug [instance_methods/method]: SelectAlternateSettingError */


// Sets the desired idle suspend timeout for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/setIdleTimeout(_:)
func (u_ USBHostInterface) SetIdleTimeoutError(idleTimeout float64, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setIdleTimeout:error:"), idleTimeout, error_)
	return rv
}/* debug [instance_methods/method]: SetIdleTimeoutError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for USBHostInterface */

// The configuration descriptor for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/configurationDescriptor
func (u_ USBHostInterface) ConfigurationDescriptor() USBConfigurationDescriptor /* not a class type */ {
	rv := objc.Send[USBConfigurationDescriptor](u_.ID, objc.Sel("configurationDescriptor"))
	return rv
}/* debug [instance_properties/getter]: configurationDescriptor */


// The current idle suspend timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/idleTimeout
func (u_ USBHostInterface) IdleTimeout() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("idleTimeout"))
	return rv
}/* debug [instance_properties/getter]: idleTimeout */


// The descriptor for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/interfaceDescriptor
func (u_ USBHostInterface) InterfaceDescriptor() USBInterfaceDescriptor /* not a class type */ {
	rv := objc.Send[USBInterfaceDescriptor](u_.ID, objc.Sel("interfaceDescriptor"))
	return rv
}/* debug [instance_properties/getter]: interfaceDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOUSBHostInterface */


