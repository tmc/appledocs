// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronSigmoidNode] class.
var (
	CNNNeuronSigmoidNodeClass     _CNNNeuronSigmoidNodeClass
	CNNNeuronSigmoidNodeClassOnce sync.Once
)

func getCNNNeuronSigmoidNodeClass() _CNNNeuronSigmoidNodeClass {
	CNNNeuronSigmoidNodeClassOnce.Do(func() {
		CNNNeuronSigmoidNodeClass = _CNNNeuronSigmoidNodeClass{objc.GetClass("MPSCNNNeuronSigmoidNode")}
	})
	return CNNNeuronSigmoidNodeClass
}

type _CNNNeuronSigmoidNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronSigmoidNode] class.
type ICNNNeuronSigmoidNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronSigmoidNodeClass) Alloc() CNNNeuronSigmoidNode {
	rv := objc.Send[CNNNeuronSigmoidNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronSigmoidNodeClass) New() CNNNeuronSigmoidNode {
	rv := objc.Send[CNNNeuronSigmoidNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronSigmoidNode) Init() CNNNeuronSigmoidNode {
	rv := objc.Send[CNNNeuronSigmoidNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronSigmoidNode) Autorelease() CNNNeuronSigmoidNode {
	rv := objc.Send[CNNNeuronSigmoidNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronSigmoidNode creates a new CNNNeuronSigmoidNode instance.
func NewCNNNeuronSigmoidNode() CNNNeuronSigmoidNode {
	return getCNNNeuronSigmoidNodeClass().New()
}





// A representation of a sigmoid neuron filter.


// A representation of a sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoidNode
type CNNNeuronSigmoidNode struct {
	CNNNeuronNode
}

// CNNNeuronSigmoidNodeFrom constructs a [CNNNeuronSigmoidNode] from an unsafe.Pointer.
//
// A representation of a sigmoid neuron filter.
func CNNNeuronSigmoidNodeFrom(ptr unsafe.Pointer) CNNNeuronSigmoidNode {
	return CNNNeuronSigmoidNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsigmoidnode/2921458-initwithsource
func NewCNNNeuronSigmoidNodeWithSource(sourceNode IImageNode) CNNNeuronSigmoidNode {
	instance := getCNNNeuronSigmoidNodeClass().Alloc()
	rv := objc.Send[CNNNeuronSigmoidNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsigmoidnode/2866467-nodewithsource
func (cc _CNNNeuronSigmoidNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















