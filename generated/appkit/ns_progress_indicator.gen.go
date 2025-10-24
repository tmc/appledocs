// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSProgressIndicator */


/* debug [class_header]: Header for NSProgressIndicator */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ProgressIndicator */
// An interface definition for the [ProgressIndicator] class.
type IProgressIndicator interface {
	IView
	
/* debug [class_interface_properties]: Properties for ProgressIndicator */
	// properties:
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	ControlTint() ControlTint
	SetControlTint(value ControlTint)
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
	Style() ProgressIndicatorStyle
	SetStyle(value ProgressIndicatorStyle)
	UsesThreadedAnimation() bool
	SetUsesThreadedAnimation(value bool)
	IsBezeled() bool
	SetIsBezeled(value bool)
	IsDisplayedWhenStopped() bool
	SetIsDisplayedWhenStopped(value bool)
	IsIndeterminate() bool
	SetIsIndeterminate(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ProgressIndicator */
	// methods:
	IncrementBy(delta float64)
	SizeToFit()
	StartAnimation(sender objc.IObject)
	StopAnimation(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ProgressIndicator */
// Alloc allocates a new instance without initialization.
func (pc _ProgressIndicatorClass) Alloc() ProgressIndicator {
	rv := objc.Send[ProgressIndicator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ProgressIndicator */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ProgressIndicator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ProgressIndicator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ProgressIndicator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ProgressIndicator */

// Advances the progress bar of a determinate progress indicator by the specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/increment(by:)
func (p_ ProgressIndicator) IncrementBy(delta float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("incrementBy:"), delta)
}/* debug [instance_methods/method]: IncrementBy */


// This action method resizes the progress indicator to an appropriate size depending on the value of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/sizeToFit()
func (p_ ProgressIndicator) SizeToFit() {
	objc.Send[objc.ID](p_.ID, objc.Sel("sizeToFit"))
}/* debug [instance_methods/method]: SizeToFit */


// Starts the animation of an indeterminate progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/startAnimation(_:)
func (p_ ProgressIndicator) StartAnimation(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startAnimation:"), sender)
}/* debug [instance_methods/method]: StartAnimation */


// Stops the animation of an indeterminate progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/stopAnimation(_:)
func (p_ ProgressIndicator) StopAnimation(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopAnimation:"), sender)
}/* debug [instance_methods/method]: StopAnimation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ProgressIndicator */

// The size of the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlSize
func (p_ ProgressIndicator) ControlSize() ControlSize {
	rv := objc.Send[ControlSize](p_.ID, objc.Sel("controlSize"))
	return rv
}/* debug [instance_properties/getter]: controlSize */


// The size of the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlSize
func (p_ ProgressIndicator) SetControlSize(value ControlSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlSize:"), value)
}/* debug [instance_properties/setter]: controlSize */


// The progress indicator’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlTint
func (p_ ProgressIndicator) ControlTint() ControlTint {
	rv := objc.Send[ControlTint](p_.ID, objc.Sel("controlTint"))
	return rv
}/* debug [instance_properties/getter]: controlTint */


// The progress indicator’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlTint
func (p_ ProgressIndicator) SetControlTint(value ControlTint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlTint:"), value)
}/* debug [instance_properties/setter]: controlTint */


// The value that indicates the current extent of the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/doubleValue
func (p_ ProgressIndicator) DoubleValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("doubleValue"))
	return rv
}/* debug [instance_properties/getter]: doubleValue */


// The value that indicates the current extent of the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/doubleValue
func (p_ ProgressIndicator) SetDoubleValue(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDoubleValue:"), value)
}/* debug [instance_properties/setter]: doubleValue */


// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isBezeled
func (p_ ProgressIndicator) Bezeled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("bezeled"))
	return rv
}/* debug [instance_properties/getter]: bezeled */


// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isBezeled
func (p_ ProgressIndicator) SetBezeled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBezeled:"), value)
}/* debug [instance_properties/setter]: bezeled */


// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isDisplayedWhenStopped
func (p_ ProgressIndicator) DisplayedWhenStopped() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("displayedWhenStopped"))
	return rv
}/* debug [instance_properties/getter]: displayedWhenStopped */


// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isDisplayedWhenStopped
func (p_ ProgressIndicator) SetDisplayedWhenStopped(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayedWhenStopped:"), value)
}/* debug [instance_properties/setter]: displayedWhenStopped */


// A Boolean that indicates whether the progress indicator is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isIndeterminate
func (p_ ProgressIndicator) Indeterminate() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("indeterminate"))
	return rv
}/* debug [instance_properties/getter]: indeterminate */


// A Boolean that indicates whether the progress indicator is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isIndeterminate
func (p_ ProgressIndicator) SetIndeterminate(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndeterminate:"), value)
}/* debug [instance_properties/setter]: indeterminate */


// The maximum value for the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/maxValue
func (p_ ProgressIndicator) MaxValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("maxValue"))
	return rv
}/* debug [instance_properties/getter]: maxValue */


// The maximum value for the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/maxValue
func (p_ ProgressIndicator) SetMaxValue(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaxValue:"), value)
}/* debug [instance_properties/setter]: maxValue */


// The minimum value for the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/minValue
func (p_ ProgressIndicator) MinValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minValue"))
	return rv
}/* debug [instance_properties/getter]: minValue */


// The minimum value for the progress indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/minValue
func (p_ ProgressIndicator) SetMinValue(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMinValue:"), value)
}/* debug [instance_properties/setter]: minValue */


// The progress object to use for updating the progress view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/observedProgress
func (p_ ProgressIndicator) ObservedProgress() foundation.Progress {
	rv := objc.Send[foundation.Progress](p_.ID, objc.Sel("observedProgress"))
	return rv
}/* debug [instance_properties/getter]: observedProgress */


// The progress object to use for updating the progress view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/observedProgress
func (p_ ProgressIndicator) SetObservedProgress(value foundation.Progress) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObservedProgress:"), value)
}/* debug [instance_properties/setter]: observedProgress */


// The style of the progress indicator (bar or spinning).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/style-swift.property
func (p_ ProgressIndicator) Style() ProgressIndicatorStyle {
	rv := objc.Send[ProgressIndicatorStyle](p_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// The style of the progress indicator (bar or spinning).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/style-swift.property
func (p_ ProgressIndicator) SetStyle(value ProgressIndicatorStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */


// A Boolean that indicates whether the progress indicator implements animation in a separate thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/usesThreadedAnimation
func (p_ ProgressIndicator) UsesThreadedAnimation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesThreadedAnimation"))
	return rv
}/* debug [instance_properties/getter]: usesThreadedAnimation */


// A Boolean that indicates whether the progress indicator implements animation in a separate thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/usesThreadedAnimation
func (p_ ProgressIndicator) SetUsesThreadedAnimation(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesThreadedAnimation:"), value)
}/* debug [instance_properties/setter]: usesThreadedAnimation */


// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isbezeled
func (p_ ProgressIndicator) IsBezeled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isBezeled"))
	return rv
}/* debug [instance_properties/getter]: isBezeled */


// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isbezeled
func (p_ ProgressIndicator) SetIsBezeled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsBezeled:"), value)
}/* debug [instance_properties/setter]: isBezeled */


// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isdisplayedwhenstopped
func (p_ ProgressIndicator) IsDisplayedWhenStopped() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDisplayedWhenStopped"))
	return rv
}/* debug [instance_properties/getter]: isDisplayedWhenStopped */


// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isdisplayedwhenstopped
func (p_ ProgressIndicator) SetIsDisplayedWhenStopped(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsDisplayedWhenStopped:"), value)
}/* debug [instance_properties/setter]: isDisplayedWhenStopped */


// A Boolean that indicates whether the progress indicator is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isindeterminate
func (p_ ProgressIndicator) IsIndeterminate() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndeterminate"))
	return rv
}/* debug [instance_properties/getter]: isIndeterminate */


// A Boolean that indicates whether the progress indicator is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicator/isindeterminate
func (p_ ProgressIndicator) SetIsIndeterminate(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndeterminate:"), value)
}/* debug [instance_properties/setter]: isIndeterminate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSProgressIndicator */



