// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayVectorLUTDequantize] class.
var (
	NDArrayVectorLUTDequantizeClass     _NDArrayVectorLUTDequantizeClass
	NDArrayVectorLUTDequantizeClassOnce sync.Once
)

func getNDArrayVectorLUTDequantizeClass() _NDArrayVectorLUTDequantizeClass {
	NDArrayVectorLUTDequantizeClassOnce.Do(func() {
		NDArrayVectorLUTDequantizeClass = _NDArrayVectorLUTDequantizeClass{objc.GetClass("MPSNDArrayVectorLUTDequantize")}
	})
	return NDArrayVectorLUTDequantizeClass
}

type _NDArrayVectorLUTDequantizeClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayVectorLUTDequantize] class.
type INDArrayVectorLUTDequantize interface {
	INDArrayMultiaryKernel
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayVectorLUTDequantize
type NDArrayVectorLUTDequantize struct {
	NDArrayMultiaryKernel
}

// NDArrayVectorLUTDequantizeFrom constructs a [NDArrayVectorLUTDequantize] from an unsafe.Pointer.
func NDArrayVectorLUTDequantizeFrom(ptr unsafe.Pointer) NDArrayVectorLUTDequantize {
	return NDArrayVectorLUTDequantize{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayVectorLUTDequantizeClass) Alloc() NDArrayVectorLUTDequantize {
	rv := objc.Send[NDArrayVectorLUTDequantize](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayVectorLUTDequantizeClass) New() NDArrayVectorLUTDequantize {
	rv := objc.Send[NDArrayVectorLUTDequantize](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayVectorLUTDequantize) Init() NDArrayVectorLUTDequantize {
	rv := objc.Send[NDArrayVectorLUTDequantize](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayVectorLUTDequantize) Autorelease() NDArrayVectorLUTDequantize {
	rv := objc.Send[NDArrayVectorLUTDequantize](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayVectorLUTDequantize creates a new NDArrayVectorLUTDequantize instance.
func NewNDArrayVectorLUTDequantize() NDArrayVectorLUTDequantize {
	return getNDArrayVectorLUTDequantizeClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayVectorLUTDequantize/init(device:axis:)
func NewNDArrayVectorLUTDequantizeWithDeviceAxis(device objc.ID, axis uint) NDArrayVectorLUTDequantize {
	instance := getNDArrayVectorLUTDequantizeClass().Alloc()
	rv := objc.Send[NDArrayVectorLUTDequantize](instance.ID, objc.Sel("initWithDevice:axis:"), device, axis)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayVectorLUTDequantize/vectorAxis
func (n_ NDArrayVectorLUTDequantize) VectorAxis() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("vectorAxis"))
	return rv
}

// SetVectorAxis sets the value of the vectorAxis property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayVectorLUTDequantize/vectorAxis
func (n_ NDArrayVectorLUTDequantize) SetVectorAxis(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setVectorAxis:"), value)
}
