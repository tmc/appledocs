// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpeechSynthesizer] class.
var (
	SpeechSynthesizerClass     _SpeechSynthesizerClass
	SpeechSynthesizerClassOnce sync.Once
)

func getSpeechSynthesizerClass() _SpeechSynthesizerClass {
	SpeechSynthesizerClassOnce.Do(func() {
		SpeechSynthesizerClass = _SpeechSynthesizerClass{objc.GetClass("NSSpeechSynthesizer")}
	})
	return SpeechSynthesizerClass
}

type _SpeechSynthesizerClass struct {
	class objc.Class
}

// An interface definition for the [SpeechSynthesizer] class.
type ISpeechSynthesizer interface {
	objectivec.IObject
	AddSpeechDictionary(speechDictionary unsafe.Pointer)
	ContinueSpeaking()
	ObjectForPropertyError(property ISpeechPropertyKey, outError unsafe.Pointer) objc.ID
	PauseSpeakingAtBoundary(boundary ISpeechBoundary)
	PhonemesFromText(text string) foundation.String
	SetObjectForPropertyError(object objectivec.IObject, property ISpeechPropertyKey, outError unsafe.Pointer) bool
	SetVoice(voice ISpeechSynthesizerVoiceName) bool
	StartSpeakingString(string_ string) bool
	StartSpeakingStringToURL(string_ string, url foundation.IURL) bool
	StopSpeaking()
	StopSpeakingAtBoundary(boundary ISpeechBoundary)
	Voice() SpeechSynthesizerVoiceName
}

// The Cocoa interface to speech synthesis in macOS.
//
// Speech synthesis, also called text-to-speech (TTS), parses text and converts it into audible speech. It offers a concurrent feedback mode that can be used in concert with or in place of traditional visual and aural notifications. For example, your application can use a speech synthesizer object to “pronounce” the text of important alert dialogs. Synthesized speech has several advantages. It can provide urgent information to users without forcing them to shift attention from their current task. And because speech doesn’t rely on visual elements for meaning, it is a crucial technology for users with vision or attention disabilities. In addition, synthesized speech can help save system resources. Because sound samples can take up large amounts of room on disk, using text in place of sampled sound is extremely efficient, and so a multimedia application might use an object to provide a narration of a QuickTime movie instead of including sampled-sound data on a movie track. When you create an instance using the default initializer ( ), the class uses the selected in System Preferences > Speech. Alternatively, you can select a specific voice for an instance by initializing it with . To begin synthesis, send either or to the instance. The former generates speech through the system’s default sound output device; the latter saves the generated speech to a file. If you wish to be notified when the current speech concludes, set the property and implement the delegate method . Speech synthesis is just one of the macOS speech technologies. The speech recognizer technology allows applications to “listen to” text spoken in U.S. English; the class is the Cocoa interface to this technology. Both technologies provide benefits for all users, and are particularly useful to those users who have difficulties seeing the screen or using the mouse and keyboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer
type SpeechSynthesizer struct {
	objectivec.Object
}

// SpeechSynthesizerFrom constructs a [SpeechSynthesizer] from an unsafe.Pointer.
//
// The Cocoa interface to speech synthesis in macOS.
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




// Initializes the receiver with a voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/init(voice:)
func NewSpeechSynthesizerWithVoice(voice ISpeechSynthesizerVoiceName) SpeechSynthesizer {
	instance := getSpeechSynthesizerClass().Alloc()
	rv := objc.Send[SpeechSynthesizer](instance.ID, objc.Sel("initWithVoice:"), voice)
	rv.Autorelease()
	return rv
}


// Provides the attribute dictionary of a voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/attributes(forVoice:)
func (sc _SpeechSynthesizerClass) AttributesForVoice(voice ISpeechSynthesizerVoiceName) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("attributesForVoice:"), voice)
	return rv
}

// Provides the identifiers of the voices available on the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/availableVoices
func (sc _SpeechSynthesizerClass) AvailableVoices() []string {
	rv := objc.Send[[]string](objc.ID(sc.class), objc.Sel("availableVoices"))
	return rv
}
// Provides the identifier of the default voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/defaultVoice
func (sc _SpeechSynthesizerClass) DefaultVoice() SpeechSynthesizerVoiceName {
	rv := objc.Send[SpeechSynthesizerVoiceName](objc.ID(sc.class), objc.Sel("defaultVoice"))
	return rv
}
// A Boolean value indicating whether any application is currently speaking through the sound output device.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/isAnyApplicationSpeaking
func (sc _SpeechSynthesizerClass) AnyApplicationSpeaking() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("anyApplicationSpeaking"))
	return rv
}
// Registers the given speech dictionary with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/addSpeechDictionary(_:)
func (s_ SpeechSynthesizer) AddSpeechDictionary(speechDictionary unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addSpeechDictionary:"), speechDictionary)
}

// Resumes synthesis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/continueSpeaking()
func (s_ SpeechSynthesizer) ContinueSpeaking() {
	objc.Send[objc.ID](s_.ID, objc.Sel("continueSpeaking"))
}

// Provides the value of a receiver’s property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/object(forProperty:)
func (s_ SpeechSynthesizer) ObjectForPropertyError(property ISpeechPropertyKey, outError unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("objectForProperty:error:"), property, outError)
	return rv
}

// Pauses synthesis in progress at a given boundary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/pauseSpeaking(at:)
func (s_ SpeechSynthesizer) PauseSpeakingAtBoundary(boundary ISpeechBoundary) {
	objc.Send[objc.ID](s_.ID, objc.Sel("pauseSpeakingAtBoundary:"), boundary)
}

// Provides the phoneme symbols generated by the given text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/phonemes(from:)
func (s_ SpeechSynthesizer) PhonemesFromText(text string) foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("phonemesFromText:"), objc.String(text))
	return rv
}

// Specifies the value of a receiver’s property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/setObject(_:forProperty:)
func (s_ SpeechSynthesizer) SetObjectForPropertyError(object objectivec.IObject, property ISpeechPropertyKey, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("setObject:forProperty:error:"), object, property, outError)
	return rv
}

// Sets the receiver’s current voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/setVoice(_:)
func (s_ SpeechSynthesizer) SetVoice(voice ISpeechSynthesizerVoiceName) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("setVoice:"), voice)
	return rv
}

// Begins speaking synthesized text through the system’s default sound output device.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/startSpeaking(_:)
func (s_ SpeechSynthesizer) StartSpeakingString(string_ string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("startSpeakingString:"), objc.String(string_))
	return rv
}

// Begins synthesizing text into a sound (AIFF) file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/startSpeaking(_:to:)
func (s_ SpeechSynthesizer) StartSpeakingStringToURL(string_ string, url foundation.IURL) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("startSpeakingString:toURL:"), objc.String(string_), url)
	return rv
}

// Stops synthesis in progress.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/stopSpeaking()
func (s_ SpeechSynthesizer) StopSpeaking() {
	objc.Send[objc.ID](s_.ID, objc.Sel("stopSpeaking"))
}

// Stops synthesis in progress at a given boundary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/stopSpeaking(at:)
func (s_ SpeechSynthesizer) StopSpeakingAtBoundary(boundary ISpeechBoundary) {
	objc.Send[objc.ID](s_.ID, objc.Sel("stopSpeakingAtBoundary:"), boundary)
}

// Returns the identifier of the receiver’s current voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/voice()
func (s_ SpeechSynthesizer) Voice() SpeechSynthesizerVoiceName {
	rv := objc.Send[SpeechSynthesizerVoiceName](s_.ID, objc.Sel("voice"))
	return rv
}

// Provides the identifiers of the voices available on the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/availableVoices
func (s_ SpeechSynthesizer) AvailableVoices() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("availableVoices"))
	return rv
}

// Provides the identifier of the default voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/defaultVoice
func (s_ SpeechSynthesizer) DefaultVoice() SpeechSynthesizerVoiceName {
	rv := objc.Send[SpeechSynthesizerVoiceName](s_.ID, objc.Sel("defaultVoice"))
	return rv
}

// The synthesizer’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/delegate
func (s_ SpeechSynthesizer) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The synthesizer’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/delegate
func (s_ SpeechSynthesizer) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value indicating whether any application is currently speaking through the sound output device.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/isAnyApplicationSpeaking
func (s_ SpeechSynthesizer) AnyApplicationSpeaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("anyApplicationSpeaking"))
	return rv
}

// Indicates whether the receiver is currently generating synthesized speech.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/isSpeaking
func (s_ SpeechSynthesizer) Speaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("speaking"))
	return rv
}

// The synthesizer’s speaking rate (words per minute).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/rate
func (s_ SpeechSynthesizer) Rate() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
// The synthesizer’s speaking rate (words per minute).

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/rate
func (s_ SpeechSynthesizer) SetRate(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:"), value)
}

// Indicates whether the receiver uses the speech feedback window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/usesFeedbackWindow
func (s_ SpeechSynthesizer) UsesFeedbackWindow() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("usesFeedbackWindow"))
	return rv
}


// SetUsesFeedbackWindow sets the value of the usesFeedbackWindow property.
// Indicates whether the receiver uses the speech feedback window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/usesFeedbackWindow
func (s_ SpeechSynthesizer) SetUsesFeedbackWindow(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUsesFeedbackWindow:"), value)
}

// The synthesizer’s speaking volume.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/volume
func (s_ SpeechSynthesizer) Volume() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}


// SetVolume sets the value of the volume property.
// The synthesizer’s speaking volume.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/volume
func (s_ SpeechSynthesizer) SetVolume(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}

// Indicates whether the receiver is currently generating synthesized speech.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/isspeaking
func (s_ SpeechSynthesizer) IsSpeaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSpeaking"))
	return rv
}


// SetIsSpeaking sets the value of the isSpeaking property.
// Indicates whether the receiver is currently generating synthesized speech.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/isspeaking
func (s_ SpeechSynthesizer) SetIsSpeaking(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSpeaking:"), value)
}


