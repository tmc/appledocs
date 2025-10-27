// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PAssetWriterDelegate is the AVAssetWriterDelegate protocol interface.
//
// A delegate protocol that defines the methods to implement to respond to asset-writing events.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVAssetWriterDelegate
type PAssetWriterDelegate interface {
	// Optional methods
	AssetWriterDidOutputSegmentDataSegmentType(writer IAVAssetWriter, segmentData foundation.foundation.INSData, segmentType AssetSegmentType)
	HasAssetWriterDidOutputSegmentDataSegmentType() bool
	AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport(writer IAVAssetWriter, segmentData foundation.foundation.INSData, segmentType AssetSegmentType, segmentReport IAVAssetSegmentReport)
	HasAssetWriterDidOutputSegmentDataSegmentTypeSegmentReport() bool
}
