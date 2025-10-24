// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionWeightsAndBiasesState */


/* debug [class_header]: Header for MPSCNNConvolutionWeightsAndBiasesState */
// The class instance for the [CNNConvolutionWeightsAndBiasesState] class.
var (
	CNNConvolutionWeightsAndBiasesStateClass     _CNNConvolutionWeightsAndBiasesStateClass
	CNNConvolutionWeightsAndBiasesStateClassOnce sync.Once
)

func getCNNConvolutionWeightsAndBiasesStateClass() _CNNConvolutionWeightsAndBiasesStateClass {
	CNNConvolutionWeightsAndBiasesStateClassOnce.Do(func() {
		CNNConvolutionWeightsAndBiasesStateClass = _CNNConvolutionWeightsAndBiasesStateClass{objc.GetClass("MPSCNNConvolutionWeightsAndBiasesState")}
	})
	return CNNConvolutionWeightsAndBiasesStateClass
}

type _CNNConvolutionWeightsAndBiasesStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionWeightsAndBiasesState */
// An interface definition for the [CNNConvolutionWeightsAndBiasesState] class.
type ICNNConvolutionWeightsAndBiasesState interface {
	IState
	
/* debug [class_interface_properties]: Properties for CNNConvolutionWeightsAndBiasesState */
	// properties:
	Biases() Buffer get /* not a class type */
	SetBiases(value Buffer get /* not a class type */)
	Weights() Buffer get /* not a class type */
	SetWeights(value Buffer get /* not a class type */)
	BiasesOffset() objectivec.IObject
	SetBiasesOffset(value objectivec.IObject)
	WeightsOffset() objectivec.IObject
	SetWeightsOffset(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionWeightsAndBiasesState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionWeightsAndBiasesState */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionWeightsAndBiasesStateClass) Alloc() CNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[CNNConvolutionWeightsAndBiasesState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionWeightsAndBiasesStateClass) New() CNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[CNNConvolutionWeightsAndBiasesState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionWeightsAndBiasesState) Init() CNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[CNNConvolutionWeightsAndBiasesState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionWeightsAndBiasesState) Autorelease() CNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[CNNConvolutionWeightsAndBiasesState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionWeightsAndBiasesState creates a new CNNConvolutionWeightsAndBiasesState instance.
func NewCNNConvolutionWeightsAndBiasesState() CNNConvolutionWeightsAndBiasesState {
	return getCNNConvolutionWeightsAndBiasesStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionWeightsAndBiasesState */
// A class that stores weights and biases.


// A class that stores weights and biases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState
type CNNConvolutionWeightsAndBiasesState struct {
	State
}

// CNNConvolutionWeightsAndBiasesStateFrom constructs a [CNNConvolutionWeightsAndBiasesState] from an unsafe.Pointer.
//
// A class that stores weights and biases.
func CNNConvolutionWeightsAndBiasesStateFrom(ptr unsafe.Pointer) CNNConvolutionWeightsAndBiasesState {
	return CNNConvolutionWeightsAndBiasesState{
		State: StateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionWeightsAndBiasesState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953004-initwithdevice
func NewCNNConvolutionWeightsAndBiasesStateWithDeviceCnnConvolutionDescriptor(device unsafe.Pointer, descriptor ICNNConvolutionDescriptor) CNNConvolutionWeightsAndBiasesState {
	instance := getCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[CNNConvolutionWeightsAndBiasesState](instance.ID, objc.Sel("initWithDevice:cnnConvolutionDescriptor:"), device, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionWeightsAndBiasesStateWithDeviceCnnConvolutionDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953008-initwithweights
func NewCNNConvolutionWeightsAndBiasesStateWithWeightsBiases(weights unsafe.Pointer, biases unsafe.Pointer) CNNConvolutionWeightsAndBiasesState {
	instance := getCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[CNNConvolutionWeightsAndBiasesState](instance.ID, objc.Sel("initWithWeights:biases:"), weights, biases)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionWeightsAndBiasesStateWithWeightsBiases */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/3325843-initwithweights
func NewCNNConvolutionWeightsAndBiasesStateWithWeightsWeightsOffsetBiasesBiasesOffsetCnnConvolutionDescriptor(weights unsafe.Pointer, weightsOffset uint, biases unsafe.Pointer, biasesOffset uint, descriptor ICNNConvolutionDescriptor) CNNConvolutionWeightsAndBiasesState {
	instance := getCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[CNNConvolutionWeightsAndBiasesState](instance.ID, objc.Sel("initWithWeights:weightsOffset:biases:biasesOffset:cnnConvolutionDescriptor:"), weights, weightsOffset, biases, biasesOffset, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionWeightsAndBiasesStateWithWeightsWeightsOffsetBiasesBiasesOffsetCnnConvolutionDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionWeightsAndBiasesState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953005-temporarycnnconvolutionweightsan
func (cc _CNNConvolutionWeightsAndBiasesStateClass) TemporaryCNNConvolutionWeightsAndBiasesState() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("temporaryCNNConvolutionWeightsAndBiasesState"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TemporaryCNNConvolutionWeightsAndBiasesState) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/temporaryCNNConvolutionWeightsAndBiasesState(with:cnnConvolutionDescriptor:)
func (cc _CNNConvolutionWeightsAndBiasesStateClass) TemporaryCNNConvolutionWeightsAndBiasesStateWithCommandBufferCnnConvolutionDescriptor(commandBuffer unsafe.Pointer, descriptor IMPSCNNConvolutionDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("temporaryCNNConvolutionWeightsAndBiasesStateWithCommandBuffer:cnnConvolutionDescriptor:"), commandBuffer, descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TemporaryCNNConvolutionWeightsAndBiasesStateWithCommandBufferCnnConvolutionDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionWeightsAndBiasesState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionWeightsAndBiasesState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionWeightsAndBiasesState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953002-biases
func (c_ CNNConvolutionWeightsAndBiasesState) Biases() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("biases"))
	return rv
}/* debug [instance_properties/getter]: biases */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953002-biases
func (c_ CNNConvolutionWeightsAndBiasesState) SetBiases(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBiases:"), value)
}/* debug [instance_properties/setter]: biases */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953006-weights
func (c_ CNNConvolutionWeightsAndBiasesState) Weights() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("weights"))
	return rv
}/* debug [instance_properties/getter]: weights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953006-weights
func (c_ CNNConvolutionWeightsAndBiasesState) SetWeights(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeights:"), value)
}/* debug [instance_properties/setter]: weights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/3325842-biasesoffset
func (c_ CNNConvolutionWeightsAndBiasesState) BiasesOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("biasesOffset"))
	return rv
}/* debug [instance_properties/getter]: biasesOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/3325842-biasesoffset
func (c_ CNNConvolutionWeightsAndBiasesState) SetBiasesOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBiasesOffset:"), value)
}/* debug [instance_properties/setter]: biasesOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/3325844-weightsoffset
func (c_ CNNConvolutionWeightsAndBiasesState) WeightsOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("weightsOffset"))
	return rv
}/* debug [instance_properties/getter]: weightsOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/3325844-weightsoffset
func (c_ CNNConvolutionWeightsAndBiasesState) SetWeightsOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeightsOffset:"), value)
}/* debug [instance_properties/setter]: weightsOffset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionWeightsAndBiasesState */


