// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixFullyConnected */


/* debug [class_header]: Header for MPSMatrixFullyConnected */
// The class instance for the [MatrixFullyConnected] class.
var (
	MatrixFullyConnectedClass     _MatrixFullyConnectedClass
	MatrixFullyConnectedClassOnce sync.Once
)

func getMatrixFullyConnectedClass() _MatrixFullyConnectedClass {
	MatrixFullyConnectedClassOnce.Do(func() {
		MatrixFullyConnectedClass = _MatrixFullyConnectedClass{objc.GetClass("MPSMatrixFullyConnected")}
	})
	return MatrixFullyConnectedClass
}

type _MatrixFullyConnectedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixFullyConnected */
// An interface definition for the [MatrixFullyConnected] class.
type IMatrixFullyConnected interface {
	IMatrixBinaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixFullyConnected */
	// properties:
	SourceOutputFeatureChannels() objectivec.IObject
	SetSourceOutputFeatureChannels(value objectivec.IObject)
	SourceInputFeatureChannels() objectivec.IObject
	SetSourceInputFeatureChannels(value objectivec.IObject)
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	SourceNumberOfFeatureVectors() objectivec.IObject
	SetSourceNumberOfFeatureVectors(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixFullyConnected */
	// methods:
	NeuronType()
	NeuronParameterB()
	SetNeuronType()
	NeuronParameterC()
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferInputMatrixWeightMatrixBiasVectorResultMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, weightMatrix IMatrix, biasVector IVector, resultMatrix IMatrix)
	NeuronParameterA()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixFullyConnected */
// Alloc allocates a new instance without initialization.
func (mc _MatrixFullyConnectedClass) Alloc() MatrixFullyConnected {
	rv := objc.Send[MatrixFullyConnected](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixFullyConnectedClass) New() MatrixFullyConnected {
	rv := objc.Send[MatrixFullyConnected](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixFullyConnected) Init() MatrixFullyConnected {
	rv := objc.Send[MatrixFullyConnected](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixFullyConnected) Autorelease() MatrixFullyConnected {
	rv := objc.Send[MatrixFullyConnected](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixFullyConnected creates a new MatrixFullyConnected instance.
func NewMatrixFullyConnected() MatrixFullyConnected {
	return getMatrixFullyConnectedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixFullyConnected */
// A kernel for applying a fully connected neural network layer.


// A kernel for applying a fully connected neural network layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected
type MatrixFullyConnected struct {
	MatrixBinaryKernel
}

// MatrixFullyConnectedFrom constructs a [MatrixFullyConnected] from an unsafe.Pointer.
//
// A kernel for applying a fully connected neural network layer.
func MatrixFullyConnectedFrom(ptr unsafe.Pointer) MatrixFullyConnected {
	return MatrixFullyConnected{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixFullyConnected */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935611-initwithcoder
func NewMatrixFullyConnectedWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) MatrixFullyConnected {
	instance := getMatrixFullyConnectedClass().Alloc()
	rv := objc.Send[MatrixFullyConnected](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixFullyConnectedWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935584-initwithdevice
func NewMatrixFullyConnectedWithDevice(device unsafe.Pointer) MatrixFullyConnected {
	instance := getMatrixFullyConnectedClass().Alloc()
	rv := objc.Send[MatrixFullyConnected](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixFullyConnectedWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixFullyConnected */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixFullyConnected */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixFullyConnected */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935588-neurontype
func (m_ MatrixFullyConnected) NeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronType"))
}/* debug [instance_methods/method]: NeuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935591-neuronparameterb
func (m_ MatrixFullyConnected) NeuronParameterB() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterB"))
}/* debug [instance_methods/method]: NeuronParameterB */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935593-setneurontype
func (m_ MatrixFullyConnected) SetNeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType"))
}/* debug [instance_methods/method]: SetNeuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935594-neuronparameterc
func (m_ MatrixFullyConnected) NeuronParameterC() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterC"))
}/* debug [instance_methods/method]: NeuronParameterC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935595-copywithzone
func (m_ MatrixFullyConnected) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935596-encode
func (m_ MatrixFullyConnected) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935596-encodetocommandbuffer
func (m_ MatrixFullyConnected) EncodeToCommandBufferInputMatrixWeightMatrixBiasVectorResultMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, weightMatrix IMatrix, biasVector IVector, resultMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:weightMatrix:biasVector:resultMatrix:"), commandBuffer, inputMatrix, weightMatrix, biasVector, resultMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputMatrixWeightMatrixBiasVectorResultMatrix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935602-neuronparametera
func (m_ MatrixFullyConnected) NeuronParameterA() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterA"))
}/* debug [instance_methods/method]: NeuronParameterA */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixFullyConnected */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935592-sourceoutputfeaturechannels
func (m_ MatrixFullyConnected) SourceOutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceOutputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceOutputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935592-sourceoutputfeaturechannels
func (m_ MatrixFullyConnected) SetSourceOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceOutputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935597-sourceinputfeaturechannels
func (m_ MatrixFullyConnected) SourceInputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceInputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935597-sourceinputfeaturechannels
func (m_ MatrixFullyConnected) SetSourceInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceInputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935608-alpha
func (m_ MatrixFullyConnected) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935608-alpha
func (m_ MatrixFullyConnected) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935609-sourcenumberoffeaturevectors
func (m_ MatrixFullyConnected) SourceNumberOfFeatureVectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}/* debug [instance_properties/getter]: sourceNumberOfFeatureVectors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935609-sourcenumberoffeaturevectors
func (m_ MatrixFullyConnected) SetSourceNumberOfFeatureVectors(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}/* debug [instance_properties/setter]: sourceNumberOfFeatureVectors */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixFullyConnected */


