// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayLUTDequantize] class.
var (
	NDArrayLUTDequantizeClass     _NDArrayLUTDequantizeClass
	NDArrayLUTDequantizeClassOnce sync.Once
)

func getNDArrayLUTDequantizeClass() _NDArrayLUTDequantizeClass {
	NDArrayLUTDequantizeClassOnce.Do(func() {
		NDArrayLUTDequantizeClass = _NDArrayLUTDequantizeClass{objc.GetClass("MPSNDArrayLUTDequantize")}
	})
	return NDArrayLUTDequantizeClass
}

type _NDArrayLUTDequantizeClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayLUTDequantize] class.
type INDArrayLUTDequantize interface {
	INDArrayMultiaryKernel
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTDequantize
type NDArrayLUTDequantize struct {
	NDArrayMultiaryKernel
}

// NDArrayLUTDequantizeFrom constructs a [NDArrayLUTDequantize] from an unsafe.Pointer.
func NDArrayLUTDequantizeFrom(ptr unsafe.Pointer) NDArrayLUTDequantize {
	return NDArrayLUTDequantize{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayLUTDequantizeClass) Alloc() NDArrayLUTDequantize {
	rv := objc.Send[NDArrayLUTDequantize](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayLUTDequantizeClass) New() NDArrayLUTDequantize {
	rv := objc.Send[NDArrayLUTDequantize](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayLUTDequantize) Init() NDArrayLUTDequantize {
	rv := objc.Send[NDArrayLUTDequantize](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayLUTDequantize) Autorelease() NDArrayLUTDequantize {
	rv := objc.Send[NDArrayLUTDequantize](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayLUTDequantize creates a new NDArrayLUTDequantize instance.
func NewNDArrayLUTDequantize() NDArrayLUTDequantize {
	return getNDArrayLUTDequantizeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTDequantize/init(device:)
func NewNDArrayLUTDequantizeWithDevice(device objc.ID) NDArrayLUTDequantize {
	instance := getNDArrayLUTDequantizeClass().Alloc()
	rv := objc.Send[NDArrayLUTDequantize](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



