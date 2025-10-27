// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PCaptureSessionDeferredStartDelegate is the AVCaptureSessionDeferredStartDelegate protocol interface.
//
// A protocol that defines the interface to respond to events about a capture session’s deferred start.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureSessionDeferredStartDelegate
type PCaptureSessionDeferredStartDelegate interface {
	// Required methods
	SessionDidRunDeferredStart(session IAVCaptureSession)
	SessionWillRunDeferredStart(session IAVCaptureSession)
}
