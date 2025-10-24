// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNNConvolutionGradientState] class.
var (
	CNNConvolutionGradientStateClass     _CNNConvolutionGradientStateClass
	CNNConvolutionGradientStateClassOnce sync.Once
)

func getCNNConvolutionGradientStateClass() _CNNConvolutionGradientStateClass {
	CNNConvolutionGradientStateClassOnce.Do(func() {
		CNNConvolutionGradientStateClass = _CNNConvolutionGradientStateClass{objc.GetClass("MPSCNNConvolutionGradientState")}
	})
	return CNNConvolutionGradientStateClass
}

type _CNNConvolutionGradientStateClass struct {
	class objc.Class
}

// An interface definition for the [CNNConvolutionGradientState] class.
type ICNNConvolutionGradientState interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type CNNConvolutionGradientState struct {
	objectivec.Object
}

// CNNConvolutionGradientStateFrom constructs a [CNNConvolutionGradientState] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func CNNConvolutionGradientStateFrom(ptr unsafe.Pointer) CNNConvolutionGradientState {
	return CNNConvolutionGradientState{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionGradientStateClass) Alloc() CNNConvolutionGradientState {
	rv := objc.Send[CNNConvolutionGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNConvolutionGradientStateClass) New() CNNConvolutionGradientState {
	rv := objc.Send[CNNConvolutionGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionGradientState) Init() CNNConvolutionGradientState {
	rv := objc.Send[CNNConvolutionGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionGradientState) Autorelease() CNNConvolutionGradientState {
	rv := objc.Send[CNNConvolutionGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionGradientState creates a new CNNConvolutionGradientState instance.
func NewCNNConvolutionGradientState() CNNConvolutionGradientState {
	return getCNNConvolutionGradientStateClass().New()
}




