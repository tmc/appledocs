// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PComputePipelineState is the MTLComputePipelineState protocol interface.
//
// An interface that represents a GPU pipeline configuration for running kernels in a compute pass.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLComputePipelineState
type PComputePipelineState interface {
	// Required methods
	FunctionHandleWithFunction(function unsafe.Pointer) unsafe.Pointer
	FunctionHandleWithBinaryFunction(function unsafe.Pointer) unsafe.Pointer
	FunctionHandleWithName(name foundation.foundation.INSString) unsafe.Pointer
	ImageblockMemoryLengthForDimensions(imageblockDimensions Size) uint
	NewComputePipelineStateWithBinaryFunctionsError(additionalBinaryFunctions []objc.ID, error_ foundation.foundation.INSError) unsafe.Pointer
	NewComputePipelineStateWithAdditionalBinaryFunctionsError(functions []objc.ID, error_ foundation.foundation.INSError) unsafe.Pointer
	NewIntersectionFunctionTableWithDescriptor(descriptor IMTLIntersectionFunctionTableDescriptor) unsafe.Pointer
	NewVisibleFunctionTableWithDescriptor(descriptor IMTLVisibleFunctionTableDescriptor) unsafe.Pointer
}
