// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An abstract class that represents a request to recognize speech from an audio source.
//
// Don’t create objects directly. Create an or object instead. Use the properties of this class to configure various aspects of your request object before you start the speech recognition process. For example, use the property to specify whether you want partial results or only the final result of speech recognition.
//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/customizedLanguageModel
func (s_ SFSpeechRecognitionRequest) CustomizedLanguageModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("customizedLanguageModel"))
	return rv
}


// SetCustomizedLanguageModel sets the value of the customizedLanguageModel property.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/customizedLanguageModel
func (s_ SFSpeechRecognitionRequest) SetCustomizedLanguageModel(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizedLanguageModel:"), value)
}

// A Boolean value that determines whether a request must keep its audio data on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/requiresOnDeviceRecognition
func (s_ SFSpeechRecognitionRequest) RequiresOnDeviceRecognition() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("requiresOnDeviceRecognition"))
	return rv
}


// SetRequiresOnDeviceRecognition sets the value of the requiresOnDeviceRecognition property.
// A Boolean value that determines whether a request must keep its audio data on the device.

//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/requiresOnDeviceRecognition
func (s_ SFSpeechRecognitionRequest) SetRequiresOnDeviceRecognition(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRequiresOnDeviceRecognition:"), value)
}

// A Boolean value that indicates whether you want intermediate results returned for each utterance.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/shouldReportPartialResults
func (s_ SFSpeechRecognitionRequest) ShouldReportPartialResults() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldReportPartialResults"))
	return rv
}


// SetShouldReportPartialResults sets the value of the shouldReportPartialResults property.
// A Boolean value that indicates whether you want intermediate results returned for each utterance.

//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/shouldReportPartialResults
func (s_ SFSpeechRecognitionRequest) SetShouldReportPartialResults(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldReportPartialResults:"), value)
}

// A Boolean value that indicates whether to add punctuation to speech recognition results.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/addspunctuation
func (s_ SFSpeechRecognitionRequest) AddsPunctuation() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("addsPunctuation"))
	return rv
}


// SetAddsPunctuation sets the value of the addsPunctuation property.
// A Boolean value that indicates whether to add punctuation to speech recognition results.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/addspunctuation
func (s_ SFSpeechRecognitionRequest) SetAddsPunctuation(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAddsPunctuation:"), value)
}

// An array of phrases that should be recognized, even if they are not in the system vocabulary.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/contextualstrings
func (s_ SFSpeechRecognitionRequest) ContextualStrings() string {
	rv := objc.Send[string](s_.ID, objc.Sel("contextualStrings"))
	return rv
}


// SetContextualStrings sets the value of the contextualStrings property.
// An array of phrases that should be recognized, even if they are not in the system vocabulary.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/contextualstrings
func (s_ SFSpeechRecognitionRequest) SetContextualStrings(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContextualStrings:"), objc.String(value))
}

// An identifier string that you use to describe the type of interaction associated with the speech recognition request.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/interactionidentifier
func (s_ SFSpeechRecognitionRequest) InteractionIdentifier() string {
	rv := objc.Send[string](s_.ID, objc.Sel("interactionIdentifier"))
	return rv
}


// SetInteractionIdentifier sets the value of the interactionIdentifier property.
// An identifier string that you use to describe the type of interaction associated with the speech recognition request.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/interactionidentifier
func (s_ SFSpeechRecognitionRequest) SetInteractionIdentifier(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInteractionIdentifier:"), objc.String(value))
}

// A value that indicates the type of speech recognition being performed.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/taskhint
func (s_ SFSpeechRecognitionRequest) TaskHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("taskHint"))
	return rv
}


// SetTaskHint sets the value of the taskHint property.
// A value that indicates the type of speech recognition being performed.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/taskhint
func (s_ SFSpeechRecognitionRequest) SetTaskHint(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTaskHint:"), value)
}



