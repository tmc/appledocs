// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corevideo"
)

// PCaptureDepthDataOutputDelegate is the AVCaptureDepthDataOutputDelegate protocol interface.
//
// Methods for receiving depth data produced by a depth capture output.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - tvOS 17.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureDepthDataOutputDelegate
type PCaptureDepthDataOutputDelegate interface {
	// Optional methods
	DepthDataOutputDidDropDepthDataTimestampConnectionReason(output IAVCaptureDepthDataOutput, depthData IAVDepthData, timestamp objc.IObject /* cross-framework: Time */, connection IAVCaptureConnection, reason CaptureOutputDataDroppedReason)
	HasDepthDataOutputDidDropDepthDataTimestampConnectionReason() bool
	DepthDataOutputDidOutputDepthDataTimestampConnection(output IAVCaptureDepthDataOutput, depthData IAVDepthData, timestamp objc.IObject /* cross-framework: Time */, connection IAVCaptureConnection)
	HasDepthDataOutputDidOutputDepthDataTimestampConnection() bool
}

// CaptureDepthDataOutputDelegate is a delegate implementation builder for the PCaptureDepthDataOutputDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureDepthDataOutputDelegate struct {
	_DepthDataOutputDidDropDepthDataTimestampConnectionReason func(output IAVCaptureDepthDataOutput, depthData IAVDepthData, timestamp objc.IObject /* cross-framework: Time */, connection IAVCaptureConnection, reason CaptureOutputDataDroppedReason)
	_DepthDataOutputDidOutputDepthDataTimestampConnection func(output IAVCaptureDepthDataOutput, depthData IAVDepthData, timestamp objc.IObject /* cross-framework: Time */, connection IAVCaptureConnection)
}

// SetDepthDataOutputDidDropDepthDataTimestampConnectionReason sets the handler for the DepthDataOutputDidDropDepthDataTimestampConnectionReason delegate method.
//
// Informs the delegate that captured depth data was not processed.
func (d *CaptureDepthDataOutputDelegate) SetDepthDataOutputDidDropDepthDataTimestampConnectionReason(f func(output IAVCaptureDepthDataOutput, depthData IAVDepthData, timestamp objc.IObject /* cross-framework: Time */, connection IAVCaptureConnection, reason CaptureOutputDataDroppedReason)) {
	d._DepthDataOutputDidDropDepthDataTimestampConnectionReason = f
}

// SetDepthDataOutputDidOutputDepthDataTimestampConnection sets the handler for the DepthDataOutputDidOutputDepthDataTimestampConnection delegate method.
//
// Provides newly captured depth data to the delegate.
func (d *CaptureDepthDataOutputDelegate) SetDepthDataOutputDidOutputDepthDataTimestampConnection(f func(output IAVCaptureDepthDataOutput, depthData IAVDepthData, timestamp objc.IObject /* cross-framework: Time */, connection IAVCaptureConnection)) {
	d._DepthDataOutputDidOutputDepthDataTimestampConnection = f
}

// DepthDataOutputDidDropDepthDataTimestampConnectionReason implements the PCaptureDepthDataOutputDelegate interface.
func (d *CaptureDepthDataOutputDelegate) DepthDataOutputDidDropDepthDataTimestampConnectionReason(output IAVCaptureDepthDataOutput, depthData IAVDepthData, timestamp objc.IObject /* cross-framework: Time */, connection IAVCaptureConnection, reason CaptureOutputDataDroppedReason) {
	if d._DepthDataOutputDidDropDepthDataTimestampConnectionReason != nil {
		d._DepthDataOutputDidDropDepthDataTimestampConnectionReason(output, depthData, timestamp, connection, reason)
	}
}

// HasDepthDataOutputDidDropDepthDataTimestampConnectionReason returns true if a handler for DepthDataOutputDidDropDepthDataTimestampConnectionReason has been set.
func (d *CaptureDepthDataOutputDelegate) HasDepthDataOutputDidDropDepthDataTimestampConnectionReason() bool {
	return d._DepthDataOutputDidDropDepthDataTimestampConnectionReason != nil
}

// DepthDataOutputDidOutputDepthDataTimestampConnection implements the PCaptureDepthDataOutputDelegate interface.
func (d *CaptureDepthDataOutputDelegate) DepthDataOutputDidOutputDepthDataTimestampConnection(output IAVCaptureDepthDataOutput, depthData IAVDepthData, timestamp objc.IObject /* cross-framework: Time */, connection IAVCaptureConnection) {
	if d._DepthDataOutputDidOutputDepthDataTimestampConnection != nil {
		d._DepthDataOutputDidOutputDepthDataTimestampConnection(output, depthData, timestamp, connection)
	}
}

// HasDepthDataOutputDidOutputDepthDataTimestampConnection returns true if a handler for DepthDataOutputDidOutputDepthDataTimestampConnection has been set.
func (d *CaptureDepthDataOutputDelegate) HasDepthDataOutputDidOutputDepthDataTimestampConnection() bool {
	return d._DepthDataOutputDidOutputDepthDataTimestampConnection != nil
}
