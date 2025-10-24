// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModel */


/* debug [class_header]: Header for MLModel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Model */
// An interface definition for the [Model] class.
type IModel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Model */
	// properties:
	Configuration() IMLModelConfiguration
	ModelDescription() IMLModelDescription
	Metadata() ModelMetadataKey /* typedef */
	SetMetadata(value ModelMetadataKey /* typedef */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Model */
	// methods:
	NewState() IState
	ParameterValueForKeyError(key IMLParameterKey, error_ objectivec.IObject) objc.ID
	Prediction()
	PredictionFromFeaturesError(input unsafe.Pointer, error_ objectivec.IObject) unsafe.Pointer
	PredictionFromFeaturesOptionsError(input unsafe.Pointer, options IMLPredictionOptions, error_ objectivec.IObject) unsafe.Pointer
	PredictionFromFeaturesUsingStateError(inputFeatures unsafe.Pointer, state IMLState, error_ objectivec.IObject) unsafe.Pointer
	PredictionFromFeaturesUsingStateOptionsError(inputFeatures unsafe.Pointer, state IMLState, options IMLPredictionOptions, error_ objectivec.IObject) unsafe.Pointer
	PredictionFromFeaturesCompletionHandler(input unsafe.Pointer, completionHandler unsafe.Pointer)
	PredictionFromFeaturesOptionsCompletionHandler(input unsafe.Pointer, options IMLPredictionOptions, completionHandler unsafe.Pointer)
	PredictionFromFeaturesUsingStateOptionsCompletionHandler(inputFeatures unsafe.Pointer, state IMLState, options IMLPredictionOptions, completionHandler unsafe.Pointer)
	PredictionsFromBatchOptionsError(inputBatch unsafe.Pointer, options IMLPredictionOptions, error_ objectivec.IObject) unsafe.Pointer
	PredictionsFromBatchError(inputBatch unsafe.Pointer, error_ objectivec.IObject) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Model */
// Alloc allocates a new instance without initialization.
func (mc _ModelClass) Alloc() Model {
	rv := objc.Send[Model](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Model */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Model */

// Creates a Core ML model instance from a compiled model file and a custom configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:configuration:)
func NewModelWithContentsOfURLConfigurationError(url objc.IObject /* cross-framework: NSURL */, configuration IMLModelConfiguration, error_ objectivec.IObject) Model {
	rv := objc.Send[Model](objc.ID(getModelClass().class), objc.Sel("modelWithContentsOfURL:configuration:error:"), url, configuration, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewModelWithContentsOfURLConfigurationError */


// Creates a Core ML model instance from a compiled model file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:)
func NewModelWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) Model {
	rv := objc.Send[Model](objc.ID(getModelClass().class), objc.Sel("modelWithContentsOfURL:error:"), url, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewModelWithContentsOfURLError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Model */

// Compile a model for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)-3nea
func (mc _ModelClass) CompileModelAtURLCompletionHandler(modelURL objc.IObject /* cross-framework: NSURL */, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("compileModelAtURL:completionHandler:"), modelURL, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompileModelAtURLCompletionHandler) */


// Compiles a model on the device to update the model in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)-6442s
func (mc _ModelClass) CompileModelAtURLError(modelURL objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) foundation.URL {
	rv := objc.Send[foundation.URL](objc.ID(mc.class), objc.Sel("compileModelAtURL:error:"), modelURL, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompileModelAtURLError) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)
func (mc _ModelClass) CompileModel() {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("compileModel"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompileModel) */


// Creates a Core ML model instance from a compiled model file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:)
func (mc _ModelClass) ModelWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("modelWithContentsOfURL:error:"), url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ModelWithContentsOfURLError) */


// Creates a Core ML model instance from a compiled model file and a custom configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:configuration:)
func (mc _ModelClass) ModelWithContentsOfURLConfigurationError(url objc.IObject /* cross-framework: NSURL */, configuration IMLModelConfiguration, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("modelWithContentsOfURL:configuration:error:"), url, configuration, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ModelWithContentsOfURLConfigurationError) */


// Construct a model asynchronously from a compiled model asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/load(_:configuration:completionHandler:)
func (mc _ModelClass) LoadModelAssetConfigurationCompletionHandler(asset IMLModelAsset, configuration IMLModelConfiguration, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("loadModelAsset:configuration:completionHandler:"), asset, configuration, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadModelAssetConfigurationCompletionHandler) */


// Creates a Core ML model instance asynchronously from a compiled model file, a custom configuration, and a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/loadContentsOfURL:configuration:completionHandler:
func (mc _ModelClass) LoadContentsOfURLConfigurationCompletionHandler(url objc.IObject /* cross-framework: NSURL */, configuration IMLModelConfiguration, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("loadContentsOfURL:configuration:completionHandler:"), url, configuration, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadContentsOfURLConfigurationCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Model */

// The list of available compute devices that the model’s prediction can use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/availableComputeDevices-42uzt
func (mc _ModelClass) AvailableComputeDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](objc.ID(mc.class), objc.Sel("availableComputeDevices"))
	return rv
}/* debug [class_properties_class/property]: availableComputeDevices */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Model */

// Creates a new state object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/newState
func (m_ Model) NewState() IState {
	rv := objc.Send[State](m_.ID, objc.Sel("newState"))
	return rv
}/* debug [instance_methods/method]: NewState */


// Returns a model parameter value for a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/parameterValue(for:)
func (m_ Model) ParameterValueForKeyError(key IMLParameterKey, error_ objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("parameterValueForKey:error:"), key, error_)
	return rv
}/* debug [instance_methods/method]: ParameterValueForKeyError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:)
func (m_ Model) Prediction() {
	objc.Send[objc.ID](m_.ID, objc.Sel("prediction"))
}/* debug [instance_methods/method]: Prediction */


// Generates a prediction from the feature values within the input feature provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:)-9y2aa
func (m_ Model) PredictionFromFeaturesError(input unsafe.Pointer, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("predictionFromFeatures:error:"), input, error_)
	return rv
}/* debug [instance_methods/method]: PredictionFromFeaturesError */


// Generates a prediction from the feature values within the input feature provider using the prediction options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:options:)-81mr6
func (m_ Model) PredictionFromFeaturesOptionsError(input unsafe.Pointer, options IMLPredictionOptions, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("predictionFromFeatures:options:error:"), input, options, error_)
	return rv
}/* debug [instance_methods/method]: PredictionFromFeaturesOptionsError */


// Run a stateful prediction synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:using:)-97bu1
func (m_ Model) PredictionFromFeaturesUsingStateError(inputFeatures unsafe.Pointer, state IMLState, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("predictionFromFeatures:usingState:error:"), inputFeatures, state, error_)
	return rv
}/* debug [instance_methods/method]: PredictionFromFeaturesUsingStateError */


// Run a stateful prediction synchronously with options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:using:options:)-v4wp
func (m_ Model) PredictionFromFeaturesUsingStateOptionsError(inputFeatures unsafe.Pointer, state IMLState, options IMLPredictionOptions, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("predictionFromFeatures:usingState:options:error:"), inputFeatures, state, options, error_)
	return rv
}/* debug [instance_methods/method]: PredictionFromFeaturesUsingStateOptionsError */


// Generates a prediction asynchronously from the feature values within the input feature provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:completionHandler:
func (m_ Model) PredictionFromFeaturesCompletionHandler(input unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("predictionFromFeatures:completionHandler:"), input, completionHandler)
}/* debug [instance_methods/method]: PredictionFromFeaturesCompletionHandler */


// Generates a prediction asynchronously from the feature values within the input feature provider using the prediction options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:options:completionHandler:
func (m_ Model) PredictionFromFeaturesOptionsCompletionHandler(input unsafe.Pointer, options IMLPredictionOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("predictionFromFeatures:options:completionHandler:"), input, options, completionHandler)
}/* debug [instance_methods/method]: PredictionFromFeaturesOptionsCompletionHandler */


// Run a stateful prediction asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:usingState:options:completionHandler:
func (m_ Model) PredictionFromFeaturesUsingStateOptionsCompletionHandler(inputFeatures unsafe.Pointer, state IMLState, options IMLPredictionOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("predictionFromFeatures:usingState:options:completionHandler:"), inputFeatures, state, options, completionHandler)
}/* debug [instance_methods/method]: PredictionFromFeaturesUsingStateOptionsCompletionHandler */


// Generates a prediction for each input feature provider within the batch provider using the prediction options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/predictions(from:options:)
func (m_ Model) PredictionsFromBatchOptionsError(inputBatch unsafe.Pointer, options IMLPredictionOptions, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("predictionsFromBatch:options:error:"), inputBatch, options, error_)
	return rv
}/* debug [instance_methods/method]: PredictionsFromBatchOptionsError */


// Generates predictions for each input feature provider within the batch provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/predictions(fromBatch:)
func (m_ Model) PredictionsFromBatchError(inputBatch unsafe.Pointer, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("predictionsFromBatch:error:"), inputBatch, error_)
	return rv
}/* debug [instance_methods/method]: PredictionsFromBatchError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Model */

// The list of available compute devices that the model’s prediction can use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/availableComputeDevices-42uzt
func (m_ Model) AvailableComputeDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("availableComputeDevices"))
	return rv
}/* debug [instance_properties/getter]: availableComputeDevices */


// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/configuration
func (m_ Model) Configuration() IMLModelConfiguration {
	rv := objc.Send[ModelConfiguration](m_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (m_ Model) ModelDescription() IMLModelDescription {
	rv := objc.Send[ModelDescription](m_.ID, objc.Sel("modelDescription"))
	return rv
}/* debug [instance_properties/getter]: modelDescription */


// A dictionary of the model’s creation information, such as its description, author, version, and license.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/metadata
func (m_ Model) Metadata() ModelMetadataKey /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// A dictionary of the model’s creation information, such as its description, author, version, and license.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/metadata
func (m_ Model) SetMetadata(value ModelMetadataKey /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadata:"), value)
}/* debug [instance_properties/setter]: metadata */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModel */


