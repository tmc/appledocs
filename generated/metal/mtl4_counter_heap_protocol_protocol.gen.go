// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTL4CounterHeap is the MTL4CounterHeap protocol interface.
//
// Represents an opaque, driver-controlled section of memory that can store GPU counter data.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4CounterHeap
type PMTL4CounterHeap interface {
	// Required methods
	InvalidateCounterRange(range_ foundation.Range)
	ResolveCounterRange(range_ foundation.Range) foundation.Data
}
