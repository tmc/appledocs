// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CLSTMLayer] class.
var (
	CLSTMLayerClass     _CLSTMLayerClass
	CLSTMLayerClassOnce sync.Once
)

func getCLSTMLayerClass() _CLSTMLayerClass {
	CLSTMLayerClassOnce.Do(func() {
		CLSTMLayerClass = _CLSTMLayerClass{objc.GetClass("MLCLSTMLayer")}
	})
	return CLSTMLayerClass
}

type _CLSTMLayerClass struct {
	class objc.Class
}

// An interface definition for the [CLSTMLayer] class.
type ICLSTMLayer interface {
	ICLayer
	// properties:
	LayerCount() int
	SetLayerCount(value int)
	Biases() IMLCTensor
	SetBiases(value IMLCTensor)
	BiasesParameters() objc.IObject /* cross-framework: CTensorParameter */
	SetBiasesParameters(value objc.IObject /* cross-framework: CTensorParameter */)
	Descriptor() CLSTMDescriptor /* not a class type */
	SetDescriptor(value CLSTMDescriptor /* not a class type */)
	GateActivations() CActivationDescriptor /* not a class type */
	SetGateActivations(value CActivationDescriptor /* not a class type */)
	HiddenWeights() IMLCTensor
	SetHiddenWeights(value IMLCTensor)
	HiddenWeightsParameters() objc.IObject /* cross-framework: CTensorParameter */
	SetHiddenWeightsParameters(value objc.IObject /* cross-framework: CTensorParameter */)
	InputWeights() IMLCTensor
	SetInputWeights(value IMLCTensor)
	InputWeightsParameters() objc.IObject /* cross-framework: CTensorParameter */
	SetInputWeightsParameters(value objc.IObject /* cross-framework: CTensorParameter */)
	OutputResultActivation() CActivationDescriptor /* not a class type */
	SetOutputResultActivation(value CActivationDescriptor /* not a class type */)
	PeepholeWeights() IMLCTensor
	SetPeepholeWeights(value IMLCTensor)
	PeepholeWeightsParameters() objc.IObject /* cross-framework: CTensorParameter */
	SetPeepholeWeightsParameters(value objc.IObject /* cross-framework: CTensorParameter */)
	// methods:
}

// A layer that represents long short-term memory (LSTM) networks.
//
// Use this class to create an LSTM layer with one of the following configurations:


// A layer that represents long short-term memory (LSTM) networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer
type CLSTMLayer struct {
	CLayer
}

// CLSTMLayerFrom constructs a [CLSTMLayer] from an unsafe.Pointer.
//
// A layer that represents long short-term memory (LSTM) networks.
func CLSTMLayerFrom(ptr unsafe.Pointer) CLSTMLayer {
	return CLSTMLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _CLSTMLayerClass) Alloc() CLSTMLayer {
	rv := objc.Send[CLSTMLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _CLSTMLayerClass) New() CLSTMLayer {
	rv := objc.Send[CLSTMLayer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ CLSTMLayer) Init() CLSTMLayer {
	rv := objc.Send[CLSTMLayer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ CLSTMLayer) Autorelease() CLSTMLayer {
	rv := objc.Send[CLSTMLayer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLSTMLayer creates a new CLSTMLayer instance.
func NewCLSTMLayer() CLSTMLayer {
	return getCLSTMLayerClass().New()
}



// The number of recurrent layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmdescriptor/layercount
func (s_ CLSTMLayer) LayerCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("layerCount"))
	return rv
}


// The number of recurrent layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmdescriptor/layercount
func (s_ CLSTMLayer) SetLayerCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLayerCount:"), value)
}


// The array of tensors that describe the bias terms you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/biases
func (s_ CLSTMLayer) Biases() IMLCTensor {
	rv := objc.Send[CTensor](s_.ID, objc.Sel("biases"))
	return rv
}


// The array of tensors that describe the bias terms you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/biases
func (s_ CLSTMLayer) SetBiases(value IMLCTensor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBiases:"), value)
}


// The biases tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/biasesparameters
func (s_ CLSTMLayer) BiasesParameters() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](s_.ID, objc.Sel("biasesParameters"))
	return rv
}


// The biases tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/biasesparameters
func (s_ CLSTMLayer) SetBiasesParameters(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBiasesParameters:"), value)
}


// The configuration object you use to create the LSTM layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/descriptor
func (s_ CLSTMLayer) Descriptor() CLSTMDescriptor /* not a class type */ {
	rv := objc.Send[STMDescriptor](s_.ID, objc.Sel("descriptor"))
	return rv
}


// The configuration object you use to create the LSTM layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/descriptor
func (s_ CLSTMLayer) SetDescriptor(value CLSTMDescriptor /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDescriptor:"), value)
}


// The array of gate activations you use for input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/gateactivations
func (s_ CLSTMLayer) GateActivations() CActivationDescriptor /* not a class type */ {
	rv := objc.Send[CActivationDescriptor](s_.ID, objc.Sel("gateActivations"))
	return rv
}


// The array of gate activations you use for input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/gateactivations
func (s_ CLSTMLayer) SetGateActivations(value CActivationDescriptor /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setGateActivations:"), value)
}


// The array of tensors that describe the hidden weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/hiddenweights
func (s_ CLSTMLayer) HiddenWeights() IMLCTensor {
	rv := objc.Send[CTensor](s_.ID, objc.Sel("hiddenWeights"))
	return rv
}


// The array of tensors that describe the hidden weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/hiddenweights
func (s_ CLSTMLayer) SetHiddenWeights(value IMLCTensor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHiddenWeights:"), value)
}


// The hidden weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/hiddenweightsparameters
func (s_ CLSTMLayer) HiddenWeightsParameters() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](s_.ID, objc.Sel("hiddenWeightsParameters"))
	return rv
}


// The hidden weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/hiddenweightsparameters
func (s_ CLSTMLayer) SetHiddenWeightsParameters(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHiddenWeightsParameters:"), value)
}


// The array of tensors that describe the input weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/inputweights
func (s_ CLSTMLayer) InputWeights() IMLCTensor {
	rv := objc.Send[CTensor](s_.ID, objc.Sel("inputWeights"))
	return rv
}


// The array of tensors that describe the input weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/inputweights
func (s_ CLSTMLayer) SetInputWeights(value IMLCTensor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInputWeights:"), value)
}


// The input weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/inputweightsparameters
func (s_ CLSTMLayer) InputWeightsParameters() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](s_.ID, objc.Sel("inputWeightsParameters"))
	return rv
}


// The input weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/inputweightsparameters
func (s_ CLSTMLayer) SetInputWeightsParameters(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInputWeightsParameters:"), value)
}


// The output activation descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/outputresultactivation
func (s_ CLSTMLayer) OutputResultActivation() CActivationDescriptor /* not a class type */ {
	rv := objc.Send[CActivationDescriptor](s_.ID, objc.Sel("outputResultActivation"))
	return rv
}


// The output activation descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/outputresultactivation
func (s_ CLSTMLayer) SetOutputResultActivation(value CActivationDescriptor /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setOutputResultActivation:"), value)
}


// The array of tensors that describe the peephole weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/peepholeweights
func (s_ CLSTMLayer) PeepholeWeights() IMLCTensor {
	rv := objc.Send[CTensor](s_.ID, objc.Sel("peepholeWeights"))
	return rv
}


// The array of tensors that describe the peephole weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/peepholeweights
func (s_ CLSTMLayer) SetPeepholeWeights(value IMLCTensor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPeepholeWeights:"), value)
}


// The peephole weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/peepholeweightsparameters
func (s_ CLSTMLayer) PeepholeWeightsParameters() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](s_.ID, objc.Sel("peepholeWeightsParameters"))
	return rv
}


// The peephole weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/peepholeweightsparameters
func (s_ CLSTMLayer) SetPeepholeWeightsParameters(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPeepholeWeightsParameters:"), value)
}



