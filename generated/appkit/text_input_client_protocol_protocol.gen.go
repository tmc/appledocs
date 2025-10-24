// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
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
	AttributedSubstringForProposedRangeActualRange(range_ corefoundation.Range, actualRange RangePointer /* not a class type */) foundation.AttributedString
	InsertTextReplacementRange(string_ objc.IObject, replacementRange corefoundation.Range)
	SetMarkedTextSelectedRangeReplacementRange(string_ objc.IObject, selectedRange corefoundation.Range, replacementRange corefoundation.Range)
	// Optional methods
	AttributedString() foundation.AttributedString
	HasAttributedString() bool
}
