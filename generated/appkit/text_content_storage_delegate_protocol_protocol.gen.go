// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTextContentStorageDelegate is the NSTextContentStorageDelegate protocol interface.
//
// The optional methods that delegates of content storage objects implement to handle content processing.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextContentStorageDelegate
type PTextContentStorageDelegate interface {
	// Optional methods
	TextContentStorageTextParagraphWithRange(textContentStorage ITextContentStorage, range_ foundation.Range) ITextParagraph
	HasTextContentStorageTextParagraphWithRange() bool
}

// TextContentStorageDelegate is a delegate implementation builder for the PTextContentStorageDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextContentStorageDelegate struct {
	_TextContentStorageTextParagraphWithRange func(textContentStorage ITextContentStorage, range_ foundation.Range) ITextParagraph
}

// SetTextContentStorageTextParagraphWithRange sets the handler for the TextContentStorageTextParagraphWithRange delegate method.
//
// Returns a custom paragraph for a range that you provide from the object’s attributed string.
func (d *TextContentStorageDelegate) SetTextContentStorageTextParagraphWithRange(f func(textContentStorage ITextContentStorage, range_ foundation.Range) ITextParagraph) {
	d._TextContentStorageTextParagraphWithRange = f
}

// TextContentStorageTextParagraphWithRange implements the PTextContentStorageDelegate interface.
func (d *TextContentStorageDelegate) TextContentStorageTextParagraphWithRange(textContentStorage ITextContentStorage, range_ foundation.Range) ITextParagraph {
	if d._TextContentStorageTextParagraphWithRange != nil {
		return d._TextContentStorageTextParagraphWithRange(textContentStorage, range_)
	}
	var zero ITextParagraph
	return zero
}

// HasTextContentStorageTextParagraphWithRange returns true if a handler for TextContentStorageTextParagraphWithRange has been set.
func (d *TextContentStorageDelegate) HasTextContentStorageTextParagraphWithRange() bool {
	return d._TextContentStorageTextParagraphWithRange != nil
}

// TextContentStorageDelegateObject wraps an existing Objective-C object that conforms to the PTextContentStorageDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TextContentStorageDelegateObject struct {
	objectivec.Object
}

// NewTextContentStorageDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTextContentStorageDelegate protocol.
func NewTextContentStorageDelegateObject(obj objectivec.Object) *TextContentStorageDelegateObject {
	return &TextContentStorageDelegateObject{obj}
}

// Make sure TextContentStorageDelegateObject implements PTextContentStorageDelegate.
var _ PTextContentStorageDelegate = (*TextContentStorageDelegateObject)(nil)

// TextContentStorageTextParagraphWithRange implements the PTextContentStorageDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextContentStorageDelegateObject) TextContentStorageTextParagraphWithRange(textContentStorage ITextContentStorage, range_ foundation.Range) ITextParagraph {
	return objc.Send[ITextParagraph](o.ID, objc.Sel("textContentStorage:textParagraphWithRange:"), textContentStorage, range_)
}

// HasTextContentStorageTextParagraphWithRange returns true; this is a placeholder for optional method checks.
func (o *TextContentStorageDelegateObject) HasTextContentStorageTextParagraphWithRange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
