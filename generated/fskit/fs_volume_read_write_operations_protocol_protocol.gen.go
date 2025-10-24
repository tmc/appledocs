// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFSVolumeReadWriteOperations is the FSVolumeReadWriteOperations protocol interface.
//
// Methods implemented for read and write operations that deliver data to and from the extension.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSVolume/ReadWriteOperations
type PFSVolumeReadWriteOperations interface {
	// Required methods
	ReadFromFileOffsetLengthIntoBufferReplyHandler(item IFSItem, offset unsafe.Pointer, length uintptr /* not a class type */, buffer IFSMutableFileDataBuffer, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: ReadFromFileOffsetLengthIntoBufferReplyHandler */
	WriteContentsToFileAtOffsetReplyHandler(contents objc.IObject /* cross-framework: NSData */, item IFSItem, offset unsafe.Pointer, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: WriteContentsToFileAtOffsetReplyHandler */
}
