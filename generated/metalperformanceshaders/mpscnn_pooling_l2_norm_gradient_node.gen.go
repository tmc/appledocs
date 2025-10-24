// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNPoolingL2NormGradientNode] class.
var (
	CNNPoolingL2NormGradientNodeClass     _CNNPoolingL2NormGradientNodeClass
	CNNPoolingL2NormGradientNodeClassOnce sync.Once
)

func getCNNPoolingL2NormGradientNodeClass() _CNNPoolingL2NormGradientNodeClass {
	CNNPoolingL2NormGradientNodeClassOnce.Do(func() {
		CNNPoolingL2NormGradientNodeClass = _CNNPoolingL2NormGradientNodeClass{objc.GetClass("MPSCNNPoolingL2NormGradientNode")}
	})
	return CNNPoolingL2NormGradientNodeClass
}

type _CNNPoolingL2NormGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingL2NormGradientNode] class.
type ICNNPoolingL2NormGradientNode interface {
	ICNNPoolingGradientNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingL2NormGradientNodeClass) Alloc() CNNPoolingL2NormGradientNode {
	rv := objc.Send[CNNPoolingL2NormGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingL2NormGradientNodeClass) New() CNNPoolingL2NormGradientNode {
	rv := objc.Send[CNNPoolingL2NormGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingL2NormGradientNode) Init() CNNPoolingL2NormGradientNode {
	rv := objc.Send[CNNPoolingL2NormGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingL2NormGradientNode) Autorelease() CNNPoolingL2NormGradientNode {
	rv := objc.Send[CNNPoolingL2NormGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingL2NormGradientNode creates a new CNNPoolingL2NormGradientNode instance.
func NewCNNPoolingL2NormGradientNode() CNNPoolingL2NormGradientNode {
	return getCNNPoolingL2NormGradientNodeClass().New()
}





// A representation of a gradient L2-norm pooling filter.


// A representation of a gradient L2-norm pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormGradientNode
type CNNPoolingL2NormGradientNode struct {
	CNNPoolingGradientNode
}

// CNNPoolingL2NormGradientNodeFrom constructs a [CNNPoolingL2NormGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient L2-norm pooling filter.
func CNNPoolingL2NormGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingL2NormGradientNode {
	return CNNPoolingL2NormGradientNode{
		CNNPoolingGradientNode: CNNPoolingGradientNodeFrom(ptr),
	}
}































