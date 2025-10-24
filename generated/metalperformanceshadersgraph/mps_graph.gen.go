// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

/* debug [class.gen.go]: Generating class MPSGraph */


/* debug [class_header]: Header for MPSGraph */
// The class instance for the [Graph] class.
var (
	GraphClass     _GraphClass
	GraphClassOnce sync.Once
)

func getGraphClass() _GraphClass {
	GraphClassOnce.Do(func() {
		GraphClass = _GraphClass{objc.GetClass("MPSGraph")}
	})
	return GraphClass
}

type _GraphClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Graph */
// An interface definition for the [Graph] class.
type IGraph interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for Graph */
	// properties:
	Options() GraphOptions
	SetOptions(value GraphOptions)
	PlaceholderTensors() []GraphTensor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Graph */
	// methods:
	AbsoluteWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	AbsoluteSquareWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	AdamWithCurrentLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName(currentLearningRateTensor IMPSGraphTensor, beta1Tensor IMPSGraphTensor, beta2Tensor IMPSGraphTensor, epsilonTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, momentumTensor IMPSGraphTensor, velocityTensor IMPSGraphTensor, maximumVelocityTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	AdamWithLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorBeta1PowerTensorBeta2PowerTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName(learningRateTensor IMPSGraphTensor, beta1Tensor IMPSGraphTensor, beta2Tensor IMPSGraphTensor, epsilonTensor IMPSGraphTensor, beta1PowerTensor IMPSGraphTensor, beta2PowerTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, momentumTensor IMPSGraphTensor, velocityTensor IMPSGraphTensor, maximumVelocityTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	AdditionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ApplyStochasticGradientDescentWithLearningRateTensorVariableGradientTensorName(learningRateTensor IMPSGraphTensor, variable IMPSGraphVariableOp, gradientTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphOperation
	ArgSortWithTensorAxisDescendingName(tensor IMPSGraphTensor, axis int, descending bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ArgSortWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ArgSortWithTensorAxisTensorDescendingName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, descending bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ArgSortWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	AssignVariableWithValueOfTensorName(variable IMPSGraphTensor, tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphOperation
	AvgPooling2DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	AvgPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	AvgPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	AvgPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BandPartWithTensorNumLowerNumUpperName(inputTensor IMPSGraphTensor, numLower int, numUpper int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BandPartWithTensorNumLowerTensorNumUpperTensorName(inputTensor IMPSGraphTensor, numLowerTensor IMPSGraphTensor, numUpperTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BatchToSpaceTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxes []foundation.Number, batchAxis int, blockDimensions []foundation.Number, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BatchToSpaceTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxesTensor IMPSGraphTensor, batchAxisTensor IMPSGraphTensor, blockDimensionsTensor IMPSGraphTensor, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BitwiseANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BitwiseLeftShiftWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BitwiseNOTWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BitwiseORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BitwisePopulationCountWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BitwiseRightShiftWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BitwiseXORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BottomKWithSourceTensorAxisKName(source IMPSGraphTensor, axis int, k uint, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	BottomKWithSourceTensorAxisTensorKTensorName(source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	BottomKWithGradientTensorSourceAxisKName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, k uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BottomKWithGradientTensorSourceAxisTensorKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BroadcastTensorToShapeName(tensor IMPSGraphTensor, shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	BroadcastTensorToShapeTensorName(tensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CallSymbolNameInputTensorsOutputTypesName(symbolName objc.IObject /* cross-framework: NSString */, inputTensors []GraphTensor, outputTypes []GraphType, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	CastTensorToTypeName(tensor IMPSGraphTensor, type_ DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CeilWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ClampWithTensorMinValueTensorMaxValueTensorName(tensor IMPSGraphTensor, minValueTensor IMPSGraphTensor, maxValueTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ColToImWithSourceTensorOutputShapeDescriptorName(source IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphImToColOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CompileWithDeviceFeedsTargetTensorsTargetOperationsCompilationDescriptor(device IMPSGraphDevice, feeds GraphTensorShapedTypeDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation, compilationDescriptor IMPSGraphCompilationDescriptor) IGraphExecutable
	ConstantWithRealPartImaginaryPart(realPart float64, imaginaryPart float64) IGraphTensor
	ConstantWithRealPartImaginaryPartDataType(realPart float64, imaginaryPart float64, dataType DataType /* not a class type */) IGraphTensor
	ConstantWithRealPartImaginaryPartShapeDataType(realPart float64, imaginaryPart float64, shape Shape /* not a class type */, dataType DataType /* not a class type */) IGraphTensor
	ComplexTensorWithRealTensorImaginaryTensorName(realTensor IMPSGraphTensor, imaginaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConcatTensorWithTensorDimensionName(tensor IMPSGraphTensor, tensor2 IMPSGraphTensor, dimensionIndex int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConcatTensorsDimensionInterleaveName(tensors []GraphTensor, dimensionIndex int, interleave bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConcatTensorsDimensionName(tensors []GraphTensor, dimensionIndex int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConjugateWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConstantWithScalarDataType(scalar float64, dataType DataType /* not a class type */) IGraphTensor
	ConstantWithScalarShapeDataType(scalar float64, shape Shape /* not a class type */, dataType DataType /* not a class type */) IGraphTensor
	ConstantWithDataShapeDataType(data objc.IObject /* cross-framework: NSData */, shape Shape /* not a class type */, dataType DataType /* not a class type */) IGraphTensor
	ControlDependencyWithOperationsDependentBlockName(operations []GraphOperation, dependentBlock GraphControlFlowDependencyBlock /* not a class type */, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	Convolution2DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, weights IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Convolution3DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, weights IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradientTensor IMPSGraphTensor, source IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradientTensor IMPSGraphTensor, source IMPSGraphTensor, outputShape IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CoordinateAlongAxisWithShapeName(axis int, shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CoordinateAlongAxisWithShapeTensorName(axis int, shapeTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CoordinateAlongAxisTensorWithShapeName(axisTensor IMPSGraphTensor, shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CoordinateAlongAxisTensorWithShapeTensorName(axisTensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeMaximumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeMaximumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeMaximumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeMinimumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeMinimumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeMinimumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeProductWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeProductWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeProductWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeProductWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeSumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeSumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeSumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	CumulativeSumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DepthToSpace2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DepthToSpace2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxisTensor IMPSGraphTensor, heightAxisTensor IMPSGraphTensor, depthAxisTensor IMPSGraphTensor, blockSize uint, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DepthwiseConvolution2DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DepthwiseConvolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DepthwiseConvolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DepthwiseConvolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DequantizeTensorLUTTensorAxisName(tensor IMPSGraphTensor, LUTTensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DequantizeTensorLUTTensorName(tensor IMPSGraphTensor, LUTTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DequantizeTensorScaleZeroPointDataTypeName(tensor IMPSGraphTensor, scale float64, zeroPoint float64, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DequantizeTensorScaleTensorDataTypeName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DequantizeTensorScaleTensorZeroPointDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPoint float64, dataType DataType /* not a class type */, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DequantizeTensorScaleTensorZeroPointTensorDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType DataType /* not a class type */, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DequantizeTensorScaleTensorZeroPointTensorDataTypeName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DivisionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DivisionNoNaNWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DropoutTensorRateTensorName(tensor IMPSGraphTensor, rate IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	DropoutTensorRateName(tensor IMPSGraphTensor, rate float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	EncodeToCommandBufferFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandBuffer metalperformanceshaders.CommandBuffer, feeds GraphTensorDataDictionary /* not a class type */, targetOperations []GraphOperation, resultsDictionary GraphTensorDataDictionary /* not a class type */, executionDescriptor IMPSGraphExecutionDescriptor)
	EncodeToCommandBufferFeedsTargetTensorsTargetOperationsExecutionDescriptor(commandBuffer metalperformanceshaders.CommandBuffer, feeds GraphTensorDataDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) GraphTensorDataDictionary /* not a class type */
	EqualWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ExpandDimsOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ExpandDimsOfTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ExpandDimsOfTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ExponentWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ExponentBase10WithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ExponentBase2WithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	FastFourierTransformWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.Number, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	FastFourierTransformWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Flatten2DTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	Flatten2DTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	FloorWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	FloorModuloWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ForLoopWithLowerBoundUpperBoundStepInitialBodyArgumentsBodyName(lowerBound IMPSGraphTensor, upperBound IMPSGraphTensor, step IMPSGraphTensor, initialBodyArguments []GraphTensor, body GraphForLoopBodyBlock /* not a class type */, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	ForLoopWithNumberOfIterationsInitialBodyArgumentsBodyName(numberOfIterations IMPSGraphTensor, initialBodyArguments []GraphTensor, body GraphForLoopBodyBlock /* not a class type */, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	GatherWithUpdatesTensorIndicesTensorAxisBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, axis uint, batchDimensions uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	GatherAlongAxisWithUpdatesTensorIndicesTensorName(axis int, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	GatherAlongAxisTensorWithUpdatesTensorIndicesTensorName(axisTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	GatherNDWithUpdatesTensorIndicesTensorBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, batchDimensions uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	GradientForPrimaryTensorWithTensorsName(primaryTensor IMPSGraphTensor, tensors []GraphTensor, name objc.IObject /* cross-framework: NSString */) foundation.IDictionary
	GreaterThanWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	GreaterThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	GRUWithSourceTensorRecurrentWeightInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, secondaryBias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdStateGradientInputWeightBiasInitStateMaskSecondaryBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, stateGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, secondaryBias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	HammingDistanceWithPrimaryTensorSecondaryTensorResultDataTypeName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, resultDataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	HermiteanToRealFFTWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.Number, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	HermiteanToRealFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	IdentityWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	IfWithPredicateTensorThenBlockElseBlockName(predicateTensor IMPSGraphTensor, thenBlock GraphIfThenElseBlock /* not a class type */, elseBlock GraphIfThenElseBlock /* not a class type */, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	ImaginaryPartOfTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ImToColWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphImToColOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	InverseOfTensorName(inputTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	IsFiniteWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	IsInfiniteWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	IsNaNWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	L2NormPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	L2NormPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LeakyReLUWithTensorAlphaName(tensor IMPSGraphTensor, alpha float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LeakyReLUWithTensorAlphaTensorName(tensor IMPSGraphTensor, alphaTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LeakyReLUGradientWithIncomingGradientSourceTensorAlphaTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, alphaTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LessThanWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LessThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LogarithmWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LogarithmBase10WithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LogarithmBase2WithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LogicalANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LogicalNANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LogicalNORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LogicalORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LogicalXNORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LogicalXORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, peephole IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdStateGradientCellGradientInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, stateGradient IMPSGraphTensor, cellGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, peephole IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaximumWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaximumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaxPooling2DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaxPooling2DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	MaxPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaxPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MaxPooling4DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	MeanOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MinimumWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MinimumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ModuloWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	MultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	NegativeWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	NonMaximumSuppressionWithBoxesTensorScoresTensorClassIndicesTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName(boxesTensor IMPSGraphTensor, scoresTensor IMPSGraphTensor, classIndicesTensor IMPSGraphTensor, IOUThreshold float32, scoreThreshold float32, perClassSuppression bool, coordinateMode GraphNonMaximumSuppressionCoordinateMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	NonMaximumSuppressionWithBoxesTensorScoresTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName(boxesTensor IMPSGraphTensor, scoresTensor IMPSGraphTensor, IOUThreshold float32, scoreThreshold float32, perClassSuppression bool, coordinateMode GraphNonMaximumSuppressionCoordinateMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	NonZeroIndicesOfTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	NormalizationBetaGradientWithIncomingGradientTensorSourceTensorReductionAxesName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	NormalizationGammaGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorReductionAxesEpsilonName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, meanTensor IMPSGraphTensor, varianceTensor IMPSGraphTensor, axes []foundation.Number, epsilon float32, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	NormalizationGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorGammaTensorGammaGradientTensorBetaGradientTensorReductionAxesEpsilonName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, meanTensor IMPSGraphTensor, varianceTensor IMPSGraphTensor, gamma IMPSGraphTensor, gammaGradient IMPSGraphTensor, betaGradient IMPSGraphTensor, axes []foundation.Number, epsilon float32, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	NormalizationWithTensorMeanTensorVarianceTensorGammaTensorBetaTensorEpsilonName(tensor IMPSGraphTensor, mean IMPSGraphTensor, variance IMPSGraphTensor, gamma IMPSGraphTensor, beta IMPSGraphTensor, epsilon float32, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	NotWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	NotEqualWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	OneHotWithIndicesTensorDepthAxisDataTypeName(indicesTensor IMPSGraphTensor, depth uint, axis uint, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	OneHotWithIndicesTensorDepthAxisDataTypeOnValueOffValueName(indicesTensor IMPSGraphTensor, depth uint, axis uint, dataType DataType /* not a class type */, onValue float64, offValue float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	OneHotWithIndicesTensorDepthAxisName(indicesTensor IMPSGraphTensor, depth uint, axis uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	OneHotWithIndicesTensorDepthDataTypeName(indicesTensor IMPSGraphTensor, depth uint, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	OneHotWithIndicesTensorDepthDataTypeOnValueOffValueName(indicesTensor IMPSGraphTensor, depth uint, dataType DataType /* not a class type */, onValue float64, offValue float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	OneHotWithIndicesTensorDepthName(indicesTensor IMPSGraphTensor, depth uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	PadGradientWithIncomingGradientTensorSourceTensorPaddingModeLeftPaddingRightPaddingName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, paddingMode GraphPaddingMode, leftPadding Shape /* not a class type */, rightPadding Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	PadTensorWithPaddingModeLeftPaddingRightPaddingConstantValueName(tensor IMPSGraphTensor, paddingMode GraphPaddingMode, leftPadding Shape /* not a class type */, rightPadding Shape /* not a class type */, constantValue float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	PlaceholderWithShapeDataTypeName(shape Shape /* not a class type */, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	PlaceholderWithShapeName(shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	PowerWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	QuantizeTensorScaleZeroPointDataTypeName(tensor IMPSGraphTensor, scale float64, zeroPoint float64, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	QuantizeTensorScaleTensorZeroPointDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPoint float64, dataType DataType /* not a class type */, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	QuantizeTensorScaleTensorZeroPointTensorDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType DataType /* not a class type */, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomPhiloxStateTensorWithCounterLowCounterHighKeyName(counterLow uint, counterHigh uint, key uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomPhiloxStateTensorWithSeedName(seed uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomTensorWithShapeDescriptorName(shape Shape /* not a class type */, descriptor IMPSGraphRandomOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomTensorWithShapeDescriptorSeedName(shape Shape /* not a class type */, descriptor IMPSGraphRandomOpDescriptor, seed uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomTensorWithShapeDescriptorStateTensorName(shape Shape /* not a class type */, descriptor IMPSGraphRandomOpDescriptor, state IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	RandomTensorWithShapeTensorDescriptorName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomTensorWithShapeTensorDescriptorSeedName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, seed uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomTensorWithShapeTensorDescriptorStateTensorName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, state IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	RandomUniformTensorWithShapeName(shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomUniformTensorWithShapeSeedName(shape Shape /* not a class type */, seed uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomUniformTensorWithShapeStateTensorName(shape Shape /* not a class type */, state IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	RandomUniformTensorWithShapeTensorName(shapeTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomUniformTensorWithShapeTensorSeedName(shapeTensor IMPSGraphTensor, seed uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RandomUniformTensorWithShapeTensorStateTensorName(shapeTensor IMPSGraphTensor, state IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	ReadVariableName(variable IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RealPartOfTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RealToHermiteanFFTWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.Number, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RealToHermiteanFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReciprocalWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReciprocalSquareRootWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionAndWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionAndWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionArgMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionArgMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionMaximumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionMaximumPropagateNaNWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionMaximumPropagateNaNWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionMinimumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionMinimumPropagateNaNWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionMinimumPropagateNaNWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionOrWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionOrWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionProductWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionProductWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionSumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReductionSumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReinterpretCastTensorToTypeName(tensor IMPSGraphTensor, type_ DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReLUWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReLUGradientWithIncomingGradientSourceTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReshapeTensorWithShapeName(tensor IMPSGraphTensor, shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReshapeTensorWithShapeTensorName(tensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeTensorSizeModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size Shape /* not a class type */, mode GraphResizeMode, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeTensorSizeTensorModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, mode GraphResizeMode, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeTensorSizeTensorModeCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, mode GraphResizeMode, centerResult bool, alignCorners bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeTensorSizeTensorScaleOffsetTensorModeLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, mode GraphResizeMode, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeTensorSizeTensorScaleTensorOffsetTensorModeName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, mode GraphResizeMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeWithGradientTensorInputModeCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, mode GraphResizeMode, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeWithGradientTensorInputScaleTensorOffsetTensorModeName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, mode GraphResizeMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeWithGradientTensorInputScaleOffsetTensorModeLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, mode GraphResizeMode, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, centerResult bool, alignCorners bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeBilinearWithTensorSizeTensorScaleOffsetTensorLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeBilinearWithTensorSizeTensorScaleTensorOffsetTensorName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeBilinearWithGradientTensorInputCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeBilinearWithGradientTensorInputScaleTensorOffsetTensorName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeBilinearWithGradientTensorInputScaleOffsetTensorLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeNearestWithTensorSizeTensorScaleOffsetTensorNearestRoundingModeLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeNearestWithTensorSizeTensorScaleTensorOffsetTensorNearestRoundingModeName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeNearestWithGradientTensorInputNearestRoundingModeCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeNearestWithGradientTensorInputScaleTensorOffsetTensorNearestRoundingModeName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ResizeNearestWithGradientTensorInputScaleOffsetTensorNearestRoundingModeLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReverseTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReverseTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ReverseTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RoundWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	RunWithFeedsTargetTensorsTargetOperations(feeds GraphTensorDataDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation) GraphTensorDataDictionary /* not a class type */
	RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(commandQueue unsafe.Pointer, feeds GraphTensorDataDictionary /* not a class type */, targetOperations []GraphOperation, resultsDictionary GraphTensorDataDictionary /* not a class type */)
	RunWithMTLCommandQueueFeedsTargetTensorsTargetOperations(commandQueue unsafe.Pointer, feeds GraphTensorDataDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation) GraphTensorDataDictionary /* not a class type */
	RunAsyncWithFeedsTargetTensorsTargetOperationsExecutionDescriptor(feeds GraphTensorDataDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) GraphTensorDataDictionary /* not a class type */
	RunAsyncWithMTLCommandQueueFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandQueue unsafe.Pointer, feeds GraphTensorDataDictionary /* not a class type */, targetOperations []GraphOperation, resultsDictionary GraphTensorDataDictionary /* not a class type */, executionDescriptor IMPSGraphExecutionDescriptor)
	RunAsyncWithMTLCommandQueueFeedsTargetTensorsTargetOperationsExecutionDescriptor(commandQueue unsafe.Pointer, feeds GraphTensorDataDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) GraphTensorDataDictionary /* not a class type */
	SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeNearestRoundingModeConstantValueName(source IMPSGraphTensor, coordinates IMPSGraphTensor, layout GraphTensorNamedDataLayout, normalizeCoordinates bool, relativeCoordinates bool, alignCorners bool, paddingMode GraphPaddingMode, nearestRoundingMode GraphResizeNearestRoundingMode, constantValue float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeSamplingModeConstantValueName(source IMPSGraphTensor, coordinates IMPSGraphTensor, layout GraphTensorNamedDataLayout, normalizeCoordinates bool, relativeCoordinates bool, alignCorners bool, paddingMode GraphPaddingMode, samplingMode GraphResizeMode, constantValue float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorMaskTensorScaleName(queryTensor IMPSGraphTensor, keyTensor IMPSGraphTensor, valueTensor IMPSGraphTensor, maskTensor IMPSGraphTensor, scale float32, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorScaleName(queryTensor IMPSGraphTensor, keyTensor IMPSGraphTensor, valueTensor IMPSGraphTensor, scale float32, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScatterWithUpdatesTensorIndicesTensorShapeAxisModeName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape Shape /* not a class type */, axis int, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScatterAlongAxisWithDataTensorUpdatesTensorIndicesTensorModeName(axis int, dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScatterAlongAxisWithUpdatesTensorIndicesTensorShapeModeName(axis int, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape Shape /* not a class type */, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScatterAlongAxisTensorWithDataTensorUpdatesTensorIndicesTensorModeName(axisTensor IMPSGraphTensor, dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScatterAlongAxisTensorWithUpdatesTensorIndicesTensorShapeModeName(axisTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape Shape /* not a class type */, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsModeName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape Shape /* not a class type */, batchDimensions uint, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape Shape /* not a class type */, batchDimensions uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScatterNDWithDataTensorUpdatesTensorIndicesTensorBatchDimensionsModeName(dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, batchDimensions uint, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ScatterWithDataTensorUpdatesTensorIndicesTensorAxisModeName(dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, axis int, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SelectWithPredicateTensorTruePredicateTensorFalsePredicateTensorName(predicateTensor IMPSGraphTensor, truePredicateTensor IMPSGraphTensor, falseSelectTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	ShapeOfTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SigmoidWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SigmoidGradientWithIncomingGradientSourceTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SignWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SignbitWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateStateGradientInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, stateGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	SliceGradientTensorFwdInShapeTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, startTensor IMPSGraphTensor, endTensor IMPSGraphTensor, strideTensor IMPSGraphTensor, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceGradientTensorFwdInShapeTensorStartTensorSizeTensorSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, startTensor IMPSGraphTensor, sizeTensor IMPSGraphTensor, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceGradientTensorFwdInShapeTensorStartsEndsStridesName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceGradientTensorFwdInShapeTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceTensorDimensionStartLengthName(tensor IMPSGraphTensor, dimensionIndex uint, start int, length int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName(tensor IMPSGraphTensor, startTensor IMPSGraphTensor, endTensor IMPSGraphTensor, strideTensor IMPSGraphTensor, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceTensorStartTensorSizeTensorSqueezeMaskName(tensor IMPSGraphTensor, startTensor IMPSGraphTensor, sizeTensor IMPSGraphTensor, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceTensorStartsEndsStridesName(tensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(tensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceUpdateDataTensorUpdateTensorStartsEndsStridesName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceUpdateDataTensorUpdateTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, startsTensor IMPSGraphTensor, endsTensor IMPSGraphTensor, stridesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorStartMaskEndMaskSqueezeMaskName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, startsTensor IMPSGraphTensor, endsTensor IMPSGraphTensor, stridesTensor IMPSGraphTensor, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SoftMaxWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SoftMaxCrossEntropyWithSourceTensorLabelsTensorAxisReductionTypeName(sourceTensor IMPSGraphTensor, labelsTensor IMPSGraphTensor, axis int, reductionType GraphLossReductionType, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SoftMaxCrossEntropyGradientWithIncomingGradientTensorSourceTensorLabelsTensorAxisReductionTypeName(gradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, labelsTensor IMPSGraphTensor, axis int, reductionType GraphLossReductionType, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SoftMaxGradientWithIncomingGradientSourceTensorAxisName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SortWithTensorAxisDescendingName(tensor IMPSGraphTensor, axis int, descending bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SortWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SortWithTensorAxisTensorDescendingName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, descending bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SortWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SpaceToDepth2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxisTensor IMPSGraphTensor, heightAxisTensor IMPSGraphTensor, depthAxisTensor IMPSGraphTensor, blockSize uint, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SpaceToBatchTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxes []foundation.Number, batchAxis int, blockDimensions []foundation.Number, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SpaceToBatchTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxesTensor IMPSGraphTensor, batchAxisTensor IMPSGraphTensor, blockDimensionsTensor IMPSGraphTensor, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SparseTensorWithDescriptorTensorsShapeName(sparseDescriptor IMPSGraphCreateSparseOpDescriptor, inputTensorArray []GraphTensor, shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SparseTensorWithTypeTensorsShapeDataTypeName(sparseStorageType GraphSparseStorageType, inputTensorArray []GraphTensor, shape Shape /* not a class type */, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SplitTensorNumSplitsAxisName(tensor IMPSGraphTensor, numSplits uint, axis int, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	SplitTensorSplitSizesAxisName(tensor IMPSGraphTensor, splitSizes []foundation.Number, axis int, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	SplitTensorSplitSizesTensorAxisName(tensor IMPSGraphTensor, splitSizesTensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	SquareWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SquareRootWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SqueezeTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SqueezeTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SqueezeTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SqueezeTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	StackTensorsAxisName(inputTensors []GraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	StencilWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphStencilOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	StochasticGradientDescentWithLearningRateTensorValuesTensorGradientTensorName(learningRateTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	SubtractionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	TileGradientWithIncomingGradientTensorSourceTensorWithMultiplierName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, multiplier Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	TileTensorWithMultiplierName(tensor IMPSGraphTensor, multiplier Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	TopKWithSourceTensorAxisKName(source IMPSGraphTensor, axis int, k uint, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	TopKWithSourceTensorAxisTensorKTensorName(source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	TopKWithSourceTensorKName(source IMPSGraphTensor, k uint, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	TopKWithSourceTensorKTensorName(source IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor
	TopKWithGradientTensorSourceKName(gradient IMPSGraphTensor, source IMPSGraphTensor, k uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	TopKWithGradientTensorSourceKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	TopKWithGradientTensorSourceAxisKName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, k uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	TopKWithGradientTensorSourceAxisTensorKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	TransposeTensorPermutationName(tensor IMPSGraphTensor, permutation []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	TransposeTensorDimensionWithDimensionName(tensor IMPSGraphTensor, dimensionIndex uint, dimensionIndex2 uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	TruncateWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	VariableWithDataShapeDataTypeName(data objc.IObject /* cross-framework: NSData */, shape Shape /* not a class type */, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	VariableFromTensorWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	VarianceOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	VarianceOfTensorMeanTensorAxesName(tensor IMPSGraphTensor, meanTensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor
	WhileWithInitialInputsBeforeAfterName(initialInputs []GraphTensor, before GraphWhileBeforeBlock /* not a class type */, after GraphWhileAfterBlock /* not a class type */, name objc.IObject /* cross-framework: NSString */) []GraphTensor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Graph */
// Alloc allocates a new instance without initialization.
func (gc _GraphClass) Alloc() Graph {
	rv := objc.Send[Graph](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphClass) New() Graph {
	rv := objc.Send[Graph](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Graph) Init() Graph {
	rv := objc.Send[Graph](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Graph) Autorelease() Graph {
	rv := objc.Send[Graph](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraph creates a new Graph instance.
func NewGraph() Graph {
	return getGraphClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Graph */
// The optimized representation of a compute graph of operations and tensors.
//
// An MPSGraph is a symbolic representation of operations to be utilized to execute compute graphs on a device.


// The optimized representation of a compute graph of operations and tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph
type Graph struct {
	GraphObject
}

// GraphFrom constructs a [Graph] from an unsafe.Pointer.
//
// The optimized representation of a compute graph of operations and tensors.
func GraphFrom(ptr unsafe.Pointer) Graph {
	return Graph{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Graph */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Graph */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Graph */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Graph */

// Returns the absolute values of the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/absolute(with:name:)
func (g_ Graph) AbsoluteWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("absoluteWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: AbsoluteWithTensorName */


// Returns the absolute square of the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/absoluteSquare(tensor:name:)
func (g_ Graph) AbsoluteSquareWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("absoluteSquareWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: AbsoluteSquareWithTensorName */


// Creates operations to apply Adam optimization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/adam(currentLearningRate:beta1:beta2:epsilon:values:momentum:velocity:maximumVelocity:gradient:name:)
func (g_ Graph) AdamWithCurrentLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName(currentLearningRateTensor IMPSGraphTensor, beta1Tensor IMPSGraphTensor, beta2Tensor IMPSGraphTensor, epsilonTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, momentumTensor IMPSGraphTensor, velocityTensor IMPSGraphTensor, maximumVelocityTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("adamWithCurrentLearningRateTensor:beta1Tensor:beta2Tensor:epsilonTensor:valuesTensor:momentumTensor:velocityTensor:maximumVelocityTensor:gradientTensor:name:"), currentLearningRateTensor, beta1Tensor, beta2Tensor, epsilonTensor, valuesTensor, momentumTensor, velocityTensor, maximumVelocityTensor, gradientTensor, name)
	return rv
}/* debug [instance_methods/method]: AdamWithCurrentLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName */


// Creates operations to apply Adam optimization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/adam(learningRate:beta1:beta2:epsilon:beta1Power:beta2Power:values:momentum:velocity:maximumVelocity:gradient:name:)
func (g_ Graph) AdamWithLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorBeta1PowerTensorBeta2PowerTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName(learningRateTensor IMPSGraphTensor, beta1Tensor IMPSGraphTensor, beta2Tensor IMPSGraphTensor, epsilonTensor IMPSGraphTensor, beta1PowerTensor IMPSGraphTensor, beta2PowerTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, momentumTensor IMPSGraphTensor, velocityTensor IMPSGraphTensor, maximumVelocityTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("adamWithLearningRateTensor:beta1Tensor:beta2Tensor:epsilonTensor:beta1PowerTensor:beta2PowerTensor:valuesTensor:momentumTensor:velocityTensor:maximumVelocityTensor:gradientTensor:name:"), learningRateTensor, beta1Tensor, beta2Tensor, epsilonTensor, beta1PowerTensor, beta2PowerTensor, valuesTensor, momentumTensor, velocityTensor, maximumVelocityTensor, gradientTensor, name)
	return rv
}/* debug [instance_methods/method]: AdamWithLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorBeta1PowerTensorBeta2PowerTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName */


// Adds two input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/addition(_:_:name:)
func (g_ Graph) AdditionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("additionWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: AdditionWithPrimaryTensorSecondaryTensorName */


// The Stochastic gradient descent performs a gradient descent where, is gradient of error wrt variable this op directly writes to the variable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/applyStochasticGradientDescent(learningRate:variable:gradient:name:)
func (g_ Graph) ApplyStochasticGradientDescentWithLearningRateTensorVariableGradientTensorName(learningRateTensor IMPSGraphTensor, variable IMPSGraphVariableOp, gradientTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphOperation {
	rv := objc.Send[GraphOperation](g_.ID, objc.Sel("applyStochasticGradientDescentWithLearningRateTensor:variable:gradientTensor:name:"), learningRateTensor, variable, gradientTensor, name)
	return rv
}/* debug [instance_methods/method]: ApplyStochasticGradientDescentWithLearningRateTensorVariableGradientTensorName */


// Computes the indices that sort the elements of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/argSort(_:axis:descending:name:)
func (g_ Graph) ArgSortWithTensorAxisDescendingName(tensor IMPSGraphTensor, axis int, descending bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("argSortWithTensor:axis:descending:name:"), tensor, axis, descending, name)
	return rv
}/* debug [instance_methods/method]: ArgSortWithTensorAxisDescendingName */


// Computes the indices that sort the elements of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/argSort(_:axis:name:)
func (g_ Graph) ArgSortWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("argSortWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ArgSortWithTensorAxisName */


// Computes the indices that sort the elements of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/argSort(_:axisTensor:descending:name:)
func (g_ Graph) ArgSortWithTensorAxisTensorDescendingName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, descending bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("argSortWithTensor:axisTensor:descending:name:"), tensor, axisTensor, descending, name)
	return rv
}/* debug [instance_methods/method]: ArgSortWithTensorAxisTensorDescendingName */


// Computes the indices that sort the elements of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/argSort(_:axisTensor:name:)
func (g_ Graph) ArgSortWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("argSortWithTensor:axisTensor:name:"), tensor, axisTensor, name)
	return rv
}/* debug [instance_methods/method]: ArgSortWithTensorAxisTensorName */


// Creates an assign operation which writes at this point of execution of the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/assign(_:tensor:name:)
func (g_ Graph) AssignVariableWithValueOfTensorName(variable IMPSGraphTensor, tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphOperation {
	rv := objc.Send[GraphOperation](g_.ID, objc.Sel("assignVariable:withValueOfTensor:name:"), variable, tensor, name)
	return rv
}/* debug [instance_methods/method]: AssignVariableWithValueOfTensorName */


// Creates a 2D average-pooling operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/avgPooling2D(withSourceTensor:descriptor:name:)
func (g_ Graph) AvgPooling2DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("avgPooling2DWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: AvgPooling2DWithSourceTensorDescriptorName */


// Creates a 2D average pooling gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/avgPooling2DGradient(withGradientTensor:sourceTensor:descriptor:name:)
func (g_ Graph) AvgPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("avgPooling2DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: AvgPooling2DGradientWithGradientTensorSourceTensorDescriptorName */


// Creates a 4D average pooling operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/avgPooling4D(_:descriptor:name:)
func (g_ Graph) AvgPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("avgPooling4DWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: AvgPooling4DWithSourceTensorDescriptorName */


// Creates an average pooling gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/avgPooling4DGradient(_:source:descriptor:name:)
func (g_ Graph) AvgPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("avgPooling4DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: AvgPooling4DGradientWithGradientTensorSourceTensorDescriptorName */


// Computes the band part of an input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bandPart(_:numLower:numUpper:name:)
func (g_ Graph) BandPartWithTensorNumLowerNumUpperName(inputTensor IMPSGraphTensor, numLower int, numUpper int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bandPartWithTensor:numLower:numUpper:name:"), inputTensor, numLower, numUpper, name)
	return rv
}/* debug [instance_methods/method]: BandPartWithTensorNumLowerNumUpperName */


// Creates the band part operation and returns the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bandPart(_:numLowerTensor:numUpperTensor:name:)
func (g_ Graph) BandPartWithTensorNumLowerTensorNumUpperTensorName(inputTensor IMPSGraphTensor, numLowerTensor IMPSGraphTensor, numUpperTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bandPartWithTensor:numLowerTensor:numUpperTensor:name:"), inputTensor, numLowerTensor, numUpperTensor, name)
	return rv
}/* debug [instance_methods/method]: BandPartWithTensorNumLowerTensorNumUpperTensorName */


// Creates a batch-to-space operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/batchToSpace(_:spatialAxes:batchAxis:blockDimensions:usePixelShuffleOrder:name:)
func (g_ Graph) BatchToSpaceTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxes []foundation.Number, batchAxis int, blockDimensions []foundation.Number, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("batchToSpaceTensor:spatialAxes:batchAxis:blockDimensions:usePixelShuffleOrder:name:"), tensor, spatialAxes, batchAxis, blockDimensions, usePixelShuffleOrder, name)
	return rv
}/* debug [instance_methods/method]: BatchToSpaceTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName */


// Creates a batch-to-space operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/batchToSpace(_:spatialAxesTensor:batchAxisTensor:blockDimensionsTensor:usePixelShuffleOrder:name:)
func (g_ Graph) BatchToSpaceTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxesTensor IMPSGraphTensor, batchAxisTensor IMPSGraphTensor, blockDimensionsTensor IMPSGraphTensor, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("batchToSpaceTensor:spatialAxesTensor:batchAxisTensor:blockDimensionsTensor:usePixelShuffleOrder:name:"), tensor, spatialAxesTensor, batchAxisTensor, blockDimensionsTensor, usePixelShuffleOrder, name)
	return rv
}/* debug [instance_methods/method]: BatchToSpaceTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName */


// Returns the elementwise bitwise AND of binary representations of two integer tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseAND(_:_:name:)
func (g_ Graph) BitwiseANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bitwiseANDWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: BitwiseANDWithPrimaryTensorSecondaryTensorName */


// Returns the elementwise left-shifted binary representations of the primary integer by the secondary tensor amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseLeftShift(_:_:name:)
func (g_ Graph) BitwiseLeftShiftWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bitwiseLeftShiftWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: BitwiseLeftShiftWithPrimaryTensorSecondaryTensorName */


// Applies the bitwise NOT operation to the input tensor element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseNOT(_:name:)
func (g_ Graph) BitwiseNOTWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bitwiseNOTWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: BitwiseNOTWithTensorName */


// Returns the elementwise bitwise OR of binary representations of two integer tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseOR(_:_:name:)
func (g_ Graph) BitwiseORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bitwiseORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: BitwiseORWithPrimaryTensorSecondaryTensorName */


// Returns the population count of the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwisePopulationCount(_:name:)
func (g_ Graph) BitwisePopulationCountWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bitwisePopulationCountWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: BitwisePopulationCountWithTensorName */


// Returns the elementwise right-shifted binary representations of the primary integer by the secondary tensor amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseRightShift(_:_:name:)
func (g_ Graph) BitwiseRightShiftWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bitwiseRightShiftWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: BitwiseRightShiftWithPrimaryTensorSecondaryTensorName */


// Returns the elementwise bitwise XOR of binary representations of two integer tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseXOR(_:_:name:)
func (g_ Graph) BitwiseXORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bitwiseXORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: BitwiseXORWithPrimaryTensorSecondaryTensorName */


// Creates a BottomK operation and returns the value and indices tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bottomK(_:axis:k:name:)
func (g_ Graph) BottomKWithSourceTensorAxisKName(source IMPSGraphTensor, axis int, k uint, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("bottomKWithSourceTensor:axis:k:name:"), source, axis, k, name)
	return rv
}/* debug [instance_methods/method]: BottomKWithSourceTensorAxisKName */


// Creates a BottomK operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bottomK(_:axisTensor:kTensor:name:)
func (g_ Graph) BottomKWithSourceTensorAxisTensorKTensorName(source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("bottomKWithSourceTensor:axisTensor:kTensor:name:"), source, axisTensor, kTensor, name)
	return rv
}/* debug [instance_methods/method]: BottomKWithSourceTensorAxisTensorKTensorName */


// Creates a BottomKGradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bottomKGradient(_:source:axis:k:name:)
func (g_ Graph) BottomKWithGradientTensorSourceAxisKName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, k uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bottomKWithGradientTensor:source:axis:k:name:"), gradient, source, axis, k, name)
	return rv
}/* debug [instance_methods/method]: BottomKWithGradientTensorSourceAxisKName */


// Creates a BottomKGradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bottomKGradient(_:source:axisTensor:kTensor:name:)
func (g_ Graph) BottomKWithGradientTensorSourceAxisTensorKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("bottomKWithGradientTensor:source:axisTensor:kTensor:name:"), gradient, source, axisTensor, kTensor, name)
	return rv
}/* debug [instance_methods/method]: BottomKWithGradientTensorSourceAxisTensorKTensorName */


// Creates a broadcast operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/broadcast(_:shape:name:)
func (g_ Graph) BroadcastTensorToShapeName(tensor IMPSGraphTensor, shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("broadcastTensor:toShape:name:"), tensor, shape, name)
	return rv
}/* debug [instance_methods/method]: BroadcastTensorToShapeName */


// Creates a broadcast operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/broadcast(_:shapeTensor:name:)
func (g_ Graph) BroadcastTensorToShapeTensorName(tensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("broadcastTensor:toShapeTensor:name:"), tensor, shapeTensor, name)
	return rv
}/* debug [instance_methods/method]: BroadcastTensorToShapeTensorName */


// Creates an operation which invokes another executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/call(symbolName:inputTensors:outputTypes:name:)
func (g_ Graph) CallSymbolNameInputTensorsOutputTypesName(symbolName objc.IObject /* cross-framework: NSString */, inputTensors []GraphTensor, outputTypes []GraphType, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("callSymbolName:inputTensors:outputTypes:name:"), symbolName, inputTensors, outputTypes, name)
	return rv
}/* debug [instance_methods/method]: CallSymbolNameInputTensorsOutputTypesName */


// Creates a cast operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cast(_:to:name:)
func (g_ Graph) CastTensorToTypeName(tensor IMPSGraphTensor, type_ DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("castTensor:toType:name:"), tensor, type_, name)
	return rv
}/* debug [instance_methods/method]: CastTensorToTypeName */


// Applies the ceiling operation to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/ceil(with:name:)
func (g_ Graph) CeilWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("ceilWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: CeilWithTensorName */


// Clamps the values in the first tensor between the corresponding values in the minimum and maximum value tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/clamp(_:min:max:name:)
func (g_ Graph) ClampWithTensorMinValueTensorMaxValueTensorName(tensor IMPSGraphTensor, minValueTensor IMPSGraphTensor, maxValueTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("clampWithTensor:minValueTensor:maxValueTensor:name:"), tensor, minValueTensor, maxValueTensor, name)
	return rv
}/* debug [instance_methods/method]: ClampWithTensorMinValueTensorMaxValueTensorName */


// Creates a column to image operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/colToIm(_:outputShape:descriptor:name:)
func (g_ Graph) ColToImWithSourceTensorOutputShapeDescriptorName(source IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphImToColOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("colToImWithSourceTensor:outputShape:descriptor:name:"), source, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: ColToImWithSourceTensorOutputShapeDescriptorName */


// Compiles the graph for the given feeds to returns the target tensor values, ensuring all target operations would be executed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/compile(with:feeds:targetTensors:targetOperations:compilationDescriptor:)
func (g_ Graph) CompileWithDeviceFeedsTargetTensorsTargetOperationsCompilationDescriptor(device IMPSGraphDevice, feeds GraphTensorShapedTypeDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation, compilationDescriptor IMPSGraphCompilationDescriptor) IGraphExecutable {
	rv := objc.Send[GraphExecutable](g_.ID, objc.Sel("compileWithDevice:feeds:targetTensors:targetOperations:compilationDescriptor:"), device, feeds, targetTensors, targetOperations, compilationDescriptor)
	return rv
}/* debug [instance_methods/method]: CompileWithDeviceFeedsTargetTensorsTargetOperationsCompilationDescriptor */


// Creates a complex constant op with the MPSDataTypeComplexFloat32 data type and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/complexConstant(realPart:imaginaryPart:)
func (g_ Graph) ConstantWithRealPartImaginaryPart(realPart float64, imaginaryPart float64) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("constantWithRealPart:imaginaryPart:"), realPart, imaginaryPart)
	return rv
}/* debug [instance_methods/method]: ConstantWithRealPartImaginaryPart */


// Creates a complex constant operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/complexConstant(realPart:imaginaryPart:dataType:)
func (g_ Graph) ConstantWithRealPartImaginaryPartDataType(realPart float64, imaginaryPart float64, dataType DataType /* not a class type */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("constantWithRealPart:imaginaryPart:dataType:"), realPart, imaginaryPart, dataType)
	return rv
}/* debug [instance_methods/method]: ConstantWithRealPartImaginaryPartDataType */


// Creates a complex constant op with a given shape and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/complexConstant(realPart:imaginaryPart:shape:dataType:)
func (g_ Graph) ConstantWithRealPartImaginaryPartShapeDataType(realPart float64, imaginaryPart float64, shape Shape /* not a class type */, dataType DataType /* not a class type */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("constantWithRealPart:imaginaryPart:shape:dataType:"), realPart, imaginaryPart, shape, dataType)
	return rv
}/* debug [instance_methods/method]: ConstantWithRealPartImaginaryPartShapeDataType */


// Returns a complex tensor from the two input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/complexTensor(realTensor:imaginaryTensor:name:)
func (g_ Graph) ComplexTensorWithRealTensorImaginaryTensorName(realTensor IMPSGraphTensor, imaginaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("complexTensorWithRealTensor:imaginaryTensor:name:"), realTensor, imaginaryTensor, name)
	return rv
}/* debug [instance_methods/method]: ComplexTensorWithRealTensorImaginaryTensorName */


// Creates a concatenation operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/concatTensor(_:with:dimension:name:)
func (g_ Graph) ConcatTensorWithTensorDimensionName(tensor IMPSGraphTensor, tensor2 IMPSGraphTensor, dimensionIndex int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("concatTensor:withTensor:dimension:name:"), tensor, tensor2, dimensionIndex, name)
	return rv
}/* debug [instance_methods/method]: ConcatTensorWithTensorDimensionName */


// Creates a concatenation operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/concatTensors(_:dimension:interleave:name:)
func (g_ Graph) ConcatTensorsDimensionInterleaveName(tensors []GraphTensor, dimensionIndex int, interleave bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("concatTensors:dimension:interleave:name:"), tensors, dimensionIndex, interleave, name)
	return rv
}/* debug [instance_methods/method]: ConcatTensorsDimensionInterleaveName */


// Creates a concatenation operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/concatTensors(_:dimension:name:)
func (g_ Graph) ConcatTensorsDimensionName(tensors []GraphTensor, dimensionIndex int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("concatTensors:dimension:name:"), tensors, dimensionIndex, name)
	return rv
}/* debug [instance_methods/method]: ConcatTensorsDimensionName */


// Returns the complex conjugate of the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/conjugate(tensor:name:)
func (g_ Graph) ConjugateWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("conjugateWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: ConjugateWithTensorName */


// Creates a constant operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/constant(_:dataType:)
func (g_ Graph) ConstantWithScalarDataType(scalar float64, dataType DataType /* not a class type */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("constantWithScalar:dataType:"), scalar, dataType)
	return rv
}/* debug [instance_methods/method]: ConstantWithScalarDataType */


// Creates a constant op with a given shape and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/constant(_:shape:dataType:)-3wa0e
func (g_ Graph) ConstantWithScalarShapeDataType(scalar float64, shape Shape /* not a class type */, dataType DataType /* not a class type */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("constantWithScalar:shape:dataType:"), scalar, shape, dataType)
	return rv
}/* debug [instance_methods/method]: ConstantWithScalarShapeDataType */


// Creates a constant op with a given shape and data, and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/constant(_:shape:dataType:)-ylr4
func (g_ Graph) ConstantWithDataShapeDataType(data objc.IObject /* cross-framework: NSData */, shape Shape /* not a class type */, dataType DataType /* not a class type */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("constantWithData:shape:dataType:"), data, shape, dataType)
	return rv
}/* debug [instance_methods/method]: ConstantWithDataShapeDataType */


// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/controlDependency(with:dependentBlock:name:)
func (g_ Graph) ControlDependencyWithOperationsDependentBlockName(operations []GraphOperation, dependentBlock GraphControlFlowDependencyBlock /* not a class type */, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("controlDependencyWithOperations:dependentBlock:name:"), operations, dependentBlock, name)
	return rv
}/* debug [instance_methods/method]: ControlDependencyWithOperationsDependentBlockName */


// Creates a 2D (forward) convolution operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution2D(_:weights:descriptor:name:)
func (g_ Graph) Convolution2DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolution2DWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, name)
	return rv
}/* debug [instance_methods/method]: Convolution2DWithSourceTensorWeightsTensorDescriptorName */


// Creates a 2D convolution gradient operation with respect to the source tensor of the forward convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution2DDataGradient(_:weights:outputShape:forwardConvolutionDescriptor:name:)
func (g_ Graph) Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolution2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradient, weights, outputShape, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName */


// Creates a 2D convolution gradient operation with respect to the source tensor of the forward convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution2DDataGradient(_:weights:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g_ Graph) Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, weights IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolution2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), gradient, weights, outputShapeTensor, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName */


// Creates a 2D convolution gradient operation with respect to the weights tensor of the forward convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution2DWeightsGradient(_:source:outputShape:forwardConvolutionDescriptor:name:)
func (g_ Graph) Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolution2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradient, source, outputShape, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName */


// Creates a 2D convolution gradient operation with respect to weights tensor of forward convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution2DWeightsGradient(_:source:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g_ Graph) Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolution2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), gradient, source, outputShapeTensor, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName */


// Creates a 3D forward convolution operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution3D(_:weights:descriptor:name:)
func (g_ Graph) Convolution3DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolution3DWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, name)
	return rv
}/* debug [instance_methods/method]: Convolution3DWithSourceTensorWeightsTensorDescriptorName */


// Creates a 3D convolution gradient operation with respect to the source tensor of the forward convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution3DDataGradient(_:weights:outputShape:forwardConvolutionDescriptor:name:)
func (g_ Graph) Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolution3DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradient, weights, outputShape, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName */


// Creates a 3D convolution gradient operation with respect to the source tensor of the forward convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution3DDataGradient(_:weights:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g_ Graph) Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, weights IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolution3DDataGradientWithIncomingGradientTensor:weightsTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), gradient, weights, outputShapeTensor, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName */


// Creates a 3D convolution gradient operation with respect to the weights tensor of the forward convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution3DWeightsGradient(_:source:outputShape:forwardConvolutionDescriptor:name:)
func (g_ Graph) Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolution3DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradient, source, outputShape, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName */


// Creates a 3D convolution gradient operation with respect to the weights tensor of the forward convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution3DWeightsGradient(_:source:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g_ Graph) Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolution3DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), gradient, source, outputShapeTensor, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName */


// Creates a convolution transpose operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2D(_:weights:outputShape:descriptor:name:)
func (g_ Graph) ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolutionTranspose2DWithSourceTensor:weightsTensor:outputShape:descriptor:name:"), source, weights, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeDescriptorName */


// Creates a convolution transpose operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2D(_:weights:outputShapeTensor:descriptor:name:)
func (g_ Graph) ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolutionTranspose2DWithSourceTensor:weightsTensor:outputShapeTensor:descriptor:name:"), source, weights, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeTensorDescriptorName */


// Creates a convolution transpose gradient operation with respect to the source tensor of convolution transpose operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2DDataGradient(_:weights:outputShape:forwardConvolutionDescriptor:name:)
func (g_ Graph) ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolutionTranspose2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradient, weights, outputShape, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName */


// Creates a convolution transpose gradient operation with respect to the source tensor of convolution transpose operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2DDataGradient(_:weights:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g_ Graph) ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolutionTranspose2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), incomingGradient, weights, outputShape, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName */


// Creates a convolution transpose gradient operation with respect to the weights tensor of the convolution transpose operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2DWeightsGradient(_:weights:outputShape:forwardConvolutionDescriptor:name:)
func (g_ Graph) ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradientTensor IMPSGraphTensor, source IMPSGraphTensor, outputShape Shape /* not a class type */, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolutionTranspose2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradientTensor, source, outputShape, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName */


// Creates a convolution transpose gradient operation with respect to the weights tensor of the convolution transpose operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2DWeightsGradient(_:weights:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g_ Graph) ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradientTensor IMPSGraphTensor, source IMPSGraphTensor, outputShape IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("convolutionTranspose2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), incomingGradientTensor, source, outputShape, forwardConvolutionDescriptor, name)
	return rv
}/* debug [instance_methods/method]: ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName */


// Creates a get-coordindate operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/coordinate(alongAxis:withShape:name:)
func (g_ Graph) CoordinateAlongAxisWithShapeName(axis int, shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("coordinateAlongAxis:withShape:name:"), axis, shape, name)
	return rv
}/* debug [instance_methods/method]: CoordinateAlongAxisWithShapeName */


// Creates a get-coordindate operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/coordinate(alongAxis:withShapeTensor:name:)
func (g_ Graph) CoordinateAlongAxisWithShapeTensorName(axis int, shapeTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("coordinateAlongAxis:withShapeTensor:name:"), axis, shapeTensor, name)
	return rv
}/* debug [instance_methods/method]: CoordinateAlongAxisWithShapeTensorName */


// Creates a get-coordindate operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/coordinate(alongAxisTensor:withShape:name:)
func (g_ Graph) CoordinateAlongAxisTensorWithShapeName(axisTensor IMPSGraphTensor, shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("coordinateAlongAxisTensor:withShape:name:"), axisTensor, shape, name)
	return rv
}/* debug [instance_methods/method]: CoordinateAlongAxisTensorWithShapeName */


// Creates a get-coordindate operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/coordinate(alongAxisTensor:withShapeTensor:name:)
func (g_ Graph) CoordinateAlongAxisTensorWithShapeTensorName(axisTensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("coordinateAlongAxisTensor:withShapeTensor:name:"), axisTensor, shapeTensor, name)
	return rv
}/* debug [instance_methods/method]: CoordinateAlongAxisTensorWithShapeTensorName */


// Computes the cumulative maximum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMaximum(_:axis:exclusive:reverse:name:)
func (g_ Graph) CumulativeMaximumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeMaximumWithTensor:axis:exclusive:reverse:name:"), tensor, axis, exclusive, reverse, name)
	return rv
}/* debug [instance_methods/method]: CumulativeMaximumWithTensorAxisExclusiveReverseName */


// Computes the cumulative maximum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMaximum(_:axis:name:)
func (g_ Graph) CumulativeMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeMaximumWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: CumulativeMaximumWithTensorAxisName */


// Computes the cumulative maximum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMaximum(_:axisTensor:exclusive:reverse:name:)
func (g_ Graph) CumulativeMaximumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeMaximumWithTensor:axisTensor:exclusive:reverse:name:"), tensor, axisTensor, exclusive, reverse, name)
	return rv
}/* debug [instance_methods/method]: CumulativeMaximumWithTensorAxisTensorExclusiveReverseName */


// Computes the cumulative maximum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMaximum(_:axisTensor:name:)
func (g_ Graph) CumulativeMaximumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeMaximumWithTensor:axisTensor:name:"), tensor, axisTensor, name)
	return rv
}/* debug [instance_methods/method]: CumulativeMaximumWithTensorAxisTensorName */


// Computes the cumulative minimum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMinimum(_:axis:exclusive:reverse:name:)
func (g_ Graph) CumulativeMinimumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeMinimumWithTensor:axis:exclusive:reverse:name:"), tensor, axis, exclusive, reverse, name)
	return rv
}/* debug [instance_methods/method]: CumulativeMinimumWithTensorAxisExclusiveReverseName */


// Computes the cumulative minimum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMinimum(_:axis:name:)
func (g_ Graph) CumulativeMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeMinimumWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: CumulativeMinimumWithTensorAxisName */


// Computes the cumulative minimum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMinimum(_:axisTensor:exclusive:reverse:name:)
func (g_ Graph) CumulativeMinimumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeMinimumWithTensor:axisTensor:exclusive:reverse:name:"), tensor, axisTensor, exclusive, reverse, name)
	return rv
}/* debug [instance_methods/method]: CumulativeMinimumWithTensorAxisTensorExclusiveReverseName */


// Computes the cumulative minimum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMinimum(_:axisTensor:name:)
func (g_ Graph) CumulativeMinimumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeMinimumWithTensor:axisTensor:name:"), tensor, axisTensor, name)
	return rv
}/* debug [instance_methods/method]: CumulativeMinimumWithTensorAxisTensorName */


// Computes the cumulative product of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeProduct(_:axis:exclusive:reverse:name:)
func (g_ Graph) CumulativeProductWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeProductWithTensor:axis:exclusive:reverse:name:"), tensor, axis, exclusive, reverse, name)
	return rv
}/* debug [instance_methods/method]: CumulativeProductWithTensorAxisExclusiveReverseName */


// Computes the cumulative product of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeProduct(_:axis:name:)
func (g_ Graph) CumulativeProductWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeProductWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: CumulativeProductWithTensorAxisName */


// Computes the cumulative product of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeProduct(_:axisTensor:exclusive:reverse:name:)
func (g_ Graph) CumulativeProductWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeProductWithTensor:axisTensor:exclusive:reverse:name:"), tensor, axisTensor, exclusive, reverse, name)
	return rv
}/* debug [instance_methods/method]: CumulativeProductWithTensorAxisTensorExclusiveReverseName */


// Computes the cumulative product of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeProduct(_:axisTensor:name:)
func (g_ Graph) CumulativeProductWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeProductWithTensor:axisTensor:name:"), tensor, axisTensor, name)
	return rv
}/* debug [instance_methods/method]: CumulativeProductWithTensorAxisTensorName */


// Computes the cumulative sum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeSum(_:axis:exclusive:reverse:name:)
func (g_ Graph) CumulativeSumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeSumWithTensor:axis:exclusive:reverse:name:"), tensor, axis, exclusive, reverse, name)
	return rv
}/* debug [instance_methods/method]: CumulativeSumWithTensorAxisExclusiveReverseName */


// Computes the cumulative sum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeSum(_:axis:name:)
func (g_ Graph) CumulativeSumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeSumWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: CumulativeSumWithTensorAxisName */


// Computes the cumulative sum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeSum(_:axisTensor:exclusive:reverse:name:)
func (g_ Graph) CumulativeSumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeSumWithTensor:axisTensor:exclusive:reverse:name:"), tensor, axisTensor, exclusive, reverse, name)
	return rv
}/* debug [instance_methods/method]: CumulativeSumWithTensorAxisTensorExclusiveReverseName */


// Computes the cumulative sum of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeSum(_:axisTensor:name:)
func (g_ Graph) CumulativeSumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("cumulativeSumWithTensor:axisTensor:name:"), tensor, axisTensor, name)
	return rv
}/* debug [instance_methods/method]: CumulativeSumWithTensorAxisTensorName */


// Creates a depth-to-space2D operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depth(toSpace2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:)
func (g_ Graph) DepthToSpace2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("depthToSpace2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:"), tensor, widthAxis, heightAxis, depthAxis, blockSize, usePixelShuffleOrder, name)
	return rv
}/* debug [instance_methods/method]: DepthToSpace2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName */


// Creates a depth-to-space2D operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depth(toSpace2DTensor:widthAxisTensor:heightAxisTensor:depthAxisTensor:blockSize:usePixelShuffleOrder:name:)
func (g_ Graph) DepthToSpace2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxisTensor IMPSGraphTensor, heightAxisTensor IMPSGraphTensor, depthAxisTensor IMPSGraphTensor, blockSize uint, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("depthToSpace2DTensor:widthAxisTensor:heightAxisTensor:depthAxisTensor:blockSize:usePixelShuffleOrder:name:"), tensor, widthAxisTensor, heightAxisTensor, depthAxisTensor, blockSize, usePixelShuffleOrder, name)
	return rv
}/* debug [instance_methods/method]: DepthToSpace2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName */


// Creates a 2D-depthwise convolution operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution2D(_:weights:descriptor:name:)
func (g_ Graph) DepthwiseConvolution2DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("depthwiseConvolution2DWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, name)
	return rv
}/* debug [instance_methods/method]: DepthwiseConvolution2DWithSourceTensorWeightsTensorDescriptorName */


// Creates a 2D-depthwise convolution gradient for data operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution2DDataGradient(_:weights:outputShape:descriptor:name:)
func (g_ Graph) DepthwiseConvolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("depthwiseConvolution2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:descriptor:name:"), incomingGradient, weights, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: DepthwiseConvolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName */


// Creates a 2D-depthwise convolution gradient for weights operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution2DWeightsGradient(_:source:outputShape:descriptor:name:)
func (g_ Graph) DepthwiseConvolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("depthwiseConvolution2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:descriptor:name:"), incomingGradient, source, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: DepthwiseConvolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName */


// Creates a 3D depthwise convolution operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution3D(_:weights:descriptor:name:)
func (g_ Graph) DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("depthwiseConvolution3DWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, name)
	return rv
}/* debug [instance_methods/method]: DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName */


// Creates a 3D depthwise convolution gradient for data operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution3DDataGradient(_:weights:outputShape:descriptor:name:)
func (g_ Graph) DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("depthwiseConvolution3DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:descriptor:name:"), incomingGradient, weights, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName */


// Creates a 3D depthwise convolution gradient for weights operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution3DWeightsGradient(_:source:outputShape:descriptor:name:)
func (g_ Graph) DepthwiseConvolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("depthwiseConvolution3DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:descriptor:name:"), incomingGradient, source, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: DepthwiseConvolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName */


// Creates a vector lookup-table based quantization operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:LUTTensor:axis:name:)
func (g_ Graph) DequantizeTensorLUTTensorAxisName(tensor IMPSGraphTensor, LUTTensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("dequantizeTensor:LUTTensor:axis:name:"), tensor, LUTTensor, axis, name)
	return rv
}/* debug [instance_methods/method]: DequantizeTensorLUTTensorAxisName */


// Creates a lookup-table based quantization operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:LUTTensor:name:)
func (g_ Graph) DequantizeTensorLUTTensorName(tensor IMPSGraphTensor, LUTTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("dequantizeTensor:LUTTensor:name:"), tensor, LUTTensor, name)
	return rv
}/* debug [instance_methods/method]: DequantizeTensorLUTTensorName */


// Creates Dequantize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:scale:zeroPoint:dataType:name:)
func (g_ Graph) DequantizeTensorScaleZeroPointDataTypeName(tensor IMPSGraphTensor, scale float64, zeroPoint float64, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("dequantizeTensor:scale:zeroPoint:dataType:name:"), tensor, scale, zeroPoint, dataType, name)
	return rv
}/* debug [instance_methods/method]: DequantizeTensorScaleZeroPointDataTypeName */


// Creates a dequantize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:scaleTensor:dataType:name:)
func (g_ Graph) DequantizeTensorScaleTensorDataTypeName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("dequantizeTensor:scaleTensor:dataType:name:"), tensor, scaleTensor, dataType, name)
	return rv
}/* debug [instance_methods/method]: DequantizeTensorScaleTensorDataTypeName */


// Creates Dequantize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:scaleTensor:zeroPoint:dataType:axis:name:)
func (g_ Graph) DequantizeTensorScaleTensorZeroPointDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPoint float64, dataType DataType /* not a class type */, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("dequantizeTensor:scaleTensor:zeroPoint:dataType:axis:name:"), tensor, scaleTensor, zeroPoint, dataType, axis, name)
	return rv
}/* debug [instance_methods/method]: DequantizeTensorScaleTensorZeroPointDataTypeAxisName */


// Creates a dequantize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:scaleTensor:zeroPointTensor:dataType:axis:name:)
func (g_ Graph) DequantizeTensorScaleTensorZeroPointTensorDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType DataType /* not a class type */, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("dequantizeTensor:scaleTensor:zeroPointTensor:dataType:axis:name:"), tensor, scaleTensor, zeroPointTensor, dataType, axis, name)
	return rv
}/* debug [instance_methods/method]: DequantizeTensorScaleTensorZeroPointTensorDataTypeAxisName */


// Creates a dequantize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:scaleTensor:zeroPointTensor:dataType:name:)
func (g_ Graph) DequantizeTensorScaleTensorZeroPointTensorDataTypeName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("dequantizeTensor:scaleTensor:zeroPointTensor:dataType:name:"), tensor, scaleTensor, zeroPointTensor, dataType, name)
	return rv
}/* debug [instance_methods/method]: DequantizeTensorScaleTensorZeroPointTensorDataTypeName */


// Divides the first input tensor by the second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/division(_:_:name:)
func (g_ Graph) DivisionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("divisionWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: DivisionWithPrimaryTensorSecondaryTensorName */


// Divides the first input tensor by the second, with the result being 0 if the denominator is 0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/divisionNoNaN(_:_:name:)
func (g_ Graph) DivisionNoNaNWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("divisionNoNaNWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: DivisionNoNaNWithPrimaryTensorSecondaryTensorName */


// Creates a dropout operation and returns the result
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dropout(_:rate:name:)-16cq4
func (g_ Graph) DropoutTensorRateTensorName(tensor IMPSGraphTensor, rate IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("dropoutTensor:rateTensor:name:"), tensor, rate, name)
	return rv
}/* debug [instance_methods/method]: DropoutTensorRateTensorName */


// Creates a dropout operation and returns the result
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dropout(_:rate:name:)-6hvf3
func (g_ Graph) DropoutTensorRateName(tensor IMPSGraphTensor, rate float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("dropoutTensor:rate:name:"), tensor, rate, name)
	return rv
}/* debug [instance_methods/method]: DropoutTensorRateName */


// Encodes the graph for the given feeds to returns the target tensor values in the results dictionary provided by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/encode(to:feeds:targetOperations:resultsDictionary:executionDescriptor:)
func (g_ Graph) EncodeToCommandBufferFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandBuffer metalperformanceshaders.CommandBuffer, feeds GraphTensorDataDictionary /* not a class type */, targetOperations []GraphOperation, resultsDictionary GraphTensorDataDictionary /* not a class type */, executionDescriptor IMPSGraphExecutionDescriptor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("encodeToCommandBuffer:feeds:targetOperations:resultsDictionary:executionDescriptor:"), commandBuffer, feeds, targetOperations, resultsDictionary, executionDescriptor)
}/* debug [instance_methods/method]: EncodeToCommandBufferFeedsTargetOperationsResultsDictionaryExecutionDescriptor */


// Encodes the graph for the given feeds to returns the target tensor values, ensuring all target operations also executed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/encode(to:feeds:targetTensors:targetOperations:executionDescriptor:)
func (g_ Graph) EncodeToCommandBufferFeedsTargetTensorsTargetOperationsExecutionDescriptor(commandBuffer metalperformanceshaders.CommandBuffer, feeds GraphTensorDataDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) GraphTensorDataDictionary /* not a class type */ {
	rv := objc.Send[GraphTensorDataDictionary](g_.ID, objc.Sel("encodeToCommandBuffer:feeds:targetTensors:targetOperations:executionDescriptor:"), commandBuffer, feeds, targetTensors, targetOperations, executionDescriptor)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferFeedsTargetTensorsTargetOperationsExecutionDescriptor */


// Returns the elementwise equality check of the input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/equal(_:_:name:)
func (g_ Graph) EqualWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("equalWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: EqualWithPrimaryTensorSecondaryTensorName */


// Creates an expand-dimensions operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/expandDims(_:axes:name:)
func (g_ Graph) ExpandDimsOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("expandDimsOfTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: ExpandDimsOfTensorAxesName */


// Creates an expand-dimensions operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/expandDims(_:axesTensor:name:)
func (g_ Graph) ExpandDimsOfTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("expandDimsOfTensor:axesTensor:name:"), tensor, axesTensor, name)
	return rv
}/* debug [instance_methods/method]: ExpandDimsOfTensorAxesTensorName */


// Creates an expand-dimensions operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/expandDims(_:axis:name:)
func (g_ Graph) ExpandDimsOfTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("expandDimsOfTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ExpandDimsOfTensorAxisName */


// Applies the natural exponent to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/exponent(with:name:)
func (g_ Graph) ExponentWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("exponentWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: ExponentWithTensorName */


// Applies an exponent with base 10 to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/exponentBase10(with:name:)
func (g_ Graph) ExponentBase10WithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("exponentBase10WithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: ExponentBase10WithTensorName */


// Applies an exponent with base 2 to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/exponentBase2(with:name:)
func (g_ Graph) ExponentBase2WithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("exponentBase2WithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: ExponentBase2WithTensorName */


// Creates a fast Fourier transform operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/fastFourierTransform(_:axes:descriptor:name:)
func (g_ Graph) FastFourierTransformWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.Number, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("fastFourierTransformWithTensor:axes:descriptor:name:"), tensor, axes, descriptor, name)
	return rv
}/* debug [instance_methods/method]: FastFourierTransformWithTensorAxesDescriptorName */


// Creates a fast Fourier transform operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/fastFourierTransform(_:axesTensor:descriptor:name:)
func (g_ Graph) FastFourierTransformWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("fastFourierTransformWithTensor:axesTensor:descriptor:name:"), tensor, axesTensor, descriptor, name)
	return rv
}/* debug [instance_methods/method]: FastFourierTransformWithTensorAxesTensorDescriptorName */


// Creates a flatten2D operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/flatten2D(_:axis:name:)
func (g_ Graph) Flatten2DTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("flatten2DTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: Flatten2DTensorAxisName */


// Creates a flatten2D operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/flatten2D(_:axisTensor:name:)
func (g_ Graph) Flatten2DTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("flatten2DTensor:axisTensor:name:"), tensor, axisTensor, name)
	return rv
}/* debug [instance_methods/method]: Flatten2DTensorAxisTensorName */


// Applies the floor operation to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/floor(with:name:)
func (g_ Graph) FloorWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("floorWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: FloorWithTensorName */


// Returns the remainder of floor divison between the primary and secondary tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/floorModulo(_:_:name:)
func (g_ Graph) FloorModuloWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("floorModuloWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: FloorModuloWithPrimaryTensorSecondaryTensorName */


// Adds a for loop operation, The lower and upper bounds specify a half-open range: the range includes the lower bound but does not include the upper bound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/for(lowerBound:upperBound:step:initialBodyArguments:body:name:)
func (g_ Graph) ForLoopWithLowerBoundUpperBoundStepInitialBodyArgumentsBodyName(lowerBound IMPSGraphTensor, upperBound IMPSGraphTensor, step IMPSGraphTensor, initialBodyArguments []GraphTensor, body GraphForLoopBodyBlock /* not a class type */, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("forLoopWithLowerBound:upperBound:step:initialBodyArguments:body:name:"), lowerBound, upperBound, step, initialBodyArguments, body, name)
	return rv
}/* debug [instance_methods/method]: ForLoopWithLowerBoundUpperBoundStepInitialBodyArgumentsBodyName */


// Adds a for loop operation, with a specific number of iterations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/for(numberOfIterations:initialBodyArguments:body:name:)
func (g_ Graph) ForLoopWithNumberOfIterationsInitialBodyArgumentsBodyName(numberOfIterations IMPSGraphTensor, initialBodyArguments []GraphTensor, body GraphForLoopBodyBlock /* not a class type */, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("forLoopWithNumberOfIterations:initialBodyArguments:body:name:"), numberOfIterations, initialBodyArguments, body, name)
	return rv
}/* debug [instance_methods/method]: ForLoopWithNumberOfIterationsInitialBodyArgumentsBodyName */


// Creates a Gather operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/gather(withUpdatesTensor:indicesTensor:axis:batchDimensions:name:)
func (g_ Graph) GatherWithUpdatesTensorIndicesTensorAxisBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, axis uint, batchDimensions uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("gatherWithUpdatesTensor:indicesTensor:axis:batchDimensions:name:"), updatesTensor, indicesTensor, axis, batchDimensions, name)
	return rv
}/* debug [instance_methods/method]: GatherWithUpdatesTensorIndicesTensorAxisBatchDimensionsName */


// Creates a GatherAlongAxis operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/gatherAlongAxis(_:updates:indices:name:)
func (g_ Graph) GatherAlongAxisWithUpdatesTensorIndicesTensorName(axis int, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("gatherAlongAxis:withUpdatesTensor:indicesTensor:name:"), axis, updatesTensor, indicesTensor, name)
	return rv
}/* debug [instance_methods/method]: GatherAlongAxisWithUpdatesTensorIndicesTensorName */


// Creates a GatherAlongAxis operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/gatherAlongAxisTensor(_:updates:indices:name:)
func (g_ Graph) GatherAlongAxisTensorWithUpdatesTensorIndicesTensorName(axisTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("gatherAlongAxisTensor:withUpdatesTensor:indicesTensor:name:"), axisTensor, updatesTensor, indicesTensor, name)
	return rv
}/* debug [instance_methods/method]: GatherAlongAxisTensorWithUpdatesTensorIndicesTensorName */


// Creates a GatherND operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/gatherND(withUpdatesTensor:indicesTensor:batchDimensions:name:)
func (g_ Graph) GatherNDWithUpdatesTensorIndicesTensorBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, batchDimensions uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("gatherNDWithUpdatesTensor:indicesTensor:batchDimensions:name:"), updatesTensor, indicesTensor, batchDimensions, name)
	return rv
}/* debug [instance_methods/method]: GatherNDWithUpdatesTensorIndicesTensorBatchDimensionsName */


// Calculates a partial derivative of primaryTensor with respect to the tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/gradients(of:with:name:)
func (g_ Graph) GradientForPrimaryTensorWithTensorsName(primaryTensor IMPSGraphTensor, tensors []GraphTensor, name objc.IObject /* cross-framework: NSString */) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("gradientForPrimaryTensor:withTensors:name:"), primaryTensor, tensors, name)
	return rv
}/* debug [instance_methods/method]: GradientForPrimaryTensorWithTensorsName */


// Checks in an elementwise manner if the first input tensor is greater than the second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/greaterThan(_:_:name:)
func (g_ Graph) GreaterThanWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("greaterThanWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: GreaterThanWithPrimaryTensorSecondaryTensorName */


// Checks in an elementwise manner if the first input tensor is greater than or equal to the second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/greaterThanOrEqualTo(_:_:name:)
func (g_ Graph) GreaterThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("greaterThanOrEqualToWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: GreaterThanOrEqualToWithPrimaryTensorSecondaryTensorName */


// Creates a GRU operation and returns the value and optionally the training state tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRU(_:recurrentWeight:inputWeight:bias:descriptor:name:)
func (g_ Graph) GRUWithSourceTensorRecurrentWeightInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("GRUWithSourceTensor:recurrentWeight:inputWeight:bias:descriptor:name:"), source, recurrentWeight, inputWeight, bias, descriptor, name)
	return rv
}/* debug [instance_methods/method]: GRUWithSourceTensorRecurrentWeightInputWeightBiasDescriptorName */


// Creates a GRU operation and returns the value and optionally the training state tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRU(_:recurrentWeight:inputWeight:bias:initState:descriptor:name:)
func (g_ Graph) GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("GRUWithSourceTensor:recurrentWeight:inputWeight:bias:initState:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, descriptor, name)
	return rv
}/* debug [instance_methods/method]: GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName */


// Creates a GRU operation and returns the value and optionally the training state tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRU(_:recurrentWeight:inputWeight:bias:initState:mask:secondaryBias:descriptor:name:)
func (g_ Graph) GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, secondaryBias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("GRUWithSourceTensor:recurrentWeight:inputWeight:bias:initState:mask:secondaryBias:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, mask, secondaryBias, descriptor, name)
	return rv
}/* debug [instance_methods/method]: GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName */


// Creates a GRU gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRUGradients(_:recurrentWeight:sourceGradient:zState:outputFwd:inputWeight:bias:descriptor:name:)
func (g_ Graph) GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("GRUGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:outputFwd:inputWeight:bias:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, outputFwd, inputWeight, bias, descriptor, name)
	return rv
}/* debug [instance_methods/method]: GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasDescriptorName */


// Creates a GRU gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRUGradients(_:recurrentWeight:sourceGradient:zState:outputFwd:inputWeight:bias:initState:descriptor:name:)
func (g_ Graph) GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("GRUGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:outputFwd:inputWeight:bias:initState:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, outputFwd, inputWeight, bias, initState, descriptor, name)
	return rv
}/* debug [instance_methods/method]: GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasInitStateDescriptorName */


// Creates a GRU gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRUGradients(_:recurrentWeight:sourceGradient:zState:outputFwd:stateGradient:inputWeight:bias:initState:mask:secondaryBias:descriptor:name:)
func (g_ Graph) GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdStateGradientInputWeightBiasInitStateMaskSecondaryBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, stateGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, secondaryBias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("GRUGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:outputFwd:stateGradient:inputWeight:bias:initState:mask:secondaryBias:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, outputFwd, stateGradient, inputWeight, bias, initState, mask, secondaryBias, descriptor, name)
	return rv
}/* debug [instance_methods/method]: GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdStateGradientInputWeightBiasInitStateMaskSecondaryBiasDescriptorName */


// Computes the hamming distance of two input tensors with support for broadcasting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/HammingDistance(primary:secondary:resultDataType:name:)
func (g_ Graph) HammingDistanceWithPrimaryTensorSecondaryTensorResultDataTypeName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, resultDataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("HammingDistanceWithPrimaryTensor:secondaryTensor:resultDataType:name:"), primaryTensor, secondaryTensor, resultDataType, name)
	return rv
}/* debug [instance_methods/method]: HammingDistanceWithPrimaryTensorSecondaryTensorResultDataTypeName */


// Creates a Hermitean-to-real fast Fourier transform operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/HermiteanToRealFFT(_:axes:descriptor:name:)
func (g_ Graph) HermiteanToRealFFTWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.Number, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("HermiteanToRealFFTWithTensor:axes:descriptor:name:"), tensor, axes, descriptor, name)
	return rv
}/* debug [instance_methods/method]: HermiteanToRealFFTWithTensorAxesDescriptorName */


// Creates a Hermitean-to-real fast Fourier transform operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/HermiteanToRealFFT(_:axesTensor:descriptor:name:)
func (g_ Graph) HermiteanToRealFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("HermiteanToRealFFTWithTensor:axesTensor:descriptor:name:"), tensor, axesTensor, descriptor, name)
	return rv
}/* debug [instance_methods/method]: HermiteanToRealFFTWithTensorAxesTensorDescriptorName */


// Copies the input tensor values into the output, behaving as an identity operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/identity(with:name:)
func (g_ Graph) IdentityWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("identityWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: IdentityWithTensorName */


// Adds an if-then-else operation to the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/if(_:then:else:name:)
func (g_ Graph) IfWithPredicateTensorThenBlockElseBlockName(predicateTensor IMPSGraphTensor, thenBlock GraphIfThenElseBlock /* not a class type */, elseBlock GraphIfThenElseBlock /* not a class type */, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("ifWithPredicateTensor:thenBlock:elseBlock:name:"), predicateTensor, thenBlock, elseBlock, name)
	return rv
}/* debug [instance_methods/method]: IfWithPredicateTensorThenBlockElseBlockName */


// Returns the imaginary part of a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/imaginaryPartOfTensor(tensor:name:)
func (g_ Graph) ImaginaryPartOfTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("imaginaryPartOfTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: ImaginaryPartOfTensorName */


// Creates an imToCol operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/imToCol(_:descriptor:name:)
func (g_ Graph) ImToColWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphImToColOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("imToColWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: ImToColWithSourceTensorDescriptorName */


// Computes the inverse of an input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/inverse(input:name:)
func (g_ Graph) InverseOfTensorName(inputTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("inverseOfTensor:name:"), inputTensor, name)
	return rv
}/* debug [instance_methods/method]: InverseOfTensorName */


// Checks if the input tensor elements are finite or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/isFinite(with:name:)
func (g_ Graph) IsFiniteWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("isFiniteWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: IsFiniteWithTensorName */


// Checks if the input tensor elements are infinite or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/isInfinite(with:name:)
func (g_ Graph) IsInfiniteWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("isInfiniteWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: IsInfiniteWithTensorName */


// Checks if the input tensor elements are or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/isNaN(with:name:)
func (g_ Graph) IsNaNWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("isNaNWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: IsNaNWithTensorName */


// Creates a 4D L2-norm pooling operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/L2NormPooling4D(_:descriptor:name:)
func (g_ Graph) L2NormPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("L2NormPooling4DWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: L2NormPooling4DWithSourceTensorDescriptorName */


// Creates a L2-Norm pooling gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/L2NormPooling4DGradient(_:source:descriptor:name:)
func (g_ Graph) L2NormPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("L2NormPooling4DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: L2NormPooling4DGradientWithGradientTensorSourceTensorDescriptorName */


// Computes the leaky rectified linear unit (ReLU) activation function on the input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/leakyReLU(with:alpha:name:)
func (g_ Graph) LeakyReLUWithTensorAlphaName(tensor IMPSGraphTensor, alpha float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("leakyReLUWithTensor:alpha:name:"), tensor, alpha, name)
	return rv
}/* debug [instance_methods/method]: LeakyReLUWithTensorAlphaName */


// Computes the leaky rectified linear unit (ReLU) activation function on the input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/leakyReLU(with:alphaTensor:name:)
func (g_ Graph) LeakyReLUWithTensorAlphaTensorName(tensor IMPSGraphTensor, alphaTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("leakyReLUWithTensor:alphaTensor:name:"), tensor, alphaTensor, name)
	return rv
}/* debug [instance_methods/method]: LeakyReLUWithTensorAlphaTensorName */


// Computes the gradient of the leaky rectified linear unit (ReLU) activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/leakyReLUGradient(withIncomingGradient:sourceTensor:alphaTensor:name:)
func (g_ Graph) LeakyReLUGradientWithIncomingGradientSourceTensorAlphaTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, alphaTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("leakyReLUGradientWithIncomingGradient:sourceTensor:alphaTensor:name:"), gradient, source, alphaTensor, name)
	return rv
}/* debug [instance_methods/method]: LeakyReLUGradientWithIncomingGradientSourceTensorAlphaTensorName */


// Checks in an elementwise manner if the first input tensor is less than the second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/lessThan(_:_:name:)
func (g_ Graph) LessThanWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("lessThanWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: LessThanWithPrimaryTensorSecondaryTensorName */


// Checks in an elementwise manner if the first input tensor is less than or equal to the second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/lessThanOrEqualTo(_:_:name:)
func (g_ Graph) LessThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("lessThanOrEqualToWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: LessThanOrEqualToWithPrimaryTensorSecondaryTensorName */


// Computes the natural logarithm to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logarithm(with:name:)
func (g_ Graph) LogarithmWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("logarithmWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: LogarithmWithTensorName */


// Computes the logarithm with base 10 to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logarithmBase10(with:name:)
func (g_ Graph) LogarithmBase10WithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("logarithmBase10WithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: LogarithmBase10WithTensorName */


// Computes the logarithm with base 2 to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logarithmBase2(with:name:)
func (g_ Graph) LogarithmBase2WithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("logarithmBase2WithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: LogarithmBase2WithTensorName */


// Returns the elementwise logical AND of the input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalAND(_:_:name:)
func (g_ Graph) LogicalANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("logicalANDWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: LogicalANDWithPrimaryTensorSecondaryTensorName */


// Returns the elementwise logical NAND of the input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalNAND(_:_:name:)
func (g_ Graph) LogicalNANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("logicalNANDWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: LogicalNANDWithPrimaryTensorSecondaryTensorName */


// Returns the elementwise logical NOR of the input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalNOR(_:_:name:)
func (g_ Graph) LogicalNORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("logicalNORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: LogicalNORWithPrimaryTensorSecondaryTensorName */


// Returns the elementwise logical OR of the input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalOR(_:_:name:)
func (g_ Graph) LogicalORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("logicalORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: LogicalORWithPrimaryTensorSecondaryTensorName */


// Returns the elementwise logical XNOR of the input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalXNOR(_:_:name:)
func (g_ Graph) LogicalXNORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("logicalXNORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: LogicalXNORWithPrimaryTensorSecondaryTensorName */


// Returns the elementwise logical XOR of the input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalXOR(_:_:name:)
func (g_ Graph) LogicalXORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("logicalXORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: LogicalXORWithPrimaryTensorSecondaryTensorName */


// Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and the training state tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTM(_:recurrentWeight:initState:initCell:descriptor:name:)
func (g_ Graph) LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("LSTMWithSourceTensor:recurrentWeight:initState:initCell:descriptor:name:"), source, recurrentWeight, initState, initCell, descriptor, name)
	return rv
}/* debug [instance_methods/method]: LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName */


// Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and the training state tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTM(_:recurrentWeight:inputWeight:bias:initState:initCell:descriptor:name:)
func (g_ Graph) LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("LSTMWithSourceTensor:recurrentWeight:inputWeight:bias:initState:initCell:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, initCell, descriptor, name)
	return rv
}/* debug [instance_methods/method]: LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName */


// Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and the training state tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTM(_:recurrentWeight:inputWeight:bias:initState:initCell:mask:peephole:descriptor:name:)
func (g_ Graph) LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, peephole IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("LSTMWithSourceTensor:recurrentWeight:inputWeight:bias:initState:initCell:mask:peephole:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, initCell, mask, peephole, descriptor, name)
	return rv
}/* debug [instance_methods/method]: LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName */


// Creates an LSTM gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTMGradients(_:recurrentWeight:sourceGradient:zState:cellOutputFwd:descriptor:name:)
func (g_ Graph) LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("LSTMGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:cellOutputFwd:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, cellOutputFwd, descriptor, name)
	return rv
}/* debug [instance_methods/method]: LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdDescriptorName */


// Creates an LSTM gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTMGradients(_:recurrentWeight:sourceGradient:zState:cellOutputFwd:inputWeight:bias:initState:initCell:descriptor:name:)
func (g_ Graph) LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("LSTMGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:cellOutputFwd:inputWeight:bias:initState:initCell:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, cellOutputFwd, inputWeight, bias, initState, initCell, descriptor, name)
	return rv
}/* debug [instance_methods/method]: LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellDescriptorName */


// Creates an LSTM gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTMGradients(_:recurrentWeight:sourceGradient:zState:cellOutputFwd:inputWeight:bias:initState:initCell:mask:descriptor:name:)
func (g_ Graph) LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("LSTMGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:cellOutputFwd:inputWeight:bias:initState:initCell:mask:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, cellOutputFwd, inputWeight, bias, initState, initCell, mask, descriptor, name)
	return rv
}/* debug [instance_methods/method]: LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellMaskDescriptorName */


// Creates an LSTM gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTMGradients(_:recurrentWeight:sourceGradient:zState:cellOutputFwd:stateGradient:cellGradient:inputWeight:bias:initState:initCell:mask:peephole:descriptor:name:)
func (g_ Graph) LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdStateGradientCellGradientInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, stateGradient IMPSGraphTensor, cellGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, peephole IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("LSTMGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:cellOutputFwd:stateGradient:cellGradient:inputWeight:bias:initState:initCell:mask:peephole:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, cellOutputFwd, stateGradient, cellGradient, inputWeight, bias, initState, initCell, mask, peephole, descriptor, name)
	return rv
}/* debug [instance_methods/method]: LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdStateGradientCellGradientInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName */


// Computes the matrix multiplication of 2 input tensors with support for broadcasting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/matrixMultiplication(primary:secondary:name:)
func (g_ Graph) MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("matrixMultiplicationWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: MatrixMultiplicationWithPrimaryTensorSecondaryTensorName */


// Returns the elementwise maximum of the input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maximum(_:_:name:)
func (g_ Graph) MaximumWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maximumWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: MaximumWithPrimaryTensorSecondaryTensorName */


// Returns the elementwise maximum of the input tensors, while propagating values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maximumWithNaNPropagation(_:_:name:)
func (g_ Graph) MaximumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maximumWithNaNPropagationWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: MaximumWithNaNPropagationWithPrimaryTensorSecondaryTensorName */


// Creates a 2D max-pooling operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2D(withSourceTensor:descriptor:name:)
func (g_ Graph) MaxPooling2DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling2DWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: MaxPooling2DWithSourceTensorDescriptorName */


// Creates a max-pooling gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DGradient(withGradientTensor:indicesTensor:outputShape:descriptor:name:)
func (g_ Graph) MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling2DGradientWithGradientTensor:indicesTensor:outputShape:descriptor:name:"), gradient, indices, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName */


// Creates a max-pooling gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DGradient(withGradientTensor:indicesTensor:outputShapeTensor:descriptor:name:)
func (g_ Graph) MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling2DGradientWithGradientTensor:indicesTensor:outputShapeTensor:descriptor:name:"), gradient, indices, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName */


// Creates a max-pooling gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DGradient(withGradientTensor:sourceTensor:descriptor:name:)
func (g_ Graph) MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling2DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName */


// Creates a 2D max-pooling operation and returns the result tensor and the corresponding indices tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)
func (g_ Graph) MaxPooling2DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("maxPooling2DReturnIndicesWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: MaxPooling2DReturnIndicesWithSourceTensorDescriptorName */


// Creates a 4D max-pooling operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4D(_:descriptor:name:)
func (g_ Graph) MaxPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling4DWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: MaxPooling4DWithSourceTensorDescriptorName */


// Creates a max-pooling gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4DGradient(_:source:descriptor:name:)
func (g_ Graph) MaxPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling4DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: MaxPooling4DGradientWithGradientTensorSourceTensorDescriptorName */


// Creates a max-pooling gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4DGradient(withGradientTensor:indicesTensor:outputShape:descriptor:name:)
func (g_ Graph) MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape Shape /* not a class type */, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling4DGradientWithGradientTensor:indicesTensor:outputShape:descriptor:name:"), gradient, indices, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName */


// Creates a max-pooling gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4DGradient(withGradientTensor:indicesTensor:outputShapeTensor:descriptor:name:)
func (g_ Graph) MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling4DGradientWithGradientTensor:indicesTensor:outputShapeTensor:descriptor:name:"), gradient, indices, outputShape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName */


// Creates a 4D max-pooling operation and returns the result tensor and the corresponding indices tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4DReturnIndices(_:descriptor:name:)
func (g_ Graph) MaxPooling4DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("maxPooling4DReturnIndicesWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}/* debug [instance_methods/method]: MaxPooling4DReturnIndicesWithSourceTensorDescriptorName */


// Returns the mean of the first input along the specified axes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/mean(of:axes:name:)
func (g_ Graph) MeanOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("meanOfTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: MeanOfTensorAxesName */


// Returns the elementwise minimum of the input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/minimum(_:_:name:)
func (g_ Graph) MinimumWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("minimumWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: MinimumWithPrimaryTensorSecondaryTensorName */


// Returns the elementwise minimum of the input tensors, while propagating values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/minimumWithNaNPropagation(_:_:name:)
func (g_ Graph) MinimumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("minimumWithNaNPropagationWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: MinimumWithNaNPropagationWithPrimaryTensorSecondaryTensorName */


// Returns the remainder obtained by dividing the first input tensor by the second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/modulo(_:_:name:)
func (g_ Graph) ModuloWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("moduloWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: ModuloWithPrimaryTensorSecondaryTensorName */


// Multiplies two input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/multiplication(_:_:name:)
func (g_ Graph) MultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("multiplicationWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: MultiplicationWithPrimaryTensorSecondaryTensorName */


// Applies negative to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/negative(with:name:)
func (g_ Graph) NegativeWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("negativeWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: NegativeWithTensorName */


// Creates a nonMaximumumSuppression operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/nonMaximumSuppression(withBoxesTensor:scoresTensor:classIndicesTensor:iouThreshold:scoreThreshold:perClassSuppression:coordinateMode:name:)
func (g_ Graph) NonMaximumSuppressionWithBoxesTensorScoresTensorClassIndicesTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName(boxesTensor IMPSGraphTensor, scoresTensor IMPSGraphTensor, classIndicesTensor IMPSGraphTensor, IOUThreshold float32, scoreThreshold float32, perClassSuppression bool, coordinateMode GraphNonMaximumSuppressionCoordinateMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("nonMaximumSuppressionWithBoxesTensor:scoresTensor:classIndicesTensor:IOUThreshold:scoreThreshold:perClassSuppression:coordinateMode:name:"), boxesTensor, scoresTensor, classIndicesTensor, IOUThreshold, scoreThreshold, perClassSuppression, coordinateMode, name)
	return rv
}/* debug [instance_methods/method]: NonMaximumSuppressionWithBoxesTensorScoresTensorClassIndicesTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName */


// Creates a nonMaximumumSuppression operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/nonMaximumSuppression(withBoxesTensor:scoresTensor:iouThreshold:scoreThreshold:perClassSuppression:coordinateMode:name:)
func (g_ Graph) NonMaximumSuppressionWithBoxesTensorScoresTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName(boxesTensor IMPSGraphTensor, scoresTensor IMPSGraphTensor, IOUThreshold float32, scoreThreshold float32, perClassSuppression bool, coordinateMode GraphNonMaximumSuppressionCoordinateMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("nonMaximumSuppressionWithBoxesTensor:scoresTensor:IOUThreshold:scoreThreshold:perClassSuppression:coordinateMode:name:"), boxesTensor, scoresTensor, IOUThreshold, scoreThreshold, perClassSuppression, coordinateMode, name)
	return rv
}/* debug [instance_methods/method]: NonMaximumSuppressionWithBoxesTensorScoresTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName */


// Computes the indices of the non-zero elements of the input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/nonZeroIndices(_:name:)
func (g_ Graph) NonZeroIndicesOfTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("nonZeroIndicesOfTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: NonZeroIndicesOfTensorName */


// Creates a normalization beta-gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/normalizationBetaGradient(withIncomingGradientTensor:sourceTensor:reductionAxes:name:)
func (g_ Graph) NormalizationBetaGradientWithIncomingGradientTensorSourceTensorReductionAxesName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("normalizationBetaGradientWithIncomingGradientTensor:sourceTensor:reductionAxes:name:"), incomingGradientTensor, sourceTensor, axes, name)
	return rv
}/* debug [instance_methods/method]: NormalizationBetaGradientWithIncomingGradientTensorSourceTensorReductionAxesName */


// Creates a normalization gamma-gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/normalizationGammaGradient(withIncomingGradientTensor:sourceTensor:mean:varianceTensor:reductionAxes:epsilon:name:)
func (g_ Graph) NormalizationGammaGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorReductionAxesEpsilonName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, meanTensor IMPSGraphTensor, varianceTensor IMPSGraphTensor, axes []foundation.Number, epsilon float32, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("normalizationGammaGradientWithIncomingGradientTensor:sourceTensor:meanTensor:varianceTensor:reductionAxes:epsilon:name:"), incomingGradientTensor, sourceTensor, meanTensor, varianceTensor, axes, epsilon, name)
	return rv
}/* debug [instance_methods/method]: NormalizationGammaGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorReductionAxesEpsilonName */


// Creates a normalization input gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/normalizationGradient(withIncomingGradientTensor:sourceTensor:mean:varianceTensor:gammaTensor:gammaGradientTensor:betaGradientTensor:reductionAxes:epsilon:name:)
func (g_ Graph) NormalizationGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorGammaTensorGammaGradientTensorBetaGradientTensorReductionAxesEpsilonName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, meanTensor IMPSGraphTensor, varianceTensor IMPSGraphTensor, gamma IMPSGraphTensor, gammaGradient IMPSGraphTensor, betaGradient IMPSGraphTensor, axes []foundation.Number, epsilon float32, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("normalizationGradientWithIncomingGradientTensor:sourceTensor:meanTensor:varianceTensor:gammaTensor:gammaGradientTensor:betaGradientTensor:reductionAxes:epsilon:name:"), incomingGradientTensor, sourceTensor, meanTensor, varianceTensor, gamma, gammaGradient, betaGradient, axes, epsilon, name)
	return rv
}/* debug [instance_methods/method]: NormalizationGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorGammaTensorGammaGradientTensorBetaGradientTensorReductionAxesEpsilonName */


// Creates a batch normalization operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/normalize(_:mean:variance:gamma:beta:epsilon:name:)
func (g_ Graph) NormalizationWithTensorMeanTensorVarianceTensorGammaTensorBetaTensorEpsilonName(tensor IMPSGraphTensor, mean IMPSGraphTensor, variance IMPSGraphTensor, gamma IMPSGraphTensor, beta IMPSGraphTensor, epsilon float32, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("normalizationWithTensor:meanTensor:varianceTensor:gammaTensor:betaTensor:epsilon:name:"), tensor, mean, variance, gamma, beta, epsilon, name)
	return rv
}/* debug [instance_methods/method]: NormalizationWithTensorMeanTensorVarianceTensorGammaTensorBetaTensorEpsilonName */


// Applies the logical NOT operation to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/not(with:name:)
func (g_ Graph) NotWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("notWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: NotWithTensorName */


// Returns the elementwise inequality check of the input tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/notEqual(_:_:name:)
func (g_ Graph) NotEqualWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("notEqualWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: NotEqualWithPrimaryTensorSecondaryTensorName */


// Creates a oneHot operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:axis:dataType:name:)
func (g_ Graph) OneHotWithIndicesTensorDepthAxisDataTypeName(indicesTensor IMPSGraphTensor, depth uint, axis uint, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("oneHotWithIndicesTensor:depth:axis:dataType:name:"), indicesTensor, depth, axis, dataType, name)
	return rv
}/* debug [instance_methods/method]: OneHotWithIndicesTensorDepthAxisDataTypeName */


// Creates a oneHot operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:axis:dataType:onValue:offValue:name:)
func (g_ Graph) OneHotWithIndicesTensorDepthAxisDataTypeOnValueOffValueName(indicesTensor IMPSGraphTensor, depth uint, axis uint, dataType DataType /* not a class type */, onValue float64, offValue float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("oneHotWithIndicesTensor:depth:axis:dataType:onValue:offValue:name:"), indicesTensor, depth, axis, dataType, onValue, offValue, name)
	return rv
}/* debug [instance_methods/method]: OneHotWithIndicesTensorDepthAxisDataTypeOnValueOffValueName */


// Creates a oneHot operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:axis:name:)
func (g_ Graph) OneHotWithIndicesTensorDepthAxisName(indicesTensor IMPSGraphTensor, depth uint, axis uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("oneHotWithIndicesTensor:depth:axis:name:"), indicesTensor, depth, axis, name)
	return rv
}/* debug [instance_methods/method]: OneHotWithIndicesTensorDepthAxisName */


// Creates a oneHot operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:dataType:name:)
func (g_ Graph) OneHotWithIndicesTensorDepthDataTypeName(indicesTensor IMPSGraphTensor, depth uint, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("oneHotWithIndicesTensor:depth:dataType:name:"), indicesTensor, depth, dataType, name)
	return rv
}/* debug [instance_methods/method]: OneHotWithIndicesTensorDepthDataTypeName */


// Creates a oneHot operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:dataType:onValue:offValue:name:)
func (g_ Graph) OneHotWithIndicesTensorDepthDataTypeOnValueOffValueName(indicesTensor IMPSGraphTensor, depth uint, dataType DataType /* not a class type */, onValue float64, offValue float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("oneHotWithIndicesTensor:depth:dataType:onValue:offValue:name:"), indicesTensor, depth, dataType, onValue, offValue, name)
	return rv
}/* debug [instance_methods/method]: OneHotWithIndicesTensorDepthDataTypeOnValueOffValueName */


// Creates a oneHot operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:name:)
func (g_ Graph) OneHotWithIndicesTensorDepthName(indicesTensor IMPSGraphTensor, depth uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("oneHotWithIndicesTensor:depth:name:"), indicesTensor, depth, name)
	return rv
}/* debug [instance_methods/method]: OneHotWithIndicesTensorDepthName */


// Creates a padding gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/padGradient(withIncomingGradientTensor:sourceTensor:paddingMode:leftPadding:rightPadding:name:)
func (g_ Graph) PadGradientWithIncomingGradientTensorSourceTensorPaddingModeLeftPaddingRightPaddingName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, paddingMode GraphPaddingMode, leftPadding Shape /* not a class type */, rightPadding Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("padGradientWithIncomingGradientTensor:sourceTensor:paddingMode:leftPadding:rightPadding:name:"), incomingGradientTensor, sourceTensor, paddingMode, leftPadding, rightPadding, name)
	return rv
}/* debug [instance_methods/method]: PadGradientWithIncomingGradientTensorSourceTensorPaddingModeLeftPaddingRightPaddingName */


// Creates a padding operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/padTensor(_:with:leftPadding:rightPadding:constantValue:name:)
func (g_ Graph) PadTensorWithPaddingModeLeftPaddingRightPaddingConstantValueName(tensor IMPSGraphTensor, paddingMode GraphPaddingMode, leftPadding Shape /* not a class type */, rightPadding Shape /* not a class type */, constantValue float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("padTensor:withPaddingMode:leftPadding:rightPadding:constantValue:name:"), tensor, paddingMode, leftPadding, rightPadding, constantValue, name)
	return rv
}/* debug [instance_methods/method]: PadTensorWithPaddingModeLeftPaddingRightPaddingConstantValueName */


// Creates a placeholder operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/placeholder(shape:dataType:name:)
func (g_ Graph) PlaceholderWithShapeDataTypeName(shape Shape /* not a class type */, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("placeholderWithShape:dataType:name:"), shape, dataType, name)
	return rv
}/* debug [instance_methods/method]: PlaceholderWithShapeDataTypeName */


// Creates a placeholder operation and returns the result tensor with the dataType of the placeholder tensor set to 32 bit float.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/placeholder(shape:name:)
func (g_ Graph) PlaceholderWithShapeName(shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("placeholderWithShape:name:"), shape, name)
	return rv
}/* debug [instance_methods/method]: PlaceholderWithShapeName */


// Returns the elementwise result of raising the first tensor to the power of the second tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/power(_:_:name:)
func (g_ Graph) PowerWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("powerWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: PowerWithPrimaryTensorSecondaryTensorName */


// Creates a Quantize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/quantize(_:scale:zeroPoint:dataType:name:)
func (g_ Graph) QuantizeTensorScaleZeroPointDataTypeName(tensor IMPSGraphTensor, scale float64, zeroPoint float64, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("quantizeTensor:scale:zeroPoint:dataType:name:"), tensor, scale, zeroPoint, dataType, name)
	return rv
}/* debug [instance_methods/method]: QuantizeTensorScaleZeroPointDataTypeName */


// Creates a Quantize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/quantize(_:scaleTensor:zeroPoint:dataType:axis:name:)
func (g_ Graph) QuantizeTensorScaleTensorZeroPointDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPoint float64, dataType DataType /* not a class type */, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("quantizeTensor:scaleTensor:zeroPoint:dataType:axis:name:"), tensor, scaleTensor, zeroPoint, dataType, axis, name)
	return rv
}/* debug [instance_methods/method]: QuantizeTensorScaleTensorZeroPointDataTypeAxisName */


// Creates a Quantize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/quantize(_:scaleTensor:zeroPointTensor:dataType:axis:name:)
func (g_ Graph) QuantizeTensorScaleTensorZeroPointTensorDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType DataType /* not a class type */, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("quantizeTensor:scaleTensor:zeroPointTensor:dataType:axis:name:"), tensor, scaleTensor, zeroPointTensor, dataType, axis, name)
	return rv
}/* debug [instance_methods/method]: QuantizeTensorScaleTensorZeroPointTensorDataTypeAxisName */


// Creates a tensor representing state using the Philox algorithm with given counter and key values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomPhiloxStateTensor(withCounterLow:counterHigh:key:name:)
func (g_ Graph) RandomPhiloxStateTensorWithCounterLowCounterHighKeyName(counterLow uint, counterHigh uint, key uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("randomPhiloxStateTensorWithCounterLow:counterHigh:key:name:"), counterLow, counterHigh, key, name)
	return rv
}/* debug [instance_methods/method]: RandomPhiloxStateTensorWithCounterLowCounterHighKeyName */


// Creates a tensor representing state using the Philox algorithm with given counter and key values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomPhiloxStateTensor(withSeed:name:)
func (g_ Graph) RandomPhiloxStateTensorWithSeedName(seed uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("randomPhiloxStateTensorWithSeed:name:"), seed, name)
	return rv
}/* debug [instance_methods/method]: RandomPhiloxStateTensorWithSeedName */


// Creates a Random op of type matching distribution in descriptor and returns random values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShape:descriptor:name:)
func (g_ Graph) RandomTensorWithShapeDescriptorName(shape Shape /* not a class type */, descriptor IMPSGraphRandomOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("randomTensorWithShape:descriptor:name:"), shape, descriptor, name)
	return rv
}/* debug [instance_methods/method]: RandomTensorWithShapeDescriptorName */


// Creates a Random op of type matching distribution in descriptor and returns random values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShape:descriptor:seed:name:)
func (g_ Graph) RandomTensorWithShapeDescriptorSeedName(shape Shape /* not a class type */, descriptor IMPSGraphRandomOpDescriptor, seed uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("randomTensorWithShape:descriptor:seed:name:"), shape, descriptor, seed, name)
	return rv
}/* debug [instance_methods/method]: RandomTensorWithShapeDescriptorSeedName */


// Creates a Random op of type matching distribution in descriptor, and returns random values and updated state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShape:descriptor:stateTensor:name:)
func (g_ Graph) RandomTensorWithShapeDescriptorStateTensorName(shape Shape /* not a class type */, descriptor IMPSGraphRandomOpDescriptor, state IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("randomTensorWithShape:descriptor:stateTensor:name:"), shape, descriptor, state, name)
	return rv
}/* debug [instance_methods/method]: RandomTensorWithShapeDescriptorStateTensorName */


// Creates a Random op of type matching distribution in descriptor and returns random values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShapeTensor:descriptor:name:)
func (g_ Graph) RandomTensorWithShapeTensorDescriptorName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("randomTensorWithShapeTensor:descriptor:name:"), shapeTensor, descriptor, name)
	return rv
}/* debug [instance_methods/method]: RandomTensorWithShapeTensorDescriptorName */


// Creates a Random op of type matching distribution in descriptor and returns random values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShapeTensor:descriptor:seed:name:)
func (g_ Graph) RandomTensorWithShapeTensorDescriptorSeedName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, seed uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("randomTensorWithShapeTensor:descriptor:seed:name:"), shapeTensor, descriptor, seed, name)
	return rv
}/* debug [instance_methods/method]: RandomTensorWithShapeTensorDescriptorSeedName */


// Creates a Random op of type matching distribution in descriptor, and returns random values and updated state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShapeTensor:descriptor:stateTensor:name:)
func (g_ Graph) RandomTensorWithShapeTensorDescriptorStateTensorName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, state IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("randomTensorWithShapeTensor:descriptor:stateTensor:name:"), shapeTensor, descriptor, state, name)
	return rv
}/* debug [instance_methods/method]: RandomTensorWithShapeTensorDescriptorStateTensorName */


// Creates a RandomUniform operation and returns random uniform values
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShape:name:)
func (g_ Graph) RandomUniformTensorWithShapeName(shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("randomUniformTensorWithShape:name:"), shape, name)
	return rv
}/* debug [instance_methods/method]: RandomUniformTensorWithShapeName */


// Creates a RandomUniform operation and returns random uniform values
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShape:seed:name:)
func (g_ Graph) RandomUniformTensorWithShapeSeedName(shape Shape /* not a class type */, seed uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("randomUniformTensorWithShape:seed:name:"), shape, seed, name)
	return rv
}/* debug [instance_methods/method]: RandomUniformTensorWithShapeSeedName */


// Creates a RandomUniform operation and returns random uniform values and updated state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShape:stateTensor:name:)
func (g_ Graph) RandomUniformTensorWithShapeStateTensorName(shape Shape /* not a class type */, state IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("randomUniformTensorWithShape:stateTensor:name:"), shape, state, name)
	return rv
}/* debug [instance_methods/method]: RandomUniformTensorWithShapeStateTensorName */


// Creates a RandomUniform operation and returns random uniform values
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShapeTensor:name:)
func (g_ Graph) RandomUniformTensorWithShapeTensorName(shapeTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("randomUniformTensorWithShapeTensor:name:"), shapeTensor, name)
	return rv
}/* debug [instance_methods/method]: RandomUniformTensorWithShapeTensorName */


// Creates a RandomUniform operation and returns random uniform values
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShapeTensor:seed:name:)
func (g_ Graph) RandomUniformTensorWithShapeTensorSeedName(shapeTensor IMPSGraphTensor, seed uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("randomUniformTensorWithShapeTensor:seed:name:"), shapeTensor, seed, name)
	return rv
}/* debug [instance_methods/method]: RandomUniformTensorWithShapeTensorSeedName */


// Creates a RandomUniform operation and returns random uniform values and updated state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShapeTensor:stateTensor:name:)
func (g_ Graph) RandomUniformTensorWithShapeTensorStateTensorName(shapeTensor IMPSGraphTensor, state IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("randomUniformTensorWithShapeTensor:stateTensor:name:"), shapeTensor, state, name)
	return rv
}/* debug [instance_methods/method]: RandomUniformTensorWithShapeTensorStateTensorName */


// Creates a read op which reads at this point of execution of the graph and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/read(_:name:)
func (g_ Graph) ReadVariableName(variable IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("readVariable:name:"), variable, name)
	return rv
}/* debug [instance_methods/method]: ReadVariableName */


// Returns the real part of a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/realPartOfTensor(tensor:name:)
func (g_ Graph) RealPartOfTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("realPartOfTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: RealPartOfTensorName */


// Creates a Real-to-Hermitean fast Fourier transform operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/realToHermiteanFFT(_:axes:descriptor:name:)
func (g_ Graph) RealToHermiteanFFTWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.Number, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("realToHermiteanFFTWithTensor:axes:descriptor:name:"), tensor, axes, descriptor, name)
	return rv
}/* debug [instance_methods/method]: RealToHermiteanFFTWithTensorAxesDescriptorName */


// Creates a Real-to-Hermitean fast Fourier transform operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/realToHermiteanFFT(_:axesTensor:descriptor:name:)
func (g_ Graph) RealToHermiteanFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("realToHermiteanFFTWithTensor:axesTensor:descriptor:name:"), tensor, axesTensor, descriptor, name)
	return rv
}/* debug [instance_methods/method]: RealToHermiteanFFTWithTensorAxesTensorDescriptorName */


// Applies the reciprocal operation to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reciprocal(with:name:)
func (g_ Graph) ReciprocalWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reciprocalWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: ReciprocalWithTensorName */


// Applies the reciprocal square root operation to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reciprocalSquareRoot(_:name:)
func (g_ Graph) ReciprocalSquareRootWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reciprocalSquareRootWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: ReciprocalSquareRootWithTensorName */


// Creates a reduction and operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionAnd(with:axes:name:)
func (g_ Graph) ReductionAndWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionAndWithTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: ReductionAndWithTensorAxesName */


// Creates a reduction and operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionAnd(with:axis:name:)
func (g_ Graph) ReductionAndWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionAndWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ReductionAndWithTensorAxisName */


// Creates a reduction argMax operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionArgMaximum(with:axis:name:)
func (g_ Graph) ReductionArgMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionArgMaximumWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ReductionArgMaximumWithTensorAxisName */


// Creates a reduction argMin operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionArgMinimum(with:axis:name:)
func (g_ Graph) ReductionArgMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionArgMinimumWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ReductionArgMinimumWithTensorAxisName */


// Creates a reduction max operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMaximum(with:axes:name:)
func (g_ Graph) ReductionMaximumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionMaximumWithTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: ReductionMaximumWithTensorAxesName */


// Creates a reduction max operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMaximum(with:axis:name:)
func (g_ Graph) ReductionMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionMaximumWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ReductionMaximumWithTensorAxisName */


// Creates a reduction max propagate NaN operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMaximumPropagateNaN(with:axes:name:)
func (g_ Graph) ReductionMaximumPropagateNaNWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionMaximumPropagateNaNWithTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: ReductionMaximumPropagateNaNWithTensorAxesName */


// Creates a reduction max propagate NaN operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMaximumPropagateNaN(with:axis:name:)
func (g_ Graph) ReductionMaximumPropagateNaNWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionMaximumPropagateNaNWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ReductionMaximumPropagateNaNWithTensorAxisName */


// Creates a reduction min operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMinimum(with:axes:name:)
func (g_ Graph) ReductionMinimumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionMinimumWithTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: ReductionMinimumWithTensorAxesName */


// Creates a reduction minimum operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMinimum(with:axis:name:)
func (g_ Graph) ReductionMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionMinimumWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ReductionMinimumWithTensorAxisName */


// Creates a reduction min propagate NaN operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMinimumPropagateNaN(with:axes:name:)
func (g_ Graph) ReductionMinimumPropagateNaNWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionMinimumPropagateNaNWithTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: ReductionMinimumPropagateNaNWithTensorAxesName */


// Creates a reduction min propagate NaN operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMinimumPropagateNaN(with:axis:name:)
func (g_ Graph) ReductionMinimumPropagateNaNWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionMinimumPropagateNaNWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ReductionMinimumPropagateNaNWithTensorAxisName */


// Creates a reduction or operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionOr(with:axes:name:)
func (g_ Graph) ReductionOrWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionOrWithTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: ReductionOrWithTensorAxesName */


// Creates a reduction or operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionOr(with:axis:name:)
func (g_ Graph) ReductionOrWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionOrWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ReductionOrWithTensorAxisName */


// Creates a reduction product operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionProduct(with:axes:name:)
func (g_ Graph) ReductionProductWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionProductWithTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: ReductionProductWithTensorAxesName */


// Creates a reduction product operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionProduct(with:axis:name:)
func (g_ Graph) ReductionProductWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionProductWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ReductionProductWithTensorAxisName */


// Creates a reduction sum operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionSum(with:axes:name:)
func (g_ Graph) ReductionSumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionSumWithTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: ReductionSumWithTensorAxesName */


// Creates a reduction sum operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionSum(with:axis:name:)
func (g_ Graph) ReductionSumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reductionSumWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: ReductionSumWithTensorAxisName */


// Creates a reinterpret cast operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reinterpretCast(_:to:name:)
func (g_ Graph) ReinterpretCastTensorToTypeName(tensor IMPSGraphTensor, type_ DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reinterpretCastTensor:toType:name:"), tensor, type_, name)
	return rv
}/* debug [instance_methods/method]: ReinterpretCastTensorToTypeName */


// Computes the ReLU (rectified linear activation unit) function with the input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reLU(with:name:)
func (g_ Graph) ReLUWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reLUWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: ReLUWithTensorName */


// Computes the gradient of the ReLU (rectified linear activation unit) function using the incoming gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reLUGradient(withIncomingGradient:sourceTensor:name:)
func (g_ Graph) ReLUGradientWithIncomingGradientSourceTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reLUGradientWithIncomingGradient:sourceTensor:name:"), gradient, source, name)
	return rv
}/* debug [instance_methods/method]: ReLUGradientWithIncomingGradientSourceTensorName */


// Creates a reshape operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reshape(_:shape:name:)
func (g_ Graph) ReshapeTensorWithShapeName(tensor IMPSGraphTensor, shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reshapeTensor:withShape:name:"), tensor, shape, name)
	return rv
}/* debug [instance_methods/method]: ReshapeTensorWithShapeName */


// Creates a reshape operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reshape(_:shapeTensor:name:)
func (g_ Graph) ReshapeTensorWithShapeTensorName(tensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reshapeTensor:withShapeTensor:name:"), tensor, shapeTensor, name)
	return rv
}/* debug [instance_methods/method]: ReshapeTensorWithShapeTensorName */


// Creates a Resize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(_:size:mode:centerResult:alignCorners:layout:name:)
func (g_ Graph) ResizeTensorSizeModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size Shape /* not a class type */, mode GraphResizeMode, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeTensor:size:mode:centerResult:alignCorners:layout:name:"), imagesTensor, size, mode, centerResult, alignCorners, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeTensorSizeModeCenterResultAlignCornersLayoutName */


// Creates a Resize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(_:sizeTensor:mode:centerResult:alignCorners:layout:name:)
func (g_ Graph) ResizeTensorSizeTensorModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, mode GraphResizeMode, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeTensor:sizeTensor:mode:centerResult:alignCorners:layout:name:"), imagesTensor, size, mode, centerResult, alignCorners, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeTensorSizeTensorModeCenterResultAlignCornersLayoutName */


// Creates a Resize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(_:sizeTensor:mode:centerResult:alignCorners:name:)
func (g_ Graph) ResizeTensorSizeTensorModeCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, mode GraphResizeMode, centerResult bool, alignCorners bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeTensor:sizeTensor:mode:centerResult:alignCorners:name:"), imagesTensor, size, mode, centerResult, alignCorners, name)
	return rv
}/* debug [instance_methods/method]: ResizeTensorSizeTensorModeCenterResultAlignCornersName */


// Resamples input images to given size using the provided scale and offset. Destination indices are computed using
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(_:sizeTensor:scaleOffsetTensor:mode:layout:name:)
func (g_ Graph) ResizeTensorSizeTensorScaleOffsetTensorModeLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, mode GraphResizeMode, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeTensor:sizeTensor:scaleOffsetTensor:mode:layout:name:"), imagesTensor, size, scaleOffset, mode, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeTensorSizeTensorScaleOffsetTensorModeLayoutName */


// Creates a Resize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(_:sizeTensor:scaleTensor:offsetTenor:mode:name:)
func (g_ Graph) ResizeTensorSizeTensorScaleTensorOffsetTensorModeName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, mode GraphResizeMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeTensor:sizeTensor:scaleTensor:offsetTensor:mode:name:"), imagesTensor, size, scale, offset, mode, name)
	return rv
}/* debug [instance_methods/method]: ResizeTensorSizeTensorScaleTensorOffsetTensorModeName */


// Creates a Resize gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(withGradientTensor:input:mode:centerResult:alignCorners:layout:name:)
func (g_ Graph) ResizeWithGradientTensorInputModeCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, mode GraphResizeMode, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeWithGradientTensor:input:mode:centerResult:alignCorners:layout:name:"), gradient, input, mode, centerResult, alignCorners, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeWithGradientTensorInputModeCenterResultAlignCornersLayoutName */


// Creates a Resize gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(withGradientTensor:input:scale:offsetTensor:mode:name:)
func (g_ Graph) ResizeWithGradientTensorInputScaleTensorOffsetTensorModeName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, mode GraphResizeMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeWithGradientTensor:input:scaleTensor:offsetTensor:mode:name:"), gradient, input, scale, offset, mode, name)
	return rv
}/* debug [instance_methods/method]: ResizeWithGradientTensorInputScaleTensorOffsetTensorModeName */


// Creates a Resize gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(withGradientTensor:input:scaleOffsetTensor:mode:layout:name:)
func (g_ Graph) ResizeWithGradientTensorInputScaleOffsetTensorModeLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, mode GraphResizeMode, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeWithGradientTensor:input:scaleOffsetTensor:mode:layout:name:"), gradient, input, scaleOffset, mode, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeWithGradientTensorInputScaleOffsetTensorModeLayoutName */


// Resamples input images to given size using bilinear sampling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(_:sizeTensor:centerResult:alignCorners:layout:name:)
func (g_ Graph) ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeBilinearWithTensor:sizeTensor:centerResult:alignCorners:layout:name:"), imagesTensor, size, centerResult, alignCorners, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersLayoutName */


// Creates a Resize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(_:sizeTensor:centerResult:alignCorners:name:)
func (g_ Graph) ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, centerResult bool, alignCorners bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeBilinearWithTensor:sizeTensor:centerResult:alignCorners:name:"), imagesTensor, size, centerResult, alignCorners, name)
	return rv
}/* debug [instance_methods/method]: ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersName */


// Resamples input images to given size using the provided scale and offset and bilinear sampling See above discussion for more details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(_:sizeTensor:scaleOffsetTensor:layout:name:)
func (g_ Graph) ResizeBilinearWithTensorSizeTensorScaleOffsetTensorLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeBilinearWithTensor:sizeTensor:scaleOffsetTensor:layout:name:"), imagesTensor, size, scaleOffset, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeBilinearWithTensorSizeTensorScaleOffsetTensorLayoutName */


// Creates a Resize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(_:sizeTensor:scaleTensor:offsetTensor:name:)
func (g_ Graph) ResizeBilinearWithTensorSizeTensorScaleTensorOffsetTensorName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeBilinearWithTensor:sizeTensor:scaleTensor:offsetTensor:name:"), imagesTensor, size, scale, offset, name)
	return rv
}/* debug [instance_methods/method]: ResizeBilinearWithTensorSizeTensorScaleTensorOffsetTensorName */


// Creates a Resize gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(withGradientTensor:input:centerResult:alignCorners:layout:name:)
func (g_ Graph) ResizeBilinearWithGradientTensorInputCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeBilinearWithGradientTensor:input:centerResult:alignCorners:layout:name:"), gradient, input, centerResult, alignCorners, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeBilinearWithGradientTensorInputCenterResultAlignCornersLayoutName */


// Creates a Resize gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(withGradientTensor:input:scale:offsetTensor:name:)
func (g_ Graph) ResizeBilinearWithGradientTensorInputScaleTensorOffsetTensorName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeBilinearWithGradientTensor:input:scaleTensor:offsetTensor:name:"), gradient, input, scale, offset, name)
	return rv
}/* debug [instance_methods/method]: ResizeBilinearWithGradientTensorInputScaleTensorOffsetTensorName */


// Creates a Resize gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(withGradientTensor:input:scaleOffsetTensor:layout:name:)
func (g_ Graph) ResizeBilinearWithGradientTensorInputScaleOffsetTensorLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeBilinearWithGradientTensor:input:scaleOffsetTensor:layout:name:"), gradient, input, scaleOffset, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeBilinearWithGradientTensorInputScaleOffsetTensorLayoutName */


// Resamples input images to given size using nearest neighbor sampling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(_:sizeTensor:nearestRoundingMode:centerResult:alignCorners:layout:name:)
func (g_ Graph) ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeNearestWithTensor:sizeTensor:nearestRoundingMode:centerResult:alignCorners:layout:name:"), imagesTensor, size, nearestRoundingMode, centerResult, alignCorners, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersLayoutName */


// Creates a Resize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(_:sizeTensor:nearestRoundingMode:centerResult:alignCorners:name:)
func (g_ Graph) ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeNearestWithTensor:sizeTensor:nearestRoundingMode:centerResult:alignCorners:name:"), imagesTensor, size, nearestRoundingMode, centerResult, alignCorners, name)
	return rv
}/* debug [instance_methods/method]: ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersName */


// Resamples input images to given size using the provided scale and offset and nearest neighbor sampling See above discussion for more details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(_:sizeTensor:scaleOffsetTensor:nearestRoundingMode:layout:name:)
func (g_ Graph) ResizeNearestWithTensorSizeTensorScaleOffsetTensorNearestRoundingModeLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeNearestWithTensor:sizeTensor:scaleOffsetTensor:nearestRoundingMode:layout:name:"), imagesTensor, size, scaleOffset, nearestRoundingMode, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeNearestWithTensorSizeTensorScaleOffsetTensorNearestRoundingModeLayoutName */


// Creates a Resize operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(_:sizeTensor:scaleTensor:offsetTensor:nearestRoundingMode:name:)
func (g_ Graph) ResizeNearestWithTensorSizeTensorScaleTensorOffsetTensorNearestRoundingModeName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeNearestWithTensor:sizeTensor:scaleTensor:offsetTensor:nearestRoundingMode:name:"), imagesTensor, size, scale, offset, nearestRoundingMode, name)
	return rv
}/* debug [instance_methods/method]: ResizeNearestWithTensorSizeTensorScaleTensorOffsetTensorNearestRoundingModeName */


// Creates a Resize gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(withGradientTensor:input:nearestRoundingMode:centerResult:alignCorners:layout:name:)
func (g_ Graph) ResizeNearestWithGradientTensorInputNearestRoundingModeCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeNearestWithGradientTensor:input:nearestRoundingMode:centerResult:alignCorners:layout:name:"), gradient, input, nearestRoundingMode, centerResult, alignCorners, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeNearestWithGradientTensorInputNearestRoundingModeCenterResultAlignCornersLayoutName */


