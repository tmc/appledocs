// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReshapeGradientNode] class.
var (
	ReshapeGradientNodeClass     _ReshapeGradientNodeClass
	ReshapeGradientNodeClassOnce sync.Once
)

func getReshapeGradientNodeClass() _ReshapeGradientNodeClass {
	ReshapeGradientNodeClassOnce.Do(func() {
		ReshapeGradientNodeClass = _ReshapeGradientNodeClass{objc.GetClass("MPSNNReshapeGradientNode")}
	})
	return ReshapeGradientNodeClass
}

type _ReshapeGradientNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReshapeGradientNode] class.
type IReshapeGradientNode interface {
	IGradientFilterNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradientNode
type ReshapeGradientNode struct {
	GradientFilterNode
}

// ReshapeGradientNodeFrom constructs a [ReshapeGradientNode] from an unsafe.Pointer.
func ReshapeGradientNodeFrom(ptr unsafe.Pointer) ReshapeGradientNode {
	return ReshapeGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReshapeGradientNodeClass) Alloc() ReshapeGradientNode {
	rv := objc.Send[ReshapeGradientNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReshapeGradientNodeClass) New() ReshapeGradientNode {
	rv := objc.Send[ReshapeGradientNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReshapeGradientNode) Init() ReshapeGradientNode {
	rv := objc.Send[ReshapeGradientNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReshapeGradientNode) Autorelease() ReshapeGradientNode {
	rv := objc.Send[ReshapeGradientNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReshapeGradientNode creates a new ReshapeGradientNode instance.
func NewReshapeGradientNode() ReshapeGradientNode {
	return getReshapeGradientNodeClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewReshapeGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient unsafe.Pointer, sourceImage unsafe.Pointer, gradientState unsafe.Pointer) ReshapeGradientNode {
	instance := getReshapeGradientNodeClass().Alloc()
	rv := objc.Send[ReshapeGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}
