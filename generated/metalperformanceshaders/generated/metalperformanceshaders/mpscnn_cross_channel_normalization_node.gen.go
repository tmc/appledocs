// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNCrossChannelNormalizationNode */


/* debug [class_header]: Header for MPSCNNCrossChannelNormalizationNode */
// The class instance for the [CNNCrossChannelNormalizationNode] class.
var (
	CNNCrossChannelNormalizationNodeClass     _CNNCrossChannelNormalizationNodeClass
	CNNCrossChannelNormalizationNodeClassOnce sync.Once
)

func getCNNCrossChannelNormalizationNodeClass() _CNNCrossChannelNormalizationNodeClass {
	CNNCrossChannelNormalizationNodeClassOnce.Do(func() {
		CNNCrossChannelNormalizationNodeClass = _CNNCrossChannelNormalizationNodeClass{objc.GetClass("MPSCNNCrossChannelNormalizationNode")}
	})
	return CNNCrossChannelNormalizationNodeClass
}

type _CNNCrossChannelNormalizationNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNCrossChannelNormalizationNode */
// An interface definition for the [CNNCrossChannelNormalizationNode] class.
type ICNNCrossChannelNormalizationNode interface {
	ICNNNormalizationNode
	
/* debug [class_interface_properties]: Properties for CNNCrossChannelNormalizationNode */
	// properties:
	KernelSizeInFeatureChannels() objectivec.IObject
	SetKernelSizeInFeatureChannels(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNCrossChannelNormalizationNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNCrossChannelNormalizationNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNCrossChannelNormalizationNodeClass) Alloc() CNNCrossChannelNormalizationNode {
	rv := objc.Send[CNNCrossChannelNormalizationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNCrossChannelNormalizationNodeClass) New() CNNCrossChannelNormalizationNode {
	rv := objc.Send[CNNCrossChannelNormalizationNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNCrossChannelNormalizationNode) Init() CNNCrossChannelNormalizationNode {
	rv := objc.Send[CNNCrossChannelNormalizationNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNCrossChannelNormalizationNode) Autorelease() CNNCrossChannelNormalizationNode {
	rv := objc.Send[CNNCrossChannelNormalizationNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNCrossChannelNormalizationNode creates a new CNNCrossChannelNormalizationNode instance.
func NewCNNCrossChannelNormalizationNode() CNNCrossChannelNormalizationNode {
	return getCNNCrossChannelNormalizationNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNCrossChannelNormalizationNode */
// A representation of a normalization kernel across feature channels.


// A representation of a normalization kernel across feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationNode
type CNNCrossChannelNormalizationNode struct {
	CNNNormalizationNode
}

// CNNCrossChannelNormalizationNodeFrom constructs a [CNNCrossChannelNormalizationNode] from an unsafe.Pointer.
//
// A representation of a normalization kernel across feature channels.
func CNNCrossChannelNormalizationNodeFrom(ptr unsafe.Pointer) CNNCrossChannelNormalizationNode {
	return CNNCrossChannelNormalizationNode{
		CNNNormalizationNode: CNNNormalizationNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNCrossChannelNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866459-initwithsource
func NewCNNCrossChannelNormalizationNodeWithSource(sourceNode IImageNode) CNNCrossChannelNormalizationNode {
	instance := getCNNCrossChannelNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalizationNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNCrossChannelNormalizationNodeWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866456-initwithsource
func NewCNNCrossChannelNormalizationNodeWithSourceKernelSize(sourceNode IImageNode, kernelSize uint) CNNCrossChannelNormalizationNode {
	instance := getCNNCrossChannelNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalizationNode](instance.ID, objc.Sel("initWithSource:kernelSize:"), sourceNode, kernelSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNCrossChannelNormalizationNodeWithSourceKernelSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNCrossChannelNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866476-nodewithsource
func (cc _CNNCrossChannelNormalizationNodeClass) NodeWithSourceKernelSize(sourceNode IImageNode, kernelSize uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:kernelSize:"), sourceNode, kernelSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceKernelSize) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNCrossChannelNormalizationNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNCrossChannelNormalizationNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNCrossChannelNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866419-kernelsizeinfeaturechannels
func (c_ CNNCrossChannelNormalizationNode) KernelSizeInFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelSizeInFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: kernelSizeInFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866419-kernelsizeinfeaturechannels
func (c_ CNNCrossChannelNormalizationNode) SetKernelSizeInFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelSizeInFeatureChannels:"), value)
}/* debug [instance_properties/setter]: kernelSizeInFeatureChannels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNCrossChannelNormalizationNode */


