// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [SampleBufferDisplayLayer] class.
type ISampleBufferDisplayLayer interface {
	ILayer
	

	// properties:
	ControlTimebase() TimebaseRef /* not a class type */
	SetControlTimebase(value TimebaseRef /* not a class type */)
	Error() foundation.foundation.INSError
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
	VideoGravity() LayerVideoGravity
	SetVideoGravity(value LayerVideoGravity)
	IsOutputObscuredDueToInsufficientExternalProtection() bool
	SetIsOutputObscuredDueToInsufficientExternalProtection(value bool)
	IsReadyForDisplay() bool
	SetIsReadyForDisplay(value bool)
	AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey() foundation.foundation.INSString


	

	// methods:


}





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

























// A timebase that determines how the layer interprets timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/controlTimebase
func (s_ SampleBufferDisplayLayer) ControlTimebase() TimebaseRef /* not a class type */ {
	rv := objc.Send[TimebaseRef](s_.ID, objc.Sel("controlTimebase"))
	return rv
}


// A timebase that determines how the layer interprets timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/controlTimebase
func (s_ SampleBufferDisplayLayer) SetControlTimebase(value TimebaseRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setControlTimebase:"), value)
}


// The error that caused the failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/error
func (s_ SampleBufferDisplayLayer) Error() foundation.foundation.INSError {
	rv := objc.Send[foundation.NSError](s_.ID, objc.Sel("error"))
	return rv
}


// A Boolean value that indicates whether the enqueued media data meets the renderer’s preroll level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/hasSufficientMediaDataForReliablePlaybackStart
func (s_ SampleBufferDisplayLayer) HasSufficientMediaDataForReliablePlaybackStart() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasSufficientMediaDataForReliablePlaybackStart"))
	return rv
}


// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/isOutputObscuredDueToInsufficientExternalProtection
func (s_ SampleBufferDisplayLayer) OutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("outputObscuredDueToInsufficientExternalProtection"))
	return rv
}


// A Boolean value that indicates whether the first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/isReadyForDisplay
func (s_ SampleBufferDisplayLayer) ReadyForDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("readyForDisplay"))
	return rv
}


// A Boolean value that indicates the readiness of the layer to accept more sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/isReadyForMoreMediaData
func (s_ SampleBufferDisplayLayer) ReadyForMoreMediaData() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("readyForMoreMediaData"))
	return rv
}


// A Boolean value that indicates whether the layer protects against screen capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsCapture
func (s_ SampleBufferDisplayLayer) PreventsCapture() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsCapture"))
	return rv
}


// A Boolean value that indicates whether the layer protects against screen capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsCapture
func (s_ SampleBufferDisplayLayer) SetPreventsCapture(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreventsCapture:"), value)
}


// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsDisplaySleepDuringVideoPlayback
func (s_ SampleBufferDisplayLayer) PreventsDisplaySleepDuringVideoPlayback() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsDisplaySleepDuringVideoPlayback"))
	return rv
}


// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsDisplaySleepDuringVideoPlayback
func (s_ SampleBufferDisplayLayer) SetPreventsDisplaySleepDuringVideoPlayback(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreventsDisplaySleepDuringVideoPlayback:"), value)
}


// A Boolean value that indicates whether the layer needs to flush its state to continue decoding frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/requiresFlushToResumeDecoding
func (s_ SampleBufferDisplayLayer) RequiresFlushToResumeDecoding() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("requiresFlushToResumeDecoding"))
	return rv
}


// An object that enqueues video sample buffers for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/sampleBufferRenderer
func (s_ SampleBufferDisplayLayer) SampleBufferRenderer() IAVSampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](s_.ID, objc.Sel("sampleBufferRenderer"))
	return rv
}


// The ability of the display layer to be used for enqueuing sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/status
func (s_ SampleBufferDisplayLayer) Status() QueuedSampleBufferRenderingStatus {
	rv := objc.Send[QueuedSampleBufferRenderingStatus](s_.ID, objc.Sel("status"))
	return rv
}


// The renderer’s timebase, which determines how the layer interprets time stamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/timebase
func (s_ SampleBufferDisplayLayer) Timebase() TimebaseRef /* not a class type */ {
	rv := objc.Send[TimebaseRef](s_.ID, objc.Sel("timebase"))
	return rv
}


// A value that indicates how the layer displays video within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/videoGravity
func (s_ SampleBufferDisplayLayer) VideoGravity() LayerVideoGravity {
	rv := objc.Send[LayerVideoGravity](s_.ID, objc.Sel("videoGravity"))
	return rv
}


// A value that indicates how the layer displays video within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/videoGravity
func (s_ SampleBufferDisplayLayer) SetVideoGravity(value LayerVideoGravity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVideoGravity:"), value)
}


// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isoutputobscuredduetoinsufficientexternalprotection
func (s_ SampleBufferDisplayLayer) IsOutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isOutputObscuredDueToInsufficientExternalProtection"))
	return rv
}


// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isoutputobscuredduetoinsufficientexternalprotection
func (s_ SampleBufferDisplayLayer) SetIsOutputObscuredDueToInsufficientExternalProtection(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsOutputObscuredDueToInsufficientExternalProtection:"), value)
}


// A Boolean value that indicates whether the first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isreadyfordisplay
func (s_ SampleBufferDisplayLayer) IsReadyForDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isReadyForDisplay"))
	return rv
}


// A Boolean value that indicates whether the first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isreadyfordisplay
func (s_ SampleBufferDisplayLayer) SetIsReadyForDisplay(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsReadyForDisplay:"), value)
}


// The key for the corresponding error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayerfailedtodecodenotificationerrorkey
func (s_ SampleBufferDisplayLayer) AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey"))
	return rv
}







