// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// CoreAudio Functions (2 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_AudioConvertNanosToHostTime func(unsafe.Pointer) unsafe.Pointer
	_AudioHardwareDestroyProcessTap func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_AudioConvertNanosToHostTime, lib, "AudioConvertNanosToHostTime")
	tryRegister(&_AudioHardwareDestroyProcessTap, lib, "AudioHardwareDestroyProcessTap")
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


// AudioConvertNanosToHostTime is a CoreAudio function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: doc://com.apple.coreaudio/documentation/CoreAudio/AudioConvertNanosToHostTime(_:)
func AudioConvertNanosToHostTime(inNanos unsafe.Pointer) unsafe.Pointer {
	return _AudioConvertNanosToHostTime(inNanos)
	}


// AudioHardwareDestroyProcessTap is a CoreAudio function. [Full Topic]
//
// Added in macOS 14.2.
//
// [Full Topic]: doc://com.apple.coreaudio/documentation/CoreAudio/AudioHardwareDestroyProcessTap(_:)
func AudioHardwareDestroyProcessTap(inTapID unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareDestroyProcessTap(inTapID)
	}



