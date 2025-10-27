// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PMenuDelegate is the NSMenuDelegate protocol interface.
//
// The optional methods implemented by delegates of   objects to manage menu display and handle some events.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSMenuDelegate
type PMenuDelegate interface {
	// Optional methods
	MenuHasKeyEquivalentForEventTargetAction(menu IMenu, event IEvent, target objectivec.IObject, action objectivec.IObject) bool
	HasMenuHasKeyEquivalentForEventTargetAction() bool
}

// MenuDelegate is a delegate implementation builder for the PMenuDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MenuDelegate struct {
	_MenuHasKeyEquivalentForEventTargetAction func(menu IMenu, event IEvent, target objectivec.IObject, action objectivec.IObject) bool
}

// SetMenuHasKeyEquivalentForEventTargetAction sets the handler for the MenuHasKeyEquivalentForEventTargetAction delegate method.
//
// Invoked to allow the delegate to return the target and action for a key-down event.
func (d *MenuDelegate) SetMenuHasKeyEquivalentForEventTargetAction(f func(menu IMenu, event IEvent, target objectivec.IObject, action objectivec.IObject) bool) {
	d._MenuHasKeyEquivalentForEventTargetAction = f
}

// MenuHasKeyEquivalentForEventTargetAction implements the PMenuDelegate interface.
func (d *MenuDelegate) MenuHasKeyEquivalentForEventTargetAction(menu IMenu, event IEvent, target objectivec.IObject, action objectivec.IObject) bool {
	if d._MenuHasKeyEquivalentForEventTargetAction != nil {
		return d._MenuHasKeyEquivalentForEventTargetAction(menu, event, target, action)
	}
	var zero bool
	return zero
}

// HasMenuHasKeyEquivalentForEventTargetAction returns true if a handler for MenuHasKeyEquivalentForEventTargetAction has been set.
func (d *MenuDelegate) HasMenuHasKeyEquivalentForEventTargetAction() bool {
	return d._MenuHasKeyEquivalentForEventTargetAction != nil
}

// MenuDelegateObject wraps an existing Objective-C object that conforms to the PMenuDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type MenuDelegateObject struct {
	objectivec.Object
}

// NewMenuDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSMenuDelegate protocol.
func NewMenuDelegateObject(obj objectivec.Object) *MenuDelegateObject {
	return &MenuDelegateObject{obj}
}

// Make sure MenuDelegateObject implements PMenuDelegate.
var _ PMenuDelegate = (*MenuDelegateObject)(nil)

// MenuHasKeyEquivalentForEventTargetAction implements the PMenuDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *MenuDelegateObject) MenuHasKeyEquivalentForEventTargetAction(menu IMenu, event IEvent, target objectivec.IObject, action objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("menuHasKeyEquivalent:forEvent:target:action:"), menu, event, target, action)
}

// HasMenuHasKeyEquivalentForEventTargetAction returns true; this is a placeholder for optional method checks.
func (o *MenuDelegateObject) HasMenuHasKeyEquivalentForEventTargetAction() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
