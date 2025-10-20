// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

//
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

// Alloc allocates a new instance without initialization.
func (nc _NDArrayMultiaryGradientKernelClass) Alloc() NDArrayMultiaryGradientKernel {
	rv := objc.Send[NDArrayMultiaryGradientKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/init(coder:device:)
func NewNDArrayMultiaryGradientKernelWithCoderDevice(coder unsafe.Pointer, device objc.ID) NDArrayMultiaryGradientKernel {
	instance := getNDArrayMultiaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayMultiaryGradientKernel](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}
