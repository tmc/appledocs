// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ComputePlan] class.
type IComputePlan interface {
	objectivec.IObject
	ComputeDeviceUsageForMLProgramOperation(operation IMLModelStructureProgramOperation) ComputePlanDeviceUsage
	ComputeDeviceUsageForNeuralNetworkLayer(layer IMLModelStructureNeuralNetworkLayer) ComputePlanDeviceUsage
	EstimatedCostOfMLProgramOperation(operation IMLModelStructureProgramOperation) ComputePlanCost
	ModelStructure() MLModelStructure
}

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

// Alloc allocates a new instance without initialization.
func (cc _ComputePlanClass) Alloc() ComputePlan {
	rv := objc.Send[ComputePlan](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Construct the compute plan of a model asynchronously given the location of its on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/loadContentsOfURL:configuration:completionHandler:
func (cc _ComputePlanClass) LoadContentsOfURLConfigurationCompletionHandler(url foundation.IURL, configuration IMLModelConfiguration, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadContentsOfURL:configuration:completionHandler:"), url, configuration, handler)
}


// Construct the compute plan of a model asynchronously given the model asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/loadModelAsset:configuration:completionHandler:
func (cc _ComputePlanClass) LoadModelAssetConfigurationCompletionHandler(asset IMLModelAsset, configuration IMLModelConfiguration, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadModelAsset:configuration:completionHandler:"), asset, configuration, handler)
}


// Returns The anticipated compute devices that would be used for executing an ML Program operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/computeDeviceUsageForMLProgramOperation:
func (c_ ComputePlan) ComputeDeviceUsageForMLProgramOperation(operation IMLModelStructureProgramOperation) ComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](c_.ID, objc.Sel("computeDeviceUsageForMLProgramOperation:"), operation)
	return rv
}


// Returns the anticipated compute devices that would be used for executing a NeuralNetwork layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/computeDeviceUsageForNeuralNetworkLayer:
func (c_ ComputePlan) ComputeDeviceUsageForNeuralNetworkLayer(layer IMLModelStructureNeuralNetworkLayer) ComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](c_.ID, objc.Sel("computeDeviceUsageForNeuralNetworkLayer:"), layer)
	return rv
}


// Returns the estimated cost of executing an ML Program operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/estimatedCostOfMLProgramOperation:
func (c_ ComputePlan) EstimatedCostOfMLProgramOperation(operation IMLModelStructureProgramOperation) ComputePlanCost {
	rv := objc.Send[ComputePlanCost](c_.ID, objc.Sel("estimatedCostOfMLProgramOperation:"), operation)
	return rv
}


// The model structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/modelStructure
func (c_ ComputePlan) ModelStructure() MLModelStructure {
	rv := objc.Send[MLModelStructure](c_.ID, objc.Sel("modelStructure"))
	return rv
}



