// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SpeechSynthesisProviderRequest] class.
var (
	SpeechSynthesisProviderRequestClass     _SpeechSynthesisProviderRequestClass
	SpeechSynthesisProviderRequestClassOnce sync.Once
)

func getSpeechSynthesisProviderRequestClass() _SpeechSynthesisProviderRequestClass {
	SpeechSynthesisProviderRequestClassOnce.Do(func() {
		SpeechSynthesisProviderRequestClass = _SpeechSynthesisProviderRequestClass{objc.GetClass("AVSpeechSynthesisProviderRequest")}
	})
	return SpeechSynthesisProviderRequestClass
}

type _SpeechSynthesisProviderRequestClass struct {
	class objc.Class
}





// An interface definition for the [SpeechSynthesisProviderRequest] class.
type ISpeechSynthesisProviderRequest interface {
	objectivec.IObject
	

	// properties:
	SsmlRepresentation() objc.IObject /* cross-framework: NSString */
	Voice() IAVSpeechSynthesisProviderVoice


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesisProviderRequestClass) Alloc() SpeechSynthesisProviderRequest {
	rv := objc.Send[SpeechSynthesisProviderRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SpeechSynthesisProviderRequestClass) New() SpeechSynthesisProviderRequest {
	rv := objc.Send[SpeechSynthesisProviderRequest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpeechSynthesisProviderRequest) Init() SpeechSynthesisProviderRequest {
	rv := objc.Send[SpeechSynthesisProviderRequest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpeechSynthesisProviderRequest) Autorelease() SpeechSynthesisProviderRequest {
	rv := objc.Send[SpeechSynthesisProviderRequest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeechSynthesisProviderRequest creates a new SpeechSynthesisProviderRequest instance.
func NewSpeechSynthesisProviderRequest() SpeechSynthesisProviderRequest {
	return getSpeechSynthesisProviderRequestClass().New()
}





// An object that represents the text to synthesize and the voice to use.


// An object that represents the text to synthesize and the voice to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest
type SpeechSynthesisProviderRequest struct {
	objectivec.Object
}

// SpeechSynthesisProviderRequestFrom constructs a [SpeechSynthesisProviderRequest] from an unsafe.Pointer.
//
// An object that represents the text to synthesize and the voice to use.
func SpeechSynthesisProviderRequestFrom(ptr unsafe.Pointer) SpeechSynthesisProviderRequest {
	return SpeechSynthesisProviderRequest{objectivec.Object{objc.ID(ptr)}}
}






// Creates a request with a voice and a description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/init(ssmlRepresentation:voice:)
func NewSpeechSynthesisProviderRequestWithSSMLRepresentationVoice(text objc.IObject /* cross-framework: NSString */, voice IAVSpeechSynthesisProviderVoice) SpeechSynthesisProviderRequest {
	instance := getSpeechSynthesisProviderRequestClass().Alloc()
	rv := objc.Send[SpeechSynthesisProviderRequest](instance.ID, objc.Sel("initWithSSMLRepresentation:voice:"), text, voice)
	rv.Autorelease()
	return rv
}






















// The description of the text to synthesize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/ssmlRepresentation
func (s_ SpeechSynthesisProviderRequest) SsmlRepresentation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("ssmlRepresentation"))
	return rv
}


// The voice to use in the speech request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/voice
func (s_ SpeechSynthesisProviderRequest) Voice() IAVSpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](s_.ID, objc.Sel("voice"))
	return rv
}







