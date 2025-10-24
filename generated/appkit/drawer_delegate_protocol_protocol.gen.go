// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
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
	DrawerDidClose(notification foundation.Notification)
	HasDrawerDidClose() bool
	DrawerDidOpen(notification foundation.Notification)
	HasDrawerDidOpen() bool
	DrawerShouldClose(sender IDrawer) bool
	HasDrawerShouldClose() bool
	DrawerShouldOpen(sender IDrawer) bool
	HasDrawerShouldOpen() bool
	DrawerWillClose(notification foundation.Notification)
	HasDrawerWillClose() bool
	DrawerWillOpen(notification foundation.Notification)
	HasDrawerWillOpen() bool
	DrawerWillResizeContentsToSize(sender IDrawer, contentSize Size /* not a class type */) Size
	HasDrawerWillResizeContentsToSize() bool
}

// DrawerDelegate is a delegate implementation builder for the PDrawerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type DrawerDelegate struct {
	_DrawerDidClose func(notification foundation.Notification)
	_DrawerDidOpen func(notification foundation.Notification)
	_DrawerShouldClose func(sender IDrawer) bool
	_DrawerShouldOpen func(sender IDrawer) bool
	_DrawerWillClose func(notification foundation.Notification)
	_DrawerWillOpen func(notification foundation.Notification)
	_DrawerWillResizeContentsToSize func(sender IDrawer, contentSize Size /* not a class type */) Size
}

// SetDrawerDidClose sets the handler for the DrawerDidClose delegate method.
//
// Notifies the delegate that the drawer has closed.
func (d *DrawerDelegate) SetDrawerDidClose(f func(notification foundation.Notification)) {
	d._DrawerDidClose = f
}

// SetDrawerDidOpen sets the handler for the DrawerDidOpen delegate method.
//
// Notifies the delegate that the drawer has opened.
func (d *DrawerDelegate) SetDrawerDidOpen(f func(notification foundation.Notification)) {
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
func (d *DrawerDelegate) SetDrawerWillClose(f func(notification foundation.Notification)) {
	d._DrawerWillClose = f
}

// SetDrawerWillOpen sets the handler for the DrawerWillOpen delegate method.
//
// Notifies the delegate that the drawer will open.
func (d *DrawerDelegate) SetDrawerWillOpen(f func(notification foundation.Notification)) {
	d._DrawerWillOpen = f
}

// SetDrawerWillResizeContentsToSize sets the handler for the DrawerWillResizeContentsToSize delegate method.
//
// Invoked when the user resizes the drawer or parent.
func (d *DrawerDelegate) SetDrawerWillResizeContentsToSize(f func(sender IDrawer, contentSize Size /* not a class type */) Size) {
	d._DrawerWillResizeContentsToSize = f
}

// DrawerDidClose implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerDidClose(notification foundation.Notification) {
	if d._DrawerDidClose != nil {
		d._DrawerDidClose(notification)
	}
}

// HasDrawerDidClose returns true if a handler for DrawerDidClose has been set.
func (d *DrawerDelegate) HasDrawerDidClose() bool {
	return d._DrawerDidClose != nil
}

// DrawerDidOpen implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerDidOpen(notification foundation.Notification) {
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
func (d *DrawerDelegate) DrawerWillClose(notification foundation.Notification) {
	if d._DrawerWillClose != nil {
		d._DrawerWillClose(notification)
	}
}

// HasDrawerWillClose returns true if a handler for DrawerWillClose has been set.
func (d *DrawerDelegate) HasDrawerWillClose() bool {
	return d._DrawerWillClose != nil
}

// DrawerWillOpen implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerWillOpen(notification foundation.Notification) {
	if d._DrawerWillOpen != nil {
		d._DrawerWillOpen(notification)
	}
}

// HasDrawerWillOpen returns true if a handler for DrawerWillOpen has been set.
func (d *DrawerDelegate) HasDrawerWillOpen() bool {
	return d._DrawerWillOpen != nil
}

// DrawerWillResizeContentsToSize implements the PDrawerDelegate interface.
func (d *DrawerDelegate) DrawerWillResizeContentsToSize(sender IDrawer, contentSize Size /* not a class type */) Size {
	if d._DrawerWillResizeContentsToSize != nil {
		return d._DrawerWillResizeContentsToSize(sender, contentSize)
	}
	var zero Size
	return zero
}

// HasDrawerWillResizeContentsToSize returns true if a handler for DrawerWillResizeContentsToSize has been set.
func (d *DrawerDelegate) HasDrawerWillResizeContentsToSize() bool {
	return d._DrawerWillResizeContentsToSize != nil
}
