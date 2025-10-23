// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ExceptionHandling Functions (1 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSExceptionHandlerResume func()
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSExceptionHandlerResume, lib, "NSExceptionHandlerResume")
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



// NSExceptionHandlerResume is a ExceptionHandling function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandlerResume()
func NSExceptionHandlerResume() {
	_NSExceptionHandlerResume()
}



