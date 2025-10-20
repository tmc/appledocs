// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionFeatureChannelsArgumentMinNode] class.
var (
	ReductionFeatureChannelsArgumentMinNodeClass     _ReductionFeatureChannelsArgumentMinNodeClass
	ReductionFeatureChannelsArgumentMinNodeClassOnce sync.Once
)

func getReductionFeatureChannelsArgumentMinNodeClass() _ReductionFeatureChannelsArgumentMinNodeClass {
	ReductionFeatureChannelsArgumentMinNodeClassOnce.Do(func() {
		ReductionFeatureChannelsArgumentMinNodeClass = _ReductionFeatureChannelsArgumentMinNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsArgumentMinNode")}
	})
	return ReductionFeatureChannelsArgumentMinNodeClass
}

type _ReductionFeatureChannelsArgumentMinNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionFeatureChannelsArgumentMinNode] class.
type IReductionFeatureChannelsArgumentMinNode interface {
	IUnaryReductionNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsArgumentMinNode
type ReductionFeatureChannelsArgumentMinNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsArgumentMinNodeFrom constructs a [ReductionFeatureChannelsArgumentMinNode] from an unsafe.Pointer.
func ReductionFeatureChannelsArgumentMinNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsArgumentMinNode {
	return ReductionFeatureChannelsArgumentMinNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsArgumentMinNodeClass) Alloc() ReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMinNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionFeatureChannelsArgumentMinNodeClass) New() ReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMinNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsArgumentMinNode) Init() ReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMinNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsArgumentMinNode) Autorelease() ReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[ReductionFeatureChannelsArgumentMinNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsArgumentMinNode creates a new ReductionFeatureChannelsArgumentMinNode instance.
func NewReductionFeatureChannelsArgumentMinNode() ReductionFeatureChannelsArgumentMinNode {
	return getReductionFeatureChannelsArgumentMinNodeClass().New()
}
