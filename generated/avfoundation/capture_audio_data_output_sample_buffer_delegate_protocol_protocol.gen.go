// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

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

// CaptureAudioDataOutputSampleBufferDelegate is a delegate implementation builder for the PCaptureAudioDataOutputSampleBufferDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureAudioDataOutputSampleBufferDelegate struct {
	_CaptureOutputDidOutputSampleBufferFromConnection func(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)
}

// SetCaptureOutputDidOutputSampleBufferFromConnection sets the handler for the CaptureOutputDidOutputSampleBufferFromConnection delegate method.
//
// Notifies the delegate that a sample buffer was written.
func (d *CaptureAudioDataOutputSampleBufferDelegate) SetCaptureOutputDidOutputSampleBufferFromConnection(f func(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)) {
	d._CaptureOutputDidOutputSampleBufferFromConnection = f
}

// CaptureOutputDidOutputSampleBufferFromConnection implements the PCaptureAudioDataOutputSampleBufferDelegate interface.
func (d *CaptureAudioDataOutputSampleBufferDelegate) CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection) {
	if d._CaptureOutputDidOutputSampleBufferFromConnection != nil {
		d._CaptureOutputDidOutputSampleBufferFromConnection(output, sampleBuffer, connection)
	}
}

// HasCaptureOutputDidOutputSampleBufferFromConnection returns true if a handler for CaptureOutputDidOutputSampleBufferFromConnection has been set.
func (d *CaptureAudioDataOutputSampleBufferDelegate) HasCaptureOutputDidOutputSampleBufferFromConnection() bool {
	return d._CaptureOutputDidOutputSampleBufferFromConnection != nil
}
