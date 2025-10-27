// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	AudioSettings() foundation.IDictionary
	SetAudioSettings(value foundation.IDictionary)
	SampleBufferCallbackQueue() objectivec.IObject
	SampleBufferDelegate() unsafe.Pointer
	SpatialAudioChannelLayoutTag() objectivec.IObject
	SetSpatialAudioChannelLayoutTag(value objectivec.IObject)


	

	// methods:
	RecommendedAudioSettingsForAssetWriterWithOutputFileType(outputFileType FileType) foundation.IDictionary
	SetSampleBufferDelegateQueue(sampleBufferDelegate unsafe.Pointer, sampleBufferCallbackQueue objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureAudioDataOutputClass) Alloc() CaptureAudioDataOutput {
	rv := objc.Send[CaptureAudioDataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





















// Specifies the recommended settings for use with an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/recommendedAudioSettingsForAssetWriter(writingTo:)
func (c_ CaptureAudioDataOutput) RecommendedAudioSettingsForAssetWriterWithOutputFileType(outputFileType FileType) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("recommendedAudioSettingsForAssetWriterWithOutputFileType:"), outputFileType)
	return rv
}


// Sets the delegate that will accept captured buffers and the dispatch queue on which the delegate will be called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/setSampleBufferDelegate(_:queue:)
func (c_ CaptureAudioDataOutput) SetSampleBufferDelegateQueue(sampleBufferDelegate unsafe.Pointer, sampleBufferCallbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBufferDelegate:queue:"), sampleBufferDelegate, sampleBufferCallbackQueue)
}







// The settings used to decode or re-encode audio before it’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/audioSettings
func (c_ CaptureAudioDataOutput) AudioSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("audioSettings"))
	return rv
}


// The settings used to decode or re-encode audio before it’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/audioSettings
func (c_ CaptureAudioDataOutput) SetAudioSettings(value foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSettings:"), value)
}


// The queue on which delegate callbacks are invoked
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/sampleBufferCallbackQueue
func (c_ CaptureAudioDataOutput) SampleBufferCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sampleBufferCallbackQueue"))
	return rv
}


// The capture object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/sampleBufferDelegate
func (c_ CaptureAudioDataOutput) SampleBufferDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sampleBufferDelegate"))
	return rv
}


// The audio channel layout tag of the audio sample buffers produced by the audio data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/spatialAudioChannelLayoutTag
func (c_ CaptureAudioDataOutput) SpatialAudioChannelLayoutTag() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("spatialAudioChannelLayoutTag"))
	return rv
}


// The audio channel layout tag of the audio sample buffers produced by the audio data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/spatialAudioChannelLayoutTag
func (c_ CaptureAudioDataOutput) SetSpatialAudioChannelLayoutTag(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpatialAudioChannelLayoutTag:"), value)
}







