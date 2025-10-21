// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpeechSynthesisProviderVoice] class.
var (
	SpeechSynthesisProviderVoiceClass     _SpeechSynthesisProviderVoiceClass
	SpeechSynthesisProviderVoiceClassOnce sync.Once
)

func getSpeechSynthesisProviderVoiceClass() _SpeechSynthesisProviderVoiceClass {
	SpeechSynthesisProviderVoiceClassOnce.Do(func() {
		SpeechSynthesisProviderVoiceClass = _SpeechSynthesisProviderVoiceClass{objc.GetClass("AVSpeechSynthesisProviderVoice")}
	})
	return SpeechSynthesisProviderVoiceClass
}

type _SpeechSynthesisProviderVoiceClass struct {
	class objc.Class
}

// An interface definition for the [SpeechSynthesisProviderVoice] class.
type ISpeechSynthesisProviderVoice interface {
	objectivec.IObject
}

// An object that represents a voice that an audio unit provides to its host.
//
// This is a voice that an provides to the system, distinct from . Use to access the underlying in the voice quality .
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice
type SpeechSynthesisProviderVoice struct {
	objectivec.Object
}

// SpeechSynthesisProviderVoiceFrom constructs a [SpeechSynthesisProviderVoice] from an unsafe.Pointer.
//
// An object that represents a voice that an audio unit provides to its host.
func SpeechSynthesisProviderVoiceFrom(ptr unsafe.Pointer) SpeechSynthesisProviderVoice {
	return SpeechSynthesisProviderVoice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesisProviderVoiceClass) Alloc() SpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SpeechSynthesisProviderVoiceClass) New() SpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpeechSynthesisProviderVoice) Init() SpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpeechSynthesisProviderVoice) Autorelease() SpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeechSynthesisProviderVoice creates a new SpeechSynthesisProviderVoice instance.
func NewSpeechSynthesisProviderVoice() SpeechSynthesisProviderVoice {
	return getSpeechSynthesisProviderVoiceClass().New()
}


// Updates the voices your app provides to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/updateSpeechVoices()
func (sc _SpeechSynthesisProviderVoiceClass) UpdateSpeechVoices() {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("updateSpeechVoices"))
}

// A list of voices the audio unit provides to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovideraudiounit/speechvoices
func (s_ SpeechSynthesisProviderVoice) SpeechVoices() AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](s_.ID, objc.Sel("speechVoices"))
	return rv
}


// SetSpeechVoices sets the value of the speechVoices property.
// A list of voices the audio unit provides to the system.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovideraudiounit/speechvoices
func (s_ SpeechSynthesisProviderVoice) SetSpeechVoices(value IAVSpeechSynthesisProviderVoice) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpeechVoices:"), value)
}

// The age of the voice, in years.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/age
func (s_ SpeechSynthesisProviderVoice) Age() int {
	rv := objc.Send[int](s_.ID, objc.Sel("age"))
	return rv
}


// SetAge sets the value of the age property.
// The age of the voice, in years.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/age
func (s_ SpeechSynthesisProviderVoice) SetAge(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAge:"), value)
}

// The gender of the voice.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/gender
func (s_ SpeechSynthesisProviderVoice) Gender() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("gender"))
	return rv
}


// SetGender sets the value of the gender property.
// The gender of the voice.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/gender
func (s_ SpeechSynthesisProviderVoice) SetGender(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setGender:"), value)
}

// The unique identifier for the voice.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/identifier
func (s_ SpeechSynthesisProviderVoice) Identifier() appkit.string {
	rv := objc.Send[appkit.string](s_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The unique identifier for the voice.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/identifier
func (s_ SpeechSynthesisProviderVoice) SetIdentifier(value appkit.string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIdentifier:"), value)
}

// The localized name of the voice.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/name
func (s_ SpeechSynthesisProviderVoice) Name() appkit.string {
	rv := objc.Send[appkit.string](s_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The localized name of the voice.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/name
func (s_ SpeechSynthesisProviderVoice) SetName(value appkit.string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setName:"), value)
}

// A list of BCP 47 codes that identify the languages the synthesizer uses.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/primarylanguages
func (s_ SpeechSynthesisProviderVoice) PrimaryLanguages() appkit.string {
	rv := objc.Send[appkit.string](s_.ID, objc.Sel("primaryLanguages"))
	return rv
}


// SetPrimaryLanguages sets the value of the primaryLanguages property.
// A list of BCP 47 codes that identify the languages the synthesizer uses.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/primarylanguages
func (s_ SpeechSynthesisProviderVoice) SetPrimaryLanguages(value appkit.string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPrimaryLanguages:"), value)
}

// A list of BCP 47 codes that identify the languages a voice supports.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/supportedlanguages
func (s_ SpeechSynthesisProviderVoice) SupportedLanguages() appkit.string {
	rv := objc.Send[appkit.string](s_.ID, objc.Sel("supportedLanguages"))
	return rv
}


// SetSupportedLanguages sets the value of the supportedLanguages property.
// A list of BCP 47 codes that identify the languages a voice supports.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/supportedlanguages
func (s_ SpeechSynthesisProviderVoice) SetSupportedLanguages(value appkit.string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSupportedLanguages:"), value)
}

// The version of the voice.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/version
func (s_ SpeechSynthesisProviderVoice) Version() appkit.string {
	rv := objc.Send[appkit.string](s_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
// The version of the voice.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/version
func (s_ SpeechSynthesisProviderVoice) SetVersion(value appkit.string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVersion:"), value)
}

// The size of the voice package on disk, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/voicesize
func (s_ SpeechSynthesisProviderVoice) VoiceSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("voiceSize"))
	return rv
}


// SetVoiceSize sets the value of the voiceSize property.
// The size of the voice package on disk, in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovidervoice/voicesize
func (s_ SpeechSynthesisProviderVoice) SetVoiceSize(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoiceSize:"), value)
}



