// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ProgressIndicator] class.
var (
	ProgressIndicatorClass     _ProgressIndicatorClass
	ProgressIndicatorClassOnce sync.Once
)

func getProgressIndicatorClass() _ProgressIndicatorClass {
	ProgressIndicatorClassOnce.Do(func() {
		ProgressIndicatorClass = _ProgressIndicatorClass{objc.GetClass("NSProgressIndicator")}
	})
	return ProgressIndicatorClass
}

type _ProgressIndicatorClass struct {
	class objc.Class
}

// An interface definition for the [ProgressIndicator] class.
type IProgressIndicator interface {
	IView
	Animate(sender objectivec.IObject)
	AnimationDelay() float64
	IncrementBy(delta float64)
	SetAnimationDelay(delay float64)
	SizeToFit()
	StartAnimation(sender objectivec.IObject)
	StopAnimation(sender objectivec.IObject)
}

// An interface that provides visual feedback to the user about the status of an ongoing task.
//
// Progress indicators can be determinate or indeterminate. A determinate indicator displays the completion percentage of a task. An indeterminate indicator shows that the app is busy without providing a visual indication of how long the task will take.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getProgressIndicatorClass().New()
}


// This action method advances the progress animation of an indeterminate progress animator by one step.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/animate:
func (p_ ProgressIndicator) Animate(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("animate:"), sender)
}

// Returns the delay, in seconds, between animation steps for an indeterminate progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/animationDelay
func (p_ ProgressIndicator) AnimationDelay() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("animationDelay"))
	return rv
}

// Advances the progress bar of a determinate progress indicator by the specified amount.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/increment(by:)
func (p_ ProgressIndicator) IncrementBy(delta float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("incrementBy:"), delta)
}

// Sets the delay, in seconds, between animation steps for an indeterminate progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/setAnimationDelay:
func (p_ ProgressIndicator) SetAnimationDelay(delay float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAnimationDelay:"), delay)
}

// This action method resizes the progress indicator to an appropriate size depending on the value of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/sizeToFit()
func (p_ ProgressIndicator) SizeToFit() {
	objc.Send[objc.ID](p_.ID, objc.Sel("sizeToFit"))
}

// Starts the animation of an indeterminate progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/startAnimation(_:)
func (p_ ProgressIndicator) StartAnimation(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startAnimation:"), sender)
}

// Stops the animation of an indeterminate progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/stopAnimation(_:)
func (p_ ProgressIndicator) StopAnimation(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopAnimation:"), sender)
}

// The size of the progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlSize
func (p_ ProgressIndicator) ControlSize() ControlSize {
	rv := objc.Send[ControlSize](p_.ID, objc.Sel("controlSize"))
	return rv
}


// SetControlSize sets the value of the controlSize property.
// The size of the progress indicator.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlSize
func (p_ ProgressIndicator) SetControlSize(value IControlSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlSize:"), value)
}

// The progress indicator’s control tint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlTint
func (p_ ProgressIndicator) ControlTint() ControlTint {
	rv := objc.Send[ControlTint](p_.ID, objc.Sel("controlTint"))
	return rv
}


// SetControlTint sets the value of the controlTint property.
// The progress indicator’s control tint.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlTint
func (p_ ProgressIndicator) SetControlTint(value IControlTint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlTint:"), value)
}

// The value that indicates the current extent of the progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/doubleValue
func (p_ ProgressIndicator) DoubleValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("doubleValue"))
	return rv
}


// SetDoubleValue sets the value of the doubleValue property.
// The value that indicates the current extent of the progress indicator.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/doubleValue
func (p_ ProgressIndicator) SetDoubleValue(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDoubleValue:"), value)
}

// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isBezeled
func (p_ ProgressIndicator) Bezeled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("bezeled"))
	return rv
}


// SetBezeled sets the value of the bezeled property.
// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isBezeled
func (p_ ProgressIndicator) SetBezeled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBezeled:"), value)
}

// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isDisplayedWhenStopped
func (p_ ProgressIndicator) DisplayedWhenStopped() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("displayedWhenStopped"))
	return rv
}


// SetDisplayedWhenStopped sets the value of the displayedWhenStopped property.
// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isDisplayedWhenStopped
func (p_ ProgressIndicator) SetDisplayedWhenStopped(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayedWhenStopped:"), value)
}

// A Boolean that indicates whether the progress indicator is indeterminate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isIndeterminate
func (p_ ProgressIndicator) Indeterminate() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("indeterminate"))
	return rv
}


// SetIndeterminate sets the value of the indeterminate property.
// A Boolean that indicates whether the progress indicator is indeterminate.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isIndeterminate
func (p_ ProgressIndicator) SetIndeterminate(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndeterminate:"), value)
}

// The maximum value for the progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/maxValue
func (p_ ProgressIndicator) MaxValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("maxValue"))
	return rv
}


// SetMaxValue sets the value of the maxValue property.
// The maximum value for the progress indicator.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/maxValue
func (p_ ProgressIndicator) SetMaxValue(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaxValue:"), value)
}

// The minimum value for the progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/minValue
func (p_ ProgressIndicator) MinValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minValue"))
	return rv
}


// SetMinValue sets the value of the minValue property.
// The minimum value for the progress indicator.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/minValue
func (p_ ProgressIndicator) SetMinValue(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMinValue:"), value)
}

// The progress object to use for updating the progress view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/observedProgress
func (p_ ProgressIndicator) ObservedProgress() foundation.Progress {
	rv := objc.Send[foundation.Progress](p_.ID, objc.Sel("observedProgress"))
	return rv
}


// SetObservedProgress sets the value of the observedProgress property.
// The progress object to use for updating the progress view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/observedProgress
func (p_ ProgressIndicator) SetObservedProgress(value foundation.IProgress) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObservedProgress:"), value)
}

// The style of the progress indicator (bar or spinning).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/style-swift.property
func (p_ ProgressIndicator) Style() ProgressIndicatorStyle {
	rv := objc.Send[ProgressIndicatorStyle](p_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// The style of the progress indicator (bar or spinning).

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/style-swift.property
func (p_ ProgressIndicator) SetStyle(value ProgressIndicatorStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}

// A Boolean that indicates whether the progress indicator implements animation in a separate thread.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/usesThreadedAnimation
func (p_ ProgressIndicator) UsesThreadedAnimation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesThreadedAnimation"))
	return rv
}


// SetUsesThreadedAnimation sets the value of the usesThreadedAnimation property.
// A Boolean that indicates whether the progress indicator implements animation in a separate thread.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/usesThreadedAnimation
func (p_ ProgressIndicator) SetUsesThreadedAnimation(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesThreadedAnimation:"), value)
}

// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isbezeled
func (p_ ProgressIndicator) IsBezeled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isBezeled"))
	return rv
}


// SetIsBezeled sets the value of the isBezeled property.
// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isbezeled
func (p_ ProgressIndicator) SetIsBezeled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsBezeled:"), value)
}

// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isdisplayedwhenstopped
func (p_ ProgressIndicator) IsDisplayedWhenStopped() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDisplayedWhenStopped"))
	return rv
}


// SetIsDisplayedWhenStopped sets the value of the isDisplayedWhenStopped property.
// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isdisplayedwhenstopped
func (p_ ProgressIndicator) SetIsDisplayedWhenStopped(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsDisplayedWhenStopped:"), value)
}

// A Boolean that indicates whether the progress indicator is indeterminate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isindeterminate
func (p_ ProgressIndicator) IsIndeterminate() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndeterminate"))
	return rv
}


// SetIsIndeterminate sets the value of the isIndeterminate property.
// A Boolean that indicates whether the progress indicator is indeterminate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isindeterminate
func (p_ ProgressIndicator) SetIsIndeterminate(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndeterminate:"), value)
}



