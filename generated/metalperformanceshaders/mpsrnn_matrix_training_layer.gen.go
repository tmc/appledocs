// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [RNNMatrixTrainingLayer] class.
var (
	RNNMatrixTrainingLayerClass     _RNNMatrixTrainingLayerClass
	RNNMatrixTrainingLayerClassOnce sync.Once
)

func getRNNMatrixTrainingLayerClass() _RNNMatrixTrainingLayerClass {
	RNNMatrixTrainingLayerClassOnce.Do(func() {
		RNNMatrixTrainingLayerClass = _RNNMatrixTrainingLayerClass{objc.GetClass("MPSRNNMatrixTrainingLayer")}
	})
	return RNNMatrixTrainingLayerClass
}

type _RNNMatrixTrainingLayerClass struct {
	class objc.Class
}





// An interface definition for the [RNNMatrixTrainingLayer] class.
type IRNNMatrixTrainingLayer interface {
	IKernel
	

	// properties:
	AccumulateWeightGradients() objectivec.IObject
	SetAccumulateWeightGradients(value objectivec.IObject)
	InputFeatureChannels() objectivec.IObject
	SetInputFeatureChannels(value objectivec.IObject)
	OutputFeatureChannels() objectivec.IObject
	SetOutputFeatureChannels(value objectivec.IObject)
	RecurrentOutputIsTemporary() objectivec.IObject
	SetRecurrentOutputIsTemporary(value objectivec.IObject)
	StoreAllIntermediateStates() objectivec.IObject
	SetStoreAllIntermediateStates(value objectivec.IObject)
	TrainingStateIsTemporary() objectivec.IObject
	SetTrainingStateIsTemporary(value objectivec.IObject)


	

	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	CreateTemporaryWeightGradientMatrices()
	CreateWeightGradientMatrices()
	CreateWeightMatrices()
	EncodeCopyWeights()
	EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset(commandBuffer unsafe.Pointer, weights unsafe.Pointer, matrixId RNNMatrixId, matrix IMatrix, copyFromWeightsToMatrix bool, matrixOffset objc.IObject /* cross-framework: MTLOrigin */)
	EncodeForwardSequence()
	EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, destinationMatrices unsafe.Pointer, trainingStates unsafe.Pointer, weights unsafe.Pointer)
	EncodeForwardSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, sourceOffsets uint, destinationMatrices unsafe.Pointer, destinationOffsets uint, trainingStates unsafe.Pointer, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer, weights unsafe.Pointer)
	EncodeGradientSequence()
	EncodeGradientSequenceToCommandBufferForwardSourcesForwardSourceOffsetsSourceGradientsSourceGradientOffsetsDestinationGradientsDestinationOffsetsWeightGradientsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer unsafe.Pointer, forwardSources unsafe.Pointer, forwardSourceOffsets uint, sourceGradients unsafe.Pointer, sourceGradientOffsets uint, destinationGradients unsafe.Pointer, destinationOffsets uint, weightGradients unsafe.Pointer, trainingStates unsafe.Pointer, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer, weights unsafe.Pointer)
	EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights(commandBuffer unsafe.Pointer, forwardSources unsafe.Pointer, sourceGradients unsafe.Pointer, destinationGradients unsafe.Pointer, weightGradients unsafe.Pointer, trainingStates unsafe.Pointer, weights unsafe.Pointer)
	CreateTemporaryWeightGradientMatricesDataTypeCommandBuffer(matricesOut unsafe.Pointer, dataType DataType, commandBuffer unsafe.Pointer)
	CreateWeightGradientMatricesDataType(matricesOut unsafe.Pointer, dataType DataType)
	CreateWeightMatricesWithMatricesOut(matricesOut unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (rc _RNNMatrixTrainingLayerClass) Alloc() RNNMatrixTrainingLayer {
	rv := objc.Send[RNNMatrixTrainingLayer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RNNMatrixTrainingLayerClass) New() RNNMatrixTrainingLayer {
	rv := objc.Send[RNNMatrixTrainingLayer](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNMatrixTrainingLayer) Init() RNNMatrixTrainingLayer {
	rv := objc.Send[RNNMatrixTrainingLayer](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNMatrixTrainingLayer) Autorelease() RNNMatrixTrainingLayer {
	rv := objc.Send[RNNMatrixTrainingLayer](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNMatrixTrainingLayer creates a new RNNMatrixTrainingLayer instance.
func NewRNNMatrixTrainingLayer() RNNMatrixTrainingLayer {
	return getRNNMatrixTrainingLayerClass().New()
}





// A layer for training recurrent neural networks on Metal Performance Shaders matrices.


// A layer for training recurrent neural networks on Metal Performance Shaders matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer
type RNNMatrixTrainingLayer struct {
	Kernel
}

// RNNMatrixTrainingLayerFrom constructs a [RNNMatrixTrainingLayer] from an unsafe.Pointer.
//
// A layer for training recurrent neural networks on Metal Performance Shaders matrices.
func RNNMatrixTrainingLayerFrom(ptr unsafe.Pointer) RNNMatrixTrainingLayer {
	return RNNMatrixTrainingLayer{
		Kernel: KernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966793-initwithcoder
func NewRNNMatrixTrainingLayerWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) RNNMatrixTrainingLayer {
	instance := getRNNMatrixTrainingLayerClass().Alloc()
	rv := objc.Send[RNNMatrixTrainingLayer](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966794-initwithdevice
func NewRNNMatrixTrainingLayerWithDeviceRnnDescriptorTrainableWeights(device unsafe.Pointer, rnnDescriptor IRNNDescriptor, trainableWeights unsafe.Pointer) RNNMatrixTrainingLayer {
	instance := getRNNMatrixTrainingLayerClass().Alloc()
	rv := objc.Send[RNNMatrixTrainingLayer](instance.ID, objc.Sel("initWithDevice:rnnDescriptor:trainableWeights:"), device, rnnDescriptor, trainableWeights)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966784-copywithzone
func (r_ RNNMatrixTrainingLayer) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966785-createtemporaryweightgradientmat
func (r_ RNNMatrixTrainingLayer) CreateTemporaryWeightGradientMatrices() {
	objc.Send[objc.ID](r_.ID, objc.Sel("createTemporaryWeightGradientMatrices"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966786-createweightgradientmatrices
func (r_ RNNMatrixTrainingLayer) CreateWeightGradientMatrices() {
	objc.Send[objc.ID](r_.ID, objc.Sel("createWeightGradientMatrices"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966787-createweightmatrices
func (r_ RNNMatrixTrainingLayer) CreateWeightMatrices() {
	objc.Send[objc.ID](r_.ID, objc.Sel("createWeightMatrices"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966788-encodecopyweights
func (r_ RNNMatrixTrainingLayer) EncodeCopyWeights() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeCopyWeights"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966788-encodecopyweightstocommandbuffer
func (r_ RNNMatrixTrainingLayer) EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset(commandBuffer unsafe.Pointer, weights unsafe.Pointer, matrixId RNNMatrixId, matrix IMatrix, copyFromWeightsToMatrix bool, matrixOffset objc.IObject /* cross-framework: MTLOrigin */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeCopyWeightsToCommandBuffer:weights:matrixId:matrix:copyFromWeightsToMatrix:matrixOffset:"), commandBuffer, weights, matrixId, matrix, copyFromWeightsToMatrix, matrixOffset)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966789-encodeforwardsequence
func (r_ RNNMatrixTrainingLayer) EncodeForwardSequence() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeForwardSequence"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966789-encodeforwardsequencetocommandbu
func (r_ RNNMatrixTrainingLayer) EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, destinationMatrices unsafe.Pointer, trainingStates unsafe.Pointer, weights unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeForwardSequenceToCommandBuffer:sourceMatrices:destinationMatrices:trainingStates:weights:"), commandBuffer, sourceMatrices, destinationMatrices, trainingStates, weights)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966790-encodeforwardsequencetocommandbu
func (r_ RNNMatrixTrainingLayer) EncodeForwardSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, sourceOffsets uint, destinationMatrices unsafe.Pointer, destinationOffsets uint, trainingStates unsafe.Pointer, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer, weights unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeForwardSequenceToCommandBuffer:sourceMatrices:sourceOffsets:destinationMatrices:destinationOffsets:trainingStates:recurrentInputState:recurrentOutputStates:weights:"), commandBuffer, sourceMatrices, sourceOffsets, destinationMatrices, destinationOffsets, trainingStates, recurrentInputState, recurrentOutputStates, weights)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966791-encodegradientsequence
func (r_ RNNMatrixTrainingLayer) EncodeGradientSequence() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeGradientSequence"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966791-encodegradientsequencetocommandb
func (r_ RNNMatrixTrainingLayer) EncodeGradientSequenceToCommandBufferForwardSourcesForwardSourceOffsetsSourceGradientsSourceGradientOffsetsDestinationGradientsDestinationOffsetsWeightGradientsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer unsafe.Pointer, forwardSources unsafe.Pointer, forwardSourceOffsets uint, sourceGradients unsafe.Pointer, sourceGradientOffsets uint, destinationGradients unsafe.Pointer, destinationOffsets uint, weightGradients unsafe.Pointer, trainingStates unsafe.Pointer, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer, weights unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeGradientSequenceToCommandBuffer:forwardSources:forwardSourceOffsets:sourceGradients:sourceGradientOffsets:destinationGradients:destinationOffsets:weightGradients:trainingStates:recurrentInputState:recurrentOutputStates:weights:"), commandBuffer, forwardSources, forwardSourceOffsets, sourceGradients, sourceGradientOffsets, destinationGradients, destinationOffsets, weightGradients, trainingStates, recurrentInputState, recurrentOutputStates, weights)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966792-encodegradientsequencetocommandb
func (r_ RNNMatrixTrainingLayer) EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights(commandBuffer unsafe.Pointer, forwardSources unsafe.Pointer, sourceGradients unsafe.Pointer, destinationGradients unsafe.Pointer, weightGradients unsafe.Pointer, trainingStates unsafe.Pointer, weights unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeGradientSequenceToCommandBuffer:forwardSources:sourceGradients:destinationGradients:weightGradients:trainingStates:weights:"), commandBuffer, forwardSources, sourceGradients, destinationGradients, weightGradients, trainingStates, weights)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/createTemporaryWeightGradientMatrices(_:dataType:commandBuffer:)
func (r_ RNNMatrixTrainingLayer) CreateTemporaryWeightGradientMatricesDataTypeCommandBuffer(matricesOut unsafe.Pointer, dataType DataType, commandBuffer unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("createTemporaryWeightGradientMatrices:dataType:commandBuffer:"), matricesOut, dataType, commandBuffer)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/createWeightGradientMatrices(_:dataType:)
func (r_ RNNMatrixTrainingLayer) CreateWeightGradientMatricesDataType(matricesOut unsafe.Pointer, dataType DataType) {
	objc.Send[objc.ID](r_.ID, objc.Sel("createWeightGradientMatrices:dataType:"), matricesOut, dataType)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/createWeightMatrices(_:)
func (r_ RNNMatrixTrainingLayer) CreateWeightMatricesWithMatricesOut(matricesOut unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("createWeightMatrices:"), matricesOut)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966783-accumulateweightgradients
func (r_ RNNMatrixTrainingLayer) AccumulateWeightGradients() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("accumulateWeightGradients"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966783-accumulateweightgradients
func (r_ RNNMatrixTrainingLayer) SetAccumulateWeightGradients(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAccumulateWeightGradients:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966795-inputfeaturechannels
func (r_ RNNMatrixTrainingLayer) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966795-inputfeaturechannels
func (r_ RNNMatrixTrainingLayer) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966796-outputfeaturechannels
func (r_ RNNMatrixTrainingLayer) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966796-outputfeaturechannels
func (r_ RNNMatrixTrainingLayer) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966797-recurrentoutputistemporary
func (r_ RNNMatrixTrainingLayer) RecurrentOutputIsTemporary() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966797-recurrentoutputistemporary
func (r_ RNNMatrixTrainingLayer) SetRecurrentOutputIsTemporary(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966798-storeallintermediatestates
func (r_ RNNMatrixTrainingLayer) StoreAllIntermediateStates() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("storeAllIntermediateStates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966798-storeallintermediatestates
func (r_ RNNMatrixTrainingLayer) SetStoreAllIntermediateStates(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStoreAllIntermediateStates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966799-trainingstateistemporary
func (r_ RNNMatrixTrainingLayer) TrainingStateIsTemporary() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("trainingStateIsTemporary"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966799-trainingstateistemporary
func (r_ RNNMatrixTrainingLayer) SetTrainingStateIsTemporary(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTrainingStateIsTemporary:"), value)
}







