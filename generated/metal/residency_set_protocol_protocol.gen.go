// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PResidencySet is the MTLResidencySet protocol interface.
//
// A collection of resource allocations that can move in and out of resident memory.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLResidencySet
type PResidencySet interface {
	// Required methods
	AddAllocation(allocation unsafe.Pointer)
	AddAllocationsCount(allocations []objc.ID, count uint)
	Commit()
	ContainsAllocation(anAllocation unsafe.Pointer) bool
	EndResidency()
	RemoveAllAllocations()
	RemoveAllocation(allocation unsafe.Pointer)
	RemoveAllocationsCount(allocations []objc.ID, count uint)
	RequestResidency()
}
