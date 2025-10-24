// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNLogSoftMaxNode] class.
var (
	CNNLogSoftMaxNodeClass     _CNNLogSoftMaxNodeClass
	CNNLogSoftMaxNodeClassOnce sync.Once
)

func getCNNLogSoftMaxNodeClass() _CNNLogSoftMaxNodeClass {
	CNNLogSoftMaxNodeClassOnce.Do(func() {
		CNNLogSoftMaxNodeClass = _CNNLogSoftMaxNodeClass{objc.GetClass("MPSCNNLogSoftMaxNode")}
	})
	return CNNLogSoftMaxNodeClass
}

type _CNNLogSoftMaxNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNLogSoftMaxNode] class.
type ICNNLogSoftMaxNode interface {
	IFilterNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNLogSoftMaxNodeClass) Alloc() CNNLogSoftMaxNode {
	rv := objc.Send[CNNLogSoftMaxNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLogSoftMaxNodeClass) New() CNNLogSoftMaxNode {
	rv := objc.Send[CNNLogSoftMaxNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLogSoftMaxNode) Init() CNNLogSoftMaxNode {
	rv := objc.Send[CNNLogSoftMaxNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLogSoftMaxNode) Autorelease() CNNLogSoftMaxNode {
	rv := objc.Send[CNNLogSoftMaxNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLogSoftMaxNode creates a new CNNLogSoftMaxNode instance.
func NewCNNLogSoftMaxNode() CNNLogSoftMaxNode {
	return getCNNLogSoftMaxNodeClass().New()
}





// A representation of a logarithmic softmax filter kernel.


// A representation of a logarithmic softmax filter kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxNode
type CNNLogSoftMaxNode struct {
	FilterNode
}

// CNNLogSoftMaxNodeFrom constructs a [CNNLogSoftMaxNode] from an unsafe.Pointer.
//
// A representation of a logarithmic softmax filter kernel.
func CNNLogSoftMaxNodeFrom(ptr unsafe.Pointer) CNNLogSoftMaxNode {
	return CNNLogSoftMaxNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxnode/2866457-initwithsource
func NewCNNLogSoftMaxNodeWithSource(sourceNode IImageNode) CNNLogSoftMaxNode {
	instance := getCNNLogSoftMaxNodeClass().Alloc()
	rv := objc.Send[CNNLogSoftMaxNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxnode/2866434-nodewithsource
func (cc _CNNLogSoftMaxNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}






















