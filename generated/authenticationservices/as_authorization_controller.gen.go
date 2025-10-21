// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AuthorizationController] class.
type IAuthorizationController interface {
	objectivec.IObject
	Cancel()
}

// A controller that manages authorization requests that a provider creates.
//
// Create authorization requests for the credential types your app supports, such as for Sign in with Apple, or for password credentials. Create an authorization controller using , supplying the authorization requests you create. Set the authorization controller’s to receive responses when requests succeed or fail, and set its so that the authorization controller can present UI. Call to present inline UI to request credentials, or or to request credentials using modal UI. calls your delegate’s methods when the request completes. Set the content type of text fields in your app’s login UI so that can detect when to offer AutoFill suggestions. Use as the content type for user name text fields, and for password fields.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationControllerClass) Alloc() AuthorizationController {
	rv := objc.Send[AuthorizationController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Cancels any active authorization requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/cancel()
func (a_ AuthorizationController) Cancel() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancel"))
}

// The authorization requests that the controller manages.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/authorizationrequests
func (a_ AuthorizationController) AuthorizationRequests() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("authorizationRequests"))
	return rv
}


// SetAuthorizationRequests sets the value of the authorizationRequests property.
// The authorization requests that the controller manages.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/authorizationrequests
func (a_ AuthorizationController) SetAuthorizationRequests(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthorizationRequests:"), value)
}

// A delegate that the authorization controller informs about the success or failure of an authorization attempt.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/delegate
func (a_ AuthorizationController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate that the authorization controller informs about the success or failure of an authorization attempt.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/delegate
func (a_ AuthorizationController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}

// An array of custom authorization methods for the user to choose.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/customAuthorizationMethods
func (a_ AuthorizationController) CustomAuthorizationMethods() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("customAuthorizationMethods"))
	return rv
}


// SetCustomAuthorizationMethods sets the value of the customAuthorizationMethods property.
// An array of custom authorization methods for the user to choose.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/customAuthorizationMethods
func (a_ AuthorizationController) SetCustomAuthorizationMethods(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomAuthorizationMethods:"), nsArray)
}

// A delegate that provides a display context in which the system can present an authorization interface to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/presentationContextProvider
func (a_ AuthorizationController) PresentationContextProvider() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("presentationContextProvider"))
	return rv
}


// SetPresentationContextProvider sets the value of the presentationContextProvider property.
// A delegate that provides a display context in which the system can present an authorization interface to the user.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/presentationContextProvider
func (a_ AuthorizationController) SetPresentationContextProvider(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPresentationContextProvider:"), value)
}



