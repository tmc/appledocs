// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"
)

// PCounterSampleBuffer is the MTLCounterSampleBuffer protocol interface.
//
// A specialized memory buffer that stores a GPU’s counter set data.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 10.15+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLCounterSampleBuffer
type PCounterSampleBuffer interface {
	// Required methods
	ResolveCounterRange(range_ corefoundation.Range) foundation.Data/* debug [protocol_interface/required_method]: ResolveCounterRange */
}
