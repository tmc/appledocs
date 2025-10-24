// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Model] class.
var (
	ModelClass     _ModelClass
	ModelClassOnce sync.Once
)

func getModelClass() _ModelClass {
	ModelClassOnce.Do(func() {
		ModelClass = _ModelClass{objc.GetClass("MLModel")}
	})
	return ModelClass
}

type _ModelClass struct {
	class objc.Class
}

// An interface definition for the [Model] class.
type IModel interface {
	objectivec.IObject
	// properties:
	Configuration() IMLModelConfiguration
	ModelDescription() IMLModelDescription
	Metadata() objc.IObject /* cross-framework: ModelMetadataKey */
	SetMetadata(value objc.IObject /* cross-framework: ModelMetadataKey */)
	// methods:
	NewState() IState
	ParameterValueForKeyError(key IMLParameterKey, error_ unsafe.Pointer) objc.ID
	Prediction()
	PredictionFromFeaturesError(input objectivec.IObject, error_ unsafe.Pointer) objc.ID
	PredictionFromFeaturesOptionsError(input objectivec.IObject, options IMLPredictionOptions, error_ unsafe.Pointer) objc.ID
	PredictionFromFeaturesCompletionHandler(input objectivec.IObject, completionHandler unsafe.Pointer)
	PredictionFromFeaturesOptionsCompletionHandler(input objectivec.IObject, options IMLPredictionOptions, completionHandler unsafe.Pointer)
	PredictionFromFeaturesUsingStateOptionsCompletionHandler(inputFeatures objectivec.IObject, state IMLState, options IMLPredictionOptions, completionHandler unsafe.Pointer)
	PredictionsFromBatchOptionsError(inputBatch objectivec.IObject, options IMLPredictionOptions, error_ unsafe.Pointer) objc.ID
	PredictionsFromBatchError(inputBatch objectivec.IObject, error_ unsafe.Pointer) objc.ID
}

// An encapsulation of all the details of your machine learning model.
//
// encapsulates a model’s prediction methods, configuration, and model description. In most cases, you can use Core ML without accessing the class directly. Instead, use the programmer-friendly wrapper class that Xcode automatically generates when you add a model (see ). If your app needs the interface, use the wrapper class’s property. With the interface, you can: Make a prediction with your app’s custom by calling or . Make multiple predictions with your app’s custom by calling or . Inspect your model’s and instances through . If your app downloads and compiles a model on the user’s device, you must use the class directly to make predictions. See .


// An encapsulation of all the details of your machine learning model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel
type Model struct {
	objectivec.Object
}

// ModelFrom constructs a [Model] from an unsafe.Pointer.
//
// An encapsulation of all the details of your machine learning model.
func ModelFrom(ptr unsafe.Pointer) Model {
	return Model{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelClass) Alloc() Model {
	rv := objc.Send[Model](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelClass) New() Model {
	rv := objc.Send[Model](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Model) Init() Model {
	rv := objc.Send[Model](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Model) Autorelease() Model {
	rv := objc.Send[Model](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModel creates a new Model instance.
func NewModel() Model {
	return getModelClass().New()
}



// Creates a Core ML model instance from a compiled model file and a custom configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:configuration:)
func NewModelWithContentsOfURLConfigurationError(url objc.IObject /* cross-framework: NSURL */, configuration IMLModelConfiguration, error_ unsafe.Pointer) Model {
	rv := objc.Send[Model](objc.ID(getModelClass().class), objc.Sel("modelWithContentsOfURL:configuration:error:"), url, configuration, error_)
	return rv
}


// Creates a Core ML model instance from a compiled model file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:)
func NewModelWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, error_ unsafe.Pointer) Model {
	rv := objc.Send[Model](objc.ID(getModelClass().class), objc.Sel("modelWithContentsOfURL:error:"), url, error_)
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)
func (mc _ModelClass) CompileModel() {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("compileModel"))
}


// Compile a model for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)-3nea
func (mc _ModelClass) CompileModelAtURLCompletionHandler(modelURL objc.IObject /* cross-framework: NSURL */, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("compileModelAtURL:completionHandler:"), modelURL, handler)
}


// Compiles a model on the device to update the model in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)-6442s
func (mc _ModelClass) CompileModelAtURLError(modelURL objc.IObject /* cross-framework: NSURL */, error_ unsafe.Pointer) objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](objc.ID(mc.class), objc.Sel("compileModelAtURL:error:"), modelURL, error_)
	return rv
}


// Creates a Core ML model instance from a compiled model file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:)
func (mc _ModelClass) ModelWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("modelWithContentsOfURL:error:"), url, error_)
	return rv
}


