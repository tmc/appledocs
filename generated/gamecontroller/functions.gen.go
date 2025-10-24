// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

/* debug [functions.gen.go]: Generating 14 functions for GameController */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// GameController Functions (14 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_GCExtendedGamepadSnapshotDataFromNSData func(unsafe.Pointer, unsafe.Pointer) bool
	_GCExtendedGamepadSnapShotDataV100FromNSData func(unsafe.Pointer, unsafe.Pointer) bool
	_GCGamepadSnapShotDataV100FromNSData func(unsafe.Pointer, unsafe.Pointer) bool
	_GCInputArcadeButtonName func(int, int) unsafe.Pointer
	_GCInputBackLeftButton func(int) unsafe.Pointer
	_GCInputBackRightButton func(int) unsafe.Pointer
	_GCMicroGamepadSnapshotDataFromNSData func(unsafe.Pointer, unsafe.Pointer) bool
	_GCMicroGamepadSnapShotDataV100FromNSData func(unsafe.Pointer, unsafe.Pointer) bool
	_NSDataFromGCExtendedGamepadSnapshotData func(unsafe.Pointer) unsafe.Pointer
	_NSDataFromGCExtendedGamepadSnapShotDataV100 func(unsafe.Pointer) unsafe.Pointer
	_NSDataFromGCGamepadSnapShotDataV100 func(unsafe.Pointer) unsafe.Pointer
	_NSDataFromGCMicroGamepadSnapshotData func(unsafe.Pointer) unsafe.Pointer
	_NSDataFromGCMicroGamepadSnapShotDataV100 func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromGCPoint2 func(GCPoint2) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_GCExtendedGamepadSnapshotDataFromNSData, lib, "GCExtendedGamepadSnapshotDataFromNSData")
	tryRegister(&_GCExtendedGamepadSnapShotDataV100FromNSData, lib, "GCExtendedGamepadSnapShotDataV100FromNSData")
	tryRegister(&_GCGamepadSnapShotDataV100FromNSData, lib, "GCGamepadSnapShotDataV100FromNSData")
	tryRegister(&_GCInputArcadeButtonName, lib, "GCInputArcadeButtonName")
	tryRegister(&_GCInputBackLeftButton, lib, "GCInputBackLeftButton")
	tryRegister(&_GCInputBackRightButton, lib, "GCInputBackRightButton")
	tryRegister(&_GCMicroGamepadSnapshotDataFromNSData, lib, "GCMicroGamepadSnapshotDataFromNSData")
	tryRegister(&_GCMicroGamepadSnapShotDataV100FromNSData, lib, "GCMicroGamepadSnapShotDataV100FromNSData")
	tryRegister(&_NSDataFromGCExtendedGamepadSnapshotData, lib, "NSDataFromGCExtendedGamepadSnapshotData")
	tryRegister(&_NSDataFromGCExtendedGamepadSnapShotDataV100, lib, "NSDataFromGCExtendedGamepadSnapShotDataV100")
	tryRegister(&_NSDataFromGCGamepadSnapShotDataV100, lib, "NSDataFromGCGamepadSnapShotDataV100")
	tryRegister(&_NSDataFromGCMicroGamepadSnapshotData, lib, "NSDataFromGCMicroGamepadSnapshotData")
	tryRegister(&_NSDataFromGCMicroGamepadSnapShotDataV100, lib, "NSDataFromGCMicroGamepadSnapShotDataV100")
	tryRegister(&_NSStringFromGCPoint2, lib, "NSStringFromGCPoint2")
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
}/* debug [functions.gen.go/function]: GCExtendedGamepadSnapshotDataFromNSData */

// Copies the recorded data from an extended gamepad snapshot into a readable structure.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.9.
// Copies the recorded data from an extended gamepad snapshot into a readable structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapShotDataV100FromNSData(_:_:)
func GCExtendedGamepadSnapShotDataV100FromNSData(snapshotData unsafe.Pointer, data unsafe.Pointer) bool {
	return _GCExtendedGamepadSnapShotDataV100FromNSData(snapshotData, data)
}/* debug [functions.gen.go/function]: GCExtendedGamepadSnapShotDataV100FromNSData */

// Copies the recorded data from a gamepad snapshot into a readable structure.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.9.
// Copies the recorded data from a gamepad snapshot into a readable structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepadSnapShotDataV100FromNSData(_:_:)
func GCGamepadSnapShotDataV100FromNSData(snapshotData unsafe.Pointer, data unsafe.Pointer) bool {
	return _GCGamepadSnapShotDataV100FromNSData(snapshotData, data)
}/* debug [functions.gen.go/function]: GCGamepadSnapShotDataV100FromNSData */

