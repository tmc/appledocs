// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationSingleSignOnProvider */


/* debug [class_header]: Header for ASAuthorizationSingleSignOnProvider */
// The class instance for the [AuthorizationSingleSignOnProvider] class.
var (
	AuthorizationSingleSignOnProviderClass     _AuthorizationSingleSignOnProviderClass
	AuthorizationSingleSignOnProviderClassOnce sync.Once
)

func getAuthorizationSingleSignOnProviderClass() _AuthorizationSingleSignOnProviderClass {
	AuthorizationSingleSignOnProviderClassOnce.Do(func() {
		AuthorizationSingleSignOnProviderClass = _AuthorizationSingleSignOnProviderClass{objc.GetClass("ASAuthorizationSingleSignOnProvider")}
	})
	return AuthorizationSingleSignOnProviderClass
}

type _AuthorizationSingleSignOnProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationSingleSignOnProvider */
// An interface definition for the [AuthorizationSingleSignOnProvider] class.
type IAuthorizationSingleSignOnProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationSingleSignOnProvider */
	// properties:
	CanPerformAuthorization() bool
	Url() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationSingleSignOnProvider */
	// methods:
	CreateRequest() IAuthorizationSingleSignOnRequest
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationSingleSignOnProvider */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSingleSignOnProviderClass) Alloc() AuthorizationSingleSignOnProvider {
	rv := objc.Send[AuthorizationSingleSignOnProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationSingleSignOnProviderClass) New() AuthorizationSingleSignOnProvider {
	rv := objc.Send[AuthorizationSingleSignOnProvider](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSingleSignOnProvider) Init() AuthorizationSingleSignOnProvider {
	rv := objc.Send[AuthorizationSingleSignOnProvider](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSingleSignOnProvider) Autorelease() AuthorizationSingleSignOnProvider {
	rv := objc.Send[AuthorizationSingleSignOnProvider](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSingleSignOnProvider creates a new AuthorizationSingleSignOnProvider instance.
func NewAuthorizationSingleSignOnProvider() AuthorizationSingleSignOnProvider {
	return getAuthorizationSingleSignOnProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationSingleSignOnProvider */
// A mechanism for generating requests to authenticate users with third-party providers.


// A mechanism for generating requests to authenticate users with third-party providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnProvider
type AuthorizationSingleSignOnProvider struct {
	objectivec.Object
}

// AuthorizationSingleSignOnProviderFrom constructs a [AuthorizationSingleSignOnProvider] from an unsafe.Pointer.
//
// A mechanism for generating requests to authenticate users with third-party providers.
func AuthorizationSingleSignOnProviderFrom(ptr unsafe.Pointer) AuthorizationSingleSignOnProvider {
	return AuthorizationSingleSignOnProvider{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationSingleSignOnProvider */

// Creates a single sign-on (SSO) authorization provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnProvider/init(identityProvider:)
func NewAuthorizationSingleSignOnProviderAuthorizationProviderWithIdentityProviderURL(url objc.IObject /* cross-framework: NSURL */) AuthorizationSingleSignOnProvider {
	rv := objc.Send[AuthorizationSingleSignOnProvider](objc.ID(getAuthorizationSingleSignOnProviderClass().class), objc.Sel("authorizationProviderWithIdentityProviderURL:"), url)
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationSingleSignOnProviderAuthorizationProviderWithIdentityProviderURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationSingleSignOnProvider */

// Creates a single sign-on (SSO) authorization provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnProvider/init(identityProvider:)
func (ac _AuthorizationSingleSignOnProviderClass) AuthorizationProviderWithIdentityProviderURL(url objc.IObject /* cross-framework: NSURL */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("authorizationProviderWithIdentityProviderURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationProviderWithIdentityProviderURL) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationSingleSignOnProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationSingleSignOnProvider */

// Creates a single sign-on (SSO) authorization request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnProvider/createRequest()
func (a_ AuthorizationSingleSignOnProvider) CreateRequest() IAuthorizationSingleSignOnRequest {
	rv := objc.Send[AuthorizationSingleSignOnRequest](a_.ID, objc.Sel("createRequest"))
	return rv
}/* debug [instance_methods/method]: CreateRequest */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationSingleSignOnProvider */

// A Boolean value that indicates if the provider is capable of performing authorization within a given configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnProvider/canPerformAuthorization
func (a_ AuthorizationSingleSignOnProvider) CanPerformAuthorization() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformAuthorization"))
	return rv
}/* debug [instance_properties/getter]: canPerformAuthorization */


// The URL of the identity provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnProvider/url
func (a_ AuthorizationSingleSignOnProvider) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationSingleSignOnProvider */


