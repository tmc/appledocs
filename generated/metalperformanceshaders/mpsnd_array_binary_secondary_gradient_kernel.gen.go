// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NDArrayBinarySecondaryGradientKernel] class.
var (
	NDArrayBinarySecondaryGradientKernelClass     _NDArrayBinarySecondaryGradientKernelClass
	NDArrayBinarySecondaryGradientKernelClassOnce sync.Once
)

func getNDArrayBinarySecondaryGradientKernelClass() _NDArrayBinarySecondaryGradientKernelClass {
	NDArrayBinarySecondaryGradientKernelClassOnce.Do(func() {
		NDArrayBinarySecondaryGradientKernelClass = _NDArrayBinarySecondaryGradientKernelClass{objc.GetClass("MPSNDArrayBinarySecondaryGradientKernel")}
	})
	return NDArrayBinarySecondaryGradientKernelClass
}

type _NDArrayBinarySecondaryGradientKernelClass struct {
	class objc.Class
}





// An interface definition for the [NDArrayBinarySecondaryGradientKernel] class.
type INDArrayBinarySecondaryGradientKernel interface {
	INDArrayMultiaryGradientKernel
	

	// properties:


	

	// methods:
	Encode()
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) INDArray
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState, destination INDArray)


}





// Alloc allocates a new instance without initialization.
func (nc _NDArrayBinarySecondaryGradientKernelClass) Alloc() NDArrayBinarySecondaryGradientKernel {
	rv := objc.Send[NDArrayBinarySecondaryGradientKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayBinarySecondaryGradientKernelClass) New() NDArrayBinarySecondaryGradientKernel {
	rv := objc.Send[NDArrayBinarySecondaryGradientKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayBinarySecondaryGradientKernel) Init() NDArrayBinarySecondaryGradientKernel {
	rv := objc.Send[NDArrayBinarySecondaryGradientKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayBinarySecondaryGradientKernel) Autorelease() NDArrayBinarySecondaryGradientKernel {
	rv := objc.Send[NDArrayBinarySecondaryGradientKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayBinarySecondaryGradientKernel creates a new NDArrayBinarySecondaryGradientKernel instance.
func NewNDArrayBinarySecondaryGradientKernel() NDArrayBinarySecondaryGradientKernel {
	return getNDArrayBinarySecondaryGradientKernelClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinarySecondaryGradientKernel
type NDArrayBinarySecondaryGradientKernel struct {
	NDArrayMultiaryGradientKernel
}

// NDArrayBinarySecondaryGradientKernelFrom constructs a [NDArrayBinarySecondaryGradientKernel] from an unsafe.Pointer.
func NDArrayBinarySecondaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayBinarySecondaryGradientKernel {
	return NDArrayBinarySecondaryGradientKernel{
		NDArrayMultiaryGradientKernel: NDArrayMultiaryGradientKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarysecondarygradientkernel/3175007-initwithcoder
func NewNDArrayBinarySecondaryGradientKernelWithCoderDevice(coder foundation.Coder, device unsafe.Pointer) NDArrayBinarySecondaryGradientKernel {
	instance := getNDArrayBinarySecondaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayBinarySecondaryGradientKernel](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarysecondarygradientkernel/3143519-initwithdevice
func NewNDArrayBinarySecondaryGradientKernelWithDevice(device unsafe.Pointer) NDArrayBinarySecondaryGradientKernel {
	instance := getNDArrayBinarySecondaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayBinarySecondaryGradientKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarysecondarygradientkernel/3143517-encode
func (n_ NDArrayBinarySecondaryGradientKernel) Encode() {
	objc.Send[objc.ID](n_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarysecondarygradientkernel/3143517-encodetocommandbuffer
func (n_ NDArrayBinarySecondaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:"), cmdBuf, primarySourceArray, secondarySourceArray, gradient, state)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarysecondarygradientkernel/3143518-encodetocommandbuffer
func (n_ NDArrayBinarySecondaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:destinationArray:"), cmdBuf, primarySourceArray, secondarySourceArray, gradient, state, destination)
}












