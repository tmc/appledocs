// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"
)

// PMTRNOCChainIssuer is the MTRNOCChainIssuer protocol interface.
//
// Availability:
//   - Mac Catalyst 16.1+ (Deprecated in 16.4)
//   - iOS 16.1+ (Deprecated in 16.4)
//   - iPadOS 16.1+ (Deprecated in 16.4)
//   - macOS 13.0+ (Deprecated in 13.3)
//   - tvOS 16.1+ (Deprecated in 16.4)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 9.1+ (Deprecated in 9.4)
//
// See: doc://com.apple.matter/documentation/Matter/MTRNOCChainIssuer
type PMTRNOCChainIssuer interface {
	// Required methods
	OnNOCChainGenerationNeededAttestationInfoOnNOCChainGenerationComplete(csrInfo ICSRInfo, attestationInfo IAttestationInfo, onNOCChainGenerationComplete unsafe.Pointer)/* debug [protocol_interface/required_method]: OnNOCChainGenerationNeededAttestationInfoOnNOCChainGenerationComplete */
}
