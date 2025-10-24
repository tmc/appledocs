// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/avfoundation"
)

// PCaptureViewDelegate is the AVCaptureViewDelegate protocol interface.
//
// The protocol that defines the methods you can implement to respond to capture view events.
//
// Availability:
//   - macOS 10.9+
//
// See: doc://com.apple.avkit/documentation/AVKit/AVCaptureViewDelegate
type PCaptureViewDelegate interface {
	// Required methods
	CaptureViewStartRecordingToFileOutput(captureView IAVCaptureView, fileOutput avfoundation.CaptureFileOutput)
}

// CaptureViewDelegate is a delegate implementation builder for the PCaptureViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureViewDelegate struct {
	_CaptureViewStartRecordingToFileOutput func(captureView IAVCaptureView, fileOutput avfoundation.CaptureFileOutput)
}

// SetCaptureViewStartRecordingToFileOutput sets the handler for the CaptureViewStartRecordingToFileOutput delegate method.
//
// Tells the delegate that the user has made a request to start a new recording.
func (d *CaptureViewDelegate) SetCaptureViewStartRecordingToFileOutput(f func(captureView IAVCaptureView, fileOutput avfoundation.CaptureFileOutput)) {
	d._CaptureViewStartRecordingToFileOutput = f
}

// CaptureViewStartRecordingToFileOutput implements the PCaptureViewDelegate interface.
func (d *CaptureViewDelegate) CaptureViewStartRecordingToFileOutput(captureView IAVCaptureView, fileOutput avfoundation.CaptureFileOutput) {
	if d._CaptureViewStartRecordingToFileOutput != nil {
		d._CaptureViewStartRecordingToFileOutput(captureView, fileOutput)
	}
}

// HasCaptureViewStartRecordingToFileOutput returns true if a handler for CaptureViewStartRecordingToFileOutput has been set.
func (d *CaptureViewDelegate) HasCaptureViewStartRecordingToFileOutput() bool {
	return d._CaptureViewStartRecordingToFileOutput != nil
}
