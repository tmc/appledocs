// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayGradientState] class.
var (
	NDArrayGradientStateClass     _NDArrayGradientStateClass
	NDArrayGradientStateClassOnce sync.Once
)

func getNDArrayGradientStateClass() _NDArrayGradientStateClass {
	NDArrayGradientStateClassOnce.Do(func() {
		NDArrayGradientStateClass = _NDArrayGradientStateClass{objc.GetClass("MPSNDArrayGradientState")}
	})
	return NDArrayGradientStateClass
}

type _NDArrayGradientStateClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayGradientState] class.
type INDArrayGradientState interface {
	IState
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGradientState
type NDArrayGradientState struct {
	State
}

// NDArrayGradientStateFrom constructs a [NDArrayGradientState] from an unsafe.Pointer.
func NDArrayGradientStateFrom(ptr unsafe.Pointer) NDArrayGradientState {
	return NDArrayGradientState{
		State: StateFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayGradientStateClass) Alloc() NDArrayGradientState {
	rv := objc.Send[NDArrayGradientState](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayGradientStateClass) New() NDArrayGradientState {
	rv := objc.Send[NDArrayGradientState](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayGradientState) Init() NDArrayGradientState {
	rv := objc.Send[NDArrayGradientState](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayGradientState) Autorelease() NDArrayGradientState {
	rv := objc.Send[NDArrayGradientState](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayGradientState creates a new NDArrayGradientState instance.
func NewNDArrayGradientState() NDArrayGradientState {
	return getNDArrayGradientStateClass().New()
}
