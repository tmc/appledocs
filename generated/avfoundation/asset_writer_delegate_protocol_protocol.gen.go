// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

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
	AssetWriterDidOutputSegmentDataSegmentType(writer IAVAssetWriter, segmentData objc.IObject /* cross-framework: NSData */, segmentType AssetSegmentType)
	HasAssetWriterDidOutputSegmentDataSegmentType() bool
	AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport(writer IAVAssetWriter, segmentData objc.IObject /* cross-framework: NSData */, segmentType AssetSegmentType, segmentReport IAVAssetSegmentReport)
	HasAssetWriterDidOutputSegmentDataSegmentTypeSegmentReport() bool
}

// AssetWriterDelegate is a delegate implementation builder for the PAssetWriterDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AssetWriterDelegate struct {
	_AssetWriterDidOutputSegmentDataSegmentType func(writer IAVAssetWriter, segmentData objc.IObject /* cross-framework: NSData */, segmentType AssetSegmentType)
	_AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport func(writer IAVAssetWriter, segmentData objc.IObject /* cross-framework: NSData */, segmentType AssetSegmentType, segmentReport IAVAssetSegmentReport)
}

// SetAssetWriterDidOutputSegmentDataSegmentType sets the handler for the AssetWriterDidOutputSegmentDataSegmentType delegate method.
//
// Tells the delegate that the asset writer output segment data.
func (d *AssetWriterDelegate) SetAssetWriterDidOutputSegmentDataSegmentType(f func(writer IAVAssetWriter, segmentData objc.IObject /* cross-framework: NSData */, segmentType AssetSegmentType)) {
	d._AssetWriterDidOutputSegmentDataSegmentType = f
}

// SetAssetWriterDidOutputSegmentDataSegmentTypeSegmentReport sets the handler for the AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport delegate method.
//
// Tells the delegate that the asset writer output segment data and a report.
func (d *AssetWriterDelegate) SetAssetWriterDidOutputSegmentDataSegmentTypeSegmentReport(f func(writer IAVAssetWriter, segmentData objc.IObject /* cross-framework: NSData */, segmentType AssetSegmentType, segmentReport IAVAssetSegmentReport)) {
	d._AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport = f
}

// AssetWriterDidOutputSegmentDataSegmentType implements the PAssetWriterDelegate interface.
func (d *AssetWriterDelegate) AssetWriterDidOutputSegmentDataSegmentType(writer IAVAssetWriter, segmentData objc.IObject /* cross-framework: NSData */, segmentType AssetSegmentType) {
	if d._AssetWriterDidOutputSegmentDataSegmentType != nil {
		d._AssetWriterDidOutputSegmentDataSegmentType(writer, segmentData, segmentType)
	}
}

// HasAssetWriterDidOutputSegmentDataSegmentType returns true if a handler for AssetWriterDidOutputSegmentDataSegmentType has been set.
func (d *AssetWriterDelegate) HasAssetWriterDidOutputSegmentDataSegmentType() bool {
	return d._AssetWriterDidOutputSegmentDataSegmentType != nil
}

// AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport implements the PAssetWriterDelegate interface.
func (d *AssetWriterDelegate) AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport(writer IAVAssetWriter, segmentData objc.IObject /* cross-framework: NSData */, segmentType AssetSegmentType, segmentReport IAVAssetSegmentReport) {
	if d._AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport != nil {
		d._AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport(writer, segmentData, segmentType, segmentReport)
	}
}

// HasAssetWriterDidOutputSegmentDataSegmentTypeSegmentReport returns true if a handler for AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport has been set.
func (d *AssetWriterDelegate) HasAssetWriterDidOutputSegmentDataSegmentTypeSegmentReport() bool {
	return d._AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport != nil
}
