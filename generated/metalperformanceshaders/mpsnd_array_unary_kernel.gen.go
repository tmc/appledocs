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
	EncodeToCommandBufferSourceArray(cmdBuf objectivec.IObject, sourceArray IMPSNDArray) NDArray
	EdgeMode() unsafe.Pointer
	Strides() unsafe.Pointer
	DilationRates() unsafe.Pointer
	SetDilationRates(value unsafe.Pointer)
	KernelSizes() unsafe.Pointer
	SetKernelSizes(value unsafe.Pointer)
	Offsets() unsafe.Pointer
	SetOffsets(value unsafe.Pointer)
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/encode(to:sourceArray:)
func (n_ NDArrayUnaryKernel) EncodeToCommandBufferSourceArray(cmdBuf objectivec.IObject, sourceArray IMPSNDArray) NDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArray:"), cmdBuf, sourceArray)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/edgeMode
func (n_ NDArrayUnaryKernel) EdgeMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("edgeMode"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/strides
func (n_ NDArrayUnaryKernel) Strides() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("strides"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/dilationrates
func (n_ NDArrayUnaryKernel) DilationRates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("dilationRates"))
	return rv
}


// SetDilationRates sets the value of the dilationRates property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/dilationrates
func (n_ NDArrayUnaryKernel) SetDilationRates(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDilationRates:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/kernelsizes
func (n_ NDArrayUnaryKernel) KernelSizes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("kernelSizes"))
	return rv
}


// SetKernelSizes sets the value of the kernelSizes property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/kernelsizes
func (n_ NDArrayUnaryKernel) SetKernelSizes(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setKernelSizes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/offsets
func (n_ NDArrayUnaryKernel) Offsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("offsets"))
	return rv
}


// SetOffsets sets the value of the offsets property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/offsets
func (n_ NDArrayUnaryKernel) SetOffsets(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOffsets:"), value)
}



