// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

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
