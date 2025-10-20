// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [SFSpeechURLRecognitionRequest] class.
type ISFSpeechURLRecognitionRequest interface {
	ISFSpeechRecognitionRequest
}

// A request to recognize speech in a recorded audio file.
//
// Use this object to perform speech recognition on the contents of an audio file. The following example shows a method that performs recognition on an audio file based on the user’s default language and prints out the transcription. Listing 1. Getting a speech recognizer and making a recognition request
//
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

// Alloc allocates a new instance without initialization.
func (sc _SFSpeechURLRecognitionRequestClass) Alloc() SFSpeechURLRecognitionRequest {
	rv := objc.Send[SFSpeechURLRecognitionRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a speech recognition request, initialized with the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechURLRecognitionRequest/init(url:)
func NewSFSpeechURLRecognitionRequestWithURL(URL unsafe.Pointer) SFSpeechURLRecognitionRequest {
	instance := getSFSpeechURLRecognitionRequestClass().Alloc()
	rv := objc.Send[SFSpeechURLRecognitionRequest](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}


// The URL of the audio file.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechURLRecognitionRequest/url
func (s_ SFSpeechURLRecognitionRequest) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("URL"))
	return rv
}


