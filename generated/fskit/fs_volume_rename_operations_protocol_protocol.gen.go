// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"
)

// PFSVolumeRenameOperations is the FSVolumeRenameOperations protocol interface.
//
// Methods and properties implemented by volumes that support renaming the volume.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSVolume/RenameOperations
type PFSVolumeRenameOperations interface {
	// Required methods
	SetVolumeNameReplyHandler(name IFSFileName, reply unsafe.Pointer)
}
