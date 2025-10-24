// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNNConvolutionTransposeGradientState] class.
var (
	CNNConvolutionTransposeGradientStateClass     _CNNConvolutionTransposeGradientStateClass
	CNNConvolutionTransposeGradientStateClassOnce sync.Once
)

func getCNNConvolutionTransposeGradientStateClass() _CNNConvolutionTransposeGradientStateClass {
	CNNConvolutionTransposeGradientStateClassOnce.Do(func() {
		CNNConvolutionTransposeGradientStateClass = _CNNConvolutionTransposeGradientStateClass{objc.GetClass("MPSCNNConvolutionTransposeGradientState")}
	})
	return CNNConvolutionTransposeGradientStateClass
}

type _CNNConvolutionTransposeGradientStateClass struct {
	class objc.Class
}

// An interface definition for the [CNNConvolutionTransposeGradientState] class.
type ICNNConvolutionTransposeGradientState interface {
	ICNNConvolutionGradientState
	// properties:
	ConvolutionTranspose() CNNConvolutionTranspose /* not a class type */
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientState
type CNNConvolutionTransposeGradientState struct {
	CNNConvolutionGradientState
}

// CNNConvolutionTransposeGradientStateFrom constructs a [CNNConvolutionTransposeGradientState] from an unsafe.Pointer.
func CNNConvolutionTransposeGradientStateFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientState{
		CNNConvolutionGradientState: CNNConvolutionGradientStateFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeGradientStateClass) Alloc() CNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNConvolutionTransposeGradientStateClass) New() CNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionTransposeGradientState) Init() CNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionTransposeGradientState) Autorelease() CNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionTransposeGradientState creates a new CNNConvolutionTransposeGradientState instance.
func NewCNNConvolutionTransposeGradientState() CNNConvolutionTransposeGradientState {
	return getCNNConvolutionTransposeGradientStateClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientState/convolutionTranspose
func (c_ CNNConvolutionTransposeGradientState) ConvolutionTranspose() CNNConvolutionTranspose /* not a class type */ {
	rv := objc.Send[CNNConvolutionTranspose](c_.ID, objc.Sel("convolutionTranspose"))
	return rv
}



