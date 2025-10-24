// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PWindowDelegate is the NSWindowDelegate protocol interface.
//
// A set of optional methods that a window’s delegate can implement to respond to events, such as window resizing, moving, exposing, and minimizing.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSWindowDelegate
type PWindowDelegate interface {
	// Optional methods
	WindowShouldDragDocumentWithEventFromWithPasteboard(window IWindow, event IEvent, dragImageLocation objc.IObject /* cross-framework: Point */, pasteboard IPasteboard) bool
	HasWindowShouldDragDocumentWithEventFromWithPasteboard() bool
	WindowShouldPopUpDocumentPathMenu(window IWindow, menu IMenu) bool
	HasWindowShouldPopUpDocumentPathMenu() bool
	WindowWillUseFullScreenContentSize(window IWindow, proposedSize objc.IObject /* cross-framework: Size */) corefoundation.Size
	HasWindowWillUseFullScreenContentSize() bool
	WindowWillUseFullScreenPresentationOptions(window IWindow, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions
	HasWindowWillUseFullScreenPresentationOptions() bool
	WindowDidBecomeKey(notification foundation.Notification)
	HasWindowDidBecomeKey() bool
	WindowDidBecomeMain(notification foundation.Notification)
	HasWindowDidBecomeMain() bool
	WindowDidEndLiveResize(notification foundation.Notification)
	HasWindowDidEndLiveResize() bool
	WindowDidEnterFullScreen(notification foundation.Notification)
	HasWindowDidEnterFullScreen() bool
	WindowDidExitFullScreen(notification foundation.Notification)
	HasWindowDidExitFullScreen() bool
	WindowDidResignKey(notification foundation.Notification)
	HasWindowDidResignKey() bool
	WindowDidResignMain(notification foundation.Notification)
	HasWindowDidResignMain() bool
	WindowDidResize(notification foundation.Notification)
	HasWindowDidResize() bool
	WindowShouldClose(sender IWindow) bool
	HasWindowShouldClose() bool
	WindowShouldZoomToFrame(window IWindow, newFrame objc.IObject /* cross-framework: Rect */) bool
	HasWindowShouldZoomToFrame() bool
	WindowWillClose(notification foundation.Notification)
	HasWindowWillClose() bool
	WindowWillEnterFullScreen(notification foundation.Notification)
	HasWindowWillEnterFullScreen() bool
	WindowWillExitFullScreen(notification foundation.Notification)
	HasWindowWillExitFullScreen() bool
	WindowWillResizeToSize(sender IWindow, frameSize objc.IObject /* cross-framework: Size */) corefoundation.Size
	HasWindowWillResizeToSize() bool
	WindowWillReturnFieldEditorToObject(sender IWindow, client objc.IObject) objc.ID
	HasWindowWillReturnFieldEditorToObject() bool
	WindowWillStartLiveResize(notification foundation.Notification)
	HasWindowWillStartLiveResize() bool
	WindowWillUseStandardFrameDefaultFrame(window IWindow, newFrame objc.IObject /* cross-framework: Rect */) corefoundation.Rect
	HasWindowWillUseStandardFrameDefaultFrame() bool
}

// WindowDelegate is a delegate implementation builder for the PWindowDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WindowDelegate struct {
	_WindowShouldDragDocumentWithEventFromWithPasteboard func(window IWindow, event IEvent, dragImageLocation objc.IObject /* cross-framework: Point */, pasteboard IPasteboard) bool
	_WindowShouldPopUpDocumentPathMenu func(window IWindow, menu IMenu) bool
	_WindowWillUseFullScreenContentSize func(window IWindow, proposedSize objc.IObject /* cross-framework: Size */) corefoundation.Size
	_WindowWillUseFullScreenPresentationOptions func(window IWindow, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions
	_WindowDidBecomeKey func(notification foundation.Notification)
	_WindowDidBecomeMain func(notification foundation.Notification)
	_WindowDidEndLiveResize func(notification foundation.Notification)
	_WindowDidEnterFullScreen func(notification foundation.Notification)
	_WindowDidExitFullScreen func(notification foundation.Notification)
	_WindowDidResignKey func(notification foundation.Notification)
	_WindowDidResignMain func(notification foundation.Notification)
	_WindowDidResize func(notification foundation.Notification)
	_WindowShouldClose func(sender IWindow) bool
	_WindowShouldZoomToFrame func(window IWindow, newFrame objc.IObject /* cross-framework: Rect */) bool
	_WindowWillClose func(notification foundation.Notification)
	_WindowWillEnterFullScreen func(notification foundation.Notification)
	_WindowWillExitFullScreen func(notification foundation.Notification)
	_WindowWillResizeToSize func(sender IWindow, frameSize objc.IObject /* cross-framework: Size */) corefoundation.Size
	_WindowWillReturnFieldEditorToObject func(sender IWindow, client objc.IObject) objc.ID
	_WindowWillStartLiveResize func(notification foundation.Notification)
	_WindowWillUseStandardFrameDefaultFrame func(window IWindow, newFrame objc.IObject /* cross-framework: Rect */) corefoundation.Rect
}

// SetWindowShouldDragDocumentWithEventFromWithPasteboard sets the handler for the WindowShouldDragDocumentWithEventFromWithPasteboard delegate method.
//
// Asks the delegate whether a user can drag the document icon from the window’s title bar.
func (d *WindowDelegate) SetWindowShouldDragDocumentWithEventFromWithPasteboard(f func(window IWindow, event IEvent, dragImageLocation objc.IObject /* cross-framework: Point */, pasteboard IPasteboard) bool) {
	d._WindowShouldDragDocumentWithEventFromWithPasteboard = f
}

// SetWindowShouldPopUpDocumentPathMenu sets the handler for the WindowShouldPopUpDocumentPathMenu delegate method.
//
// Asks the delegate whether the window displays the title pop-up menu in response to a Command-click or Control-click on its title.
func (d *WindowDelegate) SetWindowShouldPopUpDocumentPathMenu(f func(window IWindow, menu IMenu) bool) {
	d._WindowShouldPopUpDocumentPathMenu = f
}

// SetWindowWillUseFullScreenContentSize sets the handler for the WindowWillUseFullScreenContentSize delegate method.
//
// Called to allow the delegate to modify the full-screen content size.
func (d *WindowDelegate) SetWindowWillUseFullScreenContentSize(f func(window IWindow, proposedSize objc.IObject /* cross-framework: Size */) corefoundation.Size) {
	d._WindowWillUseFullScreenContentSize = f
}

// SetWindowWillUseFullScreenPresentationOptions sets the handler for the WindowWillUseFullScreenPresentationOptions delegate method.
//
// Returns the presentation options the window uses when transitioning to full-screen mode.
func (d *WindowDelegate) SetWindowWillUseFullScreenPresentationOptions(f func(window IWindow, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions) {
	d._WindowWillUseFullScreenPresentationOptions = f
}

// SetWindowDidBecomeKey sets the handler for the WindowDidBecomeKey delegate method.
//
// Tells the delegate that the window has become the key window.
func (d *WindowDelegate) SetWindowDidBecomeKey(f func(notification foundation.Notification)) {
	d._WindowDidBecomeKey = f
}

// SetWindowDidBecomeMain sets the handler for the WindowDidBecomeMain delegate method.
//
// Tells the delegate that the window has become main.
func (d *WindowDelegate) SetWindowDidBecomeMain(f func(notification foundation.Notification)) {
	d._WindowDidBecomeMain = f
}

// SetWindowDidEndLiveResize sets the handler for the WindowDidEndLiveResize delegate method.
//
// Tells the delegate that a live resize operation on the window has ended.
func (d *WindowDelegate) SetWindowDidEndLiveResize(f func(notification foundation.Notification)) {
	d._WindowDidEndLiveResize = f
}

// SetWindowDidEnterFullScreen sets the handler for the WindowDidEnterFullScreen delegate method.
//
// The window has entered full-screen mode.
func (d *WindowDelegate) SetWindowDidEnterFullScreen(f func(notification foundation.Notification)) {
	d._WindowDidEnterFullScreen = f
}

// SetWindowDidExitFullScreen sets the handler for the WindowDidExitFullScreen delegate method.
//
// The window has left full-screen mode.
func (d *WindowDelegate) SetWindowDidExitFullScreen(f func(notification foundation.Notification)) {
	d._WindowDidExitFullScreen = f
}

// SetWindowDidResignKey sets the handler for the WindowDidResignKey delegate method.
//
// Tells the delegate that the window has resigned key window status.
func (d *WindowDelegate) SetWindowDidResignKey(f func(notification foundation.Notification)) {
	d._WindowDidResignKey = f
}

// SetWindowDidResignMain sets the handler for the WindowDidResignMain delegate method.
//
// Tells the delegate that the window has resigned main window status.
func (d *WindowDelegate) SetWindowDidResignMain(f func(notification foundation.Notification)) {
	d._WindowDidResignMain = f
}

// SetWindowDidResize sets the handler for the WindowDidResize delegate method.
//
// Tells the delegate that the window has been resized.
func (d *WindowDelegate) SetWindowDidResize(f func(notification foundation.Notification)) {
	d._WindowDidResize = f
}

// SetWindowShouldClose sets the handler for the WindowShouldClose delegate method.
//
// Tells the delegate that the user has attempted to close a window or the window has received a   message.
func (d *WindowDelegate) SetWindowShouldClose(f func(sender IWindow) bool) {
	d._WindowShouldClose = f
}

// SetWindowShouldZoomToFrame sets the handler for the WindowShouldZoomToFrame delegate method.
//
// Asks the delegate whether the specified window should zoom to the specified frame.
func (d *WindowDelegate) SetWindowShouldZoomToFrame(f func(window IWindow, newFrame objc.IObject /* cross-framework: Rect */) bool) {
	d._WindowShouldZoomToFrame = f
}

// SetWindowWillClose sets the handler for the WindowWillClose delegate method.
//
// Tells the delegate that the window is about to close.
func (d *WindowDelegate) SetWindowWillClose(f func(notification foundation.Notification)) {
	d._WindowWillClose = f
}

// SetWindowWillEnterFullScreen sets the handler for the WindowWillEnterFullScreen delegate method.
//
// The window is about to enter full-screen mode.
func (d *WindowDelegate) SetWindowWillEnterFullScreen(f func(notification foundation.Notification)) {
	d._WindowWillEnterFullScreen = f
}

// SetWindowWillExitFullScreen sets the handler for the WindowWillExitFullScreen delegate method.
//
// The window is about to exit full-screen mode.
func (d *WindowDelegate) SetWindowWillExitFullScreen(f func(notification foundation.Notification)) {
	d._WindowWillExitFullScreen = f
}

// SetWindowWillResizeToSize sets the handler for the WindowWillResizeToSize delegate method.
//
// Tells the delegate that the window is being resized (whether by the user or through one of the   methods other than  ).
func (d *WindowDelegate) SetWindowWillResizeToSize(f func(sender IWindow, frameSize objc.IObject /* cross-framework: Size */) corefoundation.Size) {
	d._WindowWillResizeToSize = f
}

// SetWindowWillReturnFieldEditorToObject sets the handler for the WindowWillReturnFieldEditorToObject delegate method.
//
// Tells the delegate that the field editor for a text-displaying object has been requested.
func (d *WindowDelegate) SetWindowWillReturnFieldEditorToObject(f func(sender IWindow, client objc.IObject) objc.ID) {
	d._WindowWillReturnFieldEditorToObject = f
}

// SetWindowWillStartLiveResize sets the handler for the WindowWillStartLiveResize delegate method.
//
// Tells the delegate that the window is about to be live resized.
func (d *WindowDelegate) SetWindowWillStartLiveResize(f func(notification foundation.Notification)) {
	d._WindowWillStartLiveResize = f
}

// SetWindowWillUseStandardFrameDefaultFrame sets the handler for the WindowWillUseStandardFrameDefaultFrame delegate method.
//
// Called by  ’s   method while determining the frame a window may be zoomed to.
func (d *WindowDelegate) SetWindowWillUseStandardFrameDefaultFrame(f func(window IWindow, newFrame objc.IObject /* cross-framework: Rect */) corefoundation.Rect) {
	d._WindowWillUseStandardFrameDefaultFrame = f
}

// WindowShouldDragDocumentWithEventFromWithPasteboard implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowShouldDragDocumentWithEventFromWithPasteboard(window IWindow, event IEvent, dragImageLocation objc.IObject /* cross-framework: Point */, pasteboard IPasteboard) bool {
	if d._WindowShouldDragDocumentWithEventFromWithPasteboard != nil {
		return d._WindowShouldDragDocumentWithEventFromWithPasteboard(window, event, dragImageLocation, pasteboard)
	}
	var zero bool
	return zero
}

// HasWindowShouldDragDocumentWithEventFromWithPasteboard returns true if a handler for WindowShouldDragDocumentWithEventFromWithPasteboard has been set.
func (d *WindowDelegate) HasWindowShouldDragDocumentWithEventFromWithPasteboard() bool {
	return d._WindowShouldDragDocumentWithEventFromWithPasteboard != nil
}

// WindowShouldPopUpDocumentPathMenu implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowShouldPopUpDocumentPathMenu(window IWindow, menu IMenu) bool {
	if d._WindowShouldPopUpDocumentPathMenu != nil {
		return d._WindowShouldPopUpDocumentPathMenu(window, menu)
	}
	var zero bool
	return zero
}

