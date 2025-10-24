// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class SFSpeechAudioBufferRecognitionRequest */

/* debug [class_header]: Header for SFSpeechAudioBufferRecognitionRequest */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFSpeechAudioBufferRecognitionRequest */
// An interface definition for the [SFSpeechAudioBufferRecognitionRequest] class.
type ISFSpeechAudioBufferRecognitionRequest interface {
	ISFSpeechRecognitionRequest

	/* debug [class_interface_properties]: Properties for SFSpeechAudioBufferRecognitionRequest */
	// properties:
	NativeAudioFormat() avfaudio.AudioFormat
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFSpeechAudioBufferRecognitionRequest */
	// methods:
	AppendAudioPCMBuffer(audioPCMBuffer avfaudio.AudioPCMBuffer)
	AppendAudioSampleBuffer(sampleBuffer SampleBufferRef /* not a class type */)
	EndAudio()
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFSpeechAudioBufferRecognitionRequest */
// Alloc allocates a new instance without initialization.
func (sc _SFSpeechAudioBufferRecognitionRequestClass) Alloc() SFSpeechAudioBufferRecognitionRequest {
	rv := objc.Send[SFSpeechAudioBufferRecognitionRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFSpeechAudioBufferRecognitionRequest */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFSpeechAudioBufferRecognitionRequest */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFSpeechAudioBufferRecognitionRequest */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFSpeechAudioBufferRecognitionRequest */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFSpeechAudioBufferRecognitionRequest */

// Appends audio in the PCM format to the end of the recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest/append(_:)
func (s_ SFSpeechAudioBufferRecognitionRequest) AppendAudioPCMBuffer(audioPCMBuffer avfaudio.AudioPCMBuffer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("appendAudioPCMBuffer:"), audioPCMBuffer)
} /* debug [instance_methods/method]: AppendAudioPCMBuffer */

// Appends audio to the end of the recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest/appendAudioSampleBuffer(_:)
func (s_ SFSpeechAudioBufferRecognitionRequest) AppendAudioSampleBuffer(sampleBuffer SampleBufferRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("appendAudioSampleBuffer:"), sampleBuffer)
} /* debug [instance_methods/method]: AppendAudioSampleBuffer */

// Marks the end of audio input for the recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest/endAudio()
func (s_ SFSpeechAudioBufferRecognitionRequest) EndAudio() {
	objc.Send[objc.ID](s_.ID, objc.Sel("endAudio"))
} /* debug [instance_methods/method]: EndAudio */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFSpeechAudioBufferRecognitionRequest */

// The preferred audio format for optimal speech recognition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest/nativeAudioFormat
func (s_ SFSpeechAudioBufferRecognitionRequest) NativeAudioFormat() avfaudio.AudioFormat {
	rv := objc.Send[avfaudio.AudioFormat](s_.ID, objc.Sel("nativeAudioFormat"))
	return rv
} /* debug [instance_properties/getter]: nativeAudioFormat */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFSpeechAudioBufferRecognitionRequest */
