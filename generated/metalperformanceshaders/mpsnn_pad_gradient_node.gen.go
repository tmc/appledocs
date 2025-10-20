// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

//
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

// Alloc allocates a new instance without initialization.
func (pc _PadGradientNodeClass) Alloc() PadGradientNode {
	rv := objc.Send[PadGradientNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewPadGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient unsafe.Pointer, sourceImage unsafe.Pointer, gradientState unsafe.Pointer) PadGradientNode {
	instance := getPadGradientNodeClass().Alloc()
	rv := objc.Send[PadGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}



