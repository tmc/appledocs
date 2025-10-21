// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationSingleSignOnProvider] class.
var (
	AuthorizationSingleSignOnProviderClass     _AuthorizationSingleSignOnProviderClass
	AuthorizationSingleSignOnProviderClassOnce sync.Once
)

func getAuthorizationSingleSignOnProviderClass() _AuthorizationSingleSignOnProviderClass {
	AuthorizationSingleSignOnProviderClassOnce.Do(func() {
		AuthorizationSingleSignOnProviderClass = _AuthorizationSingleSignOnProviderClass{objc.GetClass("ASAuthorizationSingleSignOnProvider")}
	})
	return AuthorizationSingleSignOnProviderClass
}

type _AuthorizationSingleSignOnProviderClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationSingleSignOnProvider] class.
type IAuthorizationSingleSignOnProvider interface {
	objectivec.IObject
	CreateRequest() unsafe.Pointer
}

// A mechanism for generating requests to authenticate users with third-party providers.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnProvider
type AuthorizationSingleSignOnProvider struct {
	objectivec.Object
}

// AuthorizationSingleSignOnProviderFrom constructs a [AuthorizationSingleSignOnProvider] from an unsafe.Pointer.
//
// A mechanism for generating requests to authenticate users with third-party providers.
func AuthorizationSingleSignOnProviderFrom(ptr unsafe.Pointer) AuthorizationSingleSignOnProvider {
	return AuthorizationSingleSignOnProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSingleSignOnProviderClass) Alloc() AuthorizationSingleSignOnProvider {
	rv := objc.Send[AuthorizationSingleSignOnProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationSingleSignOnProviderClass) New() AuthorizationSingleSignOnProvider {
	rv := objc.Send[AuthorizationSingleSignOnProvider](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSingleSignOnProvider) Init() AuthorizationSingleSignOnProvider {
	rv := objc.Send[AuthorizationSingleSignOnProvider](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSingleSignOnProvider) Autorelease() AuthorizationSingleSignOnProvider {
	rv := objc.Send[AuthorizationSingleSignOnProvider](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSingleSignOnProvider creates a new AuthorizationSingleSignOnProvider instance.
func NewAuthorizationSingleSignOnProvider() AuthorizationSingleSignOnProvider {
	return getAuthorizationSingleSignOnProviderClass().New()
}


// Creates a single sign-on (SSO) authorization request.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnProvider/createRequest()
func (a_ AuthorizationSingleSignOnProvider) CreateRequest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("createRequest"))
	return rv
}

// A Boolean value that indicates if the provider is capable of performing authorization within a given configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnProvider/canPerformAuthorization
func (a_ AuthorizationSingleSignOnProvider) CanPerformAuthorization() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformAuthorization"))
	return rv
}

// The URL of the identity provider.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignonprovider/url
func (a_ AuthorizationSingleSignOnProvider) Url() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL of the identity provider.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignonprovider/url
func (a_ AuthorizationSingleSignOnProvider) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUrl:"), value)
}



