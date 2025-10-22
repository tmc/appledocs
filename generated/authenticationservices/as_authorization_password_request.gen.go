// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AuthorizationPasswordRequest] class.
var (
	AuthorizationPasswordRequestClass     _AuthorizationPasswordRequestClass
	AuthorizationPasswordRequestClassOnce sync.Once
)

func getAuthorizationPasswordRequestClass() _AuthorizationPasswordRequestClass {
	AuthorizationPasswordRequestClassOnce.Do(func() {
		AuthorizationPasswordRequestClass = _AuthorizationPasswordRequestClass{objc.GetClass("ASAuthorizationPasswordRequest")}
	})
	return AuthorizationPasswordRequestClass
}

type _AuthorizationPasswordRequestClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPasswordRequest] class.
type IAuthorizationPasswordRequest interface {
	IAuthorizationRequest
}

// An authorization request that uses credentials stored in the keychain.


// An authorization request that uses credentials stored in the keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPasswordRequest

type AuthorizationPasswordRequest struct {
	AuthorizationRequest
}

// AuthorizationPasswordRequestFrom constructs a [AuthorizationPasswordRequest] from an unsafe.Pointer.
//
// An authorization request that uses credentials stored in the keychain.
func AuthorizationPasswordRequestFrom(ptr unsafe.Pointer) AuthorizationPasswordRequest {
	return AuthorizationPasswordRequest{
		AuthorizationRequest: AuthorizationRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPasswordRequestClass) Alloc() AuthorizationPasswordRequest {
	rv := objc.Send[AuthorizationPasswordRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPasswordRequestClass) New() AuthorizationPasswordRequest {
	rv := objc.Send[AuthorizationPasswordRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPasswordRequest) Init() AuthorizationPasswordRequest {
	rv := objc.Send[AuthorizationPasswordRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPasswordRequest) Autorelease() AuthorizationPasswordRequest {
	rv := objc.Send[AuthorizationPasswordRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPasswordRequest creates a new AuthorizationPasswordRequest instance.
func NewAuthorizationPasswordRequest() AuthorizationPasswordRequest {
	return getAuthorizationPasswordRequestClass().New()
}




