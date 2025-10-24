// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (

	"github.com/tmc/appledocs/generated/objc"
)

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

// AnimationDelegate is a delegate implementation builder for the PAnimationDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AnimationDelegate struct {
	_AnimationDidStart func(anim IAnimation)
	_AnimationDidStopFinished func(anim IAnimation, flag bool)
}

// SetAnimationDidStart sets the handler for the AnimationDidStart delegate method.
//
// Tells the delegate the animation has started.
func (d *AnimationDelegate) SetAnimationDidStart(f func(anim IAnimation)) {
	d._AnimationDidStart = f
}

// SetAnimationDidStopFinished sets the handler for the AnimationDidStopFinished delegate method.
//
// Tells the delegate the animation has ended.
func (d *AnimationDelegate) SetAnimationDidStopFinished(f func(anim IAnimation, flag bool)) {
	d._AnimationDidStopFinished = f
}

// AnimationDidStart implements the PAnimationDelegate interface.
func (d *AnimationDelegate) AnimationDidStart(anim IAnimation) {
	if d._AnimationDidStart != nil {
		d._AnimationDidStart(anim)
	}
}

// HasAnimationDidStart returns true if a handler for AnimationDidStart has been set.
func (d *AnimationDelegate) HasAnimationDidStart() bool {
	return d._AnimationDidStart != nil
}

// AnimationDidStopFinished implements the PAnimationDelegate interface.
func (d *AnimationDelegate) AnimationDidStopFinished(anim IAnimation, flag bool) {
	if d._AnimationDidStopFinished != nil {
		d._AnimationDidStopFinished(anim, flag)
	}
}

// HasAnimationDidStopFinished returns true if a handler for AnimationDidStopFinished has been set.
func (d *AnimationDelegate) HasAnimationDidStopFinished() bool {
	return d._AnimationDidStopFinished != nil
}
