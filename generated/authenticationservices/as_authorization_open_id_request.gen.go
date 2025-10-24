// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ASAuthorizationOpenIDRequest */


/* debug [class_header]: Header for ASAuthorizationOpenIDRequest */
// The class instance for the [AuthorizationOpenIDRequest] class.
var (
	AuthorizationOpenIDRequestClass     _AuthorizationOpenIDRequestClass
	AuthorizationOpenIDRequestClassOnce sync.Once
)

func getAuthorizationOpenIDRequestClass() _AuthorizationOpenIDRequestClass {
	AuthorizationOpenIDRequestClassOnce.Do(func() {
		AuthorizationOpenIDRequestClass = _AuthorizationOpenIDRequestClass{objc.GetClass("ASAuthorizationOpenIDRequest")}
	})
	return AuthorizationOpenIDRequestClass
}

type _AuthorizationOpenIDRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationOpenIDRequest */
// An interface definition for the [AuthorizationOpenIDRequest] class.
type IAuthorizationOpenIDRequest interface {
	IAuthorizationRequest
	
/* debug [class_interface_properties]: Properties for AuthorizationOpenIDRequest */
	// properties:
	Nonce() objc.IObject /* cross-framework: NSString */
	SetNonce(value objc.IObject /* cross-framework: NSString */)
	RequestedOperation() AuthorizationOpenIDOperation /* typedef */
	SetRequestedOperation(value AuthorizationOpenIDOperation /* typedef */)
	RequestedScopes() []string
	SetRequestedScopes(value []string)
	State() objc.IObject /* cross-framework: NSString */
	SetState(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationOpenIDRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationOpenIDRequest */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationOpenIDRequestClass) Alloc() AuthorizationOpenIDRequest {
	rv := objc.Send[AuthorizationOpenIDRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationOpenIDRequestClass) New() AuthorizationOpenIDRequest {
	rv := objc.Send[AuthorizationOpenIDRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationOpenIDRequest) Init() AuthorizationOpenIDRequest {
	rv := objc.Send[AuthorizationOpenIDRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationOpenIDRequest) Autorelease() AuthorizationOpenIDRequest {
	rv := objc.Send[AuthorizationOpenIDRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationOpenIDRequest creates a new AuthorizationOpenIDRequest instance.
func NewAuthorizationOpenIDRequest() AuthorizationOpenIDRequest {
	return getAuthorizationOpenIDRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationOpenIDRequest */
// An OpenID authorization request.


// An OpenID authorization request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest
type AuthorizationOpenIDRequest struct {
	AuthorizationRequest
}

// AuthorizationOpenIDRequestFrom constructs a [AuthorizationOpenIDRequest] from an unsafe.Pointer.
//
// An OpenID authorization request.
func AuthorizationOpenIDRequestFrom(ptr unsafe.Pointer) AuthorizationOpenIDRequest {
	return AuthorizationOpenIDRequest{
		AuthorizationRequest: AuthorizationRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationOpenIDRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationOpenIDRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationOpenIDRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationOpenIDRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationOpenIDRequest */

// A string value to pass to the identity provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/nonce
func (a_ AuthorizationOpenIDRequest) Nonce() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("nonce"))
	return rv
}/* debug [instance_properties/getter]: nonce */


// A string value to pass to the identity provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/nonce
func (a_ AuthorizationOpenIDRequest) SetNonce(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNonce:"), value)
}/* debug [instance_properties/setter]: nonce */


// The OpenID authentication operation you want this request to perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedOperation
func (a_ AuthorizationOpenIDRequest) RequestedOperation() AuthorizationOpenIDOperation /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("requestedOperation"))
	return rv
}/* debug [instance_properties/getter]: requestedOperation */


// The OpenID authentication operation you want this request to perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedOperation
func (a_ AuthorizationOpenIDRequest) SetRequestedOperation(value AuthorizationOpenIDOperation /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestedOperation:"), value)
}/* debug [instance_properties/setter]: requestedOperation */


// The contact information to be requested from the user during authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedScopes
func (a_ AuthorizationOpenIDRequest) RequestedScopes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("requestedScopes"))
	return rv
}/* debug [instance_properties/getter]: requestedScopes */


// The contact information to be requested from the user during authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/requestedScopes
func (a_ AuthorizationOpenIDRequest) SetRequestedScopes(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestedScopes:"), nsArray)
}/* debug [instance_properties/setter]: requestedScopes */


// Data that’s returned to you unmodified in the corresponding credential after a successful authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/state
func (a_ AuthorizationOpenIDRequest) State() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// Data that’s returned to you unmodified in the corresponding credential after a successful authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationOpenIDRequest/state
func (a_ AuthorizationOpenIDRequest) SetState(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationOpenIDRequest */



