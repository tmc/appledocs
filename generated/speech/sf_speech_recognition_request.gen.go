// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSpeechRecognitionRequest] class.
var (
	SFSpeechRecognitionRequestClass     _SFSpeechRecognitionRequestClass
	SFSpeechRecognitionRequestClassOnce sync.Once
)

func getSFSpeechRecognitionRequestClass() _SFSpeechRecognitionRequestClass {
	SFSpeechRecognitionRequestClassOnce.Do(func() {
		SFSpeechRecognitionRequestClass = _SFSpeechRecognitionRequestClass{objc.GetClass("SFSpeechRecognitionRequest")}
	})
	return SFSpeechRecognitionRequestClass
}

type _SFSpeechRecognitionRequestClass struct {
	class objc.Class
}

// An interface definition for the [SFSpeechRecognitionRequest] class.
type ISFSpeechRecognitionRequest interface {
	objectivec.IObject
	// properties:
	CustomizedLanguageModel() ISFSpeechLanguageModelConfiguration
	SetCustomizedLanguageModel(value ISFSpeechLanguageModelConfiguration)
	ShouldReportPartialResults() bool
	SetShouldReportPartialResults(value bool)
	AddsPunctuation() bool
	SetAddsPunctuation(value bool)
	ContextualStrings() objc.IObject /* cross-framework: NSString */
	SetContextualStrings(value objc.IObject /* cross-framework: NSString */)
	InteractionIdentifier() objc.IObject /* cross-framework: NSString */
	SetInteractionIdentifier(value objc.IObject /* cross-framework: NSString */)
	RequiresOnDeviceRecognition() bool
	SetRequiresOnDeviceRecognition(value bool)
	TaskHint() SFSpeechRecognitionTaskHint
	SetTaskHint(value SFSpeechRecognitionTaskHint)
	// methods:
}

// An abstract class that represents a request to recognize speech from an audio source.
//
// Don’t create objects directly. Create an or object instead. Use the properties of this class to configure various aspects of your request object before you start the speech recognition process. For example, use the property to specify whether you want partial results or only the final result of speech recognition.


// An abstract class that represents a request to recognize speech from an audio source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest
type SFSpeechRecognitionRequest struct {
	objectivec.Object
}

// SFSpeechRecognitionRequestFrom constructs a [SFSpeechRecognitionRequest] from an unsafe.Pointer.
//
// An abstract class that represents a request to recognize speech from an audio source.
func SFSpeechRecognitionRequestFrom(ptr unsafe.Pointer) SFSpeechRecognitionRequest {
	return SFSpeechRecognitionRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSpeechRecognitionRequestClass) Alloc() SFSpeechRecognitionRequest {
	rv := objc.Send[SFSpeechRecognitionRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSpeechRecognitionRequestClass) New() SFSpeechRecognitionRequest {
	rv := objc.Send[SFSpeechRecognitionRequest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSpeechRecognitionRequest) Init() SFSpeechRecognitionRequest {
	rv := objc.Send[SFSpeechRecognitionRequest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSpeechRecognitionRequest) Autorelease() SFSpeechRecognitionRequest {
	rv := objc.Send[SFSpeechRecognitionRequest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechRecognitionRequest creates a new SFSpeechRecognitionRequest instance.
func NewSFSpeechRecognitionRequest() SFSpeechRecognitionRequest {
	return getSFSpeechRecognitionRequestClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/customizedLanguageModel
func (s_ SFSpeechRecognitionRequest) CustomizedLanguageModel() ISFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s_.ID, objc.Sel("customizedLanguageModel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/customizedLanguageModel
func (s_ SFSpeechRecognitionRequest) SetCustomizedLanguageModel(value ISFSpeechLanguageModelConfiguration) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizedLanguageModel:"), value)
}


// A Boolean value that indicates whether you want intermediate results returned for each utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/shouldReportPartialResults
func (s_ SFSpeechRecognitionRequest) ShouldReportPartialResults() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldReportPartialResults"))
	return rv
}


// A Boolean value that indicates whether you want intermediate results returned for each utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/shouldReportPartialResults
func (s_ SFSpeechRecognitionRequest) SetShouldReportPartialResults(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldReportPartialResults:"), value)
}


// A Boolean value that indicates whether to add punctuation to speech recognition results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/addspunctuation
func (s_ SFSpeechRecognitionRequest) AddsPunctuation() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("addsPunctuation"))
	return rv
}


// A Boolean value that indicates whether to add punctuation to speech recognition results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/addspunctuation
func (s_ SFSpeechRecognitionRequest) SetAddsPunctuation(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAddsPunctuation:"), value)
}


// An array of phrases that should be recognized, even if they are not in the system vocabulary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/contextualstrings
func (s_ SFSpeechRecognitionRequest) ContextualStrings() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("contextualStrings"))
	return rv
}


// An array of phrases that should be recognized, even if they are not in the system vocabulary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/contextualstrings
func (s_ SFSpeechRecognitionRequest) SetContextualStrings(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContextualStrings:"), value)
}


// An identifier string that you use to describe the type of interaction associated with the speech recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/interactionidentifier
func (s_ SFSpeechRecognitionRequest) InteractionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("interactionIdentifier"))
	return rv
}


// An identifier string that you use to describe the type of interaction associated with the speech recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/interactionidentifier
func (s_ SFSpeechRecognitionRequest) SetInteractionIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInteractionIdentifier:"), value)
}


// A Boolean value that determines whether a request must keep its audio data on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/requiresondevicerecognition
func (s_ SFSpeechRecognitionRequest) RequiresOnDeviceRecognition() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("requiresOnDeviceRecognition"))
	return rv
}


// A Boolean value that determines whether a request must keep its audio data on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/requiresondevicerecognition
func (s_ SFSpeechRecognitionRequest) SetRequiresOnDeviceRecognition(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRequiresOnDeviceRecognition:"), value)
}


// A value that indicates the type of speech recognition being performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/taskhint
func (s_ SFSpeechRecognitionRequest) TaskHint() SFSpeechRecognitionTaskHint {
	rv := objc.Send[SFSpeechRecognitionTaskHint](s_.ID, objc.Sel("taskHint"))
	return rv
}


// A value that indicates the type of speech recognition being performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/taskhint
func (s_ SFSpeechRecognitionRequest) SetTaskHint(value SFSpeechRecognitionTaskHint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTaskHint:"), value)
}



