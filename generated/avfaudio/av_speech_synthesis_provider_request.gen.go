// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSpeechSynthesisProviderRequest */


/* debug [class_header]: Header for AVSpeechSynthesisProviderRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpeechSynthesisProviderRequest */
// An interface definition for the [SpeechSynthesisProviderRequest] class.
type ISpeechSynthesisProviderRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SpeechSynthesisProviderRequest */
	// properties:
	SsmlRepresentation() objc.IObject /* cross-framework: NSString */
	Voice() IAVSpeechSynthesisProviderVoice
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpeechSynthesisProviderRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpeechSynthesisProviderRequest */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpeechSynthesisProviderRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpeechSynthesisProviderRequest */

// Creates a request with a voice and a description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/init(ssmlRepresentation:voice:)
func NewSpeechSynthesisProviderRequestWithSSMLRepresentationVoice(text objc.IObject /* cross-framework: NSString */, voice IAVSpeechSynthesisProviderVoice) SpeechSynthesisProviderRequest {
	instance := getSpeechSynthesisProviderRequestClass().Alloc()
	rv := objc.Send[SpeechSynthesisProviderRequest](instance.ID, objc.Sel("initWithSSMLRepresentation:voice:"), text, voice)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechSynthesisProviderRequestWithSSMLRepresentationVoice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpeechSynthesisProviderRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpeechSynthesisProviderRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpeechSynthesisProviderRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpeechSynthesisProviderRequest */

// The description of the text to synthesize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/ssmlRepresentation
func (s_ SpeechSynthesisProviderRequest) SsmlRepresentation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("ssmlRepresentation"))
	return rv
}/* debug [instance_properties/getter]: ssmlRepresentation */


// The voice to use in the speech request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/voice
func (s_ SpeechSynthesisProviderRequest) Voice() IAVSpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](s_.ID, objc.Sel("voice"))
	return rv
}/* debug [instance_properties/getter]: voice */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSpeechSynthesisProviderRequest */


