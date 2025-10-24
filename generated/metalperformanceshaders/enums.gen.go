// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders


// Enum types and constants

// MPSAccelerationStructureStatus - Constants that indicate an acceleration structure build state.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureStatus
type MPSAccelerationStructureStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureStatus/built
	MPSAccelerationStructureStatusBuilt MPSAccelerationStructureStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructurestatus/mpsaccelerationstructurestatusunbuilt
	MPSAccelerationStructureStatusUnbuilt MPSAccelerationStructureStatus = 0
)


// MPSAccelerationStructureUsage - Options that describe how an acceleration structure will be used.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureUsage
type MPSAccelerationStructureUsage uint

const (
	// frequentRebuild - Option indicating that the acceleration structure will be rebuilt frequently.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructureusage/2980778-frequentrebuild
	frequentRebuild MPSAccelerationStructureUsage = 0
	// refit - Option that enables support for refitting the acceleration structure after it has been built.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructureusage/2980780-refit
	refit MPSAccelerationStructureUsage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructureusage/3152575-prefercpubuild
	preferCPUBuild MPSAccelerationStructureUsage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructureusage/3152576-prefergpubuild
	preferGPUBuild MPSAccelerationStructureUsage = 0
	// MPSAccelerationStructureUsageFrequentRebuild - Option indicating that the acceleration structure will be rebuilt frequently.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureUsage/frequentRebuild
	MPSAccelerationStructureUsageFrequentRebuild MPSAccelerationStructureUsage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureUsage/MPSAccelerationStructureUsageNone
	MPSAccelerationStructureUsageNone MPSAccelerationStructureUsage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructureusage/mpsaccelerationstructureusageprefercpubuild
	MPSAccelerationStructureUsagePreferCPUBuild MPSAccelerationStructureUsage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructureusage/mpsaccelerationstructureusageprefergpubuild
	MPSAccelerationStructureUsagePreferGPUBuild MPSAccelerationStructureUsage = 0
	// MPSAccelerationStructureUsageRefit - Option that enables support for refitting the acceleration structure after it has been built.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructureusage/mpsaccelerationstructureusagerefit
	MPSAccelerationStructureUsageRefit MPSAccelerationStructureUsage = 0
)


// MPSAliasingStrategy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAliasingStrategy
type MPSAliasingStrategy uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy/3114017-aliasingreserved
	aliasingReserved MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy/3114018-default
	`default` MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy/3114020-prefernontemporarymemory
	preferNonTemporaryMemory MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy/3114021-prefertemporarymemory
	preferTemporaryMemory MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy/3114022-shallalias
	shallAlias MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy/3114023-shallnotalias
	shallNotAlias MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAliasingStrategy/aliasingReserved
	MPSAliasingStrategyAliasingReserved MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAliasingStrategy/default
	MPSAliasingStrategyDefault MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAliasingStrategy/MPSAliasingStrategyDontCare
	MPSAliasingStrategyDontCare MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy/mpsaliasingstrategyprefernontemporarymemory
	MPSAliasingStrategyPreferNonTemporaryMemory MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy/mpsaliasingstrategyprefertemporarymemory
	MPSAliasingStrategyPreferTemporaryMemory MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy/mpsaliasingstrategyshallalias
	MPSAliasingStrategyShallAlias MPSAliasingStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy/mpsaliasingstrategyshallnotalias
	MPSAliasingStrategyShallNotAlias MPSAliasingStrategy = 0
)


// MPSAlphaType - Premultiplication description for the color channels of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAlphaType
type MPSAlphaType uint

const (
	// MPSAlphaTypeAlphaIsOne - Alpha is guaranteed to be 1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAlphaType/alphaIsOne
	MPSAlphaTypeAlphaIsOne MPSAlphaType = 0
	// MPSAlphaTypeNonPremultiplied - The image is not premultiplied by alpha. Alpha is not guaranteed to be 1. ( )
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsalphatype/mpsalphatypenonpremultiplied
	MPSAlphaTypeNonPremultiplied MPSAlphaType = 0
	// MPSAlphaTypePremultiplied - The image is premultiplied by alpha. Alpha is not guaranteed to be 1. ( )
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsalphatype/mpsalphatypepremultiplied
	MPSAlphaTypePremultiplied MPSAlphaType = 0
)


// MPSBoundingBoxIntersectionTestType - Options for the intersection test type for a ray intersector bounding box.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBoundingBoxIntersectionTestType
type MPSBoundingBoxIntersectionTestType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBoundingBoxIntersectionTestType/axisAligned
	MPSBoundingBoxIntersectionTestTypeAxisAligned MPSBoundingBoxIntersectionTestType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBoundingBoxIntersectionTestType/default
	MPSBoundingBoxIntersectionTestTypeDefault MPSBoundingBoxIntersectionTestType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBoundingBoxIntersectionTestType/fast
	MPSBoundingBoxIntersectionTestTypeFast MPSBoundingBoxIntersectionTestType = 0
)


// MPSCNNBatchNormalizationFlags - Options that define how statistics are calculated during batch normalization.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationFlags
type MPSCNNBatchNormalizationFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationflags/2953945-calculatestatisticsalways
	calculateStatisticsAlways MPSCNNBatchNormalizationFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationflags/2953947-calculatestatisticsautomatic
	CalculateStatisticsAutomatic MPSCNNBatchNormalizationFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationflags/2953948-calculatestatisticsmask
	calculateStatisticsMask MPSCNNBatchNormalizationFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationflags/2953949-calculatestatisticsnever
	calculateStatisticsNever MPSCNNBatchNormalizationFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationflags/2953950-default
	Default MPSCNNBatchNormalizationFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationFlags/CalculateStatisticsAutomatic
	MPSCNNBatchNormalizationFlagsCalculateStatisticsAutomatic MPSCNNBatchNormalizationFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationFlags/Default
	MPSCNNBatchNormalizationFlagsDefault MPSCNNBatchNormalizationFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationFlags/calculateStatisticsAlways
	MPSCNNBatchNormalizationFlagsCalculateStatisticsAlways MPSCNNBatchNormalizationFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationFlags/calculateStatisticsMask
	MPSCNNBatchNormalizationFlagsCalculateStatisticsMask MPSCNNBatchNormalizationFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationFlags/calculateStatisticsNever
	MPSCNNBatchNormalizationFlagsCalculateStatisticsNever MPSCNNBatchNormalizationFlags = 0
)


// MPSCNNBinaryConvolutionFlags - Options used to control binary convolution kernels.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionFlags
type MPSCNNBinaryConvolutionFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionflags/mpscnnbinaryconvolutionflagsnone
	MPSCNNBinaryConvolutionFlagsNone MPSCNNBinaryConvolutionFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionflags/mpscnnbinaryconvolutionflagsusebetascaling
	MPSCNNBinaryConvolutionFlagsUseBetaScaling MPSCNNBinaryConvolutionFlags = 0
)


// MPSCNNBinaryConvolutionType - Options that defines what operations are used to perform binary convolution.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionType
type MPSCNNBinaryConvolutionType uint

const (
	// MPSCNNBinaryConvolutionTypeAND - A convolution type that uses input image binarization and the AND-operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionType/AND
	MPSCNNBinaryConvolutionTypeAND MPSCNNBinaryConvolutionType = 0
	// MPSCNNBinaryConvolutionTypeBinaryWeights - A convolution type that operates as a normal convolution, except that the weights are binary values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionType/binaryWeights
	MPSCNNBinaryConvolutionTypeBinaryWeights MPSCNNBinaryConvolutionType = 0
	// MPSCNNBinaryConvolutionTypeXNOR - A convolution type that uses input image binarization and the XNOR-operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutiontype/mpscnnbinaryconvolutiontypexnor
	MPSCNNBinaryConvolutionTypeXNOR MPSCNNBinaryConvolutionType = 0
)


// MPSCNNConvolutionFlags - Options used to control how kernel weights are stored and used in the CNN kernels
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionFlags
type MPSCNNConvolutionFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionFlags/none
	MPSCNNConvolutionFlagsNone MPSCNNConvolutionFlags = 0
)


// MPSCNNConvolutionGradientOption - Options that control which gradient to compute during backward propagation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientOption
type MPSCNNConvolutionGradientOption uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientoption/2942422-gradientwithweightsandbias
	gradientWithWeightsAndBias MPSCNNConvolutionGradientOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientoption/2942426-gradientwithdata
	gradientWithData MPSCNNConvolutionGradientOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientoption/2942431-all
	all MPSCNNConvolutionGradientOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientOption/all
	MPSCNNConvolutionGradientOptionAll MPSCNNConvolutionGradientOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientOption/gradientWithData
	MPSCNNConvolutionGradientOptionGradientWithData MPSCNNConvolutionGradientOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientOption/gradientWithWeightsAndBias
	MPSCNNConvolutionGradientOptionGradientWithWeightsAndBias MPSCNNConvolutionGradientOption = 0
)


// MPSCNNConvolutionWeightsLayout enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsLayout
type MPSCNNConvolutionWeightsLayout uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightslayout/mpscnnconvolutionweightslayoutohwi
	MPSCNNConvolutionWeightsLayoutOHWI MPSCNNConvolutionWeightsLayout = 0
)


// MPSCNNLossType - Constants that indicate supported loss filter types.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType
type MPSCNNLossType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType/categoricalCrossEntropy
	MPSCNNLossTypeCategoricalCrossEntropy MPSCNNLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType/cosineDistance
	MPSCNNLossTypeCosineDistance MPSCNNLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType/count
	MPSCNNLossTypeCount MPSCNNLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType/hinge
	MPSCNNLossTypeHinge MPSCNNLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType/huber
	MPSCNNLossTypeHuber MPSCNNLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType/kullbackLeiblerDivergence
	MPSCNNLossTypeKullbackLeiblerDivergence MPSCNNLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType/log
	MPSCNNLossTypeLog MPSCNNLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType/meanAbsoluteError
	MPSCNNLossTypeMeanAbsoluteError MPSCNNLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType/meanSquaredError
	MPSCNNLossTypeMeanSquaredError MPSCNNLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosstype/mpscnnlosstypesigmoidcrossentropy
	MPSCNNLossTypeSigmoidCrossEntropy MPSCNNLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosstype/mpscnnlosstypesoftmaxcrossentropy
	MPSCNNLossTypeSoftMaxCrossEntropy MPSCNNLossType = 0
)


// MPSCNNNeuronType - The types of neuron filter to append to a convolution.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronType
type MPSCNNNeuronType uint

