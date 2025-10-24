// Code generated from Apple documentation for IdentityLookup. DO NOT EDIT.

package identitylookup

import (
	"unsafe"
)

// PILMessageFilterQueryHandling is the ILMessageFilterQueryHandling protocol interface.
//
// A set of methods implemented by a Message Filter app extension to handle query requests.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.identitylookup/documentation/IdentityLookup/ILMessageFilterQueryHandling
type PILMessageFilterQueryHandling interface {
	// Required methods
	HandleQueryRequestContextCompletion(queryRequest unsafe.Pointer, context unsafe.Pointer, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: HandleQueryRequestContextCompletion */
}
