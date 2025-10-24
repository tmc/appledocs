// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"unsafe"
)

// PMTL4FXTemporalScaler is the MTL4FXTemporalScaler protocol interface.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//
// See: doc://com.apple.metalfx/documentation/MetalFX/MTL4FXTemporalScaler
type PMTL4FXTemporalScaler interface {
	// Required methods
	EncodeToCommandBuffer(commandBuffer unsafe.Pointer)/* debug [protocol_interface/required_method]: EncodeToCommandBuffer */
}
