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
}

// A distinct voice for use in speech synthesis.
//
// The primary factors that distinguish a voice in speech synthesis are language, locale, and quality. Create an instance of to select a voice that’s appropriate for the text and the language, and set it as the value of the property on an instance. The voice may optionally reflect a local variant of the language, such as Australian or South African English. For a complete list of supported languages, see .
//
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


// The gender for a voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/gender
func (s_ SpeechSynthesisVoice) Gender() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("gender"))
	return rv
}

// A BCP 47 code that contains the voice’s language and locale.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/language
func (s_ SpeechSynthesisVoice) Language() string {
	rv := objc.Send[string](s_.ID, objc.Sel("language"))
	return rv
}



