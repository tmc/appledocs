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

// An interface definition for the [ProgressIndicator] class.
type IProgressIndicator interface {
	IView
	Animate(sender objc.ID)
	AnimationDelay() float64
	IncrementBy(delta float64)
	SetAnimationDelay(delay float64)
	SizeToFit()
	StartAnimation(sender objc.ID)
	StopAnimation(sender objc.ID)
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
// Alloc allocates a new instance without initialization.
func (pc _ProgressIndicatorClass) Alloc() ProgressIndicator {
	rv := objc.Send[ProgressIndicator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _ProgressIndicatorClass) New() ProgressIndicator {
	rv := objc.Send[ProgressIndicator](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ProgressIndicator) Init() ProgressIndicator {
	rv := objc.Send[ProgressIndicator](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ProgressIndicator) Autorelease() ProgressIndicator {
	rv := objc.Send[ProgressIndicator](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProgressIndicator creates a new ProgressIndicator instance.
func NewProgressIndicator() ProgressIndicator {
	return progressIndicatorClass.New()
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


