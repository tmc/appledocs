// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [RNNImageInferenceLayer] class.
var (
	RNNImageInferenceLayerClass     _RNNImageInferenceLayerClass
	RNNImageInferenceLayerClassOnce sync.Once
)

func getRNNImageInferenceLayerClass() _RNNImageInferenceLayerClass {
	RNNImageInferenceLayerClassOnce.Do(func() {
		RNNImageInferenceLayerClass = _RNNImageInferenceLayerClass{objc.GetClass("MPSRNNImageInferenceLayer")}
	})
	return RNNImageInferenceLayerClass
}

type _RNNImageInferenceLayerClass struct {
	class objc.Class
}





// An interface definition for the [RNNImageInferenceLayer] class.
type IRNNImageInferenceLayer interface {
	ICNNKernel
	

	// properties:
	NumberOfLayers() objectivec.IObject
	SetNumberOfLayers(value objectivec.IObject)
	StoreAllIntermediateStates() objectivec.IObject
	SetStoreAllIntermediateStates(value objectivec.IObject)
	BidirectionalCombineMode() RNNBidirectionalCombineMode get set /* not a class type */
	SetBidirectionalCombineMode(value RNNBidirectionalCombineMode get set /* not a class type */)
	RecurrentOutputIsTemporary() objectivec.IObject
	SetRecurrentOutputIsTemporary(value objectivec.IObject)
	OutputFeatureChannels() objectivec.IObject
	SetOutputFeatureChannels(value objectivec.IObject)
	InputFeatureChannels() objectivec.IObject
	SetInputFeatureChannels(value objectivec.IObject)


	

	// methods:
	EncodeBidirectionalSequence()
	EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardImagesDestinationBackwardImages(commandBuffer unsafe.Pointer, sourceSequence unsafe.Pointer, destinationForwardImages unsafe.Pointer, destinationBackwardImages unsafe.Pointer)
	EncodeSequence()
	EncodeSequenceToCommandBufferSourceImagesDestinationImagesRecurrentInputStateRecurrentOutputStates(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer, destinationImages unsafe.Pointer, recurrentInputState IRNNRecurrentImageState, recurrentOutputStates unsafe.Pointer)
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject


}





