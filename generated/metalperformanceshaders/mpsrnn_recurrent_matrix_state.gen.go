// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [RNNRecurrentMatrixState] class.
var (
	RNNRecurrentMatrixStateClass     _RNNRecurrentMatrixStateClass
	RNNRecurrentMatrixStateClassOnce sync.Once
)

func getRNNRecurrentMatrixStateClass() _RNNRecurrentMatrixStateClass {
	RNNRecurrentMatrixStateClassOnce.Do(func() {
		RNNRecurrentMatrixStateClass = _RNNRecurrentMatrixStateClass{objc.GetClass("MPSRNNRecurrentMatrixState")}
	})
	return RNNRecurrentMatrixStateClass
}

type _RNNRecurrentMatrixStateClass struct {
	class objc.Class
}





// An interface definition for the [RNNRecurrentMatrixState] class.
type IRNNRecurrentMatrixState interface {
	IState
	

	// properties:


	

	// methods:
	GetRecurrentOutputMatrix()
	GetRecurrentOutputMatrixForLayerIndex(layerIndex uint) IMatrix
	GetMemoryCellMatrix()
	GetMemoryCellMatrixForLayerIndex(layerIndex uint) IMatrix


}





// Alloc allocates a new instance without initialization.
func (rc _RNNRecurrentMatrixStateClass) Alloc() RNNRecurrentMatrixState {
	rv := objc.Send[RNNRecurrentMatrixState](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RNNRecurrentMatrixStateClass) New() RNNRecurrentMatrixState {
	rv := objc.Send[RNNRecurrentMatrixState](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNRecurrentMatrixState) Init() RNNRecurrentMatrixState {
	rv := objc.Send[RNNRecurrentMatrixState](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNRecurrentMatrixState) Autorelease() RNNRecurrentMatrixState {
	rv := objc.Send[RNNRecurrentMatrixState](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNRecurrentMatrixState creates a new RNNRecurrentMatrixState instance.
func NewRNNRecurrentMatrixState() RNNRecurrentMatrixState {
	return getRNNRecurrentMatrixStateClass().New()
}





// A class holds all the data that’s passed from one sequence iteration of the matrix-based recurrent neural network layer to the next.


// A class holds all the data that’s passed from one sequence iteration of the matrix-based recurrent neural network layer to the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNRecurrentMatrixState
type RNNRecurrentMatrixState struct {
	State
}

// RNNRecurrentMatrixStateFrom constructs a [RNNRecurrentMatrixState] from an unsafe.Pointer.
//
// A class holds all the data that’s passed from one sequence iteration of the matrix-based recurrent neural network layer to the next.
func RNNRecurrentMatrixStateFrom(ptr unsafe.Pointer) RNNRecurrentMatrixState {
	return RNNRecurrentMatrixState{
		State: StateFrom(ptr),
	}
}




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate/2873339-getrecurrentoutputmatrix
func (r_ RNNRecurrentMatrixState) GetRecurrentOutputMatrix() {
	objc.Send[objc.ID](r_.ID, objc.Sel("getRecurrentOutputMatrix"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate/2873339-getrecurrentoutputmatrixforlayer
func (r_ RNNRecurrentMatrixState) GetRecurrentOutputMatrixForLayerIndex(layerIndex uint) IMatrix {
	rv := objc.Send[Matrix](r_.ID, objc.Sel("getRecurrentOutputMatrixForLayerIndex:"), layerIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate/2873390-getmemorycellmatrix
func (r_ RNNRecurrentMatrixState) GetMemoryCellMatrix() {
	objc.Send[objc.ID](r_.ID, objc.Sel("getMemoryCellMatrix"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate/2873390-getmemorycellmatrixforlayerindex
func (r_ RNNRecurrentMatrixState) GetMemoryCellMatrixForLayerIndex(layerIndex uint) IMatrix {
	rv := objc.Send[Matrix](r_.ID, objc.Sel("getMemoryCellMatrixForLayerIndex:"), layerIndex)
	return rv
}













