// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PDrawerDelegate is the NSDrawerDelegate protocol interface.
//
// A set of methods that drawer delegates implement to open, close, and resize the drawer.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSDrawerDelegate
type PDrawerDelegate interface {
	// Optional methods
	DrawerDidClose(notification foundation.foundation.INSNotification)
	HasDrawerDidClose() bool
	DrawerDidOpen(notification foundation.foundation.INSNotification)
	HasDrawerDidOpen() bool
	DrawerShouldClose(sender IDrawer) bool
	HasDrawerShouldClose() bool
	DrawerShouldOpen(sender IDrawer) bool
	HasDrawerShouldOpen() bool
	DrawerWillClose(notification foundation.foundation.INSNotification)
	HasDrawerWillClose() bool
	DrawerWillOpen(notification foundation.foundation.INSNotification)
	HasDrawerWillOpen() bool
	DrawerWillResizeContentsToSize(sender IDrawer, contentSize corefoundation.CGSize) corefoundation.CGSize
	HasDrawerWillResizeContentsToSize() bool
}

// DrawerDelegate is a delegate implementation builder for the PDrawerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type DrawerDelegate struct {
	_DrawerDidClose func(notification foundation.foundation.INSNotification)
	_DrawerDidOpen func(notification foundation.foundation.INSNotification)
	_DrawerShouldClose func(sender IDrawer) bool
	_DrawerShouldOpen func(sender IDrawer) bool
	_DrawerWillClose func(notification foundation.foundation.INSNotification)
	_DrawerWillOpen func(notification foundation.foundation.INSNotification)
	_DrawerWillResizeContentsToSize func(sender IDrawer, contentSize corefoundation.CGSize) corefoundation.CGSize
}

// SetDrawerDidClose sets the handler for the DrawerDidClose delegate method.
//
// Notifies the delegate that the drawer has closed.
func (d *DrawerDelegate) SetDrawerDidClose(f func(notification foundation.foundation.INSNotification)) {
	d._DrawerDidClose = f
}

// SetDrawerDidOpen sets the handler for the DrawerDidOpen delegate method.
//
// Notifies the delegate that the drawer has opened.
func (d *DrawerDelegate) SetDrawerDidOpen(f func(notification foundation.foundation.INSNotification)) {
	d._DrawerDidOpen = f
}

// SetDrawerShouldClose sets the handler for the DrawerShouldClose delegate method.
//
// Asks the delegate if the specified drawer should close.
func (d *DrawerDelegate) SetDrawerShouldClose(f func(sender IDrawer) bool) {
	d._DrawerShouldClose = f
}

// SetDrawerShouldOpen sets the handler for the DrawerShouldOpen delegate method.
//
// Asks the delegate if the specified drawer should open.
func (d *DrawerDelegate) SetDrawerShouldOpen(f func(sender IDrawer) bool) {
	d._DrawerShouldOpen = f
}

// SetDrawerWillClose sets the handler for the DrawerWillClose delegate method.
//
// Notifies the delegate the drawer will close.
func (d *DrawerDelegate) SetDrawerWillClose(f func(notification foundation.foundation.INSNotification)) {
	d._DrawerWillClose = f
}

// SetDrawerWillOpen sets the handler for the DrawerWillOpen delegate method.
//
// Notifies the delegate that the drawer will open.
func (d *DrawerDelegate) SetDrawerWillOpen(f func(notification foundation.foundation.INSNotification)) {
	d._DrawerWillOpen = f
}

// SetDrawerWillResizeContentsToSize sets the handler for the DrawerWillResizeContentsToSize delegate method.
//
// Invoked when the user resizes the drawer or parent.
func (d *DrawerDelegate) SetDrawerWillResizeContentsToSize(f func(sender IDrawer, contentSize corefoundation.CGSize) corefoundation.CGSize) {
	d._DrawerWillResizeContentsToSize = f
}

// DrawerDidClose implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerDidClose(notification foundation.foundation.INSNotification) {
	if d._DrawerDidClose != nil {
		d._DrawerDidClose(notification)
	}
}

// HasDrawerDidClose returns true if a handler for DrawerDidClose has been set.
func (d *DrawerDelegate) HasDrawerDidClose() bool {
	return d._DrawerDidClose != nil
}

// DrawerDidOpen implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerDidOpen(notification foundation.foundation.INSNotification) {
	if d._DrawerDidOpen != nil {
		d._DrawerDidOpen(notification)
	}
}

// HasDrawerDidOpen returns true if a handler for DrawerDidOpen has been set.
func (d *DrawerDelegate) HasDrawerDidOpen() bool {
	return d._DrawerDidOpen != nil
}

// DrawerShouldClose implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerShouldClose(sender IDrawer) bool {
	if d._DrawerShouldClose != nil {
		return d._DrawerShouldClose(sender)
	}
	var zero bool
	return zero
}

