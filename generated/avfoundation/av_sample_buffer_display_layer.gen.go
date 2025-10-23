// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/quartzcore"
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
	quartzcore.ILayer
	// properties:
	ControlTimebase() CMTimebase /* foo */
	SetControlTimebase(value CMTimebase /* foo */)
	IsOutputObscuredDueToInsufficientExternalProtection() bool /* primitive/slice/pointer */
	SetIsOutputObscuredDueToInsufficientExternalProtection(value bool /* primitive/slice/pointer */)
	IsReadyForDisplay() bool /* primitive/slice/pointer */
	SetIsReadyForDisplay(value bool /* primitive/slice/pointer */)
	PreventsAutomaticBackgroundingDuringVideoPlayback() bool /* primitive/slice/pointer */
	SetPreventsAutomaticBackgroundingDuringVideoPlayback(value bool /* primitive/slice/pointer */)
	PreventsCapture() bool /* primitive/slice/pointer */
	SetPreventsCapture(value bool /* primitive/slice/pointer */)
	PreventsDisplaySleepDuringVideoPlayback() bool /* primitive/slice/pointer */
	SetPreventsDisplaySleepDuringVideoPlayback(value bool /* primitive/slice/pointer */)
	SampleBufferRenderer() IAVSampleBufferVideoRenderer
	SetSampleBufferRenderer(value IAVSampleBufferVideoRenderer)
	VideoGravity() AVLayerVideoGravity /* foo */
	SetVideoGravity(value AVLayerVideoGravity /* foo */)
	AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey() string /* primitive/slice/pointer */
	// methods:
}

// An object that displays compressed or uncompressed video frames.


// An object that displays compressed or uncompressed video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer
type SampleBufferDisplayLayer struct {
	quartzcore.Layer
}

// SampleBufferDisplayLayerFrom constructs a [SampleBufferDisplayLayer] from an unsafe.Pointer.
//
// An object that displays compressed or uncompressed video frames.
func SampleBufferDisplayLayerFrom(ptr unsafe.Pointer) SampleBufferDisplayLayer {
	return SampleBufferDisplayLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SampleBufferDisplayLayerClass) Alloc() SampleBufferDisplayLayer {
	rv := objc.Send[SampleBufferDisplayLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A timebase that determines how the layer interprets timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/controltimebase
func (s_ SampleBufferDisplayLayer) ControlTimebase() CMTimebase /* foo */ {
	rv := objc.Send[Timebase](s_.ID, objc.Sel("controlTimebase"))
	return rv
}


// A timebase that determines how the layer interprets timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/controltimebase
func (s_ SampleBufferDisplayLayer) SetControlTimebase(value CMTimebase /* foo */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setControlTimebase:"), value)
}


// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isoutputobscuredduetoinsufficientexternalprotection
func (s_ SampleBufferDisplayLayer) IsOutputObscuredDueToInsufficientExternalProtection() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("isOutputObscuredDueToInsufficientExternalProtection"))
	return rv
}


// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isoutputobscuredduetoinsufficientexternalprotection
func (s_ SampleBufferDisplayLayer) SetIsOutputObscuredDueToInsufficientExternalProtection(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsOutputObscuredDueToInsufficientExternalProtection:"), value)
}


// A Boolean value that indicates whether the first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isreadyfordisplay
func (s_ SampleBufferDisplayLayer) IsReadyForDisplay() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("isReadyForDisplay"))
	return rv
}


// A Boolean value that indicates whether the first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isreadyfordisplay
func (s_ SampleBufferDisplayLayer) SetIsReadyForDisplay(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsReadyForDisplay:"), value)
}


// A Boolean value that indicates whether video playback prevents the system from automatically backgrounding an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventsautomaticbackgroundingduringvideoplayback
func (s_ SampleBufferDisplayLayer) PreventsAutomaticBackgroundingDuringVideoPlayback() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsAutomaticBackgroundingDuringVideoPlayback"))
	return rv
}


// A Boolean value that indicates whether video playback prevents the system from automatically backgrounding an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventsautomaticbackgroundingduringvideoplayback
func (s_ SampleBufferDisplayLayer) SetPreventsAutomaticBackgroundingDuringVideoPlayback(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreventsAutomaticBackgroundingDuringVideoPlayback:"), value)
}


// A Boolean value that indicates whether the layer protects against screen capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventscapture
func (s_ SampleBufferDisplayLayer) PreventsCapture() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsCapture"))
	return rv
}


// A Boolean value that indicates whether the layer protects against screen capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventscapture
func (s_ SampleBufferDisplayLayer) SetPreventsCapture(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreventsCapture:"), value)
}


// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventsdisplaysleepduringvideoplayback
func (s_ SampleBufferDisplayLayer) PreventsDisplaySleepDuringVideoPlayback() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsDisplaySleepDuringVideoPlayback"))
	return rv
}


// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventsdisplaysleepduringvideoplayback
func (s_ SampleBufferDisplayLayer) SetPreventsDisplaySleepDuringVideoPlayback(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreventsDisplaySleepDuringVideoPlayback:"), value)
}


// An object that enqueues video sample buffers for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/samplebufferrenderer
func (s_ SampleBufferDisplayLayer) SampleBufferRenderer() IAVSampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](s_.ID, objc.Sel("sampleBufferRenderer"))
	return rv
}


// An object that enqueues video sample buffers for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/samplebufferrenderer
func (s_ SampleBufferDisplayLayer) SetSampleBufferRenderer(value IAVSampleBufferVideoRenderer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSampleBufferRenderer:"), value)
}


// A value that indicates how the layer displays video within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/videogravity
func (s_ SampleBufferDisplayLayer) VideoGravity() AVLayerVideoGravity /* foo */ {
	rv := objc.Send[LayerVideoGravity](s_.ID, objc.Sel("videoGravity"))
	return rv
}


// A value that indicates how the layer displays video within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/videogravity
func (s_ SampleBufferDisplayLayer) SetVideoGravity(value AVLayerVideoGravity /* foo */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVideoGravity:"), value)
}


// The key for the corresponding error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayerfailedtodecodenotificationerrorkey
func (s_ SampleBufferDisplayLayer) AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](s_.ID, objc.Sel("AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey"))
	return rv
}



