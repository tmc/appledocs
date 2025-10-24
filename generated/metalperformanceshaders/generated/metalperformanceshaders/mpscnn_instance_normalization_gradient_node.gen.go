// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNInstanceNormalizationGradientNode */


/* debug [class_header]: Header for MPSCNNInstanceNormalizationGradientNode */
// The class instance for the [CNNInstanceNormalizationGradientNode] class.
var (
	CNNInstanceNormalizationGradientNodeClass     _CNNInstanceNormalizationGradientNodeClass
	CNNInstanceNormalizationGradientNodeClassOnce sync.Once
)

func getCNNInstanceNormalizationGradientNodeClass() _CNNInstanceNormalizationGradientNodeClass {
	CNNInstanceNormalizationGradientNodeClassOnce.Do(func() {
		CNNInstanceNormalizationGradientNodeClass = _CNNInstanceNormalizationGradientNodeClass{objc.GetClass("MPSCNNInstanceNormalizationGradientNode")}
	})
	return CNNInstanceNormalizationGradientNodeClass
}

type _CNNInstanceNormalizationGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNInstanceNormalizationGradientNode */
// An interface definition for the [CNNInstanceNormalizationGradientNode] class.
type ICNNInstanceNormalizationGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for CNNInstanceNormalizationGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNInstanceNormalizationGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNInstanceNormalizationGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNInstanceNormalizationGradientNodeClass) Alloc() CNNInstanceNormalizationGradientNode {
	rv := objc.Send[CNNInstanceNormalizationGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNInstanceNormalizationGradientNodeClass) New() CNNInstanceNormalizationGradientNode {
	rv := objc.Send[CNNInstanceNormalizationGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNInstanceNormalizationGradientNode) Init() CNNInstanceNormalizationGradientNode {
	rv := objc.Send[CNNInstanceNormalizationGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNInstanceNormalizationGradientNode) Autorelease() CNNInstanceNormalizationGradientNode {
	rv := objc.Send[CNNInstanceNormalizationGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNInstanceNormalizationGradientNode creates a new CNNInstanceNormalizationGradientNode instance.
func NewCNNInstanceNormalizationGradientNode() CNNInstanceNormalizationGradientNode {
	return getCNNInstanceNormalizationGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNInstanceNormalizationGradientNode */
// A representation of a gradient instance normalization kernel.


// A representation of a gradient instance normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientNode
type CNNInstanceNormalizationGradientNode struct {
	GradientFilterNode
}

// CNNInstanceNormalizationGradientNodeFrom constructs a [CNNInstanceNormalizationGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient instance normalization kernel.
func CNNInstanceNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNInstanceNormalizationGradientNode {
	return CNNInstanceNormalizationGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNInstanceNormalizationGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientnode/2951954-initwithsourcegradient
func NewCNNInstanceNormalizationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) CNNInstanceNormalizationGradientNode {
	instance := getCNNInstanceNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[CNNInstanceNormalizationGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNInstanceNormalizationGradientNodeWithSourceGradientSourceImageGradientState */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNInstanceNormalizationGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientnode/2951932-nodewithsourcegradient
func (cc _CNNInstanceNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientState) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNInstanceNormalizationGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNInstanceNormalizationGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNInstanceNormalizationGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNInstanceNormalizationGradientNode */


