// Code generated from Apple documentation for SceneKit. DO NOT EDIT.

package scenekit

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// SceneKit Functions (1 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_SCNExportJavaScriptModule func(unsafe.Pointer)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_SCNExportJavaScriptModule, lib, "SCNExportJavaScriptModule")
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


// Makes SceneKit classes and global constants available to the specified JavaScript context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 26.0.
//
// Added in macOS 10.8.
//
// [Full Topic]: doc://com.apple.scenekit/documentation/SceneKit/SCNExportJavaScriptModule(_:)
func SCNExportJavaScriptModule(context unsafe.Pointer) {
	_SCNExportJavaScriptModule(context)
	}



