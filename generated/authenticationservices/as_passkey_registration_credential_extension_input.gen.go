// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PasskeyRegistrationCredentialExtensionInput] class.
var (
	PasskeyRegistrationCredentialExtensionInputClass     _PasskeyRegistrationCredentialExtensionInputClass
	PasskeyRegistrationCredentialExtensionInputClassOnce sync.Once
)

func getPasskeyRegistrationCredentialExtensionInputClass() _PasskeyRegistrationCredentialExtensionInputClass {
	PasskeyRegistrationCredentialExtensionInputClassOnce.Do(func() {
		PasskeyRegistrationCredentialExtensionInputClass = _PasskeyRegistrationCredentialExtensionInputClass{objc.GetClass("ASPasskeyRegistrationCredentialExtensionInput")}
	})
	return PasskeyRegistrationCredentialExtensionInputClass
}

type _PasskeyRegistrationCredentialExtensionInputClass struct {
	class objc.Class
}

// An interface definition for the [PasskeyRegistrationCredentialExtensionInput] class.
type IPasskeyRegistrationCredentialExtensionInput interface {
	objectivec.IObject
}

// This class encapsulates input for various WebAuthn extensions during passkey registration.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredentialExtensionInput-c.class
type PasskeyRegistrationCredentialExtensionInput struct {
	objectivec.Object
}

// PasskeyRegistrationCredentialExtensionInputFrom constructs a [PasskeyRegistrationCredentialExtensionInput] from an unsafe.Pointer.
//
// This class encapsulates input for various WebAuthn extensions during passkey registration.
func PasskeyRegistrationCredentialExtensionInputFrom(ptr unsafe.Pointer) PasskeyRegistrationCredentialExtensionInput {
	return PasskeyRegistrationCredentialExtensionInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasskeyRegistrationCredentialExtensionInputClass) Alloc() PasskeyRegistrationCredentialExtensionInput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionInput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasskeyRegistrationCredentialExtensionInputClass) New() PasskeyRegistrationCredentialExtensionInput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionInput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyRegistrationCredentialExtensionInput) Init() PasskeyRegistrationCredentialExtensionInput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionInput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyRegistrationCredentialExtensionInput) Autorelease() PasskeyRegistrationCredentialExtensionInput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionInput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyRegistrationCredentialExtensionInput creates a new PasskeyRegistrationCredentialExtensionInput instance.
func NewPasskeyRegistrationCredentialExtensionInput() PasskeyRegistrationCredentialExtensionInput {
	return getPasskeyRegistrationCredentialExtensionInputClass().New()
}


// Input for the extension in passkey registration requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredentialExtensionInput-c.class/largeBlob
func (p_ PasskeyRegistrationCredentialExtensionInput) LargeBlob() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("largeBlob"))
	return rv
}



