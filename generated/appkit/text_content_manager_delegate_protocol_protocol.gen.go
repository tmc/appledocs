// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PTextContentManagerDelegate is the NSTextContentManagerDelegate protocol interface.
//
// The optional methods that delegates of content manager objects implement for customizing or validating text elements.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextContentManagerDelegate
type PTextContentManagerDelegate interface {
	// Optional methods
	TextContentManagerShouldEnumerateTextElementOptions(textContentManager ITextContentManager, textElement ITextElement, options TextContentManagerEnumerationOptions) bool
	HasTextContentManagerShouldEnumerateTextElementOptions() bool
	TextContentManagerTextElementAtLocation(textContentManager ITextContentManager, location objc.IObject) TextElement
	HasTextContentManagerTextElementAtLocation() bool
}

// TextContentManagerDelegate is a delegate implementation builder for the PTextContentManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextContentManagerDelegate struct {
	_TextContentManagerShouldEnumerateTextElementOptions func(textContentManager ITextContentManager, textElement ITextElement, options TextContentManagerEnumerationOptions) bool
	_TextContentManagerTextElementAtLocation func(textContentManager ITextContentManager, location objc.IObject) TextElement
}

// SetTextContentManagerShouldEnumerateTextElementOptions sets the handler for the TextContentManagerShouldEnumerateTextElementOptions delegate method.
//
// Returns a Boolean value that indicates whether the framework should skip this text element in the enumeration.
func (d *TextContentManagerDelegate) SetTextContentManagerShouldEnumerateTextElementOptions(f func(textContentManager ITextContentManager, textElement ITextElement, options TextContentManagerEnumerationOptions) bool) {
	d._TextContentManagerShouldEnumerateTextElementOptions = f
}

// SetTextContentManagerTextElementAtLocation sets the handler for the TextContentManagerTextElementAtLocation delegate method.
//
// The method the framework calls to return the text element at a specific location.
func (d *TextContentManagerDelegate) SetTextContentManagerTextElementAtLocation(f func(textContentManager ITextContentManager, location objc.IObject) TextElement) {
	d._TextContentManagerTextElementAtLocation = f
}

// TextContentManagerShouldEnumerateTextElementOptions implements the PTextContentManagerDelegate interface.
func (d *TextContentManagerDelegate) TextContentManagerShouldEnumerateTextElementOptions(textContentManager ITextContentManager, textElement ITextElement, options TextContentManagerEnumerationOptions) bool {
	if d._TextContentManagerShouldEnumerateTextElementOptions != nil {
		return d._TextContentManagerShouldEnumerateTextElementOptions(textContentManager, textElement, options)
	}
	var zero bool
	return zero
}

// HasTextContentManagerShouldEnumerateTextElementOptions returns true if a handler for TextContentManagerShouldEnumerateTextElementOptions has been set.
func (d *TextContentManagerDelegate) HasTextContentManagerShouldEnumerateTextElementOptions() bool {
	return d._TextContentManagerShouldEnumerateTextElementOptions != nil
}

// TextContentManagerTextElementAtLocation implements the PTextContentManagerDelegate interface.
func (d *TextContentManagerDelegate) TextContentManagerTextElementAtLocation(textContentManager ITextContentManager, location objc.IObject) TextElement {
	if d._TextContentManagerTextElementAtLocation != nil {
		return d._TextContentManagerTextElementAtLocation(textContentManager, location)
	}
	var zero TextElement
	return zero
}

// HasTextContentManagerTextElementAtLocation returns true if a handler for TextContentManagerTextElementAtLocation has been set.
func (d *TextContentManagerDelegate) HasTextContentManagerTextElementAtLocation() bool {
	return d._TextContentManagerTextElementAtLocation != nil
}
