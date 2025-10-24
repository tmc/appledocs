// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [MatrixNeuron] class.
type IMatrixNeuron interface {
	IMatrixUnaryKernel
	

	// properties:
	SourceInputFeatureChannels() objectivec.IObject
	SetSourceInputFeatureChannels(value objectivec.IObject)
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	SourceNumberOfFeatureVectors() objectivec.IObject
	SetSourceNumberOfFeatureVectors(value objectivec.IObject)


	

	// methods:
	NeuronParameterA()
	NeuronParameterB()
	NeuronType()
	SetNeuronType()
	NeuronParameterC()
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferInputMatrixBiasVectorResultMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, biasVector IVector, resultMatrix IMatrix)
	SetNeuronToPReLUWithParametersA(A objc.IObject /* cross-framework: NSData */)
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32)


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935600-initwithcoder
func NewMatrixNeuronWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixNeuron {
	instance := getMatrixNeuronClass().Alloc()
	rv := objc.Send[MatrixNeuron](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935603-initwithdevice
func NewMatrixNeuronWithDevice(device unsafe.Pointer) MatrixNeuron {
	instance := getMatrixNeuronClass().Alloc()
	rv := objc.Send[MatrixNeuron](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935583-neuronparametera
func (m_ MatrixNeuron) NeuronParameterA() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterA"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935585-neuronparameterb
func (m_ MatrixNeuron) NeuronParameterB() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterB"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935587-neurontype
func (m_ MatrixNeuron) NeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronType"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935590-setneurontype
func (m_ MatrixNeuron) SetNeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935598-neuronparameterc
func (m_ MatrixNeuron) NeuronParameterC() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterC"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935604-copywithzone
func (m_ MatrixNeuron) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935606-encode
func (m_ MatrixNeuron) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935606-encodetocommandbuffer
func (m_ MatrixNeuron) EncodeToCommandBufferInputMatrixBiasVectorResultMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, biasVector IVector, resultMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:biasVector:resultMatrix:"), commandBuffer, inputMatrix, biasVector, resultMatrix)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/setNeuronToPReLUWithParametersA(_:)
func (m_ MatrixNeuron) SetNeuronToPReLUWithParametersA(A objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronToPReLUWithParametersA:"), A)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/setNeuronType(_:parameterA:parameterB:parameterC:)
func (m_ MatrixNeuron) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935599-sourceinputfeaturechannels
func (m_ MatrixNeuron) SourceInputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935599-sourceinputfeaturechannels
func (m_ MatrixNeuron) SetSourceInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935605-alpha
func (m_ MatrixNeuron) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935605-alpha
func (m_ MatrixNeuron) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935607-sourcenumberoffeaturevectors
func (m_ MatrixNeuron) SourceNumberOfFeatureVectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935607-sourcenumberoffeaturevectors
func (m_ MatrixNeuron) SetSourceNumberOfFeatureVectors(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}







