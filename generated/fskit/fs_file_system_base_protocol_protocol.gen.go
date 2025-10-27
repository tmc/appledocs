// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"
)

// PFSFileSystemBase is the FSFileSystemBase protocol interface.
//
// A protocol containing functionality supplied by FSKit to file system implementations.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSFileSystemBase
type PFSFileSystemBase interface {
	// Required methods
	WipeResourceCompletionHandler(resource IFSBlockDeviceResource, completion unsafe.Pointer)
}
