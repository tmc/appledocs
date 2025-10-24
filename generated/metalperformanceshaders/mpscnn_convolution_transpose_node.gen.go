// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNConvolutionTransposeNode] class.
var (
	CNNConvolutionTransposeNodeClass     _CNNConvolutionTransposeNodeClass
	CNNConvolutionTransposeNodeClassOnce sync.Once
)

func getCNNConvolutionTransposeNodeClass() _CNNConvolutionTransposeNodeClass {
	CNNConvolutionTransposeNodeClassOnce.Do(func() {
		CNNConvolutionTransposeNodeClass = _CNNConvolutionTransposeNodeClass{objc.GetClass("MPSCNNConvolutionTransposeNode")}
	})
	return CNNConvolutionTransposeNodeClass
}

type _CNNConvolutionTransposeNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNConvolutionTransposeNode] class.
type ICNNConvolutionTransposeNode interface {
	ICNNConvolutionNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeNodeClass) Alloc() CNNConvolutionTransposeNode {
	rv := objc.Send[CNNConvolutionTransposeNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionTransposeNodeClass) New() CNNConvolutionTransposeNode {
	rv := objc.Send[CNNConvolutionTransposeNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionTransposeNode) Init() CNNConvolutionTransposeNode {
	rv := objc.Send[CNNConvolutionTransposeNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionTransposeNode) Autorelease() CNNConvolutionTransposeNode {
	rv := objc.Send[CNNConvolutionTransposeNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionTransposeNode creates a new CNNConvolutionTransposeNode instance.
func NewCNNConvolutionTransposeNode() CNNConvolutionTransposeNode {
	return getCNNConvolutionTransposeNodeClass().New()
}





// A representation of a transposed convolution.


// A representation of a transposed convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeNode
type CNNConvolutionTransposeNode struct {
	CNNConvolutionNode
}

// CNNConvolutionTransposeNodeFrom constructs a [CNNConvolutionTransposeNode] from an unsafe.Pointer.
//
// A representation of a transposed convolution.
func CNNConvolutionTransposeNodeFrom(ptr unsafe.Pointer) CNNConvolutionTransposeNode {
	return CNNConvolutionTransposeNode{
		CNNConvolutionNode: CNNConvolutionNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposenode/2942641-initwithsource
func NewCNNConvolutionTransposeNodeWithSourceConvolutionGradientStateWeights(sourceNode IImageNode, convolutionGradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) CNNConvolutionTransposeNode {
	instance := getCNNConvolutionTransposeNodeClass().Alloc()
	rv := objc.Send[CNNConvolutionTransposeNode](instance.ID, objc.Sel("initWithSource:convolutionGradientState:weights:"), sourceNode, convolutionGradientState, weights)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposenode/2942636-nodewithsource
func (cc _CNNConvolutionTransposeNodeClass) NodeWithSourceConvolutionGradientStateWeights(sourceNode IImageNode, convolutionGradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:convolutionGradientState:weights:"), sourceNode, convolutionGradientState, weights)
	return rv
}






















