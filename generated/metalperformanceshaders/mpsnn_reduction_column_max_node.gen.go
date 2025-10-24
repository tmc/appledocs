// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ReductionColumnMaxNode] class.
var (
	ReductionColumnMaxNodeClass     _ReductionColumnMaxNodeClass
	ReductionColumnMaxNodeClassOnce sync.Once
)

func getReductionColumnMaxNodeClass() _ReductionColumnMaxNodeClass {
	ReductionColumnMaxNodeClassOnce.Do(func() {
		ReductionColumnMaxNodeClass = _ReductionColumnMaxNodeClass{objc.GetClass("MPSNNReductionColumnMaxNode")}
	})
	return ReductionColumnMaxNodeClass
}

type _ReductionColumnMaxNodeClass struct {
	class objc.Class
}





// An interface definition for the [ReductionColumnMaxNode] class.
type IReductionColumnMaxNode interface {
	IUnaryReductionNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReductionColumnMaxNodeClass) Alloc() ReductionColumnMaxNode {
	rv := objc.Send[ReductionColumnMaxNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionColumnMaxNodeClass) New() ReductionColumnMaxNode {
	rv := objc.Send[ReductionColumnMaxNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionColumnMaxNode) Init() ReductionColumnMaxNode {
	rv := objc.Send[ReductionColumnMaxNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionColumnMaxNode) Autorelease() ReductionColumnMaxNode {
	rv := objc.Send[ReductionColumnMaxNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionColumnMaxNode creates a new ReductionColumnMaxNode instance.
func NewReductionColumnMaxNode() ReductionColumnMaxNode {
	return getReductionColumnMaxNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMaxNode
type ReductionColumnMaxNode struct {
	UnaryReductionNode
}

// ReductionColumnMaxNodeFrom constructs a [ReductionColumnMaxNode] from an unsafe.Pointer.
func ReductionColumnMaxNodeFrom(ptr unsafe.Pointer) ReductionColumnMaxNode {
	return ReductionColumnMaxNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}































