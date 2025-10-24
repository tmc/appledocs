// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"
)

// PFileProviderKnownFolderSupporting is the NSFileProviderKnownFolderSupporting protocol interface.
//
// A protocol that defines the interface for sharing known-folder locations with the system.
//
// Availability:
//   - macOS 15.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderKnownFolderSupporting
type PFileProviderKnownFolderSupporting interface {
	// Required methods
	GetKnownFolderLocationsCompletionHandler(knownFolders FileProviderKnownFolders, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: GetKnownFolderLocationsCompletionHandler */
}
