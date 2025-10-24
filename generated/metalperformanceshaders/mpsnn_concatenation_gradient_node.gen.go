// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ConcatenationGradientNode] class.
var (
	ConcatenationGradientNodeClass     _ConcatenationGradientNodeClass
	ConcatenationGradientNodeClassOnce sync.Once
)

func getConcatenationGradientNodeClass() _ConcatenationGradientNodeClass {
	ConcatenationGradientNodeClassOnce.Do(func() {
		ConcatenationGradientNodeClass = _ConcatenationGradientNodeClass{objc.GetClass("MPSNNConcatenationGradientNode")}
	})
	return ConcatenationGradientNodeClass
}

type _ConcatenationGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [ConcatenationGradientNode] class.
type IConcatenationGradientNode interface {
	IGradientFilterNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _ConcatenationGradientNodeClass) Alloc() ConcatenationGradientNode {
	rv := objc.Send[ConcatenationGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ConcatenationGradientNodeClass) New() ConcatenationGradientNode {
	rv := objc.Send[ConcatenationGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConcatenationGradientNode) Init() ConcatenationGradientNode {
	rv := objc.Send[ConcatenationGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConcatenationGradientNode) Autorelease() ConcatenationGradientNode {
	rv := objc.Send[ConcatenationGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConcatenationGradientNode creates a new ConcatenationGradientNode instance.
func NewConcatenationGradientNode() ConcatenationGradientNode {
	return getConcatenationGradientNodeClass().New()
}





// A representation of the results from one or more gradient kernels.


// A representation of the results from one or more gradient kernels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationGradientNode
type ConcatenationGradientNode struct {
	GradientFilterNode
}

// ConcatenationGradientNodeFrom constructs a [ConcatenationGradientNode] from an unsafe.Pointer.
//
// A representation of the results from one or more gradient kernels.
func ConcatenationGradientNodeFrom(ptr unsafe.Pointer) ConcatenationGradientNode {
	return ConcatenationGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationgradientnode/2951934-initwithsourcegradient
func NewConcatenationGradientNodeWithSourceGradientSourceImageGradientState(gradientSourceNode IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) ConcatenationGradientNode {
	instance := getConcatenationGradientNodeClass().Alloc()
	rv := objc.Send[ConcatenationGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradientSourceNode, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationgradientnode/2951948-nodewithsourcegradient
func (cc _ConcatenationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(gradientSourceNode IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), gradientSourceNode, sourceImage, gradientState)
	return rv
}






















