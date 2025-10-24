// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSpeechLanguageModelConfiguration */


/* debug [class_header]: Header for SFSpeechLanguageModelConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSpeechLanguageModelConfiguration */
// An interface definition for the [SFSpeechLanguageModelConfiguration] class.
type ISFSpeechLanguageModelConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSpeechLanguageModelConfiguration */
	// properties:
	LanguageModel() objc.IObject /* cross-framework: NSURL */
	Vocabulary() objc.IObject /* cross-framework: NSURL */
	Weight() objc.IObject /* cross-framework: NSNumber */
	CustomizedLanguageModel() ISFSpeechLanguageModelConfiguration
	SetCustomizedLanguageModel(value ISFSpeechLanguageModelConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSpeechLanguageModelConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSpeechLanguageModelConfiguration */
// Alloc allocates a new instance without initialization.
func (sc _SFSpeechLanguageModelConfigurationClass) Alloc() SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSpeechLanguageModelConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSpeechLanguageModelConfiguration */

// Creates a configuration with the location of a language model file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:)
func NewSFSpeechLanguageModelConfigurationWithLanguageModel(languageModel objc.IObject /* cross-framework: NSURL */) SFSpeechLanguageModelConfiguration {
	instance := getSFSpeechLanguageModelConfigurationClass().Alloc()
	rv := objc.Send[SFSpeechLanguageModelConfiguration](instance.ID, objc.Sel("initWithLanguageModel:"), languageModel)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSFSpeechLanguageModelConfigurationWithLanguageModel */


// Creates a configuration with the locations of language model and vocabulary files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:vocabulary:)
func NewSFSpeechLanguageModelConfigurationWithLanguageModelVocabulary(languageModel objc.IObject /* cross-framework: NSURL */, vocabulary objc.IObject /* cross-framework: NSURL */) SFSpeechLanguageModelConfiguration {
	instance := getSFSpeechLanguageModelConfigurationClass().Alloc()
	rv := objc.Send[SFSpeechLanguageModelConfiguration](instance.ID, objc.Sel("initWithLanguageModel:vocabulary:"), languageModel, vocabulary)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSFSpeechLanguageModelConfigurationWithLanguageModelVocabulary */


// Creates a configuration with the locations of language model and vocabulary files, and custom weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:vocabulary:weight:)
func NewSFSpeechLanguageModelConfigurationWithLanguageModelVocabularyWeight(languageModel objc.IObject /* cross-framework: NSURL */, vocabulary objc.IObject /* cross-framework: NSURL */, weight objc.IObject /* cross-framework: NSNumber */) SFSpeechLanguageModelConfiguration {
	instance := getSFSpeechLanguageModelConfigurationClass().Alloc()
	rv := objc.Send[SFSpeechLanguageModelConfiguration](instance.ID, objc.Sel("initWithLanguageModel:vocabulary:weight:"), languageModel, vocabulary, weight)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSFSpeechLanguageModelConfigurationWithLanguageModelVocabularyWeight */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSpeechLanguageModelConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSpeechLanguageModelConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSpeechLanguageModelConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSpeechLanguageModelConfiguration */

// The location of a compiled language model file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/languageModel
func (s_ SFSpeechLanguageModelConfiguration) LanguageModel() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("languageModel"))
	return rv
}/* debug [instance_properties/getter]: languageModel */


// The location of a compiled vocabulary file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/vocabulary
func (s_ SFSpeechLanguageModelConfiguration) Vocabulary() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("vocabulary"))
	return rv
}/* debug [instance_properties/getter]: vocabulary */


// The relative weight of the language model customization. Value must be between 0.0 and 1.0 inclusive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/weight
func (s_ SFSpeechLanguageModelConfiguration) Weight() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/customizedlanguagemodel
func (s_ SFSpeechLanguageModelConfiguration) CustomizedLanguageModel() ISFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s_.ID, objc.Sel("customizedLanguageModel"))
	return rv
}/* debug [instance_properties/getter]: customizedLanguageModel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionrequest/customizedlanguagemodel
func (s_ SFSpeechLanguageModelConfiguration) SetCustomizedLanguageModel(value ISFSpeechLanguageModelConfiguration) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizedLanguageModel:"), value)
}/* debug [instance_properties/setter]: customizedLanguageModel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSpeechLanguageModelConfiguration */


