// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNGroupNormalizationGradientNode] class.
var (
	CNNGroupNormalizationGradientNodeClass     _CNNGroupNormalizationGradientNodeClass
	CNNGroupNormalizationGradientNodeClassOnce sync.Once
)

func getCNNGroupNormalizationGradientNodeClass() _CNNGroupNormalizationGradientNodeClass {
	CNNGroupNormalizationGradientNodeClassOnce.Do(func() {
		CNNGroupNormalizationGradientNodeClass = _CNNGroupNormalizationGradientNodeClass{objc.GetClass("MPSCNNGroupNormalizationGradientNode")}
	})
	return CNNGroupNormalizationGradientNodeClass
}

type _CNNGroupNormalizationGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNGroupNormalizationGradientNode] class.
type ICNNGroupNormalizationGradientNode interface {
	IGradientFilterNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNGroupNormalizationGradientNodeClass) Alloc() CNNGroupNormalizationGradientNode {
	rv := objc.Send[CNNGroupNormalizationGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNGroupNormalizationGradientNodeClass) New() CNNGroupNormalizationGradientNode {
	rv := objc.Send[CNNGroupNormalizationGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNGroupNormalizationGradientNode) Init() CNNGroupNormalizationGradientNode {
	rv := objc.Send[CNNGroupNormalizationGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNGroupNormalizationGradientNode) Autorelease() CNNGroupNormalizationGradientNode {
	rv := objc.Send[CNNGroupNormalizationGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNGroupNormalizationGradientNode creates a new CNNGroupNormalizationGradientNode instance.
func NewCNNGroupNormalizationGradientNode() CNNGroupNormalizationGradientNode {
	return getCNNGroupNormalizationGradientNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientNode
type CNNGroupNormalizationGradientNode struct {
	GradientFilterNode
}

// CNNGroupNormalizationGradientNodeFrom constructs a [CNNGroupNormalizationGradientNode] from an unsafe.Pointer.
func CNNGroupNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNGroupNormalizationGradientNode {
	return CNNGroupNormalizationGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientnode/3152569-initwithsourcegradient
func NewCNNGroupNormalizationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) CNNGroupNormalizationGradientNode {
	instance := getCNNGroupNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[CNNGroupNormalizationGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientnode/3152570-nodewithsourcegradient
func (cc _CNNGroupNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}






















