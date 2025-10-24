// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSpeechRecognitionRequest */

/* debug [class_header]: Header for SFSpeechRecognitionRequest */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFSpeechRecognitionRequest */
// An interface definition for the [SFSpeechRecognitionRequest] class.
type ISFSpeechRecognitionRequest interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for SFSpeechRecognitionRequest */
	// properties:
	AddsPunctuation() bool
	SetAddsPunctuation(value bool)
	ContextualStrings() []string
	SetContextualStrings(value []string)
	CustomizedLanguageModel() ISFSpeechLanguageModelConfiguration
	SetCustomizedLanguageModel(value ISFSpeechLanguageModelConfiguration)
	InteractionIdentifier() objc.IObject /* cross-framework: NSString */
	SetInteractionIdentifier(value objc.IObject /* cross-framework: NSString */)
	RequiresOnDeviceRecognition() bool
	SetRequiresOnDeviceRecognition(value bool)
	ShouldReportPartialResults() bool
	SetShouldReportPartialResults(value bool)
	TaskHint() SFSpeechRecognitionTaskHint
	SetTaskHint(value SFSpeechRecognitionTaskHint)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFSpeechRecognitionRequest */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFSpeechRecognitionRequest */
// Alloc allocates a new instance without initialization.
func (sc _SFSpeechRecognitionRequestClass) Alloc() SFSpeechRecognitionRequest {
	rv := objc.Send[SFSpeechRecognitionRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFSpeechRecognitionRequest */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFSpeechRecognitionRequest */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFSpeechRecognitionRequest */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFSpeechRecognitionRequest */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFSpeechRecognitionRequest */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFSpeechRecognitionRequest */

// A Boolean value that indicates whether to add punctuation to speech recognition results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/addsPunctuation
func (s_ SFSpeechRecognitionRequest) AddsPunctuation() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("addsPunctuation"))
	return rv
} /* debug [instance_properties/getter]: addsPunctuation */

// A Boolean value that indicates whether to add punctuation to speech recognition results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/addsPunctuation
func (s_ SFSpeechRecognitionRequest) SetAddsPunctuation(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAddsPunctuation:"), value)
} /* debug [instance_properties/setter]: addsPunctuation */

// An array of phrases that should be recognized, even if they are not in the system vocabulary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/contextualStrings
func (s_ SFSpeechRecognitionRequest) ContextualStrings() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("contextualStrings"))
	return rv
} /* debug [instance_properties/getter]: contextualStrings */

// An array of phrases that should be recognized, even if they are not in the system vocabulary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/contextualStrings
func (s_ SFSpeechRecognitionRequest) SetContextualStrings(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setContextualStrings:"), nsArray)
} /* debug [instance_properties/setter]: contextualStrings */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/customizedLanguageModel
func (s_ SFSpeechRecognitionRequest) CustomizedLanguageModel() ISFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s_.ID, objc.Sel("customizedLanguageModel"))
	return rv
} /* debug [instance_properties/getter]: customizedLanguageModel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/customizedLanguageModel
func (s_ SFSpeechRecognitionRequest) SetCustomizedLanguageModel(value ISFSpeechLanguageModelConfiguration) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizedLanguageModel:"), value)
} /* debug [instance_properties/setter]: customizedLanguageModel */

// An identifier string that you use to describe the type of interaction associated with the speech recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/interactionIdentifier
func (s_ SFSpeechRecognitionRequest) InteractionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("interactionIdentifier"))
	return rv
} /* debug [instance_properties/getter]: interactionIdentifier */

// An identifier string that you use to describe the type of interaction associated with the speech recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/interactionIdentifier
func (s_ SFSpeechRecognitionRequest) SetInteractionIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInteractionIdentifier:"), value)
} /* debug [instance_properties/setter]: interactionIdentifier */

// A Boolean value that determines whether a request must keep its audio data on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/requiresOnDeviceRecognition
func (s_ SFSpeechRecognitionRequest) RequiresOnDeviceRecognition() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("requiresOnDeviceRecognition"))
	return rv
} /* debug [instance_properties/getter]: requiresOnDeviceRecognition */

// A Boolean value that determines whether a request must keep its audio data on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/requiresOnDeviceRecognition
func (s_ SFSpeechRecognitionRequest) SetRequiresOnDeviceRecognition(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRequiresOnDeviceRecognition:"), value)
} /* debug [instance_properties/setter]: requiresOnDeviceRecognition */

// A Boolean value that indicates whether you want intermediate results returned for each utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/shouldReportPartialResults
func (s_ SFSpeechRecognitionRequest) ShouldReportPartialResults() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldReportPartialResults"))
	return rv
} /* debug [instance_properties/getter]: shouldReportPartialResults */

// A Boolean value that indicates whether you want intermediate results returned for each utterance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/shouldReportPartialResults
func (s_ SFSpeechRecognitionRequest) SetShouldReportPartialResults(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldReportPartialResults:"), value)
} /* debug [instance_properties/setter]: shouldReportPartialResults */

// A value that indicates the type of speech recognition being performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/taskHint
func (s_ SFSpeechRecognitionRequest) TaskHint() SFSpeechRecognitionTaskHint {
	rv := objc.Send[SFSpeechRecognitionTaskHint](s_.ID, objc.Sel("taskHint"))
	return rv
} /* debug [instance_properties/getter]: taskHint */

// A value that indicates the type of speech recognition being performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/taskHint
func (s_ SFSpeechRecognitionRequest) SetTaskHint(value SFSpeechRecognitionTaskHint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTaskHint:"), value)
} /* debug [instance_properties/setter]: taskHint */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFSpeechRecognitionRequest */
