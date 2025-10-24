// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

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
	// properties:
	Delegate() SpeechSynthesizerDelegate /* not a class type */
	SetDelegate(value SpeechSynthesizerDelegate /* not a class type */)
	IsSpeaking() bool
	SetIsSpeaking(value bool)
	Rate() float32
	SetRate(value float32)
	UsesFeedbackWindow() bool
	SetUsesFeedbackWindow(value bool)
	Volume() float32
	SetVolume(value float32)
	// methods:
}

// The Cocoa interface to speech synthesis in macOS.
//
// Speech synthesis, also called text-to-speech (TTS), parses text and converts it into audible speech. It offers a concurrent feedback mode that can be used in concert with or in place of traditional visual and aural notifications. For example, your application can use a speech synthesizer object to “pronounce” the text of important alert dialogs. Synthesized speech has several advantages. It can provide urgent information to users without forcing them to shift attention from their current task. And because speech doesn’t rely on visual elements for meaning, it is a crucial technology for users with vision or attention disabilities. In addition, synthesized speech can help save system resources. Because sound samples can take up large amounts of room on disk, using text in place of sampled sound is extremely efficient, and so a multimedia application might use an object to provide a narration of a QuickTime movie instead of including sampled-sound data on a movie track. When you create an instance using the default initializer ( ), the class uses the selected in System Preferences > Speech. Alternatively, you can select a specific voice for an instance by initializing it with . To begin synthesis, send either or to the instance. The former generates speech through the system’s default sound output device; the latter saves the generated speech to a file. If you wish to be notified when the current speech concludes, set the property and implement the delegate method . Speech synthesis is just one of the macOS speech technologies. The speech recognizer technology allows applications to “listen to” text spoken in U.S. English; the class is the Cocoa interface to this technology. Both technologies provide benefits for all users, and are particularly useful to those users who have difficulties seeing the screen or using the mouse and keyboard.

// The Cocoa interface to speech synthesis in macOS.
//
// [Full Topic]
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

// The synthesizer’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/delegate
func (s_ SpeechSynthesizer) Delegate() SpeechSynthesizerDelegate /* not a class type */ {
	rv := objc.Send[SpeechSynthesizerDelegate](s_.ID, objc.Sel("delegate"))
	return rv
}

// The synthesizer’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/delegate
func (s_ SpeechSynthesizer) SetDelegate(value SpeechSynthesizerDelegate /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// Indicates whether the receiver is currently generating synthesized speech.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/isspeaking
func (s_ SpeechSynthesizer) IsSpeaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSpeaking"))
	return rv
}

// Indicates whether the receiver is currently generating synthesized speech.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/isspeaking
func (s_ SpeechSynthesizer) SetIsSpeaking(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSpeaking:"), value)
}

// The synthesizer’s speaking rate (words per minute).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/rate
func (s_ SpeechSynthesizer) Rate() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("rate"))
	return rv
}

// The synthesizer’s speaking rate (words per minute).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/rate
func (s_ SpeechSynthesizer) SetRate(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:"), value)
}

// Indicates whether the receiver uses the speech feedback window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/usesfeedbackwindow
func (s_ SpeechSynthesizer) UsesFeedbackWindow() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("usesFeedbackWindow"))
	return rv
}

// Indicates whether the receiver uses the speech feedback window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/usesfeedbackwindow
func (s_ SpeechSynthesizer) SetUsesFeedbackWindow(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUsesFeedbackWindow:"), value)
}

// The synthesizer’s speaking volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/volume
func (s_ SpeechSynthesizer) Volume() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}

// The synthesizer’s speaking volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/volume
func (s_ SpeechSynthesizer) SetVolume(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}
