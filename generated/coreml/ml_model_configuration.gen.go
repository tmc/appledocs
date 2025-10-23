// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelConfiguration] class.
var (
	ModelConfigurationClass     _ModelConfigurationClass
	ModelConfigurationClassOnce sync.Once
)

func getModelConfigurationClass() _ModelConfigurationClass {
	ModelConfigurationClassOnce.Do(func() {
		ModelConfigurationClass = _ModelConfigurationClass{objc.GetClass("MLModelConfiguration")}
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
	AllowLowPrecisionAccumulationOnGPU() bool /* primitive/slice/pointer. */
	SetAllowLowPrecisionAccumulationOnGPU(value bool /* primitive/slice/pointer. */)
	ComputeUnits() ComputeUnits /* not a class type */
	SetComputeUnits(value ComputeUnits /* not a class type */)
	FunctionName() string /* primitive/slice/pointer. */
	SetFunctionName(value string /* primitive/slice/pointer. */)
	ModelDisplayName() string /* primitive/slice/pointer. */
	SetModelDisplayName(value string /* primitive/slice/pointer. */)
	OptimizationHints() IOptimizationHints
	SetOptimizationHints(value IOptimizationHints)
	Parameters() IMLParameterKey
	SetParameters(value IMLParameterKey)
	PreferredMetalDevice() Device /* not a class type */
	SetPreferredMetalDevice(value Device /* not a class type */)
	// methods:
}

// The settings for creating or updating a machine learning model.
//
// Use a model configuration to: Set or override model parameters. Designate which device the model uses to make predictions, such as a GPU. Restrict the model to use a specific computational device category, such as a CPU. You typically use a model configuration instance to configure an instance as you create it with or create an . See . Configure your model parameters by setting values for each relevant in the property.


// The settings for creating or updating a machine learning model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration
type ModelConfiguration struct {
	objectivec.Object
}

// ModelConfigurationFrom constructs a [ModelConfiguration] from an unsafe.Pointer.
//
// The settings for creating or updating a machine learning model.
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



// A Boolean value that determines whether to allow low-precision accumulation on a GPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/allowlowprecisionaccumulationongpu
func (m_ ModelConfiguration) AllowLowPrecisionAccumulationOnGPU() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowLowPrecisionAccumulationOnGPU"))
	return rv
}


// A Boolean value that determines whether to allow low-precision accumulation on a GPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/allowlowprecisionaccumulationongpu
func (m_ ModelConfiguration) SetAllowLowPrecisionAccumulationOnGPU(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowLowPrecisionAccumulationOnGPU:"), value)
}


// The processing unit or units the model uses to make predictions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/computeunits
func (m_ ModelConfiguration) ComputeUnits() ComputeUnits /* not a class type */ {
	rv := objc.Send[ComputeUnits](m_.ID, objc.Sel("computeUnits"))
	return rv
}


// The processing unit or units the model uses to make predictions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/computeunits
func (m_ ModelConfiguration) SetComputeUnits(value ComputeUnits /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setComputeUnits:"), value)
}


// Function name that
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/functionname
func (m_ ModelConfiguration) FunctionName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("functionName"))
	return rv
}


// Function name that
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/functionname
func (m_ ModelConfiguration) SetFunctionName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFunctionName:"), objc.String(value))
}


// A human readable name of a model for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/modeldisplayname
func (m_ ModelConfiguration) ModelDisplayName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("modelDisplayName"))
	return rv
}


// A human readable name of a model for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/modeldisplayname
func (m_ ModelConfiguration) SetModelDisplayName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModelDisplayName:"), objc.String(value))
}


// A group of hints for CoreML to optimize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/optimizationhints-1oq0g
func (m_ ModelConfiguration) OptimizationHints() IOptimizationHints {
	rv := objc.Send[OptimizationHints](m_.ID, objc.Sel("optimizationHints"))
	return rv
}


// A group of hints for CoreML to optimize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/optimizationhints-1oq0g
func (m_ ModelConfiguration) SetOptimizationHints(value IOptimizationHints) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptimizationHints:"), value)
}


// A dictionary of configuration settings your app can override when loading a model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/parameters
func (m_ ModelConfiguration) Parameters() IMLParameterKey {
	rv := objc.Send[ParameterKey](m_.ID, objc.Sel("parameters"))
	return rv
}


// A dictionary of configuration settings your app can override when loading a model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/parameters
func (m_ ModelConfiguration) SetParameters(value IMLParameterKey) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParameters:"), value)
}


// The metal device you prefer this model use to make predictions (inference) and update the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/preferredmetaldevice
func (m_ ModelConfiguration) PreferredMetalDevice() Device /* not a class type */ {
	rv := objc.Send[Device](m_.ID, objc.Sel("preferredMetalDevice"))
	return rv
}


// The metal device you prefer this model use to make predictions (inference) and update the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/preferredmetaldevice
func (m_ ModelConfiguration) SetPreferredMetalDevice(value Device /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredMetalDevice:"), value)
}



