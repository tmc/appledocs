// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	AttributedSubstringFromRange(range_ foundation.Range) foundation.AttributedString
	CharacterIndexForPoint(point corefoundation.CGPoint) uint
	FirstRectForCharacterRange(range_ foundation.Range) corefoundation.CGRect
	HasMarkedText() bool
	MarkedRange() foundation.Range
	SelectedRange() foundation.Range
	SetMarkedTextSelectedRange(string_ objectivec.IObject, selRange foundation.Range)
	UnmarkText()
	ValidAttributesForMarkedText() foundation.Array
}
