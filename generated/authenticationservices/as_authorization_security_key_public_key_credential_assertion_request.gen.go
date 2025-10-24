// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */


/* debug [class_header]: Header for ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */
// The class instance for the [AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest] class.
var (
	AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass     _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass
	AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClassOnce sync.Once
)

func getAuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass() _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass {
	AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClassOnce.Do(func() {
		AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass = _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass{objc.GetClass("ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest")}
	})
	return AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass
}

type _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */
// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest] class.
type IAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest interface {
	IAuthorizationRequest
	
/* debug [class_interface_properties]: Properties for AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */
	// properties:
	AllowedCredentials() []AuthorizationSecurityKeyPublicKeyCredentialDescriptor
	SetAllowedCredentials(value []AuthorizationSecurityKeyPublicKeyCredentialDescriptor)
	AppID() objc.IObject /* cross-framework: NSString */
	SetAppID(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass) New() AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest) Init() AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest) Autorelease() AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest creates a new AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest instance.
func NewAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest() AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	return getAuthorizationSecurityKeyPublicKeyCredentialAssertionRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */
// A class that defines the assertion request type for security key credentials.
//
// Use this class to sign in with an existing credential on a security key.


// A class that defines the assertion request type for security key credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest
type AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest struct {
	AuthorizationRequest
}

// AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestFrom constructs a [AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest] from an unsafe.Pointer.
//
// A class that defines the assertion request type for security key credentials.
func AuthorizationSecurityKeyPublicKeyCredentialAssertionRequestFrom(ptr unsafe.Pointer) AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest {
	return AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest{
		AuthorizationRequest: AuthorizationRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */

// An array of allowed credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest/allowedCredentials
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest) AllowedCredentials() []AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	rv := objc.Send[[]AuthorizationSecurityKeyPublicKeyCredentialDescriptor](a_.ID, objc.Sel("allowedCredentials"))
	return rv
}/* debug [instance_properties/getter]: allowedCredentials */


// An array of allowed credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest/allowedCredentials
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest) SetAllowedCredentials(value []AuthorizationSecurityKeyPublicKeyCredentialDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowedCredentials:"), nsArray)
}/* debug [instance_properties/setter]: allowedCredentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest/appID
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest) AppID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("appID"))
	return rv
}/* debug [instance_properties/getter]: appID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest/appID
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest) SetAppID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAppID:"), value)
}/* debug [instance_properties/setter]: appID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationSecurityKeyPublicKeyCredentialAssertionRequest */