const (
	// MPSCNNNeuronTypeAbsolute - A neuron type indicating an absolute neuron filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronType/absolute
	MPSCNNNeuronTypeAbsolute MPSCNNNeuronType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronType/count
	MPSCNNNeuronTypeCount MPSCNNNeuronType = 0
	// MPSCNNNeuronTypeELU - A neuron type indicating a parametric exponential linear unit neuron filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronType/ELU
	MPSCNNNeuronTypeELU MPSCNNNeuronType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronType/exponential
	MPSCNNNeuronTypeExponential MPSCNNNeuronType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronType/geLU
	MPSCNNNeuronTypeGeLU MPSCNNNeuronType = 0
	// MPSCNNNeuronTypeHardSigmoid - A neuron type indicating a hard sigmoid neuron filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronType/hardSigmoid
	MPSCNNNeuronTypeHardSigmoid MPSCNNNeuronType = 0
	// MPSCNNNeuronTypeLinear - A neuron type indicating a linear neuron filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronType/linear
	MPSCNNNeuronTypeLinear MPSCNNNeuronType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronType/logarithm
	MPSCNNNeuronTypeLogarithm MPSCNNNeuronType = 0
	// MPSCNNNeuronTypeNone - A neuron type indicating no neuron filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontype/mpscnnneurontypenone
	MPSCNNNeuronTypeNone MPSCNNNeuronType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontype/mpscnnneurontypepower
	MPSCNNNeuronTypePower MPSCNNNeuronType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontype/mpscnnneurontypeprelu
	MPSCNNNeuronTypePReLU MPSCNNNeuronType = 0
	// MPSCNNNeuronTypeReLU - A neuron type indicating a rectified linear unit neuron filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontype/mpscnnneurontyperelu
	MPSCNNNeuronTypeReLU MPSCNNNeuronType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontype/mpscnnneurontyperelun
	MPSCNNNeuronTypeReLUN MPSCNNNeuronType = 0
	// MPSCNNNeuronTypeSigmoid - A neuron type indicating a sigmoid neuron filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontype/mpscnnneurontypesigmoid
	MPSCNNNeuronTypeSigmoid MPSCNNNeuronType = 0
	// MPSCNNNeuronTypeSoftPlus - A neuron type indicating a parametric softplus neuron filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontype/mpscnnneurontypesoftplus
	MPSCNNNeuronTypeSoftPlus MPSCNNNeuronType = 0
	// MPSCNNNeuronTypeSoftSign - A neuron type indicating a softsign neuron filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontype/mpscnnneurontypesoftsign
	MPSCNNNeuronTypeSoftSign MPSCNNNeuronType = 0
	// MPSCNNNeuronTypeTanH - A neuron type indicating a hyperbolic tangent neuron filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontype/mpscnnneurontypetanh
	MPSCNNNeuronTypeTanH MPSCNNNeuronType = 0
)


// MPSCNNReductionType - Constants that indicate supported reduction types.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNReductionType
type MPSCNNReductionType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNReductionType/count
	MPSCNNReductionTypeCount MPSCNNReductionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNReductionType/mean
	MPSCNNReductionTypeMean MPSCNNReductionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnreductiontype/mpscnnreductiontypenone
	MPSCNNReductionTypeNone MPSCNNReductionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnreductiontype/mpscnnreductiontypesum
	MPSCNNReductionTypeSum MPSCNNReductionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnreductiontype/mpscnnreductiontypesumbynonzeroweights
	MPSCNNReductionTypeSumByNonZeroWeights MPSCNNReductionType = 0
)


// MPSCNNWeightsQuantizationType - Options that specify the type of quantization used to generate unsigned integer weights.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNWeightsQuantizationType
type MPSCNNWeightsQuantizationType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnweightsquantizationtype/mpscnnweightsquantizationtypelinear
	MPSCNNWeightsQuantizationTypeLinear MPSCNNWeightsQuantizationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnweightsquantizationtype/mpscnnweightsquantizationtypelookuptable
	MPSCNNWeightsQuantizationTypeLookupTable MPSCNNWeightsQuantizationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnweightsquantizationtype/mpscnnweightsquantizationtypenone
	MPSCNNWeightsQuantizationTypeNone MPSCNNWeightsQuantizationType = 0
)


// MPSDataLayout - Options that define how buffer data is arranged.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataLayout
type MPSDataLayout uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataLayout/featureChannelsxHeightxWidth
	MPSDataLayoutFeatureChannelsxHeightxWidth MPSDataLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataLayout/HeightxWidthxFeatureChannels
	MPSDataLayoutHeightxWidthxFeatureChannels MPSDataLayout = 0
)


// MPSDataType - A value to specify a type of data.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType
type MPSDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/2866103-intbit
	intBit MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/alternateEncodingBit
	MPSDataTypeAlternateEncodingBit MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/bFloat16
	MPSDataTypeBFloat16 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/bool
	MPSDataTypeBool MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/complexBit
	MPSDataTypeComplexBit MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/complexFloat16
	MPSDataTypeComplexFloat16 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/complexFloat32
	MPSDataTypeComplexFloat32 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/float16
	MPSDataTypeFloat16 MPSDataType = 0
	// MPSDataTypeFloat32 - A 32-bit floating point type (single precision).
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/float32
	MPSDataTypeFloat32 MPSDataType = 0
	// MPSDataTypeFloatBit - A common bit for all floating point data types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/floatBit
	MPSDataTypeFloatBit MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/int16
	MPSDataTypeInt16 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/int2
	MPSDataTypeInt2 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/int32
	MPSDataTypeInt32 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/int4
	MPSDataTypeInt4 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/int64
	MPSDataTypeInt64 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/int8
	MPSDataTypeInt8 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/intBit
	MPSDataTypeIntBit MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType/invalid
	MPSDataTypeInvalid MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/mpsdatatypenormalizedbit
	MPSDataTypeNormalizedBit MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/mpsdatatypesignedbit
	MPSDataTypeSignedBit MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/mpsdatatypeuint16
	MPSDataTypeUInt16 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/mpsdatatypeuint2
	MPSDataTypeUInt2 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/mpsdatatypeuint32
	MPSDataTypeUInt32 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/mpsdatatypeuint4
	MPSDataTypeUInt4 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/mpsdatatypeuint64
	MPSDataTypeUInt64 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/mpsdatatypeuint8
	MPSDataTypeUInt8 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/mpsdatatypeunorm1
	MPSDataTypeUnorm1 MPSDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype/mpsdatatypeunorm8
	MPSDataTypeUnorm8 MPSDataType = 0
)


// MPSDeviceOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDeviceOptions
type MPSDeviceOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdeviceoptions/3088915-default
	Default MPSDeviceOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdeviceoptions/3088916-lowpower
	lowPower MPSDeviceOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdeviceoptions/3088917-skipremovable
	skipRemovable MPSDeviceOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDeviceOptions/Default
	MPSDeviceOptionsDefault MPSDeviceOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDeviceOptions/lowPower
	MPSDeviceOptionsLowPower MPSDeviceOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdeviceoptions/mpsdeviceoptionsskipremovable
	MPSDeviceOptionsSkipRemovable MPSDeviceOptions = 0
)


// MPSFloatDataTypeBit enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFloatDataTypeBit
type MPSFloatDataTypeBit uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFloatDataTypeBit/exponentBit
	MPSFloatDataTypeExponentBit MPSFloatDataTypeBit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFloatDataTypeBit/mantissaBit
	MPSFloatDataTypeMantissaBit MPSFloatDataTypeBit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsfloatdatatypebit/mpsfloatdatatypesignbit
	MPSFloatDataTypeSignBit MPSFloatDataTypeBit = 0
)


// MPSFloatDataTypeShift enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFloatDataTypeShift
type MPSFloatDataTypeShift uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFloatDataTypeShift/exponentShift
	MPSFloatDataTypeExponentShift MPSFloatDataTypeShift = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFloatDataTypeShift/mantissaShift
	MPSFloatDataTypeMantissaShift MPSFloatDataTypeShift = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsfloatdatatypeshift/mpsfloatdatatypesignshift
	MPSFloatDataTypeSignShift MPSFloatDataTypeShift = 0
)


// MPSImageEdgeMode - The options used to control the edge behavior of an image filter when it reads outside the bounds of a source texture.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEdgeMode
type MPSImageEdgeMode uint

const (
	// MPSImageEdgeModeClamp - Out-of-bound pixels are clamped to the nearest edge pixel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEdgeMode/clamp
	MPSImageEdgeModeClamp MPSImageEdgeMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEdgeMode/constant
	MPSImageEdgeModeConstant MPSImageEdgeMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEdgeMode/mirror
	MPSImageEdgeModeMirror MPSImageEdgeMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEdgeMode/mirrorWithEdge
	MPSImageEdgeModeMirrorWithEdge MPSImageEdgeMode = 0
	// MPSImageEdgeModeZero - Out-of-bound pixels are set to   for images without an alpha channel or   for images with an alpha channel, as defined by their pixel format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedgemode/mpsimageedgemodezero
	MPSImageEdgeModeZero MPSImageEdgeMode = 0
)


// MPSImageFeatureChannelFormat - Encodes the representation of a single channel within an image.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFeatureChannelFormat
type MPSImageFeatureChannelFormat uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFeatureChannelFormat/count
	MPSImageFeatureChannelFormatCount MPSImageFeatureChannelFormat = 0
	// MPSImageFeatureChannelFormatFloat16 - IEEE-754 16-bit floating-point type (half precision). Representable normal range is  . 11 bits of precision + exponent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFeatureChannelFormat/float16
	MPSImageFeatureChannelFormatFloat16 MPSImageFeatureChannelFormat = 0
	// MPSImageFeatureChannelFormatFloat32 - IEEE-754 32-bit floating-point type (single precision, standard   type in C). 24 bits of precision + exponent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFeatureChannelFormat/float32
	MPSImageFeatureChannelFormatFloat32 MPSImageFeatureChannelFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFeatureChannelFormat/MPSImageFeatureChannelFormat_reserved0
	MPSImageFeatureChannelFormat_reserved0 MPSImageFeatureChannelFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefeaturechannelformat/mpsimagefeaturechannelformatnone
	MPSImageFeatureChannelFormatNone MPSImageFeatureChannelFormat = 0
	// MPSImageFeatureChannelFormatUnorm16 -  type with value   and encoding  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefeaturechannelformat/mpsimagefeaturechannelformatunorm16
	MPSImageFeatureChannelFormatUnorm16 MPSImageFeatureChannelFormat = 0
	// MPSImageFeatureChannelFormatUnorm8 -  type with value   and encoding  . 
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefeaturechannelformat/mpsimagefeaturechannelformatunorm8
	MPSImageFeatureChannelFormatUnorm8 MPSImageFeatureChannelFormat = 0
)


// MPSIntersectionDataType - Options that determine the data contained in an intersection result.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType
type MPSIntersectionDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType/distance
	MPSIntersectionDataTypeDistance MPSIntersectionDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType/distancePrimitiveIndex
	MPSIntersectionDataTypeDistancePrimitiveIndex MPSIntersectionDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType/distancePrimitiveIndexBufferIndex
	MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndex MPSIntersectionDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType/distancePrimitiveIndexBufferIndexCoordinates
	MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexCoordinates MPSIntersectionDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType/distancePrimitiveIndexBufferIndexInstanceIndex
	MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexInstanceIndex MPSIntersectionDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType/distancePrimitiveIndexBufferIndexInstanceIndexCoordinates
	MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates MPSIntersectionDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType/distancePrimitiveIndexCoordinates
	MPSIntersectionDataTypeDistancePrimitiveIndexCoordinates MPSIntersectionDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType/distancePrimitiveIndexInstanceIndex
	MPSIntersectionDataTypeDistancePrimitiveIndexInstanceIndex MPSIntersectionDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType/distancePrimitiveIndexInstanceIndexCoordinates
	MPSIntersectionDataTypeDistancePrimitiveIndexInstanceIndexCoordinates MPSIntersectionDataType = 0
)


// MPSIntersectionType - Options that determine an intersection type for a ray intersector.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionType
type MPSIntersectionType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionType/any
	MPSIntersectionTypeAny MPSIntersectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiontype/mpsintersectiontypenearest
	MPSIntersectionTypeNearest MPSIntersectionType = 0
)


