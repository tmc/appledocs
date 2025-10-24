// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNConvolutionTransposeGradientStateNode] class.
var (
	CNNConvolutionTransposeGradientStateNodeClass     _CNNConvolutionTransposeGradientStateNodeClass
	CNNConvolutionTransposeGradientStateNodeClassOnce sync.Once
)

func getCNNConvolutionTransposeGradientStateNodeClass() _CNNConvolutionTransposeGradientStateNodeClass {
	CNNConvolutionTransposeGradientStateNodeClassOnce.Do(func() {
		CNNConvolutionTransposeGradientStateNodeClass = _CNNConvolutionTransposeGradientStateNodeClass{objc.GetClass("MPSCNNConvolutionTransposeGradientStateNode")}
	})
	return CNNConvolutionTransposeGradientStateNodeClass
}

type _CNNConvolutionTransposeGradientStateNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNConvolutionTransposeGradientStateNode] class.
type ICNNConvolutionTransposeGradientStateNode interface {
	ICNNConvolutionGradientStateNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeGradientStateNodeClass) Alloc() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[CNNConvolutionTransposeGradientStateNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionTransposeGradientStateNodeClass) New() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[CNNConvolutionTransposeGradientStateNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionTransposeGradientStateNode) Init() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[CNNConvolutionTransposeGradientStateNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionTransposeGradientStateNode) Autorelease() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[CNNConvolutionTransposeGradientStateNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionTransposeGradientStateNode creates a new CNNConvolutionTransposeGradientStateNode instance.
func NewCNNConvolutionTransposeGradientStateNode() CNNConvolutionTransposeGradientStateNode {
	return getCNNConvolutionTransposeGradientStateNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientStateNode
type CNNConvolutionTransposeGradientStateNode struct {
	CNNConvolutionGradientStateNode
}

// CNNConvolutionTransposeGradientStateNodeFrom constructs a [CNNConvolutionTransposeGradientStateNode] from an unsafe.Pointer.
func CNNConvolutionTransposeGradientStateNodeFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradientStateNode {
	return CNNConvolutionTransposeGradientStateNode{
		CNNConvolutionGradientStateNode: CNNConvolutionGradientStateNodeFrom(ptr),
	}
}































