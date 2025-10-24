// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNPoolingMaxGradientNode] class.
var (
	CNNPoolingMaxGradientNodeClass     _CNNPoolingMaxGradientNodeClass
	CNNPoolingMaxGradientNodeClassOnce sync.Once
)

func getCNNPoolingMaxGradientNodeClass() _CNNPoolingMaxGradientNodeClass {
	CNNPoolingMaxGradientNodeClassOnce.Do(func() {
		CNNPoolingMaxGradientNodeClass = _CNNPoolingMaxGradientNodeClass{objc.GetClass("MPSCNNPoolingMaxGradientNode")}
	})
	return CNNPoolingMaxGradientNodeClass
}

type _CNNPoolingMaxGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingMaxGradientNode] class.
type ICNNPoolingMaxGradientNode interface {
	ICNNPoolingGradientNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingMaxGradientNodeClass) Alloc() CNNPoolingMaxGradientNode {
	rv := objc.Send[CNNPoolingMaxGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingMaxGradientNodeClass) New() CNNPoolingMaxGradientNode {
	rv := objc.Send[CNNPoolingMaxGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingMaxGradientNode) Init() CNNPoolingMaxGradientNode {
	rv := objc.Send[CNNPoolingMaxGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingMaxGradientNode) Autorelease() CNNPoolingMaxGradientNode {
	rv := objc.Send[CNNPoolingMaxGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingMaxGradientNode creates a new CNNPoolingMaxGradientNode instance.
func NewCNNPoolingMaxGradientNode() CNNPoolingMaxGradientNode {
	return getCNNPoolingMaxGradientNodeClass().New()
}





// A representation of a gradient max pooling filter.


// A representation of a gradient max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxGradientNode
type CNNPoolingMaxGradientNode struct {
	CNNPoolingGradientNode
}

// CNNPoolingMaxGradientNodeFrom constructs a [CNNPoolingMaxGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient max pooling filter.
func CNNPoolingMaxGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingMaxGradientNode {
	return CNNPoolingMaxGradientNode{
		CNNPoolingGradientNode: CNNPoolingGradientNodeFrom(ptr),
	}
}