// MPSKernelOptions - The options used when creating a kernel.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernelOptions
type MPSKernelOptions uint

const (
	// allowReducedPrecision - When possible, kernels use a higher-precision data representation internally than the destination storage format to avoid excessive accumulation of computational rounding error in the result. This option advises the kernel that the destination storage format already has too much precision for what is ultimately required downstream, and the kernel may use reduced precision internally when it determines that a less precise result would yield better performance. When enabled, the performance win is often small and the precision of the result may vary by hardware and OS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskerneloptions/1618748-allowreducedprecision
	allowReducedPrecision MPSKernelOptions = 0
	// none - The default option for the kernel. Kernels created with this option will not skip any API validation and will not use reduced precision.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskerneloptions/1618816-none
	none MPSKernelOptions = 0
	// skipAPIValidation - A property that directs the kernel to perform or skip argument validation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskerneloptions/1618826-skipapivalidation
	skipAPIValidation MPSKernelOptions = 0
	// insertDebugGroups - Enabling this option will cause various kernel   methods to call the   and   methods. The debug string will be drawn from the kernel’s   property, if available, or the name of the class otherwise.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskerneloptions/1648897-insertdebuggroups
	insertDebugGroups MPSKernelOptions = 0
	// disableInternalTiling - Some kernels may automatically split up their work internally into multiple tiles. This improves performance on larger textures and reduces the amount of memory needed by the framework for temporary storage. However, if you are using your own tiling scheme to achieve similar results, your tile sizes and the framework’s choice of tile sizes may interfere with one another, causing the framework to subdivide your tiles for its own use inefficiently. Use this option to force the framework to process your data tile as a single chunk.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskerneloptions/1648950-disableinternaltiling
	disableInternalTiling MPSKernelOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskerneloptions/2889867-verbose
	verbose MPSKernelOptions = 0
	// MPSKernelOptionsAllowReducedPrecision - When possible, kernels use a higher-precision data representation internally than the destination storage format to avoid excessive accumulation of computational rounding error in the result. This option advises the kernel that the destination storage format already has too much precision for what is ultimately required downstream, and the kernel may use reduced precision internally when it determines that a less precise result would yield better performance. When enabled, the performance win is often small and the precision of the result may vary by hardware and OS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernelOptions/allowReducedPrecision
	MPSKernelOptionsAllowReducedPrecision MPSKernelOptions = 0
	// MPSKernelOptionsDisableInternalTiling - Some kernels may automatically split up their work internally into multiple tiles. This improves performance on larger textures and reduces the amount of memory needed by the framework for temporary storage. However, if you are using your own tiling scheme to achieve similar results, your tile sizes and the framework’s choice of tile sizes may interfere with one another, causing the framework to subdivide your tiles for its own use inefficiently. Use this option to force the framework to process your data tile as a single chunk.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernelOptions/disableInternalTiling
	MPSKernelOptionsDisableInternalTiling MPSKernelOptions = 0
	// MPSKernelOptionsInsertDebugGroups - Enables calling kernel encode methods.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernelOptions/insertDebugGroups
	MPSKernelOptionsInsertDebugGroups MPSKernelOptions = 0
	// MPSKernelOptionsNone - The default option for the kernel. Kernels created with this option will not skip any API validation and will not use reduced precision.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskerneloptions/mpskerneloptionsnone
	MPSKernelOptionsNone MPSKernelOptions = 0
	// MPSKernelOptionsSkipAPIValidation - A property that directs the kernel to perform or skip argument validation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskerneloptions/mpskerneloptionsskipapivalidation
	MPSKernelOptionsSkipAPIValidation MPSKernelOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskerneloptions/mpskerneloptionsverbose
	MPSKernelOptionsVerbose MPSKernelOptions = 0
)


// MPSMatrixDecompositionStatus enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionStatus
type MPSMatrixDecompositionStatus uint

const (
	// MPSMatrixDecompositionStatusFailure - A status indicating the decomposition was not able to be completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionStatus/failure
	MPSMatrixDecompositionStatusFailure MPSMatrixDecompositionStatus = 0
	// MPSMatrixDecompositionStatusNonPositiveDefinite - A status indicating a non-positive-definite pivot value was calculated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionstatus/mpsmatrixdecompositionstatusnonpositivedefinite
	MPSMatrixDecompositionStatusNonPositiveDefinite MPSMatrixDecompositionStatus = 0
	// MPSMatrixDecompositionStatusSingular - A status indicating the resulting decomposition is not suitable for use in a subsequent system solve.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionstatus/mpsmatrixdecompositionstatussingular
	MPSMatrixDecompositionStatusSingular MPSMatrixDecompositionStatus = 0
	// MPSMatrixDecompositionStatusSuccess - A status indicating the decomposition was performed successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionstatus/mpsmatrixdecompositionstatussuccess
	MPSMatrixDecompositionStatusSuccess MPSMatrixDecompositionStatus = 0
)


// MPSMatrixRandomDistribution enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistribution
type MPSMatrixRandomDistribution uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistribution/3242853-default
	`default` MPSMatrixRandomDistribution = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistribution/3242854-uniform
	uniform MPSMatrixRandomDistribution = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistribution/3547978-normal
	normal MPSMatrixRandomDistribution = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistribution/default
	MPSMatrixRandomDistributionDefault MPSMatrixRandomDistribution = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistribution/mpsmatrixrandomdistributionnormal
	MPSMatrixRandomDistributionNormal MPSMatrixRandomDistribution = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistribution/mpsmatrixrandomdistributionuniform
	MPSMatrixRandomDistributionUniform MPSMatrixRandomDistribution = 0
)


