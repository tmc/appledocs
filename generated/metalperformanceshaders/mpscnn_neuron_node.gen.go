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
	A() float32
	SetA(value float32)
	B() float32
	SetB(value float32)
	C() float32
	SetC(value float32)
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


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/a
func (c_ CNNNeuronNode) A() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("a"))
	return rv
}


// SetA sets the value of the a property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/a
func (c_ CNNNeuronNode) SetA(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setA:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/b
func (c_ CNNNeuronNode) B() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("b"))
	return rv
}


// SetB sets the value of the b property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/b
func (c_ CNNNeuronNode) SetB(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setB:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/c
func (c_ CNNNeuronNode) C() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("c"))
	return rv
}


// SetC sets the value of the c property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/c
func (c_ CNNNeuronNode) SetC(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setC:"), value)
}



