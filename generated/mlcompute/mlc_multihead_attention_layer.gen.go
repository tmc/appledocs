// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CMultiheadAttentionLayer] class.
var (
	CMultiheadAttentionLayerClass     _CMultiheadAttentionLayerClass
	CMultiheadAttentionLayerClassOnce sync.Once
)

func getCMultiheadAttentionLayerClass() _CMultiheadAttentionLayerClass {
	CMultiheadAttentionLayerClassOnce.Do(func() {
		CMultiheadAttentionLayerClass = _CMultiheadAttentionLayerClass{objc.GetClass("MLCMultiheadAttentionLayer")}
	})
	return CMultiheadAttentionLayerClass
}

type _CMultiheadAttentionLayerClass struct {
	class objc.Class
}

// An interface definition for the [CMultiheadAttentionLayer] class.
type ICMultiheadAttentionLayer interface {
	ICLayer
	// properties:
	AttentionBiases() IMLCTensor
	SetAttentionBiases(value IMLCTensor)
	Biases() IMLCTensor
	SetBiases(value IMLCTensor)
	BiasesParameters() objc.IObject /* cross-framework: CTensorParameter */
	SetBiasesParameters(value objc.IObject /* cross-framework: CTensorParameter */)
	Descriptor() CMultiheadAttentionDescriptor /* not a class type */
	SetDescriptor(value CMultiheadAttentionDescriptor /* not a class type */)
	Weights() IMLCTensor
	SetWeights(value IMLCTensor)
	WeightsParameters() objc.IObject /* cross-framework: CTensorParameter */
	SetWeightsParameters(value objc.IObject /* cross-framework: CTensorParameter */)
	// methods:
}

// A multihead, scaled dot-product attention layer that attends to one or more entries in the input key-value pairs.
//
// The dimensions of projections are as follows: ``


// A multihead, scaled dot-product attention layer that attends to one or more entries in the input key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionLayer
type CMultiheadAttentionLayer struct {
	CLayer
}

// CMultiheadAttentionLayerFrom constructs a [CMultiheadAttentionLayer] from an unsafe.Pointer.
//
// A multihead, scaled dot-product attention layer that attends to one or more entries in the input key-value pairs.
func CMultiheadAttentionLayerFrom(ptr unsafe.Pointer) CMultiheadAttentionLayer {
	return CMultiheadAttentionLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CMultiheadAttentionLayerClass) Alloc() CMultiheadAttentionLayer {
	rv := objc.Send[CMultiheadAttentionLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CMultiheadAttentionLayerClass) New() CMultiheadAttentionLayer {
	rv := objc.Send[CMultiheadAttentionLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CMultiheadAttentionLayer) Init() CMultiheadAttentionLayer {
	rv := objc.Send[CMultiheadAttentionLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CMultiheadAttentionLayer) Autorelease() CMultiheadAttentionLayer {
	rv := objc.Send[CMultiheadAttentionLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCMultiheadAttentionLayer creates a new CMultiheadAttentionLayer instance.
func NewCMultiheadAttentionLayer() CMultiheadAttentionLayer {
	return getCMultiheadAttentionLayerClass().New()
}



// The array of attention biases you use for key and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/attentionbiases
func (c_ CMultiheadAttentionLayer) AttentionBiases() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("attentionBiases"))
	return rv
}


// The array of attention biases you use for key and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/attentionbiases
func (c_ CMultiheadAttentionLayer) SetAttentionBiases(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttentionBiases:"), value)
}


// The array of biases you use for query, key, value, and output projections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/biases
func (c_ CMultiheadAttentionLayer) Biases() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("biases"))
	return rv
}


// The array of biases you use for query, key, value, and output projections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/biases
func (c_ CMultiheadAttentionLayer) SetBiases(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBiases:"), value)
}


// The array of biases tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/biasesparameters
func (c_ CMultiheadAttentionLayer) BiasesParameters() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("biasesParameters"))
	return rv
}


// The array of biases tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/biasesparameters
func (c_ CMultiheadAttentionLayer) SetBiasesParameters(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBiasesParameters:"), value)
}


// The configuration object you use to create the multi-head attention layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/descriptor
func (c_ CMultiheadAttentionLayer) Descriptor() CMultiheadAttentionDescriptor /* not a class type */ {
	rv := objc.Send[CMultiheadAttentionDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}


// The configuration object you use to create the multi-head attention layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/descriptor
func (c_ CMultiheadAttentionLayer) SetDescriptor(value CMultiheadAttentionDescriptor /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDescriptor:"), value)
}


// The array of weights you use for query, key, value, and output projections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/weights
func (c_ CMultiheadAttentionLayer) Weights() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("weights"))
	return rv
}


// The array of weights you use for query, key, value, and output projections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/weights
func (c_ CMultiheadAttentionLayer) SetWeights(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeights:"), value)
}


// The array of weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/weightsparameters
func (c_ CMultiheadAttentionLayer) WeightsParameters() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("weightsParameters"))
	return rv
}


// The array of weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcmultiheadattentionlayer/weightsparameters
func (c_ CMultiheadAttentionLayer) SetWeightsParameters(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeightsParameters:"), value)
}



