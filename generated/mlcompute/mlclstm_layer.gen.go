// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCLSTMLayer */


/* debug [class_header]: Header for MLCLSTMLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CLSTMLayer */
// An interface definition for the [CLSTMLayer] class.
type ICLSTMLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CLSTMLayer */
	// properties:
	Biases() []CTensor
	BiasesParameters() []CTensorParameter
	Descriptor() IMLCLSTMDescriptor
	GateActivations() []CActivationDescriptor
	HiddenWeights() []CTensor
	HiddenWeightsParameters() []CTensorParameter
	InputWeights() []CTensor
	InputWeightsParameters() []CTensorParameter
	OutputResultActivation() IMLCActivationDescriptor
	PeepholeWeights() []CTensor
	PeepholeWeightsParameters() []CTensorParameter
	LayerCount() int
	SetLayerCount(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CLSTMLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CLSTMLayer */
// Alloc allocates a new instance without initialization.
func (sc _CLSTMLayerClass) Alloc() CLSTMLayer {
	rv := objc.Send[CLSTMLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CLSTMLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CLSTMLayer */

// Creates an LSTM layer with the descriptor, input and hidden weights, and biases you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/init(descriptor:inputWeights:hiddenWeights:biases:)
func NewCLSTMLayerWithDescriptorInputWeightsHiddenWeightsBiases(descriptor IMLCLSTMDescriptor, inputWeights []CTensor, hiddenWeights []CTensor, biases []CTensor) CLSTMLayer {
	rv := objc.Send[CLSTMLayer](objc.ID(getCLSTMLayerClass().class), objc.Sel("layerWithDescriptor:inputWeights:hiddenWeights:biases:"), descriptor, inputWeights, hiddenWeights, biases)
	return rv
}/* debug [class_init_methods/constructor]: NewCLSTMLayerWithDescriptorInputWeightsHiddenWeightsBiases */


// Creates an LSTM layer with the descriptor, weights, and biases you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/init(descriptor:inputWeights:hiddenWeights:peepholeWeights:biases:)
func NewCLSTMLayerWithDescriptorInputWeightsHiddenWeightsPeepholeWeightsBiases(descriptor IMLCLSTMDescriptor, inputWeights []CTensor, hiddenWeights []CTensor, peepholeWeights []CTensor, biases []CTensor) CLSTMLayer {
	rv := objc.Send[CLSTMLayer](objc.ID(getCLSTMLayerClass().class), objc.Sel("layerWithDescriptor:inputWeights:hiddenWeights:peepholeWeights:biases:"), descriptor, inputWeights, hiddenWeights, peepholeWeights, biases)
	return rv
}/* debug [class_init_methods/constructor]: NewCLSTMLayerWithDescriptorInputWeightsHiddenWeightsPeepholeWeightsBiases */


// Creates an LSTM layer using the descriptor, weights, biases, gate activations, and output result activation that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/init(descriptor:inputWeights:hiddenWeights:peepholeWeights:biases:gateActivations:outputResultActivation:)
func NewCLSTMLayerWithDescriptorInputWeightsHiddenWeightsPeepholeWeightsBiasesGateActivationsOutputResultActivation(descriptor IMLCLSTMDescriptor, inputWeights []CTensor, hiddenWeights []CTensor, peepholeWeights []CTensor, biases []CTensor, gateActivations []CActivationDescriptor, outputResultActivation IMLCActivationDescriptor) CLSTMLayer {
	rv := objc.Send[CLSTMLayer](objc.ID(getCLSTMLayerClass().class), objc.Sel("layerWithDescriptor:inputWeights:hiddenWeights:peepholeWeights:biases:gateActivations:outputResultActivation:"), descriptor, inputWeights, hiddenWeights, peepholeWeights, biases, gateActivations, outputResultActivation)
	return rv
}/* debug [class_init_methods/constructor]: NewCLSTMLayerWithDescriptorInputWeightsHiddenWeightsPeepholeWeightsBiasesGateActivationsOutputResultActivation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CLSTMLayer */

// Creates an LSTM layer with the descriptor, input and hidden weights, and biases you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/init(descriptor:inputWeights:hiddenWeights:biases:)
func (sc _CLSTMLayerClass) LayerWithDescriptorInputWeightsHiddenWeightsBiases(descriptor IMLCLSTMDescriptor, inputWeights []CTensor, hiddenWeights []CTensor, biases []CTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("layerWithDescriptor:inputWeights:hiddenWeights:biases:"), descriptor, inputWeights, hiddenWeights, biases)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDescriptorInputWeightsHiddenWeightsBiases) */


// Creates an LSTM layer with the descriptor, weights, and biases you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/init(descriptor:inputWeights:hiddenWeights:peepholeWeights:biases:)
func (sc _CLSTMLayerClass) LayerWithDescriptorInputWeightsHiddenWeightsPeepholeWeightsBiases(descriptor IMLCLSTMDescriptor, inputWeights []CTensor, hiddenWeights []CTensor, peepholeWeights []CTensor, biases []CTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("layerWithDescriptor:inputWeights:hiddenWeights:peepholeWeights:biases:"), descriptor, inputWeights, hiddenWeights, peepholeWeights, biases)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDescriptorInputWeightsHiddenWeightsPeepholeWeightsBiases) */


// Creates an LSTM layer using the descriptor, weights, biases, gate activations, and output result activation that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/init(descriptor:inputWeights:hiddenWeights:peepholeWeights:biases:gateActivations:outputResultActivation:)
func (sc _CLSTMLayerClass) LayerWithDescriptorInputWeightsHiddenWeightsPeepholeWeightsBiasesGateActivationsOutputResultActivation(descriptor IMLCLSTMDescriptor, inputWeights []CTensor, hiddenWeights []CTensor, peepholeWeights []CTensor, biases []CTensor, gateActivations []CActivationDescriptor, outputResultActivation IMLCActivationDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("layerWithDescriptor:inputWeights:hiddenWeights:peepholeWeights:biases:gateActivations:outputResultActivation:"), descriptor, inputWeights, hiddenWeights, peepholeWeights, biases, gateActivations, outputResultActivation)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDescriptorInputWeightsHiddenWeightsPeepholeWeightsBiasesGateActivationsOutputResultActivation) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CLSTMLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CLSTMLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CLSTMLayer */

