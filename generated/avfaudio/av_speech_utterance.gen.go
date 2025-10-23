// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpeechUtterance] class.
var (
	SpeechUtteranceClass     _SpeechUtteranceClass
	SpeechUtteranceClassOnce sync.Once
)

func getSpeechUtteranceClass() _SpeechUtteranceClass {
	SpeechUtteranceClassOnce.Do(func() {
		SpeechUtteranceClass = _SpeechUtteranceClass{objc.GetClass("AVSpeechUtterance")}
	})
	return SpeechUtteranceClass
}

type _SpeechUtteranceClass struct {
	class objc.Class
}

// An interface definition for the [SpeechUtterance] class.
type ISpeechUtterance interface {
	objectivec.IObject
	// properties:
	PitchMultiplier() float32 /* primitive/slice/pointer. */
	SetPitchMultiplier(value float32 /* primitive/slice/pointer. */)
	PreUtteranceDelay() foundation.TimeInterval /* not a class type */
	SetPreUtteranceDelay(value foundation.TimeInterval /* not a class type */)
	Voice() IAVSpeechSynthesisVoice
	SetVoice(value IAVSpeechSynthesisVoice)
	Volume() float32 /* primitive/slice/pointer. */
	SetVolume(value float32 /* primitive/slice/pointer. */)
	AVSpeechSynthesisIPANotationAttribute() string /* primitive/slice/pointer. */
	AttributedSpeechString() objc.IObject /* cross-framework: AttributedString */
	SetAttributedSpeechString(value objc.IObject /* cross-framework: AttributedString */)
	PostUtteranceDelay() unsafe.Pointer
	SetPostUtteranceDelay(value unsafe.Pointer)
	PrefersAssistiveTechnologySettings() bool /* primitive/slice/pointer. */
	SetPrefersAssistiveTechnologySettings(value bool /* primitive/slice/pointer. */)
	Rate() float32 /* primitive/slice/pointer. */
	SetRate(value float32 /* primitive/slice/pointer. */)
	SpeechString() string /* primitive/slice/pointer. */
	SetSpeechString(value string /* primitive/slice/pointer. */)
	AVSpeechUtteranceDefaultSpeechRate() float32 /* primitive/slice/pointer. */
	AVSpeechUtteranceMaximumSpeechRate() float32 /* primitive/slice/pointer. */
	AVSpeechUtteranceMinimumSpeechRate() float32 /* primitive/slice/pointer. */
	// methods:
}

// An object that encapsulates the text for speech synthesis and parameters that affect the speech.
//
// An is the basic unit of speech synthesis. To synthesize speech, create an instance with text you want a speech synthesizer to speak. Optionally, change the , , , , , or parameters for the utterance. Pass the utterance to an instance of to begin speech, or enqueue the utterance to speak later if the synthesizer is already speaking. Split a body of text into multiple utterances if you want to apply different speech parameters. For example, you can emphasize a sentence by increasing the pitch and decreasing the rate of that utterance relative to others, or you can introduce pauses between sentences by putting each into an utterance with a leading or trailing delay. Set and use the to receive notifications when the synthesizer starts or finishes speaking an utterance. Create an utterance for each meaningful unit in a body of text if you want to receive notifications as its speech progresses.


// An object that encapsulates the text for speech synthesis and parameters that affect the speech.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance
type SpeechUtterance struct {
	objectivec.Object
}

// SpeechUtteranceFrom constructs a [SpeechUtterance] from an unsafe.Pointer.
//
// An object that encapsulates the text for speech synthesis and parameters that affect the speech.
func SpeechUtteranceFrom(ptr unsafe.Pointer) SpeechUtterance {
	return SpeechUtterance{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SpeechUtteranceClass) Alloc() SpeechUtterance {
	rv := objc.Send[SpeechUtterance](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SpeechUtteranceClass) New() SpeechUtterance {
	rv := objc.Send[SpeechUtterance](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpeechUtterance) Init() SpeechUtterance {
	rv := objc.Send[SpeechUtterance](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpeechUtterance) Autorelease() SpeechUtterance {
	rv := objc.Send[SpeechUtterance](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeechUtterance creates a new SpeechUtterance instance.
func NewSpeechUtterance() SpeechUtterance {
	return getSpeechUtteranceClass().New()
}



// The baseline pitch the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/pitchMultiplier
func (s_ SpeechUtterance) PitchMultiplier() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("pitchMultiplier"))
	return rv
}


// The baseline pitch the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/pitchMultiplier
func (s_ SpeechUtterance) SetPitchMultiplier(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPitchMultiplier:"), value)
}


// The amount of time the speech synthesizer pauses before speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/preUtteranceDelay
func (s_ SpeechUtterance) PreUtteranceDelay() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](s_.ID, objc.Sel("preUtteranceDelay"))
	return rv
}


// The amount of time the speech synthesizer pauses before speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/preUtteranceDelay
func (s_ SpeechUtterance) SetPreUtteranceDelay(value foundation.TimeInterval /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreUtteranceDelay:"), value)
}


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/voice
func (s_ SpeechUtterance) Voice() IAVSpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](s_.ID, objc.Sel("voice"))
	return rv
}


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/voice
func (s_ SpeechUtterance) SetVoice(value IAVSpeechSynthesisVoice) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoice:"), value)
}


// The volume the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/volume
func (s_ SpeechUtterance) Volume() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}


// The volume the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/volume
func (s_ SpeechUtterance) SetVolume(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}


// A string that contains International Phonetic Alphabet (IPA) symbols the speech synthesizer uses to control pronunciation of certain words or phrases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisipanotationattribute
func (s_ SpeechUtterance) AVSpeechSynthesisIPANotationAttribute() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](s_.ID, objc.Sel("AVSpeechSynthesisIPANotationAttribute"))
	return rv
}


// An attributed string that contains the text for speech synthesis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/attributedspeechstring
func (s_ SpeechUtterance) AttributedSpeechString() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[AttributedString](s_.ID, objc.Sel("attributedSpeechString"))
	return rv
}


// An attributed string that contains the text for speech synthesis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/attributedspeechstring
func (s_ SpeechUtterance) SetAttributedSpeechString(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAttributedSpeechString:"), value)
}


// The amount of time the speech synthesizer pauses after speaking an utterance before handling the next utterance in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/postutterancedelay
func (s_ SpeechUtterance) PostUtteranceDelay() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("postUtteranceDelay"))
	return rv
}


// The amount of time the speech synthesizer pauses after speaking an utterance before handling the next utterance in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/postutterancedelay
func (s_ SpeechUtterance) SetPostUtteranceDelay(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPostUtteranceDelay:"), value)
}


// A Boolean that specifies whether assistive technology settings take precedence over the property values of this utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/prefersassistivetechnologysettings
func (s_ SpeechUtterance) PrefersAssistiveTechnologySettings() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("prefersAssistiveTechnologySettings"))
	return rv
}


// A Boolean that specifies whether assistive technology settings take precedence over the property values of this utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/prefersassistivetechnologysettings
func (s_ SpeechUtterance) SetPrefersAssistiveTechnologySettings(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPrefersAssistiveTechnologySettings:"), value)
}


// The rate the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/rate
func (s_ SpeechUtterance) Rate() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("rate"))
	return rv
}


// The rate the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/rate
func (s_ SpeechUtterance) SetRate(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:"), value)
}


// A string that contains the text for speech synthesis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/speechstring
func (s_ SpeechUtterance) SpeechString() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](s_.ID, objc.Sel("speechString"))
	return rv
}


// A string that contains the text for speech synthesis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/speechstring
func (s_ SpeechUtterance) SetSpeechString(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpeechString:"), objc.String(value))
}


// The default rate the speech synthesizer uses when speaking an utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterancedefaultspeechrate
func (s_ SpeechUtterance) AVSpeechUtteranceDefaultSpeechRate() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("AVSpeechUtteranceDefaultSpeechRate"))
	return rv
}


// The maximum rate the speech synthesizer uses when speaking an utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterancemaximumspeechrate
func (s_ SpeechUtterance) AVSpeechUtteranceMaximumSpeechRate() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("AVSpeechUtteranceMaximumSpeechRate"))
	return rv
}


// The minimum rate the speech synthesizer uses when speaking an utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutteranceminimumspeechrate
func (s_ SpeechUtterance) AVSpeechUtteranceMinimumSpeechRate() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("AVSpeechUtteranceMinimumSpeechRate"))
	return rv
}




