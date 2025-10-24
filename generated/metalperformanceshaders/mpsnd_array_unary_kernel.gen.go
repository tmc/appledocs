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
	EdgeMode() ImageEdgeMode
	Strides() NDArrayOffsets /* not a class type */
	DilationRates() NDArraySizes /* not a class type */
	SetDilationRates(value NDArraySizes /* not a class type */)
	KernelSizes() NDArraySizes /* not a class type */
	SetKernelSizes(value NDArraySizes /* not a class type */)
	Offsets() NDArrayOffsets /* not a class type */
	SetOffsets(value NDArrayOffsets /* not a class type */)
	// methods:
	EncodeToCommandBufferSourceArray(cmdBuf objectivec.IObject, sourceArray IMPSNDArray) INDArray
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

// Alloc allocates a new instance without initialization.
func (nc _NDArrayUnaryKernelClass) Alloc() NDArrayUnaryKernel {
	rv := objc.Send[NDArrayUnaryKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/encode(to:sourceArray:)
func (n_ NDArrayUnaryKernel) EncodeToCommandBufferSourceArray(cmdBuf objectivec.IObject, sourceArray IMPSNDArray) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArray:"), cmdBuf, sourceArray)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/edgeMode
func (n_ NDArrayUnaryKernel) EdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](n_.ID, objc.Sel("edgeMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/strides
func (n_ NDArrayUnaryKernel) Strides() NDArrayOffsets /* not a class type */ {
	rv := objc.Send[NDArrayOffsets](n_.ID, objc.Sel("strides"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/dilationrates
func (n_ NDArrayUnaryKernel) DilationRates() NDArraySizes /* not a class type */ {
	rv := objc.Send[NDArraySizes](n_.ID, objc.Sel("dilationRates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/dilationrates
func (n_ NDArrayUnaryKernel) SetDilationRates(value NDArraySizes /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDilationRates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/kernelsizes
func (n_ NDArrayUnaryKernel) KernelSizes() NDArraySizes /* not a class type */ {
	rv := objc.Send[NDArraySizes](n_.ID, objc.Sel("kernelSizes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/kernelsizes
func (n_ NDArrayUnaryKernel) SetKernelSizes(value NDArraySizes /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setKernelSizes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/offsets
func (n_ NDArrayUnaryKernel) Offsets() NDArrayOffsets /* not a class type */ {
	rv := objc.Send[NDArrayOffsets](n_.ID, objc.Sel("offsets"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/offsets
func (n_ NDArrayUnaryKernel) SetOffsets(value NDArrayOffsets /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOffsets:"), value)
}



