// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PTextStorageDelegate is the NSTextStorageDelegate protocol interface.
//
// The optional methods that delegates of text storage objects implement to handle text-edit processing.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextStorageDelegate
type PTextStorageDelegate interface {
	// Optional methods
	TextStorageDidProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange corefoundation.Range, delta int)
	HasTextStorageDidProcessEditingRangeChangeInLength() bool
	TextStorageWillProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange corefoundation.Range, delta int)
	HasTextStorageWillProcessEditingRangeChangeInLength() bool
}

// TextStorageDelegate is a delegate implementation builder for the PTextStorageDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextStorageDelegate struct {
	_TextStorageDidProcessEditingRangeChangeInLength func(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange corefoundation.Range, delta int)
	_TextStorageWillProcessEditingRangeChangeInLength func(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange corefoundation.Range, delta int)
}

// SetTextStorageDidProcessEditingRangeChangeInLength sets the handler for the TextStorageDidProcessEditingRangeChangeInLength delegate method.
//
// The method the framework calls when a text storage object has finished processing edits.
func (d *TextStorageDelegate) SetTextStorageDidProcessEditingRangeChangeInLength(f func(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange corefoundation.Range, delta int)) {
	d._TextStorageDidProcessEditingRangeChangeInLength = f
}

// SetTextStorageWillProcessEditingRangeChangeInLength sets the handler for the TextStorageWillProcessEditingRangeChangeInLength delegate method.
//
// The method the framework calls when a text storage object is about to process edits.
func (d *TextStorageDelegate) SetTextStorageWillProcessEditingRangeChangeInLength(f func(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange corefoundation.Range, delta int)) {
	d._TextStorageWillProcessEditingRangeChangeInLength = f
}

// TextStorageDidProcessEditingRangeChangeInLength implements the PTextStorageDelegate interface.
func (d *TextStorageDelegate) TextStorageDidProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange corefoundation.Range, delta int) {
	if d._TextStorageDidProcessEditingRangeChangeInLength != nil {
		d._TextStorageDidProcessEditingRangeChangeInLength(textStorage, editedMask, editedRange, delta)
	}
}

// HasTextStorageDidProcessEditingRangeChangeInLength returns true if a handler for TextStorageDidProcessEditingRangeChangeInLength has been set.
func (d *TextStorageDelegate) HasTextStorageDidProcessEditingRangeChangeInLength() bool {
	return d._TextStorageDidProcessEditingRangeChangeInLength != nil
}

// TextStorageWillProcessEditingRangeChangeInLength implements the PTextStorageDelegate interface.
func (d *TextStorageDelegate) TextStorageWillProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange corefoundation.Range, delta int) {
	if d._TextStorageWillProcessEditingRangeChangeInLength != nil {
		d._TextStorageWillProcessEditingRangeChangeInLength(textStorage, editedMask, editedRange, delta)
	}
}

// HasTextStorageWillProcessEditingRangeChangeInLength returns true if a handler for TextStorageWillProcessEditingRangeChangeInLength has been set.
func (d *TextStorageDelegate) HasTextStorageWillProcessEditingRangeChangeInLength() bool {
	return d._TextStorageWillProcessEditingRangeChangeInLength != nil
}
