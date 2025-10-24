// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionColumnSumNode */


/* debug [class_header]: Header for MPSNNReductionColumnSumNode */
// The class instance for the [ReductionColumnSumNode] class.
var (
	ReductionColumnSumNodeClass     _ReductionColumnSumNodeClass
	ReductionColumnSumNodeClassOnce sync.Once
)

func getReductionColumnSumNodeClass() _ReductionColumnSumNodeClass {
	ReductionColumnSumNodeClassOnce.Do(func() {
		ReductionColumnSumNodeClass = _ReductionColumnSumNodeClass{objc.GetClass("MPSNNReductionColumnSumNode")}
	})
	return ReductionColumnSumNodeClass
}

type _ReductionColumnSumNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionColumnSumNode */
// An interface definition for the [ReductionColumnSumNode] class.
type IReductionColumnSumNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionColumnSumNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionColumnSumNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionColumnSumNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionColumnSumNodeClass) Alloc() ReductionColumnSumNode {
	rv := objc.Send[ReductionColumnSumNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionColumnSumNodeClass) New() ReductionColumnSumNode {
	rv := objc.Send[ReductionColumnSumNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionColumnSumNode) Init() ReductionColumnSumNode {
	rv := objc.Send[ReductionColumnSumNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionColumnSumNode) Autorelease() ReductionColumnSumNode {
	rv := objc.Send[ReductionColumnSumNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionColumnSumNode creates a new ReductionColumnSumNode instance.
func NewReductionColumnSumNode() ReductionColumnSumNode {
	return getReductionColumnSumNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionColumnSumNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnSumNode
type ReductionColumnSumNode struct {
	UnaryReductionNode
}

// ReductionColumnSumNodeFrom constructs a [ReductionColumnSumNode] from an unsafe.Pointer.
func ReductionColumnSumNodeFrom(ptr unsafe.Pointer) ReductionColumnSumNode {
	return ReductionColumnSumNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionColumnSumNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionColumnSumNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionColumnSumNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionColumnSumNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionColumnSumNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionColumnSumNode */



