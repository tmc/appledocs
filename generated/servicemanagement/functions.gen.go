// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

/* debug [functions.gen.go]: Generating 6 functions for ServiceManagement */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// ServiceManagement Functions (6 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_SMJobBless               func(StringRef, StringRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SMLoginItemSetEnabled    func(StringRef, unsafe.Pointer) unsafe.Pointer
	_SMCopyAllJobDictionaries func(StringRef) ArrayRef
	_SMJobCopyDictionary      func(StringRef, StringRef) DictionaryRef
	_SMJobRemove              func(StringRef, StringRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SMJobSubmit              func(StringRef, DictionaryRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_SMJobBless, lib, "SMJobBless")
	tryRegister(&_SMLoginItemSetEnabled, lib, "SMLoginItemSetEnabled")
	tryRegister(&_SMCopyAllJobDictionaries, lib, "SMCopyAllJobDictionaries")
	tryRegister(&_SMJobCopyDictionary, lib, "SMJobCopyDictionary")
	tryRegister(&_SMJobRemove, lib, "SMJobRemove")
	tryRegister(&_SMJobSubmit, lib, "SMJobSubmit")
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

// Submits the executable for the given label as a job to .
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.6.
// Submits the executable for the given label as a job to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMJobBless(_:_:_:_:)
func SMJobBless(domain StringRef, executableLabel StringRef, auth unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	return _SMJobBless(domain, executableLabel, auth, outError)
} /* debug [functions.gen.go/function]: SMJobBless */

// Enables a helper executable in the main app-bundle directory.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.6.
// Enables a helper executable in the main app-bundle directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMLoginItemSetEnabled(_:_:)
func SMLoginItemSetEnabled(identifier StringRef, enabled unsafe.Pointer) unsafe.Pointer {
	return _SMLoginItemSetEnabled(identifier, enabled)
} /* debug [functions.gen.go/function]: SMLoginItemSetEnabled */

// Copies the job description dictionaries for all jobs in the specified domain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.6.
// Copies the job description dictionaries for all jobs in the specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMCopyAllJobDictionaries(_:)
func SMCopyAllJobDictionaries(domain StringRef) ArrayRef {
	return _SMCopyAllJobDictionaries(domain)
} /* debug [functions.gen.go/function]: SMCopyAllJobDictionaries */

// Copies the job description dictionary for the specified job label.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.6.
// Copies the job description dictionary for the specified job label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMJobCopyDictionary(_:_:)
func SMJobCopyDictionary(domain StringRef, jobLabel StringRef) DictionaryRef {
	return _SMJobCopyDictionary(domain, jobLabel)
} /* debug [functions.gen.go/function]: SMJobCopyDictionary */

// Removes the job with the specified label from the specified domain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.6.
// Removes the job with the specified label from the specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMJobRemove(_:_:_:_:_:)
func SMJobRemove(domain StringRef, jobLabel StringRef, auth unsafe.Pointer, wait unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	return _SMJobRemove(domain, jobLabel, auth, wait, outError)
} /* debug [functions.gen.go/function]: SMJobRemove */

// Submits the specified job to the specified domain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.6.
// Submits the specified job to the specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMJobSubmit(_:_:_:_:)
func SMJobSubmit(domain StringRef, job DictionaryRef, auth unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	return _SMJobSubmit(domain, job, auth, outError)
} /* debug [functions.gen.go/function]: SMJobSubmit */
