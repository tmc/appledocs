// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

// PMTL4ArgumentTable is the MTL4ArgumentTable protocol interface.
//
// Provides a mechanism to manage and provide resource bindings for buffers, textures, sampler states and other Metal resources.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4ArgumentTable
type PMTL4ArgumentTable interface {
	// Required methods
	SetAddressAttributeStrideAtIndex(gpuAddress GPUAddress, stride uint, bindingIndex uint)
	SetAddressAtIndex(gpuAddress GPUAddress, bindingIndex uint)
	SetResourceAtBufferIndex(resourceID ResourceID, bindingIndex uint)
	SetSamplerStateAtIndex(resourceID ResourceID, bindingIndex uint)
	SetTextureAtIndex(resourceID ResourceID, bindingIndex uint)
}
