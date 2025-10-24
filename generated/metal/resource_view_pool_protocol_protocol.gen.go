// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PResourceViewPool is the MTLResourceViewPool protocol interface.
//
// Contains views over resources of a specific type, and allows you to manage those views.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLResourceViewPool
type PResourceViewPool interface {
	// Required methods
	CopyResourceViewsFromPoolSourceRangeDestinationIndex(sourcePool unsafe.Pointer, sourceRange corefoundation.Range, destinationIndex uint) MTLResourceID/* debug [protocol_interface/required_method]: CopyResourceViewsFromPoolSourceRangeDestinationIndex */
}
