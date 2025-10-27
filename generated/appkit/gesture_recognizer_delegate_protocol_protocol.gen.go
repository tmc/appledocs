// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PGestureRecognizerDelegate is the NSGestureRecognizerDelegate protocol interface.
//
// A set of methods for fine-tuning a gesture recognizer’s behavior.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizerDelegate
type PGestureRecognizerDelegate interface {
	// Optional methods
	GestureRecognizerShouldAttemptToRecognizeWithEvent(gestureRecognizer GestureRecognizer, event IEvent) bool
	HasGestureRecognizerShouldAttemptToRecognizeWithEvent() bool
	GestureRecognizerShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	HasGestureRecognizerShouldBeRequiredToFailByGestureRecognizer() bool
	GestureRecognizerShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch ITouch) bool
	HasGestureRecognizerShouldReceiveTouch() bool
	GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	HasGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer() bool
	GestureRecognizerShouldRequireFailureOfGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	HasGestureRecognizerShouldRequireFailureOfGestureRecognizer() bool
	GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool
	HasGestureRecognizerShouldBegin() bool
}

// GestureRecognizerDelegate is a delegate implementation builder for the PGestureRecognizerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type GestureRecognizerDelegate struct {
	_GestureRecognizerShouldAttemptToRecognizeWithEvent func(gestureRecognizer GestureRecognizer, event IEvent) bool
	_GestureRecognizerShouldBeRequiredToFailByGestureRecognizer func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizerShouldReceiveTouch func(gestureRecognizer GestureRecognizer, touch ITouch) bool
	_GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizerShouldRequireFailureOfGestureRecognizer func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizerShouldBegin func(gestureRecognizer GestureRecognizer) bool
}

// SetGestureRecognizerShouldAttemptToRecognizeWithEvent sets the handler for the GestureRecognizerShouldAttemptToRecognizeWithEvent delegate method.
//
// Asks the delegate if a gesture recognizer should attempt to recognize gestures for a particular event.
func (d *GestureRecognizerDelegate) SetGestureRecognizerShouldAttemptToRecognizeWithEvent(f func(gestureRecognizer GestureRecognizer, event IEvent) bool) {
	d._GestureRecognizerShouldAttemptToRecognizeWithEvent = f
}

// SetGestureRecognizerShouldBeRequiredToFailByGestureRecognizer sets the handler for the GestureRecognizerShouldBeRequiredToFailByGestureRecognizer delegate method.
//
// Asks the delegate if the current gesture recognizer must fail before another gesture recognizer is allowed to recognize its gesture.
func (d *GestureRecognizerDelegate) SetGestureRecognizerShouldBeRequiredToFailByGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	d._GestureRecognizerShouldBeRequiredToFailByGestureRecognizer = f
}

// SetGestureRecognizerShouldReceiveTouch sets the handler for the GestureRecognizerShouldReceiveTouch delegate method.
//
// Called, for a new touch, before the system calls the   method on the gesture recognizer. Return   to prevent the gesture recognizer from seeing this touch.
func (d *GestureRecognizerDelegate) SetGestureRecognizerShouldReceiveTouch(f func(gestureRecognizer GestureRecognizer, touch ITouch) bool) {
	d._GestureRecognizerShouldReceiveTouch = f
}

// SetGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer sets the handler for the GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer delegate method.
//
// Asks the delegate if two gesture recognizers should be allowed to recognize their gestures simultaneously.
func (d *GestureRecognizerDelegate) SetGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	d._GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer = f
}

// SetGestureRecognizerShouldRequireFailureOfGestureRecognizer sets the handler for the GestureRecognizerShouldRequireFailureOfGestureRecognizer delegate method.
//
// Asks the delegate if the current gesture recognizer must wait to recognize its gesture until the specified gesture recognizer fails.
func (d *GestureRecognizerDelegate) SetGestureRecognizerShouldRequireFailureOfGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	d._GestureRecognizerShouldRequireFailureOfGestureRecognizer = f
}

// SetGestureRecognizerShouldBegin sets the handler for the GestureRecognizerShouldBegin delegate method.
//
// Asks the delegate if a gesture recognizer should transition out of the Possible ( ) state.
func (d *GestureRecognizerDelegate) SetGestureRecognizerShouldBegin(f func(gestureRecognizer GestureRecognizer) bool) {
	d._GestureRecognizerShouldBegin = f
}

// GestureRecognizerShouldAttemptToRecognizeWithEvent implements the PGestureRecognizerDelegate interface.
func (d *GestureRecognizerDelegate) GestureRecognizerShouldAttemptToRecognizeWithEvent(gestureRecognizer GestureRecognizer, event IEvent) bool {
	if d._GestureRecognizerShouldAttemptToRecognizeWithEvent != nil {
		return d._GestureRecognizerShouldAttemptToRecognizeWithEvent(gestureRecognizer, event)
	}
	var zero bool
	return zero
}

// HasGestureRecognizerShouldAttemptToRecognizeWithEvent returns true if a handler for GestureRecognizerShouldAttemptToRecognizeWithEvent has been set.
func (d *GestureRecognizerDelegate) HasGestureRecognizerShouldAttemptToRecognizeWithEvent() bool {
	return d._GestureRecognizerShouldAttemptToRecognizeWithEvent != nil
}

// GestureRecognizerShouldBeRequiredToFailByGestureRecognizer implements the PGestureRecognizerDelegate interface.
func (d *GestureRecognizerDelegate) GestureRecognizerShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	if d._GestureRecognizerShouldBeRequiredToFailByGestureRecognizer != nil {
		return d._GestureRecognizerShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
	}
	var zero bool
	return zero
}

// HasGestureRecognizerShouldBeRequiredToFailByGestureRecognizer returns true if a handler for GestureRecognizerShouldBeRequiredToFailByGestureRecognizer has been set.
func (d *GestureRecognizerDelegate) HasGestureRecognizerShouldBeRequiredToFailByGestureRecognizer() bool {
	return d._GestureRecognizerShouldBeRequiredToFailByGestureRecognizer != nil
}

// GestureRecognizerShouldReceiveTouch implements the PGestureRecognizerDelegate interface.
func (d *GestureRecognizerDelegate) GestureRecognizerShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch ITouch) bool {
	if d._GestureRecognizerShouldReceiveTouch != nil {
		return d._GestureRecognizerShouldReceiveTouch(gestureRecognizer, touch)
	}
	var zero bool
	return zero
}

// HasGestureRecognizerShouldReceiveTouch returns true if a handler for GestureRecognizerShouldReceiveTouch has been set.
func (d *GestureRecognizerDelegate) HasGestureRecognizerShouldReceiveTouch() bool {
	return d._GestureRecognizerShouldReceiveTouch != nil
}

// GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer implements the PGestureRecognizerDelegate interface.
func (d *GestureRecognizerDelegate) GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	if d._GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer != nil {
		return d._GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
	}
	var zero bool
	return zero
}

// HasGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer returns true if a handler for GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer has been set.
func (d *GestureRecognizerDelegate) HasGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer() bool {
	return d._GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer != nil
}

// GestureRecognizerShouldRequireFailureOfGestureRecognizer implements the PGestureRecognizerDelegate interface.
func (d *GestureRecognizerDelegate) GestureRecognizerShouldRequireFailureOfGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	if d._GestureRecognizerShouldRequireFailureOfGestureRecognizer != nil {
		return d._GestureRecognizerShouldRequireFailureOfGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
	}
	var zero bool
	return zero
}

// HasGestureRecognizerShouldRequireFailureOfGestureRecognizer returns true if a handler for GestureRecognizerShouldRequireFailureOfGestureRecognizer has been set.
func (d *GestureRecognizerDelegate) HasGestureRecognizerShouldRequireFailureOfGestureRecognizer() bool {
	return d._GestureRecognizerShouldRequireFailureOfGestureRecognizer != nil
}

// GestureRecognizerShouldBegin implements the PGestureRecognizerDelegate interface.
func (d *GestureRecognizerDelegate) GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool {
	if d._GestureRecognizerShouldBegin != nil {
		return d._GestureRecognizerShouldBegin(gestureRecognizer)
	}
	var zero bool
	return zero
}

// HasGestureRecognizerShouldBegin returns true if a handler for GestureRecognizerShouldBegin has been set.
func (d *GestureRecognizerDelegate) HasGestureRecognizerShouldBegin() bool {
	return d._GestureRecognizerShouldBegin != nil
}

// GestureRecognizerDelegateObject wraps an existing Objective-C object that conforms to the PGestureRecognizerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type GestureRecognizerDelegateObject struct {
	objectivec.Object
}

// NewGestureRecognizerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSGestureRecognizerDelegate protocol.
func NewGestureRecognizerDelegateObject(obj objectivec.Object) *GestureRecognizerDelegateObject {
	return &GestureRecognizerDelegateObject{obj}
}

// Make sure GestureRecognizerDelegateObject implements PGestureRecognizerDelegate.
var _ PGestureRecognizerDelegate = (*GestureRecognizerDelegateObject)(nil)

// GestureRecognizerShouldAttemptToRecognizeWithEvent implements the PGestureRecognizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *GestureRecognizerDelegateObject) GestureRecognizerShouldAttemptToRecognizeWithEvent(gestureRecognizer GestureRecognizer, event IEvent) bool {
	return objc.Send[bool](o.ID, objc.Sel("gestureRecognizer:shouldAttemptToRecognizeWithEvent:"), gestureRecognizer, event)
}

// HasGestureRecognizerShouldAttemptToRecognizeWithEvent returns true; this is a placeholder for optional method checks.
func (o *GestureRecognizerDelegateObject) HasGestureRecognizerShouldAttemptToRecognizeWithEvent() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// GestureRecognizerShouldBeRequiredToFailByGestureRecognizer implements the PGestureRecognizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *GestureRecognizerDelegateObject) GestureRecognizerShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return objc.Send[bool](o.ID, objc.Sel("gestureRecognizer:shouldBeRequiredToFailByGestureRecognizer:"), gestureRecognizer, otherGestureRecognizer)
}

// HasGestureRecognizerShouldBeRequiredToFailByGestureRecognizer returns true; this is a placeholder for optional method checks.
func (o *GestureRecognizerDelegateObject) HasGestureRecognizerShouldBeRequiredToFailByGestureRecognizer() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// GestureRecognizerShouldReceiveTouch implements the PGestureRecognizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *GestureRecognizerDelegateObject) GestureRecognizerShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch ITouch) bool {
	return objc.Send[bool](o.ID, objc.Sel("gestureRecognizer:shouldReceiveTouch:"), gestureRecognizer, touch)
}

// HasGestureRecognizerShouldReceiveTouch returns true; this is a placeholder for optional method checks.
func (o *GestureRecognizerDelegateObject) HasGestureRecognizerShouldReceiveTouch() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer implements the PGestureRecognizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *GestureRecognizerDelegateObject) GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return objc.Send[bool](o.ID, objc.Sel("gestureRecognizer:shouldRecognizeSimultaneouslyWithGestureRecognizer:"), gestureRecognizer, otherGestureRecognizer)
}

// HasGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer returns true; this is a placeholder for optional method checks.
func (o *GestureRecognizerDelegateObject) HasGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// GestureRecognizerShouldRequireFailureOfGestureRecognizer implements the PGestureRecognizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *GestureRecognizerDelegateObject) GestureRecognizerShouldRequireFailureOfGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return objc.Send[bool](o.ID, objc.Sel("gestureRecognizer:shouldRequireFailureOfGestureRecognizer:"), gestureRecognizer, otherGestureRecognizer)
}

// HasGestureRecognizerShouldRequireFailureOfGestureRecognizer returns true; this is a placeholder for optional method checks.
func (o *GestureRecognizerDelegateObject) HasGestureRecognizerShouldRequireFailureOfGestureRecognizer() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// GestureRecognizerShouldBegin implements the PGestureRecognizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *GestureRecognizerDelegateObject) GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool {
	return objc.Send[bool](o.ID, objc.Sel("gestureRecognizerShouldBegin:"), gestureRecognizer)
}

// HasGestureRecognizerShouldBegin returns true; this is a placeholder for optional method checks.
func (o *GestureRecognizerDelegateObject) HasGestureRecognizerShouldBegin() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
