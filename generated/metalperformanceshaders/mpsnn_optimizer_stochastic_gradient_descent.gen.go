// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [OptimizerStochasticGradientDescent] class.
var (
	OptimizerStochasticGradientDescentClass     _OptimizerStochasticGradientDescentClass
	OptimizerStochasticGradientDescentClassOnce sync.Once
)

func getOptimizerStochasticGradientDescentClass() _OptimizerStochasticGradientDescentClass {
	OptimizerStochasticGradientDescentClassOnce.Do(func() {
		OptimizerStochasticGradientDescentClass = _OptimizerStochasticGradientDescentClass{objc.GetClass("MPSNNOptimizerStochasticGradientDescent")}
	})
	return OptimizerStochasticGradientDescentClass
}

type _OptimizerStochasticGradientDescentClass struct {
	class objc.Class
}





// An interface definition for the [OptimizerStochasticGradientDescent] class.
type IOptimizerStochasticGradientDescent interface {
	IOptimizer
	

	// properties:
	MomentumScale() objectivec.IObject
	SetMomentumScale(value objectivec.IObject)
	UseNestrovMomentum() objectivec.IObject
	SetUseNestrovMomentum(value objectivec.IObject)
	UseNesterovMomentum() objectivec.IObject
	SetUseNesterovMomentum(value objectivec.IObject)


	

	// methods:
	Encode()
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector(commandBuffer unsafe.Pointer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, resultValuesVector IVector)
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState(commandBuffer unsafe.Pointer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors unsafe.Pointer, resultState ICNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBuffer unsafe.Pointer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, resultValuesMatrix IMatrix)


}





// Alloc allocates a new instance without initialization.
func (oc _OptimizerStochasticGradientDescentClass) Alloc() OptimizerStochasticGradientDescent {
	rv := objc.Send[OptimizerStochasticGradientDescent](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OptimizerStochasticGradientDescentClass) New() OptimizerStochasticGradientDescent {
	rv := objc.Send[OptimizerStochasticGradientDescent](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OptimizerStochasticGradientDescent) Init() OptimizerStochasticGradientDescent {
	rv := objc.Send[OptimizerStochasticGradientDescent](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OptimizerStochasticGradientDescent) Autorelease() OptimizerStochasticGradientDescent {
	rv := objc.Send[OptimizerStochasticGradientDescent](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOptimizerStochasticGradientDescent creates a new OptimizerStochasticGradientDescent instance.
func NewOptimizerStochasticGradientDescent() OptimizerStochasticGradientDescent {
	return getOptimizerStochasticGradientDescentClass().New()
}





// An optimization layer that performs a gradient descent with an optional momentum update.


// An optimization layer that performs a gradient descent with an optional momentum update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent
type OptimizerStochasticGradientDescent struct {
	Optimizer
}

// OptimizerStochasticGradientDescentFrom constructs a [OptimizerStochasticGradientDescent] from an unsafe.Pointer.
//
// An optimization layer that performs a gradient descent with an optional momentum update.
func OptimizerStochasticGradientDescentFrom(ptr unsafe.Pointer) OptimizerStochasticGradientDescent {
	return OptimizerStochasticGradientDescent{
		Optimizer: OptimizerFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966741-initwithdevice
func NewOptimizerStochasticGradientDescentWithDeviceLearningRate(device unsafe.Pointer, learningRate float32) OptimizerStochasticGradientDescent {
	instance := getOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[OptimizerStochasticGradientDescent](instance.ID, objc.Sel("initWithDevice:learningRate:"), device, learningRate)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3675591-initwithdevice
func NewOptimizerStochasticGradientDescentWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor(device unsafe.Pointer, momentumScale float32, useNesterovMomentum bool, optimizerDescriptor IOptimizerDescriptor) OptimizerStochasticGradientDescent {
	instance := getOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[OptimizerStochasticGradientDescent](instance.ID, objc.Sel("initWithDevice:momentumScale:useNesterovMomentum:optimizerDescriptor:"), device, momentumScale, useNesterovMomentum, optimizerDescriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966742-initwithdevice
func NewOptimizerStochasticGradientDescentWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor(device unsafe.Pointer, momentumScale float32, useNestrovMomentum bool, optimizerDescriptor IOptimizerDescriptor) OptimizerStochasticGradientDescent {
	instance := getOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[OptimizerStochasticGradientDescent](instance.ID, objc.Sel("initWithDevice:momentumScale:useNestrovMomentum:optimizerDescriptor:"), device, momentumScale, useNestrovMomentum, optimizerDescriptor)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966740-encode
func (o_ OptimizerStochasticGradientDescent) Encode() {
	objc.Send[objc.ID](o_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966740-encodetocommandbuffer
func (o_ OptimizerStochasticGradientDescent) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector(commandBuffer unsafe.Pointer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, resultValuesVector IVector) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:resultValuesVector:"), commandBuffer, inputGradientVector, inputValuesVector, inputMomentumVector, resultValuesVector)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3013785-encodetocommandbuffer
func (o_ OptimizerStochasticGradientDescent) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:resultState:"), commandBuffer, batchNormalizationGradientState, batchNormalizationSourceState, inputMomentumVectors, resultState)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3013786-encodetocommandbuffer
func (o_ OptimizerStochasticGradientDescent) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState(commandBuffer unsafe.Pointer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors unsafe.Pointer, resultState ICNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:resultState:"), commandBuffer, convolutionGradientState, convolutionSourceState, inputMomentumVectors, resultState)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3019336-encodetocommandbuffer
func (o_ OptimizerStochasticGradientDescent) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors unsafe.Pointer, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:resultState:"), commandBuffer, batchNormalizationState, inputMomentumVectors, resultState)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3197828-encodetocommandbuffer
func (o_ OptimizerStochasticGradientDescent) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBuffer unsafe.Pointer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Send[objc.ID](o_.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:resultValuesMatrix:"), commandBuffer, inputGradientMatrix, inputValuesMatrix, inputMomentumMatrix, resultValuesMatrix)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966743-momentumscale
func (o_ OptimizerStochasticGradientDescent) MomentumScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("momentumScale"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966743-momentumscale
func (o_ OptimizerStochasticGradientDescent) SetMomentumScale(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setMomentumScale:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966744-usenestrovmomentum
func (o_ OptimizerStochasticGradientDescent) UseNestrovMomentum() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("useNestrovMomentum"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966744-usenestrovmomentum
func (o_ OptimizerStochasticGradientDescent) SetUseNestrovMomentum(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUseNestrovMomentum:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3675592-usenesterovmomentum
func (o_ OptimizerStochasticGradientDescent) UseNesterovMomentum() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("useNesterovMomentum"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3675592-usenesterovmomentum
func (o_ OptimizerStochasticGradientDescent) SetUseNesterovMomentum(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUseNesterovMomentum:"), value)
}







