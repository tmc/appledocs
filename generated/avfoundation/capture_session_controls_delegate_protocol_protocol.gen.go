// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PCaptureSessionControlsDelegate is the AVCaptureSessionControlsDelegate protocol interface.
//
// A protocol that defines the interface to respond to capture control activation and presentation events.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureSessionControlsDelegate
type PCaptureSessionControlsDelegate interface {
	// Required methods
	SessionControlsDidBecomeActive(session IAVCaptureSession)
	SessionControlsDidBecomeInactive(session IAVCaptureSession)
	SessionControlsWillEnterFullscreenAppearance(session IAVCaptureSession)
	SessionControlsWillExitFullscreenAppearance(session IAVCaptureSession)
}
