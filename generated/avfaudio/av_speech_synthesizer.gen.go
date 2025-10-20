// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

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
		SpeechSynthesizerClass = _SpeechSynthesizerClass{objc.GetClass("AVSpeechSynthesizer")}
	})
	return SpeechSynthesizerClass
}

type _SpeechSynthesizerClass struct {
	class objc.Class
}

// An interface definition for the [SpeechSynthesizer] class.
type ISpeechSynthesizer interface {
	objectivec.IObject
	SpeakUtterance(utterance unsafe.Pointer)
}

// An object that produces synthesized speech from text utterances and enables monitoring or controlling of ongoing speech.
//
// To speak some text, create an instance that contains the text and pass it to on a speech synthesizer instance. You can optionally also retrieve an and set it on the utterance’s property to have the speech synthesizer use that voice when speaking the utterance’s text. The speech synthesizer maintains a queue of utterances that it speaks. If the synthesizer isn’t speaking, calling begins speaking that utterance either immediately or after pausing for its , if necessary. If the synthesizer is speaking, the synthesizer adds utterances to a queue and speaks them in the order it receives them. After speech begins, you can use the synthesizer object to pause or stop speech. After pausing, you can resume the speech from its paused point or stop the speech entirely and remove all remaining utterances in the queue. You can monitor the speech synthesizer by examining its and properties, or by setting a delegate that conforms to . The delegate receives significant events as they occur during speech synthesis. An also controls the route where the speech plays. For more information, see Directing speech output.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer
type SpeechSynthesizer struct {
	objectivec.Object
}

// SpeechSynthesizerFrom constructs a [SpeechSynthesizer] from an unsafe.Pointer.
//
// An object that produces synthesized speech from text utterances and enables monitoring or controlling of ongoing speech.
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


// Adds the utterance you specify to the speech synthesizer’s queue.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speak(_:)
func (s_ SpeechSynthesizer) SpeakUtterance(utterance unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("speakUtterance:"), utterance)
}

// The delegate object for the speech synthesizer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/delegate
func (s_ SpeechSynthesizer) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object for the speech synthesizer.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/delegate
func (s_ SpeechSynthesizer) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}
// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isSpeaking
func (s_ SpeechSynthesizer) Speaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("speaking"))
	return rv
}



