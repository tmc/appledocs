
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ProgressIndicator] class.
var ProgressIndicatorClass _ProgressIndicatorClass

func init() {
	ProgressIndicatorClass = _ProgressIndicatorClass{objc.GetClass("NSProgressIndicator")}
}

type _ProgressIndicatorClass struct {
	objc.Class
}

// An interface definition for the [ProgressIndicator] class.
type IProgressIndicator interface {
	ID() objc.ID
	Animate(sender objc.ID)
	AnimationDelay() unsafe.Pointer
	IncrementBy(delta float64)
	SetAnimationDelay(delay unsafe.Pointer)
	SizeToFit()
	StartAnimation(sender objc.ID)
	StopAnimation(sender objc.ID)
}

type ProgressIndicator struct {
	id objc.ID
}

func ProgressIndicatorFrom(ptr unsafe.Pointer) ProgressIndicator {
	return ProgressIndicator{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ ProgressIndicator) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _ProgressIndicatorClass) Alloc() ProgressIndicator {
	rv := objc.Send[ProgressIndicator](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _ProgressIndicatorClass) New() ProgressIndicator {
	rv := objc.Send[ProgressIndicator](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewProgressIndicator creates and returns a new initialized instance.
func NewProgressIndicator() ProgressIndicator {
	return ProgressIndicatorClass.New()
}

// Init initializes the instance.
func (p_ ProgressIndicator) Init() ProgressIndicator {
	rv := objc.Send[ProgressIndicator](p_.ID(), selInit)
	return rv
}
// This action method advances the progress animation of an indeterminate progress animator by one step. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/animate:
func (p_ ProgressIndicator) Animate(sender objc.ID) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("animate:"), sender)
}
// Returns the delay, in seconds, between animation steps for an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/animationDelay
func (p_ ProgressIndicator) AnimationDelay() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID(), objc.RegisterName("animationDelay"))
	return rv
}
// Advances the progress bar of a determinate progress indicator by the specified amount. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/increment(by:)
func (p_ ProgressIndicator) IncrementBy(delta float64) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("incrementBy:"), delta)
}
// Sets the delay, in seconds, between animation steps for an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/setAnimationDelay:
func (p_ ProgressIndicator) SetAnimationDelay(delay unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setAnimationDelay:"), delay)
}
// This action method resizes the progress indicator to an appropriate size depending on the value of  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/sizeToFit()
func (p_ ProgressIndicator) SizeToFit() {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("sizeToFit"))
}
// Starts the animation of an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/startAnimation(_:)
func (p_ ProgressIndicator) StartAnimation(sender objc.ID) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("startAnimation:"), sender)
}
// Stops the animation of an indeterminate progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/stopAnimation(_:)
func (p_ ProgressIndicator) StopAnimation(sender objc.ID) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("stopAnimation:"), sender)
}
// The size of the progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/controlSize
func (p_ ProgressIndicator) ControlSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID(), objc.RegisterName("controlSize"))
	return rv
}
// SetControlSize sets the value of the controlSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/controlSize
func (p_ ProgressIndicator) SetControlSize(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setControlSize:"), value)
}
// The progress indicator’s control tint. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/controlTint
func (p_ ProgressIndicator) ControlTint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID(), objc.RegisterName("controlTint"))
	return rv
}
// SetControlTint sets the value of the controlTint property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/controlTint
func (p_ ProgressIndicator) SetControlTint(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setControlTint:"), value)
}
// The value that indicates the current extent of the progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/doubleValue
func (p_ ProgressIndicator) DoubleValue() float64 {
	rv := objc.Send[float64](p_.ID(), objc.RegisterName("doubleValue"))
	return rv
}
// SetDoubleValue sets the value of the doubleValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/doubleValue
func (p_ ProgressIndicator) SetDoubleValue(value float64) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setDoubleValue:"), value)
}
// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/isBezeled
func (p_ ProgressIndicator) Bezeled() bool {
	rv := objc.Send[bool](p_.ID(), objc.RegisterName("bezeled"))
	return rv
}
// SetBezeled sets the value of the bezeled property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/isBezeled
func (p_ ProgressIndicator) SetBezeled(value bool) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setBezeled:"), value)
}
// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/isDisplayedWhenStopped
func (p_ ProgressIndicator) DisplayedWhenStopped() bool {
	rv := objc.Send[bool](p_.ID(), objc.RegisterName("displayedWhenStopped"))
	return rv
}
// SetDisplayedWhenStopped sets the value of the displayedWhenStopped property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/isDisplayedWhenStopped
func (p_ ProgressIndicator) SetDisplayedWhenStopped(value bool) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setDisplayedWhenStopped:"), value)
}
// A Boolean that indicates whether the progress indicator is indeterminate. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/isIndeterminate
func (p_ ProgressIndicator) Indeterminate() bool {
	rv := objc.Send[bool](p_.ID(), objc.RegisterName("indeterminate"))
	return rv
}
// SetIndeterminate sets the value of the indeterminate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/isIndeterminate
func (p_ ProgressIndicator) SetIndeterminate(value bool) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setIndeterminate:"), value)
}
// The maximum value for the progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/maxValue
func (p_ ProgressIndicator) MaxValue() float64 {
	rv := objc.Send[float64](p_.ID(), objc.RegisterName("maxValue"))
	return rv
}
// SetMaxValue sets the value of the maxValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/maxValue
func (p_ ProgressIndicator) SetMaxValue(value float64) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setMaxValue:"), value)
}
// The minimum value for the progress indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/minValue
func (p_ ProgressIndicator) MinValue() float64 {
	rv := objc.Send[float64](p_.ID(), objc.RegisterName("minValue"))
	return rv
}
// SetMinValue sets the value of the minValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/minValue
func (p_ ProgressIndicator) SetMinValue(value float64) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setMinValue:"), value)
}
// The progress object to use for updating the progress view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/observedProgress
func (p_ ProgressIndicator) ObservedProgress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID(), objc.RegisterName("observedProgress"))
	return rv
}
// SetObservedProgress sets the value of the observedProgress property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/observedProgress
func (p_ ProgressIndicator) SetObservedProgress(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setObservedProgress:"), value)
}
// The style of the progress indicator (bar or spinning). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/style-swift.property
func (p_ ProgressIndicator) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID(), objc.RegisterName("style"))
	return rv
}
// SetStyle sets the value of the style property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/style-swift.property
func (p_ ProgressIndicator) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setStyle:"), value)
}
// A Boolean that indicates whether the progress indicator implements animation in a separate thread. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/usesThreadedAnimation
func (p_ ProgressIndicator) UsesThreadedAnimation() bool {
	rv := objc.Send[bool](p_.ID(), objc.RegisterName("usesThreadedAnimation"))
	return rv
}
// SetUsesThreadedAnimation sets the value of the usesThreadedAnimation property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/usesThreadedAnimation
func (p_ ProgressIndicator) SetUsesThreadedAnimation(value bool) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setUsesThreadedAnimation:"), value)
}
