// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNNConvolutionGradientNode] class.
var (
	CNNConvolutionGradientNodeClass     _CNNConvolutionGradientNodeClass
	CNNConvolutionGradientNodeClassOnce sync.Once
)

func getCNNConvolutionGradientNodeClass() _CNNConvolutionGradientNodeClass {
	CNNConvolutionGradientNodeClassOnce.Do(func() {
		CNNConvolutionGradientNodeClass = _CNNConvolutionGradientNodeClass{objc.GetClass("MPSCNNConvolutionGradientNode")}
	})
	return CNNConvolutionGradientNodeClass
}

type _CNNConvolutionGradientNodeClass struct {
	class objc.Class
}

// An interface definition for the [CNNConvolutionGradientNode] class.
type ICNNConvolutionGradientNode interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type CNNConvolutionGradientNode struct {
	objectivec.Object
}

// CNNConvolutionGradientNodeFrom constructs a [CNNConvolutionGradientNode] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func CNNConvolutionGradientNodeFrom(ptr unsafe.Pointer) CNNConvolutionGradientNode {
	return CNNConvolutionGradientNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionGradientNodeClass) Alloc() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNConvolutionGradientNodeClass) New() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionGradientNode) Init() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionGradientNode) Autorelease() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionGradientNode creates a new CNNConvolutionGradientNode instance.
func NewCNNConvolutionGradientNode() CNNConvolutionGradientNode {
	return getCNNConvolutionGradientNodeClass().New()
}




