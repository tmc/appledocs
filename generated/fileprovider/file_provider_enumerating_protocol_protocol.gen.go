// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"
)

// PFileProviderEnumerating is the NSFileProviderEnumerating protocol interface.
//
// Support for enumerating the file provider’s contents.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderEnumerating
type PFileProviderEnumerating interface {
	// Required methods
	EnumeratorForContainerItemIdentifierRequestError(containerItemIdentifier FileProviderItemIdentifier /* typedef */, request IFileProviderRequest, error_ unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: EnumeratorForContainerItemIdentifierRequestError */
}
