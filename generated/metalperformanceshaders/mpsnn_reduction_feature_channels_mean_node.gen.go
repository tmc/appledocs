// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReductionFeatureChannelsMeanNode */


/* debug [class_header]: Header for MPSNNReductionFeatureChannelsMeanNode */
// The class instance for the [ReductionFeatureChannelsMeanNode] class.
var (
	ReductionFeatureChannelsMeanNodeClass     _ReductionFeatureChannelsMeanNodeClass
	ReductionFeatureChannelsMeanNodeClassOnce sync.Once
)

func getReductionFeatureChannelsMeanNodeClass() _ReductionFeatureChannelsMeanNodeClass {
	ReductionFeatureChannelsMeanNodeClassOnce.Do(func() {
		ReductionFeatureChannelsMeanNodeClass = _ReductionFeatureChannelsMeanNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsMeanNode")}
	})
	return ReductionFeatureChannelsMeanNodeClass
}

type _ReductionFeatureChannelsMeanNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionFeatureChannelsMeanNode */
// An interface definition for the [ReductionFeatureChannelsMeanNode] class.
type IReductionFeatureChannelsMeanNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionFeatureChannelsMeanNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionFeatureChannelsMeanNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionFeatureChannelsMeanNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsMeanNodeClass) Alloc() ReductionFeatureChannelsMeanNode {
	rv := objc.Send[ReductionFeatureChannelsMeanNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionFeatureChannelsMeanNodeClass) New() ReductionFeatureChannelsMeanNode {
	rv := objc.Send[ReductionFeatureChannelsMeanNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsMeanNode) Init() ReductionFeatureChannelsMeanNode {
	rv := objc.Send[ReductionFeatureChannelsMeanNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsMeanNode) Autorelease() ReductionFeatureChannelsMeanNode {
	rv := objc.Send[ReductionFeatureChannelsMeanNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsMeanNode creates a new ReductionFeatureChannelsMeanNode instance.
func NewReductionFeatureChannelsMeanNode() ReductionFeatureChannelsMeanNode {
	return getReductionFeatureChannelsMeanNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionFeatureChannelsMeanNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsMeanNode
type ReductionFeatureChannelsMeanNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsMeanNodeFrom constructs a [ReductionFeatureChannelsMeanNode] from an unsafe.Pointer.
func ReductionFeatureChannelsMeanNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsMeanNode {
	return ReductionFeatureChannelsMeanNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionFeatureChannelsMeanNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionFeatureChannelsMeanNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionFeatureChannelsMeanNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionFeatureChannelsMeanNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionFeatureChannelsMeanNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionFeatureChannelsMeanNode */



