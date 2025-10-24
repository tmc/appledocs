// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/imagecapturecore"
)

// PIKDeviceBrowserViewDelegate is the IKDeviceBrowserViewDelegate protocol interface.
//
// The   defines the methods that the delegate of the   class can implement. All the methods are optional.
//
// Availability:
//   - macOS 10.4+
//
// See: doc://com.apple.quartz/documentation/Quartz/IKDeviceBrowserViewDelegate
type PIKDeviceBrowserViewDelegate interface {
	// Required methods
	DeviceBrowserViewSelectionDidChange(deviceBrowserView IKDeviceBrowserView, device objc.IObject)/* debug [protocol_interface/required_method]: DeviceBrowserViewSelectionDidChange */
	// Optional methods
	DeviceBrowserViewDidEncounterError(deviceBrowserView IKDeviceBrowserView, error_ objc.IObject /* cross-framework: Error */)
	HasDeviceBrowserViewDidEncounterError() bool
}

// IKDeviceBrowserViewDelegate is a delegate implementation builder for the PIKDeviceBrowserViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type IKDeviceBrowserViewDelegate struct {
	_DeviceBrowserViewDidEncounterError func(deviceBrowserView IKDeviceBrowserView, error_ objc.IObject /* cross-framework: Error */)
	_DeviceBrowserViewSelectionDidChange func(deviceBrowserView IKDeviceBrowserView, device objc.IObject)
}

// SetDeviceBrowserViewDidEncounterError sets the handler for the DeviceBrowserViewDidEncounterError delegate method.
//
// Invoked when the device browser encounters an error.
func (d *IKDeviceBrowserViewDelegate) SetDeviceBrowserViewDidEncounterError(f func(deviceBrowserView IKDeviceBrowserView, error_ objc.IObject /* cross-framework: Error */)) {
	d._DeviceBrowserViewDidEncounterError = f
}

// SetDeviceBrowserViewSelectionDidChange sets the handler for the DeviceBrowserViewSelectionDidChange delegate method.
//
// Sent to the delegate when the selection changes in the browser view.
func (d *IKDeviceBrowserViewDelegate) SetDeviceBrowserViewSelectionDidChange(f func(deviceBrowserView IKDeviceBrowserView, device objc.IObject)) {
	d._DeviceBrowserViewSelectionDidChange = f
}

// DeviceBrowserViewDidEncounterError implements the PIKDeviceBrowserViewDelegate interface.
func (d *IKDeviceBrowserViewDelegate) DeviceBrowserViewDidEncounterError(deviceBrowserView IKDeviceBrowserView, error_ objc.IObject /* cross-framework: Error */) {
	if d._DeviceBrowserViewDidEncounterError != nil {
		d._DeviceBrowserViewDidEncounterError(deviceBrowserView, error_)
	}
}

// HasDeviceBrowserViewDidEncounterError returns true if a handler for DeviceBrowserViewDidEncounterError has been set.
func (d *IKDeviceBrowserViewDelegate) HasDeviceBrowserViewDidEncounterError() bool {
	return d._DeviceBrowserViewDidEncounterError != nil
}

// DeviceBrowserViewSelectionDidChange implements the PIKDeviceBrowserViewDelegate interface.
func (d *IKDeviceBrowserViewDelegate) DeviceBrowserViewSelectionDidChange(deviceBrowserView IKDeviceBrowserView, device objc.IObject) {
	if d._DeviceBrowserViewSelectionDidChange != nil {
		d._DeviceBrowserViewSelectionDidChange(deviceBrowserView, device)
	}
}

// HasDeviceBrowserViewSelectionDidChange returns true if a handler for DeviceBrowserViewSelectionDidChange has been set.
func (d *IKDeviceBrowserViewDelegate) HasDeviceBrowserViewSelectionDidChange() bool {
	return d._DeviceBrowserViewSelectionDidChange != nil
}
