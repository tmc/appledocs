// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ServiceManagement Functions (6 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_SMCopyAllJobDictionaries func(unsafe.Pointer) unsafe.Pointer
	_SMJobBless func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SMJobCopyDictionary func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SMJobRemove func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SMJobSubmit func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SMLoginItemSetEnabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_SMCopyAllJobDictionaries, lib, "SMCopyAllJobDictionaries")
	tryRegister(&_SMJobBless, lib, "SMJobBless")
	tryRegister(&_SMJobCopyDictionary, lib, "SMJobCopyDictionary")
	tryRegister(&_SMJobRemove, lib, "SMJobRemove")
	tryRegister(&_SMJobSubmit, lib, "SMJobSubmit")
	tryRegister(&_SMLoginItemSetEnabled, lib, "SMLoginItemSetEnabled")
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



// Copies the job description dictionaries for all jobs in the specified domain. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMCopyAllJobDictionaries(_:)
func SMCopyAllJobDictionaries(domain unsafe.Pointer) unsafe.Pointer {
	return _SMCopyAllJobDictionaries(domain)
	}


// Submits the executable for the given label as a job to . [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMJobBless(_:_:_:_:)
func SMJobBless(domain unsafe.Pointer, executableLabel unsafe.Pointer, auth unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	return _SMJobBless(domain, executableLabel, auth, outError)
	}


// Copies the job description dictionary for the specified job label. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMJobCopyDictionary(_:_:)
func SMJobCopyDictionary(domain unsafe.Pointer, jobLabel unsafe.Pointer) unsafe.Pointer {
	return _SMJobCopyDictionary(domain, jobLabel)
	}


// Removes the job with the specified label from the specified domain. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMJobRemove(_:_:_:_:_:)
func SMJobRemove(domain unsafe.Pointer, jobLabel unsafe.Pointer, auth unsafe.Pointer, wait unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	return _SMJobRemove(domain, jobLabel, auth, wait, outError)
	}


// Submits the specified job to the specified domain. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMJobSubmit(_:_:_:_:)
func SMJobSubmit(domain unsafe.Pointer, job unsafe.Pointer, auth unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	return _SMJobSubmit(domain, job, auth, outError)
	}


// Enables a helper executable in the main app-bundle directory. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMLoginItemSetEnabled(_:_:)
func SMLoginItemSetEnabled(identifier unsafe.Pointer, enabled unsafe.Pointer) unsafe.Pointer {
	return _SMLoginItemSetEnabled(identifier, enabled)
	}




