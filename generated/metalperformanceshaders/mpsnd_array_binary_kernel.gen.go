// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayBinaryKernel] class.
var (
	NDArrayBinaryKernelClass     _NDArrayBinaryKernelClass
	NDArrayBinaryKernelClassOnce sync.Once
)

func getNDArrayBinaryKernelClass() _NDArrayBinaryKernelClass {
	NDArrayBinaryKernelClassOnce.Do(func() {
		NDArrayBinaryKernelClass = _NDArrayBinaryKernelClass{objc.GetClass("MPSNDArrayBinaryKernel")}
	})
	return NDArrayBinaryKernelClass
}

type _NDArrayBinaryKernelClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayBinaryKernel] class.
type INDArrayBinaryKernel interface {
	INDArrayMultiaryKernel
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArray(cmdBuf objc.ID, primarySourceArray unsafe.Pointer, secondarySourceArray unsafe.Pointer) unsafe.Pointer
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel
type NDArrayBinaryKernel struct {
	NDArrayMultiaryKernel
}

// NDArrayBinaryKernelFrom constructs a [NDArrayBinaryKernel] from an unsafe.Pointer.
func NDArrayBinaryKernelFrom(ptr unsafe.Pointer) NDArrayBinaryKernel {
	return NDArrayBinaryKernel{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayBinaryKernelClass) Alloc() NDArrayBinaryKernel {
	rv := objc.Send[NDArrayBinaryKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayBinaryKernelClass) New() NDArrayBinaryKernel {
	rv := objc.Send[NDArrayBinaryKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayBinaryKernel) Init() NDArrayBinaryKernel {
	rv := objc.Send[NDArrayBinaryKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayBinaryKernel) Autorelease() NDArrayBinaryKernel {
	rv := objc.Send[NDArrayBinaryKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayBinaryKernel creates a new NDArrayBinaryKernel instance.
func NewNDArrayBinaryKernel() NDArrayBinaryKernel {
	return getNDArrayBinaryKernelClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/init(device:)
func NewNDArrayBinaryKernelWithDevice(device objc.ID) NDArrayBinaryKernel {
	instance := getNDArrayBinaryKernelClass().Alloc()
	rv := objc.Send[NDArrayBinaryKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/encode(to:primarySourceArray:secondarySourceArray:)
func (n_ NDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArray(cmdBuf objc.ID, primarySourceArray unsafe.Pointer, secondarySourceArray unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:"), cmdBuf, primarySourceArray, secondarySourceArray)
	return rv
}


