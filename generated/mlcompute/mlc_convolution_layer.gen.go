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
	// properties:
	Biases() IMLCTensor
	SetBiases(value IMLCTensor)
	BiasesParameter() objc.IObject /* cross-framework: CTensorParameter */
	SetBiasesParameter(value objc.IObject /* cross-framework: CTensorParameter */)
	Descriptor() CConvolutionDescriptor /* not a class type */
	SetDescriptor(value CConvolutionDescriptor /* not a class type */)
	Weights() IMLCTensor
	SetWeights(value IMLCTensor)
	WeightsParameter() objc.IObject /* cross-framework: CTensorParameter */
	SetWeightsParameter(value objc.IObject /* cross-framework: CTensorParameter */)
	// methods:
}

// A layer that applies a convolution over a signal.


// A layer that applies a convolution over a signal.
//
// [Full Topic]
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



// The biases tensor you use for the convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/biases
func (c_ CConvolutionLayer) Biases() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("biases"))
	return rv
}


// The biases tensor you use for the convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/biases
func (c_ CConvolutionLayer) SetBiases(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBiases:"), value)
}


// The biases tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/biasesparameter
func (c_ CConvolutionLayer) BiasesParameter() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("biasesParameter"))
	return rv
}


// The biases tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/biasesparameter
func (c_ CConvolutionLayer) SetBiasesParameter(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBiasesParameter:"), value)
}


// The configuration object you use to create the convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/descriptor
func (c_ CConvolutionLayer) Descriptor() CConvolutionDescriptor /* not a class type */ {
	rv := objc.Send[CConvolutionDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}


// The configuration object you use to create the convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/descriptor
func (c_ CConvolutionLayer) SetDescriptor(value CConvolutionDescriptor /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDescriptor:"), value)
}


// The weights tensor you use for the convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/weights
func (c_ CConvolutionLayer) Weights() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("weights"))
	return rv
}


// The weights tensor you use for the convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/weights
func (c_ CConvolutionLayer) SetWeights(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeights:"), value)
}


// The weights tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/weightsparameter
func (c_ CConvolutionLayer) WeightsParameter() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("weightsParameter"))
	return rv
}


// The weights tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutionlayer/weightsparameter
func (c_ CConvolutionLayer) SetWeightsParameter(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeightsParameter:"), value)
}



