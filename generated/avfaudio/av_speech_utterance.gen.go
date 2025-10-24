// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSpeechUtterance */


/* debug [class_header]: Header for AVSpeechUtterance */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpeechUtterance */
// An interface definition for the [SpeechUtterance] class.
type ISpeechUtterance interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SpeechUtterance */
	// properties:
	AttributedSpeechString() foundation.AttributedString
	PitchMultiplier() float32
	SetPitchMultiplier(value float32)
	PostUtteranceDelay() float64
	SetPostUtteranceDelay(value float64)
	PrefersAssistiveTechnologySettings() bool
	SetPrefersAssistiveTechnologySettings(value bool)
	PreUtteranceDelay() float64
	SetPreUtteranceDelay(value float64)
	Rate() float32
	SetRate(value float32)
	SpeechString() objc.IObject /* cross-framework: NSString */
	Voice() IAVSpeechSynthesisVoice
	SetVoice(value IAVSpeechSynthesisVoice)
	Volume() float32
	SetVolume(value float32)
	AVSpeechSynthesisIPANotationAttribute() objc.IObject /* cross-framework: NSString */
	AVSpeechUtteranceDefaultSpeechRate() float32
	AVSpeechUtteranceMaximumSpeechRate() float32
	AVSpeechUtteranceMinimumSpeechRate() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpeechUtterance */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpeechUtterance */
// Alloc allocates a new instance without initialization.
func (sc _SpeechUtteranceClass) Alloc() SpeechUtterance {
	rv := objc.Send[SpeechUtterance](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpeechUtterance */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpeechUtterance */

// Creates an utterance with the attributed text string that you specify for the speech synthesizer to speak.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/init(attributedString:)
func NewSpeechUtteranceWithAttributedString(string_ foundation.AttributedString) SpeechUtterance {
	instance := getSpeechUtteranceClass().Alloc()
	rv := objc.Send[SpeechUtterance](instance.ID, objc.Sel("initWithAttributedString:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechUtteranceWithAttributedString */


// Creates a speech utterance with an Speech Synthesis Markup Language (SSML) string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/init(ssmlRepresentation:)
func NewSpeechUtteranceWithSSMLRepresentation(string_ objc.IObject /* cross-framework: NSString */) SpeechUtterance {
	instance := getSpeechUtteranceClass().Alloc()
	rv := objc.Send[SpeechUtterance](instance.ID, objc.Sel("initWithSSMLRepresentation:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechUtteranceWithSSMLRepresentation */


// Creates an utterance with the text string that you specify for the speech synthesizer to speak.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/init(string:)
func NewSpeechUtteranceWithString(string_ objc.IObject /* cross-framework: NSString */) SpeechUtterance {
	instance := getSpeechUtteranceClass().Alloc()
	rv := objc.Send[SpeechUtterance](instance.ID, objc.Sel("initWithString:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechUtteranceWithString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpeechUtterance */

// Creates an utterance with the attributed text string that you specify for the speech synthesizer to speak.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechUtteranceWithAttributedString:
func (sc _SpeechUtteranceClass) SpeechUtteranceWithAttributedString(string_ foundation.AttributedString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("speechUtteranceWithAttributedString:"), string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SpeechUtteranceWithAttributedString) */


// Returns a new speech utterance with an Speech Synthesis Markup Language (SSML) string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechUtteranceWithSSMLRepresentation:
func (sc _SpeechUtteranceClass) SpeechUtteranceWithSSMLRepresentation(string_ objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("speechUtteranceWithSSMLRepresentation:"), string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SpeechUtteranceWithSSMLRepresentation) */


// Creates an utterance with the text string that you specify for the speech synthesizer to speak.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechUtteranceWithString:
func (sc _SpeechUtteranceClass) SpeechUtteranceWithString(string_ objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("speechUtteranceWithString:"), string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SpeechUtteranceWithString) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpeechUtterance */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpeechUtterance */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpeechUtterance */

// An attributed string that contains the text for speech synthesis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/attributedSpeechString
func (s_ SpeechUtterance) AttributedSpeechString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](s_.ID, objc.Sel("attributedSpeechString"))
	return rv
}/* debug [instance_properties/getter]: attributedSpeechString */


// The baseline pitch the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/pitchMultiplier
func (s_ SpeechUtterance) PitchMultiplier() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("pitchMultiplier"))
	return rv
}/* debug [instance_properties/getter]: pitchMultiplier */


// The baseline pitch the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/pitchMultiplier
func (s_ SpeechUtterance) SetPitchMultiplier(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPitchMultiplier:"), value)
}/* debug [instance_properties/setter]: pitchMultiplier */


// The amount of time the speech synthesizer pauses after speaking an utterance before handling the next utterance in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/postUtteranceDelay
func (s_ SpeechUtterance) PostUtteranceDelay() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("postUtteranceDelay"))
	return rv
}/* debug [instance_properties/getter]: postUtteranceDelay */


// The amount of time the speech synthesizer pauses after speaking an utterance before handling the next utterance in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/postUtteranceDelay
func (s_ SpeechUtterance) SetPostUtteranceDelay(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPostUtteranceDelay:"), value)
}/* debug [instance_properties/setter]: postUtteranceDelay */


// A Boolean that specifies whether assistive technology settings take precedence over the property values of this utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/prefersAssistiveTechnologySettings
func (s_ SpeechUtterance) PrefersAssistiveTechnologySettings() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("prefersAssistiveTechnologySettings"))
	return rv
}/* debug [instance_properties/getter]: prefersAssistiveTechnologySettings */


// A Boolean that specifies whether assistive technology settings take precedence over the property values of this utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/prefersAssistiveTechnologySettings
func (s_ SpeechUtterance) SetPrefersAssistiveTechnologySettings(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPrefersAssistiveTechnologySettings:"), value)
}/* debug [instance_properties/setter]: prefersAssistiveTechnologySettings */


// The amount of time the speech synthesizer pauses before speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/preUtteranceDelay
func (s_ SpeechUtterance) PreUtteranceDelay() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("preUtteranceDelay"))
	return rv
}/* debug [instance_properties/getter]: preUtteranceDelay */


// The amount of time the speech synthesizer pauses before speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/preUtteranceDelay
func (s_ SpeechUtterance) SetPreUtteranceDelay(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreUtteranceDelay:"), value)
}/* debug [instance_properties/setter]: preUtteranceDelay */


// The rate the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/rate
func (s_ SpeechUtterance) Rate() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// The rate the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/rate
func (s_ SpeechUtterance) SetRate(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// A string that contains the text for speech synthesis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechString
func (s_ SpeechUtterance) SpeechString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("speechString"))
	return rv
}/* debug [instance_properties/getter]: speechString */


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/voice
func (s_ SpeechUtterance) Voice() IAVSpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](s_.ID, objc.Sel("voice"))
	return rv
}/* debug [instance_properties/getter]: voice */


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/voice
func (s_ SpeechUtterance) SetVoice(value IAVSpeechSynthesisVoice) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoice:"), value)
}/* debug [instance_properties/setter]: voice */


// The volume the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/volume
func (s_ SpeechUtterance) Volume() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}/* debug [instance_properties/getter]: volume */


// The volume the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/volume
func (s_ SpeechUtterance) SetVolume(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}/* debug [instance_properties/setter]: volume */


// A string that contains International Phonetic Alphabet (IPA) symbols the speech synthesizer uses to control pronunciation of certain words or phrases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisipanotationattribute
func (s_ SpeechUtterance) AVSpeechSynthesisIPANotationAttribute() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("AVSpeechSynthesisIPANotationAttribute"))
	return rv
}/* debug [instance_properties/getter]: AVSpeechSynthesisIPANotationAttribute */


// The default rate the speech synthesizer uses when speaking an utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterancedefaultspeechrate
func (s_ SpeechUtterance) AVSpeechUtteranceDefaultSpeechRate() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("AVSpeechUtteranceDefaultSpeechRate"))
	return rv
}/* debug [instance_properties/getter]: AVSpeechUtteranceDefaultSpeechRate */


// The maximum rate the speech synthesizer uses when speaking an utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterancemaximumspeechrate
func (s_ SpeechUtterance) AVSpeechUtteranceMaximumSpeechRate() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("AVSpeechUtteranceMaximumSpeechRate"))
	return rv
}/* debug [instance_properties/getter]: AVSpeechUtteranceMaximumSpeechRate */


// The minimum rate the speech synthesizer uses when speaking an utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutteranceminimumspeechrate
func (s_ SpeechUtterance) AVSpeechUtteranceMinimumSpeechRate() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("AVSpeechUtteranceMinimumSpeechRate"))
	return rv
}/* debug [instance_properties/getter]: AVSpeechUtteranceMinimumSpeechRate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSpeechUtterance */


