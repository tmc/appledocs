// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NDArrayUnaryGradientKernel] class.
var (
	NDArrayUnaryGradientKernelClass     _NDArrayUnaryGradientKernelClass
	NDArrayUnaryGradientKernelClassOnce sync.Once
)

func getNDArrayUnaryGradientKernelClass() _NDArrayUnaryGradientKernelClass {
	NDArrayUnaryGradientKernelClassOnce.Do(func() {
		NDArrayUnaryGradientKernelClass = _NDArrayUnaryGradientKernelClass{objc.GetClass("MPSNDArrayUnaryGradientKernel")}
	})
	return NDArrayUnaryGradientKernelClass
}

type _NDArrayUnaryGradientKernelClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayUnaryGradientKernel] class.
type INDArrayUnaryGradientKernel interface {
	INDArrayMultiaryGradientKernel
	EncodeToCommandBufferSourceArraySourceGradientGradientStateDestinationArray(cmdBuf objectivec.IObject, sourceArray IMPSNDArray, gradient IMPSNDArray, state MPSState, destination IMPSNDArray)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel
type NDArrayUnaryGradientKernel struct {
	NDArrayMultiaryGradientKernel
}

// NDArrayUnaryGradientKernelFrom constructs a [NDArrayUnaryGradientKernel] from an unsafe.Pointer.
func NDArrayUnaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayUnaryGradientKernel {
	return NDArrayUnaryGradientKernel{
		NDArrayMultiaryGradientKernel: NDArrayMultiaryGradientKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayUnaryGradientKernelClass) Alloc() NDArrayUnaryGradientKernel {
	rv := objc.Send[NDArrayUnaryGradientKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayUnaryGradientKernelClass) New() NDArrayUnaryGradientKernel {
	rv := objc.Send[NDArrayUnaryGradientKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayUnaryGradientKernel) Init() NDArrayUnaryGradientKernel {
	rv := objc.Send[NDArrayUnaryGradientKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayUnaryGradientKernel) Autorelease() NDArrayUnaryGradientKernel {
	rv := objc.Send[NDArrayUnaryGradientKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayUnaryGradientKernel creates a new NDArrayUnaryGradientKernel instance.
func NewNDArrayUnaryGradientKernel() NDArrayUnaryGradientKernel {
	return getNDArrayUnaryGradientKernelClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel/init(device:)
func NewNDArrayUnaryGradientKernelWithDevice(device objectivec.IObject) NDArrayUnaryGradientKernel {
	instance := getNDArrayUnaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayUnaryGradientKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel/encode(to:sourceArray:sourceGradient:gradientState:destinationArray:)
func (n_ NDArrayUnaryGradientKernel) EncodeToCommandBufferSourceArraySourceGradientGradientStateDestinationArray(cmdBuf objectivec.IObject, sourceArray IMPSNDArray, gradient IMPSNDArray, state MPSState, destination IMPSNDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArray:sourceGradient:gradientState:destinationArray:"), cmdBuf, sourceArray, gradient, state, destination)
}


