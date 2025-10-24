// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNBatchNormalizationGradientNode */


/* debug [class_header]: Header for MPSCNNBatchNormalizationGradientNode */
// The class instance for the [CNNBatchNormalizationGradientNode] class.
var (
	CNNBatchNormalizationGradientNodeClass     _CNNBatchNormalizationGradientNodeClass
	CNNBatchNormalizationGradientNodeClassOnce sync.Once
)

func getCNNBatchNormalizationGradientNodeClass() _CNNBatchNormalizationGradientNodeClass {
	CNNBatchNormalizationGradientNodeClassOnce.Do(func() {
		CNNBatchNormalizationGradientNodeClass = _CNNBatchNormalizationGradientNodeClass{objc.GetClass("MPSCNNBatchNormalizationGradientNode")}
	})
	return CNNBatchNormalizationGradientNodeClass
}

type _CNNBatchNormalizationGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNBatchNormalizationGradientNode */
// An interface definition for the [CNNBatchNormalizationGradientNode] class.
type ICNNBatchNormalizationGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for CNNBatchNormalizationGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNBatchNormalizationGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNBatchNormalizationGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNBatchNormalizationGradientNodeClass) Alloc() CNNBatchNormalizationGradientNode {
	rv := objc.Send[CNNBatchNormalizationGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBatchNormalizationGradientNodeClass) New() CNNBatchNormalizationGradientNode {
	rv := objc.Send[CNNBatchNormalizationGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBatchNormalizationGradientNode) Init() CNNBatchNormalizationGradientNode {
	rv := objc.Send[CNNBatchNormalizationGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBatchNormalizationGradientNode) Autorelease() CNNBatchNormalizationGradientNode {
	rv := objc.Send[CNNBatchNormalizationGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBatchNormalizationGradientNode creates a new CNNBatchNormalizationGradientNode instance.
func NewCNNBatchNormalizationGradientNode() CNNBatchNormalizationGradientNode {
	return getCNNBatchNormalizationGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNBatchNormalizationGradientNode */
// A representation of a gradient batch normalization kernel.


// A representation of a gradient batch normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradientNode
type CNNBatchNormalizationGradientNode struct {
	GradientFilterNode
}

// CNNBatchNormalizationGradientNodeFrom constructs a [CNNBatchNormalizationGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient batch normalization kernel.
func CNNBatchNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNBatchNormalizationGradientNode {
	return CNNBatchNormalizationGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNBatchNormalizationGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradientnode/2953939-initwithsourcegradient
func NewCNNBatchNormalizationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) CNNBatchNormalizationGradientNode {
	instance := getCNNBatchNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[CNNBatchNormalizationGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBatchNormalizationGradientNodeWithSourceGradientSourceImageGradientState */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNBatchNormalizationGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradientnode/2953943-nodewithsourcegradient
func (cc _CNNBatchNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientState) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNBatchNormalizationGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNBatchNormalizationGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNBatchNormalizationGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNBatchNormalizationGradientNode */


