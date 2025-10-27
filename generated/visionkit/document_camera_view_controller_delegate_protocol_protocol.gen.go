// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

package visionkit

import (

	"github.com/tmc/appledocs/generated/foundation"
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
	DocumentCameraViewControllerDidFailWithError(controller IVNDocumentCameraViewController, error_ foundation.foundation.INSError)
	HasDocumentCameraViewControllerDidFailWithError() bool
	DocumentCameraViewControllerDidFinishWithScan(controller IVNDocumentCameraViewController, scan IVNDocumentCameraScan)
	HasDocumentCameraViewControllerDidFinishWithScan() bool
	DocumentCameraViewControllerDidCancel(controller IVNDocumentCameraViewController)
	HasDocumentCameraViewControllerDidCancel() bool
}
