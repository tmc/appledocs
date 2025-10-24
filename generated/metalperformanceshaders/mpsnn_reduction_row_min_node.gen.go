// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionRowMinNode */


/* debug [class_header]: Header for MPSNNReductionRowMinNode */
// The class instance for the [ReductionRowMinNode] class.
var (
	ReductionRowMinNodeClass     _ReductionRowMinNodeClass
	ReductionRowMinNodeClassOnce sync.Once
)

func getReductionRowMinNodeClass() _ReductionRowMinNodeClass {
	ReductionRowMinNodeClassOnce.Do(func() {
		ReductionRowMinNodeClass = _ReductionRowMinNodeClass{objc.GetClass("MPSNNReductionRowMinNode")}
	})
	return ReductionRowMinNodeClass
}

type _ReductionRowMinNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionRowMinNode */
// An interface definition for the [ReductionRowMinNode] class.
type IReductionRowMinNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionRowMinNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionRowMinNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionRowMinNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionRowMinNodeClass) Alloc() ReductionRowMinNode {
	rv := objc.Send[ReductionRowMinNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionRowMinNodeClass) New() ReductionRowMinNode {
	rv := objc.Send[ReductionRowMinNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionRowMinNode) Init() ReductionRowMinNode {
	rv := objc.Send[ReductionRowMinNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionRowMinNode) Autorelease() ReductionRowMinNode {
	rv := objc.Send[ReductionRowMinNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionRowMinNode creates a new ReductionRowMinNode instance.
func NewReductionRowMinNode() ReductionRowMinNode {
	return getReductionRowMinNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionRowMinNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMinNode
type ReductionRowMinNode struct {
	UnaryReductionNode
}

// ReductionRowMinNodeFrom constructs a [ReductionRowMinNode] from an unsafe.Pointer.
func ReductionRowMinNodeFrom(ptr unsafe.Pointer) ReductionRowMinNode {
	return ReductionRowMinNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionRowMinNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionRowMinNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionRowMinNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionRowMinNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionRowMinNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionRowMinNode */



