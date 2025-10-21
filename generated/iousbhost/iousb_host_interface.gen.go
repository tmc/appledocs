// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	SelectAlternateSettingError(alternateSetting uint, error_ unsafe.Pointer) bool
}

// The class for accessing USB-related services.
//
// Use this class to create pipes, retrieve descriptors, send device requests, and enable power savings. Create an instance of the class with .
//
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
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/createMatchingDictionaryWithVendorID:productID:bcdDevice:interfaceNumber:configurationValue:interfaceClass:interfaceSubclass:interfaceProtocol:speed:productIDArray:
func (uc _USBHostInterfaceClass) CreateMatchingDictionaryWithVendorIDProductIDBcdDeviceInterfaceNumberConfigurationValueInterfaceClassInterfaceSubclassInterfaceProtocolSpeedProductIDArray(vendorID unsafe.Pointer, productID unsafe.Pointer, bcdDevice unsafe.Pointer, interfaceNumber unsafe.Pointer, configurationValue unsafe.Pointer, interfaceClass unsafe.Pointer, interfaceSubclass unsafe.Pointer, interfaceProtocol unsafe.Pointer, speed unsafe.Pointer, productIDArray objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("createMatchingDictionaryWithVendorID:productID:bcdDevice:interfaceNumber:configurationValue:interfaceClass:interfaceSubclass:interfaceProtocol:speed:productIDArray:"), vendorID, productID, bcdDevice, interfaceNumber, configurationValue, interfaceClass, interfaceSubclass, interfaceProtocol, speed, productIDArray)
	return rv
}

// Selects an alternative setting for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/selectAlternateSetting(_:)
func (u_ USBHostInterface) SelectAlternateSettingError(alternateSetting uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("selectAlternateSetting:error:"), alternateSetting, error_)
	return rv
}

// The descriptor for the interface.
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/interfaceDescriptor
func (u_ USBHostInterface) InterfaceDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("interfaceDescriptor"))
	return rv
}



