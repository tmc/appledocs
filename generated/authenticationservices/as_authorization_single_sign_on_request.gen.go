// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [AuthorizationSingleSignOnRequest] class.
type IAuthorizationSingleSignOnRequest interface {
	IAuthorizationOpenIDRequest
	AuthorizationOptions() []foundation.URLQueryItem
	SetAuthorizationOptions(value []foundation.IURLQueryItem)
	UserInterfaceEnabled() bool
	SetUserInterfaceEnabled(value bool)
	CanPerformAuthorization() bool
	SetCanPerformAuthorization(value bool)
	IsUserInterfaceEnabled() bool
	SetIsUserInterfaceEnabled(value bool)
}

// An OpenID authorization request that provides single sign-on (SSO) functionality.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSingleSignOnRequestClass) Alloc() AuthorizationSingleSignOnRequest {
	rv := objc.Send[AuthorizationSingleSignOnRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Options that control the authorization process.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnRequest/authorizationOptions
func (a_ AuthorizationSingleSignOnRequest) AuthorizationOptions() []foundation.URLQueryItem {
	rv := objc.Send[[]foundation.URLQueryItem](a_.ID, objc.Sel("authorizationOptions"))
	return rv
}


// SetAuthorizationOptions sets the value of the authorizationOptions property.
// Options that control the authorization process.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnRequest/authorizationOptions
func (a_ AuthorizationSingleSignOnRequest) SetAuthorizationOptions(value []foundation.IURLQueryItem) {
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
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthorizationOptions:"), nsArray)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnRequest/isUserInterfaceEnabled
func (a_ AuthorizationSingleSignOnRequest) UserInterfaceEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("userInterfaceEnabled"))
	return rv
}


// SetUserInterfaceEnabled sets the value of the userInterfaceEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnRequest/isUserInterfaceEnabled
func (a_ AuthorizationSingleSignOnRequest) SetUserInterfaceEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserInterfaceEnabled:"), value)
}

// A Boolean value that indicates if the provider is capable of performing authorization within a given configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignonprovider/canperformauthorization
func (a_ AuthorizationSingleSignOnRequest) CanPerformAuthorization() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformAuthorization"))
	return rv
}


// SetCanPerformAuthorization sets the value of the canPerformAuthorization property.
// A Boolean value that indicates if the provider is capable of performing authorization within a given configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignonprovider/canperformauthorization
func (a_ AuthorizationSingleSignOnRequest) SetCanPerformAuthorization(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanPerformAuthorization:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignonrequest/isuserinterfaceenabled
func (a_ AuthorizationSingleSignOnRequest) IsUserInterfaceEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isUserInterfaceEnabled"))
	return rv
}


// SetIsUserInterfaceEnabled sets the value of the isUserInterfaceEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignonrequest/isuserinterfaceenabled
func (a_ AuthorizationSingleSignOnRequest) SetIsUserInterfaceEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsUserInterfaceEnabled:"), value)
}



