// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NDArrayMultiaryKernel] class.
var (
	NDArrayMultiaryKernelClass     _NDArrayMultiaryKernelClass
	NDArrayMultiaryKernelClassOnce sync.Once
)

func getNDArrayMultiaryKernelClass() _NDArrayMultiaryKernelClass {
	NDArrayMultiaryKernelClassOnce.Do(func() {
		NDArrayMultiaryKernelClass = _NDArrayMultiaryKernelClass{objc.GetClass("MPSNDArrayMultiaryKernel")}
	})
	return NDArrayMultiaryKernelClass
}

type _NDArrayMultiaryKernelClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayMultiaryKernel] class.
type INDArrayMultiaryKernel interface {
	INDArrayMultiaryBase
	EncodeToCommandBufferSourceArraysDestinationArray(cmdBuf objectivec.IObject, sourceArrays []NDArray, destination IMPSNDArray)
	EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary(cmdBuf objectivec.IObject, sourceArrays []NDArray, outGradientState MPSState, outputStateIsTemporary bool) NDArray
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel
type NDArrayMultiaryKernel struct {
	NDArrayMultiaryBase
}

// NDArrayMultiaryKernelFrom constructs a [NDArrayMultiaryKernel] from an unsafe.Pointer.
func NDArrayMultiaryKernelFrom(ptr unsafe.Pointer) NDArrayMultiaryKernel {
	return NDArrayMultiaryKernel{
		NDArrayMultiaryBase: NDArrayMultiaryBaseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayMultiaryKernelClass) Alloc() NDArrayMultiaryKernel {
	rv := objc.Send[NDArrayMultiaryKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayMultiaryKernelClass) New() NDArrayMultiaryKernel {
	rv := objc.Send[NDArrayMultiaryKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayMultiaryKernel) Init() NDArrayMultiaryKernel {
	rv := objc.Send[NDArrayMultiaryKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayMultiaryKernel) Autorelease() NDArrayMultiaryKernel {
	rv := objc.Send[NDArrayMultiaryKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayMultiaryKernel creates a new NDArrayMultiaryKernel instance.
func NewNDArrayMultiaryKernel() NDArrayMultiaryKernel {
	return getNDArrayMultiaryKernelClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/encode(to:sourceArrays:destinationArray:)
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysDestinationArray(cmdBuf objectivec.IObject, sourceArrays []NDArray, destination IMPSNDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:destinationArray:"), cmdBuf, sourceArrays, destination)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/encode(to:sourceArrays:resultState:outputStateIsTemporary:)
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary(cmdBuf objectivec.IObject, sourceArrays []NDArray, outGradientState MPSState, outputStateIsTemporary bool) NDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:outputStateIsTemporary:"), cmdBuf, sourceArrays, outGradientState, outputStateIsTemporary)
	return rv
}



