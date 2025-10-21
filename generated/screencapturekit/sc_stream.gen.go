// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Stream] class.
var (
	StreamClass     _StreamClass
	StreamClassOnce sync.Once
)

func getStreamClass() _StreamClass {
	StreamClassOnce.Do(func() {
		StreamClass = _StreamClass{objc.GetClass("SCStream")}
	})
	return StreamClass
}

type _StreamClass struct {
	class objc.Class
}

// An interface definition for the [Stream] class.
type IStream interface {
	objectivec.IObject
	AddStreamOutputTypeSampleHandlerQueueError(output objc.ID, type_ unsafe.Pointer, sampleHandlerQueue unsafe.Pointer, error_ unsafe.Pointer) bool
	StartCaptureWithCompletionHandler(completionHandler unsafe.Pointer)
	UpdateContentFilterCompletionHandler(contentFilter unsafe.Pointer, completionHandler unsafe.Pointer)
}

// An instance that represents a stream of shareable content.
//
// Use a stream to capture video of screen content like apps and windows. Create a content stream by passing it an instance of and an object. The stream uses the filter to determine which screen content to capture, and uses the configuration data to configure the output.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream
type Stream struct {
	objectivec.Object
}

// StreamFrom constructs a [Stream] from an unsafe.Pointer.
//
// An instance that represents a stream of shareable content.
func StreamFrom(ptr unsafe.Pointer) Stream {
	return Stream{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StreamClass) Alloc() Stream {
	rv := objc.Send[Stream](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StreamClass) New() Stream {
	rv := objc.Send[Stream](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Stream) Init() Stream {
	rv := objc.Send[Stream](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Stream) Autorelease() Stream {
	rv := objc.Send[Stream](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStream creates a new Stream instance.
func NewStream() Stream {
	return getStreamClass().New()
}


// Adds a destination that receives the stream output.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/addStreamOutput(_:type:sampleHandlerQueue:)
func (s_ Stream) AddStreamOutputTypeSampleHandlerQueueError(output objc.ID, type_ unsafe.Pointer, sampleHandlerQueue unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("addStreamOutput:type:sampleHandlerQueue:error:"), output, type_, sampleHandlerQueue, error_)
	return rv
}

// Starts the stream with a callback to indicate whether it successfully starts.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/startCapture(completionHandler:)
func (s_ Stream) StartCaptureWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("startCaptureWithCompletionHandler:"), completionHandler)
}

// Updates the stream by applying a new content filter.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/updateContentFilter(_:completionHandler:)
func (s_ Stream) UpdateContentFilterCompletionHandler(contentFilter unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateContentFilter:completionHandler:"), contentFilter, completionHandler)
}

// A clock to use for output synchronization.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstream/synchronizationclock
func (s_ Stream) SynchronizationClock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("synchronizationClock"))
	return rv
}


// SetSynchronizationClock sets the value of the synchronizationClock property.
// A clock to use for output synchronization.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstream/synchronizationclock
func (s_ Stream) SetSynchronizationClock(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSynchronizationClock:"), value)
}



