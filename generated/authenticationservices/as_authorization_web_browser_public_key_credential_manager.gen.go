// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationWebBrowserPublicKeyCredentialManager] class.
var (
	AuthorizationWebBrowserPublicKeyCredentialManagerClass     _AuthorizationWebBrowserPublicKeyCredentialManagerClass
	AuthorizationWebBrowserPublicKeyCredentialManagerClassOnce sync.Once
)

func getAuthorizationWebBrowserPublicKeyCredentialManagerClass() _AuthorizationWebBrowserPublicKeyCredentialManagerClass {
	AuthorizationWebBrowserPublicKeyCredentialManagerClassOnce.Do(func() {
		AuthorizationWebBrowserPublicKeyCredentialManagerClass = _AuthorizationWebBrowserPublicKeyCredentialManagerClass{objc.GetClass("ASAuthorizationWebBrowserPublicKeyCredentialManager")}
	})
	return AuthorizationWebBrowserPublicKeyCredentialManagerClass
}

type _AuthorizationWebBrowserPublicKeyCredentialManagerClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationWebBrowserPublicKeyCredentialManager] class.
type IAuthorizationWebBrowserPublicKeyCredentialManager interface {
	objectivec.IObject
	AuthorizationStateForPlatformCredentials() unsafe.Pointer
	SetAuthorizationStateForPlatformCredentials(value unsafe.Pointer)
}

// A class that you use to request access to a person’s passkeys in a web browser, and that reports on the access status.


// A class that you use to request access to a person’s passkeys in a web browser, and that reports on the access status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager
type AuthorizationWebBrowserPublicKeyCredentialManager struct {
	objectivec.Object
}

// AuthorizationWebBrowserPublicKeyCredentialManagerFrom constructs a [AuthorizationWebBrowserPublicKeyCredentialManager] from an unsafe.Pointer.
//
// A class that you use to request access to a person’s passkeys in a web browser, and that reports on the access status.
func AuthorizationWebBrowserPublicKeyCredentialManagerFrom(ptr unsafe.Pointer) AuthorizationWebBrowserPublicKeyCredentialManager {
	return AuthorizationWebBrowserPublicKeyCredentialManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationWebBrowserPublicKeyCredentialManagerClass) Alloc() AuthorizationWebBrowserPublicKeyCredentialManager {
	rv := objc.Send[AuthorizationWebBrowserPublicKeyCredentialManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationWebBrowserPublicKeyCredentialManagerClass) New() AuthorizationWebBrowserPublicKeyCredentialManager {
	rv := objc.Send[AuthorizationWebBrowserPublicKeyCredentialManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationWebBrowserPublicKeyCredentialManager) Init() AuthorizationWebBrowserPublicKeyCredentialManager {
	rv := objc.Send[AuthorizationWebBrowserPublicKeyCredentialManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationWebBrowserPublicKeyCredentialManager) Autorelease() AuthorizationWebBrowserPublicKeyCredentialManager {
	rv := objc.Send[AuthorizationWebBrowserPublicKeyCredentialManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationWebBrowserPublicKeyCredentialManager creates a new AuthorizationWebBrowserPublicKeyCredentialManager instance.
func NewAuthorizationWebBrowserPublicKeyCredentialManager() AuthorizationWebBrowserPublicKeyCredentialManager {
	return getAuthorizationWebBrowserPublicKeyCredentialManagerClass().New()
}



// Returns a value that indicates whether the browser app has access to a person’s passkeys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationwebbrowserpublickeycredentialmanager/authorizationstateforplatformcredentials
func (a_ AuthorizationWebBrowserPublicKeyCredentialManager) AuthorizationStateForPlatformCredentials() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("authorizationStateForPlatformCredentials"))
	return rv
}


// Returns a value that indicates whether the browser app has access to a person’s passkeys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationwebbrowserpublickeycredentialmanager/authorizationstateforplatformcredentials
func (a_ AuthorizationWebBrowserPublicKeyCredentialManager) SetAuthorizationStateForPlatformCredentials(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthorizationStateForPlatformCredentials:"), value)
}