// Creates a Core ML model instance from a compiled model file and a custom configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:configuration:)
func (mc _ModelClass) ModelWithContentsOfURLConfigurationError(url objc.IObject /* cross-framework: NSURL */, configuration IMLModelConfiguration, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("modelWithContentsOfURL:configuration:error:"), url, configuration, error_)
	return rv
}


// Construct a model asynchronously from a compiled model asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/load(_:configuration:completionHandler:)
func (mc _ModelClass) LoadModelAssetConfigurationCompletionHandler(asset IMLModelAsset, configuration IMLModelConfiguration, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("loadModelAsset:configuration:completionHandler:"), asset, configuration, handler)
}


// Creates a Core ML model instance asynchronously from a compiled model file, a custom configuration, and a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/loadContentsOfURL:configuration:completionHandler:
func (mc _ModelClass) LoadContentsOfURLConfigurationCompletionHandler(url objc.IObject /* cross-framework: NSURL */, configuration IMLModelConfiguration, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("loadContentsOfURL:configuration:completionHandler:"), url, configuration, handler)
}


// The list of available compute devices that the model’s prediction can use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/availableComputeDevices-42uzt
func (mc _ModelClass) AvailableComputeDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](objc.ID(mc.class), objc.Sel("availableComputeDevices"))
	return rv
}

// Creates a new state object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/newState
func (m_ Model) NewState() IState {
	rv := objc.Send[State](m_.ID, objc.Sel("newState"))
	return rv
}


// Returns a model parameter value for a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/parameterValue(for:)
func (m_ Model) ParameterValueForKeyError(key IMLParameterKey, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("parameterValueForKey:error:"), key, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:)
func (m_ Model) Prediction() {
	objc.Send[objc.ID](m_.ID, objc.Sel("prediction"))
}


// Generates a prediction from the feature values within the input feature provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:)-9y2aa
func (m_ Model) PredictionFromFeaturesError(input objectivec.IObject, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("predictionFromFeatures:error:"), input, error_)
	return rv
}


// Generates a prediction from the feature values within the input feature provider using the prediction options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:options:)-81mr6
func (m_ Model) PredictionFromFeaturesOptionsError(input objectivec.IObject, options IMLPredictionOptions, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("predictionFromFeatures:options:error:"), input, options, error_)
	return rv
}


// Generates a prediction asynchronously from the feature values within the input feature provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:completionHandler:
func (m_ Model) PredictionFromFeaturesCompletionHandler(input objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("predictionFromFeatures:completionHandler:"), input, completionHandler)
}


// Generates a prediction asynchronously from the feature values within the input feature provider using the prediction options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:options:completionHandler:
func (m_ Model) PredictionFromFeaturesOptionsCompletionHandler(input objectivec.IObject, options IMLPredictionOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("predictionFromFeatures:options:completionHandler:"), input, options, completionHandler)
}


// Run a stateful prediction asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:usingState:options:completionHandler:
func (m_ Model) PredictionFromFeaturesUsingStateOptionsCompletionHandler(inputFeatures objectivec.IObject, state IMLState, options IMLPredictionOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("predictionFromFeatures:usingState:options:completionHandler:"), inputFeatures, state, options, completionHandler)
}


// Generates a prediction for each input feature provider within the batch provider using the prediction options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/predictions(from:options:)
func (m_ Model) PredictionsFromBatchOptionsError(inputBatch objectivec.IObject, options IMLPredictionOptions, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("predictionsFromBatch:options:error:"), inputBatch, options, error_)
	return rv
}


// Generates predictions for each input feature provider within the batch provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/predictions(fromBatch:)
func (m_ Model) PredictionsFromBatchError(inputBatch objectivec.IObject, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("predictionsFromBatch:error:"), inputBatch, error_)
	return rv
}


// The list of available compute devices that the model’s prediction can use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/availableComputeDevices-42uzt
func (m_ Model) AvailableComputeDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("availableComputeDevices"))
	return rv
}


// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/configuration
func (m_ Model) Configuration() IMLModelConfiguration {
	rv := objc.Send[ModelConfiguration](m_.ID, objc.Sel("configuration"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (m_ Model) ModelDescription() IMLModelDescription {
	rv := objc.Send[ModelDescription](m_.ID, objc.Sel("modelDescription"))
	return rv
}


// A dictionary of the model’s creation information, such as its description, author, version, and license.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/metadata
func (m_ Model) Metadata() objc.IObject /* cross-framework: ModelMetadataKey */ {
	rv := objc.Send[ModelMetadataKey](m_.ID, objc.Sel("metadata"))
	return rv
}


// A dictionary of the model’s creation information, such as its description, author, version, and license.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/metadata
func (m_ Model) SetMetadata(value objc.IObject /* cross-framework: ModelMetadataKey */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadata:"), value)
}


