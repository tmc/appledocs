// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNSpatialNormalizationNode */


/* debug [class_header]: Header for MPSCNNSpatialNormalizationNode */
// The class instance for the [CNNSpatialNormalizationNode] class.
var (
	CNNSpatialNormalizationNodeClass     _CNNSpatialNormalizationNodeClass
	CNNSpatialNormalizationNodeClassOnce sync.Once
)

func getCNNSpatialNormalizationNodeClass() _CNNSpatialNormalizationNodeClass {
	CNNSpatialNormalizationNodeClassOnce.Do(func() {
		CNNSpatialNormalizationNodeClass = _CNNSpatialNormalizationNodeClass{objc.GetClass("MPSCNNSpatialNormalizationNode")}
	})
	return CNNSpatialNormalizationNodeClass
}

type _CNNSpatialNormalizationNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNSpatialNormalizationNode */
// An interface definition for the [CNNSpatialNormalizationNode] class.
type ICNNSpatialNormalizationNode interface {
	ICNNNormalizationNode
	
/* debug [class_interface_properties]: Properties for CNNSpatialNormalizationNode */
	// properties:
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNSpatialNormalizationNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNSpatialNormalizationNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNSpatialNormalizationNodeClass) Alloc() CNNSpatialNormalizationNode {
	rv := objc.Send[CNNSpatialNormalizationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSpatialNormalizationNodeClass) New() CNNSpatialNormalizationNode {
	rv := objc.Send[CNNSpatialNormalizationNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSpatialNormalizationNode) Init() CNNSpatialNormalizationNode {
	rv := objc.Send[CNNSpatialNormalizationNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSpatialNormalizationNode) Autorelease() CNNSpatialNormalizationNode {
	rv := objc.Send[CNNSpatialNormalizationNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSpatialNormalizationNode creates a new CNNSpatialNormalizationNode instance.
func NewCNNSpatialNormalizationNode() CNNSpatialNormalizationNode {
	return getCNNSpatialNormalizationNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNSpatialNormalizationNode */
// A representation of a spatial normalization kernel.


// A representation of a spatial normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationNode
type CNNSpatialNormalizationNode struct {
	CNNNormalizationNode
}

// CNNSpatialNormalizationNodeFrom constructs a [CNNSpatialNormalizationNode] from an unsafe.Pointer.
//
// A representation of a spatial normalization kernel.
func CNNSpatialNormalizationNodeFrom(ptr unsafe.Pointer) CNNSpatialNormalizationNode {
	return CNNSpatialNormalizationNode{
		CNNNormalizationNode: CNNNormalizationNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNSpatialNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866502-initwithsource
func NewCNNSpatialNormalizationNodeWithSource(sourceNode IImageNode) CNNSpatialNormalizationNode {
	instance := getCNNSpatialNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNSpatialNormalizationNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNSpatialNormalizationNodeWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866438-initwithsource
func NewCNNSpatialNormalizationNodeWithSourceKernelSize(sourceNode IImageNode, kernelSize uint) CNNSpatialNormalizationNode {
	instance := getCNNSpatialNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNSpatialNormalizationNode](instance.ID, objc.Sel("initWithSource:kernelSize:"), sourceNode, kernelSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNSpatialNormalizationNodeWithSourceKernelSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNSpatialNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866401-nodewithsource
func (cc _CNNSpatialNormalizationNodeClass) NodeWithSourceKernelSize(sourceNode IImageNode, kernelSize uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:kernelSize:"), sourceNode, kernelSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceKernelSize) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNSpatialNormalizationNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNSpatialNormalizationNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNSpatialNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866402-kernelwidth
func (c_ CNNSpatialNormalizationNode) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866402-kernelwidth
func (c_ CNNSpatialNormalizationNode) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866424-kernelheight
func (c_ CNNSpatialNormalizationNode) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866424-kernelheight
func (c_ CNNSpatialNormalizationNode) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNSpatialNormalizationNode */


