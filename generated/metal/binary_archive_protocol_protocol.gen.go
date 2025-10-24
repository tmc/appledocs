// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PBinaryArchive is the MTLBinaryArchive protocol interface.
//
// A container for pipeline state descriptors and their associated compiled shader code.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLBinaryArchive
type PBinaryArchive interface {
	// Required methods
	AddComputePipelineFunctionsWithDescriptorError(descriptor IMTLComputePipelineDescriptor, error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: AddComputePipelineFunctionsWithDescriptorError */
	AddFunctionWithDescriptorLibraryError(descriptor IMTLFunctionDescriptor, library unsafe.Pointer, error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: AddFunctionWithDescriptorLibraryError */
	AddLibraryWithDescriptorError(descriptor IMTLStitchedLibraryDescriptor, error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: AddLibraryWithDescriptorError */
	AddMeshRenderPipelineFunctionsWithDescriptorError(descriptor IMTLMeshRenderPipelineDescriptor, error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: AddMeshRenderPipelineFunctionsWithDescriptorError */
	AddRenderPipelineFunctionsWithDescriptorError(descriptor IMTLRenderPipelineDescriptor, error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: AddRenderPipelineFunctionsWithDescriptorError */
	AddTileRenderPipelineFunctionsWithDescriptorError(descriptor IMTLTileRenderPipelineDescriptor, error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: AddTileRenderPipelineFunctionsWithDescriptorError */
	SerializeToURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: SerializeToURLError */
}
