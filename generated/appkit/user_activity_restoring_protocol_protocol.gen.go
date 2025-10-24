// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PUserActivityRestoring is the NSUserActivityRestoring protocol interface.
//
// A protocol that marks classes to restore the state of your app to continue a user activity.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSUserActivityRestoring
type PUserActivityRestoring interface {
	// Required methods
	RestoreUserActivityState(userActivity foundation.UserActivity)
}
