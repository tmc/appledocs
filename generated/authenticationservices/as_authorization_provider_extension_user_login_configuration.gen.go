// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationProviderExtensionUserLoginConfiguration */


/* debug [class_header]: Header for ASAuthorizationProviderExtensionUserLoginConfiguration */
// The class instance for the [AuthorizationProviderExtensionUserLoginConfiguration] class.
var (
	AuthorizationProviderExtensionUserLoginConfigurationClass     _AuthorizationProviderExtensionUserLoginConfigurationClass
	AuthorizationProviderExtensionUserLoginConfigurationClassOnce sync.Once
)

func getAuthorizationProviderExtensionUserLoginConfigurationClass() _AuthorizationProviderExtensionUserLoginConfigurationClass {
	AuthorizationProviderExtensionUserLoginConfigurationClassOnce.Do(func() {
		AuthorizationProviderExtensionUserLoginConfigurationClass = _AuthorizationProviderExtensionUserLoginConfigurationClass{objc.GetClass("ASAuthorizationProviderExtensionUserLoginConfiguration")}
	})
	return AuthorizationProviderExtensionUserLoginConfigurationClass
}

type _AuthorizationProviderExtensionUserLoginConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationProviderExtensionUserLoginConfiguration */
// An interface definition for the [AuthorizationProviderExtensionUserLoginConfiguration] class.
type IAuthorizationProviderExtensionUserLoginConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationProviderExtensionUserLoginConfiguration */
	// properties:
	LoginUserName() objc.IObject /* cross-framework: NSString */
	SetLoginUserName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationProviderExtensionUserLoginConfiguration */
	// methods:
	SetCustomAssertionRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomAssertionRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomLoginRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
	SetCustomLoginRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationProviderExtensionUserLoginConfiguration */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionUserLoginConfigurationClass) Alloc() AuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationProviderExtensionUserLoginConfigurationClass) New() AuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationProviderExtensionUserLoginConfiguration) Init() AuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationProviderExtensionUserLoginConfiguration) Autorelease() AuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationProviderExtensionUserLoginConfiguration creates a new AuthorizationProviderExtensionUserLoginConfiguration instance.
func NewAuthorizationProviderExtensionUserLoginConfiguration() AuthorizationProviderExtensionUserLoginConfiguration {
	return getAuthorizationProviderExtensionUserLoginConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationProviderExtensionUserLoginConfiguration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserLoginConfiguration
type AuthorizationProviderExtensionUserLoginConfiguration struct {
	objectivec.Object
}

// AuthorizationProviderExtensionUserLoginConfigurationFrom constructs a [AuthorizationProviderExtensionUserLoginConfiguration] from an unsafe.Pointer.
func AuthorizationProviderExtensionUserLoginConfigurationFrom(ptr unsafe.Pointer) AuthorizationProviderExtensionUserLoginConfiguration {
	return AuthorizationProviderExtensionUserLoginConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationProviderExtensionUserLoginConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserLoginConfiguration/init(loginUserName:)
func NewAuthorizationProviderExtensionUserLoginConfigurationWithLoginUserName(loginUserName objc.IObject /* cross-framework: NSString */) AuthorizationProviderExtensionUserLoginConfiguration {
	instance := getAuthorizationProviderExtensionUserLoginConfigurationClass().Alloc()
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](instance.ID, objc.Sel("initWithLoginUserName:"), loginUserName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationProviderExtensionUserLoginConfigurationWithLoginUserName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationProviderExtensionUserLoginConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationProviderExtensionUserLoginConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationProviderExtensionUserLoginConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserLoginConfiguration/setCustomAssertionRequestBodyClaims(_:)
func (a_ AuthorizationProviderExtensionUserLoginConfiguration) SetCustomAssertionRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomAssertionRequestBodyClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomAssertionRequestBodyClaimsReturningError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserLoginConfiguration/setCustomAssertionRequestHeaderClaims(_:)
func (a_ AuthorizationProviderExtensionUserLoginConfiguration) SetCustomAssertionRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomAssertionRequestHeaderClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomAssertionRequestHeaderClaimsReturningError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserLoginConfiguration/setCustomLoginRequestBodyClaims(_:)
func (a_ AuthorizationProviderExtensionUserLoginConfiguration) SetCustomLoginRequestBodyClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomLoginRequestBodyClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomLoginRequestBodyClaimsReturningError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserLoginConfiguration/setCustomLoginRequestHeaderClaims(_:)
func (a_ AuthorizationProviderExtensionUserLoginConfiguration) SetCustomLoginRequestHeaderClaimsReturningError(claims foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomLoginRequestHeaderClaims:returningError:"), claims, error_)
	return rv
}/* debug [instance_methods/method]: SetCustomLoginRequestHeaderClaimsReturningError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationProviderExtensionUserLoginConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserLoginConfiguration/loginUserName
func (a_ AuthorizationProviderExtensionUserLoginConfiguration) LoginUserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("loginUserName"))
	return rv
}/* debug [instance_properties/getter]: loginUserName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserLoginConfiguration/loginUserName
func (a_ AuthorizationProviderExtensionUserLoginConfiguration) SetLoginUserName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginUserName:"), value)
}/* debug [instance_properties/setter]: loginUserName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationProviderExtensionUserLoginConfiguration */


