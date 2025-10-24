// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

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

// CaptureMetadataOutputObjectsDelegate is a delegate implementation builder for the PCaptureMetadataOutputObjectsDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureMetadataOutputObjectsDelegate struct {
	_CaptureOutputDidOutputMetadataObjectsFromConnection func(output IAVCaptureOutput, metadataObjects []MetadataObject, connection IAVCaptureConnection)
}

// SetCaptureOutputDidOutputMetadataObjectsFromConnection sets the handler for the CaptureOutputDidOutputMetadataObjectsFromConnection delegate method.
//
// Informs the delegate that the capture output object emitted new metadata objects.
func (d *CaptureMetadataOutputObjectsDelegate) SetCaptureOutputDidOutputMetadataObjectsFromConnection(f func(output IAVCaptureOutput, metadataObjects []MetadataObject, connection IAVCaptureConnection)) {
	d._CaptureOutputDidOutputMetadataObjectsFromConnection = f
}

// CaptureOutputDidOutputMetadataObjectsFromConnection implements the PCaptureMetadataOutputObjectsDelegate interface.
func (d *CaptureMetadataOutputObjectsDelegate) CaptureOutputDidOutputMetadataObjectsFromConnection(output IAVCaptureOutput, metadataObjects []MetadataObject, connection IAVCaptureConnection) {
	if d._CaptureOutputDidOutputMetadataObjectsFromConnection != nil {
		d._CaptureOutputDidOutputMetadataObjectsFromConnection(output, metadataObjects, connection)
	}
}

// HasCaptureOutputDidOutputMetadataObjectsFromConnection returns true if a handler for CaptureOutputDidOutputMetadataObjectsFromConnection has been set.
func (d *CaptureMetadataOutputObjectsDelegate) HasCaptureOutputDidOutputMetadataObjectsFromConnection() bool {
	return d._CaptureOutputDidOutputMetadataObjectsFromConnection != nil
}
