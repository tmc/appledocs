// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PCaptureAudioDataOutputSampleBufferDelegate is the AVCaptureAudioDataOutputSampleBufferDelegate protocol interface.
//
// Methods for receiving audio sample data from an audio capture.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 17.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureAudioDataOutputSampleBufferDelegate
type PCaptureAudioDataOutputSampleBufferDelegate interface {
	// Optional methods
	CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)
	HasCaptureOutputDidOutputSampleBufferFromConnection() bool
}
