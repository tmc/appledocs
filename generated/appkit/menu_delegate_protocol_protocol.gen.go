// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	MenuHasKeyEquivalentForEventTargetAction(menu IMenu, event IEvent, target unsafe.Pointer, action unsafe.Pointer) bool
	HasMenuHasKeyEquivalentForEventTargetAction() bool
}

// MenuDelegate is a delegate implementation builder for the PMenuDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MenuDelegate struct {
	_MenuHasKeyEquivalentForEventTargetAction func(menu IMenu, event IEvent, target unsafe.Pointer, action unsafe.Pointer) bool
}

// SetMenuHasKeyEquivalentForEventTargetAction sets the handler for the MenuHasKeyEquivalentForEventTargetAction delegate method.
//
// Invoked to allow the delegate to return the target and action for a key-down event.
func (d *MenuDelegate) SetMenuHasKeyEquivalentForEventTargetAction(f func(menu IMenu, event IEvent, target unsafe.Pointer, action unsafe.Pointer) bool) {
	d._MenuHasKeyEquivalentForEventTargetAction = f
}

// MenuHasKeyEquivalentForEventTargetAction implements the PMenuDelegate interface.
func (d *MenuDelegate) MenuHasKeyEquivalentForEventTargetAction(menu IMenu, event IEvent, target unsafe.Pointer, action unsafe.Pointer) bool {
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
