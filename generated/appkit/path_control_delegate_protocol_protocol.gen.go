// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PPathControlDelegate is the NSPathControlDelegate protocol interface.
//
// A set of methods that can be implemented by the delegate of a path control object to support dragging to and from the control.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSPathControlDelegate
type PPathControlDelegate interface {
	// Optional methods
	PathControlAcceptDrop(pathControl IPathControl, info unsafe.Pointer) bool
	HasPathControlAcceptDrop() bool
	PathControlShouldDragPathComponentCellWithPasteboard(pathControl IPathControl, pathComponentCell IPathComponentCell, pasteboard IPasteboard) bool
	HasPathControlShouldDragPathComponentCellWithPasteboard() bool
	PathControlShouldDragItemWithPasteboard(pathControl IPathControl, pathItem IPathControlItem, pasteboard IPasteboard) bool
	HasPathControlShouldDragItemWithPasteboard() bool
	PathControlValidateDrop(pathControl IPathControl, info unsafe.Pointer) DragOperation
	HasPathControlValidateDrop() bool
	PathControlWillDisplayOpenPanel(pathControl IPathControl, openPanel IOpenPanel)
	HasPathControlWillDisplayOpenPanel() bool
	PathControlWillPopUpMenu(pathControl IPathControl, menu IMenu)
	HasPathControlWillPopUpMenu() bool
}

// PathControlDelegate is a delegate implementation builder for the PPathControlDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PathControlDelegate struct {
	_PathControlAcceptDrop func(pathControl IPathControl, info unsafe.Pointer) bool
	_PathControlShouldDragPathComponentCellWithPasteboard func(pathControl IPathControl, pathComponentCell IPathComponentCell, pasteboard IPasteboard) bool
	_PathControlShouldDragItemWithPasteboard func(pathControl IPathControl, pathItem IPathControlItem, pasteboard IPasteboard) bool
	_PathControlValidateDrop func(pathControl IPathControl, info unsafe.Pointer) DragOperation
	_PathControlWillDisplayOpenPanel func(pathControl IPathControl, openPanel IOpenPanel)
	_PathControlWillPopUpMenu func(pathControl IPathControl, menu IMenu)
}

// SetPathControlAcceptDrop sets the handler for the PathControlAcceptDrop delegate method.
//
// Implement this method to accept previously validated contents dropped onto the control.
func (d *PathControlDelegate) SetPathControlAcceptDrop(f func(pathControl IPathControl, info unsafe.Pointer) bool) {
	d._PathControlAcceptDrop = f
}

// SetPathControlShouldDragPathComponentCellWithPasteboard sets the handler for the PathControlShouldDragPathComponentCellWithPasteboard delegate method.
//
// Implement this method to enable dragging from the control.
func (d *PathControlDelegate) SetPathControlShouldDragPathComponentCellWithPasteboard(f func(pathControl IPathControl, pathComponentCell IPathComponentCell, pasteboard IPasteboard) bool) {
	d._PathControlShouldDragPathComponentCellWithPasteboard = f
}

// SetPathControlShouldDragItemWithPasteboard sets the handler for the PathControlShouldDragItemWithPasteboard delegate method.
func (d *PathControlDelegate) SetPathControlShouldDragItemWithPasteboard(f func(pathControl IPathControl, pathItem IPathControlItem, pasteboard IPasteboard) bool) {
	d._PathControlShouldDragItemWithPasteboard = f
}

// SetPathControlValidateDrop sets the handler for the PathControlValidateDrop delegate method.
//
// Implement this method to enable dragging onto the control.
func (d *PathControlDelegate) SetPathControlValidateDrop(f func(pathControl IPathControl, info unsafe.Pointer) DragOperation) {
	d._PathControlValidateDrop = f
}

// SetPathControlWillDisplayOpenPanel sets the handler for the PathControlWillDisplayOpenPanel delegate method.
//
// Implement this method to customize the Open panel shown by a pop-up–style path.
func (d *PathControlDelegate) SetPathControlWillDisplayOpenPanel(f func(pathControl IPathControl, openPanel IOpenPanel)) {
	d._PathControlWillDisplayOpenPanel = f
}

