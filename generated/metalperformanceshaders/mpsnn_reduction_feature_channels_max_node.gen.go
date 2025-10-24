// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionFeatureChannelsMaxNode] class.
var (
	ReductionFeatureChannelsMaxNodeClass     _ReductionFeatureChannelsMaxNodeClass
	ReductionFeatureChannelsMaxNodeClassOnce sync.Once
)

func getReductionFeatureChannelsMaxNodeClass() _ReductionFeatureChannelsMaxNodeClass {
	ReductionFeatureChannelsMaxNodeClassOnce.Do(func() {
		ReductionFeatureChannelsMaxNodeClass = _ReductionFeatureChannelsMaxNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsMaxNode")}
	})
	return ReductionFeatureChannelsMaxNodeClass
}

type _ReductionFeatureChannelsMaxNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionFeatureChannelsMaxNode] class.
type IReductionFeatureChannelsMaxNode interface {
	IUnaryReductionNode
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsMaxNode
type ReductionFeatureChannelsMaxNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsMaxNodeFrom constructs a [ReductionFeatureChannelsMaxNode] from an unsafe.Pointer.
func ReductionFeatureChannelsMaxNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsMaxNode {
	return ReductionFeatureChannelsMaxNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsMaxNodeClass) Alloc() ReductionFeatureChannelsMaxNode {
	rv := objc.Send[ReductionFeatureChannelsMaxNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionFeatureChannelsMaxNodeClass) New() ReductionFeatureChannelsMaxNode {
	rv := objc.Send[ReductionFeatureChannelsMaxNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsMaxNode) Init() ReductionFeatureChannelsMaxNode {
	rv := objc.Send[ReductionFeatureChannelsMaxNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsMaxNode) Autorelease() ReductionFeatureChannelsMaxNode {
	rv := objc.Send[ReductionFeatureChannelsMaxNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsMaxNode creates a new ReductionFeatureChannelsMaxNode instance.
func NewReductionFeatureChannelsMaxNode() ReductionFeatureChannelsMaxNode {
	return getReductionFeatureChannelsMaxNodeClass().New()
}




