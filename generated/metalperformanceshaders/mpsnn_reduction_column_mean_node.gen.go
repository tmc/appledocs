// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ReductionColumnMeanNode] class.
var (
	ReductionColumnMeanNodeClass     _ReductionColumnMeanNodeClass
	ReductionColumnMeanNodeClassOnce sync.Once
)

func getReductionColumnMeanNodeClass() _ReductionColumnMeanNodeClass {
	ReductionColumnMeanNodeClassOnce.Do(func() {
		ReductionColumnMeanNodeClass = _ReductionColumnMeanNodeClass{objc.GetClass("MPSNNReductionColumnMeanNode")}
	})
	return ReductionColumnMeanNodeClass
}

type _ReductionColumnMeanNodeClass struct {
	class objc.Class
}





// An interface definition for the [ReductionColumnMeanNode] class.
type IReductionColumnMeanNode interface {
	IUnaryReductionNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReductionColumnMeanNodeClass) Alloc() ReductionColumnMeanNode {
	rv := objc.Send[ReductionColumnMeanNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionColumnMeanNodeClass) New() ReductionColumnMeanNode {
	rv := objc.Send[ReductionColumnMeanNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionColumnMeanNode) Init() ReductionColumnMeanNode {
	rv := objc.Send[ReductionColumnMeanNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionColumnMeanNode) Autorelease() ReductionColumnMeanNode {
	rv := objc.Send[ReductionColumnMeanNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionColumnMeanNode creates a new ReductionColumnMeanNode instance.
func NewReductionColumnMeanNode() ReductionColumnMeanNode {
	return getReductionColumnMeanNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMeanNode
type ReductionColumnMeanNode struct {
	UnaryReductionNode
}

// ReductionColumnMeanNodeFrom constructs a [ReductionColumnMeanNode] from an unsafe.Pointer.
func ReductionColumnMeanNodeFrom(ptr unsafe.Pointer) ReductionColumnMeanNode {
	return ReductionColumnMeanNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}































