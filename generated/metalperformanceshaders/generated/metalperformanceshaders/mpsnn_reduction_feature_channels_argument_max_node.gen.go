// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionFeatureChannelsArgumentMaxNode */


/* debug [class_header]: Header for MPSNNReductionFeatureChannelsArgumentMaxNode */
// The class instance for the [ReductionFeatureChannelsArgumentMaxNode] class.
var (
	ReductionFeatureChannelsArgumentMaxNodeClass     _ReductionFeatureChannelsArgumentMaxNodeClass
	ReductionFeatureChannelsArgumentMaxNodeClassOnce sync.Once
)

func getReductionFeatureChannelsArgumentMaxNodeClass() _ReductionFeatureChannelsArgumentMaxNodeClass {
	ReductionFeatureChannelsArgumentMaxNodeClassOnce.Do(func() {
		ReductionFeatureChannelsArgumentMaxNodeClass = _ReductionFeatureChannelsArgumentMaxNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsArgumentMaxNode")}
	})
	return ReductionFeatureChannelsArgumentMaxNodeClass
}

type _ReductionFeatureChannelsArgumentMaxNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionFeatureChannelsArgumentMaxNode */
// An interface definition for the [ReductionFeatureChannelsArgumentMaxNode] class.
type IReductionFeatureChannelsArgumentMaxNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionFeatureChannelsArgumentMaxNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionFeatureChannelsArgumentMaxNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionFeatureChannelsArgumentMaxNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsArgumentMaxNodeClass) Alloc() ReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMaxNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionFeatureChannelsArgumentMaxNodeClass) New() ReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMaxNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsArgumentMaxNode) Init() ReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMaxNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsArgumentMaxNode) Autorelease() ReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMaxNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsArgumentMaxNode creates a new ReductionFeatureChannelsArgumentMaxNode instance.
func NewReductionFeatureChannelsArgumentMaxNode() ReductionFeatureChannelsArgumentMaxNode {
	return getReductionFeatureChannelsArgumentMaxNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionFeatureChannelsArgumentMaxNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsArgumentMaxNode
type ReductionFeatureChannelsArgumentMaxNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsArgumentMaxNodeFrom constructs a [ReductionFeatureChannelsArgumentMaxNode] from an unsafe.Pointer.
func ReductionFeatureChannelsArgumentMaxNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsArgumentMaxNode {
	return ReductionFeatureChannelsArgumentMaxNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionFeatureChannelsArgumentMaxNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionFeatureChannelsArgumentMaxNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionFeatureChannelsArgumentMaxNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionFeatureChannelsArgumentMaxNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionFeatureChannelsArgumentMaxNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionFeatureChannelsArgumentMaxNode */



