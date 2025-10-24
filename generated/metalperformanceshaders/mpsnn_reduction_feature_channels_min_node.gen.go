// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionFeatureChannelsMinNode */


/* debug [class_header]: Header for MPSNNReductionFeatureChannelsMinNode */
// The class instance for the [ReductionFeatureChannelsMinNode] class.
var (
	ReductionFeatureChannelsMinNodeClass     _ReductionFeatureChannelsMinNodeClass
	ReductionFeatureChannelsMinNodeClassOnce sync.Once
)

func getReductionFeatureChannelsMinNodeClass() _ReductionFeatureChannelsMinNodeClass {
	ReductionFeatureChannelsMinNodeClassOnce.Do(func() {
		ReductionFeatureChannelsMinNodeClass = _ReductionFeatureChannelsMinNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsMinNode")}
	})
	return ReductionFeatureChannelsMinNodeClass
}

type _ReductionFeatureChannelsMinNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionFeatureChannelsMinNode */
// An interface definition for the [ReductionFeatureChannelsMinNode] class.
type IReductionFeatureChannelsMinNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionFeatureChannelsMinNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionFeatureChannelsMinNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionFeatureChannelsMinNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsMinNodeClass) Alloc() ReductionFeatureChannelsMinNode {
	rv := objc.Send[ReductionFeatureChannelsMinNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionFeatureChannelsMinNodeClass) New() ReductionFeatureChannelsMinNode {
	rv := objc.Send[ReductionFeatureChannelsMinNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsMinNode) Init() ReductionFeatureChannelsMinNode {
	rv := objc.Send[ReductionFeatureChannelsMinNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsMinNode) Autorelease() ReductionFeatureChannelsMinNode {
	rv := objc.Send[ReductionFeatureChannelsMinNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsMinNode creates a new ReductionFeatureChannelsMinNode instance.
func NewReductionFeatureChannelsMinNode() ReductionFeatureChannelsMinNode {
	return getReductionFeatureChannelsMinNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionFeatureChannelsMinNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsMinNode
type ReductionFeatureChannelsMinNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsMinNodeFrom constructs a [ReductionFeatureChannelsMinNode] from an unsafe.Pointer.
func ReductionFeatureChannelsMinNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsMinNode {
	return ReductionFeatureChannelsMinNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionFeatureChannelsMinNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionFeatureChannelsMinNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionFeatureChannelsMinNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionFeatureChannelsMinNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionFeatureChannelsMinNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionFeatureChannelsMinNode */



