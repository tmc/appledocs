// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayGather] class.
var (
	NDArrayGatherClass     _NDArrayGatherClass
	NDArrayGatherClassOnce sync.Once
)

func getNDArrayGatherClass() _NDArrayGatherClass {
	NDArrayGatherClassOnce.Do(func() {
		NDArrayGatherClass = _NDArrayGatherClass{objc.GetClass("MPSNDArrayGather")}
	})
	return NDArrayGatherClass
}

type _NDArrayGatherClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayGather] class.
type INDArrayGather interface {
	INDArrayBinaryKernel
	Axis() int
	SetAxis(value int)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGather
type NDArrayGather struct {
	NDArrayBinaryKernel
}

// NDArrayGatherFrom constructs a [NDArrayGather] from an unsafe.Pointer.
func NDArrayGatherFrom(ptr unsafe.Pointer) NDArrayGather {
	return NDArrayGather{
		NDArrayBinaryKernel: NDArrayBinaryKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayGatherClass) Alloc() NDArrayGather {
	rv := objc.Send[NDArrayGather](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayGatherClass) New() NDArrayGather {
	rv := objc.Send[NDArrayGather](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayGather) Init() NDArrayGather {
	rv := objc.Send[NDArrayGather](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayGather) Autorelease() NDArrayGather {
	rv := objc.Send[NDArrayGather](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayGather creates a new NDArrayGather instance.
func NewNDArrayGather() NDArrayGather {
	return getNDArrayGatherClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygather/axis
func (n_ NDArrayGather) Axis() int {
	rv := objc.Send[int](n_.ID, objc.Sel("axis"))
	return rv
}


// SetAxis sets the value of the axis property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygather/axis
func (n_ NDArrayGather) SetAxis(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAxis:"), value)
}



