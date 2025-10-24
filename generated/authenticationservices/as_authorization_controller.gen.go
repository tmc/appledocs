// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationController */


/* debug [class_header]: Header for ASAuthorizationController */
// The class instance for the [AuthorizationController] class.
var (
	AuthorizationControllerClass     _AuthorizationControllerClass
	AuthorizationControllerClassOnce sync.Once
)

func getAuthorizationControllerClass() _AuthorizationControllerClass {
	AuthorizationControllerClassOnce.Do(func() {
		AuthorizationControllerClass = _AuthorizationControllerClass{objc.GetClass("ASAuthorizationController")}
	})
	return AuthorizationControllerClass
}

type _AuthorizationControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationController */
// An interface definition for the [AuthorizationController] class.
type IAuthorizationController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationController */
	// properties:
	AuthorizationRequests() []AuthorizationRequest
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	PresentationContextProvider() unsafe.Pointer
	SetPresentationContextProvider(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationController */
	// methods:
	Cancel()
	PerformRequests()
	PerformRequestsWithOptions(options AuthorizationControllerRequestOptions)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationController */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationControllerClass) Alloc() AuthorizationController {
	rv := objc.Send[AuthorizationController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationControllerClass) New() AuthorizationController {
	rv := objc.Send[AuthorizationController](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationController) Init() AuthorizationController {
	rv := objc.Send[AuthorizationController](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationController) Autorelease() AuthorizationController {
	rv := objc.Send[AuthorizationController](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationController creates a new AuthorizationController instance.
func NewAuthorizationController() AuthorizationController {
	return getAuthorizationControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationController */
// A controller that manages authorization requests that a provider creates.
//
// Create authorization requests for the credential types your app supports, such as for Sign in with Apple, or for password credentials. Create an authorization controller using , supplying the authorization requests you create. Set the authorization controller’s to receive responses when requests succeed or fail, and set its so that the authorization controller can present UI. Call to present inline UI to request credentials, or or to request credentials using modal UI. calls your delegate’s methods when the request completes. Set the content type of text fields in your app’s login UI so that can detect when to offer AutoFill suggestions. Use as the content type for user name text fields, and for password fields.


// A controller that manages authorization requests that a provider creates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController
type AuthorizationController struct {
	objectivec.Object
}

// AuthorizationControllerFrom constructs a [AuthorizationController] from an unsafe.Pointer.
//
// A controller that manages authorization requests that a provider creates.
func AuthorizationControllerFrom(ptr unsafe.Pointer) AuthorizationController {
	return AuthorizationController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationController */

// Creates a controller from a collection of authorization requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/init(authorizationRequests:)
func NewAuthorizationControllerWithAuthorizationRequests(authorizationRequests []AuthorizationRequest) AuthorizationController {
	instance := getAuthorizationControllerClass().Alloc()
	rv := objc.Send[AuthorizationController](instance.ID, objc.Sel("initWithAuthorizationRequests:"), authorizationRequests)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationControllerWithAuthorizationRequests */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationController */

// Cancels any active authorization requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/cancel()
func (a_ AuthorizationController) Cancel() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Starts the specified authorization flows during controller initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/performRequests()
func (a_ AuthorizationController) PerformRequests() {
	objc.Send[objc.ID](a_.ID, objc.Sel("performRequests"))
}/* debug [instance_methods/method]: PerformRequests */


// Starts the specified authorization flows during controller initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/performRequests(options:)
func (a_ AuthorizationController) PerformRequestsWithOptions(options AuthorizationControllerRequestOptions) {
	objc.Send[objc.ID](a_.ID, objc.Sel("performRequestsWithOptions:"), options)
}/* debug [instance_methods/method]: PerformRequestsWithOptions */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationController */

// The authorization requests that the controller manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/authorizationRequests
func (a_ AuthorizationController) AuthorizationRequests() []AuthorizationRequest {
	rv := objc.Send[[]AuthorizationRequest](a_.ID, objc.Sel("authorizationRequests"))
	return rv
}/* debug [instance_properties/getter]: authorizationRequests */


// A delegate that the authorization controller informs about the success or failure of an authorization attempt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/delegate
func (a_ AuthorizationController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// A delegate that the authorization controller informs about the success or failure of an authorization attempt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/delegate
func (a_ AuthorizationController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A delegate that provides a display context in which the system can present an authorization interface to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/presentationContextProvider
func (a_ AuthorizationController) PresentationContextProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("presentationContextProvider"))
	return rv
}/* debug [instance_properties/getter]: presentationContextProvider */


// A delegate that provides a display context in which the system can present an authorization interface to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/presentationContextProvider
func (a_ AuthorizationController) SetPresentationContextProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPresentationContextProvider:"), value)
}/* debug [instance_properties/setter]: presentationContextProvider */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationController */


