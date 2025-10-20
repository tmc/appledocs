// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest] class.
var (
	AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass     _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass
	AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClassOnce sync.Once
)

func getAuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass() _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass {
	AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClassOnce.Do(func() {
		AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass = _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass{objc.GetClass("ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest")}
	})
	return AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass
}

type _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest] class.
type IAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest interface {
	IAuthorizationRequest
}

// A class that defines the assertion request type for security key credentials.
//
// Use this class to sign in with an existing credential on a security key.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest
type AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest struct {
	AuthorizationRequest
}

// AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestFrom constructs a [AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest] from an unsafe.Pointer.
//
// A class that defines the assertion request type for security key credentials.
func AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestFrom(ptr unsafe.Pointer) AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	return AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest{
		AuthorizationRequest: AuthorizationRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass) New() AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest) Init() AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest) Autorelease() AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest creates a new AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest instance.
func NewAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest() AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	return getAuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass().New()
}




