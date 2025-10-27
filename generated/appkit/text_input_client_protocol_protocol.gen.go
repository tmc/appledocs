// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTextInputClient is the NSTextInputClient protocol interface.
//
// A set of methods that text views need to implement to interact properly with the text input management system.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextInputClient
type PTextInputClient interface {
	// Required methods
	AttributedSubstringForProposedRangeActualRange(range_ foundation.Range, actualRange RangePointer /* not a class type */) foundation.AttributedString
	InsertTextReplacementRange(string_ objectivec.IObject, replacementRange foundation.Range)
	SetMarkedTextSelectedRangeReplacementRange(string_ objectivec.IObject, selectedRange foundation.Range, replacementRange foundation.Range)
	// Optional methods
	AttributedString() foundation.AttributedString
	HasAttributedString() bool
}
