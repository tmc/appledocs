// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFileProviderIncrementalContentFetching is the NSFileProviderIncrementalContentFetching protocol interface.
//
// Support for fetching changes to the item’s content.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderIncrementalContentFetching
type PFileProviderIncrementalContentFetching interface {
	// Required methods
	FetchContentsForItemWithIdentifierVersionUsingExistingContentsAtURLExistingVersionRequestCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, requestedVersion IFileProviderItemVersion, existingContents objc.IObject /* cross-framework: NSURL */, existingVersion IFileProviderItemVersion, request IFileProviderRequest, completionHandler unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: FetchContentsForItemWithIdentifierVersionUsingExistingContentsAtURLExistingVersionRequestCompletionHandler */
}
