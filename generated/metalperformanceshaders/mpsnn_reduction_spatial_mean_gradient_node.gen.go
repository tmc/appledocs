// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReductionSpatialMeanGradientNode */


/* debug [class_header]: Header for MPSNNReductionSpatialMeanGradientNode */
// The class instance for the [ReductionSpatialMeanGradientNode] class.
var (
	ReductionSpatialMeanGradientNodeClass     _ReductionSpatialMeanGradientNodeClass
	ReductionSpatialMeanGradientNodeClassOnce sync.Once
)

func getReductionSpatialMeanGradientNodeClass() _ReductionSpatialMeanGradientNodeClass {
	ReductionSpatialMeanGradientNodeClassOnce.Do(func() {
		ReductionSpatialMeanGradientNodeClass = _ReductionSpatialMeanGradientNodeClass{objc.GetClass("MPSNNReductionSpatialMeanGradientNode")}
	})
	return ReductionSpatialMeanGradientNodeClass
}

type _ReductionSpatialMeanGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReductionSpatialMeanGradientNode */
// An interface definition for the [ReductionSpatialMeanGradientNode] class.
type IReductionSpatialMeanGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for ReductionSpatialMeanGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReductionSpatialMeanGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReductionSpatialMeanGradientNode */
// Alloc allocates a new instance without initialization.
func (rc _ReductionSpatialMeanGradientNodeClass) Alloc() ReductionSpatialMeanGradientNode {
	rv := objc.Send[ReductionSpatialMeanGradientNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionSpatialMeanGradientNodeClass) New() ReductionSpatialMeanGradientNode {
	rv := objc.Send[ReductionSpatialMeanGradientNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionSpatialMeanGradientNode) Init() ReductionSpatialMeanGradientNode {
	rv := objc.Send[ReductionSpatialMeanGradientNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionSpatialMeanGradientNode) Autorelease() ReductionSpatialMeanGradientNode {
	rv := objc.Send[ReductionSpatialMeanGradientNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionSpatialMeanGradientNode creates a new ReductionSpatialMeanGradientNode instance.
func NewReductionSpatialMeanGradientNode() ReductionSpatialMeanGradientNode {
	return getReductionSpatialMeanGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReductionSpatialMeanGradientNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionSpatialMeanGradientNode
type ReductionSpatialMeanGradientNode struct {
	GradientFilterNode
}

// ReductionSpatialMeanGradientNodeFrom constructs a [ReductionSpatialMeanGradientNode] from an unsafe.Pointer.
func ReductionSpatialMeanGradientNodeFrom(ptr unsafe.Pointer) ReductionSpatialMeanGradientNode {
	return ReductionSpatialMeanGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReductionSpatialMeanGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionspatialmeangradientnode/3037413-initwithsourcegradient
func NewReductionSpatialMeanGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) ReductionSpatialMeanGradientNode {
	instance := getReductionSpatialMeanGradientNodeClass().Alloc()
	rv := objc.Send[ReductionSpatialMeanGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReductionSpatialMeanGradientNodeWithSourceGradientSourceImageGradientState */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReductionSpatialMeanGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionspatialmeangradientnode/3037414-nodewithsourcegradient
func (rc _ReductionSpatialMeanGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientState) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReductionSpatialMeanGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReductionSpatialMeanGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReductionSpatialMeanGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReductionSpatialMeanGradientNode */


