// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PRecordingOutputDelegate is the SCRecordingOutputDelegate protocol interface.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - macOS 15.0+
//
// See: doc://com.apple.screencapturekit/documentation/ScreenCaptureKit/SCRecordingOutputDelegate
type PRecordingOutputDelegate interface {
	// Optional methods
	RecordingOutputDidFailWithError(recordingOutput ISCRecordingOutput, error_ objc.IObject /* cross-framework: Error */)
	HasRecordingOutputDidFailWithError() bool
	RecordingOutputDidFinishRecording(recordingOutput ISCRecordingOutput)
	HasRecordingOutputDidFinishRecording() bool
	RecordingOutputDidStartRecording(recordingOutput ISCRecordingOutput)
	HasRecordingOutputDidStartRecording() bool
}

// RecordingOutputDelegate is a delegate implementation builder for the PRecordingOutputDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type RecordingOutputDelegate struct {
	_RecordingOutputDidFailWithError func(recordingOutput ISCRecordingOutput, error_ objc.IObject /* cross-framework: Error */)
	_RecordingOutputDidFinishRecording func(recordingOutput ISCRecordingOutput)
	_RecordingOutputDidStartRecording func(recordingOutput ISCRecordingOutput)
}

// SetRecordingOutputDidFailWithError sets the handler for the RecordingOutputDidFailWithError delegate method.
func (d *RecordingOutputDelegate) SetRecordingOutputDidFailWithError(f func(recordingOutput ISCRecordingOutput, error_ objc.IObject /* cross-framework: Error */)) {
	d._RecordingOutputDidFailWithError = f
}

// SetRecordingOutputDidFinishRecording sets the handler for the RecordingOutputDidFinishRecording delegate method.
func (d *RecordingOutputDelegate) SetRecordingOutputDidFinishRecording(f func(recordingOutput ISCRecordingOutput)) {
	d._RecordingOutputDidFinishRecording = f
}

// SetRecordingOutputDidStartRecording sets the handler for the RecordingOutputDidStartRecording delegate method.
func (d *RecordingOutputDelegate) SetRecordingOutputDidStartRecording(f func(recordingOutput ISCRecordingOutput)) {
	d._RecordingOutputDidStartRecording = f
}

// RecordingOutputDidFailWithError implements the PRecordingOutputDelegate interface.
func (d *RecordingOutputDelegate) RecordingOutputDidFailWithError(recordingOutput ISCRecordingOutput, error_ objc.IObject /* cross-framework: Error */) {
	if d._RecordingOutputDidFailWithError != nil {
		d._RecordingOutputDidFailWithError(recordingOutput, error_)
	}
}

// HasRecordingOutputDidFailWithError returns true if a handler for RecordingOutputDidFailWithError has been set.
func (d *RecordingOutputDelegate) HasRecordingOutputDidFailWithError() bool {
	return d._RecordingOutputDidFailWithError != nil
}

// RecordingOutputDidFinishRecording implements the PRecordingOutputDelegate interface.
func (d *RecordingOutputDelegate) RecordingOutputDidFinishRecording(recordingOutput ISCRecordingOutput) {
	if d._RecordingOutputDidFinishRecording != nil {
		d._RecordingOutputDidFinishRecording(recordingOutput)
	}
}

// HasRecordingOutputDidFinishRecording returns true if a handler for RecordingOutputDidFinishRecording has been set.
func (d *RecordingOutputDelegate) HasRecordingOutputDidFinishRecording() bool {
	return d._RecordingOutputDidFinishRecording != nil
}

// RecordingOutputDidStartRecording implements the PRecordingOutputDelegate interface.
func (d *RecordingOutputDelegate) RecordingOutputDidStartRecording(recordingOutput ISCRecordingOutput) {
	if d._RecordingOutputDidStartRecording != nil {
		d._RecordingOutputDidStartRecording(recordingOutput)
	}
}

// HasRecordingOutputDidStartRecording returns true if a handler for RecordingOutputDidStartRecording has been set.
func (d *RecordingOutputDelegate) HasRecordingOutputDidStartRecording() bool {
	return d._RecordingOutputDidStartRecording != nil
}
