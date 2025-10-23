// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpeechSynthesizer] class.
var (
	SpeechSynthesizerClass     _SpeechSynthesizerClass
	SpeechSynthesizerClassOnce sync.Once
)

func getSpeechSynthesizerClass() _SpeechSynthesizerClass {
	SpeechSynthesizerClassOnce.Do(func() {
		SpeechSynthesizerClass = _SpeechSynthesizerClass{objc.GetClass("AVSpeechSynthesizer")}
	})
	return SpeechSynthesizerClass
}

type _SpeechSynthesizerClass struct {
	class objc.Class
}

// An interface definition for the [SpeechSynthesizer] class.
type ISpeechSynthesizer interface {
	objectivec.IObject
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Paused() bool
	Speaking() bool
	MixToTelephonyUplink() bool
	SetMixToTelephonyUplink(value bool)
	OutputChannels() []unsafe.Pointer
	SetOutputChannels(value []unsafe.Pointer)
	UsesApplicationAudioSession() bool
	SetUsesApplicationAudioSession(value bool)
	IsPaused() bool
	SetIsPaused(value bool)
	IsSpeaking() bool
	SetIsSpeaking(value bool)
	PreUtteranceDelay() unsafe.Pointer
	SetPreUtteranceDelay(value unsafe.Pointer)
	Voice() IAVSpeechSynthesisVoice
	SetVoice(value IAVSpeechSynthesisVoice)
	ContinueSpeaking() bool
	PauseSpeakingAtBoundary(boundary AVSpeechBoundary) bool
	SpeakUtterance(utterance IAVSpeechUtterance)
	StopSpeakingAtBoundary(boundary AVSpeechBoundary) bool
	WriteUtteranceToBufferCallback(utterance IAVSpeechUtterance, bufferCallback unsafe.Pointer)
	WriteUtteranceToBufferCallbackToMarkerCallback(utterance IAVSpeechUtterance, bufferCallback unsafe.Pointer, markerCallback unsafe.Pointer)
}

// An object that produces synthesized speech from text utterances and enables monitoring or controlling of ongoing speech.
//
// To speak some text, create an instance that contains the text and pass it to on a speech synthesizer instance. You can optionally also retrieve an and set it on the utterance’s property to have the speech synthesizer use that voice when speaking the utterance’s text. The speech synthesizer maintains a queue of utterances that it speaks. If the synthesizer isn’t speaking, calling begins speaking that utterance either immediately or after pausing for its , if necessary. If the synthesizer is speaking, the synthesizer adds utterances to a queue and speaks them in the order it receives them. After speech begins, you can use the synthesizer object to pause or stop speech. After pausing, you can resume the speech from its paused point or stop the speech entirely and remove all remaining utterances in the queue. You can monitor the speech synthesizer by examining its and properties, or by setting a delegate that conforms to . The delegate receives significant events as they occur during speech synthesis. An also controls the route where the speech plays. For more information, see Directing speech output.


// An object that produces synthesized speech from text utterances and enables monitoring or controlling of ongoing speech.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer
type SpeechSynthesizer struct {
	objectivec.Object
}

// SpeechSynthesizerFrom constructs a [SpeechSynthesizer] from an unsafe.Pointer.
//
// An object that produces synthesized speech from text utterances and enables monitoring or controlling of ongoing speech.
func SpeechSynthesizerFrom(ptr unsafe.Pointer) SpeechSynthesizer {
	return SpeechSynthesizer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesizerClass) Alloc() SpeechSynthesizer {
	rv := objc.Send[SpeechSynthesizer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SpeechSynthesizerClass) New() SpeechSynthesizer {
	rv := objc.Send[SpeechSynthesizer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpeechSynthesizer) Init() SpeechSynthesizer {
	rv := objc.Send[SpeechSynthesizer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpeechSynthesizer) Autorelease() SpeechSynthesizer {
	rv := objc.Send[SpeechSynthesizer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeechSynthesizer creates a new SpeechSynthesizer instance.
func NewSpeechSynthesizer() SpeechSynthesizer {
	return getSpeechSynthesizerClass().New()
}



// Prompts the user to authorize your app to use personal voices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/requestPersonalVoiceAuthorization(completionHandler:)
func (sc _SpeechSynthesizerClass) RequestPersonalVoiceAuthorizationWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("requestPersonalVoiceAuthorizationWithCompletionHandler:"), handler)
}


// Your app’s authorization to use personal voices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/personalVoiceAuthorizationStatus-swift.type.property
func (sc _SpeechSynthesizerClass) PersonalVoiceAuthorizationStatus() AVSpeechSynthesisPersonalVoiceAuthorizationStatus {
	rv := objc.Send[AVSpeechSynthesisPersonalVoiceAuthorizationStatus](objc.ID(sc.class), objc.Sel("personalVoiceAuthorizationStatus"))
	return rv
}

// Resumes speech from its paused point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/continueSpeaking()
func (s_ SpeechSynthesizer) ContinueSpeaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("continueSpeaking"))
	return rv
}


// Pauses speech at the boundary you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/pauseSpeaking(at:)
func (s_ SpeechSynthesizer) PauseSpeakingAtBoundary(boundary AVSpeechBoundary) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("pauseSpeakingAtBoundary:"), boundary)
	return rv
}


