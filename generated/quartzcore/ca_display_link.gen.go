// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	AddToRunLoopForMode(runloop unsafe.Pointer, mode unsafe.Pointer)
	Invalidate()
	RemoveFromRunLoopForMode(runloop unsafe.Pointer, mode unsafe.Pointer)
}

// A timer object that allows your app to synchronize its drawing to the refresh rate of the display.
//
// Your app initializes a new display link by providing a target object and a selector to call when the system updates the screen. To synchronize your display loop with the display, your application adds it to a run loop using the method. Once you associate the display link with a run loop, the system calls the selector on the target when the screen’s contents need to update. The target can read the display link’s property to retrieve the time the system displayed the previous frame. For example, an app that displays movies might use to calculate which video frame to display next. An app that performs its own animations might use to determine where and how visible objects appear in the upcoming frame. The property provides the amount of time between frames at the . To calculate the actual frame duration, use - . You can use this value in your app to calculate the frame rate of the display, the approximate time the system displays the next frame, and to adjust the drawing behavior so that the next frame is ready in time to display. Your app can disable notifications by setting to . Also, if your app can’t provide frames in the time the system provides, you may want to choose a slower frame rate. An app with a slower but consistent frame rate appears smoother to the user than an app that skips frames. You can define the number of frames per second by setting . When your app finishes with a display link, call to remove it from all run loops and to disassociate it from the target. The code listing below shows how to create a display link and add it to the current run loop. The display link invokes the step function, which prints the target timestamp with each screen update. You shouldn’t subclass .
//
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




// Creates a display link for a target that calls its selector.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/init(target:selector:)
func NewDisplayLinkWithTargetSelector(target objc.ID, sel objc.SEL) DisplayLink {
	rv := objc.Send[DisplayLink](objc.ID(getDisplayLinkClass().class), objc.Sel("displayLinkWithTarget:selector:"), target, sel)
	return rv
}


// Creates a display link for a target that calls its selector.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/init(target:selector:)
func (dc _DisplayLinkClass) DisplayLinkWithTargetSelector(target objc.ID, sel objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("displayLinkWithTarget:selector:"), target, sel)
	return rv
}

// Registers the display link with a run loop.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/add(to:forMode:)
func (d_ DisplayLink) AddToRunLoopForMode(runloop unsafe.Pointer, mode unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addToRunLoop:forMode:"), runloop, mode)
}

// Removes the display link from all run loop modes.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/invalidate()
func (d_ DisplayLink) Invalidate() {
	objc.Send[objc.ID](d_.ID, objc.Sel("invalidate"))
}

// Removes the display link from the run loop for the given mode.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/remove(from:forMode:)
func (d_ DisplayLink) RemoveFromRunLoopForMode(runloop unsafe.Pointer, mode unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("removeFromRunLoop:forMode:"), runloop, mode)
}

// The time interval between screen refresh updates.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/duration
func (d_ DisplayLink) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("duration"))
	return rv
}

// The number of frames that must pass before the display link notifies the target again.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/frameInterval
func (d_ DisplayLink) FrameInterval() int {
	rv := objc.Send[int](d_.ID, objc.Sel("frameInterval"))
	return rv
}


// SetFrameInterval sets the value of the frameInterval property.
// The number of frames that must pass before the display link notifies the target again.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/frameInterval
func (d_ DisplayLink) SetFrameInterval(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFrameInterval:"), value)
}

// A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/isPaused
func (d_ DisplayLink) Paused() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("paused"))
	return rv
}


// SetPaused sets the value of the paused property.
// A Boolean value that indicates whether the system suspends the display link’s notifications to the target.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/isPaused
func (d_ DisplayLink) SetPaused(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPaused:"), value)
}

// A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFrameRateRange
func (d_ DisplayLink) PreferredFrameRateRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("preferredFrameRateRange"))
	return rv
}


// SetPreferredFrameRateRange sets the value of the preferredFrameRateRange property.
// A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFrameRateRange
func (d_ DisplayLink) SetPreferredFrameRateRange(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreferredFrameRateRange:"), value)
}

// A frequency your app prefers for frame updates, affecting how often the system invokes your delegate’s callback.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFramesPerSecond
func (d_ DisplayLink) PreferredFramesPerSecond() int {
	rv := objc.Send[int](d_.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}


// SetPreferredFramesPerSecond sets the value of the preferredFramesPerSecond property.
// A frequency your app prefers for frame updates, affecting how often the system invokes your delegate’s callback.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFramesPerSecond
func (d_ DisplayLink) SetPreferredFramesPerSecond(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreferredFramesPerSecond:"), value)
}

// The time interval that represents when the next frame displays.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/targetTimestamp
func (d_ DisplayLink) TargetTimestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("targetTimestamp"))
	return rv
}

// The time interval that represents when the last frame displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/timestamp
func (d_ DisplayLink) Timestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("timestamp"))
	return rv
}


