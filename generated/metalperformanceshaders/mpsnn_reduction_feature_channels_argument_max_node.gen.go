// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionFeatureChannelsArgumentMaxNode] class.
var (
	ReductionFeatureChannelsArgumentMaxNodeClass     _ReductionFeatureChannelsArgumentMaxNodeClass
	ReductionFeatureChannelsArgumentMaxNodeClassOnce sync.Once
)

func getReductionFeatureChannelsArgumentMaxNodeClass() _ReductionFeatureChannelsArgumentMaxNodeClass {
	ReductionFeatureChannelsArgumentMaxNodeClassOnce.Do(func() {
		ReductionFeatureChannelsArgumentMaxNodeClass = _ReductionFeatureChannelsArgumentMaxNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsArgumentMaxNode")}
	})
	return ReductionFeatureChannelsArgumentMaxNodeClass
}

type _ReductionFeatureChannelsArgumentMaxNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionFeatureChannelsArgumentMaxNode] class.
type IReductionFeatureChannelsArgumentMaxNode interface {
	IUnaryReductionNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsArgumentMaxNode
type ReductionFeatureChannelsArgumentMaxNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsArgumentMaxNodeFrom constructs a [ReductionFeatureChannelsArgumentMaxNode] from an unsafe.Pointer.
func ReductionFeatureChannelsArgumentMaxNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsArgumentMaxNode {
	return ReductionFeatureChannelsArgumentMaxNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsArgumentMaxNodeClass) Alloc() ReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMaxNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionFeatureChannelsArgumentMaxNodeClass) New() ReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMaxNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsArgumentMaxNode) Init() ReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMaxNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsArgumentMaxNode) Autorelease() ReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMaxNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsArgumentMaxNode creates a new ReductionFeatureChannelsArgumentMaxNode instance.
func NewReductionFeatureChannelsArgumentMaxNode() ReductionFeatureChannelsArgumentMaxNode {
	return getReductionFeatureChannelsArgumentMaxNodeClass().New()
}
