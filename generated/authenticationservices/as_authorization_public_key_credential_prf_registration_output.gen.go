// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationPublicKeyCredentialPRFRegistrationOutput] class.
var (
	AuthorizationPublicKeyCredentialPRFRegistrationOutputClass     _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass
	AuthorizationPublicKeyCredentialPRFRegistrationOutputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialPRFRegistrationOutputClass() _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass {
	AuthorizationPublicKeyCredentialPRFRegistrationOutputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialPRFRegistrationOutputClass = _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialPRFRegistrationOutput")}
	})
	return AuthorizationPublicKeyCredentialPRFRegistrationOutputClass
}

type _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPublicKeyCredentialPRFRegistrationOutput] class.
type IAuthorizationPublicKeyCredentialPRFRegistrationOutput interface {
	objectivec.IObject
	First() foundation.NSData
	IsSupported() bool
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationOutput-c.class
type AuthorizationPublicKeyCredentialPRFRegistrationOutput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialPRFRegistrationOutputFrom constructs a [AuthorizationPublicKeyCredentialPRFRegistrationOutput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialPRFRegistrationOutputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	return AuthorizationPublicKeyCredentialPRFRegistrationOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass) Alloc() AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass) New() AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationOutput) Init() AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationOutput) Autorelease() AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialPRFRegistrationOutput creates a new AuthorizationPublicKeyCredentialPRFRegistrationOutput instance.
func NewAuthorizationPublicKeyCredentialPRFRegistrationOutput() AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	return getAuthorizationPublicKeyCredentialPRFRegistrationOutputClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationOutput-c.class/first
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationOutput) First() foundation.NSData {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("first"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationOutput-c.class/isSupported
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationOutput) IsSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSupported"))
	return rv
}



