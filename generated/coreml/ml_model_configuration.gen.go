// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelConfiguration */


/* debug [class_header]: Header for MLModelConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelConfiguration */
// An interface definition for the [ModelConfiguration] class.
type IModelConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelConfiguration */
	// properties:
	AllowLowPrecisionAccumulationOnGPU() bool
	SetAllowLowPrecisionAccumulationOnGPU(value bool)
	ComputeUnits() ComputeUnits
	SetComputeUnits(value ComputeUnits)
	FunctionName() objc.IObject /* cross-framework: NSString */
	SetFunctionName(value objc.IObject /* cross-framework: NSString */)
	ModelDisplayName() objc.IObject /* cross-framework: NSString */
	SetModelDisplayName(value objc.IObject /* cross-framework: NSString */)
	OptimizationHints() IMLOptimizationHints
	SetOptimizationHints(value IMLOptimizationHints)
	Parameters() foundation.IDictionary
	SetParameters(value foundation.IDictionary)
	PreferredMetalDevice() unsafe.Pointer
	SetPreferredMetalDevice(value unsafe.Pointer)
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelConfiguration */

// A Boolean value that determines whether to allow low-precision accumulation on a GPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/allowLowPrecisionAccumulationOnGPU
func (m_ ModelConfiguration) AllowLowPrecisionAccumulationOnGPU() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowLowPrecisionAccumulationOnGPU"))
	return rv
}/* debug [instance_properties/getter]: allowLowPrecisionAccumulationOnGPU */


// A Boolean value that determines whether to allow low-precision accumulation on a GPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/allowLowPrecisionAccumulationOnGPU
func (m_ ModelConfiguration) SetAllowLowPrecisionAccumulationOnGPU(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowLowPrecisionAccumulationOnGPU:"), value)
}/* debug [instance_properties/setter]: allowLowPrecisionAccumulationOnGPU */


// The processing unit or units the model uses to make predictions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/computeUnits
func (m_ ModelConfiguration) ComputeUnits() ComputeUnits {
	rv := objc.Send[ComputeUnits](m_.ID, objc.Sel("computeUnits"))
	return rv
}/* debug [instance_properties/getter]: computeUnits */


// The processing unit or units the model uses to make predictions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/computeUnits
func (m_ ModelConfiguration) SetComputeUnits(value ComputeUnits) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setComputeUnits:"), value)
}/* debug [instance_properties/setter]: computeUnits */


// Function name that will use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/functionName
func (m_ ModelConfiguration) FunctionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("functionName"))
	return rv
}/* debug [instance_properties/getter]: functionName */


// Function name that will use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/functionName
func (m_ ModelConfiguration) SetFunctionName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFunctionName:"), value)
}/* debug [instance_properties/setter]: functionName */


// A human readable name of a model for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/modelDisplayName
func (m_ ModelConfiguration) ModelDisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("modelDisplayName"))
	return rv
}/* debug [instance_properties/getter]: modelDisplayName */


// A human readable name of a model for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/modelDisplayName
func (m_ ModelConfiguration) SetModelDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModelDisplayName:"), value)
}/* debug [instance_properties/setter]: modelDisplayName */


// A group of hints for CoreML to optimize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/optimizationHints-81u6f
func (m_ ModelConfiguration) OptimizationHints() IMLOptimizationHints {
	rv := objc.Send[OptimizationHints](m_.ID, objc.Sel("optimizationHints"))
	return rv
}/* debug [instance_properties/getter]: optimizationHints */


// A group of hints for CoreML to optimize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/optimizationHints-81u6f
func (m_ ModelConfiguration) SetOptimizationHints(value IMLOptimizationHints) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptimizationHints:"), value)
}/* debug [instance_properties/setter]: optimizationHints */


// A dictionary of configuration settings your app can override when loading a model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/parameters
func (m_ ModelConfiguration) Parameters() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("parameters"))
	return rv
}/* debug [instance_properties/getter]: parameters */


// A dictionary of configuration settings your app can override when loading a model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/parameters
func (m_ ModelConfiguration) SetParameters(value foundation.IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParameters:"), value)
}/* debug [instance_properties/setter]: parameters */


// The metal device you prefer this model use to make predictions (inference) and update the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/preferredMetalDevice
func (m_ ModelConfiguration) PreferredMetalDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("preferredMetalDevice"))
	return rv
}/* debug [instance_properties/getter]: preferredMetalDevice */


// The metal device you prefer this model use to make predictions (inference) and update the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/preferredMetalDevice
func (m_ ModelConfiguration) SetPreferredMetalDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredMetalDevice:"), value)
}/* debug [instance_properties/setter]: preferredMetalDevice */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelConfiguration */



