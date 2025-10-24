// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionColumnSumNode] class.
var (
	ReductionColumnSumNodeClass     _ReductionColumnSumNodeClass
	ReductionColumnSumNodeClassOnce sync.Once
)

func getReductionColumnSumNodeClass() _ReductionColumnSumNodeClass {
	ReductionColumnSumNodeClassOnce.Do(func() {
		ReductionColumnSumNodeClass = _ReductionColumnSumNodeClass{objc.GetClass("MPSNNReductionColumnSumNode")}
	})
	return ReductionColumnSumNodeClass
}

type _ReductionColumnSumNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionColumnSumNode] class.
type IReductionColumnSumNode interface {
	IUnaryReductionNode
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnSumNode
type ReductionColumnSumNode struct {
	UnaryReductionNode
}

// ReductionColumnSumNodeFrom constructs a [ReductionColumnSumNode] from an unsafe.Pointer.
func ReductionColumnSumNodeFrom(ptr unsafe.Pointer) ReductionColumnSumNode {
	return ReductionColumnSumNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionColumnSumNodeClass) Alloc() ReductionColumnSumNode {
	rv := objc.Send[ReductionColumnSumNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionColumnSumNodeClass) New() ReductionColumnSumNode {
	rv := objc.Send[ReductionColumnSumNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionColumnSumNode) Init() ReductionColumnSumNode {
	rv := objc.Send[ReductionColumnSumNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionColumnSumNode) Autorelease() ReductionColumnSumNode {
	rv := objc.Send[ReductionColumnSumNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionColumnSumNode creates a new ReductionColumnSumNode instance.
func NewReductionColumnSumNode() ReductionColumnSumNode {
	return getReductionColumnSumNodeClass().New()
}




