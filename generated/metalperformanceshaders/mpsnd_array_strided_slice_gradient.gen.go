// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayStridedSliceGradient] class.
var (
	NDArrayStridedSliceGradientClass     _NDArrayStridedSliceGradientClass
	NDArrayStridedSliceGradientClassOnce sync.Once
)

func getNDArrayStridedSliceGradientClass() _NDArrayStridedSliceGradientClass {
	NDArrayStridedSliceGradientClassOnce.Do(func() {
		NDArrayStridedSliceGradientClass = _NDArrayStridedSliceGradientClass{objc.GetClass("MPSNDArrayStridedSliceGradient")}
	})
	return NDArrayStridedSliceGradientClass
}

type _NDArrayStridedSliceGradientClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayStridedSliceGradient] class.
type INDArrayStridedSliceGradient interface {
	INDArrayUnaryGradientKernel
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSliceGradient
type NDArrayStridedSliceGradient struct {
	NDArrayUnaryGradientKernel
}

// NDArrayStridedSliceGradientFrom constructs a [NDArrayStridedSliceGradient] from an unsafe.Pointer.
func NDArrayStridedSliceGradientFrom(ptr unsafe.Pointer) NDArrayStridedSliceGradient {
	return NDArrayStridedSliceGradient{
		NDArrayUnaryGradientKernel: NDArrayUnaryGradientKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayStridedSliceGradientClass) Alloc() NDArrayStridedSliceGradient {
	rv := objc.Send[NDArrayStridedSliceGradient](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayStridedSliceGradientClass) New() NDArrayStridedSliceGradient {
	rv := objc.Send[NDArrayStridedSliceGradient](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayStridedSliceGradient) Init() NDArrayStridedSliceGradient {
	rv := objc.Send[NDArrayStridedSliceGradient](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayStridedSliceGradient) Autorelease() NDArrayStridedSliceGradient {
	rv := objc.Send[NDArrayStridedSliceGradient](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayStridedSliceGradient creates a new NDArrayStridedSliceGradient instance.
func NewNDArrayStridedSliceGradient() NDArrayStridedSliceGradient {
	return getNDArrayStridedSliceGradientClass().New()
}




