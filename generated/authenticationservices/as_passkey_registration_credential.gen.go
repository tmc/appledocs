// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PasskeyRegistrationCredential] class.
var (
	PasskeyRegistrationCredentialClass     _PasskeyRegistrationCredentialClass
	PasskeyRegistrationCredentialClassOnce sync.Once
)

func getPasskeyRegistrationCredentialClass() _PasskeyRegistrationCredentialClass {
	PasskeyRegistrationCredentialClassOnce.Do(func() {
		PasskeyRegistrationCredentialClass = _PasskeyRegistrationCredentialClass{objc.GetClass("ASPasskeyRegistrationCredential")}
	})
	return PasskeyRegistrationCredentialClass
}

type _PasskeyRegistrationCredentialClass struct {
	class objc.Class
}

// An interface definition for the [PasskeyRegistrationCredential] class.
type IPasskeyRegistrationCredential interface {
	objectivec.IObject
}

// A passkey registration credential.
//
// Create a passkey registration credential to provide a response to a passkey registration request from your credential provider extension. Call , passing your passkey registration credential.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential
type PasskeyRegistrationCredential struct {
	objectivec.Object
}

// PasskeyRegistrationCredentialFrom constructs a [PasskeyRegistrationCredential] from an unsafe.Pointer.
//
// A passkey registration credential.
func PasskeyRegistrationCredentialFrom(ptr unsafe.Pointer) PasskeyRegistrationCredential {
	return PasskeyRegistrationCredential{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasskeyRegistrationCredentialClass) Alloc() PasskeyRegistrationCredential {
	rv := objc.Send[PasskeyRegistrationCredential](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasskeyRegistrationCredentialClass) New() PasskeyRegistrationCredential {
	rv := objc.Send[PasskeyRegistrationCredential](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyRegistrationCredential) Init() PasskeyRegistrationCredential {
	rv := objc.Send[PasskeyRegistrationCredential](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyRegistrationCredential) Autorelease() PasskeyRegistrationCredential {
	rv := objc.Send[PasskeyRegistrationCredential](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyRegistrationCredential creates a new PasskeyRegistrationCredential instance.
func NewPasskeyRegistrationCredential() PasskeyRegistrationCredential {
	return getPasskeyRegistrationCredentialClass().New()
}




