// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"
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
	TextContentStorageTextParagraphWithRange(textContentStorage ITextContentStorage, range_ corefoundation.Range) TextParagraph
	HasTextContentStorageTextParagraphWithRange() bool
}

// TextContentStorageDelegate is a delegate implementation builder for the PTextContentStorageDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextContentStorageDelegate struct {
	_TextContentStorageTextParagraphWithRange func(textContentStorage ITextContentStorage, range_ corefoundation.Range) TextParagraph
}

// SetTextContentStorageTextParagraphWithRange sets the handler for the TextContentStorageTextParagraphWithRange delegate method.
//
// Returns a custom paragraph for a range that you provide from the object’s attributed string.
func (d *TextContentStorageDelegate) SetTextContentStorageTextParagraphWithRange(f func(textContentStorage ITextContentStorage, range_ corefoundation.Range) TextParagraph) {
	d._TextContentStorageTextParagraphWithRange = f
}

// TextContentStorageTextParagraphWithRange implements the PTextContentStorageDelegate interface.
func (d *TextContentStorageDelegate) TextContentStorageTextParagraphWithRange(textContentStorage ITextContentStorage, range_ corefoundation.Range) TextParagraph {
	if d._TextContentStorageTextParagraphWithRange != nil {
		return d._TextContentStorageTextParagraphWithRange(textContentStorage, range_)
	}
	var zero TextParagraph
	return zero
}

// HasTextContentStorageTextParagraphWithRange returns true if a handler for TextContentStorageTextParagraphWithRange has been set.
func (d *TextContentStorageDelegate) HasTextContentStorageTextParagraphWithRange() bool {
	return d._TextContentStorageTextParagraphWithRange != nil
}
