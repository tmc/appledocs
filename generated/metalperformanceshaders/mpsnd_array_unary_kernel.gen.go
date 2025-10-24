// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NDArrayUnaryKernel] class.
var (
	NDArrayUnaryKernelClass     _NDArrayUnaryKernelClass
	NDArrayUnaryKernelClassOnce sync.Once
)

func getNDArrayUnaryKernelClass() _NDArrayUnaryKernelClass {
	NDArrayUnaryKernelClassOnce.Do(func() {
		NDArrayUnaryKernelClass = _NDArrayUnaryKernelClass{objc.GetClass("MPSNDArrayUnaryKernel")}
	})
	return NDArrayUnaryKernelClass
}

type _NDArrayUnaryKernelClass struct {
	class objc.Class
}





// An interface definition for the [NDArrayUnaryKernel] class.
type INDArrayUnaryKernel interface {
	INDArrayMultiaryKernel
	

	// properties:
	DilationRates() NDArraySizes get /* not a class type */
	SetDilationRates(value NDArraySizes get /* not a class type */)
	EdgeMode() ImageEdgeMode get /* not a class type */
	SetEdgeMode(value ImageEdgeMode get /* not a class type */)
	KernelSizes() NDArraySizes get /* not a class type */
	SetKernelSizes(value NDArraySizes get /* not a class type */)
	Offsets() NDArrayOffsets get /* not a class type */
	SetOffsets(value NDArrayOffsets get /* not a class type */)
	Strides() NDArrayOffsets get /* not a class type */
	SetStrides(value NDArrayOffsets get /* not a class type */)


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceArray(cmdBuf unsafe.Pointer, sourceArray INDArray) INDArray
	EncodeToCommandBufferSourceArrayDestinationArray(cmdBuf unsafe.Pointer, sourceArray INDArray, destination INDArray)
	EncodeToCommandBufferSourceArrayResultStateDestinationArray(cmdBuf unsafe.Pointer, sourceArray INDArray, outGradientState IState, destination INDArray)
	EncodeToCommandBufferSourceArrayResultStateOutputStateIsTemporary(cmdBuf unsafe.Pointer, sourceArray INDArray, outGradientState objectivec.IObject, outputStateIsTemporary bool) INDArray


}





// Alloc allocates a new instance without initialization.
func (nc _NDArrayUnaryKernelClass) Alloc() NDArrayUnaryKernel {
	rv := objc.Send[NDArrayUnaryKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayUnaryKernelClass) New() NDArrayUnaryKernel {
	rv := objc.Send[NDArrayUnaryKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayUnaryKernel) Init() NDArrayUnaryKernel {
	rv := objc.Send[NDArrayUnaryKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayUnaryKernel) Autorelease() NDArrayUnaryKernel {
	rv := objc.Send[NDArrayUnaryKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayUnaryKernel creates a new NDArrayUnaryKernel instance.
func NewNDArrayUnaryKernel() NDArrayUnaryKernel {
	return getNDArrayUnaryKernelClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel
type NDArrayUnaryKernel struct {
	NDArrayMultiaryKernel
}

// NDArrayUnaryKernelFrom constructs a [NDArrayUnaryKernel] from an unsafe.Pointer.
func NDArrayUnaryKernelFrom(ptr unsafe.Pointer) NDArrayUnaryKernel {
	return NDArrayUnaryKernel{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3175012-initwithcoder
func NewNDArrayUnaryKernelWithCoderDevice(coder foundation.Coder, device unsafe.Pointer) NDArrayUnaryKernel {
	instance := getNDArrayUnaryKernelClass().Alloc()
	rv := objc.Send[NDArrayUnaryKernel](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143540-initwithdevice
func NewNDArrayUnaryKernelWithDevice(device unsafe.Pointer) NDArrayUnaryKernel {
	instance := getNDArrayUnaryKernelClass().Alloc()
	rv := objc.Send[NDArrayUnaryKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143536-encode
func (n_ NDArrayUnaryKernel) Encode() {
	objc.Send[objc.ID](n_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143536-encodetocommandbuffer
func (n_ NDArrayUnaryKernel) EncodeToCommandBufferSourceArray(cmdBuf unsafe.Pointer, sourceArray INDArray) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArray:"), cmdBuf, sourceArray)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143537-encodetocommandbuffer
func (n_ NDArrayUnaryKernel) EncodeToCommandBufferSourceArrayDestinationArray(cmdBuf unsafe.Pointer, sourceArray INDArray, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArray:destinationArray:"), cmdBuf, sourceArray, destination)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143538-encodetocommandbuffer
func (n_ NDArrayUnaryKernel) EncodeToCommandBufferSourceArrayResultStateDestinationArray(cmdBuf unsafe.Pointer, sourceArray INDArray, outGradientState IState, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArray:resultState:destinationArray:"), cmdBuf, sourceArray, outGradientState, destination)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143539-encodetocommandbuffer
func (n_ NDArrayUnaryKernel) EncodeToCommandBufferSourceArrayResultStateOutputStateIsTemporary(cmdBuf unsafe.Pointer, sourceArray INDArray, outGradientState objectivec.IObject, outputStateIsTemporary bool) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArray:resultState:outputStateIsTemporary:"), cmdBuf, sourceArray, outGradientState, outputStateIsTemporary)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143534-dilationrates
func (n_ NDArrayUnaryKernel) DilationRates() NDArraySizes get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("dilationRates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143534-dilationrates
func (n_ NDArrayUnaryKernel) SetDilationRates(value NDArraySizes get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDilationRates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143535-edgemode
func (n_ NDArrayUnaryKernel) EdgeMode() ImageEdgeMode get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("edgeMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143535-edgemode
func (n_ NDArrayUnaryKernel) SetEdgeMode(value ImageEdgeMode get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEdgeMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143541-kernelsizes
func (n_ NDArrayUnaryKernel) KernelSizes() NDArraySizes get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("kernelSizes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143541-kernelsizes
func (n_ NDArrayUnaryKernel) SetKernelSizes(value NDArraySizes get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setKernelSizes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143542-offsets
func (n_ NDArrayUnaryKernel) Offsets() NDArrayOffsets get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("offsets"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143542-offsets
func (n_ NDArrayUnaryKernel) SetOffsets(value NDArrayOffsets get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOffsets:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143543-strides
func (n_ NDArrayUnaryKernel) Strides() NDArrayOffsets get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("strides"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143543-strides
func (n_ NDArrayUnaryKernel) SetStrides(value NDArrayOffsets get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setStrides:"), value)
}







