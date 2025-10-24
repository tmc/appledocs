// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronTanHNode] class.
var (
	CNNNeuronTanHNodeClass     _CNNNeuronTanHNodeClass
	CNNNeuronTanHNodeClassOnce sync.Once
)

func getCNNNeuronTanHNodeClass() _CNNNeuronTanHNodeClass {
	CNNNeuronTanHNodeClassOnce.Do(func() {
		CNNNeuronTanHNodeClass = _CNNNeuronTanHNodeClass{objc.GetClass("MPSCNNNeuronTanHNode")}
	})
	return CNNNeuronTanHNodeClass
}

type _CNNNeuronTanHNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronTanHNode] class.
type ICNNNeuronTanHNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronTanHNodeClass) Alloc() CNNNeuronTanHNode {
	rv := objc.Send[CNNNeuronTanHNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronTanHNodeClass) New() CNNNeuronTanHNode {
	rv := objc.Send[CNNNeuronTanHNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronTanHNode) Init() CNNNeuronTanHNode {
	rv := objc.Send[CNNNeuronTanHNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronTanHNode) Autorelease() CNNNeuronTanHNode {
	rv := objc.Send[CNNNeuronTanHNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronTanHNode creates a new CNNNeuronTanHNode instance.
func NewCNNNeuronTanHNode() CNNNeuronTanHNode {
	return getCNNNeuronTanHNodeClass().New()
}





// A representation of a hyperbolic tangent neuron filter.


// A representation of a hyperbolic tangent neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanHNode
type CNNNeuronTanHNode struct {
	CNNNeuronNode
}

// CNNNeuronTanHNodeFrom constructs a [CNNNeuronTanHNode] from an unsafe.Pointer.
//
// A representation of a hyperbolic tangent neuron filter.
func CNNNeuronTanHNodeFrom(ptr unsafe.Pointer) CNNNeuronTanHNode {
	return CNNNeuronTanHNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2921465-initwithsource
func NewCNNNeuronTanHNodeWithSource(sourceNode IImageNode) CNNNeuronTanHNode {
	instance := getCNNNeuronTanHNodeClass().Alloc()
	rv := objc.Send[CNNNeuronTanHNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2866481-initwithsource
func NewCNNNeuronTanHNodeWithSourceAB(sourceNode IImageNode, a float32, b float32) CNNNeuronTanHNode {
	instance := getCNNNeuronTanHNodeClass().Alloc()
	rv := objc.Send[CNNNeuronTanHNode](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2866416-nodewithsource
func (cc _CNNNeuronTanHNodeClass) NodeWithSourceAB(sourceNode IImageNode, a float32, b float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2921451-nodewithsource
func (cc _CNNNeuronTanHNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















