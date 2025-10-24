// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/imagecapturecore"
)

// PIKCameraDeviceViewDelegate is the IKCameraDeviceViewDelegate protocol interface.
//
// The   protocol is adopted by the delegate of the   class. It allows downloading of camera content, handling selection changes, and handling errors.
//
// Availability:
//   - macOS 10.4+
//
// See: doc://com.apple.quartz/documentation/Quartz/IKCameraDeviceViewDelegate
type PIKCameraDeviceViewDelegate interface {
	// Optional methods
	CameraDeviceViewDidDownloadFileLocationFileDataError(cameraDeviceView IKCameraDeviceView, file objc.IObject, url objc.IObject /* cross-framework: NSURL */, data objc.IObject /* cross-framework: NSData */, error_ objc.IObject /* cross-framework: Error */)
	HasCameraDeviceViewDidDownloadFileLocationFileDataError() bool
	CameraDeviceViewDidEncounterError(cameraDeviceView IKCameraDeviceView, error_ objc.IObject /* cross-framework: Error */)
	HasCameraDeviceViewDidEncounterError() bool
	CameraDeviceViewSelectionDidChange(cameraDeviceView IKCameraDeviceView)
	HasCameraDeviceViewSelectionDidChange() bool
}

// IKCameraDeviceViewDelegate is a delegate implementation builder for the PIKCameraDeviceViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type IKCameraDeviceViewDelegate struct {
	_CameraDeviceViewDidDownloadFileLocationFileDataError func(cameraDeviceView IKCameraDeviceView, file objc.IObject, url objc.IObject /* cross-framework: NSURL */, data objc.IObject /* cross-framework: NSData */, error_ objc.IObject /* cross-framework: Error */)
	_CameraDeviceViewDidEncounterError func(cameraDeviceView IKCameraDeviceView, error_ objc.IObject /* cross-framework: Error */)
	_CameraDeviceViewSelectionDidChange func(cameraDeviceView IKCameraDeviceView)
}

// SetCameraDeviceViewDidDownloadFileLocationFileDataError sets the handler for the CameraDeviceViewDidDownloadFileLocationFileDataError delegate method.
//
// Invoked for each file that is downloaded from the camera device.
func (d *IKCameraDeviceViewDelegate) SetCameraDeviceViewDidDownloadFileLocationFileDataError(f func(cameraDeviceView IKCameraDeviceView, file objc.IObject, url objc.IObject /* cross-framework: NSURL */, data objc.IObject /* cross-framework: NSData */, error_ objc.IObject /* cross-framework: Error */)) {
	d._CameraDeviceViewDidDownloadFileLocationFileDataError = f
}

// SetCameraDeviceViewDidEncounterError sets the handler for the CameraDeviceViewDidEncounterError delegate method.
//
// Invoked when the camera encounters an error.
func (d *IKCameraDeviceViewDelegate) SetCameraDeviceViewDidEncounterError(f func(cameraDeviceView IKCameraDeviceView, error_ objc.IObject /* cross-framework: Error */)) {
	d._CameraDeviceViewDidEncounterError = f
}

// SetCameraDeviceViewSelectionDidChange sets the handler for the CameraDeviceViewSelectionDidChange delegate method.
//
// Invoked when the selection changed.
func (d *IKCameraDeviceViewDelegate) SetCameraDeviceViewSelectionDidChange(f func(cameraDeviceView IKCameraDeviceView)) {
	d._CameraDeviceViewSelectionDidChange = f
}

// CameraDeviceViewDidDownloadFileLocationFileDataError implements the PIKCameraDeviceViewDelegate interface.
func (d *IKCameraDeviceViewDelegate) CameraDeviceViewDidDownloadFileLocationFileDataError(cameraDeviceView IKCameraDeviceView, file objc.IObject, url objc.IObject /* cross-framework: NSURL */, data objc.IObject /* cross-framework: NSData */, error_ objc.IObject /* cross-framework: Error */) {
	if d._CameraDeviceViewDidDownloadFileLocationFileDataError != nil {
		d._CameraDeviceViewDidDownloadFileLocationFileDataError(cameraDeviceView, file, url, data, error_)
	}
}

// HasCameraDeviceViewDidDownloadFileLocationFileDataError returns true if a handler for CameraDeviceViewDidDownloadFileLocationFileDataError has been set.
func (d *IKCameraDeviceViewDelegate) HasCameraDeviceViewDidDownloadFileLocationFileDataError() bool {
	return d._CameraDeviceViewDidDownloadFileLocationFileDataError != nil
}

// CameraDeviceViewDidEncounterError implements the PIKCameraDeviceViewDelegate interface.
func (d *IKCameraDeviceViewDelegate) CameraDeviceViewDidEncounterError(cameraDeviceView IKCameraDeviceView, error_ objc.IObject /* cross-framework: Error */) {
	if d._CameraDeviceViewDidEncounterError != nil {
		d._CameraDeviceViewDidEncounterError(cameraDeviceView, error_)
	}
}

// HasCameraDeviceViewDidEncounterError returns true if a handler for CameraDeviceViewDidEncounterError has been set.
func (d *IKCameraDeviceViewDelegate) HasCameraDeviceViewDidEncounterError() bool {
	return d._CameraDeviceViewDidEncounterError != nil
}

// CameraDeviceViewSelectionDidChange implements the PIKCameraDeviceViewDelegate interface.
func (d *IKCameraDeviceViewDelegate) CameraDeviceViewSelectionDidChange(cameraDeviceView IKCameraDeviceView) {
	if d._CameraDeviceViewSelectionDidChange != nil {
		d._CameraDeviceViewSelectionDidChange(cameraDeviceView)
	}
}

// HasCameraDeviceViewSelectionDidChange returns true if a handler for CameraDeviceViewSelectionDidChange has been set.
func (d *IKCameraDeviceViewDelegate) HasCameraDeviceViewSelectionDidChange() bool {
	return d._CameraDeviceViewSelectionDidChange != nil
}
