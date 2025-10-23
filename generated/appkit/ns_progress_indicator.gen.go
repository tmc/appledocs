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
	ControlSize() unsafe.Pointer
	SetControlSize(value unsafe.Pointer)
	ControlTint() unsafe.Pointer
	SetControlTint(value unsafe.Pointer)
	DoubleValue() float64
	SetDoubleValue(value float64)
	Bezeled() bool
	SetBezeled(value bool)
	DisplayedWhenStopped() bool
	SetDisplayedWhenStopped(value bool)
	Indeterminate() bool
	SetIndeterminate(value bool)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	ObservedProgress() foundation.Progress
	SetObservedProgress(value foundation.Progress)
	Style() NSProgressIndicatorStyle
	SetStyle(value NSProgressIndicatorStyle)
	UsesThreadedAnimation() bool
	SetUsesThreadedAnimation(value bool)
	IsBezeled() bool
	SetIsBezeled(value bool)
	IsDisplayedWhenStopped() bool
	SetIsDisplayedWhenStopped(value bool)
	IsIndeterminate() bool
	SetIsIndeterminate(value bool)
	IncrementBy(delta float64)
	SizeToFit()
	StartAnimation(sender objectivec.IObject)
	StopAnimation(sender objectivec.IObject)
}

// An interface that provides visual feedback to the user about the status of an ongoing task.
//
// Progress indicators can be determinate or indeterminate. A determinate indicator displays the completion percentage of a task. An indeterminate indicator shows that the app is busy without providing a visual indication of how long the task will take.


// An interface that provides visual feedback to the user about the status of an ongoing task.
//
// [Full Topic]
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



// Advances the progress bar of a determinate progress indicator by the specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/increment(by:)
func (p_ ProgressIndicator) IncrementBy(delta float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("incrementBy:"), delta)
}


// This action method resizes the progress indicator to an appropriate size depending on the value of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/sizeToFit()
func (p_ ProgressIndicator) SizeToFit() {
	objc.Send[objc.ID](p_.ID, objc.Sel("sizeToFit"))
}


// Starts the animation of an indeterminate progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/startAnimation(_:)
func (p_ ProgressIndicator) StartAnimation(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startAnimation:"), sender)
}


// Stops the animation of an indeterminate progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/stopAnimation(_:)
func (p_ ProgressIndicator) StopAnimation(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopAnimation:"), sender)
}


// The size of the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlSize
func (p_ ProgressIndicator) ControlSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("controlSize"))
	return rv
}


// The size of the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlSize
func (p_ ProgressIndicator) SetControlSize(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlSize:"), value)
}


// The progress indicator’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlTint
func (p_ ProgressIndicator) ControlTint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("controlTint"))
	return rv
}


// The progress indicator’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlTint
func (p_ ProgressIndicator) SetControlTint(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlTint:"), value)
}


// The value that indicates the current extent of the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/doubleValue
func (p_ ProgressIndicator) DoubleValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("doubleValue"))
	return rv
}


// The value that indicates the current extent of the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/doubleValue
func (p_ ProgressIndicator) SetDoubleValue(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDoubleValue:"), value)
}


// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isBezeled
func (p_ ProgressIndicator) Bezeled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("bezeled"))
	return rv
}


// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isBezeled
func (p_ ProgressIndicator) SetBezeled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBezeled:"), value)
}


// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isDisplayedWhenStopped
func (p_ ProgressIndicator) DisplayedWhenStopped() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("displayedWhenStopped"))
	return rv
}


// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isDisplayedWhenStopped
func (p_ ProgressIndicator) SetDisplayedWhenStopped(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayedWhenStopped:"), value)
}


// A Boolean that indicates whether the progress indicator is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isIndeterminate
func (p_ ProgressIndicator) Indeterminate() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("indeterminate"))
	return rv
}


// A Boolean that indicates whether the progress indicator is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isIndeterminate
func (p_ ProgressIndicator) SetIndeterminate(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndeterminate:"), value)
}


// The maximum value for the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/maxValue
func (p_ ProgressIndicator) MaxValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("maxValue"))
	return rv
}


// The maximum value for the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/maxValue
func (p_ ProgressIndicator) SetMaxValue(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaxValue:"), value)
}


// The minimum value for the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/minValue
func (p_ ProgressIndicator) MinValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minValue"))
	return rv
}


// The minimum value for the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/minValue
func (p_ ProgressIndicator) SetMinValue(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMinValue:"), value)
}


// The progress object to use for updating the progress view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/observedProgress
func (p_ ProgressIndicator) ObservedProgress() foundation.Progress {
	rv := objc.Send[foundation.Progress](p_.ID, objc.Sel("observedProgress"))
	return rv
}


// The progress object to use for updating the progress view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/observedProgress
func (p_ ProgressIndicator) SetObservedProgress(value foundation.Progress) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObservedProgress:"), value)
}


// The style of the progress indicator (bar or spinning).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/style-swift.property
func (p_ ProgressIndicator) Style() NSProgressIndicatorStyle {
	rv := objc.Send[NSProgressIndicatorStyle](p_.ID, objc.Sel("style"))
	return rv
}


// The style of the progress indicator (bar or spinning).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/style-swift.property
func (p_ ProgressIndicator) SetStyle(value NSProgressIndicatorStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}


// A Boolean that indicates whether the progress indicator implements animation in a separate thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/usesThreadedAnimation
func (p_ ProgressIndicator) UsesThreadedAnimation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesThreadedAnimation"))
	return rv
}


// A Boolean that indicates whether the progress indicator implements animation in a separate thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/usesThreadedAnimation
func (p_ ProgressIndicator) SetUsesThreadedAnimation(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesThreadedAnimation:"), value)
}


// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isbezeled
func (p_ ProgressIndicator) IsBezeled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isBezeled"))
	return rv
}


// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isbezeled
func (p_ ProgressIndicator) SetIsBezeled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsBezeled:"), value)
}


// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isdisplayedwhenstopped
func (p_ ProgressIndicator) IsDisplayedWhenStopped() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDisplayedWhenStopped"))
	return rv
}


// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isdisplayedwhenstopped
func (p_ ProgressIndicator) SetIsDisplayedWhenStopped(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsDisplayedWhenStopped:"), value)
}


// A Boolean that indicates whether the progress indicator is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isindeterminate
func (p_ ProgressIndicator) IsIndeterminate() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndeterminate"))
	return rv
}


// A Boolean that indicates whether the progress indicator is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isindeterminate
func (p_ ProgressIndicator) SetIsIndeterminate(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndeterminate:"), value)
}



