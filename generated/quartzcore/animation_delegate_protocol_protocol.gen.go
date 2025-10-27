// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

// PAnimationDelegate is the CAAnimationDelegate protocol interface.
//
// Methods your app can implement to respond when animations start and stop.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.quartzcore/documentation/QuartzCore/CAAnimationDelegate
type PAnimationDelegate interface {
	// Optional methods
	AnimationDidStart(anim IAnimation)
	HasAnimationDidStart() bool
	AnimationDidStopFinished(anim IAnimation, flag bool)
	HasAnimationDidStopFinished() bool
}
