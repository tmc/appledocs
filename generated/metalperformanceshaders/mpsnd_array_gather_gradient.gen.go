// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayGatherGradient] class.
var (
	NDArrayGatherGradientClass     _NDArrayGatherGradientClass
	NDArrayGatherGradientClassOnce sync.Once
)

func getNDArrayGatherGradientClass() _NDArrayGatherGradientClass {
	NDArrayGatherGradientClassOnce.Do(func() {
		NDArrayGatherGradientClass = _NDArrayGatherGradientClass{objc.GetClass("MPSNDArrayGatherGradient")}
	})
	return NDArrayGatherGradientClass
}

type _NDArrayGatherGradientClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayGatherGradient] class.
type INDArrayGatherGradient interface {
	INDArrayBinaryPrimaryGradientKernel
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGatherGradient
type NDArrayGatherGradient struct {
	NDArrayBinaryPrimaryGradientKernel
}

// NDArrayGatherGradientFrom constructs a [NDArrayGatherGradient] from an unsafe.Pointer.
func NDArrayGatherGradientFrom(ptr unsafe.Pointer) NDArrayGatherGradient {
	return NDArrayGatherGradient{
		NDArrayBinaryPrimaryGradientKernel: NDArrayBinaryPrimaryGradientKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayGatherGradientClass) Alloc() NDArrayGatherGradient {
	rv := objc.Send[NDArrayGatherGradient](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayGatherGradientClass) New() NDArrayGatherGradient {
	rv := objc.Send[NDArrayGatherGradient](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayGatherGradient) Init() NDArrayGatherGradient {
	rv := objc.Send[NDArrayGatherGradient](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayGatherGradient) Autorelease() NDArrayGatherGradient {
	rv := objc.Send[NDArrayGatherGradient](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayGatherGradient creates a new NDArrayGatherGradient instance.
func NewNDArrayGatherGradient() NDArrayGatherGradient {
	return getNDArrayGatherGradientClass().New()
}




