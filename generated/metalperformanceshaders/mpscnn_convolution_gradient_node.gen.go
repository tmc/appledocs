// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNConvolutionGradientNode] class.
var (
	CNNConvolutionGradientNodeClass     _CNNConvolutionGradientNodeClass
	CNNConvolutionGradientNodeClassOnce sync.Once
)

func getCNNConvolutionGradientNodeClass() _CNNConvolutionGradientNodeClass {
	CNNConvolutionGradientNodeClassOnce.Do(func() {
		CNNConvolutionGradientNodeClass = _CNNConvolutionGradientNodeClass{objc.GetClass("MPSCNNConvolutionGradientNode")}
	})
	return CNNConvolutionGradientNodeClass
}

type _CNNConvolutionGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNConvolutionGradientNode] class.
type ICNNConvolutionGradientNode interface {
	IGradientFilterNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionGradientNodeClass) Alloc() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionGradientNodeClass) New() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionGradientNode) Init() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionGradientNode) Autorelease() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionGradientNode creates a new CNNConvolutionGradientNode instance.
func NewCNNConvolutionGradientNode() CNNConvolutionGradientNode {
	return getCNNConvolutionGradientNodeClass().New()
}





// A representation of a gradient convolution kernel.


// A representation of a gradient convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientNode
type CNNConvolutionGradientNode struct {
	GradientFilterNode
}

// CNNConvolutionGradientNodeFrom constructs a [CNNConvolutionGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient convolution kernel.
func CNNConvolutionGradientNodeFrom(ptr unsafe.Pointer) CNNConvolutionGradientNode {
	return CNNConvolutionGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientnode/2947999-initwithsourcegradient
func NewCNNConvolutionGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IImageNode, sourceImage IImageNode, gradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) CNNConvolutionGradientNode {
	instance := getCNNConvolutionGradientNodeClass().Alloc()
	rv := objc.Send[CNNConvolutionGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientnode/2947984-nodewithsourcegradient
func (cc _CNNConvolutionGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IImageNode, sourceImage IImageNode, gradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return rv
}






















