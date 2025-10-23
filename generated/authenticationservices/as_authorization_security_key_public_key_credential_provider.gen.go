// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationSecurityKeyPublicKeyCredentialProvider] class.
var (
	AuthorizationSecurityKeyPublicKeyCredentialProviderClass     _AuthorizationSecurityKeyPublicKeyCredentialProviderClass
	AuthorizationSecurityKeyPublicKeyCredentialProviderClassOnce sync.Once
)

func getAuthorizationSecurityKeyPublicKeyCredentialProviderClass() _AuthorizationSecurityKeyPublicKeyCredentialProviderClass {
	AuthorizationSecurityKeyPublicKeyCredentialProviderClassOnce.Do(func() {
		AuthorizationSecurityKeyPublicKeyCredentialProviderClass = _AuthorizationSecurityKeyPublicKeyCredentialProviderClass{objc.GetClass("ASAuthorizationSecurityKeyPublicKeyCredentialProvider")}
	})
	return AuthorizationSecurityKeyPublicKeyCredentialProviderClass
}

type _AuthorizationSecurityKeyPublicKeyCredentialProviderClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialProvider] class.
type IAuthorizationSecurityKeyPublicKeyCredentialProvider interface {
	objectivec.IObject
	// properties:
	RelyingPartyIdentifier() string /* primitive/slice/pointer. */
	SetRelyingPartyIdentifier(value string /* primitive/slice/pointer. */)
	// methods:
}

// A mechanism for providing public key credential requests to an app or service with a physical security key.
//
// The credential provider accesses public-private key pairs stored on a physical security key for registration or authentication with a relying party. Instantiate this object, passing in the relying party identifier for the credentials.


// A mechanism for providing public key credential requests to an app or service with a physical security key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialProvider
type AuthorizationSecurityKeyPublicKeyCredentialProvider struct {
	objectivec.Object
}

// AuthorizationSecurityKeyPublicKeyCredentialProviderFrom constructs a [AuthorizationSecurityKeyPublicKeyCredentialProvider] from an unsafe.Pointer.
//
// A mechanism for providing public key credential requests to an app or service with a physical security key.
func AuthorizationSecurityKeyPublicKeyCredentialProviderFrom(ptr unsafe.Pointer) AuthorizationSecurityKeyPublicKeyCredentialProvider {
	return AuthorizationSecurityKeyPublicKeyCredentialProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialProviderClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialProvider {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialProviderClass) New() AuthorizationSecurityKeyPublicKeyCredentialProvider {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialProvider](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialProvider) Init() AuthorizationSecurityKeyPublicKeyCredentialProvider {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialProvider](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialProvider) Autorelease() AuthorizationSecurityKeyPublicKeyCredentialProvider {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialProvider](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSecurityKeyPublicKeyCredentialProvider creates a new AuthorizationSecurityKeyPublicKeyCredentialProvider instance.
func NewAuthorizationSecurityKeyPublicKeyCredentialProvider() AuthorizationSecurityKeyPublicKeyCredentialProvider {
	return getAuthorizationSecurityKeyPublicKeyCredentialProviderClass().New()
}



// The domain name of the service to authorize against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsecuritykeypublickeycredentialprovider/relyingpartyidentifier
func (a_ AuthorizationSecurityKeyPublicKeyCredentialProvider) RelyingPartyIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("relyingPartyIdentifier"))
	return rv
}


// The domain name of the service to authorize against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsecuritykeypublickeycredentialprovider/relyingpartyidentifier
func (a_ AuthorizationSecurityKeyPublicKeyCredentialProvider) SetRelyingPartyIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRelyingPartyIdentifier:"), objc.String(value))
}



