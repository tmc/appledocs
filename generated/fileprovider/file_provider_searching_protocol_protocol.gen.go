// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"
)

// PFileProviderSearching is the NSFileProviderSearching protocol interface.
//
// A protocol you implement to support searching in your file provider.
//
// Availability:
//   - macOS 26.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderSearching
type PFileProviderSearching interface {
	// Required methods
	SearchEnumeratorForStringSearchRequest(request IFileProviderStringSearchRequest) unsafe.Pointer/* debug [protocol_interface/required_method]: SearchEnumeratorForStringSearchRequest */
}
