// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixBatchNormalizationGradient] class.
var (
	MatrixBatchNormalizationGradientClass     _MatrixBatchNormalizationGradientClass
	MatrixBatchNormalizationGradientClassOnce sync.Once
)

func getMatrixBatchNormalizationGradientClass() _MatrixBatchNormalizationGradientClass {
	MatrixBatchNormalizationGradientClassOnce.Do(func() {
		MatrixBatchNormalizationGradientClass = _MatrixBatchNormalizationGradientClass{objc.GetClass("MPSMatrixBatchNormalizationGradient")}
	})
	return MatrixBatchNormalizationGradientClass
}

type _MatrixBatchNormalizationGradientClass struct {
	class objc.Class
}





// An interface definition for the [MatrixBatchNormalizationGradient] class.
type IMatrixBatchNormalizationGradient interface {
	IMatrixBinaryKernel
	

	// properties:
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	SourceInputFeatureChannels() objectivec.IObject
	SetSourceInputFeatureChannels(value objectivec.IObject)
	SourceNumberOfFeatureVectors() objectivec.IObject
	SetSourceNumberOfFeatureVectors(value objectivec.IObject)


	

	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferGradientMatrixInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultGradientForDataMatrixResultGradientForGammaVectorResultGradientForBetaVector(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForGammaVector IVector, resultGradientForBetaVector IVector)
	NeuronParameterA()
	NeuronParameterB()
	NeuronParameterC()
	NeuronType()
	SetNeuronType()
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32)


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixBatchNormalizationGradientClass) Alloc() MatrixBatchNormalizationGradient {
	rv := objc.Send[MatrixBatchNormalizationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixBatchNormalizationGradientClass) New() MatrixBatchNormalizationGradient {
	rv := objc.Send[MatrixBatchNormalizationGradient](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixBatchNormalizationGradient) Init() MatrixBatchNormalizationGradient {
	rv := objc.Send[MatrixBatchNormalizationGradient](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixBatchNormalizationGradient) Autorelease() MatrixBatchNormalizationGradient {
	rv := objc.Send[MatrixBatchNormalizationGradient](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixBatchNormalizationGradient creates a new MatrixBatchNormalizationGradient instance.
func NewMatrixBatchNormalizationGradient() MatrixBatchNormalizationGradient {
	return getMatrixBatchNormalizationGradientClass().New()
}





// A batch normalization gradient kernel that operates on matrices.


// A batch normalization gradient kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient
type MatrixBatchNormalizationGradient struct {
	MatrixBinaryKernel
}

// MatrixBatchNormalizationGradientFrom constructs a [MatrixBatchNormalizationGradient] from an unsafe.Pointer.
//
// A batch normalization gradient kernel that operates on matrices.
func MatrixBatchNormalizationGradientFrom(ptr unsafe.Pointer) MatrixBatchNormalizationGradient {
	return MatrixBatchNormalizationGradient{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980747-initwithcoder
func NewMatrixBatchNormalizationGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixBatchNormalizationGradient {
	instance := getMatrixBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[MatrixBatchNormalizationGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980748-initwithdevice
func NewMatrixBatchNormalizationGradientWithDevice(device unsafe.Pointer) MatrixBatchNormalizationGradient {
	instance := getMatrixBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[MatrixBatchNormalizationGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980744-copywithzone
func (m_ MatrixBatchNormalizationGradient) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980745-encode
func (m_ MatrixBatchNormalizationGradient) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980745-encodetocommandbuffer
func (m_ MatrixBatchNormalizationGradient) EncodeToCommandBufferGradientMatrixInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultGradientForDataMatrixResultGradientForGammaVectorResultGradientForBetaVector(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForGammaVector IVector, resultGradientForBetaVector IVector) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:gradientMatrix:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultGradientForDataMatrix:resultGradientForGammaVector:resultGradientForBetaVector:"), commandBuffer, gradientMatrix, inputMatrix, meanVector, varianceVector, gammaVector, betaVector, resultGradientForDataMatrix, resultGradientForGammaVector, resultGradientForBetaVector)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980749-neuronparametera
func (m_ MatrixBatchNormalizationGradient) NeuronParameterA() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterA"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980750-neuronparameterb
func (m_ MatrixBatchNormalizationGradient) NeuronParameterB() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterB"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980751-neuronparameterc
func (m_ MatrixBatchNormalizationGradient) NeuronParameterC() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterC"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980752-neurontype
func (m_ MatrixBatchNormalizationGradient) NeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronType"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980753-setneurontype
func (m_ MatrixBatchNormalizationGradient) SetNeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/setNeuronType(_:parameterA:parameterB:parameterC:)
func (m_ MatrixBatchNormalizationGradient) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980746-epsilon
func (m_ MatrixBatchNormalizationGradient) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980746-epsilon
func (m_ MatrixBatchNormalizationGradient) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpsilon:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980754-sourceinputfeaturechannels
func (m_ MatrixBatchNormalizationGradient) SourceInputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980754-sourceinputfeaturechannels
func (m_ MatrixBatchNormalizationGradient) SetSourceInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980755-sourcenumberoffeaturevectors
func (m_ MatrixBatchNormalizationGradient) SourceNumberOfFeatureVectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980755-sourcenumberoffeaturevectors
func (m_ MatrixBatchNormalizationGradient) SetSourceNumberOfFeatureVectors(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}







