// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronHardSigmoidNode] class.
var (
	CNNNeuronHardSigmoidNodeClass     _CNNNeuronHardSigmoidNodeClass
	CNNNeuronHardSigmoidNodeClassOnce sync.Once
)

func getCNNNeuronHardSigmoidNodeClass() _CNNNeuronHardSigmoidNodeClass {
	CNNNeuronHardSigmoidNodeClassOnce.Do(func() {
		CNNNeuronHardSigmoidNodeClass = _CNNNeuronHardSigmoidNodeClass{objc.GetClass("MPSCNNNeuronHardSigmoidNode")}
	})
	return CNNNeuronHardSigmoidNodeClass
}

type _CNNNeuronHardSigmoidNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronHardSigmoidNode] class.
type ICNNNeuronHardSigmoidNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronHardSigmoidNodeClass) Alloc() CNNNeuronHardSigmoidNode {
	rv := objc.Send[CNNNeuronHardSigmoidNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronHardSigmoidNodeClass) New() CNNNeuronHardSigmoidNode {
	rv := objc.Send[CNNNeuronHardSigmoidNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronHardSigmoidNode) Init() CNNNeuronHardSigmoidNode {
	rv := objc.Send[CNNNeuronHardSigmoidNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronHardSigmoidNode) Autorelease() CNNNeuronHardSigmoidNode {
	rv := objc.Send[CNNNeuronHardSigmoidNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronHardSigmoidNode creates a new CNNNeuronHardSigmoidNode instance.
func NewCNNNeuronHardSigmoidNode() CNNNeuronHardSigmoidNode {
	return getCNNNeuronHardSigmoidNodeClass().New()
}





// A representation of a hard sigmoid neuron filter.


// A representation of a hard sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoidNode
type CNNNeuronHardSigmoidNode struct {
	CNNNeuronNode
}

// CNNNeuronHardSigmoidNodeFrom constructs a [CNNNeuronHardSigmoidNode] from an unsafe.Pointer.
//
// A representation of a hard sigmoid neuron filter.
func CNNNeuronHardSigmoidNodeFrom(ptr unsafe.Pointer) CNNNeuronHardSigmoidNode {
	return CNNNeuronHardSigmoidNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronhardsigmoidnode/2921455-initwithsource
func NewCNNNeuronHardSigmoidNodeWithSource(sourceNode IImageNode) CNNNeuronHardSigmoidNode {
	instance := getCNNNeuronHardSigmoidNodeClass().Alloc()
	rv := objc.Send[CNNNeuronHardSigmoidNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronhardsigmoidnode/2875181-initwithsource
func NewCNNNeuronHardSigmoidNodeWithSourceAB(sourceNode IImageNode, a float32, b float32) CNNNeuronHardSigmoidNode {
	instance := getCNNNeuronHardSigmoidNodeClass().Alloc()
	rv := objc.Send[CNNNeuronHardSigmoidNode](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronhardsigmoidnode/2875986-nodewithsource
func (cc _CNNNeuronHardSigmoidNodeClass) NodeWithSourceAB(sourceNode IImageNode, a float32, b float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronhardsigmoidnode/2921453-nodewithsource
func (cc _CNNNeuronHardSigmoidNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















