// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationProviderExtensionLoginConfiguration */


/* debug [class_header]: Header for ASAuthorizationProviderExtensionLoginConfiguration */
// The class instance for the [AuthorizationProviderExtensionLoginConfiguration] class.
var (
	AuthorizationProviderExtensionLoginConfigurationClass     _AuthorizationProviderExtensionLoginConfigurationClass
	AuthorizationProviderExtensionLoginConfigurationClassOnce sync.Once
)

func getAuthorizationProviderExtensionLoginConfigurationClass() _AuthorizationProviderExtensionLoginConfigurationClass {
	AuthorizationProviderExtensionLoginConfigurationClassOnce.Do(func() {
		AuthorizationProviderExtensionLoginConfigurationClass = _AuthorizationProviderExtensionLoginConfigurationClass{objc.GetClass("ASAuthorizationProviderExtensionLoginConfiguration")}
	})
	return AuthorizationProviderExtensionLoginConfigurationClass
}

type _AuthorizationProviderExtensionLoginConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationProviderExtensionLoginConfiguration */
// An interface definition for the [AuthorizationProviderExtensionLoginConfiguration] class.
type IAuthorizationProviderExtensionLoginConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationProviderExtensionLoginConfiguration */
	// properties:
	AccountDisplayName() objc.IObject /* cross-framework: NSString */
	SetAccountDisplayName(value objc.IObject /* cross-framework: NSString */)
	AdditionalAuthorizationScopes() objc.IObject /* cross-framework: NSString */
	SetAdditionalAuthorizationScopes(value objc.IObject /* cross-framework: NSString */)
	AdditionalScopes() objc.IObject /* cross-framework: NSString */
	SetAdditionalScopes(value objc.IObject /* cross-framework: NSString */)
	Audience() objc.IObject /* cross-framework: NSString */
	SetAudience(value objc.IObject /* cross-framework: NSString */)
	ClientID() objc.IObject /* cross-framework: NSString */
	CustomFederationUserPreauthenticationRequestValues() []foundation.URLQueryItem
	SetCustomFederationUserPreauthenticationRequestValues(value []foundation.URLQueryItem)
	CustomKeyExchangeRequestValues() []foundation.URLQueryItem
	SetCustomKeyExchangeRequestValues(value []foundation.URLQueryItem)
	CustomKeyRequestValues() []foundation.URLQueryItem
	SetCustomKeyRequestValues(value []foundation.URLQueryItem)
	CustomLoginRequestValues() []foundation.URLQueryItem
	SetCustomLoginRequestValues(value []foundation.URLQueryItem)
	CustomNonceRequestValues() []foundation.URLQueryItem
	SetCustomNonceRequestValues(value []foundation.URLQueryItem)
	CustomRefreshRequestValues() []foundation.URLQueryItem
	SetCustomRefreshRequestValues(value []foundation.URLQueryItem)
	CustomRequestJWTParameterName() objc.IObject /* cross-framework: NSString */
	SetCustomRequestJWTParameterName(value objc.IObject /* cross-framework: NSString */)
	DeviceContext() objc.IObject /* cross-framework: NSData */
	SetDeviceContext(value objc.IObject /* cross-framework: NSData */)
	FederationMEXURL() objc.IObject /* cross-framework: NSURL */
	SetFederationMEXURL(value objc.IObject /* cross-framework: NSURL */)
	FederationMEXURLKeypath() objc.IObject /* cross-framework: NSString */
	SetFederationMEXURLKeypath(value objc.IObject /* cross-framework: NSString */)
	FederationPredicate() objc.IObject /* cross-framework: NSString */
	SetFederationPredicate(value objc.IObject /* cross-framework: NSString */)
	FederationRequestURN() objc.IObject /* cross-framework: NSString */
	SetFederationRequestURN(value objc.IObject /* cross-framework: NSString */)
	FederationType() AuthorizationProviderExtensionFederationType
	SetFederationType(value AuthorizationProviderExtensionFederationType)
	FederationUserPreauthenticationURL() objc.IObject /* cross-framework: NSURL */
	SetFederationUserPreauthenticationURL(value objc.IObject /* cross-framework: NSURL */)
	GroupRequestClaimName() objc.IObject /* cross-framework: NSString */
	SetGroupRequestClaimName(value objc.IObject /* cross-framework: NSString */)
	GroupResponseClaimName() objc.IObject /* cross-framework: NSString */
	SetGroupResponseClaimName(value objc.IObject /* cross-framework: NSString */)
	HpkeAuthPublicKey() unsafe.Pointer
	SetHpkeAuthPublicKey(value unsafe.Pointer)
	HpkePreSharedKey() objc.IObject /* cross-framework: NSData */
	SetHpkePreSharedKey(value objc.IObject /* cross-framework: NSData */)
	HpkePreSharedKeyID() objc.IObject /* cross-framework: NSData */
	SetHpkePreSharedKeyID(value objc.IObject /* cross-framework: NSData */)
	IncludePreviousRefreshTokenInLoginRequest() bool
	SetIncludePreviousRefreshTokenInLoginRequest(value bool)
	InvalidCredentialPredicate() objc.IObject /* cross-framework: NSString */
	SetInvalidCredentialPredicate(value objc.IObject /* cross-framework: NSString */)
	Issuer() objc.IObject /* cross-framework: NSString */
	JwksEndpointURL() objc.IObject /* cross-framework: NSURL */
	SetJwksEndpointURL(value objc.IObject /* cross-framework: NSURL */)
	JwksTrustedRootCertificates() objc.IObject /* cross-framework: NSArray */
	SetJwksTrustedRootCertificates(value objc.IObject /* cross-framework: NSArray */)
	KerberosTicketMappings() []AuthorizationProviderExtensionKerberosMapping
	SetKerberosTicketMappings(value []AuthorizationProviderExtensionKerberosMapping)
	KeyEndpointURL() objc.IObject /* cross-framework: NSURL */
	SetKeyEndpointURL(value objc.IObject /* cross-framework: NSURL */)
	LoginRequestEncryptionAlgorithm() AuthorizationProviderExtensionEncryptionAlgorithm /* typedef */
	SetLoginRequestEncryptionAlgorithm(value AuthorizationProviderExtensionEncryptionAlgorithm /* typedef */)
	LoginRequestEncryptionAPVPrefix() objc.IObject /* cross-framework: NSData */
	SetLoginRequestEncryptionAPVPrefix(value objc.IObject /* cross-framework: NSData */)
	LoginRequestEncryptionPublicKey() unsafe.Pointer
	SetLoginRequestEncryptionPublicKey(value unsafe.Pointer)
	LoginRequestHPKEPreSharedKey() objc.IObject /* cross-framework: NSData */
	SetLoginRequestHPKEPreSharedKey(value objc.IObject /* cross-framework: NSData */)
	LoginRequestHPKEPreSharedKeyID() objc.IObject /* cross-framework: NSData */
	SetLoginRequestHPKEPreSharedKeyID(value objc.IObject /* cross-framework: NSData */)
	NonceEndpointURL() objc.IObject /* cross-framework: NSURL */
	SetNonceEndpointURL(value objc.IObject /* cross-framework: NSURL */)
	NonceResponseKeypath() objc.IObject /* cross-framework: NSString */
	SetNonceResponseKeypath(value objc.IObject /* cross-framework: NSString */)
	PreviousRefreshTokenClaimName() objc.IObject /* cross-framework: NSString */
	SetPreviousRefreshTokenClaimName(value objc.IObject /* cross-framework: NSString */)
	RefreshEndpointURL() objc.IObject /* cross-framework: NSURL */
	SetRefreshEndpointURL(value objc.IObject /* cross-framework: NSURL */)
	ServerNonceClaimName() objc.IObject /* cross-framework: NSString */
	SetServerNonceClaimName(value objc.IObject /* cross-framework: NSString */)
	TokenEndpointURL() objc.IObject /* cross-framework: NSURL */
	SetTokenEndpointURL(value objc.IObject /* cross-framework: NSURL */)
	UniqueIdentifierClaimName() objc.IObject /* cross-framework: NSString */
	SetUniqueIdentifierClaimName(value objc.IObject /* cross-framework: NSString */)
	UserSecureEnclaveKeyBiometricPolicy() AuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy
	SetUserSecureEnclaveKeyBiometricPolicy(value AuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationProviderExtensionLoginConfiguration */
	// methods:
	SetCustomAssertionRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomAssertionRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomKeyExchangeRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomKeyExchangeRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomKeyRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomKeyRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomLoginRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomLoginRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomRefreshRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomRefreshRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationProviderExtensionLoginConfiguration */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionLoginConfigurationClass) Alloc() AuthorizationProviderExtensionLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionLoginConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationProviderExtensionLoginConfigurationClass) New() AuthorizationProviderExtensionLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionLoginConfiguration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationProviderExtensionLoginConfiguration) Init() AuthorizationProviderExtensionLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionLoginConfiguration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationProviderExtensionLoginConfiguration) Autorelease() AuthorizationProviderExtensionLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionLoginConfiguration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationProviderExtensionLoginConfiguration creates a new AuthorizationProviderExtensionLoginConfiguration instance.
func NewAuthorizationProviderExtensionLoginConfiguration() AuthorizationProviderExtensionLoginConfiguration {
	return getAuthorizationProviderExtensionLoginConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationProviderExtensionLoginConfiguration */
// An interface for configuring platform single sign-on.
//
// This class provides login configuration information for platform single sign-on.


// An interface for configuring platform single sign-on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration
type AuthorizationProviderExtensionLoginConfiguration struct {
	objectivec.Object
}

// AuthorizationProviderExtensionLoginConfigurationFrom constructs a [AuthorizationProviderExtensionLoginConfiguration] from an unsafe.Pointer.
//
// An interface for configuring platform single sign-on.
func AuthorizationProviderExtensionLoginConfigurationFrom(ptr unsafe.Pointer) AuthorizationProviderExtensionLoginConfiguration {
	return AuthorizationProviderExtensionLoginConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationProviderExtensionLoginConfiguration */

// Creates a configuration with the required values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/init(clientID:issuer:tokenEndpointURL:jwksEndpointURL:audience:)
func NewAuthorizationProviderExtensionLoginConfigurationWithClientIDIssuerTokenEndpointURLJwksEndpointURLAudience(clientID objc.IObject /* cross-framework: NSString */, issuer objc.IObject /* cross-framework: NSString */, tokenEndpointURL objc.IObject /* cross-framework: NSURL */, jwksEndpointURL objc.IObject /* cross-framework: NSURL */, audience objc.IObject /* cross-framework: NSString */) AuthorizationProviderExtensionLoginConfiguration {
	instance := getAuthorizationProviderExtensionLoginConfigurationClass().Alloc()
	rv := objc.Send[AuthorizationProviderExtensionLoginConfiguration](instance.ID, objc.Sel("initWithClientID:issuer:tokenEndpointURL:jwksEndpointURL:audience:"), clientID, issuer, tokenEndpointURL, jwksEndpointURL, audience)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationProviderExtensionLoginConfigurationWithClientIDIssuerTokenEndpointURLJwksEndpointURLAudience */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationProviderExtensionLoginConfiguration */

// Creates a login configuration using the OpenID configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/configuration(openIDConfigurationURL:clientID:issuer:completion:)
func (ac _AuthorizationProviderExtensionLoginConfigurationClass) ConfigurationWithOpenIDConfigurationURLClientIDIssuerCompletion(openIDConfigurationURL objc.IObject /* cross-framework: NSURL */, clientID objc.IObject /* cross-framework: NSString */, issuer objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("configurationWithOpenIDConfigurationURL:clientID:issuer:completion:"), openIDConfigurationURL, clientID, issuer, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithOpenIDConfigurationURLClientIDIssuerCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationProviderExtensionLoginConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationProviderExtensionLoginConfiguration */

// Adds the custom claims to the embedded assertion request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/setCustomAssertionRequestBodyClaims(_:)
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomAssertionRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomAssertionRequestBodyClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomAssertionRequestBodyClaimsReturningError */


// Adds the custom claims to the embedded assertion request header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/setCustomAssertionRequestHeaderClaims(_:)
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomAssertionRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomAssertionRequestHeaderClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomAssertionRequestHeaderClaimsReturningError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/setCustomKeyExchangeRequestBodyClaims(_:)
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomKeyExchangeRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomKeyExchangeRequestBodyClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomKeyExchangeRequestBodyClaimsReturningError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/setCustomKeyExchangeRequestHeaderClaims(_:)
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomKeyExchangeRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomKeyExchangeRequestHeaderClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomKeyExchangeRequestHeaderClaimsReturningError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/setCustomKeyRequestBodyClaims(_:)
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomKeyRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomKeyRequestBodyClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomKeyRequestBodyClaimsReturningError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/setCustomKeyRequestHeaderClaims(_:)
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomKeyRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomKeyRequestHeaderClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomKeyRequestHeaderClaimsReturningError */


// Adds the custom claims to the login request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/setCustomLoginRequestBodyClaims(_:)
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomLoginRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomLoginRequestBodyClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomLoginRequestBodyClaimsReturningError */


// Adds the custom claims to the login request header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/setCustomLoginRequestHeaderClaims(_:)
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomLoginRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomLoginRequestHeaderClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomLoginRequestHeaderClaimsReturningError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/setCustomRefreshRequestBodyClaims(_:)
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomRefreshRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomRefreshRequestBodyClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomRefreshRequestBodyClaimsReturningError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/setCustomRefreshRequestHeaderClaims(_:)
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomRefreshRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomRefreshRequestHeaderClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomRefreshRequestHeaderClaimsReturningError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationProviderExtensionLoginConfiguration */

// The display name for the account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/accountDisplayName
func (a_ AuthorizationProviderExtensionLoginConfiguration) AccountDisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("accountDisplayName"))
	return rv
}/* debug [instance_properties/getter]: accountDisplayName */


// The display name for the account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/accountDisplayName
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetAccountDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccountDisplayName:"), value)
}/* debug [instance_properties/setter]: accountDisplayName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/additionalAuthorizationScopes
func (a_ AuthorizationProviderExtensionLoginConfiguration) AdditionalAuthorizationScopes() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("additionalAuthorizationScopes"))
	return rv
}/* debug [instance_properties/getter]: additionalAuthorizationScopes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/additionalAuthorizationScopes
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetAdditionalAuthorizationScopes(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdditionalAuthorizationScopes:"), value)
}/* debug [instance_properties/setter]: additionalAuthorizationScopes */


// A set of extra scopes to add to the base for the authentication request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/additionalScopes
func (a_ AuthorizationProviderExtensionLoginConfiguration) AdditionalScopes() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("additionalScopes"))
	return rv
}/* debug [instance_properties/getter]: additionalScopes */


// A set of extra scopes to add to the base for the authentication request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/additionalScopes
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetAdditionalScopes(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdditionalScopes:"), value)
}/* debug [instance_properties/setter]: additionalScopes */


// The audience for validation and requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/audience
func (a_ AuthorizationProviderExtensionLoginConfiguration) Audience() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("audience"))
	return rv
}/* debug [instance_properties/getter]: audience */


// The audience for validation and requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/audience
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetAudience(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudience:"), value)
}/* debug [instance_properties/setter]: audience */


// The identifier for the client at the identity provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/clientID
func (a_ AuthorizationProviderExtensionLoginConfiguration) ClientID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("clientID"))
	return rv
}/* debug [instance_properties/getter]: clientID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customFederationUserPreauthenticationRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) CustomFederationUserPreauthenticationRequestValues() []foundation.URLQueryItem {
	rv := objc.Send[[]foundation.URLQueryItem](a_.ID, objc.Sel("customFederationUserPreauthenticationRequestValues"))
	return rv
}/* debug [instance_properties/getter]: customFederationUserPreauthenticationRequestValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customFederationUserPreauthenticationRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomFederationUserPreauthenticationRequestValues(value []foundation.URLQueryItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomFederationUserPreauthenticationRequestValues:"), nsArray)
}/* debug [instance_properties/setter]: customFederationUserPreauthenticationRequestValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customKeyExchangeRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) CustomKeyExchangeRequestValues() []foundation.URLQueryItem {
	rv := objc.Send[[]foundation.URLQueryItem](a_.ID, objc.Sel("customKeyExchangeRequestValues"))
	return rv
}/* debug [instance_properties/getter]: customKeyExchangeRequestValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customKeyExchangeRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomKeyExchangeRequestValues(value []foundation.URLQueryItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomKeyExchangeRequestValues:"), nsArray)
}/* debug [instance_properties/setter]: customKeyExchangeRequestValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customKeyRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) CustomKeyRequestValues() []foundation.URLQueryItem {
	rv := objc.Send[[]foundation.URLQueryItem](a_.ID, objc.Sel("customKeyRequestValues"))
	return rv
}/* debug [instance_properties/getter]: customKeyRequestValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customKeyRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomKeyRequestValues(value []foundation.URLQueryItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomKeyRequestValues:"), nsArray)
}/* debug [instance_properties/setter]: customKeyRequestValues */


// Provider-supplied values to add to the login POST request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customLoginRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) CustomLoginRequestValues() []foundation.URLQueryItem {
	rv := objc.Send[[]foundation.URLQueryItem](a_.ID, objc.Sel("customLoginRequestValues"))
	return rv
}/* debug [instance_properties/getter]: customLoginRequestValues */


// Provider-supplied values to add to the login POST request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customLoginRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomLoginRequestValues(value []foundation.URLQueryItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomLoginRequestValues:"), nsArray)
}/* debug [instance_properties/setter]: customLoginRequestValues */


// Custom values to add to the server nonce POST request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customNonceRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) CustomNonceRequestValues() []foundation.URLQueryItem {
	rv := objc.Send[[]foundation.URLQueryItem](a_.ID, objc.Sel("customNonceRequestValues"))
	return rv
}/* debug [instance_properties/getter]: customNonceRequestValues */


// Custom values to add to the server nonce POST request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customNonceRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomNonceRequestValues(value []foundation.URLQueryItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomNonceRequestValues:"), nsArray)
}/* debug [instance_properties/setter]: customNonceRequestValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customRefreshRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) CustomRefreshRequestValues() []foundation.URLQueryItem {
	rv := objc.Send[[]foundation.URLQueryItem](a_.ID, objc.Sel("customRefreshRequestValues"))
	return rv
}/* debug [instance_properties/getter]: customRefreshRequestValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customRefreshRequestValues
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomRefreshRequestValues(value []foundation.URLQueryItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomRefreshRequestValues:"), nsArray)
}/* debug [instance_properties/setter]: customRefreshRequestValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customRequestJWTParameterName
func (a_ AuthorizationProviderExtensionLoginConfiguration) CustomRequestJWTParameterName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("customRequestJWTParameterName"))
	return rv
}/* debug [instance_properties/getter]: customRequestJWTParameterName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/customRequestJWTParameterName
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetCustomRequestJWTParameterName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomRequestJWTParameterName:"), value)
}/* debug [instance_properties/setter]: customRequestJWTParameterName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/deviceContext
func (a_ AuthorizationProviderExtensionLoginConfiguration) DeviceContext() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("deviceContext"))
	return rv
}/* debug [instance_properties/getter]: deviceContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/deviceContext
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetDeviceContext(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDeviceContext:"), value)
}/* debug [instance_properties/setter]: deviceContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationMEXURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) FederationMEXURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("federationMEXURL"))
	return rv
}/* debug [instance_properties/getter]: federationMEXURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationMEXURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetFederationMEXURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFederationMEXURL:"), value)
}/* debug [instance_properties/setter]: federationMEXURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationMEXURLKeypath
func (a_ AuthorizationProviderExtensionLoginConfiguration) FederationMEXURLKeypath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("federationMEXURLKeypath"))
	return rv
}/* debug [instance_properties/getter]: federationMEXURLKeypath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationMEXURLKeypath
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetFederationMEXURLKeypath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFederationMEXURLKeypath:"), value)
}/* debug [instance_properties/setter]: federationMEXURLKeypath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationPredicate
func (a_ AuthorizationProviderExtensionLoginConfiguration) FederationPredicate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("federationPredicate"))
	return rv
}/* debug [instance_properties/getter]: federationPredicate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationPredicate
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetFederationPredicate(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFederationPredicate:"), value)
}/* debug [instance_properties/setter]: federationPredicate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationRequestURN
func (a_ AuthorizationProviderExtensionLoginConfiguration) FederationRequestURN() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("federationRequestURN"))
	return rv
}/* debug [instance_properties/getter]: federationRequestURN */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationRequestURN
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetFederationRequestURN(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFederationRequestURN:"), value)
}/* debug [instance_properties/setter]: federationRequestURN */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationType-swift.property
func (a_ AuthorizationProviderExtensionLoginConfiguration) FederationType() AuthorizationProviderExtensionFederationType {
	rv := objc.Send[AuthorizationProviderExtensionFederationType](a_.ID, objc.Sel("federationType"))
	return rv
}/* debug [instance_properties/getter]: federationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationType-swift.property
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetFederationType(value AuthorizationProviderExtensionFederationType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFederationType:"), value)
}/* debug [instance_properties/setter]: federationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationUserPreauthenticationURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) FederationUserPreauthenticationURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("federationUserPreauthenticationURL"))
	return rv
}/* debug [instance_properties/getter]: federationUserPreauthenticationURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/federationUserPreauthenticationURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetFederationUserPreauthenticationURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFederationUserPreauthenticationURL:"), value)
}/* debug [instance_properties/setter]: federationUserPreauthenticationURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/groupRequestClaimName
func (a_ AuthorizationProviderExtensionLoginConfiguration) GroupRequestClaimName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("groupRequestClaimName"))
	return rv
}/* debug [instance_properties/getter]: groupRequestClaimName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/groupRequestClaimName
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetGroupRequestClaimName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGroupRequestClaimName:"), value)
}/* debug [instance_properties/setter]: groupRequestClaimName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/groupResponseClaimName
func (a_ AuthorizationProviderExtensionLoginConfiguration) GroupResponseClaimName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("groupResponseClaimName"))
	return rv
}/* debug [instance_properties/getter]: groupResponseClaimName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/groupResponseClaimName
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetGroupResponseClaimName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGroupResponseClaimName:"), value)
}/* debug [instance_properties/setter]: groupResponseClaimName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/hpkeAuthPublicKey
func (a_ AuthorizationProviderExtensionLoginConfiguration) HpkeAuthPublicKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("hpkeAuthPublicKey"))
	return rv
}/* debug [instance_properties/getter]: hpkeAuthPublicKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/hpkeAuthPublicKey
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetHpkeAuthPublicKey(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHpkeAuthPublicKey:"), value)
}/* debug [instance_properties/setter]: hpkeAuthPublicKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/hpkePreSharedKey
func (a_ AuthorizationProviderExtensionLoginConfiguration) HpkePreSharedKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("hpkePreSharedKey"))
	return rv
}/* debug [instance_properties/getter]: hpkePreSharedKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/hpkePreSharedKey
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetHpkePreSharedKey(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHpkePreSharedKey:"), value)
}/* debug [instance_properties/setter]: hpkePreSharedKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/hpkePreSharedKeyID
func (a_ AuthorizationProviderExtensionLoginConfiguration) HpkePreSharedKeyID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("hpkePreSharedKeyID"))
	return rv
}/* debug [instance_properties/getter]: hpkePreSharedKeyID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/hpkePreSharedKeyID
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetHpkePreSharedKeyID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHpkePreSharedKeyID:"), value)
}/* debug [instance_properties/setter]: hpkePreSharedKeyID */


