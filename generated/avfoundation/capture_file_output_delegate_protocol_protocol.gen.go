// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PCaptureFileOutputDelegate is the AVCaptureFileOutputDelegate protocol interface.
//
// Methods for monitoring or controlling the output of a media file capture.
//
// Availability:
//   - macOS 10.7+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureFileOutputDelegate
type PCaptureFileOutputDelegate interface {
	// Required methods
	CaptureOutputShouldProvideSampleAccurateRecordingStart(output IAVCaptureOutput) bool
	// Optional methods
	CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureFileOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)
	HasCaptureOutputDidOutputSampleBufferFromConnection() bool
}
