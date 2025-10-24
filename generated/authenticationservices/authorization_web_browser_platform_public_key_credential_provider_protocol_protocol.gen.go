// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PAuthorizationWebBrowserPlatformPublicKeyCredentialProvider is the ASAuthorizationWebBrowserPlatformPublicKeyCredentialProvider protocol interface.
//
// A protocol for creating passkey requests.
//
// Availability:
//   - Mac Catalyst 16.6+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - macOS 13.5+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASAuthorizationWebBrowserPlatformPublicKeyCredentialProvider-1c8cl
type PAuthorizationWebBrowserPlatformPublicKeyCredentialProvider interface {
	// Required methods
	CreateCredentialAssertionRequestWithClientData(clientData IASPublicKeyCredentialClientData) AuthorizationPlatformPublicKeyCredentialAssertionRequest/* debug [protocol_interface/required_method]: CreateCredentialAssertionRequestWithClientData */
	CreateCredentialRegistrationRequestWithClientDataNameUserID(clientData IASPublicKeyCredentialClientData, name objc.IObject /* cross-framework: NSString */, userID objc.IObject /* cross-framework: NSData */) AuthorizationPlatformPublicKeyCredentialRegistrationRequest/* debug [protocol_interface/required_method]: CreateCredentialRegistrationRequestWithClientDataNameUserID */
	CreateCredentialRegistrationRequestWithClientDataNameUserIDRequestStyle(clientData IASPublicKeyCredentialClientData, name objc.IObject /* cross-framework: NSString */, userID objc.IObject /* cross-framework: NSData */, requestStyle AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle) AuthorizationPlatformPublicKeyCredentialRegistrationRequest/* debug [protocol_interface/required_method]: CreateCredentialRegistrationRequestWithClientDataNameUserIDRequestStyle */
}
