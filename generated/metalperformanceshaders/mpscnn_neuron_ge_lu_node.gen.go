// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronGeLUNode] class.
var (
	CNNNeuronGeLUNodeClass     _CNNNeuronGeLUNodeClass
	CNNNeuronGeLUNodeClassOnce sync.Once
)

func getCNNNeuronGeLUNodeClass() _CNNNeuronGeLUNodeClass {
	CNNNeuronGeLUNodeClassOnce.Do(func() {
		CNNNeuronGeLUNodeClass = _CNNNeuronGeLUNodeClass{objc.GetClass("MPSCNNNeuronGeLUNode")}
	})
	return CNNNeuronGeLUNodeClass
}

type _CNNNeuronGeLUNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronGeLUNode] class.
type ICNNNeuronGeLUNode interface {
	ICNNNeuronNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronGeLUNodeClass) Alloc() CNNNeuronGeLUNode {
	rv := objc.Send[CNNNeuronGeLUNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronGeLUNodeClass) New() CNNNeuronGeLUNode {
	rv := objc.Send[CNNNeuronGeLUNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronGeLUNode) Init() CNNNeuronGeLUNode {
	rv := objc.Send[CNNNeuronGeLUNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronGeLUNode) Autorelease() CNNNeuronGeLUNode {
	rv := objc.Send[CNNNeuronGeLUNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronGeLUNode creates a new CNNNeuronGeLUNode instance.
func NewCNNNeuronGeLUNode() CNNNeuronGeLUNode {
	return getCNNNeuronGeLUNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGeLUNode
type CNNNeuronGeLUNode struct {
	CNNNeuronNode
}

// CNNNeuronGeLUNodeFrom constructs a [CNNNeuronGeLUNode] from an unsafe.Pointer.
func CNNNeuronGeLUNodeFrom(ptr unsafe.Pointer) CNNNeuronGeLUNode {
	return CNNNeuronGeLUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongelunode/3237266-initwithsource
func NewCNNNeuronGeLUNodeWithSource(sourceNode IImageNode) CNNNeuronGeLUNode {
	instance := getCNNNeuronGeLUNodeClass().Alloc()
	rv := objc.Send[CNNNeuronGeLUNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongelunode/3237267-nodewithsource
func (cc _CNNNeuronGeLUNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















