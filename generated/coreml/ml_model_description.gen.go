// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelDescription */


/* debug [class_header]: Header for MLModelDescription */
// The class instance for the [ModelDescription] class.
var (
	ModelDescriptionClass     _ModelDescriptionClass
	ModelDescriptionClassOnce sync.Once
)

func getModelDescriptionClass() _ModelDescriptionClass {
	ModelDescriptionClassOnce.Do(func() {
		ModelDescriptionClass = _ModelDescriptionClass{objc.GetClass("MLModelDescription")}
	})
	return ModelDescriptionClass
}

type _ModelDescriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelDescription */
// An interface definition for the [ModelDescription] class.
type IModelDescription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelDescription */
	// properties:
	ClassLabels() []objc.ID
	InputDescriptionsByName() foundation.IDictionary
	IsUpdatable() bool
	Metadata() foundation.IDictionary
	OutputDescriptionsByName() foundation.IDictionary
	ParameterDescriptionsByKey() foundation.IDictionary
	PredictedFeatureName() objc.IObject /* cross-framework: NSString */
	PredictedProbabilitiesName() objc.IObject /* cross-framework: NSString */
	StateDescriptionsByName() foundation.IDictionary
	TrainingInputDescriptionsByName() foundation.IDictionary
	Configuration() IMLModelConfiguration
	SetConfiguration(value IMLModelConfiguration)
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelDescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelDescription */
// Alloc allocates a new instance without initialization.
func (mc _ModelDescriptionClass) Alloc() ModelDescription {
	rv := objc.Send[ModelDescription](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _ModelDescriptionClass) New() ModelDescription {
	rv := objc.Send[ModelDescription](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelDescription) Init() ModelDescription {
	rv := objc.Send[ModelDescription](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelDescription) Autorelease() ModelDescription {
	rv := objc.Send[ModelDescription](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelDescription creates a new ModelDescription instance.
func NewModelDescription() ModelDescription {
	return getModelDescriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelDescription */
// Information about a model, primarily the input and output format for each feature the model expects, and optional metadata.


// Information about a model, primarily the input and output format for each feature the model expects, and optional metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription
type ModelDescription struct {
	objectivec.Object
}

// ModelDescriptionFrom constructs a [ModelDescription] from an unsafe.Pointer.
//
// Information about a model, primarily the input and output format for each feature the model expects, and optional metadata.
func ModelDescriptionFrom(ptr unsafe.Pointer) ModelDescription {
	return ModelDescription{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelDescription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelDescription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelDescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelDescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelDescription */

// An array of labels, which can be either strings or a numbers, for classifier models.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/classLabels
func (m_ ModelDescription) ClassLabels() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("classLabels"))
	return rv
}/* debug [instance_properties/getter]: classLabels */


// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/inputDescriptionsByName
func (m_ ModelDescription) InputDescriptionsByName() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("inputDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: inputDescriptionsByName */


// A Boolean value that indicates whether you can update the model with additional training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/isUpdatable
func (m_ ModelDescription) IsUpdatable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUpdatable"))
	return rv
}/* debug [instance_properties/getter]: isUpdatable */


// A dictionary of the model’s creation information, such as its description, author, version, and license.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/metadata
func (m_ ModelDescription) Metadata() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (m_ ModelDescription) OutputDescriptionsByName() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: outputDescriptionsByName */


// A dictionary of the descriptions for the model’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/parameterDescriptionsByKey
func (m_ ModelDescription) ParameterDescriptionsByKey() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("parameterDescriptionsByKey"))
	return rv
}/* debug [instance_properties/getter]: parameterDescriptionsByKey */


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (m_ ModelDescription) PredictedFeatureName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("predictedFeatureName"))
	return rv
}/* debug [instance_properties/getter]: predictedFeatureName */


// The name of the feature output description for all probabilities of a prediction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedProbabilitiesName
func (m_ ModelDescription) PredictedProbabilitiesName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("predictedProbabilitiesName"))
	return rv
}/* debug [instance_properties/getter]: predictedProbabilitiesName */


// Description of the state features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/stateDescriptionsByName
func (m_ ModelDescription) StateDescriptionsByName() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("stateDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: stateDescriptionsByName */


// A dictionary of the training input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/trainingInputDescriptionsByName
func (m_ ModelDescription) TrainingInputDescriptionsByName() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("trainingInputDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: trainingInputDescriptionsByName */


// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (m_ ModelDescription) Configuration() IMLModelConfiguration {
	rv := objc.Send[ModelConfiguration](m_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (m_ ModelDescription) SetConfiguration(value IMLModelConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m_ ModelDescription) ModelDescription() IMLModelDescription {
	rv := objc.Send[ModelDescription](m_.ID, objc.Sel("modelDescription"))
	return rv
}/* debug [instance_properties/getter]: modelDescription */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m_ ModelDescription) SetModelDescription(value IMLModelDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModelDescription:"), value)
}/* debug [instance_properties/setter]: modelDescription */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelDescription */



