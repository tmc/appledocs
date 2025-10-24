// Code generated from Apple documentation for IdentityLookup. DO NOT EDIT.

package identitylookup

import (
	"unsafe"
)

// PILMessageFilterCapabilitiesQueryHandling is the ILMessageFilterCapabilitiesQueryHandling protocol interface.
//
// A set of methods implemented by a Message Filter app extension to handle capabilities query requests.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.identitylookup/documentation/IdentityLookup/ILMessageFilterCapabilitiesQueryHandling
type PILMessageFilterCapabilitiesQueryHandling interface {
	// Required methods
	HandleCapabilitiesQueryRequestContextCompletion(capabilitiesQueryRequest unsafe.Pointer, context unsafe.Pointer, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: HandleCapabilitiesQueryRequestContextCompletion */
}
