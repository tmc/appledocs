// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CConvolutionLayer] class.
var (
	CConvolutionLayerClass     _CConvolutionLayerClass
	CConvolutionLayerClassOnce sync.Once
)

func getCConvolutionLayerClass() _CConvolutionLayerClass {
	CConvolutionLayerClassOnce.Do(func() {
		CConvolutionLayerClass = _CConvolutionLayerClass{objc.GetClass("MLCConvolutionLayer")}
	})
	return CConvolutionLayerClass
}

type _CConvolutionLayerClass struct {
	class objc.Class
}

// An interface definition for the [CConvolutionLayer] class.
type ICConvolutionLayer interface {
	ICLayer
}

// A layer that applies a convolution over a signal.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionLayer
type CConvolutionLayer struct {
	CLayer
}

// CConvolutionLayerFrom constructs a [CConvolutionLayer] from an unsafe.Pointer.
//
// A layer that applies a convolution over a signal.
func CConvolutionLayerFrom(ptr unsafe.Pointer) CConvolutionLayer {
	return CConvolutionLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CConvolutionLayerClass) Alloc() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CConvolutionLayerClass) New() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CConvolutionLayer) Init() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CConvolutionLayer) Autorelease() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCConvolutionLayer creates a new CConvolutionLayer instance.
func NewCConvolutionLayer() CConvolutionLayer {
	return getCConvolutionLayerClass().New()
}


// The weights tensor you use for the convolution layer.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/weights
func (c_ CConvolutionLayer) Weights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("weights"))
	return rv
}


// SetWeights sets the value of the weights property.
// The weights tensor you use for the convolution layer.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/weights
func (c_ CConvolutionLayer) SetWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeights:"), value)
}

// The biases tensor parameter you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/biasesparameter
func (c_ CConvolutionLayer) BiasesParameter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("biasesParameter"))
	return rv
}


// SetBiasesParameter sets the value of the biasesParameter property.
// The biases tensor parameter you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/biasesparameter
func (c_ CConvolutionLayer) SetBiasesParameter(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBiasesParameter:"), value)
}

// The weights tensor parameter you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/weightsparameter
func (c_ CConvolutionLayer) WeightsParameter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("weightsParameter"))
	return rv
}


// SetWeightsParameter sets the value of the weightsParameter property.
// The weights tensor parameter you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/weightsparameter
func (c_ CConvolutionLayer) SetWeightsParameter(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeightsParameter:"), value)
}

// The biases tensor you use for the convolution layer.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/biases
func (c_ CConvolutionLayer) Biases() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("biases"))
	return rv
}


// SetBiases sets the value of the biases property.
// The biases tensor you use for the convolution layer.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/biases
func (c_ CConvolutionLayer) SetBiases(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBiases:"), value)
}

// The configuration object you use to create the convolution layer.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/descriptor
func (c_ CConvolutionLayer) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("descriptor"))
	return rv
}


// SetDescriptor sets the value of the descriptor property.
// The configuration object you use to create the convolution layer.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/descriptor
func (c_ CConvolutionLayer) SetDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDescriptor:"), value)
}



