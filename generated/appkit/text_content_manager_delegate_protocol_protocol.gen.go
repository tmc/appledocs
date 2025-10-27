// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
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
	TextContentManagerTextElementAtLocation(textContentManager ITextContentManager, location unsafe.Pointer) ITextElement
	HasTextContentManagerTextElementAtLocation() bool
}

// TextContentManagerDelegate is a delegate implementation builder for the PTextContentManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextContentManagerDelegate struct {
	_TextContentManagerShouldEnumerateTextElementOptions func(textContentManager ITextContentManager, textElement ITextElement, options TextContentManagerEnumerationOptions) bool
	_TextContentManagerTextElementAtLocation func(textContentManager ITextContentManager, location unsafe.Pointer) ITextElement
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
func (d *TextContentManagerDelegate) SetTextContentManagerTextElementAtLocation(f func(textContentManager ITextContentManager, location unsafe.Pointer) ITextElement) {
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
func (d *TextContentManagerDelegate) TextContentManagerTextElementAtLocation(textContentManager ITextContentManager, location unsafe.Pointer) ITextElement {
	if d._TextContentManagerTextElementAtLocation != nil {
		return d._TextContentManagerTextElementAtLocation(textContentManager, location)
	}
	var zero ITextElement
	return zero
}

// HasTextContentManagerTextElementAtLocation returns true if a handler for TextContentManagerTextElementAtLocation has been set.
func (d *TextContentManagerDelegate) HasTextContentManagerTextElementAtLocation() bool {
	return d._TextContentManagerTextElementAtLocation != nil
}

// TextContentManagerDelegateObject wraps an existing Objective-C object that conforms to the PTextContentManagerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TextContentManagerDelegateObject struct {
	objectivec.Object
}

// NewTextContentManagerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTextContentManagerDelegate protocol.
func NewTextContentManagerDelegateObject(obj objectivec.Object) *TextContentManagerDelegateObject {
	return &TextContentManagerDelegateObject{obj}
}

// Make sure TextContentManagerDelegateObject implements PTextContentManagerDelegate.
var _ PTextContentManagerDelegate = (*TextContentManagerDelegateObject)(nil)

// TextContentManagerShouldEnumerateTextElementOptions implements the PTextContentManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextContentManagerDelegateObject) TextContentManagerShouldEnumerateTextElementOptions(textContentManager ITextContentManager, textElement ITextElement, options TextContentManagerEnumerationOptions) bool {
	return objc.Send[bool](o.ID, objc.Sel("textContentManager:shouldEnumerateTextElement:options:"), textContentManager, textElement, options)
}

// HasTextContentManagerShouldEnumerateTextElementOptions returns true; this is a placeholder for optional method checks.
func (o *TextContentManagerDelegateObject) HasTextContentManagerShouldEnumerateTextElementOptions() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextContentManagerTextElementAtLocation implements the PTextContentManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextContentManagerDelegateObject) TextContentManagerTextElementAtLocation(textContentManager ITextContentManager, location unsafe.Pointer) ITextElement {
	return objc.Send[ITextElement](o.ID, objc.Sel("textContentManager:textElementAtLocation:"), textContentManager, location)
}

// HasTextContentManagerTextElementAtLocation returns true; this is a placeholder for optional method checks.
func (o *TextContentManagerDelegateObject) HasTextContentManagerTextElementAtLocation() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
