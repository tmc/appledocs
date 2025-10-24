// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

package visionkit

import (
	"github.com/tmc/appledocs/generated/objc"
)

// PDocumentCameraViewControllerDelegate is the VNDocumentCameraViewControllerDelegate protocol interface.
//
// A delegate protocol through which the document camera returns its scanned results.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.VisionKit/documentation/VisionKit/VNDocumentCameraViewControllerDelegate
type PDocumentCameraViewControllerDelegate interface {
	// Optional methods
	DocumentCameraViewControllerDidFailWithError(controller IVNDocumentCameraViewController, error_ objc.IObject /* cross-framework: Error */)
	HasDocumentCameraViewControllerDidFailWithError() bool
	DocumentCameraViewControllerDidFinishWithScan(controller IVNDocumentCameraViewController, scan IVNDocumentCameraScan)
	HasDocumentCameraViewControllerDidFinishWithScan() bool
	DocumentCameraViewControllerDidCancel(controller IVNDocumentCameraViewController)
	HasDocumentCameraViewControllerDidCancel() bool
}

// DocumentCameraViewControllerDelegate is a delegate implementation builder for the PDocumentCameraViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type DocumentCameraViewControllerDelegate struct {
	_DocumentCameraViewControllerDidFailWithError  func(controller IVNDocumentCameraViewController, error_ objc.IObject /* cross-framework: Error */)
	_DocumentCameraViewControllerDidFinishWithScan func(controller IVNDocumentCameraViewController, scan IVNDocumentCameraScan)
	_DocumentCameraViewControllerDidCancel         func(controller IVNDocumentCameraViewController)
}

// SetDocumentCameraViewControllerDidFailWithError sets the handler for the DocumentCameraViewControllerDidFailWithError delegate method.
//
// Tells the delegate that document scanning failed while the camera view controller was active.
func (d *DocumentCameraViewControllerDelegate) SetDocumentCameraViewControllerDidFailWithError(f func(controller IVNDocumentCameraViewController, error_ objc.IObject /* cross-framework: Error */)) {
	d._DocumentCameraViewControllerDidFailWithError = f
}

// SetDocumentCameraViewControllerDidFinishWithScan sets the handler for the DocumentCameraViewControllerDidFinishWithScan delegate method.
//
// Tells the delegate that the user successfully saved a scanned document from the document camera.
func (d *DocumentCameraViewControllerDelegate) SetDocumentCameraViewControllerDidFinishWithScan(f func(controller IVNDocumentCameraViewController, scan IVNDocumentCameraScan)) {
	d._DocumentCameraViewControllerDidFinishWithScan = f
}

// SetDocumentCameraViewControllerDidCancel sets the handler for the DocumentCameraViewControllerDidCancel delegate method.
//
// Tells the delegate that the user canceled out of the document scanner camera.
func (d *DocumentCameraViewControllerDelegate) SetDocumentCameraViewControllerDidCancel(f func(controller IVNDocumentCameraViewController)) {
	d._DocumentCameraViewControllerDidCancel = f
}

// DocumentCameraViewControllerDidFailWithError implements the PDocumentCameraViewControllerDelegate interface.
func (d *DocumentCameraViewControllerDelegate) DocumentCameraViewControllerDidFailWithError(controller IVNDocumentCameraViewController, error_ objc.IObject /* cross-framework: Error */) {
	if d._DocumentCameraViewControllerDidFailWithError != nil {
		d._DocumentCameraViewControllerDidFailWithError(controller, error_)
	}
}

// HasDocumentCameraViewControllerDidFailWithError returns true if a handler for DocumentCameraViewControllerDidFailWithError has been set.
func (d *DocumentCameraViewControllerDelegate) HasDocumentCameraViewControllerDidFailWithError() bool {
	return d._DocumentCameraViewControllerDidFailWithError != nil
}

// DocumentCameraViewControllerDidFinishWithScan implements the PDocumentCameraViewControllerDelegate interface.
func (d *DocumentCameraViewControllerDelegate) DocumentCameraViewControllerDidFinishWithScan(controller IVNDocumentCameraViewController, scan IVNDocumentCameraScan) {
	if d._DocumentCameraViewControllerDidFinishWithScan != nil {
		d._DocumentCameraViewControllerDidFinishWithScan(controller, scan)
	}
}

// HasDocumentCameraViewControllerDidFinishWithScan returns true if a handler for DocumentCameraViewControllerDidFinishWithScan has been set.
func (d *DocumentCameraViewControllerDelegate) HasDocumentCameraViewControllerDidFinishWithScan() bool {
	return d._DocumentCameraViewControllerDidFinishWithScan != nil
}

// DocumentCameraViewControllerDidCancel implements the PDocumentCameraViewControllerDelegate interface.
func (d *DocumentCameraViewControllerDelegate) DocumentCameraViewControllerDidCancel(controller IVNDocumentCameraViewController) {
	if d._DocumentCameraViewControllerDidCancel != nil {
		d._DocumentCameraViewControllerDidCancel(controller)
	}
}

// HasDocumentCameraViewControllerDidCancel returns true if a handler for DocumentCameraViewControllerDidCancel has been set.
func (d *DocumentCameraViewControllerDelegate) HasDocumentCameraViewControllerDidCancel() bool {
	return d._DocumentCameraViewControllerDidCancel != nil
}
