// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"
)

// PFSVolumeAccessCheckOperations is the FSVolumeAccessCheckOperations protocol interface.
//
// Methods and properties implemented by volumes that want to enforce access check operations.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSVolume/AccessCheckOperations
type PFSVolumeAccessCheckOperations interface {
	// Required methods
	CheckAccessToItemRequestedAccessReplyHandler(theItem IFSItem, access FSAccessMask, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: CheckAccessToItemRequestedAccessReplyHandler */
}
