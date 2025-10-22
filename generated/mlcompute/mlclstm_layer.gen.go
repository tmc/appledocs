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
	LayerCount() int
	SetLayerCount(value int)
	Biases() MLCTensor
	SetBiases(value IMLCTensor)
	BiasesParameters() MLCTensorParameter
	SetBiasesParameters(value IMLCTensorParameter)
	Descriptor() unsafe.Pointer
	SetDescriptor(value unsafe.Pointer)
	GateActivations() unsafe.Pointer
	SetGateActivations(value unsafe.Pointer)
	HiddenWeights() MLCTensor
	SetHiddenWeights(value IMLCTensor)
	HiddenWeightsParameters() MLCTensorParameter
	SetHiddenWeightsParameters(value IMLCTensorParameter)
	InputWeights() MLCTensor
	SetInputWeights(value IMLCTensor)
	InputWeightsParameters() MLCTensorParameter
	SetInputWeightsParameters(value IMLCTensorParameter)
	OutputResultActivation() unsafe.Pointer
	SetOutputResultActivation(value unsafe.Pointer)
	PeepholeWeights() MLCTensor
	SetPeepholeWeights(value IMLCTensor)
	PeepholeWeightsParameters() MLCTensorParameter
	SetPeepholeWeightsParameters(value IMLCTensorParameter)
}

// A layer that represents long short-term memory (LSTM) networks.
//
// Use this class to create an LSTM layer with one of the following configurations:
//
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
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmdescriptor/layercount
func (s_ CLSTMLayer) LayerCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("layerCount"))
	return rv
}


// SetLayerCount sets the value of the layerCount property.
// The number of recurrent layers.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmdescriptor/layercount
func (s_ CLSTMLayer) SetLayerCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLayerCount:"), value)
}

// The array of tensors that describe the bias terms you use for the input, hidden, cell, and output gates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/biases
func (s_ CLSTMLayer) Biases() MLCTensor {
	rv := objc.Send[MLCTensor](s_.ID, objc.Sel("biases"))
	return rv
}


// SetBiases sets the value of the biases property.
// The array of tensors that describe the bias terms you use for the input, hidden, cell, and output gates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/biases
func (s_ CLSTMLayer) SetBiases(value IMLCTensor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBiases:"), value)
}

// The biases tensor parameters you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/biasesparameters
func (s_ CLSTMLayer) BiasesParameters() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](s_.ID, objc.Sel("biasesParameters"))
	return rv
}


// SetBiasesParameters sets the value of the biasesParameters property.
// The biases tensor parameters you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/biasesparameters
func (s_ CLSTMLayer) SetBiasesParameters(value IMLCTensorParameter) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBiasesParameters:"), value)
}

// The configuration object you use to create the LSTM layer.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/descriptor
func (s_ CLSTMLayer) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("descriptor"))
	return rv
}


// SetDescriptor sets the value of the descriptor property.
// The configuration object you use to create the LSTM layer.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/descriptor
func (s_ CLSTMLayer) SetDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDescriptor:"), value)
}

// The array of gate activations you use for input, hidden, cell, and output gates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/gateactivations
func (s_ CLSTMLayer) GateActivations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("gateActivations"))
	return rv
}


// SetGateActivations sets the value of the gateActivations property.
// The array of gate activations you use for input, hidden, cell, and output gates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/gateactivations
func (s_ CLSTMLayer) SetGateActivations(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setGateActivations:"), value)
}

// The array of tensors that describe the hidden weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/hiddenweights
func (s_ CLSTMLayer) HiddenWeights() MLCTensor {
	rv := objc.Send[MLCTensor](s_.ID, objc.Sel("hiddenWeights"))
	return rv
}


// SetHiddenWeights sets the value of the hiddenWeights property.
// The array of tensors that describe the hidden weights you use for the input, hidden, cell, and output gates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/hiddenweights
func (s_ CLSTMLayer) SetHiddenWeights(value IMLCTensor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHiddenWeights:"), value)
}

// The hidden weights tensor parameters you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/hiddenweightsparameters
func (s_ CLSTMLayer) HiddenWeightsParameters() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](s_.ID, objc.Sel("hiddenWeightsParameters"))
	return rv
}


// SetHiddenWeightsParameters sets the value of the hiddenWeightsParameters property.
// The hidden weights tensor parameters you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/hiddenweightsparameters
func (s_ CLSTMLayer) SetHiddenWeightsParameters(value IMLCTensorParameter) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHiddenWeightsParameters:"), value)
}

// The array of tensors that describe the input weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/inputweights
func (s_ CLSTMLayer) InputWeights() MLCTensor {
	rv := objc.Send[MLCTensor](s_.ID, objc.Sel("inputWeights"))
	return rv
}


// SetInputWeights sets the value of the inputWeights property.
// The array of tensors that describe the input weights you use for the input, hidden, cell, and output gates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/inputweights
func (s_ CLSTMLayer) SetInputWeights(value IMLCTensor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInputWeights:"), value)
}

// The input weights tensor parameters you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/inputweightsparameters
func (s_ CLSTMLayer) InputWeightsParameters() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](s_.ID, objc.Sel("inputWeightsParameters"))
	return rv
}


// SetInputWeightsParameters sets the value of the inputWeightsParameters property.
// The input weights tensor parameters you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/inputweightsparameters
func (s_ CLSTMLayer) SetInputWeightsParameters(value IMLCTensorParameter) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInputWeightsParameters:"), value)
}

// The output activation descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/outputresultactivation
func (s_ CLSTMLayer) OutputResultActivation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("outputResultActivation"))
	return rv
}


// SetOutputResultActivation sets the value of the outputResultActivation property.
// The output activation descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/outputresultactivation
func (s_ CLSTMLayer) SetOutputResultActivation(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setOutputResultActivation:"), value)
}

// The array of tensors that describe the peephole weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/peepholeweights
func (s_ CLSTMLayer) PeepholeWeights() MLCTensor {
	rv := objc.Send[MLCTensor](s_.ID, objc.Sel("peepholeWeights"))
	return rv
}


// SetPeepholeWeights sets the value of the peepholeWeights property.
// The array of tensors that describe the peephole weights you use for the input, hidden, cell, and output gates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/peepholeweights
func (s_ CLSTMLayer) SetPeepholeWeights(value IMLCTensor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPeepholeWeights:"), value)
}

// The peephole weights tensor parameters you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/peepholeweightsparameters
func (s_ CLSTMLayer) PeepholeWeightsParameters() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](s_.ID, objc.Sel("peepholeWeightsParameters"))
	return rv
}


// SetPeepholeWeightsParameters sets the value of the peepholeWeightsParameters property.
// The peephole weights tensor parameters you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmlayer/peepholeweightsparameters
func (s_ CLSTMLayer) SetPeepholeWeightsParameters(value IMLCTensorParameter) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPeepholeWeightsParameters:"), value)
}



