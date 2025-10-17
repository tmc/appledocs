// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ProgressIndicator] class.
var progressIndicatorClass = _ProgressIndicatorClass{objc.GetClass("NSProgressIndicator")}

type _ProgressIndicatorClass struct {
	class objc.Class
}

// An interface that provides visual feedback to the user about the status of an ongoing task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator

type ProgressIndicator struct {
	View
}

// ProgressIndicatorFrom constructs a [ProgressIndicator] from an unsafe.Pointer.
//
// An interface that provides visual feedback to the user about the status of an ongoing task.
func ProgressIndicatorFrom(ptr unsafe.Pointer) ProgressIndicator {
	return ProgressIndicator{
		View: ViewFrom(ptr),
	}
}

// This action method advances the progress animation of an indeterminate progress animator by one step. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/animate:
func (p_ ProgressIndicator) Animate(sender objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("animate:"), sender)
}
// Returns the delay, in seconds, between animation steps for an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/animationDelay
func (p_ ProgressIndicator) AnimationDelay() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("animationDelay"))
	return rv
}
// Advances the progress bar of a determinate progress indicator by the specified amount. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/increment(by:)
func (p_ ProgressIndicator) IncrementBy(delta float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("incrementBy:"), delta)
}
// Sets the delay, in seconds, between animation steps for an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/setAnimationDelay:
func (p_ ProgressIndicator) SetAnimationDelay(delay float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAnimationDelay:"), delay)
}
// This action method resizes the progress indicator to an appropriate size depending on the value of . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/sizeToFit()
func (p_ ProgressIndicator) SizeToFit() {
	objc.Send[objc.ID](p_.ID, objc.Sel("sizeToFit"))
}
// Starts the animation of an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/startAnimation(_:)
func (p_ ProgressIndicator) StartAnimation(sender objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startAnimation:"), sender)
}
// Stops the animation of an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/stopAnimation(_:)
func (p_ ProgressIndicator) StopAnimation(sender objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopAnimation:"), sender)
}


