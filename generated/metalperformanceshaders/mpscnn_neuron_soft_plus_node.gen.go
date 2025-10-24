// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronSoftPlusNode] class.
var (
	CNNNeuronSoftPlusNodeClass     _CNNNeuronSoftPlusNodeClass
	CNNNeuronSoftPlusNodeClassOnce sync.Once
)

func getCNNNeuronSoftPlusNodeClass() _CNNNeuronSoftPlusNodeClass {
	CNNNeuronSoftPlusNodeClassOnce.Do(func() {
		CNNNeuronSoftPlusNodeClass = _CNNNeuronSoftPlusNodeClass{objc.GetClass("MPSCNNNeuronSoftPlusNode")}
	})
	return CNNNeuronSoftPlusNodeClass
}

type _CNNNeuronSoftPlusNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronSoftPlusNode] class.
type ICNNNeuronSoftPlusNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronSoftPlusNodeClass) Alloc() CNNNeuronSoftPlusNode {
	rv := objc.Send[CNNNeuronSoftPlusNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronSoftPlusNodeClass) New() CNNNeuronSoftPlusNode {
	rv := objc.Send[CNNNeuronSoftPlusNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronSoftPlusNode) Init() CNNNeuronSoftPlusNode {
	rv := objc.Send[CNNNeuronSoftPlusNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronSoftPlusNode) Autorelease() CNNNeuronSoftPlusNode {
	rv := objc.Send[CNNNeuronSoftPlusNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronSoftPlusNode creates a new CNNNeuronSoftPlusNode instance.
func NewCNNNeuronSoftPlusNode() CNNNeuronSoftPlusNode {
	return getCNNNeuronSoftPlusNodeClass().New()
}





// A representation of a parametric softplus neuron filter.


// A representation of a parametric softplus neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlusNode
type CNNNeuronSoftPlusNode struct {
	CNNNeuronNode
}

// CNNNeuronSoftPlusNodeFrom constructs a [CNNNeuronSoftPlusNode] from an unsafe.Pointer.
//
// A representation of a parametric softplus neuron filter.
func CNNNeuronSoftPlusNodeFrom(ptr unsafe.Pointer) CNNNeuronSoftPlusNode {
	return CNNNeuronSoftPlusNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2921457-initwithsource
func NewCNNNeuronSoftPlusNodeWithSource(sourceNode IImageNode) CNNNeuronSoftPlusNode {
	instance := getCNNNeuronSoftPlusNodeClass().Alloc()
	rv := objc.Send[CNNNeuronSoftPlusNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2866413-initwithsource
func NewCNNNeuronSoftPlusNodeWithSourceAB(sourceNode IImageNode, a float32, b float32) CNNNeuronSoftPlusNode {
	instance := getCNNNeuronSoftPlusNodeClass().Alloc()
	rv := objc.Send[CNNNeuronSoftPlusNode](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2866450-nodewithsource
func (cc _CNNNeuronSoftPlusNodeClass) NodeWithSourceAB(sourceNode IImageNode, a float32, b float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2921449-nodewithsource
func (cc _CNNNeuronSoftPlusNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















