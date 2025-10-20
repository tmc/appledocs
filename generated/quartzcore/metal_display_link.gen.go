// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetalDisplayLink] class.
var (
	metalDisplayLinkClass     _MetalDisplayLinkClass
	metalDisplayLinkClassOnce sync.Once
)

func getMetalDisplayLinkClass() _MetalDisplayLinkClass {
	metalDisplayLinkClassOnce.Do(func() {
		metalDisplayLinkClass = _MetalDisplayLinkClass{objc.GetClass("CAMetalDisplayLink")}
	})
	return metalDisplayLinkClass
}

type _MetalDisplayLinkClass struct {
	class objc.Class
}

// An interface definition for the [MetalDisplayLink] class.
type IMetalDisplayLink interface {
	objectivec.IObject
}

// A class your Metal app uses to register for callbacks to synchronize its animations for a display.
//
// instances are a specialized way to interact with variable-rate displays when you need more control over the timing window to render your app’s frames. Controlling the timing window and rendering delay for frames can help you achieve smoother frame rates and avoid visual artifacts. Your app initializes a new Metal display link by providing a target . Set this instance’s property to an implementation that encodes the rendering work for Metal to perform. With a set delegate, synchronize the display with a run loop to perform rendering on by calling the method. Once you associate the display link with a run loop, the system calls the delegate’s method to request new frames. This method receives update requests based on the and of the display link. The system makes a best effort to make callbacks at appropriate times. Your app should complete any commits to the Metal device’s for rendering the display layer before calling on a drawable element. Your app can disable notifications by setting to . When your app finishes with a display link, call to remove it from all run loops and the target.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink
type MetalDisplayLink struct {
	objectivec.Object
}

// MetalDisplayLinkFrom constructs a [MetalDisplayLink] from an unsafe.Pointer.
//
// A class your Metal app uses to register for callbacks to synchronize its animations for a display.
func MetalDisplayLinkFrom(ptr unsafe.Pointer) MetalDisplayLink {
	return MetalDisplayLink{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetalDisplayLinkClass) Alloc() MetalDisplayLink {
	rv := objc.Send[MetalDisplayLink](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetalDisplayLinkClass) New() MetalDisplayLink {
	rv := objc.Send[MetalDisplayLink](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetalDisplayLink) Init() MetalDisplayLink {
	rv := objc.Send[MetalDisplayLink](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetalDisplayLink) Autorelease() MetalDisplayLink {
	rv := objc.Send[MetalDisplayLink](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetalDisplayLink creates a new MetalDisplayLink instance.
func NewMetalDisplayLink() MetalDisplayLink {
	return getMetalDisplayLinkClass().New()
}


// An instance of a type your app implements that responds to the system’s callbacks.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/delegate
func (m_ MetalDisplayLink) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}

// SetDelegate sets the value of the delegate property.
// An instance of a type your app implements that responds to the system’s callbacks.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/delegate
func (m_ MetalDisplayLink) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}
// A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/preferredFrameRateRange
func (m_ MetalDisplayLink) PreferredFrameRateRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("preferredFrameRateRange"))
	return rv
}

// SetPreferredFrameRateRange sets the value of the preferredFrameRateRange property.
// A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/preferredFrameRateRange
func (m_ MetalDisplayLink) SetPreferredFrameRateRange(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredFrameRateRange:"), value)
}


