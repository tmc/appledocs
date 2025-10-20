// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionSpatialMeanGradientNode] class.
var (
	ReductionSpatialMeanGradientNodeClass     _ReductionSpatialMeanGradientNodeClass
	ReductionSpatialMeanGradientNodeClassOnce sync.Once
)

func getReductionSpatialMeanGradientNodeClass() _ReductionSpatialMeanGradientNodeClass {
	ReductionSpatialMeanGradientNodeClassOnce.Do(func() {
		ReductionSpatialMeanGradientNodeClass = _ReductionSpatialMeanGradientNodeClass{objc.GetClass("MPSNNReductionSpatialMeanGradientNode")}
	})
	return ReductionSpatialMeanGradientNodeClass
}

type _ReductionSpatialMeanGradientNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionSpatialMeanGradientNode] class.
type IReductionSpatialMeanGradientNode interface {
	IGradientFilterNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionSpatialMeanGradientNode
type ReductionSpatialMeanGradientNode struct {
	GradientFilterNode
}

// ReductionSpatialMeanGradientNodeFrom constructs a [ReductionSpatialMeanGradientNode] from an unsafe.Pointer.
func ReductionSpatialMeanGradientNodeFrom(ptr unsafe.Pointer) ReductionSpatialMeanGradientNode {
	return ReductionSpatialMeanGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionSpatialMeanGradientNodeClass) Alloc() ReductionSpatialMeanGradientNode {
	rv := objc.Send[ReductionSpatialMeanGradientNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionSpatialMeanGradientNodeClass) New() ReductionSpatialMeanGradientNode {
	rv := objc.Send[ReductionSpatialMeanGradientNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionSpatialMeanGradientNode) Init() ReductionSpatialMeanGradientNode {
	rv := objc.Send[ReductionSpatialMeanGradientNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionSpatialMeanGradientNode) Autorelease() ReductionSpatialMeanGradientNode {
	rv := objc.Send[ReductionSpatialMeanGradientNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionSpatialMeanGradientNode creates a new ReductionSpatialMeanGradientNode instance.
func NewReductionSpatialMeanGradientNode() ReductionSpatialMeanGradientNode {
	return getReductionSpatialMeanGradientNodeClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionSpatialMeanGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (rc _ReductionSpatialMeanGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient unsafe.Pointer, sourceImage unsafe.Pointer, gradientState unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}
