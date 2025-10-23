// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpeechSynthesisVoice] class.
var (
	SpeechSynthesisVoiceClass     _SpeechSynthesisVoiceClass
	SpeechSynthesisVoiceClassOnce sync.Once
)

func getSpeechSynthesisVoiceClass() _SpeechSynthesisVoiceClass {
	SpeechSynthesisVoiceClassOnce.Do(func() {
		SpeechSynthesisVoiceClass = _SpeechSynthesisVoiceClass{objc.GetClass("AVSpeechSynthesisVoice")}
	})
	return SpeechSynthesisVoiceClass
}

type _SpeechSynthesisVoiceClass struct {
	class objc.Class
}

// An interface definition for the [SpeechSynthesisVoice] class.
type ISpeechSynthesisVoice interface {
	objectivec.IObject
	AudioFileSettings() string
	SetAudioFileSettings(value string)
	Gender() unsafe.Pointer
	SetGender(value unsafe.Pointer)
	Identifier() string
	SetIdentifier(value string)
	Language() string
	SetLanguage(value string)
	Name() string
	SetName(value string)
	Quality() unsafe.Pointer
	SetQuality(value unsafe.Pointer)
	VoiceTraits() unsafe.Pointer
	SetVoiceTraits(value unsafe.Pointer)
	AVSpeechSynthesisVoiceIdentifierAlex() string
	Voice() IAVSpeechSynthesisVoice
	SetVoice(value IAVSpeechSynthesisVoice)
}

// A distinct voice for use in speech synthesis.
//
// The primary factors that distinguish a voice in speech synthesis are language, locale, and quality. Create an instance of to select a voice that’s appropriate for the text and the language, and set it as the value of the property on an instance. The voice may optionally reflect a local variant of the language, such as Australian or South African English. For a complete list of supported languages, see .


// A distinct voice for use in speech synthesis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice
type SpeechSynthesisVoice struct {
	objectivec.Object
}

// SpeechSynthesisVoiceFrom constructs a [SpeechSynthesisVoice] from an unsafe.Pointer.
//
// A distinct voice for use in speech synthesis.
func SpeechSynthesisVoiceFrom(ptr unsafe.Pointer) SpeechSynthesisVoice {
	return SpeechSynthesisVoice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesisVoiceClass) Alloc() SpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SpeechSynthesisVoiceClass) New() SpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpeechSynthesisVoice) Init() SpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpeechSynthesisVoice) Autorelease() SpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeechSynthesisVoice creates a new SpeechSynthesisVoice instance.
func NewSpeechSynthesisVoice() SpeechSynthesisVoice {
	return getSpeechSynthesisVoiceClass().New()
}



// A dictionary that contains audio file settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/audiofilesettings
func (s_ SpeechSynthesisVoice) AudioFileSettings() string {
	rv := objc.Send[string](s_.ID, objc.Sel("audioFileSettings"))
	return rv
}


// A dictionary that contains audio file settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/audiofilesettings
func (s_ SpeechSynthesisVoice) SetAudioFileSettings(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAudioFileSettings:"), objc.String(value))
}


// The gender for a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/gender
func (s_ SpeechSynthesisVoice) Gender() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("gender"))
	return rv
}


// The gender for a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/gender
func (s_ SpeechSynthesisVoice) SetGender(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setGender:"), value)
}


// The unique identifier of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/identifier
func (s_ SpeechSynthesisVoice) Identifier() string {
	rv := objc.Send[string](s_.ID, objc.Sel("identifier"))
	return rv
}


// The unique identifier of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/identifier
func (s_ SpeechSynthesisVoice) SetIdentifier(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// A BCP 47 code that contains the voice’s language and locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/language
func (s_ SpeechSynthesisVoice) Language() string {
	rv := objc.Send[string](s_.ID, objc.Sel("language"))
	return rv
}


// A BCP 47 code that contains the voice’s language and locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/language
func (s_ SpeechSynthesisVoice) SetLanguage(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLanguage:"), objc.String(value))
}


// The name of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/name
func (s_ SpeechSynthesisVoice) Name() string {
	rv := objc.Send[string](s_.ID, objc.Sel("name"))
	return rv
}


// The name of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/name
func (s_ SpeechSynthesisVoice) SetName(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setName:"), objc.String(value))
}


// The speech quality of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/quality
func (s_ SpeechSynthesisVoice) Quality() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("quality"))
	return rv
}


// The speech quality of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/quality
func (s_ SpeechSynthesisVoice) SetQuality(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setQuality:"), value)
}


// The traits of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/voicetraits
func (s_ SpeechSynthesisVoice) VoiceTraits() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("voiceTraits"))
	return rv
}


// The traits of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoice/voicetraits
func (s_ SpeechSynthesisVoice) SetVoiceTraits(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoiceTraits:"), value)
}


// The voice that the system identifies as Alex.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoiceidentifieralex
func (s_ SpeechSynthesisVoice) AVSpeechSynthesisVoiceIdentifierAlex() string {
	rv := objc.Send[string](s_.ID, objc.Sel("AVSpeechSynthesisVoiceIdentifierAlex"))
	return rv
}


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s_ SpeechSynthesisVoice) Voice() IAVSpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](s_.ID, objc.Sel("voice"))
	return rv
}


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s_ SpeechSynthesisVoice) SetVoice(value IAVSpeechSynthesisVoice) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoice:"), value)
}



