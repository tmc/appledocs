// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	ContentViewAtIndexEffectiveCharacterRange(index uint, outRange RangePointer /* not a class type */) IView
	HasContentViewAtIndexEffectiveCharacterRange() bool
	DidReplaceCharacters()
	HasDidReplaceCharacters() bool
	DrawCharactersInRangeForContentView(range_ foundation.Range, view IView)
	HasDrawCharactersInRangeForContentView() bool
	RectsForCharacterRange(range_ foundation.Range) []foundation.Value
	HasRectsForCharacterRange() bool
	ReplaceCharactersInRangeWithString(range_ foundation.Range, string_ foundation.foundation.INSString)
	HasReplaceCharactersInRangeWithString() bool
	ScrollRangeToVisible(range_ foundation.Range)
	HasScrollRangeToVisible() bool
	ShouldReplaceCharactersInRangesWithStrings(ranges []foundation.Value, strings []string) bool
	HasShouldReplaceCharactersInRangesWithStrings() bool
	StringAtIndexEffectiveRangeEndsWithSearchBoundary(characterIndex uint, outRange RangePointer /* not a class type */, outFlag objectivec.IObject) foundation.String
	HasStringAtIndexEffectiveRangeEndsWithSearchBoundary() bool
	StringLength() uint
	HasStringLength() bool
}
