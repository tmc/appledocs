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
	Nonce() string
	SetNonce(value string)
	RequestedOperation() AuthorizationOpenIDOperation
	SetRequestedOperation(value IAuthorizationOpenIDOperation)
	RequestedScopes() []string
	SetRequestedScopes(value []string)
	State() string
	SetState(value string)
}

// An OpenID authorization request.


// An OpenID authorization request.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/nonce
func (a_ AuthorizationOpenIDRequest) Nonce() string {
	rv := objc.Send[string](a_.ID, objc.Sel("nonce"))
	return rv
}


// A string value to pass to the identity provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/nonce
func (a_ AuthorizationOpenIDRequest) SetNonce(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNonce:"), objc.String(value))
}


// The OpenID authentication operation you want this request to perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedOperation
func (a_ AuthorizationOpenIDRequest) RequestedOperation() AuthorizationOpenIDOperation {
	rv := objc.Send[AuthorizationOpenIDOperation](a_.ID, objc.Sel("requestedOperation"))
	return rv
}


// The OpenID authentication operation you want this request to perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedOperation
func (a_ AuthorizationOpenIDRequest) SetRequestedOperation(value IAuthorizationOpenIDOperation) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestedOperation:"), value)
}


// The contact information to be requested from the user during authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedScopes
func (a_ AuthorizationOpenIDRequest) RequestedScopes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("requestedScopes"))
	return rv
}


// The contact information to be requested from the user during authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedScopes
func (a_ AuthorizationOpenIDRequest) SetRequestedScopes(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestedScopes:"), nsArray)
}


// Data that’s returned to you unmodified in the corresponding credential after a successful authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/state
func (a_ AuthorizationOpenIDRequest) State() string {
	rv := objc.Send[string](a_.ID, objc.Sel("state"))
	return rv
}


// Data that’s returned to you unmodified in the corresponding credential after a successful authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/state
func (a_ AuthorizationOpenIDRequest) SetState(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setState:"), objc.String(value))
}



