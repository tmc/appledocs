// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// POpenSavePanelDelegate is the NSOpenSavePanelDelegate protocol interface.
//
// A set of methods for managing interactions with an open or save panel.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSOpenSavePanelDelegate
type POpenSavePanelDelegate interface {
	// Optional methods
	PanelDidChangeToDirectoryURL(sender objc.IObject, url objc.IObject /* cross-framework: NSURL */)
	HasPanelDidChangeToDirectoryURL() bool
}

// OpenSavePanelDelegate is a delegate implementation builder for the POpenSavePanelDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type OpenSavePanelDelegate struct {
	_PanelDidChangeToDirectoryURL func(sender objc.IObject, url objc.IObject /* cross-framework: NSURL */)
}

// SetPanelDidChangeToDirectoryURL sets the handler for the PanelDidChangeToDirectoryURL delegate method.
//
// Tells the delegate that the user changed the selected directory to the directory located at the specified URL.
func (d *OpenSavePanelDelegate) SetPanelDidChangeToDirectoryURL(f func(sender objc.IObject, url objc.IObject /* cross-framework: NSURL */)) {
	d._PanelDidChangeToDirectoryURL = f
}

// PanelDidChangeToDirectoryURL implements the POpenSavePanelDelegate interface.
func (d *OpenSavePanelDelegate) PanelDidChangeToDirectoryURL(sender objc.IObject, url objc.IObject /* cross-framework: NSURL */) {
	if d._PanelDidChangeToDirectoryURL != nil {
		d._PanelDidChangeToDirectoryURL(sender, url)
	}
}

// HasPanelDidChangeToDirectoryURL returns true if a handler for PanelDidChangeToDirectoryURL has been set.
func (d *OpenSavePanelDelegate) HasPanelDidChangeToDirectoryURL() bool {
	return d._PanelDidChangeToDirectoryURL != nil
}
