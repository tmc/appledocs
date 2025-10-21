// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PasskeyAssertionCredential] class.
var (
	PasskeyAssertionCredentialClass     _PasskeyAssertionCredentialClass
	PasskeyAssertionCredentialClassOnce sync.Once
)

func getPasskeyAssertionCredentialClass() _PasskeyAssertionCredentialClass {
	PasskeyAssertionCredentialClassOnce.Do(func() {
		PasskeyAssertionCredentialClass = _PasskeyAssertionCredentialClass{objc.GetClass("ASPasskeyAssertionCredential")}
	})
	return PasskeyAssertionCredentialClass
}

type _PasskeyAssertionCredentialClass struct {
	class objc.Class
}

// An interface definition for the [PasskeyAssertionCredential] class.
type IPasskeyAssertionCredential interface {
	objectivec.IObject
}

// A passkey assertion credential.
//
// Create a passkey assertion credential to provide a response to a passkey authentication challenge from your credential provider extension. Call , passing your passkey assertion credential.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential
type PasskeyAssertionCredential struct {
	objectivec.Object
}

// PasskeyAssertionCredentialFrom constructs a [PasskeyAssertionCredential] from an unsafe.Pointer.
//
// A passkey assertion credential.
func PasskeyAssertionCredentialFrom(ptr unsafe.Pointer) PasskeyAssertionCredential {
	return PasskeyAssertionCredential{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasskeyAssertionCredentialClass) Alloc() PasskeyAssertionCredential {
	rv := objc.Send[PasskeyAssertionCredential](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasskeyAssertionCredentialClass) New() PasskeyAssertionCredential {
	rv := objc.Send[PasskeyAssertionCredential](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyAssertionCredential) Init() PasskeyAssertionCredential {
	rv := objc.Send[PasskeyAssertionCredential](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyAssertionCredential) Autorelease() PasskeyAssertionCredential {
	rv := objc.Send[PasskeyAssertionCredential](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyAssertionCredential creates a new PasskeyAssertionCredential instance.
func NewPasskeyAssertionCredential() PasskeyAssertionCredential {
	return getPasskeyAssertionCredentialClass().New()
}


// The cryptographic signature of this credential.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/signature
func (p_ PasskeyAssertionCredential) Signature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("signature"))
	return rv
}



