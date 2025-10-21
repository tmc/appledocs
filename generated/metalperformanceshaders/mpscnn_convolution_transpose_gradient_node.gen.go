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
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientNode
type CNNConvolutionTransposeGradientNode struct {
	objectivec.Object
}

// CNNConvolutionTransposeGradientNodeFrom constructs a [CNNConvolutionTransposeGradientNode] from an unsafe.Pointer.
func CNNConvolutionTransposeGradientNodeFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradientNode {
	return CNNConvolutionTransposeGradientNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeGradientNodeClass) Alloc() CNNConvolutionTransposeGradientNode {
	rv := objc.Send[CNNConvolutionTransposeGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientNode/nodeWithSourceGradient:sourceImage:convolutionTransposeGradientState:weights:
func (cc _CNNConvolutionTransposeGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionTransposeGradientStateNode, weights objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:convolutionTransposeGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return rv
}



