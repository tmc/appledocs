// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NDArrayMultiaryGradientKernel] class.
var (
	NDArrayMultiaryGradientKernelClass     _NDArrayMultiaryGradientKernelClass
	NDArrayMultiaryGradientKernelClassOnce sync.Once
)

func getNDArrayMultiaryGradientKernelClass() _NDArrayMultiaryGradientKernelClass {
	NDArrayMultiaryGradientKernelClassOnce.Do(func() {
		NDArrayMultiaryGradientKernelClass = _NDArrayMultiaryGradientKernelClass{objc.GetClass("MPSNDArrayMultiaryGradientKernel")}
	})
	return NDArrayMultiaryGradientKernelClass
}

type _NDArrayMultiaryGradientKernelClass struct {
	class objc.Class
}





// An interface definition for the [NDArrayMultiaryGradientKernel] class.
type INDArrayMultiaryGradientKernel interface {
	INDArrayMultiaryBase
	

	// properties:


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceArraysSourceGradientGradientState(cmdBuf unsafe.Pointer, sources unsafe.Pointer, gradient INDArray, state IState) INDArray
	EncodeToCommandBufferSourceArraysSourceGradientGradientStateDestinationArray(cmdBuf unsafe.Pointer, sources unsafe.Pointer, gradient INDArray, state IState, destination INDArray)


}





// Alloc allocates a new instance without initialization.
func (nc _NDArrayMultiaryGradientKernelClass) Alloc() NDArrayMultiaryGradientKernel {
	rv := objc.Send[NDArrayMultiaryGradientKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayMultiaryGradientKernelClass) New() NDArrayMultiaryGradientKernel {
	rv := objc.Send[NDArrayMultiaryGradientKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayMultiaryGradientKernel) Init() NDArrayMultiaryGradientKernel {
	rv := objc.Send[NDArrayMultiaryGradientKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayMultiaryGradientKernel) Autorelease() NDArrayMultiaryGradientKernel {
	rv := objc.Send[NDArrayMultiaryGradientKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayMultiaryGradientKernel creates a new NDArrayMultiaryGradientKernel instance.
func NewNDArrayMultiaryGradientKernel() NDArrayMultiaryGradientKernel {
	return getNDArrayMultiaryGradientKernelClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel
type NDArrayMultiaryGradientKernel struct {
	NDArrayMultiaryBase
}

// NDArrayMultiaryGradientKernelFrom constructs a [NDArrayMultiaryGradientKernel] from an unsafe.Pointer.
func NDArrayMultiaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayMultiaryGradientKernel {
	return NDArrayMultiaryGradientKernel{
		NDArrayMultiaryBase: NDArrayMultiaryBaseFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3175008-initwithcoder
func NewNDArrayMultiaryGradientKernelWithCoderDevice(coder foundation.Coder, device unsafe.Pointer) NDArrayMultiaryGradientKernel {
	instance := getNDArrayMultiaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayMultiaryGradientKernel](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143524-initwithdevice
func NewNDArrayMultiaryGradientKernelWithDeviceSourceCountSourceGradientIndex(device unsafe.Pointer, count uint, sourceGradientIndex uint) NDArrayMultiaryGradientKernel {
	instance := getNDArrayMultiaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayMultiaryGradientKernel](instance.ID, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), device, count, sourceGradientIndex)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143522-encode
func (n_ NDArrayMultiaryGradientKernel) Encode() {
	objc.Send[objc.ID](n_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143522-encodetocommandbuffer
func (n_ NDArrayMultiaryGradientKernel) EncodeToCommandBufferSourceArraysSourceGradientGradientState(cmdBuf unsafe.Pointer, sources unsafe.Pointer, gradient INDArray, state IState) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:sourceGradient:gradientState:"), cmdBuf, sources, gradient, state)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143523-encodetocommandbuffer
func (n_ NDArrayMultiaryGradientKernel) EncodeToCommandBufferSourceArraysSourceGradientGradientStateDestinationArray(cmdBuf unsafe.Pointer, sources unsafe.Pointer, gradient INDArray, state IState, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:sourceGradient:gradientState:destinationArray:"), cmdBuf, sources, gradient, state, destination)
}












