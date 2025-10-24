// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNArithmetic] class.
var (
	CNNArithmeticClass     _CNNArithmeticClass
	CNNArithmeticClassOnce sync.Once
)

func getCNNArithmeticClass() _CNNArithmeticClass {
	CNNArithmeticClassOnce.Do(func() {
		CNNArithmeticClass = _CNNArithmeticClass{objc.GetClass("MPSCNNArithmetic")}
	})
	return CNNArithmeticClass
}

type _CNNArithmeticClass struct {
	class objc.Class
}





// An interface definition for the [CNNArithmetic] class.
type ICNNArithmetic interface {
	ICNNBinaryKernel
	

	// properties:
	SecondaryScale() objectivec.IObject
	SetSecondaryScale(value objectivec.IObject)
	MaximumValue() objectivec.IObject
	SetMaximumValue(value objectivec.IObject)
	Bias() objectivec.IObject
	SetBias(value objectivec.IObject)
	MinimumValue() objectivec.IObject
	SetMinimumValue(value objectivec.IObject)
	PrimaryScale() objectivec.IObject
	SetPrimaryScale(value objectivec.IObject)
	PrimaryStrideInFeatureChannels() objectivec.IObject
	SetPrimaryStrideInFeatureChannels(value objectivec.IObject)
	SecondaryStrideInFeatureChannels() objectivec.IObject
	SetSecondaryStrideInFeatureChannels(value objectivec.IObject)


	

	// methods:
	Encode()
	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, destinationState ICNNArithmeticGradientState, destinationImage IImage)
	EncodeBatch()
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationImages(commandBuffer unsafe.Pointer, primaryImages ImageBatch /* not a class type */, secondaryImages ImageBatch /* not a class type */, destinationStates CNNArithmeticGradientStateBatch /* not a class type */, destinationImages ImageBatch /* not a class type */)


}





// Alloc allocates a new instance without initialization.
func (cc _CNNArithmeticClass) Alloc() CNNArithmetic {
	rv := objc.Send[CNNArithmetic](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNArithmeticClass) New() CNNArithmetic {
	rv := objc.Send[CNNArithmetic](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNArithmetic) Init() CNNArithmetic {
	rv := objc.Send[CNNArithmetic](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNArithmetic) Autorelease() CNNArithmetic {
	rv := objc.Send[CNNArithmetic](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNArithmetic creates a new CNNArithmetic instance.
func NewCNNArithmetic() CNNArithmetic {
	return getCNNArithmeticClass().New()
}





// The base class for arithmetic operators.


// The base class for arithmetic operators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic
type CNNArithmetic struct {
	CNNBinaryKernel
}

// CNNArithmeticFrom constructs a [CNNArithmetic] from an unsafe.Pointer.
//
// The base class for arithmetic operators.
func CNNArithmeticFrom(ptr unsafe.Pointer) CNNArithmetic {
	return CNNArithmetic{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2954876-encode
func (c_ CNNArithmetic) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2954876-encodetocommandbuffer
func (c_ CNNArithmetic) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, destinationState ICNNArithmeticGradientState, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationState:destinationImage:"), commandBuffer, primaryImage, secondaryImage, destinationState, destinationImage)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2954877-encodebatch
func (c_ CNNArithmetic) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2954877-encodebatchtocommandbuffer
func (c_ CNNArithmetic) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationImages(commandBuffer unsafe.Pointer, primaryImages ImageBatch /* not a class type */, secondaryImages ImageBatch /* not a class type */, destinationStates CNNArithmeticGradientStateBatch /* not a class type */, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationStates:destinationImages:"), commandBuffer, primaryImages, secondaryImages, destinationStates, destinationImages)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942497-secondaryscale
func (c_ CNNArithmetic) SecondaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryScale"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942497-secondaryscale
func (c_ CNNArithmetic) SetSecondaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryScale:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942498-maximumvalue
func (c_ CNNArithmetic) MaximumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("maximumValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942498-maximumvalue
func (c_ CNNArithmetic) SetMaximumValue(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942499-bias
func (c_ CNNArithmetic) Bias() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("bias"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942499-bias
func (c_ CNNArithmetic) SetBias(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBias:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942502-minimumvalue
func (c_ CNNArithmetic) MinimumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("minimumValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942502-minimumvalue
func (c_ CNNArithmetic) SetMinimumValue(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942509-primaryscale
func (c_ CNNArithmetic) PrimaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryScale"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942509-primaryscale
func (c_ CNNArithmetic) SetPrimaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryScale:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2947963-primarystrideinfeaturechannels
func (c_ CNNArithmetic) PrimaryStrideInFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryStrideInFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2947963-primarystrideinfeaturechannels
func (c_ CNNArithmetic) SetPrimaryStrideInFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryStrideInFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2947964-secondarystrideinfeaturechannels
func (c_ CNNArithmetic) SecondaryStrideInFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2947964-secondarystrideinfeaturechannels
func (c_ CNNArithmetic) SetSecondaryStrideInFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}