// A Boolean value that indicates whether to include the previous refresh token in the authentation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/includePreviousRefreshTokenInLoginRequest
func (a_ AuthorizationProviderExtensionLoginConfiguration) IncludePreviousRefreshTokenInLoginRequest() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("includePreviousRefreshTokenInLoginRequest"))
	return rv
}/* debug [instance_properties/getter]: includePreviousRefreshTokenInLoginRequest */


// A Boolean value that indicates whether to include the previous refresh token in the authentation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/includePreviousRefreshTokenInLoginRequest
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetIncludePreviousRefreshTokenInLoginRequest(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIncludePreviousRefreshTokenInLoginRequest:"), value)
}/* debug [instance_properties/setter]: includePreviousRefreshTokenInLoginRequest */


// The predicate string that identifies invalid credential errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/invalidCredentialPredicate
func (a_ AuthorizationProviderExtensionLoginConfiguration) InvalidCredentialPredicate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("invalidCredentialPredicate"))
	return rv
}/* debug [instance_properties/getter]: invalidCredentialPredicate */


// The predicate string that identifies invalid credential errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/invalidCredentialPredicate
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetInvalidCredentialPredicate(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInvalidCredentialPredicate:"), value)
}/* debug [instance_properties/setter]: invalidCredentialPredicate */


