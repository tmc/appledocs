// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

// Enum types and constants
// MPSGraphDeploymentPlatform - The options available to a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeploymentPlatform
type GraphDeploymentPlatform uint

const (
	// GraphDeploymentPlatformVisionOS - Deployment target for visionOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeploymentPlatform/visionOS
	GraphDeploymentPlatformVisionOS GraphDeploymentPlatform = 0
)

// MPSGraphDeviceType - The device type.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeviceType
type GraphDeviceType uint

const (
	// GraphDeviceTypeMetal - Device of type Metal
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeviceType/metal
	GraphDeviceTypeMetal GraphDeviceType = 0
)

// MPSGraphExecutionStage - Execution events that can be used with shared events.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionStage
type GraphExecutionStage uint

const (
	// GraphExecutionStageCompleted - stage when execution of the graph completes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionStage/completed
	GraphExecutionStageCompleted GraphExecutionStage = 0
)

// MPSGraphFFTScalingMode - The scaling modes for Fourier transform operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTScalingMode
type GraphFFTScalingMode uint

const (
	// GraphFFTScalingModeUnitary - Scales the FFT result with reciprocal square root of the total FFT size over all transformed dimensions, resulting in signal strength conserving transformation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTScalingMode/unitary
	GraphFFTScalingModeUnitary GraphFFTScalingMode = 0
)

// MPSGraphLossReductionType - The type of the reduction the graph applies in the loss operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType
type GraphLossReductionType uint

const (
	// GraphLossReductionTypeAxis - Computes the loss without reduction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType/axis
	GraphLossReductionTypeAxis GraphLossReductionType = 0
	// GraphLossReductionTypeNone - Computes the loss without reduction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType/none
	GraphLossReductionTypeNone GraphLossReductionType = 0
)

// MPSGraphNonMaximumSuppressionCoordinateMode - The non-maximum suppression coordinate mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphNonMaximumSuppressionCoordinateMode
type GraphNonMaximumSuppressionCoordinateMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphNonMaximumSuppressionCoordinateMode/centersWidthFirst
	GraphNonMaximumSuppressionCoordinateModeCentersWidthFirst GraphNonMaximumSuppressionCoordinateMode = 0
)

// MPSGraphOptimization - The optimization levels to trade compilation time for even more runtime performance by running more passes.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimization
type GraphOptimization uint

const (
	// GraphOptimizationLevel1 - Graph performs additional Optimizations, like using the placement pass to dispatch across different HW blocks like the NeuralEngine and CPU along with the GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimization/level1
	GraphOptimizationLevel1 GraphOptimization = 0
)

// MPSGraphOptimizationProfile - The optimization profile used as a heuristic as the graph compiler optimizes the network.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimizationProfile
type GraphOptimizationProfile uint

const (
	// GraphOptimizationProfilePerformance - Default, graph optimized for performance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimizationProfile/performance
	GraphOptimizationProfilePerformance GraphOptimizationProfile = 0
	// GraphOptimizationProfilePowerEfficiency - Graph optimized for power efficiency.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimizationProfile/powerEfficiency
	GraphOptimizationProfilePowerEfficiency GraphOptimizationProfile = 0
)

// MPSGraphOptions - The options available to a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptions
type GraphOptions uint

const (
	// GraphOptionsDefault - The framework uses these options as default if not overriden.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptions/default
	GraphOptionsDefault GraphOptions = 0
	// GraphOptionsNone - No Options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptions/none
	GraphOptionsNone GraphOptions = 0
)

// MPSGraphPaddingMode - The tensor padding mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode
type GraphPaddingMode uint

const (
	// GraphPaddingModeAntiPeriodic - Anti Periodic 
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode/antiPeriodic
	GraphPaddingModeAntiPeriodic GraphPaddingMode = 0
	// GraphPaddingModeSymmetric - Symmetric
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode/symmetric
	GraphPaddingModeSymmetric GraphPaddingMode = 0
)

// MPSGraphPaddingStyle - The tensor padding style.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle
type GraphPaddingStyle uint

const (
	// GraphPaddingStyleTF_SAME - TF_SAME
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle/TF_SAME
	GraphPaddingStyleTF_SAME GraphPaddingStyle = 0
	// GraphPaddingStyleExplicit - Explicit
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle/explicit
	GraphPaddingStyleExplicit GraphPaddingStyle = 0
)

// MPSGraphPoolingReturnIndicesMode - The flattening mode for returned indices with max-pooling.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode
type GraphPoolingReturnIndicesMode uint

const (
	// GraphPoolingReturnIndicesLocalFlatten3D - Returns indices within pooling window, flattened in 3 innernost dimensions. eg: HWC in NHWC.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/localFlatten3D
	GraphPoolingReturnIndicesLocalFlatten3D GraphPoolingReturnIndicesMode = 0
	// GraphPoolingReturnIndicesNone - No indices returned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/none
	GraphPoolingReturnIndicesNone GraphPoolingReturnIndicesMode = 0
)

