// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationAppleIDProvider */


/* debug [class_header]: Header for ASAuthorizationAppleIDProvider */
// The class instance for the [AuthorizationAppleIDProvider] class.
var (
	AuthorizationAppleIDProviderClass     _AuthorizationAppleIDProviderClass
	AuthorizationAppleIDProviderClassOnce sync.Once
)

func getAuthorizationAppleIDProviderClass() _AuthorizationAppleIDProviderClass {
	AuthorizationAppleIDProviderClassOnce.Do(func() {
		AuthorizationAppleIDProviderClass = _AuthorizationAppleIDProviderClass{objc.GetClass("ASAuthorizationAppleIDProvider")}
	})
	return AuthorizationAppleIDProviderClass
}

type _AuthorizationAppleIDProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationAppleIDProvider */
// An interface definition for the [AuthorizationAppleIDProvider] class.
type IAuthorizationAppleIDProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationAppleIDProvider */
	// properties:
	User() objc.IObject /* cross-framework: NSString */
	SetUser(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationAppleIDProvider */
	// methods:
	CreateRequest() IAuthorizationAppleIDRequest
	GetCredentialStateForUserIDCompletion(userID objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationAppleIDProvider */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationAppleIDProviderClass) Alloc() AuthorizationAppleIDProvider {
	rv := objc.Send[AuthorizationAppleIDProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationAppleIDProviderClass) New() AuthorizationAppleIDProvider {
	rv := objc.Send[AuthorizationAppleIDProvider](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationAppleIDProvider) Init() AuthorizationAppleIDProvider {
	rv := objc.Send[AuthorizationAppleIDProvider](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationAppleIDProvider) Autorelease() AuthorizationAppleIDProvider {
	rv := objc.Send[AuthorizationAppleIDProvider](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationAppleIDProvider creates a new AuthorizationAppleIDProvider instance.
func NewAuthorizationAppleIDProvider() AuthorizationAppleIDProvider {
	return getAuthorizationAppleIDProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationAppleIDProvider */
// A mechanism for generating requests to authenticate users based on their Apple ID.
//
// You use a provider to create a request ( ), which you then use to initialize a controller ( ) that performs the request: On success, the controller’s delegate receives an authorization ( ) containing a credential ( ) that has an opaque identifier. You can use that identifier to later check the user’s credential state—for example, to see if authorization has been revoked—by calling the method:


// A mechanism for generating requests to authenticate users based on their Apple ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider
type AuthorizationAppleIDProvider struct {
	objectivec.Object
}

// AuthorizationAppleIDProviderFrom constructs a [AuthorizationAppleIDProvider] from an unsafe.Pointer.
//
// A mechanism for generating requests to authenticate users based on their Apple ID.
func AuthorizationAppleIDProviderFrom(ptr unsafe.Pointer) AuthorizationAppleIDProvider {
	return AuthorizationAppleIDProvider{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationAppleIDProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationAppleIDProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationAppleIDProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationAppleIDProvider */

// Creates a new Apple ID authorization request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider/createRequest()
func (a_ AuthorizationAppleIDProvider) CreateRequest() IAuthorizationAppleIDRequest {
	rv := objc.Send[AuthorizationAppleIDRequest](a_.ID, objc.Sel("createRequest"))
	return rv
}/* debug [instance_methods/method]: CreateRequest */


// Returns the credential state for the given user in a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider/getCredentialState(forUserID:completion:)
func (a_ AuthorizationAppleIDProvider) GetCredentialStateForUserIDCompletion(userID objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("getCredentialStateForUserID:completion:"), userID, completion)
}/* debug [instance_methods/method]: GetCredentialStateForUserIDCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationAppleIDProvider */

// An identifier for the authenticated user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/user
func (a_ AuthorizationAppleIDProvider) User() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("user"))
	return rv
}/* debug [instance_properties/getter]: user */


// An identifier for the authenticated user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/user
func (a_ AuthorizationAppleIDProvider) SetUser(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUser:"), value)
}/* debug [instance_properties/setter]: user */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationAppleIDProvider */



