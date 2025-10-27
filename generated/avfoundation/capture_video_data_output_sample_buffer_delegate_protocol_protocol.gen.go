// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PCaptureVideoDataOutputSampleBufferDelegate is the AVCaptureVideoDataOutputSampleBufferDelegate protocol interface.
//
// Methods for receiving sample buffers from, and monitoring the status of, a video data output.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 17.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureVideoDataOutputSampleBufferDelegate
type PCaptureVideoDataOutputSampleBufferDelegate interface {
	// Optional methods
	CaptureOutputDidDropSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)
	HasCaptureOutputDidDropSampleBufferFromConnection() bool
	CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)
	HasCaptureOutputDidOutputSampleBufferFromConnection() bool
}