// HasWindowShouldPopUpDocumentPathMenu returns true if a handler for WindowShouldPopUpDocumentPathMenu has been set.
func (d *WindowDelegate) HasWindowShouldPopUpDocumentPathMenu() bool {
	return d._WindowShouldPopUpDocumentPathMenu != nil
}

// WindowWillUseFullScreenContentSize implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowWillUseFullScreenContentSize(window IWindow, proposedSize objc.IObject /* cross-framework: Size */) corefoundation.Size {
	if d._WindowWillUseFullScreenContentSize != nil {
		return d._WindowWillUseFullScreenContentSize(window, proposedSize)
	}
	var zero corefoundation.Size
	return zero
}

// HasWindowWillUseFullScreenContentSize returns true if a handler for WindowWillUseFullScreenContentSize has been set.
func (d *WindowDelegate) HasWindowWillUseFullScreenContentSize() bool {
	return d._WindowWillUseFullScreenContentSize != nil
}

// WindowWillUseFullScreenPresentationOptions implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowWillUseFullScreenPresentationOptions(window IWindow, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions {
	if d._WindowWillUseFullScreenPresentationOptions != nil {
		return d._WindowWillUseFullScreenPresentationOptions(window, proposedOptions)
	}
	var zero ApplicationPresentationOptions
	return zero
}

