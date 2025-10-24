// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPreviewingController is the QLPreviewingController protocol interface.
//
// For view based previews, the view controller that implements the QLPreviewingController protocol must at least implement one of the two following methods:   -[QLPreviewingController preparePreviewOfSearchableItemWithIdentifier:queryString:completionHandler:], to generate previews for Spotlight searchable items.   -[QLPreviewingController preparePreviewOfFileAtURL:completionHandler:], to generate previews for file URLs.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.quicklook/documentation/QuickLook/QLPreviewingController
type PPreviewingController interface {
	// Optional methods
	PreparePreviewOfFileAtURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, handler unsafe.Pointer)
	HasPreparePreviewOfFileAtURLCompletionHandler() bool
	PreparePreviewOfSearchableItemWithIdentifierQueryStringCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, queryString objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	HasPreparePreviewOfSearchableItemWithIdentifierQueryStringCompletionHandler() bool
	ProvidePreviewForFileRequestCompletionHandler(request IQLFilePreviewRequest, handler unsafe.Pointer)
	HasProvidePreviewForFileRequestCompletionHandler() bool
}
