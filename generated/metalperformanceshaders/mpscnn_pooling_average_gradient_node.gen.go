// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNPoolingAverageGradientNode] class.
var (
	CNNPoolingAverageGradientNodeClass     _CNNPoolingAverageGradientNodeClass
	CNNPoolingAverageGradientNodeClassOnce sync.Once
)

func getCNNPoolingAverageGradientNodeClass() _CNNPoolingAverageGradientNodeClass {
	CNNPoolingAverageGradientNodeClassOnce.Do(func() {
		CNNPoolingAverageGradientNodeClass = _CNNPoolingAverageGradientNodeClass{objc.GetClass("MPSCNNPoolingAverageGradientNode")}
	})
	return CNNPoolingAverageGradientNodeClass
}

type _CNNPoolingAverageGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingAverageGradientNode] class.
type ICNNPoolingAverageGradientNode interface {
	ICNNPoolingGradientNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingAverageGradientNodeClass) Alloc() CNNPoolingAverageGradientNode {
	rv := objc.Send[CNNPoolingAverageGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingAverageGradientNodeClass) New() CNNPoolingAverageGradientNode {
	rv := objc.Send[CNNPoolingAverageGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingAverageGradientNode) Init() CNNPoolingAverageGradientNode {
	rv := objc.Send[CNNPoolingAverageGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingAverageGradientNode) Autorelease() CNNPoolingAverageGradientNode {
	rv := objc.Send[CNNPoolingAverageGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingAverageGradientNode creates a new CNNPoolingAverageGradientNode instance.
func NewCNNPoolingAverageGradientNode() CNNPoolingAverageGradientNode {
	return getCNNPoolingAverageGradientNodeClass().New()
}





// A representation of a gradient average pooling filter.


// A representation of a gradient average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradientNode
type CNNPoolingAverageGradientNode struct {
	CNNPoolingGradientNode
}

// CNNPoolingAverageGradientNodeFrom constructs a [CNNPoolingAverageGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient average pooling filter.
func CNNPoolingAverageGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingAverageGradientNode {
	return CNNPoolingAverageGradientNode{
		CNNPoolingGradientNode: CNNPoolingGradientNodeFrom(ptr),
	}
}































