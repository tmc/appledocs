// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AuthorizationOpenIDRequest] class.
var (
	AuthorizationOpenIDRequestClass     _AuthorizationOpenIDRequestClass
	AuthorizationOpenIDRequestClassOnce sync.Once
)

func getAuthorizationOpenIDRequestClass() _AuthorizationOpenIDRequestClass {
	AuthorizationOpenIDRequestClassOnce.Do(func() {
		AuthorizationOpenIDRequestClass = _AuthorizationOpenIDRequestClass{objc.GetClass("ASAuthorizationOpenIDRequest")}
	})
	return AuthorizationOpenIDRequestClass
}

type _AuthorizationOpenIDRequestClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationOpenIDRequest] class.
type IAuthorizationOpenIDRequest interface {
	IAuthorizationRequest
}

// An OpenID authorization request.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest
type AuthorizationOpenIDRequest struct {
	AuthorizationRequest
}

// AuthorizationOpenIDRequestFrom constructs a [AuthorizationOpenIDRequest] from an unsafe.Pointer.
//
// An OpenID authorization request.
func AuthorizationOpenIDRequestFrom(ptr unsafe.Pointer) AuthorizationOpenIDRequest {
	return AuthorizationOpenIDRequest{
		AuthorizationRequest: AuthorizationRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationOpenIDRequestClass) Alloc() AuthorizationOpenIDRequest {
	rv := objc.Send[AuthorizationOpenIDRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationOpenIDRequestClass) New() AuthorizationOpenIDRequest {
	rv := objc.Send[AuthorizationOpenIDRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationOpenIDRequest) Init() AuthorizationOpenIDRequest {
	rv := objc.Send[AuthorizationOpenIDRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationOpenIDRequest) Autorelease() AuthorizationOpenIDRequest {
	rv := objc.Send[AuthorizationOpenIDRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationOpenIDRequest creates a new AuthorizationOpenIDRequest instance.
func NewAuthorizationOpenIDRequest() AuthorizationOpenIDRequest {
	return getAuthorizationOpenIDRequestClass().New()
}


// A string value to pass to the identity provider.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/nonce
func (a_ AuthorizationOpenIDRequest) Nonce() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("nonce"))
	return rv
}


// SetNonce sets the value of the nonce property.
// A string value to pass to the identity provider.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/nonce
func (a_ AuthorizationOpenIDRequest) SetNonce(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNonce:"), value)
}
// The OpenID authentication operation you want this request to perform.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedOperation
func (a_ AuthorizationOpenIDRequest) RequestedOperation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("requestedOperation"))
	return rv
}


// SetRequestedOperation sets the value of the requestedOperation property.
// The OpenID authentication operation you want this request to perform.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedOperation
func (a_ AuthorizationOpenIDRequest) SetRequestedOperation(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestedOperation:"), value)
}
// The contact information to be requested from the user during authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedScopes
func (a_ AuthorizationOpenIDRequest) RequestedScopes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("requestedScopes"))
	return rv
}


// SetRequestedScopes sets the value of the requestedScopes property.
// The contact information to be requested from the user during authentication.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedScopes
func (a_ AuthorizationOpenIDRequest) SetRequestedScopes(value []string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestedScopes:"), value)
}
// Data that’s returned to you unmodified in the corresponding credential after a successful authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/state
func (a_ AuthorizationOpenIDRequest) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// Data that’s returned to you unmodified in the corresponding credential after a successful authentication.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/state
func (a_ AuthorizationOpenIDRequest) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setState:"), value)
}


