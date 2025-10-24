// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionColumnMeanNode */


/* debug [class_header]: Header for MPSNNReductionColumnMeanNode */
// The class instance for the [ReductionColumnMeanNode] class.
var (
	ReductionColumnMeanNodeClass     _ReductionColumnMeanNodeClass
	ReductionColumnMeanNodeClassOnce sync.Once
)

func getReductionColumnMeanNodeClass() _ReductionColumnMeanNodeClass {
	ReductionColumnMeanNodeClassOnce.Do(func() {
		ReductionColumnMeanNodeClass = _ReductionColumnMeanNodeClass{objc.GetClass("MPSNNReductionColumnMeanNode")}
	})
	return ReductionColumnMeanNodeClass
}

type _ReductionColumnMeanNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionColumnMeanNode */
// An interface definition for the [ReductionColumnMeanNode] class.
type IReductionColumnMeanNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionColumnMeanNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionColumnMeanNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionColumnMeanNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionColumnMeanNodeClass) Alloc() ReductionColumnMeanNode {
	rv := objc.Send[ReductionColumnMeanNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionColumnMeanNodeClass) New() ReductionColumnMeanNode {
	rv := objc.Send[ReductionColumnMeanNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionColumnMeanNode) Init() ReductionColumnMeanNode {
	rv := objc.Send[ReductionColumnMeanNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionColumnMeanNode) Autorelease() ReductionColumnMeanNode {
	rv := objc.Send[ReductionColumnMeanNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionColumnMeanNode creates a new ReductionColumnMeanNode instance.
func NewReductionColumnMeanNode() ReductionColumnMeanNode {
	return getReductionColumnMeanNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionColumnMeanNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMeanNode
type ReductionColumnMeanNode struct {
	UnaryReductionNode
}

// ReductionColumnMeanNodeFrom constructs a [ReductionColumnMeanNode] from an unsafe.Pointer.
func ReductionColumnMeanNodeFrom(ptr unsafe.Pointer) ReductionColumnMeanNode {
	return ReductionColumnMeanNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionColumnMeanNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionColumnMeanNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionColumnMeanNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionColumnMeanNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionColumnMeanNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionColumnMeanNode */



