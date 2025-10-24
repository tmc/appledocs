// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionRowMeanNode */


/* debug [class_header]: Header for MPSNNReductionRowMeanNode */
// The class instance for the [ReductionRowMeanNode] class.
var (
	ReductionRowMeanNodeClass     _ReductionRowMeanNodeClass
	ReductionRowMeanNodeClassOnce sync.Once
)

func getReductionRowMeanNodeClass() _ReductionRowMeanNodeClass {
	ReductionRowMeanNodeClassOnce.Do(func() {
		ReductionRowMeanNodeClass = _ReductionRowMeanNodeClass{objc.GetClass("MPSNNReductionRowMeanNode")}
	})
	return ReductionRowMeanNodeClass
}

type _ReductionRowMeanNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionRowMeanNode */
// An interface definition for the [ReductionRowMeanNode] class.
type IReductionRowMeanNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionRowMeanNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionRowMeanNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionRowMeanNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionRowMeanNodeClass) Alloc() ReductionRowMeanNode {
	rv := objc.Send[ReductionRowMeanNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionRowMeanNodeClass) New() ReductionRowMeanNode {
	rv := objc.Send[ReductionRowMeanNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionRowMeanNode) Init() ReductionRowMeanNode {
	rv := objc.Send[ReductionRowMeanNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionRowMeanNode) Autorelease() ReductionRowMeanNode {
	rv := objc.Send[ReductionRowMeanNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionRowMeanNode creates a new ReductionRowMeanNode instance.
func NewReductionRowMeanNode() ReductionRowMeanNode {
	return getReductionRowMeanNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionRowMeanNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMeanNode
type ReductionRowMeanNode struct {
	UnaryReductionNode
}

// ReductionRowMeanNodeFrom constructs a [ReductionRowMeanNode] from an unsafe.Pointer.
func ReductionRowMeanNodeFrom(ptr unsafe.Pointer) ReductionRowMeanNode {
	return ReductionRowMeanNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionRowMeanNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionRowMeanNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionRowMeanNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionRowMeanNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionRowMeanNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionRowMeanNode */



