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
	AddAllocation(allocation unsafe.Pointer)/* debug [protocol_interface/required_method]: AddAllocation */
	AddAllocationsCount(allocations []objc.ID, count uint)/* debug [protocol_interface/required_method]: AddAllocationsCount */
	Commit()/* debug [protocol_interface/required_method]: Commit */
	ContainsAllocation(anAllocation unsafe.Pointer) bool/* debug [protocol_interface/required_method]: ContainsAllocation */
	EndResidency()/* debug [protocol_interface/required_method]: EndResidency */
	RemoveAllAllocations()/* debug [protocol_interface/required_method]: RemoveAllAllocations */
	RemoveAllocation(allocation unsafe.Pointer)/* debug [protocol_interface/required_method]: RemoveAllocation */
	RemoveAllocationsCount(allocations []objc.ID, count uint)/* debug [protocol_interface/required_method]: RemoveAllocationsCount */
	RequestResidency()/* debug [protocol_interface/required_method]: RequestResidency */
}
