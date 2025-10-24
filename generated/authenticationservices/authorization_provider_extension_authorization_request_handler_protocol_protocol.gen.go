// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

// PAuthorizationProviderExtensionAuthorizationRequestHandler is the ASAuthorizationProviderExtensionAuthorizationRequestHandler protocol interface.
//
// An interface through which a single sign-on (SSO) authentication provider extension handles authentication requests.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequestHandler
type PAuthorizationProviderExtensionAuthorizationRequestHandler interface {
	// Required methods
	BeginAuthorizationWithRequest(request IASAuthorizationProviderExtensionAuthorizationRequest)/* debug [protocol_interface/required_method]: BeginAuthorizationWithRequest */
	// Optional methods
	CancelAuthorizationWithRequest(request IASAuthorizationProviderExtensionAuthorizationRequest)
	HasCancelAuthorizationWithRequest() bool
}
