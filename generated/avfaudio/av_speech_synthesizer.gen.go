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
	SpeakUtterance(utterance unsafe.Pointer)
}

// An object that produces synthesized speech from text utterances and enables monitoring or controlling of ongoing speech.
//
// To speak some text, create an instance that contains the text and pass it to on a speech synthesizer instance. You can optionally also retrieve an and set it on the utterance’s property to have the speech synthesizer use that voice when speaking the utterance’s text. The speech synthesizer maintains a queue of utterances that it speaks. If the synthesizer isn’t speaking, calling begins speaking that utterance either immediately or after pausing for its , if necessary. If the synthesizer is speaking, the synthesizer adds utterances to a queue and speaks them in the order it receives them. After speech begins, you can use the synthesizer object to pause or stop speech. After pausing, you can resume the speech from its paused point or stop the speech entirely and remove all remaining utterances in the queue. You can monitor the speech synthesizer by examining its and properties, or by setting a delegate that conforms to . The delegate receives significant events as they occur during speech synthesis. An also controls the route where the speech plays. For more information, see Directing speech output.
//
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


// Adds the utterance you specify to the speech synthesizer’s queue.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speak(_:)
func (s_ SpeechSynthesizer) SpeakUtterance(utterance unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("speakUtterance:"), utterance)
}

// A Boolean value that indicates whether a speech synthesizer is in a paused state.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/ispaused
func (s_ SpeechSynthesizer) IsPaused() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isPaused"))
	return rv
}


// SetIsPaused sets the value of the isPaused property.
// A Boolean value that indicates whether a speech synthesizer is in a paused state.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/ispaused
func (s_ SpeechSynthesizer) SetIsPaused(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsPaused:"), value)
}

// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s_ SpeechSynthesizer) Voice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("voice"))
	return rv
}


// SetVoice sets the value of the voice property.
// The voice the speech synthesizer uses when speaking the utterance.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s_ SpeechSynthesizer) SetVoice(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoice:"), value)
}

// A Boolean value that specifies whether to send synthesized speech to an active call.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/mixtotelephonyuplink
func (s_ SpeechSynthesizer) MixToTelephonyUplink() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("mixToTelephonyUplink"))
	return rv
}


// SetMixToTelephonyUplink sets the value of the mixToTelephonyUplink property.
// A Boolean value that specifies whether to send synthesized speech to an active call.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/mixtotelephonyuplink
func (s_ SpeechSynthesizer) SetMixToTelephonyUplink(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMixToTelephonyUplink:"), value)
}

// An array of audio session channels to route generated speech.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/outputchannels
func (s_ SpeechSynthesizer) OutputChannels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("outputChannels"))
	return rv
}


// SetOutputChannels sets the value of the outputChannels property.
// An array of audio session channels to route generated speech.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/outputchannels
func (s_ SpeechSynthesizer) SetOutputChannels(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setOutputChannels:"), value)
}

// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/isspeaking
func (s_ SpeechSynthesizer) IsSpeaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSpeaking"))
	return rv
}


// SetIsSpeaking sets the value of the isSpeaking property.
// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/isspeaking
func (s_ SpeechSynthesizer) SetIsSpeaking(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSpeaking:"), value)
}

// The amount of time the speech synthesizer pauses before speaking the utterance.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/preutterancedelay
func (s_ SpeechSynthesizer) PreUtteranceDelay() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("preUtteranceDelay"))
	return rv
}


// SetPreUtteranceDelay sets the value of the preUtteranceDelay property.
// The amount of time the speech synthesizer pauses before speaking the utterance.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/preutterancedelay
func (s_ SpeechSynthesizer) SetPreUtteranceDelay(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreUtteranceDelay:"), value)
}

// A Boolean value that specifies whether the app manages the audio session.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/usesapplicationaudiosession
func (s_ SpeechSynthesizer) UsesApplicationAudioSession() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("usesApplicationAudioSession"))
	return rv
}


// SetUsesApplicationAudioSession sets the value of the usesApplicationAudioSession property.
// A Boolean value that specifies whether the app manages the audio session.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/usesapplicationaudiosession
func (s_ SpeechSynthesizer) SetUsesApplicationAudioSession(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUsesApplicationAudioSession:"), value)
}

// The delegate object for the speech synthesizer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/delegate
func (s_ SpeechSynthesizer) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object for the speech synthesizer.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/delegate
func (s_ SpeechSynthesizer) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isSpeaking
func (s_ SpeechSynthesizer) Speaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("speaking"))
	return rv
}