// Creates a Resize gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(withGradientTensor:input:scale:offsetTensor:nearestRoundingMode:name:)
func (g_ Graph) ResizeNearestWithGradientTensorInputScaleTensorOffsetTensorNearestRoundingModeName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeNearestWithGradientTensor:input:scaleTensor:offsetTensor:nearestRoundingMode:name:"), gradient, input, scale, offset, nearestRoundingMode, name)
	return rv
}/* debug [instance_methods/method]: ResizeNearestWithGradientTensorInputScaleTensorOffsetTensorNearestRoundingModeName */


// Creates a Resize gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(withGradientTensor:input:scaleOffsetTensor:nearestRoundingMode:layout:name:)
func (g_ Graph) ResizeNearestWithGradientTensorInputScaleOffsetTensorNearestRoundingModeLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, nearestRoundingMode GraphResizeNearestRoundingMode, layout GraphTensorNamedDataLayout, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("resizeNearestWithGradientTensor:input:scaleOffsetTensor:nearestRoundingMode:layout:name:"), gradient, input, scaleOffset, nearestRoundingMode, layout, name)
	return rv
}/* debug [instance_methods/method]: ResizeNearestWithGradientTensorInputScaleOffsetTensorNearestRoundingModeLayoutName */


// Creates a reverse operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reverse(_:axes:name:)
func (g_ Graph) ReverseTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reverseTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: ReverseTensorAxesName */


// Creates a reverse operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reverse(_:axesTensor:name:)
func (g_ Graph) ReverseTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reverseTensor:axesTensor:name:"), tensor, axesTensor, name)
	return rv
}/* debug [instance_methods/method]: ReverseTensorAxesTensorName */


// Creates a reverse operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reverse(_:name:)
func (g_ Graph) ReverseTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("reverseTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: ReverseTensorName */


// Rounds the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/round(with:name:)
func (g_ Graph) RoundWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("roundWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: RoundWithTensorName */


// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/run(feeds:targetTensors:targetOperations:)
func (g_ Graph) RunWithFeedsTargetTensorsTargetOperations(feeds GraphTensorDataDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation) GraphTensorDataDictionary /* not a class type */ {
	rv := objc.Send[GraphTensorDataDictionary](g_.ID, objc.Sel("runWithFeeds:targetTensors:targetOperations:"), feeds, targetTensors, targetOperations)
	return rv
}/* debug [instance_methods/method]: RunWithFeedsTargetTensorsTargetOperations */


// Runs the graph for the given feeds and returns the target tensor values in the results dictionary provided by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/run(with:feeds:targetOperations:resultsDictionary:)
func (g_ Graph) RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(commandQueue unsafe.Pointer, feeds GraphTensorDataDictionary /* not a class type */, targetOperations []GraphOperation, resultsDictionary GraphTensorDataDictionary /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("runWithMTLCommandQueue:feeds:targetOperations:resultsDictionary:"), commandQueue, feeds, targetOperations, resultsDictionary)
}/* debug [instance_methods/method]: RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary */


// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/run(with:feeds:targetTensors:targetOperations:)
func (g_ Graph) RunWithMTLCommandQueueFeedsTargetTensorsTargetOperations(commandQueue unsafe.Pointer, feeds GraphTensorDataDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation) GraphTensorDataDictionary /* not a class type */ {
	rv := objc.Send[GraphTensorDataDictionary](g_.ID, objc.Sel("runWithMTLCommandQueue:feeds:targetTensors:targetOperations:"), commandQueue, feeds, targetTensors, targetOperations)
	return rv
}/* debug [instance_methods/method]: RunWithMTLCommandQueueFeedsTargetTensorsTargetOperations */


// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/runAsync(feeds:targetTensors:targetOperations:executionDescriptor:)
func (g_ Graph) RunAsyncWithFeedsTargetTensorsTargetOperationsExecutionDescriptor(feeds GraphTensorDataDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) GraphTensorDataDictionary /* not a class type */ {
	rv := objc.Send[GraphTensorDataDictionary](g_.ID, objc.Sel("runAsyncWithFeeds:targetTensors:targetOperations:executionDescriptor:"), feeds, targetTensors, targetOperations, executionDescriptor)
	return rv
}/* debug [instance_methods/method]: RunAsyncWithFeedsTargetTensorsTargetOperationsExecutionDescriptor */


// Encodes the graph for the given feeds to returns the target tensor values in the results dictionary provided by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/runAsync(with:feeds:targetOperations:resultsDictionary:executionDescriptor:)
func (g_ Graph) RunAsyncWithMTLCommandQueueFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandQueue unsafe.Pointer, feeds GraphTensorDataDictionary /* not a class type */, targetOperations []GraphOperation, resultsDictionary GraphTensorDataDictionary /* not a class type */, executionDescriptor IMPSGraphExecutionDescriptor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("runAsyncWithMTLCommandQueue:feeds:targetOperations:resultsDictionary:executionDescriptor:"), commandQueue, feeds, targetOperations, resultsDictionary, executionDescriptor)
}/* debug [instance_methods/method]: RunAsyncWithMTLCommandQueueFeedsTargetOperationsResultsDictionaryExecutionDescriptor */


// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/runAsync(with:feeds:targetTensors:targetOperations:executionDescriptor:)
func (g_ Graph) RunAsyncWithMTLCommandQueueFeedsTargetTensorsTargetOperationsExecutionDescriptor(commandQueue unsafe.Pointer, feeds GraphTensorDataDictionary /* not a class type */, targetTensors []GraphTensor, targetOperations []GraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) GraphTensorDataDictionary /* not a class type */ {
	rv := objc.Send[GraphTensorDataDictionary](g_.ID, objc.Sel("runAsyncWithMTLCommandQueue:feeds:targetTensors:targetOperations:executionDescriptor:"), commandQueue, feeds, targetTensors, targetOperations, executionDescriptor)
	return rv
}/* debug [instance_methods/method]: RunAsyncWithMTLCommandQueueFeedsTargetTensorsTargetOperationsExecutionDescriptor */


// Samples a tensor using the coordinates provided, using nearest neighbor sampling with specified rounding mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sampleGrid(withSourceTensor:coordinateTensor:layout:normalizeCoordinates:relativeCoordinates:alignCorners:paddingMode:nearestRoundingMode:constantValue:name:)
func (g_ Graph) SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeNearestRoundingModeConstantValueName(source IMPSGraphTensor, coordinates IMPSGraphTensor, layout GraphTensorNamedDataLayout, normalizeCoordinates bool, relativeCoordinates bool, alignCorners bool, paddingMode GraphPaddingMode, nearestRoundingMode GraphResizeNearestRoundingMode, constantValue float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sampleGridWithSourceTensor:coordinateTensor:layout:normalizeCoordinates:relativeCoordinates:alignCorners:paddingMode:nearestRoundingMode:constantValue:name:"), source, coordinates, layout, normalizeCoordinates, relativeCoordinates, alignCorners, paddingMode, nearestRoundingMode, constantValue, name)
	return rv
}/* debug [instance_methods/method]: SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeNearestRoundingModeConstantValueName */


// Samples a tensor using the coordinates provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sampleGrid(withSourceTensor:coordinateTensor:layout:normalizeCoordinates:relativeCoordinates:alignCorners:paddingMode:samplingMode:constantValue:name:)
func (g_ Graph) SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeSamplingModeConstantValueName(source IMPSGraphTensor, coordinates IMPSGraphTensor, layout GraphTensorNamedDataLayout, normalizeCoordinates bool, relativeCoordinates bool, alignCorners bool, paddingMode GraphPaddingMode, samplingMode GraphResizeMode, constantValue float64, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sampleGridWithSourceTensor:coordinateTensor:layout:normalizeCoordinates:relativeCoordinates:alignCorners:paddingMode:samplingMode:constantValue:name:"), source, coordinates, layout, normalizeCoordinates, relativeCoordinates, alignCorners, paddingMode, samplingMode, constantValue, name)
	return rv
}/* debug [instance_methods/method]: SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeSamplingModeConstantValueName */


// Creates a scaled dot product attention (SDPA) operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scaledDotProductAttention(query:key:value:mask:scale:name:)
func (g_ Graph) ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorMaskTensorScaleName(queryTensor IMPSGraphTensor, keyTensor IMPSGraphTensor, valueTensor IMPSGraphTensor, maskTensor IMPSGraphTensor, scale float32, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scaledDotProductAttentionWithQueryTensor:keyTensor:valueTensor:maskTensor:scale:name:"), queryTensor, keyTensor, valueTensor, maskTensor, scale, name)
	return rv
}/* debug [instance_methods/method]: ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorMaskTensorScaleName */


// Creates a scaled dot product attention (SDPA) operation (without a mask) and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scaledDotProductAttention(query:key:value:scale:name:)
func (g_ Graph) ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorScaleName(queryTensor IMPSGraphTensor, keyTensor IMPSGraphTensor, valueTensor IMPSGraphTensor, scale float32, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scaledDotProductAttentionWithQueryTensor:keyTensor:valueTensor:scale:name:"), queryTensor, keyTensor, valueTensor, scale, name)
	return rv
}/* debug [instance_methods/method]: ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorScaleName */


// Creates a Scatter operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatter(_:indices:shape:axis:mode:name:)
func (g_ Graph) ScatterWithUpdatesTensorIndicesTensorShapeAxisModeName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape Shape /* not a class type */, axis int, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scatterWithUpdatesTensor:indicesTensor:shape:axis:mode:name:"), updatesTensor, indicesTensor, shape, axis, mode, name)
	return rv
}/* debug [instance_methods/method]: ScatterWithUpdatesTensorIndicesTensorShapeAxisModeName */


// Creates a ScatterAlongAxis operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterAlongAxis(_:data:updates:indices:mode:name:)
func (g_ Graph) ScatterAlongAxisWithDataTensorUpdatesTensorIndicesTensorModeName(axis int, dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scatterAlongAxis:withDataTensor:updatesTensor:indicesTensor:mode:name:"), axis, dataTensor, updatesTensor, indicesTensor, mode, name)
	return rv
}/* debug [instance_methods/method]: ScatterAlongAxisWithDataTensorUpdatesTensorIndicesTensorModeName */


// Creates a ScatterAlongAxis operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterAlongAxis(_:updates:indices:shape:mode:name:)
func (g_ Graph) ScatterAlongAxisWithUpdatesTensorIndicesTensorShapeModeName(axis int, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape Shape /* not a class type */, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scatterAlongAxis:withUpdatesTensor:indicesTensor:shape:mode:name:"), axis, updatesTensor, indicesTensor, shape, mode, name)
	return rv
}/* debug [instance_methods/method]: ScatterAlongAxisWithUpdatesTensorIndicesTensorShapeModeName */


// Creates a ScatterAlongAxis operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterAlongAxisTensor(_:data:updates:indices:mode:name:)
func (g_ Graph) ScatterAlongAxisTensorWithDataTensorUpdatesTensorIndicesTensorModeName(axisTensor IMPSGraphTensor, dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scatterAlongAxisTensor:withDataTensor:updatesTensor:indicesTensor:mode:name:"), axisTensor, dataTensor, updatesTensor, indicesTensor, mode, name)
	return rv
}/* debug [instance_methods/method]: ScatterAlongAxisTensorWithDataTensorUpdatesTensorIndicesTensorModeName */


// Creates a ScatterAlongAxis operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterAlongAxisTensor(_:updates:indices:shape:mode:name:)
func (g_ Graph) ScatterAlongAxisTensorWithUpdatesTensorIndicesTensorShapeModeName(axisTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape Shape /* not a class type */, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scatterAlongAxisTensor:withUpdatesTensor:indicesTensor:shape:mode:name:"), axisTensor, updatesTensor, indicesTensor, shape, mode, name)
	return rv
}/* debug [instance_methods/method]: ScatterAlongAxisTensorWithUpdatesTensorIndicesTensorShapeModeName */


// Creates a ScatterND operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterND(withUpdatesTensor:indicesTensor:shape:batchDimensions:mode:name:)
func (g_ Graph) ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsModeName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape Shape /* not a class type */, batchDimensions uint, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scatterNDWithUpdatesTensor:indicesTensor:shape:batchDimensions:mode:name:"), updatesTensor, indicesTensor, shape, batchDimensions, mode, name)
	return rv
}/* debug [instance_methods/method]: ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsModeName */


// Creates a ScatterND operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterND(withUpdatesTensor:indicesTensor:shape:batchDimensions:name:)
func (g_ Graph) ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape Shape /* not a class type */, batchDimensions uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scatterNDWithUpdatesTensor:indicesTensor:shape:batchDimensions:name:"), updatesTensor, indicesTensor, shape, batchDimensions, name)
	return rv
}/* debug [instance_methods/method]: ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsName */


// Creates a ScatterND operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterNDWithData(_:updates:indices:batchDimensions:mode:name:)
func (g_ Graph) ScatterNDWithDataTensorUpdatesTensorIndicesTensorBatchDimensionsModeName(dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, batchDimensions uint, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scatterNDWithDataTensor:updatesTensor:indicesTensor:batchDimensions:mode:name:"), dataTensor, updatesTensor, indicesTensor, batchDimensions, mode, name)
	return rv
}/* debug [instance_methods/method]: ScatterNDWithDataTensorUpdatesTensorIndicesTensorBatchDimensionsModeName */


// Creates a Scatter operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterWithData(_:updates:indices:axis:mode:name:)
func (g_ Graph) ScatterWithDataTensorUpdatesTensorIndicesTensorAxisModeName(dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, axis int, mode GraphScatterMode, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("scatterWithDataTensor:updatesTensor:indicesTensor:axis:mode:name:"), dataTensor, updatesTensor, indicesTensor, axis, mode, name)
	return rv
}/* debug [instance_methods/method]: ScatterWithDataTensorUpdatesTensorIndicesTensorAxisModeName */


// Selects values from either the true or false predicate tensor, depending on the values in the first input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/select(predicate:trueTensor:falseTensor:name:)
func (g_ Graph) SelectWithPredicateTensorTruePredicateTensorFalsePredicateTensorName(predicateTensor IMPSGraphTensor, truePredicateTensor IMPSGraphTensor, falseSelectTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("selectWithPredicateTensor:truePredicateTensor:falsePredicateTensor:name:"), predicateTensor, truePredicateTensor, falseSelectTensor, name)
	return rv
}/* debug [instance_methods/method]: SelectWithPredicateTensorTruePredicateTensorFalsePredicateTensorName */


// Creates a shape-of operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/shapeOf(_:name:)
func (g_ Graph) ShapeOfTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("shapeOfTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: ShapeOfTensorName */


// Computes the sigmoid operation on an input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sigmoid(with:name:)
func (g_ Graph) SigmoidWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sigmoidWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: SigmoidWithTensorName */


// Computes the gradient of the sigmoid function using the incoming gradient tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sigmoidGradient(withIncomingGradient:sourceTensor:name:)
func (g_ Graph) SigmoidGradientWithIncomingGradientSourceTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sigmoidGradientWithIncomingGradient:sourceTensor:name:"), gradient, source, name)
	return rv
}/* debug [instance_methods/method]: SigmoidGradientWithIncomingGradientSourceTensorName */


// Returns the sign of the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sign(with:name:)
func (g_ Graph) SignWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("signWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: SignWithTensorName */


// Returns the sign bit of the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/signbit(with:name:)
func (g_ Graph) SignbitWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("signbitWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: SignbitWithTensorName */


// Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNN(_:recurrentWeight:initState:descriptor:name:)
func (g_ Graph) SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("singleGateRNNWithSourceTensor:recurrentWeight:initState:descriptor:name:"), source, recurrentWeight, initState, descriptor, name)
	return rv
}/* debug [instance_methods/method]: SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName */


// Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNN(_:recurrentWeight:inputWeight:bias:initState:descriptor:name:)
func (g_ Graph) SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("singleGateRNNWithSourceTensor:recurrentWeight:inputWeight:bias:initState:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, descriptor, name)
	return rv
}/* debug [instance_methods/method]: SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName */


// Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNN(_:recurrentWeight:inputWeight:bias:initState:mask:descriptor:name:)
func (g_ Graph) SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("singleGateRNNWithSourceTensor:recurrentWeight:inputWeight:bias:initState:mask:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, mask, descriptor, name)
	return rv
}/* debug [instance_methods/method]: SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName */


// Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNNGradients(_:recurrentWeight:sourceGradient:zState:initState:descriptor:name:)
func (g_ Graph) SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("singleGateRNNGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:initState:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, initState, descriptor, name)
	return rv
}/* debug [instance_methods/method]: SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInitStateDescriptorName */


// Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNNGradients(_:recurrentWeight:sourceGradient:zState:inputWeight:bias:initState:descriptor:name:)
func (g_ Graph) SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("singleGateRNNGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:inputWeight:bias:initState:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, inputWeight, bias, initState, descriptor, name)
	return rv
}/* debug [instance_methods/method]: SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateDescriptorName */


// Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNNGradients(_:recurrentWeight:sourceGradient:zState:inputWeight:bias:initState:mask:descriptor:name:)
func (g_ Graph) SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("singleGateRNNGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:inputWeight:bias:initState:mask:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, inputWeight, bias, initState, mask, descriptor, name)
	return rv
}/* debug [instance_methods/method]: SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateMaskDescriptorName */


// Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNNGradients(_:recurrentWeight:sourceGradient:zState:stateGradient:inputWeight:bias:initState:mask:descriptor:name:)
func (g_ Graph) SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateStateGradientInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, stateGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("singleGateRNNGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:stateGradient:inputWeight:bias:initState:mask:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, stateGradient, inputWeight, bias, initState, mask, descriptor, name)
	return rv
}/* debug [instance_methods/method]: SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateStateGradientInputWeightBiasInitStateMaskDescriptorName */


// Creates a strided-slice gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceGradientTensor(_:fwdInShapeTensor:start:end:strideTensor:startMask:endMask:squeezeMask:name:)
func (g_ Graph) SliceGradientTensorFwdInShapeTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, startTensor IMPSGraphTensor, endTensor IMPSGraphTensor, strideTensor IMPSGraphTensor, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceGradientTensor:fwdInShapeTensor:startTensor:endTensor:strideTensor:startMask:endMask:squeezeMask:name:"), inputGradientTensor, fwdInShapeTensor, startTensor, endTensor, strideTensor, startMask, endMask, squeezeMask, name)
	return rv
}/* debug [instance_methods/method]: SliceGradientTensorFwdInShapeTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName */


// Creates a slice gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceGradientTensor(_:fwdInShapeTensor:start:sizeTensor:squeezeMask:name:)
func (g_ Graph) SliceGradientTensorFwdInShapeTensorStartTensorSizeTensorSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, startTensor IMPSGraphTensor, sizeTensor IMPSGraphTensor, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceGradientTensor:fwdInShapeTensor:startTensor:sizeTensor:squeezeMask:name:"), inputGradientTensor, fwdInShapeTensor, startTensor, sizeTensor, squeezeMask, name)
	return rv
}/* debug [instance_methods/method]: SliceGradientTensorFwdInShapeTensorStartTensorSizeTensorSqueezeMaskName */


// Creates a strided-slice gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceGradientTensor(_:fwdInShapeTensor:starts:ends:strides:name:)
func (g_ Graph) SliceGradientTensorFwdInShapeTensorStartsEndsStridesName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceGradientTensor:fwdInShapeTensor:starts:ends:strides:name:"), inputGradientTensor, fwdInShapeTensor, starts, ends, strides, name)
	return rv
}/* debug [instance_methods/method]: SliceGradientTensorFwdInShapeTensorStartsEndsStridesName */


// Creates a strided-slice gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceGradientTensor(_:fwdInShapeTensor:starts:ends:strides:startMask:endMask:squeezeMask:name:)
func (g_ Graph) SliceGradientTensorFwdInShapeTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceGradientTensor:fwdInShapeTensor:starts:ends:strides:startMask:endMask:squeezeMask:name:"), inputGradientTensor, fwdInShapeTensor, starts, ends, strides, startMask, endMask, squeezeMask, name)
	return rv
}/* debug [instance_methods/method]: SliceGradientTensorFwdInShapeTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName */


// Creates a slice operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceTensor(_:dimension:start:length:name:)
func (g_ Graph) SliceTensorDimensionStartLengthName(tensor IMPSGraphTensor, dimensionIndex uint, start int, length int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceTensor:dimension:start:length:name:"), tensor, dimensionIndex, start, length, name)
	return rv
}/* debug [instance_methods/method]: SliceTensorDimensionStartLengthName */


// Creates a strided-slice operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceTensor(_:start:end:strideTensor:startMask:endMask:squeezeMask:name:)
func (g_ Graph) SliceTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName(tensor IMPSGraphTensor, startTensor IMPSGraphTensor, endTensor IMPSGraphTensor, strideTensor IMPSGraphTensor, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceTensor:startTensor:endTensor:strideTensor:startMask:endMask:squeezeMask:name:"), tensor, startTensor, endTensor, strideTensor, startMask, endMask, squeezeMask, name)
	return rv
}/* debug [instance_methods/method]: SliceTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName */


// Creates a slice operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceTensor(_:start:sizeTensor:squeezeMask:name:)
func (g_ Graph) SliceTensorStartTensorSizeTensorSqueezeMaskName(tensor IMPSGraphTensor, startTensor IMPSGraphTensor, sizeTensor IMPSGraphTensor, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceTensor:startTensor:sizeTensor:squeezeMask:name:"), tensor, startTensor, sizeTensor, squeezeMask, name)
	return rv
}/* debug [instance_methods/method]: SliceTensorStartTensorSizeTensorSqueezeMaskName */


// Creates a strided-slice operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceTensor(_:starts:ends:strides:name:)
func (g_ Graph) SliceTensorStartsEndsStridesName(tensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceTensor:starts:ends:strides:name:"), tensor, starts, ends, strides, name)
	return rv
}/* debug [instance_methods/method]: SliceTensorStartsEndsStridesName */


// Creates a strided-slice operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceTensor(_:starts:ends:strides:startMask:endMask:squeezeMask:name:)
func (g_ Graph) SliceTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(tensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceTensor:starts:ends:strides:startMask:endMask:squeezeMask:name:"), tensor, starts, ends, strides, startMask, endMask, squeezeMask, name)
	return rv
}/* debug [instance_methods/method]: SliceTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName */


// Creates a strided-slice update operation with zero masks and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceUpdateDataTensor(_:update:starts:ends:strides:name:)
func (g_ Graph) SliceUpdateDataTensorUpdateTensorStartsEndsStridesName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceUpdateDataTensor:updateTensor:starts:ends:strides:name:"), dataTensor, updateTensor, starts, ends, strides, name)
	return rv
}/* debug [instance_methods/method]: SliceUpdateDataTensorUpdateTensorStartsEndsStridesName */


// Creates a strided-slice update operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceUpdateDataTensor(_:update:starts:ends:strides:startMask:endMask:squeezeMask:name:)
func (g_ Graph) SliceUpdateDataTensorUpdateTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, starts []foundation.Number, ends []foundation.Number, strides []foundation.Number, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceUpdateDataTensor:updateTensor:starts:ends:strides:startMask:endMask:squeezeMask:name:"), dataTensor, updateTensor, starts, ends, strides, startMask, endMask, squeezeMask, name)
	return rv
}/* debug [instance_methods/method]: SliceUpdateDataTensorUpdateTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName */


// Creates a strided-slice update operation with zero masks and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceUpdateDataTensor(_:update:startsTensor:endsTensor:stridesTensor:name:)
func (g_ Graph) SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, startsTensor IMPSGraphTensor, endsTensor IMPSGraphTensor, stridesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceUpdateDataTensor:updateTensor:startsTensor:endsTensor:stridesTensor:name:"), dataTensor, updateTensor, startsTensor, endsTensor, stridesTensor, name)
	return rv
}/* debug [instance_methods/method]: SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorName */


// Creates a strided-slice update operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceUpdateDataTensor(_:update:startsTensor:endsTensor:stridesTensor:startMask:endMask:squeezeMask:name:)
func (g_ Graph) SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorStartMaskEndMaskSqueezeMaskName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, startsTensor IMPSGraphTensor, endsTensor IMPSGraphTensor, stridesTensor IMPSGraphTensor, startMask uint32 /* not a class type */, endMask uint32 /* not a class type */, squeezeMask uint32 /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sliceUpdateDataTensor:updateTensor:startsTensor:endsTensor:stridesTensor:startMask:endMask:squeezeMask:name:"), dataTensor, updateTensor, startsTensor, endsTensor, stridesTensor, startMask, endMask, squeezeMask, name)
	return rv
}/* debug [instance_methods/method]: SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorStartMaskEndMaskSqueezeMaskName */


// Computes the softmax function on the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/softMax(with:axis:name:)
func (g_ Graph) SoftMaxWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("softMaxWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: SoftMaxWithTensorAxisName */


// Creates a softmax cross-entropy loss operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/softMaxCrossEntropy(_:labels:axis:reuctionType:name:)
func (g_ Graph) SoftMaxCrossEntropyWithSourceTensorLabelsTensorAxisReductionTypeName(sourceTensor IMPSGraphTensor, labelsTensor IMPSGraphTensor, axis int, reductionType GraphLossReductionType, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("softMaxCrossEntropyWithSourceTensor:labelsTensor:axis:reductionType:name:"), sourceTensor, labelsTensor, axis, reductionType, name)
	return rv
}/* debug [instance_methods/method]: SoftMaxCrossEntropyWithSourceTensorLabelsTensorAxisReductionTypeName */


// Creates the gradient of a softmax cross-entropy loss operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/softMaxCrossEntropyGradient(_:source:labels:axis:reuctionType:name:)
func (g_ Graph) SoftMaxCrossEntropyGradientWithIncomingGradientTensorSourceTensorLabelsTensorAxisReductionTypeName(gradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, labelsTensor IMPSGraphTensor, axis int, reductionType GraphLossReductionType, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("softMaxCrossEntropyGradientWithIncomingGradientTensor:sourceTensor:labelsTensor:axis:reductionType:name:"), gradientTensor, sourceTensor, labelsTensor, axis, reductionType, name)
	return rv
}/* debug [instance_methods/method]: SoftMaxCrossEntropyGradientWithIncomingGradientTensorSourceTensorLabelsTensorAxisReductionTypeName */


// Computes the gradient of the softmax function along the specified axis using the incoming gradient tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/softMaxGradient(withIncomingGradient:sourceTensor:axis:name:)
func (g_ Graph) SoftMaxGradientWithIncomingGradientSourceTensorAxisName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("softMaxGradientWithIncomingGradient:sourceTensor:axis:name:"), gradient, source, axis, name)
	return rv
}/* debug [instance_methods/method]: SoftMaxGradientWithIncomingGradientSourceTensorAxisName */


// Sorts the elements of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sort(_:axis:descending:name:)
func (g_ Graph) SortWithTensorAxisDescendingName(tensor IMPSGraphTensor, axis int, descending bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sortWithTensor:axis:descending:name:"), tensor, axis, descending, name)
	return rv
}/* debug [instance_methods/method]: SortWithTensorAxisDescendingName */


// Sorts the elements of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sort(_:axis:name:)
func (g_ Graph) SortWithTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sortWithTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: SortWithTensorAxisName */


// Sorts the elements of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sort(_:axisTensor:descending:name:)
func (g_ Graph) SortWithTensorAxisTensorDescendingName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, descending bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sortWithTensor:axisTensor:descending:name:"), tensor, axisTensor, descending, name)
	return rv
}/* debug [instance_methods/method]: SortWithTensorAxisTensorDescendingName */


// Sorts the elements of the input tensor along the specified axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sort(_:axisTensor:name:)
func (g_ Graph) SortWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sortWithTensor:axisTensor:name:"), tensor, axisTensor, name)
	return rv
}/* debug [instance_methods/method]: SortWithTensorAxisTensorName */


// Creates a space-to-depth2D operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/space(toDepth2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:)
func (g_ Graph) SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("spaceToDepth2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:"), tensor, widthAxis, heightAxis, depthAxis, blockSize, usePixelShuffleOrder, name)
	return rv
}/* debug [instance_methods/method]: SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName */


// Creates a space-to-depth2D operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/space(toDepth2DTensor:widthAxisTensor:heightAxisTensor:depthAxisTensor:blockSize:usePixelShuffleOrder:name:)
func (g_ Graph) SpaceToDepth2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxisTensor IMPSGraphTensor, heightAxisTensor IMPSGraphTensor, depthAxisTensor IMPSGraphTensor, blockSize uint, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("spaceToDepth2DTensor:widthAxisTensor:heightAxisTensor:depthAxisTensor:blockSize:usePixelShuffleOrder:name:"), tensor, widthAxisTensor, heightAxisTensor, depthAxisTensor, blockSize, usePixelShuffleOrder, name)
	return rv
}/* debug [instance_methods/method]: SpaceToDepth2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName */


// Creates a space-to-batch operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/spaceToBatch(_:spatialAxes:batchAxis:blockDimensions:usePixelShuffleOrder:name:)
func (g_ Graph) SpaceToBatchTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxes []foundation.Number, batchAxis int, blockDimensions []foundation.Number, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("spaceToBatchTensor:spatialAxes:batchAxis:blockDimensions:usePixelShuffleOrder:name:"), tensor, spatialAxes, batchAxis, blockDimensions, usePixelShuffleOrder, name)
	return rv
}/* debug [instance_methods/method]: SpaceToBatchTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName */


// Creates a space-to-batch operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/spaceToBatch(_:spatialAxesTensor:batchAxisTensor:blockDimensionsTensor:usePixelShuffleOrder:name:)
func (g_ Graph) SpaceToBatchTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxesTensor IMPSGraphTensor, batchAxisTensor IMPSGraphTensor, blockDimensionsTensor IMPSGraphTensor, usePixelShuffleOrder bool, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("spaceToBatchTensor:spatialAxesTensor:batchAxisTensor:blockDimensionsTensor:usePixelShuffleOrder:name:"), tensor, spatialAxesTensor, batchAxisTensor, blockDimensionsTensor, usePixelShuffleOrder, name)
	return rv
}/* debug [instance_methods/method]: SpaceToBatchTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName */


// Creates a sparse tensor representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sparseTensor(sparseTensorWithDescriptor:tensors:shape:name:)
func (g_ Graph) SparseTensorWithDescriptorTensorsShapeName(sparseDescriptor IMPSGraphCreateSparseOpDescriptor, inputTensorArray []GraphTensor, shape Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sparseTensorWithDescriptor:tensors:shape:name:"), sparseDescriptor, inputTensorArray, shape, name)
	return rv
}/* debug [instance_methods/method]: SparseTensorWithDescriptorTensorsShapeName */


// Creates a sparse tensor representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sparseTensor(sparseTensorWithType:tensors:shape:dataType:name:)
func (g_ Graph) SparseTensorWithTypeTensorsShapeDataTypeName(sparseStorageType GraphSparseStorageType, inputTensorArray []GraphTensor, shape Shape /* not a class type */, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("sparseTensorWithType:tensors:shape:dataType:name:"), sparseStorageType, inputTensorArray, shape, dataType, name)
	return rv
}/* debug [instance_methods/method]: SparseTensorWithTypeTensorsShapeDataTypeName */


// Creates a split operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/split(_:numSplits:axis:name:)
func (g_ Graph) SplitTensorNumSplitsAxisName(tensor IMPSGraphTensor, numSplits uint, axis int, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("splitTensor:numSplits:axis:name:"), tensor, numSplits, axis, name)
	return rv
}/* debug [instance_methods/method]: SplitTensorNumSplitsAxisName */


// Creates a split operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/split(_:splitSizes:axis:name:)
func (g_ Graph) SplitTensorSplitSizesAxisName(tensor IMPSGraphTensor, splitSizes []foundation.Number, axis int, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("splitTensor:splitSizes:axis:name:"), tensor, splitSizes, axis, name)
	return rv
}/* debug [instance_methods/method]: SplitTensorSplitSizesAxisName */


// Creates a split operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/split(_:splitSizesTensor:axis:name:)
func (g_ Graph) SplitTensorSplitSizesTensorAxisName(tensor IMPSGraphTensor, splitSizesTensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("splitTensor:splitSizesTensor:axis:name:"), tensor, splitSizesTensor, axis, name)
	return rv
}/* debug [instance_methods/method]: SplitTensorSplitSizesTensorAxisName */


// Applies the square operation to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/square(with:name:)
func (g_ Graph) SquareWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("squareWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: SquareWithTensorName */


// Applies the square root operation to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/squareRoot(with:name:)
func (g_ Graph) SquareRootWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("squareRootWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: SquareRootWithTensorName */


// Creates a squeeze operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/squeeze(_:axes:name:)
func (g_ Graph) SqueezeTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("squeezeTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: SqueezeTensorAxesName */


// Creates a squeeze operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/squeeze(_:axesTensor:name:)
func (g_ Graph) SqueezeTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("squeezeTensor:axesTensor:name:"), tensor, axesTensor, name)
	return rv
}/* debug [instance_methods/method]: SqueezeTensorAxesTensorName */


// Creates a squeeze operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/squeeze(_:axis:name:)
func (g_ Graph) SqueezeTensorAxisName(tensor IMPSGraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("squeezeTensor:axis:name:"), tensor, axis, name)
	return rv
}/* debug [instance_methods/method]: SqueezeTensorAxisName */


// Creates a squeeze operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/squeeze(_:name:)
func (g_ Graph) SqueezeTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("squeezeTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: SqueezeTensorName */


// Creates a stack operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/stack(_:axis:name:)
func (g_ Graph) StackTensorsAxisName(inputTensors []GraphTensor, axis int, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("stackTensors:axis:name:"), inputTensors, axis, name)
	return rv
}/* debug [instance_methods/method]: StackTensorsAxisName */


// Creates a stencil operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/stencil(withSourceTensor:weightsTensor:descriptor:name:)
func (g_ Graph) StencilWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphStencilOpDescriptor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("stencilWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, name)
	return rv
}/* debug [instance_methods/method]: StencilWithSourceTensorWeightsTensorDescriptorName */


// The Stochastic gradient descent performs a gradient descent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/stochasticGradientDescent(learningRate:values:gradient:name:)
func (g_ Graph) StochasticGradientDescentWithLearningRateTensorValuesTensorGradientTensorName(learningRateTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("stochasticGradientDescentWithLearningRateTensor:valuesTensor:gradientTensor:name:"), learningRateTensor, valuesTensor, gradientTensor, name)
	return rv
}/* debug [instance_methods/method]: StochasticGradientDescentWithLearningRateTensorValuesTensorGradientTensorName */


// Subtracts the second input tensor from the first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/subtraction(_:_:name:)
func (g_ Graph) SubtractionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("subtractionWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, name)
	return rv
}/* debug [instance_methods/method]: SubtractionWithPrimaryTensorSecondaryTensorName */


// Creates a tile gradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/tileGradient(withIncomingGradientTensor:sourceTensor:withMultiplier:name:)
func (g_ Graph) TileGradientWithIncomingGradientTensorSourceTensorWithMultiplierName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, multiplier Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("tileGradientWithIncomingGradientTensor:sourceTensor:withMultiplier:name:"), incomingGradientTensor, sourceTensor, multiplier, name)
	return rv
}/* debug [instance_methods/method]: TileGradientWithIncomingGradientTensorSourceTensorWithMultiplierName */


// Creates a tile operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/tileTensor(_:withMultiplier:name:)
func (g_ Graph) TileTensorWithMultiplierName(tensor IMPSGraphTensor, multiplier Shape /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("tileTensor:withMultiplier:name:"), tensor, multiplier, name)
	return rv
}/* debug [instance_methods/method]: TileTensorWithMultiplierName */


// Creates a TopK operation and returns the value and indices tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topK(_:axis:k:name:)
func (g_ Graph) TopKWithSourceTensorAxisKName(source IMPSGraphTensor, axis int, k uint, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("topKWithSourceTensor:axis:k:name:"), source, axis, k, name)
	return rv
}/* debug [instance_methods/method]: TopKWithSourceTensorAxisKName */


// Creates a TopK operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topK(_:axisTensor:kTensor:name:)
func (g_ Graph) TopKWithSourceTensorAxisTensorKTensorName(source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("topKWithSourceTensor:axisTensor:kTensor:name:"), source, axisTensor, kTensor, name)
	return rv
}/* debug [instance_methods/method]: TopKWithSourceTensorAxisTensorKTensorName */


// Creates a TopK operation and returns the value and indices tensors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topK(_:k:name:)
func (g_ Graph) TopKWithSourceTensorKName(source IMPSGraphTensor, k uint, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("topKWithSourceTensor:k:name:"), source, k, name)
	return rv
}/* debug [instance_methods/method]: TopKWithSourceTensorKName */


// Creates a TopK operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topK(_:kTensor:name:)
func (g_ Graph) TopKWithSourceTensorKTensorName(source IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("topKWithSourceTensor:kTensor:name:"), source, kTensor, name)
	return rv
}/* debug [instance_methods/method]: TopKWithSourceTensorKTensorName */


// Creates a TopKGradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topKGradient(_:input:k:name:)
func (g_ Graph) TopKWithGradientTensorSourceKName(gradient IMPSGraphTensor, source IMPSGraphTensor, k uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("topKWithGradientTensor:source:k:name:"), gradient, source, k, name)
	return rv
}/* debug [instance_methods/method]: TopKWithGradientTensorSourceKName */


// Creates a TopKGradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topKGradient(_:input:kTensor:name:)
func (g_ Graph) TopKWithGradientTensorSourceKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("topKWithGradientTensor:source:kTensor:name:"), gradient, source, kTensor, name)
	return rv
}/* debug [instance_methods/method]: TopKWithGradientTensorSourceKTensorName */


// Creates a TopKGradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topKGradient(_:source:axis:k:name:)
func (g_ Graph) TopKWithGradientTensorSourceAxisKName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, k uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("topKWithGradientTensor:source:axis:k:name:"), gradient, source, axis, k, name)
	return rv
}/* debug [instance_methods/method]: TopKWithGradientTensorSourceAxisKName */


// Creates a TopKGradient operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topKGradient(_:source:axisTensor:kTensor:name:)
func (g_ Graph) TopKWithGradientTensorSourceAxisTensorKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("topKWithGradientTensor:source:axisTensor:kTensor:name:"), gradient, source, axisTensor, kTensor, name)
	return rv
}/* debug [instance_methods/method]: TopKWithGradientTensorSourceAxisTensorKTensorName */


// Creates a permutation operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/transpose(_:permutation:name:)
func (g_ Graph) TransposeTensorPermutationName(tensor IMPSGraphTensor, permutation []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("transposeTensor:permutation:name:"), tensor, permutation, name)
	return rv
}/* debug [instance_methods/method]: TransposeTensorPermutationName */


// Creates a transpose operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/transposeTensor(_:dimension:withDimension:name:)
func (g_ Graph) TransposeTensorDimensionWithDimensionName(tensor IMPSGraphTensor, dimensionIndex uint, dimensionIndex2 uint, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("transposeTensor:dimension:withDimension:name:"), tensor, dimensionIndex, dimensionIndex2, name)
	return rv
}/* debug [instance_methods/method]: TransposeTensorDimensionWithDimensionName */


// Applies the truncate operation to the input tensor elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/truncate(_:name:)
func (g_ Graph) TruncateWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("truncateWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: TruncateWithTensorName */


// Creates a variable operation and returns the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/variable(with:shape:dataType:name:)
func (g_ Graph) VariableWithDataShapeDataTypeName(data objc.IObject /* cross-framework: NSData */, shape Shape /* not a class type */, dataType DataType /* not a class type */, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("variableWithData:shape:dataType:name:"), data, shape, dataType, name)
	return rv
}/* debug [instance_methods/method]: VariableWithDataShapeDataTypeName */


// Creates a variable from an input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/variableFromTensor(_:name:)
func (g_ Graph) VariableFromTensorWithTensorName(tensor IMPSGraphTensor, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("variableFromTensorWithTensor:name:"), tensor, name)
	return rv
}/* debug [instance_methods/method]: VariableFromTensorWithTensorName */


// Returns the variance of the first input along the specified axes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/variance(of:axes:name:)
func (g_ Graph) VarianceOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("varianceOfTensor:axes:name:"), tensor, axes, name)
	return rv
}/* debug [instance_methods/method]: VarianceOfTensorAxesName */


// Returns the variance of the first input along the specified axes when the mean has been precomputed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/variance(of:mean:axes:name:)
func (g_ Graph) VarianceOfTensorMeanTensorAxesName(tensor IMPSGraphTensor, meanTensor IMPSGraphTensor, axes []foundation.Number, name objc.IObject /* cross-framework: NSString */) IGraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("varianceOfTensor:meanTensor:axes:name:"), tensor, meanTensor, axes, name)
	return rv
}/* debug [instance_methods/method]: VarianceOfTensorMeanTensorAxesName */


// Adds a while loop operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/while(initialInputs:before:after:name:)
func (g_ Graph) WhileWithInitialInputsBeforeAfterName(initialInputs []GraphTensor, before GraphWhileBeforeBlock /* not a class type */, after GraphWhileAfterBlock /* not a class type */, name objc.IObject /* cross-framework: NSString */) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("whileWithInitialInputs:before:after:name:"), initialInputs, before, after, name)
	return rv
}/* debug [instance_methods/method]: WhileWithInitialInputsBeforeAfterName */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Graph */

// Options for the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/options
func (g_ Graph) Options() GraphOptions {
	rv := objc.Send[GraphOptions](g_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// Options for the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/options
func (g_ Graph) SetOptions(value GraphOptions) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOptions:"), value)
}/* debug [instance_properties/setter]: options */


// Array of all the placeholder tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/placeholderTensors
func (g_ Graph) PlaceholderTensors() []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("placeholderTensors"))
	return rv
}/* debug [instance_properties/getter]: placeholderTensors */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraph */


