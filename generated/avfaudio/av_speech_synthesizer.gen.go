// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSpeechSynthesizer */


/* debug [class_header]: Header for AVSpeechSynthesizer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpeechSynthesizer */
// An interface definition for the [SpeechSynthesizer] class.
type ISpeechSynthesizer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SpeechSynthesizer */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Paused() bool
	Speaking() bool
	IsPaused() bool
	SetIsPaused(value bool)
	IsSpeaking() bool
	SetIsSpeaking(value bool)
	PreUtteranceDelay() float64
	SetPreUtteranceDelay(value float64)
	Voice() IAVSpeechSynthesisVoice
	SetVoice(value IAVSpeechSynthesisVoice)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpeechSynthesizer */
	// methods:
	ContinueSpeaking() bool
	PauseSpeakingAtBoundary(boundary SpeechBoundary) bool
	SpeakUtterance(utterance IAVSpeechUtterance)
	StopSpeakingAtBoundary(boundary SpeechBoundary) bool
	WriteUtteranceToBufferCallback(utterance IAVSpeechUtterance, bufferCallback SpeechSynthesizerBufferCallback /* not a class type */)
	WriteUtteranceToBufferCallbackToMarkerCallback(utterance IAVSpeechUtterance, bufferCallback SpeechSynthesizerBufferCallback /* not a class type */, markerCallback SpeechSynthesizerMarkerCallback /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpeechSynthesizer */
// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesizerClass) Alloc() SpeechSynthesizer {
	rv := objc.Send[SpeechSynthesizer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpeechSynthesizer */
// An object that produces synthesized speech from text utterances and enables monitoring or controlling of ongoing speech.
//
// To speak some text, create an instance that contains the text and pass it to on a speech synthesizer instance. You can optionally also retrieve an and set it on the utterance’s property to have the speech synthesizer use that voice when speaking the utterance’s text. The speech synthesizer maintains a queue of utterances that it speaks. If the synthesizer isn’t speaking, calling begins speaking that utterance either immediately or after pausing for its , if necessary. If the synthesizer is speaking, the synthesizer adds utterances to a queue and speaks them in the order it receives them. After speech begins, you can use the synthesizer object to pause or stop speech. After pausing, you can resume the speech from its paused point or stop the speech entirely and remove all remaining utterances in the queue. You can monitor the speech synthesizer by examining its and properties, or by setting a delegate that conforms to . The delegate receives significant events as they occur during speech synthesis. An also controls the route where the speech plays. For more information, see Directing speech output.


// An object that produces synthesized speech from text utterances and enables monitoring or controlling of ongoing speech.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpeechSynthesizer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpeechSynthesizer */

// Prompts the user to authorize your app to use personal voices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/requestPersonalVoiceAuthorization(completionHandler:)
func (sc _SpeechSynthesizerClass) RequestPersonalVoiceAuthorizationWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("requestPersonalVoiceAuthorizationWithCompletionHandler:"), handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestPersonalVoiceAuthorizationWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpeechSynthesizer */

// Your app’s authorization to use personal voices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/personalVoiceAuthorizationStatus-swift.type.property
func (sc _SpeechSynthesizerClass) PersonalVoiceAuthorizationStatus() SpeechSynthesisPersonalVoiceAuthorizationStatus {
	rv := objc.Send[SpeechSynthesisPersonalVoiceAuthorizationStatus](objc.ID(sc.class), objc.Sel("personalVoiceAuthorizationStatus"))
	return rv
}/* debug [class_properties_class/property]: personalVoiceAuthorizationStatus */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpeechSynthesizer */

// Resumes speech from its paused point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/continueSpeaking()
func (s_ SpeechSynthesizer) ContinueSpeaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("continueSpeaking"))
	return rv
}/* debug [instance_methods/method]: ContinueSpeaking */


// Pauses speech at the boundary you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/pauseSpeaking(at:)
func (s_ SpeechSynthesizer) PauseSpeakingAtBoundary(boundary SpeechBoundary) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("pauseSpeakingAtBoundary:"), boundary)
	return rv
}/* debug [instance_methods/method]: PauseSpeakingAtBoundary */


// Adds the utterance you specify to the speech synthesizer’s queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speak(_:)
func (s_ SpeechSynthesizer) SpeakUtterance(utterance IAVSpeechUtterance) {
	objc.Send[objc.ID](s_.ID, objc.Sel("speakUtterance:"), utterance)
}/* debug [instance_methods/method]: SpeakUtterance */


// Stops speech at the boundary you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/stopSpeaking(at:)
func (s_ SpeechSynthesizer) StopSpeakingAtBoundary(boundary SpeechBoundary) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("stopSpeakingAtBoundary:"), boundary)
	return rv
}/* debug [instance_methods/method]: StopSpeakingAtBoundary */


// Generates speech for the utterance and invokes the callback with the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/write(_:toBufferCallback:)
func (s_ SpeechSynthesizer) WriteUtteranceToBufferCallback(utterance IAVSpeechUtterance, bufferCallback SpeechSynthesizerBufferCallback /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("writeUtterance:toBufferCallback:"), utterance, bufferCallback)
}/* debug [instance_methods/method]: WriteUtteranceToBufferCallback */


// Generates audio buffers and associated metadata for storage or further speech synthesis processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/write(_:toBufferCallback:toMarkerCallback:)
func (s_ SpeechSynthesizer) WriteUtteranceToBufferCallbackToMarkerCallback(utterance IAVSpeechUtterance, bufferCallback SpeechSynthesizerBufferCallback /* not a class type */, markerCallback SpeechSynthesizerMarkerCallback /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("writeUtterance:toBufferCallback:toMarkerCallback:"), utterance, bufferCallback, markerCallback)
}/* debug [instance_methods/method]: WriteUtteranceToBufferCallbackToMarkerCallback */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpeechSynthesizer */

// The delegate object for the speech synthesizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/delegate
func (s_ SpeechSynthesizer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object for the speech synthesizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/delegate
func (s_ SpeechSynthesizer) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether a speech synthesizer is in a paused state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isPaused
func (s_ SpeechSynthesizer) Paused() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("paused"))
	return rv
}/* debug [instance_properties/getter]: paused */


// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isSpeaking
func (s_ SpeechSynthesizer) Speaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("speaking"))
	return rv
}/* debug [instance_properties/getter]: speaking */


// Your app’s authorization to use personal voices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/personalVoiceAuthorizationStatus-swift.type.property
func (s_ SpeechSynthesizer) PersonalVoiceAuthorizationStatus() SpeechSynthesisPersonalVoiceAuthorizationStatus {
	rv := objc.Send[SpeechSynthesisPersonalVoiceAuthorizationStatus](s_.ID, objc.Sel("personalVoiceAuthorizationStatus"))
	return rv
}/* debug [instance_properties/getter]: personalVoiceAuthorizationStatus */


// A Boolean value that indicates whether a speech synthesizer is in a paused state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/ispaused
func (s_ SpeechSynthesizer) IsPaused() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isPaused"))
	return rv
}/* debug [instance_properties/getter]: isPaused */


// A Boolean value that indicates whether a speech synthesizer is in a paused state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/ispaused
func (s_ SpeechSynthesizer) SetIsPaused(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsPaused:"), value)
}/* debug [instance_properties/setter]: isPaused */


// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/isspeaking
func (s_ SpeechSynthesizer) IsSpeaking() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSpeaking"))
	return rv
}/* debug [instance_properties/getter]: isSpeaking */


// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/isspeaking
func (s_ SpeechSynthesizer) SetIsSpeaking(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSpeaking:"), value)
}/* debug [instance_properties/setter]: isSpeaking */


// The amount of time the speech synthesizer pauses before speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/preutterancedelay
func (s_ SpeechSynthesizer) PreUtteranceDelay() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("preUtteranceDelay"))
	return rv
}/* debug [instance_properties/getter]: preUtteranceDelay */


// The amount of time the speech synthesizer pauses before speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/preutterancedelay
func (s_ SpeechSynthesizer) SetPreUtteranceDelay(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreUtteranceDelay:"), value)
}/* debug [instance_properties/setter]: preUtteranceDelay */


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s_ SpeechSynthesizer) Voice() IAVSpeechSynthesisVoice {
	rv := objc.Send[SpeechSynthesisVoice](s_.ID, objc.Sel("voice"))
	return rv
}/* debug [instance_properties/getter]: voice */


// The voice the speech synthesizer uses when speaking the utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s_ SpeechSynthesizer) SetVoice(value IAVSpeechSynthesisVoice) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoice:"), value)
}/* debug [instance_properties/setter]: voice */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSpeechSynthesizer */


