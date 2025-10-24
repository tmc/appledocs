// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PResourceStateCommandEncoder is the MTLResourceStateCommandEncoder protocol interface.
//
// An encoder that encodes commands that modify resource configurations.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLResourceStateCommandEncoder
type PResourceStateCommandEncoder interface {
	// Required methods
	MoveTextureMappingsFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceTexture unsafe.Pointer, sourceSlice uint, sourceLevel uint, sourceOrigin objc.IObject /* cross-framework: MTLOrigin */, sourceSize objc.IObject /* cross-framework: MTLSize */, destinationTexture unsafe.Pointer, destinationSlice uint, destinationLevel uint, destinationOrigin objc.IObject /* cross-framework: MTLOrigin */)/* debug [protocol_interface/required_method]: MoveTextureMappingsFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin */
	UpdateFence(fence unsafe.Pointer)/* debug [protocol_interface/required_method]: UpdateFence */
	WaitForFence(fence unsafe.Pointer)/* debug [protocol_interface/required_method]: WaitForFence */
	// Optional methods
	UpdateTextureMappingModeIndirectBufferIndirectBufferOffset(texture unsafe.Pointer, mode SparseTextureMappingMode, indirectBuffer unsafe.Pointer, indirectBufferOffset uint)
	HasUpdateTextureMappingModeIndirectBufferIndirectBufferOffset() bool
	UpdateTextureMappingModeRegionMipLevelSlice(texture unsafe.Pointer, mode SparseTextureMappingMode, region objc.IObject /* cross-framework: MTLRegion */, mipLevel uint, slice uint)
	HasUpdateTextureMappingModeRegionMipLevelSlice() bool
	UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions(texture unsafe.Pointer, mode SparseTextureMappingMode, regions []objc.IObject /* cross-framework: MTLRegion */, mipLevels []uint, slices []uint, numRegions uint)
	HasUpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions() bool
}
