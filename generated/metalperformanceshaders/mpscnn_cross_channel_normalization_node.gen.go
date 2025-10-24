// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CNNCrossChannelNormalizationNode] class.
type ICNNCrossChannelNormalizationNode interface {
	ICNNNormalizationNode
	

	// properties:
	KernelSizeInFeatureChannels() objectivec.IObject
	SetKernelSizeInFeatureChannels(value objectivec.IObject)


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866459-initwithsource
func NewCNNCrossChannelNormalizationNodeWithSource(sourceNode IImageNode) CNNCrossChannelNormalizationNode {
	instance := getCNNCrossChannelNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalizationNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866456-initwithsource
func NewCNNCrossChannelNormalizationNodeWithSourceKernelSize(sourceNode IImageNode, kernelSize uint) CNNCrossChannelNormalizationNode {
	instance := getCNNCrossChannelNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalizationNode](instance.ID, objc.Sel("initWithSource:kernelSize:"), sourceNode, kernelSize)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866476-nodewithsource
func (cc _CNNCrossChannelNormalizationNodeClass) NodeWithSourceKernelSize(sourceNode IImageNode, kernelSize uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:kernelSize:"), sourceNode, kernelSize)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866419-kernelsizeinfeaturechannels
func (c_ CNNCrossChannelNormalizationNode) KernelSizeInFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelSizeInFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866419-kernelsizeinfeaturechannels
func (c_ CNNCrossChannelNormalizationNode) SetKernelSizeInFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelSizeInFeatureChannels:"), value)
}







