// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNLocalContrastNormalizationNode */


/* debug [class_header]: Header for MPSCNNLocalContrastNormalizationNode */
// The class instance for the [CNNLocalContrastNormalizationNode] class.
var (
	CNNLocalContrastNormalizationNodeClass     _CNNLocalContrastNormalizationNodeClass
	CNNLocalContrastNormalizationNodeClassOnce sync.Once
)

func getCNNLocalContrastNormalizationNodeClass() _CNNLocalContrastNormalizationNodeClass {
	CNNLocalContrastNormalizationNodeClassOnce.Do(func() {
		CNNLocalContrastNormalizationNodeClass = _CNNLocalContrastNormalizationNodeClass{objc.GetClass("MPSCNNLocalContrastNormalizationNode")}
	})
	return CNNLocalContrastNormalizationNodeClass
}

type _CNNLocalContrastNormalizationNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNLocalContrastNormalizationNode */
// An interface definition for the [CNNLocalContrastNormalizationNode] class.
type ICNNLocalContrastNormalizationNode interface {
	ICNNNormalizationNode
	
/* debug [class_interface_properties]: Properties for CNNLocalContrastNormalizationNode */
	// properties:
	Pm() objectivec.IObject
	SetPm(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	Ps() objectivec.IObject
	SetPs(value objectivec.IObject)
	P0() objectivec.IObject
	SetP0(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNLocalContrastNormalizationNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNLocalContrastNormalizationNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNLocalContrastNormalizationNodeClass) Alloc() CNNLocalContrastNormalizationNode {
	rv := objc.Send[CNNLocalContrastNormalizationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLocalContrastNormalizationNodeClass) New() CNNLocalContrastNormalizationNode {
	rv := objc.Send[CNNLocalContrastNormalizationNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLocalContrastNormalizationNode) Init() CNNLocalContrastNormalizationNode {
	rv := objc.Send[CNNLocalContrastNormalizationNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLocalContrastNormalizationNode) Autorelease() CNNLocalContrastNormalizationNode {
	rv := objc.Send[CNNLocalContrastNormalizationNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLocalContrastNormalizationNode creates a new CNNLocalContrastNormalizationNode instance.
func NewCNNLocalContrastNormalizationNode() CNNLocalContrastNormalizationNode {
	return getCNNLocalContrastNormalizationNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNLocalContrastNormalizationNode */
// A representation of a local-contrast normalization kernel.


// A representation of a local-contrast normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode
type CNNLocalContrastNormalizationNode struct {
	CNNNormalizationNode
}

// CNNLocalContrastNormalizationNodeFrom constructs a [CNNLocalContrastNormalizationNode] from an unsafe.Pointer.
//
// A representation of a local-contrast normalization kernel.
func CNNLocalContrastNormalizationNodeFrom(ptr unsafe.Pointer) CNNLocalContrastNormalizationNode {
	return CNNLocalContrastNormalizationNode{
		CNNNormalizationNode: CNNNormalizationNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNLocalContrastNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866454-initwithsource
func NewCNNLocalContrastNormalizationNodeWithSource(sourceNode IImageNode) CNNLocalContrastNormalizationNode {
	instance := getCNNLocalContrastNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNLocalContrastNormalizationNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLocalContrastNormalizationNodeWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866473-initwithsource
func NewCNNLocalContrastNormalizationNodeWithSourceKernelSize(sourceNode IImageNode, kernelSize uint) CNNLocalContrastNormalizationNode {
	instance := getCNNLocalContrastNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNLocalContrastNormalizationNode](instance.ID, objc.Sel("initWithSource:kernelSize:"), sourceNode, kernelSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLocalContrastNormalizationNodeWithSourceKernelSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNLocalContrastNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866439-nodewithsource
func (cc _CNNLocalContrastNormalizationNodeClass) NodeWithSourceKernelSize(sourceNode IImageNode, kernelSize uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:kernelSize:"), sourceNode, kernelSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceKernelSize) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNLocalContrastNormalizationNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNLocalContrastNormalizationNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNLocalContrastNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866404-pm
func (c_ CNNLocalContrastNormalizationNode) Pm() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("pm"))
	return rv
}/* debug [instance_properties/getter]: pm */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866404-pm
func (c_ CNNLocalContrastNormalizationNode) SetPm(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPm:"), value)
}/* debug [instance_properties/setter]: pm */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866441-kernelwidth
func (c_ CNNLocalContrastNormalizationNode) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866441-kernelwidth
func (c_ CNNLocalContrastNormalizationNode) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866485-kernelheight
func (c_ CNNLocalContrastNormalizationNode) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866485-kernelheight
func (c_ CNNLocalContrastNormalizationNode) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866500-ps
func (c_ CNNLocalContrastNormalizationNode) Ps() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("ps"))
	return rv
}/* debug [instance_properties/getter]: ps */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866500-ps
func (c_ CNNLocalContrastNormalizationNode) SetPs(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPs:"), value)
}/* debug [instance_properties/setter]: ps */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866510-p0
func (c_ CNNLocalContrastNormalizationNode) P0() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("p0"))
	return rv
}/* debug [instance_properties/getter]: p0 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866510-p0
func (c_ CNNLocalContrastNormalizationNode) SetP0(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setP0:"), value)
}/* debug [instance_properties/setter]: p0 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNLocalContrastNormalizationNode */


