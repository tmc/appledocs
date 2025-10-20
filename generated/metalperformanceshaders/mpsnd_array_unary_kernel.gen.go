// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	EncodeToCommandBufferSourceArray(cmdBuf objc.ID, sourceArray unsafe.Pointer) unsafe.Pointer
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
func (n_ NDArrayUnaryKernel) EncodeToCommandBufferSourceArray(cmdBuf objc.ID, sourceArray unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArray:"), cmdBuf, sourceArray)
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



