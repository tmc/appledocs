// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixFullyConnectedGradient] class.
var (
	MatrixFullyConnectedGradientClass     _MatrixFullyConnectedGradientClass
	MatrixFullyConnectedGradientClassOnce sync.Once
)

func getMatrixFullyConnectedGradientClass() _MatrixFullyConnectedGradientClass {
	MatrixFullyConnectedGradientClassOnce.Do(func() {
		MatrixFullyConnectedGradientClass = _MatrixFullyConnectedGradientClass{objc.GetClass("MPSMatrixFullyConnectedGradient")}
	})
	return MatrixFullyConnectedGradientClass
}

type _MatrixFullyConnectedGradientClass struct {
	class objc.Class
}





// An interface definition for the [MatrixFullyConnectedGradient] class.
type IMatrixFullyConnectedGradient interface {
	IMatrixBinaryKernel
	

	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	SourceInputFeatureChannels() objectivec.IObject
	SetSourceInputFeatureChannels(value objectivec.IObject)
	SourceNumberOfFeatureVectors() objectivec.IObject
	SetSourceNumberOfFeatureVectors(value objectivec.IObject)
	SourceOutputFeatureChannels() objectivec.IObject
	SetSourceOutputFeatureChannels(value objectivec.IObject)


	

	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	EncodeForData()
	EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, weightMatrix IMatrix, resultGradientForDataMatrix IMatrix)
	EncodeForWeightsAndBias()
	EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, inputMatrix IMatrix, resultGradientForWeightMatrix IMatrix, resultGradientForBiasVector IVector)


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixFullyConnectedGradientClass) Alloc() MatrixFullyConnectedGradient {
	rv := objc.Send[MatrixFullyConnectedGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixFullyConnectedGradientClass) New() MatrixFullyConnectedGradient {
	rv := objc.Send[MatrixFullyConnectedGradient](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixFullyConnectedGradient) Init() MatrixFullyConnectedGradient {
	rv := objc.Send[MatrixFullyConnectedGradient](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixFullyConnectedGradient) Autorelease() MatrixFullyConnectedGradient {
	rv := objc.Send[MatrixFullyConnectedGradient](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixFullyConnectedGradient creates a new MatrixFullyConnectedGradient instance.
func NewMatrixFullyConnectedGradient() MatrixFullyConnectedGradient {
	return getMatrixFullyConnectedGradientClass().New()
}





// A kernel for applying a fully gradient connected neural network layer.


// A kernel for applying a fully gradient connected neural network layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient
type MatrixFullyConnectedGradient struct {
	MatrixBinaryKernel
}

// MatrixFullyConnectedGradientFrom constructs a [MatrixFullyConnectedGradient] from an unsafe.Pointer.
//
// A kernel for applying a fully gradient connected neural network layer.
func MatrixFullyConnectedGradientFrom(ptr unsafe.Pointer) MatrixFullyConnectedGradient {
	return MatrixFullyConnectedGradient{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966667-initwithcoder
func NewMatrixFullyConnectedGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixFullyConnectedGradient {
	instance := getMatrixFullyConnectedGradientClass().Alloc()
	rv := objc.Send[MatrixFullyConnectedGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966668-initwithdevice
func NewMatrixFullyConnectedGradientWithDevice(device unsafe.Pointer) MatrixFullyConnectedGradient {
	instance := getMatrixFullyConnectedGradientClass().Alloc()
	rv := objc.Send[MatrixFullyConnectedGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966664-copywithzone
func (m_ MatrixFullyConnectedGradient) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966665-encodefordata
func (m_ MatrixFullyConnectedGradient) EncodeForData() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeForData"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966665-encodegradientfordatatocommandbu
func (m_ MatrixFullyConnectedGradient) EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, weightMatrix IMatrix, resultGradientForDataMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeGradientForDataToCommandBuffer:gradientMatrix:weightMatrix:resultGradientForDataMatrix:"), commandBuffer, gradientMatrix, weightMatrix, resultGradientForDataMatrix)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966666-encodeforweightsandbias
func (m_ MatrixFullyConnectedGradient) EncodeForWeightsAndBias() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeForWeightsAndBias"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966666-encodegradientforweightsandbiast
func (m_ MatrixFullyConnectedGradient) EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, inputMatrix IMatrix, resultGradientForWeightMatrix IMatrix, resultGradientForBiasVector IVector) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeGradientForWeightsAndBiasToCommandBuffer:gradientMatrix:inputMatrix:resultGradientForWeightMatrix:resultGradientForBiasVector:"), commandBuffer, gradientMatrix, inputMatrix, resultGradientForWeightMatrix, resultGradientForBiasVector)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966663-alpha
func (m_ MatrixFullyConnectedGradient) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966663-alpha
func (m_ MatrixFullyConnectedGradient) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966669-sourceinputfeaturechannels
func (m_ MatrixFullyConnectedGradient) SourceInputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966669-sourceinputfeaturechannels
func (m_ MatrixFullyConnectedGradient) SetSourceInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966670-sourcenumberoffeaturevectors
func (m_ MatrixFullyConnectedGradient) SourceNumberOfFeatureVectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966670-sourcenumberoffeaturevectors
func (m_ MatrixFullyConnectedGradient) SetSourceNumberOfFeatureVectors(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966671-sourceoutputfeaturechannels
func (m_ MatrixFullyConnectedGradient) SourceOutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceOutputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966671-sourceoutputfeaturechannels
func (m_ MatrixFullyConnectedGradient) SetSourceOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceOutputFeatureChannels:"), value)
}







