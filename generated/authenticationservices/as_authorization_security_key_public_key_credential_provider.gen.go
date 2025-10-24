// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationSecurityKeyPublicKeyCredentialProvider */


/* debug [class_header]: Header for ASAuthorizationSecurityKeyPublicKeyCredentialProvider */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationSecurityKeyPublicKeyCredentialProvider */
// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialProvider] class.
type IAuthorizationSecurityKeyPublicKeyCredentialProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationSecurityKeyPublicKeyCredentialProvider */
	// properties:
	RelyingPartyIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationSecurityKeyPublicKeyCredentialProvider */
	// methods:
	CreateCredentialAssertionRequestWithChallenge(challenge objc.IObject /* cross-framework: NSData */) IAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest
	CreateCredentialRegistrationRequestWithChallengeDisplayNameNameUserID(challenge objc.IObject /* cross-framework: NSData */, displayName objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */, userID objc.IObject /* cross-framework: NSData */) IAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationSecurityKeyPublicKeyCredentialProvider */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialProviderClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialProvider {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationSecurityKeyPublicKeyCredentialProvider */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationSecurityKeyPublicKeyCredentialProvider */

// Creates the object with a relying party identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialProvider/init(relyingPartyIdentifier:)
func NewAuthorizationSecurityKeyPublicKeyCredentialProviderWithRelyingPartyIdentifier(relyingPartyIdentifier objc.IObject /* cross-framework: NSString */) AuthorizationSecurityKeyPublicKeyCredentialProvider {
	instance := getAuthorizationSecurityKeyPublicKeyCredentialProviderClass().Alloc()
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialProvider](instance.ID, objc.Sel("initWithRelyingPartyIdentifier:"), relyingPartyIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationSecurityKeyPublicKeyCredentialProviderWithRelyingPartyIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationSecurityKeyPublicKeyCredentialProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationSecurityKeyPublicKeyCredentialProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationSecurityKeyPublicKeyCredentialProvider */

// Creates an assertion request with a challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialProvider/createCredentialAssertionRequest(challenge:)
func (a_ AuthorizationSecurityKeyPublicKeyCredentialProvider) CreateCredentialAssertionRequestWithChallenge(challenge objc.IObject /* cross-framework: NSData */) IAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest](a_.ID, objc.Sel("createCredentialAssertionRequestWithChallenge:"), challenge)
	return rv
}/* debug [instance_methods/method]: CreateCredentialAssertionRequestWithChallenge */


// Creates an assertion request with a challenge, display name, and user ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialProvider/createCredentialRegistrationRequest(challenge:displayName:name:userID:)
func (a_ AuthorizationSecurityKeyPublicKeyCredentialProvider) CreateCredentialRegistrationRequestWithChallengeDisplayNameNameUserID(challenge objc.IObject /* cross-framework: NSData */, displayName objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */, userID objc.IObject /* cross-framework: NSData */) IAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest](a_.ID, objc.Sel("createCredentialRegistrationRequestWithChallenge:displayName:name:userID:"), challenge, displayName, name, userID)
	return rv
}/* debug [instance_methods/method]: CreateCredentialRegistrationRequestWithChallengeDisplayNameNameUserID */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationSecurityKeyPublicKeyCredentialProvider */

// The domain name of the service to authorize against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialProvider/relyingPartyIdentifier
func (a_ AuthorizationSecurityKeyPublicKeyCredentialProvider) RelyingPartyIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("relyingPartyIdentifier"))
	return rv
}/* debug [instance_properties/getter]: relyingPartyIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationSecurityKeyPublicKeyCredentialProvider */


