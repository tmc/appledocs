// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Matter Functions (8 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MTRAttributeNameForID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MTRClusterNameForID func(unsafe.Pointer) unsafe.Pointer
	_MTRDeviceControllerStorageClasses func() unsafe.Pointer
	_MTREventNameForID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MTRRequestCommandNameForID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MTRResponseCommandNameForID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MTRSetLogCallback func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MTRSetMessageReliabilityParameters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MTRAttributeNameForID, lib, "MTRAttributeNameForID")
	tryRegister(&_MTRClusterNameForID, lib, "MTRClusterNameForID")
	tryRegister(&_MTRDeviceControllerStorageClasses, lib, "MTRDeviceControllerStorageClasses")
	tryRegister(&_MTREventNameForID, lib, "MTREventNameForID")
	tryRegister(&_MTRRequestCommandNameForID, lib, "MTRRequestCommandNameForID")
	tryRegister(&_MTRResponseCommandNameForID, lib, "MTRResponseCommandNameForID")
	tryRegister(&_MTRSetLogCallback, lib, "MTRSetLogCallback")
	tryRegister(&_MTRSetMessageReliabilityParameters, lib, "MTRSetMessageReliabilityParameters")
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



// MTRAttributeNameForID is a Matter function. [Full Topic]
//
// Added in macOS 14.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeNameForID(_:_:)
func MTRAttributeNameForID(clusterID unsafe.Pointer, attributeID unsafe.Pointer) unsafe.Pointer {
	return _MTRAttributeNameForID(clusterID, attributeID)
	}


// MTRClusterNameForID is a Matter function. [Full Topic]
//
// Added in macOS 14.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterNameForID(_:)
func MTRClusterNameForID(clusterID unsafe.Pointer) unsafe.Pointer {
	return _MTRClusterNameForID(clusterID)
	}


// MTRDeviceControllerStorageClasses is a Matter function. [Full Topic]
//
// Added in macOS 14.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStorageClasses()
func MTRDeviceControllerStorageClasses() unsafe.Pointer {
	return _MTRDeviceControllerStorageClasses()
	}


// Resolve Matter event IDs into a descriptive string. [Full Topic]
//
// Added in macOS 15.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventNameForID(_:_:)
func MTREventNameForID(clusterID unsafe.Pointer, eventID unsafe.Pointer) unsafe.Pointer {
	return _MTREventNameForID(clusterID, eventID)
	}


// Resolve Matter request (client to server) command IDs into a descriptive string. [Full Topic]
//
// Added in macOS 15.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRequestCommandNameForID(_:_:)
func MTRRequestCommandNameForID(clusterID unsafe.Pointer, commandID unsafe.Pointer) unsafe.Pointer {
	return _MTRRequestCommandNameForID(clusterID, commandID)
	}


// Resolve Matter response (server to client) command IDs into a descriptive string. [Full Topic]
//
// Added in macOS 15.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRResponseCommandNameForID(_:_:)
func MTRResponseCommandNameForID(clusterID unsafe.Pointer, commandID unsafe.Pointer) unsafe.Pointer {
	return _MTRResponseCommandNameForID(clusterID, commandID)
	}


// MTRSetLogCallback is a Matter function. [Full Topic]
//
// Added in macOS 13.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetLogCallback(_:_:)
func MTRSetLogCallback(logTypeThreshold unsafe.Pointer, callback unsafe.Pointer) {
	_MTRSetLogCallback(logTypeThreshold, callback)
	}


// MTRSetMessageReliabilityParameters is a Matter function. [Full Topic]
//
// Added in macOS 14.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetMessageReliabilityParameters(_:_:_:_:)
func MTRSetMessageReliabilityParameters(idleRetransmitMs unsafe.Pointer, activeRetransmitMs unsafe.Pointer, activeThresholdMs unsafe.Pointer, additionalRetransmitDelayMs unsafe.Pointer) {
	_MTRSetMessageReliabilityParameters(idleRetransmitMs, activeRetransmitMs, activeThresholdMs, additionalRetransmitDelayMs)
	}




