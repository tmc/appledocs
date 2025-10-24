// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCStream */


/* debug [class_header]: Header for SCStream */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Stream */
// An interface definition for the [Stream] class.
type IStream interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Stream */
	// properties:
	SynchronizationClock() ClockRef /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Stream */
	// methods:
	AddRecordingOutputError(recordingOutput ISCRecordingOutput, error_ unsafe.Pointer) bool
	AddStreamOutputTypeSampleHandlerQueueError(output unsafe.Pointer, type_ StreamOutputType, sampleHandlerQueue unsafe.Pointer, error_ unsafe.Pointer) bool
	RemoveRecordingOutputError(recordingOutput ISCRecordingOutput, error_ unsafe.Pointer) bool
	RemoveStreamOutputTypeError(output unsafe.Pointer, type_ StreamOutputType, error_ unsafe.Pointer) bool
	StartCaptureWithCompletionHandler(completionHandler unsafe.Pointer)
	StopCaptureWithCompletionHandler(completionHandler unsafe.Pointer)
	UpdateConfigurationCompletionHandler(streamConfig ISCStreamConfiguration, completionHandler unsafe.Pointer)
	UpdateContentFilterCompletionHandler(contentFilter ISCContentFilter, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Stream */
// Alloc allocates a new instance without initialization.
func (sc _StreamClass) Alloc() Stream {
	rv := objc.Send[Stream](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Stream */
// An instance that represents a stream of shareable content.
//
// Use a stream to capture video of screen content like apps and windows. Create a content stream by passing it an instance of and an object. The stream uses the filter to determine which screen content to capture, and uses the configuration data to configure the output.


// An instance that represents a stream of shareable content.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Stream */

// Creates a stream with a content filter and configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/init(filter:configuration:delegate:)
func NewStreamWithFilterConfigurationDelegate(contentFilter ISCContentFilter, streamConfig ISCStreamConfiguration, delegate unsafe.Pointer) Stream {
	instance := getStreamClass().Alloc()
	rv := objc.Send[Stream](instance.ID, objc.Sel("initWithFilter:configuration:delegate:"), contentFilter, streamConfig, delegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewStreamWithFilterConfigurationDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Stream */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Stream */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Stream */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/addRecordingOutput(_:)
func (s_ Stream) AddRecordingOutputError(recordingOutput ISCRecordingOutput, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("addRecordingOutput:error:"), recordingOutput, error_)
	return rv
}/* debug [instance_methods/method]: AddRecordingOutputError */


// Adds a destination that receives the stream output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/addStreamOutput(_:type:sampleHandlerQueue:)
func (s_ Stream) AddStreamOutputTypeSampleHandlerQueueError(output unsafe.Pointer, type_ StreamOutputType, sampleHandlerQueue unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("addStreamOutput:type:sampleHandlerQueue:error:"), output, type_, sampleHandlerQueue, error_)
	return rv
}/* debug [instance_methods/method]: AddStreamOutputTypeSampleHandlerQueueError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/removeRecordingOutput(_:)
func (s_ Stream) RemoveRecordingOutputError(recordingOutput ISCRecordingOutput, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("removeRecordingOutput:error:"), recordingOutput, error_)
	return rv
}/* debug [instance_methods/method]: RemoveRecordingOutputError */


// Removes a destination from receiving stream output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/removeStreamOutput(_:type:)
func (s_ Stream) RemoveStreamOutputTypeError(output unsafe.Pointer, type_ StreamOutputType, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("removeStreamOutput:type:error:"), output, type_, error_)
	return rv
}/* debug [instance_methods/method]: RemoveStreamOutputTypeError */


// Starts the stream with a callback to indicate whether it successfully starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/startCapture(completionHandler:)
func (s_ Stream) StartCaptureWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("startCaptureWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: StartCaptureWithCompletionHandler */


// Stops the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/stopCapture(completionHandler:)
func (s_ Stream) StopCaptureWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("stopCaptureWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: StopCaptureWithCompletionHandler */


// Updates the stream with a new configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/updateConfiguration(_:completionHandler:)
func (s_ Stream) UpdateConfigurationCompletionHandler(streamConfig ISCStreamConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateConfiguration:completionHandler:"), streamConfig, completionHandler)
}/* debug [instance_methods/method]: UpdateConfigurationCompletionHandler */


// Updates the stream by applying a new content filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/updateContentFilter(_:completionHandler:)
func (s_ Stream) UpdateContentFilterCompletionHandler(contentFilter ISCContentFilter, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateContentFilter:completionHandler:"), contentFilter, completionHandler)
}/* debug [instance_methods/method]: UpdateContentFilterCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Stream */

// A clock to use for output synchronization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/synchronizationClock
func (s_ Stream) SynchronizationClock() ClockRef /* not a class type */ {
	rv := objc.Send[ClockRef](s_.ID, objc.Sel("synchronizationClock"))
	return rv
}/* debug [instance_properties/getter]: synchronizationClock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCStream */


