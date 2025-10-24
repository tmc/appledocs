// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

// PAccountAuthenticationModificationControllerPresentationContextProviding is the ASAccountAuthenticationModificationControllerPresentationContextProviding protocol interface.
//
// An interface you implement to coordinate presentation of the user interface when modifying an account’s authentication properties.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASAccountAuthenticationModificationControllerPresentationContextProviding
type PAccountAuthenticationModificationControllerPresentationContextProviding interface {
	// Required methods
	PresentationAnchorForAccountAuthenticationModificationController(controller IASAccountAuthenticationModificationController) PresentationAnchor/* debug [protocol_interface/required_method]: PresentationAnchorForAccountAuthenticationModificationController */
}
