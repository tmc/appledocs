// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"
)

// PFSVolumeOpenCloseOperations is the FSVolumeOpenCloseOperations protocol interface.
//
// Methods and properties implemented by volumes that want to receive open and close calls for each item.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSVolume/OpenCloseOperations
type PFSVolumeOpenCloseOperations interface {
	// Required methods
	CloseItemKeepingModesReplyHandler(item IFSItem, modes FSVolumeOpenModes, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: CloseItemKeepingModesReplyHandler */
	OpenItemWithModesReplyHandler(item IFSItem, modes FSVolumeOpenModes, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: OpenItemWithModesReplyHandler */
}
