// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationProviderExtensionLoginManager */


/* debug [class_header]: Header for ASAuthorizationProviderExtensionLoginManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationProviderExtensionLoginManager */
// An interface definition for the [AuthorizationProviderExtensionLoginManager] class.
type IAuthorizationProviderExtensionLoginManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationProviderExtensionLoginManager */
	// properties:
	AuthenticationMethod() AuthorizationProviderExtensionAuthenticationMethod
	ExtensionData() objc.IObject /* cross-framework: NSDictionary */
	DeviceRegistered() bool
	UserRegistered() bool
	LoginConfiguration() IASAuthorizationProviderExtensionLoginConfiguration
	LoginUserName() objc.IObject /* cross-framework: NSString */
	SetLoginUserName(value objc.IObject /* cross-framework: NSString */)
	RegistrationToken() objc.IObject /* cross-framework: NSString */
	SsoTokens() objc.IObject /* cross-framework: NSDictionary */
	SetSsoTokens(value objc.IObject /* cross-framework: NSDictionary */)
	UserLoginConfiguration() IASAuthorizationProviderExtensionUserLoginConfiguration
	IsDeviceRegistered() bool
	SetIsDeviceRegistered(value bool)
	IsUserRegistered() bool
	SetIsUserRegistered(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationProviderExtensionLoginManager */
	// methods:
	AttestKeyClientDataHashCompletion(keyType AuthorizationProviderExtensionKeyType, clientDataHash objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer)
	AttestPendingKeyClientDataHashCompletion(keyType AuthorizationProviderExtensionKeyType, clientDataHash objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer)
	BeginKeyRotationForKeyType(keyType AuthorizationProviderExtensionKeyType) unsafe.Pointer
	CompleteKeyRotationForKeyType(keyType AuthorizationProviderExtensionKeyType)
	DecryptionKeysNeedRepair()
	DeviceRegistrationsNeedsRepair()
	CopyIdentityForKeyType(keyType AuthorizationProviderExtensionKeyType) unsafe.Pointer
	CopyKeyForKeyType(keyType AuthorizationProviderExtensionKeyType) unsafe.Pointer
	PresentRegistrationViewControllerWithCompletion(completion unsafe.Pointer)
	ResetDeviceKeys()
	ResetKeys()
	ResetUserSecureEnclaveKey()
	SaveCertificateKeyType(certificate unsafe.Pointer, keyType AuthorizationProviderExtensionKeyType)
	SaveLoginConfigurationError(loginConfiguration IASAuthorizationProviderExtensionLoginConfiguration, error_ unsafe.Pointer) bool
	SaveUserLoginConfigurationError(userLoginConfiguration IASAuthorizationProviderExtensionUserLoginConfiguration, error_ unsafe.Pointer) bool
	UserNeedsReauthenticationWithCompletion(completion unsafe.Pointer)
	UserRegistrationsNeedsRepair()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationProviderExtensionLoginManager */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionLoginManagerClass) Alloc() AuthorizationProviderExtensionLoginManager {
	rv := objc.Send[AuthorizationProviderExtensionLoginManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationProviderExtensionLoginManager */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationProviderExtensionLoginManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationProviderExtensionLoginManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationProviderExtensionLoginManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationProviderExtensionLoginManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/attestKey:clientDataHash:completion:
func (a_ AuthorizationProviderExtensionLoginManager) AttestKeyClientDataHashCompletion(keyType AuthorizationProviderExtensionKeyType, clientDataHash objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("attestKey:clientDataHash:completion:"), keyType, clientDataHash, completion)
}/* debug [instance_methods/method]: AttestKeyClientDataHashCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/attestPendingKey:clientDataHash:completion:
func (a_ AuthorizationProviderExtensionLoginManager) AttestPendingKeyClientDataHashCompletion(keyType AuthorizationProviderExtensionKeyType, clientDataHash objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("attestPendingKey:clientDataHash:completion:"), keyType, clientDataHash, completion)
}/* debug [instance_methods/method]: AttestPendingKeyClientDataHashCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/beginKeyRotation(_:)
func (a_ AuthorizationProviderExtensionLoginManager) BeginKeyRotationForKeyType(keyType AuthorizationProviderExtensionKeyType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("beginKeyRotationForKeyType:"), keyType)
	return rv
}/* debug [instance_methods/method]: BeginKeyRotationForKeyType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/completeKeyRotation(_:)
func (a_ AuthorizationProviderExtensionLoginManager) CompleteKeyRotationForKeyType(keyType AuthorizationProviderExtensionKeyType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("completeKeyRotationForKeyType:"), keyType)
}/* debug [instance_methods/method]: CompleteKeyRotationForKeyType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/decryptionKeysNeedRepair()
func (a_ AuthorizationProviderExtensionLoginManager) DecryptionKeysNeedRepair() {
	objc.Send[objc.ID](a_.ID, objc.Sel("decryptionKeysNeedRepair"))
}/* debug [instance_methods/method]: DecryptionKeysNeedRepair */


// Invokes the device registration to run again so the current user can repair it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/deviceRegistrationsNeedsRepair()
func (a_ AuthorizationProviderExtensionLoginManager) DeviceRegistrationsNeedsRepair() {
	objc.Send[objc.ID](a_.ID, objc.Sel("deviceRegistrationsNeedsRepair"))
}/* debug [instance_methods/method]: DeviceRegistrationsNeedsRepair */


// Retrieves the identity for the specified platform single sign-on key type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/identity(for:)
func (a_ AuthorizationProviderExtensionLoginManager) CopyIdentityForKeyType(keyType AuthorizationProviderExtensionKeyType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("copyIdentityForKeyType:"), keyType)
	return rv
}/* debug [instance_methods/method]: CopyIdentityForKeyType */


// Retrieves the key for the specified platform single sign-on key type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/key(for:)
func (a_ AuthorizationProviderExtensionLoginManager) CopyKeyForKeyType(keyType AuthorizationProviderExtensionKeyType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("copyKeyForKeyType:"), keyType)
	return rv
}/* debug [instance_methods/method]: CopyKeyForKeyType */


// Requests platform single sign-on to show the extension’s view controller to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/presentRegistrationViewController(completion:)
func (a_ AuthorizationProviderExtensionLoginManager) PresentRegistrationViewControllerWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("presentRegistrationViewControllerWithCompletion:"), completion)
}/* debug [instance_methods/method]: PresentRegistrationViewControllerWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/resetDeviceKeys()
func (a_ AuthorizationProviderExtensionLoginManager) ResetDeviceKeys() {
	objc.Send[objc.ID](a_.ID, objc.Sel("resetDeviceKeys"))
}/* debug [instance_methods/method]: ResetDeviceKeys */


// Creates new encryption, signing, and Secure Enclave keys for the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/resetKeys()
func (a_ AuthorizationProviderExtensionLoginManager) ResetKeys() {
	objc.Send[objc.ID](a_.ID, objc.Sel("resetKeys"))
}/* debug [instance_methods/method]: ResetKeys */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/resetUserSecureEnclaveKey()
func (a_ AuthorizationProviderExtensionLoginManager) ResetUserSecureEnclaveKey() {
	objc.Send[objc.ID](a_.ID, objc.Sel("resetUserSecureEnclaveKey"))
}/* debug [instance_methods/method]: ResetUserSecureEnclaveKey */


// Saves the provided certificate for the key type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/saveCertificate(_:keyType:)
func (a_ AuthorizationProviderExtensionLoginManager) SaveCertificateKeyType(certificate unsafe.Pointer, keyType AuthorizationProviderExtensionKeyType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("saveCertificate:keyType:"), certificate, keyType)
}/* debug [instance_methods/method]: SaveCertificateKeyType */


// Saves or replaces the login configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/saveLoginConfiguration(_:)
func (a_ AuthorizationProviderExtensionLoginManager) SaveLoginConfigurationError(loginConfiguration IASAuthorizationProviderExtensionLoginConfiguration, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("saveLoginConfiguration:error:"), loginConfiguration, error_)
	return rv
}/* debug [instance_methods/method]: SaveLoginConfigurationError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/saveUserLoginConfiguration(_:)
func (a_ AuthorizationProviderExtensionLoginManager) SaveUserLoginConfigurationError(userLoginConfiguration IASAuthorizationProviderExtensionUserLoginConfiguration, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("saveUserLoginConfiguration:error:"), userLoginConfiguration, error_)
	return rv
}/* debug [instance_methods/method]: SaveUserLoginConfigurationError */


// Requests platform single sign-on to reauthenticate the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/userNeedsReauthentication(completion:)
func (a_ AuthorizationProviderExtensionLoginManager) UserNeedsReauthenticationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("userNeedsReauthenticationWithCompletion:"), completion)
}/* debug [instance_methods/method]: UserNeedsReauthenticationWithCompletion */


// Invokes the user registration to run again so the current user can repair it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/userRegistrationsNeedsRepair()
func (a_ AuthorizationProviderExtensionLoginManager) UserRegistrationsNeedsRepair() {
	objc.Send[objc.ID](a_.ID, objc.Sel("userRegistrationsNeedsRepair"))
}/* debug [instance_methods/method]: UserRegistrationsNeedsRepair */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationProviderExtensionLoginManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/authenticationMethod
func (a_ AuthorizationProviderExtensionLoginManager) AuthenticationMethod() AuthorizationProviderExtensionAuthenticationMethod {
	rv := objc.Send[AuthorizationProviderExtensionAuthenticationMethod](a_.ID, objc.Sel("authenticationMethod"))
	return rv
}/* debug [instance_properties/getter]: authenticationMethod */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/extensionData
func (a_ AuthorizationProviderExtensionLoginManager) ExtensionData() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](a_.ID, objc.Sel("extensionData"))
	return rv
}/* debug [instance_properties/getter]: extensionData */


// A Boolean value that indicates whether the device completes registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/isDeviceRegistered
func (a_ AuthorizationProviderExtensionLoginManager) DeviceRegistered() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("deviceRegistered"))
	return rv
}/* debug [instance_properties/getter]: deviceRegistered */


// A Boolean value that indicates whether the user completes registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/isUserRegistered
func (a_ AuthorizationProviderExtensionLoginManager) UserRegistered() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("userRegistered"))
	return rv
}/* debug [instance_properties/getter]: userRegistered */


// The current login configuration for the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/loginConfiguration
func (a_ AuthorizationProviderExtensionLoginManager) LoginConfiguration() IASAuthorizationProviderExtensionLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionLoginConfiguration](a_.ID, objc.Sel("loginConfiguration"))
	return rv
}/* debug [instance_properties/getter]: loginConfiguration */


// The user name to use when authenticating with the identity provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/loginUserName
func (a_ AuthorizationProviderExtensionLoginManager) LoginUserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("loginUserName"))
	return rv
}/* debug [instance_properties/getter]: loginUserName */


// The user name to use when authenticating with the identity provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/loginUserName
func (a_ AuthorizationProviderExtensionLoginManager) SetLoginUserName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoginUserName:"), value)
}/* debug [instance_properties/setter]: loginUserName */


// The device registration token from the mobile device management profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/registrationToken
func (a_ AuthorizationProviderExtensionLoginManager) RegistrationToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("registrationToken"))
	return rv
}/* debug [instance_properties/getter]: registrationToken */


// The single sign-on response tokens for the current user and extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/ssoTokens
func (a_ AuthorizationProviderExtensionLoginManager) SsoTokens() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](a_.ID, objc.Sel("ssoTokens"))
	return rv
}/* debug [instance_properties/getter]: ssoTokens */


// The single sign-on response tokens for the current user and extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/ssoTokens
func (a_ AuthorizationProviderExtensionLoginManager) SetSsoTokens(value objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSsoTokens:"), value)
}/* debug [instance_properties/setter]: ssoTokens */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/userLoginConfiguration
func (a_ AuthorizationProviderExtensionLoginManager) UserLoginConfiguration() IASAuthorizationProviderExtensionUserLoginConfiguration {
	rv := objc.Send[AuthorizationProviderExtensionUserLoginConfiguration](a_.ID, objc.Sel("userLoginConfiguration"))
	return rv
}/* debug [instance_properties/getter]: userLoginConfiguration */


// A Boolean value that indicates whether the device completes registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isdeviceregistered
func (a_ AuthorizationProviderExtensionLoginManager) IsDeviceRegistered() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isDeviceRegistered"))
	return rv
}/* debug [instance_properties/getter]: isDeviceRegistered */


// A Boolean value that indicates whether the device completes registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isdeviceregistered
func (a_ AuthorizationProviderExtensionLoginManager) SetIsDeviceRegistered(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsDeviceRegistered:"), value)
}/* debug [instance_properties/setter]: isDeviceRegistered */


// A Boolean value that indicates whether the user completes registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isuserregistered
func (a_ AuthorizationProviderExtensionLoginManager) IsUserRegistered() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isUserRegistered"))
	return rv
}/* debug [instance_properties/getter]: isUserRegistered */


// A Boolean value that indicates whether the user completes registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionloginmanager/isuserregistered
func (a_ AuthorizationProviderExtensionLoginManager) SetIsUserRegistered(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsUserRegistered:"), value)
}/* debug [instance_properties/setter]: isUserRegistered */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationProviderExtensionLoginManager */



