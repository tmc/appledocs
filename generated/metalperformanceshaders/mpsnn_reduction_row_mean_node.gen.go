// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionRowMeanNode] class.
var (
	ReductionRowMeanNodeClass     _ReductionRowMeanNodeClass
	ReductionRowMeanNodeClassOnce sync.Once
)

func getReductionRowMeanNodeClass() _ReductionRowMeanNodeClass {
	ReductionRowMeanNodeClassOnce.Do(func() {
		ReductionRowMeanNodeClass = _ReductionRowMeanNodeClass{objc.GetClass("MPSNNReductionRowMeanNode")}
	})
	return ReductionRowMeanNodeClass
}

type _ReductionRowMeanNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionRowMeanNode] class.
type IReductionRowMeanNode interface {
	IUnaryReductionNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMeanNode
type ReductionRowMeanNode struct {
	UnaryReductionNode
}

// ReductionRowMeanNodeFrom constructs a [ReductionRowMeanNode] from an unsafe.Pointer.
func ReductionRowMeanNodeFrom(ptr unsafe.Pointer) ReductionRowMeanNode {
	return ReductionRowMeanNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionRowMeanNodeClass) Alloc() ReductionRowMeanNode {
	rv := objc.Send[ReductionRowMeanNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionRowMeanNodeClass) New() ReductionRowMeanNode {
	rv := objc.Send[ReductionRowMeanNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionRowMeanNode) Init() ReductionRowMeanNode {
	rv := objc.Send[ReductionRowMeanNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionRowMeanNode) Autorelease() ReductionRowMeanNode {
	rv := objc.Send[ReductionRowMeanNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionRowMeanNode creates a new ReductionRowMeanNode instance.
func NewReductionRowMeanNode() ReductionRowMeanNode {
	return getReductionRowMeanNodeClass().New()
}
