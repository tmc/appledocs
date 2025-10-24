// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNPoolingMaxNode] class.
var (
	CNNPoolingMaxNodeClass     _CNNPoolingMaxNodeClass
	CNNPoolingMaxNodeClassOnce sync.Once
)

func getCNNPoolingMaxNodeClass() _CNNPoolingMaxNodeClass {
	CNNPoolingMaxNodeClassOnce.Do(func() {
		CNNPoolingMaxNodeClass = _CNNPoolingMaxNodeClass{objc.GetClass("MPSCNNPoolingMaxNode")}
	})
	return CNNPoolingMaxNodeClass
}

type _CNNPoolingMaxNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingMaxNode] class.
type ICNNPoolingMaxNode interface {
	ICNNPoolingNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingMaxNodeClass) Alloc() CNNPoolingMaxNode {
	rv := objc.Send[CNNPoolingMaxNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingMaxNodeClass) New() CNNPoolingMaxNode {
	rv := objc.Send[CNNPoolingMaxNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingMaxNode) Init() CNNPoolingMaxNode {
	rv := objc.Send[CNNPoolingMaxNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingMaxNode) Autorelease() CNNPoolingMaxNode {
	rv := objc.Send[CNNPoolingMaxNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingMaxNode creates a new CNNPoolingMaxNode instance.
func NewCNNPoolingMaxNode() CNNPoolingMaxNode {
	return getCNNPoolingMaxNodeClass().New()
}





// A representation of a max pooling filter.


// A representation of a max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxNode
type CNNPoolingMaxNode struct {
	CNNPoolingNode
}

// CNNPoolingMaxNodeFrom constructs a [CNNPoolingMaxNode] from an unsafe.Pointer.
//
// A representation of a max pooling filter.
func CNNPoolingMaxNodeFrom(ptr unsafe.Pointer) CNNPoolingMaxNode {
	return CNNPoolingMaxNode{
		CNNPoolingNode: CNNPoolingNodeFrom(ptr),
	}
}































