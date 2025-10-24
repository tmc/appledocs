// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNSoftMaxGradientNode */


/* debug [class_header]: Header for MPSCNNSoftMaxGradientNode */
// The class instance for the [CNNSoftMaxGradientNode] class.
var (
	CNNSoftMaxGradientNodeClass     _CNNSoftMaxGradientNodeClass
	CNNSoftMaxGradientNodeClassOnce sync.Once
)

func getCNNSoftMaxGradientNodeClass() _CNNSoftMaxGradientNodeClass {
	CNNSoftMaxGradientNodeClassOnce.Do(func() {
		CNNSoftMaxGradientNodeClass = _CNNSoftMaxGradientNodeClass{objc.GetClass("MPSCNNSoftMaxGradientNode")}
	})
	return CNNSoftMaxGradientNodeClass
}

type _CNNSoftMaxGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNSoftMaxGradientNode */
// An interface definition for the [CNNSoftMaxGradientNode] class.
type ICNNSoftMaxGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for CNNSoftMaxGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNSoftMaxGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNSoftMaxGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNSoftMaxGradientNodeClass) Alloc() CNNSoftMaxGradientNode {
	rv := objc.Send[CNNSoftMaxGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSoftMaxGradientNodeClass) New() CNNSoftMaxGradientNode {
	rv := objc.Send[CNNSoftMaxGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSoftMaxGradientNode) Init() CNNSoftMaxGradientNode {
	rv := objc.Send[CNNSoftMaxGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSoftMaxGradientNode) Autorelease() CNNSoftMaxGradientNode {
	rv := objc.Send[CNNSoftMaxGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSoftMaxGradientNode creates a new CNNSoftMaxGradientNode instance.
func NewCNNSoftMaxGradientNode() CNNSoftMaxGradientNode {
	return getCNNSoftMaxGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNSoftMaxGradientNode */
// A representation of a gradient softmax filter.


// A representation of a gradient softmax filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradientNode
type CNNSoftMaxGradientNode struct {
	GradientFilterNode
}

// CNNSoftMaxGradientNodeFrom constructs a [CNNSoftMaxGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient softmax filter.
func CNNSoftMaxGradientNodeFrom(ptr unsafe.Pointer) CNNSoftMaxGradientNode {
	return CNNSoftMaxGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNSoftMaxGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradientnode/2948039-initwithsourcegradient
func NewCNNSoftMaxGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) CNNSoftMaxGradientNode {
	instance := getCNNSoftMaxGradientNodeClass().Alloc()
	rv := objc.Send[CNNSoftMaxGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNSoftMaxGradientNodeWithSourceGradientSourceImageGradientState */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNSoftMaxGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradientnode/2947995-nodewithsourcegradient
func (cc _CNNSoftMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientState) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNSoftMaxGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNSoftMaxGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNSoftMaxGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNSoftMaxGradientNode */


