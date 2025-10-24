// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFileProviderServicing is the NSFileProviderServicing protocol interface.
//
// Support for providing a custom service source.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderServicing
type PFileProviderServicing interface {
	// Required methods
	SupportedServiceSourcesForItemIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: SupportedServiceSourcesForItemIdentifierCompletionHandler */
}
