// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTL4Archive is the MTL4Archive protocol interface.
//
// A read-only container that stores pipeline states from a shader compiler.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4Archive
type PMTL4Archive interface {
	// Required methods
	NewBinaryFunctionWithDescriptorError(descriptor IMTL4BinaryFunctionDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorError(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewComputePipelineStateWithDescriptorError(descriptor IMTL4ComputePipelineDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorError(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewRenderPipelineStateWithDescriptorError(descriptor IMTL4PipelineDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
}
