// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

/* debug [enums.gen.go]: Generating 19 enums for MLCompute */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MLCPaddingPolicy (3 cases) */
// MLCPaddingPolicy - A padding policy that you specify for a convolution or pooling layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingPolicy-14ba7
type MLCPaddingPolicy int

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingPolicy-14ba7/MLCPaddingPolicySame
	MLCPaddingPolicySame MLCPaddingPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingPolicy-14ba7/MLCPaddingPolicyUsePaddingSize
	MLCPaddingPolicyUsePaddingSize MLCPaddingPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingPolicy-14ba7/MLCPaddingPolicyValid
	MLCPaddingPolicyValid MLCPaddingPolicy = 0
)

/* debug [enums.gen.go]: Processing enum MLCActivationType (22 cases) */
// MLCActivationType - An activation type that you specify for an activation descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType
type MLCActivationType uint

const (
	// MLCActivationTypeAbsolute - An activation type that implements the absolute activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/absolute
	MLCActivationTypeAbsolute MLCActivationType = 0
	// MLCActivationTypeCELU - An activation type that implements the CELU activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/celu
	MLCActivationTypeCELU MLCActivationType = 0
	// MLCActivationTypeClamp - An activation type that implements the clamp activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/clamp
	MLCActivationTypeClamp MLCActivationType = 0
	// MLCActivationTypeELU - An activation type that implements the exponential linear unit activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/elu
	MLCActivationTypeELU MLCActivationType = 0
	// MLCActivationTypeGELU - An activation type that implements the gaussian error linear unit activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/gelu
	MLCActivationTypeGELU MLCActivationType = 0
	// MLCActivationTypeHardShrink - An activation type that implements the hard shrink activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/hardShrink
	MLCActivationTypeHardShrink MLCActivationType = 0
	// MLCActivationTypeHardSigmoid - An activation type that implements the hard sigmoid activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/hardSigmoid
	MLCActivationTypeHardSigmoid MLCActivationType = 0
	// MLCActivationTypeHardSwish - An activation type that implements the hard swish activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/hardSwish
	MLCActivationTypeHardSwish MLCActivationType = 0
	// MLCActivationTypeLinear - An activation type that implements the linear activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/linear
	MLCActivationTypeLinear MLCActivationType = 0
	// MLCActivationTypeLogSigmoid - An activation type that implements the log sigmoid activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/logSigmoid
	MLCActivationTypeLogSigmoid MLCActivationType = 0
	// MLCActivationTypeCount - The count of activation types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/MLCActivationTypeCount
	MLCActivationTypeCount MLCActivationType = 0
	// MLCActivationTypeNone - An activation type that implements the identity function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/none
	MLCActivationTypeNone MLCActivationType = 0
	// MLCActivationTypeReLU - An activation type that implements the rectified linear unit activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/relu
	MLCActivationTypeReLU MLCActivationType = 0
	// MLCActivationTypeReLUN - An activation type that implements the ReLUN activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/relun
	MLCActivationTypeReLUN MLCActivationType = 0
	// MLCActivationTypeSELU - An activation type that implements the scaled exponential linear unit activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/selu
	MLCActivationTypeSELU MLCActivationType = 0
	// MLCActivationTypeSigmoid - An activation type that implements the sigmoid activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/sigmoid
	MLCActivationTypeSigmoid MLCActivationType = 0
	// MLCActivationTypeSoftPlus - An activation type that implements the soft plus activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/softPlus
	MLCActivationTypeSoftPlus MLCActivationType = 0
	// MLCActivationTypeSoftShrink - An activation type that implements the soft shrink activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/softShrink
	MLCActivationTypeSoftShrink MLCActivationType = 0
	// MLCActivationTypeSoftSign - An activation type that implements the parametric soft sign activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/softSign
	MLCActivationTypeSoftSign MLCActivationType = 0
	// MLCActivationTypeTanh - An activation type that implements the hyperbolic tangent activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/tanh
	MLCActivationTypeTanh MLCActivationType = 0
	// MLCActivationTypeTanhShrink - An activation type that implements the hyperbolic tangent shrink activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/tanhShrink
	MLCActivationTypeTanhShrink MLCActivationType = 0
	// MLCActivationTypeThreshold - An activation type that implements the threshold activation function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/threshold
	MLCActivationTypeThreshold MLCActivationType = 0
)

/* debug [enums.gen.go]: Processing enum MLCArithmeticOperation (31 cases) */
// MLCArithmeticOperation - Constants that describe an arithmetic operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation
type MLCArithmeticOperation uint

const (
	// MLCArithmeticOperationAcos - Calculates the element-wise inverse cosine of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/acos
	MLCArithmeticOperationAcos MLCArithmeticOperation = 0
	// MLCArithmeticOperationAcosh - Calculates the element-wise inverse hyperbolic cosine of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/acosh
	MLCArithmeticOperationAcosh MLCArithmeticOperation = 0
	// MLCArithmeticOperationAdd - Calculates the element-wise sum of the inputs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/add
	MLCArithmeticOperationAdd MLCArithmeticOperation = 0
	// MLCArithmeticOperationAsin - Calculates the element-wise inverse sine of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/asin
	MLCArithmeticOperationAsin MLCArithmeticOperation = 0
	// MLCArithmeticOperationAsinh - Calculates the element-wise inverse hyperbolic sine of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/asinh
	MLCArithmeticOperationAsinh MLCArithmeticOperation = 0
	// MLCArithmeticOperationAtan - Calculates the element-wise inverse tangent of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/atan
	MLCArithmeticOperationAtan MLCArithmeticOperation = 0
	// MLCArithmeticOperationAtanh - Calculates the element-wise inverse hyperbolic tangent of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/atanh
	MLCArithmeticOperationAtanh MLCArithmeticOperation = 0
	// MLCArithmeticOperationCeil - Calculates the element-wise ceiling of the inputs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/ceil
	MLCArithmeticOperationCeil MLCArithmeticOperation = 0
	// MLCArithmeticOperationCos - Calculates the element-wise cosine of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/cos
	MLCArithmeticOperationCos MLCArithmeticOperation = 0
	// MLCArithmeticOperationCosh - Calculates the element-wise hyperbolic cosine of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/cosh
	MLCArithmeticOperationCosh MLCArithmeticOperation = 0
	// MLCArithmeticOperationDivide - Calculates the element-wise division of the inputs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/divide
	MLCArithmeticOperationDivide MLCArithmeticOperation = 0
	// MLCArithmeticOperationDivideNoNaN - Calculates the element-wise division of the inputs, and returns   if the denominator is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/divideNoNaN
	MLCArithmeticOperationDivideNoNaN MLCArithmeticOperation = 0
	// MLCArithmeticOperationExp - Calculates the element-wise result of the exponent raised to the power of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/exp
	MLCArithmeticOperationExp MLCArithmeticOperation = 0
	// MLCArithmeticOperationExp2 - Calculates the element-wise result of the number 2 raised to the power of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/exp2
	MLCArithmeticOperationExp2 MLCArithmeticOperation = 0
	// MLCArithmeticOperationFloor - Calculates the element-wise floor of the inputs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/floor
	MLCArithmeticOperationFloor MLCArithmeticOperation = 0
	// MLCArithmeticOperationLog - Calculates the element-wise natural logarithm of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/log
	MLCArithmeticOperationLog MLCArithmeticOperation = 0
	// MLCArithmeticOperationLog2 - Calculates the element-wise base 2 logarithm of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/log2
	MLCArithmeticOperationLog2 MLCArithmeticOperation = 0
	// MLCArithmeticOperationMax - Calculates the element-wise maximum the inputs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/max
	MLCArithmeticOperationMax MLCArithmeticOperation = 0
	// MLCArithmeticOperationMin - Calculates the element-wise minimum of the inputs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/min
	MLCArithmeticOperationMin MLCArithmeticOperation = 0
	// MLCArithmeticOperationCount - The total number of arithmetic operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/MLCArithmeticOperationCount
	MLCArithmeticOperationCount MLCArithmeticOperation = 0
	// MLCArithmeticOperationMultiply - Calculates the element-wise product of the inputs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/multiply
	MLCArithmeticOperationMultiply MLCArithmeticOperation = 0
	// MLCArithmeticOperationMultiplyNoNaN - Calculates the element-wise product of the inputs, and returns   when the result isn’t a number or infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/multiplyNoNaN
	MLCArithmeticOperationMultiplyNoNaN MLCArithmeticOperation = 0
	// MLCArithmeticOperationPow - Calculates the element-wise first input raised to the power of the second input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/pow
	MLCArithmeticOperationPow MLCArithmeticOperation = 0
	// MLCArithmeticOperationRound - Calculates the element-wise rounding of the inputs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/round
	MLCArithmeticOperationRound MLCArithmeticOperation = 0
	// MLCArithmeticOperationRsqrt - Calculates the element-wise reciprocal of the square root of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/rsqrt
	MLCArithmeticOperationRsqrt MLCArithmeticOperation = 0
	// MLCArithmeticOperationSin - Calculates the element-wise sine of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/sin
	MLCArithmeticOperationSin MLCArithmeticOperation = 0
	// MLCArithmeticOperationSinh - Calculates the element-wise hyperbolic sine of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/sinh
	MLCArithmeticOperationSinh MLCArithmeticOperation = 0
	// MLCArithmeticOperationSqrt - Calculates the element-wise square root of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/sqrt
	MLCArithmeticOperationSqrt MLCArithmeticOperation = 0
	// MLCArithmeticOperationSubtract - Calculates the element-wise difference between the inputs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/subtract
	MLCArithmeticOperationSubtract MLCArithmeticOperation = 0
	// MLCArithmeticOperationTan - Calculates the element-wise tangent of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/tan
	MLCArithmeticOperationTan MLCArithmeticOperation = 0
	// MLCArithmeticOperationTanh - Calculates the element-wise hyperbolic tangent of the input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/tanh
	MLCArithmeticOperationTanh MLCArithmeticOperation = 0
)

/* debug [enums.gen.go]: Processing enum MLCComparisonOperation (13 cases) */
// MLCComparisonOperation - A comparison operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation
type MLCComparisonOperation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/equal
	MLCComparisonOperationEqual MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/greater
	MLCComparisonOperationGreater MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/greaterOrEqual
	MLCComparisonOperationGreaterOrEqual MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/less
	MLCComparisonOperationLess MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/lessOrEqual
	MLCComparisonOperationLessOrEqual MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/logicalAND
	MLCComparisonOperationLogicalAND MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/logicalNAND
	MLCComparisonOperationLogicalNAND MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/logicalNOR
	MLCComparisonOperationLogicalNOR MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/logicalNOT
	MLCComparisonOperationLogicalNOT MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/logicalOR
	MLCComparisonOperationLogicalOR MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/logicalXOR
	MLCComparisonOperationLogicalXOR MLCComparisonOperation = 0
	// MLCComparisonOperationCount - A number that represents the operation count.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/MLCComparisonOperationCount
	MLCComparisonOperationCount MLCComparisonOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/notEqual
	MLCComparisonOperationNotEqual MLCComparisonOperation = 0
)

/* debug [enums.gen.go]: Processing enum MLCConvolutionType (3 cases) */
// MLCConvolutionType - The convolution type specified for a convolution layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionType
type MLCConvolutionType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionType/depthwise
	MLCConvolutionTypeDepthwise MLCConvolutionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionType/standard
	MLCConvolutionTypeStandard MLCConvolutionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionType/transposed
	MLCConvolutionTypeTransposed MLCConvolutionType = 0
)

/* debug [enums.gen.go]: Processing enum MLCDataType (9 cases) */
// MLCDataType - A tensor data type.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType
type MLCDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType/boolean
	MLCDataTypeBoolean MLCDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType/float16
	MLCDataTypeFloat16 MLCDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType/float32
	MLCDataTypeFloat32 MLCDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType/int32
	MLCDataTypeInt32 MLCDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType/int64
	MLCDataTypeInt64 MLCDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType/int8
	MLCDataTypeInt8 MLCDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType/MLCDataTypeCount
	MLCDataTypeCount MLCDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType/MLCDataTypeInvalid
	MLCDataTypeInvalid MLCDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType/uint8
	MLCDataTypeUInt8 MLCDataType = 0
)

/* debug [enums.gen.go]: Processing enum MLCDeviceType (5 cases) */
// MLCDeviceType - A device type for execution of a neural network.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDeviceType
type MLCDeviceType uint

const (
	// MLCDeviceTypeANE - A device type that represents the Apple Neural Engine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDeviceType/ane
	MLCDeviceTypeANE MLCDeviceType = 0
	// MLCDeviceTypeAny - A device type that represents either the CPU or GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDeviceType/any
	MLCDeviceTypeAny MLCDeviceType = 0
	// MLCDeviceTypeCPU - A device type that represents the CPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDeviceType/cpu
	MLCDeviceTypeCPU MLCDeviceType = 0
	// MLCDeviceTypeGPU - A device type that represents the GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDeviceType/gpu
	MLCDeviceTypeGPU MLCDeviceType = 0
	// MLCDeviceTypeCount - A number that represents the number of device types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDeviceType/MLCDeviceTypeCount
	MLCDeviceTypeCount MLCDeviceType = 0
)

/* debug [enums.gen.go]: Processing enum MLCExecutionOptions (6 cases) */
// MLCExecutionOptions - A bitmask that specifies the options you use when executing a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCExecutionOptions
type MLCExecutionOptions uint

const (
	// MLCExecutionOptionsForwardForInference - The option to execute the forward pass for inference only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCExecutionOptions/forwardForInference
	MLCExecutionOptionsForwardForInference MLCExecutionOptions = 0
	// MLCExecutionOptionsNone - The option to execute the graph in the most efficient way possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCExecutionOptions/MLCExecutionOptionsNone
	MLCExecutionOptionsNone MLCExecutionOptions = 0
	// MLCExecutionOptionsPerLayerProfiling - The option to enable additional per-layer profiling information using signposts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCExecutionOptions/perLayerProfiling
	MLCExecutionOptionsPerLayerProfiling MLCExecutionOptions = 0
	// MLCExecutionOptionsProfiling - The option to return profiling information in the callback before returning from execution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCExecutionOptions/profiling
	MLCExecutionOptionsProfiling MLCExecutionOptions = 0
	// MLCExecutionOptionsSkipWritingInputDataToDevice - The option to skip writing input data to device memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCExecutionOptions/skipWritingInputDataToDevice
	MLCExecutionOptionsSkipWritingInputDataToDevice MLCExecutionOptions = 0
	// MLCExecutionOptionsSynchronous - The option to execute the graph synchronously.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCExecutionOptions/synchronous
	MLCExecutionOptionsSynchronous MLCExecutionOptions = 0
)

/* debug [enums.gen.go]: Processing enum MLCGradientClippingType (3 cases) */
// MLCGradientClippingType - A clipping type the system applies to a gradient.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGradientClippingType
type MLCGradientClippingType uint

const (
	// MLCGradientClippingTypeByGlobalNorm - An option that clips by global norm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGradientClippingType/byGlobalNorm
	MLCGradientClippingTypeByGlobalNorm MLCGradientClippingType = 0
	// MLCGradientClippingTypeByNorm - An option that clips by norm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGradientClippingType/byNorm
	MLCGradientClippingTypeByNorm MLCGradientClippingType = 0
	// MLCGradientClippingTypeByValue - An option that clips by value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGradientClippingType/byValue
	MLCGradientClippingTypeByValue MLCGradientClippingType = 0
)

/* debug [enums.gen.go]: Processing enum MLCGraphCompilationOptions (5 cases) */
// MLCGraphCompilationOptions - A bitmask that specifies the options you use when compiling a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraphCompilationOptions
type MLCGraphCompilationOptions uint

const (
	// MLCGraphCompilationOptionsComputeAllGradients - The option to compute all gradients during graph compilation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraphCompilationOptions/computeAllGradients
	MLCGraphCompilationOptionsComputeAllGradients MLCGraphCompilationOptions = 0
	// MLCGraphCompilationOptionsDebugLayers - The option to debug layers during graph compilation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraphCompilationOptions/debugLayers
	MLCGraphCompilationOptionsDebugLayers MLCGraphCompilationOptions = 0
	// MLCGraphCompilationOptionsDisableLayerFusion - The option to disable layer fusion during graph compilation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraphCompilationOptions/disableLayerFusion
	MLCGraphCompilationOptionsDisableLayerFusion MLCGraphCompilationOptions = 0
	// MLCGraphCompilationOptionsLinkGraphs - The option to link graphs during graph compilation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraphCompilationOptions/linkGraphs
	MLCGraphCompilationOptionsLinkGraphs MLCGraphCompilationOptions = 0
	// MLCGraphCompilationOptionsNone - The default option for graph compilation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraphCompilationOptions/MLCGraphCompilationOptionsNone
	MLCGraphCompilationOptionsNone MLCGraphCompilationOptions = 0
)

/* debug [enums.gen.go]: Processing enum MLCLossType (10 cases) */
// MLCLossType - A loss function.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType
type MLCLossType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/categoricalCrossEntropy
	MLCLossTypeCategoricalCrossEntropy MLCLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/cosineDistance
	MLCLossTypeCosineDistance MLCLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/hinge
	MLCLossTypeHinge MLCLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/huber
	MLCLossTypeHuber MLCLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/log
	MLCLossTypeLog MLCLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/meanAbsoluteError
	MLCLossTypeMeanAbsoluteError MLCLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/meanSquaredError
	MLCLossTypeMeanSquaredError MLCLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/MLCLossTypeCount
	MLCLossTypeCount MLCLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/sigmoidCrossEntropy
	MLCLossTypeSigmoidCrossEntropy MLCLossType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/softmaxCrossEntropy
	MLCLossTypeSoftmaxCrossEntropy MLCLossType = 0
)

/* debug [enums.gen.go]: Processing enum MLCLSTMResultMode (2 cases) */
// MLCLSTMResultMode - Constants that describe the result of an LSTM layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMResultMode
type MLCLSTMResultMode uint

const (
	// MLCLSTMResultModeOutput - A result mode that indicates the layer produces a single result tensor that represents the final output of the LSTM.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMResultMode/output
	MLCLSTMResultModeOutput MLCLSTMResultMode = 0
	// MLCLSTMResultModeOutputAndStates - A result mode that indicates the layer produces three result tensors that represent the final output of the LSTM, the last hidden state, and the cell state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMResultMode/outputAndStates
	MLCLSTMResultModeOutputAndStates MLCLSTMResultMode = 0
)

/* debug [enums.gen.go]: Processing enum MLCPaddingType (4 cases) */
// MLCPaddingType - A padding type that you specify for a padding layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingType
type MLCPaddingType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingType/constant
	MLCPaddingTypeConstant MLCPaddingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingType/reflect
	MLCPaddingTypeReflect MLCPaddingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingType/symmetric
	MLCPaddingTypeSymmetric MLCPaddingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingType/zero
	MLCPaddingTypeZero MLCPaddingType = 0
)

/* debug [enums.gen.go]: Processing enum MLCPoolingType (4 cases) */
// MLCPoolingType - A pooling function type for a pooling layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingType-8hrit
type MLCPoolingType int

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingType-8hrit/MLCPoolingTypeAverage
	MLCPoolingTypeAverage MLCPoolingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingType-8hrit/MLCPoolingTypeCount
	MLCPoolingTypeCount MLCPoolingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingType-8hrit/MLCPoolingTypeL2Norm
	MLCPoolingTypeL2Norm MLCPoolingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingType-8hrit/MLCPoolingTypeMax
	MLCPoolingTypeMax MLCPoolingType = 0
)

