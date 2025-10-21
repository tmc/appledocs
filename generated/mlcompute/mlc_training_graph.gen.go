// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CTrainingGraph] class.
var (
	CTrainingGraphClass     _CTrainingGraphClass
	CTrainingGraphClassOnce sync.Once
)

func getCTrainingGraphClass() _CTrainingGraphClass {
	CTrainingGraphClassOnce.Do(func() {
		CTrainingGraphClass = _CTrainingGraphClass{objc.GetClass("MLCTrainingGraph")}
	})
	return CTrainingGraphClass
}

type _CTrainingGraphClass struct {
	class objc.Class
}

// An interface definition for the [CTrainingGraph] class.
type ICTrainingGraph interface {
	ICGraph
	AddInputsLossLabels(inputs unsafe.Pointer, lossLabels unsafe.Pointer) bool
	AddInputsLossLabelsLossLabelWeights(inputs unsafe.Pointer, lossLabels unsafe.Pointer, lossLabelWeights unsafe.Pointer) bool
	AddOutputs(outputs unsafe.Pointer) bool
	AllocateUserGradientForTensor(tensor IMLCTensor) CTensor
	BindOptimizerDataDeviceDataWithTensor(data []CTensorData, deviceData []CTensorOptimizerDeviceData, tensor IMLCTensor) bool
	CompileWithOptionsDevice(options CGraphCompilationOptions, device IMLCDevice) bool
	CompileWithOptionsDeviceInputTensorsInputTensorsData(options CGraphCompilationOptions, device IMLCDevice, inputTensors unsafe.Pointer, inputTensorsData unsafe.Pointer) bool
	CompileOptimizer(optimizer unsafe.Pointer) bool
	ExecuteWithInputsDataLossLabelsDataLossLabelWeightsDataBatchSizeOptionsCompletionHandler(inputsData unsafe.Pointer, lossLabelsData unsafe.Pointer, lossLabelWeightsData unsafe.Pointer, batchSize uint, options CExecutionOptions, completionHandler unsafe.Pointer) bool
	ExecuteWithInputsDataLossLabelsDataLossLabelWeightsDataOutputsDataBatchSizeOptionsCompletionHandler(inputsData unsafe.Pointer, lossLabelsData unsafe.Pointer, lossLabelWeightsData unsafe.Pointer, outputsData unsafe.Pointer, batchSize uint, options CExecutionOptions, completionHandler unsafe.Pointer) bool
	ExecuteForwardWithBatchSizeOptionsCompletionHandler(batchSize uint, options CExecutionOptions, completionHandler unsafe.Pointer) bool
	ExecuteForwardWithBatchSizeOptionsOutputsDataCompletionHandler(batchSize uint, options CExecutionOptions, outputsData unsafe.Pointer, completionHandler unsafe.Pointer) bool
	ExecuteGradientWithBatchSizeOptionsCompletionHandler(batchSize uint, options CExecutionOptions, completionHandler unsafe.Pointer) bool
	ExecuteGradientWithBatchSizeOptionsOutputsDataCompletionHandler(batchSize uint, options CExecutionOptions, outputsData unsafe.Pointer, completionHandler unsafe.Pointer) bool
	ExecuteOptimizerUpdateWithOptionsCompletionHandler(options CExecutionOptions, completionHandler unsafe.Pointer) bool
	GradientDataForParameterLayer(parameter IMLCTensor, layer IMLCLayer) foundation.Data
	GradientTensorForInput(input IMLCTensor) CTensor
	LinkWithGraphs(graphs []CTrainingGraph) bool
	ResultGradientTensorsForLayer(layer IMLCLayer) []CTensor
	SetTrainingTensorParameters(parameters []CTensorParameter) bool
	SourceGradientTensorsForLayer(layer IMLCLayer) []CTensor
	StopGradientForTensors(tensors []CTensor) bool
	SynchronizeUpdates()
}

// A training graph that you create from one or more graph objects plus additional layers you add directly to the training graph.
//
// The framework provides a family of graph-execution methods to execute a full training iteration, and methods to execute the forward pass, the gradient pass, and optimizer update, individually. Use one of the   methods to execute a full training iteration to accelerate an ML model represented as a single training graph. Use one of the  ,  , or  to accelerate an ML library that separates the forward pass, gradient pass, and optimizer update as separate phases.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph
type CTrainingGraph struct {
	CGraph
}

