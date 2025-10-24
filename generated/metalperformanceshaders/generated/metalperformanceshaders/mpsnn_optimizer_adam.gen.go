// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNOptimizerAdam */


/* debug [class_header]: Header for MPSNNOptimizerAdam */
// The class instance for the [OptimizerAdam] class.
var (
	OptimizerAdamClass     _OptimizerAdamClass
	OptimizerAdamClassOnce sync.Once
)

func getOptimizerAdamClass() _OptimizerAdamClass {
	OptimizerAdamClassOnce.Do(func() {
		OptimizerAdamClass = _OptimizerAdamClass{objc.GetClass("MPSNNOptimizerAdam")}
	})
	return OptimizerAdamClass
}

type _OptimizerAdamClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OptimizerAdam */
// An interface definition for the [OptimizerAdam] class.
type IOptimizerAdam interface {
	IOptimizer
	
/* debug [class_interface_properties]: Properties for OptimizerAdam */
	// properties:
	Beta1() objectivec.IObject
	SetBeta1(value objectivec.IObject)
	Beta2() objectivec.IObject
	SetBeta2(value objectivec.IObject)
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	TimeStep() objectivec.IObject
	SetTimeStep(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OptimizerAdam */
	// methods:
	Encode()
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector(commandBuffer unsafe.Pointer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, resultValuesVector IVector)
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer unsafe.Pointer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, resultState ICNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, maximumVelocityVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, maximumVelocityVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer unsafe.Pointer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, maximumVelocityVectors unsafe.Pointer, resultState ICNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector(commandBuffer unsafe.Pointer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, maximumVelocityVector IVector, resultValuesVector IVector)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix(commandBuffer unsafe.Pointer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, maximumVelocityMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBuffer unsafe.Pointer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, resultValuesMatrix IMatrix)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OptimizerAdam */
// Alloc allocates a new instance without initialization.
func (oc _OptimizerAdamClass) Alloc() OptimizerAdam {
	rv := objc.Send[OptimizerAdam](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OptimizerAdamClass) New() OptimizerAdam {
	rv := objc.Send[OptimizerAdam](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OptimizerAdam) Init() OptimizerAdam {
	rv := objc.Send[OptimizerAdam](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OptimizerAdam) Autorelease() OptimizerAdam {
	rv := objc.Send[OptimizerAdam](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOptimizerAdam creates a new OptimizerAdam instance.
func NewOptimizerAdam() OptimizerAdam {
	return getOptimizerAdamClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OptimizerAdam */
// An optimization layer that performs an Adam pdate.


// An optimization layer that performs an Adam pdate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam
type OptimizerAdam struct {
	Optimizer
}

// OptimizerAdamFrom constructs a [OptimizerAdam] from an unsafe.Pointer.
//
// An optimization layer that performs an Adam pdate.
func OptimizerAdamFrom(ptr unsafe.Pointer) OptimizerAdam {
	return OptimizerAdam{
		Optimizer: OptimizerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OptimizerAdam */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966718-initwithdevice
func NewOptimizerAdamWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor(device unsafe.Pointer, beta1 float64, beta2 float64, epsilon float32, timeStep uint, optimizerDescriptor IOptimizerDescriptor) OptimizerAdam {
	instance := getOptimizerAdamClass().Alloc()
	rv := objc.Send[OptimizerAdam](instance.ID, objc.Sel("initWithDevice:beta1:beta2:epsilon:timeStep:optimizerDescriptor:"), device, beta1, beta2, epsilon, timeStep, optimizerDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOptimizerAdamWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966719-initwithdevice
func NewOptimizerAdamWithDeviceLearningRate(device unsafe.Pointer, learningRate float32) OptimizerAdam {
	instance := getOptimizerAdamClass().Alloc()
	rv := objc.Send[OptimizerAdam](instance.ID, objc.Sel("initWithDevice:learningRate:"), device, learningRate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOptimizerAdamWithDeviceLearningRate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OptimizerAdam */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OptimizerAdam */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OptimizerAdam */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966716-encode
func (o_ OptimizerAdam) Encode() {
	objc.Send[objc.ID](o_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966716-encodetocommandbuffer
func (o_ OptimizerAdam) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector(commandBuffer unsafe.Pointer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, resultValuesVector IVector) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:resultValuesVector:"), commandBuffer, inputGradientVector, inputValuesVector, inputMomentumVector, inputVelocityVector, resultValuesVector)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3013781-encodetocommandbuffer
func (o_ OptimizerAdam) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:resultState:"), commandBuffer, batchNormalizationGradientState, batchNormalizationSourceState, inputMomentumVectors, inputVelocityVectors, resultState)
}/* debug [instance_methods/method]: EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3013782-encodetocommandbuffer
func (o_ OptimizerAdam) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer unsafe.Pointer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, resultState ICNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:resultState:"), commandBuffer, convolutionGradientState, convolutionSourceState, inputMomentumVectors, inputVelocityVectors, resultState)
}/* debug [instance_methods/method]: EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3019334-encodetocommandbuffer
func (o_ OptimizerAdam) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:resultState:"), commandBuffer, batchNormalizationState, inputMomentumVectors, inputVelocityVectors, resultState)
}/* debug [instance_methods/method]: EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175013-encodetocommandbuffer
func (o_ OptimizerAdam) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, maximumVelocityVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), commandBuffer, batchNormalizationGradientState, batchNormalizationSourceState, inputMomentumVectors, inputVelocityVectors, maximumVelocityVectors, resultState)
}/* debug [instance_methods/method]: EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175014-encodetocommandbuffer
func (o_ OptimizerAdam) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, maximumVelocityVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), commandBuffer, batchNormalizationState, inputMomentumVectors, inputVelocityVectors, maximumVelocityVectors, resultState)
}/* debug [instance_methods/method]: EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175015-encodetocommandbuffer
func (o_ OptimizerAdam) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer unsafe.Pointer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors unsafe.Pointer, inputVelocityVectors unsafe.Pointer, maximumVelocityVectors unsafe.Pointer, resultState ICNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), commandBuffer, convolutionGradientState, convolutionSourceState, inputMomentumVectors, inputVelocityVectors, maximumVelocityVectors, resultState)
}/* debug [instance_methods/method]: EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175016-encodetocommandbuffer
func (o_ OptimizerAdam) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector(commandBuffer unsafe.Pointer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, maximumVelocityVector IVector, resultValuesVector IVector) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:maximumVelocityVector:resultValuesVector:"), commandBuffer, inputGradientVector, inputValuesVector, inputMomentumVector, inputVelocityVector, maximumVelocityVector, resultValuesVector)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3197825-encodetocommandbuffer
func (o_ OptimizerAdam) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix(commandBuffer unsafe.Pointer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, maximumVelocityMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:maximumVelocityMatrix:resultValuesMatrix:"), commandBuffer, inputGradientMatrix, inputValuesMatrix, inputMomentumMatrix, inputVelocityMatrix, maximumVelocityMatrix, resultValuesMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3197826-encodetocommandbuffer
func (o_ OptimizerAdam) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBuffer unsafe.Pointer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:resultValuesMatrix:"), commandBuffer, inputGradientMatrix, inputValuesMatrix, inputMomentumMatrix, inputVelocityMatrix, resultValuesMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OptimizerAdam */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966714-beta1
func (o_ OptimizerAdam) Beta1() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("beta1"))
	return rv
}/* debug [instance_properties/getter]: beta1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966714-beta1
func (o_ OptimizerAdam) SetBeta1(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBeta1:"), value)
}/* debug [instance_properties/setter]: beta1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966715-beta2
func (o_ OptimizerAdam) Beta2() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("beta2"))
	return rv
}/* debug [instance_properties/getter]: beta2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966715-beta2
func (o_ OptimizerAdam) SetBeta2(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBeta2:"), value)
}/* debug [instance_properties/setter]: beta2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966717-epsilon
func (o_ OptimizerAdam) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966717-epsilon
func (o_ OptimizerAdam) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966720-timestep
func (o_ OptimizerAdam) TimeStep() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("timeStep"))
	return rv
}/* debug [instance_properties/getter]: timeStep */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966720-timestep
func (o_ OptimizerAdam) SetTimeStep(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setTimeStep:"), value)
}/* debug [instance_properties/setter]: timeStep */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNOptimizerAdam */


