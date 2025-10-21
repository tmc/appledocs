// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CoreMediaIO Functions (4 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CMIODeviceStartStream func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOObjectIsPropertySettable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOSampleBufferGetSequenceNumber func(unsafe.Pointer) unsafe.Pointer
	_CMIOStreamDeckJog func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CMIODeviceStartStream, lib, "CMIODeviceStartStream")
	tryRegister(&_CMIOObjectIsPropertySettable, lib, "CMIOObjectIsPropertySettable")
	tryRegister(&_CMIOSampleBufferGetSequenceNumber, lib, "CMIOSampleBufferGetSequenceNumber")
	tryRegister(&_CMIOStreamDeckJog, lib, "CMIOStreamDeckJog")
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



// CMIODeviceStartStream is a CoreMediaIO function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStartStream(_:_:)
func CMIODeviceStartStream(deviceID unsafe.Pointer, streamID unsafe.Pointer) unsafe.Pointer {
	return _CMIODeviceStartStream(deviceID, streamID)
	}


// CMIOObjectIsPropertySettable is a CoreMediaIO function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectIsPropertySettable(_:_:_:)
func CMIOObjectIsPropertySettable(objectID unsafe.Pointer, address unsafe.Pointer, isSettable unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectIsPropertySettable(objectID, address, isSettable)
	}


// CMIOSampleBufferGetSequenceNumber is a CoreMediaIO function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferGetSequenceNumber
func CMIOSampleBufferGetSequenceNumber(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMIOSampleBufferGetSequenceNumber(sbuf)
	}


// CMIOStreamDeckJog is a CoreMediaIO function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckJog(_:_:)
func CMIOStreamDeckJog(streamID unsafe.Pointer, speed unsafe.Pointer) unsafe.Pointer {
	return _CMIOStreamDeckJog(streamID, speed)
	}




