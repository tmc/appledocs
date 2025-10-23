// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationRequest] class.
var (
	AuthorizationRequestClass     _AuthorizationRequestClass
	AuthorizationRequestClassOnce sync.Once
)

func getAuthorizationRequestClass() _AuthorizationRequestClass {
	AuthorizationRequestClassOnce.Do(func() {
		AuthorizationRequestClass = _AuthorizationRequestClass{objc.GetClass("ASAuthorizationRequest")}
	})
	return AuthorizationRequestClass
}

type _AuthorizationRequestClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationRequest] class.
type IAuthorizationRequest interface {
	objectivec.IObject
}

// A parent class referenced by other AuthenticationServices classes.


// A parent class referenced by other AuthenticationServices classes. [Full Topic]
type AuthorizationRequest struct {
	objectivec.Object
}

// AuthorizationRequestFrom constructs a [AuthorizationRequest] from an unsafe.Pointer.
//
// A parent class referenced by other AuthenticationServices classes.
func AuthorizationRequestFrom(ptr unsafe.Pointer) AuthorizationRequest {
	return AuthorizationRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationRequestClass) Alloc() AuthorizationRequest {
	rv := objc.Send[AuthorizationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationRequestClass) New() AuthorizationRequest {
	rv := objc.Send[AuthorizationRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationRequest) Init() AuthorizationRequest {
	rv := objc.Send[AuthorizationRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationRequest) Autorelease() AuthorizationRequest {
	rv := objc.Send[AuthorizationRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationRequest creates a new AuthorizationRequest instance.
func NewAuthorizationRequest() AuthorizationRequest {
	return getAuthorizationRequestClass().New()
}