// The array of tensors that describe the bias terms you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/biases
func (s_ CLSTMLayer) Biases() []CTensor {
	rv := objc.Send[[]CTensor](s_.ID, objc.Sel("biases"))
	return rv
}/* debug [instance_properties/getter]: biases */


// The biases tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/biasesParameters
func (s_ CLSTMLayer) BiasesParameters() []CTensorParameter {
	rv := objc.Send[[]CTensorParameter](s_.ID, objc.Sel("biasesParameters"))
	return rv
}/* debug [instance_properties/getter]: biasesParameters */


// The configuration object you use to create the LSTM layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/descriptor
func (s_ CLSTMLayer) Descriptor() IMLCLSTMDescriptor {
	rv := objc.Send[CLSTMDescriptor](s_.ID, objc.Sel("descriptor"))
	return rv
}/* debug [instance_properties/getter]: descriptor */


// The array of gate activations you use for input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/gateActivations
func (s_ CLSTMLayer) GateActivations() []CActivationDescriptor {
	rv := objc.Send[[]CActivationDescriptor](s_.ID, objc.Sel("gateActivations"))
	return rv
}/* debug [instance_properties/getter]: gateActivations */


// The array of tensors that describe the hidden weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/hiddenWeights
func (s_ CLSTMLayer) HiddenWeights() []CTensor {
	rv := objc.Send[[]CTensor](s_.ID, objc.Sel("hiddenWeights"))
	return rv
}/* debug [instance_properties/getter]: hiddenWeights */


// The hidden weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/hiddenWeightsParameters
func (s_ CLSTMLayer) HiddenWeightsParameters() []CTensorParameter {
	rv := objc.Send[[]CTensorParameter](s_.ID, objc.Sel("hiddenWeightsParameters"))
	return rv
}/* debug [instance_properties/getter]: hiddenWeightsParameters */


// The array of tensors that describe the input weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/inputWeights
func (s_ CLSTMLayer) InputWeights() []CTensor {
	rv := objc.Send[[]CTensor](s_.ID, objc.Sel("inputWeights"))
	return rv
}/* debug [instance_properties/getter]: inputWeights */


// The input weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/inputWeightsParameters
func (s_ CLSTMLayer) InputWeightsParameters() []CTensorParameter {
	rv := objc.Send[[]CTensorParameter](s_.ID, objc.Sel("inputWeightsParameters"))
	return rv
}/* debug [instance_properties/getter]: inputWeightsParameters */


// The output activation descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/outputResultActivation
func (s_ CLSTMLayer) OutputResultActivation() IMLCActivationDescriptor {
	rv := objc.Send[CActivationDescriptor](s_.ID, objc.Sel("outputResultActivation"))
	return rv
}/* debug [instance_properties/getter]: outputResultActivation */


// The array of tensors that describe the peephole weights you use for the input, hidden, cell, and output gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/peepholeWeights
func (s_ CLSTMLayer) PeepholeWeights() []CTensor {
	rv := objc.Send[[]CTensor](s_.ID, objc.Sel("peepholeWeights"))
	return rv
}/* debug [instance_properties/getter]: peepholeWeights */


// The peephole weights tensor parameters you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer/peepholeWeightsParameters
func (s_ CLSTMLayer) PeepholeWeightsParameters() []CTensorParameter {
	rv := objc.Send[[]CTensorParameter](s_.ID, objc.Sel("peepholeWeightsParameters"))
	return rv
}/* debug [instance_properties/getter]: peepholeWeightsParameters */


// The number of recurrent layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmdescriptor/layercount
func (s_ CLSTMLayer) LayerCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("layerCount"))
	return rv
}/* debug [instance_properties/getter]: layerCount */


// The number of recurrent layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclstmdescriptor/layercount
func (s_ CLSTMLayer) SetLayerCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLayerCount:"), value)
}/* debug [instance_properties/setter]: layerCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCLSTMLayer */


