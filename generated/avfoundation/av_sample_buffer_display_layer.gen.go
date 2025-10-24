// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSampleBufferDisplayLayer */


/* debug [class_header]: Header for AVSampleBufferDisplayLayer */
// The class instance for the [SampleBufferDisplayLayer] class.
var (
	SampleBufferDisplayLayerClass     _SampleBufferDisplayLayerClass
	SampleBufferDisplayLayerClassOnce sync.Once
)

func getSampleBufferDisplayLayerClass() _SampleBufferDisplayLayerClass {
	SampleBufferDisplayLayerClassOnce.Do(func() {
		SampleBufferDisplayLayerClass = _SampleBufferDisplayLayerClass{objc.GetClass("AVSampleBufferDisplayLayer")}
	})
	return SampleBufferDisplayLayerClass
}

type _SampleBufferDisplayLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SampleBufferDisplayLayer */
// An interface definition for the [SampleBufferDisplayLayer] class.
type ISampleBufferDisplayLayer interface {
	ILayer
	
/* debug [class_interface_properties]: Properties for SampleBufferDisplayLayer */
	// properties:
	ControlTimebase() TimebaseRef /* not a class type */
	SetControlTimebase(value TimebaseRef /* not a class type */)
	Error() Error
	HasSufficientMediaDataForReliablePlaybackStart() bool
	OutputObscuredDueToInsufficientExternalProtection() bool
	ReadyForDisplay() bool
	ReadyForMoreMediaData() bool
	PreventsCapture() bool
	SetPreventsCapture(value bool)
	PreventsDisplaySleepDuringVideoPlayback() bool
	SetPreventsDisplaySleepDuringVideoPlayback(value bool)
	RequiresFlushToResumeDecoding() bool
	SampleBufferRenderer() IAVSampleBufferVideoRenderer
	Status() QueuedSampleBufferRenderingStatus
	Timebase() TimebaseRef /* not a class type */
	VideoGravity() LayerVideoGravity /* typedef */
	SetVideoGravity(value LayerVideoGravity /* typedef */)
	IsOutputObscuredDueToInsufficientExternalProtection() bool
	SetIsOutputObscuredDueToInsufficientExternalProtection(value bool)
	IsReadyForDisplay() bool
	SetIsReadyForDisplay(value bool)
	AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SampleBufferDisplayLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SampleBufferDisplayLayer */
// Alloc allocates a new instance without initialization.
func (sc _SampleBufferDisplayLayerClass) Alloc() SampleBufferDisplayLayer {
	rv := objc.Send[SampleBufferDisplayLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SampleBufferDisplayLayerClass) New() SampleBufferDisplayLayer {
	rv := objc.Send[SampleBufferDisplayLayer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferDisplayLayer) Init() SampleBufferDisplayLayer {
	rv := objc.Send[SampleBufferDisplayLayer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferDisplayLayer) Autorelease() SampleBufferDisplayLayer {
	rv := objc.Send[SampleBufferDisplayLayer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferDisplayLayer creates a new SampleBufferDisplayLayer instance.
func NewSampleBufferDisplayLayer() SampleBufferDisplayLayer {
	return getSampleBufferDisplayLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SampleBufferDisplayLayer */
// An object that displays compressed or uncompressed video frames.


// An object that displays compressed or uncompressed video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer
type SampleBufferDisplayLayer struct {
	Layer
}

// SampleBufferDisplayLayerFrom constructs a [SampleBufferDisplayLayer] from an unsafe.Pointer.
//
// An object that displays compressed or uncompressed video frames.
func SampleBufferDisplayLayerFrom(ptr unsafe.Pointer) SampleBufferDisplayLayer {
	return SampleBufferDisplayLayer{
		Layer: LayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SampleBufferDisplayLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SampleBufferDisplayLayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SampleBufferDisplayLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SampleBufferDisplayLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SampleBufferDisplayLayer */

// A timebase that determines how the layer interprets timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/controlTimebase
func (s_ SampleBufferDisplayLayer) ControlTimebase() TimebaseRef /* not a class type */ {
	rv := objc.Send[TimebaseRef](s_.ID, objc.Sel("controlTimebase"))
	return rv
}/* debug [instance_properties/getter]: controlTimebase */


// A timebase that determines how the layer interprets timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/controlTimebase
func (s_ SampleBufferDisplayLayer) SetControlTimebase(value TimebaseRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setControlTimebase:"), value)
}/* debug [instance_properties/setter]: controlTimebase */


// The error that caused the failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/error
func (s_ SampleBufferDisplayLayer) Error() Error {
	rv := objc.Send[Error](s_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// A Boolean value that indicates whether the enqueued media data meets the renderer’s preroll level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/hasSufficientMediaDataForReliablePlaybackStart
func (s_ SampleBufferDisplayLayer) HasSufficientMediaDataForReliablePlaybackStart() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasSufficientMediaDataForReliablePlaybackStart"))
	return rv
}/* debug [instance_properties/getter]: hasSufficientMediaDataForReliablePlaybackStart */


// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/isOutputObscuredDueToInsufficientExternalProtection
func (s_ SampleBufferDisplayLayer) OutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("outputObscuredDueToInsufficientExternalProtection"))
	return rv
}/* debug [instance_properties/getter]: outputObscuredDueToInsufficientExternalProtection */


// A Boolean value that indicates whether the first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/isReadyForDisplay
func (s_ SampleBufferDisplayLayer) ReadyForDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("readyForDisplay"))
	return rv
}/* debug [instance_properties/getter]: readyForDisplay */


// A Boolean value that indicates the readiness of the layer to accept more sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/isReadyForMoreMediaData
func (s_ SampleBufferDisplayLayer) ReadyForMoreMediaData() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("readyForMoreMediaData"))
	return rv
}/* debug [instance_properties/getter]: readyForMoreMediaData */


// A Boolean value that indicates whether the layer protects against screen capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsCapture
func (s_ SampleBufferDisplayLayer) PreventsCapture() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsCapture"))
	return rv
}/* debug [instance_properties/getter]: preventsCapture */


// A Boolean value that indicates whether the layer protects against screen capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsCapture
func (s_ SampleBufferDisplayLayer) SetPreventsCapture(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreventsCapture:"), value)
}/* debug [instance_properties/setter]: preventsCapture */


// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsDisplaySleepDuringVideoPlayback
func (s_ SampleBufferDisplayLayer) PreventsDisplaySleepDuringVideoPlayback() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsDisplaySleepDuringVideoPlayback"))
	return rv
}/* debug [instance_properties/getter]: preventsDisplaySleepDuringVideoPlayback */


// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsDisplaySleepDuringVideoPlayback
func (s_ SampleBufferDisplayLayer) SetPreventsDisplaySleepDuringVideoPlayback(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreventsDisplaySleepDuringVideoPlayback:"), value)
}/* debug [instance_properties/setter]: preventsDisplaySleepDuringVideoPlayback */


// A Boolean value that indicates whether the layer needs to flush its state to continue decoding frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/requiresFlushToResumeDecoding
func (s_ SampleBufferDisplayLayer) RequiresFlushToResumeDecoding() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("requiresFlushToResumeDecoding"))
	return rv
}/* debug [instance_properties/getter]: requiresFlushToResumeDecoding */


