// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PCaptureMetadataOutputObjectsDelegate is the AVCaptureMetadataOutputObjectsDelegate protocol interface.
//
// Methods for receiving metadata produced by a metadata capture output.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 13.0+
//   - tvOS 17.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureMetadataOutputObjectsDelegate
type PCaptureMetadataOutputObjectsDelegate interface {
	// Optional methods
	CaptureOutputDidOutputMetadataObjectsFromConnection(output IAVCaptureOutput, metadataObjects []MetadataObject, connection IAVCaptureConnection)
	HasCaptureOutputDidOutputMetadataObjectsFromConnection() bool
}
