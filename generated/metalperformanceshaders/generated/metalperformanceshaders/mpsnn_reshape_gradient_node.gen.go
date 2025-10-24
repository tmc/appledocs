// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReshapeGradientNode */


/* debug [class_header]: Header for MPSNNReshapeGradientNode */
// The class instance for the [ReshapeGradientNode] class.
var (
	ReshapeGradientNodeClass     _ReshapeGradientNodeClass
	ReshapeGradientNodeClassOnce sync.Once
)

func getReshapeGradientNodeClass() _ReshapeGradientNodeClass {
	ReshapeGradientNodeClassOnce.Do(func() {
		ReshapeGradientNodeClass = _ReshapeGradientNodeClass{objc.GetClass("MPSNNReshapeGradientNode")}
	})
	return ReshapeGradientNodeClass
}

type _ReshapeGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReshapeGradientNode */
// An interface definition for the [ReshapeGradientNode] class.
type IReshapeGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for ReshapeGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReshapeGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReshapeGradientNode */
// Alloc allocates a new instance without initialization.
func (rc _ReshapeGradientNodeClass) Alloc() ReshapeGradientNode {
	rv := objc.Send[ReshapeGradientNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReshapeGradientNodeClass) New() ReshapeGradientNode {
	rv := objc.Send[ReshapeGradientNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReshapeGradientNode) Init() ReshapeGradientNode {
	rv := objc.Send[ReshapeGradientNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReshapeGradientNode) Autorelease() ReshapeGradientNode {
	rv := objc.Send[ReshapeGradientNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReshapeGradientNode creates a new ReshapeGradientNode instance.
func NewReshapeGradientNode() ReshapeGradientNode {
	return getReshapeGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReshapeGradientNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradientNode
type ReshapeGradientNode struct {
	GradientFilterNode
}

// ReshapeGradientNodeFrom constructs a [ReshapeGradientNode] from an unsafe.Pointer.
func ReshapeGradientNodeFrom(ptr unsafe.Pointer) ReshapeGradientNode {
	return ReshapeGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReshapeGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapegradientnode/3037417-initwithsourcegradient
func NewReshapeGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) ReshapeGradientNode {
	instance := getReshapeGradientNodeClass().Alloc()
	rv := objc.Send[ReshapeGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReshapeGradientNodeWithSourceGradientSourceImageGradientState */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReshapeGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapegradientnode/3037418-nodewithsourcegradient
func (rc _ReshapeGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientState) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReshapeGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReshapeGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReshapeGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReshapeGradientNode */


