// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelConfiguration] class.
var (
	ModelConfigurationClass     _ModelConfigurationClass
	ModelConfigurationClassOnce sync.Once
)

func getModelConfigurationClass() _ModelConfigurationClass {
	ModelConfigurationClassOnce.Do(func() {
		ModelConfigurationClass = _ModelConfigurationClass{objc.GetClass("NLModelConfiguration")}
	})
	return ModelConfigurationClass
}

type _ModelConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [ModelConfiguration] class.
type IModelConfiguration interface {
	objectivec.IObject
	// properties:
	Language() objc.IObject /* cross-framework: Language */
	Revision() uint
	Type() ModelType
	Configuration() INLModelConfiguration
	SetConfiguration(value INLModelConfiguration)
	// methods:
}

// The configuration parameters of a natural language model.


// The configuration parameters of a natural language model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration
type ModelConfiguration struct {
	objectivec.Object
}

// ModelConfigurationFrom constructs a [ModelConfiguration] from an unsafe.Pointer.
//
// The configuration parameters of a natural language model.
func ModelConfigurationFrom(ptr unsafe.Pointer) ModelConfiguration {
	return ModelConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelConfigurationClass) Alloc() ModelConfiguration {
	rv := objc.Send[ModelConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelConfigurationClass) New() ModelConfiguration {
	rv := objc.Send[ModelConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelConfiguration) Init() ModelConfiguration {
	rv := objc.Send[ModelConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelConfiguration) Autorelease() ModelConfiguration {
	rv := objc.Send[ModelConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelConfiguration creates a new ModelConfiguration instance.
func NewModelConfiguration() ModelConfiguration {
	return getModelConfigurationClass().New()
}



// Returns the current Natural Language framework version in the OS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/currentRevision(for:)
func (mc _ModelConfigurationClass) CurrentRevisionForType(type_ ModelType) uint {
	rv := objc.Send[uint](objc.ID(mc.class), objc.Sel("currentRevisionForType:"), type_)
	return rv
}


// Returns the versions of the Natural Language framework the OS supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/supportedRevisions(for:)
func (mc _ModelConfigurationClass) SupportedRevisionsForType(type_ ModelType) objc.IObject /* cross-framework: IndexSet */ {
	rv := objc.Send[foundation.IndexSet](objc.ID(mc.class), objc.Sel("supportedRevisionsForType:"), type_)
	return rv
}


// The language the model supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/language
func (m_ ModelConfiguration) Language() objc.IObject /* cross-framework: Language */ {
	rv := objc.Send[Language](m_.ID, objc.Sel("language"))
	return rv
}


// The version of the Natural Language framework that trained the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/revision
func (m_ ModelConfiguration) Revision() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("revision"))
	return rv
}


// The natural language model type of the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/type
func (m_ ModelConfiguration) Type() ModelType {
	rv := objc.Send[ModelType](m_.ID, objc.Sel("type"))
	return rv
}


// A configuration describing the natural language model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/naturallanguage/nlmodel/configuration
func (m_ ModelConfiguration) Configuration() INLModelConfiguration {
	rv := objc.Send[ModelConfiguration](m_.ID, objc.Sel("configuration"))
	return rv
}


// A configuration describing the natural language model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/naturallanguage/nlmodel/configuration
func (m_ ModelConfiguration) SetConfiguration(value INLModelConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConfiguration:"), value)
}



