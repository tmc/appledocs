// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronReLUNNode] class.
var (
	CNNNeuronReLUNNodeClass     _CNNNeuronReLUNNodeClass
	CNNNeuronReLUNNodeClassOnce sync.Once
)

func getCNNNeuronReLUNNodeClass() _CNNNeuronReLUNNodeClass {
	CNNNeuronReLUNNodeClassOnce.Do(func() {
		CNNNeuronReLUNNodeClass = _CNNNeuronReLUNNodeClass{objc.GetClass("MPSCNNNeuronReLUNNode")}
	})
	return CNNNeuronReLUNNodeClass
}

type _CNNNeuronReLUNNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronReLUNNode] class.
type ICNNNeuronReLUNNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronReLUNNodeClass) Alloc() CNNNeuronReLUNNode {
	rv := objc.Send[CNNNeuronReLUNNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronReLUNNodeClass) New() CNNNeuronReLUNNode {
	rv := objc.Send[CNNNeuronReLUNNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronReLUNNode) Init() CNNNeuronReLUNNode {
	rv := objc.Send[CNNNeuronReLUNNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronReLUNNode) Autorelease() CNNNeuronReLUNNode {
	rv := objc.Send[CNNNeuronReLUNNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronReLUNNode creates a new CNNNeuronReLUNNode instance.
func NewCNNNeuronReLUNNode() CNNNeuronReLUNNode {
	return getCNNNeuronReLUNNodeClass().New()
}





// A representation a ReLUN neuron filter.


// A representation a ReLUN neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNNode
type CNNNeuronReLUNNode struct {
	CNNNeuronNode
}

// CNNNeuronReLUNNodeFrom constructs a [CNNNeuronReLUNNode] from an unsafe.Pointer.
//
// A representation a ReLUN neuron filter.
func CNNNeuronReLUNNodeFrom(ptr unsafe.Pointer) CNNNeuronReLUNNode {
	return CNNNeuronReLUNNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode/2921593-initwithsource
func NewCNNNeuronReLUNNodeWithSource(sourceNode IImageNode) CNNNeuronReLUNNode {
	instance := getCNNNeuronReLUNNodeClass().Alloc()
	rv := objc.Send[CNNNeuronReLUNNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode/2921596-initwithsource
func NewCNNNeuronReLUNNodeWithSourceAB(sourceNode IImageNode, a float32, b float32) CNNNeuronReLUNNode {
	instance := getCNNNeuronReLUNNodeClass().Alloc()
	rv := objc.Send[CNNNeuronReLUNNode](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode/2921590-nodewithsource
func (cc _CNNNeuronReLUNNodeClass) NodeWithSourceAB(sourceNode IImageNode, a float32, b float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode/2921594-nodewithsource
func (cc _CNNNeuronReLUNNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















