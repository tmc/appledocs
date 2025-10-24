// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNPoolingAverageNode] class.
var (
	CNNPoolingAverageNodeClass     _CNNPoolingAverageNodeClass
	CNNPoolingAverageNodeClassOnce sync.Once
)

func getCNNPoolingAverageNodeClass() _CNNPoolingAverageNodeClass {
	CNNPoolingAverageNodeClassOnce.Do(func() {
		CNNPoolingAverageNodeClass = _CNNPoolingAverageNodeClass{objc.GetClass("MPSCNNPoolingAverageNode")}
	})
	return CNNPoolingAverageNodeClass
}

type _CNNPoolingAverageNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingAverageNode] class.
type ICNNPoolingAverageNode interface {
	ICNNPoolingNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingAverageNodeClass) Alloc() CNNPoolingAverageNode {
	rv := objc.Send[CNNPoolingAverageNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingAverageNodeClass) New() CNNPoolingAverageNode {
	rv := objc.Send[CNNPoolingAverageNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingAverageNode) Init() CNNPoolingAverageNode {
	rv := objc.Send[CNNPoolingAverageNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingAverageNode) Autorelease() CNNPoolingAverageNode {
	rv := objc.Send[CNNPoolingAverageNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingAverageNode creates a new CNNPoolingAverageNode instance.
func NewCNNPoolingAverageNode() CNNPoolingAverageNode {
	return getCNNPoolingAverageNodeClass().New()
}





// A representation of an average pooling filter.


// A representation of an average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageNode
type CNNPoolingAverageNode struct {
	CNNPoolingNode
}

// CNNPoolingAverageNodeFrom constructs a [CNNPoolingAverageNode] from an unsafe.Pointer.
//
// A representation of an average pooling filter.
func CNNPoolingAverageNodeFrom(ptr unsafe.Pointer) CNNPoolingAverageNode {
	return CNNPoolingAverageNode{
		CNNPoolingNode: CNNPoolingNodeFrom(ptr),
	}
}































