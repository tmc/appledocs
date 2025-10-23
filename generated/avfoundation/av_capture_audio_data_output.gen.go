// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CaptureAudioDataOutput] class.
var (
	CaptureAudioDataOutputClass     _CaptureAudioDataOutputClass
	CaptureAudioDataOutputClassOnce sync.Once
)

func getCaptureAudioDataOutputClass() _CaptureAudioDataOutputClass {
	CaptureAudioDataOutputClassOnce.Do(func() {
		CaptureAudioDataOutputClass = _CaptureAudioDataOutputClass{objc.GetClass("AVCaptureAudioDataOutput")}
	})
	return CaptureAudioDataOutputClass
}

type _CaptureAudioDataOutputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureAudioDataOutput] class.
type ICaptureAudioDataOutput interface {
	ICaptureOutput
	// properties:
	AudioSettings() objc.IObject /* cross-framework: NSString */
	SetAudioSettings(value objc.IObject /* cross-framework: NSString */)
	SampleBufferCallbackQueue() unsafe.Pointer
	SetSampleBufferCallbackQueue(value unsafe.Pointer)
	SampleBufferDelegate() CaptureAudioDataOutputSampleBufferDelegate /* not a class type */
	SetSampleBufferDelegate(value CaptureAudioDataOutputSampleBufferDelegate /* not a class type */)
	SpatialAudioChannelLayoutTag() unsafe.Pointer
	SetSpatialAudioChannelLayoutTag(value unsafe.Pointer)
	// methods:
}

// A capture output that records audio and provides access to audio sample buffers as they are recorded.


// A capture output that records audio and provides access to audio sample buffers as they are recorded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput
type CaptureAudioDataOutput struct {
	CaptureOutput
}

// CaptureAudioDataOutputFrom constructs a [CaptureAudioDataOutput] from an unsafe.Pointer.
//
// A capture output that records audio and provides access to audio sample buffers as they are recorded.
func CaptureAudioDataOutputFrom(ptr unsafe.Pointer) CaptureAudioDataOutput {
	return CaptureAudioDataOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureAudioDataOutputClass) Alloc() CaptureAudioDataOutput {
	rv := objc.Send[CaptureAudioDataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureAudioDataOutputClass) New() CaptureAudioDataOutput {
	rv := objc.Send[CaptureAudioDataOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureAudioDataOutput) Init() CaptureAudioDataOutput {
	rv := objc.Send[CaptureAudioDataOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureAudioDataOutput) Autorelease() CaptureAudioDataOutput {
	rv := objc.Send[CaptureAudioDataOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureAudioDataOutput creates a new CaptureAudioDataOutput instance.
func NewCaptureAudioDataOutput() CaptureAudioDataOutput {
	return getCaptureAudioDataOutputClass().New()
}



// The settings used to decode or re-encode audio before it’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/audiosettings
func (c_ CaptureAudioDataOutput) AudioSettings() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("audioSettings"))
	return rv
}


// The settings used to decode or re-encode audio before it’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/audiosettings
func (c_ CaptureAudioDataOutput) SetAudioSettings(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSettings:"), value)
}


// The queue on which delegate callbacks are invoked
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/samplebuffercallbackqueue
func (c_ CaptureAudioDataOutput) SampleBufferCallbackQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sampleBufferCallbackQueue"))
	return rv
}


// The queue on which delegate callbacks are invoked
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/samplebuffercallbackqueue
func (c_ CaptureAudioDataOutput) SetSampleBufferCallbackQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBufferCallbackQueue:"), value)
}


// The capture object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/samplebufferdelegate
func (c_ CaptureAudioDataOutput) SampleBufferDelegate() CaptureAudioDataOutputSampleBufferDelegate /* not a class type */ {
	rv := objc.Send[CaptureAudioDataOutputSampleBufferDelegate](c_.ID, objc.Sel("sampleBufferDelegate"))
	return rv
}


// The capture object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/samplebufferdelegate
func (c_ CaptureAudioDataOutput) SetSampleBufferDelegate(value CaptureAudioDataOutputSampleBufferDelegate /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBufferDelegate:"), value)
}


// The audio channel layout tag of the audio sample buffers produced by the audio data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/spatialaudiochannellayouttag
func (c_ CaptureAudioDataOutput) SpatialAudioChannelLayoutTag() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("spatialAudioChannelLayoutTag"))
	return rv
}


// The audio channel layout tag of the audio sample buffers produced by the audio data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/spatialaudiochannellayouttag
func (c_ CaptureAudioDataOutput) SetSpatialAudioChannelLayoutTag(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpatialAudioChannelLayoutTag:"), value)
}



