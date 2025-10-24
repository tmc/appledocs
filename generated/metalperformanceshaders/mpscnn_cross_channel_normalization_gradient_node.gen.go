// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNCrossChannelNormalizationGradientNode] class.
var (
	CNNCrossChannelNormalizationGradientNodeClass     _CNNCrossChannelNormalizationGradientNodeClass
	CNNCrossChannelNormalizationGradientNodeClassOnce sync.Once
)

func getCNNCrossChannelNormalizationGradientNodeClass() _CNNCrossChannelNormalizationGradientNodeClass {
	CNNCrossChannelNormalizationGradientNodeClassOnce.Do(func() {
		CNNCrossChannelNormalizationGradientNodeClass = _CNNCrossChannelNormalizationGradientNodeClass{objc.GetClass("MPSCNNCrossChannelNormalizationGradientNode")}
	})
	return CNNCrossChannelNormalizationGradientNodeClass
}

type _CNNCrossChannelNormalizationGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNCrossChannelNormalizationGradientNode] class.
type ICNNCrossChannelNormalizationGradientNode interface {
	IGradientFilterNode
	

	// properties:
	KernelSize() objectivec.IObject
	SetKernelSize(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNCrossChannelNormalizationGradientNodeClass) Alloc() CNNCrossChannelNormalizationGradientNode {
	rv := objc.Send[CNNCrossChannelNormalizationGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNCrossChannelNormalizationGradientNodeClass) New() CNNCrossChannelNormalizationGradientNode {
	rv := objc.Send[CNNCrossChannelNormalizationGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNCrossChannelNormalizationGradientNode) Init() CNNCrossChannelNormalizationGradientNode {
	rv := objc.Send[CNNCrossChannelNormalizationGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNCrossChannelNormalizationGradientNode) Autorelease() CNNCrossChannelNormalizationGradientNode {
	rv := objc.Send[CNNCrossChannelNormalizationGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNCrossChannelNormalizationGradientNode creates a new CNNCrossChannelNormalizationGradientNode instance.
func NewCNNCrossChannelNormalizationGradientNode() CNNCrossChannelNormalizationGradientNode {
	return getCNNCrossChannelNormalizationGradientNodeClass().New()
}





// A representation of a gradient normalization kernel applied across feature channels.


// A representation of a gradient normalization kernel applied across feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradientNode
type CNNCrossChannelNormalizationGradientNode struct {
	GradientFilterNode
}

// CNNCrossChannelNormalizationGradientNodeFrom constructs a [CNNCrossChannelNormalizationGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient normalization kernel applied across feature channels.
func CNNCrossChannelNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNCrossChannelNormalizationGradientNode {
	return CNNCrossChannelNormalizationGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradientnode/2948043-initwithsourcegradient
func NewCNNCrossChannelNormalizationGradientNodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelSize uint) CNNCrossChannelNormalizationGradientNode {
	instance := getCNNCrossChannelNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalizationGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelSize:"), sourceGradient, sourceImage, gradientState, kernelSize)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradientnode/2948032-nodewithsourcegradient
func (cc _CNNCrossChannelNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelSize uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelSize:"), sourceGradient, sourceImage, gradientState, kernelSize)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradientnode/2948049-kernelsize
func (c_ CNNCrossChannelNormalizationGradientNode) KernelSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradientnode/2948049-kernelsize
func (c_ CNNCrossChannelNormalizationGradientNode) SetKernelSize(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelSize:"), value)
}







