// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

/* debug [functions.gen.go]: Generating 95 functions for CoreMIDI */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CoreMIDI Functions (95 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MIDIBluetoothDriverActivateAllConnections func() unsafe.Pointer
	_MIDIBluetoothDriverDisconnect func(StringRef) unsafe.Pointer
	_MIDIClientCreate func(StringRef, MIDINotifyProc, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIClientCreateWithBlock func(StringRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIClientDispose func(MIDIClientRef) unsafe.Pointer
	_MIDIDestinationCreate func(MIDIClientRef, StringRef, MIDIReadProc, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDestinationCreateWithBlock func(MIDIClientRef, StringRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDestinationCreateWithProtocol func(MIDIClientRef, StringRef, MIDIProtocolID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceAddEntity func(MIDIDeviceRef, StringRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceCreate func(MIDIDriverRef, StringRef, StringRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceDispose func(MIDIDeviceRef) unsafe.Pointer
	_MIDIDeviceGetEntity func(MIDIDeviceRef, unsafe.Pointer) MIDIEntityRef
	_MIDIDeviceGetNumberOfEntities func(MIDIDeviceRef) unsafe.Pointer
	_MIDIDeviceListAddDevice func(MIDIDeviceListRef, MIDIDeviceRef) unsafe.Pointer
	_MIDIDeviceListDispose func(MIDIDeviceListRef) unsafe.Pointer
	_MIDIDeviceListGetDevice func(MIDIDeviceListRef, unsafe.Pointer) MIDIDeviceRef
	_MIDIDeviceListGetNumberOfDevices func(MIDIDeviceListRef) unsafe.Pointer
	_MIDIDeviceNewEntity func(MIDIDeviceRef, StringRef, MIDIProtocolID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceRemoveEntity func(MIDIDeviceRef, MIDIEntityRef) unsafe.Pointer
	_MIDIDriverEnableMonitoring func(MIDIDriverRef, unsafe.Pointer) unsafe.Pointer
	_MIDIEndpointDispose func(MIDIEndpointRef) unsafe.Pointer
	_MIDIEndpointGetEntity func(MIDIEndpointRef, unsafe.Pointer) unsafe.Pointer
	_MIDIEndpointGetRefCons func(MIDIEndpointRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEndpointSetRefCons func(MIDIEndpointRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEntityAddOrRemoveEndpoints func(MIDIEntityRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEntityGetDestination func(MIDIEntityRef, unsafe.Pointer) MIDIEndpointRef
	_MIDIEntityGetDevice func(MIDIEntityRef, unsafe.Pointer) unsafe.Pointer
	_MIDIEntityGetNumberOfDestinations func(MIDIEntityRef) unsafe.Pointer
	_MIDIEntityGetNumberOfSources func(MIDIEntityRef) unsafe.Pointer
	_MIDIEntityGetSource func(MIDIEntityRef, unsafe.Pointer) MIDIEndpointRef
	_MIDIEventListAdd func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, MIDITimeStamp, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEventListForEachEvent func(unsafe.Pointer, MIDIEventVisitor, unsafe.Pointer)
	_MIDIEventListInit func(unsafe.Pointer, MIDIProtocolID) unsafe.Pointer
	_MIDIEventPacketSysexBytesForGroup func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIExternalDeviceCreate func(StringRef, StringRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDIFlushOutput func(MIDIEndpointRef) unsafe.Pointer
	_MIDIGetDestination func(unsafe.Pointer) MIDIEndpointRef
	_MIDIGetDevice func(unsafe.Pointer) MIDIDeviceRef
	_MIDIGetDriverDeviceList func(MIDIDriverRef) MIDIDeviceListRef
	_MIDIGetDriverIORunLoop func() RunLoopRef
	_MIDIGetExternalDevice func(unsafe.Pointer) MIDIDeviceRef
	_MIDIGetNumberOfDestinations func() unsafe.Pointer
	_MIDIGetNumberOfDevices func() unsafe.Pointer
	_MIDIGetNumberOfExternalDevices func() unsafe.Pointer
	_MIDIGetNumberOfSources func() unsafe.Pointer
	_MIDIGetSerialPortDrivers func(unsafe.Pointer) unsafe.Pointer
	_MIDIGetSerialPortOwner func(StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDIGetSource func(unsafe.Pointer) MIDIEndpointRef
	_MIDIInputPortCreate func(MIDIClientRef, StringRef, MIDIReadProc, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIInputPortCreateWithBlock func(MIDIClientRef, StringRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIInputPortCreateWithProtocol func(MIDIClientRef, StringRef, MIDIProtocolID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIObjectFindByUniqueID func(MIDIUniqueID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIObjectGetDataProperty func(MIDIObjectRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDIObjectGetDictionaryProperty func(MIDIObjectRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDIObjectGetIntegerProperty func(MIDIObjectRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDIObjectGetProperties func(MIDIObjectRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIObjectGetStringProperty func(MIDIObjectRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDIObjectRemoveProperty func(MIDIObjectRef, StringRef) unsafe.Pointer
	_MIDIObjectSetDataProperty func(MIDIObjectRef, StringRef, DataRef) unsafe.Pointer
	_MIDIObjectSetDictionaryProperty func(MIDIObjectRef, StringRef, DictionaryRef) unsafe.Pointer
	_MIDIObjectSetIntegerProperty func(MIDIObjectRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDIObjectSetStringProperty func(MIDIObjectRef, StringRef, StringRef) unsafe.Pointer
	_MIDIOutputPortCreate func(MIDIClientRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDIPacketListAdd func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, MIDITimeStamp, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIPacketListInit func(unsafe.Pointer) unsafe.Pointer
	_MIDIPortConnectSource func(MIDIPortRef, MIDIEndpointRef, unsafe.Pointer) unsafe.Pointer
	_MIDIPortDisconnectSource func(MIDIPortRef, MIDIEndpointRef) unsafe.Pointer
	_MIDIPortDispose func(MIDIPortRef) unsafe.Pointer
	_MIDIReceived func(MIDIEndpointRef, unsafe.Pointer) unsafe.Pointer
	_MIDIReceivedEventList func(MIDIEndpointRef, unsafe.Pointer) unsafe.Pointer
	_MIDIRestart func() unsafe.Pointer
	_MIDISend func(MIDIPortRef, MIDIEndpointRef, unsafe.Pointer) unsafe.Pointer
	_MIDISendEventList func(MIDIPortRef, MIDIEndpointRef, unsafe.Pointer) unsafe.Pointer
	_MIDISendSysex func(unsafe.Pointer) unsafe.Pointer
	_MIDISendUMPSysex func(unsafe.Pointer) unsafe.Pointer
	_MIDISendUMPSysex8 func(unsafe.Pointer) unsafe.Pointer
	_MIDISetSerialPortOwner func(StringRef, StringRef) unsafe.Pointer
	_MIDISetupAddDevice func(MIDIDeviceRef) unsafe.Pointer
	_MIDISetupAddExternalDevice func(MIDIDeviceRef) unsafe.Pointer
	_MIDISetupCreate func(unsafe.Pointer) unsafe.Pointer
	_MIDISetupDispose func(MIDISetupRef) unsafe.Pointer
	_MIDISetupFromData func(DataRef, unsafe.Pointer) unsafe.Pointer
	_MIDISetupGetCurrent func(unsafe.Pointer) unsafe.Pointer
	_MIDISetupInstall func(MIDISetupRef) unsafe.Pointer
	_MIDISetupRemoveDevice func(MIDIDeviceRef) unsafe.Pointer
	_MIDISetupRemoveExternalDevice func(MIDIDeviceRef) unsafe.Pointer
	_MIDISetupToData func(MIDISetupRef, unsafe.Pointer) unsafe.Pointer
	_MIDISourceCreate func(MIDIClientRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDISourceCreateWithProtocol func(MIDIClientRef, StringRef, MIDIProtocolID, unsafe.Pointer) unsafe.Pointer
	_MIDIThruConnectionCreate func(StringRef, DataRef, unsafe.Pointer) unsafe.Pointer
	_MIDIThruConnectionDispose func(MIDIThruConnectionRef) unsafe.Pointer
	_MIDIThruConnectionFind func(StringRef, unsafe.Pointer) unsafe.Pointer
	_MIDIThruConnectionGetParams func(MIDIThruConnectionRef, unsafe.Pointer) unsafe.Pointer
	_MIDIThruConnectionParamsInitialize func(unsafe.Pointer)
	_MIDIThruConnectionSetParams func(MIDIThruConnectionRef, DataRef) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MIDIBluetoothDriverActivateAllConnections, lib, "MIDIBluetoothDriverActivateAllConnections")
	tryRegister(&_MIDIBluetoothDriverDisconnect, lib, "MIDIBluetoothDriverDisconnect")
	tryRegister(&_MIDIClientCreate, lib, "MIDIClientCreate")
	tryRegister(&_MIDIClientCreateWithBlock, lib, "MIDIClientCreateWithBlock")
	tryRegister(&_MIDIClientDispose, lib, "MIDIClientDispose")
	tryRegister(&_MIDIDestinationCreate, lib, "MIDIDestinationCreate")
	tryRegister(&_MIDIDestinationCreateWithBlock, lib, "MIDIDestinationCreateWithBlock")
	tryRegister(&_MIDIDestinationCreateWithProtocol, lib, "MIDIDestinationCreateWithProtocol")
	tryRegister(&_MIDIDeviceAddEntity, lib, "MIDIDeviceAddEntity")
	tryRegister(&_MIDIDeviceCreate, lib, "MIDIDeviceCreate")
	tryRegister(&_MIDIDeviceDispose, lib, "MIDIDeviceDispose")
	tryRegister(&_MIDIDeviceGetEntity, lib, "MIDIDeviceGetEntity")
	tryRegister(&_MIDIDeviceGetNumberOfEntities, lib, "MIDIDeviceGetNumberOfEntities")
	tryRegister(&_MIDIDeviceListAddDevice, lib, "MIDIDeviceListAddDevice")
	tryRegister(&_MIDIDeviceListDispose, lib, "MIDIDeviceListDispose")
	tryRegister(&_MIDIDeviceListGetDevice, lib, "MIDIDeviceListGetDevice")
	tryRegister(&_MIDIDeviceListGetNumberOfDevices, lib, "MIDIDeviceListGetNumberOfDevices")
	tryRegister(&_MIDIDeviceNewEntity, lib, "MIDIDeviceNewEntity")
	tryRegister(&_MIDIDeviceRemoveEntity, lib, "MIDIDeviceRemoveEntity")
	tryRegister(&_MIDIDriverEnableMonitoring, lib, "MIDIDriverEnableMonitoring")
	tryRegister(&_MIDIEndpointDispose, lib, "MIDIEndpointDispose")
	tryRegister(&_MIDIEndpointGetEntity, lib, "MIDIEndpointGetEntity")
	tryRegister(&_MIDIEndpointGetRefCons, lib, "MIDIEndpointGetRefCons")
	tryRegister(&_MIDIEndpointSetRefCons, lib, "MIDIEndpointSetRefCons")
	tryRegister(&_MIDIEntityAddOrRemoveEndpoints, lib, "MIDIEntityAddOrRemoveEndpoints")
	tryRegister(&_MIDIEntityGetDestination, lib, "MIDIEntityGetDestination")
	tryRegister(&_MIDIEntityGetDevice, lib, "MIDIEntityGetDevice")
	tryRegister(&_MIDIEntityGetNumberOfDestinations, lib, "MIDIEntityGetNumberOfDestinations")
	tryRegister(&_MIDIEntityGetNumberOfSources, lib, "MIDIEntityGetNumberOfSources")
	tryRegister(&_MIDIEntityGetSource, lib, "MIDIEntityGetSource")
	tryRegister(&_MIDIEventListAdd, lib, "MIDIEventListAdd")
	tryRegister(&_MIDIEventListForEachEvent, lib, "MIDIEventListForEachEvent")
	tryRegister(&_MIDIEventListInit, lib, "MIDIEventListInit")
	tryRegister(&_MIDIEventPacketSysexBytesForGroup, lib, "MIDIEventPacketSysexBytesForGroup")
	tryRegister(&_MIDIExternalDeviceCreate, lib, "MIDIExternalDeviceCreate")
	tryRegister(&_MIDIFlushOutput, lib, "MIDIFlushOutput")
	tryRegister(&_MIDIGetDestination, lib, "MIDIGetDestination")
	tryRegister(&_MIDIGetDevice, lib, "MIDIGetDevice")
	tryRegister(&_MIDIGetDriverDeviceList, lib, "MIDIGetDriverDeviceList")
	tryRegister(&_MIDIGetDriverIORunLoop, lib, "MIDIGetDriverIORunLoop")
	tryRegister(&_MIDIGetExternalDevice, lib, "MIDIGetExternalDevice")
	tryRegister(&_MIDIGetNumberOfDestinations, lib, "MIDIGetNumberOfDestinations")
	tryRegister(&_MIDIGetNumberOfDevices, lib, "MIDIGetNumberOfDevices")
	tryRegister(&_MIDIGetNumberOfExternalDevices, lib, "MIDIGetNumberOfExternalDevices")
	tryRegister(&_MIDIGetNumberOfSources, lib, "MIDIGetNumberOfSources")
	tryRegister(&_MIDIGetSerialPortDrivers, lib, "MIDIGetSerialPortDrivers")
	tryRegister(&_MIDIGetSerialPortOwner, lib, "MIDIGetSerialPortOwner")
	tryRegister(&_MIDIGetSource, lib, "MIDIGetSource")
	tryRegister(&_MIDIInputPortCreate, lib, "MIDIInputPortCreate")
	tryRegister(&_MIDIInputPortCreateWithBlock, lib, "MIDIInputPortCreateWithBlock")
	tryRegister(&_MIDIInputPortCreateWithProtocol, lib, "MIDIInputPortCreateWithProtocol")
	tryRegister(&_MIDIObjectFindByUniqueID, lib, "MIDIObjectFindByUniqueID")
	tryRegister(&_MIDIObjectGetDataProperty, lib, "MIDIObjectGetDataProperty")
	tryRegister(&_MIDIObjectGetDictionaryProperty, lib, "MIDIObjectGetDictionaryProperty")
	tryRegister(&_MIDIObjectGetIntegerProperty, lib, "MIDIObjectGetIntegerProperty")
	tryRegister(&_MIDIObjectGetProperties, lib, "MIDIObjectGetProperties")
	tryRegister(&_MIDIObjectGetStringProperty, lib, "MIDIObjectGetStringProperty")
	tryRegister(&_MIDIObjectRemoveProperty, lib, "MIDIObjectRemoveProperty")
	tryRegister(&_MIDIObjectSetDataProperty, lib, "MIDIObjectSetDataProperty")
	tryRegister(&_MIDIObjectSetDictionaryProperty, lib, "MIDIObjectSetDictionaryProperty")
	tryRegister(&_MIDIObjectSetIntegerProperty, lib, "MIDIObjectSetIntegerProperty")
	tryRegister(&_MIDIObjectSetStringProperty, lib, "MIDIObjectSetStringProperty")
	tryRegister(&_MIDIOutputPortCreate, lib, "MIDIOutputPortCreate")
	tryRegister(&_MIDIPacketListAdd, lib, "MIDIPacketListAdd")
	tryRegister(&_MIDIPacketListInit, lib, "MIDIPacketListInit")
	tryRegister(&_MIDIPortConnectSource, lib, "MIDIPortConnectSource")
	tryRegister(&_MIDIPortDisconnectSource, lib, "MIDIPortDisconnectSource")
	tryRegister(&_MIDIPortDispose, lib, "MIDIPortDispose")
	tryRegister(&_MIDIReceived, lib, "MIDIReceived")
	tryRegister(&_MIDIReceivedEventList, lib, "MIDIReceivedEventList")
	tryRegister(&_MIDIRestart, lib, "MIDIRestart")
	tryRegister(&_MIDISend, lib, "MIDISend")
	tryRegister(&_MIDISendEventList, lib, "MIDISendEventList")
	tryRegister(&_MIDISendSysex, lib, "MIDISendSysex")
	tryRegister(&_MIDISendUMPSysex, lib, "MIDISendUMPSysex")
	tryRegister(&_MIDISendUMPSysex8, lib, "MIDISendUMPSysex8")
	tryRegister(&_MIDISetSerialPortOwner, lib, "MIDISetSerialPortOwner")
	tryRegister(&_MIDISetupAddDevice, lib, "MIDISetupAddDevice")
	tryRegister(&_MIDISetupAddExternalDevice, lib, "MIDISetupAddExternalDevice")
	tryRegister(&_MIDISetupCreate, lib, "MIDISetupCreate")
	tryRegister(&_MIDISetupDispose, lib, "MIDISetupDispose")
	tryRegister(&_MIDISetupFromData, lib, "MIDISetupFromData")
	tryRegister(&_MIDISetupGetCurrent, lib, "MIDISetupGetCurrent")
	tryRegister(&_MIDISetupInstall, lib, "MIDISetupInstall")
	tryRegister(&_MIDISetupRemoveDevice, lib, "MIDISetupRemoveDevice")
	tryRegister(&_MIDISetupRemoveExternalDevice, lib, "MIDISetupRemoveExternalDevice")
	tryRegister(&_MIDISetupToData, lib, "MIDISetupToData")
	tryRegister(&_MIDISourceCreate, lib, "MIDISourceCreate")
	tryRegister(&_MIDISourceCreateWithProtocol, lib, "MIDISourceCreateWithProtocol")
	tryRegister(&_MIDIThruConnectionCreate, lib, "MIDIThruConnectionCreate")
	tryRegister(&_MIDIThruConnectionDispose, lib, "MIDIThruConnectionDispose")
	tryRegister(&_MIDIThruConnectionFind, lib, "MIDIThruConnectionFind")
	tryRegister(&_MIDIThruConnectionGetParams, lib, "MIDIThruConnectionGetParams")
	tryRegister(&_MIDIThruConnectionParamsInitialize, lib, "MIDIThruConnectionParamsInitialize")
	tryRegister(&_MIDIThruConnectionSetParams, lib, "MIDIThruConnectionSetParams")
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



// Promote all active Bluetooth connections into an online MIDI device capable of input and output.
//
// Added in macOS 13.0.
// Promote all active Bluetooth connections into an online MIDI device capable of input and output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIBluetoothDriverActivateAllConnections()
func MIDIBluetoothDriverActivateAllConnections() unsafe.Pointer {
	return _MIDIBluetoothDriverActivateAllConnections()
}/* debug [functions.gen.go/function]: MIDIBluetoothDriverActivateAllConnections */

// Disconnect the Bluetooth MIDI driver from a Bluetooth Low Energy MIDI peripheral.
//
// Added in macOS 13.0.
// Disconnect the Bluetooth MIDI driver from a Bluetooth Low Energy MIDI peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIBluetoothDriverDisconnect(_:)
func MIDIBluetoothDriverDisconnect(uuid StringRef) unsafe.Pointer {
	return _MIDIBluetoothDriverDisconnect(uuid)
}/* debug [functions.gen.go/function]: MIDIBluetoothDriverDisconnect */

// Creates a MIDI client.
//
// Added in macOS 10.0.
// Creates a MIDI client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIClientCreate(_:_:_:_:)
func MIDIClientCreate(name StringRef, notifyProc MIDINotifyProc, notifyRefCon unsafe.Pointer, outClient unsafe.Pointer) unsafe.Pointer {
	return _MIDIClientCreate(name, notifyProc, notifyRefCon, outClient)
}/* debug [functions.gen.go/function]: MIDIClientCreate */

// Creates a MIDI client with a callback block.
//
// Added in macOS 10.11.
// Creates a MIDI client with a callback block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIClientCreateWithBlock(_:_:_:)
func MIDIClientCreateWithBlock(name StringRef, outClient unsafe.Pointer, notifyBlock unsafe.Pointer) unsafe.Pointer {
	return _MIDIClientCreateWithBlock(name, outClient, notifyBlock)
}/* debug [functions.gen.go/function]: MIDIClientCreateWithBlock */

// Disposes of a MIDI client.
//
// Added in macOS 10.0.
// Disposes of a MIDI client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIClientDispose(_:)
func MIDIClientDispose(client MIDIClientRef) unsafe.Pointer {
	return _MIDIClientDispose(client)
}/* debug [functions.gen.go/function]: MIDIClientDispose */

// Creates a virtual destination in a client.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
// Creates a virtual destination in a client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDestinationCreate(_:_:_:_:_:)
func MIDIDestinationCreate(client MIDIClientRef, name StringRef, readProc MIDIReadProc, refCon unsafe.Pointer, outDest unsafe.Pointer) unsafe.Pointer {
	return _MIDIDestinationCreate(client, name, readProc, refCon, outDest)
}/* debug [functions.gen.go/function]: MIDIDestinationCreate */

// Creates a virtual destination in a client.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.11.
// Creates a virtual destination in a client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDestinationCreateWithBlock(_:_:_:_:)
func MIDIDestinationCreateWithBlock(client MIDIClientRef, name StringRef, outDest unsafe.Pointer, readBlock unsafe.Pointer) unsafe.Pointer {
	return _MIDIDestinationCreateWithBlock(client, name, outDest, readBlock)
}/* debug [functions.gen.go/function]: MIDIDestinationCreateWithBlock */

// Creates a virtual destination in a client.
//
// Added in macOS 11.0.
// Creates a virtual destination in a client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDestinationCreateWithProtocol(_:_:_:_:_:)
func MIDIDestinationCreateWithProtocol(client MIDIClientRef, name StringRef, protocol_ MIDIProtocolID, outDest unsafe.Pointer, readBlock unsafe.Pointer) unsafe.Pointer {
	return _MIDIDestinationCreateWithProtocol(client, name, protocol_, outDest, readBlock)
}/* debug [functions.gen.go/function]: MIDIDestinationCreateWithProtocol */

// Specifies one of the entities that make up a device.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
// Specifies one of the entities that make up a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceAddEntity(_:_:_:_:_:_:)
func MIDIDeviceAddEntity(device MIDIDeviceRef, name StringRef, embedded unsafe.Pointer, numSourceEndpoints unsafe.Pointer, numDestinationEndpoints unsafe.Pointer, newEntity unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceAddEntity(device, name, embedded, numSourceEndpoints, numDestinationEndpoints, newEntity)
}/* debug [functions.gen.go/function]: MIDIDeviceAddEntity */

// Creates a new device object that corresponds to the available hardware.
//
// Added in macOS 10.0.
// Creates a new device object that corresponds to the available hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceCreate(_:_:_:_:_:)
func MIDIDeviceCreate(owner MIDIDriverRef, name StringRef, manufacturer StringRef, model StringRef, outDevice unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceCreate(owner, name, manufacturer, model, outDevice)
}/* debug [functions.gen.go/function]: MIDIDeviceCreate */

// Disposes of a MIDI device.
//
// Added in macOS 10.3.
// Disposes of a MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceDispose(_:)
func MIDIDeviceDispose(device MIDIDeviceRef) unsafe.Pointer {
	return _MIDIDeviceDispose(device)
}/* debug [functions.gen.go/function]: MIDIDeviceDispose */

// Returns the device’s entity at a specific index.
//
// Added in macOS 10.0.
// Returns the device’s entity at a specific index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceGetEntity(_:_:)
func MIDIDeviceGetEntity(device MIDIDeviceRef, entityIndex0 unsafe.Pointer) MIDIEntityRef {
	return _MIDIDeviceGetEntity(device, entityIndex0)
}/* debug [functions.gen.go/function]: MIDIDeviceGetEntity */

// Returns the number of entities in a device.
//
// Added in macOS 10.0.
// Returns the number of entities in a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceGetNumberOfEntities(_:)
func MIDIDeviceGetNumberOfEntities(device MIDIDeviceRef) unsafe.Pointer {
	return _MIDIDeviceGetNumberOfEntities(device)
}/* debug [functions.gen.go/function]: MIDIDeviceGetNumberOfEntities */

// Adds the specified device to the device list.
//
// Added in macOS 10.0.
// Adds the specified device to the device list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListAddDevice(_:_:)
func MIDIDeviceListAddDevice(devList MIDIDeviceListRef, dev MIDIDeviceRef) unsafe.Pointer {
	return _MIDIDeviceListAddDevice(devList, dev)
}/* debug [functions.gen.go/function]: MIDIDeviceListAddDevice */

// Disposes of a device list, but not its devices.
//
// Added in macOS 10.1.
// Disposes of a device list, but not its devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListDispose(_:)
func MIDIDeviceListDispose(devList MIDIDeviceListRef) unsafe.Pointer {
	return _MIDIDeviceListDispose(devList)
}/* debug [functions.gen.go/function]: MIDIDeviceListDispose */

// Retrieves a MIDI device from a device list.
//
// Added in macOS 10.0.
// Retrieves a MIDI device from a device list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListGetDevice(_:_:)
func MIDIDeviceListGetDevice(devList MIDIDeviceListRef, index0 unsafe.Pointer) MIDIDeviceRef {
	return _MIDIDeviceListGetDevice(devList, index0)
}/* debug [functions.gen.go/function]: MIDIDeviceListGetDevice */

// Retrieves the number of devices in a device list.
//
// Added in macOS 10.0.
// Retrieves the number of devices in a device list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListGetNumberOfDevices(_:)
func MIDIDeviceListGetNumberOfDevices(devList MIDIDeviceListRef) unsafe.Pointer {
	return _MIDIDeviceListGetNumberOfDevices(devList)
}/* debug [functions.gen.go/function]: MIDIDeviceListGetNumberOfDevices */

// Adds a new entity to a device.
//
// Added in macOS 11.0.
// Adds a new entity to a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceNewEntity(_:_:_:_:_:_:_:)
func MIDIDeviceNewEntity(device MIDIDeviceRef, name StringRef, protocol_ MIDIProtocolID, embedded unsafe.Pointer, numSourceEndpoints unsafe.Pointer, numDestinationEndpoints unsafe.Pointer, newEntity unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceNewEntity(device, name, protocol_, embedded, numSourceEndpoints, numDestinationEndpoints, newEntity)
}/* debug [functions.gen.go/function]: MIDIDeviceNewEntity */

// Removes an entity from a device.
//
// Added in macOS 10.1.
// Removes an entity from a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceRemoveEntity(_:_:)
func MIDIDeviceRemoveEntity(device MIDIDeviceRef, entity MIDIEntityRef) unsafe.Pointer {
	return _MIDIDeviceRemoveEntity(device, entity)
}/* debug [functions.gen.go/function]: MIDIDeviceRemoveEntity */

// Enables monitoring of all outgoing MIDI packets.
//
// Added in macOS 10.1.
// Enables monitoring of all outgoing MIDI packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDriverEnableMonitoring(_:_:)
func MIDIDriverEnableMonitoring(driver MIDIDriverRef, enabled unsafe.Pointer) unsafe.Pointer {
	return _MIDIDriverEnableMonitoring(driver, enabled)
}/* debug [functions.gen.go/function]: MIDIDriverEnableMonitoring */

// Disposes of a virtual source or destination.
//
// Added in macOS 10.0.
// Disposes of a virtual source or destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointDispose(_:)
func MIDIEndpointDispose(endpt MIDIEndpointRef) unsafe.Pointer {
	return _MIDIEndpointDispose(endpt)
}/* debug [functions.gen.go/function]: MIDIEndpointDispose */

// Returns an endpoint’s entity.
//
// Added in macOS 10.2.
// Returns an endpoint’s entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointGetEntity(_:_:)
func MIDIEndpointGetEntity(inEndpoint MIDIEndpointRef, outEntity unsafe.Pointer) unsafe.Pointer {
	return _MIDIEndpointGetEntity(inEndpoint, outEntity)
}/* debug [functions.gen.go/function]: MIDIEndpointGetEntity */

// Returns contextual data assigned to an endpoint.
//
// Added in macOS 10.0.
// Returns contextual data assigned to an endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointGetRefCons(_:_:_:)
func MIDIEndpointGetRefCons(endpt MIDIEndpointRef, ref1 unsafe.Pointer, ref2 unsafe.Pointer) unsafe.Pointer {
	return _MIDIEndpointGetRefCons(endpt, ref1, ref2)
}/* debug [functions.gen.go/function]: MIDIEndpointGetRefCons */

// Sets contextual data on an endpoint.
//
// Added in macOS 10.0.
// Sets contextual data on an endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointSetRefCons(_:_:_:)
func MIDIEndpointSetRefCons(endpt MIDIEndpointRef, ref1 unsafe.Pointer, ref2 unsafe.Pointer) unsafe.Pointer {
	return _MIDIEndpointSetRefCons(endpt, ref1, ref2)
}/* debug [functions.gen.go/function]: MIDIEndpointSetRefCons */

// Adds or removes an entity’s endpoints.
//
// Added in macOS 10.2.
// Adds or removes an entity’s endpoints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityAddOrRemoveEndpoints(_:_:_:)
func MIDIEntityAddOrRemoveEndpoints(entity MIDIEntityRef, numSourceEndpoints unsafe.Pointer, numDestinationEndpoints unsafe.Pointer) unsafe.Pointer {
	return _MIDIEntityAddOrRemoveEndpoints(entity, numSourceEndpoints, numDestinationEndpoints)
}/* debug [functions.gen.go/function]: MIDIEntityAddOrRemoveEndpoints */

// Returns one of an entity’s destinations.
//
// Added in macOS 10.0.
// Returns one of an entity’s destinations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetDestination(_:_:)
func MIDIEntityGetDestination(entity MIDIEntityRef, destIndex0 unsafe.Pointer) MIDIEndpointRef {
	return _MIDIEntityGetDestination(entity, destIndex0)
}/* debug [functions.gen.go/function]: MIDIEntityGetDestination */

// Returns an entity’s device.
//
// Added in macOS 10.2.
// Returns an entity’s device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetDevice(_:_:)
func MIDIEntityGetDevice(inEntity MIDIEntityRef, outDevice unsafe.Pointer) unsafe.Pointer {
	return _MIDIEntityGetDevice(inEntity, outDevice)
}/* debug [functions.gen.go/function]: MIDIEntityGetDevice */

// Returns the number of destinations in an entity.
//
// Added in macOS 10.0.
// Returns the number of destinations in an entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetNumberOfDestinations(_:)
func MIDIEntityGetNumberOfDestinations(entity MIDIEntityRef) unsafe.Pointer {
	return _MIDIEntityGetNumberOfDestinations(entity)
}/* debug [functions.gen.go/function]: MIDIEntityGetNumberOfDestinations */

// Returns the number of sources in an entity.
//
// Added in macOS 10.0.
// Returns the number of sources in an entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetNumberOfSources(_:)
func MIDIEntityGetNumberOfSources(entity MIDIEntityRef) unsafe.Pointer {
	return _MIDIEntityGetNumberOfSources(entity)
}/* debug [functions.gen.go/function]: MIDIEntityGetNumberOfSources */

// Returns one of an entity’s sources.
//
// Added in macOS 10.0.
// Returns one of an entity’s sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetSource(_:_:)
func MIDIEntityGetSource(entity MIDIEntityRef, sourceIndex0 unsafe.Pointer) MIDIEndpointRef {
	return _MIDIEntityGetSource(entity, sourceIndex0)
}/* debug [functions.gen.go/function]: MIDIEntityGetSource */

// Adds an event to an event list.
//
// Added in macOS 11.0.
// Adds an event to an event list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventListAdd(_:_:_:_:_:_:)
func MIDIEventListAdd(evtlist unsafe.Pointer, listSize unsafe.Pointer, curPacket unsafe.Pointer, time MIDITimeStamp, wordCount unsafe.Pointer, words unsafe.Pointer) unsafe.Pointer {
	return _MIDIEventListAdd(evtlist, listSize, curPacket, time, wordCount, words)
}/* debug [functions.gen.go/function]: MIDIEventListAdd */

// MIDIEventListForEachEvent is a CoreMIDI function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventListForEachEvent(_:_:_:)
func MIDIEventListForEachEvent(evtlist unsafe.Pointer, visitor MIDIEventVisitor, visitorContext unsafe.Pointer) {
	_MIDIEventListForEachEvent(evtlist, visitor, visitorContext)
}/* debug [functions.gen.go/function]: MIDIEventListForEachEvent */

// Initializes an event list.
//
// Added in macOS 11.0.
// Initializes an event list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventListInit(_:_:)
func MIDIEventListInit(evtlist unsafe.Pointer, protocol_ MIDIProtocolID) unsafe.Pointer {
	return _MIDIEventListInit(evtlist, protocol_)
}/* debug [functions.gen.go/function]: MIDIEventListInit */

// Gets MIDI 1.0 system-exclusive (SysEx) bytes on the indicated group.
//
// Added in macOS 14.0.
// Gets MIDI 1.0 system-exclusive (SysEx) bytes on the indicated group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventPacketSysexBytesForGroup(_:_:_:)
func MIDIEventPacketSysexBytesForGroup(pkt unsafe.Pointer, groupIndex unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _MIDIEventPacketSysexBytesForGroup(pkt, groupIndex, outData)
}/* debug [functions.gen.go/function]: MIDIEventPacketSysexBytesForGroup */

// Creates an external MIDI device.
//
// Added in macOS 10.1.
// Creates an external MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIExternalDeviceCreate(_:_:_:_:)
func MIDIExternalDeviceCreate(name StringRef, manufacturer StringRef, model StringRef, outDevice unsafe.Pointer) unsafe.Pointer {
	return _MIDIExternalDeviceCreate(name, manufacturer, model, outDevice)
}/* debug [functions.gen.go/function]: MIDIExternalDeviceCreate */

// Cancels all pending events that were previously scheduled to send.
//
// Added in macOS 10.1.
// Cancels all pending events that were previously scheduled to send.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIFlushOutput(_:)
func MIDIFlushOutput(dest MIDIEndpointRef) unsafe.Pointer {
	return _MIDIFlushOutput(dest)
}/* debug [functions.gen.go/function]: MIDIFlushOutput */

// Returns a destination in the system.
//
// Added in macOS 10.0.
// Returns a destination in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDestination(_:)
func MIDIGetDestination(destIndex0 unsafe.Pointer) MIDIEndpointRef {
	return _MIDIGetDestination(destIndex0)
}/* debug [functions.gen.go/function]: MIDIGetDestination */

// Returns a device from the system.
//
// Added in macOS 10.0.
// Returns a device from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDevice(_:)
func MIDIGetDevice(deviceIndex0 unsafe.Pointer) MIDIDeviceRef {
	return _MIDIGetDevice(deviceIndex0)
}/* debug [functions.gen.go/function]: MIDIGetDevice */

// Returns the list of driver-created devices in the current MIDI setup.
//
// Added in macOS 10.1.
// Returns the list of driver-created devices in the current MIDI setup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDriverDeviceList(_:)
func MIDIGetDriverDeviceList(driver MIDIDriverRef) MIDIDeviceListRef {
	return _MIDIGetDriverDeviceList(driver)
}/* debug [functions.gen.go/function]: MIDIGetDriverDeviceList */

// Returns the server’s driver I/O thread.
//
// Added in macOS 10.0.
// Returns the server’s driver I/O thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDriverIORunLoop()
func MIDIGetDriverIORunLoop() RunLoopRef {
	return _MIDIGetDriverIORunLoop()
}/* debug [functions.gen.go/function]: MIDIGetDriverIORunLoop */

// Returns one of the external devices in the system.
//
// Added in macOS 10.1.
// Returns one of the external devices in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetExternalDevice(_:)
func MIDIGetExternalDevice(deviceIndex0 unsafe.Pointer) MIDIDeviceRef {
	return _MIDIGetExternalDevice(deviceIndex0)
}/* debug [functions.gen.go/function]: MIDIGetExternalDevice */

// Returns the number of destinations in the system.
//
// Added in macOS 10.0.
// Returns the number of destinations in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfDestinations()
func MIDIGetNumberOfDestinations() unsafe.Pointer {
	return _MIDIGetNumberOfDestinations()
}/* debug [functions.gen.go/function]: MIDIGetNumberOfDestinations */

// Returns the number of devices in the system.
//
// Added in macOS 10.0.
// Returns the number of devices in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfDevices()
func MIDIGetNumberOfDevices() unsafe.Pointer {
	return _MIDIGetNumberOfDevices()
}/* debug [functions.gen.go/function]: MIDIGetNumberOfDevices */

// Returns the number of external MIDI devices in the system.
//
// Added in macOS 10.1.
// Returns the number of external MIDI devices in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfExternalDevices()
func MIDIGetNumberOfExternalDevices() unsafe.Pointer {
	return _MIDIGetNumberOfExternalDevices()
}/* debug [functions.gen.go/function]: MIDIGetNumberOfExternalDevices */

// Returns the number of sources in the system.
//
// Added in macOS 10.0.
// Returns the number of sources in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfSources()
func MIDIGetNumberOfSources() unsafe.Pointer {
	return _MIDIGetNumberOfSources()
}/* debug [functions.gen.go/function]: MIDIGetNumberOfSources */

// Returns a list of installed MIDI drivers for serial port MIDI devices.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
// Returns a list of installed MIDI drivers for serial port MIDI devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetSerialPortDrivers
func MIDIGetSerialPortDrivers(outDriverNames unsafe.Pointer) unsafe.Pointer {
	return _MIDIGetSerialPortDrivers(outDriverNames)
}/* debug [functions.gen.go/function]: MIDIGetSerialPortDrivers */

// Returns the MIDI driver that owns a serial port.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
// Returns the MIDI driver that owns a serial port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetSerialPortOwner
func MIDIGetSerialPortOwner(portName StringRef, outDriverName unsafe.Pointer) unsafe.Pointer {
	return _MIDIGetSerialPortOwner(portName, outDriverName)
}/* debug [functions.gen.go/function]: MIDIGetSerialPortOwner */

// Returns a source in the system.
//
// Added in macOS 10.0.
// Returns a source in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetSource(_:)
func MIDIGetSource(sourceIndex0 unsafe.Pointer) MIDIEndpointRef {
	return _MIDIGetSource(sourceIndex0)
}/* debug [functions.gen.go/function]: MIDIGetSource */

// Creates an input port through which the client may receive incoming MIDI messages from any MIDI source.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
// Creates an input port through which the client may receive incoming MIDI messages from any MIDI source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIInputPortCreate(_:_:_:_:_:)
func MIDIInputPortCreate(client MIDIClientRef, portName StringRef, readProc MIDIReadProc, refCon unsafe.Pointer, outPort unsafe.Pointer) unsafe.Pointer {
	return _MIDIInputPortCreate(client, portName, readProc, refCon, outPort)
}/* debug [functions.gen.go/function]: MIDIInputPortCreate */

// Creates an input port through which the client may receive incoming MIDI messages from any MIDI source.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.11.
// Creates an input port through which the client may receive incoming MIDI messages from any MIDI source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIInputPortCreateWithBlock(_:_:_:_:)
func MIDIInputPortCreateWithBlock(client MIDIClientRef, portName StringRef, outPort unsafe.Pointer, readBlock unsafe.Pointer) unsafe.Pointer {
	return _MIDIInputPortCreateWithBlock(client, portName, outPort, readBlock)
}/* debug [functions.gen.go/function]: MIDIInputPortCreateWithBlock */

// Creates an input port through which the client may receive incoming MIDI messages from any MIDI source.
//
// Added in macOS 11.0.
// Creates an input port through which the client may receive incoming MIDI messages from any MIDI source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIInputPortCreateWithProtocol(_:_:_:_:_:)
func MIDIInputPortCreateWithProtocol(client MIDIClientRef, portName StringRef, protocol_ MIDIProtocolID, outPort unsafe.Pointer, receiveBlock unsafe.Pointer) unsafe.Pointer {
	return _MIDIInputPortCreateWithProtocol(client, portName, protocol_, outPort, receiveBlock)
}/* debug [functions.gen.go/function]: MIDIInputPortCreateWithProtocol */

// Locates a device, entity, or endpoint by its unique identifier.
//
// Added in macOS 10.2.
// Locates a device, entity, or endpoint by its unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectFindByUniqueID(_:_:_:)
func MIDIObjectFindByUniqueID(inUniqueID MIDIUniqueID, outObject unsafe.Pointer, outObjectType unsafe.Pointer) unsafe.Pointer {
	return _MIDIObjectFindByUniqueID(inUniqueID, outObject, outObjectType)
}/* debug [functions.gen.go/function]: MIDIObjectFindByUniqueID */

// Gets an object’s data-type property.
//
// Added in macOS 10.0.
// Gets an object’s data-type property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectGetDataProperty(_:_:_:)
func MIDIObjectGetDataProperty(obj MIDIObjectRef, propertyID StringRef, outData unsafe.Pointer) unsafe.Pointer {
	return _MIDIObjectGetDataProperty(obj, propertyID, outData)
}/* debug [functions.gen.go/function]: MIDIObjectGetDataProperty */

// Gets an object’s dictionary-type property.
//
// Added in macOS 10.2.
// Gets an object’s dictionary-type property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectGetDictionaryProperty(_:_:_:)
func MIDIObjectGetDictionaryProperty(obj MIDIObjectRef, propertyID StringRef, outDict unsafe.Pointer) unsafe.Pointer {
	return _MIDIObjectGetDictionaryProperty(obj, propertyID, outDict)
}/* debug [functions.gen.go/function]: MIDIObjectGetDictionaryProperty */

// Gets an object’s integer-type property.
//
// Added in macOS 10.0.
// Gets an object’s integer-type property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectGetIntegerProperty(_:_:_:)
func MIDIObjectGetIntegerProperty(obj MIDIObjectRef, propertyID StringRef, outValue unsafe.Pointer) unsafe.Pointer {
	return _MIDIObjectGetIntegerProperty(obj, propertyID, outValue)
}/* debug [functions.gen.go/function]: MIDIObjectGetIntegerProperty */

// Returns all properties of an object.
//
// Added in macOS 10.1.
// Returns all properties of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectGetProperties(_:_:_:)
func MIDIObjectGetProperties(obj MIDIObjectRef, outProperties unsafe.Pointer, deep unsafe.Pointer) unsafe.Pointer {
	return _MIDIObjectGetProperties(obj, outProperties, deep)
}/* debug [functions.gen.go/function]: MIDIObjectGetProperties */

// Gets an object’s string-type property.
//
// Added in macOS 10.0.
// Gets an object’s string-type property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectGetStringProperty(_:_:_:)
func MIDIObjectGetStringProperty(obj MIDIObjectRef, propertyID StringRef, str unsafe.Pointer) unsafe.Pointer {
	return _MIDIObjectGetStringProperty(obj, propertyID, str)
}/* debug [functions.gen.go/function]: MIDIObjectGetStringProperty */

// Removes an object’s property.
//
// Added in macOS 10.2.
// Removes an object’s property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectRemoveProperty(_:_:)
func MIDIObjectRemoveProperty(obj MIDIObjectRef, propertyID StringRef) unsafe.Pointer {
	return _MIDIObjectRemoveProperty(obj, propertyID)
}/* debug [functions.gen.go/function]: MIDIObjectRemoveProperty */

// Sets an object’s data-type property.
//
// Added in macOS 10.0.
// Sets an object’s data-type property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectSetDataProperty(_:_:_:)
func MIDIObjectSetDataProperty(obj MIDIObjectRef, propertyID StringRef, data DataRef) unsafe.Pointer {
	return _MIDIObjectSetDataProperty(obj, propertyID, data)
}/* debug [functions.gen.go/function]: MIDIObjectSetDataProperty */

// Sets an object’s dictionary-type property.
//
// Added in macOS 10.2.
// Sets an object’s dictionary-type property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectSetDictionaryProperty(_:_:_:)
func MIDIObjectSetDictionaryProperty(obj MIDIObjectRef, propertyID StringRef, dict DictionaryRef) unsafe.Pointer {
	return _MIDIObjectSetDictionaryProperty(obj, propertyID, dict)
}/* debug [functions.gen.go/function]: MIDIObjectSetDictionaryProperty */

// Sets an object’s integer-type property.
//
// Added in macOS 10.0.
// Sets an object’s integer-type property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectSetIntegerProperty(_:_:_:)
func MIDIObjectSetIntegerProperty(obj MIDIObjectRef, propertyID StringRef, value unsafe.Pointer) unsafe.Pointer {
	return _MIDIObjectSetIntegerProperty(obj, propertyID, value)
}/* debug [functions.gen.go/function]: MIDIObjectSetIntegerProperty */

// Sets an object’s string-type property.
//
// Added in macOS 10.0.
// Sets an object’s string-type property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectSetStringProperty(_:_:_:)
func MIDIObjectSetStringProperty(obj MIDIObjectRef, propertyID StringRef, str StringRef) unsafe.Pointer {
	return _MIDIObjectSetStringProperty(obj, propertyID, str)
}/* debug [functions.gen.go/function]: MIDIObjectSetStringProperty */

// Creates an output port through which a client sends outgoing MIDI messages to any MIDI destination.
//
// Added in macOS 10.0.
// Creates an output port through which a client sends outgoing MIDI messages to any MIDI destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIOutputPortCreate(_:_:_:)
func MIDIOutputPortCreate(client MIDIClientRef, portName StringRef, outPort unsafe.Pointer) unsafe.Pointer {
	return _MIDIOutputPortCreate(client, portName, outPort)
}/* debug [functions.gen.go/function]: MIDIOutputPortCreate */

// Adds a MIDI event to a MIDIPacketList.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
// Adds a MIDI event to a MIDIPacketList.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPacketListAdd(_:_:_:_:_:_:)
func MIDIPacketListAdd(pktlist unsafe.Pointer, listSize unsafe.Pointer, curPacket unsafe.Pointer, time MIDITimeStamp, nData unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _MIDIPacketListAdd(pktlist, listSize, curPacket, time, nData, data)
}/* debug [functions.gen.go/function]: MIDIPacketListAdd */

// Prepares a MIDIPacketList to be built up dynamically.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
// Prepares a MIDIPacketList to be built up dynamically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPacketListInit(_:)
func MIDIPacketListInit(pktlist unsafe.Pointer) unsafe.Pointer {
	return _MIDIPacketListInit(pktlist)
}/* debug [functions.gen.go/function]: MIDIPacketListInit */

// Makes a connection from a source to a client input port.
//
// Added in macOS 10.0.
// Makes a connection from a source to a client input port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPortConnectSource(_:_:_:)
func MIDIPortConnectSource(port MIDIPortRef, source MIDIEndpointRef, connRefCon unsafe.Pointer) unsafe.Pointer {
	return _MIDIPortConnectSource(port, source, connRefCon)
}/* debug [functions.gen.go/function]: MIDIPortConnectSource */

// Closes a previously established source-to-input port connection.
//
// Added in macOS 10.0.
// Closes a previously established source-to-input port connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPortDisconnectSource(_:_:)
func MIDIPortDisconnectSource(port MIDIPortRef, source MIDIEndpointRef) unsafe.Pointer {
	return _MIDIPortDisconnectSource(port, source)
}/* debug [functions.gen.go/function]: MIDIPortDisconnectSource */

// Disposes of a MIDI port.
//
// Added in macOS 10.0.
// Disposes of a MIDI port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPortDispose(_:)
func MIDIPortDispose(port MIDIPortRef) unsafe.Pointer {
	return _MIDIPortDispose(port)
}/* debug [functions.gen.go/function]: MIDIPortDispose */

// Distributes incoming MIDI from a source to the client input ports which are connected to that source.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
// Distributes incoming MIDI from a source to the client input ports which are connected to that source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIReceived(_:_:)
func MIDIReceived(src MIDIEndpointRef, pktlist unsafe.Pointer) unsafe.Pointer {
	return _MIDIReceived(src, pktlist)
}/* debug [functions.gen.go/function]: MIDIReceived */

// Distributes incoming MIDI events from a source to its connected client input ports.
//
// Added in macOS 11.0.
// Distributes incoming MIDI events from a source to its connected client input ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIReceivedEventList(_:_:)
func MIDIReceivedEventList(src MIDIEndpointRef, evtlist unsafe.Pointer) unsafe.Pointer {
	return _MIDIReceivedEventList(src, evtlist)
}/* debug [functions.gen.go/function]: MIDIReceivedEventList */

// Stops and restarts MIDI I/O.
//
// Added in macOS 10.1.
// Stops and restarts MIDI I/O.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIRestart()
func MIDIRestart() unsafe.Pointer {
	return _MIDIRestart()
}/* debug [functions.gen.go/function]: MIDIRestart */

// Sends MIDI to a destination.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
// Sends MIDI to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISend(_:_:_:)
func MIDISend(port MIDIPortRef, dest MIDIEndpointRef, pktlist unsafe.Pointer) unsafe.Pointer {
	return _MIDISend(port, dest, pktlist)
}/* debug [functions.gen.go/function]: MIDISend */

// Sends MIDI events to a destination.
//
// Added in macOS 11.0.
// Sends MIDI events to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISendEventList(_:_:_:)
func MIDISendEventList(port MIDIPortRef, dest MIDIEndpointRef, evtlist unsafe.Pointer) unsafe.Pointer {
	return _MIDISendEventList(port, dest, evtlist)
}/* debug [functions.gen.go/function]: MIDISendEventList */

// Asynchronously sends a single system-exclusive (SysEx) event.
//
// Added in macOS 10.0.
// Asynchronously sends a single system-exclusive (SysEx) event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISendSysex(_:)
func MIDISendSysex(request unsafe.Pointer) unsafe.Pointer {
	return _MIDISendSysex(request)
}/* debug [functions.gen.go/function]: MIDISendSysex */

// Asynchronously sends a single universal MIDI packet (UMP) system-exclusive (SysEx) event.
//
// Added in macOS 14.0.
// Asynchronously sends a single universal MIDI packet (UMP) system-exclusive (SysEx) event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISendUMPSysex(_:)
func MIDISendUMPSysex(umpRequest unsafe.Pointer) unsafe.Pointer {
	return _MIDISendUMPSysex(umpRequest)
}/* debug [functions.gen.go/function]: MIDISendUMPSysex */

// Asynchronously sends a single universal MIDI packet (UMP) system-exclusive (SysEx) event with an 8-bit message.
//
// Added in macOS 14.0.
// Asynchronously sends a single universal MIDI packet (UMP) system-exclusive (SysEx) event with an 8-bit message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISendUMPSysex8(_:)
func MIDISendUMPSysex8(umpRequest unsafe.Pointer) unsafe.Pointer {
	return _MIDISendUMPSysex8(umpRequest)
}/* debug [functions.gen.go/function]: MIDISendUMPSysex8 */

// Specifies the MIDI driver that owns a serial port.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
// Specifies the MIDI driver that owns a serial port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetSerialPortOwner
func MIDISetSerialPortOwner(portName StringRef, driverName StringRef) unsafe.Pointer {
	return _MIDISetSerialPortOwner(portName, driverName)
}/* debug [functions.gen.go/function]: MIDISetSerialPortOwner */

// Adds a driver-owned MIDI device to the current MIDI setup.
//
// Added in macOS 10.1.
// Adds a driver-owned MIDI device to the current MIDI setup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupAddDevice(_:)
func MIDISetupAddDevice(device MIDIDeviceRef) unsafe.Pointer {
	return _MIDISetupAddDevice(device)
}/* debug [functions.gen.go/function]: MIDISetupAddDevice */

// Adds an external MIDI device to the current MIDI setup.
//
// Added in macOS 10.1.
// Adds an external MIDI device to the current MIDI setup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupAddExternalDevice(_:)
func MIDISetupAddExternalDevice(device MIDIDeviceRef) unsafe.Pointer {
	return _MIDISetupAddExternalDevice(device)
}/* debug [functions.gen.go/function]: MIDISetupAddExternalDevice */

// Queries drivers to discover what hardware is available.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Queries drivers to discover what hardware is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupCreate
func MIDISetupCreate(outSetup unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupCreate(outSetup)
}/* debug [functions.gen.go/function]: MIDISetupCreate */

// Disposes the specified setup object.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Disposes the specified setup object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupDispose
func MIDISetupDispose(setup MIDISetupRef) unsafe.Pointer {
	return _MIDISetupDispose(setup)
}/* debug [functions.gen.go/function]: MIDISetupDispose */

// Creates a MIDISetup object from an XML stream.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Creates a MIDISetup object from an XML stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupFromData
func MIDISetupFromData(data DataRef, outSetup unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupFromData(data, outSetup)
}/* debug [functions.gen.go/function]: MIDISetupFromData */

// Returns the system’s current MIDISetup.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Returns the system’s current MIDISetup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupGetCurrent
func MIDISetupGetCurrent(outSetup unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupGetCurrent(outSetup)
}/* debug [functions.gen.go/function]: MIDISetupGetCurrent */

// Installs a MIDISetup as the system’s current state.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Installs a MIDISetup as the system’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupInstall
func MIDISetupInstall(setup MIDISetupRef) unsafe.Pointer {
	return _MIDISetupInstall(setup)
}/* debug [functions.gen.go/function]: MIDISetupInstall */

// Removes a driver-owned MIDI device from the current MIDI setup.
//
// Added in macOS 10.1.
// Removes a driver-owned MIDI device from the current MIDI setup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupRemoveDevice(_:)
func MIDISetupRemoveDevice(device MIDIDeviceRef) unsafe.Pointer {
	return _MIDISetupRemoveDevice(device)
}/* debug [functions.gen.go/function]: MIDISetupRemoveDevice */

// Removes an external MIDI device from the current MIDI setup.
//
// Added in macOS 10.1.
// Removes an external MIDI device from the current MIDI setup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupRemoveExternalDevice(_:)
func MIDISetupRemoveExternalDevice(device MIDIDeviceRef) unsafe.Pointer {
	return _MIDISetupRemoveExternalDevice(device)
}/* debug [functions.gen.go/function]: MIDISetupRemoveExternalDevice */

// Creates an XML representation of a MIDISetup object.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Creates an XML representation of a MIDISetup object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupToData
func MIDISetupToData(setup MIDISetupRef, outData unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupToData(setup, outData)
}/* debug [functions.gen.go/function]: MIDISetupToData */

// Creates a virtual source in a client.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
// Creates a virtual source in a client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISourceCreate(_:_:_:)
func MIDISourceCreate(client MIDIClientRef, name StringRef, outSrc unsafe.Pointer) unsafe.Pointer {
	return _MIDISourceCreate(client, name, outSrc)
}/* debug [functions.gen.go/function]: MIDISourceCreate */

// Creates a virtual source in a client.
//
// Added in macOS 11.0.
// Creates a virtual source in a client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISourceCreateWithProtocol(_:_:_:_:)
func MIDISourceCreateWithProtocol(client MIDIClientRef, name StringRef, protocol_ MIDIProtocolID, outSrc unsafe.Pointer) unsafe.Pointer {
	return _MIDISourceCreateWithProtocol(client, name, protocol_, outSrc)
}/* debug [functions.gen.go/function]: MIDISourceCreateWithProtocol */

// Creates a MIDI thru connection.
//
// Added in macOS 10.2.
// Creates a MIDI thru connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionCreate(_:_:_:)
func MIDIThruConnectionCreate(inPersistentOwnerID StringRef, inConnectionParams DataRef, outConnection unsafe.Pointer) unsafe.Pointer {
	return _MIDIThruConnectionCreate(inPersistentOwnerID, inConnectionParams, outConnection)
}/* debug [functions.gen.go/function]: MIDIThruConnectionCreate */

// Disposes a MIDI thru connection.
//
// Added in macOS 10.2.
// Disposes a MIDI thru connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionDispose(_:)
func MIDIThruConnectionDispose(connection MIDIThruConnectionRef) unsafe.Pointer {
	return _MIDIThruConnectionDispose(connection)
}/* debug [functions.gen.go/function]: MIDIThruConnectionDispose */

// Finds the persistent thru connections for the specified client.
//
// Added in macOS 10.2.
// Finds the persistent thru connections for the specified client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionFind(_:_:)
func MIDIThruConnectionFind(inPersistentOwnerID StringRef, outConnectionList unsafe.Pointer) unsafe.Pointer {
	return _MIDIThruConnectionFind(inPersistentOwnerID, outConnectionList)
}/* debug [functions.gen.go/function]: MIDIThruConnectionFind */

// Returns the thru connection’s parameters.
//
// Added in macOS 10.2.
// Returns the thru connection’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionGetParams(_:_:)
func MIDIThruConnectionGetParams(connection MIDIThruConnectionRef, outConnectionParams unsafe.Pointer) unsafe.Pointer {
	return _MIDIThruConnectionGetParams(connection, outConnectionParams)
}/* debug [functions.gen.go/function]: MIDIThruConnectionGetParams */

// Initializes a parameters object with its default values.
//
// Added in macOS 10.2.
// Initializes a parameters object with its default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionParamsInitialize(_:)
func MIDIThruConnectionParamsInitialize(inConnectionParams unsafe.Pointer) {
	_MIDIThruConnectionParamsInitialize(inConnectionParams)
}/* debug [functions.gen.go/function]: MIDIThruConnectionParamsInitialize */

// Updates a thru connection’s parameters.
//
// Added in macOS 10.2.
// Updates a thru connection’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionSetParams(_:_:)
func MIDIThruConnectionSetParams(connection MIDIThruConnectionRef, inConnectionParams DataRef) unsafe.Pointer {
	return _MIDIThruConnectionSetParams(connection, inConnectionParams)
}/* debug [functions.gen.go/function]: MIDIThruConnectionSetParams */




