// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	ReshapeWithCommandEncoderCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray(encoder objc.ID, cmdBuf objc.ID, sourceArray unsafe.Pointer, numberOfDimensions uint, dimensionSizes unsafe.Pointer, destinationArray unsafe.Pointer) unsafe.Pointer
}

//
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

// Alloc allocates a new instance without initialization.
func (nc _NDArrayIdentityClass) Alloc() NDArrayIdentity {
	rv := objc.Send[NDArrayIdentity](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayIdentity/reshape(with:commandBuffer:sourceArray:dimensionCount:dimensionSizes:destinationArray:)
func (n_ NDArrayIdentity) ReshapeWithCommandEncoderCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray(encoder objc.ID, cmdBuf objc.ID, sourceArray unsafe.Pointer, numberOfDimensions uint, dimensionSizes unsafe.Pointer, destinationArray unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("reshapeWithCommandEncoder:commandBuffer:sourceArray:dimensionCount:dimensionSizes:destinationArray:"), encoder, cmdBuf, sourceArray, numberOfDimensions, dimensionSizes, destinationArray)
	return rv
}
