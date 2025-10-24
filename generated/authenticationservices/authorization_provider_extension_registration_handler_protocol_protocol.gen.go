// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PAuthorizationProviderExtensionRegistrationHandler is the ASAuthorizationProviderExtensionRegistrationHandler protocol interface.
//
// An interface through which a single sign-on (SSO) authentication provider extension registers users and devices for platform SSO.
//
// Availability:
//   - macOS 13.0+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRegistrationHandler
type PAuthorizationProviderExtensionRegistrationHandler interface {
	// Required methods
	BeginDeviceRegistrationUsingLoginManagerOptionsCompletion(loginManager IASAuthorizationProviderExtensionLoginManager, options AuthorizationProviderExtensionRequestOptions, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: BeginDeviceRegistrationUsingLoginManagerOptionsCompletion */
	BeginUserRegistrationUsingLoginManagerUserNameAuthenticationMethodOptionsCompletion(loginManager IASAuthorizationProviderExtensionLoginManager, userName objc.IObject /* cross-framework: NSString */, authenticationMethod AuthorizationProviderExtensionAuthenticationMethod, options AuthorizationProviderExtensionRequestOptions, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: BeginUserRegistrationUsingLoginManagerUserNameAuthenticationMethodOptionsCompletion */
	// Optional methods
	DisplayNamesForGroupsLoginManagerCompletion(groups []string, loginManager IASAuthorizationProviderExtensionLoginManager, completion unsafe.Pointer)
	HasDisplayNamesForGroupsLoginManagerCompletion() bool
	KeyWillRotateForKeyTypeNewKeyLoginManagerCompletion(keyType AuthorizationProviderExtensionKeyType, newKey unsafe.Pointer, loginManager IASAuthorizationProviderExtensionLoginManager, completion unsafe.Pointer)
	HasKeyWillRotateForKeyTypeNewKeyLoginManagerCompletion() bool
	ProfilePictureForUserUsingLoginManagerCompletion(loginManager IASAuthorizationProviderExtensionLoginManager, completion unsafe.Pointer)
	HasProfilePictureForUserUsingLoginManagerCompletion() bool
	ProtocolVersion() AuthorizationProviderExtensionPlatformSSOProtocolVersion
	HasProtocolVersion() bool
	RegistrationDidCancel()
	HasRegistrationDidCancel() bool
	RegistrationDidComplete()
	HasRegistrationDidComplete() bool
	SupportedGrantTypes() AuthorizationProviderExtensionSupportedGrantTypes
	HasSupportedGrantTypes() bool
}
