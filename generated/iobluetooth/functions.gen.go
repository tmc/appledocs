// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// IOBluetooth Functions (64 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_IOBluetoothAddSCOAudioDevice func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothFindNumberOfRegistryEntriesOfClassName func(unsafe.Pointer) unsafe.Pointer
	_IOBluetoothGetUniqueFileNameAndPath func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothIgnoreHIDDevice func(unsafe.Pointer)
	_IOBluetoothIsFileAppleDesignatedPIMData func(unsafe.Pointer) unsafe.Pointer
	_IOBluetoothL2CAPChannelRegisterForChannelCloseNotification func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothNSStringFromDeviceAddress func(unsafe.Pointer) unsafe.Pointer
	_IOBluetoothNSStringFromDeviceAddressColon func(unsafe.Pointer) unsafe.Pointer
	_IOBluetoothNSStringToDeviceAddress func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothNumberOfAvailableHIDDevices func() unsafe.Pointer
	_IOBluetoothNumberOfKeyboardHIDDevices func() unsafe.Pointer
	_IOBluetoothNumberOfPointingHIDDevices func() unsafe.Pointer
	_IOBluetoothNumberOfTabletHIDDevices func() unsafe.Pointer
	_IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothOBEXSessionOpenTransportConnection func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothPackData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothPackDataList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothRemoveIgnoredHIDDevice func(unsafe.Pointer)
	_IOBluetoothRemoveSCOAudioDevice func(unsafe.Pointer) unsafe.Pointer
	_IOBluetoothUnpackData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothUnpackDataList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothUserNotificationUnregister func(unsafe.Pointer)
	_OBEXAddApplicationParameterHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddAuthorizationChallengeHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddAuthorizationResponseHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddBodyHeader func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXAddByteSequenceHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddConnectionIDHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddCountHeader func(uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddDescriptionHeader func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXAddHTTPHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddLengthHeader func(uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddNameHeader func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXAddObjectClassHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddTargetHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddTime4ByteHeader func(uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddTimeISOHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddTypeHeader func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXAddUserDefinedHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXAddWhoHeader func(unsafe.Pointer, uint32, unsafe.Pointer) unsafe.Pointer
	_OBEXCreateVCard func(unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32) unsafe.Pointer
	_OBEXCreateVEvent func(unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32) unsafe.Pointer
	_OBEXGetHeaders func(unsafe.Pointer, uintptr) unsafe.Pointer
	_OBEXHeadersToBytes func(unsafe.Pointer) unsafe.Pointer
	_OBEXSessionAbort func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionAbortResponse func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionConnect func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionConnectResponse func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionDelete func(unsafe.Pointer) unsafe.Pointer
	_OBEXSessionDisconnect func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionDisconnectResponse func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionGet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionGetAvailableCommandPayloadLength func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionGetAvailableCommandResponsePayloadLength func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionGetMaxPacketLength func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionGetResponse func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionHasOpenOBEXConnection func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionPut func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionPutResponse func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionSetPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionSetPathResponse func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OBEXSessionSetServerCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_IOBluetoothAddSCOAudioDevice, lib, "IOBluetoothAddSCOAudioDevice")
	tryRegister(&_IOBluetoothFindNumberOfRegistryEntriesOfClassName, lib, "IOBluetoothFindNumberOfRegistryEntriesOfClassName")
	tryRegister(&_IOBluetoothGetUniqueFileNameAndPath, lib, "IOBluetoothGetUniqueFileNameAndPath")
	tryRegister(&_IOBluetoothIgnoreHIDDevice, lib, "IOBluetoothIgnoreHIDDevice")
	tryRegister(&_IOBluetoothIsFileAppleDesignatedPIMData, lib, "IOBluetoothIsFileAppleDesignatedPIMData")
	tryRegister(&_IOBluetoothL2CAPChannelRegisterForChannelCloseNotification, lib, "IOBluetoothL2CAPChannelRegisterForChannelCloseNotification")
	tryRegister(&_IOBluetoothNSStringFromDeviceAddress, lib, "IOBluetoothNSStringFromDeviceAddress")
	tryRegister(&_IOBluetoothNSStringFromDeviceAddressColon, lib, "IOBluetoothNSStringFromDeviceAddressColon")
	tryRegister(&_IOBluetoothNSStringToDeviceAddress, lib, "IOBluetoothNSStringToDeviceAddress")
	tryRegister(&_IOBluetoothNumberOfAvailableHIDDevices, lib, "IOBluetoothNumberOfAvailableHIDDevices")
	tryRegister(&_IOBluetoothNumberOfKeyboardHIDDevices, lib, "IOBluetoothNumberOfKeyboardHIDDevices")
	tryRegister(&_IOBluetoothNumberOfPointingHIDDevices, lib, "IOBluetoothNumberOfPointingHIDDevices")
	tryRegister(&_IOBluetoothNumberOfTabletHIDDevices, lib, "IOBluetoothNumberOfTabletHIDDevices")
	tryRegister(&_IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber, lib, "IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber")
	tryRegister(&_IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef, lib, "IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef")
	tryRegister(&_IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel, lib, "IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel")
	tryRegister(&_IOBluetoothOBEXSessionOpenTransportConnection, lib, "IOBluetoothOBEXSessionOpenTransportConnection")
	tryRegister(&_IOBluetoothPackData, lib, "IOBluetoothPackData")
	tryRegister(&_IOBluetoothPackDataList, lib, "IOBluetoothPackDataList")
	tryRegister(&_IOBluetoothRemoveIgnoredHIDDevice, lib, "IOBluetoothRemoveIgnoredHIDDevice")
	tryRegister(&_IOBluetoothRemoveSCOAudioDevice, lib, "IOBluetoothRemoveSCOAudioDevice")
	tryRegister(&_IOBluetoothUnpackData, lib, "IOBluetoothUnpackData")
	tryRegister(&_IOBluetoothUnpackDataList, lib, "IOBluetoothUnpackDataList")
	tryRegister(&_IOBluetoothUserNotificationUnregister, lib, "IOBluetoothUserNotificationUnregister")
	tryRegister(&_OBEXAddApplicationParameterHeader, lib, "OBEXAddApplicationParameterHeader")
	tryRegister(&_OBEXAddAuthorizationChallengeHeader, lib, "OBEXAddAuthorizationChallengeHeader")
	tryRegister(&_OBEXAddAuthorizationResponseHeader, lib, "OBEXAddAuthorizationResponseHeader")
	tryRegister(&_OBEXAddBodyHeader, lib, "OBEXAddBodyHeader")
	tryRegister(&_OBEXAddByteSequenceHeader, lib, "OBEXAddByteSequenceHeader")
	tryRegister(&_OBEXAddConnectionIDHeader, lib, "OBEXAddConnectionIDHeader")
	tryRegister(&_OBEXAddCountHeader, lib, "OBEXAddCountHeader")
	tryRegister(&_OBEXAddDescriptionHeader, lib, "OBEXAddDescriptionHeader")
	tryRegister(&_OBEXAddHTTPHeader, lib, "OBEXAddHTTPHeader")
	tryRegister(&_OBEXAddLengthHeader, lib, "OBEXAddLengthHeader")
	tryRegister(&_OBEXAddNameHeader, lib, "OBEXAddNameHeader")
	tryRegister(&_OBEXAddObjectClassHeader, lib, "OBEXAddObjectClassHeader")
	tryRegister(&_OBEXAddTargetHeader, lib, "OBEXAddTargetHeader")
	tryRegister(&_OBEXAddTime4ByteHeader, lib, "OBEXAddTime4ByteHeader")
	tryRegister(&_OBEXAddTimeISOHeader, lib, "OBEXAddTimeISOHeader")
	tryRegister(&_OBEXAddTypeHeader, lib, "OBEXAddTypeHeader")
	tryRegister(&_OBEXAddUserDefinedHeader, lib, "OBEXAddUserDefinedHeader")
	tryRegister(&_OBEXAddWhoHeader, lib, "OBEXAddWhoHeader")
	tryRegister(&_OBEXCreateVCard, lib, "OBEXCreateVCard")
	tryRegister(&_OBEXCreateVEvent, lib, "OBEXCreateVEvent")
	tryRegister(&_OBEXGetHeaders, lib, "OBEXGetHeaders")
	tryRegister(&_OBEXHeadersToBytes, lib, "OBEXHeadersToBytes")
	tryRegister(&_OBEXSessionAbort, lib, "OBEXSessionAbort")
	tryRegister(&_OBEXSessionAbortResponse, lib, "OBEXSessionAbortResponse")
	tryRegister(&_OBEXSessionConnect, lib, "OBEXSessionConnect")
	tryRegister(&_OBEXSessionConnectResponse, lib, "OBEXSessionConnectResponse")
	tryRegister(&_OBEXSessionDelete, lib, "OBEXSessionDelete")
	tryRegister(&_OBEXSessionDisconnect, lib, "OBEXSessionDisconnect")
	tryRegister(&_OBEXSessionDisconnectResponse, lib, "OBEXSessionDisconnectResponse")
	tryRegister(&_OBEXSessionGet, lib, "OBEXSessionGet")
	tryRegister(&_OBEXSessionGetAvailableCommandPayloadLength, lib, "OBEXSessionGetAvailableCommandPayloadLength")
	tryRegister(&_OBEXSessionGetAvailableCommandResponsePayloadLength, lib, "OBEXSessionGetAvailableCommandResponsePayloadLength")
	tryRegister(&_OBEXSessionGetMaxPacketLength, lib, "OBEXSessionGetMaxPacketLength")
	tryRegister(&_OBEXSessionGetResponse, lib, "OBEXSessionGetResponse")
	tryRegister(&_OBEXSessionHasOpenOBEXConnection, lib, "OBEXSessionHasOpenOBEXConnection")
	tryRegister(&_OBEXSessionPut, lib, "OBEXSessionPut")
	tryRegister(&_OBEXSessionPutResponse, lib, "OBEXSessionPutResponse")
	tryRegister(&_OBEXSessionSetPath, lib, "OBEXSessionSetPath")
	tryRegister(&_OBEXSessionSetPathResponse, lib, "OBEXSessionSetPathResponse")
	tryRegister(&_OBEXSessionSetServerCallback, lib, "OBEXSessionSetServerCallback")
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



// Creates a persistent audio driver that will route audio data to/from the specified device.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.0.
// Creates a persistent audio driver that will route audio data to/from the specified device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothAddSCOAudioDevice
func IOBluetoothAddSCOAudioDevice(device unsafe.Pointer, configDict unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothAddSCOAudioDevice(device, configDict)
}

// The number of registry entries with a device classname.

// The number of registry entries with a device classname.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothFindNumberOfRegistryEntriesOfClassName(_:)
func IOBluetoothFindNumberOfRegistryEntriesOfClassName(deviceType unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothFindNumberOfRegistryEntriesOfClassName(deviceType)
}

// IOBluetoothGetUniqueFileNameAndPath is a IOBluetooth function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothGetUniqueFileNameAndPath(_:_:)
func IOBluetoothGetUniqueFileNameAndPath(inName unsafe.Pointer, inPath unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothGetUniqueFileNameAndPath(inName, inPath)
}

// Hints that the macOS Bluetooth software should ignore a HID device that connects up.

// Hints that the macOS Bluetooth software should ignore a HID device that connects up.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothIgnoreHIDDevice(_:)
func IOBluetoothIgnoreHIDDevice(device unsafe.Pointer) {
	_IOBluetoothIgnoreHIDDevice(device)
}

// Apple designated PIM data is classified as: .vcard, .vcal, .vcf, .vnote, .vmsg, .vcs

// Apple designated PIM data is classified as: .vcard, .vcal, .vcf, .vnote, .vmsg, .vcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothIsFileAppleDesignatedPIMData(_:)
func IOBluetoothIsFileAppleDesignatedPIMData(inFileName unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothIsFileAppleDesignatedPIMData(inFileName)
}

// Allows a client to register for a channel close notification.

// Allows a client to register for a channel close notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelRegisterForChannelCloseNotification(_:_:_:)
func IOBluetoothL2CAPChannelRegisterForChannelCloseNotification(channel unsafe.Pointer, callback unsafe.Pointer, inRefCon unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothL2CAPChannelRegisterForChannelCloseNotification(channel, callback, inRefCon)
}

// Convenience routine to take a device address structure and create an NSString.

// Convenience routine to take a device address structure and create an NSString.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNSStringFromDeviceAddress(_:)
func IOBluetoothNSStringFromDeviceAddress(deviceAddress unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothNSStringFromDeviceAddress(deviceAddress)
}

// IOBluetoothNSStringFromDeviceAddressColon is a IOBluetooth function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNSStringFromDeviceAddressColon(_:)
func IOBluetoothNSStringFromDeviceAddressColon(deviceAddress unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothNSStringFromDeviceAddressColon(deviceAddress)
}

// Convenience routine to take an NSString and turn it into a BluetoothDeviceAddress structure.

// Convenience routine to take an NSString and turn it into a BluetoothDeviceAddress structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNSStringToDeviceAddress(_:_:)
func IOBluetoothNSStringToDeviceAddress(inNameString unsafe.Pointer, outDeviceAddress unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothNSStringToDeviceAddress(inNameString, outDeviceAddress)
}

// Returns total number of HID devices on the system (Bluetooth + USB)

// Returns total number of HID devices on the system (Bluetooth + USB)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfAvailableHIDDevices()
func IOBluetoothNumberOfAvailableHIDDevices() unsafe.Pointer {
	return _IOBluetoothNumberOfAvailableHIDDevices()
}

// Returns number of keyboard HID devices on the system (Bluetooth + USB)

// Returns number of keyboard HID devices on the system (Bluetooth + USB)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfKeyboardHIDDevices()
func IOBluetoothNumberOfKeyboardHIDDevices() unsafe.Pointer {
	return _IOBluetoothNumberOfKeyboardHIDDevices()
}

// Returns number of “pointing” HID devices on the system (Bluetooth + USB)

// Returns number of “pointing” HID devices on the system (Bluetooth + USB)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfPointingHIDDevices()
func IOBluetoothNumberOfPointingHIDDevices() unsafe.Pointer {
	return _IOBluetoothNumberOfPointingHIDDevices()
}

// Returns number of “Tablet” HID devices on the system (Bluetooth + USB)

// Returns number of “Tablet” HID devices on the system (Bluetooth + USB)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfTabletHIDDevices()
func IOBluetoothNumberOfTabletHIDDevices() unsafe.Pointer {
	return _IOBluetoothNumberOfTabletHIDDevices()
}

// Create an OBEX session with a device ref and an RFCOMM channel ID. This allows you to bypass the browser if you already know the SDP information.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Create an OBEX session with a device ref and an RFCOMM channel ID. This allows you to bypass the browser if you already know the SDP information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber
func IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber(inDeviceRef unsafe.Pointer, inChannelID unsafe.Pointer, outSessionRef unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber(inDeviceRef, inChannelID, outSessionRef)
}

// Create an OBEX session with a service ref, usually obtained from the device browser.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Create an OBEX session with a service ref, usually obtained from the device browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef
func IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef(inSDPServiceRef unsafe.Pointer, outSessionRef unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef(inSDPServiceRef, outSessionRef)
}

// Create an OBEX session with an IOBluetoothRFCOMMchannel. This implies you are creating a OBEX SERVER session that will dole out info to remote Bluetooth clients.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Create an OBEX session with an IOBluetoothRFCOMMchannel. This implies you are creating a OBEX SERVER session that will dole out info to remote Bluetooth clients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel
func IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel(inRFCOMMChannelRef unsafe.Pointer, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer, outSessionRef unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel(inRFCOMMChannelRef, inCallback, inUserRefCon, outSessionRef)
}

// IOBluetoothOBEXSessionOpenTransportConnection is a IOBluetooth function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionOpenTransportConnection
func IOBluetoothOBEXSessionOpenTransportConnection(inSessionRef unsafe.Pointer, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothOBEXSessionOpenTransportConnection(inSessionRef, inCallback, inUserRefCon)
}

// Packs a variable amount of parameters into a buffer according to a printf-style format string.

// Packs a variable amount of parameters into a buffer according to a printf-style format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPackData
func IOBluetoothPackData(ioBuffer unsafe.Pointer, inFormat unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothPackData(ioBuffer, inFormat)
}

// IOBluetoothPackDataList is a IOBluetooth function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPackDataList(_:_:_:)
func IOBluetoothPackDataList(ioBuffer unsafe.Pointer, inFormat unsafe.Pointer, inArgs unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothPackDataList(ioBuffer, inFormat, inArgs)
}

// The counterpart to the above IOBluetoothIgnoreHIDDevice() API.

// The counterpart to the above IOBluetoothIgnoreHIDDevice() API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRemoveIgnoredHIDDevice(_:)
func IOBluetoothRemoveIgnoredHIDDevice(device unsafe.Pointer) {
	_IOBluetoothRemoveIgnoredHIDDevice(device)
}

// Removes a persistent audio driver for a device that had already been added using IOBluetoothAddAudioDevice().
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.0.
// Removes a persistent audio driver for a device that had already been added using IOBluetoothAddAudioDevice().
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRemoveSCOAudioDevice
func IOBluetoothRemoveSCOAudioDevice(device unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothRemoveSCOAudioDevice(device)
}

// Unpacks a variable amount of data from a buffer into a variable number of parameters according to a printf-style format string.

// Unpacks a variable amount of data from a buffer into a variable number of parameters according to a printf-style format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUnpackData
func IOBluetoothUnpackData(inBufferSize unsafe.Pointer, inBuffer unsafe.Pointer, inFormat unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothUnpackData(inBufferSize, inBuffer, inFormat)
}

// IOBluetoothUnpackDataList is a IOBluetooth function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUnpackDataList(_:_:_:_:)
func IOBluetoothUnpackDataList(inBufferSize unsafe.Pointer, inBuffer unsafe.Pointer, inFormat unsafe.Pointer, inArgs unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothUnpackDataList(inBufferSize, inBuffer, inFormat, inArgs)
}

// Unregisters the target notification.

// Unregisters the target notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUserNotificationUnregister(_:)
func IOBluetoothUserNotificationUnregister(notificationRef unsafe.Pointer) {
	_IOBluetoothUserNotificationUnregister(notificationRef)
}

// Add bytes representing an application parameter to a dictionary of OBEX headers.

// Add bytes representing an application parameter to a dictionary of OBEX headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddApplicationParameterHeader(_:_:_:)
func OBEXAddApplicationParameterHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddApplicationParameterHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Add an authorization challenge header to a dictionary of OBEXheaders.

// Add an authorization challenge header to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddAuthorizationChallengeHeader(_:_:_:)
func OBEXAddAuthorizationChallengeHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddAuthorizationChallengeHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Add an authorization Response header to a dictionary of OBEXheaders.

// Add an authorization Response header to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddAuthorizationResponseHeader(_:_:_:)
func OBEXAddAuthorizationResponseHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddAuthorizationResponseHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Add bytes of data to a dictionary of OBEXheaders.

// Add bytes of data to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddBodyHeader(_:_:_:_:)
func OBEXAddBodyHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, isEndOfBody unsafe.Pointer, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddBodyHeader(inHeaderData, inHeaderDataLength, isEndOfBody, dictRef)
}

// Add a byte sequence header to a dictionary of OBEXheaders.

// Add a byte sequence header to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddByteSequenceHeader(_:_:_:)
func OBEXAddByteSequenceHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddByteSequenceHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Add bytes representing a connection ID to a dictionary of OBEX headers.

// Add bytes representing a connection ID to a dictionary of OBEX headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddConnectionIDHeader(_:_:_:)
func OBEXAddConnectionIDHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddConnectionIDHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Add a CFStringRef to a dictionary of OBEXheaders.

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddCountHeader(_:_:)
func OBEXAddCountHeader(count uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddCountHeader(count, dictRef)
}

// Add a CFStringRef to a dictionary of OBEXheaders.

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddDescriptionHeader(_:_:)
func OBEXAddDescriptionHeader(description unsafe.Pointer, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddDescriptionHeader(description, dictRef)
}

// Add bytes of data to a dictionary of OBEXheaders.

// Add bytes of data to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddHTTPHeader(_:_:_:)
func OBEXAddHTTPHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddHTTPHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Add a CFStringRef to a dictionary of OBEXheaders.

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddLengthHeader(_:_:)
func OBEXAddLengthHeader(length uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddLengthHeader(length, dictRef)
}

// Add a CFStringRef to a dictionary of OBEXheaders.

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddNameHeader(_:_:)
func OBEXAddNameHeader(name unsafe.Pointer, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddNameHeader(name, dictRef)
}

// Add an object class header to a dictionary of OBEXheaders.

// Add an object class header to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddObjectClassHeader(_:_:_:)
func OBEXAddObjectClassHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddObjectClassHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Add bytes of data to a dictionary of OBEXheaders.

// Add bytes of data to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTargetHeader(_:_:_:)
func OBEXAddTargetHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddTargetHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Add a CFStringRef to a dictionary of OBEXheaders.

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTime4ByteHeader(_:_:)
func OBEXAddTime4ByteHeader(time4Byte uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddTime4ByteHeader(time4Byte, dictRef)
}

// Add bytes to a dictionary of OBEXheaders.

// Add bytes to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTimeISOHeader(_:_:_:)
func OBEXAddTimeISOHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddTimeISOHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Add a CFStringRef to a dictionary of OBEXheaders.

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTypeHeader(_:_:)
func OBEXAddTypeHeader(type_ unsafe.Pointer, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddTypeHeader(type_, dictRef)
}

// Add a user-defined custom header to a dictionary of OBEXheaders.

// Add a user-defined custom header to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddUserDefinedHeader(_:_:_:)
func OBEXAddUserDefinedHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddUserDefinedHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Add bytes of data to a dictionary of OBEXheaders.

// Add bytes of data to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddWhoHeader(_:_:_:)
func OBEXAddWhoHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXAddWhoHeader(inHeaderData, inHeaderDataLength, dictRef)
}

// Creates a formatted vCard, ready to be sent over OBEX or whatever.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Creates a formatted vCard, ready to be sent over OBEX or whatever.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXCreateVCard
func OBEXCreateVCard(inFirstName unsafe.Pointer, inFirstNameLength uint32, inLastName unsafe.Pointer, inLastNameLength uint32, inFriendlyName unsafe.Pointer, inFriendlyNameLength uint32, inNameCharset unsafe.Pointer, inNameCharsetLength uint32, inHomePhone unsafe.Pointer, inHomePhoneLength uint32, inWorkPhone unsafe.Pointer, inWorkPhoneLength uint32, inCellPhone unsafe.Pointer, inCellPhoneLength uint32, inFaxPhone unsafe.Pointer, inFaxPhoneLength uint32, inEMailAddress unsafe.Pointer, inEMailAddressLength uint32, inEMailAddressCharset unsafe.Pointer, inEMailAddressCharsetLength uint32, inOrganization unsafe.Pointer, inOrganizationLength uint32, inOrganizationCharset unsafe.Pointer, inOrganizationCharsetLength uint32, inTitle unsafe.Pointer, inTitleLength uint32, inTitleCharset unsafe.Pointer, inTitleCharsetLength uint32) unsafe.Pointer {
	return _OBEXCreateVCard(inFirstName, inFirstNameLength, inLastName, inLastNameLength, inFriendlyName, inFriendlyNameLength, inNameCharset, inNameCharsetLength, inHomePhone, inHomePhoneLength, inWorkPhone, inWorkPhoneLength, inCellPhone, inCellPhoneLength, inFaxPhone, inFaxPhoneLength, inEMailAddress, inEMailAddressLength, inEMailAddressCharset, inEMailAddressCharsetLength, inOrganization, inOrganizationLength, inOrganizationCharset, inOrganizationCharsetLength, inTitle, inTitleLength, inTitleCharset, inTitleCharsetLength)
}

// Creates a formatted vEvent, ready to be sent over OBEX or whatever. You probably will embed the output in a vCalendar event.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Creates a formatted vEvent, ready to be sent over OBEX or whatever. You probably will embed the output in a vCalendar event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXCreateVEvent
func OBEXCreateVEvent(inCharset unsafe.Pointer, inCharsetLength uint32, inEncoding unsafe.Pointer, inEncodingLength uint32, inEventStartDate unsafe.Pointer, inEventStartDateLength uint32, inEventEndDate unsafe.Pointer, inEventEndDateLength uint32, inAlarmDate unsafe.Pointer, inAlarmDateLength uint32, inCategory unsafe.Pointer, inCategoryLength uint32, inSummary unsafe.Pointer, inSummaryLength uint32, inLocation unsafe.Pointer, inLocationLength uint32, inXIRMCLUID unsafe.Pointer, inXIRMCLUIDLength uint32) unsafe.Pointer {
	return _OBEXCreateVEvent(inCharset, inCharsetLength, inEncoding, inEncodingLength, inEventStartDate, inEventStartDateLength, inEventEndDate, inEventEndDateLength, inAlarmDate, inAlarmDateLength, inCategory, inCategoryLength, inSummary, inSummaryLength, inLocation, inLocationLength, inXIRMCLUID, inXIRMCLUIDLength)
}

// Take a data blob and looks for OBEX headers.

// Take a data blob and looks for OBEX headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXGetHeaders(_:_:)
func OBEXGetHeaders(inData unsafe.Pointer, inDataSize uintptr) unsafe.Pointer {
	return _OBEXGetHeaders(inData, inDataSize)
}

// Converts a dictionary of headers to a data pointer, from which you can extract as bytes and pass to the OBEX command/response functions.

// Converts a dictionary of headers to a data pointer, from which you can extract as bytes and pass to the OBEX command/response functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXHeadersToBytes(_:)
func OBEXHeadersToBytes(dictionaryOfHeaders unsafe.Pointer) unsafe.Pointer {
	return _OBEXHeadersToBytes(dictionaryOfHeaders)
}

// Send an abort command to a remote OBEX server.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send an abort command to a remote OBEX server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionAbort
func OBEXSessionAbort(inSessionRef unsafe.Pointer, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionAbort(inSessionRef, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}

// Send a response to a abort command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a abort command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionAbortResponse
func OBEXSessionAbortResponse(inSessionRef unsafe.Pointer, inResponseOpCode unsafe.Pointer, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionAbortResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}

// Establishes an OBEX connection to the target device for the session. If a transport connection is not open yet, it will be opened if possible.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Establishes an OBEX connection to the target device for the session. If a transport connection is not open yet, it will be opened if possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionConnect
func OBEXSessionConnect(inSessionRef unsafe.Pointer, inFlags unsafe.Pointer, inMaxPacketLength unsafe.Pointer, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionConnect(inSessionRef, inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}

// Send a response to a connect command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a connect command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionConnectResponse
func OBEXSessionConnectResponse(inSessionRef unsafe.Pointer, inResponseOpCode unsafe.Pointer, inFlags unsafe.Pointer, inMaxPacketLength unsafe.Pointer, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionConnectResponse(inSessionRef, inResponseOpCode, inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}

// Destroy an OBEX session. If connections are open, they will (eventually) be terminated for you.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Destroy an OBEX session. If connections are open, they will (eventually) be terminated for you.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionDelete
func OBEXSessionDelete(inSessionRef unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionDelete(inSessionRef)
}

// Send a disconnect command to a remote OBEX server.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a disconnect command to a remote OBEX server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionDisconnect
func OBEXSessionDisconnect(inSessionRef unsafe.Pointer, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionDisconnect(inSessionRef, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}

// Send a response to a disconnect command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a disconnect command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionDisconnectResponse
func OBEXSessionDisconnectResponse(inSessionRef unsafe.Pointer, inResponseOpCode unsafe.Pointer, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionDisconnectResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}

// Send a get command to a remote OBEX server.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a get command to a remote OBEX server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGet
func OBEXSessionGet(inSessionRef unsafe.Pointer, inIsFinalChunk unsafe.Pointer, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionGet(inSessionRef, inIsFinalChunk, inHeadersData, inHeadersDataLength, inCallback, inUserRefCon)
}

// Gets space available for your data for a particular command response you are trying to send.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Gets space available for your data for a particular command response you are trying to send.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetAvailableCommandPayloadLength
func OBEXSessionGetAvailableCommandPayloadLength(inSessionRef unsafe.Pointer, inOpCode unsafe.Pointer, outLength unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionGetAvailableCommandPayloadLength(inSessionRef, inOpCode, outLength)
}

// Gets space available for your data for a particular command response you are trying to send.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Gets space available for your data for a particular command response you are trying to send.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetAvailableCommandResponsePayloadLength
func OBEXSessionGetAvailableCommandResponsePayloadLength(inSessionRef unsafe.Pointer, inOpCode unsafe.Pointer, outLength unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionGetAvailableCommandResponsePayloadLength(inSessionRef, inOpCode, outLength)
}

// Gets current max packet length.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Gets current max packet length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetMaxPacketLength
func OBEXSessionGetMaxPacketLength(inSessionRef unsafe.Pointer, outLength unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionGetMaxPacketLength(inSessionRef, outLength)
}

// Send a response to a get command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a get command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetResponse
func OBEXSessionGetResponse(inSessionRef unsafe.Pointer, inResponseOpCode unsafe.Pointer, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionGetResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}

// Allows you to test the session for an open OBEX connection for a particular session.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Allows you to test the session for an open OBEX connection for a particular session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionHasOpenOBEXConnection
func OBEXSessionHasOpenOBEXConnection(inSessionRef unsafe.Pointer, outIsConnected unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionHasOpenOBEXConnection(inSessionRef, outIsConnected)
}

// Send a put command to a remote OBEX server.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a put command to a remote OBEX server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionPut
func OBEXSessionPut(inSessionRef unsafe.Pointer, inIsFinalChunk unsafe.Pointer, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inBodyData unsafe.Pointer, inBodyDataLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionPut(inSessionRef, inIsFinalChunk, inHeadersData, inHeadersDataLength, inBodyData, inBodyDataLength, inCallback, inUserRefCon)
}

// Send a response to a put command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a put command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionPutResponse
func OBEXSessionPutResponse(inSessionRef unsafe.Pointer, inResponseOpCode unsafe.Pointer, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionPutResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}

// Send a set path command to a remote OBEX server.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a set path command to a remote OBEX server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionSetPath
func OBEXSessionSetPath(inSessionRef unsafe.Pointer, inFlags unsafe.Pointer, inConstants unsafe.Pointer, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionSetPath(inSessionRef, inFlags, inConstants, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}

// Send a response to a set path command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a set path command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionSetPathResponse
func OBEXSessionSetPathResponse(inSessionRef unsafe.Pointer, inResponseOpCode unsafe.Pointer, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionSetPathResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}

// OBEXSessionSetServerCallback is a IOBluetooth function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionSetServerCallback
func OBEXSessionSetServerCallback(inSessionRef unsafe.Pointer, inCallback unsafe.Pointer, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	return _OBEXSessionSetServerCallback(inSessionRef, inCallback, inUserRefCon)
}



