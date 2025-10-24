// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ASAuthorizationSingleSignOnRequest */


/* debug [class_header]: Header for ASAuthorizationSingleSignOnRequest */
// The class instance for the [AuthorizationSingleSignOnRequest] class.
var (
	AuthorizationSingleSignOnRequestClass     _AuthorizationSingleSignOnRequestClass
	AuthorizationSingleSignOnRequestClassOnce sync.Once
)

func getAuthorizationSingleSignOnRequestClass() _AuthorizationSingleSignOnRequestClass {
	AuthorizationSingleSignOnRequestClassOnce.Do(func() {
		AuthorizationSingleSignOnRequestClass = _AuthorizationSingleSignOnRequestClass{objc.GetClass("ASAuthorizationSingleSignOnRequest")}
	})
	return AuthorizationSingleSignOnRequestClass
}

type _AuthorizationSingleSignOnRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationSingleSignOnRequest */
// An interface definition for the [AuthorizationSingleSignOnRequest] class.
type IAuthorizationSingleSignOnRequest interface {
	IAuthorizationOpenIDRequest
	
/* debug [class_interface_properties]: Properties for AuthorizationSingleSignOnRequest */
	// properties:
	AuthorizationOptions() []foundation.URLQueryItem
	SetAuthorizationOptions(value []foundation.URLQueryItem)
	UserInterfaceEnabled() bool
	SetUserInterfaceEnabled(value bool)
	CanPerformAuthorization() bool
	SetCanPerformAuthorization(value bool)
	IsUserInterfaceEnabled() bool
	SetIsUserInterfaceEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationSingleSignOnRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationSingleSignOnRequest */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSingleSignOnRequestClass) Alloc() AuthorizationSingleSignOnRequest {
	rv := objc.Send[AuthorizationSingleSignOnRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationSingleSignOnRequestClass) New() AuthorizationSingleSignOnRequest {
	rv := objc.Send[AuthorizationSingleSignOnRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSingleSignOnRequest) Init() AuthorizationSingleSignOnRequest {
	rv := objc.Send[AuthorizationSingleSignOnRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSingleSignOnRequest) Autorelease() AuthorizationSingleSignOnRequest {
	rv := objc.Send[AuthorizationSingleSignOnRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSingleSignOnRequest creates a new AuthorizationSingleSignOnRequest instance.
func NewAuthorizationSingleSignOnRequest() AuthorizationSingleSignOnRequest {
	return getAuthorizationSingleSignOnRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationSingleSignOnRequest */
// An OpenID authorization request that provides single sign-on (SSO) functionality.


// An OpenID authorization request that provides single sign-on (SSO) functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnRequest
type AuthorizationSingleSignOnRequest struct {
	AuthorizationOpenIDRequest
}

// AuthorizationSingleSignOnRequestFrom constructs a [AuthorizationSingleSignOnRequest] from an unsafe.Pointer.
//
// An OpenID authorization request that provides single sign-on (SSO) functionality.
func AuthorizationSingleSignOnRequestFrom(ptr unsafe.Pointer) AuthorizationSingleSignOnRequest {
	return AuthorizationSingleSignOnRequest{
		AuthorizationOpenIDRequest: AuthorizationOpenIDRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationSingleSignOnRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationSingleSignOnRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationSingleSignOnRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationSingleSignOnRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationSingleSignOnRequest */

// Options that control the authorization process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnRequest/authorizationOptions
func (a_ AuthorizationSingleSignOnRequest) AuthorizationOptions() []foundation.URLQueryItem {
	rv := objc.Send[[]foundation.URLQueryItem](a_.ID, objc.Sel("authorizationOptions"))
	return rv
}/* debug [instance_properties/getter]: authorizationOptions */


// Options that control the authorization process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnRequest/authorizationOptions
func (a_ AuthorizationSingleSignOnRequest) SetAuthorizationOptions(value []foundation.URLQueryItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthorizationOptions:"), nsArray)
}/* debug [instance_properties/setter]: authorizationOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnRequest/isUserInterfaceEnabled
func (a_ AuthorizationSingleSignOnRequest) UserInterfaceEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("userInterfaceEnabled"))
	return rv
}/* debug [instance_properties/getter]: userInterfaceEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnRequest/isUserInterfaceEnabled
func (a_ AuthorizationSingleSignOnRequest) SetUserInterfaceEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserInterfaceEnabled:"), value)
}/* debug [instance_properties/setter]: userInterfaceEnabled */


// A Boolean value that indicates if the provider is capable of performing authorization within a given configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignonprovider/canperformauthorization
func (a_ AuthorizationSingleSignOnRequest) CanPerformAuthorization() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformAuthorization"))
	return rv
}/* debug [instance_properties/getter]: canPerformAuthorization */


// A Boolean value that indicates if the provider is capable of performing authorization within a given configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignonprovider/canperformauthorization
func (a_ AuthorizationSingleSignOnRequest) SetCanPerformAuthorization(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanPerformAuthorization:"), value)
}/* debug [instance_properties/setter]: canPerformAuthorization */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignonrequest/isuserinterfaceenabled
func (a_ AuthorizationSingleSignOnRequest) IsUserInterfaceEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isUserInterfaceEnabled"))
	return rv
}/* debug [instance_properties/getter]: isUserInterfaceEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignonrequest/isuserinterfaceenabled
func (a_ AuthorizationSingleSignOnRequest) SetIsUserInterfaceEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsUserInterfaceEnabled:"), value)
}/* debug [instance_properties/setter]: isUserInterfaceEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationSingleSignOnRequest */



