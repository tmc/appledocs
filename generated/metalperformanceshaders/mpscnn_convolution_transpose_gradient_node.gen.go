// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNConvolutionTransposeGradientNode] class.
var (
	CNNConvolutionTransposeGradientNodeClass     _CNNConvolutionTransposeGradientNodeClass
	CNNConvolutionTransposeGradientNodeClassOnce sync.Once
)

func getCNNConvolutionTransposeGradientNodeClass() _CNNConvolutionTransposeGradientNodeClass {
	CNNConvolutionTransposeGradientNodeClassOnce.Do(func() {
		CNNConvolutionTransposeGradientNodeClass = _CNNConvolutionTransposeGradientNodeClass{objc.GetClass("MPSCNNConvolutionTransposeGradientNode")}
	})
	return CNNConvolutionTransposeGradientNodeClass
}

type _CNNConvolutionTransposeGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNConvolutionTransposeGradientNode] class.
type ICNNConvolutionTransposeGradientNode interface {
	ICNNConvolutionGradientNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeGradientNodeClass) Alloc() CNNConvolutionTransposeGradientNode {
	rv := objc.Send[CNNConvolutionTransposeGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionTransposeGradientNodeClass) New() CNNConvolutionTransposeGradientNode {
	rv := objc.Send[CNNConvolutionTransposeGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionTransposeGradientNode) Init() CNNConvolutionTransposeGradientNode {
	rv := objc.Send[CNNConvolutionTransposeGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionTransposeGradientNode) Autorelease() CNNConvolutionTransposeGradientNode {
	rv := objc.Send[CNNConvolutionTransposeGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionTransposeGradientNode creates a new CNNConvolutionTransposeGradientNode instance.
func NewCNNConvolutionTransposeGradientNode() CNNConvolutionTransposeGradientNode {
	return getCNNConvolutionTransposeGradientNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientNode
type CNNConvolutionTransposeGradientNode struct {
	CNNConvolutionGradientNode
}

// CNNConvolutionTransposeGradientNodeFrom constructs a [CNNConvolutionTransposeGradientNode] from an unsafe.Pointer.
func CNNConvolutionTransposeGradientNodeFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradientNode {
	return CNNConvolutionTransposeGradientNode{
		CNNConvolutionGradientNode: CNNConvolutionGradientNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientnode/3143550-initwithsourcegradient
func NewCNNConvolutionTransposeGradientNodeWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient IImageNode, sourceImage IImageNode, gradientState ICNNConvolutionTransposeGradientStateNode, weights unsafe.Pointer) CNNConvolutionTransposeGradientNode {
	instance := getCNNConvolutionTransposeGradientNodeClass().Alloc()
	rv := objc.Send[CNNConvolutionTransposeGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionTransposeGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientnode/3143551-nodewithsourcegradient
func (cc _CNNConvolutionTransposeGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient IImageNode, sourceImage IImageNode, gradientState ICNNConvolutionTransposeGradientStateNode, weights unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:convolutionTransposeGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return rv
}






















