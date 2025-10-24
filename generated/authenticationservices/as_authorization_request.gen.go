// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationRequest */


/* debug [class_header]: Header for ASAuthorizationRequest */
// The class instance for the [AuthorizationRequest] class.
var (
	AuthorizationRequestClass     _AuthorizationRequestClass
	AuthorizationRequestClassOnce sync.Once
)

func getAuthorizationRequestClass() _AuthorizationRequestClass {
	AuthorizationRequestClassOnce.Do(func() {
		AuthorizationRequestClass = _AuthorizationRequestClass{objc.GetClass("ASAuthorizationRequest")}
	})
	return AuthorizationRequestClass
}

type _AuthorizationRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationRequest */
// An interface definition for the [AuthorizationRequest] class.
type IAuthorizationRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationRequest */
	// properties:
	Provider() unsafe.Pointer
	AuthorizationRequests() IASAuthorizationRequest
	SetAuthorizationRequests(value IASAuthorizationRequest)
	CustomAuthorizationMethods() AuthorizationCustomMethod /* typedef */
	SetCustomAuthorizationMethods(value AuthorizationCustomMethod /* typedef */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationRequest */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationRequestClass) Alloc() AuthorizationRequest {
	rv := objc.Send[AuthorizationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationRequestClass) New() AuthorizationRequest {
	rv := objc.Send[AuthorizationRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationRequest) Init() AuthorizationRequest {
	rv := objc.Send[AuthorizationRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationRequest) Autorelease() AuthorizationRequest {
	rv := objc.Send[AuthorizationRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationRequest creates a new AuthorizationRequest instance.
func NewAuthorizationRequest() AuthorizationRequest {
	return getAuthorizationRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationRequest */
// A base class for different kinds of authorization requests.
//
// Use one of the concrete requests, like , , or . You typically generate one of these using the corresponding provider, which is an instance of , , or , respectively.


// A base class for different kinds of authorization requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationRequest
type AuthorizationRequest struct {
	objectivec.Object
}

// AuthorizationRequestFrom constructs a [AuthorizationRequest] from an unsafe.Pointer.
//
// A base class for different kinds of authorization requests.
func AuthorizationRequestFrom(ptr unsafe.Pointer) AuthorizationRequest {
	return AuthorizationRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationRequest */

// The provider servicing the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationRequest/provider
func (a_ AuthorizationRequest) Provider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("provider"))
	return rv
}/* debug [instance_properties/getter]: provider */


// The authorization requests that the controller manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/authorizationrequests
func (a_ AuthorizationRequest) AuthorizationRequests() IASAuthorizationRequest {
	rv := objc.Send[AuthorizationRequest](a_.ID, objc.Sel("authorizationRequests"))
	return rv
}/* debug [instance_properties/getter]: authorizationRequests */


// The authorization requests that the controller manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/authorizationrequests
func (a_ AuthorizationRequest) SetAuthorizationRequests(value IASAuthorizationRequest) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthorizationRequests:"), value)
}/* debug [instance_properties/setter]: authorizationRequests */


// An array of custom authorization methods for the user to choose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/customauthorizationmethods
func (a_ AuthorizationRequest) CustomAuthorizationMethods() AuthorizationCustomMethod /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("customAuthorizationMethods"))
	return rv
}/* debug [instance_properties/getter]: customAuthorizationMethods */


// An array of custom authorization methods for the user to choose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/customauthorizationmethods
func (a_ AuthorizationRequest) SetCustomAuthorizationMethods(value AuthorizationCustomMethod /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomAuthorizationMethods:"), value)
}/* debug [instance_properties/setter]: customAuthorizationMethods */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationRequest */



