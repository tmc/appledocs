// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTRKeypair is the MTRKeypair protocol interface.
//
// Availability:
//   - Mac Catalyst 16.1+
//   - iOS 16.1+
//   - iPadOS 16.1+
//   - macOS 13.0+
//   - tvOS 16.1+
//   - visionOS 1.0+
//   - watchOS 9.1+
//
// See: doc://com.apple.matter/documentation/Matter/MTRKeypair
type PMTRKeypair interface {
	// Optional methods
	CopyPublicKey() unsafe.Pointer
	HasCopyPublicKey() bool
	PublicKey() unsafe.Pointer
	HasPublicKey() bool
	SignMessageECDSA_DER(message objc.IObject /* cross-framework: NSData */) foundation.Data
	HasSignMessageECDSA_DER() bool
	SignMessageECDSA_RAW(message objc.IObject /* cross-framework: NSData */) foundation.Data
	HasSignMessageECDSA_RAW() bool
}
