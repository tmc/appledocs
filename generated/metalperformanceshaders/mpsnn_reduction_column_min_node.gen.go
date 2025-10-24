// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionColumnMinNode */


/* debug [class_header]: Header for MPSNNReductionColumnMinNode */
// The class instance for the [ReductionColumnMinNode] class.
var (
	ReductionColumnMinNodeClass     _ReductionColumnMinNodeClass
	ReductionColumnMinNodeClassOnce sync.Once
)

func getReductionColumnMinNodeClass() _ReductionColumnMinNodeClass {
	ReductionColumnMinNodeClassOnce.Do(func() {
		ReductionColumnMinNodeClass = _ReductionColumnMinNodeClass{objc.GetClass("MPSNNReductionColumnMinNode")}
	})
	return ReductionColumnMinNodeClass
}

type _ReductionColumnMinNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionColumnMinNode */
// An interface definition for the [ReductionColumnMinNode] class.
type IReductionColumnMinNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionColumnMinNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionColumnMinNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionColumnMinNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionColumnMinNodeClass) Alloc() ReductionColumnMinNode {
	rv := objc.Send[ReductionColumnMinNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionColumnMinNodeClass) New() ReductionColumnMinNode {
	rv := objc.Send[ReductionColumnMinNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionColumnMinNode) Init() ReductionColumnMinNode {
	rv := objc.Send[ReductionColumnMinNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionColumnMinNode) Autorelease() ReductionColumnMinNode {
	rv := objc.Send[ReductionColumnMinNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionColumnMinNode creates a new ReductionColumnMinNode instance.
func NewReductionColumnMinNode() ReductionColumnMinNode {
	return getReductionColumnMinNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionColumnMinNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMinNode
type ReductionColumnMinNode struct {
	UnaryReductionNode
}

// ReductionColumnMinNodeFrom constructs a [ReductionColumnMinNode] from an unsafe.Pointer.
func ReductionColumnMinNodeFrom(ptr unsafe.Pointer) ReductionColumnMinNode {
	return ReductionColumnMinNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionColumnMinNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionColumnMinNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionColumnMinNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionColumnMinNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionColumnMinNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionColumnMinNode */



