// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronExponentialNode] class.
var (
	CNNNeuronExponentialNodeClass     _CNNNeuronExponentialNodeClass
	CNNNeuronExponentialNodeClassOnce sync.Once
)

func getCNNNeuronExponentialNodeClass() _CNNNeuronExponentialNodeClass {
	CNNNeuronExponentialNodeClassOnce.Do(func() {
		CNNNeuronExponentialNodeClass = _CNNNeuronExponentialNodeClass{objc.GetClass("MPSCNNNeuronExponentialNode")}
	})
	return CNNNeuronExponentialNodeClass
}

type _CNNNeuronExponentialNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronExponentialNode] class.
type ICNNNeuronExponentialNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronExponentialNodeClass) Alloc() CNNNeuronExponentialNode {
	rv := objc.Send[CNNNeuronExponentialNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronExponentialNodeClass) New() CNNNeuronExponentialNode {
	rv := objc.Send[CNNNeuronExponentialNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronExponentialNode) Init() CNNNeuronExponentialNode {
	rv := objc.Send[CNNNeuronExponentialNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronExponentialNode) Autorelease() CNNNeuronExponentialNode {
	rv := objc.Send[CNNNeuronExponentialNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronExponentialNode creates a new CNNNeuronExponentialNode instance.
func NewCNNNeuronExponentialNode() CNNNeuronExponentialNode {
	return getCNNNeuronExponentialNodeClass().New()
}





// A representation of an exponential neuron filter.


// A representation of an exponential neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponentialNode
type CNNNeuronExponentialNode struct {
	CNNNeuronNode
}

// CNNNeuronExponentialNodeFrom constructs a [CNNNeuronExponentialNode] from an unsafe.Pointer.
//
// A representation of an exponential neuron filter.
func CNNNeuronExponentialNodeFrom(ptr unsafe.Pointer) CNNNeuronExponentialNode {
	return CNNNeuronExponentialNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponentialnode/2951936-initwithsource
func NewCNNNeuronExponentialNodeWithSource(sourceNode IImageNode) CNNNeuronExponentialNode {
	instance := getCNNNeuronExponentialNodeClass().Alloc()
	rv := objc.Send[CNNNeuronExponentialNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponentialnode/2951933-initwithsource
func NewCNNNeuronExponentialNodeWithSourceABC(sourceNode IImageNode, a float32, b float32, c float32) CNNNeuronExponentialNode {
	instance := getCNNNeuronExponentialNodeClass().Alloc()
	rv := objc.Send[CNNNeuronExponentialNode](instance.ID, objc.Sel("initWithSource:a:b:c:"), sourceNode, a, b, c)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponentialnode/2951950-nodewithsource
func (cc _CNNNeuronExponentialNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponentialnode/2951959-nodewithsource
func (cc _CNNNeuronExponentialNodeClass) NodeWithSourceABC(sourceNode IImageNode, a float32, b float32, c float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:b:c:"), sourceNode, a, b, c)
	return rv
}






















