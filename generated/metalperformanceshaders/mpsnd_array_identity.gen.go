// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NDArrayIdentity] class.
var (
	NDArrayIdentityClass     _NDArrayIdentityClass
	NDArrayIdentityClassOnce sync.Once
)

func getNDArrayIdentityClass() _NDArrayIdentityClass {
	NDArrayIdentityClassOnce.Do(func() {
		NDArrayIdentityClass = _NDArrayIdentityClass{objc.GetClass("MPSNDArrayIdentity")}
	})
	return NDArrayIdentityClass
}

type _NDArrayIdentityClass struct {
	class objc.Class
}





// An interface definition for the [NDArrayIdentity] class.
type INDArrayIdentity interface {
	INDArrayUnaryKernel
	

	// properties:


	

	// methods:
	Reshape()
	ReshapeWithCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray(cmdBuf unsafe.Pointer, sourceArray INDArray, numberOfDimensions uint, dimensionSizes uint, destinationArray INDArray) INDArray
	ReshapeWithCommandBufferSourceArrayShapeDestinationArray(cmdBuf unsafe.Pointer, sourceArray INDArray, shape Shape /* not a class type */, destinationArray INDArray) INDArray
	ReshapeWithCommandEncoderCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray(encoder unsafe.Pointer, cmdBuf unsafe.Pointer, sourceArray INDArray, numberOfDimensions uint, dimensionSizes uint, destinationArray INDArray) INDArray
	ReshapeWithCommandEncoderCommandBufferSourceArrayShapeDestinationArray(encoder unsafe.Pointer, cmdBuf unsafe.Pointer, sourceArray INDArray, shape Shape /* not a class type */, destinationArray INDArray) INDArray


}





// Alloc allocates a new instance without initialization.
func (nc _NDArrayIdentityClass) Alloc() NDArrayIdentity {
	rv := objc.Send[NDArrayIdentity](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayIdentityClass) New() NDArrayIdentity {
	rv := objc.Send[NDArrayIdentity](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayIdentity) Init() NDArrayIdentity {
	rv := objc.Send[NDArrayIdentity](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayIdentity) Autorelease() NDArrayIdentity {
	rv := objc.Send[NDArrayIdentity](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayIdentity creates a new NDArrayIdentity instance.
func NewNDArrayIdentity() NDArrayIdentity {
	return getNDArrayIdentityClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayIdentity
type NDArrayIdentity struct {
	NDArrayUnaryKernel
}

// NDArrayIdentityFrom constructs a [NDArrayIdentity] from an unsafe.Pointer.
func NDArrayIdentityFrom(ptr unsafe.Pointer) NDArrayIdentity {
	return NDArrayIdentity{
		NDArrayUnaryKernel: NDArrayUnaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayidentity/4438555-initwithdevice
func NewNDArrayIdentityWithDevice(device unsafe.Pointer) NDArrayIdentity {
	instance := getNDArrayIdentityClass().Alloc()
	rv := objc.Send[NDArrayIdentity](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayidentity/4438556-reshape
func (n_ NDArrayIdentity) Reshape() {
	objc.Send[objc.ID](n_.ID, objc.Sel("reshape"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayidentity/4438556-reshapewithcommandbuffer
func (n_ NDArrayIdentity) ReshapeWithCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray(cmdBuf unsafe.Pointer, sourceArray INDArray, numberOfDimensions uint, dimensionSizes uint, destinationArray INDArray) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("reshapeWithCommandBuffer:sourceArray:dimensionCount:dimensionSizes:destinationArray:"), cmdBuf, sourceArray, numberOfDimensions, dimensionSizes, destinationArray)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayidentity/4438557-reshapewithcommandbuffer
func (n_ NDArrayIdentity) ReshapeWithCommandBufferSourceArrayShapeDestinationArray(cmdBuf unsafe.Pointer, sourceArray INDArray, shape Shape /* not a class type */, destinationArray INDArray) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("reshapeWithCommandBuffer:sourceArray:shape:destinationArray:"), cmdBuf, sourceArray, shape, destinationArray)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayidentity/4438558-reshapewithcommandencoder
func (n_ NDArrayIdentity) ReshapeWithCommandEncoderCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray(encoder unsafe.Pointer, cmdBuf unsafe.Pointer, sourceArray INDArray, numberOfDimensions uint, dimensionSizes uint, destinationArray INDArray) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("reshapeWithCommandEncoder:commandBuffer:sourceArray:dimensionCount:dimensionSizes:destinationArray:"), encoder, cmdBuf, sourceArray, numberOfDimensions, dimensionSizes, destinationArray)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayidentity/4438559-reshapewithcommandencoder
func (n_ NDArrayIdentity) ReshapeWithCommandEncoderCommandBufferSourceArrayShapeDestinationArray(encoder unsafe.Pointer, cmdBuf unsafe.Pointer, sourceArray INDArray, shape Shape /* not a class type */, destinationArray INDArray) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("reshapeWithCommandEncoder:commandBuffer:sourceArray:shape:destinationArray:"), encoder, cmdBuf, sourceArray, shape, destinationArray)
	return rv
}