// An object that enqueues video sample buffers for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/sampleBufferRenderer
func (s_ SampleBufferDisplayLayer) SampleBufferRenderer() IAVSampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](s_.ID, objc.Sel("sampleBufferRenderer"))
	return rv
}/* debug [instance_properties/getter]: sampleBufferRenderer */


// The ability of the display layer to be used for enqueuing sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/status
func (s_ SampleBufferDisplayLayer) Status() QueuedSampleBufferRenderingStatus {
	rv := objc.Send[QueuedSampleBufferRenderingStatus](s_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// The renderer’s timebase, which determines how the layer interprets time stamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/timebase
func (s_ SampleBufferDisplayLayer) Timebase() TimebaseRef /* not a class type */ {
	rv := objc.Send[TimebaseRef](s_.ID, objc.Sel("timebase"))
	return rv
}/* debug [instance_properties/getter]: timebase */


// A value that indicates how the layer displays video within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/videoGravity
func (s_ SampleBufferDisplayLayer) VideoGravity() LayerVideoGravity /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("videoGravity"))
	return rv
}/* debug [instance_properties/getter]: videoGravity */


// A value that indicates how the layer displays video within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/videoGravity
func (s_ SampleBufferDisplayLayer) SetVideoGravity(value LayerVideoGravity /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVideoGravity:"), value)
}/* debug [instance_properties/setter]: videoGravity */


// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isoutputobscuredduetoinsufficientexternalprotection
func (s_ SampleBufferDisplayLayer) IsOutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isOutputObscuredDueToInsufficientExternalProtection"))
	return rv
}/* debug [instance_properties/getter]: isOutputObscuredDueToInsufficientExternalProtection */


// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isoutputobscuredduetoinsufficientexternalprotection
func (s_ SampleBufferDisplayLayer) SetIsOutputObscuredDueToInsufficientExternalProtection(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsOutputObscuredDueToInsufficientExternalProtection:"), value)
}/* debug [instance_properties/setter]: isOutputObscuredDueToInsufficientExternalProtection */


// A Boolean value that indicates whether the first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isreadyfordisplay
func (s_ SampleBufferDisplayLayer) IsReadyForDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isReadyForDisplay"))
	return rv
}/* debug [instance_properties/getter]: isReadyForDisplay */


// A Boolean value that indicates whether the first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isreadyfordisplay
func (s_ SampleBufferDisplayLayer) SetIsReadyForDisplay(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsReadyForDisplay:"), value)
}/* debug [instance_properties/setter]: isReadyForDisplay */


// The key for the corresponding error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayerfailedtodecodenotificationerrorkey
func (s_ SampleBufferDisplayLayer) AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey"))
	return rv
}/* debug [instance_properties/getter]: AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSampleBufferDisplayLayer */


