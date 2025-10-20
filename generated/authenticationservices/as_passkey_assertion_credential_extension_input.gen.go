// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PasskeyAssertionCredentialExtensionInput] class.
var (
	PasskeyAssertionCredentialExtensionInputClass     _PasskeyAssertionCredentialExtensionInputClass
	PasskeyAssertionCredentialExtensionInputClassOnce sync.Once
)

func getPasskeyAssertionCredentialExtensionInputClass() _PasskeyAssertionCredentialExtensionInputClass {
	PasskeyAssertionCredentialExtensionInputClassOnce.Do(func() {
		PasskeyAssertionCredentialExtensionInputClass = _PasskeyAssertionCredentialExtensionInputClass{objc.GetClass("ASPasskeyAssertionCredentialExtensionInput")}
	})
	return PasskeyAssertionCredentialExtensionInputClass
}

type _PasskeyAssertionCredentialExtensionInputClass struct {
	class objc.Class
}

// An interface definition for the [PasskeyAssertionCredentialExtensionInput] class.
type IPasskeyAssertionCredentialExtensionInput interface {
	objectivec.IObject
}

// This class encapsulates input for various WebAuthn extensions during passkey assertion.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredentialExtensionInput-c.class
type PasskeyAssertionCredentialExtensionInput struct {
	objectivec.Object
}

// PasskeyAssertionCredentialExtensionInputFrom constructs a [PasskeyAssertionCredentialExtensionInput] from an unsafe.Pointer.
//
// This class encapsulates input for various WebAuthn extensions during passkey assertion.
func PasskeyAssertionCredentialExtensionInputFrom(ptr unsafe.Pointer) PasskeyAssertionCredentialExtensionInput {
	return PasskeyAssertionCredentialExtensionInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasskeyAssertionCredentialExtensionInputClass) Alloc() PasskeyAssertionCredentialExtensionInput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionInput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasskeyAssertionCredentialExtensionInputClass) New() PasskeyAssertionCredentialExtensionInput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionInput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyAssertionCredentialExtensionInput) Init() PasskeyAssertionCredentialExtensionInput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionInput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyAssertionCredentialExtensionInput) Autorelease() PasskeyAssertionCredentialExtensionInput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionInput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyAssertionCredentialExtensionInput creates a new PasskeyAssertionCredentialExtensionInput instance.
func NewPasskeyAssertionCredentialExtensionInput() PasskeyAssertionCredentialExtensionInput {
	return getPasskeyAssertionCredentialExtensionInputClass().New()
}


// Input for the extension in passkey assertion requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredentialExtensionInput-c.class/largeBlob
func (p_ PasskeyAssertionCredentialExtensionInput) LargeBlob() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("largeBlob"))
	return rv
}



