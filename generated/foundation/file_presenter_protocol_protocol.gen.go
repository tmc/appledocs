// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
)

// PFilePresenter is the NSFilePresenter protocol interface.
//
// The interface a file coordinator uses to inform an object presenting a file about changes to that file made elsewhere in the system.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSFilePresenter
type PFilePresenter interface {
	// Optional methods
	AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler unsafe.Pointer)
	HasAccommodatePresentedItemDeletionWithCompletionHandler() bool
	AccommodatePresentedSubitemDeletionAtURLCompletionHandler(url IURL, completionHandler unsafe.Pointer)
	HasAccommodatePresentedSubitemDeletionAtURLCompletionHandler() bool
	PresentedItemDidChange()
	HasPresentedItemDidChange() bool
	PresentedItemDidChangeUbiquityAttributes(attributes unsafe.Pointer)
	HasPresentedItemDidChangeUbiquityAttributes() bool
	PresentedItemDidMoveToURL(newURL IURL)
	HasPresentedItemDidMoveToURL() bool
	PresentedSubitemAtURLDidMoveToURL(oldURL IURL, newURL IURL)
	HasPresentedSubitemAtURLDidMoveToURL() bool
	PresentedSubitemDidAppearAtURL(url IURL)
	HasPresentedSubitemDidAppearAtURL() bool
	PresentedSubitemDidChangeAtURL(url IURL)
	HasPresentedSubitemDidChangeAtURL() bool
	RelinquishPresentedItemToReader(reader unsafe.Pointer)
	HasRelinquishPresentedItemToReader() bool
	RelinquishPresentedItemToWriter(writer unsafe.Pointer)
	HasRelinquishPresentedItemToWriter() bool
	SavePresentedItemChangesWithCompletionHandler(completionHandler unsafe.Pointer)
	HasSavePresentedItemChangesWithCompletionHandler() bool
}
