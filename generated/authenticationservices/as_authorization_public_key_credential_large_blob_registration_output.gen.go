// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput] class.
var (
	AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass     _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass
	AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass() _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass {
	AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass = _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput")}
	})
	return AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass
}

type _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput] class.
type IAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput-c.class
type AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputFrom constructs a [AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	return AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass) Alloc() AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass) New() AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput) Init() AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput) Autorelease() AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput creates a new AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput instance.
func NewAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput() AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	return getAuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput-c.class/isSupported
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput) IsSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSupported"))
	return rv
}



