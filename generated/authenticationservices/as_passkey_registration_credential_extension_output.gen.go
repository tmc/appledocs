// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PasskeyRegistrationCredentialExtensionOutput] class.
var (
	PasskeyRegistrationCredentialExtensionOutputClass     _PasskeyRegistrationCredentialExtensionOutputClass
	PasskeyRegistrationCredentialExtensionOutputClassOnce sync.Once
)

func getPasskeyRegistrationCredentialExtensionOutputClass() _PasskeyRegistrationCredentialExtensionOutputClass {
	PasskeyRegistrationCredentialExtensionOutputClassOnce.Do(func() {
		PasskeyRegistrationCredentialExtensionOutputClass = _PasskeyRegistrationCredentialExtensionOutputClass{objc.GetClass("ASPasskeyRegistrationCredentialExtensionOutput")}
	})
	return PasskeyRegistrationCredentialExtensionOutputClass
}

type _PasskeyRegistrationCredentialExtensionOutputClass struct {
	class objc.Class
}

// An interface definition for the [PasskeyRegistrationCredentialExtensionOutput] class.
type IPasskeyRegistrationCredentialExtensionOutput interface {
	objectivec.IObject
}

// This class encapsulates output for various WebAuthn extensions used during passkey registration.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredentialExtensionOutput-c.class
type PasskeyRegistrationCredentialExtensionOutput struct {
	objectivec.Object
}

// PasskeyRegistrationCredentialExtensionOutputFrom constructs a [PasskeyRegistrationCredentialExtensionOutput] from an unsafe.Pointer.
//
// This class encapsulates output for various WebAuthn extensions used during passkey registration.
func PasskeyRegistrationCredentialExtensionOutputFrom(ptr unsafe.Pointer) PasskeyRegistrationCredentialExtensionOutput {
	return PasskeyRegistrationCredentialExtensionOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasskeyRegistrationCredentialExtensionOutputClass) Alloc() PasskeyRegistrationCredentialExtensionOutput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasskeyRegistrationCredentialExtensionOutputClass) New() PasskeyRegistrationCredentialExtensionOutput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyRegistrationCredentialExtensionOutput) Init() PasskeyRegistrationCredentialExtensionOutput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyRegistrationCredentialExtensionOutput) Autorelease() PasskeyRegistrationCredentialExtensionOutput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyRegistrationCredentialExtensionOutput creates a new PasskeyRegistrationCredentialExtensionOutput instance.
func NewPasskeyRegistrationCredentialExtensionOutput() PasskeyRegistrationCredentialExtensionOutput {
	return getPasskeyRegistrationCredentialExtensionOutputClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredentialExtensionOutput-c.class/initWithLargeBlobOutput:
func NewPasskeyRegistrationCredentialExtensionOutputWithLargeBlobOutput(largeBlob unsafe.Pointer) PasskeyRegistrationCredentialExtensionOutput {
	instance := getPasskeyRegistrationCredentialExtensionOutputClass().Alloc()
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](instance.ID, objc.Sel("initWithLargeBlobOutput:"), largeBlob)
	rv.Autorelease()
	return rv
}


// Output for operation during passkey registration.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredentialExtensionOutput-c.class/largeBlobRegistrationOutput
func (p_ PasskeyRegistrationCredentialExtensionOutput) LargeBlobRegistrationOutput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("largeBlobRegistrationOutput"))
	return rv
}


