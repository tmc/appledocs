// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayGatherGradientState] class.
var (
	NDArrayGatherGradientStateClass     _NDArrayGatherGradientStateClass
	NDArrayGatherGradientStateClassOnce sync.Once
)

func getNDArrayGatherGradientStateClass() _NDArrayGatherGradientStateClass {
	NDArrayGatherGradientStateClassOnce.Do(func() {
		NDArrayGatherGradientStateClass = _NDArrayGatherGradientStateClass{objc.GetClass("MPSNDArrayGatherGradientState")}
	})
	return NDArrayGatherGradientStateClass
}

type _NDArrayGatherGradientStateClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayGatherGradientState] class.
type INDArrayGatherGradientState interface {
	INDArrayGradientState
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGatherGradientState
type NDArrayGatherGradientState struct {
	NDArrayGradientState
}

// NDArrayGatherGradientStateFrom constructs a [NDArrayGatherGradientState] from an unsafe.Pointer.
func NDArrayGatherGradientStateFrom(ptr unsafe.Pointer) NDArrayGatherGradientState {
	return NDArrayGatherGradientState{
		NDArrayGradientState: NDArrayGradientStateFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayGatherGradientStateClass) Alloc() NDArrayGatherGradientState {
	rv := objc.Send[NDArrayGatherGradientState](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayGatherGradientStateClass) New() NDArrayGatherGradientState {
	rv := objc.Send[NDArrayGatherGradientState](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayGatherGradientState) Init() NDArrayGatherGradientState {
	rv := objc.Send[NDArrayGatherGradientState](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayGatherGradientState) Autorelease() NDArrayGatherGradientState {
	rv := objc.Send[NDArrayGatherGradientState](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayGatherGradientState creates a new NDArrayGatherGradientState instance.
func NewNDArrayGatherGradientState() NDArrayGatherGradientState {
	return getNDArrayGatherGradientStateClass().New()
}




