// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AuthorizationAppleIDProvider] class.
var (
	AuthorizationAppleIDProviderClass     _AuthorizationAppleIDProviderClass
	AuthorizationAppleIDProviderClassOnce sync.Once
)

func getAuthorizationAppleIDProviderClass() _AuthorizationAppleIDProviderClass {
	AuthorizationAppleIDProviderClassOnce.Do(func() {
		AuthorizationAppleIDProviderClass = _AuthorizationAppleIDProviderClass{objc.GetClass("ASAuthorizationAppleIDProvider")}
	})
	return AuthorizationAppleIDProviderClass
}

type _AuthorizationAppleIDProviderClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationAppleIDProvider] class.
type IAuthorizationAppleIDProvider interface {
	objectivec.IObject
	GetCredentialStateForUserIDCompletion(userID string, completion unsafe.Pointer)
}

// A mechanism for generating requests to authenticate users based on their Apple ID.
//
// You use a provider to create a request ( ), which you then use to initialize a controller ( ) that performs the request: On success, the controller’s delegate receives an authorization ( ) containing a credential ( ) that has an opaque identifier. You can use that identifier to later check the user’s credential state—for example, to see if authorization has been revoked—by calling the method:
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider
type AuthorizationAppleIDProvider struct {
	objectivec.Object
}

// AuthorizationAppleIDProviderFrom constructs a [AuthorizationAppleIDProvider] from an unsafe.Pointer.
//
// A mechanism for generating requests to authenticate users based on their Apple ID.
func AuthorizationAppleIDProviderFrom(ptr unsafe.Pointer) AuthorizationAppleIDProvider {
	return AuthorizationAppleIDProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationAppleIDProviderClass) Alloc() AuthorizationAppleIDProvider {
	rv := objc.Send[AuthorizationAppleIDProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationAppleIDProviderClass) New() AuthorizationAppleIDProvider {
	rv := objc.Send[AuthorizationAppleIDProvider](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationAppleIDProvider) Init() AuthorizationAppleIDProvider {
	rv := objc.Send[AuthorizationAppleIDProvider](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationAppleIDProvider) Autorelease() AuthorizationAppleIDProvider {
	rv := objc.Send[AuthorizationAppleIDProvider](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationAppleIDProvider creates a new AuthorizationAppleIDProvider instance.
func NewAuthorizationAppleIDProvider() AuthorizationAppleIDProvider {
	return getAuthorizationAppleIDProviderClass().New()
}


// Returns the credential state for the given user in a completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider/getCredentialState(forUserID:completion:)
func (a_ AuthorizationAppleIDProvider) GetCredentialStateForUserIDCompletion(userID string, completion unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("getCredentialStateForUserID:completion:"), objc.String(userID), completion)
}



