// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionRowMaxNode] class.
var (
	ReductionRowMaxNodeClass     _ReductionRowMaxNodeClass
	ReductionRowMaxNodeClassOnce sync.Once
)

func getReductionRowMaxNodeClass() _ReductionRowMaxNodeClass {
	ReductionRowMaxNodeClassOnce.Do(func() {
		ReductionRowMaxNodeClass = _ReductionRowMaxNodeClass{objc.GetClass("MPSNNReductionRowMaxNode")}
	})
	return ReductionRowMaxNodeClass
}

type _ReductionRowMaxNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionRowMaxNode] class.
type IReductionRowMaxNode interface {
	IUnaryReductionNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMaxNode
type ReductionRowMaxNode struct {
	UnaryReductionNode
}

// ReductionRowMaxNodeFrom constructs a [ReductionRowMaxNode] from an unsafe.Pointer.
func ReductionRowMaxNodeFrom(ptr unsafe.Pointer) ReductionRowMaxNode {
	return ReductionRowMaxNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionRowMaxNodeClass) Alloc() ReductionRowMaxNode {
	rv := objc.Send[ReductionRowMaxNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionRowMaxNodeClass) New() ReductionRowMaxNode {
	rv := objc.Send[ReductionRowMaxNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionRowMaxNode) Init() ReductionRowMaxNode {
	rv := objc.Send[ReductionRowMaxNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionRowMaxNode) Autorelease() ReductionRowMaxNode {
	rv := objc.Send[ReductionRowMaxNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionRowMaxNode creates a new ReductionRowMaxNode instance.
func NewReductionRowMaxNode() ReductionRowMaxNode {
	return getReductionRowMaxNodeClass().New()
}
