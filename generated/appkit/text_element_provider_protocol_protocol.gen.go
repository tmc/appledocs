// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	EnumerateTextElementsFromLocationOptionsUsingBlock(textLocation objc.IObject, options TextContentManagerEnumerationOptions, block unsafe.Pointer) objc.ID
	ReplaceContentsInRangeWithTextElements(range_ ITextRange, textElements []TextElement)
	SynchronizeToBackingStore(completionHandler unsafe.Pointer)
	// Optional methods
	AdjustedRangeFromRangeForEditingTextSelection(textRange ITextRange, forEditingTextSelection bool) TextRange
	HasAdjustedRangeFromRangeForEditingTextSelection() bool
	LocationFromLocationWithOffset(location objc.IObject, offset int) objc.ID
	HasLocationFromLocationWithOffset() bool
	OffsetFromLocationToLocation(from objc.IObject, to objc.IObject) int
	HasOffsetFromLocationToLocation() bool
}
