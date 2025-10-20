// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CoreMIDI Functions (72 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MIDIBluetoothDriverActivateAllConnections func() unsafe.Pointer
	_MIDIBluetoothDriverDisconnect func(unsafe.Pointer) unsafe.Pointer
	_MIDIClientCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIClientCreateWithBlock func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIClientDispose func(unsafe.Pointer) unsafe.Pointer
	_MIDIDestinationCreateWithProtocol func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceAddEntity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceDispose func(unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceGetEntity func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceGetNumberOfEntities func(unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceListAddDevice func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceListDispose func(unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceListGetDevice func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceListGetNumberOfDevices func(unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceNewEntity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDeviceRemoveEntity func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIDriverEnableMonitoring func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEndpointDispose func(unsafe.Pointer) unsafe.Pointer
	_MIDIEndpointGetEntity func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEndpointGetRefCons func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEndpointSetRefCons func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEntityAddOrRemoveEndpoints func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEntityGetDestination func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEntityGetDevice func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEntityGetNumberOfDestinations func(unsafe.Pointer) unsafe.Pointer
	_MIDIEntityGetNumberOfSources func(unsafe.Pointer) unsafe.Pointer
	_MIDIEntityGetSource func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEventListAdd func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEventListForEachEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIEventListInit func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIExternalDeviceCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIFlushOutput func(unsafe.Pointer) unsafe.Pointer
	_MIDIGetDestination func(unsafe.Pointer) unsafe.Pointer
	_MIDIGetDevice func(unsafe.Pointer) unsafe.Pointer
	_MIDIGetDriverDeviceList func(unsafe.Pointer) unsafe.Pointer
	_MIDIGetDriverIORunLoop func() unsafe.Pointer
	_MIDIGetExternalDevice func(unsafe.Pointer) unsafe.Pointer
	_MIDIGetNumberOfDestinations func() unsafe.Pointer
	_MIDIGetNumberOfDevices func() unsafe.Pointer
	_MIDIGetNumberOfExternalDevices func() unsafe.Pointer
	_MIDIGetNumberOfSources func() unsafe.Pointer
	_MIDIGetSerialPortDrivers func(unsafe.Pointer) unsafe.Pointer
	_MIDIGetSerialPortOwner func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIGetSource func(unsafe.Pointer) unsafe.Pointer
	_MIDIInputPortCreateWithProtocol func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIObjectFindByUniqueID func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIOutputPortCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIPortConnectSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIPortDisconnectSource func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIPortDispose func(unsafe.Pointer) unsafe.Pointer
	_MIDIReceivedEventList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIRestart func() unsafe.Pointer
	_MIDISendEventList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDISetSerialPortOwner func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDISetupAddDevice func(unsafe.Pointer) unsafe.Pointer
	_MIDISetupAddExternalDevice func(unsafe.Pointer) unsafe.Pointer
	_MIDISetupCreate func(unsafe.Pointer) unsafe.Pointer
	_MIDISetupDispose func(unsafe.Pointer) unsafe.Pointer
	_MIDISetupFromData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDISetupGetCurrent func(unsafe.Pointer) unsafe.Pointer
	_MIDISetupInstall func(unsafe.Pointer) unsafe.Pointer
	_MIDISetupRemoveDevice func(unsafe.Pointer) unsafe.Pointer
	_MIDISetupRemoveExternalDevice func(unsafe.Pointer) unsafe.Pointer
	_MIDISetupToData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDISourceCreateWithProtocol func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIThruConnectionCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIThruConnectionDispose func(unsafe.Pointer) unsafe.Pointer
	_MIDIThruConnectionFind func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIThruConnectionGetParams func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MIDIThruConnectionParamsInitialize func(unsafe.Pointer) unsafe.Pointer
	_MIDIThruConnectionSetParams func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
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
	tryRegister(&_MIDIInputPortCreateWithProtocol, lib, "MIDIInputPortCreateWithProtocol")
	tryRegister(&_MIDIObjectFindByUniqueID, lib, "MIDIObjectFindByUniqueID")
	tryRegister(&_MIDIOutputPortCreate, lib, "MIDIOutputPortCreate")
	tryRegister(&_MIDIPortConnectSource, lib, "MIDIPortConnectSource")
	tryRegister(&_MIDIPortDisconnectSource, lib, "MIDIPortDisconnectSource")
	tryRegister(&_MIDIPortDispose, lib, "MIDIPortDispose")
	tryRegister(&_MIDIReceivedEventList, lib, "MIDIReceivedEventList")
	tryRegister(&_MIDIRestart, lib, "MIDIRestart")
	tryRegister(&_MIDISendEventList, lib, "MIDISendEventList")
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



// Promote all active Bluetooth connections into an online MIDI device capable of input and output. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIBluetoothDriverActivateAllConnections()
func MIDIBluetoothDriverActivateAllConnections() unsafe.Pointer {
	return _MIDIBluetoothDriverActivateAllConnections()
	}


// Disconnect the Bluetooth MIDI driver from a Bluetooth Low Energy MIDI peripheral. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIBluetoothDriverDisconnect(_:)
func MIDIBluetoothDriverDisconnect(uuid unsafe.Pointer) unsafe.Pointer {
	return _MIDIBluetoothDriverDisconnect(uuid)
	}


// Creates a MIDI client. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIClientCreate(_:_:_:_:)
func MIDIClientCreate(name unsafe.Pointer, notifyProc unsafe.Pointer, notifyRefCon unsafe.Pointer, outClient unsafe.Pointer) unsafe.Pointer {
	return _MIDIClientCreate(name, notifyProc, notifyRefCon, outClient)
	}


// Creates a MIDI client with a callback block. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIClientCreateWithBlock(_:_:_:)
func MIDIClientCreateWithBlock(name unsafe.Pointer, outClient unsafe.Pointer, notifyBlock unsafe.Pointer) unsafe.Pointer {
	return _MIDIClientCreateWithBlock(name, outClient, notifyBlock)
	}


// Disposes of a MIDI client. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIClientDispose(_:)
func MIDIClientDispose(client unsafe.Pointer) unsafe.Pointer {
	return _MIDIClientDispose(client)
	}


// Creates a virtual destination in a client. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDestinationCreateWithProtocol(_:_:_:_:_:)
func MIDIDestinationCreateWithProtocol(client unsafe.Pointer, name unsafe.Pointer, protocol unsafe.Pointer, outDest unsafe.Pointer, readBlock unsafe.Pointer) unsafe.Pointer {
	return _MIDIDestinationCreateWithProtocol(client, name, protocol, outDest, readBlock)
	}


// Specifies one of the entities that make up a device. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceAddEntity(_:_:_:_:_:_:)
func MIDIDeviceAddEntity(device unsafe.Pointer, name unsafe.Pointer, embedded unsafe.Pointer, numSourceEndpoints unsafe.Pointer, numDestinationEndpoints unsafe.Pointer, newEntity unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceAddEntity(device, name, embedded, numSourceEndpoints, numDestinationEndpoints, newEntity)
	}


// Creates a new device object that corresponds to the available hardware. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceCreate(_:_:_:_:_:)
func MIDIDeviceCreate(owner unsafe.Pointer, name unsafe.Pointer, manufacturer unsafe.Pointer, model unsafe.Pointer, outDevice unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceCreate(owner, name, manufacturer, model, outDevice)
	}


// Disposes of a MIDI device. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceDispose(_:)
func MIDIDeviceDispose(device unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceDispose(device)
	}


// Returns the device’s entity at a specific index. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceGetEntity(_:_:)
func MIDIDeviceGetEntity(device unsafe.Pointer, entityIndex0 unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceGetEntity(device, entityIndex0)
	}


// Returns the number of entities in a device. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceGetNumberOfEntities(_:)
func MIDIDeviceGetNumberOfEntities(device unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceGetNumberOfEntities(device)
	}


// Adds the specified device to the device list. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListAddDevice(_:_:)
func MIDIDeviceListAddDevice(devList unsafe.Pointer, dev unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceListAddDevice(devList, dev)
	}


// Disposes of a device list, but not its devices. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListDispose(_:)
func MIDIDeviceListDispose(devList unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceListDispose(devList)
	}


// Retrieves a MIDI device from a device list. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListGetDevice(_:_:)
func MIDIDeviceListGetDevice(devList unsafe.Pointer, index0 unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceListGetDevice(devList, index0)
	}


// Retrieves the number of devices in a device list. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListGetNumberOfDevices(_:)
func MIDIDeviceListGetNumberOfDevices(devList unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceListGetNumberOfDevices(devList)
	}


// Adds a new entity to a device. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceNewEntity(_:_:_:_:_:_:_:)
func MIDIDeviceNewEntity(device unsafe.Pointer, name unsafe.Pointer, protocol unsafe.Pointer, embedded unsafe.Pointer, numSourceEndpoints unsafe.Pointer, numDestinationEndpoints unsafe.Pointer, newEntity unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceNewEntity(device, name, protocol, embedded, numSourceEndpoints, numDestinationEndpoints, newEntity)
	}


// Removes an entity from a device. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceRemoveEntity(_:_:)
func MIDIDeviceRemoveEntity(device unsafe.Pointer, entity unsafe.Pointer) unsafe.Pointer {
	return _MIDIDeviceRemoveEntity(device, entity)
	}


// Enables monitoring of all outgoing MIDI packets. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDriverEnableMonitoring(_:_:)
func MIDIDriverEnableMonitoring(driver unsafe.Pointer, enabled unsafe.Pointer) unsafe.Pointer {
	return _MIDIDriverEnableMonitoring(driver, enabled)
	}


// Disposes of a virtual source or destination. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointDispose(_:)
func MIDIEndpointDispose(endpt unsafe.Pointer) unsafe.Pointer {
	return _MIDIEndpointDispose(endpt)
	}


// Returns an endpoint’s entity. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointGetEntity(_:_:)
func MIDIEndpointGetEntity(inEndpoint unsafe.Pointer, outEntity unsafe.Pointer) unsafe.Pointer {
	return _MIDIEndpointGetEntity(inEndpoint, outEntity)
	}


// Returns contextual data assigned to an endpoint. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointGetRefCons(_:_:_:)
func MIDIEndpointGetRefCons(endpt unsafe.Pointer, ref1 unsafe.Pointer, ref2 unsafe.Pointer) unsafe.Pointer {
	return _MIDIEndpointGetRefCons(endpt, ref1, ref2)
	}


// Sets contextual data on an endpoint. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointSetRefCons(_:_:_:)
func MIDIEndpointSetRefCons(endpt unsafe.Pointer, ref1 unsafe.Pointer, ref2 unsafe.Pointer) unsafe.Pointer {
	return _MIDIEndpointSetRefCons(endpt, ref1, ref2)
	}


// Adds or removes an entity’s endpoints. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityAddOrRemoveEndpoints(_:_:_:)
func MIDIEntityAddOrRemoveEndpoints(entity unsafe.Pointer, numSourceEndpoints unsafe.Pointer, numDestinationEndpoints unsafe.Pointer) unsafe.Pointer {
	return _MIDIEntityAddOrRemoveEndpoints(entity, numSourceEndpoints, numDestinationEndpoints)
	}


// Returns one of an entity’s destinations. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetDestination(_:_:)
func MIDIEntityGetDestination(entity unsafe.Pointer, destIndex0 unsafe.Pointer) unsafe.Pointer {
	return _MIDIEntityGetDestination(entity, destIndex0)
	}


// Returns an entity’s device. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetDevice(_:_:)
func MIDIEntityGetDevice(inEntity unsafe.Pointer, outDevice unsafe.Pointer) unsafe.Pointer {
	return _MIDIEntityGetDevice(inEntity, outDevice)
	}


// Returns the number of destinations in an entity. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetNumberOfDestinations(_:)
func MIDIEntityGetNumberOfDestinations(entity unsafe.Pointer) unsafe.Pointer {
	return _MIDIEntityGetNumberOfDestinations(entity)
	}


// Returns the number of sources in an entity. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetNumberOfSources(_:)
func MIDIEntityGetNumberOfSources(entity unsafe.Pointer) unsafe.Pointer {
	return _MIDIEntityGetNumberOfSources(entity)
	}


// Returns one of an entity’s sources. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetSource(_:_:)
func MIDIEntityGetSource(entity unsafe.Pointer, sourceIndex0 unsafe.Pointer) unsafe.Pointer {
	return _MIDIEntityGetSource(entity, sourceIndex0)
	}


// Adds an event to an event list. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventListAdd(_:_:_:_:_:_:)
func MIDIEventListAdd(evtlist unsafe.Pointer, listSize unsafe.Pointer, curPacket unsafe.Pointer, time unsafe.Pointer, wordCount unsafe.Pointer, words unsafe.Pointer) unsafe.Pointer {
	return _MIDIEventListAdd(evtlist, listSize, curPacket, time, wordCount, words)
	}


// MIDIEventListForEachEvent is a CoreMIDI function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventListForEachEvent(_:_:_:)
func MIDIEventListForEachEvent(evtlist unsafe.Pointer, visitor unsafe.Pointer, visitorContext unsafe.Pointer) {
	_MIDIEventListForEachEvent(evtlist, visitor, visitorContext)
	}


// Initializes an event list. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventListInit(_:_:)
func MIDIEventListInit(evtlist unsafe.Pointer, protocol unsafe.Pointer) unsafe.Pointer {
	return _MIDIEventListInit(evtlist, protocol)
	}


// Creates an external MIDI device. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIExternalDeviceCreate(_:_:_:_:)
func MIDIExternalDeviceCreate(name unsafe.Pointer, manufacturer unsafe.Pointer, model unsafe.Pointer, outDevice unsafe.Pointer) unsafe.Pointer {
	return _MIDIExternalDeviceCreate(name, manufacturer, model, outDevice)
	}


// Cancels all pending events that were previously scheduled to send. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIFlushOutput(_:)
func MIDIFlushOutput(dest unsafe.Pointer) unsafe.Pointer {
	return _MIDIFlushOutput(dest)
	}


// Returns a destination in the system. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDestination(_:)
func MIDIGetDestination(destIndex0 unsafe.Pointer) unsafe.Pointer {
	return _MIDIGetDestination(destIndex0)
	}


// Returns a device from the system. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDevice(_:)
func MIDIGetDevice(deviceIndex0 unsafe.Pointer) unsafe.Pointer {
	return _MIDIGetDevice(deviceIndex0)
	}


// Returns the list of driver-created devices in the current MIDI setup. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDriverDeviceList(_:)
func MIDIGetDriverDeviceList(driver unsafe.Pointer) unsafe.Pointer {
	return _MIDIGetDriverDeviceList(driver)
	}


// Returns the server’s driver I/O thread. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDriverIORunLoop()
func MIDIGetDriverIORunLoop() unsafe.Pointer {
	return _MIDIGetDriverIORunLoop()
	}


// Returns one of the external devices in the system. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetExternalDevice(_:)
func MIDIGetExternalDevice(deviceIndex0 unsafe.Pointer) unsafe.Pointer {
	return _MIDIGetExternalDevice(deviceIndex0)
	}


// Returns the number of destinations in the system. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfDestinations()
func MIDIGetNumberOfDestinations() unsafe.Pointer {
	return _MIDIGetNumberOfDestinations()
	}


// Returns the number of devices in the system. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfDevices()
func MIDIGetNumberOfDevices() unsafe.Pointer {
	return _MIDIGetNumberOfDevices()
	}


// Returns the number of external MIDI devices in the system. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfExternalDevices()
func MIDIGetNumberOfExternalDevices() unsafe.Pointer {
	return _MIDIGetNumberOfExternalDevices()
	}


// Returns the number of sources in the system. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfSources()
func MIDIGetNumberOfSources() unsafe.Pointer {
	return _MIDIGetNumberOfSources()
	}


// Returns a list of installed MIDI drivers for serial port MIDI devices. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetSerialPortDrivers
func MIDIGetSerialPortDrivers(outDriverNames unsafe.Pointer) unsafe.Pointer {
	return _MIDIGetSerialPortDrivers(outDriverNames)
	}


// Returns the MIDI driver that owns a serial port. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetSerialPortOwner
func MIDIGetSerialPortOwner(portName unsafe.Pointer, outDriverName unsafe.Pointer) unsafe.Pointer {
	return _MIDIGetSerialPortOwner(portName, outDriverName)
	}


// Returns a source in the system. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIGetSource(_:)
func MIDIGetSource(sourceIndex0 unsafe.Pointer) unsafe.Pointer {
	return _MIDIGetSource(sourceIndex0)
	}


// Creates an input port through which the client may receive incoming MIDI messages from any MIDI source. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIInputPortCreateWithProtocol(_:_:_:_:_:)
func MIDIInputPortCreateWithProtocol(client unsafe.Pointer, portName unsafe.Pointer, protocol unsafe.Pointer, outPort unsafe.Pointer, receiveBlock unsafe.Pointer) unsafe.Pointer {
	return _MIDIInputPortCreateWithProtocol(client, portName, protocol, outPort, receiveBlock)
	}


// Locates a device, entity, or endpoint by its unique identifier. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectFindByUniqueID(_:_:_:)
func MIDIObjectFindByUniqueID(inUniqueID unsafe.Pointer, outObject unsafe.Pointer, outObjectType unsafe.Pointer) unsafe.Pointer {
	return _MIDIObjectFindByUniqueID(inUniqueID, outObject, outObjectType)
	}


// Creates an output port through which a client sends outgoing MIDI messages to any MIDI destination. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIOutputPortCreate(_:_:_:)
func MIDIOutputPortCreate(client unsafe.Pointer, portName unsafe.Pointer, outPort unsafe.Pointer) unsafe.Pointer {
	return _MIDIOutputPortCreate(client, portName, outPort)
	}


// Makes a connection from a source to a client input port. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPortConnectSource(_:_:_:)
func MIDIPortConnectSource(port unsafe.Pointer, source unsafe.Pointer, connRefCon unsafe.Pointer) unsafe.Pointer {
	return _MIDIPortConnectSource(port, source, connRefCon)
	}


// Closes a previously established source-to-input port connection. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPortDisconnectSource(_:_:)
func MIDIPortDisconnectSource(port unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	return _MIDIPortDisconnectSource(port, source)
	}


// Disposes of a MIDI port. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPortDispose(_:)
func MIDIPortDispose(port unsafe.Pointer) unsafe.Pointer {
	return _MIDIPortDispose(port)
	}


// Distributes incoming MIDI events from a source to its connected client input ports. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIReceivedEventList(_:_:)
func MIDIReceivedEventList(src unsafe.Pointer, evtlist unsafe.Pointer) unsafe.Pointer {
	return _MIDIReceivedEventList(src, evtlist)
	}


// Stops and restarts MIDI I/O. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIRestart()
func MIDIRestart() unsafe.Pointer {
	return _MIDIRestart()
	}


// Sends MIDI events to a destination. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISendEventList(_:_:_:)
func MIDISendEventList(port unsafe.Pointer, dest unsafe.Pointer, evtlist unsafe.Pointer) unsafe.Pointer {
	return _MIDISendEventList(port, dest, evtlist)
	}


// Specifies the MIDI driver that owns a serial port. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetSerialPortOwner
func MIDISetSerialPortOwner(portName unsafe.Pointer, driverName unsafe.Pointer) unsafe.Pointer {
	return _MIDISetSerialPortOwner(portName, driverName)
	}


// Adds a driver-owned MIDI device to the current MIDI setup. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupAddDevice(_:)
func MIDISetupAddDevice(device unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupAddDevice(device)
	}


// Adds an external MIDI device to the current MIDI setup. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupAddExternalDevice(_:)
func MIDISetupAddExternalDevice(device unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupAddExternalDevice(device)
	}


// Queries drivers to discover what hardware is available. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupCreate
func MIDISetupCreate(outSetup unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupCreate(outSetup)
	}


// Disposes the specified setup object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupDispose
func MIDISetupDispose(setup unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupDispose(setup)
	}


// Creates a MIDISetup object from an XML stream. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupFromData
func MIDISetupFromData(data unsafe.Pointer, outSetup unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupFromData(data, outSetup)
	}


// Returns the system’s current MIDISetup. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupGetCurrent
func MIDISetupGetCurrent(outSetup unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupGetCurrent(outSetup)
	}


// Installs a MIDISetup as the system’s current state. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupInstall
func MIDISetupInstall(setup unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupInstall(setup)
	}


// Removes a driver-owned MIDI device from the current MIDI setup. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupRemoveDevice(_:)
func MIDISetupRemoveDevice(device unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupRemoveDevice(device)
	}


// Removes an external MIDI device from the current MIDI setup. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupRemoveExternalDevice(_:)
func MIDISetupRemoveExternalDevice(device unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupRemoveExternalDevice(device)
	}


// Creates an XML representation of a MIDISetup object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupToData
func MIDISetupToData(setup unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _MIDISetupToData(setup, outData)
	}


// Creates a virtual source in a client. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISourceCreateWithProtocol(_:_:_:_:)
func MIDISourceCreateWithProtocol(client unsafe.Pointer, name unsafe.Pointer, protocol unsafe.Pointer, outSrc unsafe.Pointer) unsafe.Pointer {
	return _MIDISourceCreateWithProtocol(client, name, protocol, outSrc)
	}


// Creates a MIDI thru connection. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionCreate(_:_:_:)
func MIDIThruConnectionCreate(inPersistentOwnerID unsafe.Pointer, inConnectionParams unsafe.Pointer, outConnection unsafe.Pointer) unsafe.Pointer {
	return _MIDIThruConnectionCreate(inPersistentOwnerID, inConnectionParams, outConnection)
	}


// Disposes a MIDI thru connection. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionDispose(_:)
func MIDIThruConnectionDispose(connection unsafe.Pointer) unsafe.Pointer {
	return _MIDIThruConnectionDispose(connection)
	}


// Finds the persistent thru connections for the specified client. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionFind(_:_:)
func MIDIThruConnectionFind(inPersistentOwnerID unsafe.Pointer, outConnectionList unsafe.Pointer) unsafe.Pointer {
	return _MIDIThruConnectionFind(inPersistentOwnerID, outConnectionList)
	}


// Returns the thru connection’s parameters. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionGetParams(_:_:)
func MIDIThruConnectionGetParams(connection unsafe.Pointer, outConnectionParams unsafe.Pointer) unsafe.Pointer {
	return _MIDIThruConnectionGetParams(connection, outConnectionParams)
	}


// Initializes a parameters object with its default values. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionParamsInitialize(_:)
func MIDIThruConnectionParamsInitialize(inConnectionParams unsafe.Pointer) {
	_MIDIThruConnectionParamsInitialize(inConnectionParams)
	}


// Updates a thru connection’s parameters. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionSetParams(_:_:)
func MIDIThruConnectionSetParams(connection unsafe.Pointer, inConnectionParams unsafe.Pointer) unsafe.Pointer {
	return _MIDIThruConnectionSetParams(connection, inConnectionParams)
	}




