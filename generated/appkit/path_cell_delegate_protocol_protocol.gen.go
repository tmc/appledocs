// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PPathCellDelegate is the NSPathCellDelegate protocol interface.
//
// A set of methods that enable the delegate of a path cell object to customize the Open panel or pop-up menu of a path whose style is set to  .
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSPathCellDelegate
type PPathCellDelegate interface {
	// Optional methods
	PathCellWillDisplayOpenPanel(pathCell IPathCell, openPanel IOpenPanel)
	HasPathCellWillDisplayOpenPanel() bool
	PathCellWillPopUpMenu(pathCell IPathCell, menu IMenu)
	HasPathCellWillPopUpMenu() bool
}

// PathCellDelegate is a delegate implementation builder for the PPathCellDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PathCellDelegate struct {
	_PathCellWillDisplayOpenPanel func(pathCell IPathCell, openPanel IOpenPanel)
	_PathCellWillPopUpMenu func(pathCell IPathCell, menu IMenu)
}

// SetPathCellWillDisplayOpenPanel sets the handler for the PathCellWillDisplayOpenPanel delegate method.
//
// Implement this method to customize the Open panel shown by a pop-up–style path.
func (d *PathCellDelegate) SetPathCellWillDisplayOpenPanel(f func(pathCell IPathCell, openPanel IOpenPanel)) {
	d._PathCellWillDisplayOpenPanel = f
}

// SetPathCellWillPopUpMenu sets the handler for the PathCellWillPopUpMenu delegate method.
//
// Implement this method to customize the menu of a pop-up–style path.
func (d *PathCellDelegate) SetPathCellWillPopUpMenu(f func(pathCell IPathCell, menu IMenu)) {
	d._PathCellWillPopUpMenu = f
}

// PathCellWillDisplayOpenPanel implements the PPathCellDelegate interface.
func (d *PathCellDelegate) PathCellWillDisplayOpenPanel(pathCell IPathCell, openPanel IOpenPanel) {
	if d._PathCellWillDisplayOpenPanel != nil {
		d._PathCellWillDisplayOpenPanel(pathCell, openPanel)
	}
}

// HasPathCellWillDisplayOpenPanel returns true if a handler for PathCellWillDisplayOpenPanel has been set.
func (d *PathCellDelegate) HasPathCellWillDisplayOpenPanel() bool {
	return d._PathCellWillDisplayOpenPanel != nil
}

// PathCellWillPopUpMenu implements the PPathCellDelegate interface.
func (d *PathCellDelegate) PathCellWillPopUpMenu(pathCell IPathCell, menu IMenu) {
	if d._PathCellWillPopUpMenu != nil {
		d._PathCellWillPopUpMenu(pathCell, menu)
	}
}

// HasPathCellWillPopUpMenu returns true if a handler for PathCellWillPopUpMenu has been set.
func (d *PathCellDelegate) HasPathCellWillPopUpMenu() bool {
	return d._PathCellWillPopUpMenu != nil
}

// PathCellDelegateObject wraps an existing Objective-C object that conforms to the PPathCellDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type PathCellDelegateObject struct {
	objectivec.Object
}

// NewPathCellDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSPathCellDelegate protocol.
func NewPathCellDelegateObject(obj objectivec.Object) *PathCellDelegateObject {
	return &PathCellDelegateObject{obj}
}

// Make sure PathCellDelegateObject implements PPathCellDelegate.
var _ PPathCellDelegate = (*PathCellDelegateObject)(nil)

// PathCellWillDisplayOpenPanel implements the PPathCellDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PathCellDelegateObject) PathCellWillDisplayOpenPanel(pathCell IPathCell, openPanel IOpenPanel) {
	objc.Send[objc.ID](o.ID, objc.Sel("pathCell:willDisplayOpenPanel:"), pathCell, openPanel)
}

// HasPathCellWillDisplayOpenPanel returns true; this is a placeholder for optional method checks.
func (o *PathCellDelegateObject) HasPathCellWillDisplayOpenPanel() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// PathCellWillPopUpMenu implements the PPathCellDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PathCellDelegateObject) PathCellWillPopUpMenu(pathCell IPathCell, menu IMenu) {
	objc.Send[objc.ID](o.ID, objc.Sel("pathCell:willPopUpMenu:"), pathCell, menu)
}

// HasPathCellWillPopUpMenu returns true; this is a placeholder for optional method checks.
func (o *PathCellDelegateObject) HasPathCellWillPopUpMenu() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