// Adds the utterance you specify to the speech synthesizer’s queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speak(_:)
func (s_ SpeechSynthesizer) SpeakUtterance(utterance IAVSpeechUtterance) {
	objc.Send[objc.ID](s_.ID, objc.Sel("speakUtterance:"), utterance)
}


// Stops speech at the boundary you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/stopSpeaking(at:)
func (s_ SpeechSynthesizer) StopSpeakingAtBoundary(boundary AVSpeechBoundary) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("stopSpeakingAtBoundary:"), boundary)
	return rv
}


// Generates speech for the utterance and invokes the callback with the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/write(_:toBufferCallback:)
func (s_ SpeechSynthesizer) WriteUtteranceToBufferCallback(utterance IAVSpeechUtterance, bufferCallback unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("writeUtterance:toBufferCallback:"), utterance, bufferCallback)
}


// Generates audio buffers and associated metadata for storage or further speech synthesis processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/write(_:toBufferCallback:toMarkerCallback:)
func (s_ SpeechSynthesizer) WriteUtteranceToBufferCallbackToMarkerCallback(utterance IAVSpeechUtterance, bufferCallback unsafe.Pointer, markerCallback unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("writeUtterance:toBufferCallback:toMarkerCallback:"), utterance, bufferCallback, markerCallback)
}


// The delegate object for the speech synthesizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/delegate
func (s_ SpeechSynthesizer) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate object for the speech synthesizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/delegate
func (s_ SpeechSynthesizer) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether a speech synthesizer is in a paused state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isPaused
func (s_ SpeechSynthesizer) Paused() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("paused"))
	return rv
}


// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isSpeaking
func (s_ SpeechSynthesizer) Speaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("speaking"))
	return rv
}


// A Boolean value that specifies whether to send synthesized speech to an active call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/mixToTelephonyUplink
func (s_ SpeechSynthesizer) MixToTelephonyUplink() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("mixToTelephonyUplink"))
	return rv
}


// A Boolean value that specifies whether to send synthesized speech to an active call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/mixToTelephonyUplink
func (s_ SpeechSynthesizer) SetMixToTelephonyUplink(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMixToTelephonyUplink:"), value)
}


// An array of audio session channels to route generated speech.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/outputChannels
func (s_ SpeechSynthesizer) OutputChannels() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](s_.ID, objc.Sel("outputChannels"))
	return rv
}


// An array of audio session channels to route generated speech.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/outputChannels
func (s_ SpeechSynthesizer) SetOutputChannels(value []unsafe.Pointer) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setOutputChannels:"), nsArray)
}


// Your app’s authorization to use personal voices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/personalVoiceAuthorizationStatus-swift.type.property
func (s_ SpeechSynthesizer) PersonalVoiceAuthorizationStatus() AVSpeechSynthesisPersonalVoiceAuthorizationStatus {
	rv := objc.Send[AVSpeechSynthesisPersonalVoiceAuthorizationStatus](s_.ID, objc.Sel("personalVoiceAuthorizationStatus"))
	return rv
}


// A Boolean value that specifies whether the app manages the audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/usesApplicationAudioSession
func (s_ SpeechSynthesizer) UsesApplicationAudioSession() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("usesApplicationAudioSession"))
	return rv
}


// A Boolean value that specifies whether the app manages the audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/usesApplicationAudioSession
func (s_ SpeechSynthesizer) SetUsesApplicationAudioSession(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUsesApplicationAudioSession:"), value)
}


// A Boolean value that indicates whether a speech synthesizer is in a paused state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/ispaused
func (s_ SpeechSynthesizer) IsPaused() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isPaused"))
	return rv
}


// A Boolean value that indicates whether a speech synthesizer is in a paused state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/ispaused
func (s_ SpeechSynthesizer) SetIsPaused(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsPaused:"), value)
}


// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/isspeaking
func (s_ SpeechSynthesizer) IsSpeaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSpeaking"))
	return rv
}


// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/isspeaking
func (s_ SpeechSynthesizer) SetIsSpeaking(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSpeaking:"), value)
}


// The amount of time the speech synthesizer pauses before speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/preutterancedelay
func (s_ SpeechSynthesizer) PreUtteranceDelay() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("preUtteranceDelay"))
	return rv
}


// The amount of time the speech synthesizer pauses before speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/preutterancedelay
func (s_ SpeechSynthesizer) SetPreUtteranceDelay(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreUtteranceDelay:"), value)
}


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s_ SpeechSynthesizer) Voice() IAVSpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](s_.ID, objc.Sel("voice"))
	return rv
}


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s_ SpeechSynthesizer) SetVoice(value IAVSpeechSynthesisVoice) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoice:"), value)
}



