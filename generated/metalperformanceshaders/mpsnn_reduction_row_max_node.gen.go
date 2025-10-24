// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionRowMaxNode */


/* debug [class_header]: Header for MPSNNReductionRowMaxNode */
// The class instance for the [ReductionRowMaxNode] class.
var (
	ReductionRowMaxNodeClass     _ReductionRowMaxNodeClass
	ReductionRowMaxNodeClassOnce sync.Once
)

func getReductionRowMaxNodeClass() _ReductionRowMaxNodeClass {
	ReductionRowMaxNodeClassOnce.Do(func() {
		ReductionRowMaxNodeClass = _ReductionRowMaxNodeClass{objc.GetClass("MPSNNReductionRowMaxNode")}
	})
	return ReductionRowMaxNodeClass
}

type _ReductionRowMaxNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionRowMaxNode */
// An interface definition for the [ReductionRowMaxNode] class.
type IReductionRowMaxNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionRowMaxNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionRowMaxNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionRowMaxNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionRowMaxNodeClass) Alloc() ReductionRowMaxNode {
	rv := objc.Send[ReductionRowMaxNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionRowMaxNodeClass) New() ReductionRowMaxNode {
	rv := objc.Send[ReductionRowMaxNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionRowMaxNode) Init() ReductionRowMaxNode {
	rv := objc.Send[ReductionRowMaxNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionRowMaxNode) Autorelease() ReductionRowMaxNode {
	rv := objc.Send[ReductionRowMaxNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionRowMaxNode creates a new ReductionRowMaxNode instance.
func NewReductionRowMaxNode() ReductionRowMaxNode {
	return getReductionRowMaxNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionRowMaxNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMaxNode
type ReductionRowMaxNode struct {
	UnaryReductionNode
}

// ReductionRowMaxNodeFrom constructs a [ReductionRowMaxNode] from an unsafe.Pointer.
func ReductionRowMaxNodeFrom(ptr unsafe.Pointer) ReductionRowMaxNode {
	return ReductionRowMaxNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionRowMaxNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionRowMaxNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionRowMaxNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionRowMaxNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionRowMaxNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionRowMaxNode */



