// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFileProviderServiceSource is the NSFileProviderServiceSource protocol interface.
//
// A service that provides a custom communication channel between the host app and the File Provider extension.
//
// Availability:
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderServiceSource
type PFileProviderServiceSource interface {
	// Required methods
	MakeListenerEndpointAndReturnError(error_ unsafe.Pointer) foundation.XPCListenerEndpoint/* debug [protocol_interface/required_method]: MakeListenerEndpointAndReturnError */
}
