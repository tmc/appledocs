// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [RNNMatrixTrainingState] class.
var (
	RNNMatrixTrainingStateClass     _RNNMatrixTrainingStateClass
	RNNMatrixTrainingStateClassOnce sync.Once
)

func getRNNMatrixTrainingStateClass() _RNNMatrixTrainingStateClass {
	RNNMatrixTrainingStateClassOnce.Do(func() {
		RNNMatrixTrainingStateClass = _RNNMatrixTrainingStateClass{objc.GetClass("MPSRNNMatrixTrainingState")}
	})
	return RNNMatrixTrainingStateClass
}

type _RNNMatrixTrainingStateClass struct {
	class objc.Class
}





// An interface definition for the [RNNMatrixTrainingState] class.
type IRNNMatrixTrainingState interface {
	IState
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _RNNMatrixTrainingStateClass) Alloc() RNNMatrixTrainingState {
	rv := objc.Send[RNNMatrixTrainingState](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RNNMatrixTrainingStateClass) New() RNNMatrixTrainingState {
	rv := objc.Send[RNNMatrixTrainingState](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNMatrixTrainingState) Init() RNNMatrixTrainingState {
	rv := objc.Send[RNNMatrixTrainingState](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNMatrixTrainingState) Autorelease() RNNMatrixTrainingState {
	rv := objc.Send[RNNMatrixTrainingState](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNMatrixTrainingState creates a new RNNMatrixTrainingState instance.
func NewRNNMatrixTrainingState() RNNMatrixTrainingState {
	return getRNNMatrixTrainingStateClass().New()
}





// A class that holds data from a forward pass to be used in a backward pass.


// A class that holds data from a forward pass to be used in a backward pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingState
type RNNMatrixTrainingState struct {
	State
}

// RNNMatrixTrainingStateFrom constructs a [RNNMatrixTrainingState] from an unsafe.Pointer.
//
// A class that holds data from a forward pass to be used in a backward pass.
func RNNMatrixTrainingStateFrom(ptr unsafe.Pointer) RNNMatrixTrainingState {
	return RNNMatrixTrainingState{
		State: StateFrom(ptr),
	}
}































