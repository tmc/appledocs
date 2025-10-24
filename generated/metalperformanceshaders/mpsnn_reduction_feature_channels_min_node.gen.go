// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ReductionFeatureChannelsMinNode] class.
var (
	ReductionFeatureChannelsMinNodeClass     _ReductionFeatureChannelsMinNodeClass
	ReductionFeatureChannelsMinNodeClassOnce sync.Once
)

func getReductionFeatureChannelsMinNodeClass() _ReductionFeatureChannelsMinNodeClass {
	ReductionFeatureChannelsMinNodeClassOnce.Do(func() {
		ReductionFeatureChannelsMinNodeClass = _ReductionFeatureChannelsMinNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsMinNode")}
	})
	return ReductionFeatureChannelsMinNodeClass
}

type _ReductionFeatureChannelsMinNodeClass struct {
	class objc.Class
}





// An interface definition for the [ReductionFeatureChannelsMinNode] class.
type IReductionFeatureChannelsMinNode interface {
	IUnaryReductionNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsMinNodeClass) Alloc() ReductionFeatureChannelsMinNode {
	rv := objc.Send[ReductionFeatureChannelsMinNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReductionFeatureChannelsMinNodeClass) New() ReductionFeatureChannelsMinNode {
	rv := objc.Send[ReductionFeatureChannelsMinNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsMinNode) Init() ReductionFeatureChannelsMinNode {
	rv := objc.Send[ReductionFeatureChannelsMinNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsMinNode) Autorelease() ReductionFeatureChannelsMinNode {
	rv := objc.Send[ReductionFeatureChannelsMinNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsMinNode creates a new ReductionFeatureChannelsMinNode instance.
func NewReductionFeatureChannelsMinNode() ReductionFeatureChannelsMinNode {
	return getReductionFeatureChannelsMinNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsMinNode
type ReductionFeatureChannelsMinNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsMinNodeFrom constructs a [ReductionFeatureChannelsMinNode] from an unsafe.Pointer.
func ReductionFeatureChannelsMinNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsMinNode {
	return ReductionFeatureChannelsMinNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}































