// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronPReLUNode] class.
var (
	CNNNeuronPReLUNodeClass     _CNNNeuronPReLUNodeClass
	CNNNeuronPReLUNodeClassOnce sync.Once
)

func getCNNNeuronPReLUNodeClass() _CNNNeuronPReLUNodeClass {
	CNNNeuronPReLUNodeClassOnce.Do(func() {
		CNNNeuronPReLUNodeClass = _CNNNeuronPReLUNodeClass{objc.GetClass("MPSCNNNeuronPReLUNode")}
	})
	return CNNNeuronPReLUNodeClass
}

type _CNNNeuronPReLUNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronPReLUNode] class.
type ICNNNeuronPReLUNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronPReLUNodeClass) Alloc() CNNNeuronPReLUNode {
	rv := objc.Send[CNNNeuronPReLUNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronPReLUNodeClass) New() CNNNeuronPReLUNode {
	rv := objc.Send[CNNNeuronPReLUNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronPReLUNode) Init() CNNNeuronPReLUNode {
	rv := objc.Send[CNNNeuronPReLUNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronPReLUNode) Autorelease() CNNNeuronPReLUNode {
	rv := objc.Send[CNNNeuronPReLUNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronPReLUNode creates a new CNNNeuronPReLUNode instance.
func NewCNNNeuronPReLUNode() CNNNeuronPReLUNode {
	return getCNNNeuronPReLUNodeClass().New()
}





// A representation a PReLU neuron filter.


// A representation a PReLU neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPReLUNode
type CNNNeuronPReLUNode struct {
	CNNNeuronNode
}

// CNNNeuronPReLUNodeFrom constructs a [CNNNeuronPReLUNode] from an unsafe.Pointer.
//
// A representation a PReLU neuron filter.
func CNNNeuronPReLUNodeFrom(ptr unsafe.Pointer) CNNNeuronPReLUNode {
	return CNNNeuronPReLUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronprelunode/2921595-initwithsource
func NewCNNNeuronPReLUNodeWithSourceAData(sourceNode IImageNode, aData foundation.Data) CNNNeuronPReLUNode {
	instance := getCNNNeuronPReLUNodeClass().Alloc()
	rv := objc.Send[CNNNeuronPReLUNode](instance.ID, objc.Sel("initWithSource:aData:"), sourceNode, aData)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronprelunode/2921597-nodewithsource
func (cc _CNNNeuronPReLUNodeClass) NodeWithSourceAData(sourceNode IImageNode, aData foundation.Data) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:aData:"), sourceNode, aData)
	return rv
}






















