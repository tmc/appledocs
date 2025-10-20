// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionFeatureChannelsMeanNode] class.
var (
	ReductionFeatureChannelsMeanNodeClass     _ReductionFeatureChannelsMeanNodeClass
	ReductionFeatureChannelsMeanNodeClassOnce sync.Once
)

func getReductionFeatureChannelsMeanNodeClass() _ReductionFeatureChannelsMeanNodeClass {
	ReductionFeatureChannelsMeanNodeClassOnce.Do(func() {
		ReductionFeatureChannelsMeanNodeClass = _ReductionFeatureChannelsMeanNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsMeanNode")}
	})
	return ReductionFeatureChannelsMeanNodeClass
}

type _ReductionFeatureChannelsMeanNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionFeatureChannelsMeanNode] class.
type IReductionFeatureChannelsMeanNode interface {
	IUnaryReductionNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsMeanNode
type ReductionFeatureChannelsMeanNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsMeanNodeFrom constructs a [ReductionFeatureChannelsMeanNode] from an unsafe.Pointer.
func ReductionFeatureChannelsMeanNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsMeanNode {
	return ReductionFeatureChannelsMeanNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsMeanNodeClass) Alloc() ReductionFeatureChannelsMeanNode {
	rv := objc.Send[ReductionFeatureChannelsMeanNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionFeatureChannelsMeanNodeClass) New() ReductionFeatureChannelsMeanNode {
	rv := objc.Send[ReductionFeatureChannelsMeanNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsMeanNode) Init() ReductionFeatureChannelsMeanNode {
	rv := objc.Send[ReductionFeatureChannelsMeanNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsMeanNode) Autorelease() ReductionFeatureChannelsMeanNode {
	rv := objc.Send[ReductionFeatureChannelsMeanNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsMeanNode creates a new ReductionFeatureChannelsMeanNode instance.
func NewReductionFeatureChannelsMeanNode() ReductionFeatureChannelsMeanNode {
	return getReductionFeatureChannelsMeanNodeClass().New()
}




