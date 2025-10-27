// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PVideoCompositing is the AVVideoCompositing protocol interface.
//
// A protocol that defines the methods custom video compositors must implement.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVVideoCompositing
type PVideoCompositing interface {
	// Required methods
	RenderContextChanged(newRenderContext IAVVideoCompositionRenderContext)
	StartVideoCompositionRequest(asyncVideoCompositionRequest IAVAsynchronousVideoCompositionRequest)
	// Optional methods
	AnticipateRenderingUsingHint(renderHint IAVVideoCompositionRenderHint)
	HasAnticipateRenderingUsingHint() bool
	CancelAllPendingVideoCompositionRequests()
	HasCancelAllPendingVideoCompositionRequests() bool
	PrerollForRenderingUsingHint(renderHint IAVVideoCompositionRenderHint)
	HasPrerollForRenderingUsingHint() bool
}
