// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLComputePlan */


/* debug [class_header]: Header for MLComputePlan */
// The class instance for the [ComputePlan] class.
var (
	ComputePlanClass     _ComputePlanClass
	ComputePlanClassOnce sync.Once
)

func getComputePlanClass() _ComputePlanClass {
	ComputePlanClassOnce.Do(func() {
		ComputePlanClass = _ComputePlanClass{objc.GetClass("MLComputePlan")}
	})
	return ComputePlanClass
}

type _ComputePlanClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComputePlan */
// An interface definition for the [ComputePlan] class.
type IComputePlan interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ComputePlan */
	// properties:
	ModelStructure() IMLModelStructure
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComputePlan */
	// methods:
	ComputeDeviceUsageForMLProgramOperation(operation IMLModelStructureProgramOperation) IComputePlanDeviceUsage
	ComputeDeviceUsageForNeuralNetworkLayer(layer IMLModelStructureNeuralNetworkLayer) IComputePlanDeviceUsage
	EstimatedCostOfMLProgramOperation(operation IMLModelStructureProgramOperation) IComputePlanCost
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComputePlan */
// Alloc allocates a new instance without initialization.
func (cc _ComputePlanClass) Alloc() ComputePlan {
	rv := objc.Send[ComputePlan](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ComputePlanClass) New() ComputePlan {
	rv := objc.Send[ComputePlan](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePlan) Init() ComputePlan {
	rv := objc.Send[ComputePlan](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePlan) Autorelease() ComputePlan {
	rv := objc.Send[ComputePlan](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePlan creates a new ComputePlan instance.
func NewComputePlan() ComputePlan {
	return getComputePlanClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComputePlan */
// A class describing the plan for executing a model.
//
// The application can use the plan to estimate the necessary cost and resources of the model before running the predictions.


// A class describing the plan for executing a model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw
type ComputePlan struct {
	objectivec.Object
}

// ComputePlanFrom constructs a [ComputePlan] from an unsafe.Pointer.
//
// A class describing the plan for executing a model.
func ComputePlanFrom(ptr unsafe.Pointer) ComputePlan {
	return ComputePlan{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComputePlan *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComputePlan */

// Construct the compute plan of a model asynchronously given the location of its on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/loadContentsOfURL:configuration:completionHandler:
func (cc _ComputePlanClass) LoadContentsOfURLConfigurationCompletionHandler(url objc.IObject /* cross-framework: NSURL */, configuration IMLModelConfiguration, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadContentsOfURL:configuration:completionHandler:"), url, configuration, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadContentsOfURLConfigurationCompletionHandler) */


// Construct the compute plan of a model asynchronously given the model asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/loadModelAsset:configuration:completionHandler:
func (cc _ComputePlanClass) LoadModelAssetConfigurationCompletionHandler(asset IMLModelAsset, configuration IMLModelConfiguration, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadModelAsset:configuration:completionHandler:"), asset, configuration, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadModelAssetConfigurationCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComputePlan */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComputePlan */

// Returns The anticipated compute devices that would be used for executing an ML Program operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/computeDeviceUsageForMLProgramOperation:
func (c_ ComputePlan) ComputeDeviceUsageForMLProgramOperation(operation IMLModelStructureProgramOperation) IComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](c_.ID, objc.Sel("computeDeviceUsageForMLProgramOperation:"), operation)
	return rv
}/* debug [instance_methods/method]: ComputeDeviceUsageForMLProgramOperation */


// Returns the anticipated compute devices that would be used for executing a NeuralNetwork layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/computeDeviceUsageForNeuralNetworkLayer:
func (c_ ComputePlan) ComputeDeviceUsageForNeuralNetworkLayer(layer IMLModelStructureNeuralNetworkLayer) IComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](c_.ID, objc.Sel("computeDeviceUsageForNeuralNetworkLayer:"), layer)
	return rv
}/* debug [instance_methods/method]: ComputeDeviceUsageForNeuralNetworkLayer */


// Returns the estimated cost of executing an ML Program operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/estimatedCostOfMLProgramOperation:
func (c_ ComputePlan) EstimatedCostOfMLProgramOperation(operation IMLModelStructureProgramOperation) IComputePlanCost {
	rv := objc.Send[ComputePlanCost](c_.ID, objc.Sel("estimatedCostOfMLProgramOperation:"), operation)
	return rv
}/* debug [instance_methods/method]: EstimatedCostOfMLProgramOperation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComputePlan */

// The model structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/modelStructure
func (c_ ComputePlan) ModelStructure() IMLModelStructure {
	rv := objc.Send[ModelStructure](c_.ID, objc.Sel("modelStructure"))
	return rv
}/* debug [instance_properties/getter]: modelStructure */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLComputePlan */



