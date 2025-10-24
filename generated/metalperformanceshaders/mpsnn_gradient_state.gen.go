// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [GradientState] class.
var (
	GradientStateClass     _GradientStateClass
	GradientStateClassOnce sync.Once
)

func getGradientStateClass() _GradientStateClass {
	GradientStateClassOnce.Do(func() {
		GradientStateClass = _GradientStateClass{objc.GetClass("MPSNNGradientState")}
	})
	return GradientStateClass
}

type _GradientStateClass struct {
	class objc.Class
}





// An interface definition for the [GradientState] class.
type IGradientState interface {
	IState
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GradientStateClass) Alloc() GradientState {
	rv := objc.Send[GradientState](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GradientStateClass) New() GradientState {
	rv := objc.Send[GradientState](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GradientState) Init() GradientState {
	rv := objc.Send[GradientState](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GradientState) Autorelease() GradientState {
	rv := objc.Send[GradientState](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGradientState creates a new GradientState instance.
func NewGradientState() GradientState {
	return getGradientStateClass().New()
}





// A class representing the state of a gradient kernel when it was encoded.


// A class representing the state of a gradient kernel when it was encoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientState
type GradientState struct {
	State
}

// GradientStateFrom constructs a [GradientState] from an unsafe.Pointer.
//
// A class representing the state of a gradient kernel when it was encoded.
func GradientStateFrom(ptr unsafe.Pointer) GradientState {
	return GradientState{
		State: StateFrom(ptr),
	}
}































