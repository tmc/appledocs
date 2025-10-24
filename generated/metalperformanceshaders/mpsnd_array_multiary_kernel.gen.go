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
	

	// properties:


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceArrays(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer) INDArray
	EncodeToCommandBufferSourceArraysDestinationArray(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, destination INDArray)
	EncodeToCommandBufferSourceArraysResultStateDestinationArray(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, outGradientState IState, destination INDArray)
	EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, outGradientState objectivec.IObject, outputStateIsTemporary bool) INDArray
	EncodeToCommandEncoderCommandBufferSourceArraysDestinationArray(encoder unsafe.Pointer, commandBuffer unsafe.Pointer, sourceArrays unsafe.Pointer, destination INDArray)


}





// Alloc allocates a new instance without initialization.
func (nc _NDArrayMultiaryKernelClass) Alloc() NDArrayMultiaryKernel {
	rv := objc.Send[NDArrayMultiaryKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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







// [Full Topic]
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3175009-initwithcoder
func NewNDArrayMultiaryKernelWithCoderDevice(coder foundation.Coder, device unsafe.Pointer) NDArrayMultiaryKernel {
	instance := getNDArrayMultiaryKernelClass().Alloc()
	rv := objc.Send[NDArrayMultiaryKernel](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3175010-initwithdevice
func NewNDArrayMultiaryKernelWithDeviceSourceCount(device unsafe.Pointer, count uint) NDArrayMultiaryKernel {
	instance := getNDArrayMultiaryKernelClass().Alloc()
	rv := objc.Send[NDArrayMultiaryKernel](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143525-encode
func (n_ NDArrayMultiaryKernel) Encode() {
	objc.Send[objc.ID](n_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143525-encodetocommandbuffer
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArrays(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:"), cmdBuf, sourceArrays)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143526-encodetocommandbuffer
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysDestinationArray(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:destinationArray:"), cmdBuf, sourceArrays, destination)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143527-encodetocommandbuffer
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysResultStateDestinationArray(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, outGradientState IState, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:destinationArray:"), cmdBuf, sourceArrays, outGradientState, destination)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143528-encodetocommandbuffer
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, outGradientState objectivec.IObject, outputStateIsTemporary bool) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:outputStateIsTemporary:"), cmdBuf, sourceArrays, outGradientState, outputStateIsTemporary)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/4462738-encodetocommandencoder
func (n_ NDArrayMultiaryKernel) EncodeToCommandEncoderCommandBufferSourceArraysDestinationArray(encoder unsafe.Pointer, commandBuffer unsafe.Pointer, sourceArrays unsafe.Pointer, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandEncoder:commandBuffer:sourceArrays:destinationArray:"), encoder, commandBuffer, sourceArrays, destination)
}