// HasDrawerShouldClose returns true if a handler for DrawerShouldClose has been set.
func (d *DrawerDelegate) HasDrawerShouldClose() bool {
	return d._DrawerShouldClose != nil
}

// DrawerShouldOpen implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerShouldOpen(sender IDrawer) bool {
	if d._DrawerShouldOpen != nil {
		return d._DrawerShouldOpen(sender)
	}
	var zero bool
	return zero
}

// HasDrawerShouldOpen returns true if a handler for DrawerShouldOpen has been set.
func (d *DrawerDelegate) HasDrawerShouldOpen() bool {
	return d._DrawerShouldOpen != nil
}

// DrawerWillClose implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerWillClose(notification foundation.foundation.INSNotification) {
	if d._DrawerWillClose != nil {
		d._DrawerWillClose(notification)
	}
}

// HasDrawerWillClose returns true if a handler for DrawerWillClose has been set.
func (d *DrawerDelegate) HasDrawerWillClose() bool {
	return d._DrawerWillClose != nil
}

// DrawerWillOpen implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerWillOpen(notification foundation.foundation.INSNotification) {
	if d._DrawerWillOpen != nil {
		d._DrawerWillOpen(notification)
	}
}

// HasDrawerWillOpen returns true if a handler for DrawerWillOpen has been set.
func (d *DrawerDelegate) HasDrawerWillOpen() bool {
	return d._DrawerWillOpen != nil
}

// DrawerWillResizeContentsToSize implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerWillResizeContentsToSize(sender IDrawer, contentSize corefoundation.CGSize) corefoundation.CGSize {
	if d._DrawerWillResizeContentsToSize != nil {
		return d._DrawerWillResizeContentsToSize(sender, contentSize)
	}
	var zero corefoundation.CGSize
	return zero
}

// HasDrawerWillResizeContentsToSize returns true if a handler for DrawerWillResizeContentsToSize has been set.
func (d *DrawerDelegate) HasDrawerWillResizeContentsToSize() bool {
	return d._DrawerWillResizeContentsToSize != nil
}

// DrawerDelegateObject wraps an existing Objective-C object that conforms to the PDrawerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type DrawerDelegateObject struct {
	objectivec.Object
}

// NewDrawerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSDrawerDelegate protocol.
func NewDrawerDelegateObject(obj objectivec.Object) *DrawerDelegateObject {
	return &DrawerDelegateObject{obj}
}

// Make sure DrawerDelegateObject implements PDrawerDelegate.
var _ PDrawerDelegate = (*DrawerDelegateObject)(nil)

// DrawerDidClose implements the PDrawerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *DrawerDelegateObject) DrawerDidClose(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("drawerDidClose:"), notification)
}

// HasDrawerDidClose returns true; this is a placeholder for optional method checks.
func (o *DrawerDelegateObject) HasDrawerDidClose() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DrawerDidOpen implements the PDrawerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *DrawerDelegateObject) DrawerDidOpen(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("drawerDidOpen:"), notification)
}

// HasDrawerDidOpen returns true; this is a placeholder for optional method checks.
func (o *DrawerDelegateObject) HasDrawerDidOpen() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DrawerShouldClose implements the PDrawerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *DrawerDelegateObject) DrawerShouldClose(sender IDrawer) bool {
	return objc.Send[bool](o.ID, objc.Sel("drawerShouldClose:"), sender)
}

// HasDrawerShouldClose returns true; this is a placeholder for optional method checks.
func (o *DrawerDelegateObject) HasDrawerShouldClose() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DrawerShouldOpen implements the PDrawerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *DrawerDelegateObject) DrawerShouldOpen(sender IDrawer) bool {
	return objc.Send[bool](o.ID, objc.Sel("drawerShouldOpen:"), sender)
}

// HasDrawerShouldOpen returns true; this is a placeholder for optional method checks.
func (o *DrawerDelegateObject) HasDrawerShouldOpen() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DrawerWillClose implements the PDrawerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *DrawerDelegateObject) DrawerWillClose(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("drawerWillClose:"), notification)
}

// HasDrawerWillClose returns true; this is a placeholder for optional method checks.
func (o *DrawerDelegateObject) HasDrawerWillClose() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DrawerWillOpen implements the PDrawerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *DrawerDelegateObject) DrawerWillOpen(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("drawerWillOpen:"), notification)
}

// HasDrawerWillOpen returns true; this is a placeholder for optional method checks.
func (o *DrawerDelegateObject) HasDrawerWillOpen() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DrawerWillResizeContentsToSize implements the PDrawerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *DrawerDelegateObject) DrawerWillResizeContentsToSize(sender IDrawer, contentSize corefoundation.CGSize) corefoundation.CGSize {
	return objc.Send[corefoundation.CGSize](o.ID, objc.Sel("drawerWillResizeContents:toSize:"), sender, contentSize)
}

// HasDrawerWillResizeContentsToSize returns true; this is a placeholder for optional method checks.
func (o *DrawerDelegateObject) HasDrawerWillResizeContentsToSize() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
