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
	BeginKeyRotationForKeyType(keyType unsafe.Pointer) unsafe.Pointer
	CompleteKeyRotationForKeyType(keyType unsafe.Pointer)
	SaveCertificateKeyType(certificate unsafe.Pointer, keyType unsafe.Pointer)
	AuthenticationMethod() unsafe.Pointer
	SetAuthenticationMethod(value unsafe.Pointer)
	ExtensionData() unsafe.Pointer
	SetExtensionData(value unsafe.Pointer)
	IsDeviceRegistered() bool
	SetIsDeviceRegistered(value bool)
	IsUserRegistered() bool
	SetIsUserRegistered(value bool)
	LoginConfiguration() unsafe.Pointer
	SetLoginConfiguration(value unsafe.Pointer)
	LoginUserName() string
	SetLoginUserName(value string)
	RegistrationToken() string
	SetRegistrationToken(value string)
	SsoTokens() unsafe.Pointer
	SetSsoTokens(value unsafe.Pointer)
	UserLoginConfiguration() ASAuthorizationProviderExtensionUserLoginConfiguration
	SetUserLoginConfiguration(value IASAuthorizationProviderExtensionUserLoginConfiguration)
}

// An interface to maintain platform single sign-on (SSO) during authentication and registration.
//
// Use this class to perform registration and authentication tasks, and to repair registrations.
//
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


//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/beginKeyRotation(_:)
func (a_ AuthorizationProviderExtensionLoginManager) BeginKeyRotationForKeyType(keyType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("beginKeyRotationForKeyType:"), keyType)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/completeKeyRotation(_:)
func (a_ AuthorizationProviderExtensionLoginManager) CompleteKeyRotationForKeyType(keyType unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("completeKeyRotationForKeyType:"), keyType)
}

// Saves the provided certificate for the key type.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/saveCertificate(_:keyType:)
func (a_ AuthorizationProviderExtensionLoginManager) SaveCertificateKeyType(certificate unsafe.Pointer, keyType unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("saveCertificate:keyType:"), certificate, keyType)
}

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/authenticationmethod
func (a_ AuthorizationProviderExtensionLoginManager) AuthenticationMethod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("authenticationMethod"))
	return rv
}


// SetAuthenticationMethod sets the value of the authenticationMethod property.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/authenticationmethod
func (a_ AuthorizationProviderExtensionLoginManager) SetAuthenticationMethod(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthenticationMethod:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/extensiondata
func (a_ AuthorizationProviderExtensionLoginManager) ExtensionData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("extensionData"))
	return rv
}


// SetExtensionData sets the value of the extensionData property.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/extensiondata
func (a_ AuthorizationProviderExtensionLoginManager) SetExtensionData(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExtensionData:"), value)
}

// A Boolean value that indicates whether the device completes registration.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isdeviceregistered
func (a_ AuthorizationProviderExtensionLoginManager) IsDeviceRegistered() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isDeviceRegistered"))
	return rv
}


// SetIsDeviceRegistered sets the value of the isDeviceRegistered property.
// A Boolean value that indicates whether the device completes registration.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isdeviceregistered
func (a_ AuthorizationProviderExtensionLoginManager) SetIsDeviceRegistered(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsDeviceRegistered:"), value)
}

// A Boolean value that indicates whether the user completes registration.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isuserregistered
func (a_ AuthorizationProviderExtensionLoginManager) IsUserRegistered() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isUserRegistered"))
	return rv
}


// SetIsUserRegistered sets the value of the isUserRegistered property.
// A Boolean value that indicates whether the user completes registration.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isuserregistered
func (a_ AuthorizationProviderExtensionLoginManager) SetIsUserRegistered(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsUserRegistered:"), value)
}

// The current login configuration for the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/loginconfiguration
func (a_ AuthorizationProviderExtensionLoginManager) LoginConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("loginConfiguration"))
	return rv
}


// SetLoginConfiguration sets the value of the loginConfiguration property.
// The current login configuration for the extension.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/loginconfiguration
func (a_ AuthorizationProviderExtensionLoginManager) SetLoginConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginConfiguration:"), value)
}

// The user name to use when authenticating with the identity provider.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/loginusername
func (a_ AuthorizationProviderExtensionLoginManager) LoginUserName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("loginUserName"))
	return rv
}


// SetLoginUserName sets the value of the loginUserName property.
// The user name to use when authenticating with the identity provider.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/loginusername
func (a_ AuthorizationProviderExtensionLoginManager) SetLoginUserName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginUserName:"), objc.String(value))
}

// The device registration token from the mobile device management profile.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/registrationtoken
func (a_ AuthorizationProviderExtensionLoginManager) RegistrationToken() string {
	rv := objc.Send[string](a_.ID, objc.Sel("registrationToken"))
	return rv
}


// SetRegistrationToken sets the value of the registrationToken property.
// The device registration token from the mobile device management profile.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/registrationtoken
func (a_ AuthorizationProviderExtensionLoginManager) SetRegistrationToken(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRegistrationToken:"), objc.String(value))
}

// The single sign-on response tokens for the current user and extension.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/ssotokens
func (a_ AuthorizationProviderExtensionLoginManager) SsoTokens() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("ssoTokens"))
	return rv
}


// SetSsoTokens sets the value of the ssoTokens property.
// The single sign-on response tokens for the current user and extension.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/ssotokens
func (a_ AuthorizationProviderExtensionLoginManager) SetSsoTokens(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSsoTokens:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/userloginconfiguration
func (a_ AuthorizationProviderExtensionLoginManager) UserLoginConfiguration() ASAuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[ASAuthorizationProviderExtensionUserLoginConfiguration](a_.ID, objc.Sel("userLoginConfiguration"))
	return rv
}


// SetUserLoginConfiguration sets the value of the userLoginConfiguration property.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/userloginconfiguration
func (a_ AuthorizationProviderExtensionLoginManager) SetUserLoginConfiguration(value IASAuthorizationProviderExtensionUserLoginConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserLoginConfiguration:"), value)
}



