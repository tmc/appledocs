// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSpeechLanguageModelConfiguration] class.
var (
	SFSpeechLanguageModelConfigurationClass     _SFSpeechLanguageModelConfigurationClass
	SFSpeechLanguageModelConfigurationClassOnce sync.Once
)

func getSFSpeechLanguageModelConfigurationClass() _SFSpeechLanguageModelConfigurationClass {
	SFSpeechLanguageModelConfigurationClassOnce.Do(func() {
		SFSpeechLanguageModelConfigurationClass = _SFSpeechLanguageModelConfigurationClass{objc.GetClass("SFSpeechLanguageModelConfiguration")}
	})
	return SFSpeechLanguageModelConfigurationClass
}

type _SFSpeechLanguageModelConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [SFSpeechLanguageModelConfiguration] class.
type ISFSpeechLanguageModelConfiguration interface {
	objectivec.IObject
	// properties:
	LanguageModel() objc.IObject /* cross-framework: NSURL */
	Vocabulary() objc.IObject /* cross-framework: NSURL */
	Weight() objc.IObject /* cross-framework: NSNumber */
	CustomizedLanguageModel() ISFSpeechLanguageModelConfiguration
	SetCustomizedLanguageModel(value ISFSpeechLanguageModelConfiguration)
	// methods:
}

// An object describing the location of a custom language model and specialized vocabulary.
//
// Pass this object to to indicate where that method should create the custom language model file, and to or to indicate where the system should find that model to use.


// An object describing the location of a custom language model and specialized vocabulary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration
type SFSpeechLanguageModelConfiguration struct {
	objectivec.Object
}

// SFSpeechLanguageModelConfigurationFrom constructs a [SFSpeechLanguageModelConfiguration] from an unsafe.Pointer.
//
// An object describing the location of a custom language model and specialized vocabulary.
func SFSpeechLanguageModelConfigurationFrom(ptr unsafe.Pointer) SFSpeechLanguageModelConfiguration {
	return SFSpeechLanguageModelConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSpeechLanguageModelConfigurationClass) Alloc() SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSpeechLanguageModelConfigurationClass) New() SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSpeechLanguageModelConfiguration) Init() SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSpeechLanguageModelConfiguration) Autorelease() SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechLanguageModelConfiguration creates a new SFSpeechLanguageModelConfiguration instance.
func NewSFSpeechLanguageModelConfiguration() SFSpeechLanguageModelConfiguration {
	return getSFSpeechLanguageModelConfigurationClass().New()
}



// Creates a configuration with the location of a language model file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:)
func NewSFSpeechLanguageModelConfigurationWithLanguageModel(languageModel objc.IObject /* cross-framework: NSURL */) SFSpeechLanguageModelConfiguration {
	instance := getSFSpeechLanguageModelConfigurationClass().Alloc()
	rv := objc.Send[SFSpeechLanguageModelConfiguration](instance.ID, objc.Sel("initWithLanguageModel:"), languageModel)
	rv.Autorelease()
	return rv
}


// Creates a configuration with the locations of language model and vocabulary files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:vocabulary:)
func NewSFSpeechLanguageModelConfigurationWithLanguageModelVocabulary(languageModel objc.IObject /* cross-framework: NSURL */, vocabulary objc.IObject /* cross-framework: NSURL */) SFSpeechLanguageModelConfiguration {
	instance := getSFSpeechLanguageModelConfigurationClass().Alloc()
	rv := objc.Send[SFSpeechLanguageModelConfiguration](instance.ID, objc.Sel("initWithLanguageModel:vocabulary:"), languageModel, vocabulary)
	rv.Autorelease()
	return rv
}


// Creates a configuration with the locations of language model and vocabulary files, and custom weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:vocabulary:weight:)
func NewSFSpeechLanguageModelConfigurationWithLanguageModelVocabularyWeight(languageModel objc.IObject /* cross-framework: NSURL */, vocabulary objc.IObject /* cross-framework: NSURL */, weight objc.IObject /* cross-framework: NSNumber */) SFSpeechLanguageModelConfiguration {
	instance := getSFSpeechLanguageModelConfigurationClass().Alloc()
	rv := objc.Send[SFSpeechLanguageModelConfiguration](instance.ID, objc.Sel("initWithLanguageModel:vocabulary:weight:"), languageModel, vocabulary, weight)
	rv.Autorelease()
	return rv
}



// The location of a compiled language model file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/languageModel
func (s_ SFSpeechLanguageModelConfiguration) LanguageModel() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("languageModel"))
	return rv
}


// The location of a compiled vocabulary file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/vocabulary
func (s_ SFSpeechLanguageModelConfiguration) Vocabulary() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("vocabulary"))
	return rv
}


// The relative weight of the language model customization. Value must be between 0.0 and 1.0 inclusive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/weight
func (s_ SFSpeechLanguageModelConfiguration) Weight() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("weight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/customizedlanguagemodel
func (s_ SFSpeechLanguageModelConfiguration) CustomizedLanguageModel() ISFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s_.ID, objc.Sel("customizedLanguageModel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/customizedlanguagemodel
func (s_ SFSpeechLanguageModelConfiguration) SetCustomizedLanguageModel(value ISFSpeechLanguageModelConfiguration) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizedLanguageModel:"), value)
}


