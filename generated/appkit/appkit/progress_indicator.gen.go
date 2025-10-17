// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ProgressIndicator] class.
var ProgressIndicatorClass objc.Class

func init() {
	ProgressIndicatorClass = objc.GetClass("NSProgressIndicator")
}

type ProgressIndicator struct {
	objc.ID
}

func ProgressIndicatorFrom(ptr unsafe.Pointer) ProgressIndicator {
	return ProgressIndicator{
		ID: objc.ID(ptr),
	}
}


// This action method advances the progress animation of an indeterminate progress animator by one step. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/animate:
func (p_ ProgressIndicator) Animate(sender objc.ID) {
	sel := objc.RegisterName("animate:")
	p_.ID.Send(sel, sender)
}
// Returns the delay, in seconds, between animation steps for an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/animationDelay
func (p_ ProgressIndicator) AnimationDelay() foundation.TimeInterval {
	sel := objc.RegisterName("animationDelay")
	ret := p_.ID.Send(sel)
	return foundation.TimeInterval(ret)
}
// Advances the progress bar of a determinate progress indicator by the specified amount. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/increment(by:)
func (p_ ProgressIndicator) IncrementBy(delta float64) {
	sel := objc.RegisterName("incrementBy:")
	p_.ID.Send(sel, delta)
}
// Sets the delay, in seconds, between animation steps for an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/setAnimationDelay:
func (p_ ProgressIndicator) SetAnimationDelay(delay foundation.TimeInterval) {
	sel := objc.RegisterName("setAnimationDelay:")
	p_.ID.Send(sel, delay)
}
// This action method resizes the progress indicator to an appropriate size depending on the value of  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/sizeToFit()
func (p_ ProgressIndicator) SizeToFit() {
	sel := objc.RegisterName("sizeToFit")
	p_.ID.Send(sel)
}
// Starts the animation of an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/startAnimation(_:)
func (p_ ProgressIndicator) StartAnimation(sender objc.ID) {
	sel := objc.RegisterName("startAnimation:")
	p_.ID.Send(sel, sender)
}
// Stops the animation of an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/stopAnimation(_:)
func (p_ ProgressIndicator) StopAnimation(sender objc.ID) {
	sel := objc.RegisterName("stopAnimation:")
	p_.ID.Send(sel, sender)
}