// MPSNDArrayQuantizationScheme enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizationScheme
type MPSNDArrayQuantizationScheme uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationscheme/4446145-typeaffine
	typeAffine MPSNDArrayQuantizationScheme = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationscheme/4446146-typelut
	typeLUT MPSNDArrayQuantizationScheme = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationscheme/4446147-none
	none MPSNDArrayQuantizationScheme = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationscheme/mpsndarrayquantizationtypeaffine
	MPSNDArrayQuantizationTypeAffine MPSNDArrayQuantizationScheme = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationscheme/mpsndarrayquantizationtypelut
	MPSNDArrayQuantizationTypeLUT MPSNDArrayQuantizationScheme = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationscheme/mpsndarrayquantizationtypenone
	MPSNDArrayQuantizationTypeNone MPSNDArrayQuantizationScheme = 0
)


// MPSNNComparisonType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonType
type MPSNNComparisonType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisontype/3037378-equal
	equal MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisontype/3037379-greater
	greater MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisontype/3037380-greaterorequal
	greaterOrEqual MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisontype/3037381-less
	less MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisontype/3037382-lessorequal
	lessOrEqual MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisontype/3037383-notequal
	notEqual MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonType/equal
	MPSNNComparisonTypeEqual MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonType/greater
	MPSNNComparisonTypeGreater MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonType/greaterOrEqual
	MPSNNComparisonTypeGreaterOrEqual MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonType/less
	MPSNNComparisonTypeLess MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonType/lessOrEqual
	MPSNNComparisonTypeLessOrEqual MPSNNComparisonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisontype/mpsnncomparisontypenotequal
	MPSNNComparisonTypeNotEqual MPSNNComparisonType = 0
)


// MPSNNConvolutionAccumulatorPrecisionOption - Options that specify convolution accumulator precision.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConvolutionAccumulatorPrecisionOption
type MPSNNConvolutionAccumulatorPrecisionOption uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconvolutionaccumulatorprecisionoption/2942457-float
	float MPSNNConvolutionAccumulatorPrecisionOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconvolutionaccumulatorprecisionoption/2942458-half
	half MPSNNConvolutionAccumulatorPrecisionOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConvolutionAccumulatorPrecisionOption/float
	MPSNNConvolutionAccumulatorPrecisionOptionFloat MPSNNConvolutionAccumulatorPrecisionOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConvolutionAccumulatorPrecisionOption/half
	MPSNNConvolutionAccumulatorPrecisionOptionHalf MPSNNConvolutionAccumulatorPrecisionOption = 0
)


// MPSNNPaddingMethod - Options that define a graph’s padding.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod
type MPSNNPaddingMethod uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2866951-addremaindertobottomright
	addRemainderToBottomRight MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2866952-custom
	custom MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867003-size_reserved
	size_reserved MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867015-alignmask
	alignMask MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867061-addremaindertobottomleft
	addRemainderToBottomLeft MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867062-alignbottomright
	alignBottomRight MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867069-sizesame
	sizeSame MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867095-sizemask
	sizeMask MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867100-sizefull
	sizeFull MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867108-addremaindertomask
	addRemainderToMask MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867109-aligntopleft
	alignTopLeft MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867116-align_reserved
	align_reserved MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867161-addremaindertotopright
	addRemainderToTopRight MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867573-centered
	centered MPSNNPaddingMethod = 0
	// validOnly - A padding method where result values are only produced for the area that is guaranteed to have all of its input values defined   
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867574-validonly
	validOnly MPSNNPaddingMethod = 0
	// topLeft - A padding method where leftover padding is added to the top or left side of image as appropriate.	  
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2867575-topleft
	topLeft MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/2890785-excludeedges
	excludeEdges MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/3020690-customwhitelistfornodefusion
	customWhitelistForNodeFusion MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/3763056-customallowfornodefusion
	customAllowForNodeFusion MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/addRemainderToBottomLeft
	MPSNNPaddingMethodAddRemainderToBottomLeft MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/addRemainderToBottomRight
	MPSNNPaddingMethodAddRemainderToBottomRight MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/addRemainderToMask
	MPSNNPaddingMethodAddRemainderToMask MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/addRemainderToTopRight
	MPSNNPaddingMethodAddRemainderToTopRight MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/alignBottomRight
	MPSNNPaddingMethodAlignBottomRight MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/alignMask
	MPSNNPaddingMethodAlignMask MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/alignTopLeft
	MPSNNPaddingMethodAlignTopLeft MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/align_reserved
	MPSNNPaddingMethodAlign_reserved MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/centered
	MPSNNPaddingMethodAlignCentered MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/custom
	MPSNNPaddingMethodCustom MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/customAllowForNodeFusion
	MPSNNPaddingMethodCustomAllowForNodeFusion MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/customWhitelistForNodeFusion
	MPSNNPaddingMethodCustomWhitelistForNodeFusion MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod/excludeEdges
	MPSNNPaddingMethodExcludeEdges MPSNNPaddingMethod = 0
	// MPSNNPaddingMethodAddRemainderToTopLeft - A padding method where leftover padding is added to the top or left side of image as appropriate.	  
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/mpsnnpaddingmethodaddremaindertotopleft
	MPSNNPaddingMethodAddRemainderToTopLeft MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/mpsnnpaddingmethodsize_reserved
	MPSNNPaddingMethodSize_reserved MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/mpsnnpaddingmethodsizefull
	MPSNNPaddingMethodSizeFull MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/mpsnnpaddingmethodsizemask
	MPSNNPaddingMethodSizeMask MPSNNPaddingMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/mpsnnpaddingmethodsizesame
	MPSNNPaddingMethodSizeSame MPSNNPaddingMethod = 0
	// MPSNNPaddingMethodSizeValidOnly - A padding method where result values are only produced for the area that is guaranteed to have all of its input values defined   
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod/mpsnnpaddingmethodsizevalidonly
	MPSNNPaddingMethodSizeValidOnly MPSNNPaddingMethod = 0
)


