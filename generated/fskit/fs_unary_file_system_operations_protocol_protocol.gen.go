// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"
)

// PFSUnaryFileSystemOperations is the FSUnaryFileSystemOperations protocol interface.
//
// Operations performed by a unary file system.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSUnaryFileSystemOperations
type PFSUnaryFileSystemOperations interface {
	// Required methods
	LoadResourceOptionsReplyHandler(resource IFSResource, options IFSTaskOptions, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: LoadResourceOptionsReplyHandler */
	ProbeResourceReplyHandler(resource IFSResource, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: ProbeResourceReplyHandler */
	UnloadResourceOptionsReplyHandler(resource IFSResource, options IFSTaskOptions, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: UnloadResourceOptionsReplyHandler */
	// Optional methods
	DidFinishLoading()
	HasDidFinishLoading() bool
}
