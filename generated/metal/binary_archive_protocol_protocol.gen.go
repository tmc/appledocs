// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
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
	AddComputePipelineFunctionsWithDescriptorError(descriptor IMTLComputePipelineDescriptor, error_ foundation.foundation.INSError) bool
	AddFunctionWithDescriptorLibraryError(descriptor IMTLFunctionDescriptor, library unsafe.Pointer, error_ foundation.foundation.INSError) bool
	AddLibraryWithDescriptorError(descriptor IMTLStitchedLibraryDescriptor, error_ foundation.foundation.INSError) bool
	AddMeshRenderPipelineFunctionsWithDescriptorError(descriptor IMTLMeshRenderPipelineDescriptor, error_ foundation.foundation.INSError) bool
	AddRenderPipelineFunctionsWithDescriptorError(descriptor IMTLRenderPipelineDescriptor, error_ foundation.foundation.INSError) bool
	AddTileRenderPipelineFunctionsWithDescriptorError(descriptor IMTLTileRenderPipelineDescriptor, error_ foundation.foundation.INSError) bool
	SerializeToURLError(url foundation.foundation.INSURL, error_ foundation.foundation.INSError) bool
}
