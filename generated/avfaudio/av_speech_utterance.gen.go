// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An object that encapsulates the text for speech synthesis and parameters that affect the speech.
//
// An is the basic unit of speech synthesis. To synthesize speech, create an instance with text you want a speech synthesizer to speak. Optionally, change the , , , , , or parameters for the utterance. Pass the utterance to an instance of to begin speech, or enqueue the utterance to speak later if the synthesizer is already speaking. Split a body of text into multiple utterances if you want to apply different speech parameters. For example, you can emphasize a sentence by increasing the pitch and decreasing the rate of that utterance relative to others, or you can introduce pauses between sentences by putting each into an utterance with a leading or trailing delay. Set and use the to receive notifications when the synthesizer starts or finishes speaking an utterance. Create an utterance for each meaningful unit in a body of text if you want to receive notifications as its speech progresses.
//
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


// Creates an utterance with the attributed text string that you specify for the speech synthesizer to speak.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechUtteranceWithAttributedString:
func (sc _SpeechUtteranceClass) SpeechUtteranceWithAttributedString(string_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("speechUtteranceWithAttributedString:"), string_)
	return rv
}

// An attributed string that contains the text for speech synthesis.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/attributedSpeechString
func (s_ SpeechUtterance) AttributedSpeechString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("attributedSpeechString"))
	return rv
}

// The baseline pitch the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/pitchMultiplier
func (s_ SpeechUtterance) PitchMultiplier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("pitchMultiplier"))
	return rv
}


// SetPitchMultiplier sets the value of the pitchMultiplier property.
// The baseline pitch the speech synthesizer uses when speaking the utterance.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/pitchMultiplier
func (s_ SpeechUtterance) SetPitchMultiplier(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPitchMultiplier:"), value)
}
// The amount of time the speech synthesizer pauses after speaking an utterance before handling the next utterance in the queue.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/postUtteranceDelay
func (s_ SpeechUtterance) PostUtteranceDelay() TimeInterval {
	rv := objc.Send[TimeInterval](s_.ID, objc.Sel("postUtteranceDelay"))
	return rv
}


// SetPostUtteranceDelay sets the value of the postUtteranceDelay property.
// The amount of time the speech synthesizer pauses after speaking an utterance before handling the next utterance in the queue.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/postUtteranceDelay
func (s_ SpeechUtterance) SetPostUtteranceDelay(value TimeInterval) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPostUtteranceDelay:"), value)
}
// The amount of time the speech synthesizer pauses before speaking the utterance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/preUtteranceDelay
func (s_ SpeechUtterance) PreUtteranceDelay() TimeInterval {
	rv := objc.Send[TimeInterval](s_.ID, objc.Sel("preUtteranceDelay"))
	return rv
}


// SetPreUtteranceDelay sets the value of the preUtteranceDelay property.
// The amount of time the speech synthesizer pauses before speaking the utterance.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/preUtteranceDelay
func (s_ SpeechUtterance) SetPreUtteranceDelay(value TimeInterval) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreUtteranceDelay:"), value)
}
// The rate the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/rate
func (s_ SpeechUtterance) Rate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
// The rate the speech synthesizer uses when speaking the utterance.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/rate
func (s_ SpeechUtterance) SetRate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:"), value)
}
// A string that contains the text for speech synthesis.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechString
func (s_ SpeechUtterance) SpeechString() string {
	rv := objc.Send[string](s_.ID, objc.Sel("speechString"))
	return rv
}

// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/voice
func (s_ SpeechUtterance) Voice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("voice"))
	return rv
}


// SetVoice sets the value of the voice property.
// The voice the speech synthesizer uses when speaking the utterance.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/voice
func (s_ SpeechUtterance) SetVoice(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoice:"), value)
}
// The volume the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/volume
func (s_ SpeechUtterance) Volume() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("volume"))
	return rv
}


// SetVolume sets the value of the volume property.
// The volume the speech synthesizer uses when speaking the utterance.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/volume
func (s_ SpeechUtterance) SetVolume(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}


