// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayBinaryPrimaryGradientKernel] class.
var (
	NDArrayBinaryPrimaryGradientKernelClass     _NDArrayBinaryPrimaryGradientKernelClass
	NDArrayBinaryPrimaryGradientKernelClassOnce sync.Once
)

func getNDArrayBinaryPrimaryGradientKernelClass() _NDArrayBinaryPrimaryGradientKernelClass {
	NDArrayBinaryPrimaryGradientKernelClassOnce.Do(func() {
		NDArrayBinaryPrimaryGradientKernelClass = _NDArrayBinaryPrimaryGradientKernelClass{objc.GetClass("MPSNDArrayBinaryPrimaryGradientKernel")}
	})
	return NDArrayBinaryPrimaryGradientKernelClass
}

type _NDArrayBinaryPrimaryGradientKernelClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayBinaryPrimaryGradientKernel] class.
type INDArrayBinaryPrimaryGradientKernel interface {
	INDArrayMultiaryGradientKernel
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf objc.ID, primarySourceArray unsafe.Pointer, secondarySourceArray unsafe.Pointer, gradient unsafe.Pointer, state unsafe.Pointer) unsafe.Pointer
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray(cmdBuf objc.ID, primarySourceArray unsafe.Pointer, secondarySourceArray unsafe.Pointer, gradient unsafe.Pointer, state unsafe.Pointer, destination unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel
type NDArrayBinaryPrimaryGradientKernel struct {
	NDArrayMultiaryGradientKernel
}

// NDArrayBinaryPrimaryGradientKernelFrom constructs a [NDArrayBinaryPrimaryGradientKernel] from an unsafe.Pointer.
func NDArrayBinaryPrimaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayBinaryPrimaryGradientKernel {
	return NDArrayBinaryPrimaryGradientKernel{
		NDArrayMultiaryGradientKernel: NDArrayMultiaryGradientKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayBinaryPrimaryGradientKernelClass) Alloc() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[NDArrayBinaryPrimaryGradientKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayBinaryPrimaryGradientKernelClass) New() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[NDArrayBinaryPrimaryGradientKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayBinaryPrimaryGradientKernel) Init() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[NDArrayBinaryPrimaryGradientKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayBinaryPrimaryGradientKernel) Autorelease() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[NDArrayBinaryPrimaryGradientKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayBinaryPrimaryGradientKernel creates a new NDArrayBinaryPrimaryGradientKernel instance.
func NewNDArrayBinaryPrimaryGradientKernel() NDArrayBinaryPrimaryGradientKernel {
	return getNDArrayBinaryPrimaryGradientKernelClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel/encode(to:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:)
func (n_ NDArrayBinaryPrimaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf objc.ID, primarySourceArray unsafe.Pointer, secondarySourceArray unsafe.Pointer, gradient unsafe.Pointer, state unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:"), cmdBuf, primarySourceArray, secondarySourceArray, gradient, state)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel/encode(to:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:destinationArray:)
func (n_ NDArrayBinaryPrimaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray(cmdBuf objc.ID, primarySourceArray unsafe.Pointer, secondarySourceArray unsafe.Pointer, gradient unsafe.Pointer, state unsafe.Pointer, destination unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:destinationArray:"), cmdBuf, primarySourceArray, secondarySourceArray, gradient, state, destination)
}
