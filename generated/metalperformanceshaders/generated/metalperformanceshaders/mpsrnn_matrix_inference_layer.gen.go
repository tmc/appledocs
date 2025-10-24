// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSRNNMatrixInferenceLayer */


/* debug [class_header]: Header for MPSRNNMatrixInferenceLayer */
// The class instance for the [RNNMatrixInferenceLayer] class.
var (
	RNNMatrixInferenceLayerClass     _RNNMatrixInferenceLayerClass
	RNNMatrixInferenceLayerClassOnce sync.Once
)

func getRNNMatrixInferenceLayerClass() _RNNMatrixInferenceLayerClass {
	RNNMatrixInferenceLayerClassOnce.Do(func() {
		RNNMatrixInferenceLayerClass = _RNNMatrixInferenceLayerClass{objc.GetClass("MPSRNNMatrixInferenceLayer")}
	})
	return RNNMatrixInferenceLayerClass
}

type _RNNMatrixInferenceLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RNNMatrixInferenceLayer */
// An interface definition for the [RNNMatrixInferenceLayer] class.
type IRNNMatrixInferenceLayer interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for RNNMatrixInferenceLayer */
	// properties:
	RecurrentOutputIsTemporary() objectivec.IObject
	SetRecurrentOutputIsTemporary(value objectivec.IObject)
	StoreAllIntermediateStates() objectivec.IObject
	SetStoreAllIntermediateStates(value objectivec.IObject)
	BidirectionalCombineMode() RNNBidirectionalCombineMode get set /* not a class type */
	SetBidirectionalCombineMode(value RNNBidirectionalCombineMode get set /* not a class type */)
	NumberOfLayers() objectivec.IObject
	SetNumberOfLayers(value objectivec.IObject)
	OutputFeatureChannels() objectivec.IObject
	SetOutputFeatureChannels(value objectivec.IObject)
	InputFeatureChannels() objectivec.IObject
	SetInputFeatureChannels(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RNNMatrixInferenceLayer */
	// methods:
	EncodeBidirectionalSequence()
	EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices(commandBuffer unsafe.Pointer, sourceSequence unsafe.Pointer, destinationForwardMatrices unsafe.Pointer, destinationBackwardMatrices unsafe.Pointer)
	EncodeSequence()
	EncodeSequenceToCommandBufferSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, destinationMatrices unsafe.Pointer, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer)
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	EncodeSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, sourceOffsets uint, destinationMatrices unsafe.Pointer, destinationOffsets uint, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RNNMatrixInferenceLayer */
// Alloc allocates a new instance without initialization.
func (rc _RNNMatrixInferenceLayerClass) Alloc() RNNMatrixInferenceLayer {
	rv := objc.Send[RNNMatrixInferenceLayer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RNNMatrixInferenceLayerClass) New() RNNMatrixInferenceLayer {
	rv := objc.Send[RNNMatrixInferenceLayer](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNMatrixInferenceLayer) Init() RNNMatrixInferenceLayer {
	rv := objc.Send[RNNMatrixInferenceLayer](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNMatrixInferenceLayer) Autorelease() RNNMatrixInferenceLayer {
	rv := objc.Send[RNNMatrixInferenceLayer](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNMatrixInferenceLayer creates a new RNNMatrixInferenceLayer instance.
func NewRNNMatrixInferenceLayer() RNNMatrixInferenceLayer {
	return getRNNMatrixInferenceLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RNNMatrixInferenceLayer */
// A recurrent neural network layer for inference on Metal Performance Shaders matrices.
//
// The specifies a recurrent neural network layer for inference on objects. Two types of recurrent layers are supported: —Operates with convolutions on images. —Operates on matrices. You can use to implement the latter by using 1 x 1 images, but due to image size restrictions and performance, is the better choice for linear recurrent layers. is initialized using either of the following: A single instance, which further specifies the recurrent network layer. An array of instances, which specifies a stack of recurrent layers that can operate in parallel a subset of the inputs in a sequence of inputs and recurrent outputs. Stacks with bidirectionally traversing encode functions don’t support starting from a previous set of recurrent states. However, you can achieve this effect by defining two separate unidirectional stacks of layers, running the same input sequence on them separately (one forward and one backward), and ultimately combining the two result sequences.


// A recurrent neural network layer for inference on Metal Performance Shaders matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer
type RNNMatrixInferenceLayer struct {
	Kernel
}

// RNNMatrixInferenceLayerFrom constructs a [RNNMatrixInferenceLayer] from an unsafe.Pointer.
//
// A recurrent neural network layer for inference on Metal Performance Shaders matrices.
func RNNMatrixInferenceLayerFrom(ptr unsafe.Pointer) RNNMatrixInferenceLayer {
	return RNNMatrixInferenceLayer{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RNNMatrixInferenceLayer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865745-initwithcoder
func NewRNNMatrixInferenceLayerWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) RNNMatrixInferenceLayer {
	instance := getRNNMatrixInferenceLayerClass().Alloc()
	rv := objc.Send[RNNMatrixInferenceLayer](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRNNMatrixInferenceLayerWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865704-initwithdevice
func NewRNNMatrixInferenceLayerWithDeviceRnnDescriptor(device unsafe.Pointer, rnnDescriptor IRNNDescriptor) RNNMatrixInferenceLayer {
	instance := getRNNMatrixInferenceLayerClass().Alloc()
	rv := objc.Send[RNNMatrixInferenceLayer](instance.ID, objc.Sel("initWithDevice:rnnDescriptor:"), device, rnnDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRNNMatrixInferenceLayerWithDeviceRnnDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865751-initwithdevice
func NewRNNMatrixInferenceLayerWithDeviceRnnDescriptors(device unsafe.Pointer, rnnDescriptors unsafe.Pointer) RNNMatrixInferenceLayer {
	instance := getRNNMatrixInferenceLayerClass().Alloc()
	rv := objc.Send[RNNMatrixInferenceLayer](instance.ID, objc.Sel("initWithDevice:rnnDescriptors:"), device, rnnDescriptors)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRNNMatrixInferenceLayerWithDeviceRnnDescriptors */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RNNMatrixInferenceLayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RNNMatrixInferenceLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RNNMatrixInferenceLayer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865698-encodebidirectionalsequence
func (r_ RNNMatrixInferenceLayer) EncodeBidirectionalSequence() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeBidirectionalSequence"))
}/* debug [instance_methods/method]: EncodeBidirectionalSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865698-encodebidirectionalsequencetocom
func (r_ RNNMatrixInferenceLayer) EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices(commandBuffer unsafe.Pointer, sourceSequence unsafe.Pointer, destinationForwardMatrices unsafe.Pointer, destinationBackwardMatrices unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeBidirectionalSequenceToCommandBuffer:sourceSequence:destinationForwardMatrices:destinationBackwardMatrices:"), commandBuffer, sourceSequence, destinationForwardMatrices, destinationBackwardMatrices)
}/* debug [instance_methods/method]: EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865705-encodesequence
func (r_ RNNMatrixInferenceLayer) EncodeSequence() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeSequence"))
}/* debug [instance_methods/method]: EncodeSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865705-encodesequencetocommandbuffer
func (r_ RNNMatrixInferenceLayer) EncodeSequenceToCommandBufferSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, destinationMatrices unsafe.Pointer, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeSequenceToCommandBuffer:sourceMatrices:destinationMatrices:recurrentInputState:recurrentOutputStates:"), commandBuffer, sourceMatrices, destinationMatrices, recurrentInputState, recurrentOutputStates)
}/* debug [instance_methods/method]: EncodeSequenceToCommandBufferSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865746-copywithzone
func (r_ RNNMatrixInferenceLayer) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2966781-encodesequencetocommandbuffer
func (r_ RNNMatrixInferenceLayer) EncodeSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates(commandBuffer unsafe.Pointer, sourceMatrices unsafe.Pointer, sourceOffsets uint, destinationMatrices unsafe.Pointer, destinationOffsets uint, recurrentInputState IRNNRecurrentMatrixState, recurrentOutputStates unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeSequenceToCommandBuffer:sourceMatrices:sourceOffsets:destinationMatrices:destinationOffsets:recurrentInputState:recurrentOutputStates:"), commandBuffer, sourceMatrices, sourceOffsets, destinationMatrices, destinationOffsets, recurrentInputState, recurrentOutputStates)
}/* debug [instance_methods/method]: EncodeSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RNNMatrixInferenceLayer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865714-recurrentoutputistemporary
func (r_ RNNMatrixInferenceLayer) RecurrentOutputIsTemporary() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}/* debug [instance_properties/getter]: recurrentOutputIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865714-recurrentoutputistemporary
func (r_ RNNMatrixInferenceLayer) SetRecurrentOutputIsTemporary(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}/* debug [instance_properties/setter]: recurrentOutputIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865729-storeallintermediatestates
func (r_ RNNMatrixInferenceLayer) StoreAllIntermediateStates() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("storeAllIntermediateStates"))
	return rv
}/* debug [instance_properties/getter]: storeAllIntermediateStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865729-storeallintermediatestates
func (r_ RNNMatrixInferenceLayer) SetStoreAllIntermediateStates(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStoreAllIntermediateStates:"), value)
}/* debug [instance_properties/setter]: storeAllIntermediateStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865739-bidirectionalcombinemode
func (r_ RNNMatrixInferenceLayer) BidirectionalCombineMode() RNNBidirectionalCombineMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("bidirectionalCombineMode"))
	return rv
}/* debug [instance_properties/getter]: bidirectionalCombineMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2865739-bidirectionalcombinemode
func (r_ RNNMatrixInferenceLayer) SetBidirectionalCombineMode(value RNNBidirectionalCombineMode get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBidirectionalCombineMode:"), value)
}/* debug [instance_properties/setter]: bidirectionalCombineMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2873347-numberoflayers
func (r_ RNNMatrixInferenceLayer) NumberOfLayers() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("numberOfLayers"))
	return rv
}/* debug [instance_properties/getter]: numberOfLayers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2873347-numberoflayers
func (r_ RNNMatrixInferenceLayer) SetNumberOfLayers(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNumberOfLayers:"), value)
}/* debug [instance_properties/setter]: numberOfLayers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2890142-outputfeaturechannels
func (r_ RNNMatrixInferenceLayer) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2890142-outputfeaturechannels
func (r_ RNNMatrixInferenceLayer) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2890143-inputfeaturechannels
func (r_ RNNMatrixInferenceLayer) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixinferencelayer/2890143-inputfeaturechannels
func (r_ RNNMatrixInferenceLayer) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: inputFeatureChannels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSRNNMatrixInferenceLayer */


