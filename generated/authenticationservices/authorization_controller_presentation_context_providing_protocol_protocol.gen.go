// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

// PAuthorizationControllerPresentationContextProviding is the ASAuthorizationControllerPresentationContextProviding protocol interface.
//
// An interface the controller uses to ask a delegate for a presentation context.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASAuthorizationControllerPresentationContextProviding
type PAuthorizationControllerPresentationContextProviding interface {
	// Required methods
	PresentationAnchorForAuthorizationController(controller IASAuthorizationController) PresentationAnchor/* debug [protocol_interface/required_method]: PresentationAnchorForAuthorizationController */
}
