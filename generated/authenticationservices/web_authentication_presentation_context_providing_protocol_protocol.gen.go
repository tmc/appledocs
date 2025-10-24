// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

// PWebAuthenticationPresentationContextProviding is the ASWebAuthenticationPresentationContextProviding protocol interface.
//
// An interface the session uses to ask a delegate for a presentation context.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASWebAuthenticationPresentationContextProviding
type PWebAuthenticationPresentationContextProviding interface {
	// Required methods
	PresentationAnchorForWebAuthenticationSession(session IASWebAuthenticationSession) PresentationAnchor/* debug [protocol_interface/required_method]: PresentationAnchorForWebAuthenticationSession */
}
