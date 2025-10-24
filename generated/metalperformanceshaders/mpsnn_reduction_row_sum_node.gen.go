// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ReductionRowSumNode] class.
var (
	ReductionRowSumNodeClass     _ReductionRowSumNodeClass
	ReductionRowSumNodeClassOnce sync.Once
)

func getReductionRowSumNodeClass() _ReductionRowSumNodeClass {
	ReductionRowSumNodeClassOnce.Do(func() {
		ReductionRowSumNodeClass = _ReductionRowSumNodeClass{objc.GetClass("MPSNNReductionRowSumNode")}
	})
	return ReductionRowSumNodeClass
}

type _ReductionRowSumNodeClass struct {
	class objc.Class
}





// An interface definition for the [ReductionRowSumNode] class.
type IReductionRowSumNode interface {
	IUnaryReductionNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReductionRowSumNodeClass) Alloc() ReductionRowSumNode {
	rv := objc.Send[ReductionRowSumNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionRowSumNodeClass) New() ReductionRowSumNode {
	rv := objc.Send[ReductionRowSumNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionRowSumNode) Init() ReductionRowSumNode {
	rv := objc.Send[ReductionRowSumNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionRowSumNode) Autorelease() ReductionRowSumNode {
	rv := objc.Send[ReductionRowSumNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionRowSumNode creates a new ReductionRowSumNode instance.
func NewReductionRowSumNode() ReductionRowSumNode {
	return getReductionRowSumNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowSumNode
type ReductionRowSumNode struct {
	UnaryReductionNode
}

// ReductionRowSumNodeFrom constructs a [ReductionRowSumNode] from an unsafe.Pointer.
func ReductionRowSumNodeFrom(ptr unsafe.Pointer) ReductionRowSumNode {
	return ReductionRowSumNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}































