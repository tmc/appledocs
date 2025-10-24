// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

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

// CaptureFileOutputDelegate is a delegate implementation builder for the PCaptureFileOutputDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureFileOutputDelegate struct {
	_CaptureOutputDidOutputSampleBufferFromConnection func(output IAVCaptureFileOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)
	_CaptureOutputShouldProvideSampleAccurateRecordingStart func(output IAVCaptureOutput) bool
}

// SetCaptureOutputDidOutputSampleBufferFromConnection sets the handler for the CaptureOutputDidOutputSampleBufferFromConnection delegate method.
//
// Gives the delegate the opportunity to inspect samples as they are received by the output and start and stop recording at exact times.
func (d *CaptureFileOutputDelegate) SetCaptureOutputDidOutputSampleBufferFromConnection(f func(output IAVCaptureFileOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection)) {
	d._CaptureOutputDidOutputSampleBufferFromConnection = f
}

// SetCaptureOutputShouldProvideSampleAccurateRecordingStart sets the handler for the CaptureOutputShouldProvideSampleAccurateRecordingStart delegate method.
//
// Allows a client to opt in to frame accurate recording in  .
func (d *CaptureFileOutputDelegate) SetCaptureOutputShouldProvideSampleAccurateRecordingStart(f func(output IAVCaptureOutput) bool) {
	d._CaptureOutputShouldProvideSampleAccurateRecordingStart = f
}

// CaptureOutputDidOutputSampleBufferFromConnection implements the PCaptureFileOutputDelegate interface.
func (d *CaptureFileOutputDelegate) CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureFileOutput, sampleBuffer SampleBufferRef /* not a class type */, connection IAVCaptureConnection) {
	if d._CaptureOutputDidOutputSampleBufferFromConnection != nil {
		d._CaptureOutputDidOutputSampleBufferFromConnection(output, sampleBuffer, connection)
	}
}

// HasCaptureOutputDidOutputSampleBufferFromConnection returns true if a handler for CaptureOutputDidOutputSampleBufferFromConnection has been set.
func (d *CaptureFileOutputDelegate) HasCaptureOutputDidOutputSampleBufferFromConnection() bool {
	return d._CaptureOutputDidOutputSampleBufferFromConnection != nil
}

// CaptureOutputShouldProvideSampleAccurateRecordingStart implements the PCaptureFileOutputDelegate interface.
func (d *CaptureFileOutputDelegate) CaptureOutputShouldProvideSampleAccurateRecordingStart(output IAVCaptureOutput) bool {
	if d._CaptureOutputShouldProvideSampleAccurateRecordingStart != nil {
		return d._CaptureOutputShouldProvideSampleAccurateRecordingStart(output)
	}
	var zero bool
	return zero
}

// HasCaptureOutputShouldProvideSampleAccurateRecordingStart returns true if a handler for CaptureOutputShouldProvideSampleAccurateRecordingStart has been set.
func (d *CaptureFileOutputDelegate) HasCaptureOutputShouldProvideSampleAccurateRecordingStart() bool {
	return d._CaptureOutputShouldProvideSampleAccurateRecordingStart != nil
}
