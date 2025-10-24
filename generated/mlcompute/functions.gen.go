// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// MLCompute Functions (1 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MLCPaddingPolicyDebugDescription func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MLCPaddingPolicyDebugDescription, lib, "MLCPaddingPolicyDebugDescription")
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



// A textual description of the padding policy, suitable for debugging.
//
// Added in macOS 11.0.
// A textual description of the padding policy, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingPolicyDebugDescription
func MLCPaddingPolicyDebugDescription(paddingPolicy unsafe.Pointer) unsafe.Pointer {
	return _MLCPaddingPolicyDebugDescription(paddingPolicy)
}



