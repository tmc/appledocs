// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNFullyConnectedGradientNode] class.
var (
	CNNFullyConnectedGradientNodeClass     _CNNFullyConnectedGradientNodeClass
	CNNFullyConnectedGradientNodeClassOnce sync.Once
)

func getCNNFullyConnectedGradientNodeClass() _CNNFullyConnectedGradientNodeClass {
	CNNFullyConnectedGradientNodeClassOnce.Do(func() {
		CNNFullyConnectedGradientNodeClass = _CNNFullyConnectedGradientNodeClass{objc.GetClass("MPSCNNFullyConnectedGradientNode")}
	})
	return CNNFullyConnectedGradientNodeClass
}

type _CNNFullyConnectedGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNFullyConnectedGradientNode] class.
type ICNNFullyConnectedGradientNode interface {
	ICNNConvolutionGradientNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNFullyConnectedGradientNodeClass) Alloc() CNNFullyConnectedGradientNode {
	rv := objc.Send[CNNFullyConnectedGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNFullyConnectedGradientNodeClass) New() CNNFullyConnectedGradientNode {
	rv := objc.Send[CNNFullyConnectedGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNFullyConnectedGradientNode) Init() CNNFullyConnectedGradientNode {
	rv := objc.Send[CNNFullyConnectedGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNFullyConnectedGradientNode) Autorelease() CNNFullyConnectedGradientNode {
	rv := objc.Send[CNNFullyConnectedGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNFullyConnectedGradientNode creates a new CNNFullyConnectedGradientNode instance.
func NewCNNFullyConnectedGradientNode() CNNFullyConnectedGradientNode {
	return getCNNFullyConnectedGradientNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradientNode
type CNNFullyConnectedGradientNode struct {
	CNNConvolutionGradientNode
}

// CNNFullyConnectedGradientNodeFrom constructs a [CNNFullyConnectedGradientNode] from an unsafe.Pointer.
func CNNFullyConnectedGradientNodeFrom(ptr unsafe.Pointer) CNNFullyConnectedGradientNode {
	return CNNFullyConnectedGradientNode{
		CNNConvolutionGradientNode: CNNConvolutionGradientNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradientnode/3152566-initwithsourcegradient
func NewCNNFullyConnectedGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IImageNode, sourceImage IImageNode, gradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) CNNFullyConnectedGradientNode {
	instance := getCNNFullyConnectedGradientNodeClass().Alloc()
	rv := objc.Send[CNNFullyConnectedGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradientnode/3152567-nodewithsourcegradient
func (cc _CNNFullyConnectedGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IImageNode, sourceImage IImageNode, gradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return rv
}






















