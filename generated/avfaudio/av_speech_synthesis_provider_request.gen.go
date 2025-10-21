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
}

// An object that represents the text to synthesize and the voice to use.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesisProviderRequestClass) Alloc() SpeechSynthesisProviderRequest {
	rv := objc.Send[SpeechSynthesisProviderRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The description of the text to synthesize.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisproviderrequest/ssmlrepresentation
func (s_ SpeechSynthesisProviderRequest) SsmlRepresentation() string {
	rv := objc.Send[string](s_.ID, objc.Sel("ssmlRepresentation"))
	return rv
}


// SetSsmlRepresentation sets the value of the ssmlRepresentation property.
// The description of the text to synthesize.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisproviderrequest/ssmlrepresentation
func (s_ SpeechSynthesisProviderRequest) SetSsmlRepresentation(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSsmlRepresentation:"), objc.String(value))
}

// The voice to use in the speech request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/voice
func (s_ SpeechSynthesisProviderRequest) Voice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("voice"))
	return rv
}



