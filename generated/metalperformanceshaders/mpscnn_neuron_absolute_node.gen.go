// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronAbsoluteNode] class.
var (
	CNNNeuronAbsoluteNodeClass     _CNNNeuronAbsoluteNodeClass
	CNNNeuronAbsoluteNodeClassOnce sync.Once
)

func getCNNNeuronAbsoluteNodeClass() _CNNNeuronAbsoluteNodeClass {
	CNNNeuronAbsoluteNodeClassOnce.Do(func() {
		CNNNeuronAbsoluteNodeClass = _CNNNeuronAbsoluteNodeClass{objc.GetClass("MPSCNNNeuronAbsoluteNode")}
	})
	return CNNNeuronAbsoluteNodeClass
}

type _CNNNeuronAbsoluteNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronAbsoluteNode] class.
type ICNNNeuronAbsoluteNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronAbsoluteNodeClass) Alloc() CNNNeuronAbsoluteNode {
	rv := objc.Send[CNNNeuronAbsoluteNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronAbsoluteNodeClass) New() CNNNeuronAbsoluteNode {
	rv := objc.Send[CNNNeuronAbsoluteNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronAbsoluteNode) Init() CNNNeuronAbsoluteNode {
	rv := objc.Send[CNNNeuronAbsoluteNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronAbsoluteNode) Autorelease() CNNNeuronAbsoluteNode {
	rv := objc.Send[CNNNeuronAbsoluteNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronAbsoluteNode creates a new CNNNeuronAbsoluteNode instance.
func NewCNNNeuronAbsoluteNode() CNNNeuronAbsoluteNode {
	return getCNNNeuronAbsoluteNodeClass().New()
}





// A representation of an absolute neuron filter.


// A representation of an absolute neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsoluteNode
type CNNNeuronAbsoluteNode struct {
	CNNNeuronNode
}

// CNNNeuronAbsoluteNodeFrom constructs a [CNNNeuronAbsoluteNode] from an unsafe.Pointer.
//
// A representation of an absolute neuron filter.
func CNNNeuronAbsoluteNodeFrom(ptr unsafe.Pointer) CNNNeuronAbsoluteNode {
	return CNNNeuronAbsoluteNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronabsolutenode/2921448-initwithsource
func NewCNNNeuronAbsoluteNodeWithSource(sourceNode IImageNode) CNNNeuronAbsoluteNode {
	instance := getCNNNeuronAbsoluteNodeClass().Alloc()
	rv := objc.Send[CNNNeuronAbsoluteNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronabsolutenode/2866431-nodewithsource
func (cc _CNNNeuronAbsoluteNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















