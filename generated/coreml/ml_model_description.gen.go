// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ModelDescription] class.
type IModelDescription interface {
	objectivec.IObject
	

	// properties:
	ClassLabels() []objc.ID
	InputDescriptionsByName() foundation.IDictionary
	IsUpdatable() bool
	Metadata() foundation.IDictionary
	OutputDescriptionsByName() foundation.IDictionary
	ParameterDescriptionsByKey() foundation.IDictionary
	PredictedFeatureName() foundation.foundation.INSString
	PredictedProbabilitiesName() foundation.foundation.INSString
	StateDescriptionsByName() foundation.IDictionary
	TrainingInputDescriptionsByName() foundation.IDictionary
	Configuration() IMLModelConfiguration
	SetConfiguration(value IMLModelConfiguration)
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)


	

	// methods:


}





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

























// An array of labels, which can be either strings or a numbers, for classifier models.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/classLabels
func (m_ ModelDescription) ClassLabels() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("classLabels"))
	return rv
}


// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/inputDescriptionsByName
func (m_ ModelDescription) InputDescriptionsByName() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("inputDescriptionsByName"))
	return rv
}


// A Boolean value that indicates whether you can update the model with additional training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/isUpdatable
func (m_ ModelDescription) IsUpdatable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUpdatable"))
	return rv
}


// A dictionary of the model’s creation information, such as its description, author, version, and license.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/metadata
func (m_ ModelDescription) Metadata() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("metadata"))
	return rv
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (m_ ModelDescription) OutputDescriptionsByName() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}


// A dictionary of the descriptions for the model’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/parameterDescriptionsByKey
func (m_ ModelDescription) ParameterDescriptionsByKey() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("parameterDescriptionsByKey"))
	return rv
}


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (m_ ModelDescription) PredictedFeatureName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("predictedFeatureName"))
	return rv
}


// The name of the feature output description for all probabilities of a prediction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedProbabilitiesName
func (m_ ModelDescription) PredictedProbabilitiesName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("predictedProbabilitiesName"))
	return rv
}


// Description of the state features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/stateDescriptionsByName
func (m_ ModelDescription) StateDescriptionsByName() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("stateDescriptionsByName"))
	return rv
}


// A dictionary of the training input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/trainingInputDescriptionsByName
func (m_ ModelDescription) TrainingInputDescriptionsByName() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("trainingInputDescriptionsByName"))
	return rv
}


// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (m_ ModelDescription) Configuration() IMLModelConfiguration {
	rv := objc.Send[ModelConfiguration](m_.ID, objc.Sel("configuration"))
	return rv
}


// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (m_ ModelDescription) SetConfiguration(value IMLModelConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConfiguration:"), value)
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m_ ModelDescription) ModelDescription() IMLModelDescription {
	rv := objc.Send[ModelDescription](m_.ID, objc.Sel("modelDescription"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m_ ModelDescription) SetModelDescription(value IMLModelDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModelDescription:"), value)
}








