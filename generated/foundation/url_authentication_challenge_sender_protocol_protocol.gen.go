// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// PURLAuthenticationChallengeSender is the NSURLAuthenticationChallengeSender protocol interface.
//
// The   protocol represents the interface that the sender of an authentication challenge must implement.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/URLAuthenticationChallengeSender
type PURLAuthenticationChallengeSender interface {
	// Required methods
	CancelAuthenticationChallenge(challenge IURLAuthenticationChallenge)/* debug [protocol_interface/required_method]: CancelAuthenticationChallenge */
	ContinueWithoutCredentialForAuthenticationChallenge(challenge IURLAuthenticationChallenge)/* debug [protocol_interface/required_method]: ContinueWithoutCredentialForAuthenticationChallenge */
	UseCredentialForAuthenticationChallenge(credential IURLCredential, challenge IURLAuthenticationChallenge)/* debug [protocol_interface/required_method]: UseCredentialForAuthenticationChallenge */
	// Optional methods
	PerformDefaultHandlingForAuthenticationChallenge(challenge IURLAuthenticationChallenge)
	HasPerformDefaultHandlingForAuthenticationChallenge() bool
	RejectProtectionSpaceAndContinueWithChallenge(challenge IURLAuthenticationChallenge)
	HasRejectProtectionSpaceAndContinueWithChallenge() bool
}
