// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReductionFeatureChannelsSumNode] class.
var (
	ReductionFeatureChannelsSumNodeClass     _ReductionFeatureChannelsSumNodeClass
	ReductionFeatureChannelsSumNodeClassOnce sync.Once
)

func getReductionFeatureChannelsSumNodeClass() _ReductionFeatureChannelsSumNodeClass {
	ReductionFeatureChannelsSumNodeClassOnce.Do(func() {
		ReductionFeatureChannelsSumNodeClass = _ReductionFeatureChannelsSumNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsSumNode")}
	})
	return ReductionFeatureChannelsSumNodeClass
}

type _ReductionFeatureChannelsSumNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReductionFeatureChannelsSumNode] class.
type IReductionFeatureChannelsSumNode interface {
	IUnaryReductionNode
	// properties:
	Weight() float32
	SetWeight(value float32)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsSumNode
type ReductionFeatureChannelsSumNode struct {
	UnaryReductionNode
}

// ReductionFeatureChannelsSumNodeFrom constructs a [ReductionFeatureChannelsSumNode] from an unsafe.Pointer.
func ReductionFeatureChannelsSumNodeFrom(ptr unsafe.Pointer) ReductionFeatureChannelsSumNode {
	return ReductionFeatureChannelsSumNode{
		UnaryReductionNode: UnaryReductionNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReductionFeatureChannelsSumNodeClass) Alloc() ReductionFeatureChannelsSumNode {
	rv := objc.Send[ReductionFeatureChannelsSumNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReductionFeatureChannelsSumNodeClass) New() ReductionFeatureChannelsSumNode {
	rv := objc.Send[ReductionFeatureChannelsSumNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReductionFeatureChannelsSumNode) Init() ReductionFeatureChannelsSumNode {
	rv := objc.Send[ReductionFeatureChannelsSumNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReductionFeatureChannelsSumNode) Autorelease() ReductionFeatureChannelsSumNode {
	rv := objc.Send[ReductionFeatureChannelsSumNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReductionFeatureChannelsSumNode creates a new ReductionFeatureChannelsSumNode instance.
func NewReductionFeatureChannelsSumNode() ReductionFeatureChannelsSumNode {
	return getReductionFeatureChannelsSumNodeClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsSumNode/weight
func (r_ ReductionFeatureChannelsSumNode) Weight() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("weight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsSumNode/weight
func (r_ ReductionFeatureChannelsSumNode) SetWeight(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setWeight:"), value)
}



