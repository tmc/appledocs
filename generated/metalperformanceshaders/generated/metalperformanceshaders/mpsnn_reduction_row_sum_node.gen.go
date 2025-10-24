// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionRowSumNode */


/* debug [class_header]: Header for MPSNNReductionRowSumNode */
// The class instance for the [ReductionRowSumNode] class.
var (
	ReductionRowSumNodeClass     _ReductionRowSumNodeClass
	ReductionRowSumNodeClassOnce sync.Once
)

func getReductionRowSumNodeClass() _ReductionRowSumNodeClass {
	ReductionRowSumNodeClassOnce.Do(func() {
		ReductionRowSumNodeClass = _ReductionRowSumNodeClass{objc.GetClass("MPSNNReductionRowSumNode")}
	})
	return ReductionRowSumNodeClass
}

type _ReductionRowSumNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionRowSumNode */
// An interface definition for the [ReductionRowSumNode] class.
type IReductionRowSumNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionRowSumNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionRowSumNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionRowSumNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionRowSumNodeClass) Alloc() ReductionRowSumNode {
	rv := objc.Send[ReductionRowSumNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionRowSumNodeClass) New() ReductionRowSumNode {
	rv := objc.Send[ReductionRowSumNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionRowSumNode) Init() ReductionRowSumNode {
	rv := objc.Send[ReductionRowSumNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionRowSumNode) Autorelease() ReductionRowSumNode {
	rv := objc.Send[ReductionRowSumNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionRowSumNode creates a new ReductionRowSumNode instance.
func NewReductionRowSumNode() ReductionRowSumNode {
	return getReductionRowSumNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionRowSumNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowSumNode
type ReductionRowSumNode struct {
	UnaryReductionNode
}

// ReductionRowSumNodeFrom constructs a [ReductionRowSumNode] from an unsafe.Pointer.
func ReductionRowSumNodeFrom(ptr unsafe.Pointer) ReductionRowSumNode {
	return ReductionRowSumNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionRowSumNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionRowSumNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionRowSumNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionRowSumNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionRowSumNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionRowSumNode */



