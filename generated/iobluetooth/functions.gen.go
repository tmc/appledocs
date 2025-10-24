// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

/* debug [functions.gen.go]: Generating 64 functions for IOBluetooth */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// IOBluetooth Functions (64 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_IOBluetoothAddSCOAudioDevice func(BluetoothDeviceRef, DictionaryRef) int
	_IOBluetoothFindNumberOfRegistryEntriesOfClassName func(unsafe.Pointer) unsafe.Pointer
	_IOBluetoothGetUniqueFileNameAndPath func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothIgnoreHIDDevice func(BluetoothDeviceRef)
	_IOBluetoothIsFileAppleDesignatedPIMData func(unsafe.Pointer) unsafe.Pointer
	_IOBluetoothL2CAPChannelRegisterForChannelCloseNotification func(BluetoothL2CAPChannelRef, BluetoothUserNotificationCallback, unsafe.Pointer) BluetoothUserNotificationRef
	_IOBluetoothNSStringFromDeviceAddress func(unsafe.Pointer) unsafe.Pointer
	_IOBluetoothNSStringFromDeviceAddressColon func(unsafe.Pointer) unsafe.Pointer
	_IOBluetoothNSStringToDeviceAddress func(unsafe.Pointer, unsafe.Pointer) int
	_IOBluetoothNumberOfAvailableHIDDevices func() unsafe.Pointer
	_IOBluetoothNumberOfKeyboardHIDDevices func() unsafe.Pointer
	_IOBluetoothNumberOfPointingHIDDevices func() unsafe.Pointer
	_IOBluetoothNumberOfTabletHIDDevices func() unsafe.Pointer
	_IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel func(BluetoothRFCOMMChannelRef, OBEXSessionEventCallback, unsafe.Pointer, unsafe.Pointer) OBEXError
	_IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber func(BluetoothDeviceRef, BluetoothRFCOMMChannelID, unsafe.Pointer) OBEXError
	_IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef func(BluetoothSDPServiceRecordRef, unsafe.Pointer) OBEXError
	_IOBluetoothOBEXSessionOpenTransportConnection func(OBEXSessionRef, BluetoothOBEXSessionOpenConnectionCallback, unsafe.Pointer) OBEXError
	_IOBluetoothPackData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothPackDataList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothRemoveIgnoredHIDDevice func(BluetoothDeviceRef)
	_IOBluetoothRemoveSCOAudioDevice func(BluetoothDeviceRef) int
	_IOBluetoothUnpackData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothUnpackDataList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBluetoothUserNotificationUnregister func(BluetoothUserNotificationRef)
	_OBEXAddApplicationParameterHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXAddAuthorizationChallengeHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXAddAuthorizationResponseHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXAddBodyHeader func(unsafe.Pointer, uint32, unsafe.Pointer, MutableDictionaryRef) OBEXError
	_OBEXAddByteSequenceHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXAddConnectionIDHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXAddCountHeader func(uint32, MutableDictionaryRef) OBEXError
	_OBEXAddDescriptionHeader func(StringRef, MutableDictionaryRef) OBEXError
	_OBEXAddHTTPHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXAddLengthHeader func(uint32, MutableDictionaryRef) OBEXError
	_OBEXAddNameHeader func(StringRef, MutableDictionaryRef) OBEXError
	_OBEXAddObjectClassHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXAddTargetHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXAddTime4ByteHeader func(uint32, MutableDictionaryRef) OBEXError
	_OBEXAddTimeISOHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXAddTypeHeader func(StringRef, MutableDictionaryRef) OBEXError
	_OBEXAddUserDefinedHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXAddWhoHeader func(unsafe.Pointer, uint32, MutableDictionaryRef) OBEXError
	_OBEXCreateVCard func(unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32) DataRef
	_OBEXCreateVEvent func(unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32) DataRef
	_OBEXGetHeaders func(unsafe.Pointer, uintptr) DictionaryRef
	_OBEXHeadersToBytes func(DictionaryRef) MutableDataRef
	_OBEXSessionAbort func(OBEXSessionRef, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionAbortResponse func(OBEXSessionRef, OBEXOpCode, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionConnect func(OBEXSessionRef, OBEXFlags, OBEXMaxPacketLength, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionConnectResponse func(OBEXSessionRef, OBEXOpCode, OBEXFlags, OBEXMaxPacketLength, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionDelete func(OBEXSessionRef) OBEXError
	_OBEXSessionDisconnect func(OBEXSessionRef, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionDisconnectResponse func(OBEXSessionRef, OBEXOpCode, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionGet func(OBEXSessionRef, unsafe.Pointer, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionGetAvailableCommandPayloadLength func(OBEXSessionRef, OBEXOpCode, unsafe.Pointer) OBEXError
	_OBEXSessionGetAvailableCommandResponsePayloadLength func(OBEXSessionRef, OBEXOpCode, unsafe.Pointer) OBEXError
	_OBEXSessionGetMaxPacketLength func(OBEXSessionRef, unsafe.Pointer) OBEXError
	_OBEXSessionGetResponse func(OBEXSessionRef, OBEXOpCode, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionHasOpenOBEXConnection func(OBEXSessionRef, unsafe.Pointer) OBEXError
	_OBEXSessionPut func(OBEXSessionRef, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionPutResponse func(OBEXSessionRef, OBEXOpCode, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionSetPath func(OBEXSessionRef, OBEXFlags, OBEXConstants, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionSetPathResponse func(OBEXSessionRef, OBEXOpCode, unsafe.Pointer, uintptr, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
	_OBEXSessionSetServerCallback func(OBEXSessionRef, OBEXSessionEventCallback, unsafe.Pointer) OBEXError
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
	tryRegister(&_IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel, lib, "IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel")
	tryRegister(&_IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber, lib, "IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber")
	tryRegister(&_IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef, lib, "IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef")
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
func IOBluetoothAddSCOAudioDevice(device BluetoothDeviceRef, configDict DictionaryRef) int {
	return _IOBluetoothAddSCOAudioDevice(device, configDict)
}/* debug [functions.gen.go/function]: IOBluetoothAddSCOAudioDevice */

// The number of registry entries with a device classname.
//
// Added in macOS .
// The number of registry entries with a device classname.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothFindNumberOfRegistryEntriesOfClassName(_:)
func IOBluetoothFindNumberOfRegistryEntriesOfClassName(deviceType unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothFindNumberOfRegistryEntriesOfClassName(deviceType)
}/* debug [functions.gen.go/function]: IOBluetoothFindNumberOfRegistryEntriesOfClassName */

// IOBluetoothGetUniqueFileNameAndPath is a IOBluetooth function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothGetUniqueFileNameAndPath(_:_:)
func IOBluetoothGetUniqueFileNameAndPath(inName unsafe.Pointer, inPath unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothGetUniqueFileNameAndPath(inName, inPath)
}/* debug [functions.gen.go/function]: IOBluetoothGetUniqueFileNameAndPath */

// Hints that the macOS Bluetooth software should ignore a HID device that connects up.
//
// Added in macOS .
// Hints that the macOS Bluetooth software should ignore a HID device that connects up.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothIgnoreHIDDevice(_:)
func IOBluetoothIgnoreHIDDevice(device BluetoothDeviceRef) {
	_IOBluetoothIgnoreHIDDevice(device)
}/* debug [functions.gen.go/function]: IOBluetoothIgnoreHIDDevice */

// Apple designated PIM data is classified as: .vcard, .vcal, .vcf, .vnote, .vmsg, .vcs
//
// Added in macOS .
// Apple designated PIM data is classified as: .vcard, .vcal, .vcf, .vnote, .vmsg, .vcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothIsFileAppleDesignatedPIMData(_:)
func IOBluetoothIsFileAppleDesignatedPIMData(inFileName unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothIsFileAppleDesignatedPIMData(inFileName)
}/* debug [functions.gen.go/function]: IOBluetoothIsFileAppleDesignatedPIMData */

// Allows a client to register for a channel close notification.
//
// Added in macOS .
// Allows a client to register for a channel close notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelRegisterForChannelCloseNotification(_:_:_:)
func IOBluetoothL2CAPChannelRegisterForChannelCloseNotification(channel BluetoothL2CAPChannelRef, callback BluetoothUserNotificationCallback, inRefCon unsafe.Pointer) BluetoothUserNotificationRef {
	return _IOBluetoothL2CAPChannelRegisterForChannelCloseNotification(channel, callback, inRefCon)
}/* debug [functions.gen.go/function]: IOBluetoothL2CAPChannelRegisterForChannelCloseNotification */

// Convenience routine to take a device address structure and create an NSString.
//
// Added in macOS .
// Convenience routine to take a device address structure and create an NSString.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNSStringFromDeviceAddress(_:)
func IOBluetoothNSStringFromDeviceAddress(deviceAddress unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothNSStringFromDeviceAddress(deviceAddress)
}/* debug [functions.gen.go/function]: IOBluetoothNSStringFromDeviceAddress */

// IOBluetoothNSStringFromDeviceAddressColon is a IOBluetooth function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNSStringFromDeviceAddressColon(_:)
func IOBluetoothNSStringFromDeviceAddressColon(deviceAddress unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothNSStringFromDeviceAddressColon(deviceAddress)
}/* debug [functions.gen.go/function]: IOBluetoothNSStringFromDeviceAddressColon */

// Convenience routine to take an NSString and turn it into a BluetoothDeviceAddress structure.
//
// Added in macOS .
// Convenience routine to take an NSString and turn it into a BluetoothDeviceAddress structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNSStringToDeviceAddress(_:_:)
func IOBluetoothNSStringToDeviceAddress(inNameString unsafe.Pointer, outDeviceAddress unsafe.Pointer) int {
	return _IOBluetoothNSStringToDeviceAddress(inNameString, outDeviceAddress)
}/* debug [functions.gen.go/function]: IOBluetoothNSStringToDeviceAddress */

// Returns total number of HID devices on the system (Bluetooth + USB)
//
// Added in macOS .
// Returns total number of HID devices on the system (Bluetooth + USB)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfAvailableHIDDevices()
func IOBluetoothNumberOfAvailableHIDDevices() unsafe.Pointer {
	return _IOBluetoothNumberOfAvailableHIDDevices()
}/* debug [functions.gen.go/function]: IOBluetoothNumberOfAvailableHIDDevices */

// Returns number of keyboard HID devices on the system (Bluetooth + USB)
//
// Added in macOS .
// Returns number of keyboard HID devices on the system (Bluetooth + USB)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfKeyboardHIDDevices()
func IOBluetoothNumberOfKeyboardHIDDevices() unsafe.Pointer {
	return _IOBluetoothNumberOfKeyboardHIDDevices()
}/* debug [functions.gen.go/function]: IOBluetoothNumberOfKeyboardHIDDevices */

// Returns number of “pointing” HID devices on the system (Bluetooth + USB)
//
// Added in macOS .
// Returns number of “pointing” HID devices on the system (Bluetooth + USB)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfPointingHIDDevices()
func IOBluetoothNumberOfPointingHIDDevices() unsafe.Pointer {
	return _IOBluetoothNumberOfPointingHIDDevices()
}/* debug [functions.gen.go/function]: IOBluetoothNumberOfPointingHIDDevices */

// Returns number of “Tablet” HID devices on the system (Bluetooth + USB)
//
// Added in macOS .
// Returns number of “Tablet” HID devices on the system (Bluetooth + USB)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfTabletHIDDevices()
func IOBluetoothNumberOfTabletHIDDevices() unsafe.Pointer {
	return _IOBluetoothNumberOfTabletHIDDevices()
}/* debug [functions.gen.go/function]: IOBluetoothNumberOfTabletHIDDevices */

// Create an OBEX session with an IOBluetoothRFCOMMchannel. This implies you are creating a OBEX SERVER session that will dole out info to remote Bluetooth clients.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Create an OBEX session with an IOBluetoothRFCOMMchannel. This implies you are creating a OBEX SERVER session that will dole out info to remote Bluetooth clients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel
func IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel(inRFCOMMChannelRef BluetoothRFCOMMChannelRef, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer, outSessionRef unsafe.Pointer) OBEXError {
	return _IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel(inRFCOMMChannelRef, inCallback, inUserRefCon, outSessionRef)
}/* debug [functions.gen.go/function]: IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel */

// Create an OBEX session with a device ref and an RFCOMM channel ID. This allows you to bypass the browser if you already know the SDP information.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Create an OBEX session with a device ref and an RFCOMM channel ID. This allows you to bypass the browser if you already know the SDP information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber
func IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber(inDeviceRef BluetoothDeviceRef, inChannelID BluetoothRFCOMMChannelID, outSessionRef unsafe.Pointer) OBEXError {
	return _IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber(inDeviceRef, inChannelID, outSessionRef)
}/* debug [functions.gen.go/function]: IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber */

// Create an OBEX session with a service ref, usually obtained from the device browser.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Create an OBEX session with a service ref, usually obtained from the device browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef
func IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef(inSDPServiceRef BluetoothSDPServiceRecordRef, outSessionRef unsafe.Pointer) OBEXError {
	return _IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef(inSDPServiceRef, outSessionRef)
}/* debug [functions.gen.go/function]: IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef */

// IOBluetoothOBEXSessionOpenTransportConnection is a IOBluetooth function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionOpenTransportConnection
func IOBluetoothOBEXSessionOpenTransportConnection(inSessionRef OBEXSessionRef, inCallback BluetoothOBEXSessionOpenConnectionCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _IOBluetoothOBEXSessionOpenTransportConnection(inSessionRef, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: IOBluetoothOBEXSessionOpenTransportConnection */

// Packs a variable amount of parameters into a buffer according to a printf-style format string.
//
// Added in macOS .
// Packs a variable amount of parameters into a buffer according to a printf-style format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPackData
func IOBluetoothPackData(ioBuffer unsafe.Pointer, inFormat unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothPackData(ioBuffer, inFormat)
}/* debug [functions.gen.go/function]: IOBluetoothPackData */

// IOBluetoothPackDataList is a IOBluetooth function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPackDataList(_:_:_:)
func IOBluetoothPackDataList(ioBuffer unsafe.Pointer, inFormat unsafe.Pointer, inArgs unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothPackDataList(ioBuffer, inFormat, inArgs)
}/* debug [functions.gen.go/function]: IOBluetoothPackDataList */

// The counterpart to the above IOBluetoothIgnoreHIDDevice() API.
//
// Added in macOS .
// The counterpart to the above IOBluetoothIgnoreHIDDevice() API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRemoveIgnoredHIDDevice(_:)
func IOBluetoothRemoveIgnoredHIDDevice(device BluetoothDeviceRef) {
	_IOBluetoothRemoveIgnoredHIDDevice(device)
}/* debug [functions.gen.go/function]: IOBluetoothRemoveIgnoredHIDDevice */

// Removes a persistent audio driver for a device that had already been added using IOBluetoothAddAudioDevice().
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.0.
// Removes a persistent audio driver for a device that had already been added using IOBluetoothAddAudioDevice().
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRemoveSCOAudioDevice
func IOBluetoothRemoveSCOAudioDevice(device BluetoothDeviceRef) int {
	return _IOBluetoothRemoveSCOAudioDevice(device)
}/* debug [functions.gen.go/function]: IOBluetoothRemoveSCOAudioDevice */

// Unpacks a variable amount of data from a buffer into a variable number of parameters according to a printf-style format string.
//
// Added in macOS .
// Unpacks a variable amount of data from a buffer into a variable number of parameters according to a printf-style format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUnpackData
func IOBluetoothUnpackData(inBufferSize unsafe.Pointer, inBuffer unsafe.Pointer, inFormat unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothUnpackData(inBufferSize, inBuffer, inFormat)
}/* debug [functions.gen.go/function]: IOBluetoothUnpackData */

// IOBluetoothUnpackDataList is a IOBluetooth function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUnpackDataList(_:_:_:_:)
func IOBluetoothUnpackDataList(inBufferSize unsafe.Pointer, inBuffer unsafe.Pointer, inFormat unsafe.Pointer, inArgs unsafe.Pointer) unsafe.Pointer {
	return _IOBluetoothUnpackDataList(inBufferSize, inBuffer, inFormat, inArgs)
}/* debug [functions.gen.go/function]: IOBluetoothUnpackDataList */

// Unregisters the target notification.
//
// Added in macOS .
// Unregisters the target notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUserNotificationUnregister(_:)
func IOBluetoothUserNotificationUnregister(notificationRef BluetoothUserNotificationRef) {
	_IOBluetoothUserNotificationUnregister(notificationRef)
}/* debug [functions.gen.go/function]: IOBluetoothUserNotificationUnregister */

// Add bytes representing an application parameter to a dictionary of OBEX headers.
//
// Added in macOS .
// Add bytes representing an application parameter to a dictionary of OBEX headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddApplicationParameterHeader(_:_:_:)
func OBEXAddApplicationParameterHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddApplicationParameterHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddApplicationParameterHeader */

// Add an authorization challenge header to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add an authorization challenge header to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddAuthorizationChallengeHeader(_:_:_:)
func OBEXAddAuthorizationChallengeHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddAuthorizationChallengeHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddAuthorizationChallengeHeader */

// Add an authorization Response header to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add an authorization Response header to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddAuthorizationResponseHeader(_:_:_:)
func OBEXAddAuthorizationResponseHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddAuthorizationResponseHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddAuthorizationResponseHeader */

// Add bytes of data to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add bytes of data to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddBodyHeader(_:_:_:_:)
func OBEXAddBodyHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, isEndOfBody unsafe.Pointer, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddBodyHeader(inHeaderData, inHeaderDataLength, isEndOfBody, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddBodyHeader */

// Add a byte sequence header to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add a byte sequence header to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddByteSequenceHeader(_:_:_:)
func OBEXAddByteSequenceHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddByteSequenceHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddByteSequenceHeader */

// Add bytes representing a connection ID to a dictionary of OBEX headers.
//
// Added in macOS .
// Add bytes representing a connection ID to a dictionary of OBEX headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddConnectionIDHeader(_:_:_:)
func OBEXAddConnectionIDHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddConnectionIDHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddConnectionIDHeader */

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddCountHeader(_:_:)
func OBEXAddCountHeader(count uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddCountHeader(count, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddCountHeader */

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddDescriptionHeader(_:_:)
func OBEXAddDescriptionHeader(description StringRef, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddDescriptionHeader(description, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddDescriptionHeader */

// Add bytes of data to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add bytes of data to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddHTTPHeader(_:_:_:)
func OBEXAddHTTPHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddHTTPHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddHTTPHeader */

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddLengthHeader(_:_:)
func OBEXAddLengthHeader(length uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddLengthHeader(length, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddLengthHeader */

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddNameHeader(_:_:)
func OBEXAddNameHeader(name StringRef, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddNameHeader(name, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddNameHeader */

// Add an object class header to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add an object class header to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddObjectClassHeader(_:_:_:)
func OBEXAddObjectClassHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddObjectClassHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddObjectClassHeader */

// Add bytes of data to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add bytes of data to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTargetHeader(_:_:_:)
func OBEXAddTargetHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddTargetHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddTargetHeader */

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTime4ByteHeader(_:_:)
func OBEXAddTime4ByteHeader(time4Byte uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddTime4ByteHeader(time4Byte, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddTime4ByteHeader */

// Add bytes to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add bytes to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTimeISOHeader(_:_:_:)
func OBEXAddTimeISOHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddTimeISOHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddTimeISOHeader */

// Add a CFStringRef to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add a CFStringRef to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTypeHeader(_:_:)
func OBEXAddTypeHeader(type_ StringRef, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddTypeHeader(type_, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddTypeHeader */

// Add a user-defined custom header to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add a user-defined custom header to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddUserDefinedHeader(_:_:_:)
func OBEXAddUserDefinedHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddUserDefinedHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddUserDefinedHeader */

// Add bytes of data to a dictionary of OBEXheaders.
//
// Added in macOS .
// Add bytes of data to a dictionary of OBEXheaders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAddWhoHeader(_:_:_:)
func OBEXAddWhoHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef MutableDictionaryRef) OBEXError {
	return _OBEXAddWhoHeader(inHeaderData, inHeaderDataLength, dictRef)
}/* debug [functions.gen.go/function]: OBEXAddWhoHeader */

// Creates a formatted vCard, ready to be sent over OBEX or whatever.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Creates a formatted vCard, ready to be sent over OBEX or whatever.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXCreateVCard
func OBEXCreateVCard(inFirstName unsafe.Pointer, inFirstNameLength uint32, inLastName unsafe.Pointer, inLastNameLength uint32, inFriendlyName unsafe.Pointer, inFriendlyNameLength uint32, inNameCharset unsafe.Pointer, inNameCharsetLength uint32, inHomePhone unsafe.Pointer, inHomePhoneLength uint32, inWorkPhone unsafe.Pointer, inWorkPhoneLength uint32, inCellPhone unsafe.Pointer, inCellPhoneLength uint32, inFaxPhone unsafe.Pointer, inFaxPhoneLength uint32, inEMailAddress unsafe.Pointer, inEMailAddressLength uint32, inEMailAddressCharset unsafe.Pointer, inEMailAddressCharsetLength uint32, inOrganization unsafe.Pointer, inOrganizationLength uint32, inOrganizationCharset unsafe.Pointer, inOrganizationCharsetLength uint32, inTitle unsafe.Pointer, inTitleLength uint32, inTitleCharset unsafe.Pointer, inTitleCharsetLength uint32) DataRef {
	return _OBEXCreateVCard(inFirstName, inFirstNameLength, inLastName, inLastNameLength, inFriendlyName, inFriendlyNameLength, inNameCharset, inNameCharsetLength, inHomePhone, inHomePhoneLength, inWorkPhone, inWorkPhoneLength, inCellPhone, inCellPhoneLength, inFaxPhone, inFaxPhoneLength, inEMailAddress, inEMailAddressLength, inEMailAddressCharset, inEMailAddressCharsetLength, inOrganization, inOrganizationLength, inOrganizationCharset, inOrganizationCharsetLength, inTitle, inTitleLength, inTitleCharset, inTitleCharsetLength)
}/* debug [functions.gen.go/function]: OBEXCreateVCard */

// Creates a formatted vEvent, ready to be sent over OBEX or whatever. You probably will embed the output in a vCalendar event.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Creates a formatted vEvent, ready to be sent over OBEX or whatever. You probably will embed the output in a vCalendar event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXCreateVEvent
func OBEXCreateVEvent(inCharset unsafe.Pointer, inCharsetLength uint32, inEncoding unsafe.Pointer, inEncodingLength uint32, inEventStartDate unsafe.Pointer, inEventStartDateLength uint32, inEventEndDate unsafe.Pointer, inEventEndDateLength uint32, inAlarmDate unsafe.Pointer, inAlarmDateLength uint32, inCategory unsafe.Pointer, inCategoryLength uint32, inSummary unsafe.Pointer, inSummaryLength uint32, inLocation unsafe.Pointer, inLocationLength uint32, inXIRMCLUID unsafe.Pointer, inXIRMCLUIDLength uint32) DataRef {
	return _OBEXCreateVEvent(inCharset, inCharsetLength, inEncoding, inEncodingLength, inEventStartDate, inEventStartDateLength, inEventEndDate, inEventEndDateLength, inAlarmDate, inAlarmDateLength, inCategory, inCategoryLength, inSummary, inSummaryLength, inLocation, inLocationLength, inXIRMCLUID, inXIRMCLUIDLength)
}/* debug [functions.gen.go/function]: OBEXCreateVEvent */

// Take a data blob and looks for OBEX headers.
//
// Added in macOS .
// Take a data blob and looks for OBEX headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXGetHeaders(_:_:)
func OBEXGetHeaders(inData unsafe.Pointer, inDataSize uintptr) DictionaryRef {
	return _OBEXGetHeaders(inData, inDataSize)
}/* debug [functions.gen.go/function]: OBEXGetHeaders */

// Converts a dictionary of headers to a data pointer, from which you can extract as bytes and pass to the OBEX command/response functions.
//
// Added in macOS .
// Converts a dictionary of headers to a data pointer, from which you can extract as bytes and pass to the OBEX command/response functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXHeadersToBytes(_:)
func OBEXHeadersToBytes(dictionaryOfHeaders DictionaryRef) MutableDataRef {
	return _OBEXHeadersToBytes(dictionaryOfHeaders)
}/* debug [functions.gen.go/function]: OBEXHeadersToBytes */

// Send an abort command to a remote OBEX server.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send an abort command to a remote OBEX server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionAbort
func OBEXSessionAbort(inSessionRef OBEXSessionRef, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionAbort(inSessionRef, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionAbort */

// Send a response to a abort command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a abort command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionAbortResponse
func OBEXSessionAbortResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionAbortResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionAbortResponse */

// Establishes an OBEX connection to the target device for the session. If a transport connection is not open yet, it will be opened if possible.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Establishes an OBEX connection to the target device for the session. If a transport connection is not open yet, it will be opened if possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionConnect
func OBEXSessionConnect(inSessionRef OBEXSessionRef, inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionConnect(inSessionRef, inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionConnect */

// Send a response to a connect command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a connect command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionConnectResponse
func OBEXSessionConnectResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionConnectResponse(inSessionRef, inResponseOpCode, inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionConnectResponse */

// Destroy an OBEX session. If connections are open, they will (eventually) be terminated for you.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Destroy an OBEX session. If connections are open, they will (eventually) be terminated for you.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionDelete
func OBEXSessionDelete(inSessionRef OBEXSessionRef) OBEXError {
	return _OBEXSessionDelete(inSessionRef)
}/* debug [functions.gen.go/function]: OBEXSessionDelete */

// Send a disconnect command to a remote OBEX server.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a disconnect command to a remote OBEX server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionDisconnect
func OBEXSessionDisconnect(inSessionRef OBEXSessionRef, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionDisconnect(inSessionRef, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionDisconnect */

// Send a response to a disconnect command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a disconnect command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionDisconnectResponse
func OBEXSessionDisconnectResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionDisconnectResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionDisconnectResponse */

// Send a get command to a remote OBEX server.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a get command to a remote OBEX server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGet
func OBEXSessionGet(inSessionRef OBEXSessionRef, inIsFinalChunk unsafe.Pointer, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionGet(inSessionRef, inIsFinalChunk, inHeadersData, inHeadersDataLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionGet */

// Gets space available for your data for a particular command response you are trying to send.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Gets space available for your data for a particular command response you are trying to send.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetAvailableCommandPayloadLength
func OBEXSessionGetAvailableCommandPayloadLength(inSessionRef OBEXSessionRef, inOpCode OBEXOpCode, outLength unsafe.Pointer) OBEXError {
	return _OBEXSessionGetAvailableCommandPayloadLength(inSessionRef, inOpCode, outLength)
}/* debug [functions.gen.go/function]: OBEXSessionGetAvailableCommandPayloadLength */

// Gets space available for your data for a particular command response you are trying to send.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Gets space available for your data for a particular command response you are trying to send.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetAvailableCommandResponsePayloadLength
func OBEXSessionGetAvailableCommandResponsePayloadLength(inSessionRef OBEXSessionRef, inOpCode OBEXOpCode, outLength unsafe.Pointer) OBEXError {
	return _OBEXSessionGetAvailableCommandResponsePayloadLength(inSessionRef, inOpCode, outLength)
}/* debug [functions.gen.go/function]: OBEXSessionGetAvailableCommandResponsePayloadLength */

// Gets current max packet length.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Gets current max packet length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetMaxPacketLength
func OBEXSessionGetMaxPacketLength(inSessionRef OBEXSessionRef, outLength unsafe.Pointer) OBEXError {
	return _OBEXSessionGetMaxPacketLength(inSessionRef, outLength)
}/* debug [functions.gen.go/function]: OBEXSessionGetMaxPacketLength */

// Send a response to a get command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a get command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetResponse
func OBEXSessionGetResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionGetResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionGetResponse */

// Allows you to test the session for an open OBEX connection for a particular session.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Allows you to test the session for an open OBEX connection for a particular session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionHasOpenOBEXConnection
func OBEXSessionHasOpenOBEXConnection(inSessionRef OBEXSessionRef, outIsConnected unsafe.Pointer) OBEXError {
	return _OBEXSessionHasOpenOBEXConnection(inSessionRef, outIsConnected)
}/* debug [functions.gen.go/function]: OBEXSessionHasOpenOBEXConnection */

// Send a put command to a remote OBEX server.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a put command to a remote OBEX server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionPut
func OBEXSessionPut(inSessionRef OBEXSessionRef, inIsFinalChunk unsafe.Pointer, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inBodyData unsafe.Pointer, inBodyDataLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionPut(inSessionRef, inIsFinalChunk, inHeadersData, inHeadersDataLength, inBodyData, inBodyDataLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionPut */

// Send a response to a put command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a put command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionPutResponse
func OBEXSessionPutResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionPutResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionPutResponse */

// Send a set path command to a remote OBEX server.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a set path command to a remote OBEX server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionSetPath
func OBEXSessionSetPath(inSessionRef OBEXSessionRef, inFlags OBEXFlags, inConstants OBEXConstants, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionSetPath(inSessionRef, inFlags, inConstants, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionSetPath */

// Send a response to a set path command to the remote client.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Send a response to a set path command to the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionSetPathResponse
func OBEXSessionSetPathResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionSetPathResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionSetPathResponse */

// OBEXSessionSetServerCallback is a IOBluetooth function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionSetServerCallback
func OBEXSessionSetServerCallback(inSessionRef OBEXSessionRef, inCallback OBEXSessionEventCallback, inUserRefCon unsafe.Pointer) OBEXError {
	return _OBEXSessionSetServerCallback(inSessionRef, inCallback, inUserRefCon)
}/* debug [functions.gen.go/function]: OBEXSessionSetServerCallback */




