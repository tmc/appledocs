// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PHapticFeedbackPerformer is the NSHapticFeedbackPerformer protocol interface.
//
// A set of methods and constants that a haptic feedback performer implements.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSHapticFeedbackPerformer
type PHapticFeedbackPerformer interface {
	// Required methods
	PerformFeedbackPatternPerformanceTime(pattern HapticFeedbackPattern, performanceTime HapticFeedbackPerformanceTime)
}
