// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

/* debug [functions.gen.go]: Generating 13 functions for MLCompute */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// MLCompute Functions (13 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MLCActivationTypeDebugDescription func(CActivationType) unsafe.Pointer
	_MLCArithmeticOperationDebugDescription func(CArithmeticOperation) unsafe.Pointer
	_MLCComparisonOperationDebugDescription func(CComparisonOperation) unsafe.Pointer
	_MLCConvolutionTypeDebugDescription func(CConvolutionType) unsafe.Pointer
	_MLCGradientClippingTypeDebugDescription func(CGradientClippingType) unsafe.Pointer
	_MLCLossTypeDebugDescription func(CLossType) unsafe.Pointer
	_MLCLSTMResultModeDebugDescription func(CLSTMResultMode) unsafe.Pointer
	_MLCPaddingPolicyDebugDescription func(CPaddingPolicy) unsafe.Pointer
	_MLCPaddingTypeDebugDescription func(CPaddingType) unsafe.Pointer
	_MLCPoolingTypeDebugDescription func(CPoolingType) unsafe.Pointer
	_MLCReductionTypeDebugDescription func(CReductionType) unsafe.Pointer
	_MLCSampleModeDebugDescription func(CSampleMode) unsafe.Pointer
	_MLCSoftmaxOperationDebugDescription func(CSoftmaxOperation) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MLCActivationTypeDebugDescription, lib, "MLCActivationTypeDebugDescription")
	tryRegister(&_MLCArithmeticOperationDebugDescription, lib, "MLCArithmeticOperationDebugDescription")
	tryRegister(&_MLCComparisonOperationDebugDescription, lib, "MLCComparisonOperationDebugDescription")
	tryRegister(&_MLCConvolutionTypeDebugDescription, lib, "MLCConvolutionTypeDebugDescription")
	tryRegister(&_MLCGradientClippingTypeDebugDescription, lib, "MLCGradientClippingTypeDebugDescription")
	tryRegister(&_MLCLossTypeDebugDescription, lib, "MLCLossTypeDebugDescription")
	tryRegister(&_MLCLSTMResultModeDebugDescription, lib, "MLCLSTMResultModeDebugDescription")
	tryRegister(&_MLCPaddingPolicyDebugDescription, lib, "MLCPaddingPolicyDebugDescription")
	tryRegister(&_MLCPaddingTypeDebugDescription, lib, "MLCPaddingTypeDebugDescription")
	tryRegister(&_MLCPoolingTypeDebugDescription, lib, "MLCPoolingTypeDebugDescription")
	tryRegister(&_MLCReductionTypeDebugDescription, lib, "MLCReductionTypeDebugDescription")
	tryRegister(&_MLCSampleModeDebugDescription, lib, "MLCSampleModeDebugDescription")
	tryRegister(&_MLCSoftmaxOperationDebugDescription, lib, "MLCSoftmaxOperationDebugDescription")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// A textual description of the activation type, suitable for debugging.
//
// Added in macOS 11.0.
// A textual description of the activation type, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationType/debugDescription
func MLCActivationTypeDebugDescription(activationType CActivationType) unsafe.Pointer {
	return _MLCActivationTypeDebugDescription(activationType)
}/* debug [functions.gen.go/function]: MLCActivationTypeDebugDescription */

// A textual description of the arithmetic operation you use for debugging.
//
// Added in macOS 11.0.
// A textual description of the arithmetic operation you use for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticOperation/debugDescription
func MLCArithmeticOperationDebugDescription(operation CArithmeticOperation) unsafe.Pointer {
	return _MLCArithmeticOperationDebugDescription(operation)
}/* debug [functions.gen.go/function]: MLCArithmeticOperationDebugDescription */

// A textual description of the comparison operation, suitable for debugging.
//
// Added in macOS 11.3.
// A textual description of the comparison operation, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonOperation/debugDescription
func MLCComparisonOperationDebugDescription(operation CComparisonOperation) unsafe.Pointer {
	return _MLCComparisonOperationDebugDescription(operation)
}/* debug [functions.gen.go/function]: MLCComparisonOperationDebugDescription */

// A textual description of the convolution type, suitable for debugging.
//
// Added in macOS 11.0.
// A textual description of the convolution type, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionType/debugDescription
func MLCConvolutionTypeDebugDescription(convolutionType CConvolutionType) unsafe.Pointer {
	return _MLCConvolutionTypeDebugDescription(convolutionType)
}/* debug [functions.gen.go/function]: MLCConvolutionTypeDebugDescription */

// A textual description of the gradient clipping type, suitable for debugging.
//
// Added in macOS 12.0.
// A textual description of the gradient clipping type, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGradientClippingType/debugDescription
func MLCGradientClippingTypeDebugDescription(gradientClippingType CGradientClippingType) unsafe.Pointer {
	return _MLCGradientClippingTypeDebugDescription(gradientClippingType)
}/* debug [functions.gen.go/function]: MLCGradientClippingTypeDebugDescription */

// A textual description of the loss type, suitable for debugging.
//
// Added in macOS 11.0.
// A textual description of the loss type, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossType/debugDescription
func MLCLossTypeDebugDescription(lossType CLossType) unsafe.Pointer {
	return _MLCLossTypeDebugDescription(lossType)
}/* debug [functions.gen.go/function]: MLCLossTypeDebugDescription */

// A textual description of the LSTM result mode you use for debugging.
//
// Added in macOS 11.0.
// A textual description of the LSTM result mode you use for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMResultMode/debugDescription
func MLCLSTMResultModeDebugDescription(mode CLSTMResultMode) unsafe.Pointer {
	return _MLCLSTMResultModeDebugDescription(mode)
}/* debug [functions.gen.go/function]: MLCLSTMResultModeDebugDescription */

// A textual description of the padding policy, suitable for debugging.
//
// Added in macOS 11.0.
// A textual description of the padding policy, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingPolicyDebugDescription
func MLCPaddingPolicyDebugDescription(paddingPolicy CPaddingPolicy) unsafe.Pointer {
	return _MLCPaddingPolicyDebugDescription(paddingPolicy)
}/* debug [functions.gen.go/function]: MLCPaddingPolicyDebugDescription */

// A textual description of the padding type, suitable for debugging.
//
// Added in macOS 11.0.
// A textual description of the padding type, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingType/debugDescription
func MLCPaddingTypeDebugDescription(paddingType CPaddingType) unsafe.Pointer {
	return _MLCPaddingTypeDebugDescription(paddingType)
}/* debug [functions.gen.go/function]: MLCPaddingTypeDebugDescription */

// A textual description of the pooling type, suitable for debugging.
//
// Added in macOS 11.0.
// A textual description of the pooling type, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingTypeDebugDescription
func MLCPoolingTypeDebugDescription(poolingType CPoolingType) unsafe.Pointer {
	return _MLCPoolingTypeDebugDescription(poolingType)
}/* debug [functions.gen.go/function]: MLCPoolingTypeDebugDescription */

// A textual description of the reduction operation you use for debugging.
//
// Added in macOS 11.0.
// A textual description of the reduction operation you use for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionType/debugDescription
func MLCReductionTypeDebugDescription(reductionType CReductionType) unsafe.Pointer {
	return _MLCReductionTypeDebugDescription(reductionType)
}/* debug [functions.gen.go/function]: MLCReductionTypeDebugDescription */

// A textual description of the sample mode, suitable for debugging.
//
// Added in macOS 11.0.
// A textual description of the sample mode, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSampleMode/debugDescription
func MLCSampleModeDebugDescription(mode CSampleMode) unsafe.Pointer {
	return _MLCSampleModeDebugDescription(mode)
}/* debug [functions.gen.go/function]: MLCSampleModeDebugDescription */

// A textual description of the softmax operation, suitable for debugging.
//
// Added in macOS 11.0.
// A textual description of the softmax operation, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxOperation/debugDescription
func MLCSoftmaxOperationDebugDescription(operation CSoftmaxOperation) unsafe.Pointer {
	return _MLCSoftmaxOperationDebugDescription(operation)
}/* debug [functions.gen.go/function]: MLCSoftmaxOperationDebugDescription */




