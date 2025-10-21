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
}

// An object that displays compressed or uncompressed video frames.
//
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
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/controltimebase
func (s_ SampleBufferDisplayLayer) ControlTimebase() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("controlTimebase"))
	return rv
}


// SetControlTimebase sets the value of the controlTimebase property.
// A timebase that determines how the layer interprets timestamps.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/controltimebase
func (s_ SampleBufferDisplayLayer) SetControlTimebase(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setControlTimebase:"), value)
}

// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isoutputobscuredduetoinsufficientexternalprotection
func (s_ SampleBufferDisplayLayer) IsOutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isOutputObscuredDueToInsufficientExternalProtection"))
	return rv
}


// SetIsOutputObscuredDueToInsufficientExternalProtection sets the value of the isOutputObscuredDueToInsufficientExternalProtection property.
// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isoutputobscuredduetoinsufficientexternalprotection
func (s_ SampleBufferDisplayLayer) SetIsOutputObscuredDueToInsufficientExternalProtection(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsOutputObscuredDueToInsufficientExternalProtection:"), value)
}

// A Boolean value that indicates whether the first video frame is ready for display.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isreadyfordisplay
func (s_ SampleBufferDisplayLayer) IsReadyForDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isReadyForDisplay"))
	return rv
}


// SetIsReadyForDisplay sets the value of the isReadyForDisplay property.
// A Boolean value that indicates whether the first video frame is ready for display.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/isreadyfordisplay
func (s_ SampleBufferDisplayLayer) SetIsReadyForDisplay(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsReadyForDisplay:"), value)
}

// A Boolean value that indicates whether video playback prevents the system from automatically backgrounding an app.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventsautomaticbackgroundingduringvideoplayback
func (s_ SampleBufferDisplayLayer) PreventsAutomaticBackgroundingDuringVideoPlayback() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsAutomaticBackgroundingDuringVideoPlayback"))
	return rv
}


// SetPreventsAutomaticBackgroundingDuringVideoPlayback sets the value of the preventsAutomaticBackgroundingDuringVideoPlayback property.
// A Boolean value that indicates whether video playback prevents the system from automatically backgrounding an app.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventsautomaticbackgroundingduringvideoplayback
func (s_ SampleBufferDisplayLayer) SetPreventsAutomaticBackgroundingDuringVideoPlayback(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreventsAutomaticBackgroundingDuringVideoPlayback:"), value)
}

// A Boolean value that indicates whether the layer protects against screen capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventscapture
func (s_ SampleBufferDisplayLayer) PreventsCapture() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsCapture"))
	return rv
}


// SetPreventsCapture sets the value of the preventsCapture property.
// A Boolean value that indicates whether the layer protects against screen capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventscapture
func (s_ SampleBufferDisplayLayer) SetPreventsCapture(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreventsCapture:"), value)
}

// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventsdisplaysleepduringvideoplayback
func (s_ SampleBufferDisplayLayer) PreventsDisplaySleepDuringVideoPlayback() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsDisplaySleepDuringVideoPlayback"))
	return rv
}


// SetPreventsDisplaySleepDuringVideoPlayback sets the value of the preventsDisplaySleepDuringVideoPlayback property.
// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/preventsdisplaysleepduringvideoplayback
func (s_ SampleBufferDisplayLayer) SetPreventsDisplaySleepDuringVideoPlayback(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreventsDisplaySleepDuringVideoPlayback:"), value)
}

// An object that enqueues video sample buffers for rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/samplebufferrenderer
func (s_ SampleBufferDisplayLayer) SampleBufferRenderer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sampleBufferRenderer"))
	return rv
}


// SetSampleBufferRenderer sets the value of the sampleBufferRenderer property.
// An object that enqueues video sample buffers for rendering.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/samplebufferrenderer
func (s_ SampleBufferDisplayLayer) SetSampleBufferRenderer(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSampleBufferRenderer:"), value)
}

// A value that indicates how the layer displays video within its bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/videogravity
func (s_ SampleBufferDisplayLayer) VideoGravity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("videoGravity"))
	return rv
}


// SetVideoGravity sets the value of the videoGravity property.
// A value that indicates how the layer displays video within its bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/videogravity
func (s_ SampleBufferDisplayLayer) SetVideoGravity(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVideoGravity:"), value)
}

// The key for the corresponding error.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayerfailedtodecodenotificationerrorkey
func (s_ SampleBufferDisplayLayer) AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey() string {
	rv := objc.Send[string](s_.ID, objc.Sel("AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey"))
	return rv
}



