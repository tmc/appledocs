// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

/* debug [functions.gen.go]: Generating 1 functions for PassKit */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// PassKit Functions (1 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_PKPayLaterValidateAmount func(unsafe.Pointer, unsafe.Pointer)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_PKPayLaterValidateAmount, lib, "PKPayLaterValidateAmount")
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



// Checks if the framework can display Apple Pay Later visual merchandising widget information for the given amount and currency.

// Checks if the framework can display Apple Pay Later visual merchandising widget information for the given amount and currency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPayLaterValidateAmount
func PKPayLaterValidateAmount(amount unsafe.Pointer, currencyCode unsafe.Pointer) {
	_PKPayLaterValidateAmount(amount, currencyCode)
}/* debug [functions.gen.go/function]: PKPayLaterValidateAmount */




