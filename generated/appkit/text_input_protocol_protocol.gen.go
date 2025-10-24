// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PTextInput is the NSTextInput protocol interface.
//
// A set of methods that text views need to implement to interact properly with the text input management system.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextInput
type PTextInput interface {
	// Required methods
	AttributedSubstringFromRange(range_ corefoundation.Range) foundation.AttributedString
	CharacterIndexForPoint(point objc.IObject /* cross-framework: Point */) uint
	FirstRectForCharacterRange(range_ corefoundation.Range) corefoundation.Rect
	HasMarkedText() bool
	MarkedRange() corefoundation.Range
	SelectedRange() corefoundation.Range
	SetMarkedTextSelectedRange(string_ objc.IObject, selRange corefoundation.Range)
	UnmarkText()
	ValidAttributesForMarkedText() foundation.Array
}
