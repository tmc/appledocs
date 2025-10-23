// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	objectivec.IObject
}

// A parent class referenced by other AuthenticationServices classes.


// A parent class referenced by other AuthenticationServices classes. [Full Topic]
type AuthorizationOpenIDRequest struct {
	objectivec.Object
}

// AuthorizationOpenIDRequestFrom constructs a [AuthorizationOpenIDRequest] from an unsafe.Pointer.
//
// A parent class referenced by other AuthenticationServices classes.
func AuthorizationOpenIDRequestFrom(ptr unsafe.Pointer) AuthorizationOpenIDRequest {
	return AuthorizationOpenIDRequest{objectivec.Object{objc.ID(ptr)}}
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




