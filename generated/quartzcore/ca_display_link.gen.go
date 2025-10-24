// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CADisplayLink */


/* debug [class_header]: Header for CADisplayLink */
// The class instance for the [DisplayLink] class.
var (
	DisplayLinkClass     _DisplayLinkClass
	DisplayLinkClassOnce sync.Once
)

func getDisplayLinkClass() _DisplayLinkClass {
	DisplayLinkClassOnce.Do(func() {
		DisplayLinkClass = _DisplayLinkClass{objc.GetClass("CADisplayLink")}
	})
	return DisplayLinkClass
}

type _DisplayLinkClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DisplayLink */
// An interface definition for the [DisplayLink] class.
type IDisplayLink interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DisplayLink */
	// properties:
	Duration() float64
	Paused() bool
	SetPaused(value bool)
	PreferredFrameRateRange() objc.IObject /* cross-framework: CAFrameRateRange */
	SetPreferredFrameRateRange(value objc.IObject /* cross-framework: CAFrameRateRange */)
	TargetTimestamp() float64
	Timestamp() float64
	IsPaused() bool
	SetIsPaused(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DisplayLink */
	// methods:
	AddToRunLoopForMode(runloop foundation.RunLoop, mode RunLoopMode /* not a class type */)
	Invalidate()
	RemoveFromRunLoopForMode(runloop foundation.RunLoop, mode RunLoopMode /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DisplayLink */
// Alloc allocates a new instance without initialization.
func (dc _DisplayLinkClass) Alloc() DisplayLink {
	rv := objc.Send[DisplayLink](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DisplayLinkClass) New() DisplayLink {
	rv := objc.Send[DisplayLink](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DisplayLink) Init() DisplayLink {
	rv := objc.Send[DisplayLink](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DisplayLink) Autorelease() DisplayLink {
	rv := objc.Send[DisplayLink](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDisplayLink creates a new DisplayLink instance.
func NewDisplayLink() DisplayLink {
	return getDisplayLinkClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DisplayLink */
// A timer object that allows your app to synchronize its drawing to the refresh rate of the display.
//
// Your app initializes a new display link by providing a target object and a selector to call when the system updates the screen. To synchronize your display loop with the display, your application adds it to a run loop using the method. Once you associate the display link with a run loop, the system calls the selector on the target when the screen’s contents need to update. The target can read the display link’s property to retrieve the time the system displayed the previous frame. For example, an app that displays movies might use to calculate which video frame to display next. An app that performs its own animations might use to determine where and how visible objects appear in the upcoming frame. The property provides the amount of time between frames at the . To calculate the actual frame duration, use - . You can use this value in your app to calculate the frame rate of the display, the approximate time the system displays the next frame, and to adjust the drawing behavior so that the next frame is ready in time to display. Your app can disable notifications by setting to . Also, if your app can’t provide frames in the time the system provides, you may want to choose a slower frame rate. An app with a slower but consistent frame rate appears smoother to the user than an app that skips frames. You can define the number of frames per second by setting . When your app finishes with a display link, call to remove it from all run loops and to disassociate it from the target. The code listing below shows how to create a display link and add it to the current run loop. The display link invokes the step function, which prints the target timestamp with each screen update. You shouldn’t subclass .


// A timer object that allows your app to synchronize its drawing to the refresh rate of the display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink
type DisplayLink struct {
	objectivec.Object
}

// DisplayLinkFrom constructs a [DisplayLink] from an unsafe.Pointer.
//
// A timer object that allows your app to synchronize its drawing to the refresh rate of the display.
func DisplayLinkFrom(ptr unsafe.Pointer) DisplayLink {
	return DisplayLink{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DisplayLink */

// Creates a display link for a target that calls its selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/init(target:selector:)
func NewDisplayLinkWithTargetSelector(target objc.IObject, sel objc.SEL) DisplayLink {
	rv := objc.Send[DisplayLink](objc.ID(getDisplayLinkClass().class), objc.Sel("displayLinkWithTarget:selector:"), target, sel)
	return rv
}/* debug [class_init_methods/constructor]: NewDisplayLinkWithTargetSelector */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DisplayLink */

// Creates a display link for a target that calls its selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/init(target:selector:)
func (dc _DisplayLinkClass) DisplayLinkWithTargetSelector(target objc.IObject, sel objc.SEL) IDisplayLink {
	rv := objc.Send[DisplayLink](objc.ID(dc.class), objc.Sel("displayLinkWithTarget:selector:"), target, sel)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DisplayLinkWithTargetSelector) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DisplayLink */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DisplayLink */

// Registers the display link with a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/add(to:forMode:)
func (d_ DisplayLink) AddToRunLoopForMode(runloop foundation.RunLoop, mode RunLoopMode /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addToRunLoop:forMode:"), runloop, mode)
}/* debug [instance_methods/method]: AddToRunLoopForMode */


// Removes the display link from all run loop modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/invalidate()
func (d_ DisplayLink) Invalidate() {
	objc.Send[objc.ID](d_.ID, objc.Sel("invalidate"))
}/* debug [instance_methods/method]: Invalidate */


// Removes the display link from the run loop for the given mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/remove(from:forMode:)
func (d_ DisplayLink) RemoveFromRunLoopForMode(runloop foundation.RunLoop, mode RunLoopMode /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("removeFromRunLoop:forMode:"), runloop, mode)
}/* debug [instance_methods/method]: RemoveFromRunLoopForMode */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DisplayLink */

// The time interval between screen refresh updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/duration
func (d_ DisplayLink) Duration() float64 {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/isPaused
func (d_ DisplayLink) Paused() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("paused"))
	return rv
}/* debug [instance_properties/getter]: paused */


// A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/isPaused
func (d_ DisplayLink) SetPaused(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPaused:"), value)
}/* debug [instance_properties/setter]: paused */


// A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFrameRateRange
func (d_ DisplayLink) PreferredFrameRateRange() objc.IObject /* cross-framework: CAFrameRateRange */ {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("preferredFrameRateRange"))
	return rv
}/* debug [instance_properties/getter]: preferredFrameRateRange */


// A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFrameRateRange
func (d_ DisplayLink) SetPreferredFrameRateRange(value objc.IObject /* cross-framework: CAFrameRateRange */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreferredFrameRateRange:"), value)
}/* debug [instance_properties/setter]: preferredFrameRateRange */


// The time interval that represents when the next frame displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/targetTimestamp
func (d_ DisplayLink) TargetTimestamp() float64 {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("targetTimestamp"))
	return rv
}/* debug [instance_properties/getter]: targetTimestamp */


// The time interval that represents when the last frame displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/timestamp
func (d_ DisplayLink) Timestamp() float64 {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("timestamp"))
	return rv
}/* debug [instance_properties/getter]: timestamp */


// A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/ispaused
func (d_ DisplayLink) IsPaused() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isPaused"))
	return rv
}/* debug [instance_properties/getter]: isPaused */


// A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/ispaused
func (d_ DisplayLink) SetIsPaused(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsPaused:"), value)
}/* debug [instance_properties/setter]: isPaused */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CADisplayLink */


