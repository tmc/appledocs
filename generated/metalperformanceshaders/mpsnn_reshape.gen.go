// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Reshape] class.
var (
	ReshapeClass     _ReshapeClass
	ReshapeClassOnce sync.Once
)

func getReshapeClass() _ReshapeClass {
	ReshapeClassOnce.Do(func() {
		ReshapeClass = _ReshapeClass{objc.GetClass("MPSNNReshape")}
	})
	return ReshapeClass
}

type _ReshapeClass struct {
	class objc.Class
}





// An interface definition for the [Reshape] class.
type IReshape interface {
	ICNNKernel
	

	// properties:


	

	// methods:
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, outStates StateBatch /* not a class type */, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) ImageBatch /* not a class type */
	EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) ImageBatch /* not a class type */
	Encode()
	EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImage IImage, outState objectivec.IObject, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IImage
	EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImage IImage, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IImage


}





// Alloc allocates a new instance without initialization.
func (rc _ReshapeClass) Alloc() Reshape {
	rv := objc.Send[Reshape](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReshapeClass) New() Reshape {
	rv := objc.Send[Reshape](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Reshape) Init() Reshape {
	rv := objc.Send[Reshape](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Reshape) Autorelease() Reshape {
	rv := objc.Send[Reshape](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReshape creates a new Reshape instance.
func NewReshape() Reshape {
	return getReshapeClass().New()
}





// The base class for reshape operations.


// The base class for reshape operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshape
type Reshape struct {
	CNNKernel
}

// ReshapeFrom constructs a [Reshape] from an unsafe.Pointer.
//
// The base class for reshape operations.
func ReshapeFrom(ptr unsafe.Pointer) Reshape {
	return Reshape{
		CNNKernel: CNNKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/2951930-initwithcoder
func NewReshapeWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) Reshape {
	instance := getReshapeClass().Alloc()
	rv := objc.Send[Reshape](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/2951929-initwithdevice
func NewReshapeWithDevice(device unsafe.Pointer) Reshape {
	instance := getReshapeClass().Alloc()
	rv := objc.Send[Reshape](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547989-encodebatch
func (r_ Reshape) EncodeBatch() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547989-encodebatchtocommandbuffer
func (r_ Reshape) EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, outStates StateBatch /* not a class type */, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](r_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImages, outStates, isTemporary, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547990-encodebatchtocommandbuffer
func (r_ Reshape) EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](r_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImages, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547991-encode
func (r_ Reshape) Encode() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547991-encodetocommandbuffer
func (r_ Reshape) EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImage IImage, outState objectivec.IObject, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationState:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImage, outState, isTemporary, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547992-encodetocommandbuffer
func (r_ Reshape) EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImage IImage, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImage, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}












