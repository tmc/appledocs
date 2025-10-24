// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class SFSpeechURLRecognitionRequest */

/* debug [class_header]: Header for SFSpeechURLRecognitionRequest */
// The class instance for the [SFSpeechURLRecognitionRequest] class.
var (
	SFSpeechURLRecognitionRequestClass     _SFSpeechURLRecognitionRequestClass
	SFSpeechURLRecognitionRequestClassOnce sync.Once
)

func getSFSpeechURLRecognitionRequestClass() _SFSpeechURLRecognitionRequestClass {
	SFSpeechURLRecognitionRequestClassOnce.Do(func() {
		SFSpeechURLRecognitionRequestClass = _SFSpeechURLRecognitionRequestClass{objc.GetClass("SFSpeechURLRecognitionRequest")}
	})
	return SFSpeechURLRecognitionRequestClass
}

type _SFSpeechURLRecognitionRequestClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFSpeechURLRecognitionRequest */
// An interface definition for the [SFSpeechURLRecognitionRequest] class.
type ISFSpeechURLRecognitionRequest interface {
	ISFSpeechRecognitionRequest

	/* debug [class_interface_properties]: Properties for SFSpeechURLRecognitionRequest */
	// properties:
	URL() objc.IObject /* cross-framework: NSURL */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFSpeechURLRecognitionRequest */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFSpeechURLRecognitionRequest */
// Alloc allocates a new instance without initialization.
func (sc _SFSpeechURLRecognitionRequestClass) Alloc() SFSpeechURLRecognitionRequest {
	rv := objc.Send[SFSpeechURLRecognitionRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSpeechURLRecognitionRequestClass) New() SFSpeechURLRecognitionRequest {
	rv := objc.Send[SFSpeechURLRecognitionRequest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSpeechURLRecognitionRequest) Init() SFSpeechURLRecognitionRequest {
	rv := objc.Send[SFSpeechURLRecognitionRequest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSpeechURLRecognitionRequest) Autorelease() SFSpeechURLRecognitionRequest {
	rv := objc.Send[SFSpeechURLRecognitionRequest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechURLRecognitionRequest creates a new SFSpeechURLRecognitionRequest instance.
func NewSFSpeechURLRecognitionRequest() SFSpeechURLRecognitionRequest {
	return getSFSpeechURLRecognitionRequestClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFSpeechURLRecognitionRequest */
// A request to recognize speech in a recorded audio file.
//
// Use this object to perform speech recognition on the contents of an audio file. The following example shows a method that performs recognition on an audio file based on the user’s default language and prints out the transcription. Listing 1. Getting a speech recognizer and making a recognition request

// A request to recognize speech in a recorded audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechURLRecognitionRequest
type SFSpeechURLRecognitionRequest struct {
	SFSpeechRecognitionRequest
}

// SFSpeechURLRecognitionRequestFrom constructs a [SFSpeechURLRecognitionRequest] from an unsafe.Pointer.
//
// A request to recognize speech in a recorded audio file.
func SFSpeechURLRecognitionRequestFrom(ptr unsafe.Pointer) SFSpeechURLRecognitionRequest {
	return SFSpeechURLRecognitionRequest{
		SFSpeechRecognitionRequest: SFSpeechRecognitionRequestFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFSpeechURLRecognitionRequest */

// Creates a speech recognition request, initialized with the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechURLRecognitionRequest/init(url:)
func NewSFSpeechURLRecognitionRequestWithURL(URL objc.IObject /* cross-framework: NSURL */) SFSpeechURLRecognitionRequest {
	instance := getSFSpeechURLRecognitionRequestClass().Alloc()
	rv := objc.Send[SFSpeechURLRecognitionRequest](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewSFSpeechURLRecognitionRequestWithURL */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFSpeechURLRecognitionRequest */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFSpeechURLRecognitionRequest */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFSpeechURLRecognitionRequest */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFSpeechURLRecognitionRequest */

// The URL of the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechURLRecognitionRequest/url
func (s_ SFSpeechURLRecognitionRequest) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("URL"))
	return rv
} /* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFSpeechURLRecognitionRequest */
