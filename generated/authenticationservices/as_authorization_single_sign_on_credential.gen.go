// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationSingleSignOnCredential */


/* debug [class_header]: Header for ASAuthorizationSingleSignOnCredential */
// The class instance for the [AuthorizationSingleSignOnCredential] class.
var (
	AuthorizationSingleSignOnCredentialClass     _AuthorizationSingleSignOnCredentialClass
	AuthorizationSingleSignOnCredentialClassOnce sync.Once
)

func getAuthorizationSingleSignOnCredentialClass() _AuthorizationSingleSignOnCredentialClass {
	AuthorizationSingleSignOnCredentialClassOnce.Do(func() {
		AuthorizationSingleSignOnCredentialClass = _AuthorizationSingleSignOnCredentialClass{objc.GetClass("ASAuthorizationSingleSignOnCredential")}
	})
	return AuthorizationSingleSignOnCredentialClass
}

type _AuthorizationSingleSignOnCredentialClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationSingleSignOnCredential */
// An interface definition for the [AuthorizationSingleSignOnCredential] class.
type IAuthorizationSingleSignOnCredential interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationSingleSignOnCredential */
	// properties:
	AccessToken() objc.IObject /* cross-framework: NSData */
	AuthenticatedResponse() foundation.HTTPURLResponse
	AuthorizedScopes() []string
	IdentityToken() objc.IObject /* cross-framework: NSData */
	PrivateKeys() objc.IObject /* cross-framework: NSArray */
	State() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationSingleSignOnCredential */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationSingleSignOnCredential */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSingleSignOnCredentialClass) Alloc() AuthorizationSingleSignOnCredential {
	rv := objc.Send[AuthorizationSingleSignOnCredential](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationSingleSignOnCredentialClass) New() AuthorizationSingleSignOnCredential {
	rv := objc.Send[AuthorizationSingleSignOnCredential](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSingleSignOnCredential) Init() AuthorizationSingleSignOnCredential {
	rv := objc.Send[AuthorizationSingleSignOnCredential](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSingleSignOnCredential) Autorelease() AuthorizationSingleSignOnCredential {
	rv := objc.Send[AuthorizationSingleSignOnCredential](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSingleSignOnCredential creates a new AuthorizationSingleSignOnCredential instance.
func NewAuthorizationSingleSignOnCredential() AuthorizationSingleSignOnCredential {
	return getAuthorizationSingleSignOnCredentialClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationSingleSignOnCredential */
// A credential that results from a successful single sign-on (SSO) authentication.


// A credential that results from a successful single sign-on (SSO) authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnCredential
type AuthorizationSingleSignOnCredential struct {
	objectivec.Object
}

// AuthorizationSingleSignOnCredentialFrom constructs a [AuthorizationSingleSignOnCredential] from an unsafe.Pointer.
//
// A credential that results from a successful single sign-on (SSO) authentication.
func AuthorizationSingleSignOnCredentialFrom(ptr unsafe.Pointer) AuthorizationSingleSignOnCredential {
	return AuthorizationSingleSignOnCredential{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationSingleSignOnCredential *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationSingleSignOnCredential */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationSingleSignOnCredential */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationSingleSignOnCredential */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationSingleSignOnCredential */

// An access token used to get an identity token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnCredential/accessToken
func (a_ AuthorizationSingleSignOnCredential) AccessToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("accessToken"))
	return rv
}/* debug [instance_properties/getter]: accessToken */


// The complete response authentication, including technology-specific values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnCredential/authenticatedResponse
func (a_ AuthorizationSingleSignOnCredential) AuthenticatedResponse() foundation.HTTPURLResponse {
	rv := objc.Send[foundation.HTTPURLResponse](a_.ID, objc.Sel("authenticatedResponse"))
	return rv
}/* debug [instance_properties/getter]: authenticatedResponse */


// The contact information the user authorized your app to access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnCredential/authorizedScopes
func (a_ AuthorizationSingleSignOnCredential) AuthorizedScopes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("authorizedScopes"))
	return rv
}/* debug [instance_properties/getter]: authorizedScopes */


// A JSON Web Token (JWT) that securely communicates information about the user to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnCredential/identityToken
func (a_ AuthorizationSingleSignOnCredential) IdentityToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("identityToken"))
	return rv
}/* debug [instance_properties/getter]: identityToken */


// An array of private keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnCredential/privateKeys
func (a_ AuthorizationSingleSignOnCredential) PrivateKeys() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("privateKeys"))
	return rv
}/* debug [instance_properties/getter]: privateKeys */


// An arbitrary string that your app provided to the request that generated this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnCredential/state
func (a_ AuthorizationSingleSignOnCredential) State() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationSingleSignOnCredential */



