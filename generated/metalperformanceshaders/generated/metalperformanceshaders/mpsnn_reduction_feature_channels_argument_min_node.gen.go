// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionFeatureChannelsArgumentMinNode */


/* debug [class_header]: Header for MPSNNReductionFeatureChannelsArgumentMinNode */
// The class instance for the [ReductionFeatureChannelsArgumentMinNode] class.
var (
	ReductionFeatureChannelsArgumentMinNodeClass     _ReductionFeatureChannelsArgumentMinNodeClass
	ReductionFeatureChannelsArgumentMinNodeClassOnce sync.Once
)

func getReductionFeatureChannelsArgumentMinNodeClass() _ReductionFeatureChannelsArgumentMinNodeClass {
	ReductionFeatureChannelsArgumentMinNodeClassOnce.Do(func() {
		ReductionFeatureChannelsArgumentMinNodeClass = _ReductionFeatureChannelsArgumentMinNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsArgumentMinNode")}
	})
	return ReductionFeatureChannelsArgumentMinNodeClass
}

type _ReductionFeatureChannelsArgumentMinNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionFeatureChannelsArgumentMinNode */
// An interface definition for the [ReductionFeatureChannelsArgumentMinNode] class.
type IReductionFeatureChannelsArgumentMinNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionFeatureChannelsArgumentMinNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionFeatureChannelsArgumentMinNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionFeatureChannelsArgumentMinNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsArgumentMinNodeClass) Alloc() ReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMinNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionFeatureChannelsArgumentMinNodeClass) New() ReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMinNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsArgumentMinNode) Init() ReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMinNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsArgumentMinNode) Autorelease() ReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMinNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsArgumentMinNode creates a new ReductionFeatureChannelsArgumentMinNode instance.
func NewReductionFeatureChannelsArgumentMinNode() ReductionFeatureChannelsArgumentMinNode {
	return getReductionFeatureChannelsArgumentMinNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionFeatureChannelsArgumentMinNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsArgumentMinNode
type ReductionFeatureChannelsArgumentMinNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsArgumentMinNodeFrom constructs a [ReductionFeatureChannelsArgumentMinNode] from an unsafe.Pointer.
func ReductionFeatureChannelsArgumentMinNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsArgumentMinNode {
	return ReductionFeatureChannelsArgumentMinNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionFeatureChannelsArgumentMinNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionFeatureChannelsArgumentMinNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionFeatureChannelsArgumentMinNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionFeatureChannelsArgumentMinNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionFeatureChannelsArgumentMinNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionFeatureChannelsArgumentMinNode */



