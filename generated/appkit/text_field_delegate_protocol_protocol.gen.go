// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTextFieldDelegate is the NSTextFieldDelegate protocol interface.
//
// A protocol that a text field delegate can use to control its field editor action menu.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextFieldDelegate
type PTextFieldDelegate interface {
	// Optional methods
	TextFieldTextViewCandidatesForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult
	HasTextFieldTextViewCandidatesForSelectedRange() bool
	TextFieldTextViewShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool
	HasTextFieldTextViewShouldSelectCandidateAtIndex() bool
}

// TextFieldDelegate is a delegate implementation builder for the PTextFieldDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextFieldDelegate struct {
	_TextFieldTextViewCandidatesForSelectedRange func(textField ITextField, textView ITextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult
	_TextFieldTextViewShouldSelectCandidateAtIndex func(textField ITextField, textView ITextView, index uint) bool
}

// SetTextFieldTextViewCandidatesForSelectedRange sets the handler for the TextFieldTextViewCandidatesForSelectedRange delegate method.
func (d *TextFieldDelegate) SetTextFieldTextViewCandidatesForSelectedRange(f func(textField ITextField, textView ITextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult) {
	d._TextFieldTextViewCandidatesForSelectedRange = f
}

// SetTextFieldTextViewShouldSelectCandidateAtIndex sets the handler for the TextFieldTextViewShouldSelectCandidateAtIndex delegate method.
func (d *TextFieldDelegate) SetTextFieldTextViewShouldSelectCandidateAtIndex(f func(textField ITextField, textView ITextView, index uint) bool) {
	d._TextFieldTextViewShouldSelectCandidateAtIndex = f
}

// TextFieldTextViewCandidatesForSelectedRange implements the PTextFieldDelegate interface.
func (d *TextFieldDelegate) TextFieldTextViewCandidatesForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	if d._TextFieldTextViewCandidatesForSelectedRange != nil {
		return d._TextFieldTextViewCandidatesForSelectedRange(textField, textView, candidates, selectedRange)
	}
	var zero []foundation.TextCheckingResult
	return zero
}

// HasTextFieldTextViewCandidatesForSelectedRange returns true if a handler for TextFieldTextViewCandidatesForSelectedRange has been set.
func (d *TextFieldDelegate) HasTextFieldTextViewCandidatesForSelectedRange() bool {
	return d._TextFieldTextViewCandidatesForSelectedRange != nil
}

// TextFieldTextViewShouldSelectCandidateAtIndex implements the PTextFieldDelegate interface.
func (d *TextFieldDelegate) TextFieldTextViewShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool {
	if d._TextFieldTextViewShouldSelectCandidateAtIndex != nil {
		return d._TextFieldTextViewShouldSelectCandidateAtIndex(textField, textView, index)
	}
	var zero bool
	return zero
}

// HasTextFieldTextViewShouldSelectCandidateAtIndex returns true if a handler for TextFieldTextViewShouldSelectCandidateAtIndex has been set.
func (d *TextFieldDelegate) HasTextFieldTextViewShouldSelectCandidateAtIndex() bool {
	return d._TextFieldTextViewShouldSelectCandidateAtIndex != nil
}

// TextFieldDelegateObject wraps an existing Objective-C object that conforms to the PTextFieldDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TextFieldDelegateObject struct {
	objectivec.Object
}

// NewTextFieldDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTextFieldDelegate protocol.
func NewTextFieldDelegateObject(obj objectivec.Object) *TextFieldDelegateObject {
	return &TextFieldDelegateObject{obj}
}

// Make sure TextFieldDelegateObject implements PTextFieldDelegate.
var _ PTextFieldDelegate = (*TextFieldDelegateObject)(nil)

// TextFieldTextViewCandidatesForSelectedRange implements the PTextFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextFieldDelegateObject) TextFieldTextViewCandidatesForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	return objc.Send[[]foundation.TextCheckingResult](o.ID, objc.Sel("textField:textView:candidates:forSelectedRange:"), textField, textView, candidates, selectedRange)
}

// HasTextFieldTextViewCandidatesForSelectedRange returns true; this is a placeholder for optional method checks.
func (o *TextFieldDelegateObject) HasTextFieldTextViewCandidatesForSelectedRange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextFieldTextViewShouldSelectCandidateAtIndex implements the PTextFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextFieldDelegateObject) TextFieldTextViewShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool {
	return objc.Send[bool](o.ID, objc.Sel("textField:textView:shouldSelectCandidateAtIndex:"), textField, textView, index)
}

// HasTextFieldTextViewShouldSelectCandidateAtIndex returns true; this is a placeholder for optional method checks.
func (o *TextFieldDelegateObject) HasTextFieldTextViewShouldSelectCandidateAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
