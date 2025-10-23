// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [USBHostInterface] class.
type IUSBHostInterface interface {
	IUSBHostObject
	// properties:
	ConfigurationDescriptor() unsafe.Pointer
	SetConfigurationDescriptor(value unsafe.Pointer)
	IdleTimeout() unsafe.Pointer
	SetIdleTimeout(value unsafe.Pointer)
	InterfaceDescriptor() unsafe.Pointer
	SetInterfaceDescriptor(value unsafe.Pointer)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (uc _USBHostInterfaceClass) Alloc() USBHostInterface {
	rv := objc.Send[USBHostInterface](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a matching dictionary to find a USB interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/createMatchingDictionaryWithVendorID:productID:bcdDevice:interfaceNumber:configurationValue:interfaceClass:interfaceSubclass:interfaceProtocol:speed:productIDArray:
func (uc _USBHostInterfaceClass) CreateMatchingDictionaryWithVendorIDProductIDBcdDeviceInterfaceNumberConfigurationValueInterfaceClassInterfaceSubclassInterfaceProtocolSpeedProductIDArray(vendorID foundation.Number, productID foundation.Number, bcdDevice foundation.Number, interfaceNumber foundation.Number, configurationValue foundation.Number, interfaceClass foundation.Number, interfaceSubclass foundation.Number, interfaceProtocol foundation.Number, speed foundation.Number, productIDArray objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("createMatchingDictionaryWithVendorID:productID:bcdDevice:interfaceNumber:configurationValue:interfaceClass:interfaceSubclass:interfaceProtocol:speed:productIDArray:"), vendorID, productID, bcdDevice, interfaceNumber, configurationValue, interfaceClass, interfaceSubclass, interfaceProtocol, speed, productIDArray)
	return rv
}


// The configuration descriptor for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostinterface/configurationdescriptor
func (u_ USBHostInterface) ConfigurationDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("configurationDescriptor"))
	return rv
}


// The configuration descriptor for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostinterface/configurationdescriptor
func (u_ USBHostInterface) SetConfigurationDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConfigurationDescriptor:"), value)
}


// The current idle suspend timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostinterface/idletimeout
func (u_ USBHostInterface) IdleTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("idleTimeout"))
	return rv
}


// The current idle suspend timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostinterface/idletimeout
func (u_ USBHostInterface) SetIdleTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIdleTimeout:"), value)
}


// The descriptor for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostinterface/interfacedescriptor
func (u_ USBHostInterface) InterfaceDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("interfaceDescriptor"))
	return rv
}


// The descriptor for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostinterface/interfacedescriptor
func (u_ USBHostInterface) SetInterfaceDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInterfaceDescriptor:"), value)
}



