// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// GameController Functions (2 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_GCExtendedGamepadSnapshotDataFromNSData func(unsafe.Pointer, unsafe.Pointer) bool
	_GCMicroGamepadSnapshotDataFromNSData func(unsafe.Pointer, unsafe.Pointer) bool
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_GCExtendedGamepadSnapshotDataFromNSData, lib, "GCExtendedGamepadSnapshotDataFromNSData")
	tryRegister(&_GCMicroGamepadSnapshotDataFromNSData, lib, "GCMicroGamepadSnapshotDataFromNSData")
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



// GCExtendedGamepadSnapshotDataFromNSData is a GameController function.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.15.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshotDataFromNSData(_:_:)

func GCExtendedGamepadSnapshotDataFromNSData(snapshotData unsafe.Pointer, data unsafe.Pointer) bool {
	return _GCExtendedGamepadSnapshotDataFromNSData(snapshotData, data)
	}


// GCMicroGamepadSnapshotDataFromNSData is a GameController function.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.15.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshotDataFromNSData(_:_:)

func GCMicroGamepadSnapshotDataFromNSData(snapshotData unsafe.Pointer, data unsafe.Pointer) bool {
	return _GCMicroGamepadSnapshotDataFromNSData(snapshotData, data)
	}




