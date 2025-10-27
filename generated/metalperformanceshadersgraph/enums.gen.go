// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph


// Enum types and constants

// MPSGraphDeploymentPlatform - The options available to a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeploymentPlatform
type MPSGraphDeploymentPlatform uint

const (
	// MPSGraphDeploymentPlatformIOS - Deployment target for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeploymentPlatform/iOS
	MPSGraphDeploymentPlatformIOS MPSGraphDeploymentPlatform = 0
	// MPSGraphDeploymentPlatformMacOS - Deployment platofmr for macOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeploymentPlatform/macOS
	MPSGraphDeploymentPlatformMacOS MPSGraphDeploymentPlatform = 0
	// MPSGraphDeploymentPlatformTvOS - Deployment target for tvOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeploymentPlatform/tvOS
	MPSGraphDeploymentPlatformTvOS MPSGraphDeploymentPlatform = 0
	// MPSGraphDeploymentPlatformVisionOS - Deployment target for visionOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeploymentPlatform/visionOS
	MPSGraphDeploymentPlatformVisionOS MPSGraphDeploymentPlatform = 0
)


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
	// MPSGraphFFTScalingModeNone - Computes the FFT result with no scaling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTScalingMode/none
	MPSGraphFFTScalingModeNone MPSGraphFFTScalingMode = 0
	// MPSGraphFFTScalingModeSize - Scales the FFT result with reciprocal of the total FFT size over all transformed dimensions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTScalingMode/size
	MPSGraphFFTScalingModeSize MPSGraphFFTScalingMode = 0
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
	// MPSGraphLossReductionTypeMean - Reduces the loss down to a scalar with a mean operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType/mean
	MPSGraphLossReductionTypeMean MPSGraphLossReductionType = 0
	// MPSGraphLossReductionTypeNone - Computes the loss without reduction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType/none
	MPSGraphLossReductionTypeNone MPSGraphLossReductionType = 0
	// MPSGraphLossReductionTypeSum - Reduces the loss down to a scalar with a sum operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType/sum
	MPSGraphLossReductionTypeSum MPSGraphLossReductionType = 0
)


// MPSGraphNonMaximumSuppressionCoordinateMode - The non-maximum suppression coordinate mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphNonMaximumSuppressionCoordinateMode
type MPSGraphNonMaximumSuppressionCoordinateMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphNonMaximumSuppressionCoordinateMode/centersHeightFirst
	MPSGraphNonMaximumSuppressionCoordinateModeCentersHeightFirst MPSGraphNonMaximumSuppressionCoordinateMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphNonMaximumSuppressionCoordinateMode/centersWidthFirst
	MPSGraphNonMaximumSuppressionCoordinateModeCentersWidthFirst MPSGraphNonMaximumSuppressionCoordinateMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphNonMaximumSuppressionCoordinateMode/cornersWidthFirst
	MPSGraphNonMaximumSuppressionCoordinateModeCornersWidthFirst MPSGraphNonMaximumSuppressionCoordinateMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphNonMaximumSuppressionCoordinateMode/explicit
	MPSGraphNonMaximumSuppressionCoordinateModeCornersHeightFirst MPSGraphNonMaximumSuppressionCoordinateMode = 0
)


// MPSGraphOptimization - The optimization levels to trade compilation time for even more runtime performance by running more passes.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimization
type MPSGraphOptimization uint

const (
	// MPSGraphOptimizationLevel0 - Graph performs core optimizations only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimization/level0
	MPSGraphOptimizationLevel0 MPSGraphOptimization = 0
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
	// MPSGraphOptimizationProfilePowerEfficiency - Graph optimized for power efficiency.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimizationProfile/powerEfficiency
	MPSGraphOptimizationProfilePowerEfficiency MPSGraphOptimizationProfile = 0
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
	// MPSGraphOptionsNone - No Options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptions/none
	MPSGraphOptionsNone MPSGraphOptions = 0
	// MPSGraphOptionsSynchronizeResults - The graph synchronizes results to the CPU using a blit encoder if on a discrete GPU at the end of execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptions/synchronizeResults
	MPSGraphOptionsSynchronizeResults MPSGraphOptions = 0
	// MPSGraphOptionsVerbose - The framework prints more logging info.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptions/verbose
	MPSGraphOptionsVerbose MPSGraphOptions = 0
)


// MPSGraphPaddingMode - The tensor padding mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode
type MPSGraphPaddingMode uint

const (
	// MPSGraphPaddingModeAntiPeriodic - Anti Periodic 
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode/antiPeriodic
	MPSGraphPaddingModeAntiPeriodic MPSGraphPaddingMode = 0
	// MPSGraphPaddingModeClampToEdge - ClampToEdge (PyTorch ReplicationPad)
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode/clampToEdge
	MPSGraphPaddingModeClampToEdge MPSGraphPaddingMode = 0
	// MPSGraphPaddingModeConstant - Constant
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode/constant
	MPSGraphPaddingModeConstant MPSGraphPaddingMode = 0
	// MPSGraphPaddingModePeriodic - Periodic 
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode/periodic
	MPSGraphPaddingModePeriodic MPSGraphPaddingMode = 0
	// MPSGraphPaddingModeReflect - Reflect
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode/reflect
	MPSGraphPaddingModeReflect MPSGraphPaddingMode = 0
	// MPSGraphPaddingModeSymmetric - Symmetric
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode/symmetric
	MPSGraphPaddingModeSymmetric MPSGraphPaddingMode = 0
	// MPSGraphPaddingModeZero - Zero
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode/zero
	MPSGraphPaddingModeZero MPSGraphPaddingMode = 0
)


// MPSGraphPaddingStyle - The tensor padding style.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle
type MPSGraphPaddingStyle uint

const (
	// MPSGraphPaddingStyleExplicit - Explicit
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle/explicit
	MPSGraphPaddingStyleExplicit MPSGraphPaddingStyle = 0
	// MPSGraphPaddingStyleExplicitOffset - TF_VALID
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle/explicitOffset
	MPSGraphPaddingStyleExplicitOffset MPSGraphPaddingStyle = 0
	// MPSGraphPaddingStyleONNX_SAME_LOWER - Explicit offsets
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle/ONNX_SAME_LOWER
	MPSGraphPaddingStyleONNX_SAME_LOWER MPSGraphPaddingStyle = 0
	// MPSGraphPaddingStyleTF_SAME - TF_SAME
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle/TF_SAME
	MPSGraphPaddingStyleTF_SAME MPSGraphPaddingStyle = 0
	// MPSGraphPaddingStyleTF_VALID - ONNX_SAME_LOWER
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle/TF_VALID
	MPSGraphPaddingStyleTF_VALID MPSGraphPaddingStyle = 0
)


// MPSGraphPoolingReturnIndicesMode - The flattening mode for returned indices with max-pooling.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode
type MPSGraphPoolingReturnIndicesMode uint

const (
	// MPSGraphPoolingReturnIndicesGlobalFlatten1D - Returns indices flattened in inner most (last) dimension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/globalFlatten1D
	MPSGraphPoolingReturnIndicesGlobalFlatten1D MPSGraphPoolingReturnIndicesMode = 0
	// MPSGraphPoolingReturnIndicesGlobalFlatten2D - Returns indices flattened in 2 innermost dimensions. eg: HW in NCHW.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/globalFlatten2D
	MPSGraphPoolingReturnIndicesGlobalFlatten2D MPSGraphPoolingReturnIndicesMode = 0
	// MPSGraphPoolingReturnIndicesGlobalFlatten3D - Returns indices flattened in 3 innernost dimensions. eg: HWC in NHWC.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/globalFlatten3D
	MPSGraphPoolingReturnIndicesGlobalFlatten3D MPSGraphPoolingReturnIndicesMode = 0
	// MPSGraphPoolingReturnIndicesGlobalFlatten4D - Returns indices flattened in 4 innermost dimensions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/globalFlatten4D
	MPSGraphPoolingReturnIndicesGlobalFlatten4D MPSGraphPoolingReturnIndicesMode = 0
	// MPSGraphPoolingReturnIndicesLocalFlatten1D - Returns indices within pooling window, flattened in inner most dimension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/localFlatten1D
	MPSGraphPoolingReturnIndicesLocalFlatten1D MPSGraphPoolingReturnIndicesMode = 0
	// MPSGraphPoolingReturnIndicesLocalFlatten2D - Returns indices within pooling window, flattened in 2 innermost dimensions. eg: HW in NCHW.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/localFlatten2D
	MPSGraphPoolingReturnIndicesLocalFlatten2D MPSGraphPoolingReturnIndicesMode = 0
	// MPSGraphPoolingReturnIndicesLocalFlatten3D - Returns indices within pooling window, flattened in 3 innernost dimensions. eg: HWC in NHWC.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/localFlatten3D
	MPSGraphPoolingReturnIndicesLocalFlatten3D MPSGraphPoolingReturnIndicesMode = 0
	// MPSGraphPoolingReturnIndicesLocalFlatten4D - Returns indices within pooling window, flattened in 4 innermost dimensions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/localFlatten4D
	MPSGraphPoolingReturnIndicesLocalFlatten4D MPSGraphPoolingReturnIndicesMode = 0
	// MPSGraphPoolingReturnIndicesNone - No indices returned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode/none
	MPSGraphPoolingReturnIndicesNone MPSGraphPoolingReturnIndicesMode = 0
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
	// MPSGraphRNNActivationNone - Defines a pass through activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation/none
	MPSGraphRNNActivationNone MPSGraphRNNActivation = 0
	// MPSGraphRNNActivationRelu - Defines a ReLU activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation/relu
	MPSGraphRNNActivationRelu MPSGraphRNNActivation = 0
	// MPSGraphRNNActivationSigmoid - Defines a Sigmoid activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation/sigmoid
	MPSGraphRNNActivationSigmoid MPSGraphRNNActivation = 0
	// MPSGraphRNNActivationTanh - Defines a Tanh activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation/tanh
	MPSGraphRNNActivationTanh MPSGraphRNNActivation = 0
)


