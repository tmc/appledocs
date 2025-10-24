// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixBatchNormalization */


/* debug [class_header]: Header for MPSMatrixBatchNormalization */
// The class instance for the [MatrixBatchNormalization] class.
var (
	MatrixBatchNormalizationClass     _MatrixBatchNormalizationClass
	MatrixBatchNormalizationClassOnce sync.Once
)

func getMatrixBatchNormalizationClass() _MatrixBatchNormalizationClass {
	MatrixBatchNormalizationClassOnce.Do(func() {
		MatrixBatchNormalizationClass = _MatrixBatchNormalizationClass{objc.GetClass("MPSMatrixBatchNormalization")}
	})
	return MatrixBatchNormalizationClass
}

type _MatrixBatchNormalizationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixBatchNormalization */
// An interface definition for the [MatrixBatchNormalization] class.
type IMatrixBatchNormalization interface {
	IMatrixUnaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixBatchNormalization */
	// properties:
	ComputeStatistics() objectivec.IObject
	SetComputeStatistics(value objectivec.IObject)
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	SourceInputFeatureChannels() objectivec.IObject
	SetSourceInputFeatureChannels(value objectivec.IObject)
	SourceNumberOfFeatureVectors() objectivec.IObject
	SetSourceNumberOfFeatureVectors(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixBatchNormalization */
	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultMatrix IMatrix)
	NeuronParameterA()
	NeuronParameterB()
	NeuronParameterC()
	NeuronType()
	SetNeuronType()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixBatchNormalization */
// Alloc allocates a new instance without initialization.
func (mc _MatrixBatchNormalizationClass) Alloc() MatrixBatchNormalization {
	rv := objc.Send[MatrixBatchNormalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixBatchNormalizationClass) New() MatrixBatchNormalization {
	rv := objc.Send[MatrixBatchNormalization](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixBatchNormalization) Init() MatrixBatchNormalization {
	rv := objc.Send[MatrixBatchNormalization](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixBatchNormalization) Autorelease() MatrixBatchNormalization {
	rv := objc.Send[MatrixBatchNormalization](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixBatchNormalization creates a new MatrixBatchNormalization instance.
func NewMatrixBatchNormalization() MatrixBatchNormalization {
	return getMatrixBatchNormalizationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixBatchNormalization */
// A batch normalization kernel that operates on matrices.


// A batch normalization kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization
type MatrixBatchNormalization struct {
	MatrixUnaryKernel
}

// MatrixBatchNormalizationFrom constructs a [MatrixBatchNormalization] from an unsafe.Pointer.
//
// A batch normalization kernel that operates on matrices.
func MatrixBatchNormalizationFrom(ptr unsafe.Pointer) MatrixBatchNormalization {
	return MatrixBatchNormalization{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixBatchNormalization */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980734-initwithcoder
func NewMatrixBatchNormalizationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixBatchNormalization {
	instance := getMatrixBatchNormalizationClass().Alloc()
	rv := objc.Send[MatrixBatchNormalization](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixBatchNormalizationWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980735-initwithdevice
func NewMatrixBatchNormalizationWithDevice(device unsafe.Pointer) MatrixBatchNormalization {
	instance := getMatrixBatchNormalizationClass().Alloc()
	rv := objc.Send[MatrixBatchNormalization](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixBatchNormalizationWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixBatchNormalization */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixBatchNormalization */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixBatchNormalization */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980731-copywithzone
func (m_ MatrixBatchNormalization) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980732-encode
func (m_ MatrixBatchNormalization) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980732-encodetocommandbuffer
func (m_ MatrixBatchNormalization) EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultMatrix:"), commandBuffer, inputMatrix, meanVector, varianceVector, gammaVector, betaVector, resultMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980736-neuronparametera
func (m_ MatrixBatchNormalization) NeuronParameterA() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterA"))
}/* debug [instance_methods/method]: NeuronParameterA */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980737-neuronparameterb
func (m_ MatrixBatchNormalization) NeuronParameterB() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterB"))
}/* debug [instance_methods/method]: NeuronParameterB */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980738-neuronparameterc
func (m_ MatrixBatchNormalization) NeuronParameterC() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronParameterC"))
}/* debug [instance_methods/method]: NeuronParameterC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980739-neurontype
func (m_ MatrixBatchNormalization) NeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronType"))
}/* debug [instance_methods/method]: NeuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980740-setneurontype
func (m_ MatrixBatchNormalization) SetNeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType"))
}/* debug [instance_methods/method]: SetNeuronType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixBatchNormalization */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980730-computestatistics
func (m_ MatrixBatchNormalization) ComputeStatistics() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("computeStatistics"))
	return rv
}/* debug [instance_properties/getter]: computeStatistics */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980730-computestatistics
func (m_ MatrixBatchNormalization) SetComputeStatistics(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setComputeStatistics:"), value)
}/* debug [instance_properties/setter]: computeStatistics */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980733-epsilon
func (m_ MatrixBatchNormalization) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980733-epsilon
func (m_ MatrixBatchNormalization) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980741-sourceinputfeaturechannels
func (m_ MatrixBatchNormalization) SourceInputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceInputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980741-sourceinputfeaturechannels
func (m_ MatrixBatchNormalization) SetSourceInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceInputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980742-sourcenumberoffeaturevectors
func (m_ MatrixBatchNormalization) SourceNumberOfFeatureVectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}/* debug [instance_properties/getter]: sourceNumberOfFeatureVectors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980742-sourcenumberoffeaturevectors
func (m_ MatrixBatchNormalization) SetSourceNumberOfFeatureVectors(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}/* debug [instance_properties/setter]: sourceNumberOfFeatureVectors */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixBatchNormalization */


