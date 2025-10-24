// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronPowerNode] class.
var (
	CNNNeuronPowerNodeClass     _CNNNeuronPowerNodeClass
	CNNNeuronPowerNodeClassOnce sync.Once
)

func getCNNNeuronPowerNodeClass() _CNNNeuronPowerNodeClass {
	CNNNeuronPowerNodeClassOnce.Do(func() {
		CNNNeuronPowerNodeClass = _CNNNeuronPowerNodeClass{objc.GetClass("MPSCNNNeuronPowerNode")}
	})
	return CNNNeuronPowerNodeClass
}

type _CNNNeuronPowerNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronPowerNode] class.
type ICNNNeuronPowerNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronPowerNodeClass) Alloc() CNNNeuronPowerNode {
	rv := objc.Send[CNNNeuronPowerNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronPowerNodeClass) New() CNNNeuronPowerNode {
	rv := objc.Send[CNNNeuronPowerNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronPowerNode) Init() CNNNeuronPowerNode {
	rv := objc.Send[CNNNeuronPowerNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronPowerNode) Autorelease() CNNNeuronPowerNode {
	rv := objc.Send[CNNNeuronPowerNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronPowerNode creates a new CNNNeuronPowerNode instance.
func NewCNNNeuronPowerNode() CNNNeuronPowerNode {
	return getCNNNeuronPowerNodeClass().New()
}





// A representation of a power neuron filter.


// A representation of a power neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPowerNode
type CNNNeuronPowerNode struct {
	CNNNeuronNode
}

// CNNNeuronPowerNodeFrom constructs a [CNNNeuronPowerNode] from an unsafe.Pointer.
//
// A representation of a power neuron filter.
func CNNNeuronPowerNodeFrom(ptr unsafe.Pointer) CNNNeuronPowerNode {
	return CNNNeuronPowerNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronpowernode/2951937-initwithsource
func NewCNNNeuronPowerNodeWithSource(sourceNode IImageNode) CNNNeuronPowerNode {
	instance := getCNNNeuronPowerNodeClass().Alloc()
	rv := objc.Send[CNNNeuronPowerNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronpowernode/2951946-initwithsource
func NewCNNNeuronPowerNodeWithSourceABC(sourceNode IImageNode, a float32, b float32, c float32) CNNNeuronPowerNode {
	instance := getCNNNeuronPowerNodeClass().Alloc()
	rv := objc.Send[CNNNeuronPowerNode](instance.ID, objc.Sel("initWithSource:a:b:c:"), sourceNode, a, b, c)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronpowernode/2951951-nodewithsource
func (cc _CNNNeuronPowerNodeClass) NodeWithSourceABC(sourceNode IImageNode, a float32, b float32, c float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:b:c:"), sourceNode, a, b, c)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronpowernode/2951958-nodewithsource
func (cc _CNNNeuronPowerNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















