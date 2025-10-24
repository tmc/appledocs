// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixFullyConnectedGradient */


/* debug [class_header]: Header for MPSMatrixFullyConnectedGradient */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixFullyConnectedGradient */
// An interface definition for the [MatrixFullyConnectedGradient] class.
type IMatrixFullyConnectedGradient interface {
	IMatrixBinaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixFullyConnectedGradient */
	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	SourceInputFeatureChannels() objectivec.IObject
	SetSourceInputFeatureChannels(value objectivec.IObject)
	SourceNumberOfFeatureVectors() objectivec.IObject
	SetSourceNumberOfFeatureVectors(value objectivec.IObject)
	SourceOutputFeatureChannels() objectivec.IObject
	SetSourceOutputFeatureChannels(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixFullyConnectedGradient */
	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	EncodeForData()
	EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, weightMatrix IMatrix, resultGradientForDataMatrix IMatrix)
	EncodeForWeightsAndBias()
	EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, inputMatrix IMatrix, resultGradientForWeightMatrix IMatrix, resultGradientForBiasVector IVector)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixFullyConnectedGradient */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixFullyConnectedGradient */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixFullyConnectedGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966667-initwithcoder
func NewMatrixFullyConnectedGradientWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) MatrixFullyConnectedGradient {
	instance := getMatrixFullyConnectedGradientClass().Alloc()
	rv := objc.Send[MatrixFullyConnectedGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixFullyConnectedGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966668-initwithdevice
func NewMatrixFullyConnectedGradientWithDevice(device unsafe.Pointer) MatrixFullyConnectedGradient {
	instance := getMatrixFullyConnectedGradientClass().Alloc()
	rv := objc.Send[MatrixFullyConnectedGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixFullyConnectedGradientWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixFullyConnectedGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixFullyConnectedGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixFullyConnectedGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966664-copywithzone
func (m_ MatrixFullyConnectedGradient) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966665-encodefordata
func (m_ MatrixFullyConnectedGradient) EncodeForData() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeForData"))
}/* debug [instance_methods/method]: EncodeForData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966665-encodegradientfordatatocommandbu
func (m_ MatrixFullyConnectedGradient) EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, weightMatrix IMatrix, resultGradientForDataMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeGradientForDataToCommandBuffer:gradientMatrix:weightMatrix:resultGradientForDataMatrix:"), commandBuffer, gradientMatrix, weightMatrix, resultGradientForDataMatrix)
}/* debug [instance_methods/method]: EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966666-encodeforweightsandbias
func (m_ MatrixFullyConnectedGradient) EncodeForWeightsAndBias() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeForWeightsAndBias"))
}/* debug [instance_methods/method]: EncodeForWeightsAndBias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966666-encodegradientforweightsandbiast
func (m_ MatrixFullyConnectedGradient) EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, inputMatrix IMatrix, resultGradientForWeightMatrix IMatrix, resultGradientForBiasVector IVector) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeGradientForWeightsAndBiasToCommandBuffer:gradientMatrix:inputMatrix:resultGradientForWeightMatrix:resultGradientForBiasVector:"), commandBuffer, gradientMatrix, inputMatrix, resultGradientForWeightMatrix, resultGradientForBiasVector)
}/* debug [instance_methods/method]: EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixFullyConnectedGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966663-alpha
func (m_ MatrixFullyConnectedGradient) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966663-alpha
func (m_ MatrixFullyConnectedGradient) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966669-sourceinputfeaturechannels
func (m_ MatrixFullyConnectedGradient) SourceInputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceInputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966669-sourceinputfeaturechannels
func (m_ MatrixFullyConnectedGradient) SetSourceInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceInputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966670-sourcenumberoffeaturevectors
func (m_ MatrixFullyConnectedGradient) SourceNumberOfFeatureVectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}/* debug [instance_properties/getter]: sourceNumberOfFeatureVectors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966670-sourcenumberoffeaturevectors
func (m_ MatrixFullyConnectedGradient) SetSourceNumberOfFeatureVectors(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}/* debug [instance_properties/setter]: sourceNumberOfFeatureVectors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966671-sourceoutputfeaturechannels
func (m_ MatrixFullyConnectedGradient) SourceOutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceOutputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceOutputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966671-sourceoutputfeaturechannels
func (m_ MatrixFullyConnectedGradient) SetSourceOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceOutputFeatureChannels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixFullyConnectedGradient */


