// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSRNNMatrixTrainingLayer */


/* debug [class_header]: Header for MPSRNNMatrixTrainingLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RNNMatrixTrainingLayer */
// An interface definition for the [RNNMatrixTrainingLayer] class.
type IRNNMatrixTrainingLayer interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for RNNMatrixTrainingLayer */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RNNMatrixTrainingLayer */
	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	CreateTemporaryWeightGradientMatrices()
	CreateWeightGradientMatrices()
	CreateWeightMatrices()
	EncodeCopyWeights()
	EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset(commandBuffer unsafe.Pointer, weights unsafe.Pointer, matrixId RNNMatrixId, matrix IMatrix, copyFromWeightsToMatrix bool, matrixOffset Origin /* not a class type */)
	EncodeForwardSequence()
	EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, destinationMatrices unsafe.Pointer, trainingStates unsafe.Pointer, weights unsafe.Pointer)
	EncodeForwardSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, sourceOffsets uint, destinationMatrices unsafe.Pointer, destinationOffsets uint, trainingStates unsafe.Pointer, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer, weights unsafe.Pointer)
	EncodeGradientSequence()
	EncodeGradientSequenceToCommandBufferForwardSourcesForwardSourceOffsetsSourceGradientsSourceGradientOffsetsDestinationGradientsDestinationOffsetsWeightGradientsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer unsafe.Pointer, forwardSources unsafe.Pointer, forwardSourceOffsets uint, sourceGradients unsafe.Pointer, sourceGradientOffsets uint, destinationGradients unsafe.Pointer, destinationOffsets uint, weightGradients unsafe.Pointer, trainingStates unsafe.Pointer, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer, weights unsafe.Pointer)
	EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights(commandBuffer unsafe.Pointer, forwardSources unsafe.Pointer, sourceGradients unsafe.Pointer, destinationGradients unsafe.Pointer, weightGradients unsafe.Pointer, trainingStates unsafe.Pointer, weights unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RNNMatrixTrainingLayer */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RNNMatrixTrainingLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RNNMatrixTrainingLayer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966793-initwithcoder
func NewRNNMatrixTrainingLayerWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) RNNMatrixTrainingLayer {
	instance := getRNNMatrixTrainingLayerClass().Alloc()
	rv := objc.Send[RNNMatrixTrainingLayer](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRNNMatrixTrainingLayerWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966794-initwithdevice
func NewRNNMatrixTrainingLayerWithDeviceRnnDescriptorTrainableWeights(device unsafe.Pointer, rnnDescriptor IRNNDescriptor, trainableWeights unsafe.Pointer) RNNMatrixTrainingLayer {
	instance := getRNNMatrixTrainingLayerClass().Alloc()
	rv := objc.Send[RNNMatrixTrainingLayer](instance.ID, objc.Sel("initWithDevice:rnnDescriptor:trainableWeights:"), device, rnnDescriptor, trainableWeights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRNNMatrixTrainingLayerWithDeviceRnnDescriptorTrainableWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RNNMatrixTrainingLayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RNNMatrixTrainingLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RNNMatrixTrainingLayer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966784-copywithzone
func (r_ RNNMatrixTrainingLayer) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966785-createtemporaryweightgradientmat
func (r_ RNNMatrixTrainingLayer) CreateTemporaryWeightGradientMatrices() {
	objc.Send[objc.ID](r_.ID, objc.Sel("createTemporaryWeightGradientMatrices"))
}/* debug [instance_methods/method]: CreateTemporaryWeightGradientMatrices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966786-createweightgradientmatrices
func (r_ RNNMatrixTrainingLayer) CreateWeightGradientMatrices() {
	objc.Send[objc.ID](r_.ID, objc.Sel("createWeightGradientMatrices"))
}/* debug [instance_methods/method]: CreateWeightGradientMatrices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966787-createweightmatrices
func (r_ RNNMatrixTrainingLayer) CreateWeightMatrices() {
	objc.Send[objc.ID](r_.ID, objc.Sel("createWeightMatrices"))
}/* debug [instance_methods/method]: CreateWeightMatrices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966788-encodecopyweights
func (r_ RNNMatrixTrainingLayer) EncodeCopyWeights() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeCopyWeights"))
}/* debug [instance_methods/method]: EncodeCopyWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966788-encodecopyweightstocommandbuffer
func (r_ RNNMatrixTrainingLayer) EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset(commandBuffer unsafe.Pointer, weights unsafe.Pointer, matrixId RNNMatrixId, matrix IMatrix, copyFromWeightsToMatrix bool, matrixOffset Origin /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeCopyWeightsToCommandBuffer:weights:matrixId:matrix:copyFromWeightsToMatrix:matrixOffset:"), commandBuffer, weights, matrixId, matrix, copyFromWeightsToMatrix, matrixOffset)
}/* debug [instance_methods/method]: EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966789-encodeforwardsequence
func (r_ RNNMatrixTrainingLayer) EncodeForwardSequence() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeForwardSequence"))
}/* debug [instance_methods/method]: EncodeForwardSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966789-encodeforwardsequencetocommandbu
func (r_ RNNMatrixTrainingLayer) EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, destinationMatrices unsafe.Pointer, trainingStates unsafe.Pointer, weights unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeForwardSequenceToCommandBuffer:sourceMatrices:destinationMatrices:trainingStates:weights:"), commandBuffer, sourceMatrices, destinationMatrices, trainingStates, weights)
}/* debug [instance_methods/method]: EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966790-encodeforwardsequencetocommandbu
func (r_ RNNMatrixTrainingLayer) EncodeForwardSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, sourceOffsets uint, destinationMatrices unsafe.Pointer, destinationOffsets uint, trainingStates unsafe.Pointer, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer, weights unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeForwardSequenceToCommandBuffer:sourceMatrices:sourceOffsets:destinationMatrices:destinationOffsets:trainingStates:recurrentInputState:recurrentOutputStates:weights:"), commandBuffer, sourceMatrices, sourceOffsets, destinationMatrices, destinationOffsets, trainingStates, recurrentInputState, recurrentOutputStates, weights)
}/* debug [instance_methods/method]: EncodeForwardSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966791-encodegradientsequence
func (r_ RNNMatrixTrainingLayer) EncodeGradientSequence() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeGradientSequence"))
}/* debug [instance_methods/method]: EncodeGradientSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966791-encodegradientsequencetocommandb
func (r_ RNNMatrixTrainingLayer) EncodeGradientSequenceToCommandBufferForwardSourcesForwardSourceOffsetsSourceGradientsSourceGradientOffsetsDestinationGradientsDestinationOffsetsWeightGradientsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer unsafe.Pointer, forwardSources unsafe.Pointer, forwardSourceOffsets uint, sourceGradients unsafe.Pointer, sourceGradientOffsets uint, destinationGradients unsafe.Pointer, destinationOffsets uint, weightGradients unsafe.Pointer, trainingStates unsafe.Pointer, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer, weights unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeGradientSequenceToCommandBuffer:forwardSources:forwardSourceOffsets:sourceGradients:sourceGradientOffsets:destinationGradients:destinationOffsets:weightGradients:trainingStates:recurrentInputState:recurrentOutputStates:weights:"), commandBuffer, forwardSources, forwardSourceOffsets, sourceGradients, sourceGradientOffsets, destinationGradients, destinationOffsets, weightGradients, trainingStates, recurrentInputState, recurrentOutputStates, weights)
}/* debug [instance_methods/method]: EncodeGradientSequenceToCommandBufferForwardSourcesForwardSourceOffsetsSourceGradientsSourceGradientOffsetsDestinationGradientsDestinationOffsetsWeightGradientsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966792-encodegradientsequencetocommandb
func (r_ RNNMatrixTrainingLayer) EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights(commandBuffer unsafe.Pointer, forwardSources unsafe.Pointer, sourceGradients unsafe.Pointer, destinationGradients unsafe.Pointer, weightGradients unsafe.Pointer, trainingStates unsafe.Pointer, weights unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeGradientSequenceToCommandBuffer:forwardSources:sourceGradients:destinationGradients:weightGradients:trainingStates:weights:"), commandBuffer, forwardSources, sourceGradients, destinationGradients, weightGradients, trainingStates, weights)
}/* debug [instance_methods/method]: EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RNNMatrixTrainingLayer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966783-accumulateweightgradients
func (r_ RNNMatrixTrainingLayer) AccumulateWeightGradients() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("accumulateWeightGradients"))
	return rv
}/* debug [instance_properties/getter]: accumulateWeightGradients */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966783-accumulateweightgradients
func (r_ RNNMatrixTrainingLayer) SetAccumulateWeightGradients(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAccumulateWeightGradients:"), value)
}/* debug [instance_properties/setter]: accumulateWeightGradients */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966795-inputfeaturechannels
func (r_ RNNMatrixTrainingLayer) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966795-inputfeaturechannels
func (r_ RNNMatrixTrainingLayer) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966796-outputfeaturechannels
func (r_ RNNMatrixTrainingLayer) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966796-outputfeaturechannels
func (r_ RNNMatrixTrainingLayer) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966797-recurrentoutputistemporary
func (r_ RNNMatrixTrainingLayer) RecurrentOutputIsTemporary() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}/* debug [instance_properties/getter]: recurrentOutputIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966797-recurrentoutputistemporary
func (r_ RNNMatrixTrainingLayer) SetRecurrentOutputIsTemporary(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}/* debug [instance_properties/setter]: recurrentOutputIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966798-storeallintermediatestates
func (r_ RNNMatrixTrainingLayer) StoreAllIntermediateStates() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("storeAllIntermediateStates"))
	return rv
}/* debug [instance_properties/getter]: storeAllIntermediateStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966798-storeallintermediatestates
func (r_ RNNMatrixTrainingLayer) SetStoreAllIntermediateStates(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStoreAllIntermediateStates:"), value)
}/* debug [instance_properties/setter]: storeAllIntermediateStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966799-trainingstateistemporary
func (r_ RNNMatrixTrainingLayer) TrainingStateIsTemporary() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("trainingStateIsTemporary"))
	return rv
}/* debug [instance_properties/getter]: trainingStateIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966799-trainingstateistemporary
func (r_ RNNMatrixTrainingLayer) SetTrainingStateIsTemporary(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTrainingStateIsTemporary:"), value)
}/* debug [instance_properties/setter]: trainingStateIsTemporary */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSRNNMatrixTrainingLayer */


