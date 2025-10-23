// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationProviderExtensionLoginManager] class.
var (
	AuthorizationProviderExtensionLoginManagerClass     _AuthorizationProviderExtensionLoginManagerClass
	AuthorizationProviderExtensionLoginManagerClassOnce sync.Once
)

func getAuthorizationProviderExtensionLoginManagerClass() _AuthorizationProviderExtensionLoginManagerClass {
	AuthorizationProviderExtensionLoginManagerClassOnce.Do(func() {
		AuthorizationProviderExtensionLoginManagerClass = _AuthorizationProviderExtensionLoginManagerClass{objc.GetClass("ASAuthorizationProviderExtensionLoginManager")}
	})
	return AuthorizationProviderExtensionLoginManagerClass
}

type _AuthorizationProviderExtensionLoginManagerClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationProviderExtensionLoginManager] class.
type IAuthorizationProviderExtensionLoginManager interface {
	objectivec.IObject
	// properties:
	AuthenticationMethod() AuthorizationProviderExtensionAuthenticationMethod /* not a class type */
	SetAuthenticationMethod(value AuthorizationProviderExtensionAuthenticationMethod /* not a class type */)
	ExtensionData() unsafe.Pointer
	SetExtensionData(value unsafe.Pointer)
	IsDeviceRegistered() bool /* primitive/slice/pointer. */
	SetIsDeviceRegistered(value bool /* primitive/slice/pointer. */)
	IsUserRegistered() bool /* primitive/slice/pointer. */
	SetIsUserRegistered(value bool /* primitive/slice/pointer. */)
	LoginConfiguration() AuthorizationProviderExtensionLoginConfiguration /* not a class type */
	SetLoginConfiguration(value AuthorizationProviderExtensionLoginConfiguration /* not a class type */)
	LoginUserName() string /* primitive/slice/pointer. */
	SetLoginUserName(value string /* primitive/slice/pointer. */)
	RegistrationToken() string /* primitive/slice/pointer. */
	SetRegistrationToken(value string /* primitive/slice/pointer. */)
	SsoTokens() unsafe.Pointer
	SetSsoTokens(value unsafe.Pointer)
	UserLoginConfiguration() IASAuthorizationProviderExtensionUserLoginConfiguration
	SetUserLoginConfiguration(value IASAuthorizationProviderExtensionUserLoginConfiguration)
	// methods:
	BeginKeyRotationForKeyType(keyType AuthorizationProviderExtensionKeyType /* not a class type */) unsafe.Pointer
	CompleteKeyRotationForKeyType(keyType AuthorizationProviderExtensionKeyType /* not a class type */)
}

// An interface to maintain platform single sign-on (SSO) during authentication and registration.
//
// Use this class to perform registration and authentication tasks, and to repair registrations.


// An interface to maintain platform single sign-on (SSO) during authentication and registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager
type AuthorizationProviderExtensionLoginManager struct {
	objectivec.Object
}

// AuthorizationProviderExtensionLoginManagerFrom constructs a [AuthorizationProviderExtensionLoginManager] from an unsafe.Pointer.
//
// An interface to maintain platform single sign-on (SSO) during authentication and registration.
func AuthorizationProviderExtensionLoginManagerFrom(ptr unsafe.Pointer) AuthorizationProviderExtensionLoginManager {
	return AuthorizationProviderExtensionLoginManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionLoginManagerClass) Alloc() AuthorizationProviderExtensionLoginManager {
	rv := objc.Send[AuthorizationProviderExtensionLoginManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationProviderExtensionLoginManagerClass) New() AuthorizationProviderExtensionLoginManager {
	rv := objc.Send[AuthorizationProviderExtensionLoginManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationProviderExtensionLoginManager) Init() AuthorizationProviderExtensionLoginManager {
	rv := objc.Send[AuthorizationProviderExtensionLoginManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationProviderExtensionLoginManager) Autorelease() AuthorizationProviderExtensionLoginManager {
	rv := objc.Send[AuthorizationProviderExtensionLoginManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationProviderExtensionLoginManager creates a new AuthorizationProviderExtensionLoginManager instance.
func NewAuthorizationProviderExtensionLoginManager() AuthorizationProviderExtensionLoginManager {
	return getAuthorizationProviderExtensionLoginManagerClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/beginKeyRotation(_:)
func (a_ AuthorizationProviderExtensionLoginManager) BeginKeyRotationForKeyType(keyType AuthorizationProviderExtensionKeyType /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("beginKeyRotationForKeyType:"), keyType)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/completeKeyRotation(_:)
func (a_ AuthorizationProviderExtensionLoginManager) CompleteKeyRotationForKeyType(keyType AuthorizationProviderExtensionKeyType /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("completeKeyRotationForKeyType:"), keyType)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/authenticationmethod
func (a_ AuthorizationProviderExtensionLoginManager) AuthenticationMethod() AuthorizationProviderExtensionAuthenticationMethod /* not a class type */ {
	rv := objc.Send[AuthorizationProviderExtensionAuthenticationMethod](a_.ID, objc.Sel("authenticationMethod"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/authenticationmethod
func (a_ AuthorizationProviderExtensionLoginManager) SetAuthenticationMethod(value AuthorizationProviderExtensionAuthenticationMethod /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthenticationMethod:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/extensiondata
func (a_ AuthorizationProviderExtensionLoginManager) ExtensionData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("extensionData"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/extensiondata
func (a_ AuthorizationProviderExtensionLoginManager) SetExtensionData(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExtensionData:"), value)
}


// A Boolean value that indicates whether the device completes registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isdeviceregistered
func (a_ AuthorizationProviderExtensionLoginManager) IsDeviceRegistered() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isDeviceRegistered"))
	return rv
}


// A Boolean value that indicates whether the device completes registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isdeviceregistered
func (a_ AuthorizationProviderExtensionLoginManager) SetIsDeviceRegistered(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsDeviceRegistered:"), value)
}


// A Boolean value that indicates whether the user completes registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isuserregistered
func (a_ AuthorizationProviderExtensionLoginManager) IsUserRegistered() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isUserRegistered"))
	return rv
}


// A Boolean value that indicates whether the user completes registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isuserregistered
func (a_ AuthorizationProviderExtensionLoginManager) SetIsUserRegistered(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsUserRegistered:"), value)
}


// The current login configuration for the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/loginconfiguration
func (a_ AuthorizationProviderExtensionLoginManager) LoginConfiguration() AuthorizationProviderExtensionLoginConfiguration /* not a class type */ {
	rv := objc.Send[AuthorizationProviderExtensionLoginConfiguration](a_.ID, objc.Sel("loginConfiguration"))
	return rv
}


// The current login configuration for the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/loginconfiguration
func (a_ AuthorizationProviderExtensionLoginManager) SetLoginConfiguration(value AuthorizationProviderExtensionLoginConfiguration /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginConfiguration:"), value)
}


// The user name to use when authenticating with the identity provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/loginusername
func (a_ AuthorizationProviderExtensionLoginManager) LoginUserName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("loginUserName"))
	return rv
}


// The user name to use when authenticating with the identity provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/loginusername
func (a_ AuthorizationProviderExtensionLoginManager) SetLoginUserName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginUserName:"), objc.String(value))
}


// The device registration token from the mobile device management profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/registrationtoken
func (a_ AuthorizationProviderExtensionLoginManager) RegistrationToken() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("registrationToken"))
	return rv
}


// The device registration token from the mobile device management profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/registrationtoken
func (a_ AuthorizationProviderExtensionLoginManager) SetRegistrationToken(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRegistrationToken:"), objc.String(value))
}


// The single sign-on response tokens for the current user and extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/ssotokens
func (a_ AuthorizationProviderExtensionLoginManager) SsoTokens() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("ssoTokens"))
	return rv
}


// The single sign-on response tokens for the current user and extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/ssotokens
func (a_ AuthorizationProviderExtensionLoginManager) SetSsoTokens(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSsoTokens:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/userloginconfiguration
func (a_ AuthorizationProviderExtensionLoginManager) UserLoginConfiguration() IASAuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](a_.ID, objc.Sel("userLoginConfiguration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/userloginconfiguration
func (a_ AuthorizationProviderExtensionLoginManager) SetUserLoginConfiguration(value IASAuthorizationProviderExtensionUserLoginConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserLoginConfiguration:"), value)
}



