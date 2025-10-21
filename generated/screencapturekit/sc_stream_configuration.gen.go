// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StreamConfiguration] class.
var (
	StreamConfigurationClass     _StreamConfigurationClass
	StreamConfigurationClassOnce sync.Once
)

func getStreamConfigurationClass() _StreamConfigurationClass {
	StreamConfigurationClassOnce.Do(func() {
		StreamConfigurationClass = _StreamConfigurationClass{objc.GetClass("SCStreamConfiguration")}
	})
	return StreamConfigurationClass
}

type _StreamConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [StreamConfiguration] class.
type IStreamConfiguration interface {
	objectivec.IObject
}

// An instance that provides the output configuration for a stream.
//
// Creating an instance of this class provides a default configuration for a stream. Only configure its properties if you need to customize the output.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration
type StreamConfiguration struct {
	objectivec.Object
}

// StreamConfigurationFrom constructs a [StreamConfiguration] from an unsafe.Pointer.
//
// An instance that provides the output configuration for a stream.
func StreamConfigurationFrom(ptr unsafe.Pointer) StreamConfiguration {
	return StreamConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StreamConfigurationClass) Alloc() StreamConfiguration {
	rv := objc.Send[StreamConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StreamConfigurationClass) New() StreamConfiguration {
	rv := objc.Send[StreamConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StreamConfiguration) Init() StreamConfiguration {
	rv := objc.Send[StreamConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StreamConfiguration) Autorelease() StreamConfiguration {
	rv := objc.Send[StreamConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStreamConfiguration creates a new StreamConfiguration instance.
func NewStreamConfiguration() StreamConfiguration {
	return getStreamConfigurationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureDynamicRange
func (s_ StreamConfiguration) CaptureDynamicRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("captureDynamicRange"))
	return rv
}


// SetCaptureDynamicRange sets the value of the captureDynamicRange property.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureDynamicRange
func (s_ StreamConfiguration) SetCaptureDynamicRange(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaptureDynamicRange:"), value)
}

// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreGlobalClipSingleWindow
func (s_ StreamConfiguration) IgnoreGlobalClipSingleWindow() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreGlobalClipSingleWindow"))
	return rv
}


// SetIgnoreGlobalClipSingleWindow sets the value of the ignoreGlobalClipSingleWindow property.
// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreGlobalClipSingleWindow
func (s_ StreamConfiguration) SetIgnoreGlobalClipSingleWindow(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreGlobalClipSingleWindow:"), value)
}

// The maximum number of frames for the queue to store.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/queueDepth
func (s_ StreamConfiguration) QueueDepth() int {
	rv := objc.Send[int](s_.ID, objc.Sel("queueDepth"))
	return rv
}


// SetQueueDepth sets the value of the queueDepth property.
// The maximum number of frames for the queue to store.

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/queueDepth
func (s_ StreamConfiguration) SetQueueDepth(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setQueueDepth:"), value)
}



