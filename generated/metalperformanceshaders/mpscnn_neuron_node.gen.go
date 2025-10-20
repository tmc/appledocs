// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNNNeuronNode] class.
var (
	CNNNeuronNodeClass     _CNNNeuronNodeClass
	CNNNeuronNodeClassOnce sync.Once
)

func getCNNNeuronNodeClass() _CNNNeuronNodeClass {
	CNNNeuronNodeClassOnce.Do(func() {
		CNNNeuronNodeClass = _CNNNeuronNodeClass{objc.GetClass("MPSCNNNeuronNode")}
	})
	return CNNNeuronNodeClass
}

type _CNNNeuronNodeClass struct {
	class objc.Class
}

// An interface definition for the [CNNNeuronNode] class.
type ICNNNeuronNode interface {
	IFilterNode
}

// The virtual base class for MPS CNN neuron nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode
type CNNNeuronNode struct {
	FilterNode
}

// CNNNeuronNodeFrom constructs a [CNNNeuronNode] from an unsafe.Pointer.
//
// The virtual base class for MPS CNN neuron nodes.
func CNNNeuronNodeFrom(ptr unsafe.Pointer) CNNNeuronNode {
	return CNNNeuronNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronNodeClass) Alloc() CNNNeuronNode {
	rv := objc.Send[CNNNeuronNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNNeuronNodeClass) New() CNNNeuronNode {
	rv := objc.Send[CNNNeuronNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronNode) Init() CNNNeuronNode {
	rv := objc.Send[CNNNeuronNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronNode) Autorelease() CNNNeuronNode {
	rv := objc.Send[CNNNeuronNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronNode creates a new CNNNeuronNode instance.
func NewCNNNeuronNode() CNNNeuronNode {
	return getCNNNeuronNodeClass().New()
}
