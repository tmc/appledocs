// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PadGradientNode] class.
var (
	PadGradientNodeClass     _PadGradientNodeClass
	PadGradientNodeClassOnce sync.Once
)

func getPadGradientNodeClass() _PadGradientNodeClass {
	PadGradientNodeClassOnce.Do(func() {
		PadGradientNodeClass = _PadGradientNodeClass{objc.GetClass("MPSNNPadGradientNode")}
	})
	return PadGradientNodeClass
}

type _PadGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [PadGradientNode] class.
type IPadGradientNode interface {
	IGradientFilterNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PadGradientNodeClass) Alloc() PadGradientNode {
	rv := objc.Send[PadGradientNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PadGradientNodeClass) New() PadGradientNode {
	rv := objc.Send[PadGradientNode](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PadGradientNode) Init() PadGradientNode {
	rv := objc.Send[PadGradientNode](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PadGradientNode) Autorelease() PadGradientNode {
	rv := objc.Send[PadGradientNode](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPadGradientNode creates a new PadGradientNode instance.
func NewPadGradientNode() PadGradientNode {
	return getPadGradientNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradientNode
type PadGradientNode struct {
	GradientFilterNode
}

// PadGradientNodeFrom constructs a [PadGradientNode] from an unsafe.Pointer.
func PadGradientNodeFrom(ptr unsafe.Pointer) PadGradientNode {
	return PadGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadgradientnode/3037391-initwithsourcegradient
func NewPadGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) PadGradientNode {
	instance := getPadGradientNodeClass().Alloc()
	rv := objc.Send[PadGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadgradientnode/3037392-nodewithsourcegradient
func (pc _PadGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}






















