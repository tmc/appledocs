// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionColumnMinNode] class.
var (
	ReductionColumnMinNodeClass     _ReductionColumnMinNodeClass
	ReductionColumnMinNodeClassOnce sync.Once
)

func getReductionColumnMinNodeClass() _ReductionColumnMinNodeClass {
	ReductionColumnMinNodeClassOnce.Do(func() {
		ReductionColumnMinNodeClass = _ReductionColumnMinNodeClass{objc.GetClass("MPSNNReductionColumnMinNode")}
	})
	return ReductionColumnMinNodeClass
}

type _ReductionColumnMinNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionColumnMinNode] class.
type IReductionColumnMinNode interface {
	IUnaryReductionNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMinNode
type ReductionColumnMinNode struct {
	UnaryReductionNode
}

// ReductionColumnMinNodeFrom constructs a [ReductionColumnMinNode] from an unsafe.Pointer.
func ReductionColumnMinNodeFrom(ptr unsafe.Pointer) ReductionColumnMinNode {
	return ReductionColumnMinNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionColumnMinNodeClass) Alloc() ReductionColumnMinNode {
	rv := objc.Send[ReductionColumnMinNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionColumnMinNodeClass) New() ReductionColumnMinNode {
	rv := objc.Send[ReductionColumnMinNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionColumnMinNode) Init() ReductionColumnMinNode {
	rv := objc.Send[ReductionColumnMinNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionColumnMinNode) Autorelease() ReductionColumnMinNode {
	rv := objc.Send[ReductionColumnMinNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionColumnMinNode creates a new ReductionColumnMinNode instance.
func NewReductionColumnMinNode() ReductionColumnMinNode {
	return getReductionColumnMinNodeClass().New()
}
