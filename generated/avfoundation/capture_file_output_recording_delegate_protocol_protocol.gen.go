// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(output IAVCaptureFileOutput, outputFileURL foundation.foundation.INSURL, connections []CaptureConnection, error_ foundation.foundation.INSError)
	// Optional methods
	CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL foundation.foundation.INSURL, connections []CaptureConnection)
	HasCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections() bool
	CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL foundation.foundation.INSURL, connections []CaptureConnection)
	HasCaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections() bool
	CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL foundation.foundation.INSURL, connections []CaptureConnection)
	HasCaptureOutputDidStartRecordingToOutputFileAtURLFromConnections() bool
	CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections(output IAVCaptureFileOutput, fileURL foundation.foundation.INSURL, startPTS objectivec.IObject, connections []CaptureConnection)
	HasCaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections() bool
	CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(output IAVCaptureFileOutput, fileURL foundation.foundation.INSURL, connections []CaptureConnection, error_ foundation.foundation.INSError)
	HasCaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError() bool
}
