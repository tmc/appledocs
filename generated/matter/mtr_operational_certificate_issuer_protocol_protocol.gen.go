// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"
)

// PMTROperationalCertificateIssuer is the MTROperationalCertificateIssuer protocol interface.
//
// Availability:
//   - Mac Catalyst 16.4+
//   - iOS 16.4+
//   - iPadOS 16.4+
//   - macOS 13.3+
//   - tvOS 16.4+
//   - visionOS 1.0+
//   - watchOS 9.4+
//
// See: doc://com.apple.matter/documentation/Matter/MTROperationalCertificateIssuer
type PMTROperationalCertificateIssuer interface {
	// Required methods
	IssueOperationalCertificateForRequestAttestationInfoControllerCompletion(csrInfo IMTROperationalCSRInfo, attestationInfo IMTRDeviceAttestationInfo, controller IMTRDeviceController, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: IssueOperationalCertificateForRequestAttestationInfoControllerCompletion */
}
