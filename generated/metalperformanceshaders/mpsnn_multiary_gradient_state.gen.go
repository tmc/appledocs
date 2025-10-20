// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MultiaryGradientState] class.
var (
	MultiaryGradientStateClass     _MultiaryGradientStateClass
	MultiaryGradientStateClassOnce sync.Once
)

func getMultiaryGradientStateClass() _MultiaryGradientStateClass {
	MultiaryGradientStateClassOnce.Do(func() {
		MultiaryGradientStateClass = _MultiaryGradientStateClass{objc.GetClass("MPSNNMultiaryGradientState")}
	})
	return MultiaryGradientStateClass
}

type _MultiaryGradientStateClass struct {
	class objc.Class
}

// An interface definition for the [MultiaryGradientState] class.
type IMultiaryGradientState interface {
	IState
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiaryGradientState
type MultiaryGradientState struct {
	State
}

// MultiaryGradientStateFrom constructs a [MultiaryGradientState] from an unsafe.Pointer.
func MultiaryGradientStateFrom(ptr unsafe.Pointer) MultiaryGradientState {
	return MultiaryGradientState{
		State: StateFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MultiaryGradientStateClass) Alloc() MultiaryGradientState {
	rv := objc.Send[MultiaryGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MultiaryGradientStateClass) New() MultiaryGradientState {
	rv := objc.Send[MultiaryGradientState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiaryGradientState) Init() MultiaryGradientState {
	rv := objc.Send[MultiaryGradientState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiaryGradientState) Autorelease() MultiaryGradientState {
	rv := objc.Send[MultiaryGradientState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiaryGradientState creates a new MultiaryGradientState instance.
func NewMultiaryGradientState() MultiaryGradientState {
	return getMultiaryGradientStateClass().New()
}