// Returns the name of the arcade stick button at the specified location.
//
// Added in macOS 13.0.
// Returns the name of the arcade stick button at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCInputArcadeButtonName
func GCInputArcadeButtonName(row int, column int) unsafe.Pointer {
	return _GCInputArcadeButtonName(row, column)
}/* debug [functions.gen.go/function]: GCInputArcadeButtonName */

// Returns the name of the back left button at the specified location.
//
// Added in macOS 14.4.
// Returns the name of the back left button at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCInputBackLeftButton
func GCInputBackLeftButton(position int) unsafe.Pointer {
	return _GCInputBackLeftButton(position)
}/* debug [functions.gen.go/function]: GCInputBackLeftButton */

// Returns the name of the back right button at the specified location.
//
// Added in macOS 14.4.
// Returns the name of the back right button at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCInputBackRightButton
func GCInputBackRightButton(position int) unsafe.Pointer {
	return _GCInputBackRightButton(position)
}/* debug [functions.gen.go/function]: GCInputBackRightButton */

// GCMicroGamepadSnapshotDataFromNSData is a GameController function.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshotDataFromNSData(_:_:)
func GCMicroGamepadSnapshotDataFromNSData(snapshotData unsafe.Pointer, data unsafe.Pointer) bool {
	return _GCMicroGamepadSnapshotDataFromNSData(snapshotData, data)
}/* debug [functions.gen.go/function]: GCMicroGamepadSnapshotDataFromNSData */

// Copies the recorded data from a micro gamepad snapshot into a readable structure.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.11.
// Copies the recorded data from a micro gamepad snapshot into a readable structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapShotDataV100FromNSData(_:_:)
func GCMicroGamepadSnapShotDataV100FromNSData(snapshotData unsafe.Pointer, data unsafe.Pointer) bool {
	return _GCMicroGamepadSnapShotDataV100FromNSData(snapshotData, data)
}/* debug [functions.gen.go/function]: GCMicroGamepadSnapShotDataV100FromNSData */

// NSDataFromGCExtendedGamepadSnapshotData is a GameController function.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/NSDataFromGCExtendedGamepadSnapshotData(_:)
func NSDataFromGCExtendedGamepadSnapshotData(snapshotData unsafe.Pointer) unsafe.Pointer {
	return _NSDataFromGCExtendedGamepadSnapshotData(snapshotData)
}/* debug [functions.gen.go/function]: NSDataFromGCExtendedGamepadSnapshotData */

// Encapsulates the controller data from an extended gamepad structure into a data object.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.9.
// Encapsulates the controller data from an extended gamepad structure into a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/NSDataFromGCExtendedGamepadSnapShotDataV100(_:)
func NSDataFromGCExtendedGamepadSnapShotDataV100(snapshotData unsafe.Pointer) unsafe.Pointer {
	return _NSDataFromGCExtendedGamepadSnapShotDataV100(snapshotData)
}/* debug [functions.gen.go/function]: NSDataFromGCExtendedGamepadSnapShotDataV100 */

// Encapsulates the controller data from a gamepad structure into a data object.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.9.
// Encapsulates the controller data from a gamepad structure into a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/NSDataFromGCGamepadSnapShotDataV100(_:)
func NSDataFromGCGamepadSnapShotDataV100(snapshotData unsafe.Pointer) unsafe.Pointer {
	return _NSDataFromGCGamepadSnapShotDataV100(snapshotData)
}/* debug [functions.gen.go/function]: NSDataFromGCGamepadSnapShotDataV100 */

// NSDataFromGCMicroGamepadSnapshotData is a GameController function.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/NSDataFromGCMicroGamepadSnapshotData(_:)
func NSDataFromGCMicroGamepadSnapshotData(snapshotData unsafe.Pointer) unsafe.Pointer {
	return _NSDataFromGCMicroGamepadSnapshotData(snapshotData)
}/* debug [functions.gen.go/function]: NSDataFromGCMicroGamepadSnapshotData */

// Encapsulates the controller data from a micro gamepad structure into a data object.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.11.
// Encapsulates the controller data from a micro gamepad structure into a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/NSDataFromGCMicroGamepadSnapShotDataV100(_:)
func NSDataFromGCMicroGamepadSnapShotDataV100(snapshotData unsafe.Pointer) unsafe.Pointer {
	return _NSDataFromGCMicroGamepadSnapShotDataV100(snapshotData)
}/* debug [functions.gen.go/function]: NSDataFromGCMicroGamepadSnapShotDataV100 */

// Returns a string representation of a point.
//
// Added in macOS 14.3.
// Returns a string representation of a point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/NSStringFromGCPoint2(_:)
func NSStringFromGCPoint2(point GCPoint2) unsafe.Pointer {
	return _NSStringFromGCPoint2(point)
}/* debug [functions.gen.go/function]: NSStringFromGCPoint2 */