/* debug [enums.gen.go]: Processing enum MLCRandomInitializerType (5 cases) */
// MLCRandomInitializerType - An initializer type you use to create a tensor with random data.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType
type MLCRandomInitializerType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType/glorotUniform
	MLCRandomInitializerTypeGlorotUniform MLCRandomInitializerType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType/MLCRandomInitializerTypeCount
	MLCRandomInitializerTypeCount MLCRandomInitializerType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType/MLCRandomInitializerTypeInvalid
	MLCRandomInitializerTypeInvalid MLCRandomInitializerType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType/uniform
	MLCRandomInitializerTypeUniform MLCRandomInitializerType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType/xavier
	MLCRandomInitializerTypeXavier MLCRandomInitializerType = 0
)

/* debug [enums.gen.go]: Processing enum MLCReductionType (11 cases) */
// MLCReductionType - Constants that describe a reduction operation type.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType
type MLCReductionType uint

const (
	// MLCReductionTypeAll - A reduction operation that applies to all dimensions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/all
	MLCReductionTypeAll MLCReductionType = 0
	// MLCReductionTypeAny - A reduction operation that applies to any dimension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/any
	MLCReductionTypeAny MLCReductionType = 0
	// MLCReductionTypeArgMax - A reduction operation that applies to the maximum dimension you specify.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/argMax
	MLCReductionTypeArgMax MLCReductionType = 0
	// MLCReductionTypeArgMin - A reduction operation that applies to the minimum dimension you specify.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/argMin
	MLCReductionTypeArgMin MLCReductionType = 0
	// MLCReductionTypeL1Norm - A reduction operation that applies a lasso regularization penalty.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/l1Norm
	MLCReductionTypeL1Norm MLCReductionType = 0
	// MLCReductionTypeMax - A reduction operation that applies to the maximum dimension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/max
	MLCReductionTypeMax MLCReductionType = 0
	// MLCReductionTypeMean - A reduction operation that applies to the mean of the dimensions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/mean
	MLCReductionTypeMean MLCReductionType = 0
	// MLCReductionTypeMin - A reduction operation that applies to the minimum dimension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/min
	MLCReductionTypeMin MLCReductionType = 0
	// MLCReductionTypeCount - The total number of reduction operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/MLCReductionTypeCount
	MLCReductionTypeCount MLCReductionType = 0
	// MLCReductionTypeNone - A reduction operation that applies no reduction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/none
	MLCReductionTypeNone MLCReductionType = 0
	// MLCReductionTypeSum - A reduction operation that applies to the sum of the dimensions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/sum
	MLCReductionTypeSum MLCReductionType = 0
)

/* debug [enums.gen.go]: Processing enum MLCRegularizationType (3 cases) */
// MLCRegularizationType - A regularization function to use with an optimizer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRegularizationType
type MLCRegularizationType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRegularizationType/l1
	MLCRegularizationTypeL1 MLCRegularizationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRegularizationType/l2
	MLCRegularizationTypeL2 MLCRegularizationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRegularizationType/none
	MLCRegularizationTypeNone MLCRegularizationType = 0
)

/* debug [enums.gen.go]: Processing enum MLCSampleMode (2 cases) */
// MLCSampleMode - A sampling mode for an upsample layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSampleMode
type MLCSampleMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSampleMode/linear
	MLCSampleModeLinear MLCSampleMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSampleMode/nearest
	MLCSampleModeNearest MLCSampleMode = 0
)

/* debug [enums.gen.go]: Processing enum MLCSoftmaxOperation (2 cases) */
// MLCSoftmaxOperation - A softmax operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxOperation
type MLCSoftmaxOperation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxOperation/logSoftmax
	MLCSoftmaxOperationLogSoftmax MLCSoftmaxOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxOperation/softmax
	MLCSoftmaxOperationSoftmax MLCSoftmaxOperation = 0
)


