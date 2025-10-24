// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFileProviderPartialContentFetching is the NSFileProviderPartialContentFetching protocol interface.
//
// Support for fetching part of a file’s content.
//
// Availability:
//   - macOS 12.3+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderPartialContentFetching
type PFileProviderPartialContentFetching interface {
	// Required methods
	FetchPartialContentsForItemWithIdentifierVersionRequestMinimalRangeAligningToOptionsCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, requestedVersion IFileProviderItemVersion, request IFileProviderRequest, requestedRange corefoundation.Range, alignment uint, options FileProviderFetchContentsOptions, completionHandler unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: FetchPartialContentsForItemWithIdentifierVersionRequestMinimalRangeAligningToOptionsCompletionHandler */
}
