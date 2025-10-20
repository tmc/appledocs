// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNNNeuronGeLUNode] class.
var (
	CNNNeuronGeLUNodeClass     _CNNNeuronGeLUNodeClass
	CNNNeuronGeLUNodeClassOnce sync.Once
)

func getCNNNeuronGeLUNodeClass() _CNNNeuronGeLUNodeClass {
	CNNNeuronGeLUNodeClassOnce.Do(func() {
		CNNNeuronGeLUNodeClass = _CNNNeuronGeLUNodeClass{objc.GetClass("MPSCNNNeuronGeLUNode")}
	})
	return CNNNeuronGeLUNodeClass
}

type _CNNNeuronGeLUNodeClass struct {
	class objc.Class
}

// An interface definition for the [CNNNeuronGeLUNode] class.
type ICNNNeuronGeLUNode interface {
	ICNNNeuronNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGeLUNode
type CNNNeuronGeLUNode struct {
	CNNNeuronNode
}

// CNNNeuronGeLUNodeFrom constructs a [CNNNeuronGeLUNode] from an unsafe.Pointer.
func CNNNeuronGeLUNodeFrom(ptr unsafe.Pointer) CNNNeuronGeLUNode {
	return CNNNeuronGeLUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronGeLUNodeClass) Alloc() CNNNeuronGeLUNode {
	rv := objc.Send[CNNNeuronGeLUNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNNeuronGeLUNodeClass) New() CNNNeuronGeLUNode {
	rv := objc.Send[CNNNeuronGeLUNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronGeLUNode) Init() CNNNeuronGeLUNode {
	rv := objc.Send[CNNNeuronGeLUNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronGeLUNode) Autorelease() CNNNeuronGeLUNode {
	rv := objc.Send[CNNNeuronGeLUNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronGeLUNode creates a new CNNNeuronGeLUNode instance.
func NewCNNNeuronGeLUNode() CNNNeuronGeLUNode {
	return getCNNNeuronGeLUNodeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGeLUNode/init(source:)
func NewCNNNeuronGeLUNodeWithSource(sourceNode unsafe.Pointer) CNNNeuronGeLUNode {
	instance := getCNNNeuronGeLUNodeClass().Alloc()
	rv := objc.Send[CNNNeuronGeLUNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGeLUNode/nodeWithSource:
func (cc _CNNNeuronGeLUNodeClass) NodeWithSource(sourceNode unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}


