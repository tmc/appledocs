// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronELUNode] class.
var (
	CNNNeuronELUNodeClass     _CNNNeuronELUNodeClass
	CNNNeuronELUNodeClassOnce sync.Once
)

func getCNNNeuronELUNodeClass() _CNNNeuronELUNodeClass {
	CNNNeuronELUNodeClassOnce.Do(func() {
		CNNNeuronELUNodeClass = _CNNNeuronELUNodeClass{objc.GetClass("MPSCNNNeuronELUNode")}
	})
	return CNNNeuronELUNodeClass
}

type _CNNNeuronELUNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronELUNode] class.
type ICNNNeuronELUNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronELUNodeClass) Alloc() CNNNeuronELUNode {
	rv := objc.Send[CNNNeuronELUNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronELUNodeClass) New() CNNNeuronELUNode {
	rv := objc.Send[CNNNeuronELUNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronELUNode) Init() CNNNeuronELUNode {
	rv := objc.Send[CNNNeuronELUNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronELUNode) Autorelease() CNNNeuronELUNode {
	rv := objc.Send[CNNNeuronELUNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronELUNode creates a new CNNNeuronELUNode instance.
func NewCNNNeuronELUNode() CNNNeuronELUNode {
	return getCNNNeuronELUNodeClass().New()
}





// A representation of a parametric ELU neuron filter.


// A representation of a parametric ELU neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELUNode
type CNNNeuronELUNode struct {
	CNNNeuronNode
}

// CNNNeuronELUNodeFrom constructs a [CNNNeuronELUNode] from an unsafe.Pointer.
//
// A representation of a parametric ELU neuron filter.
func CNNNeuronELUNodeFrom(ptr unsafe.Pointer) CNNNeuronELUNode {
	return CNNNeuronELUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronelunode/2921447-initwithsource
func NewCNNNeuronELUNodeWithSource(sourceNode IImageNode) CNNNeuronELUNode {
	instance := getCNNNeuronELUNodeClass().Alloc()
	rv := objc.Send[CNNNeuronELUNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronelunode/2921454-initwithsource
func NewCNNNeuronELUNodeWithSourceA(sourceNode IImageNode, a float32) CNNNeuronELUNode {
	instance := getCNNNeuronELUNodeClass().Alloc()
	rv := objc.Send[CNNNeuronELUNode](instance.ID, objc.Sel("initWithSource:a:"), sourceNode, a)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronelunode/2866463-nodewithsource
func (cc _CNNNeuronELUNodeClass) NodeWithSourceA(sourceNode IImageNode, a float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:"), sourceNode, a)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronelunode/2921452-nodewithsource
func (cc _CNNNeuronELUNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