// SetPathControlWillPopUpMenu sets the handler for the PathControlWillPopUpMenu delegate method.
//
// Implement this method to customize the menu of a pop-up–style path.
func (d *PathControlDelegate) SetPathControlWillPopUpMenu(f func(pathControl IPathControl, menu IMenu)) {
	d._PathControlWillPopUpMenu = f
}

// PathControlAcceptDrop implements the PPathControlDelegate interface.
func (d *PathControlDelegate) PathControlAcceptDrop(pathControl IPathControl, info unsafe.Pointer) bool {
	if d._PathControlAcceptDrop != nil {
		return d._PathControlAcceptDrop(pathControl, info)
	}
	var zero bool
	return zero
}

// HasPathControlAcceptDrop returns true if a handler for PathControlAcceptDrop has been set.
func (d *PathControlDelegate) HasPathControlAcceptDrop() bool {
	return d._PathControlAcceptDrop != nil
}

// PathControlShouldDragPathComponentCellWithPasteboard implements the PPathControlDelegate interface.
func (d *PathControlDelegate) PathControlShouldDragPathComponentCellWithPasteboard(pathControl IPathControl, pathComponentCell IPathComponentCell, pasteboard IPasteboard) bool {
	if d._PathControlShouldDragPathComponentCellWithPasteboard != nil {
		return d._PathControlShouldDragPathComponentCellWithPasteboard(pathControl, pathComponentCell, pasteboard)
	}
	var zero bool
	return zero
}

// HasPathControlShouldDragPathComponentCellWithPasteboard returns true if a handler for PathControlShouldDragPathComponentCellWithPasteboard has been set.
func (d *PathControlDelegate) HasPathControlShouldDragPathComponentCellWithPasteboard() bool {
	return d._PathControlShouldDragPathComponentCellWithPasteboard != nil
}

// PathControlShouldDragItemWithPasteboard implements the PPathControlDelegate interface.
func (d *PathControlDelegate) PathControlShouldDragItemWithPasteboard(pathControl IPathControl, pathItem IPathControlItem, pasteboard IPasteboard) bool {
	if d._PathControlShouldDragItemWithPasteboard != nil {
		return d._PathControlShouldDragItemWithPasteboard(pathControl, pathItem, pasteboard)
	}
	var zero bool
	return zero
}

// HasPathControlShouldDragItemWithPasteboard returns true if a handler for PathControlShouldDragItemWithPasteboard has been set.
func (d *PathControlDelegate) HasPathControlShouldDragItemWithPasteboard() bool {
	return d._PathControlShouldDragItemWithPasteboard != nil
}

// PathControlValidateDrop implements the PPathControlDelegate interface.
func (d *PathControlDelegate) PathControlValidateDrop(pathControl IPathControl, info unsafe.Pointer) DragOperation {
	if d._PathControlValidateDrop != nil {
		return d._PathControlValidateDrop(pathControl, info)
	}
	var zero DragOperation
	return zero
}

// HasPathControlValidateDrop returns true if a handler for PathControlValidateDrop has been set.
func (d *PathControlDelegate) HasPathControlValidateDrop() bool {
	return d._PathControlValidateDrop != nil
}

// PathControlWillDisplayOpenPanel implements the PPathControlDelegate interface.
func (d *PathControlDelegate) PathControlWillDisplayOpenPanel(pathControl IPathControl, openPanel IOpenPanel) {
	if d._PathControlWillDisplayOpenPanel != nil {
		d._PathControlWillDisplayOpenPanel(pathControl, openPanel)
	}
}

// HasPathControlWillDisplayOpenPanel returns true if a handler for PathControlWillDisplayOpenPanel has been set.
func (d *PathControlDelegate) HasPathControlWillDisplayOpenPanel() bool {
	return d._PathControlWillDisplayOpenPanel != nil
}

// PathControlWillPopUpMenu implements the PPathControlDelegate interface.
func (d *PathControlDelegate) PathControlWillPopUpMenu(pathControl IPathControl, menu IMenu) {
	if d._PathControlWillPopUpMenu != nil {
		d._PathControlWillPopUpMenu(pathControl, menu)
	}
}

// HasPathControlWillPopUpMenu returns true if a handler for PathControlWillPopUpMenu has been set.
func (d *PathControlDelegate) HasPathControlWillPopUpMenu() bool {
	return d._PathControlWillPopUpMenu != nil
}
