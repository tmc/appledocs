// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

//
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

// Alloc allocates a new instance without initialization.
func (nc _NDArrayBinarySecondaryGradientKernelClass) Alloc() NDArrayBinarySecondaryGradientKernel {
	rv := objc.Send[NDArrayBinarySecondaryGradientKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinarySecondaryGradientKernel/init(device:)
func NewNDArrayBinarySecondaryGradientKernelWithDevice(device objc.ID) NDArrayBinarySecondaryGradientKernel {
	instance := getNDArrayBinarySecondaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayBinarySecondaryGradientKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



