// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
)

// The class instance for the [SFSpeechAudioBufferRecognitionRequest] class.
var (
	SFSpeechAudioBufferRecognitionRequestClass     _SFSpeechAudioBufferRecognitionRequestClass
	SFSpeechAudioBufferRecognitionRequestClassOnce sync.Once
)

func getSFSpeechAudioBufferRecognitionRequestClass() _SFSpeechAudioBufferRecognitionRequestClass {
	SFSpeechAudioBufferRecognitionRequestClassOnce.Do(func() {
		SFSpeechAudioBufferRecognitionRequestClass = _SFSpeechAudioBufferRecognitionRequestClass{objc.GetClass("SFSpeechAudioBufferRecognitionRequest")}
	})
	return SFSpeechAudioBufferRecognitionRequestClass
}

type _SFSpeechAudioBufferRecognitionRequestClass struct {
	class objc.Class
}

// An interface definition for the [SFSpeechAudioBufferRecognitionRequest] class.
type ISFSpeechAudioBufferRecognitionRequest interface {
	ISFSpeechRecognitionRequest
	// properties:
	NativeAudioFormat() objc.IObject /* cross-framework: AudioFormat */
	SetNativeAudioFormat(value objc.IObject /* cross-framework: AudioFormat */)
	// methods:
}

// A request to recognize speech from captured audio content, such as audio from the device’s microphone.
//
// Use an object to perform speech recognition on live audio, or on a set of existing audio buffers. For example, use this request object to route audio from a device’s microphone to the speech recognizer. The request object contains no audio initially. As you capture audio, call or to add audio samples to the request object. The speech recognizer continuously analyzes the audio you appended, stopping only when you call the method. You must call explicitly to stop the speech recognition process. For a complete example of how to use audio buffers with speech recognition, see .


// A request to recognize speech from captured audio content, such as audio from the device’s microphone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest
type SFSpeechAudioBufferRecognitionRequest struct {
	SFSpeechRecognitionRequest
}

// SFSpeechAudioBufferRecognitionRequestFrom constructs a [SFSpeechAudioBufferRecognitionRequest] from an unsafe.Pointer.
//
// A request to recognize speech from captured audio content, such as audio from the device’s microphone.
func SFSpeechAudioBufferRecognitionRequestFrom(ptr unsafe.Pointer) SFSpeechAudioBufferRecognitionRequest {
	return SFSpeechAudioBufferRecognitionRequest{
		SFSpeechRecognitionRequest: SFSpeechRecognitionRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSpeechAudioBufferRecognitionRequestClass) Alloc() SFSpeechAudioBufferRecognitionRequest {
	rv := objc.Send[SFSpeechAudioBufferRecognitionRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSpeechAudioBufferRecognitionRequestClass) New() SFSpeechAudioBufferRecognitionRequest {
	rv := objc.Send[SFSpeechAudioBufferRecognitionRequest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSpeechAudioBufferRecognitionRequest) Init() SFSpeechAudioBufferRecognitionRequest {
	rv := objc.Send[SFSpeechAudioBufferRecognitionRequest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSpeechAudioBufferRecognitionRequest) Autorelease() SFSpeechAudioBufferRecognitionRequest {
	rv := objc.Send[SFSpeechAudioBufferRecognitionRequest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechAudioBufferRecognitionRequest creates a new SFSpeechAudioBufferRecognitionRequest instance.
func NewSFSpeechAudioBufferRecognitionRequest() SFSpeechAudioBufferRecognitionRequest {
	return getSFSpeechAudioBufferRecognitionRequestClass().New()
}



// The preferred audio format for optimal speech recognition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechaudiobufferrecognitionrequest/nativeaudioformat
func (s_ SFSpeechAudioBufferRecognitionRequest) NativeAudioFormat() objc.IObject /* cross-framework: AudioFormat */ {
	rv := objc.Send[avfaudio.AudioFormat](s_.ID, objc.Sel("nativeAudioFormat"))
	return rv
}


// The preferred audio format for optimal speech recognition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechaudiobufferrecognitionrequest/nativeaudioformat
func (s_ SFSpeechAudioBufferRecognitionRequest) SetNativeAudioFormat(value objc.IObject /* cross-framework: AudioFormat */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNativeAudioFormat:"), value)
}



