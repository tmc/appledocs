// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"unsafe"
)

// PMTL4FXSpatialScaler is the MTL4FXSpatialScaler protocol interface.
//
// An upscaling effect that generates a higher resolution texture in a render pass by spatially analyzing an input texture.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metalfx/documentation/MetalFX/MTL4FXSpatialScaler
type PMTL4FXSpatialScaler interface {
	// Required methods
	EncodeToCommandBuffer(commandBuffer unsafe.Pointer)
}
