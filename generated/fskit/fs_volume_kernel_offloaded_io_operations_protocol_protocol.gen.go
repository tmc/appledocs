// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/coretelephony"
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
	BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler(file IFSItem, offset unsafe.Pointer, length uintptr /* not a class type */, flags FSBlockmapFlags, operationID FSOperationID /* typedef */, packer IFSExtentPacker, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler */
	CompleteIOForFileOffsetLengthStatusFlagsOperationIDReplyHandler(file IFSItem, offset unsafe.Pointer, length uintptr /* not a class type */, status objc.IObject /* cross-framework: Error */, flags FSCompleteIOFlags, operationID FSOperationID /* typedef */, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: CompleteIOForFileOffsetLengthStatusFlagsOperationIDReplyHandler */
	CreateFileNamedInDirectoryAttributesPackerReplyHandler(name IFSFileName, directory IFSItem, attributes IFSItemSetAttributesRequest, packer IFSExtentPacker, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: CreateFileNamedInDirectoryAttributesPackerReplyHandler */
	LookupItemNamedInDirectoryPackerReplyHandler(name IFSFileName, directory IFSItem, packer IFSExtentPacker, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: LookupItemNamedInDirectoryPackerReplyHandler */
	// Optional methods
	PreallocateSpaceForFileAtOffsetLengthFlagsPackerReplyHandler(file IFSItem, offset unsafe.Pointer, length uintptr /* not a class type */, flags FSPreallocateFlags, packer IFSExtentPacker, reply unsafe.Pointer)
	HasPreallocateSpaceForFileAtOffsetLengthFlagsPackerReplyHandler() bool
}
