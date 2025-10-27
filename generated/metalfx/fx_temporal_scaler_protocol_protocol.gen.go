// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"unsafe"
)

// PFXTemporalScaler is the MTLFXTemporalScaler protocol interface.
//
// An upscaling effect that generates a higher resolution texture in a render pass by analyzing multiple input textures over time.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//
// See: doc://com.apple.metalfx/documentation/MetalFX/MTLFXTemporalScaler
type PFXTemporalScaler interface {
	// Required methods
	EncodeToCommandBuffer(commandBuffer unsafe.Pointer)
}
