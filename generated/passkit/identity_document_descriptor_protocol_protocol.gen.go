// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

// PIdentityDocumentDescriptor is the PKIdentityDocumentDescriptor protocol interface.
//
// A type that describes the structure and behavior of an identity document.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.passkit/documentation/PassKit/PKIdentityDocumentDescriptor
type PIdentityDocumentDescriptor interface {
	// Required methods
	AddElementsWithIntentToStore(elements []IdentityElement /* not a class type */, intentToStore IdentityIntentToStore /* not a class type */)/* debug [protocol_interface/required_method]: AddElementsWithIntentToStore */
	IntentToStoreForElement(element IdentityElement /* not a class type */) IdentityIntentToStore/* debug [protocol_interface/required_method]: IntentToStoreForElement */
}
