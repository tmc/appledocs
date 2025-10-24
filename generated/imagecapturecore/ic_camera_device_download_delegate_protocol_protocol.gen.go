// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PICCameraDeviceDownloadDelegate is the ICCameraDeviceDownloadDelegate protocol interface.
//
// Methods for managing camera file downloads.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// See: doc://com.apple.imagecapturecore/documentation/ImageCaptureCore/ICCameraDeviceDownloadDelegate
type PICCameraDeviceDownloadDelegate interface {
	// Required methods
	DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes(file ICCameraFile, downloadedBytes unsafe.Pointer, maxBytes unsafe.Pointer)/* debug [protocol_interface/required_method]: DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes */
	// Optional methods
	DidReceiveDownloadProgress()
	HasDidReceiveDownloadProgress() bool
	DidDownloadFile()
	HasDidDownloadFile() bool
	DidDownloadFileErrorOptionsContextInfo(file ICCameraFile, error_ objc.IObject /* cross-framework: Error */, options foundation.IDictionary, contextInfo unsafe.Pointer)
	HasDidDownloadFileErrorOptionsContextInfo() bool
}

// ICCameraDeviceDownloadDelegate is a delegate implementation builder for the PICCameraDeviceDownloadDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ICCameraDeviceDownloadDelegate struct {
	_DidReceiveDownloadProgress func()
	_DidDownloadFile func()
	_DidDownloadFileErrorOptionsContextInfo func(file ICCameraFile, error_ objc.IObject /* cross-framework: Error */, options foundation.IDictionary, contextInfo unsafe.Pointer)
	_DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes func(file ICCameraFile, downloadedBytes unsafe.Pointer, maxBytes unsafe.Pointer)
}

// SetDidReceiveDownloadProgress sets the handler for the DidReceiveDownloadProgress delegate method.
//
// Updates the delegate about the status of the download.
func (d *ICCameraDeviceDownloadDelegate) SetDidReceiveDownloadProgress(f func()) {
	d._DidReceiveDownloadProgress = f
}

// SetDidDownloadFile sets the handler for the DidDownloadFile delegate method.
//
// Tells the delegate that the requested download has completed.
func (d *ICCameraDeviceDownloadDelegate) SetDidDownloadFile(f func()) {
	d._DidDownloadFile = f
}

// SetDidDownloadFileErrorOptionsContextInfo sets the handler for the DidDownloadFileErrorOptionsContextInfo delegate method.
//
// Tells the delegate that the requested download has completed.
func (d *ICCameraDeviceDownloadDelegate) SetDidDownloadFileErrorOptionsContextInfo(f func(file ICCameraFile, error_ objc.IObject /* cross-framework: Error */, options foundation.IDictionary, contextInfo unsafe.Pointer)) {
	d._DidDownloadFileErrorOptionsContextInfo = f
}

// SetDidReceiveDownloadProgressForFileDownloadedBytesMaxBytes sets the handler for the DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes delegate method.
//
// Updates the delegate about the status of the download.
func (d *ICCameraDeviceDownloadDelegate) SetDidReceiveDownloadProgressForFileDownloadedBytesMaxBytes(f func(file ICCameraFile, downloadedBytes unsafe.Pointer, maxBytes unsafe.Pointer)) {
	d._DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes = f
}

// DidReceiveDownloadProgress implements the PICCameraDeviceDownloadDelegate interface.
func (d *ICCameraDeviceDownloadDelegate) DidReceiveDownloadProgress() {
	if d._DidReceiveDownloadProgress != nil {
		d._DidReceiveDownloadProgress()
	}
}

// HasDidReceiveDownloadProgress returns true if a handler for DidReceiveDownloadProgress has been set.
func (d *ICCameraDeviceDownloadDelegate) HasDidReceiveDownloadProgress() bool {
	return d._DidReceiveDownloadProgress != nil
}

// DidDownloadFile implements the PICCameraDeviceDownloadDelegate interface.
func (d *ICCameraDeviceDownloadDelegate) DidDownloadFile() {
	if d._DidDownloadFile != nil {
		d._DidDownloadFile()
	}
}

// HasDidDownloadFile returns true if a handler for DidDownloadFile has been set.
func (d *ICCameraDeviceDownloadDelegate) HasDidDownloadFile() bool {
	return d._DidDownloadFile != nil
}

// DidDownloadFileErrorOptionsContextInfo implements the PICCameraDeviceDownloadDelegate interface.
func (d *ICCameraDeviceDownloadDelegate) DidDownloadFileErrorOptionsContextInfo(file ICCameraFile, error_ objc.IObject /* cross-framework: Error */, options foundation.IDictionary, contextInfo unsafe.Pointer) {
	if d._DidDownloadFileErrorOptionsContextInfo != nil {
		d._DidDownloadFileErrorOptionsContextInfo(file, error_, options, contextInfo)
	}
}

// HasDidDownloadFileErrorOptionsContextInfo returns true if a handler for DidDownloadFileErrorOptionsContextInfo has been set.
func (d *ICCameraDeviceDownloadDelegate) HasDidDownloadFileErrorOptionsContextInfo() bool {
	return d._DidDownloadFileErrorOptionsContextInfo != nil
}

// DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes implements the PICCameraDeviceDownloadDelegate interface.
func (d *ICCameraDeviceDownloadDelegate) DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes(file ICCameraFile, downloadedBytes unsafe.Pointer, maxBytes unsafe.Pointer) {
	if d._DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes != nil {
		d._DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes(file, downloadedBytes, maxBytes)
	}
}

// HasDidReceiveDownloadProgressForFileDownloadedBytesMaxBytes returns true if a handler for DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes has been set.
func (d *ICCameraDeviceDownloadDelegate) HasDidReceiveDownloadProgressForFileDownloadedBytesMaxBytes() bool {
	return d._DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes != nil
}
