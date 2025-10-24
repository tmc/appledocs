// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronReLUNode] class.
var (
	CNNNeuronReLUNodeClass     _CNNNeuronReLUNodeClass
	CNNNeuronReLUNodeClassOnce sync.Once
)

func getCNNNeuronReLUNodeClass() _CNNNeuronReLUNodeClass {
	CNNNeuronReLUNodeClassOnce.Do(func() {
		CNNNeuronReLUNodeClass = _CNNNeuronReLUNodeClass{objc.GetClass("MPSCNNNeuronReLUNode")}
	})
	return CNNNeuronReLUNodeClass
}

type _CNNNeuronReLUNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronReLUNode] class.
type ICNNNeuronReLUNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronReLUNodeClass) Alloc() CNNNeuronReLUNode {
	rv := objc.Send[CNNNeuronReLUNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronReLUNodeClass) New() CNNNeuronReLUNode {
	rv := objc.Send[CNNNeuronReLUNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronReLUNode) Init() CNNNeuronReLUNode {
	rv := objc.Send[CNNNeuronReLUNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronReLUNode) Autorelease() CNNNeuronReLUNode {
	rv := objc.Send[CNNNeuronReLUNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronReLUNode creates a new CNNNeuronReLUNode instance.
func NewCNNNeuronReLUNode() CNNNeuronReLUNode {
	return getCNNNeuronReLUNodeClass().New()
}





// A representation a ReLU neuron filter.


// A representation a ReLU neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNode
type CNNNeuronReLUNode struct {
	CNNNeuronNode
}

// CNNNeuronReLUNodeFrom constructs a [CNNNeuronReLUNode] from an unsafe.Pointer.
//
// A representation a ReLU neuron filter.
func CNNNeuronReLUNodeFrom(ptr unsafe.Pointer) CNNNeuronReLUNode {
	return CNNNeuronReLUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode/2921464-initwithsource
func NewCNNNeuronReLUNodeWithSource(sourceNode IImageNode) CNNNeuronReLUNode {
	instance := getCNNNeuronReLUNodeClass().Alloc()
	rv := objc.Send[CNNNeuronReLUNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode/2921462-initwithsource
func NewCNNNeuronReLUNodeWithSourceA(sourceNode IImageNode, a float32) CNNNeuronReLUNode {
	instance := getCNNNeuronReLUNodeClass().Alloc()
	rv := objc.Send[CNNNeuronReLUNode](instance.ID, objc.Sel("initWithSource:a:"), sourceNode, a)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode/2866494-nodewithsource
func (cc _CNNNeuronReLUNodeClass) NodeWithSourceA(sourceNode IImageNode, a float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:"), sourceNode, a)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode/2921460-nodewithsource
func (cc _CNNNeuronReLUNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















