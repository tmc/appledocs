// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNOptimizerRMSProp */


/* debug [class_header]: Header for MPSNNOptimizerRMSProp */
// The class instance for the [OptimizerRMSProp] class.
var (
	OptimizerRMSPropClass     _OptimizerRMSPropClass
	OptimizerRMSPropClassOnce sync.Once
)

func getOptimizerRMSPropClass() _OptimizerRMSPropClass {
	OptimizerRMSPropClassOnce.Do(func() {
		OptimizerRMSPropClass = _OptimizerRMSPropClass{objc.GetClass("MPSNNOptimizerRMSProp")}
	})
	return OptimizerRMSPropClass
}

type _OptimizerRMSPropClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OptimizerRMSProp */
// An interface definition for the [OptimizerRMSProp] class.
type IOptimizerRMSProp interface {
	IOptimizer
	
/* debug [class_interface_properties]: Properties for OptimizerRMSProp */
	// properties:
	Decay() objectivec.IObject
	SetDecay(value objectivec.IObject)
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OptimizerRMSProp */
	// methods:
	Encode()
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector(commandBuffer unsafe.Pointer, inputGradientVector IVector, inputValuesVector IVector, inputSumOfSquaresVector IVector, resultValuesVector IVector)
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputSumOfSquaresVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState(commandBuffer unsafe.Pointer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputSumOfSquaresVectors unsafe.Pointer, resultState ICNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState, inputSumOfSquaresVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBuffer unsafe.Pointer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputSumOfSquaresMatrix IMatrix, resultValuesMatrix IMatrix)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OptimizerRMSProp */
// Alloc allocates a new instance without initialization.
func (oc _OptimizerRMSPropClass) Alloc() OptimizerRMSProp {
	rv := objc.Send[OptimizerRMSProp](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OptimizerRMSPropClass) New() OptimizerRMSProp {
	rv := objc.Send[OptimizerRMSProp](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OptimizerRMSProp) Init() OptimizerRMSProp {
	rv := objc.Send[OptimizerRMSProp](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OptimizerRMSProp) Autorelease() OptimizerRMSProp {
	rv := objc.Send[OptimizerRMSProp](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOptimizerRMSProp creates a new OptimizerRMSProp instance.
func NewOptimizerRMSProp() OptimizerRMSProp {
	return getOptimizerRMSPropClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OptimizerRMSProp */
// An optimization layer that performs a root mean square propagation update.


// An optimization layer that performs a root mean square propagation update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp
type OptimizerRMSProp struct {
	Optimizer
}

// OptimizerRMSPropFrom constructs a [OptimizerRMSProp] from an unsafe.Pointer.
//
// An optimization layer that performs a root mean square propagation update.
func OptimizerRMSPropFrom(ptr unsafe.Pointer) OptimizerRMSProp {
	return OptimizerRMSProp{
		Optimizer: OptimizerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OptimizerRMSProp */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966737-initwithdevice
func NewOptimizerRMSPropWithDeviceDecayEpsilonOptimizerDescriptor(device unsafe.Pointer, decay float64, epsilon float32, optimizerDescriptor IOptimizerDescriptor) OptimizerRMSProp {
	instance := getOptimizerRMSPropClass().Alloc()
	rv := objc.Send[OptimizerRMSProp](instance.ID, objc.Sel("initWithDevice:decay:epsilon:optimizerDescriptor:"), device, decay, epsilon, optimizerDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOptimizerRMSPropWithDeviceDecayEpsilonOptimizerDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966738-initwithdevice
func NewOptimizerRMSPropWithDeviceLearningRate(device unsafe.Pointer, learningRate float32) OptimizerRMSProp {
	instance := getOptimizerRMSPropClass().Alloc()
	rv := objc.Send[OptimizerRMSProp](instance.ID, objc.Sel("initWithDevice:learningRate:"), device, learningRate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOptimizerRMSPropWithDeviceLearningRate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OptimizerRMSProp */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OptimizerRMSProp */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OptimizerRMSProp */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966735-encode
func (o_ OptimizerRMSProp) Encode() {
	objc.Send[objc.ID](o_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966735-encodetocommandbuffer
func (o_ OptimizerRMSProp) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector(commandBuffer unsafe.Pointer, inputGradientVector IVector, inputValuesVector IVector, inputSumOfSquaresVector IVector, resultValuesVector IVector) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputSumOfSquaresVector:resultValuesVector:"), commandBuffer, inputGradientVector, inputValuesVector, inputSumOfSquaresVector, resultValuesVector)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3013783-encodetocommandbuffer
func (o_ OptimizerRMSProp) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputSumOfSquaresVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputSumOfSquaresVectors:resultState:"), commandBuffer, batchNormalizationGradientState, batchNormalizationSourceState, inputSumOfSquaresVectors, resultState)
}/* debug [instance_methods/method]: EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3013784-encodetocommandbuffer
func (o_ OptimizerRMSProp) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState(commandBuffer unsafe.Pointer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputSumOfSquaresVectors unsafe.Pointer, resultState ICNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputSumOfSquaresVectors:resultState:"), commandBuffer, convolutionGradientState, convolutionSourceState, inputSumOfSquaresVectors, resultState)
}/* debug [instance_methods/method]: EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3019335-encodetocommandbuffer
func (o_ OptimizerRMSProp) EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState, inputSumOfSquaresVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputSumOfSquaresVectors:resultState:"), commandBuffer, batchNormalizationState, inputSumOfSquaresVectors, resultState)
}/* debug [instance_methods/method]: EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3197827-encodetocommandbuffer
func (o_ OptimizerRMSProp) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBuffer unsafe.Pointer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputSumOfSquaresMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:resultValuesMatrix:"), commandBuffer, inputGradientMatrix, inputValuesMatrix, inputSumOfSquaresMatrix, resultValuesMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OptimizerRMSProp */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966734-decay
func (o_ OptimizerRMSProp) Decay() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("decay"))
	return rv
}/* debug [instance_properties/getter]: decay */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966734-decay
func (o_ OptimizerRMSProp) SetDecay(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDecay:"), value)
}/* debug [instance_properties/setter]: decay */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966736-epsilon
func (o_ OptimizerRMSProp) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966736-epsilon
func (o_ OptimizerRMSProp) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNOptimizerRMSProp */


