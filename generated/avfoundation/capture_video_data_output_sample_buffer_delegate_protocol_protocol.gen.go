// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

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

// CaptureVideoDataOutputSampleBufferDelegate is a delegate implementation builder for the PCaptureVideoDataOutputSampleBufferDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureVideoDataOutputSampleBufferDelegate struct {
	_CaptureOutputDidDropSampleBufferFromConnection func(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)
	_CaptureOutputDidOutputSampleBufferFromConnection func(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)
}

// SetCaptureOutputDidDropSampleBufferFromConnection sets the handler for the CaptureOutputDidDropSampleBufferFromConnection delegate method.
//
// Notifies the delegate that a video frame was discarded.
func (d *CaptureVideoDataOutputSampleBufferDelegate) SetCaptureOutputDidDropSampleBufferFromConnection(f func(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)) {
	d._CaptureOutputDidDropSampleBufferFromConnection = f
}

// SetCaptureOutputDidOutputSampleBufferFromConnection sets the handler for the CaptureOutputDidOutputSampleBufferFromConnection delegate method.
//
// Notifies the delegate that a new video frame was written.
func (d *CaptureVideoDataOutputSampleBufferDelegate) SetCaptureOutputDidOutputSampleBufferFromConnection(f func(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)) {
	d._CaptureOutputDidOutputSampleBufferFromConnection = f
}

// CaptureOutputDidDropSampleBufferFromConnection implements the PCaptureVideoDataOutputSampleBufferDelegate interface.
func (d *CaptureVideoDataOutputSampleBufferDelegate) CaptureOutputDidDropSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection) {
	if d._CaptureOutputDidDropSampleBufferFromConnection != nil {
		d._CaptureOutputDidDropSampleBufferFromConnection(output, sampleBuffer, connection)
	}
}

// HasCaptureOutputDidDropSampleBufferFromConnection returns true if a handler for CaptureOutputDidDropSampleBufferFromConnection has been set.
func (d *CaptureVideoDataOutputSampleBufferDelegate) HasCaptureOutputDidDropSampleBufferFromConnection() bool {
	return d._CaptureOutputDidDropSampleBufferFromConnection != nil
}

// CaptureOutputDidOutputSampleBufferFromConnection implements the PCaptureVideoDataOutputSampleBufferDelegate interface.
func (d *CaptureVideoDataOutputSampleBufferDelegate) CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection) {
	if d._CaptureOutputDidOutputSampleBufferFromConnection != nil {
		d._CaptureOutputDidOutputSampleBufferFromConnection(output, sampleBuffer, connection)
	}
}

// HasCaptureOutputDidOutputSampleBufferFromConnection returns true if a handler for CaptureOutputDidOutputSampleBufferFromConnection has been set.
func (d *CaptureVideoDataOutputSampleBufferDelegate) HasCaptureOutputDidOutputSampleBufferFromConnection() bool {
	return d._CaptureOutputDidOutputSampleBufferFromConnection != nil
}
