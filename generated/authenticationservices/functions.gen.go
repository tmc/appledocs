// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

/* debug [functions.gen.go]: Generating 1 functions for AuthenticationServices */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// AuthenticationServices Functions (1 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_ASAuthorizationAllSupportedPublicKeyCredentialDescriptorTransports func() []unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_ASAuthorizationAllSupportedPublicKeyCredentialDescriptorTransports, lib, "ASAuthorizationAllSupportedPublicKeyCredentialDescriptorTransports")
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



// An array of currently supported transport types.
//
// Added in macOS 12.0.
// An array of currently supported transport types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor/Transport/allSupported
func ASAuthorizationAllSupportedPublicKeyCredentialDescriptorTransports() []unsafe.Pointer {
	return _ASAuthorizationAllSupportedPublicKeyCredentialDescriptorTransports()
}/* debug [functions.gen.go/function]: ASAuthorizationAllSupportedPublicKeyCredentialDescriptorTransports */




