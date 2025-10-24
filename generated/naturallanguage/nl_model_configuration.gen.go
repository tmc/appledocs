// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NLModelConfiguration */


/* debug [class_header]: Header for NLModelConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelConfiguration */
// An interface definition for the [ModelConfiguration] class.
type IModelConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelConfiguration */
	// properties:
	Language() Language /* typedef */
	Revision() uint
	Type() ModelType
	Configuration() INLModelConfiguration
	SetConfiguration(value INLModelConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelConfiguration */
// Alloc allocates a new instance without initialization.
func (mc _ModelConfigurationClass) Alloc() ModelConfiguration {
	rv := objc.Send[ModelConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelConfiguration */

// Returns the current Natural Language framework version in the OS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/currentRevision(for:)
func (mc _ModelConfigurationClass) CurrentRevisionForType(type_ ModelType) uint {
	rv := objc.Send[uint](objc.ID(mc.class), objc.Sel("currentRevisionForType:"), type_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentRevisionForType) */


// Returns the versions of the Natural Language framework the OS supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/supportedRevisions(for:)
func (mc _ModelConfigurationClass) SupportedRevisionsForType(type_ ModelType) foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](objc.ID(mc.class), objc.Sel("supportedRevisionsForType:"), type_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedRevisionsForType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelConfiguration */

// The language the model supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/language
func (m_ ModelConfiguration) Language() Language /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("language"))
	return rv
}/* debug [instance_properties/getter]: language */


// The version of the Natural Language framework that trained the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/revision
func (m_ ModelConfiguration) Revision() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("revision"))
	return rv
}/* debug [instance_properties/getter]: revision */


// The natural language model type of the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/type
func (m_ ModelConfiguration) Type() ModelType {
	rv := objc.Send[ModelType](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// A configuration describing the natural language model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/naturallanguage/nlmodel/configuration
func (m_ ModelConfiguration) Configuration() INLModelConfiguration {
	rv := objc.Send[ModelConfiguration](m_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// A configuration describing the natural language model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/naturallanguage/nlmodel/configuration
func (m_ ModelConfiguration) SetConfiguration(value INLModelConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NLModelConfiguration */






