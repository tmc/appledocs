// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PCaptureDataOutputSynchronizerDelegate is the AVCaptureDataOutputSynchronizerDelegate protocol interface.
//
// Methods for receiving captured data from multiple capture outputs synchronized to the same timestamp.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - tvOS 17.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureDataOutputSynchronizerDelegate
type PCaptureDataOutputSynchronizerDelegate interface {
	// Required methods
	DataOutputSynchronizerDidOutputSynchronizedDataCollection(synchronizer IAVCaptureDataOutputSynchronizer, synchronizedDataCollection IAVCaptureSynchronizedDataCollection)
}