// CTrainingGraphFrom constructs a [CTrainingGraph] from an unsafe.Pointer.
//
// A training graph that you create from one or more graph objects plus additional layers you add directly to the training graph.
func CTrainingGraphFrom(ptr unsafe.Pointer) CTrainingGraph {
	return CTrainingGraph{
		CGraph: CGraphFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CTrainingGraphClass) Alloc() CTrainingGraph {
	rv := objc.Send[CTrainingGraph](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CTrainingGraphClass) New() CTrainingGraph {
	rv := objc.Send[CTrainingGraph](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTrainingGraph) Init() CTrainingGraph {
	rv := objc.Send[CTrainingGraph](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTrainingGraph) Autorelease() CTrainingGraph {
	rv := objc.Send[CTrainingGraph](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTrainingGraph creates a new CTrainingGraph instance.
func NewCTrainingGraph() CTrainingGraph {
	return getCTrainingGraphClass().New()
}




// Creates a training graph with the layers from the graph objects, loss layer, and optimizer you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/init(graphObjects:lossLayer:optimizer:)
func NewCTrainingGraphWithGraphObjectsLossLayerOptimizer(graphObjects []CGraph, lossLayer IMLCLayer, optimizer unsafe.Pointer) CTrainingGraph {
	rv := objc.Send[CTrainingGraph](objc.ID(getCTrainingGraphClass().class), objc.Sel("graphWithGraphObjects:lossLayer:optimizer:"), graphObjects, lossLayer, optimizer)
	return rv
}


// Creates a training graph with the layers from the graph objects, loss layer, and optimizer you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/init(graphObjects:lossLayer:optimizer:)
func (cc _CTrainingGraphClass) GraphWithGraphObjectsLossLayerOptimizer(graphObjects []CGraph, lossLayer IMLCLayer, optimizer unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("graphWithGraphObjects:lossLayer:optimizer:"), graphObjects, lossLayer, optimizer)
	return rv
}

// Adds the inputs and loss label inputs that you specify to the training graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/addInputs(_:lossLabels:)
func (c_ CTrainingGraph) AddInputsLossLabels(inputs unsafe.Pointer, lossLabels unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("addInputs:lossLabels:"), inputs, lossLabels)
	return rv
}

// Adds the inputs, loss labels, and loss label weights that you specify to the training graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/addInputs(_:lossLabels:lossLabelWeights:)
func (c_ CTrainingGraph) AddInputsLossLabelsLossLabelWeights(inputs unsafe.Pointer, lossLabels unsafe.Pointer, lossLabelWeights unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("addInputs:lossLabels:lossLabelWeights:"), inputs, lossLabels, lossLabelWeights)
	return rv
}

// Adds the outputs to the training graph you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/addOutputs(_:)
func (c_ CTrainingGraph) AddOutputs(outputs unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("addOutputs:"), outputs)
	return rv
}

// Allocates an entry for a gradient for the result tensor you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/allocateUserGradient(for:)
func (c_ CTrainingGraph) AllocateUserGradientForTensor(tensor IMLCTensor) CTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("allocateUserGradientForTensor:"), tensor)
	return rv
}

// Associates the optimizer and device data you specify along with the tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/bindOptimizerData(_:deviceData:with:)
func (c_ CTrainingGraph) BindOptimizerDataDeviceDataWithTensor(data []CTensorData, deviceData []CTensorOptimizerDeviceData, tensor IMLCTensor) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bindOptimizerData:deviceData:withTensor:"), data, deviceData, tensor)
	return rv
}

// Compiles the training graph for the options and device you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/compile(options:device:)
func (c_ CTrainingGraph) CompileWithOptionsDevice(options CGraphCompilationOptions, device IMLCDevice) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("compileWithOptions:device:"), options, device)
	return rv
}

// Compiles the training graph for the options, device, and input tensors you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/compile(options:device:inputTensors:inputTensorsData:)
func (c_ CTrainingGraph) CompileWithOptionsDeviceInputTensorsInputTensorsData(options CGraphCompilationOptions, device IMLCDevice, inputTensors unsafe.Pointer, inputTensorsData unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("compileWithOptions:device:inputTensors:inputTensorsData:"), options, device, inputTensors, inputTensorsData)
	return rv
}

// Compiles the optimizer to use with a training graph you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/compileOptimizer(_:)
func (c_ CTrainingGraph) CompileOptimizer(optimizer unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("compileOptimizer:"), optimizer)
	return rv
}

// Executes the training graph with the input data, batch size, execution options, and completion handler you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/execute(inputsData:lossLabelsData:lossLabelWeightsData:batchSize:options:completionHandler:)
func (c_ CTrainingGraph) ExecuteWithInputsDataLossLabelsDataLossLabelWeightsDataBatchSizeOptionsCompletionHandler(inputsData unsafe.Pointer, lossLabelsData unsafe.Pointer, lossLabelWeightsData unsafe.Pointer, batchSize uint, options CExecutionOptions, completionHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("executeWithInputsData:lossLabelsData:lossLabelWeightsData:batchSize:options:completionHandler:"), inputsData, lossLabelsData, lossLabelWeightsData, batchSize, options, completionHandler)
	return rv
}

// Executes the training graph with the input data, output data, batch size, execution options, and completion handler that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/execute(inputsData:lossLabelsData:lossLabelWeightsData:outputsData:batchSize:options:completionHandler:)
func (c_ CTrainingGraph) ExecuteWithInputsDataLossLabelsDataLossLabelWeightsDataOutputsDataBatchSizeOptionsCompletionHandler(inputsData unsafe.Pointer, lossLabelsData unsafe.Pointer, lossLabelWeightsData unsafe.Pointer, outputsData unsafe.Pointer, batchSize uint, options CExecutionOptions, completionHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("executeWithInputsData:lossLabelsData:lossLabelWeightsData:outputsData:batchSize:options:completionHandler:"), inputsData, lossLabelsData, lossLabelWeightsData, outputsData, batchSize, options, completionHandler)
	return rv
}

