// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"
)

// PFSVolumeItemDeactivation is the FSVolumeItemDeactivation protocol interface.
//
// Methods and properties implemented by volumes that support deactivating items.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSVolume/ItemDeactivation
type PFSVolumeItemDeactivation interface {
	// Required methods
	DeactivateItemReplyHandler(item IFSItem, reply unsafe.Pointer)
}
