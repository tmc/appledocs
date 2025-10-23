// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AuthorizationAppleIDRequest] class.
var (
	AuthorizationAppleIDRequestClass     _AuthorizationAppleIDRequestClass
	AuthorizationAppleIDRequestClassOnce sync.Once
)

func getAuthorizationAppleIDRequestClass() _AuthorizationAppleIDRequestClass {
	AuthorizationAppleIDRequestClassOnce.Do(func() {
		AuthorizationAppleIDRequestClass = _AuthorizationAppleIDRequestClass{objc.GetClass("ASAuthorizationAppleIDRequest")}
	})
	return AuthorizationAppleIDRequestClass
}

type _AuthorizationAppleIDRequestClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationAppleIDRequest] class.
type IAuthorizationAppleIDRequest interface {
	IAuthorizationOpenIDRequest
	// properties:
	User() string /* primitive/slice/pointer. */
	SetUser(value string /* primitive/slice/pointer. */)
	// methods:
}

// An OpenID authorization request that relies on the user’s Apple ID.


// An OpenID authorization request that relies on the user’s Apple ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDRequest
type AuthorizationAppleIDRequest struct {
	AuthorizationOpenIDRequest
}

// AuthorizationAppleIDRequestFrom constructs a [AuthorizationAppleIDRequest] from an unsafe.Pointer.
//
// An OpenID authorization request that relies on the user’s Apple ID.
func AuthorizationAppleIDRequestFrom(ptr unsafe.Pointer) AuthorizationAppleIDRequest {
	return AuthorizationAppleIDRequest{
		AuthorizationOpenIDRequest: AuthorizationOpenIDRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationAppleIDRequestClass) Alloc() AuthorizationAppleIDRequest {
	rv := objc.Send[AuthorizationAppleIDRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationAppleIDRequestClass) New() AuthorizationAppleIDRequest {
	rv := objc.Send[AuthorizationAppleIDRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationAppleIDRequest) Init() AuthorizationAppleIDRequest {
	rv := objc.Send[AuthorizationAppleIDRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationAppleIDRequest) Autorelease() AuthorizationAppleIDRequest {
	rv := objc.Send[AuthorizationAppleIDRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationAppleIDRequest creates a new AuthorizationAppleIDRequest instance.
func NewAuthorizationAppleIDRequest() AuthorizationAppleIDRequest {
	return getAuthorizationAppleIDRequestClass().New()
}



// An identifier associated with the user’s Apple ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidrequest/user
func (a_ AuthorizationAppleIDRequest) User() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("user"))
	return rv
}


// An identifier associated with the user’s Apple ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidrequest/user
func (a_ AuthorizationAppleIDRequest) SetUser(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUser:"), objc.String(value))
}



