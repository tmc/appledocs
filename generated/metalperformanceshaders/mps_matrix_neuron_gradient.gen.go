// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixNeuronGradient */


/* debug [class_header]: Header for MPSMatrixNeuronGradient */
// The class instance for the [MatrixNeuronGradient] class.
var (
	MatrixNeuronGradientClass     _MatrixNeuronGradientClass
	MatrixNeuronGradientClassOnce sync.Once
)

func getMatrixNeuronGradientClass() _MatrixNeuronGradientClass {
	MatrixNeuronGradientClassOnce.Do(func() {
		MatrixNeuronGradientClass = _MatrixNeuronGradientClass{objc.GetClass("MPSMatrixNeuronGradient")}
	})
	return MatrixNeuronGradientClass
}

type _MatrixNeuronGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixNeuronGradient */
// An interface definition for the [MatrixNeuronGradient] class.
type IMatrixNeuronGradient interface {
	IMatrixBinaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixNeuronGradient */
	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	SourceInputFeatureChannels() objectivec.IObject
	SetSourceInputFeatureChannels(value objectivec.IObject)
	SourceNumberOfFeatureVectors() objectivec.IObject
	SetSourceNumberOfFeatureVectors(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixNeuronGradient */
	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, inputMatrix IMatrix, biasVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForBiasVector IVector)
	NeuronParameterA()
	NeuronParameterB()
	NeuronParameterC()
	NeuronType()
	SetNeuronToPReLUWithParametersA()
	SetNeuronType()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixNeuronGradient */
// Alloc allocates a new instance without initialization.
func (mc _MatrixNeuronGradientClass) Alloc() MatrixNeuronGradient {
	rv := objc.Send[MatrixNeuronGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixNeuronGradientClass) New() MatrixNeuronGradient {
	rv := objc.Send[MatrixNeuronGradient](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixNeuronGradient) Init() MatrixNeuronGradient {
	rv := objc.Send[MatrixNeuronGradient](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixNeuronGradient) Autorelease() MatrixNeuronGradient {
	rv := objc.Send[MatrixNeuronGradient](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixNeuronGradient creates a new MatrixNeuronGradient instance.
func NewMatrixNeuronGradient() MatrixNeuronGradient {
	return getMatrixNeuronGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixNeuronGradient */
// A gradient neuron activation kernel that operates on matrices.


// A gradient neuron activation kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient
type MatrixNeuronGradient struct {
	MatrixBinaryKernel
}

// MatrixNeuronGradientFrom constructs a [MatrixNeuronGradient] from an unsafe.Pointer.
//
// A gradient neuron activation kernel that operates on matrices.
func MatrixNeuronGradientFrom(ptr unsafe.Pointer) MatrixNeuronGradient {
	return MatrixNeuronGradient{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixNeuronGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966676-initwithcoder
func NewMatrixNeuronGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixNeuronGradient {
	instance := getMatrixNeuronGradientClass().Alloc()
	rv := objc.Send[MatrixNeuronGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixNeuronGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966677-initwithdevice
func NewMatrixNeuronGradientWithDevice(device unsafe.Pointer) MatrixNeuronGradient {
	instance := getMatrixNeuronGradientClass().Alloc()
	rv := objc.Send[MatrixNeuronGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixNeuronGradientWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixNeuronGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixNeuronGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixNeuronGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966674-copywithzone
func (m_ MatrixNeuronGradient) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966675-encode
func (m_ MatrixNeuronGradient) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966675-encodetocommandbuffer
func (m_ MatrixNeuronGradient) EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, inputMatrix IMatrix, biasVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForBiasVector IVector) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:gradientMatrix:inputMatrix:biasVector:resultGradientForDataMatrix:resultGradientForBiasVector:"), commandBuffer, gradientMatrix, inputMatrix, biasVector, resultGradientForDataMatrix, resultGradientForBiasVector)
}/* debug [instance_methods/method]: EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966678-neuronparametera
func (m_ MatrixNeuronGradient) NeuronParameterA() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterA"))
}/* debug [instance_methods/method]: NeuronParameterA */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966679-neuronparameterb
func (m_ MatrixNeuronGradient) NeuronParameterB() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterB"))
}/* debug [instance_methods/method]: NeuronParameterB */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966680-neuronparameterc
func (m_ MatrixNeuronGradient) NeuronParameterC() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterC"))
}/* debug [instance_methods/method]: NeuronParameterC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966681-neurontype
func (m_ MatrixNeuronGradient) NeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronType"))
}/* debug [instance_methods/method]: NeuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966682-setneurontopreluwithparametersa
func (m_ MatrixNeuronGradient) SetNeuronToPReLUWithParametersA() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronToPReLUWithParametersA"))
}/* debug [instance_methods/method]: SetNeuronToPReLUWithParametersA */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966683-setneurontype
func (m_ MatrixNeuronGradient) SetNeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType"))
}/* debug [instance_methods/method]: SetNeuronType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixNeuronGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966673-alpha
func (m_ MatrixNeuronGradient) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966673-alpha
func (m_ MatrixNeuronGradient) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966684-sourceinputfeaturechannels
func (m_ MatrixNeuronGradient) SourceInputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceInputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966684-sourceinputfeaturechannels
func (m_ MatrixNeuronGradient) SetSourceInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceInputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966685-sourcenumberoffeaturevectors
func (m_ MatrixNeuronGradient) SourceNumberOfFeatureVectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}/* debug [instance_properties/getter]: sourceNumberOfFeatureVectors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966685-sourcenumberoffeaturevectors
func (m_ MatrixNeuronGradient) SetSourceNumberOfFeatureVectors(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}/* debug [instance_properties/setter]: sourceNumberOfFeatureVectors */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixNeuronGradient */


