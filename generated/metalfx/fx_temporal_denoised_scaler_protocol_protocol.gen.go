// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"unsafe"
)

// PFXTemporalDenoisedScaler is the MTLFXTemporalDenoisedScaler protocol interface.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 26.0+
//   - tvOS 18.0+
//
// See: doc://com.apple.metalfx/documentation/MetalFX/MTLFXTemporalDenoisedScaler
type PFXTemporalDenoisedScaler interface {
	// Required methods
	EncodeToCommandBuffer(commandBuffer unsafe.Pointer)/* debug [protocol_interface/required_method]: EncodeToCommandBuffer */
}