// HasWindowWillUseFullScreenPresentationOptions returns true if a handler for WindowWillUseFullScreenPresentationOptions has been set.
func (d *WindowDelegate) HasWindowWillUseFullScreenPresentationOptions() bool {
	return d._WindowWillUseFullScreenPresentationOptions != nil
}

// WindowDidBecomeKey implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowDidBecomeKey(notification foundation.Notification) {
	if d._WindowDidBecomeKey != nil {
		d._WindowDidBecomeKey(notification)
	}
}

// HasWindowDidBecomeKey returns true if a handler for WindowDidBecomeKey has been set.
func (d *WindowDelegate) HasWindowDidBecomeKey() bool {
	return d._WindowDidBecomeKey != nil
}

// WindowDidBecomeMain implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowDidBecomeMain(notification foundation.Notification) {
	if d._WindowDidBecomeMain != nil {
		d._WindowDidBecomeMain(notification)
	}
}

// HasWindowDidBecomeMain returns true if a handler for WindowDidBecomeMain has been set.
func (d *WindowDelegate) HasWindowDidBecomeMain() bool {
	return d._WindowDidBecomeMain != nil
}

// WindowDidEndLiveResize implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowDidEndLiveResize(notification foundation.Notification) {
	if d._WindowDidEndLiveResize != nil {
		d._WindowDidEndLiveResize(notification)
	}
}

// HasWindowDidEndLiveResize returns true if a handler for WindowDidEndLiveResize has been set.
func (d *WindowDelegate) HasWindowDidEndLiveResize() bool {
	return d._WindowDidEndLiveResize != nil
}

// WindowDidEnterFullScreen implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowDidEnterFullScreen(notification foundation.Notification) {
	if d._WindowDidEnterFullScreen != nil {
		d._WindowDidEnterFullScreen(notification)
	}
}

// HasWindowDidEnterFullScreen returns true if a handler for WindowDidEnterFullScreen has been set.
func (d *WindowDelegate) HasWindowDidEnterFullScreen() bool {
	return d._WindowDidEnterFullScreen != nil
}

// WindowDidExitFullScreen implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowDidExitFullScreen(notification foundation.Notification) {
	if d._WindowDidExitFullScreen != nil {
		d._WindowDidExitFullScreen(notification)
	}
}

// HasWindowDidExitFullScreen returns true if a handler for WindowDidExitFullScreen has been set.
func (d *WindowDelegate) HasWindowDidExitFullScreen() bool {
	return d._WindowDidExitFullScreen != nil
}

// WindowDidResignKey implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowDidResignKey(notification foundation.Notification) {
	if d._WindowDidResignKey != nil {
		d._WindowDidResignKey(notification)
	}
}

