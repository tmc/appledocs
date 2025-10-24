// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

// PWebAuthenticationSessionWebBrowserSessionHandling is the ASWebAuthenticationSessionWebBrowserSessionHandling protocol interface.
//
// An interface that a session handler implements to handle login requests from an app.
//
// Availability:
//   - macOS 10.15+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASWebAuthenticationSessionWebBrowserSessionHandling
type PWebAuthenticationSessionWebBrowserSessionHandling interface {
	// Required methods
	BeginHandlingWebAuthenticationSessionRequest(request IASWebAuthenticationSessionRequest)/* debug [protocol_interface/required_method]: BeginHandlingWebAuthenticationSessionRequest */
	CancelWebAuthenticationSessionRequest(request IASWebAuthenticationSessionRequest)/* debug [protocol_interface/required_method]: CancelWebAuthenticationSessionRequest */
}
