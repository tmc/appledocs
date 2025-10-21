// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// QuickLook Functions (1 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_QLPreviewRequestGetGeneratorBundle func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_QLPreviewRequestGetGeneratorBundle, lib, "QLPreviewRequestGetGeneratorBundle")
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



// Get the bundle of the generator receiving the preview request. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestGetGeneratorBundle(_:)
func QLPreviewRequestGetGeneratorBundle(preview unsafe.Pointer) unsafe.Pointer {
	return _QLPreviewRequestGetGeneratorBundle(preview)
	}




