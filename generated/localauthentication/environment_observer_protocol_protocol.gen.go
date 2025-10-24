// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

// PEnvironmentObserver is the LAEnvironmentObserver protocol interface.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - visionOS 2.0+
//   - watchOS 11.0+
//
// See: doc://com.apple.localauthentication/documentation/LocalAuthentication/LAEnvironment/Observer
type PEnvironmentObserver interface {
	// Optional methods
	EnvironmentStateDidChangeFromOldState(environment ILAEnvironment, oldState ILAEnvironmentState)
	HasEnvironmentStateDidChangeFromOldState() bool
}