// HasWindowDidResignKey returns true if a handler for WindowDidResignKey has been set.
func (d *WindowDelegate) HasWindowDidResignKey() bool {
	return d._WindowDidResignKey != nil
}

// WindowDidResignMain implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowDidResignMain(notification foundation.Notification) {
	if d._WindowDidResignMain != nil {
		d._WindowDidResignMain(notification)
	}
}

// HasWindowDidResignMain returns true if a handler for WindowDidResignMain has been set.
func (d *WindowDelegate) HasWindowDidResignMain() bool {
	return d._WindowDidResignMain != nil
}

// WindowDidResize implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowDidResize(notification foundation.Notification) {
	if d._WindowDidResize != nil {
		d._WindowDidResize(notification)
	}
}

// HasWindowDidResize returns true if a handler for WindowDidResize has been set.
func (d *WindowDelegate) HasWindowDidResize() bool {
	return d._WindowDidResize != nil
}

// WindowShouldClose implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowShouldClose(sender IWindow) bool {
	if d._WindowShouldClose != nil {
		return d._WindowShouldClose(sender)
	}
	var zero bool
	return zero
}

// HasWindowShouldClose returns true if a handler for WindowShouldClose has been set.
func (d *WindowDelegate) HasWindowShouldClose() bool {
	return d._WindowShouldClose != nil
}

// WindowShouldZoomToFrame implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowShouldZoomToFrame(window IWindow, newFrame objc.IObject /* cross-framework: Rect */) bool {
	if d._WindowShouldZoomToFrame != nil {
		return d._WindowShouldZoomToFrame(window, newFrame)
	}
	var zero bool
	return zero
}

