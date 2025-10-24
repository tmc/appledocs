// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionSpatialMeanNode */


/* debug [class_header]: Header for MPSNNReductionSpatialMeanNode */
// The class instance for the [ReductionSpatialMeanNode] class.
var (
	ReductionSpatialMeanNodeClass     _ReductionSpatialMeanNodeClass
	ReductionSpatialMeanNodeClassOnce sync.Once
)

func getReductionSpatialMeanNodeClass() _ReductionSpatialMeanNodeClass {
	ReductionSpatialMeanNodeClassOnce.Do(func() {
		ReductionSpatialMeanNodeClass = _ReductionSpatialMeanNodeClass{objc.GetClass("MPSNNReductionSpatialMeanNode")}
	})
	return ReductionSpatialMeanNodeClass
}

type _ReductionSpatialMeanNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionSpatialMeanNode */
// An interface definition for the [ReductionSpatialMeanNode] class.
type IReductionSpatialMeanNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionSpatialMeanNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionSpatialMeanNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionSpatialMeanNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionSpatialMeanNodeClass) Alloc() ReductionSpatialMeanNode {
	rv := objc.Send[ReductionSpatialMeanNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionSpatialMeanNodeClass) New() ReductionSpatialMeanNode {
	rv := objc.Send[ReductionSpatialMeanNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionSpatialMeanNode) Init() ReductionSpatialMeanNode {
	rv := objc.Send[ReductionSpatialMeanNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionSpatialMeanNode) Autorelease() ReductionSpatialMeanNode {
	rv := objc.Send[ReductionSpatialMeanNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionSpatialMeanNode creates a new ReductionSpatialMeanNode instance.
func NewReductionSpatialMeanNode() ReductionSpatialMeanNode {
	return getReductionSpatialMeanNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionSpatialMeanNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionSpatialMeanNode
type ReductionSpatialMeanNode struct {
	UnaryReductionNode
}

// ReductionSpatialMeanNodeFrom constructs a [ReductionSpatialMeanNode] from an unsafe.Pointer.
func ReductionSpatialMeanNodeFrom(ptr unsafe.Pointer) ReductionSpatialMeanNode {
	return ReductionSpatialMeanNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionSpatialMeanNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionSpatialMeanNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionSpatialMeanNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionSpatialMeanNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionSpatialMeanNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionSpatialMeanNode */



