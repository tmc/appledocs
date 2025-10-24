// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNSoftMaxNode] class.
var (
	CNNSoftMaxNodeClass     _CNNSoftMaxNodeClass
	CNNSoftMaxNodeClassOnce sync.Once
)

func getCNNSoftMaxNodeClass() _CNNSoftMaxNodeClass {
	CNNSoftMaxNodeClassOnce.Do(func() {
		CNNSoftMaxNodeClass = _CNNSoftMaxNodeClass{objc.GetClass("MPSCNNSoftMaxNode")}
	})
	return CNNSoftMaxNodeClass
}

type _CNNSoftMaxNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNSoftMaxNode] class.
type ICNNSoftMaxNode interface {
	IFilterNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNSoftMaxNodeClass) Alloc() CNNSoftMaxNode {
	rv := objc.Send[CNNSoftMaxNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSoftMaxNodeClass) New() CNNSoftMaxNode {
	rv := objc.Send[CNNSoftMaxNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSoftMaxNode) Init() CNNSoftMaxNode {
	rv := objc.Send[CNNSoftMaxNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSoftMaxNode) Autorelease() CNNSoftMaxNode {
	rv := objc.Send[CNNSoftMaxNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSoftMaxNode creates a new CNNSoftMaxNode instance.
func NewCNNSoftMaxNode() CNNSoftMaxNode {
	return getCNNSoftMaxNodeClass().New()
}





// A representation of a softmax filter.


// A representation of a softmax filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxNode
type CNNSoftMaxNode struct {
	FilterNode
}

// CNNSoftMaxNodeFrom constructs a [CNNSoftMaxNode] from an unsafe.Pointer.
//
// A representation of a softmax filter.
func CNNSoftMaxNodeFrom(ptr unsafe.Pointer) CNNSoftMaxNode {
	return CNNSoftMaxNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxnode/2866408-initwithsource
func NewCNNSoftMaxNodeWithSource(sourceNode IImageNode) CNNSoftMaxNode {
	instance := getCNNSoftMaxNodeClass().Alloc()
	rv := objc.Send[CNNSoftMaxNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxnode/2866455-nodewithsource
func (cc _CNNSoftMaxNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















