// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixNeuron */


/* debug [class_header]: Header for MPSMatrixNeuron */
// The class instance for the [MatrixNeuron] class.
var (
	MatrixNeuronClass     _MatrixNeuronClass
	MatrixNeuronClassOnce sync.Once
)

func getMatrixNeuronClass() _MatrixNeuronClass {
	MatrixNeuronClassOnce.Do(func() {
		MatrixNeuronClass = _MatrixNeuronClass{objc.GetClass("MPSMatrixNeuron")}
	})
	return MatrixNeuronClass
}

type _MatrixNeuronClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixNeuron */
// An interface definition for the [MatrixNeuron] class.
type IMatrixNeuron interface {
	IMatrixUnaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixNeuron */
	// properties:
	SourceInputFeatureChannels() objectivec.IObject
	SetSourceInputFeatureChannels(value objectivec.IObject)
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	SourceNumberOfFeatureVectors() objectivec.IObject
	SetSourceNumberOfFeatureVectors(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixNeuron */
	// methods:
	NeuronParameterA()
	NeuronParameterB()
	NeuronType()
	SetNeuronType()
	NeuronParameterC()
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferInputMatrixBiasVectorResultMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, biasVector IVector, resultMatrix IMatrix)
	SetNeuronToPReLUWithParametersA()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixNeuron */
// Alloc allocates a new instance without initialization.
func (mc _MatrixNeuronClass) Alloc() MatrixNeuron {
	rv := objc.Send[MatrixNeuron](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixNeuronClass) New() MatrixNeuron {
	rv := objc.Send[MatrixNeuron](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixNeuron) Init() MatrixNeuron {
	rv := objc.Send[MatrixNeuron](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixNeuron) Autorelease() MatrixNeuron {
	rv := objc.Send[MatrixNeuron](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixNeuron creates a new MatrixNeuron instance.
func NewMatrixNeuron() MatrixNeuron {
	return getMatrixNeuronClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixNeuron */
// A neuron activation kernel that operates on matrices.


// A neuron activation kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron
type MatrixNeuron struct {
	MatrixUnaryKernel
}

// MatrixNeuronFrom constructs a [MatrixNeuron] from an unsafe.Pointer.
//
// A neuron activation kernel that operates on matrices.
func MatrixNeuronFrom(ptr unsafe.Pointer) MatrixNeuron {
	return MatrixNeuron{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixNeuron */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935600-initwithcoder
func NewMatrixNeuronWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) MatrixNeuron {
	instance := getMatrixNeuronClass().Alloc()
	rv := objc.Send[MatrixNeuron](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixNeuronWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935603-initwithdevice
func NewMatrixNeuronWithDevice(device unsafe.Pointer) MatrixNeuron {
	instance := getMatrixNeuronClass().Alloc()
	rv := objc.Send[MatrixNeuron](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixNeuronWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixNeuron */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixNeuron */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixNeuron */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935583-neuronparametera
func (m_ MatrixNeuron) NeuronParameterA() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterA"))
}/* debug [instance_methods/method]: NeuronParameterA */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935585-neuronparameterb
func (m_ MatrixNeuron) NeuronParameterB() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterB"))
}/* debug [instance_methods/method]: NeuronParameterB */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935587-neurontype
func (m_ MatrixNeuron) NeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronType"))
}/* debug [instance_methods/method]: NeuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935590-setneurontype
func (m_ MatrixNeuron) SetNeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType"))
}/* debug [instance_methods/method]: SetNeuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935598-neuronparameterc
func (m_ MatrixNeuron) NeuronParameterC() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterC"))
}/* debug [instance_methods/method]: NeuronParameterC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935604-copywithzone
func (m_ MatrixNeuron) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935606-encode
func (m_ MatrixNeuron) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935606-encodetocommandbuffer
func (m_ MatrixNeuron) EncodeToCommandBufferInputMatrixBiasVectorResultMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, biasVector IVector, resultMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:biasVector:resultMatrix:"), commandBuffer, inputMatrix, biasVector, resultMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputMatrixBiasVectorResultMatrix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935610-setneurontopreluwithparametersa
func (m_ MatrixNeuron) SetNeuronToPReLUWithParametersA() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronToPReLUWithParametersA"))
}/* debug [instance_methods/method]: SetNeuronToPReLUWithParametersA */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixNeuron */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935599-sourceinputfeaturechannels
func (m_ MatrixNeuron) SourceInputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceInputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935599-sourceinputfeaturechannels
func (m_ MatrixNeuron) SetSourceInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceInputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935605-alpha
func (m_ MatrixNeuron) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935605-alpha
func (m_ MatrixNeuron) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935607-sourcenumberoffeaturevectors
func (m_ MatrixNeuron) SourceNumberOfFeatureVectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}/* debug [instance_properties/getter]: sourceNumberOfFeatureVectors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935607-sourcenumberoffeaturevectors
func (m_ MatrixNeuron) SetSourceNumberOfFeatureVectors(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}/* debug [instance_properties/setter]: sourceNumberOfFeatureVectors */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixNeuron */


