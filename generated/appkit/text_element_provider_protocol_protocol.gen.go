// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
)

// PTextElementProvider is the NSTextElementProvider protocol interface.
//
// A protocol the text content manager and its concrete subclasses conform to, which defines the interface for interacting with custom content types of a text document.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextElementProvider
type PTextElementProvider interface {
	// Required methods
	EnumerateTextElementsFromLocationOptionsUsingBlock(textLocation unsafe.Pointer, options TextContentManagerEnumerationOptions, block unsafe.Pointer) unsafe.Pointer
	ReplaceContentsInRangeWithTextElements(range_ ITextRange, textElements []TextElement)
	SynchronizeToBackingStore(completionHandler unsafe.Pointer)
	// Optional methods
	AdjustedRangeFromRangeForEditingTextSelection(textRange ITextRange, forEditingTextSelection bool) ITextRange
	HasAdjustedRangeFromRangeForEditingTextSelection() bool
	LocationFromLocationWithOffset(location unsafe.Pointer, offset int) unsafe.Pointer
	HasLocationFromLocationWithOffset() bool
	OffsetFromLocationToLocation(from unsafe.Pointer, to unsafe.Pointer) int
	HasOffsetFromLocationToLocation() bool
}
