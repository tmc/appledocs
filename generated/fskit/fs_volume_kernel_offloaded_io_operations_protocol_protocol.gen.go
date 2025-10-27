// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFSVolumeKernelOffloadedIOOperations is the FSVolumeKernelOffloadedIOOperations protocol interface.
//
// Methods and properties implemented by volumes that use kernel-offloaded I/O to achieve higher file transfer performance.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSVolumeKernelOffloadedIOOperations
type PFSVolumeKernelOffloadedIOOperations interface {
	// Required methods
	BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler(file IFSItem, offset unsafe.Pointer, length uintptr /* not a class type */, flags FSBlockmapFlags, operationID FSOperationID, packer IFSExtentPacker, reply unsafe.Pointer)
	CompleteIOForFileOffsetLengthStatusFlagsOperationIDReplyHandler(file IFSItem, offset unsafe.Pointer, length uintptr /* not a class type */, status foundation.foundation.INSError, flags FSCompleteIOFlags, operationID FSOperationID, reply unsafe.Pointer)
	CreateFileNamedInDirectoryAttributesPackerReplyHandler(name IFSFileName, directory IFSItem, attributes IFSItemSetAttributesRequest, packer IFSExtentPacker, reply unsafe.Pointer)
	LookupItemNamedInDirectoryPackerReplyHandler(name IFSFileName, directory IFSItem, packer IFSExtentPacker, reply unsafe.Pointer)
	// Optional methods
	PreallocateSpaceForFileAtOffsetLengthFlagsPackerReplyHandler(file IFSItem, offset unsafe.Pointer, length uintptr /* not a class type */, flags FSPreallocateFlags, packer IFSExtentPacker, reply unsafe.Pointer)
	HasPreallocateSpaceForFileAtOffsetLengthFlagsPackerReplyHandler() bool
}
