// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	PanelDidChangeToDirectoryURL(sender objectivec.IObject, url foundation.foundation.INSURL)
	HasPanelDidChangeToDirectoryURL() bool
}

// OpenSavePanelDelegate is a delegate implementation builder for the POpenSavePanelDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type OpenSavePanelDelegate struct {
	_PanelDidChangeToDirectoryURL func(sender objectivec.IObject, url foundation.foundation.INSURL)
}

// SetPanelDidChangeToDirectoryURL sets the handler for the PanelDidChangeToDirectoryURL delegate method.
//
// Tells the delegate that the user changed the selected directory to the directory located at the specified URL.
func (d *OpenSavePanelDelegate) SetPanelDidChangeToDirectoryURL(f func(sender objectivec.IObject, url foundation.foundation.INSURL)) {
	d._PanelDidChangeToDirectoryURL = f
}

// PanelDidChangeToDirectoryURL implements the POpenSavePanelDelegate interface.
func (d *OpenSavePanelDelegate) PanelDidChangeToDirectoryURL(sender objectivec.IObject, url foundation.foundation.INSURL) {
	if d._PanelDidChangeToDirectoryURL != nil {
		d._PanelDidChangeToDirectoryURL(sender, url)
	}
}

// HasPanelDidChangeToDirectoryURL returns true if a handler for PanelDidChangeToDirectoryURL has been set.
func (d *OpenSavePanelDelegate) HasPanelDidChangeToDirectoryURL() bool {
	return d._PanelDidChangeToDirectoryURL != nil
}

// OpenSavePanelDelegateObject wraps an existing Objective-C object that conforms to the POpenSavePanelDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type OpenSavePanelDelegateObject struct {
	objectivec.Object
}

// NewOpenSavePanelDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSOpenSavePanelDelegate protocol.
func NewOpenSavePanelDelegateObject(obj objectivec.Object) *OpenSavePanelDelegateObject {
	return &OpenSavePanelDelegateObject{obj}
}

// Make sure OpenSavePanelDelegateObject implements POpenSavePanelDelegate.
var _ POpenSavePanelDelegate = (*OpenSavePanelDelegateObject)(nil)

// PanelDidChangeToDirectoryURL implements the POpenSavePanelDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OpenSavePanelDelegateObject) PanelDidChangeToDirectoryURL(sender objectivec.IObject, url foundation.foundation.INSURL) {
	objc.Send[objc.ID](o.ID, objc.Sel("panel:didChangeToDirectoryURL:"), sender, url)
}

// HasPanelDidChangeToDirectoryURL returns true; this is a placeholder for optional method checks.
func (o *OpenSavePanelDelegateObject) HasPanelDidChangeToDirectoryURL() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
