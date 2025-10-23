// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationPasswordProvider] class.
var (
	AuthorizationPasswordProviderClass     _AuthorizationPasswordProviderClass
	AuthorizationPasswordProviderClassOnce sync.Once
)

func getAuthorizationPasswordProviderClass() _AuthorizationPasswordProviderClass {
	AuthorizationPasswordProviderClassOnce.Do(func() {
		AuthorizationPasswordProviderClass = _AuthorizationPasswordProviderClass{objc.GetClass("ASAuthorizationPasswordProvider")}
	})
	return AuthorizationPasswordProviderClass
}

type _AuthorizationPasswordProviderClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPasswordProvider] class.
type IAuthorizationPasswordProvider interface {
	objectivec.IObject
	// properties:
	// methods:
	CreateRequest() IAuthorizationPasswordRequest
}

// A mechanism for generating requests to perform keychain credential sharing.


// A mechanism for generating requests to perform keychain credential sharing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPasswordProvider
type AuthorizationPasswordProvider struct {
	objectivec.Object
}

// AuthorizationPasswordProviderFrom constructs a [AuthorizationPasswordProvider] from an unsafe.Pointer.
//
// A mechanism for generating requests to perform keychain credential sharing.
func AuthorizationPasswordProviderFrom(ptr unsafe.Pointer) AuthorizationPasswordProvider {
	return AuthorizationPasswordProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPasswordProviderClass) Alloc() AuthorizationPasswordProvider {
	rv := objc.Send[AuthorizationPasswordProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPasswordProviderClass) New() AuthorizationPasswordProvider {
	rv := objc.Send[AuthorizationPasswordProvider](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPasswordProvider) Init() AuthorizationPasswordProvider {
	rv := objc.Send[AuthorizationPasswordProvider](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPasswordProvider) Autorelease() AuthorizationPasswordProvider {
	rv := objc.Send[AuthorizationPasswordProvider](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPasswordProvider creates a new AuthorizationPasswordProvider instance.
func NewAuthorizationPasswordProvider() AuthorizationPasswordProvider {
	return getAuthorizationPasswordProviderClass().New()
}



// Creates a new password authorization request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPasswordProvider/createRequest()
func (a_ AuthorizationPasswordProvider) CreateRequest() IAuthorizationPasswordRequest {
	rv := objc.Send[AuthorizationPasswordRequest](a_.ID, objc.Sel("createRequest"))
	return rv
}



