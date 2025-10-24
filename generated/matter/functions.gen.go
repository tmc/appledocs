// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

/* debug [functions.gen.go]: Generating 5 functions for Matter */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Matter Functions (5 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MTRAttributeNameForID func(MTRClusterIDType, unsafe.Pointer) unsafe.Pointer
	_MTREventNameForID func(MTRClusterIDType, unsafe.Pointer) unsafe.Pointer
	_MTRRequestCommandNameForID func(MTRClusterIDType, unsafe.Pointer) unsafe.Pointer
	_MTRResponseCommandNameForID func(MTRClusterIDType, unsafe.Pointer) unsafe.Pointer
	_MTRSetLogCallback func(unsafe.Pointer, unsafe.Pointer)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MTRAttributeNameForID, lib, "MTRAttributeNameForID")
	tryRegister(&_MTREventNameForID, lib, "MTREventNameForID")
	tryRegister(&_MTRRequestCommandNameForID, lib, "MTRRequestCommandNameForID")
	tryRegister(&_MTRResponseCommandNameForID, lib, "MTRResponseCommandNameForID")
	tryRegister(&_MTRSetLogCallback, lib, "MTRSetLogCallback")
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



// MTRAttributeNameForID is a Matter function.
//
// Added in macOS 14.6.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeNameForID(_:_:)
func MTRAttributeNameForID(clusterID MTRClusterIDType, attributeID unsafe.Pointer) unsafe.Pointer {
	return _MTRAttributeNameForID(clusterID, attributeID)
}/* debug [functions.gen.go/function]: MTRAttributeNameForID */

// Resolve Matter event IDs into a descriptive string.
//
// Added in macOS 15.2.
// Resolve Matter event IDs into a descriptive string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventNameForID(_:_:)
func MTREventNameForID(clusterID MTRClusterIDType, eventID unsafe.Pointer) unsafe.Pointer {
	return _MTREventNameForID(clusterID, eventID)
}/* debug [functions.gen.go/function]: MTREventNameForID */

// Resolve Matter request (client to server) command IDs into a descriptive string.
//
// Added in macOS 15.2.
// Resolve Matter request (client to server) command IDs into a descriptive string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRequestCommandNameForID(_:_:)
func MTRRequestCommandNameForID(clusterID MTRClusterIDType, commandID unsafe.Pointer) unsafe.Pointer {
	return _MTRRequestCommandNameForID(clusterID, commandID)
}/* debug [functions.gen.go/function]: MTRRequestCommandNameForID */

// Resolve Matter response (server to client) command IDs into a descriptive string.
//
// Added in macOS 15.2.
// Resolve Matter response (server to client) command IDs into a descriptive string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRResponseCommandNameForID(_:_:)
func MTRResponseCommandNameForID(clusterID MTRClusterIDType, commandID unsafe.Pointer) unsafe.Pointer {
	return _MTRResponseCommandNameForID(clusterID, commandID)
}/* debug [functions.gen.go/function]: MTRResponseCommandNameForID */

// MTRSetLogCallback is a Matter function.
//
// Added in macOS 13.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetLogCallback(_:_:)
func MTRSetLogCallback(logTypeThreshold unsafe.Pointer, callback unsafe.Pointer) {
	_MTRSetLogCallback(logTypeThreshold, callback)
}/* debug [functions.gen.go/function]: MTRSetLogCallback */




