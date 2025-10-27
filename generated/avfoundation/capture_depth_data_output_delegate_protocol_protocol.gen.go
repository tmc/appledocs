// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objectivec"
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
	DepthDataOutputDidDropDepthDataTimestampConnectionReason(output IAVCaptureDepthDataOutput, depthData IAVDepthData, timestamp objectivec.IObject, connection IAVCaptureConnection, reason CaptureOutputDataDroppedReason)
	HasDepthDataOutputDidDropDepthDataTimestampConnectionReason() bool
	DepthDataOutputDidOutputDepthDataTimestampConnection(output IAVCaptureDepthDataOutput, depthData IAVDepthData, timestamp objectivec.IObject, connection IAVCaptureConnection)
	HasDepthDataOutputDidOutputDepthDataTimestampConnection() bool
}