// Executes the forward pass of the training graph with the batch size, execution options, and completion handler you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/executeForward(batchSize:options:completionHandler:)
func (c_ CTrainingGraph) ExecuteForwardWithBatchSizeOptionsCompletionHandler(batchSize uint, options CExecutionOptions, completionHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("executeForwardWithBatchSize:options:completionHandler:"), batchSize, options, completionHandler)
	return rv
}

// Executes the forward pass of the training graph with the batch size, execution options, output data, and completion handler you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/executeForward(batchSize:options:outputsData:completionHandler:)
func (c_ CTrainingGraph) ExecuteForwardWithBatchSizeOptionsOutputsDataCompletionHandler(batchSize uint, options CExecutionOptions, outputsData unsafe.Pointer, completionHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("executeForwardWithBatchSize:options:outputsData:completionHandler:"), batchSize, options, outputsData, completionHandler)
	return rv
}

// Executes the gradient pass of the training graph with the batch size, execution options, and completion handler you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/executeGradient(batchSize:options:completionHandler:)
func (c_ CTrainingGraph) ExecuteGradientWithBatchSizeOptionsCompletionHandler(batchSize uint, options CExecutionOptions, completionHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("executeGradientWithBatchSize:options:completionHandler:"), batchSize, options, completionHandler)
	return rv
}

// Executes the gradient pass of the training graph with the batch size, execution options, output data, and completion handler you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/executeGradient(batchSize:options:outputsData:completionHandler:)
func (c_ CTrainingGraph) ExecuteGradientWithBatchSizeOptionsOutputsDataCompletionHandler(batchSize uint, options CExecutionOptions, outputsData unsafe.Pointer, completionHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("executeGradientWithBatchSize:options:outputsData:completionHandler:"), batchSize, options, outputsData, completionHandler)
	return rv
}

// Executes the optimizer update pass of the training graph with the execution options and completion handler you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/executeOptimizerUpdate(options:completionHandler:)
func (c_ CTrainingGraph) ExecuteOptimizerUpdateWithOptionsCompletionHandler(options CExecutionOptions, completionHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("executeOptimizerUpdateWithOptions:completionHandler:"), options, completionHandler)
	return rv
}

// Gets the gradient data for the trainable parameter and associated layer you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/gradientData(forParameter:layer:)
func (c_ CTrainingGraph) GradientDataForParameterLayer(parameter IMLCTensor, layer IMLCLayer) foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("gradientDataForParameter:layer:"), parameter, layer)
	return rv
}

// Gets the gradient tensor for the input tensor you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/gradientTensor(forInput:)
func (c_ CTrainingGraph) GradientTensorForInput(input IMLCTensor) CTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("gradientTensorForInput:"), input)
	return rv
}

// Links the training graphs you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/link(with:)
func (c_ CTrainingGraph) LinkWithGraphs(graphs []CTrainingGraph) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("linkWithGraphs:"), graphs)
	return rv
}

// Gets the result gradient tensors for the layer in the training graph you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/resultGradientTensors(for:)
func (c_ CTrainingGraph) ResultGradientTensorsForLayer(layer IMLCLayer) []CTensor {
	rv := objc.Send[[]CTensor](c_.ID, objc.Sel("resultGradientTensorsForLayer:"), layer)
	return rv
}

// Sets the input tensor parameters, which the optimizer then updates.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/setTrainingTensorParameters(_:)
func (c_ CTrainingGraph) SetTrainingTensorParameters(parameters []CTensorParameter) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setTrainingTensorParameters:"), parameters)
	return rv
}

// Gets the source gradient tensors for the layer in the training graph you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/sourceGradientTensors(for:)
func (c_ CTrainingGraph) SourceGradientTensorsForLayer(layer IMLCLayer) []CTensor {
	rv := objc.Send[[]CTensor](c_.ID, objc.Sel("sourceGradientTensorsForLayer:"), layer)
	return rv
}

// Adds the tensors that you specify, to indicate which contributions the graph excludes when computing gradients during gradient pass.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/stopGradient(for:)
func (c_ CTrainingGraph) StopGradientForTensors(tensors []CTensor) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("stopGradientForTensors:"), tensors)
	return rv
}

// Synchronizes updates from device memory.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/synchronizeUpdates()
func (c_ CTrainingGraph) SynchronizeUpdates() {
	objc.Send[objc.ID](c_.ID, objc.Sel("synchronizeUpdates"))
}

// The device memory size in bytes for all intermediate tensors for forward, gradient passes, and optimizer updates for all layers in the training graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/deviceMemorySize
func (c_ CTrainingGraph) DeviceMemorySize() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("deviceMemorySize"))
	return rv
}

// The optimizer to use with the training graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/optimizer
func (c_ CTrainingGraph) Optimizer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("optimizer"))
	return rv
}


