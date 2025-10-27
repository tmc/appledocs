// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"unsafe"
)

// PMTL4FXTemporalDenoisedScaler is the MTL4FXTemporalDenoisedScaler protocol interface.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//
// See: doc://com.apple.metalfx/documentation/MetalFX/MTL4FXTemporalDenoisedScaler
type PMTL4FXTemporalDenoisedScaler interface {
	// Required methods
	EncodeToCommandBuffer(commandBuffer unsafe.Pointer)
}
