// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAnimationDelegate is the NSAnimationDelegate protocol interface.
//
// A set of optional methods implemented by delegates of   objects.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAnimationDelegate
type PAnimationDelegate interface {
	// Optional methods
	AnimationDidReachProgressMark(animation IAnimation, progress objc.IObject /* cross-framework: AnimationProgress */)
	HasAnimationDidReachProgressMark() bool
	AnimationValueForProgress(animation IAnimation, progress objc.IObject /* cross-framework: AnimationProgress */) float32
	HasAnimationValueForProgress() bool
	AnimationDidEnd(animation IAnimation)
	HasAnimationDidEnd() bool
	AnimationDidStop(animation IAnimation)
	HasAnimationDidStop() bool
	AnimationShouldStart(animation IAnimation) bool
	HasAnimationShouldStart() bool
}

// AnimationDelegate is a delegate implementation builder for the PAnimationDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AnimationDelegate struct {
	_AnimationDidReachProgressMark func(animation IAnimation, progress objc.IObject /* cross-framework: AnimationProgress */)
	_AnimationValueForProgress func(animation IAnimation, progress objc.IObject /* cross-framework: AnimationProgress */) float32
	_AnimationDidEnd func(animation IAnimation)
	_AnimationDidStop func(animation IAnimation)
	_AnimationShouldStart func(animation IAnimation) bool
}

// SetAnimationDidReachProgressMark sets the handler for the AnimationDidReachProgressMark delegate method.
//
// Sent to the delegate when an animation reaches a specific progress mark.
func (d *AnimationDelegate) SetAnimationDidReachProgressMark(f func(animation IAnimation, progress objc.IObject /* cross-framework: AnimationProgress */)) {
	d._AnimationDidReachProgressMark = f
}

// SetAnimationValueForProgress sets the handler for the AnimationValueForProgress delegate method.
//
// Requests a custom curve value for the current progress value.
func (d *AnimationDelegate) SetAnimationValueForProgress(f func(animation IAnimation, progress objc.IObject /* cross-framework: AnimationProgress */) float32) {
	d._AnimationValueForProgress = f
}

// SetAnimationDidEnd sets the handler for the AnimationDidEnd delegate method.
//
// Sent to the delegate when the specified animation completes its run.
func (d *AnimationDelegate) SetAnimationDidEnd(f func(animation IAnimation)) {
	d._AnimationDidEnd = f
}

// SetAnimationDidStop sets the handler for the AnimationDidStop delegate method.
//
// Sent to the delegate when the specified animation is stopped before it completes its run.
func (d *AnimationDelegate) SetAnimationDidStop(f func(animation IAnimation)) {
	d._AnimationDidStop = f
}

// SetAnimationShouldStart sets the handler for the AnimationShouldStart delegate method.
//
// Sent to the delegate just after an animation is started.
func (d *AnimationDelegate) SetAnimationShouldStart(f func(animation IAnimation) bool) {
	d._AnimationShouldStart = f
}

// AnimationDidReachProgressMark implements the PAnimationDelegate interface.
func (d *AnimationDelegate) AnimationDidReachProgressMark(animation IAnimation, progress objc.IObject /* cross-framework: AnimationProgress */) {
	if d._AnimationDidReachProgressMark != nil {
		d._AnimationDidReachProgressMark(animation, progress)
	}
}

// HasAnimationDidReachProgressMark returns true if a handler for AnimationDidReachProgressMark has been set.
func (d *AnimationDelegate) HasAnimationDidReachProgressMark() bool {
	return d._AnimationDidReachProgressMark != nil
}

// AnimationValueForProgress implements the PAnimationDelegate interface.
func (d *AnimationDelegate) AnimationValueForProgress(animation IAnimation, progress objc.IObject /* cross-framework: AnimationProgress */) float32 {
	if d._AnimationValueForProgress != nil {
		return d._AnimationValueForProgress(animation, progress)
	}
	var zero float32
	return zero
}

// HasAnimationValueForProgress returns true if a handler for AnimationValueForProgress has been set.
func (d *AnimationDelegate) HasAnimationValueForProgress() bool {
	return d._AnimationValueForProgress != nil
}

// AnimationDidEnd implements the PAnimationDelegate interface.
func (d *AnimationDelegate) AnimationDidEnd(animation IAnimation) {
	if d._AnimationDidEnd != nil {
		d._AnimationDidEnd(animation)
	}
}

// HasAnimationDidEnd returns true if a handler for AnimationDidEnd has been set.
func (d *AnimationDelegate) HasAnimationDidEnd() bool {
	return d._AnimationDidEnd != nil
}

// AnimationDidStop implements the PAnimationDelegate interface.
func (d *AnimationDelegate) AnimationDidStop(animation IAnimation) {
	if d._AnimationDidStop != nil {
		d._AnimationDidStop(animation)
	}
}

// HasAnimationDidStop returns true if a handler for AnimationDidStop has been set.
func (d *AnimationDelegate) HasAnimationDidStop() bool {
	return d._AnimationDidStop != nil
}

// AnimationShouldStart implements the PAnimationDelegate interface.
func (d *AnimationDelegate) AnimationShouldStart(animation IAnimation) bool {
	if d._AnimationShouldStart != nil {
		return d._AnimationShouldStart(animation)
	}
	var zero bool
	return zero
}

// HasAnimationShouldStart returns true if a handler for AnimationShouldStart has been set.
func (d *AnimationDelegate) HasAnimationShouldStart() bool {
	return d._AnimationShouldStart != nil
}
