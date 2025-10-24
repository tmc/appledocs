// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CNNLogSoftMaxGradientNode] class.
type ICNNLogSoftMaxGradientNode interface {
	IGradientFilterNode
	

	// properties:


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradientnode/2947971-initwithsourcegradient
func NewCNNLogSoftMaxGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) CNNLogSoftMaxGradientNode {
	instance := getCNNLogSoftMaxGradientNodeClass().Alloc()
	rv := objc.Send[CNNLogSoftMaxGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradientnode/2947974-nodewithsourcegradient
func (cc _CNNLogSoftMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}






















