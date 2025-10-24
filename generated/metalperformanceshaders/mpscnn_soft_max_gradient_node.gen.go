// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CNNSoftMaxGradientNode] class.
type ICNNSoftMaxGradientNode interface {
	IGradientFilterNode
	

	// properties:


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradientnode/2948039-initwithsourcegradient
func NewCNNSoftMaxGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) CNNSoftMaxGradientNode {
	instance := getCNNSoftMaxGradientNodeClass().Alloc()
	rv := objc.Send[CNNSoftMaxGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradientnode/2947995-nodewithsourcegradient
func (cc _CNNSoftMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}






















