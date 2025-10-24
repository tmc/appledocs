// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPreviewingController is the QLPreviewingController protocol interface.
//
// A protocol for implementing a custom controller to create previews of files.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.quicklookui/documentation/QuickLookUI/QLPreviewingController
type PPreviewingController interface {
	// Optional methods
	PreparePreviewOfFileAtURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, handler unsafe.Pointer)
	HasPreparePreviewOfFileAtURLCompletionHandler() bool
	PreparePreviewOfSearchableItemWithIdentifierQueryStringCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, queryString objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	HasPreparePreviewOfSearchableItemWithIdentifierQueryStringCompletionHandler() bool
	ProvidePreviewForFileRequestCompletionHandler(request IQLFilePreviewRequest, handler unsafe.Pointer)
	HasProvidePreviewForFileRequestCompletionHandler() bool
}
