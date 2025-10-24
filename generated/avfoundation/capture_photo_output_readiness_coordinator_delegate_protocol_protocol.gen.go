// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCapturePhotoOutputReadinessCoordinatorDelegate is the AVCapturePhotoOutputReadinessCoordinatorDelegate protocol interface.
//
// A delegate protocol to receive updates about a photo output’s capture readiness.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinatorDelegate
type PCapturePhotoOutputReadinessCoordinatorDelegate interface {
	// Optional methods
	ReadinessCoordinatorCaptureReadinessDidChange(coordinator IAVCapturePhotoOutputReadinessCoordinator, captureReadiness CapturePhotoOutputCaptureReadiness)
	HasReadinessCoordinatorCaptureReadinessDidChange() bool
}

// CapturePhotoOutputReadinessCoordinatorDelegate is a delegate implementation builder for the PCapturePhotoOutputReadinessCoordinatorDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CapturePhotoOutputReadinessCoordinatorDelegate struct {
	_ReadinessCoordinatorCaptureReadinessDidChange func(coordinator IAVCapturePhotoOutputReadinessCoordinator, captureReadiness CapturePhotoOutputCaptureReadiness)
}

// SetReadinessCoordinatorCaptureReadinessDidChange sets the handler for the ReadinessCoordinatorCaptureReadinessDidChange delegate method.
//
// Tells the delegate that the capture readiness state of a photo output changed.
func (d *CapturePhotoOutputReadinessCoordinatorDelegate) SetReadinessCoordinatorCaptureReadinessDidChange(f func(coordinator IAVCapturePhotoOutputReadinessCoordinator, captureReadiness CapturePhotoOutputCaptureReadiness)) {
	d._ReadinessCoordinatorCaptureReadinessDidChange = f
}

// ReadinessCoordinatorCaptureReadinessDidChange implements the PCapturePhotoOutputReadinessCoordinatorDelegate interface.
func (d *CapturePhotoOutputReadinessCoordinatorDelegate) ReadinessCoordinatorCaptureReadinessDidChange(coordinator IAVCapturePhotoOutputReadinessCoordinator, captureReadiness CapturePhotoOutputCaptureReadiness) {
	if d._ReadinessCoordinatorCaptureReadinessDidChange != nil {
		d._ReadinessCoordinatorCaptureReadinessDidChange(coordinator, captureReadiness)
	}
}

// HasReadinessCoordinatorCaptureReadinessDidChange returns true if a handler for ReadinessCoordinatorCaptureReadinessDidChange has been set.
func (d *CapturePhotoOutputReadinessCoordinatorDelegate) HasReadinessCoordinatorCaptureReadinessDidChange() bool {
	return d._ReadinessCoordinatorCaptureReadinessDidChange != nil
}
