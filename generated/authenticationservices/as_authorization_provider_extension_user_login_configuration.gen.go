// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationProviderExtensionUserLoginConfiguration] class.
var (
	AuthorizationProviderExtensionUserLoginConfigurationClass     _AuthorizationProviderExtensionUserLoginConfigurationClass
	AuthorizationProviderExtensionUserLoginConfigurationClassOnce sync.Once
)

func getAuthorizationProviderExtensionUserLoginConfigurationClass() _AuthorizationProviderExtensionUserLoginConfigurationClass {
	AuthorizationProviderExtensionUserLoginConfigurationClassOnce.Do(func() {
		AuthorizationProviderExtensionUserLoginConfigurationClass = _AuthorizationProviderExtensionUserLoginConfigurationClass{objc.GetClass("ASAuthorizationProviderExtensionUserLoginConfiguration")}
	})
	return AuthorizationProviderExtensionUserLoginConfigurationClass
}

type _AuthorizationProviderExtensionUserLoginConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationProviderExtensionUserLoginConfiguration] class.
type IAuthorizationProviderExtensionUserLoginConfiguration interface {
	objectivec.IObject
	SetCustomAssertionRequestHeaderClaimsReturningError(claims unsafe.Pointer, error_ unsafe.Pointer) bool
	LoginUserName() string
	SetLoginUserName(value string)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserLoginConfiguration

type AuthorizationProviderExtensionUserLoginConfiguration struct {
	objectivec.Object
}

// AuthorizationProviderExtensionUserLoginConfigurationFrom constructs a [AuthorizationProviderExtensionUserLoginConfiguration] from an unsafe.Pointer.
func AuthorizationProviderExtensionUserLoginConfigurationFrom(ptr unsafe.Pointer) AuthorizationProviderExtensionUserLoginConfiguration {
	return AuthorizationProviderExtensionUserLoginConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionUserLoginConfigurationClass) Alloc() AuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationProviderExtensionUserLoginConfigurationClass) New() AuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationProviderExtensionUserLoginConfiguration) Init() AuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationProviderExtensionUserLoginConfiguration) Autorelease() AuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationProviderExtensionUserLoginConfiguration creates a new AuthorizationProviderExtensionUserLoginConfiguration instance.
func NewAuthorizationProviderExtensionUserLoginConfiguration() AuthorizationProviderExtensionUserLoginConfiguration {
	return getAuthorizationProviderExtensionUserLoginConfigurationClass().New()
}




// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserLoginConfiguration/setCustomAssertionRequestHeaderClaims(_:)

func (a_ AuthorizationProviderExtensionUserLoginConfiguration) SetCustomAssertionRequestHeaderClaimsReturningError(claims unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCustomAssertionRequestHeaderClaims:returningError:"), claims, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionuserloginconfiguration/loginusername

func (a_ AuthorizationProviderExtensionUserLoginConfiguration) LoginUserName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("loginUserName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionuserloginconfiguration/loginusername

func (a_ AuthorizationProviderExtensionUserLoginConfiguration) SetLoginUserName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginUserName:"), objc.String(value))
}



