// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objectivec"
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
	NewBinaryFunctionWithDescriptorError(descriptor IMTL4BinaryFunctionDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBinaryFunctionWithDescriptorError */
	NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorError(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorError */
	NewComputePipelineStateWithDescriptorError(descriptor IMTL4ComputePipelineDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithDescriptorError */
	NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorError(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorError */
	NewRenderPipelineStateWithDescriptorError(descriptor IMTL4PipelineDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithDescriptorError */
}
