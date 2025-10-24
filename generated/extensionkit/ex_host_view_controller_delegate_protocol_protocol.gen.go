// Code generated from Apple documentation for ExtensionKit. DO NOT EDIT.

package extensionkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PEXHostViewControllerDelegate is the EXHostViewControllerDelegate protocol interface.
//
// An interface you use to track the activation and deactivation of an app extension.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 13.0+
//
// See: doc://com.apple.extensionkit/documentation/ExtensionKit/EXHostViewControllerDelegate
type PEXHostViewControllerDelegate interface {
	// Optional methods
	HostViewControllerDidActivate()
	HasHostViewControllerDidActivate() bool
	HostViewControllerWillDeactivate()
	HasHostViewControllerWillDeactivate() bool
}

// EXHostViewControllerDelegate is a delegate implementation builder for the PEXHostViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type EXHostViewControllerDelegate struct {
	_HostViewControllerDidActivate func()
	_HostViewControllerWillDeactivate func()
}

// SetHostViewControllerDidActivate sets the handler for the HostViewControllerDidActivate delegate method.
//
// Tells the host that the app extension is active and ready to accept an XPC connection.
func (d *EXHostViewControllerDelegate) SetHostViewControllerDidActivate(f func()) {
	d._HostViewControllerDidActivate = f
}

// SetHostViewControllerWillDeactivate sets the handler for the HostViewControllerWillDeactivate delegate method.
//
// Tells the host that the app extension disconnected and is no longer available.
func (d *EXHostViewControllerDelegate) SetHostViewControllerWillDeactivate(f func()) {
	d._HostViewControllerWillDeactivate = f
}

// HostViewControllerDidActivate implements the PEXHostViewControllerDelegate interface.
func (d *EXHostViewControllerDelegate) HostViewControllerDidActivate() {
	if d._HostViewControllerDidActivate != nil {
		d._HostViewControllerDidActivate()
	}
}

// HasHostViewControllerDidActivate returns true if a handler for HostViewControllerDidActivate has been set.
func (d *EXHostViewControllerDelegate) HasHostViewControllerDidActivate() bool {
	return d._HostViewControllerDidActivate != nil
}

// HostViewControllerWillDeactivate implements the PEXHostViewControllerDelegate interface.
func (d *EXHostViewControllerDelegate) HostViewControllerWillDeactivate() {
	if d._HostViewControllerWillDeactivate != nil {
		d._HostViewControllerWillDeactivate()
	}
}

// HasHostViewControllerWillDeactivate returns true if a handler for HostViewControllerWillDeactivate has been set.
func (d *EXHostViewControllerDelegate) HasHostViewControllerWillDeactivate() bool {
	return d._HostViewControllerWillDeactivate != nil
}
