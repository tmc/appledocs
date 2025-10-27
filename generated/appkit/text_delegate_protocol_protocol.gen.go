// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTextDelegate is the NSTextDelegate protocol interface.
//
// A set of optional methods implemented by the delegate of an   object to edit text and change text formats.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextDelegate
type PTextDelegate interface {
	// Optional methods
	TextDidBeginEditing(notification foundation.foundation.INSNotification)
	HasTextDidBeginEditing() bool
	TextDidChange(notification foundation.foundation.INSNotification)
	HasTextDidChange() bool
	TextDidEndEditing(notification foundation.foundation.INSNotification)
	HasTextDidEndEditing() bool
	TextShouldBeginEditing(textObject IText) bool
	HasTextShouldBeginEditing() bool
	TextShouldEndEditing(textObject IText) bool
	HasTextShouldEndEditing() bool
}

// TextDelegate is a delegate implementation builder for the PTextDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextDelegate struct {
	_TextDidBeginEditing func(notification foundation.foundation.INSNotification)
	_TextDidChange func(notification foundation.foundation.INSNotification)
	_TextDidEndEditing func(notification foundation.foundation.INSNotification)
	_TextShouldBeginEditing func(textObject IText) bool
	_TextShouldEndEditing func(textObject IText) bool
}

// SetTextDidBeginEditing sets the handler for the TextDidBeginEditing delegate method.
//
// Informs the delegate that the text object has begun editing (that the user has begun changing it).
func (d *TextDelegate) SetTextDidBeginEditing(f func(notification foundation.foundation.INSNotification)) {
	d._TextDidBeginEditing = f
}

// SetTextDidChange sets the handler for the TextDidChange delegate method.
//
// Informs the delegate that the text object has changed its characters or formatting attributes.
func (d *TextDelegate) SetTextDidChange(f func(notification foundation.foundation.INSNotification)) {
	d._TextDidChange = f
}

// SetTextDidEndEditing sets the handler for the TextDidEndEditing delegate method.
//
// Informs the delegate that the text object has finished editing (that it has resigned first responder status).
func (d *TextDelegate) SetTextDidEndEditing(f func(notification foundation.foundation.INSNotification)) {
	d._TextDidEndEditing = f
}

// SetTextShouldBeginEditing sets the handler for the TextShouldBeginEditing delegate method.
//
// Invoked when a text object begins to change its text, this method requests permission for   to begin editing.
func (d *TextDelegate) SetTextShouldBeginEditing(f func(textObject IText) bool) {
	d._TextShouldBeginEditing = f
}

// SetTextShouldEndEditing sets the handler for the TextShouldEndEditing delegate method.
//
// Invoked from a text object’s implementation of  , this method requests permission for   to end editing.
func (d *TextDelegate) SetTextShouldEndEditing(f func(textObject IText) bool) {
	d._TextShouldEndEditing = f
}

// TextDidBeginEditing implements the PTextDelegate interface.
func (d *TextDelegate) TextDidBeginEditing(notification foundation.foundation.INSNotification) {
	if d._TextDidBeginEditing != nil {
		d._TextDidBeginEditing(notification)
	}
}

// HasTextDidBeginEditing returns true if a handler for TextDidBeginEditing has been set.
func (d *TextDelegate) HasTextDidBeginEditing() bool {
	return d._TextDidBeginEditing != nil
}

// TextDidChange implements the PTextDelegate interface.
func (d *TextDelegate) TextDidChange(notification foundation.foundation.INSNotification) {
	if d._TextDidChange != nil {
		d._TextDidChange(notification)
	}
}

// HasTextDidChange returns true if a handler for TextDidChange has been set.
func (d *TextDelegate) HasTextDidChange() bool {
	return d._TextDidChange != nil
}

// TextDidEndEditing implements the PTextDelegate interface.
func (d *TextDelegate) TextDidEndEditing(notification foundation.foundation.INSNotification) {
	if d._TextDidEndEditing != nil {
		d._TextDidEndEditing(notification)
	}
}

// HasTextDidEndEditing returns true if a handler for TextDidEndEditing has been set.
func (d *TextDelegate) HasTextDidEndEditing() bool {
	return d._TextDidEndEditing != nil
}

// TextShouldBeginEditing implements the PTextDelegate interface.
func (d *TextDelegate) TextShouldBeginEditing(textObject IText) bool {
	if d._TextShouldBeginEditing != nil {
		return d._TextShouldBeginEditing(textObject)
	}
	var zero bool
	return zero
}

// HasTextShouldBeginEditing returns true if a handler for TextShouldBeginEditing has been set.
func (d *TextDelegate) HasTextShouldBeginEditing() bool {
	return d._TextShouldBeginEditing != nil
}

// TextShouldEndEditing implements the PTextDelegate interface.
func (d *TextDelegate) TextShouldEndEditing(textObject IText) bool {
	if d._TextShouldEndEditing != nil {
		return d._TextShouldEndEditing(textObject)
	}
	var zero bool
	return zero
}

// HasTextShouldEndEditing returns true if a handler for TextShouldEndEditing has been set.
func (d *TextDelegate) HasTextShouldEndEditing() bool {
	return d._TextShouldEndEditing != nil
}

// TextDelegateObject wraps an existing Objective-C object that conforms to the PTextDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TextDelegateObject struct {
	objectivec.Object
}

// NewTextDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTextDelegate protocol.
func NewTextDelegateObject(obj objectivec.Object) *TextDelegateObject {
	return &TextDelegateObject{obj}
}

// Make sure TextDelegateObject implements PTextDelegate.
var _ PTextDelegate = (*TextDelegateObject)(nil)

// TextDidBeginEditing implements the PTextDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextDelegateObject) TextDidBeginEditing(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("textDidBeginEditing:"), notification)
}

// HasTextDidBeginEditing returns true; this is a placeholder for optional method checks.
func (o *TextDelegateObject) HasTextDidBeginEditing() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextDidChange implements the PTextDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextDelegateObject) TextDidChange(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("textDidChange:"), notification)
}

// HasTextDidChange returns true; this is a placeholder for optional method checks.
func (o *TextDelegateObject) HasTextDidChange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextDidEndEditing implements the PTextDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextDelegateObject) TextDidEndEditing(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("textDidEndEditing:"), notification)
}

// HasTextDidEndEditing returns true; this is a placeholder for optional method checks.
func (o *TextDelegateObject) HasTextDidEndEditing() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextShouldBeginEditing implements the PTextDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextDelegateObject) TextShouldBeginEditing(textObject IText) bool {
	return objc.Send[bool](o.ID, objc.Sel("textShouldBeginEditing:"), textObject)
}

// HasTextShouldBeginEditing returns true; this is a placeholder for optional method checks.
func (o *TextDelegateObject) HasTextShouldBeginEditing() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextShouldEndEditing implements the PTextDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextDelegateObject) TextShouldEndEditing(textObject IText) bool {
	return objc.Send[bool](o.ID, objc.Sel("textShouldEndEditing:"), textObject)
}

// HasTextShouldEndEditing returns true; this is a placeholder for optional method checks.
func (o *TextDelegateObject) HasTextShouldEndEditing() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
