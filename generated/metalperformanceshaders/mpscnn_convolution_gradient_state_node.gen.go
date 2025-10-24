// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNNConvolutionGradientStateNode] class.
var (
	CNNConvolutionGradientStateNodeClass     _CNNConvolutionGradientStateNodeClass
	CNNConvolutionGradientStateNodeClassOnce sync.Once
)

func getCNNConvolutionGradientStateNodeClass() _CNNConvolutionGradientStateNodeClass {
	CNNConvolutionGradientStateNodeClassOnce.Do(func() {
		CNNConvolutionGradientStateNodeClass = _CNNConvolutionGradientStateNodeClass{objc.GetClass("MPSCNNConvolutionGradientStateNode")}
	})
	return CNNConvolutionGradientStateNodeClass
}

type _CNNConvolutionGradientStateNodeClass struct {
	class objc.Class
}

// An interface definition for the [CNNConvolutionGradientStateNode] class.
type ICNNConvolutionGradientStateNode interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type CNNConvolutionGradientStateNode struct {
	objectivec.Object
}

// CNNConvolutionGradientStateNodeFrom constructs a [CNNConvolutionGradientStateNode] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func CNNConvolutionGradientStateNodeFrom(ptr unsafe.Pointer) CNNConvolutionGradientStateNode {
	return CNNConvolutionGradientStateNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionGradientStateNodeClass) Alloc() CNNConvolutionGradientStateNode {
	rv := objc.Send[CNNConvolutionGradientStateNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNConvolutionGradientStateNodeClass) New() CNNConvolutionGradientStateNode {
	rv := objc.Send[CNNConvolutionGradientStateNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionGradientStateNode) Init() CNNConvolutionGradientStateNode {
	rv := objc.Send[CNNConvolutionGradientStateNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionGradientStateNode) Autorelease() CNNConvolutionGradientStateNode {
	rv := objc.Send[CNNConvolutionGradientStateNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionGradientStateNode creates a new CNNConvolutionGradientStateNode instance.
func NewCNNConvolutionGradientStateNode() CNNConvolutionGradientStateNode {
	return getCNNConvolutionGradientStateNodeClass().New()
}




