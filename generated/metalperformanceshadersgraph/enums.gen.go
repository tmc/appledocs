// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

// Enum types and constants
// MPSGraphDeploymentPlatform - The options available to a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeploymentPlatform
type MPSGraphDeploymentPlatform uint

// MPSGraphDeviceType - The device type.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeviceType
type MPSGraphDeviceType uint

const (
	// MPSGraphDeviceTypeMetal - Device of type Metal
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeviceType/metal
	MPSGraphDeviceTypeMetal MPSGraphDeviceType = 0
)

// MPSGraphExecutionStage - Execution events that can be used with shared events.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionStage
type MPSGraphExecutionStage uint

const (
	// MPSGraphExecutionStageCompleted - stage when execution of the graph completes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionStage/completed
	MPSGraphExecutionStageCompleted MPSGraphExecutionStage = 0
)

// MPSGraphFFTScalingMode - The scaling modes for Fourier transform operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTScalingMode
type MPSGraphFFTScalingMode uint

const (
	// MPSGraphFFTScalingModeUnitary - Scales the FFT result with reciprocal square root of the total FFT size over all transformed dimensions, resulting in signal strength conserving transformation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTScalingMode/unitary
	MPSGraphFFTScalingModeUnitary MPSGraphFFTScalingMode = 0
)

// MPSGraphLossReductionType - The type of the reduction the graph applies in the loss operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType
type MPSGraphLossReductionType uint

const (
	// MPSGraphLossReductionTypeAxis - Computes the loss without reduction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType/axis
	MPSGraphLossReductionTypeAxis MPSGraphLossReductionType = 0
)

// MPSGraphNonMaximumSuppressionCoordinateMode - The non-maximum suppression coordinate mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphNonMaximumSuppressionCoordinateMode
type MPSGraphNonMaximumSuppressionCoordinateMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphNonMaximumSuppressionCoordinateMode/centersWidthFirst
	MPSGraphNonMaximumSuppressionCoordinateModeCentersWidthFirst MPSGraphNonMaximumSuppressionCoordinateMode = 0
)

// MPSGraphOptimization - The optimization levels to trade compilation time for even more runtime performance by running more passes.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimization
type MPSGraphOptimization uint

const (
	// MPSGraphOptimizationLevel1 - Graph performs additional Optimizations, like using the placement pass to dispatch across different HW blocks like the NeuralEngine and CPU along with the GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimization/level1
	MPSGraphOptimizationLevel1 MPSGraphOptimization = 0
)

// MPSGraphOptimizationProfile - The optimization profile used as a heuristic as the graph compiler optimizes the network.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimizationProfile
type MPSGraphOptimizationProfile uint

const (
	// MPSGraphOptimizationProfilePerformance - Default, graph optimized for performance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimizationProfile/performance
	MPSGraphOptimizationProfilePerformance MPSGraphOptimizationProfile = 0
)

// MPSGraphOptions - The options available to a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptions
type MPSGraphOptions uint

const (
	// MPSGraphOptionsDefault - The framework uses these options as default if not overriden.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptions/default
	MPSGraphOptionsDefault MPSGraphOptions = 0
)

// MPSGraphPaddingMode - The tensor padding mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode
type MPSGraphPaddingMode uint

// MPSGraphPaddingStyle - The tensor padding style.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle
type MPSGraphPaddingStyle uint

const (
	// MPSGraphPaddingStyleExplicit - Explicit
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle/explicit
	MPSGraphPaddingStyleExplicit MPSGraphPaddingStyle = 0
)

// MPSGraphPoolingReturnIndicesMode - The flattening mode for returned indices with max-pooling.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode
type MPSGraphPoolingReturnIndicesMode uint

const (
	// MPSGraphPoolingReturnIndicesLocalFlatten3D - Returns indices within pooling window, flattened in 3 innernost dimensions. eg: HWC in NHWC.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/localFlatten3D
	MPSGraphPoolingReturnIndicesLocalFlatten3D MPSGraphPoolingReturnIndicesMode = 0
)

// MPSGraphRNNActivation - The activation modes for RNN operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation
type MPSGraphRNNActivation uint

const (
	// MPSGraphRNNActivationHardSigmoid - Defines a Hard sigmoid activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation/hardSigmoid
	MPSGraphRNNActivationHardSigmoid MPSGraphRNNActivation = 0
)

// MPSGraphRandomDistribution - The distributions supported by random operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomDistribution
type MPSGraphRandomDistribution uint

const (
	// MPSGraphRandomDistributionUniform - The uniform distribution, with samples drawn uniformly from [min, max) for float types, and [min, max] for integer types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomDistribution/uniform
	MPSGraphRandomDistributionUniform MPSGraphRandomDistribution = 0
)

// MPSGraphRandomNormalSamplingMethod - The sampling method to use when generating values in the normal distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomNormalSamplingMethod
type MPSGraphRandomNormalSamplingMethod uint

const (
	// MPSGraphRandomNormalSamplingBoxMuller - Use Box Muller transform to convert uniform values to values in the normal distribution. For bounded distributions this is a rejection sampling method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomNormalSamplingMethod/boxMuller
	MPSGraphRandomNormalSamplingBoxMuller MPSGraphRandomNormalSamplingMethod = 0
)

// MPSGraphReducedPrecisionFastMath - MPSGraph could use these reduced precision paths to deliver faster math, but it is not guaranteed.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReducedPrecisionFastMath
type MPSGraphReducedPrecisionFastMath uint

// MPSGraphReductionMode - The reduction mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode
type MPSGraphReductionMode uint

const (
	// MPSGraphReductionModeMax - Max
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode/max
	MPSGraphReductionModeMax MPSGraphReductionMode = 0
)

// MPSGraphResizeMode - The resize mode to use for resizing.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeMode
type MPSGraphResizeMode uint

// MPSGraphResizeNearestRoundingMode - The rounding mode to use when using nearest resize mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode
type MPSGraphResizeNearestRoundingMode uint

const (
	// MPSGraphResizeNearestRoundingModeRoundToOdd - Rounds values to the nearest integer value, with 0.5f rounding toward the closest odd value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode/roundToOdd
	MPSGraphResizeNearestRoundingModeRoundToOdd MPSGraphResizeNearestRoundingMode = 0
)

// MPSGraphScatterMode - The scatter mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode
type MPSGraphScatterMode uint

const (
	// MPSGraphScatterModeAdd - Add
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/add
	MPSGraphScatterModeAdd MPSGraphScatterMode = 0
	// MPSGraphScatterModeMin - Minimum
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/min
	MPSGraphScatterModeMin MPSGraphScatterMode = 0
)

// MPSGraphSparseStorageType - The sparse storage options in the Metal Performance Shaders Graph framework.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSparseStorageType
type MPSGraphSparseStorageType uint

const (
	// MPSGraphSparseStorageCOO - COO Storage
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSparseStorageType/COO
	MPSGraphSparseStorageCOO MPSGraphSparseStorageType = 0
)

// MPSGraphTensorNamedDataLayout - The tensor layout.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout
type MPSGraphTensorNamedDataLayout uint

const (
	// MPSGraphTensorNamedDataLayoutOIDHW - LayoutOIDHW
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/OIDHW
	MPSGraphTensorNamedDataLayoutOIDHW MPSGraphTensorNamedDataLayout = 0
)


