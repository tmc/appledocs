// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronSoftSignNode] class.
var (
	CNNNeuronSoftSignNodeClass     _CNNNeuronSoftSignNodeClass
	CNNNeuronSoftSignNodeClassOnce sync.Once
)

func getCNNNeuronSoftSignNodeClass() _CNNNeuronSoftSignNodeClass {
	CNNNeuronSoftSignNodeClassOnce.Do(func() {
		CNNNeuronSoftSignNodeClass = _CNNNeuronSoftSignNodeClass{objc.GetClass("MPSCNNNeuronSoftSignNode")}
	})
	return CNNNeuronSoftSignNodeClass
}

type _CNNNeuronSoftSignNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronSoftSignNode] class.
type ICNNNeuronSoftSignNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronSoftSignNodeClass) Alloc() CNNNeuronSoftSignNode {
	rv := objc.Send[CNNNeuronSoftSignNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronSoftSignNodeClass) New() CNNNeuronSoftSignNode {
	rv := objc.Send[CNNNeuronSoftSignNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronSoftSignNode) Init() CNNNeuronSoftSignNode {
	rv := objc.Send[CNNNeuronSoftSignNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronSoftSignNode) Autorelease() CNNNeuronSoftSignNode {
	rv := objc.Send[CNNNeuronSoftSignNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronSoftSignNode creates a new CNNNeuronSoftSignNode instance.
func NewCNNNeuronSoftSignNode() CNNNeuronSoftSignNode {
	return getCNNNeuronSoftSignNodeClass().New()
}





// A representation of a softsign neuron filter.


// A representation of a softsign neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftSignNode
type CNNNeuronSoftSignNode struct {
	CNNNeuronNode
}

// CNNNeuronSoftSignNodeFrom constructs a [CNNNeuronSoftSignNode] from an unsafe.Pointer.
//
// A representation of a softsign neuron filter.
func CNNNeuronSoftSignNodeFrom(ptr unsafe.Pointer) CNNNeuronSoftSignNode {
	return CNNNeuronSoftSignNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftsignnode/2921463-initwithsource
func NewCNNNeuronSoftSignNodeWithSource(sourceNode IImageNode) CNNNeuronSoftSignNode {
	instance := getCNNNeuronSoftSignNodeClass().Alloc()
	rv := objc.Send[CNNNeuronSoftSignNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftsignnode/2866428-nodewithsource
func (cc _CNNNeuronSoftSignNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















