// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CEmbeddingLayer] class.
var (
	CEmbeddingLayerClass     _CEmbeddingLayerClass
	CEmbeddingLayerClassOnce sync.Once
)

func getCEmbeddingLayerClass() _CEmbeddingLayerClass {
	CEmbeddingLayerClassOnce.Do(func() {
		CEmbeddingLayerClass = _CEmbeddingLayerClass{objc.GetClass("MLCEmbeddingLayer")}
	})
	return CEmbeddingLayerClass
}

type _CEmbeddingLayerClass struct {
	class objc.Class
}

// An interface definition for the [CEmbeddingLayer] class.
type ICEmbeddingLayer interface {
	ICLayer
	Descriptor() unsafe.Pointer
	SetDescriptor(value unsafe.Pointer)
	Weights() MLCTensor
	SetWeights(value IMLCTensor)
	WeightsParameter() MLCTensorParameter
	SetWeightsParameter(value IMLCTensorParameter)
}

// A layer that stores a word embedding.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingLayer
type CEmbeddingLayer struct {
	CLayer
}

// CEmbeddingLayerFrom constructs a [CEmbeddingLayer] from an unsafe.Pointer.
//
// A layer that stores a word embedding.
func CEmbeddingLayerFrom(ptr unsafe.Pointer) CEmbeddingLayer {
	return CEmbeddingLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CEmbeddingLayerClass) Alloc() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CEmbeddingLayerClass) New() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CEmbeddingLayer) Init() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CEmbeddingLayer) Autorelease() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCEmbeddingLayer creates a new CEmbeddingLayer instance.
func NewCEmbeddingLayer() CEmbeddingLayer {
	return getCEmbeddingLayerClass().New()
}


// The configuration object you use to create the embedding layer.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcembeddinglayer/descriptor
func (c_ CEmbeddingLayer) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("descriptor"))
	return rv
}


// SetDescriptor sets the value of the descriptor property.
// The configuration object you use to create the embedding layer.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcembeddinglayer/descriptor
func (c_ CEmbeddingLayer) SetDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDescriptor:"), value)
}

// The weights tensor that contains the word embedding.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcembeddinglayer/weights
func (c_ CEmbeddingLayer) Weights() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("weights"))
	return rv
}


// SetWeights sets the value of the weights property.
// The weights tensor that contains the word embedding.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcembeddinglayer/weights
func (c_ CEmbeddingLayer) SetWeights(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeights:"), value)
}

// The tensor parameter that describes the weights for the optimizer update.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcembeddinglayer/weightsparameter
func (c_ CEmbeddingLayer) WeightsParameter() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](c_.ID, objc.Sel("weightsParameter"))
	return rv
}


// SetWeightsParameter sets the value of the weightsParameter property.
// The tensor parameter that describes the weights for the optimizer update.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcembeddinglayer/weightsparameter
func (c_ CEmbeddingLayer) SetWeightsParameter(value IMLCTensorParameter) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeightsParameter:"), value)
}



