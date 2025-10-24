// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [MatrixNeuronGradient] class.
type IMatrixNeuronGradient interface {
	IMatrixBinaryKernel
	

	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	SourceInputFeatureChannels() objectivec.IObject
	SetSourceInputFeatureChannels(value objectivec.IObject)
	SourceNumberOfFeatureVectors() objectivec.IObject
	SetSourceNumberOfFeatureVectors(value objectivec.IObject)


	

	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, inputMatrix IMatrix, biasVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForBiasVector IVector)
	NeuronParameterA()
	NeuronParameterB()
	NeuronParameterC()
	NeuronType()
	SetNeuronType()
	SetNeuronToPReLUWithParametersA(A objc.IObject /* cross-framework: NSData */)
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32)


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966676-initwithcoder
func NewMatrixNeuronGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixNeuronGradient {
	instance := getMatrixNeuronGradientClass().Alloc()
	rv := objc.Send[MatrixNeuronGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966677-initwithdevice
func NewMatrixNeuronGradientWithDevice(device unsafe.Pointer) MatrixNeuronGradient {
	instance := getMatrixNeuronGradientClass().Alloc()
	rv := objc.Send[MatrixNeuronGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966674-copywithzone
func (m_ MatrixNeuronGradient) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966675-encode
func (m_ MatrixNeuronGradient) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966675-encodetocommandbuffer
func (m_ MatrixNeuronGradient) EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, inputMatrix IMatrix, biasVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForBiasVector IVector) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:gradientMatrix:inputMatrix:biasVector:resultGradientForDataMatrix:resultGradientForBiasVector:"), commandBuffer, gradientMatrix, inputMatrix, biasVector, resultGradientForDataMatrix, resultGradientForBiasVector)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966678-neuronparametera
func (m_ MatrixNeuronGradient) NeuronParameterA() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterA"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966679-neuronparameterb
func (m_ MatrixNeuronGradient) NeuronParameterB() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterB"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966680-neuronparameterc
func (m_ MatrixNeuronGradient) NeuronParameterC() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterC"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966681-neurontype
func (m_ MatrixNeuronGradient) NeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronType"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966683-setneurontype
func (m_ MatrixNeuronGradient) SetNeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/setNeuronToPReLUWithParametersA(_:)
func (m_ MatrixNeuronGradient) SetNeuronToPReLUWithParametersA(A objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronToPReLUWithParametersA:"), A)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/setNeuronType(_:parameterA:parameterB:parameterC:)
func (m_ MatrixNeuronGradient) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966673-alpha
func (m_ MatrixNeuronGradient) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966673-alpha
func (m_ MatrixNeuronGradient) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966684-sourceinputfeaturechannels
func (m_ MatrixNeuronGradient) SourceInputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966684-sourceinputfeaturechannels
func (m_ MatrixNeuronGradient) SetSourceInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966685-sourcenumberoffeaturevectors
func (m_ MatrixNeuronGradient) SourceNumberOfFeatureVectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966685-sourcenumberoffeaturevectors
func (m_ MatrixNeuronGradient) SetSourceNumberOfFeatureVectors(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}







