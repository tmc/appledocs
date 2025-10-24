// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

/* debug [functions.gen.go]: Generating 43 functions for IOUSBHost */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// IOUSBHost Functions (43 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_IOUSBGetBillboardDescriptor func(unsafe.Pointer) unsafe.Pointer
	_IOUSBGetConfigurationMaxPowerMilliAmps func(uint32, unsafe.Pointer) uint32
	_IOUSBGetContainerIDDescriptor func(unsafe.Pointer) unsafe.Pointer
	_IOUSBGetEndpointAddress func(unsafe.Pointer) uint8
	_IOUSBGetEndpointBurstSize func(uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) uint32
	_IOUSBGetEndpointDirection func(unsafe.Pointer) uint8
	_IOUSBGetEndpointIntervalEncodedMicroframes func(uint32, unsafe.Pointer) uint32
	_IOUSBGetEndpointIntervalFrames func(uint32, unsafe.Pointer) uint32
	_IOUSBGetEndpointIntervalMicroframes func(uint32, unsafe.Pointer) uint32
	_IOUSBGetEndpointMaxPacketSize func(uint32, unsafe.Pointer) uint16
	_IOUSBGetEndpointMaxStreams func(uint32, unsafe.Pointer, unsafe.Pointer) uint32
	_IOUSBGetEndpointMaxStreamsEncoded func(uint32, unsafe.Pointer, unsafe.Pointer) uint32
	_IOUSBGetEndpointMult func(uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) uint8
	_IOUSBGetEndpointNumber func(unsafe.Pointer) uint8
	_IOUSBGetEndpointSynchronizationType func(unsafe.Pointer) uint8
	_IOUSBGetEndpointType func(unsafe.Pointer) uint8
	_IOUSBGetEndpointUsageType func(unsafe.Pointer) uint8
	_IOUSBGetNextAssociatedDescriptor func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBGetNextAssociatedDescriptorWithType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBGetNextCapabilityDescriptor func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBGetNextCapabilityDescriptorWithType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBGetNextDescriptor func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBGetNextDescriptorWithType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBGetNextEndpointDescriptor func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBGetNextInterfaceAssociationDescriptor func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBGetNextInterfaceDescriptor func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBGetPlatformCapabilityDescriptor func(unsafe.Pointer) unsafe.Pointer
	_IOUSBGetPlatformCapabilityDescriptorWithUUID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBGetSuperSpeedDeviceCapabilityDescriptor func(unsafe.Pointer) unsafe.Pointer
	_IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor func(unsafe.Pointer) unsafe.Pointer
	_IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor func(unsafe.Pointer) unsafe.Pointer
	_IOUSBHostCIControllerStateToString func(USBHostCIControllerState) unsafe.Pointer
	_IOUSBHostCIDeviceSpeedToString func(USBHostCIDeviceSpeed) unsafe.Pointer
	_IOUSBHostCIDeviceStateToString func(USBHostCIDeviceState) unsafe.Pointer
	_IOUSBHostCIEndpointStateToString func(USBHostCIEndpointState) unsafe.Pointer
	_IOUSBHostCIExceptionTypeToString func(USBHostCIExceptionType) unsafe.Pointer
	_IOUSBHostCILinkStateEnabled func(unsafe.Pointer) bool
	_IOUSBHostCILinkStateToString func(USBHostCILinkState) unsafe.Pointer
	_IOUSBHostCIMessageStatusFromIOReturn func(int) USBHostCIMessageStatus
	_IOUSBHostCIMessageStatusToIOReturn func(USBHostCIMessageStatus) int
	_IOUSBHostCIMessageStatusToString func(USBHostCIMessageStatus) unsafe.Pointer
	_IOUSBHostCIMessageTypeToString func(USBHostCIMessageType) unsafe.Pointer
	_IOUSBHostCIPortStateToString func(USBHostCIPortState) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_IOUSBGetBillboardDescriptor, lib, "IOUSBGetBillboardDescriptor")
	tryRegister(&_IOUSBGetConfigurationMaxPowerMilliAmps, lib, "IOUSBGetConfigurationMaxPowerMilliAmps")
	tryRegister(&_IOUSBGetContainerIDDescriptor, lib, "IOUSBGetContainerIDDescriptor")
	tryRegister(&_IOUSBGetEndpointAddress, lib, "IOUSBGetEndpointAddress")
	tryRegister(&_IOUSBGetEndpointBurstSize, lib, "IOUSBGetEndpointBurstSize")
	tryRegister(&_IOUSBGetEndpointDirection, lib, "IOUSBGetEndpointDirection")
	tryRegister(&_IOUSBGetEndpointIntervalEncodedMicroframes, lib, "IOUSBGetEndpointIntervalEncodedMicroframes")
	tryRegister(&_IOUSBGetEndpointIntervalFrames, lib, "IOUSBGetEndpointIntervalFrames")
	tryRegister(&_IOUSBGetEndpointIntervalMicroframes, lib, "IOUSBGetEndpointIntervalMicroframes")
	tryRegister(&_IOUSBGetEndpointMaxPacketSize, lib, "IOUSBGetEndpointMaxPacketSize")
	tryRegister(&_IOUSBGetEndpointMaxStreams, lib, "IOUSBGetEndpointMaxStreams")
	tryRegister(&_IOUSBGetEndpointMaxStreamsEncoded, lib, "IOUSBGetEndpointMaxStreamsEncoded")
	tryRegister(&_IOUSBGetEndpointMult, lib, "IOUSBGetEndpointMult")
	tryRegister(&_IOUSBGetEndpointNumber, lib, "IOUSBGetEndpointNumber")
	tryRegister(&_IOUSBGetEndpointSynchronizationType, lib, "IOUSBGetEndpointSynchronizationType")
	tryRegister(&_IOUSBGetEndpointType, lib, "IOUSBGetEndpointType")
	tryRegister(&_IOUSBGetEndpointUsageType, lib, "IOUSBGetEndpointUsageType")
	tryRegister(&_IOUSBGetNextAssociatedDescriptor, lib, "IOUSBGetNextAssociatedDescriptor")
	tryRegister(&_IOUSBGetNextAssociatedDescriptorWithType, lib, "IOUSBGetNextAssociatedDescriptorWithType")
	tryRegister(&_IOUSBGetNextCapabilityDescriptor, lib, "IOUSBGetNextCapabilityDescriptor")
	tryRegister(&_IOUSBGetNextCapabilityDescriptorWithType, lib, "IOUSBGetNextCapabilityDescriptorWithType")
	tryRegister(&_IOUSBGetNextDescriptor, lib, "IOUSBGetNextDescriptor")
	tryRegister(&_IOUSBGetNextDescriptorWithType, lib, "IOUSBGetNextDescriptorWithType")
	tryRegister(&_IOUSBGetNextEndpointDescriptor, lib, "IOUSBGetNextEndpointDescriptor")
	tryRegister(&_IOUSBGetNextInterfaceAssociationDescriptor, lib, "IOUSBGetNextInterfaceAssociationDescriptor")
	tryRegister(&_IOUSBGetNextInterfaceDescriptor, lib, "IOUSBGetNextInterfaceDescriptor")
	tryRegister(&_IOUSBGetPlatformCapabilityDescriptor, lib, "IOUSBGetPlatformCapabilityDescriptor")
	tryRegister(&_IOUSBGetPlatformCapabilityDescriptorWithUUID, lib, "IOUSBGetPlatformCapabilityDescriptorWithUUID")
	tryRegister(&_IOUSBGetSuperSpeedDeviceCapabilityDescriptor, lib, "IOUSBGetSuperSpeedDeviceCapabilityDescriptor")
	tryRegister(&_IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor, lib, "IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor")
	tryRegister(&_IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor, lib, "IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor")
	tryRegister(&_IOUSBHostCIControllerStateToString, lib, "IOUSBHostCIControllerStateToString")
	tryRegister(&_IOUSBHostCIDeviceSpeedToString, lib, "IOUSBHostCIDeviceSpeedToString")
	tryRegister(&_IOUSBHostCIDeviceStateToString, lib, "IOUSBHostCIDeviceStateToString")
	tryRegister(&_IOUSBHostCIEndpointStateToString, lib, "IOUSBHostCIEndpointStateToString")
	tryRegister(&_IOUSBHostCIExceptionTypeToString, lib, "IOUSBHostCIExceptionTypeToString")
	tryRegister(&_IOUSBHostCILinkStateEnabled, lib, "IOUSBHostCILinkStateEnabled")
	tryRegister(&_IOUSBHostCILinkStateToString, lib, "IOUSBHostCILinkStateToString")
	tryRegister(&_IOUSBHostCIMessageStatusFromIOReturn, lib, "IOUSBHostCIMessageStatusFromIOReturn")
	tryRegister(&_IOUSBHostCIMessageStatusToIOReturn, lib, "IOUSBHostCIMessageStatusToIOReturn")
	tryRegister(&_IOUSBHostCIMessageStatusToString, lib, "IOUSBHostCIMessageStatusToString")
	tryRegister(&_IOUSBHostCIMessageTypeToString, lib, "IOUSBHostCIMessageTypeToString")
	tryRegister(&_IOUSBHostCIPortStateToString, lib, "IOUSBHostCIPortStateToString")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Obtains the first billboard capability descriptor in a BOS descriptor.
//
// Added in macOS 10.15.
// Obtains the first billboard capability descriptor in a BOS descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetBillboardDescriptor(_:)
func IOUSBGetBillboardDescriptor(bosDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetBillboardDescriptor(bosDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetBillboardDescriptor */

// Obtains the maximum bus current that a configuration descriptor requires.
//
// Added in macOS 10.15.
// Obtains the maximum bus current that a configuration descriptor requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetConfigurationMaxPowerMilliAmps(_:_:)
func IOUSBGetConfigurationMaxPowerMilliAmps(usbDeviceSpeed uint32, descriptor unsafe.Pointer) uint32 {
	return _IOUSBGetConfigurationMaxPowerMilliAmps(usbDeviceSpeed, descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetConfigurationMaxPowerMilliAmps */

// Obtains the first container ID capability descriptor in a BOS descriptor.
//
// Added in macOS 10.15.
// Obtains the first container ID capability descriptor in a BOS descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetContainerIDDescriptor(_:)
func IOUSBGetContainerIDDescriptor(bosDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetContainerIDDescriptor(bosDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetContainerIDDescriptor */

// Obtains the direction and number of an endpoint from an endpoint descriptor.
//
// Added in macOS 10.15.
// Obtains the direction and number of an endpoint from an endpoint descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointAddress(_:)
func IOUSBGetEndpointAddress(descriptor unsafe.Pointer) uint8 {
	return _IOUSBGetEndpointAddress(descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointAddress */

// IOUSBGetEndpointBurstSize is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointBurstSize(_:_:_:_:)
func IOUSBGetEndpointBurstSize(usbDeviceSpeed uint32, descriptor unsafe.Pointer, companionDescriptor unsafe.Pointer, sspCompanionDescriptor unsafe.Pointer) uint32 {
	return _IOUSBGetEndpointBurstSize(usbDeviceSpeed, descriptor, companionDescriptor, sspCompanionDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointBurstSize */

// Obtains the direction of an endpoint from an endpoint descriptor.
//
// Added in macOS 10.15.
// Obtains the direction of an endpoint from an endpoint descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointDirection(_:)
func IOUSBGetEndpointDirection(descriptor unsafe.Pointer) uint8 {
	return _IOUSBGetEndpointDirection(descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointDirection */

// Obtains the interval of an endpoint descriptor.
//
// Added in macOS 10.15.
// Obtains the interval of an endpoint descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointIntervalEncodedMicroframes(_:_:)
func IOUSBGetEndpointIntervalEncodedMicroframes(usbDeviceSpeed uint32, descriptor unsafe.Pointer) uint32 {
	return _IOUSBGetEndpointIntervalEncodedMicroframes(usbDeviceSpeed, descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointIntervalEncodedMicroframes */

// Obtains the interval of an endpoint descriptor.
//
// Added in macOS 10.15.
// Obtains the interval of an endpoint descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointIntervalFrames(_:_:)
func IOUSBGetEndpointIntervalFrames(usbDeviceSpeed uint32, descriptor unsafe.Pointer) uint32 {
	return _IOUSBGetEndpointIntervalFrames(usbDeviceSpeed, descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointIntervalFrames */

// Obtains the interval of an endpoint descriptor.
//
// Added in macOS 10.15.
// Obtains the interval of an endpoint descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointIntervalMicroframes(_:_:)
func IOUSBGetEndpointIntervalMicroframes(usbDeviceSpeed uint32, descriptor unsafe.Pointer) uint32 {
	return _IOUSBGetEndpointIntervalMicroframes(usbDeviceSpeed, descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointIntervalMicroframes */

// Obtains the maximum packet size from an endpoint descriptor.
//
// Added in macOS 10.15.
// Obtains the maximum packet size from an endpoint descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointMaxPacketSize(_:_:)
func IOUSBGetEndpointMaxPacketSize(usbDeviceSpeed uint32, descriptor unsafe.Pointer) uint16 {
	return _IOUSBGetEndpointMaxPacketSize(usbDeviceSpeed, descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointMaxPacketSize */

// Obtains the number of supported streams.
//
// Added in macOS 10.15.
// Obtains the number of supported streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointMaxStreams(_:_:_:)
func IOUSBGetEndpointMaxStreams(usbDeviceSpeed uint32, descriptor unsafe.Pointer, companionDescriptor unsafe.Pointer) uint32 {
	return _IOUSBGetEndpointMaxStreams(usbDeviceSpeed, descriptor, companionDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointMaxStreams */

// Obtains the number of streams that an endpoint supports.
//
// Added in macOS 10.15.
// Obtains the number of streams that an endpoint supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointMaxStreamsEncoded(_:_:_:)
func IOUSBGetEndpointMaxStreamsEncoded(usbDeviceSpeed uint32, descriptor unsafe.Pointer, companionDescriptor unsafe.Pointer) uint32 {
	return _IOUSBGetEndpointMaxStreamsEncoded(usbDeviceSpeed, descriptor, companionDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointMaxStreamsEncoded */

// IOUSBGetEndpointMult is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointMult(_:_:_:_:)
func IOUSBGetEndpointMult(usbDeviceSpeed uint32, descriptor unsafe.Pointer, companionDescriptor unsafe.Pointer, sspCompanionDescriptor unsafe.Pointer) uint8 {
	return _IOUSBGetEndpointMult(usbDeviceSpeed, descriptor, companionDescriptor, sspCompanionDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointMult */

// Obtains the number of an endpoint from an endpoint descriptor.
//
// Added in macOS 10.15.
// Obtains the number of an endpoint from an endpoint descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointNumber(_:)
func IOUSBGetEndpointNumber(descriptor unsafe.Pointer) uint8 {
	return _IOUSBGetEndpointNumber(descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointNumber */

// IOUSBGetEndpointSynchronizationType is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointSynchronizationType(_:)
func IOUSBGetEndpointSynchronizationType(descriptor unsafe.Pointer) uint8 {
	return _IOUSBGetEndpointSynchronizationType(descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointSynchronizationType */

// Obtains the type of an endpoint from an endpoint descriptor.
//
// Added in macOS 10.15.
// Obtains the type of an endpoint from an endpoint descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointType(_:)
func IOUSBGetEndpointType(descriptor unsafe.Pointer) uint8 {
	return _IOUSBGetEndpointType(descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointType */

// IOUSBGetEndpointUsageType is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointUsageType(_:)
func IOUSBGetEndpointUsageType(descriptor unsafe.Pointer) uint8 {
	return _IOUSBGetEndpointUsageType(descriptor)
}/* debug [functions.gen.go/function]: IOUSBGetEndpointUsageType */

// Obtains the next associated descriptor in a configuration descriptor.
//
// Added in macOS 10.15.
// Obtains the next associated descriptor in a configuration descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextAssociatedDescriptor(_:_:_:)
func IOUSBGetNextAssociatedDescriptor(configurationDescriptor unsafe.Pointer, parentDescriptor unsafe.Pointer, currentDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetNextAssociatedDescriptor(configurationDescriptor, parentDescriptor, currentDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetNextAssociatedDescriptor */

// Obtains the next associated descriptor in a configuration descriptor and matches the type.
//
// Added in macOS 10.15.
// Obtains the next associated descriptor in a configuration descriptor and matches the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextAssociatedDescriptorWithType(_:_:_:_:)
func IOUSBGetNextAssociatedDescriptorWithType(configurationDescriptor unsafe.Pointer, parentDescriptor unsafe.Pointer, currentDescriptor unsafe.Pointer, type_ unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetNextAssociatedDescriptorWithType(configurationDescriptor, parentDescriptor, currentDescriptor, type_)
}/* debug [functions.gen.go/function]: IOUSBGetNextAssociatedDescriptorWithType */

// Obtains the next device capability descriptor in a BOS descriptor.
//
// Added in macOS 10.15.
// Obtains the next device capability descriptor in a BOS descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextCapabilityDescriptor(_:_:)
func IOUSBGetNextCapabilityDescriptor(bosDescriptor unsafe.Pointer, currentDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetNextCapabilityDescriptor(bosDescriptor, currentDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetNextCapabilityDescriptor */

// Obtains the next descriptor matching a specific type within a BOS descriptor.
//
// Added in macOS 10.15.
// Obtains the next descriptor matching a specific type within a BOS descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextCapabilityDescriptorWithType(_:_:_:)
func IOUSBGetNextCapabilityDescriptorWithType(bosDescriptor unsafe.Pointer, currentDescriptor unsafe.Pointer, type_ unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetNextCapabilityDescriptorWithType(bosDescriptor, currentDescriptor, type_)
}/* debug [functions.gen.go/function]: IOUSBGetNextCapabilityDescriptorWithType */

// Obtains the next descriptor in a configuration descriptor.
//
// Added in macOS 10.15.
// Obtains the next descriptor in a configuration descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextDescriptor(_:_:)
func IOUSBGetNextDescriptor(configurationDescriptor unsafe.Pointer, currentDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetNextDescriptor(configurationDescriptor, currentDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetNextDescriptor */

// Obtains the next descriptor in a configuration descriptor that matches the type.
//
// Added in macOS 10.15.
// Obtains the next descriptor in a configuration descriptor that matches the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextDescriptorWithType(_:_:_:)
func IOUSBGetNextDescriptorWithType(configurationDescriptor unsafe.Pointer, currentDescriptor unsafe.Pointer, type_ unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetNextDescriptorWithType(configurationDescriptor, currentDescriptor, type_)
}/* debug [functions.gen.go/function]: IOUSBGetNextDescriptorWithType */

// Obtains the next endpoint descriptor for an interface descriptor.
//
// Added in macOS 10.15.
// Obtains the next endpoint descriptor for an interface descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextEndpointDescriptor(_:_:_:)
func IOUSBGetNextEndpointDescriptor(configurationDescriptor unsafe.Pointer, interfaceDescriptor unsafe.Pointer, currentDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetNextEndpointDescriptor(configurationDescriptor, interfaceDescriptor, currentDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetNextEndpointDescriptor */

// Obtains the next interface association descriptor in a configuration descriptor.
//
// Added in macOS 10.15.
// Obtains the next interface association descriptor in a configuration descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextInterfaceAssociationDescriptor(_:_:)
func IOUSBGetNextInterfaceAssociationDescriptor(configurationDescriptor unsafe.Pointer, currentDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetNextInterfaceAssociationDescriptor(configurationDescriptor, currentDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetNextInterfaceAssociationDescriptor */

// Obtains the next interface descriptor in a configuration descriptor.
//
// Added in macOS 10.15.
// Obtains the next interface descriptor in a configuration descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextInterfaceDescriptor(_:_:)
func IOUSBGetNextInterfaceDescriptor(configurationDescriptor unsafe.Pointer, currentDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetNextInterfaceDescriptor(configurationDescriptor, currentDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetNextInterfaceDescriptor */

// IOUSBGetPlatformCapabilityDescriptor is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetPlatformCapabilityDescriptor(_:)
func IOUSBGetPlatformCapabilityDescriptor(bosDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetPlatformCapabilityDescriptor(bosDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetPlatformCapabilityDescriptor */

// IOUSBGetPlatformCapabilityDescriptorWithUUID is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetPlatformCapabilityDescriptorWithUUID(_:_:)
func IOUSBGetPlatformCapabilityDescriptorWithUUID(bosDescriptor unsafe.Pointer, uuid unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetPlatformCapabilityDescriptorWithUUID(bosDescriptor, uuid)
}/* debug [functions.gen.go/function]: IOUSBGetPlatformCapabilityDescriptorWithUUID */

// Obtains the first SuperSpeed capability descriptor in a BOS descriptor.
//
// Added in macOS 10.15.
// Obtains the first SuperSpeed capability descriptor in a BOS descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetSuperSpeedDeviceCapabilityDescriptor(_:)
func IOUSBGetSuperSpeedDeviceCapabilityDescriptor(bosDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetSuperSpeedDeviceCapabilityDescriptor(bosDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetSuperSpeedDeviceCapabilityDescriptor */

// IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor(_:)
func IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor(bosDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor(bosDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor */

// Obtains the first USB 2.0 extension capability descriptor in a BOS descriptor.
//
// Added in macOS 10.15.
// Obtains the first USB 2.0 extension capability descriptor in a BOS descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor(_:)
func IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor(bosDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor(bosDescriptor)
}/* debug [functions.gen.go/function]: IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor */

// IOUSBHostCIControllerStateToString is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateToString(_:)
func IOUSBHostCIControllerStateToString(controllerState USBHostCIControllerState) unsafe.Pointer {
	return _IOUSBHostCIControllerStateToString(controllerState)
}/* debug [functions.gen.go/function]: IOUSBHostCIControllerStateToString */

// IOUSBHostCIDeviceSpeedToString is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceSpeedToString(_:)
func IOUSBHostCIDeviceSpeedToString(speed USBHostCIDeviceSpeed) unsafe.Pointer {
	return _IOUSBHostCIDeviceSpeedToString(speed)
}/* debug [functions.gen.go/function]: IOUSBHostCIDeviceSpeedToString */

// IOUSBHostCIDeviceStateToString is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateToString(_:)
func IOUSBHostCIDeviceStateToString(deviceState USBHostCIDeviceState) unsafe.Pointer {
	return _IOUSBHostCIDeviceStateToString(deviceState)
}/* debug [functions.gen.go/function]: IOUSBHostCIDeviceStateToString */

// IOUSBHostCIEndpointStateToString is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateToString(_:)
func IOUSBHostCIEndpointStateToString(endpointState USBHostCIEndpointState) unsafe.Pointer {
	return _IOUSBHostCIEndpointStateToString(endpointState)
}/* debug [functions.gen.go/function]: IOUSBHostCIEndpointStateToString */

// IOUSBHostCIExceptionTypeToString is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIExceptionTypeToString(_:)
func IOUSBHostCIExceptionTypeToString(exceptionType USBHostCIExceptionType) unsafe.Pointer {
	return _IOUSBHostCIExceptionTypeToString(exceptionType)
}/* debug [functions.gen.go/function]: IOUSBHostCIExceptionTypeToString */

// IOUSBHostCILinkStateEnabled is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCILinkStateEnabled(_:)
func IOUSBHostCILinkStateEnabled(linkState unsafe.Pointer) bool {
	return _IOUSBHostCILinkStateEnabled(linkState)
}/* debug [functions.gen.go/function]: IOUSBHostCILinkStateEnabled */

// IOUSBHostCILinkStateToString is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCILinkStateToString(_:)
func IOUSBHostCILinkStateToString(linkState USBHostCILinkState) unsafe.Pointer {
	return _IOUSBHostCILinkStateToString(linkState)
}/* debug [functions.gen.go/function]: IOUSBHostCILinkStateToString */

// IOUSBHostCIMessageStatusFromIOReturn is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageStatusFromIOReturn(_:)
func IOUSBHostCIMessageStatusFromIOReturn(status int) USBHostCIMessageStatus {
	return _IOUSBHostCIMessageStatusFromIOReturn(status)
}/* debug [functions.gen.go/function]: IOUSBHostCIMessageStatusFromIOReturn */

// IOUSBHostCIMessageStatusToIOReturn is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageStatusToIOReturn(_:)
func IOUSBHostCIMessageStatusToIOReturn(status USBHostCIMessageStatus) int {
	return _IOUSBHostCIMessageStatusToIOReturn(status)
}/* debug [functions.gen.go/function]: IOUSBHostCIMessageStatusToIOReturn */

// IOUSBHostCIMessageStatusToString is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageStatusToString(_:)
func IOUSBHostCIMessageStatusToString(status USBHostCIMessageStatus) unsafe.Pointer {
	return _IOUSBHostCIMessageStatusToString(status)
}/* debug [functions.gen.go/function]: IOUSBHostCIMessageStatusToString */

// IOUSBHostCIMessageTypeToString is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageTypeToString(_:)
func IOUSBHostCIMessageTypeToString(type_ USBHostCIMessageType) unsafe.Pointer {
	return _IOUSBHostCIMessageTypeToString(type_)
}/* debug [functions.gen.go/function]: IOUSBHostCIMessageTypeToString */

// IOUSBHostCIPortStateToString is a IOUSBHost function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateToString(_:)
func IOUSBHostCIPortStateToString(portState USBHostCIPortState) unsafe.Pointer {
	return _IOUSBHostCIPortStateToString(portState)
}/* debug [functions.gen.go/function]: IOUSBHostCIPortStateToString */




