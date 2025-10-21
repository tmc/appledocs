// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

// Enum types and constants
// MLCComparisonOperation - A comparison operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation
type CComparisonOperation uint

// MLCDataType - A tensor data type.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType
type CDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDataType/float32
	CDataTypeFloat32 CDataType = 0
)

// MLCDeviceType - A device type for execution of a neural network.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDeviceType
type CDeviceType uint

// MLCExecutionOptions - A bitmask that specifies the options you use when executing a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCExecutionOptions
type CExecutionOptions uint

// MLCGraphCompilationOptions - A bitmask that specifies the options you use when compiling a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraphCompilationOptions
type CGraphCompilationOptions uint

// MLCPaddingPolicy - A padding policy that you specify for a convolution or pooling layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingPolicy-14ba7
type CPaddingPolicy uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingPolicy-14ba7/MLCPaddingPolicySame
	CPaddingPolicySame CPaddingPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingPolicy-14ba7/MLCPaddingPolicyUsePaddingSize
	CPaddingPolicyUsePaddingSize CPaddingPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingPolicy-14ba7/MLCPaddingPolicyValid
	CPaddingPolicyValid CPaddingPolicy = 0
)

// MLCPaddingType - A padding type that you specify for a padding layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingType
type CPaddingType uint

// MLCRandomInitializerType - An initializer type you use to create a tensor with random data.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType
type CRandomInitializerType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType/MLCRandomInitializerTypeCount
	CRandomInitializerTypeCount CRandomInitializerType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType/MLCRandomInitializerTypeInvalid
	CRandomInitializerTypeInvalid CRandomInitializerType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType/glorotUniform
	CRandomInitializerTypeGlorotUniform CRandomInitializerType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType/uniform
	CRandomInitializerTypeUniform CRandomInitializerType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRandomInitializerType/xavier
	CRandomInitializerTypeXavier CRandomInitializerType = 0
)

// MLCReductionType - Constants that describe a reduction operation type.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType
type CReductionType uint

const (
	// CReductionTypeNone - A reduction operation that applies no reduction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/none
	CReductionTypeNone CReductionType = 0
	// CReductionTypeSum - A reduction operation that applies to the sum of the dimensions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/sum
	CReductionTypeSum CReductionType = 0
)

// MLCSoftmaxOperation - A softmax operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxOperation
type CSoftmaxOperation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxOperation/logSoftmax
	CSoftmaxOperationLogSoftmax CSoftmaxOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxOperation/softmax
	CSoftmaxOperationSoftmax CSoftmaxOperation = 0
)


