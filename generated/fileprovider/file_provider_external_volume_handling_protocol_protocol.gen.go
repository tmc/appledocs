// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"
)

// PFileProviderExternalVolumeHandling is the NSFileProviderExternalVolumeHandling protocol interface.
//
// A protocol that defines the interface for handling external volumes.
//
// Availability:
//   - macOS 15.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderExternalVolumeHandling
type PFileProviderExternalVolumeHandling interface {
	// Required methods
	ShouldConnectExternalDomainWithCompletionHandler(completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: ShouldConnectExternalDomainWithCompletionHandler */
}