// MPSGraphRNNActivation - The activation modes for RNN operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation
type GraphRNNActivation uint

const (
	// GraphRNNActivationHardSigmoid - Defines a Hard sigmoid activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation/hardSigmoid
	GraphRNNActivationHardSigmoid GraphRNNActivation = 0
	// GraphRNNActivationNone - Defines a pass through activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation/none
	GraphRNNActivationNone GraphRNNActivation = 0
	// GraphRNNActivationRelu - Defines a ReLU activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation/relu
	GraphRNNActivationRelu GraphRNNActivation = 0
)

// MPSGraphRandomDistribution - The distributions supported by random operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomDistribution
type GraphRandomDistribution uint

const (
	// GraphRandomDistributionTruncatedNormal - The normal distribution defined by mean and standard deviation, truncated to the range [min, max)
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomDistribution/truncatedNormal
	GraphRandomDistributionTruncatedNormal GraphRandomDistribution = 0
	// GraphRandomDistributionUniform - The uniform distribution, with samples drawn uniformly from [min, max) for float types, and [min, max] for integer types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomDistribution/uniform
	GraphRandomDistributionUniform GraphRandomDistribution = 0
)

// MPSGraphRandomNormalSamplingMethod - The sampling method to use when generating values in the normal distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomNormalSamplingMethod
type GraphRandomNormalSamplingMethod uint

const (
	// GraphRandomNormalSamplingBoxMuller - Use Box Muller transform to convert uniform values to values in the normal distribution. For bounded distributions this is a rejection sampling method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomNormalSamplingMethod/boxMuller
	GraphRandomNormalSamplingBoxMuller GraphRandomNormalSamplingMethod = 0
	// GraphRandomNormalSamplingInvCDF - Use inverse erf to convert uniform values to values in the normal distribution
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomNormalSamplingMethod/invCDF
	GraphRandomNormalSamplingInvCDF GraphRandomNormalSamplingMethod = 0
)

// MPSGraphReducedPrecisionFastMath - MPSGraph could use these reduced precision paths to deliver faster math, but it is not guaranteed.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReducedPrecisionFastMath
type GraphReducedPrecisionFastMath uint

// MPSGraphReductionMode - The reduction mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode
type GraphReductionMode uint

const (
	// GraphReductionModeMax - Max
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode/max
	GraphReductionModeMax GraphReductionMode = 0
	// GraphReductionModeSum - Sum
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode/sum
	GraphReductionModeSum GraphReductionMode = 0
)

// MPSGraphResizeMode - The resize mode to use for resizing.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeMode
type GraphResizeMode uint

// MPSGraphResizeNearestRoundingMode - The rounding mode to use when using nearest resize mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode
type GraphResizeNearestRoundingMode uint

const (
	// GraphResizeNearestRoundingModeFloor - Rounds values toward -inf.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode/floor
	GraphResizeNearestRoundingModeFloor GraphResizeNearestRoundingMode = 0
	// GraphResizeNearestRoundingModeRoundToOdd - Rounds values to the nearest integer value, with 0.5f rounding toward the closest odd value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode/roundToOdd
	GraphResizeNearestRoundingModeRoundToOdd GraphResizeNearestRoundingMode = 0
)

// MPSGraphScatterMode - The scatter mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode
type GraphScatterMode uint

const (
	// GraphScatterModeAdd - Add
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/add
	GraphScatterModeAdd GraphScatterMode = 0
	// GraphScatterModeMax - Maximum
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/max
	GraphScatterModeMax GraphScatterMode = 0
	// GraphScatterModeMin - Minimum
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/min
	GraphScatterModeMin GraphScatterMode = 0
)

// MPSGraphSparseStorageType - The sparse storage options in the Metal Performance Shaders Graph framework.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSparseStorageType
type GraphSparseStorageType uint

const (
	// GraphSparseStorageCOO - COO Storage
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSparseStorageType/COO
	GraphSparseStorageCOO GraphSparseStorageType = 0
	// GraphSparseStorageCSC - CSC Storage
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSparseStorageType/CSC
	GraphSparseStorageCSC GraphSparseStorageType = 0
)

// MPSGraphTensorNamedDataLayout - The tensor layout.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout
type GraphTensorNamedDataLayout uint

const (
	// GraphTensorNamedDataLayoutNCDHW - LayoutNCDHW
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/NCDHW
	GraphTensorNamedDataLayoutNCDHW GraphTensorNamedDataLayout = 0
	// GraphTensorNamedDataLayoutNDHWC - LayoutNDHWC
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/NDHWC
	GraphTensorNamedDataLayoutNDHWC GraphTensorNamedDataLayout = 0
	// GraphTensorNamedDataLayoutOIDHW - LayoutOIDHW
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/OIDHW
	GraphTensorNamedDataLayoutOIDHW GraphTensorNamedDataLayout = 0
)


