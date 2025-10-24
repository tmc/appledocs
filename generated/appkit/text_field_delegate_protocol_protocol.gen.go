// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
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
	TextFieldTextViewCandidatesForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.TextCheckingResult, selectedRange corefoundation.Range) []foundation.TextCheckingResult
	HasTextFieldTextViewCandidatesForSelectedRange() bool
	TextFieldTextViewShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool
	HasTextFieldTextViewShouldSelectCandidateAtIndex() bool
}

// TextFieldDelegate is a delegate implementation builder for the PTextFieldDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextFieldDelegate struct {
	_TextFieldTextViewCandidatesForSelectedRange func(textField ITextField, textView ITextView, candidates []foundation.TextCheckingResult, selectedRange corefoundation.Range) []foundation.TextCheckingResult
	_TextFieldTextViewShouldSelectCandidateAtIndex func(textField ITextField, textView ITextView, index uint) bool
}

// SetTextFieldTextViewCandidatesForSelectedRange sets the handler for the TextFieldTextViewCandidatesForSelectedRange delegate method.
func (d *TextFieldDelegate) SetTextFieldTextViewCandidatesForSelectedRange(f func(textField ITextField, textView ITextView, candidates []foundation.TextCheckingResult, selectedRange corefoundation.Range) []foundation.TextCheckingResult) {
	d._TextFieldTextViewCandidatesForSelectedRange = f
}

// SetTextFieldTextViewShouldSelectCandidateAtIndex sets the handler for the TextFieldTextViewShouldSelectCandidateAtIndex delegate method.
func (d *TextFieldDelegate) SetTextFieldTextViewShouldSelectCandidateAtIndex(f func(textField ITextField, textView ITextView, index uint) bool) {
	d._TextFieldTextViewShouldSelectCandidateAtIndex = f
}

// TextFieldTextViewCandidatesForSelectedRange implements the PTextFieldDelegate interface.
func (d *TextFieldDelegate) TextFieldTextViewCandidatesForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.TextCheckingResult, selectedRange corefoundation.Range) []foundation.TextCheckingResult {
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