// Alloc allocates a new instance without initialization.
func (rc _RNNImageInferenceLayerClass) Alloc() RNNImageInferenceLayer {
	rv := objc.Send[RNNImageInferenceLayer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RNNImageInferenceLayerClass) New() RNNImageInferenceLayer {
	rv := objc.Send[RNNImageInferenceLayer](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNImageInferenceLayer) Init() RNNImageInferenceLayer {
	rv := objc.Send[RNNImageInferenceLayer](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNImageInferenceLayer) Autorelease() RNNImageInferenceLayer {
	rv := objc.Send[RNNImageInferenceLayer](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNImageInferenceLayer creates a new RNNImageInferenceLayer instance.
func NewRNNImageInferenceLayer() RNNImageInferenceLayer {
	return getRNNImageInferenceLayerClass().New()
}





// A recurrent neural network layer for inference on Metal Performance Shaders images.
//
// The specifies a recurrent neural network layer for inference on objects. Two types of recurrent layers are supported: —Operates with convolutions on images. —Operates on matrices. You can use to implement the latter by using 1 x 1 images, but due to image size restrictions and performance, is the better choice for linear recurrent layers. is initialized using either of the following: A single instance, which further specifies the recurrent network layer. An array of instances, which specifies a stack of recurrent layers that can operate in parallel a subset of the inputs in a sequence of inputs and recurrent outputs. Stacks with bidirectionally traversing encode functions don’t support starting from a previous set of recurrent states. However, you can achieve this effect by defining two separate unidirectional stacks of layers, running the same input sequence on them separately (one forward and one backward), and ultimately combining the two result sequences.


// A recurrent neural network layer for inference on Metal Performance Shaders images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer
type RNNImageInferenceLayer struct {
	CNNKernel
}

// RNNImageInferenceLayerFrom constructs a [RNNImageInferenceLayer] from an unsafe.Pointer.
//
// A recurrent neural network layer for inference on Metal Performance Shaders images.
func RNNImageInferenceLayerFrom(ptr unsafe.Pointer) RNNImageInferenceLayer {
	return RNNImageInferenceLayer{
		CNNKernel: CNNKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865743-initwithcoder
func NewRNNImageInferenceLayerWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) RNNImageInferenceLayer {
	instance := getRNNImageInferenceLayerClass().Alloc()
	rv := objc.Send[RNNImageInferenceLayer](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865691-initwithdevice
func NewRNNImageInferenceLayerWithDeviceRnnDescriptor(device unsafe.Pointer, rnnDescriptor IRNNDescriptor) RNNImageInferenceLayer {
	instance := getRNNImageInferenceLayerClass().Alloc()
	rv := objc.Send[RNNImageInferenceLayer](instance.ID, objc.Sel("initWithDevice:rnnDescriptor:"), device, rnnDescriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865682-initwithdevice
func NewRNNImageInferenceLayerWithDeviceRnnDescriptors(device unsafe.Pointer, rnnDescriptors unsafe.Pointer) RNNImageInferenceLayer {
	instance := getRNNImageInferenceLayerClass().Alloc()
	rv := objc.Send[RNNImageInferenceLayer](instance.ID, objc.Sel("initWithDevice:rnnDescriptors:"), device, rnnDescriptors)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865693-encodebidirectionalsequence
func (r_ RNNImageInferenceLayer) EncodeBidirectionalSequence() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeBidirectionalSequence"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865693-encodebidirectionalsequencetocom
func (r_ RNNImageInferenceLayer) EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardImagesDestinationBackwardImages(commandBuffer unsafe.Pointer, sourceSequence unsafe.Pointer, destinationForwardImages unsafe.Pointer, destinationBackwardImages unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeBidirectionalSequenceToCommandBuffer:sourceSequence:destinationForwardImages:destinationBackwardImages:"), commandBuffer, sourceSequence, destinationForwardImages, destinationBackwardImages)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865717-encodesequence
func (r_ RNNImageInferenceLayer) EncodeSequence() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeSequence"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865717-encodesequencetocommandbuffer
func (r_ RNNImageInferenceLayer) EncodeSequenceToCommandBufferSourceImagesDestinationImagesRecurrentInputStateRecurrentOutputStates(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer, destinationImages unsafe.Pointer, recurrentInputState IRNNRecurrentImageState, recurrentOutputStates unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeSequenceToCommandBuffer:sourceImages:destinationImages:recurrentInputState:recurrentOutputStates:"), commandBuffer, sourceImages, destinationImages, recurrentInputState, recurrentOutputStates)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865728-copywithzone
func (r_ RNNImageInferenceLayer) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865697-numberoflayers
func (r_ RNNImageInferenceLayer) NumberOfLayers() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("numberOfLayers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865697-numberoflayers
func (r_ RNNImageInferenceLayer) SetNumberOfLayers(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNumberOfLayers:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865706-storeallintermediatestates
func (r_ RNNImageInferenceLayer) StoreAllIntermediateStates() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("storeAllIntermediateStates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865706-storeallintermediatestates
func (r_ RNNImageInferenceLayer) SetStoreAllIntermediateStates(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStoreAllIntermediateStates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865737-bidirectionalcombinemode
func (r_ RNNImageInferenceLayer) BidirectionalCombineMode() RNNBidirectionalCombineMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("bidirectionalCombineMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865737-bidirectionalcombinemode
func (r_ RNNImageInferenceLayer) SetBidirectionalCombineMode(value RNNBidirectionalCombineMode get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBidirectionalCombineMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865749-recurrentoutputistemporary
func (r_ RNNImageInferenceLayer) RecurrentOutputIsTemporary() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865749-recurrentoutputistemporary
func (r_ RNNImageInferenceLayer) SetRecurrentOutputIsTemporary(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2890140-outputfeaturechannels
func (r_ RNNImageInferenceLayer) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2890140-outputfeaturechannels
func (r_ RNNImageInferenceLayer) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2890141-inputfeaturechannels
func (r_ RNNImageInferenceLayer) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2890141-inputfeaturechannels
func (r_ RNNImageInferenceLayer) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputFeatureChannels:"), value)
}







