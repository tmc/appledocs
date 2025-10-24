// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

/* debug [functions.gen.go]: Generating 1 functions for StoreKit */
import (
	"github.com/ebitengine/purego"
)

// StoreKit Functions (1 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_SKTerminateForInvalidReceipt func()
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_SKTerminateForInvalidReceipt, lib, "SKTerminateForInvalidReceipt")
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

// Terminates an app if the license to use the app has expired.
//
// Added in macOS 10.14.
// Terminates an app if the license to use the app has expired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKTerminateForInvalidReceipt()
func SKTerminateForInvalidReceipt() {
	_SKTerminateForInvalidReceipt()
} /* debug [functions.gen.go/function]: SKTerminateForInvalidReceipt */