// The issuer of the identity token that the identity provider returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/issuer
func (a_ AuthorizationProviderExtensionLoginConfiguration) Issuer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("issuer"))
	return rv
}/* debug [instance_properties/getter]: issuer */


// The JSON Web Key Set endpoint URL for keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/jwksEndpointURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) JwksEndpointURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("jwksEndpointURL"))
	return rv
}/* debug [instance_properties/getter]: jwksEndpointURL */


// The JSON Web Key Set endpoint URL for keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/jwksEndpointURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetJwksEndpointURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setJwksEndpointURL:"), value)
}/* debug [instance_properties/setter]: jwksEndpointURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/jwksTrustedRootCertificates-605jm
func (a_ AuthorizationProviderExtensionLoginConfiguration) JwksTrustedRootCertificates() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("jwksTrustedRootCertificates"))
	return rv
}/* debug [instance_properties/getter]: jwksTrustedRootCertificates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/jwksTrustedRootCertificates-605jm
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetJwksTrustedRootCertificates(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setJwksTrustedRootCertificates:"), value)
}/* debug [instance_properties/setter]: jwksTrustedRootCertificates */


// The set of ticket mappings the system uses to import Kerberos tickets from the single sign-on token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/kerberosTicketMappings
func (a_ AuthorizationProviderExtensionLoginConfiguration) KerberosTicketMappings() []AuthorizationProviderExtensionKerberosMapping {
	rv := objc.Send[[]AuthorizationProviderExtensionKerberosMapping](a_.ID, objc.Sel("kerberosTicketMappings"))
	return rv
}/* debug [instance_properties/getter]: kerberosTicketMappings */


// The set of ticket mappings the system uses to import Kerberos tickets from the single sign-on token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/kerberosTicketMappings
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetKerberosTicketMappings(value []AuthorizationProviderExtensionKerberosMapping) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setKerberosTicketMappings:"), nsArray)
}/* debug [instance_properties/setter]: kerberosTicketMappings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/keyEndpointURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) KeyEndpointURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("keyEndpointURL"))
	return rv
}/* debug [instance_properties/getter]: keyEndpointURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/keyEndpointURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetKeyEndpointURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKeyEndpointURL:"), value)
}/* debug [instance_properties/setter]: keyEndpointURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/loginRequestEncryptionAlgorithm
func (a_ AuthorizationProviderExtensionLoginConfiguration) LoginRequestEncryptionAlgorithm() AuthorizationProviderExtensionEncryptionAlgorithm /* typedef */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("loginRequestEncryptionAlgorithm"))
	return rv
}/* debug [instance_properties/getter]: loginRequestEncryptionAlgorithm */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/loginRequestEncryptionAlgorithm
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetLoginRequestEncryptionAlgorithm(value AuthorizationProviderExtensionEncryptionAlgorithm /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginRequestEncryptionAlgorithm:"), value)
}/* debug [instance_properties/setter]: loginRequestEncryptionAlgorithm */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/loginRequestEncryptionAPVPrefix
func (a_ AuthorizationProviderExtensionLoginConfiguration) LoginRequestEncryptionAPVPrefix() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("loginRequestEncryptionAPVPrefix"))
	return rv
}/* debug [instance_properties/getter]: loginRequestEncryptionAPVPrefix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/loginRequestEncryptionAPVPrefix
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetLoginRequestEncryptionAPVPrefix(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginRequestEncryptionAPVPrefix:"), value)
}/* debug [instance_properties/setter]: loginRequestEncryptionAPVPrefix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/loginRequestEncryptionPublicKey
func (a_ AuthorizationProviderExtensionLoginConfiguration) LoginRequestEncryptionPublicKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("loginRequestEncryptionPublicKey"))
	return rv
}/* debug [instance_properties/getter]: loginRequestEncryptionPublicKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/loginRequestEncryptionPublicKey
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetLoginRequestEncryptionPublicKey(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginRequestEncryptionPublicKey:"), value)
}/* debug [instance_properties/setter]: loginRequestEncryptionPublicKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/loginRequestHPKEPreSharedKey
func (a_ AuthorizationProviderExtensionLoginConfiguration) LoginRequestHPKEPreSharedKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("loginRequestHPKEPreSharedKey"))
	return rv
}/* debug [instance_properties/getter]: loginRequestHPKEPreSharedKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/loginRequestHPKEPreSharedKey
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetLoginRequestHPKEPreSharedKey(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginRequestHPKEPreSharedKey:"), value)
}/* debug [instance_properties/setter]: loginRequestHPKEPreSharedKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/loginRequestHPKEPreSharedKeyID
func (a_ AuthorizationProviderExtensionLoginConfiguration) LoginRequestHPKEPreSharedKeyID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("loginRequestHPKEPreSharedKeyID"))
	return rv
}/* debug [instance_properties/getter]: loginRequestHPKEPreSharedKeyID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/loginRequestHPKEPreSharedKeyID
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetLoginRequestHPKEPreSharedKeyID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginRequestHPKEPreSharedKeyID:"), value)
}/* debug [instance_properties/setter]: loginRequestHPKEPreSharedKeyID */


// The URL to retrieve a one-time use value from the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/nonceEndpointURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) NonceEndpointURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("nonceEndpointURL"))
	return rv
}/* debug [instance_properties/getter]: nonceEndpointURL */


// The URL to retrieve a one-time use value from the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/nonceEndpointURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetNonceEndpointURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNonceEndpointURL:"), value)
}/* debug [instance_properties/setter]: nonceEndpointURL */


// The keypath in the response that contains the one-time use value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/nonceResponseKeypath
func (a_ AuthorizationProviderExtensionLoginConfiguration) NonceResponseKeypath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("nonceResponseKeypath"))
	return rv
}/* debug [instance_properties/getter]: nonceResponseKeypath */


// The keypath in the response that contains the one-time use value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/nonceResponseKeypath
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetNonceResponseKeypath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNonceResponseKeypath:"), value)
}/* debug [instance_properties/setter]: nonceResponseKeypath */


// The claim name for the previous single sign-on token value in the authentication request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/previousRefreshTokenClaimName
func (a_ AuthorizationProviderExtensionLoginConfiguration) PreviousRefreshTokenClaimName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("previousRefreshTokenClaimName"))
	return rv
}/* debug [instance_properties/getter]: previousRefreshTokenClaimName */


// The claim name for the previous single sign-on token value in the authentication request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/previousRefreshTokenClaimName
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetPreviousRefreshTokenClaimName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreviousRefreshTokenClaimName:"), value)
}/* debug [instance_properties/setter]: previousRefreshTokenClaimName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/refreshEndpointURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) RefreshEndpointURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("refreshEndpointURL"))
	return rv
}/* debug [instance_properties/getter]: refreshEndpointURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/refreshEndpointURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetRefreshEndpointURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRefreshEndpointURL:"), value)
}/* debug [instance_properties/setter]: refreshEndpointURL */


// The name of the claim to include in authentication requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/serverNonceClaimName
func (a_ AuthorizationProviderExtensionLoginConfiguration) ServerNonceClaimName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("serverNonceClaimName"))
	return rv
}/* debug [instance_properties/getter]: serverNonceClaimName */


// The name of the claim to include in authentication requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/serverNonceClaimName
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetServerNonceClaimName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setServerNonceClaimName:"), value)
}/* debug [instance_properties/setter]: serverNonceClaimName */


// The token endpoint URL for login requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/tokenEndpointURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) TokenEndpointURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("tokenEndpointURL"))
	return rv
}/* debug [instance_properties/getter]: tokenEndpointURL */


// The token endpoint URL for login requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/tokenEndpointURL
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetTokenEndpointURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTokenEndpointURL:"), value)
}/* debug [instance_properties/setter]: tokenEndpointURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/uniqueIdentifierClaimName
func (a_ AuthorizationProviderExtensionLoginConfiguration) UniqueIdentifierClaimName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("uniqueIdentifierClaimName"))
	return rv
}/* debug [instance_properties/getter]: uniqueIdentifierClaimName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/uniqueIdentifierClaimName
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetUniqueIdentifierClaimName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUniqueIdentifierClaimName:"), value)
}/* debug [instance_properties/setter]: uniqueIdentifierClaimName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/userSecureEnclaveKeyBiometricPolicy-swift.property
func (a_ AuthorizationProviderExtensionLoginConfiguration) UserSecureEnclaveKeyBiometricPolicy() AuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy {
	rv := objc.Send[AuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy](a_.ID, objc.Sel("userSecureEnclaveKeyBiometricPolicy"))
	return rv
}/* debug [instance_properties/getter]: userSecureEnclaveKeyBiometricPolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/userSecureEnclaveKeyBiometricPolicy-swift.property
func (a_ AuthorizationProviderExtensionLoginConfiguration) SetUserSecureEnclaveKeyBiometricPolicy(value AuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserSecureEnclaveKeyBiometricPolicy:"), value)
}/* debug [instance_properties/setter]: userSecureEnclaveKeyBiometricPolicy */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationProviderExtensionLoginConfiguration */


