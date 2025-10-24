// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CTrainingGraph] class.
var (
	CTrainingGraphClass     _CTrainingGraphClass
	CTrainingGraphClassOnce sync.Once
)

func getCTrainingGraphClass() _CTrainingGraphClass {
	CTrainingGraphClassOnce.Do(func() {
		CTrainingGraphClass = _CTrainingGraphClass{objc.GetClass("MLCTrainingGraph")}
	})
	return CTrainingGraphClass
}

type _CTrainingGraphClass struct {
	class objc.Class
}

// An interface definition for the [CTrainingGraph] class.
type ICTrainingGraph interface {
	ICGraph
	// properties:
	DeviceMemorySize() int
	SetDeviceMemorySize(value int)
	Optimizer() COptimizer /* not a class type */
	SetOptimizer(value COptimizer /* not a class type */)
	// methods:
}

// A training graph that you create from one or more graph objects plus additional layers you add directly to the training graph.
//
// The framework provides a family of graph-execution methods to execute a full training iteration, and methods to execute the forward pass, the gradient pass, and optimizer update, individually. Use one of the   methods to execute a full training iteration to accelerate an ML model represented as a single training graph. Use one of the  ,  , or  to accelerate an ML library that separates the forward pass, gradient pass, and optimizer update as separate phases.


// A training graph that you create from one or more graph objects plus additional layers you add directly to the training graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph
type CTrainingGraph struct {
	CGraph
}

// CTrainingGraphFrom constructs a [CTrainingGraph] from an unsafe.Pointer.
//
// A training graph that you create from one or more graph objects plus additional layers you add directly to the training graph.
func CTrainingGraphFrom(ptr unsafe.Pointer) CTrainingGraph {
	return CTrainingGraph{
		CGraph: CGraphFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CTrainingGraphClass) Alloc() CTrainingGraph {
	rv := objc.Send[CTrainingGraph](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CTrainingGraphClass) New() CTrainingGraph {
	rv := objc.Send[CTrainingGraph](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTrainingGraph) Init() CTrainingGraph {
	rv := objc.Send[CTrainingGraph](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTrainingGraph) Autorelease() CTrainingGraph {
	rv := objc.Send[CTrainingGraph](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTrainingGraph creates a new CTrainingGraph instance.
func NewCTrainingGraph() CTrainingGraph {
	return getCTrainingGraphClass().New()
}



// The device memory size in bytes for all intermediate tensors for forward, gradient passes, and optimizer updates for all layers in the training graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctraininggraph/devicememorysize
func (c_ CTrainingGraph) DeviceMemorySize() int {
	rv := objc.Send[int](c_.ID, objc.Sel("deviceMemorySize"))
	return rv
}


// The device memory size in bytes for all intermediate tensors for forward, gradient passes, and optimizer updates for all layers in the training graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctraininggraph/devicememorysize
func (c_ CTrainingGraph) SetDeviceMemorySize(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeviceMemorySize:"), value)
}


// The optimizer to use with the training graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctraininggraph/optimizer
func (c_ CTrainingGraph) Optimizer() COptimizer /* not a class type */ {
	rv := objc.Send[COptimizer](c_.ID, objc.Sel("optimizer"))
	return rv
}


// The optimizer to use with the training graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctraininggraph/optimizer
func (c_ CTrainingGraph) SetOptimizer(value COptimizer /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOptimizer:"), value)
}



