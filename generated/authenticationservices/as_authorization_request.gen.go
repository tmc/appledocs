// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AuthorizationRequest] class.
type IAuthorizationRequest interface {
	objectivec.IObject
	AuthorizationRequests() ASAuthorizationRequest
	SetAuthorizationRequests(value IASAuthorizationRequest)
	CustomAuthorizationMethods() AuthorizationCustomMethod
	SetCustomAuthorizationMethods(value IAuthorizationCustomMethod)
	Provider() unsafe.Pointer
	SetProvider(value unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationRequestClass) Alloc() AuthorizationRequest {
	rv := objc.Send[AuthorizationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The authorization requests that the controller manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/authorizationrequests

func (a_ AuthorizationRequest) AuthorizationRequests() ASAuthorizationRequest {
	rv := objc.Send[ASAuthorizationRequest](a_.ID, objc.Sel("authorizationRequests"))
	return rv
}


// The authorization requests that the controller manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/authorizationrequests

func (a_ AuthorizationRequest) SetAuthorizationRequests(value IASAuthorizationRequest) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthorizationRequests:"), value)
}


// An array of custom authorization methods for the user to choose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/customauthorizationmethods

func (a_ AuthorizationRequest) CustomAuthorizationMethods() AuthorizationCustomMethod {
	rv := objc.Send[AuthorizationCustomMethod](a_.ID, objc.Sel("customAuthorizationMethods"))
	return rv
}


// An array of custom authorization methods for the user to choose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationcontroller/customauthorizationmethods

func (a_ AuthorizationRequest) SetCustomAuthorizationMethods(value IAuthorizationCustomMethod) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomAuthorizationMethods:"), value)
}


// The provider servicing the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationrequest/provider

func (a_ AuthorizationRequest) Provider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("provider"))
	return rv
}


// The provider servicing the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationrequest/provider

func (a_ AuthorizationRequest) SetProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProvider:"), value)
}



