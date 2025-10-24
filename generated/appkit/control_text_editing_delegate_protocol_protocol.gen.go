// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PControlTextEditingDelegate is the NSControlTextEditingDelegate protocol interface.
//
// A set of optional methods implemented by delegates of   subclasses to respond to editing actions.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSControlTextEditingDelegate
type PControlTextEditingDelegate interface {
	// Optional methods
	ControlDidFailToFormatStringErrorDescription(control IControl, string_ objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: NSString */) bool
	HasControlDidFailToFormatStringErrorDescription() bool
	ControlDidFailToValidatePartialStringErrorDescription(control IControl, string_ objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: NSString */)
	HasControlDidFailToValidatePartialStringErrorDescription() bool
	ControlIsValidObject(control IControl, obj objc.IObject) bool
	HasControlIsValidObject() bool
	ControlTextShouldBeginEditing(control IControl, fieldEditor IText) bool
	HasControlTextShouldBeginEditing() bool
	ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange corefoundation.Range, index int) []string
	HasControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool
	ControlTextViewDoCommandBySelector(control IControl, textView ITextView, commandSelector objc.SEL) bool
	HasControlTextViewDoCommandBySelector() bool
	ControlTextShouldEndEditing(control IControl, fieldEditor IText) bool
	HasControlTextShouldEndEditing() bool
	ControlTextDidBeginEditing(obj foundation.Notification)
	HasControlTextDidBeginEditing() bool
	ControlTextDidChange(obj foundation.Notification)
	HasControlTextDidChange() bool
	ControlTextDidEndEditing(obj foundation.Notification)
	HasControlTextDidEndEditing() bool
}

// ControlTextEditingDelegate is a delegate implementation builder for the PControlTextEditingDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ControlTextEditingDelegate struct {
	_ControlDidFailToFormatStringErrorDescription func(control IControl, string_ objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: NSString */) bool
	_ControlDidFailToValidatePartialStringErrorDescription func(control IControl, string_ objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: NSString */)
	_ControlIsValidObject func(control IControl, obj objc.IObject) bool
	_ControlTextShouldBeginEditing func(control IControl, fieldEditor IText) bool
	_ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem func(control IControl, textView ITextView, words []string, charRange corefoundation.Range, index int) []string
	_ControlTextViewDoCommandBySelector func(control IControl, textView ITextView, commandSelector objc.SEL) bool
	_ControlTextShouldEndEditing func(control IControl, fieldEditor IText) bool
	_ControlTextDidBeginEditing func(obj foundation.Notification)
	_ControlTextDidChange func(obj foundation.Notification)
	_ControlTextDidEndEditing func(obj foundation.Notification)
}

// SetControlDidFailToFormatStringErrorDescription sets the handler for the ControlDidFailToFormatStringErrorDescription delegate method.
//
// Invoked when the formatter for the cell belonging to the specified control cannot convert a string to an underlying object.
func (d *ControlTextEditingDelegate) SetControlDidFailToFormatStringErrorDescription(f func(control IControl, string_ objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: NSString */) bool) {
	d._ControlDidFailToFormatStringErrorDescription = f
}

// SetControlDidFailToValidatePartialStringErrorDescription sets the handler for the ControlDidFailToValidatePartialStringErrorDescription delegate method.
//
// Invoked when the formatter for the cell belonging to   (or selected cell) rejects a partial string a user is typing into the cell.
func (d *ControlTextEditingDelegate) SetControlDidFailToValidatePartialStringErrorDescription(f func(control IControl, string_ objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: NSString */)) {
	d._ControlDidFailToValidatePartialStringErrorDescription = f
}

// SetControlIsValidObject sets the handler for the ControlIsValidObject delegate method.
//
// Invoked when the insertion point leaves a cell belonging to the specified control, but before the value of the cell’s object is displayed.
func (d *ControlTextEditingDelegate) SetControlIsValidObject(f func(control IControl, obj objc.IObject) bool) {
	d._ControlIsValidObject = f
}

// SetControlTextShouldBeginEditing sets the handler for the ControlTextShouldBeginEditing delegate method.
//
// Invoked when the user tries to enter a character in a cell of a control that allows editing of text (such as a text field or form field).
func (d *ControlTextEditingDelegate) SetControlTextShouldBeginEditing(f func(control IControl, fieldEditor IText) bool) {
	d._ControlTextShouldBeginEditing = f
}

// SetControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem sets the handler for the ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem delegate method.
//
// Invoked to allow you to control the list of proposed text completions generated by text fields and other controls.
func (d *ControlTextEditingDelegate) SetControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(f func(control IControl, textView ITextView, words []string, charRange corefoundation.Range, index int) []string) {
	d._ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem = f
}

// SetControlTextViewDoCommandBySelector sets the handler for the ControlTextViewDoCommandBySelector delegate method.
//
// Invoked when users press keys with predefined bindings in a cell of the specified control.
func (d *ControlTextEditingDelegate) SetControlTextViewDoCommandBySelector(f func(control IControl, textView ITextView, commandSelector objc.SEL) bool) {
	d._ControlTextViewDoCommandBySelector = f
}

// SetControlTextShouldEndEditing sets the handler for the ControlTextShouldEndEditing delegate method.
//
// Invoked when the insertion point tries to leave a cell of the control that has been edited.
func (d *ControlTextEditingDelegate) SetControlTextShouldEndEditing(f func(control IControl, fieldEditor IText) bool) {
	d._ControlTextShouldEndEditing = f
}

// SetControlTextDidBeginEditing sets the handler for the ControlTextDidBeginEditing delegate method.
//
// Tells the delegate that the control started editing its text content.
func (d *ControlTextEditingDelegate) SetControlTextDidBeginEditing(f func(obj foundation.Notification)) {
	d._ControlTextDidBeginEditing = f
}

// SetControlTextDidChange sets the handler for the ControlTextDidChange delegate method.
//
// Tells the delegate that the control made changes to its text content.
func (d *ControlTextEditingDelegate) SetControlTextDidChange(f func(obj foundation.Notification)) {
	d._ControlTextDidChange = f
}

// SetControlTextDidEndEditing sets the handler for the ControlTextDidEndEditing delegate method.
//
// Tells the delegate that the control finished editing its text content and committed the changes.
func (d *ControlTextEditingDelegate) SetControlTextDidEndEditing(f func(obj foundation.Notification)) {
	d._ControlTextDidEndEditing = f
}

// ControlDidFailToFormatStringErrorDescription implements the PControlTextEditingDelegate interface.
func (d *ControlTextEditingDelegate) ControlDidFailToFormatStringErrorDescription(control IControl, string_ objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: NSString */) bool {
	if d._ControlDidFailToFormatStringErrorDescription != nil {
		return d._ControlDidFailToFormatStringErrorDescription(control, string_, error_)
	}
	var zero bool
	return zero
}

// HasControlDidFailToFormatStringErrorDescription returns true if a handler for ControlDidFailToFormatStringErrorDescription has been set.
func (d *ControlTextEditingDelegate) HasControlDidFailToFormatStringErrorDescription() bool {
	return d._ControlDidFailToFormatStringErrorDescription != nil
}

// ControlDidFailToValidatePartialStringErrorDescription implements the PControlTextEditingDelegate interface.
func (d *ControlTextEditingDelegate) ControlDidFailToValidatePartialStringErrorDescription(control IControl, string_ objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: NSString */) {
	if d._ControlDidFailToValidatePartialStringErrorDescription != nil {
		d._ControlDidFailToValidatePartialStringErrorDescription(control, string_, error_)
	}
}

// HasControlDidFailToValidatePartialStringErrorDescription returns true if a handler for ControlDidFailToValidatePartialStringErrorDescription has been set.
func (d *ControlTextEditingDelegate) HasControlDidFailToValidatePartialStringErrorDescription() bool {
	return d._ControlDidFailToValidatePartialStringErrorDescription != nil
}

// ControlIsValidObject implements the PControlTextEditingDelegate interface.
func (d *ControlTextEditingDelegate) ControlIsValidObject(control IControl, obj objc.IObject) bool {
	if d._ControlIsValidObject != nil {
		return d._ControlIsValidObject(control, obj)
	}
	var zero bool
	return zero
}

// HasControlIsValidObject returns true if a handler for ControlIsValidObject has been set.
func (d *ControlTextEditingDelegate) HasControlIsValidObject() bool {
	return d._ControlIsValidObject != nil
}

// ControlTextShouldBeginEditing implements the PControlTextEditingDelegate interface.
func (d *ControlTextEditingDelegate) ControlTextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	if d._ControlTextShouldBeginEditing != nil {
		return d._ControlTextShouldBeginEditing(control, fieldEditor)
	}
	var zero bool
	return zero
}

// HasControlTextShouldBeginEditing returns true if a handler for ControlTextShouldBeginEditing has been set.
func (d *ControlTextEditingDelegate) HasControlTextShouldBeginEditing() bool {
	return d._ControlTextShouldBeginEditing != nil
}

// ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem implements the PControlTextEditingDelegate interface.
func (d *ControlTextEditingDelegate) ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange corefoundation.Range, index int) []string {
	if d._ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem != nil {
		return d._ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control, textView, words, charRange, index)
	}
	var zero []string
	return zero
}

// HasControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem returns true if a handler for ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem has been set.
func (d *ControlTextEditingDelegate) HasControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool {
	return d._ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem != nil
}

// ControlTextViewDoCommandBySelector implements the PControlTextEditingDelegate interface.
func (d *ControlTextEditingDelegate) ControlTextViewDoCommandBySelector(control IControl, textView ITextView, commandSelector objc.SEL) bool {
	if d._ControlTextViewDoCommandBySelector != nil {
		return d._ControlTextViewDoCommandBySelector(control, textView, commandSelector)
	}
	var zero bool
	return zero
}

// HasControlTextViewDoCommandBySelector returns true if a handler for ControlTextViewDoCommandBySelector has been set.
func (d *ControlTextEditingDelegate) HasControlTextViewDoCommandBySelector() bool {
	return d._ControlTextViewDoCommandBySelector != nil
}

// ControlTextShouldEndEditing implements the PControlTextEditingDelegate interface.
func (d *ControlTextEditingDelegate) ControlTextShouldEndEditing(control IControl, fieldEditor IText) bool {
	if d._ControlTextShouldEndEditing != nil {
		return d._ControlTextShouldEndEditing(control, fieldEditor)
	}
	var zero bool
	return zero
}

// HasControlTextShouldEndEditing returns true if a handler for ControlTextShouldEndEditing has been set.
func (d *ControlTextEditingDelegate) HasControlTextShouldEndEditing() bool {
	return d._ControlTextShouldEndEditing != nil
}

// ControlTextDidBeginEditing implements the PControlTextEditingDelegate interface.
func (d *ControlTextEditingDelegate) ControlTextDidBeginEditing(obj foundation.Notification) {
	if d._ControlTextDidBeginEditing != nil {
		d._ControlTextDidBeginEditing(obj)
	}
}

// HasControlTextDidBeginEditing returns true if a handler for ControlTextDidBeginEditing has been set.
func (d *ControlTextEditingDelegate) HasControlTextDidBeginEditing() bool {
	return d._ControlTextDidBeginEditing != nil
}

// ControlTextDidChange implements the PControlTextEditingDelegate interface.
func (d *ControlTextEditingDelegate) ControlTextDidChange(obj foundation.Notification) {
	if d._ControlTextDidChange != nil {
		d._ControlTextDidChange(obj)
	}
}

// HasControlTextDidChange returns true if a handler for ControlTextDidChange has been set.
func (d *ControlTextEditingDelegate) HasControlTextDidChange() bool {
	return d._ControlTextDidChange != nil
}

// ControlTextDidEndEditing implements the PControlTextEditingDelegate interface.
func (d *ControlTextEditingDelegate) ControlTextDidEndEditing(obj foundation.Notification) {
	if d._ControlTextDidEndEditing != nil {
		d._ControlTextDidEndEditing(obj)
	}
}

// HasControlTextDidEndEditing returns true if a handler for ControlTextDidEndEditing has been set.
func (d *ControlTextEditingDelegate) HasControlTextDidEndEditing() bool {
	return d._ControlTextDidEndEditing != nil
}
