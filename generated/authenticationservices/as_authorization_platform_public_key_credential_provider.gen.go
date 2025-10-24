// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPlatformPublicKeyCredentialProvider */


/* debug [class_header]: Header for ASAuthorizationPlatformPublicKeyCredentialProvider */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPlatformPublicKeyCredentialProvider */
// An interface definition for the [AuthorizationPlatformPublicKeyCredentialProvider] class.
type IAuthorizationPlatformPublicKeyCredentialProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPlatformPublicKeyCredentialProvider */
	// properties:
	RelyingPartyIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPlatformPublicKeyCredentialProvider */
	// methods:
	CreateCredentialAssertionRequestWithChallenge(challenge objc.IObject /* cross-framework: NSData */) IAuthorizationPlatformPublicKeyCredentialAssertionRequest
	CreateCredentialRegistrationRequestWithChallengeNameUserID(challenge objc.IObject /* cross-framework: NSData */, name objc.IObject /* cross-framework: NSString */, userID objc.IObject /* cross-framework: NSData */) IAuthorizationPlatformPublicKeyCredentialRegistrationRequest
	CreateCredentialRegistrationRequestWithChallengeNameUserIDRequestStyle(challenge objc.IObject /* cross-framework: NSData */, name objc.IObject /* cross-framework: NSString */, userID objc.IObject /* cross-framework: NSData */, requestStyle AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle) IAuthorizationPlatformPublicKeyCredentialRegistrationRequest
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPlatformPublicKeyCredentialProvider */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPlatformPublicKeyCredentialProviderClass) Alloc() AuthorizationPlatformPublicKeyCredentialProvider {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPlatformPublicKeyCredentialProvider */
// A mechanism for providing public key credential requests to an app or service with iCloud Keychain.
//
// The credential provider accesses public-private key pairs stored in iCloud Keychain for registration or authentication with a relying party. Instantiate this object, passing in the relying party identifier for the credentials.


// A mechanism for providing public key credential requests to an app or service with iCloud Keychain.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPlatformPublicKeyCredentialProvider */

// Creates the object with a relying party identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialProvider/init(relyingPartyIdentifier:)
func NewAuthorizationPlatformPublicKeyCredentialProviderWithRelyingPartyIdentifier(relyingPartyIdentifier objc.IObject /* cross-framework: NSString */) AuthorizationPlatformPublicKeyCredentialProvider {
	instance := getAuthorizationPlatformPublicKeyCredentialProviderClass().Alloc()
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialProvider](instance.ID, objc.Sel("initWithRelyingPartyIdentifier:"), relyingPartyIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationPlatformPublicKeyCredentialProviderWithRelyingPartyIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPlatformPublicKeyCredentialProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPlatformPublicKeyCredentialProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPlatformPublicKeyCredentialProvider */

// Creates an assertion request with a challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialProvider/createCredentialAssertionRequest(challenge:)
func (a_ AuthorizationPlatformPublicKeyCredentialProvider) CreateCredentialAssertionRequestWithChallenge(challenge objc.IObject /* cross-framework: NSData */) IAuthorizationPlatformPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialAssertionRequest](a_.ID, objc.Sel("createCredentialAssertionRequestWithChallenge:"), challenge)
	return rv
}/* debug [instance_methods/method]: CreateCredentialAssertionRequestWithChallenge */


// Creates an assertion request with a challenge, name, and user ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialProvider/createCredentialRegistrationRequest(challenge:name:userID:)
func (a_ AuthorizationPlatformPublicKeyCredentialProvider) CreateCredentialRegistrationRequestWithChallengeNameUserID(challenge objc.IObject /* cross-framework: NSData */, name objc.IObject /* cross-framework: NSString */, userID objc.IObject /* cross-framework: NSData */) IAuthorizationPlatformPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistrationRequest](a_.ID, objc.Sel("createCredentialRegistrationRequestWithChallenge:name:userID:"), challenge, name, userID)
	return rv
}/* debug [instance_methods/method]: CreateCredentialRegistrationRequestWithChallengeNameUserID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialProvider/createCredentialRegistrationRequest(challenge:name:userID:requestStyle:)
func (a_ AuthorizationPlatformPublicKeyCredentialProvider) CreateCredentialRegistrationRequestWithChallengeNameUserIDRequestStyle(challenge objc.IObject /* cross-framework: NSData */, name objc.IObject /* cross-framework: NSString */, userID objc.IObject /* cross-framework: NSData */, requestStyle AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle) IAuthorizationPlatformPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistrationRequest](a_.ID, objc.Sel("createCredentialRegistrationRequestWithChallenge:name:userID:requestStyle:"), challenge, name, userID, requestStyle)
	return rv
}/* debug [instance_methods/method]: CreateCredentialRegistrationRequestWithChallengeNameUserIDRequestStyle */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPlatformPublicKeyCredentialProvider */

// The domain name of the service to register or authorize against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialProvider/relyingPartyIdentifier
func (a_ AuthorizationPlatformPublicKeyCredentialProvider) RelyingPartyIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("relyingPartyIdentifier"))
	return rv
}/* debug [instance_properties/getter]: relyingPartyIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPlatformPublicKeyCredentialProvider */


