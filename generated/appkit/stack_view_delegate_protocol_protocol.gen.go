// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PStackViewDelegate is the NSStackViewDelegate protocol interface.
//
// A set of methods you use to respond to a stack view detaching and reattaching views.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSStackViewDelegate
type PStackViewDelegate interface {
	// Optional methods
	StackViewDidReattachViews(stackView IStackView, views []View)
	HasStackViewDidReattachViews() bool
	StackViewWillDetachViews(stackView IStackView, views []View)
	HasStackViewWillDetachViews() bool
}

// StackViewDelegate is a delegate implementation builder for the PStackViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type StackViewDelegate struct {
	_StackViewDidReattachViews func(stackView IStackView, views []View)
	_StackViewWillDetachViews func(stackView IStackView, views []View)
}

// SetStackViewDidReattachViews sets the handler for the StackViewDidReattachViews delegate method.
//
// Called when the stack view has automatically reattached one or more previously-detached views.
func (d *StackViewDelegate) SetStackViewDidReattachViews(f func(stackView IStackView, views []View)) {
	d._StackViewDidReattachViews = f
}

// SetStackViewWillDetachViews sets the handler for the StackViewWillDetachViews delegate method.
//
// Called when the stack view is about to automatically detach one or more of its views.
func (d *StackViewDelegate) SetStackViewWillDetachViews(f func(stackView IStackView, views []View)) {
	d._StackViewWillDetachViews = f
}

// StackViewDidReattachViews implements the PStackViewDelegate interface.
func (d *StackViewDelegate) StackViewDidReattachViews(stackView IStackView, views []View) {
	if d._StackViewDidReattachViews != nil {
		d._StackViewDidReattachViews(stackView, views)
	}
}

// HasStackViewDidReattachViews returns true if a handler for StackViewDidReattachViews has been set.
func (d *StackViewDelegate) HasStackViewDidReattachViews() bool {
	return d._StackViewDidReattachViews != nil
}

// StackViewWillDetachViews implements the PStackViewDelegate interface.
func (d *StackViewDelegate) StackViewWillDetachViews(stackView IStackView, views []View) {
	if d._StackViewWillDetachViews != nil {
		d._StackViewWillDetachViews(stackView, views)
	}
}

// HasStackViewWillDetachViews returns true if a handler for StackViewWillDetachViews has been set.
func (d *StackViewDelegate) HasStackViewWillDetachViews() bool {
	return d._StackViewWillDetachViews != nil
}

// StackViewDelegateObject wraps an existing Objective-C object that conforms to the PStackViewDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type StackViewDelegateObject struct {
	objectivec.Object
}

// NewStackViewDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSStackViewDelegate protocol.
func NewStackViewDelegateObject(obj objectivec.Object) *StackViewDelegateObject {
	return &StackViewDelegateObject{obj}
}

// Make sure StackViewDelegateObject implements PStackViewDelegate.
var _ PStackViewDelegate = (*StackViewDelegateObject)(nil)

// StackViewDidReattachViews implements the PStackViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *StackViewDelegateObject) StackViewDidReattachViews(stackView IStackView, views []View) {
	objc.Send[objc.ID](o.ID, objc.Sel("stackView:didReattachViews:"), stackView, views)
}

// HasStackViewDidReattachViews returns true; this is a placeholder for optional method checks.
func (o *StackViewDelegateObject) HasStackViewDidReattachViews() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// StackViewWillDetachViews implements the PStackViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *StackViewDelegateObject) StackViewWillDetachViews(stackView IStackView, views []View) {
	objc.Send[objc.ID](o.ID, objc.Sel("stackView:willDetachViews:"), stackView, views)
}

// HasStackViewWillDetachViews returns true; this is a placeholder for optional method checks.
func (o *StackViewDelegateObject) HasStackViewWillDetachViews() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
