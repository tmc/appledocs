// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSRNNRecurrentMatrixState */


/* debug [class_header]: Header for MPSRNNRecurrentMatrixState */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RNNRecurrentMatrixState */
// An interface definition for the [RNNRecurrentMatrixState] class.
type IRNNRecurrentMatrixState interface {
	IState
	
/* debug [class_interface_properties]: Properties for RNNRecurrentMatrixState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RNNRecurrentMatrixState */
	// methods:
	GetRecurrentOutputMatrix()
	GetRecurrentOutputMatrixForLayerIndex(layerIndex uint) IMatrix
	GetMemoryCellMatrix()
	GetMemoryCellMatrixForLayerIndex(layerIndex uint) IMatrix
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RNNRecurrentMatrixState */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RNNRecurrentMatrixState */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RNNRecurrentMatrixState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RNNRecurrentMatrixState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RNNRecurrentMatrixState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RNNRecurrentMatrixState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate/2873339-getrecurrentoutputmatrix
func (r_ RNNRecurrentMatrixState) GetRecurrentOutputMatrix() {
	objc.Send[objc.ID](r_.ID, objc.Sel("getRecurrentOutputMatrix"))
}/* debug [instance_methods/method]: GetRecurrentOutputMatrix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate/2873339-getrecurrentoutputmatrixforlayer
func (r_ RNNRecurrentMatrixState) GetRecurrentOutputMatrixForLayerIndex(layerIndex uint) IMatrix {
	rv := objc.Send[Matrix](r_.ID, objc.Sel("getRecurrentOutputMatrixForLayerIndex:"), layerIndex)
	return rv
}/* debug [instance_methods/method]: GetRecurrentOutputMatrixForLayerIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate/2873390-getmemorycellmatrix
func (r_ RNNRecurrentMatrixState) GetMemoryCellMatrix() {
	objc.Send[objc.ID](r_.ID, objc.Sel("getMemoryCellMatrix"))
}/* debug [instance_methods/method]: GetMemoryCellMatrix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate/2873390-getmemorycellmatrixforlayerindex
func (r_ RNNRecurrentMatrixState) GetMemoryCellMatrixForLayerIndex(layerIndex uint) IMatrix {
	rv := objc.Send[Matrix](r_.ID, objc.Sel("getMemoryCellMatrixForLayerIndex:"), layerIndex)
	return rv
}/* debug [instance_methods/method]: GetMemoryCellMatrixForLayerIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RNNRecurrentMatrixState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSRNNRecurrentMatrixState */



