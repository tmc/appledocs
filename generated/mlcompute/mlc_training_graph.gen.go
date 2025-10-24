// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLCTrainingGraph */


/* debug [class_header]: Header for MLCTrainingGraph */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CTrainingGraph */
// An interface definition for the [CTrainingGraph] class.
type ICTrainingGraph interface {
	ICGraph
	
/* debug [class_interface_properties]: Properties for CTrainingGraph */
	// properties:
	DeviceMemorySize() uint
	Optimizer() IMLCOptimizer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CTrainingGraph */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CTrainingGraph */
// Alloc allocates a new instance without initialization.
func (cc _CTrainingGraphClass) Alloc() CTrainingGraph {
	rv := objc.Send[CTrainingGraph](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CTrainingGraph */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CTrainingGraph */

// Creates a training graph with the layers from the graph objects, loss layer, and optimizer you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/init(graphObjects:lossLayer:optimizer:)
func NewCTrainingGraphWithGraphObjectsLossLayerOptimizer(graphObjects []CGraph, lossLayer IMLCLayer, optimizer IMLCOptimizer) CTrainingGraph {
	rv := objc.Send[CTrainingGraph](objc.ID(getCTrainingGraphClass().class), objc.Sel("graphWithGraphObjects:lossLayer:optimizer:"), graphObjects, lossLayer, optimizer)
	return rv
}/* debug [class_init_methods/constructor]: NewCTrainingGraphWithGraphObjectsLossLayerOptimizer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CTrainingGraph */

// Creates a training graph with the layers from the graph objects, loss layer, and optimizer you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/init(graphObjects:lossLayer:optimizer:)
func (cc _CTrainingGraphClass) GraphWithGraphObjectsLossLayerOptimizer(graphObjects []CGraph, lossLayer IMLCLayer, optimizer IMLCOptimizer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("graphWithGraphObjects:lossLayer:optimizer:"), graphObjects, lossLayer, optimizer)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphWithGraphObjectsLossLayerOptimizer) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CTrainingGraph */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CTrainingGraph */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CTrainingGraph */

// The device memory size in bytes for all intermediate tensors for forward, gradient passes, and optimizer updates for all layers in the training graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/deviceMemorySize
func (c_ CTrainingGraph) DeviceMemorySize() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("deviceMemorySize"))
	return rv
}/* debug [instance_properties/getter]: deviceMemorySize */


// The optimizer to use with the training graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTrainingGraph/optimizer
func (c_ CTrainingGraph) Optimizer() IMLCOptimizer {
	rv := objc.Send[COptimizer](c_.ID, objc.Sel("optimizer"))
	return rv
}/* debug [instance_properties/getter]: optimizer */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCTrainingGraph */


