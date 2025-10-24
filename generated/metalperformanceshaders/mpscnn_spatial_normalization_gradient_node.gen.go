// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNSpatialNormalizationGradientNode */


/* debug [class_header]: Header for MPSCNNSpatialNormalizationGradientNode */
// The class instance for the [CNNSpatialNormalizationGradientNode] class.
var (
	CNNSpatialNormalizationGradientNodeClass     _CNNSpatialNormalizationGradientNodeClass
	CNNSpatialNormalizationGradientNodeClassOnce sync.Once
)

func getCNNSpatialNormalizationGradientNodeClass() _CNNSpatialNormalizationGradientNodeClass {
	CNNSpatialNormalizationGradientNodeClassOnce.Do(func() {
		CNNSpatialNormalizationGradientNodeClass = _CNNSpatialNormalizationGradientNodeClass{objc.GetClass("MPSCNNSpatialNormalizationGradientNode")}
	})
	return CNNSpatialNormalizationGradientNodeClass
}

type _CNNSpatialNormalizationGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNSpatialNormalizationGradientNode */
// An interface definition for the [CNNSpatialNormalizationGradientNode] class.
type ICNNSpatialNormalizationGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for CNNSpatialNormalizationGradientNode */
	// properties:
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNSpatialNormalizationGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNSpatialNormalizationGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNSpatialNormalizationGradientNodeClass) Alloc() CNNSpatialNormalizationGradientNode {
	rv := objc.Send[CNNSpatialNormalizationGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSpatialNormalizationGradientNodeClass) New() CNNSpatialNormalizationGradientNode {
	rv := objc.Send[CNNSpatialNormalizationGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSpatialNormalizationGradientNode) Init() CNNSpatialNormalizationGradientNode {
	rv := objc.Send[CNNSpatialNormalizationGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSpatialNormalizationGradientNode) Autorelease() CNNSpatialNormalizationGradientNode {
	rv := objc.Send[CNNSpatialNormalizationGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSpatialNormalizationGradientNode creates a new CNNSpatialNormalizationGradientNode instance.
func NewCNNSpatialNormalizationGradientNode() CNNSpatialNormalizationGradientNode {
	return getCNNSpatialNormalizationGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNSpatialNormalizationGradientNode */
// A representation of a gradient spatial normalization kernel.


// A representation of a gradient spatial normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode
type CNNSpatialNormalizationGradientNode struct {
	GradientFilterNode
}

// CNNSpatialNormalizationGradientNodeFrom constructs a [CNNSpatialNormalizationGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient spatial normalization kernel.
func CNNSpatialNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNSpatialNormalizationGradientNode {
	return CNNSpatialNormalizationGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNSpatialNormalizationGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948009-initwithsourcegradient
func NewCNNSpatialNormalizationGradientNodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelSize uint) CNNSpatialNormalizationGradientNode {
	instance := getCNNSpatialNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[CNNSpatialNormalizationGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelSize:"), sourceGradient, sourceImage, gradientState, kernelSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNSpatialNormalizationGradientNodeWithSourceGradientSourceImageGradientStateKernelSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNSpatialNormalizationGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2947978-nodewithsourcegradient
func (cc _CNNSpatialNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelSize uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelSize:"), sourceGradient, sourceImage, gradientState, kernelSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientStateKernelSize) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNSpatialNormalizationGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNSpatialNormalizationGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNSpatialNormalizationGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2947968-delta
func (c_ CNNSpatialNormalizationGradientNode) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2947968-delta
func (c_ CNNSpatialNormalizationGradientNode) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948006-beta
func (c_ CNNSpatialNormalizationGradientNode) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948006-beta
func (c_ CNNSpatialNormalizationGradientNode) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}/* debug [instance_properties/setter]: beta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948013-kernelwidth
func (c_ CNNSpatialNormalizationGradientNode) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948013-kernelwidth
func (c_ CNNSpatialNormalizationGradientNode) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948020-kernelheight
func (c_ CNNSpatialNormalizationGradientNode) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948020-kernelheight
func (c_ CNNSpatialNormalizationGradientNode) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948027-alpha
func (c_ CNNSpatialNormalizationGradientNode) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948027-alpha
func (c_ CNNSpatialNormalizationGradientNode) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNSpatialNormalizationGradientNode */


