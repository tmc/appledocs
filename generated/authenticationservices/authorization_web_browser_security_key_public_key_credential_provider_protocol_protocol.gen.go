// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PAuthorizationWebBrowserSecurityKeyPublicKeyCredentialProvider is the ASAuthorizationWebBrowserSecurityKeyPublicKeyCredentialProvider protocol interface.
//
// Availability:
//   - Mac Catalyst 17.4+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - macOS 14.4+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASAuthorizationWebBrowserSecurityKeyPublicKeyCredentialProvider-4gzot
type PAuthorizationWebBrowserSecurityKeyPublicKeyCredentialProvider interface {
	// Required methods
	CreateCredentialAssertionRequestWithClientData(clientData IASPublicKeyCredentialClientData) AuthorizationSecurityKeyPublicKeyCredentialAssertionRequest/* debug [protocol_interface/required_method]: CreateCredentialAssertionRequestWithClientData */
	CreateCredentialRegistrationRequestWithClientDataDisplayNameNameUserID(clientData IASPublicKeyCredentialClientData, displayName objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */, userID objc.IObject /* cross-framework: NSData */) AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest/* debug [protocol_interface/required_method]: CreateCredentialRegistrationRequestWithClientDataDisplayNameNameUserID */
}