// MPSGraphRandomDistribution - The distributions supported by random operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomDistribution
type MPSGraphRandomDistribution uint

const (
	// MPSGraphRandomDistributionNormal - The normal distribution defined by mean and standard deviation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomDistribution/normal
	MPSGraphRandomDistributionNormal MPSGraphRandomDistribution = 0
	// MPSGraphRandomDistributionTruncatedNormal - The normal distribution defined by mean and standard deviation, truncated to the range [min, max)
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomDistribution/truncatedNormal
	MPSGraphRandomDistributionTruncatedNormal MPSGraphRandomDistribution = 0
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
	// MPSGraphRandomNormalSamplingInvCDF - Use inverse erf to convert uniform values to values in the normal distribution
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomNormalSamplingMethod/invCDF
	MPSGraphRandomNormalSamplingInvCDF MPSGraphRandomNormalSamplingMethod = 0
)


// MPSGraphReducedPrecisionFastMath - MPSGraph could use these reduced precision paths to deliver faster math, but it is not guaranteed.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReducedPrecisionFastMath
type MPSGraphReducedPrecisionFastMath uint

const (
	// MPSGraphReducedPrecisionFastMathAllowFP16Conv2DWinogradTransformIntermediate - Execute winograd transform intermediate as FP16.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReducedPrecisionFastMath/allowFP16Conv2DWinogradTransformIntermediate
	MPSGraphReducedPrecisionFastMathAllowFP16Conv2DWinogradTransformIntermediate MPSGraphReducedPrecisionFastMath = 0
	// MPSGraphReducedPrecisionFastMathAllowFP16Intermediates - Curated list allowing intermediates for multi-pass GPU kernels to be FP16.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReducedPrecisionFastMath/allowFP16Intermediates
	MPSGraphReducedPrecisionFastMathAllowFP16Intermediates MPSGraphReducedPrecisionFastMath = 0
	// MPSGraphReducedPrecisionFastMathDefault - Default selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReducedPrecisionFastMath/MPSGraphReducedPrecisionFastMathDefault
	MPSGraphReducedPrecisionFastMathDefault MPSGraphReducedPrecisionFastMath = 0
	// MPSGraphReducedPrecisionFastMathNone - Full precision math with maximum accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReducedPrecisionFastMath/none
	MPSGraphReducedPrecisionFastMathNone MPSGraphReducedPrecisionFastMath = 0
)


// MPSGraphReductionMode - The reduction mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode
type MPSGraphReductionMode uint

const (
	// MPSGraphReductionModeArgumentMax - Argument Max
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode/argumentMax
	MPSGraphReductionModeArgumentMax MPSGraphReductionMode = 0
	// MPSGraphReductionModeArgumentMin - Argument Min
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode/argumentMin
	MPSGraphReductionModeArgumentMin MPSGraphReductionMode = 0
	// MPSGraphReductionModeMax - Max
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode/max
	MPSGraphReductionModeMax MPSGraphReductionMode = 0
	// MPSGraphReductionModeMin - Min
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode/min
	MPSGraphReductionModeMin MPSGraphReductionMode = 0
	// MPSGraphReductionModeProduct - Product
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode/product
	MPSGraphReductionModeProduct MPSGraphReductionMode = 0
	// MPSGraphReductionModeSum - Sum
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode/sum
	MPSGraphReductionModeSum MPSGraphReductionMode = 0
)


// MPSGraphResizeMode - The resize mode to use for resizing.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeMode
type MPSGraphResizeMode uint

const (
	// MPSGraphResizeBilinear - Samples the 4 neighbors to the pixel coordinate and uses bilinear interpolation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeMode/bilinear
	MPSGraphResizeBilinear MPSGraphResizeMode = 0
	// MPSGraphResizeNearest - Samples the nearest neighbor to the pixel coordinate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeMode/nearest
	MPSGraphResizeNearest MPSGraphResizeMode = 0
)


// MPSGraphResizeNearestRoundingMode - The rounding mode to use when using nearest resize mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode
type MPSGraphResizeNearestRoundingMode uint

const (
	// MPSGraphResizeNearestRoundingModeCeil - Rounds values toward +inf.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode/ceil
	MPSGraphResizeNearestRoundingModeCeil MPSGraphResizeNearestRoundingMode = 0
	// MPSGraphResizeNearestRoundingModeFloor - Rounds values toward -inf.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode/floor
	MPSGraphResizeNearestRoundingModeFloor MPSGraphResizeNearestRoundingMode = 0
	// MPSGraphResizeNearestRoundingModeRoundPreferCeil - Rounds values to the nearest integer value, with 0.5f offset rounding toward +inf.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode/roundPreferCeil
	MPSGraphResizeNearestRoundingModeRoundPreferCeil MPSGraphResizeNearestRoundingMode = 0
	// MPSGraphResizeNearestRoundingModeRoundPreferFloor - Rounds values to the nearest integer value, with 0.5f rounding toward -inf.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode/roundPreferFloor
	MPSGraphResizeNearestRoundingModeRoundPreferFloor MPSGraphResizeNearestRoundingMode = 0
	// MPSGraphResizeNearestRoundingModeRoundToEven - Rounds values to the nearest integer value, with 0.5f rounding toward the closest even value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode/roundToEven
	MPSGraphResizeNearestRoundingModeRoundToEven MPSGraphResizeNearestRoundingMode = 0
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
	// MPSGraphScatterModeDiv - Divide
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/div
	MPSGraphScatterModeDiv MPSGraphScatterMode = 0
	// MPSGraphScatterModeMax - Maximum
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/max
	MPSGraphScatterModeMax MPSGraphScatterMode = 0
	// MPSGraphScatterModeMin - Minimum
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/min
	MPSGraphScatterModeMin MPSGraphScatterMode = 0
	// MPSGraphScatterModeMul - Multiply
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/mul
	MPSGraphScatterModeMul MPSGraphScatterMode = 0
	// MPSGraphScatterModeSet - Set
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/set
	MPSGraphScatterModeSet MPSGraphScatterMode = 0
	// MPSGraphScatterModeSub - Sub
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode/sub
	MPSGraphScatterModeSub MPSGraphScatterMode = 0
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
	// MPSGraphSparseStorageCSC - CSC Storage
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSparseStorageType/CSC
	MPSGraphSparseStorageCSC MPSGraphSparseStorageType = 0
	// MPSGraphSparseStorageCSR - CSR Storage
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSparseStorageType/CSR
	MPSGraphSparseStorageCSR MPSGraphSparseStorageType = 0
)


// MPSGraphTensorNamedDataLayout - The tensor layout.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout
type MPSGraphTensorNamedDataLayout uint

const (
	// MPSGraphTensorNamedDataLayoutCHW - LayoutCHW
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/CHW
	MPSGraphTensorNamedDataLayoutCHW MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutDHWIO - LayoutDHWIO
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/DHWIO
	MPSGraphTensorNamedDataLayoutDHWIO MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutHW - LayoutHW
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/HW
	MPSGraphTensorNamedDataLayoutHW MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutHWC - LayoutHWC
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/HWC
	MPSGraphTensorNamedDataLayoutHWC MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutHWIO - LayoutHWIO
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/HWIO
	MPSGraphTensorNamedDataLayoutHWIO MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutNCDHW - LayoutNCDHW
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/NCDHW
	MPSGraphTensorNamedDataLayoutNCDHW MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutNCHW - LayoutNCHW
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/NCHW
	MPSGraphTensorNamedDataLayoutNCHW MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutNDHWC - LayoutNDHWC
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/NDHWC
	MPSGraphTensorNamedDataLayoutNDHWC MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutNHWC - LayoutNHWC
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/NHWC
	MPSGraphTensorNamedDataLayoutNHWC MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutOIDHW - LayoutOIDHW
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/OIDHW
	MPSGraphTensorNamedDataLayoutOIDHW MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutOIHW - LayoutOIHW
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout/OIHW
	MPSGraphTensorNamedDataLayoutOIHW MPSGraphTensorNamedDataLayout = 0
)


