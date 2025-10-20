// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LossGradientNode] class.
var (
	LossGradientNodeClass     _LossGradientNodeClass
	LossGradientNodeClassOnce sync.Once
)

func getLossGradientNodeClass() _LossGradientNodeClass {
	LossGradientNodeClassOnce.Do(func() {
		LossGradientNodeClass = _LossGradientNodeClass{objc.GetClass("MPSNNLossGradientNode")}
	})
	return LossGradientNodeClass
}

type _LossGradientNodeClass struct {
	class objc.Class
}

// An interface definition for the [LossGradientNode] class.
type ILossGradientNode interface {
	IGradientFilterNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode
type LossGradientNode struct {
	GradientFilterNode
}

// LossGradientNodeFrom constructs a [LossGradientNode] from an unsafe.Pointer.
func LossGradientNodeFrom(ptr unsafe.Pointer) LossGradientNode {
	return LossGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (lc _LossGradientNodeClass) Alloc() LossGradientNode {
	rv := objc.Send[LossGradientNode](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LossGradientNodeClass) New() LossGradientNode {
	rv := objc.Send[LossGradientNode](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LossGradientNode) Init() LossGradientNode {
	rv := objc.Send[LossGradientNode](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LossGradientNode) Autorelease() LossGradientNode {
	rv := objc.Send[LossGradientNode](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLossGradientNode creates a new LossGradientNode instance.
func NewLossGradientNode() LossGradientNode {
	return getLossGradientNodeClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:
func (lc _LossGradientNodeClass) NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes unsafe.Pointer, gradientState unsafe.Pointer, descriptor unsafe.Pointer, isLabelsGradientFilter bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceNodes, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/epsilon
func (l_ LossGradientNode) Epsilon() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("epsilon"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/reductionType
func (l_ LossGradientNode) ReductionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("reductionType"))
	return rv
}