// HasWindowShouldZoomToFrame returns true if a handler for WindowShouldZoomToFrame has been set.
func (d *WindowDelegate) HasWindowShouldZoomToFrame() bool {
	return d._WindowShouldZoomToFrame != nil
}

// WindowWillClose implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowWillClose(notification foundation.Notification) {
	if d._WindowWillClose != nil {
		d._WindowWillClose(notification)
	}
}

// HasWindowWillClose returns true if a handler for WindowWillClose has been set.
func (d *WindowDelegate) HasWindowWillClose() bool {
	return d._WindowWillClose != nil
}

// WindowWillEnterFullScreen implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowWillEnterFullScreen(notification foundation.Notification) {
	if d._WindowWillEnterFullScreen != nil {
		d._WindowWillEnterFullScreen(notification)
	}
}

// HasWindowWillEnterFullScreen returns true if a handler for WindowWillEnterFullScreen has been set.
func (d *WindowDelegate) HasWindowWillEnterFullScreen() bool {
	return d._WindowWillEnterFullScreen != nil
}

// WindowWillExitFullScreen implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowWillExitFullScreen(notification foundation.Notification) {
	if d._WindowWillExitFullScreen != nil {
		d._WindowWillExitFullScreen(notification)
	}
}

// HasWindowWillExitFullScreen returns true if a handler for WindowWillExitFullScreen has been set.
func (d *WindowDelegate) HasWindowWillExitFullScreen() bool {
	return d._WindowWillExitFullScreen != nil
}

// WindowWillResizeToSize implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowWillResizeToSize(sender IWindow, frameSize objc.IObject /* cross-framework: Size */) corefoundation.Size {
	if d._WindowWillResizeToSize != nil {
		return d._WindowWillResizeToSize(sender, frameSize)
	}
	var zero corefoundation.Size
	return zero
}

// HasWindowWillResizeToSize returns true if a handler for WindowWillResizeToSize has been set.
func (d *WindowDelegate) HasWindowWillResizeToSize() bool {
	return d._WindowWillResizeToSize != nil
}

// WindowWillReturnFieldEditorToObject implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowWillReturnFieldEditorToObject(sender IWindow, client objc.IObject) objc.ID {
	if d._WindowWillReturnFieldEditorToObject != nil {
		return d._WindowWillReturnFieldEditorToObject(sender, client)
	}
	var zero objc.ID
	return zero
}

// HasWindowWillReturnFieldEditorToObject returns true if a handler for WindowWillReturnFieldEditorToObject has been set.
func (d *WindowDelegate) HasWindowWillReturnFieldEditorToObject() bool {
	return d._WindowWillReturnFieldEditorToObject != nil
}

// WindowWillStartLiveResize implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowWillStartLiveResize(notification foundation.Notification) {
	if d._WindowWillStartLiveResize != nil {
		d._WindowWillStartLiveResize(notification)
	}
}

// HasWindowWillStartLiveResize returns true if a handler for WindowWillStartLiveResize has been set.
func (d *WindowDelegate) HasWindowWillStartLiveResize() bool {
	return d._WindowWillStartLiveResize != nil
}

// WindowWillUseStandardFrameDefaultFrame implements the PWindowDelegate interface.
func (d *WindowDelegate) WindowWillUseStandardFrameDefaultFrame(window IWindow, newFrame objc.IObject /* cross-framework: Rect */) corefoundation.Rect {
	if d._WindowWillUseStandardFrameDefaultFrame != nil {
		return d._WindowWillUseStandardFrameDefaultFrame(window, newFrame)
	}
	var zero corefoundation.Rect
	return zero
}

// HasWindowWillUseStandardFrameDefaultFrame returns true if a handler for WindowWillUseStandardFrameDefaultFrame has been set.
func (d *WindowDelegate) HasWindowWillUseStandardFrameDefaultFrame() bool {
	return d._WindowWillUseStandardFrameDefaultFrame != nil
}
