// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionSpatialMeanNode] class.
var (
	ReductionSpatialMeanNodeClass     _ReductionSpatialMeanNodeClass
	ReductionSpatialMeanNodeClassOnce sync.Once
)

func getReductionSpatialMeanNodeClass() _ReductionSpatialMeanNodeClass {
	ReductionSpatialMeanNodeClassOnce.Do(func() {
		ReductionSpatialMeanNodeClass = _ReductionSpatialMeanNodeClass{objc.GetClass("MPSNNReductionSpatialMeanNode")}
	})
	return ReductionSpatialMeanNodeClass
}

type _ReductionSpatialMeanNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionSpatialMeanNode] class.
type IReductionSpatialMeanNode interface {
	IUnaryReductionNode
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionSpatialMeanNode
type ReductionSpatialMeanNode struct {
	UnaryReductionNode
}

// ReductionSpatialMeanNodeFrom constructs a [ReductionSpatialMeanNode] from an unsafe.Pointer.
func ReductionSpatialMeanNodeFrom(ptr unsafe.Pointer) ReductionSpatialMeanNode {
	return ReductionSpatialMeanNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionSpatialMeanNodeClass) Alloc() ReductionSpatialMeanNode {
	rv := objc.Send[ReductionSpatialMeanNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionSpatialMeanNodeClass) New() ReductionSpatialMeanNode {
	rv := objc.Send[ReductionSpatialMeanNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionSpatialMeanNode) Init() ReductionSpatialMeanNode {
	rv := objc.Send[ReductionSpatialMeanNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionSpatialMeanNode) Autorelease() ReductionSpatialMeanNode {
	rv := objc.Send[ReductionSpatialMeanNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionSpatialMeanNode creates a new ReductionSpatialMeanNode instance.
func NewReductionSpatialMeanNode() ReductionSpatialMeanNode {
	return getReductionSpatialMeanNodeClass().New()
}




