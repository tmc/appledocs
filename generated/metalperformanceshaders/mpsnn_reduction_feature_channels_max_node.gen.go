// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionFeatureChannelsMaxNode */


/* debug [class_header]: Header for MPSNNReductionFeatureChannelsMaxNode */
// The class instance for the [ReductionFeatureChannelsMaxNode] class.
var (
	ReductionFeatureChannelsMaxNodeClass     _ReductionFeatureChannelsMaxNodeClass
	ReductionFeatureChannelsMaxNodeClassOnce sync.Once
)

func getReductionFeatureChannelsMaxNodeClass() _ReductionFeatureChannelsMaxNodeClass {
	ReductionFeatureChannelsMaxNodeClassOnce.Do(func() {
		ReductionFeatureChannelsMaxNodeClass = _ReductionFeatureChannelsMaxNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsMaxNode")}
	})
	return ReductionFeatureChannelsMaxNodeClass
}

type _ReductionFeatureChannelsMaxNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionFeatureChannelsMaxNode */
// An interface definition for the [ReductionFeatureChannelsMaxNode] class.
type IReductionFeatureChannelsMaxNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionFeatureChannelsMaxNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionFeatureChannelsMaxNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionFeatureChannelsMaxNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsMaxNodeClass) Alloc() ReductionFeatureChannelsMaxNode {
	rv := objc.Send[ReductionFeatureChannelsMaxNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionFeatureChannelsMaxNodeClass) New() ReductionFeatureChannelsMaxNode {
	rv := objc.Send[ReductionFeatureChannelsMaxNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsMaxNode) Init() ReductionFeatureChannelsMaxNode {
	rv := objc.Send[ReductionFeatureChannelsMaxNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsMaxNode) Autorelease() ReductionFeatureChannelsMaxNode {
	rv := objc.Send[ReductionFeatureChannelsMaxNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsMaxNode creates a new ReductionFeatureChannelsMaxNode instance.
func NewReductionFeatureChannelsMaxNode() ReductionFeatureChannelsMaxNode {
	return getReductionFeatureChannelsMaxNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionFeatureChannelsMaxNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsMaxNode
type ReductionFeatureChannelsMaxNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsMaxNodeFrom constructs a [ReductionFeatureChannelsMaxNode] from an unsafe.Pointer.
func ReductionFeatureChannelsMaxNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsMaxNode {
	return ReductionFeatureChannelsMaxNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionFeatureChannelsMaxNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionFeatureChannelsMaxNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionFeatureChannelsMaxNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionFeatureChannelsMaxNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionFeatureChannelsMaxNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionFeatureChannelsMaxNode */



