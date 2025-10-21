// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AuthorizationPlatformPublicKeyCredentialProvider] class.
var (
	AuthorizationPlatformPublicKeyCredentialProviderClass     _AuthorizationPlatformPublicKeyCredentialProviderClass
	AuthorizationPlatformPublicKeyCredentialProviderClassOnce sync.Once
)

func getAuthorizationPlatformPublicKeyCredentialProviderClass() _AuthorizationPlatformPublicKeyCredentialProviderClass {
	AuthorizationPlatformPublicKeyCredentialProviderClassOnce.Do(func() {
		AuthorizationPlatformPublicKeyCredentialProviderClass = _AuthorizationPlatformPublicKeyCredentialProviderClass{objc.GetClass("ASAuthorizationPlatformPublicKeyCredentialProvider")}
	})
	return AuthorizationPlatformPublicKeyCredentialProviderClass
}

type _AuthorizationPlatformPublicKeyCredentialProviderClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPlatformPublicKeyCredentialProvider] class.
type IAuthorizationPlatformPublicKeyCredentialProvider interface {
	objectivec.IObject
}

// A mechanism for providing public key credential requests to an app or service with iCloud Keychain.
//
// The credential provider accesses public-private key pairs stored in iCloud Keychain for registration or authentication with a relying party. Instantiate this object, passing in the relying party identifier for the credentials.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialProvider
type AuthorizationPlatformPublicKeyCredentialProvider struct {
	objectivec.Object
}

// AuthorizationPlatformPublicKeyCredentialProviderFrom constructs a [AuthorizationPlatformPublicKeyCredentialProvider] from an unsafe.Pointer.
//
// A mechanism for providing public key credential requests to an app or service with iCloud Keychain.
func AuthorizationPlatformPublicKeyCredentialProviderFrom(ptr unsafe.Pointer) AuthorizationPlatformPublicKeyCredentialProvider {
	return AuthorizationPlatformPublicKeyCredentialProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPlatformPublicKeyCredentialProviderClass) Alloc() AuthorizationPlatformPublicKeyCredentialProvider {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPlatformPublicKeyCredentialProviderClass) New() AuthorizationPlatformPublicKeyCredentialProvider {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialProvider](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPlatformPublicKeyCredentialProvider) Init() AuthorizationPlatformPublicKeyCredentialProvider {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialProvider](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPlatformPublicKeyCredentialProvider) Autorelease() AuthorizationPlatformPublicKeyCredentialProvider {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialProvider](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPlatformPublicKeyCredentialProvider creates a new AuthorizationPlatformPublicKeyCredentialProvider instance.
func NewAuthorizationPlatformPublicKeyCredentialProvider() AuthorizationPlatformPublicKeyCredentialProvider {
	return getAuthorizationPlatformPublicKeyCredentialProviderClass().New()
}




