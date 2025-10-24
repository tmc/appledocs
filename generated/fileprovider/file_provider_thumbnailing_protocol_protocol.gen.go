// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFileProviderThumbnailing is the NSFileProviderThumbnailing protocol interface.
//
// Support for item thumbnails.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderThumbnailing
type PFileProviderThumbnailing interface {
	// Required methods
	FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers []string, size corefoundation.CGSize, perThumbnailCompletionHandler unsafe.Pointer, completionHandler unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler */
}