// MPSNNRegularizationType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNRegularizationType
type MPSNNRegularizationType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNRegularizationType/L1
	MPSNNRegularizationTypeL1 MPSNNRegularizationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNRegularizationType/L2
	MPSNNRegularizationTypeL2 MPSNNRegularizationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnregularizationtype/mpsnnregularizationtypenone
	MPSNNRegularizationTypeNone MPSNNRegularizationType = 0
)


// MPSNNTrainingStyle - Options that control how graph nodes are trained.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainingStyle
type MPSNNTrainingStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnntrainingstyle/2952961-updatedevicecpu
	updateDeviceCPU MPSNNTrainingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnntrainingstyle/2952962-updatedevicenone
	UpdateDeviceNone MPSNNTrainingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnntrainingstyle/2952963-updatedevicegpu
	updateDeviceGPU MPSNNTrainingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainingStyle/UpdateDeviceNone
	MPSNNTrainingStyleUpdateDeviceNone MPSNNTrainingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnntrainingstyle/mpsnntrainingstyleupdatedevicecpu
	MPSNNTrainingStyleUpdateDeviceCPU MPSNNTrainingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnntrainingstyle/mpsnntrainingstyleupdatedevicegpu
	MPSNNTrainingStyleUpdateDeviceGPU MPSNNTrainingStyle = 0
)


// MPSPolygonType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonType
type MPSPolygonType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygontype/mpspolygontypequadrilateral
	MPSPolygonTypeQuadrilateral MPSPolygonType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygontype/mpspolygontypetriangle
	MPSPolygonTypeTriangle MPSPolygonType = 0
)


// MPSPurgeableState - The purgeable state of an image’s underlying texture.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPurgeableState
type MPSPurgeableState uint

const (
	// MPSPurgeableStateAllocationDeferred - The image’s underlying texture hasn’t been allocated yet. Attempts to set another purgeable state using the   method will be ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPurgeableState/allocationDeferred
	MPSPurgeableStateAllocationDeferred MPSPurgeableState = 0
	// MPSPurgeableStateEmpty - The contents of the resource are or will be discarded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPurgeableState/empty
	MPSPurgeableStateEmpty MPSPurgeableState = 0
	// MPSPurgeableStateKeepCurrent - The current state is queried but doesn’t change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPurgeableState/keepCurrent
	MPSPurgeableStateKeepCurrent MPSPurgeableState = 0
	// MPSPurgeableStateNonVolatile - Equivalent to the   value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspurgeablestate/mpspurgeablestatenonvolatile
	MPSPurgeableStateNonVolatile MPSPurgeableState = 0
	// MPSPurgeableStateVolatile - Equivalent to the   value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspurgeablestate/mpspurgeablestatevolatile
	MPSPurgeableStateVolatile MPSPurgeableState = 0
)


// MPSRayDataType - Options for the data type for an intersector ray.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayDataType
type MPSRayDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraydatatype/mpsraydatatypeorigindirection
	MPSRayDataTypeOriginDirection MPSRayDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraydatatype/mpsraydatatypeoriginmaskdirectionmaxdistance
	MPSRayDataTypeOriginMaskDirectionMaxDistance MPSRayDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraydatatype/mpsraydatatypeoriginmindistancedirectionmaxdistance
	MPSRayDataTypeOriginMinDistanceDirectionMaxDistance MPSRayDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraydatatype/mpsraydatatypepackedorigindirection
	MPSRayDataTypePackedOriginDirection MPSRayDataType = 0
)


// MPSRayMaskOperator enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOperator
type MPSRayMaskOperator uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOperator/and
	MPSRayMaskOperatorAnd MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOperator/equal
	MPSRayMaskOperatorEqual MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOperator/greaterThan
	MPSRayMaskOperatorGreaterThan MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOperator/greaterThanOrEqualTo
	MPSRayMaskOperatorGreaterThanOrEqualTo MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOperator/lessThan
	MPSRayMaskOperatorLessThan MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOperator/lessThanOrEqualTo
	MPSRayMaskOperatorLessThanOrEqualTo MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoperator/mpsraymaskoperatornotand
	MPSRayMaskOperatorNotAnd MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoperator/mpsraymaskoperatornotequal
	MPSRayMaskOperatorNotEqual MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoperator/mpsraymaskoperatornotor
	MPSRayMaskOperatorNotOr MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoperator/mpsraymaskoperatornotxor
	MPSRayMaskOperatorNotXor MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoperator/mpsraymaskoperatoror
	MPSRayMaskOperatorOr MPSRayMaskOperator = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoperator/mpsraymaskoperatorxor
	MPSRayMaskOperatorXor MPSRayMaskOperator = 0
)


// MPSRayMaskOptions - Options for ray intersector mask options.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOptions
type MPSRayMaskOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoptions/2980817-instance
	instance MPSRayMaskOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoptions/2980819-primitive
	primitive MPSRayMaskOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOptions/instance
	MPSRayMaskOptionInstance MPSRayMaskOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOptions/MPSRayMaskOptionNone
	MPSRayMaskOptionNone MPSRayMaskOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoptions/mpsraymaskoptionprimitive
	MPSRayMaskOptionPrimitive MPSRayMaskOptions = 0
)


// MPSRNNBidirectionalCombineMode - Modes that define how two images or matrices are combined.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNBidirectionalCombineMode
type MPSRNNBidirectionalCombineMode uint

const (
	// MPSRNNBidirectionalCombineModeAdd - A mode in which two sequences are summed to form a single output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNBidirectionalCombineMode/add
	MPSRNNBidirectionalCombineModeAdd MPSRNNBidirectionalCombineMode = 0
	// MPSRNNBidirectionalCombineModeConcatenate - A mode in which two sequences are concatenated along the feature channels to form a single output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNBidirectionalCombineMode/concatenate
	MPSRNNBidirectionalCombineModeConcatenate MPSRNNBidirectionalCombineMode = 0
	// MPSRNNBidirectionalCombineModeNone - A mode in which two sequences are kept separate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnbidirectionalcombinemode/mpsrnnbidirectionalcombinemodenone
	MPSRNNBidirectionalCombineModeNone MPSRNNBidirectionalCombineMode = 0
)


// MPSRNNMatrixId - Options that define which matrix is copied in or out of a trainable RNN layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId
type MPSRNNMatrixId uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/gruInputGateBiasTerms
	MPSRNNMatrixIdGRUInputGateBiasTerms MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/gruInputGateInputWeights
	MPSRNNMatrixIdGRUInputGateInputWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/gruInputGateRecurrentWeights
	MPSRNNMatrixIdGRUInputGateRecurrentWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/gruOutputGateBiasTerms
	MPSRNNMatrixIdGRUOutputGateBiasTerms MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/gruOutputGateInputGateWeights
	MPSRNNMatrixIdGRUOutputGateInputGateWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/gruOutputGateInputWeights
	MPSRNNMatrixIdGRUOutputGateInputWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/gruOutputGateRecurrentWeights
	MPSRNNMatrixIdGRUOutputGateRecurrentWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/gruRecurrentGateBiasTerms
	MPSRNNMatrixIdGRURecurrentGateBiasTerms MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/gruRecurrentGateInputWeights
	MPSRNNMatrixIdGRURecurrentGateInputWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/gruRecurrentGateRecurrentWeights
	MPSRNNMatrixIdGRURecurrentGateRecurrentWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmForgetGateBiasTerms
	MPSRNNMatrixIdLSTMForgetGateBiasTerms MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmForgetGateInputWeights
	MPSRNNMatrixIdLSTMForgetGateInputWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmForgetGateMemoryWeights
	MPSRNNMatrixIdLSTMForgetGateMemoryWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmForgetGateRecurrentWeights
	MPSRNNMatrixIdLSTMForgetGateRecurrentWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmInputGateBiasTerms
	MPSRNNMatrixIdLSTMInputGateBiasTerms MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmInputGateInputWeights
	MPSRNNMatrixIdLSTMInputGateInputWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmInputGateMemoryWeights
	MPSRNNMatrixIdLSTMInputGateMemoryWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmInputGateRecurrentWeights
	MPSRNNMatrixIdLSTMInputGateRecurrentWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmMemoryGateBiasTerms
	MPSRNNMatrixIdLSTMMemoryGateBiasTerms MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmMemoryGateInputWeights
	MPSRNNMatrixIdLSTMMemoryGateInputWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmMemoryGateMemoryWeights
	MPSRNNMatrixIdLSTMMemoryGateMemoryWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmMemoryGateRecurrentWeights
	MPSRNNMatrixIdLSTMMemoryGateRecurrentWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmOutputGateBiasTerms
	MPSRNNMatrixIdLSTMOutputGateBiasTerms MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmOutputGateInputWeights
	MPSRNNMatrixIdLSTMOutputGateInputWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmOutputGateMemoryWeights
	MPSRNNMatrixIdLSTMOutputGateMemoryWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/lstmOutputGateRecurrentWeights
	MPSRNNMatrixIdLSTMOutputGateRecurrentWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId/MPSRNNMatrixId_count
	MPSRNNMatrixId_count MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixid/mpsrnnmatrixidsinglegatebiasterms
	MPSRNNMatrixIdSingleGateBiasTerms MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixid/mpsrnnmatrixidsinglegateinputweights
	MPSRNNMatrixIdSingleGateInputWeights MPSRNNMatrixId = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixid/mpsrnnmatrixidsinglegaterecurrentweights
	MPSRNNMatrixIdSingleGateRecurrentWeights MPSRNNMatrixId = 0
)


// MPSRNNSequenceDirection - Directions that a sequence of inputs can be processed by a recurrent neural network layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSequenceDirection
type MPSRNNSequenceDirection uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSequenceDirection/backward
	MPSRNNSequenceDirectionBackward MPSRNNSequenceDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSequenceDirection/forward
	MPSRNNSequenceDirectionForward MPSRNNSequenceDirection = 0
)


// MPSStateResourceType - Options for the underlying resource type for a state object.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceType
type MPSStateResourceType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceType/buffer
	MPSStateResourceTypeBuffer MPSStateResourceType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcetype/mpsstateresourcetypenone
	MPSStateResourceTypeNone MPSStateResourceType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcetype/mpsstateresourcetypetexture
	MPSStateResourceTypeTexture MPSStateResourceType = 0
)


// MPSTemporalWeighting enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalWeighting
type MPSTemporalWeighting uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalWeighting/average
	MPSTemporalWeightingAverage MPSTemporalWeighting = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalWeighting/exponentialMovingAverage
	MPSTemporalWeightingExponentialMovingAverage MPSTemporalWeighting = 0
)


// MPSTransformType - Constants that indicate instance transformation types.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTransformType
type MPSTransformType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTransformType/float4x4
	MPSTransformTypeFloat4x4 MPSTransformType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTransformType/identity
	MPSTransformTypeIdentity MPSTransformType = 0
)


// MPSTriangleIntersectionTestType - Options for the ray-triangle intersection test.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTriangleIntersectionTestType
type MPSTriangleIntersectionTestType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTriangleIntersectionTestType/default
	MPSTriangleIntersectionTestTypeDefault MPSTriangleIntersectionTestType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstriangleintersectiontesttype/mpstriangleintersectiontesttypewatertight
	MPSTriangleIntersectionTestTypeWatertight MPSTriangleIntersectionTestType = 0
)


