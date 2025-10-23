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
	InputDescriptionsByName() unsafe.Pointer
	Metadata() unsafe.Pointer
	OutputDescriptionsByName() unsafe.Pointer
	PredictedFeatureName() string
	PredictedProbabilitiesName() string
	StateDescriptionsByName() unsafe.Pointer
	Configuration() MLModelConfiguration
	SetConfiguration(value IMLModelConfiguration)
	ModelDescription() MLModelDescription
	SetModelDescription(value IMLModelDescription)
	ClassLabels() unsafe.Pointer
	SetClassLabels(value unsafe.Pointer)
	IsUpdatable() bool
	SetIsUpdatable(value bool)
	ParameterDescriptionsByKey() unsafe.Pointer
	SetParameterDescriptionsByKey(value unsafe.Pointer)
	TrainingInputDescriptionsByName() MLFeatureDescription
	SetTrainingInputDescriptionsByName(value IMLFeatureDescription)
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

// Alloc allocates a new instance without initialization.
func (mc _ModelDescriptionClass) Alloc() ModelDescription {
	rv := objc.Send[ModelDescription](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/inputDescriptionsByName
func (m_ ModelDescription) InputDescriptionsByName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("inputDescriptionsByName"))
	return rv
}


// A dictionary of the model’s creation information, such as its description, author, version, and license.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/metadata
func (m_ ModelDescription) Metadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("metadata"))
	return rv
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (m_ ModelDescription) OutputDescriptionsByName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (m_ ModelDescription) PredictedFeatureName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("predictedFeatureName"))
	return rv
}


// The name of the feature output description for all probabilities of a prediction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedProbabilitiesName
func (m_ ModelDescription) PredictedProbabilitiesName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("predictedProbabilitiesName"))
	return rv
}


// Description of the state features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/stateDescriptionsByName
func (m_ ModelDescription) StateDescriptionsByName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("stateDescriptionsByName"))
	return rv
}


// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (m_ ModelDescription) Configuration() MLModelConfiguration {
	rv := objc.Send[MLModelConfiguration](m_.ID, objc.Sel("configuration"))
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
func (m_ ModelDescription) ModelDescription() MLModelDescription {
	rv := objc.Send[MLModelDescription](m_.ID, objc.Sel("modelDescription"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m_ ModelDescription) SetModelDescription(value IMLModelDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModelDescription:"), value)
}


// An array of labels, which can be either strings or a numbers, for classifier models.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/classlabels
func (m_ ModelDescription) ClassLabels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("classLabels"))
	return rv
}


// An array of labels, which can be either strings or a numbers, for classifier models.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/classlabels
func (m_ ModelDescription) SetClassLabels(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClassLabels:"), value)
}


// A Boolean value that indicates whether you can update the model with additional training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/isupdatable
func (m_ ModelDescription) IsUpdatable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUpdatable"))
	return rv
}


// A Boolean value that indicates whether you can update the model with additional training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/isupdatable
func (m_ ModelDescription) SetIsUpdatable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsUpdatable:"), value)
}


// A dictionary of the descriptions for the model’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/parameterdescriptionsbykey
func (m_ ModelDescription) ParameterDescriptionsByKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("parameterDescriptionsByKey"))
	return rv
}


// A dictionary of the descriptions for the model’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/parameterdescriptionsbykey
func (m_ ModelDescription) SetParameterDescriptionsByKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParameterDescriptionsByKey:"), value)
}


// A dictionary of the training input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/traininginputdescriptionsbyname
func (m_ ModelDescription) TrainingInputDescriptionsByName() MLFeatureDescription {
	rv := objc.Send[MLFeatureDescription](m_.ID, objc.Sel("trainingInputDescriptionsByName"))
	return rv
}


// A dictionary of the training input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/traininginputdescriptionsbyname
func (m_ ModelDescription) SetTrainingInputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrainingInputDescriptionsByName:"), value)
}



