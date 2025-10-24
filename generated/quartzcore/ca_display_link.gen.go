// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [DisplayLink] class.
type IDisplayLink interface {
	objectivec.IObject
	// properties:
	PreferredFrameRateRange() FrameRateRange /* not a class type */
	SetPreferredFrameRateRange(value FrameRateRange /* not a class type */)
	Duration() float64
	SetDuration(value float64)
	FrameInterval() int
	SetFrameInterval(value int)
	IsPaused() bool
	SetIsPaused(value bool)
	PreferredFramesPerSecond() int
	SetPreferredFramesPerSecond(value int)
	TargetTimestamp() float64
	SetTargetTimestamp(value float64)
	Timestamp() float64
	SetTimestamp(value float64)
	// methods:
	AddToRunLoopForMode(runloop objc.IObject /* cross-framework: RunLoop */, mode RunLoopMode /* not a class type */)
}

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

// Alloc allocates a new instance without initialization.
func (dc _DisplayLinkClass) Alloc() DisplayLink {
	rv := objc.Send[DisplayLink](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Registers the display link with a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/add(to:forMode:)
func (d_ DisplayLink) AddToRunLoopForMode(runloop objc.IObject /* cross-framework: RunLoop */, mode RunLoopMode /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addToRunLoop:forMode:"), runloop, mode)
}


// A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFrameRateRange
func (d_ DisplayLink) PreferredFrameRateRange() FrameRateRange /* not a class type */ {
	rv := objc.Send[FrameRateRange](d_.ID, objc.Sel("preferredFrameRateRange"))
	return rv
}


// A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFrameRateRange
func (d_ DisplayLink) SetPreferredFrameRateRange(value FrameRateRange /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreferredFrameRateRange:"), value)
}


// The time interval between screen refresh updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/duration
func (d_ DisplayLink) Duration() float64 {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("duration"))
	return rv
}


// The time interval between screen refresh updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/duration
func (d_ DisplayLink) SetDuration(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDuration:"), value)
}


// The number of frames that must pass before the display link notifies the target again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/frameinterval
func (d_ DisplayLink) FrameInterval() int {
	rv := objc.Send[int](d_.ID, objc.Sel("frameInterval"))
	return rv
}


// The number of frames that must pass before the display link notifies the target again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/frameinterval
func (d_ DisplayLink) SetFrameInterval(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFrameInterval:"), value)
}


// A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/ispaused
func (d_ DisplayLink) IsPaused() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isPaused"))
	return rv
}


// A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/ispaused
func (d_ DisplayLink) SetIsPaused(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsPaused:"), value)
}


// A frequency your app prefers for frame updates, affecting how often the system invokes your delegate’s callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/preferredframespersecond
func (d_ DisplayLink) PreferredFramesPerSecond() int {
	rv := objc.Send[int](d_.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}


// A frequency your app prefers for frame updates, affecting how often the system invokes your delegate’s callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/preferredframespersecond
func (d_ DisplayLink) SetPreferredFramesPerSecond(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreferredFramesPerSecond:"), value)
}


// The time interval that represents when the next frame displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/targettimestamp
func (d_ DisplayLink) TargetTimestamp() float64 {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("targetTimestamp"))
	return rv
}


// The time interval that represents when the next frame displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/targettimestamp
func (d_ DisplayLink) SetTargetTimestamp(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTargetTimestamp:"), value)
}


// The time interval that represents when the last frame displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/timestamp
func (d_ DisplayLink) Timestamp() float64 {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("timestamp"))
	return rv
}


// The time interval that represents when the last frame displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cadisplaylink/timestamp
func (d_ DisplayLink) SetTimestamp(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimestamp:"), value)
}



