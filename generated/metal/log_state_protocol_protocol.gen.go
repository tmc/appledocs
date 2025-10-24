// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PLogState is the MTLLogState protocol interface.
//
// A container for shader log messages.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLLogState
type PLogState interface {
	// Required methods
	AddLogHandler(block unsafe.Pointer)/* debug [protocol_interface/required_method]: AddLogHandler */
}
