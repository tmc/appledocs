// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest] class.
var (
	AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass     _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass
	AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClassOnce sync.Once
)

func getAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass() _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass {
	AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClassOnce.Do(func() {
		AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass = _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass{objc.GetClass("ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest")}
	})
	return AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass
}

type _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest] class.
type IAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest interface {
	IAuthorizationRequest
}

// The object for registering a new security key credential.
//
// Create an instance of this class when registering for a new credential using security key authorization.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest
type AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest struct {
	AuthorizationRequest
}

// AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestFrom constructs a [AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest] from an unsafe.Pointer.
//
// The object for registering a new security key credential.
func AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestFrom(ptr unsafe.Pointer) AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	return AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest{
		AuthorizationRequest: AuthorizationRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass) New() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) Init() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) Autorelease() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest creates a new AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest instance.
func NewAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	return getAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass().New()
}




