// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionColumnMaxNode */


/* debug [class_header]: Header for MPSNNReductionColumnMaxNode */
// The class instance for the [ReductionColumnMaxNode] class.
var (
	ReductionColumnMaxNodeClass     _ReductionColumnMaxNodeClass
	ReductionColumnMaxNodeClassOnce sync.Once
)

func getReductionColumnMaxNodeClass() _ReductionColumnMaxNodeClass {
	ReductionColumnMaxNodeClassOnce.Do(func() {
		ReductionColumnMaxNodeClass = _ReductionColumnMaxNodeClass{objc.GetClass("MPSNNReductionColumnMaxNode")}
	})
	return ReductionColumnMaxNodeClass
}

type _ReductionColumnMaxNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionColumnMaxNode */
// An interface definition for the [ReductionColumnMaxNode] class.
type IReductionColumnMaxNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionColumnMaxNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionColumnMaxNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionColumnMaxNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionColumnMaxNodeClass) Alloc() ReductionColumnMaxNode {
	rv := objc.Send[ReductionColumnMaxNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionColumnMaxNodeClass) New() ReductionColumnMaxNode {
	rv := objc.Send[ReductionColumnMaxNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionColumnMaxNode) Init() ReductionColumnMaxNode {
	rv := objc.Send[ReductionColumnMaxNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionColumnMaxNode) Autorelease() ReductionColumnMaxNode {
	rv := objc.Send[ReductionColumnMaxNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionColumnMaxNode creates a new ReductionColumnMaxNode instance.
func NewReductionColumnMaxNode() ReductionColumnMaxNode {
	return getReductionColumnMaxNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionColumnMaxNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMaxNode
type ReductionColumnMaxNode struct {
	UnaryReductionNode
}

// ReductionColumnMaxNodeFrom constructs a [ReductionColumnMaxNode] from an unsafe.Pointer.
func ReductionColumnMaxNodeFrom(ptr unsafe.Pointer) ReductionColumnMaxNode {
	return ReductionColumnMaxNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionColumnMaxNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionColumnMaxNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionColumnMaxNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionColumnMaxNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionColumnMaxNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionColumnMaxNode */



