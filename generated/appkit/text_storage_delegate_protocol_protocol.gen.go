// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	TextStorageDidProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	HasTextStorageDidProcessEditingRangeChangeInLength() bool
	TextStorageWillProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	HasTextStorageWillProcessEditingRangeChangeInLength() bool
}

// TextStorageDelegate is a delegate implementation builder for the PTextStorageDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextStorageDelegate struct {
	_TextStorageDidProcessEditingRangeChangeInLength func(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	_TextStorageWillProcessEditingRangeChangeInLength func(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
}

// SetTextStorageDidProcessEditingRangeChangeInLength sets the handler for the TextStorageDidProcessEditingRangeChangeInLength delegate method.
//
// The method the framework calls when a text storage object has finished processing edits.
func (d *TextStorageDelegate) SetTextStorageDidProcessEditingRangeChangeInLength(f func(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)) {
	d._TextStorageDidProcessEditingRangeChangeInLength = f
}

// SetTextStorageWillProcessEditingRangeChangeInLength sets the handler for the TextStorageWillProcessEditingRangeChangeInLength delegate method.
//
// The method the framework calls when a text storage object is about to process edits.
func (d *TextStorageDelegate) SetTextStorageWillProcessEditingRangeChangeInLength(f func(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)) {
	d._TextStorageWillProcessEditingRangeChangeInLength = f
}

// TextStorageDidProcessEditingRangeChangeInLength implements the PTextStorageDelegate interface.
func (d *TextStorageDelegate) TextStorageDidProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	if d._TextStorageDidProcessEditingRangeChangeInLength != nil {
		d._TextStorageDidProcessEditingRangeChangeInLength(textStorage, editedMask, editedRange, delta)
	}
}

// HasTextStorageDidProcessEditingRangeChangeInLength returns true if a handler for TextStorageDidProcessEditingRangeChangeInLength has been set.
func (d *TextStorageDelegate) HasTextStorageDidProcessEditingRangeChangeInLength() bool {
	return d._TextStorageDidProcessEditingRangeChangeInLength != nil
}

// TextStorageWillProcessEditingRangeChangeInLength implements the PTextStorageDelegate interface.
func (d *TextStorageDelegate) TextStorageWillProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	if d._TextStorageWillProcessEditingRangeChangeInLength != nil {
		d._TextStorageWillProcessEditingRangeChangeInLength(textStorage, editedMask, editedRange, delta)
	}
}

// HasTextStorageWillProcessEditingRangeChangeInLength returns true if a handler for TextStorageWillProcessEditingRangeChangeInLength has been set.
func (d *TextStorageDelegate) HasTextStorageWillProcessEditingRangeChangeInLength() bool {
	return d._TextStorageWillProcessEditingRangeChangeInLength != nil
}

// TextStorageDelegateObject wraps an existing Objective-C object that conforms to the PTextStorageDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TextStorageDelegateObject struct {
	objectivec.Object
}

// NewTextStorageDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTextStorageDelegate protocol.
func NewTextStorageDelegateObject(obj objectivec.Object) *TextStorageDelegateObject {
	return &TextStorageDelegateObject{obj}
}

// Make sure TextStorageDelegateObject implements PTextStorageDelegate.
var _ PTextStorageDelegate = (*TextStorageDelegateObject)(nil)

// TextStorageDidProcessEditingRangeChangeInLength implements the PTextStorageDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextStorageDelegateObject) TextStorageDidProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.Send[objc.ID](o.ID, objc.Sel("textStorage:didProcessEditing:range:changeInLength:"), textStorage, editedMask, editedRange, delta)
}

// HasTextStorageDidProcessEditingRangeChangeInLength returns true; this is a placeholder for optional method checks.
func (o *TextStorageDelegateObject) HasTextStorageDidProcessEditingRangeChangeInLength() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextStorageWillProcessEditingRangeChangeInLength implements the PTextStorageDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextStorageDelegateObject) TextStorageWillProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.Send[objc.ID](o.ID, objc.Sel("textStorage:willProcessEditing:range:changeInLength:"), textStorage, editedMask, editedRange, delta)
}

// HasTextStorageWillProcessEditingRangeChangeInLength returns true; this is a placeholder for optional method checks.
func (o *TextStorageDelegateObject) HasTextStorageWillProcessEditingRangeChangeInLength() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
