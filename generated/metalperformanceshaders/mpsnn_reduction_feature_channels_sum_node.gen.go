// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReductionFeatureChannelsSumNode */


/* debug [class_header]: Header for MPSNNReductionFeatureChannelsSumNode */
// The class instance for the [ReductionFeatureChannelsSumNode] class.
var (
	ReductionFeatureChannelsSumNodeClass     _ReductionFeatureChannelsSumNodeClass
	ReductionFeatureChannelsSumNodeClassOnce sync.Once
)

func getReductionFeatureChannelsSumNodeClass() _ReductionFeatureChannelsSumNodeClass {
	ReductionFeatureChannelsSumNodeClassOnce.Do(func() {
		ReductionFeatureChannelsSumNodeClass = _ReductionFeatureChannelsSumNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsSumNode")}
	})
	return ReductionFeatureChannelsSumNodeClass
}

type _ReductionFeatureChannelsSumNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionFeatureChannelsSumNode */
// An interface definition for the [ReductionFeatureChannelsSumNode] class.
type IReductionFeatureChannelsSumNode interface {
	IUnaryReductionNode
	
/* debug [class_interface_properties]: Properties for ReductionFeatureChannelsSumNode */
	// properties:
	Weight() objectivec.IObject
	SetWeight(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionFeatureChannelsSumNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionFeatureChannelsSumNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsSumNodeClass) Alloc() ReductionFeatureChannelsSumNode {
	rv := objc.Send[ReductionFeatureChannelsSumNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionFeatureChannelsSumNodeClass) New() ReductionFeatureChannelsSumNode {
	rv := objc.Send[ReductionFeatureChannelsSumNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsSumNode) Init() ReductionFeatureChannelsSumNode {
	rv := objc.Send[ReductionFeatureChannelsSumNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsSumNode) Autorelease() ReductionFeatureChannelsSumNode {
	rv := objc.Send[ReductionFeatureChannelsSumNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsSumNode creates a new ReductionFeatureChannelsSumNode instance.
func NewReductionFeatureChannelsSumNode() ReductionFeatureChannelsSumNode {
	return getReductionFeatureChannelsSumNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionFeatureChannelsSumNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsSumNode
type ReductionFeatureChannelsSumNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsSumNodeFrom constructs a [ReductionFeatureChannelsSumNode] from an unsafe.Pointer.
func ReductionFeatureChannelsSumNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsSumNode {
	return ReductionFeatureChannelsSumNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionFeatureChannelsSumNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionFeatureChannelsSumNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionFeatureChannelsSumNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionFeatureChannelsSumNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionFeatureChannelsSumNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionfeaturechannelssumnode/3037407-weight
func (r_ ReductionFeatureChannelsSumNode) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionfeaturechannelssumnode/3037407-weight
func (r_ ReductionFeatureChannelsSumNode) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setWeight:"), value)
}/* debug [instance_properties/setter]: weight */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionFeatureChannelsSumNode */



