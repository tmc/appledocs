// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corevideo"

	"github.com/tmc/appledocs/generated/foundation"
)

// PCaptureFileOutputRecordingDelegate is the AVCaptureFileOutputRecordingDelegate protocol interface.
//
// Methods for responding to events that occur while recording captured media to a file.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 17.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureFileOutputRecordingDelegate
type PCaptureFileOutputRecordingDelegate interface {
	// Required methods
	CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(output IAVCaptureFileOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection, error_ Error)/* debug [protocol_interface/required_method]: CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError */
	// Optional methods
	CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection)
	HasCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections() bool
	CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection)
	HasCaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections() bool
	CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection)
	HasCaptureOutputDidStartRecordingToOutputFileAtURLFromConnections() bool
	CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, startPTS objc.IObject /* cross-framework: Time */, connections []CaptureConnection)
	HasCaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections() bool
	CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection, error_ Error)
	HasCaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError() bool
}

// CaptureFileOutputRecordingDelegate is a delegate implementation builder for the PCaptureFileOutputRecordingDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureFileOutputRecordingDelegate struct {
	_CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections func(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection)
	_CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections func(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection)
	_CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections func(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection)
	_CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections func(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, startPTS objc.IObject /* cross-framework: Time */, connections []CaptureConnection)
	_CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError func(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection, error_ Error)
	_CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError func(output IAVCaptureFileOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection, error_ Error)
}

// SetCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections sets the handler for the CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections delegate method.
//
// Called whenever the output is recording to a file and successfully pauses the recording at the request of a client.
func (d *CaptureFileOutputRecordingDelegate) SetCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(f func(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection)) {
	d._CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections = f
}

// SetCaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections sets the handler for the CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections delegate method.
//
// Called whenever the output, at the request of the client, successfully resumes a file recording that was paused.
func (d *CaptureFileOutputRecordingDelegate) SetCaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(f func(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection)) {
	d._CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections = f
}

// SetCaptureOutputDidStartRecordingToOutputFileAtURLFromConnections sets the handler for the CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections delegate method.
//
// Informs the delegate when the output has started writing to a file.
func (d *CaptureFileOutputRecordingDelegate) SetCaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(f func(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection)) {
	d._CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections = f
}

// SetCaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections sets the handler for the CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections delegate method.
func (d *CaptureFileOutputRecordingDelegate) SetCaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections(f func(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, startPTS objc.IObject /* cross-framework: Time */, connections []CaptureConnection)) {
	d._CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections = f
}

// SetCaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError sets the handler for the CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError delegate method.
//
// Informs the delegate when the output will stop writing new samples to a file.
func (d *CaptureFileOutputRecordingDelegate) SetCaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(f func(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection, error_ Error)) {
	d._CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError = f
}

// SetCaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError sets the handler for the CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError delegate method.
//
// Informs the delegate when all pending data has been written to an output file.
func (d *CaptureFileOutputRecordingDelegate) SetCaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(f func(output IAVCaptureFileOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection, error_ Error)) {
	d._CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError = f
}

// CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections implements the PCaptureFileOutputRecordingDelegate interface.
func (d *CaptureFileOutputRecordingDelegate) CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection) {
	if d._CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections != nil {
		d._CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output, fileURL, connections)
	}
}

// HasCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections returns true if a handler for CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections has been set.
func (d *CaptureFileOutputRecordingDelegate) HasCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections() bool {
	return d._CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections != nil
}

// CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections implements the PCaptureFileOutputRecordingDelegate interface.
func (d *CaptureFileOutputRecordingDelegate) CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection) {
	if d._CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections != nil {
		d._CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(output, fileURL, connections)
	}
}

// HasCaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections returns true if a handler for CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections has been set.
func (d *CaptureFileOutputRecordingDelegate) HasCaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections() bool {
	return d._CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections != nil
}

// CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections implements the PCaptureFileOutputRecordingDelegate interface.
func (d *CaptureFileOutputRecordingDelegate) CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection) {
	if d._CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections != nil {
		d._CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(output, fileURL, connections)
	}
}

// HasCaptureOutputDidStartRecordingToOutputFileAtURLFromConnections returns true if a handler for CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections has been set.
func (d *CaptureFileOutputRecordingDelegate) HasCaptureOutputDidStartRecordingToOutputFileAtURLFromConnections() bool {
	return d._CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections != nil
}

// CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections implements the PCaptureFileOutputRecordingDelegate interface.
func (d *CaptureFileOutputRecordingDelegate) CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, startPTS objc.IObject /* cross-framework: Time */, connections []CaptureConnection) {
	if d._CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections != nil {
		d._CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections(output, fileURL, startPTS, connections)
	}
}

// HasCaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections returns true if a handler for CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections has been set.
func (d *CaptureFileOutputRecordingDelegate) HasCaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections() bool {
	return d._CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections != nil
}

// CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError implements the PCaptureFileOutputRecordingDelegate interface.
func (d *CaptureFileOutputRecordingDelegate) CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(output IAVCaptureFileOutput, fileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection, error_ Error) {
	if d._CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError != nil {
		d._CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(output, fileURL, connections, error_)
	}
}

// HasCaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError returns true if a handler for CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError has been set.
func (d *CaptureFileOutputRecordingDelegate) HasCaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError() bool {
	return d._CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError != nil
}

// CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError implements the PCaptureFileOutputRecordingDelegate interface.
func (d *CaptureFileOutputRecordingDelegate) CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(output IAVCaptureFileOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, connections []CaptureConnection, error_ Error) {
	if d._CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError != nil {
		d._CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(output, outputFileURL, connections, error_)
	}
}

// HasCaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError returns true if a handler for CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError has been set.
func (d *CaptureFileOutputRecordingDelegate) HasCaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError() bool {
	return d._CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError != nil
}
