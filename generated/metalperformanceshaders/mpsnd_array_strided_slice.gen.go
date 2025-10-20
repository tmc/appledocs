// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayStridedSlice] class.
var (
	NDArrayStridedSliceClass     _NDArrayStridedSliceClass
	NDArrayStridedSliceClassOnce sync.Once
)

func getNDArrayStridedSliceClass() _NDArrayStridedSliceClass {
	NDArrayStridedSliceClassOnce.Do(func() {
		NDArrayStridedSliceClass = _NDArrayStridedSliceClass{objc.GetClass("MPSNDArrayStridedSlice")}
	})
	return NDArrayStridedSliceClass
}

type _NDArrayStridedSliceClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayStridedSlice] class.
type INDArrayStridedSlice interface {
	INDArrayUnaryKernel
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSlice
type NDArrayStridedSlice struct {
	NDArrayUnaryKernel
}

// NDArrayStridedSliceFrom constructs a [NDArrayStridedSlice] from an unsafe.Pointer.
func NDArrayStridedSliceFrom(ptr unsafe.Pointer) NDArrayStridedSlice {
	return NDArrayStridedSlice{
		NDArrayUnaryKernel: NDArrayUnaryKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayStridedSliceClass) Alloc() NDArrayStridedSlice {
	rv := objc.Send[NDArrayStridedSlice](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayStridedSliceClass) New() NDArrayStridedSlice {
	rv := objc.Send[NDArrayStridedSlice](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayStridedSlice) Init() NDArrayStridedSlice {
	rv := objc.Send[NDArrayStridedSlice](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayStridedSlice) Autorelease() NDArrayStridedSlice {
	rv := objc.Send[NDArrayStridedSlice](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayStridedSlice creates a new NDArrayStridedSlice instance.
func NewNDArrayStridedSlice() NDArrayStridedSlice {
	return getNDArrayStridedSliceClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSlice/strides
func (n_ NDArrayStridedSlice) Strides() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("strides"))
	return rv
}

// SetStrides sets the value of the strides property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSlice/strides
func (n_ NDArrayStridedSlice) SetStrides(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setStrides:"), value)
}
