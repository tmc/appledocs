// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [SpeechSynthesisProviderAudioUnit] class.
var (
	SpeechSynthesisProviderAudioUnitClass     _SpeechSynthesisProviderAudioUnitClass
	SpeechSynthesisProviderAudioUnitClassOnce sync.Once
)

func getSpeechSynthesisProviderAudioUnitClass() _SpeechSynthesisProviderAudioUnitClass {
	SpeechSynthesisProviderAudioUnitClassOnce.Do(func() {
		SpeechSynthesisProviderAudioUnitClass = _SpeechSynthesisProviderAudioUnitClass{objc.GetClass("AVSpeechSynthesisProviderAudioUnit")}
	})
	return SpeechSynthesisProviderAudioUnitClass
}

type _SpeechSynthesisProviderAudioUnitClass struct {
	class objc.Class
}





// An interface definition for the [SpeechSynthesisProviderAudioUnit] class.
type ISpeechSynthesisProviderAudioUnit interface {
	IAudioUnit
	

	// properties:
	SpeechSynthesisOutputMetadataBlock() SpeechSynthesisProviderOutputBlock /* not a class type */
	SetSpeechSynthesisOutputMetadataBlock(value SpeechSynthesisProviderOutputBlock /* not a class type */)
	SpeechVoices() []SpeechSynthesisProviderVoice
	SetSpeechVoices(value []SpeechSynthesisProviderVoice)


	

	// methods:
	CancelSpeechRequest()
	SynthesizeSpeechRequest(speechRequest IAVSpeechSynthesisProviderRequest)


}





// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesisProviderAudioUnitClass) Alloc() SpeechSynthesisProviderAudioUnit {
	rv := objc.Send[SpeechSynthesisProviderAudioUnit](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SpeechSynthesisProviderAudioUnitClass) New() SpeechSynthesisProviderAudioUnit {
	rv := objc.Send[SpeechSynthesisProviderAudioUnit](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpeechSynthesisProviderAudioUnit) Init() SpeechSynthesisProviderAudioUnit {
	rv := objc.Send[SpeechSynthesisProviderAudioUnit](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpeechSynthesisProviderAudioUnit) Autorelease() SpeechSynthesisProviderAudioUnit {
	rv := objc.Send[SpeechSynthesisProviderAudioUnit](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeechSynthesisProviderAudioUnit creates a new SpeechSynthesisProviderAudioUnit instance.
func NewSpeechSynthesisProviderAudioUnit() SpeechSynthesisProviderAudioUnit {
	return getSpeechSynthesisProviderAudioUnitClass().New()
}





// An object that generates speech from text.
//
// Use a speech synthesizer audio unit to generate audio buffers that contain speech for a given voice and speech markup. The audio unit receives an as input, and extracts audio buffers through the render block. Use to provide metadata as an array of . The system scans and loads voices for audio unit extensions of this type, and the voices it provides are available for use in and accessibility technologies like VoiceOver and Speak Screen.


// An object that generates speech from text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit
type SpeechSynthesisProviderAudioUnit struct {
	AudioUnit
}

// SpeechSynthesisProviderAudioUnitFrom constructs a [SpeechSynthesisProviderAudioUnit] from an unsafe.Pointer.
//
// An object that generates speech from text.
func SpeechSynthesisProviderAudioUnitFrom(ptr unsafe.Pointer) SpeechSynthesisProviderAudioUnit {
	return SpeechSynthesisProviderAudioUnit{
		AudioUnit: AudioUnitFrom(ptr),
	}
}




















// Informs the audio unit to discard the speech request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit/cancelSpeechRequest()
func (s_ SpeechSynthesisProviderAudioUnit) CancelSpeechRequest() {
	objc.Send[objc.ID](s_.ID, objc.Sel("cancelSpeechRequest"))
}


// Sets the text to synthesize and the voice to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit/synthesizeSpeechRequest(_:)
func (s_ SpeechSynthesisProviderAudioUnit) SynthesizeSpeechRequest(speechRequest IAVSpeechSynthesisProviderRequest) {
	objc.Send[objc.ID](s_.ID, objc.Sel("synthesizeSpeechRequest:"), speechRequest)
}







// A block that subclasses use to send marker information to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit/speechSynthesisOutputMetadataBlock
func (s_ SpeechSynthesisProviderAudioUnit) SpeechSynthesisOutputMetadataBlock() SpeechSynthesisProviderOutputBlock /* not a class type */ {
	rv := objc.Send[SpeechSynthesisProviderOutputBlock](s_.ID, objc.Sel("speechSynthesisOutputMetadataBlock"))
	return rv
}


// A block that subclasses use to send marker information to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit/speechSynthesisOutputMetadataBlock
func (s_ SpeechSynthesisProviderAudioUnit) SetSpeechSynthesisOutputMetadataBlock(value SpeechSynthesisProviderOutputBlock /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpeechSynthesisOutputMetadataBlock:"), value)
}


// A list of voices the audio unit provides to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit/speechVoices
func (s_ SpeechSynthesisProviderAudioUnit) SpeechVoices() []SpeechSynthesisProviderVoice {
	rv := objc.Send[[]SpeechSynthesisProviderVoice](s_.ID, objc.Sel("speechVoices"))
	return rv
}


// A list of voices the audio unit provides to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit/speechVoices
func (s_ SpeechSynthesisProviderAudioUnit) SetSpeechVoices(value []SpeechSynthesisProviderVoice) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpeechVoices:"), nsArray)
}








