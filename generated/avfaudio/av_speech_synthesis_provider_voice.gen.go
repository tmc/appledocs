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
	

	// properties:
	Age() int
	SetAge(value int)
	Gender() SpeechSynthesisVoiceGender
	SetGender(value SpeechSynthesisVoiceGender)
	Identifier() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	PrimaryLanguages() []string
	SupportedLanguages() []string
	Version() objc.IObject /* cross-framework: NSString */
	SetVersion(value objc.IObject /* cross-framework: NSString */)
	VoiceSize() int64
	SetVoiceSize(value int64)
	SpeechVoices() IAVSpeechSynthesisProviderVoice
	SetSpeechVoices(value IAVSpeechSynthesisProviderVoice)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesisProviderVoiceClass) Alloc() SpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that represents a voice that an audio unit provides to its host.
//
// This is a voice that an provides to the system, distinct from . Use to access the underlying in the voice quality .


// An object that represents a voice that an audio unit provides to its host.
//
// [Full Topic]
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






// Creates a voice with a name, an identifier, and language information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/init(name:identifier:primaryLanguages:supportedLanguages:)
func NewSpeechSynthesisProviderVoiceWithNameIdentifierPrimaryLanguagesSupportedLanguages(name objc.IObject /* cross-framework: NSString */, identifier objc.IObject /* cross-framework: NSString */, primaryLanguages []string, supportedLanguages []string) SpeechSynthesisProviderVoice {
	instance := getSpeechSynthesisProviderVoiceClass().Alloc()
	rv := objc.Send[SpeechSynthesisProviderVoice](instance.ID, objc.Sel("initWithName:identifier:primaryLanguages:supportedLanguages:"), name, identifier, primaryLanguages, supportedLanguages)
	rv.Autorelease()
	return rv
}







// Updates the voices your app provides to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/updateSpeechVoices()
func (sc _SpeechSynthesisProviderVoiceClass) UpdateSpeechVoices() {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("updateSpeechVoices"))
}

















// The age of the voice, in years.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/age
func (s_ SpeechSynthesisProviderVoice) Age() int {
	rv := objc.Send[int](s_.ID, objc.Sel("age"))
	return rv
}


// The age of the voice, in years.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/age
func (s_ SpeechSynthesisProviderVoice) SetAge(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAge:"), value)
}


// The gender of the voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/gender
func (s_ SpeechSynthesisProviderVoice) Gender() SpeechSynthesisVoiceGender {
	rv := objc.Send[SpeechSynthesisVoiceGender](s_.ID, objc.Sel("gender"))
	return rv
}


// The gender of the voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/gender
func (s_ SpeechSynthesisProviderVoice) SetGender(value SpeechSynthesisVoiceGender) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setGender:"), value)
}


// The unique identifier for the voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/identifier
func (s_ SpeechSynthesisProviderVoice) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("identifier"))
	return rv
}


// The localized name of the voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/name
func (s_ SpeechSynthesisProviderVoice) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("name"))
	return rv
}


// A list of BCP 47 codes that identify the languages the synthesizer uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/primaryLanguages
func (s_ SpeechSynthesisProviderVoice) PrimaryLanguages() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("primaryLanguages"))
	return rv
}


// A list of BCP 47 codes that identify the languages a voice supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/supportedLanguages
func (s_ SpeechSynthesisProviderVoice) SupportedLanguages() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("supportedLanguages"))
	return rv
}


// The version of the voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/version
func (s_ SpeechSynthesisProviderVoice) Version() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("version"))
	return rv
}


// The version of the voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/version
func (s_ SpeechSynthesisProviderVoice) SetVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVersion:"), value)
}


// The size of the voice package on disk, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/voiceSize
func (s_ SpeechSynthesisProviderVoice) VoiceSize() int64 {
	rv := objc.Send[int64](s_.ID, objc.Sel("voiceSize"))
	return rv
}


// The size of the voice package on disk, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/voiceSize
func (s_ SpeechSynthesisProviderVoice) SetVoiceSize(value int64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoiceSize:"), value)
}


// A list of voices the audio unit provides to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovideraudiounit/speechvoices
func (s_ SpeechSynthesisProviderVoice) SpeechVoices() IAVSpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](s_.ID, objc.Sel("speechVoices"))
	return rv
}


// A list of voices the audio unit provides to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovideraudiounit/speechvoices
func (s_ SpeechSynthesisProviderVoice) SetSpeechVoices(value IAVSpeechSynthesisProviderVoice) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpeechVoices:"), value)
}







