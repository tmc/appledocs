// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"
)

// PFSVolumePreallocateOperations is the FSVolumePreallocateOperations protocol interface.
//
// Methods and properties implemented by volumes that want to offer preallocation functions.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSVolume/PreallocateOperations
type PFSVolumePreallocateOperations interface {
	// Required methods
	PreallocateSpaceForItemAtOffsetLengthFlagsReplyHandler(item IFSItem, offset unsafe.Pointer, length uintptr /* not a class type */, flags FSPreallocateFlags, reply unsafe.Pointer)
}
