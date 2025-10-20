// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionRowMinNode] class.
var (
	ReductionRowMinNodeClass     _ReductionRowMinNodeClass
	ReductionRowMinNodeClassOnce sync.Once
)

func getReductionRowMinNodeClass() _ReductionRowMinNodeClass {
	ReductionRowMinNodeClassOnce.Do(func() {
		ReductionRowMinNodeClass = _ReductionRowMinNodeClass{objc.GetClass("MPSNNReductionRowMinNode")}
	})
	return ReductionRowMinNodeClass
}

type _ReductionRowMinNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionRowMinNode] class.
type IReductionRowMinNode interface {
	IUnaryReductionNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMinNode
type ReductionRowMinNode struct {
	UnaryReductionNode
}

// ReductionRowMinNodeFrom constructs a [ReductionRowMinNode] from an unsafe.Pointer.
func ReductionRowMinNodeFrom(ptr unsafe.Pointer) ReductionRowMinNode {
	return ReductionRowMinNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionRowMinNodeClass) Alloc() ReductionRowMinNode {
	rv := objc.Send[ReductionRowMinNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionRowMinNodeClass) New() ReductionRowMinNode {
	rv := objc.Send[ReductionRowMinNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionRowMinNode) Init() ReductionRowMinNode {
	rv := objc.Send[ReductionRowMinNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionRowMinNode) Autorelease() ReductionRowMinNode {
	rv := objc.Send[ReductionRowMinNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionRowMinNode creates a new ReductionRowMinNode instance.
func NewReductionRowMinNode() ReductionRowMinNode {
	return getReductionRowMinNodeClass().New()
}




