// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSpeechSynthesisVoice */


/* debug [class_header]: Header for AVSpeechSynthesisVoice */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpeechSynthesisVoice */
// An interface definition for the [SpeechSynthesisVoice] class.
type ISpeechSynthesisVoice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SpeechSynthesisVoice */
	// properties:
	AudioFileSettings() foundation.IDictionary
	Gender() SpeechSynthesisVoiceGender
	Identifier() objc.IObject /* cross-framework: NSString */
	Language() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	Quality() SpeechSynthesisVoiceQuality
	VoiceTraits() SpeechSynthesisVoiceTraits
	AVSpeechSynthesisVoiceIdentifierAlex() objc.IObject /* cross-framework: NSString */
	Voice() IAVSpeechSynthesisVoice
	SetVoice(value IAVSpeechSynthesisVoice)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpeechSynthesisVoice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpeechSynthesisVoice */
// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesisVoiceClass) Alloc() SpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpeechSynthesisVoice */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpeechSynthesisVoice */

// Retrieves a voice for the identifier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/init(identifier:)
func NewSpeechSynthesisVoiceWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) SpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](objc.ID(getSpeechSynthesisVoiceClass().class), objc.Sel("voiceWithIdentifier:"), identifier)
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechSynthesisVoiceWithIdentifier */


// Retrieves a voice for the BCP 47 code language code you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/init(language:)
func NewSpeechSynthesisVoiceWithLanguage(languageCode objc.IObject /* cross-framework: NSString */) SpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](objc.ID(getSpeechSynthesisVoiceClass().class), objc.Sel("voiceWithLanguage:"), languageCode)
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechSynthesisVoiceWithLanguage */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpeechSynthesisVoice */

// Returns the language and locale code for the user’s current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/currentLanguageCode()
func (sc _SpeechSynthesisVoiceClass) CurrentLanguageCode() foundation.String {
	rv := objc.Send[foundation.String](objc.ID(sc.class), objc.Sel("currentLanguageCode"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentLanguageCode) */


// Retrieves a voice for the identifier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/init(identifier:)
func (sc _SpeechSynthesisVoiceClass) VoiceWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) ISpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](objc.ID(sc.class), objc.Sel("voiceWithIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VoiceWithIdentifier) */


// Retrieves a voice for the BCP 47 code language code you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/init(language:)
func (sc _SpeechSynthesisVoiceClass) VoiceWithLanguage(languageCode objc.IObject /* cross-framework: NSString */) ISpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](objc.ID(sc.class), objc.Sel("voiceWithLanguage:"), languageCode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VoiceWithLanguage) */


// Retrieves all available voices on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/speechVoices()
func (sc _SpeechSynthesisVoiceClass) SpeechVoices() []SpeechSynthesisVoice {
	rv := objc.Send[[]SpeechSynthesisVoice](objc.ID(sc.class), objc.Sel("speechVoices"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SpeechVoices) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpeechSynthesisVoice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpeechSynthesisVoice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpeechSynthesisVoice */

// A dictionary that contains audio file settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/audioFileSettings
func (s_ SpeechSynthesisVoice) AudioFileSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("audioFileSettings"))
	return rv
}/* debug [instance_properties/getter]: audioFileSettings */


// The gender for a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/gender
func (s_ SpeechSynthesisVoice) Gender() SpeechSynthesisVoiceGender {
	rv := objc.Send[SpeechSynthesisVoiceGender](s_.ID, objc.Sel("gender"))
	return rv
}/* debug [instance_properties/getter]: gender */


// The unique identifier of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/identifier
func (s_ SpeechSynthesisVoice) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A BCP 47 code that contains the voice’s language and locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/language
func (s_ SpeechSynthesisVoice) Language() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("language"))
	return rv
}/* debug [instance_properties/getter]: language */


// The name of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/name
func (s_ SpeechSynthesisVoice) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The speech quality of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/quality
func (s_ SpeechSynthesisVoice) Quality() SpeechSynthesisVoiceQuality {
	rv := objc.Send[SpeechSynthesisVoiceQuality](s_.ID, objc.Sel("quality"))
	return rv
}/* debug [instance_properties/getter]: quality */


// The traits of a voice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/voiceTraits
func (s_ SpeechSynthesisVoice) VoiceTraits() SpeechSynthesisVoiceTraits {
	rv := objc.Send[SpeechSynthesisVoiceTraits](s_.ID, objc.Sel("voiceTraits"))
	return rv
}/* debug [instance_properties/getter]: voiceTraits */


// The voice that the system identifies as Alex.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoiceidentifieralex
func (s_ SpeechSynthesisVoice) AVSpeechSynthesisVoiceIdentifierAlex() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("AVSpeechSynthesisVoiceIdentifierAlex"))
	return rv
}/* debug [instance_properties/getter]: AVSpeechSynthesisVoiceIdentifierAlex */


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s_ SpeechSynthesisVoice) Voice() IAVSpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](s_.ID, objc.Sel("voice"))
	return rv
}/* debug [instance_properties/getter]: voice */


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s_ SpeechSynthesisVoice) SetVoice(value IAVSpeechSynthesisVoice) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoice:"), value)
}/* debug [instance_properties/setter]: voice */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSpeechSynthesisVoice */


