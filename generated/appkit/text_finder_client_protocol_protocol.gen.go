// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
)

// PTextFinderClient is the NSTextFinderClient protocol interface.
//
// A set of methods implemented by objects that support searching using the   class and the in-window text find bar.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextFinderClient
type PTextFinderClient interface {
	// Optional methods
	ContentViewAtIndexEffectiveCharacterRange(index uint, outRange RangePointer /* not a class type */) View
	HasContentViewAtIndexEffectiveCharacterRange() bool
	DidReplaceCharacters()
	HasDidReplaceCharacters() bool
	DrawCharactersInRangeForContentView(range_ corefoundation.Range, view IView)
	HasDrawCharactersInRangeForContentView() bool
	RectsForCharacterRange(range_ corefoundation.Range) []foundation.Value
	HasRectsForCharacterRange() bool
	ReplaceCharactersInRangeWithString(range_ corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */)
	HasReplaceCharactersInRangeWithString() bool
	ScrollRangeToVisible(range_ corefoundation.Range)
	HasScrollRangeToVisible() bool
	ShouldReplaceCharactersInRangesWithStrings(ranges []foundation.Value, strings []string) bool
	HasShouldReplaceCharactersInRangesWithStrings() bool
	StringAtIndexEffectiveRangeEndsWithSearchBoundary(characterIndex uint, outRange RangePointer /* not a class type */, outFlag unsafe.Pointer) foundation.String
	HasStringAtIndexEffectiveRangeEndsWithSearchBoundary() bool
	StringLength() uint
	HasStringLength() bool
}
