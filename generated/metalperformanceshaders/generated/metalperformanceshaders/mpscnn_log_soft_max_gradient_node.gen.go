// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNLogSoftMaxGradientNode */


/* debug [class_header]: Header for MPSCNNLogSoftMaxGradientNode */
// The class instance for the [CNNLogSoftMaxGradientNode] class.
var (
	CNNLogSoftMaxGradientNodeClass     _CNNLogSoftMaxGradientNodeClass
	CNNLogSoftMaxGradientNodeClassOnce sync.Once
)

func getCNNLogSoftMaxGradientNodeClass() _CNNLogSoftMaxGradientNodeClass {
	CNNLogSoftMaxGradientNodeClassOnce.Do(func() {
		CNNLogSoftMaxGradientNodeClass = _CNNLogSoftMaxGradientNodeClass{objc.GetClass("MPSCNNLogSoftMaxGradientNode")}
	})
	return CNNLogSoftMaxGradientNodeClass
}

type _CNNLogSoftMaxGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNLogSoftMaxGradientNode */
// An interface definition for the [CNNLogSoftMaxGradientNode] class.
type ICNNLogSoftMaxGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for CNNLogSoftMaxGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNLogSoftMaxGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNLogSoftMaxGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNLogSoftMaxGradientNodeClass) Alloc() CNNLogSoftMaxGradientNode {
	rv := objc.Send[CNNLogSoftMaxGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLogSoftMaxGradientNodeClass) New() CNNLogSoftMaxGradientNode {
	rv := objc.Send[CNNLogSoftMaxGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLogSoftMaxGradientNode) Init() CNNLogSoftMaxGradientNode {
	rv := objc.Send[CNNLogSoftMaxGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLogSoftMaxGradientNode) Autorelease() CNNLogSoftMaxGradientNode {
	rv := objc.Send[CNNLogSoftMaxGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLogSoftMaxGradientNode creates a new CNNLogSoftMaxGradientNode instance.
func NewCNNLogSoftMaxGradientNode() CNNLogSoftMaxGradientNode {
	return getCNNLogSoftMaxGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNLogSoftMaxGradientNode */
// A representation of a gradient logarithmic softmax filter kernel.


// A representation of a gradient logarithmic softmax filter kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradientNode
type CNNLogSoftMaxGradientNode struct {
	GradientFilterNode
}

// CNNLogSoftMaxGradientNodeFrom constructs a [CNNLogSoftMaxGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient logarithmic softmax filter kernel.
func CNNLogSoftMaxGradientNodeFrom(ptr unsafe.Pointer) CNNLogSoftMaxGradientNode {
	return CNNLogSoftMaxGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNLogSoftMaxGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradientnode/2947971-initwithsourcegradient
func NewCNNLogSoftMaxGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) CNNLogSoftMaxGradientNode {
	instance := getCNNLogSoftMaxGradientNodeClass().Alloc()
	rv := objc.Send[CNNLogSoftMaxGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLogSoftMaxGradientNodeWithSourceGradientSourceImageGradientState */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNLogSoftMaxGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradientnode/2947974-nodewithsourcegradient
func (cc _CNNLogSoftMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientState) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNLogSoftMaxGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNLogSoftMaxGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNLogSoftMaxGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNLogSoftMaxGradientNode */


